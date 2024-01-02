CREATE DATABASE snip_snap
  ENCODING 'UTF8'
  LC_COLLATE = 'en_US.UTF-8'
  LC_CTYPE = 'en_US.UTF-8';

CREATE TABLE snippets (
  id SERIAL,
  title VARCHAR(100) NOT NULL,
  content TEXT NOT NULL,
  created TIMESTAMPTZ NOT NULL,
  expires TIMESTAMPTZ NOT NULL
);

CREATE INDEX idx_snippets_created ON snippets(created);

INSERT INTO snippets (title, content, created, expires) VALUES (
  'An old silent pond',
  E'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō', 
  NOW(),
  NOW() + INTERVAL '365 days'
);
INSERT INTO snippets (title, content, created, expires) VALUES (
  'Over the wintry forest',
  E'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki', 
  NOW(),
  NOW() + INTERVAL '365 days'
);
INSERT INTO snippets (title, content, created, expires) VALUES (
  'First autumn morning',
  E'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo', 
  NOW(),
  NOW() + INTERVAL '365 days'
);