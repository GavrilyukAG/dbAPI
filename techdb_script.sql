DROP TABLE IF EXISTS votes CASCADE;
DROP TABLE IF EXISTS posts CASCADE;
DROP TABLE IF EXISTS threads CASCADE;
DROP TABLE IF EXISTS forums CASCADE;
DROP TABLE IF EXISTS users CASCADE;

CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE "users" (
    nickname CITEXT PRIMARY KEY,
    about TEXT,
    email CITEXT UNIQUE NOT NULL,
    fullname CITEXT
);

CREATE TABLE forums (
    slug CITEXT PRIMARY KEY,-- UNIQUE NOT NULL,
    posts INTEGER DEFAULT 0,
    threads INTEGER DEFAULT 0,
    title TEXT,
    "user" CITEXT REFERENCES "users" (nickname)
);

CREATE TABLE threads (
    id SERIAL PRIMARY KEY,
    author CITEXT REFERENCES "users" (nickname),
    "forum" CITEXT REFERENCES "forums" (slug),
    slug CITEXT UNIQUE,
    title TEXT NOT NULL,
    message TEXT,
    created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    votes INTEGER DEFAULT 0
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    author CITEXT REFERENCES "users" (nickname),
    "forum" CITEXT REFERENCES "forums" (slug),
    "thread" INTEGER REFERENCES "threads" (id),
    message TEXT,
    created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    isEdited BOOLEAN,
    parent INTEGER DEFAULT 0,
    "path" INTEGER[]
);

CREATE TABLE votes (
    nickname CITEXT REFERENCES "users" (nickname),
    thread INTEGER REFERENCES "threads" (id),
    voice INTEGER
);

-- tpforumdb=# select id, path, parent from posts where path && (select array[path[1]] as parent_id from posts where path && array[2016]) order by path;
-- select * from (select id, path, row_number () over (order by path) from posts) p where row_number > ...

-- WITH p AS (select id, path, row_number () over (order by path) from posts)
-- SELECT id, path from p
-- where p.row_number > (select row_number from p where p.id=2016);
