CREATE TABLE users (
  id SERIAL,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE,
  hashed_password CHAR(60) NOT NULL,
  created TIMESTAMPTZ NOT NULL
);