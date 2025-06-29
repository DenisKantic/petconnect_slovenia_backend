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
    
├── middleware/
    ├── logger.go

├── models/
    ├── user.go

├── routes/
    ├── register.go

├── utils/
    ├── logger.go

├── .env
├── docker-compose.yaml
```





    