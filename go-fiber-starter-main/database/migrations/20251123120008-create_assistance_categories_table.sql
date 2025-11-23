-- +migrate Up
-- Table Definition ----------------------------------------------
CREATE TABLE IF NOT EXISTS assistance_categories (
  id   SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

-- +migrate Down
DROP TABLE "assistance_categories";
