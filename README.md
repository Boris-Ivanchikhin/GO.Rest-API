# GO.Rest-API
# Programming in the Go language. Rest API Examples

По мотивам курса УЦ Специалист “Программирование на языке GO. Проектирование REST API”.
Пример реализации Rest-API:
  - в примере (пока) используется HTTP router Джулиана Шмидта (https://github.com/julienschmidt/httprouter);
  - логгирование реализовано на базе logrus (модуль pkg\logging\logging.go);
  - синглтон взят из модуля sync (например, в модуле pkg\datasource\datasource.go);
  - также реализовано чтение конфигурации приложения из файла config.yml (модуль \internal\config\config.go). Для чего используется пакет cleanenv  ("github.com/ilyakaznacheev/cleanenv").

При написании примера предпочтения отдавались в пользу простым/небольшим пакетам.

P.S.
Использованы материалы youtube канала "The Art of Development" (https://www.youtube.com/c/TheArtofDevelopment).