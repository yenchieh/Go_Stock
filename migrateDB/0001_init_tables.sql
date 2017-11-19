-- +goose Up
CREATE TABLE account(
  id SERIAL PRIMARY KEY NOT NULL,
  name TEXT NULL,
  email TEXT NOT NULL,
  password TEXT NOT NULL,
  created_at DATE,
  updated_at DATE
);

-- +goose Down
DROP TABLE account;