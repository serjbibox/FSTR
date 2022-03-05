<h1 align="center">REST API для проекта ФСТР</h1>
<h2 align="center"><a href="http://propane-facet-342315.ue.r.appspot.com/index.html#/" >Документация Swagger</a></h2>


0. [Разделительная черта](#Разделительная-черта)
1. [Заголовки](#Заголовки)

# POST 
### POST /submitData
### Request
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

#Разделительная-черта

### Response
#### Response body
```
Status 200 OK
{
  "id": "176",
  "message": "OK"
}
```
#### Response header
```
 content-length: 28 
 content-type: application/json 
 date: Sat,05 Mar 2022 10:46:32 GMT 
 server: Google Frontend 
 x-cloud-trace-context: 94ebc906a4f7734366bfbe987cb7fbe4;o=1 
```
### Bad request
```
Status:
400 нехватка полей
503 ошибка при выполнении операции
Body:
{"message" : "причина ошибки"}
```

# GET /submitData
### Request
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
  
  
  #Заголовки
  
#### Request URL
```
http://propane-facet-342315.ue.r.appspot.com/submitData
```
# Response
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



