# Задание 31 server_app

## Настройки 
Файлы конфигурации находятся в файле [config.yaml](configs/config.yaml).

Сервер подключается к базе данных redis версии 6
``` yaml
redis:
  addr: "localhost"
  port: "6379"
  pass: ""
```

## Запуск

Флаг port назначает порт серверу, по умолчанию 8080

    go run cmd/main.go -port=8080

## Сборка
    go build cmd/main.go