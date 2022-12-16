# Xtrame with Go (Golang) REST API
make a point api for booking cinema ticket application

Used libraries:
- [gin](https://github.com/gin-gonic)
- [gorm](https://gorm.io/docs/)
- [godotenv](https://pkg.go.dev/github.com/joho/godotenv?tab=doc)
- [validator](github.com/go-playground/validator/v10)

## Run Locally
Create `.env` at root, i.e.
```
DATABASE_URL=postgresql://${{ PGUSER }}:${{ PGPASSWORD }}@${{ PGHOST }}:${{ PGPORT }}/${{ PGDATABASE }}
PGHOST=your_local
PGPORT=5432
PGUSER=postgres
PGPASSWORD=your_password
PGDATABASE=your_db_name
PORT=8080
```

Setup Db after create database in your postgres
```
CREATE TYPE role AS ENUM ('admin', 'user')
```

Run 
```
go run cmd/main.go
```