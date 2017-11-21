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
  symbol CHAR(10) UNIQUE NOT NULL,
  name TEXT NOT NULL,
  last_sale NUMERIC(4, 4),
  market_cap NUMERIC(15, 2),
  ipo_year int,
  sector TEXT NULL,
  industry TEXT NULL,
  summary_quote TEXT
);

-- +goose Down
DROP TABLE account;