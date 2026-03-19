# Go-notifier

HTTP-приложение для отправки отложенных уведомлений через очередь RabbitMQ на go.

## Оглавление

- [Запуск](#Запуск)
- [Окружение](#Окружение)

# Окружение

В [.env](.env) можно задать TIMEZONE, а также логин и пароль для rabbitmq:managment

# Запуск

```bash
# Запуск сервисов
docker-compose up -d
```

После запуска можно мониторить сервисы:

- Rabbit: http://localhost:15672/#/queues
- Notifier: http://localhost:port/swagger/
