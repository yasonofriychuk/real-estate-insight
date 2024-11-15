# real-estate-insight

## Команды

Импорт данных из OSM файлов. Для этого нужно положить `RU.osm.pbf` в [osmfiles](osmfiles). Скачать его можно здесь https://download.geofabrik.de/russia.html
```shell
docker exec -it app ./cron
```

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