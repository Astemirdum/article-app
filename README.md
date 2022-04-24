# (REST API) Go app for publishing and working with articles

## The following concepts are applied in the app:
- Development of Web Applications in Go, following the REST API design.
- Using  <a href="https://github.com/gin-gonic/gin">gin-gonic/gin</a>  framework.
- Clean Architecture approach in building application structure (dependency injection)
- Using Postgres database. Generation of migration files.
- Using the database using the library <a href="https://github.com/jmoiron/sqlx">sqlx</a>.
- Application configuration using <a href="https://github.com/spf13/viper">spf13/viper</a> library. Working with .env environment variables.
- Authorization and authentication using a stateless approach (JWT).
- Graceful Shutdown.
- Run from Docker.

### build app:
```
make build
```

### run app:
```
make run
```

### run tests:
```
make test
```

