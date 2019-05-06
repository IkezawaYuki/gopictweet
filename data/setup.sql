-- DROP TABLE tweets CASCADE;
-- DROP TABLE users CASCADE;
-- DROP TABLE sessions CASCADE;
-- DROP TABLE comments CASCADE;

CREATE TABLE tweets(
    id serial PRIMARY KEY,
    uuid varchar(64) not null unique,
    text VARCHAR(255),
    image VARCHAR(255),
    created_at TIMESTAMP not NULL
);

CREATE TABLE users(
    id serial PRIMARY KEY,
    uuid varchar(64) not null unique,
    nickname VARCHAR(255),
    email VARCHAR(255) not null,
    password VARCHAR(255) not null,
    created_at TIMESTAMP not null
);


CREATE TABLE sessions (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  email      VARCHAR(255),
  user_id    integer REFERENCES users(id),
  created_at timestamp not null   
);



CREATE TABLE comments(
    id serial PRIMARY KEY,
    uuid varchar(64) not null unique,
    user_id INTEGER REFERENCES users(id),
    tweet_id INTEGER REFERENCES tweets(id),
    text VARCHAR(255),
    created_at TIMESTAMP not null
);

