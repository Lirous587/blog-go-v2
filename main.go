package main

import (
	"blog/cache"
	"blog/dao/mysql"
	"blog/dao/redis"
	"blog/logger"
	"blog/pkg/snowflake"
	ticker "blog/pkg/tickerTask"
	"blog/routers"
	"blog/setting"
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
	if err := mysql.Init(); err != nil {
		fmt.Printf("init mysql failed! err:%v", err)
		return
	}

	//5.初始化redis
	if err := redis.Init(); err != nil {
		fmt.Printf("init redis failed err:%v", err)
		return
	}
	defer redis.Close()

	//6.初始计时器
	ticker.Init()

	//7初始缓存
	cache.Init()

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
