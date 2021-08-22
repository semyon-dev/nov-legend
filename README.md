## Основной реализованный функционал:

* Получение, создание и поиск меток
* Получение, создание и фильтры маршрутов
* Регистрация, авторизация
* Расчет расстония от меток к местоположению
* Перевод контента через Yandex Translate AI 

## Особенность проекта в следующем:

* Подбор маршрута под конкретного пользователя
* Геймификация - получение опыта за посещение мест

## Стек бэкенда

* Go 1.16
* JWT
* MongoDB
* Yandex Translate API

## ЗАПУСК

развертывание сервиса производится на любой операционной системе
требуется установленный язык Golang;

`go run app/main.go`

## Пример .env файла

```
# JWT secret:
ACCESS_SECRET=secret 
MONGO_URL="mongodb://127.0.0.1:27017/?compressors=zlib&readPreference=primary&ssl=false"
PORT=8080
# для яндекс API, не обязательно:
FOLDER_ID= 
IAM_TOKEN=
```

## Установка зависимостей проекта

Установка зависимостей проекта происходит автоматически при попытке собрать бинарник

## РАЗРАБОТЧИК

Новиков Семен Сергеевич t.me/semyon_dev
