package web

import (
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	//"github.com/goccy/go-json"
	"github.com/akynazh/bepusdt/app/config"
	"github.com/akynazh/bepusdt/app/help"
	"github.com/akynazh/bepusdt/app/log"
)

func Start() {
	gin.SetMode(gin.ReleaseMode)

	listen := config.GetListen()
	r := gin.New()
	r.Static("/img", "./static/img")
	r.Static("/css", "./static/css")
	r.Static("/js", "./static/js")
	r.LoadHTMLGlob("./templates/*")
	r.Use(gin.LoggerWithWriter(log.GetWriter()), gin.Recovery())
	r.Use(func(ctx *gin.Context) {
		// 解析请求地址
		var _host = "http://" + ctx.Request.Host
		if ctx.Request.TLS != nil {
			_host = "https://" + ctx.Request.Host
		}
		_host = config.GetAppUri(_host)

		ctx.Set("HTTP_HOST", _host)
	})
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "皮卡数字商店",
			"url":   "https://t.me/zh_pay_bot",
		})
	})

	// ==== 支付相关=====
	payRoute := r.Group("/pay")
	{
		// 收银台
		payRoute.GET("/checkout-counter/:trade_id", CheckoutCounter)
		// 状态检测
		payRoute.GET("/check-status/:trade_id", CheckStatus)
	}

	// 创建订单
	orderRoute := r.Group("/api/v1/order")
	{
		orderRoute.Use(func(ctx *gin.Context) {
			_data, err := ctx.GetRawData()
			if err != nil {
				log.Error(err.Error())
				ctx.JSON(400, gin.H{"error": err.Error()})
				ctx.Abort()
			}

			m := make(map[string]any)
			//err = json.Unmarshal(_data, &m)
			err = sonic.Unmarshal(_data, &m)
			if err != nil {
				log.Error(err.Error())
				ctx.JSON(400, gin.H{"error": err.Error()})
				ctx.Abort()
			}

			sign, ok := m["signature"]
			if !ok {
				log.Warn("signature not found", m)
				ctx.JSON(400, gin.H{"error": "signature not found"})
				ctx.Abort()
			}

			if help.GenerateSignature(m, config.GetAuthToken()) != sign {
				log.Warn("签名错误", m)
				ctx.JSON(400, gin.H{"error": "签名错误"})
				ctx.Abort()
			}

			ctx.Set("data", m)
		})
		orderRoute.POST("/create-transaction", CreateTransaction)
	}

	log.Info("Web 启动 Listen: ", listen)
	err := r.Run(listen)
	if err != nil {

		log.Error(err.Error())
	}
}
