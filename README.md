# real-estate-insight

## Запуск проекта

1. Заполнить .env файл на основе .env.example
2. Разверуть контейнеры `docker-compose up --build`
3. Если сервис запускается впервые, то импортировать geo-данные. 
Для этого нужно положить [RU.osm.pbf](https://download.geofabrik.de/russia.html) в [osmfiles](osmfiles). 
После чего запустить импорт `docker exec -it app ./cron`. Количество импортируемых нод будет выводиться в консоль, 
на момент написания инструкции их количество составляет ~1_400_000
4. Для проверки работы сервиса можно использовать запросы в [api.http](http%2Fapi.http) или посмотреть PWA по адресу http://127.0.0.1:8080/


## Helpers

Запуск контейнеров
```shell
docker-compose up --build
```

Остановка контейнеров с очисткой volume
```shell
docker-compose down -v
```

Форматирование/линтеры
```shell
golangci-lint run --fix
```

Импорт данных
```shell
docker exec -it app ./cron
```

- Добавить отображение радиуса
- Считать фактическое расстояние до объекта?
