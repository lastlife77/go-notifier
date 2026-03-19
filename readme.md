# Go-notifier

HTTP-приложение для отправки отложенных уведомлений через очередь RabbitMQ на go.

## Оглавление

- [Запуск](#Запуск)

# Запуск

```bash
# Запуск сервисов
docker-compose up -d
```

После запуска можно мониторить сервисы:

- Rabbit: http://localhost:15672/#/queues (Логин и пароль задаются в [.env](.env))
- Notifier: http://localhost:port/swagger/
