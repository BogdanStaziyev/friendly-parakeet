# Test Routes

## Ping

`http://localhost:8080/api/ping`

## Endpoints

### Coordinates

`Find all: Get http://localhost:8080/api/v1/coordinates`

`Find one: Get http://localhost:8080/api/v1/coordinates/{id}`
			
`Create: Post http://localhost:8080/api/v1/coordinates/add`

Request structure
```json
{
    "mt": 116,
    "axis": "178 оси",
    "horizon": "1350 метр",
    "x": 25960.772,
    "y": 35375.685
}
```
      
`Update: Put http://localhost:8080/api/v1/coordinates/update`
```json
{
    "id": 11,
    "mt": 116,
    "axis": "178 оси",
    "horizon": "1350 метр",
    "x": 25960.772,
    "y": 35375.685
}
```

`Delete: Delete http://localhost:8080/api/v1/coordinates/{id}`
      
`Invert: Get http://localhost:8080/api/v1/coordinates/{firstId}/{secondId}`

## Users

`Create user: POST http://localhost:8080/api/v1/user`

Request structure:
```json
{
    "name":"name",
    "email":"Email@email.com",
    "password":"password",
    "role_id": 1
}
```

Return structure:
```json
{
  "id": 1,
  "name": "name",
  "email": "Email111@email.com",
  "role_id": 1,
  "created_date": "2022-08-16T11:07:53.080809Z",
  "updated_date": "0001-01-01T00:00:00Z",
  "deleted_date": "0001-01-01T00:00:00Z"
}
```

`Update user: PUT http://localhost:8080/api/v1/user/{id}`

Request structure:
```json
{
    "name":"name",
    "email":"Email@email.com",
    "role_id": 1
}
```

Return structure:
```json
{
    "id": 11,
    "name": "name",
    "email": "Email@email.com",
    "role_id": 1,
    "created_date": "2022-08-16T11:10:58.346794Z",
    "updated_date": "2022-08-16T11:12:19.458738Z",
    "deleted_date": "0001-01-01T00:00:00Z"
}
```

`Pageinate users: GET http://localhost:8080/api/v1/user/?page=1&pageSize=20&showDeleted=0`

`Get one user: GET http://localhost:8080/api/v1/user/{id}?showDeleted=0`

`Delete user: DELETE http://localhost:8080/api/v1/user/{id}`

`LogIn user: POST http://localhost:8080/api/v1/user/login`

Request structure:
```json
{
    "email":"Email@email.com",
    "password":"pwd"
}
```

Return structure:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX3JvbGUiOjEsInRva2VuX2lkIjoyLCJleHAiOjE2NjA2NTE2MTB9.Qzwx6nZRfmiQSDUL9_MvGprCTFPYw9VKZbT17gAgXLg"
}
```

`LogOut user: POST http://localhost:8080/api/v1/user/logout`

`Check user authorization: GET http://localhost:8080/api/v1/user/checkauth`





﻿# startUp
# Створити додаток для обчислення зворотньої геодезичної задачі ✅	
# _ обчислення примикання методом зєднувального трикутника 
# _ обчислення площ складу діленням на прості фігури
# _ обчислення нівелірних та теодолітних ходів з записом точок координат горизонтів і осей в БД
