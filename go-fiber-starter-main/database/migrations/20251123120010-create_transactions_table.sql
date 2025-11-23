-- +migrate Up
-- Table Definition ----------------------------------------------
CREATE TABLE IF NOT EXISTS transactions (
  id                    SERIAL PRIMARY KEY,
  uuid                  UUID NOT NULL DEFAULT gen_random_uuid(),
  client_id             INTEGER NOT NULL,
  beneficiary_id        INTEGER,
  date                  DATE,
  program_id            INTEGER,
  is_returning          BOOLEAN DEFAULT FALSE,
  assisted_through_id   INTEGER,
  admission_mode_id     INTEGER,
  client_age            INTEGER,
  bene_age              INTEGER,
  amount_needed         NUMERIC(14,2),
  amount_provided       NUMERIC(14,2),
  problem_presented     TEXT,
  assessment            TEXT,
  CONSTRAINT fk_transactions_client FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE RESTRICT,
  CONSTRAINT fk_transactions_beneficiary FOREIGN KEY (beneficiary_id) REFERENCES beneficiaries(id) ON DELETE SET NULL,
  CONSTRAINT fk_transactions_program FOREIGN KEY (program_id) REFERENCES programs(id) ON DELETE SET NULL,
  CONSTRAINT fk_transactions_assisted_through FOREIGN KEY (assisted_through_id) REFERENCES assisted_throughs(id) ON DELETE SET NULL,
  CONSTRAINT fk_transactions_admission_mode FOREIGN KEY (admission_mode_id) REFERENCES admission_modes(id) ON DELETE SET NULL
);

-- Indices -------------------------------------------------------
CREATE INDEX IF NOT EXISTS idx_transactions_uuid ON transactions (uuid);
CREATE INDEX IF NOT EXISTS idx_transactions_id ON transactions (id);

-- +migrate Down
DROP TABLE "transactions";
