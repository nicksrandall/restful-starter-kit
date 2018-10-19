-- +migrate Up
CREATE TABLE artist (
  id UUID PRIMARY KEY NOT NULL,
  name TEXT NOT NULL
);

-- +migrate Down
DROP TABLE artist;

