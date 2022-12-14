
CREATE SCHEMA IF NOT EXISTS hack;

CREATE TABLE IF NOT EXISTS hack.person (
    id bigserial PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL
);

CREATE TABLE IF NOT EXISTS hack.invoice (
    id bigserial PRIMARY KEY,
    title text NOT NULL,
    amount numeric NOT NULL
);
