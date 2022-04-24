# REST API Для Публикации и Работы с Articles на Go

## В работе применены следующие концепции:
- Разработка Веб-Приложений на Go, следуя дизайну REST API.
- Работа с фреймворком <a href="https://github.com/gin-gonic/gin">gin-gonic/gin</a>.
- Подход Чистой Архитектуры в построении структуры приложения. Dependency injection.
- Работа с БД Postgres. Генерация файлов миграций.
- Работа с БД, используя библиотеку <a href="https://github.com/jmoiron/sqlx">sqlx</a>.
- Конфигурация приложения с помощью библиотеки <a href="https://github.com/spf13/viper">spf13/viper</a>. Работа с переменными окружения .env.
- Авторизация и аутентификация используя stateless подход. Работа с JWT. Применение Middleware при идентификации.
- Graceful Shutdown.
- Запуск из Docker.

Файл с конфигурацией лежит в папке configs (conn db and ports)

### Для coздания приложения:

```
make build
```

### Для запуска приложения:

```
make run
```
