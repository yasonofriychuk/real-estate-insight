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
- Ссылки на внешние сервисы (Карты логотипы?)
- Считать фактическое расстояние до объекта?
- Выдать доступ всем?
- Математическое обеспечение ()
- Тепловая карта

```sql
WITH bbox AS (
    SELECT ST_Transform(ST_MakeEnvelope(:minLon, :minLat, :maxLon, :maxLat, 4326), 3857) AS geom
),
     hexgrid AS (
         SELECT (ST_HexagonGrid(:cellSizeMeters, bbox.geom)).geom AS geom FROM bbox
     ),
     filtered_nodes AS (
         SELECT way,
                CASE
                    WHEN tags -> 'hospital' = 'yes' THEN :w_hospital
                    WHEN tags -> 'sport' = 'yes' THEN :w_sport
                    WHEN tags -> 'shop' = 'yes' THEN :w:shop
                    WHEN tags -> 'kindergarten' = 'yes' THEN :w_kindergarten
                    WHEN tags -> 'bus_stop' = 'yes' THEN :w_bus_stop
                    WHEN tags -> 'school' = 'yes' THEN :w_school
                    ELSE 0
                    END AS weight
         FROM osm_node
         WHERE ST_Contains(ST_MakeEnvelope(:minLon, :minLat, :maxLon, :maxLat, 4326), way)
     ),
     node_geoms AS (
         SELECT ST_Transform(way, 3857) AS geom, weight FROM filtered_nodes WHERE weight > 0
     ),
     aggregated AS (
         SELECT h.geom, SUM(n.weight) AS total_weight FROM hexgrid h
                                                               LEFT JOIN node_geoms n  ON ST_Intersects(h.geom, n.geom)
         GROUP BY h.geom
     )

SELECT
    ST_AsGeoJSON(ST_Transform(geom, 4326))::json AS geometry,
    total_weight
FROM aggregated;

```