package main

import (
	"blog/logger"
	"blog/pkg/snowflake"
	"blog/routers"
	"blog/setting"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	// 加载配置文件
	if err := setting.Init(); err != nil {
		fmt.Println("init setting failed!")
	}

	// 初始化日志
	if err := logger.Init(viper.GetString("app.mode")); err != nil {
		return
	}

	defer func() {
		//退出前写入全部日志
		if err := zap.L().Sync(); err != nil {
			fmt.Printf(" zap.L().Sync() failed,err:%v", err)
		}
	}()

	// 初始化雪花算法
	if err := snowflake.Init(viper.GetString("app.start_time"), viper.GetInt64("app.machine_id")); err != nil {
		fmt.Printf("snowflake init failed,err:%v", err)
		return
	}
	// 注册路由
	r := routers.SetupRouter(viper.GetString("app.mode"))

	port := fmt.Sprintf(":%d", viper.GetInt("app.port"))
	//err := r.RunTLS(port, "ssl/server.crt", "ssl/server.key")
	err := r.Run(port)

	if err != nil {
		fmt.Printf("run server failed,err:%v", err)
		return
	}
}
