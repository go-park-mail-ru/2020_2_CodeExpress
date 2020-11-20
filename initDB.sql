\c musicexpress;

DROP TABLE IF EXISTS artists, users, albums, tracks, genres, track_genre, user_track, session CASCADE;

CREATE TABLE artists (
    id serial NOT NULL PRIMARY KEY,
    name varchar(100) NOT NULL UNIQUE,
    description text NOT NULL, --TODO: NOT NULL? везде так
    poster varchar(100) NOT NULL DEFAULT '',
    avatar varchar(100) NOT NULL DEFAULT ''
);

CREATE TABLE users (
    id serial NOT NULL PRIMARY KEY,
    name varchar(64) NOT NULL UNIQUE,
    email varchar(64) NOT NULL UNIQUE,
    password varchar(64) NOT NULL,
    avatar varchar(255) NOT NULL DEFAULT ''
);

CREATE TABLE albums (
    id serial NOT NULL PRIMARY KEY,
    artist_id int NOT NULL REFERENCES artists(id) ON DELETE CASCADE,
    title varchar(100) NOT NULL,
    poster varchar(100) NOT NULL DEFAULT ''
);

CREATE TABLE tracks (
    id serial NOT NULL PRIMARY KEY,
    album_id int NOT NULL REFERENCES albums(id) ON DELETE CASCADE,
    title varchar(100) NOT NULL,
    duration int NOT NULL DEFAULT 0,
    index int NOT NULL DEFAULT 0,
    audio varchar(100) NOT NULL DEFAULT ''
);

CREATE TABLE genres (
    id serial NOT NULL PRIMARY KEY,
    name varchar(100) NOT NULL UNIQUE
);

CREATE TABLE track_genre (
    track_id int NOT NULL REFERENCES tracks(id) ON DELETE CASCADE,
    genre_id int NOT NULL REFERENCES genres(id) ON DELETE CASCADE,
    PRIMARY KEY(track_id, genre_id)
);

CREATE TABLE user_track (
    user_id int NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    track_id int NOT NULL REFERENCES tracks(id) ON DELETE CASCADE,
    PRIMARY KEY(user_id, track_id)
);

CREATE TABLE session (--TODO: sessions
    id varchar(64) NOT NULL PRIMARY KEY,
    userID int NOT NULL REFERENCES users(id) ON DELETE CASCADE,--TODO: user_id
    expire date NOT NULL--TODO: expires
);

CREATE OR REPLACE FUNCTION  count_index() RETURNS trigger AS $emp_stamp$
    BEGIN
        new.index := count(*) + 1 from tracks where tracks.album_id = new.album_id;
        Return new;
    END;
$emp_stamp$ LANGUAGE plpgsql;

CREATE TRIGGER emp_stamp BEFORE INSERT ON tracks
    FOR EACH ROW EXECUTE PROCEDURE count_index();

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO meuser;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO meuser;
