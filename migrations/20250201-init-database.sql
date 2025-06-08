CREATE EXTENSION IF NOT EXISTS postgis;
CREATE EXTENSION IF NOT EXISTS hstore;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

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
    location_id BIGINT references location not null,
    meta jsonb default '{}'::jsonb,

    created_at timestamp default current_timestamp NOT NULL,
    updated_at timestamp default current_timestamp NOT NULL,
    deleted_at timestamp
);

CREATE TYPE location_type AS ENUM('country', 'city', 'region');

CREATE TABLE IF NOT EXISTS location (
    id BIGINT PRIMARY KEY,
    region_id BIGINT references location(id),
    country_id BIGINT references location(id),
    name TEXT not null,
    loc_type location_type not null
);

CREATE TABLE IF NOT EXISTS profile (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT unique not null,
    password_hash TEXT not null,

    created_at timestamp default current_timestamp NOT NULL
);

CREATE TABLE IF NOT EXISTS selection (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    profile_id uuid not null references profile(id),
    name TEXT not null default '',
    comment TEXT not null default '',
    form jsonb default '{}'::jsonb,

    created_at timestamp default current_timestamp NOT NULL,
    updated_at timestamp default current_timestamp NOT NULL
);

CREATE TABLE IF NOT EXISTS favorite_selection_development (
    development_id BIGINT not null references development(id) ON DELETE cascade,
    selection_id uuid not null references selection(id) ON delete cascade,

    primary key (selection_id, development_id)
);
