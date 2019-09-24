package main

import (
	"github.com/oldthreefeng/devops-api/common"
	_ "github.com/oldthreefeng/devops-api/routers"

	"github.com/astaxie/beego"
)

func main() {

	// 是否启用 定时生成验证密码 功能
	if ok, _ := beego.AppConfig.Bool("authpassword::enableCrontabAuthPassword"); ok {
		common.CronGenAuthPassword()
	}

	// 是否启用 定时清除验证密码 功能
	if ok, _ := beego.AppConfig.Bool("authpassword::enableManualGenAuthPassword"); ok {
		common.CronClearAuthPassword()
	}

	// 初始化获取命令行参数
	common.InitCli()

}

/*
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoicm9vdCIsInVwZGF0ZVRpbWUiOjE1NjkzMDkxMzh9.lZSHf7subGSwyqqF4K9wQEjSwI1qqfpwgFfl6HKRMlw

 For < louis > token only shows once, keep in mind!!!
         eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoibG91aXMiLCJ1cGRhdGVUaW1lIjoxNTY5MzA5MjA4fQ.U0EmmJ_jimuKakwPp-sA3PtxiIfBjyndEycuyrBouL8
*/