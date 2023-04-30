# chat-bot-go

A lightweight ChatGPT telegram bot that can run fully locally with no need for any additional operations. You can run it on your own computer without configuring servers or domains


## Install
#### Download
download the latest release from [here](https://github.com/TBXark/chat-bot-go/releases)

#### Build
1. install [go](https://golang.org/)
2. `go install github.com/TBXark/chat-bot-go@latest`


## Run
1. add the config file `config.json` to the same directory as the executable file or set `--config` parameter to the config file path
2. run the executable file


## Configuration
```json

{
  "database": {
    "type": "sqlite3",
    "path": "file:db_test.sqlite?cache=shared&_fk=1"
  },
  "openai": {
    "key": "sk-",
    "model": "gpt-3.5-turbo"
  },
  "telegram": {
    "token": "123456:abc",
    "admin":[],
    "available_chat": [
      {
        "chat_id": [123, -123],
        "params":{
          "init_message": "You are a chat bot",
          "extra_params": {
            "temperature": 0.9
          }
        }
      }
    ]
  }
}
```