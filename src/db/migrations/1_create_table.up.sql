CREATE TABLE IF NOT EXISTS users(
    id         UUID PRIMARY KEY NOT NULL,
    name       VARCHAR(50)             NOT NULL,
    created_at TIMESTAMP        NOT NULL
);

CREATE TABLE IF NOT EXISTS posts(
    id         UUID PRIMARY KEY NOT NULL,
    user_id    UUID                     NOT NULL,
    title      VARCHAR(50)              NOT NULL,
    content    VARCHAR(1000)              NOT NULL,
    created_at TIMESTAMP        NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
);