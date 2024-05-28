# Базовый образ для Go (используйте актуальную версию, например, 1.21)
FROM golang:1.21.5

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем исходный код
COPY . ./
COPY .env ./
# Сборка приложения
RUN go build -o server ./cmd/api/main.go

# Указываем порт
EXPOSE 3000

# Запуск приложения
CMD ["./server"]
