# REST API "Gym Management System" на Go

Приложение реализует сервис отправки методов GET, POST, DELETE, PUT для взаимодействия с базой данных, которая описана в <a href="https://github.com/igorgofman/gms-app/blob/master/schema/000001_init.up.sql">данном файле</a>. Дополнительно реализовано авторизация и аутентификация, с помощью jwt.

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

Если как и я, используете Git Bash в Windows и у вас нет команды make в терминале, советую этот <a href="https://gist.github.com/evanwill/0207876c3243bbb6863e65ec5dc3f058#make">гайд на установку</a>.



### P.s.
Inspired by <a href="https://github.com/zhashkevych/todo-app">Maksim Zhashkevych</a>.