-- +migrate Up
CREATE TABLE artist (
  id INT PRIMARY KEY NOT NULL,
  name TEXT NOT NULL
);

-- +migrate Down
DROP TABLE artist;

