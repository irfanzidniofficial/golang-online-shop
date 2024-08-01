# Onlien Shop Project

1. Run docker for PostgreSQL

```
docker run --name postgresql -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -e POSTGRES_DB=golang_online_shop -d -p 5432:5432 postgres:16
```

2. Export enviroment variable

```
export DB_URI=postgres://user:password@localhost:5432/golang_online_shop?sslmode=disable
export ADMIN_SECRET=secret
```

3. Run program

```
go run main.go
```
