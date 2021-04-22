CREATE TABLE "topic" (
    "id"    SERIAL PRIMARY KEY,
    "name"  varchar NOT NULL UNIQUE
);

CREATE TYPE status AS ENUM ('draft', 'deleted', 'published');

CREATE TABLE "news" (
    "id"        SERIAL PRIMARY KEY,
    "title"     varchar(120) NOT NULL UNIQUE,
    "author"    varchar NOT NULL,
    "status"    status NOT NULL,
    "published_date" DATE,
    "article"   TEXT NOT NULL,
    "topic"     varchar NOT NULL,
    FOREIGN KEY ("topic") REFERENCES topic(name) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE "tag" (
    "id"    SERIAL PRIMARY KEY,
    "name"  varchar NOT NULL UNIQUE
);

CREATE TABLE "news_tag" (
    "tag_id"  INT REFERENCES "tag" ("id") ON UPDATE CASCADE ON DELETE CASCADE,
    "news_id" INT REFERENCES "news" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
