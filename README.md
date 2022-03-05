<h1 align="center">REST API для проекта ФСТР</h1>
<h2 align="center"><a href="http://propane-facet-342315.ue.r.appspot.com/index.html#/" >Документация Swagger</a></h2>

- [Создание записи](#Создание-записи)
- [Запрос записей](#Запрос-записей)
- [Запрос записи по ID](#Запрос-записи-по-ID)
- [Запрос статуса записи по ID](#Запрос-статуса-записи-по-ID)
- [Редактирование записи](#Редактирование-записи)
- [Bad request](#Bad-request)
- [Инструкция по деплою](#Инструкция-по-деплою)

# Создание записи
## POST /submitData
Добавляет новую карточку объекта в БД. 
Возвращает ID добавленной записи.
#### Request
```
curl -X 'POST' \
  'http://propane-facet-342315.ue.r.appspot.com/submitData' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{request body}'
  ```
#### Request URL
```

http://propane-facet-342315.ue.r.appspot.com/submitData
```
 #### Request body
 ```
 {
  "add_time": "2021-09-22 13:18:13",
  "beautyTitle": "пер. ",
  "connect": " ",
  "coords": {
    "height": "1200",
    "latitude": "45.3842",
    "longitude": "7.1525"
  },
  "images": [
    {
      "title": "Спуск. Фото №99",
      "url": "https://img.icons8.com/ios/2x/roller-skating.png"
    }
  ],
  "level": {
    "autumn": "1A",
    "spring": " ",
    "summer": "1A",
    "winter": " "
  },
  "other_titles": "1",
  "pereval_id": "125",
  "title": "Туя-Ашуу",
  "type": "pass",
  "user": {
    "email": "dd@aa.ru",
    "fam": "Скворцов",
    "id": "11234",
    "name": "Иван",
    "otc": "Кожедубович",
    "phone": "+744434555"
  }
}
```
#### Response
#### Response body
```
Status 200 OK
{
  "id": "176",
  "message": "OK"
}
```
#### Response headers
```
 content-length: 28 
 content-type: application/json 
 date: Sat,05 Mar 2022 10:46:32 GMT 
 server: Google Frontend 
 x-cloud-trace-context: 94ebc906a4f7734366bfbe987cb7fbe4;o=1 
```

# Запрос записей
## GET /submitData
Возвращает все записи пользователя
Данные для фильтрации в теле запроса. Фильтрация по ФИО, почта, телефон.
#### Request
```

curl -X 'GET' \
  'http://propane-facet-342315.ue.r.appspot.com/submitData/' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "email": "sample@sample.com",  
  "fam": "Иванов",
  "name": "Иван",
  "otc": "Иванович",
  "phone": "+71234567890"
}'
//Requared fields: 
"email"
OR
"phone"
OR
"fam" AND "name" AND "otc" ()

  ```
#### Request URL
```
http://propane-facet-342315.ue.r.appspot.com/submitData
```
#### Response
#### Response body
```
[{
        "pereval_id": "125",
        "beautyTitle": "пер. ",
        "title": "2",
        "other_titles": "1",
        "connect": " ",
        "add_time": "2021-09-22 13:18:13",
        "coords": {
            "latitude": "45.3842",
            "longitude": "7.1525",
            "height": "1200"
        },
        "type": "pass",
        "level": {
            "winter": " ",
            "summer": "1A",
            "autumn": "1A",
            "spring": " "
        },
        "user": {
            "id": "11234",
            "email": "sample@sample.com",
            "phone": "+71234567890",
            "fam": "Иванов",
            "name": "Иван",
            "otc": "Иванович"
        },
        "images": [
            {
                "title": "Спуск. Фото №106",
                "blob": "data:image/png;base64,imgdata"
            }
        ]
    }]

```

# Запрос записи по ID
## GET /submitData/{id}
Возвращает запись по ID БД.
#### Request
```
curl -X 'GET' \
  'http://propane-facet-342315.ue.r.appspot.com/submitData/174' \
  -H 'accept: application/json'
 ```
#### Request URL
```
http://propane-facet-342315.ue.r.appspot.com/submitData/174
```
#### Response
#### Response body
```
{
  "pereval_id": "125",
  "beautyTitle": "пер. ",
  "title": "Туя-Ашуу",
  "other_titles": "1",
  "connect": " ",
  "add_time": "2021-09-22 13:18:13",
  "coords": {
    "latitude": "45.3842",
    "longitude": "7.1525",
    "height": "1200"
  },
  "type": "pass",
  "level": {
    "winter": " ",
    "summer": "1A",
    "autumn": "1A",
    "spring": " "
  },
  "user": {
    "id": "11234",
    "email": "dd@aa.ru",
    "phone": "+744434555",
    "fam": "Скворцов",
    "name": "Иван",
    "otc": "Кожедубович"
  },
  "images": [
    {
      "title": "Спуск. Фото №99",
      "blob": "data:image/png;base64,imageblob"
    }
  ]
}

```
#### Response headers
```
 content-length: 3263 
 content-type: application/json 
 date: Sat,05 Mar 2022 11:34:55 GMT 
 server: Google Frontend 
 x-cloud-trace-context: a736ad2d9837ee3001a1517cb36ea32a
```

# Запрос статуса записи по ID
## GET /submitData/{id}/status
Возвращает статус модерирования записи.
#### Request
```
curl -X 'GET' \
  'http://propane-facet-342315.ue.r.appspot.com/submitData/174/status' \
  -H 'accept: application/json'
 ```
#### Request URL
```
http://propane-facet-342315.ue.r.appspot.com/submitData/174/status
```
#### Response
#### Response body
```
{
  "id": "174",
  "status": "new"
}
```
#### Response headers
```
 content-length: 28 
 content-type: application/json 
 date: Sat,05 Mar 2022 11:37:51 GMT 
 server: Google Frontend 
 x-cloud-trace-context: fd03f60da5e25048558afcedf6ee7197;o=1 
```

# Редактирование записи
## PUT /submitData/{id}
Обновляет запись по ID. Для обновления недоступны поля: "name", "fam", "otc", "email", "phone".
#### Request
```
curl -X 'PUT' \
  'http://propane-facet-342315.ue.r.appspot.com/submitData/174' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{request body}'
 ```
#### Request URL
```
http://propane-facet-342315.ue.r.appspot.com/submitData/174
```
 #### Request body
 ```
 {
  "add_time": "2021-09-22 13:18:13",
  "beautyTitle": "пер. ",
  "connect": " ",
  "coords": {
    "height": "1200",
    "latitude": "45.3842",
    "longitude": "7.1525"
  },
  "images": [
    {
      "title": "Спуск. Фото №99",
      "url": "https://img.icons8.com/ios/2x/roller-skating.png"
    }
  ],
  "level": {
    "autumn": "1A",
    "spring": " ",
    "summer": "1A",
    "winter": " "
  },
  "other_titles": "1",
  "pereval_id": "125",
  "title": "Туя-Ашуу",
  "type": "pass",
  "user": {
    "email": "dd@aa.ru",
    "fam": "Скворцов",
    "id": "11234",
    "name": "Иван",
    "otc": "Кожедубович",
    "phone": "+744434555"
  }
}
 ```

#### Response
#### Response body
```
{
  "id": "174",
  "message": "OK"
}

```
#### Response headers
```
 content-length: 28 
 content-type: application/json 
 date: Sat,05 Mar 2022 11:38:49 GMT 
 server: Google Frontend 
 x-cloud-trace-context: c2fbdc552ac9809af8677af9cc075dac;o=1 
```
# Bad request
## Bad request
```
Status:
400 нехватка полей
503 ошибка при выполнении операции
Body:
{"message" : "причина ошибки"}
```

# Инструкция по деплою
### Деплой на Google Cloud
- зарегистрировать аккаунт Google Cloud
- создать инстанс
- Подключить Cloud Shell для Golang
- Выбрать проект: $ gcloud config set project \ project-name
- Клонировать репозиторий: $ git clone https://github.com/serjbibox/FSTR
- Выбрать workspace $ cd \ cloud-deploy-tutorials/tutorials/base \ && cloudshell workspace
- Настроить GKE кластер и инфраструктуру $ ./setup.sh
- Проверить настройки $ gcloud container clusters list
- Настроить файл app.Yaml, подключить переменные среды, для доступа к БД
- Протестировать приложение в Cloud Shell
- Можно деплоить: $ gcloud app deploy
