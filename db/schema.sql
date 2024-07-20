CREATE TABLE IF NOT EXISTS users
(
    id         UUID PRIMARY KEY NOT NULL,
    name       TEXT             NOT NULL,
    project_v2 TEXT
);

CREATE TABLE IF NOT EXISTS repositories
(
    id         UUID PRIMARY KEY NOT NULL,
    owner      TEXT             NOT NULL,
    name       TEXT             NOT NULL,
    created_at TIMESTAMP        NOT NULL,
    FOREIGN KEY (owner) REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS issues
(
    id         UUID PRIMARY KEY NOT NULL,
    url        TEXT             NOT NULL,
    title      TEXT             NOT NULL,
    closed     INTEGER          NOT NULL DEFAULT 0,
    number     INTEGER          NOT NULL,
    repository TEXT             NOT NULL,
    CHECK (closed IN (0, 1)),
    FOREIGN KEY (repository) REFERENCES repositories (id)
);

CREATE TABLE IF NOT EXISTS projects
(
    id    UUID PRIMARY KEY NOT NULL,
    title TEXT             NOT NULL,
    url   TEXT             NOT NULL,
    owner TEXT             NOT NULL,
    FOREIGN KEY (owner) REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS pullrequests
(
    id            UUID PRIMARY KEY NOT NULL,
    base_ref_name TEXT             NOT NULL,
    closed        INTEGER          NOT NULL DEFAULT 0,
    head_ref_name TEXT             NOT NULL,
    url           TEXT             NOT NULL,
    number        INTEGER          NOT NULL,
    repository    TEXT             NOT NULL,
    CHECK (closed IN (0, 1)),
    FOREIGN KEY (repository) REFERENCES repositories (id)
);

CREATE TABLE IF NOT EXISTS projectcards
(
    id          UUID PRIMARY KEY NOT NULL,
    project     TEXT             NOT NULL,
    issue       TEXT,
    pullrequest TEXT,
    FOREIGN KEY (project) REFERENCES projects (id),
    FOREIGN KEY (issue) REFERENCES issues (id),
    FOREIGN KEY (pullrequest) REFERENCES pullrequests (id),
    CHECK (issue IS NOT NULL OR pullrequest IS NOT NULL)
);
