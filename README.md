# Software design for petconnect

### MVC approach (except for view in which will be used VueJS)

#### version 1.0.0
#### Last Updated: 26/06/2025 (European Date)

---
Technologies used:

- Backend:
    - Golang (Gin framework)
    - Docker (Docker Compose)
    - Redis
    - PostgreSQL
    - ORM Golang GORM

## MVC folder structure

```text
├── cmd/
    ├──main.go
    
├── config/
    ├── config.go
    ├── env.go

├── controllers/
    ├── auth_controller.go

├── middlewares/
    ├── cors_middleware.go
    ├── auth_middleware.go

├── models/
    ├── user.go
    ├── post.go
├── docs/
    ├── v1/
    
├── database/
    postgres.go

├── .env
├── Dockerfile
├── docker-compose.yaml
```

implement design pattern
factory

implement algorithm in golang 




    