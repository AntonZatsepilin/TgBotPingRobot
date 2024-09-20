# Устанавливаем базовый образ для Go
FROM golang:1.19-alpine

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файл go.mod и go.sum, чтобы установить зависимости
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем весь исходный код в рабочую директорию
COPY . .

# Собираем Go-приложение
RUN go build -o main .

# Устанавливаем переменные окружения, которые можно переопределить в Docker Compose
ENV TELEGRAM_BOT_TOKEN="your_default_bot_token"
ENV TELEGRAM_CHAT_ID="your_default_chat_id"

# Указываем команду для запуска приложения
CMD ["./main"]
