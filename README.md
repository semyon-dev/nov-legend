# nov-legend

## Стек бэкенда

* Go 1.16
* JWT
* MongoDB
* Yandex Translate API

## Пример .env файла

```
ACCESS_SECRET=secret
MONGO_URL="mongodb://127.0.0.1:27017/?compressors=zlib&readPreference=primary&ssl=false"
PORT=8080
FOLDER_ID=
IAM_TOKEN=
```

## Запуск проекта

`go run app/main.go`
