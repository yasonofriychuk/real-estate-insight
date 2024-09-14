# Используем минимальный образ PostgreSQL с PostGIS
FROM postgis/postgis:15-3.3

# Устанавливаем необходимые пакеты для osm2pgsql и netcat (для ожидания)
RUN apt-get update && apt-get install -y \
    osm2pgsql \
    netcat \
    && rm -rf /var/lib/apt/lists/*

# Копируем скрипты
COPY scripts/import_osm.sh /usr/local/bin/import_osm.sh

# Делаем скрипты исполняемыми
RUN chmod +x /usr/local/bin/import_osm.sh