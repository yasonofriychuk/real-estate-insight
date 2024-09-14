#!/bin/bash

# Установка переменных окружения для подключения к PostgreSQL
export PGUSER=osmuser
export PGPASSWORD=osmpassword
export PGHOST=postgis
export PGPORT=5432
export PGDATABASE=osm

# Директории с OSM PBF файлами и плоским файлом узлов
OSM_DIR="/osmfiles/append"
COMPLETED_DIR="/osmfiles/completed"
FLAT_NODES="/osmfiles/flat_nodes.bin"
INIT_FILE="/osmfiles/init/empty.osm"  # Файл для инициализации
STYLE_FILE="/osmfiles/custom.style"  # Кастомный файл стиля

# Убедитесь, что hstore расширение активировано
psql -U osmuser -d osm -h postgis -c "CREATE EXTENSION IF NOT EXISTS hstore;"

# Инициализация базы данных - создание таблиц с пустым файлом
if [ -f "$INIT_FILE" ]; then
  echo "Initializing database with empty OSM file..."
  osm2pgsql --slim --hstore --create --flat-nodes "$FLAT_NODES" \
            --style "$STYLE_FILE" "$INIT_FILE"
else
  echo "Initialization file $INIT_FILE not found!"
  exit 1
fi

# Проверяем, есть ли файлы в директории для импорта данных
if [ -d "$OSM_DIR" ]; then
  # Убедитесь, что папка для завершенных файлов существует
  if [ ! -d "$COMPLETED_DIR" ]; then
    mkdir -p "$COMPLETED_DIR"
  fi

  for file in $OSM_DIR/*.osm.pbf; do
    if [ -f "$file" ]; then
      echo "Importing $file ..."
      osm2pgsql --slim --hstore --append --flat-nodes "$FLAT_NODES" \
                --style "$STYLE_FILE" --cache 2000 --number-processes 4 "$file"

      # Перемещение обработанного файла в папку completed
      echo "Moving $file to completed directory..."
      mv "$file" "$COMPLETED_DIR/"
    else
      echo "No .osm.pbf files found in $OSM_DIR."
    fi
  done
else
  echo "Directory $OSM_DIR does not exist."
  exit 1
fi

# Удаление остаточного файла flat_nodes.bin
if [ -f "$FLAT_NODES" ]; then
  echo "Removing flat nodes file..."
  rm "$FLAT_NODES"
fi

echo "OSM import completed."
