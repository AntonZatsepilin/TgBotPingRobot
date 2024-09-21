# TgBotPingRobot

### Launch

Create an .env file like this in the root directory of your project:

``` .env
TELEGRAM_BOT_TOKEN=...
TELEGRAM_CHAT_ID=...
```

Change these parameters in the main to suit yourself:

``` go
const (
	INTERVAL        = time.Second * 10
	REQUEST_TIMEOUT = time.Second * 2
	WORKERS_COUNT   = 3
)

var urls = []string{
	"https://gb.com/AntonZatsepilin",
	"https://vk.com/antoshka_zac",
	"https://tlgg.ru/@zzwwmp",
	"https://google.com/",
	"https://golang.org/",
}
```

Enter this command in a terminal running Docker:
```
docker-compose up --build
```
