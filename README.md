# real-estate-insight

## Команды

Запуск контейнеров
```shell
docker-compose up --build -d
```

Остановка контейнеров с очисткой volume
```shell
docker-compose down -v
```

Форматирование/линтеры
```shell
golangci-lint run --fix
```