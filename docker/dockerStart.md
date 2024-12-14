# Instractions

- Выполнить сборку образов и запустить контейнеры.
````
docker compose up --build -d

````
- Сборка и запуск
````
go build

go run ./main.go

````
# Миграции Goolang-migrate CLI
```
https://github.com/golang-migrate/migrate
```

- Создание миграций
````
migrate create -ext=sql -dir=app/db/migrations -seq createTaskTasble
````
- Запуск миграций
````
migrate -path=app/db/migrations -database "postgresql://admin:admin@localhost:5432/reports?sslmode=disable" -verbose up
````
- Откат миграций
````
migrate -path=app/db/migrations -database "postgresql://admin:admin@localhost:5432/reports?sslmode=disable" -verbose down
````