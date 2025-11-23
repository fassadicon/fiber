-- +migrate Up
-- Table Definition ----------------------------------------------
CREATE TABLE IF NOT EXISTS programs (
  id   SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

-- +migrate Down
DROP TABLE "programs";
