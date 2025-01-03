// main.go
// go build -v -o bepusdt ./main && pm2 restart bepusdt

package main

import (
	"fmt"
	"github.com/akynazh/bepusdt/app/config"
	"github.com/akynazh/bepusdt/app/model"
	"github.com/akynazh/bepusdt/app/monitor"
	"github.com/akynazh/bepusdt/app/web"
	"os"
	"os/signal"
	"runtime"
)

const Version = "4.0.0"

func main() {
	if err := model.Init(); err != nil {

		panic("数据库初始化失败：" + err.Error())
	}

	if config.GetTGBotToken() == "" || config.GetTGBotAdminId() == "" {

		panic("请配置参数 TG_BOT_TOKEN 和 TG_BOT_ADMIN_ID")
	}

	// 启动机器人
	go monitor.BotStart(Version)

	// 启动汇率监控
	go monitor.OkxUsdtRateStart()

	// 启动交易监控
	go monitor.TradeStart()

	// 启动回调监控
	go monitor.NotifyStart()

	// 启动web服务
	go web.Start()

	fmt.Println("Bepusdt 启动成功，当前版本：" + Version)
	// 优雅退出
	{
		var signals = make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt, os.Kill)
		<-signals
		runtime.GC()
	}
}
