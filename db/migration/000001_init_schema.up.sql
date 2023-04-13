CREATE TABLE "accounts" (
    "id" INTEGER PRIMARY KEY,
    "fullname" VARCHAR NOT NULL,
    "nickname" VARCHAR NOT NULL,
    "username" VARCHAR NOT NULL,
    "email" VARCHAR NOT NULL,
    "password" VARCHAR NOT NULL,
    "role" ENUM ("admin", "user") NOT NULL,
    "age" INTEGER NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMP NOT NULL DEFAULT (now())
);

CREATE TABLE "tickets" (
    "id" INTEGER PRIMARY KEY,
    "category_id" INTEGER NOT NULL,
    "title" VARCHAR NOT NULL,
    "description" VARCHAR NOT NULL,
    "price" FLOAT NOT NULL,
    "seat" INTEGER NOT NULL,
    "duration" TIMESTAMP,
    "booked_at" TIMESTAMP NOT NULL DEFAULT (now()),
    "created_at" TIMESTAMP NOT NULL DEFAULT (now()),
    "deleted_at" TIMESTAMP NOT NULL DEFAULT (now())
    -- FOREIGN KEY ("category_id") REFERENCES "categories" ("id")
);

CREATE TABLE "categories" (
    "id" INTEGER PRIMARY KEY,
    "type" VARCHAR NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMP NOT NULL DEFAULT (now()),
    "deleted_at" TIMESTAMP NOT NULL DEFAULT (now())
);

ALTER TABLE "tickets" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");