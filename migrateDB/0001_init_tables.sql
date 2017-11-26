-- +goose Up
CREATE TABLE account(
  id SERIAL PRIMARY KEY NOT NULL,
  name TEXT NULL,
  email TEXT NOT NULL,
  password TEXT NOT NULL,
  created_at DATE,
  updated_at DATE
);

CREATE TABLE company_list(
  id SERIAL PRIMARY KEY NOT NULL,
  symbol TEXT UNIQUE NOT NULL,
  name TEXT NOT NULL,
  last_sale FLOAT,
  market_cap FLOAT,
  ipo_year TEXT NULL,
  sector TEXT NULL,
  industry TEXT NULL,
  summary_quote TEXT
);

-- +goose Down
DROP TABLE account;