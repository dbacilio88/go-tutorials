CREATE TABLE "users"
(
    "id"         bigserial PRIMARY KEY,
    "username"   varchar   NOT NULL,
    "password"   varchar   NOT NULL,
    "role"       varchar   NOT NULL,
    "created_at" timestamptz NOT NULL default (now())
);