-- +migrate Up
-- Table Definition ----------------------------------------------
CREATE TABLE IF NOT EXISTS client_beneficiary (
  client_id      INTEGER NOT NULL,
  beneficiary_id INTEGER NOT NULL,
  PRIMARY KEY (client_id, beneficiary_id),
  CONSTRAINT fk_cb_client FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE,
  CONSTRAINT fk_cb_beneficiary FOREIGN KEY (beneficiary_id) REFERENCES beneficiaries(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE "client_beneficiary";
