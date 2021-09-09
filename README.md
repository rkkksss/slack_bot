# Slack bot 🦆


## Описание
Этот бот позволяет отправлять сообщения из json файла в каналы Slack

## Пример использования

Пусть файл config.json лежит в одной папке с main.go и имеет такое содержание:
```JSON
{
    "bot_token": "your_bot_token",
    "channels": [
        {
            "channel": "test1",
            "text": "Hello, world!"
        },
        {
            "channel": "test2",
            "text": "Hello, world?"
        },
        {
            "channel": "test3",
            "text": "Hello, world :)"
        }
    ]
} 
```
**Не забудьте вставить необходимый токен для бота.**

Тогда запуск скрипта будет иметь вид:
```
go run main.go config.json
```
В результате бот отправит в канал test1 сообщение "Hello, world!", в канал test2 сообщение "Hello, world?" и в канал test3 сообщение "Hello, world :)"
