package main

import (
	"MetaCult_Insight/src/configs"
	"MetaCult_Insight/src/dao/mysql"
	"MetaCult_Insight/src/internal/router"
	"fmt"
	"log"
)

func main() {

	config := configs.InitConfig()
	if config == nil {
		panic(any("配置文件加载错误!"))
	}

	if err := mysql.InitMysql(config); err != nil {
		panic(any("MySQL init error: " + err.Error()))
	}

	r := router.InitRouter()

	go func() {

		addr := fmt.Sprintf(":%s", config.Port)
		if err := r.Run(addr); err != nil {
			log.Fatalf(" [ERROR] ServerRun:%s err:%v\n", addr, err)
		}
	}()
}
