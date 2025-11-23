-- +migrate Up
-- Table Definition ----------------------------------------------
CREATE TABLE IF NOT EXISTS clients (
  id              SERIAL PRIMARY KEY,
  uuid            UUID NOT NULL DEFAULT gen_random_uuid(),
  first_name      VARCHAR(255) NOT NULL,
  last_name       VARCHAR(255) NOT NULL,
  middle_name     VARCHAR(255),
  suffix          VARCHAR(50),
  birthdate       DATE,
  pcn             VARCHAR(100),
  sex             BOOLEAN,
  mobile_number   VARCHAR(50),
  civil_status_id INTEGER,
  occupation_id   INTEGER,
  monthly_salary  NUMERIC(14,2)
);

-- Indices -------------------------------------------------------
CREATE UNIQUE INDEX IF NOT EXISTS idx_clients_id ON clients (id);
CREATE UNIQUE INDEX IF NOT EXISTS idx_clients_uuid ON clients (uuid);

-- +migrate Down
DROP TABLE "clients";
