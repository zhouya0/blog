package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouya0/blog/internal/model"
	"github.com/zhouya0/blog/pkg/logger"
	"log"
	"net/http"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/zhouya0/blog/global"
	"github.com/zhouya0/blog/internal/routers"
	"github.com/zhouya0/blog/pkg/setting"
)



func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	//err = setupDBEngine()
	//if err != nil {
	//	log.Fatalf("init.setupDBEngine err: %v", err)
	//}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

}


// @title 博客系统
// @version 1.0
// @describe Yao的博客
// @termsOfService http://github.com/zhouya0/blog
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr: ":" + global.ServerSetting.HttpPort,
		Handler: router,
		ReadTimeout: global.ServerSetting.ReadTimeout,
		WriteTimeout: global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	global.Logger.Infof("%s: blog", "yao")
	s.ListenAndServe()
}


func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err !=nil {
		return err
	}
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{Filename: global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize: 600,
	MaxAge: 10,
	LocalTime: true,}, "", log.LstdFlags).WithCaller(2)
	return nil
}