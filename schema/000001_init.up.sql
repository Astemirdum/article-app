CREATE TABLE users (
    id serial                   NOT NULL UNIQUE,
    name varchar(50)            NOT NULL,
    hash_password varchar(255)  NOT NULL UNIQUE,
    email varchar(50)           NOT NULL UNIQUE
);

CREATE TABLE articles (
    id serial           NOT NULL UNIQUE,
    title varchar(256)  NOT NULL,
    thesis text         NOT NULL,
    pub_time time       NOT NULL DEFAULT NOW()
);

CREATE TABLE user_article (
    id serial                               NOT NULL UNIQUE,
    user_id int REFERENCES users(id)        ON DELETE CASCADE NOT NULL,
    article_id int REFERENCES articles(id)  ON DELETE CASCADE NOT NULL
);
