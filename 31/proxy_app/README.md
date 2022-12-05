# Задание 31 proxy_app

## Настройка
Файлы конфигурации находятся в файле [config_proxy.yaml](configs/config_proxy.yaml).

Настройка приложений (отсчет с нуля)
``` yaml
proxy:
  apps:
    '0': http://localhost:8080
    '1': http://localhost:8082
    '2': http://localhost:8086
```

## Запуск

Флаг port назначает порт серверу, по умолчанию 9000

    go run cmd/main.go -port=9000

## Сборка
    go build cmd/main.go
