# bepusdt

## CONFIG

```
# 订单有效期，单位秒，默认 600
EXPIRE_TIME=1800
# USDT 汇率，默认空，获取 Okx 交易所的汇率(每分钟同步一次)， 7.4 表示固定 7.4、～1.02 表示最新汇率上浮 2%、～0.97 表示最新汇率下浮 3%、+0.3 表示最新加 0.3、-0.2 表示最新减 0.2，以此类推；如参数错误则使用固定值 6.4
USDT_RATE=7.5
# 认证 Token，对接会用到这个参数，默认 123234
AUTH_TOKEN=
# 服务器 HTTP 监听地址，默认 :8080
LISTEN=:8080
# 是否需要网络确认，禁用可以提高回调速度，启用则可以防止交易失败，默认 0
TRADE_IS_CONFIRMED=0
# 应用访问地址，留空则系统自动获取，前端收银台会用到，默认空
APP_URI=https://xxx.xxx
# 启动时需要添加的钱包地址，多个请用半角符逗号分开，默认空
WALLET_ADDRESS=xxx
# Telegram Bot Token，默认空
TG_BOT_TOKEN=xxx
# Telegram Bot 管理员 ID，默认空
TG_BOT_ADMIN_ID=xxx
# Telegram 群组 ID，设置之后机器人会将交易消息会推送到此群，默认空
TG_BOT_GROUP_ID=
# 可选 TRON_SCAN 或 TRON_GRID，默认 TRON_SCAN
TRON_SERVER_API=TRON_GRID
# TRONSCAN API KEY
TRON_SCAN_API_KEY=
# TRONGRID API KEY
TRON_GRID_API_KEY=xxx
```

## RUN

```
go build -v -o bepusdt ./main && pm2 restart bepusdt
```

## TODO

- [ ] Support callback api
- [ ] Integrate store into telegram bot
