CREATE EXTENSION IF NOT EXISTS hstore;

CREATE TABLE IF NOT EXISTS osm_node (
    osm_id BIGINT PRIMARY KEY,
    name TEXT,
    tags hstore,
    way GEOMETRY(Point, 4326)
);

CREATE INDEX IF NOT EXISTS idx_osm_node_way ON osm_node USING GIST (way);