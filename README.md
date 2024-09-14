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

Импорт osm в базу данных
```shell
docker-compose exec postgis bash -c "/usr/local/bin/import_osm.sh"
```
