-- +migrate Up
-- Table Definition ----------------------------------------------
CREATE TABLE IF NOT EXISTS assisted_throughs (
  id   SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

-- +migrate Down
DROP TABLE "assisted_throughs";
