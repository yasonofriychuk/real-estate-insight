CREATE EXTENSION IF NOT EXISTS postgis;
CREATE EXTENSION IF NOT EXISTS hstore;

CREATE TABLE IF NOT EXISTS osm_node (
    osm_id BIGINT PRIMARY KEY,
    name TEXT,
    tags hstore,
    way GEOMETRY(Point, 4326)
);

CREATE INDEX IF NOT EXISTS idx_osm_node_way ON osm_node USING GIST (way);

CREATE TABLE IF NOT EXISTS development (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    cords GEOMETRY(Point, 4326),
    meta jsonb default '{}'::jsonb,

    created_at timestamp default current_timestamp NOT NULL,
    updated_at timestamp default current_timestamp NOT NULL,
    deleted_at timestamp
);