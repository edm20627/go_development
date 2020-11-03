DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id   serial primary key,
  name text not null,
  age  integer not null
);