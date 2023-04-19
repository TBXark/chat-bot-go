# chat-bot-go

A lightweight ChatGPT telegram bot that can run fully locally with no need for any additional operations. You can run it on your own computer without configuring servers or domains

```json

{
  "database": {
    "type": "sqlite3",
    "path": "file:db_test.sqlite?cache=shared&_fk=1"
  },
  "openai": {
    "key": "sk-"
  },
  "telegram": {
    "token": "123456:abc",
    "admin":[],
    "available_chat": [
      {
        "chat_id": [123, -123]
      }
    ]
  }
}
```