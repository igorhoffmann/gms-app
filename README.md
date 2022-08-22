# REST API "Gym Management System" на Go

Приложение реализует сервис отправки методов GET, POST, DELETE, PUT для взаимодействия с базой данных, которая описана в <a href="https://github.com/igorgofman/gms-app/blob/master/schema/000001_init.up.sql">данном файле</a>. Дополнительно реализована авторизация и аутентификация, с помощью jwt.

## API:

### POST /auth/sign-up

Creates new user 

##### Example Input: 
```
{
	"username": "UncleBob",
	"password": "cleanArch"
} 
```


### POST /auth/sign-in

Request to get JWT Token based on user credentials

##### Example Input: 
```
{
	"username": "UncleBob",
	"password": "cleanArch"
} 
```

##### Example Response: 
```
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzEwMzgyMjQuNzQ0MzI0MiwidXNlciI6eyJJRCI6IjAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMCIsIlVzZXJuYW1lIjoiemhhc2hrZXZ5Y2giLCJQYXNzd29yZCI6IjQyODYwMTc5ZmFiMTQ2YzZiZDAyNjlkMDViZTM0ZWNmYmY5Zjk3YjUifX0.3dsyKJQ-HZJxdvBMui0Mzgw6yb6If9aB8imGhxMOjsk"
} 
```

### POST /api/info

Creates new info

##### Example Input: 

###### Case instructor:
```
{
      "first_name": "Marlon",
      "last_name": "Brando",
      "relationship": "istructor",
      "phone": "+380000000000",
      "date_of_birth": "03.04.1998",
      "salary": 3000
} 
```
###### Case member:
```
{
      "first_name": "Steven",
      "last_name": "Jobs",
      "relationship": "member",
      "phone": "+380000000001",
      "date_of_birth": "07.08.2000",
      "membership_id": 1,
      "expires_at": "16.11.2022"
} 
```

### GET /api/info

Returns all user info

##### Example Response: 
```
{
	"info": [
            {
                "id": 1
                "first_name": "Marlon",
                "last_name": "Brando",
                "middle_name": "",
                "relationship": "istructor",
                "phone": "+380000000000",
                "date_of_birth": "03.04.1998",
                "date_of_registry": "16.08.2022",
                "hire_date": "16.08.2022",
                "salary": 3000
            }
    ]
} 
```

### DELETE /api/info

Deletes info by ID:

##### Example Input: 
```
{
	"id": "1"
} 
```


## Requirements
- go 1.18
- docker & docker-compose
- postman (для тестирования)


## Для запуска приложения:

```
make build && make run
```

Если приложение запускается впервые, необходимо применить миграции к базе данных:

```
make migrate
```

## <a href="https://github.com/igorgofman/DB-CNTU/blob/main/info-backup.sql">Тестовые данные</a>

Можно просто скопировать и вставить содержимое файла в терминале, открыв нужную базу данных.

Если используете Git Bash в Windows и у вас нет команды make в терминале, советую этот <a href="https://gist.github.com/evanwill/0207876c3243bbb6863e65ec5dc3f058#make">гайд</a>.



### P.s.
Inspired by <a href="https://github.com/zhashkevych/todo-app">Maksim Zhashkevych</a>.
