CREATE TABLE "topic" (
    "id"    SERIAL PRIMARY KEY,
    "name"  varchar(20) NOT NULL UNIQUE
);

CREATE TYPE status AS ENUM ('draft', 'deleted', 'published');

CREATE TABLE "news" (
    "id"        SERIAL PRIMARY KEY,
    "title"     varchar(120) NOT NULL UNIQUE,
    "author"    varchar NOT NULL,
    "status"    status NOT NULL,
    "published_date" DATE,
    "article" T EXT NOT NULL,
);

CREATE TABLE "tag" (
    "id"    SERIAL PRIMARY KEY,
    "name"  varchar NOT NULL UNIQUE,
);

CREATE TABLE "news_tag" (
    "tag_id"  INT REFERENCES "tag" ("id") ON UPDATE CASCADE ON DELETE CASCADE,
    "news_id" INT REFERENCES "news" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);

ALTER TABLE "news" ADD FOREIGN KEY ("topic_id") REFERENCES "topic" ("id") ON DELETE CASCADE;
