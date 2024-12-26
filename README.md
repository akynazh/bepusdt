# bepusdt

## CONFIG

```
# .env

# 订单有效期，单位秒
EXPIRE_TIME=600
# USDT 汇率，获取 OKX 交易所的汇率(每分钟同步一次)， 7.4 表示固定 7.4、～1.02 表示最新汇率上浮 2%、～0.97 表示最新汇率下浮 3%、+0.3 表示最新加 0.3、-0.2 表示最新减 0.2，以此类推；如参数错误则使用固定值 6.4
USDT_RATE=7.4
# 交易认证 Token
AUTH_TOKEN=xxxxxxxxxxxxxx
# 服务器 HTTP 监听地址
LISTEN=:8080
# 是否需要网络确认，禁用(0)可以提高回调速度，启用(1)则可以防止交易失败
TRADE_IS_CONFIRMED=0
# 应用访问地址，留空则系统自动获取，前端收银台会用到
APP_URI=https://xxx.xxx
# 启动时需要添加的钱包地址，多个请用半角符逗号分开
WALLET_ADDRESS=xxxxxxxxxxxxxx
# 钱包地址对应的二维码图片
WALLET_PHOTO=https://xxx.xxx/xxx/xxx.png
# Telegram Bot Token
TG_BOT_TOKEN=xxxxxxxxxxxxxx
# Telegram Bot 管理员 ID
TG_BOT_ADMIN_ID=xxxxxxxxxxxxxx
# TRONGRID API KEY
TRON_GRID_API_KEY=xxxxxxxxxxxxxx
# 订单完成后的回调接口
NOTIFY_URL=https://xxx.xxx/notify
# 订单完成后的重定向接口
REDIRECT_URL=https://xxx.xxx/redirect
# Redis 缓存配置
REDIS_HOST=127.0.0.1
REDIS_PORT=6379
REDIS_PASSWORD=xxxxxxxxxxxxxx
```

## RUN

```
go build -v -o bepusdt ./main && pm2 restart bepusdt
```

## TODO

- [ ] Support callback api
- [ ] Integrate store into telegram bot
- [ ] Use TRONGRID totally
- [ ] Remove telegram group push function
