CREATE TABLE snippets (
  id SERIAL,
  title VARCHAR(100) NOT NULL,
  content TEXT NOT NULL,
  created TIMESTAMPTZ NOT NULL,
  expires TIMESTAMPTZ NOT NULL
);

CREATE INDEX idx_snippets_created ON snippets(created);

CREATE TABLE users (
  id SERIAL,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE,
  hashed_password CHAR(60) NOT NULL,
  created TIMESTAMPTZ NOT NULL
);

INSERT INTO users (name, email, hashed_password, created) VALUES (
  'Alice Jones', 
  'alice@example.com', 
  '$2a$12$NuTjWXm3KKntReFwyBVHyuf/to.HEwTy.eS206TNfkGfr6HzGJSWG', 
  '2022-01-01 10:00:00'
);