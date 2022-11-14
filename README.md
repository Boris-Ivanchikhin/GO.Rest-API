### GO.Rest-API
#### Programming in the Go language. Rest API user service

***

  - GET  /users       -> list of users: 200, 404, 500;
  - GET  /users/:id   -> user by id: 200, 404, 500;
  - POST /users/:id   -> create user: 204, 4xx, Header Location: url;
  - PUT  /users/:id   -> fully update user: 204/200, 404, 400, 500;
  - PATCH /users/:id  -> partially update user: 204/200, 404, 400, 500;
  - DELETE /users/:id -> delete user by id: 204, 404, 400.

***

По мотивам курса [Go (Golang) для веб - разработки "УЦ Специалист"](https://www.specialist.ru/track/t-go "www.specialist.ru").
Rest-API user service:
  - использован [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver#mongodb-go-driver "github.com");
  - в примере используется HTTP router [Джулиана Шмидта](https://github.com/julienschmidt/httprouter "github.com");
  - применён логгер [logrus](https://github.com/sirupsen/logrus "github.com");
  - синглтон реализован на базе модуля sync;
  - также реализовано чтение конфигурации приложения из файла config.yml (модуль \internal\config\config.go). Для чего используется пакет [cleanenv](https://github.com/ilyakaznacheev/cleanenv "github.com").

***

Также использованы материалы канала [The Art of Development](https://www.youtube.com/c/TheArtofDevelopment "youtube").