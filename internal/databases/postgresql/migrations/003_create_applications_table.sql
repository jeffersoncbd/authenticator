CREATE TABLE IF NOT EXISTS applications (
    "id"        uuid            PRIMARY KEY         NOT NULL        DEFAULT gen_random_uuid(),
    "name"      VARCHAR(255)                        NOT NULL        UNIQUE
);

---- create above / drop below ----

DROP TABLE IF EXISTS applications;
