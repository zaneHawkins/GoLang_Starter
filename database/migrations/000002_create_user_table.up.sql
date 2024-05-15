CREATE TABLE IF NOT EXISTS users
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    first_name varChar(255) NOT NULL,
    last_name varChar(255) NOT NULL,
    email varChar(255) NOT NULL,
    password varChar(255) NOT NULL,
    date_joined TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status boolean NOT NULL DEFAULT FALSE,
    PRIMARY KEY(id),
    UNIQUE (email)
);