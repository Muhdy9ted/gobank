CREATE TABLE users (
    username varchar NOT NULL PRIMARY KEY,
    hashed_password varchar NOT NULL,
    full_name varchar NOT NULL,
    email varchar UNIQUE NOT NULL,
    password_changed_at timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    created_at timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "accounts" ("owner");

CREATE UNIQUE INDEX ON "accounts" ("owner", "currency");

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");