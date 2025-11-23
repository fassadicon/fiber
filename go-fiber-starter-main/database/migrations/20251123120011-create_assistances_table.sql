-- +migrate Up
-- Table Definition ----------------------------------------------
CREATE TABLE IF NOT EXISTS assistances (
  id                    SERIAL PRIMARY KEY,
  uuid                  UUID NOT NULL DEFAULT gen_random_uuid(),
  transaction_id        INTEGER NOT NULL,
  assistance_type_id    INTEGER NOT NULL,
  amount_needed         NUMERIC(14,2),
  amount_provided       NUMERIC(14,2),
  purpose               VARCHAR(1024),
  release_mode_id       INTEGER,
  diagnosis             VARCHAR(1024),
  social_worker_id      INTEGER,
  approving_officer_id  INTEGER,
  CONSTRAINT fk_assistances_transaction FOREIGN KEY (transaction_id) REFERENCES transactions(id) ON DELETE CASCADE,
  CONSTRAINT fk_assistances_type FOREIGN KEY (assistance_type_id) REFERENCES assistance_types(id) ON DELETE RESTRICT
);

-- Indices -------------------------------------------------------
CREATE INDEX IF NOT EXISTS idx_assistances_uuid ON assistances (uuid);
CREATE INDEX IF NOT EXISTS idx_assistances_id ON assistances (id);

-- +migrate Down
DROP TABLE "assistances";
