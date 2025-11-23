-- +migrate Up
-- Table Definition ----------------------------------------------
CREATE TABLE IF NOT EXISTS beneficiaries (
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
CREATE INDEX IF NOT EXISTS idx_beneficiaries_uuid ON beneficiaries (uuid);
CREATE INDEX IF NOT EXISTS idx_beneficiaries_id ON beneficiaries (id);

-- +migrate Down
DROP TABLE "beneficiaries";

