# Software design for petconnect

### MVC approach (except for view in which will be used VueJS)

#### version 1.0.0
#### Last Updated: 29/06/2025 (DD/MM/YYYY)

---
Technologies used:

- Backend:
    - Golang (Gin framework)
    - Docker (Docker Compose)
    - Redis
    - PostgreSQL
    - ORM Golang GORM

- Frontend:
  - VueJS 3
  - Vuetify framework
  - Axios

### To run documentation in the browser for the desired function, use the command
```
godoc -http=:(port number)
for example:
godoc -http=:6060 
```
After that go in the browser and use the url
```
http://localhost:6060/pkg/slovenia_petconnect/(target folder)
```

## MVC folder structure

```text
├── cmd/
    ├──main.go
    
├── config/
    ├── config.go

├── controllers/
    ├── Register.go

├── database/
    database.go

├── docs/
    ├── v1/
        ├── HTTP_errors.md
    
├── middleware/
    ├── logger.go

├── models/
    ├── user.go

├── routes/
    ├── register.go

├── tests/
    ├── register_manual_test.go
    ├── database_test.go
├── utils/
    ├── logger.go
    ├── hashedPassword.go

├── .env
├── docker-compose.yaml
```





    