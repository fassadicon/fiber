-- +migrate Up
-- Table Definition ----------------------------------------------
CREATE TABLE IF NOT EXISTS admission_modes (
  id   SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

-- +migrate Down
DROP TABLE "admission_modes";
