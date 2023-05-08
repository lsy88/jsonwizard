package main

import (
	"fmt"
	"github.com/lsy88/jsonwizard/core"
	"github.com/lsy88/jsonwizard/global"
)

func main() {
	global.JW_VP = core.Viper() //初始化viper
	core.Zap()                  //初始化zap日志库
	core.Gorm()                 //初始化数据库
	if global.JW_DB.DB != nil {
		core.RegisterTables(global.JW_DB.DB) //初始化表
		//程序结束前关闭数据库连接
		db, _ := global.JW_DB.DB.DB()
		defer db.Close()
	}
	fmt.Println("----")
	core.RunWindowsServer()
}
