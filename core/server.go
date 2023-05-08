//go:build windows

package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lsy88/jsonwizard/global"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.JW_CONFIG.System.UseMultipoint || global.JW_CONFIG.System.UseRedis {
		// 初始化redis服务
		Redis()
	}
	
	// 从db加载jwt数据
	//if global.GVA_DB != nil {
	//	system.LoadAll()
	//}
	
	Router := Routers()
	//Router.Static("/form-generator", "./resource/page")
	
	address := fmt.Sprintf(":%d", global.JW_CONFIG.System.Port)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.JW_LOG.Info("server run success on ", zap.String("address", address))
	
	fmt.Printf(`
	欢迎使用 jsonwizard
	当前版本:v1.0

     ██╗███████╗ ██████╗ ███╗   ██╗██╗    ██╗██╗███████╗ █████╗ ██████╗ ██████╗
     ██║██╔════╝██╔═══██╗████╗  ██║██║    ██║██║╚══███╔╝██╔══██╗██╔══██╗██╔══██╗
     ██║███████╗██║   ██║██╔██╗ ██║██║ █╗ ██║██║  ███╔╝ ███████║██████╔╝██║  ██║
██   ██║╚════██║██║   ██║██║╚██╗██║██║███╗██║██║ ███╔╝  ██╔══██║██╔══██╗██║  ██║
╚█████╔╝███████║╚██████╔╝██║ ╚████║╚███╔███╔╝██║███████╗██║  ██║██║  ██║██████╔╝
 ╚════╝ ╚══════╝ ╚═════╝ ╚═╝  ╚═══╝ ╚══╝╚══╝ ╚═╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═════╝

`)
	global.JW_LOG.Error(s.ListenAndServe().Error())
}

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
