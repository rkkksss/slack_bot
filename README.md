# Slack bot 🦆


## Описание
Этот бот позволяет отправлять сообщения из json файла в каналы Slack

## Использование

В одной папке со скриптом должен лежать json файл с настройками каналов и сообщений 

```
go run main.go config.json
  ```

## Пример использования

Есть json файл с таким содержанием:
 ``` 
{
    "bot_token": "xoxb-2448831610886-2479463301456-IpDy8GjNSQhZKr3zLpCwZANf",
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

Сохранить данные вы хотите в папку res.
Тогда запуск скрипта будет иметь такой вид:
```
go run main.go config.json
```
В результате бот отправит сообщения в нужные каналы из json файла
