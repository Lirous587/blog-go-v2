package main

import (
	"blog/cache"
	"blog/logger"
	"blog/pkg/snowflake"
	"blog/repository"
	"blog/routers"
	"blog/setting"
	"blog/ticker"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	//1.加载配置文件
	if err := setting.Init(); err != nil {
		fmt.Println("init setting failed!")
	}

	//2.初始化日志
	if err := logger.Init(viper.GetString("app.mode")); err != nil {
		return
	}

	defer func() {
		//退出前写入全部日志
		if err := zap.L().Sync(); err != nil {
			fmt.Printf(" zap.L().Sync() failed,err:%v", err)
		}
	}()

	//3.初始化雪花算法
	if err := snowflake.Init(viper.GetString("app.start_time"), viper.GetInt64("app.machine_id")); err != nil {
		fmt.Printf("snowflake init failed,err:%v", err)
		return
	}

	//4.初始化mysql
	if err := repository.Init(); err != nil {
		fmt.Printf("init mysql failed! err:%v", err)
		return
	}

	//5.初始化redis
	if err := cache.Init(); err != nil {
		fmt.Printf("init redis failed err:%v", err)
		return
	}
	defer cache.Close()

	//6.初始计时器
	ticker.Init()

	//8.注册路由
	r := routers.SetupRouter(viper.GetString("app.mode"))

	port := fmt.Sprintf(":%d", viper.GetInt("app.port"))
	//err := r.RunTLS(port, "ssl/server.crt", "ssl/server.key")
	err := r.Run(port)

	if err != nil {
		fmt.Printf("run server failed,err:%v", err)
		return
	}

}
