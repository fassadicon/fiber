-- +migrate Up
-- Table Definition ----------------------------------------------
CREATE TABLE IF NOT EXISTS assistance_types (
  id                     SERIAL PRIMARY KEY,
  assistance_category_id INTEGER NOT NULL,
  name                   VARCHAR(255) NOT NULL,
  CONSTRAINT fk_assisttype_category FOREIGN KEY (assistance_category_id) REFERENCES assistance_categories(id) ON DELETE RESTRICT
);

-- Indices -------------------------------------------------------
CREATE INDEX IF NOT EXISTS idx_assistancetype_category_id ON assistance_types (assistance_category_id);
CREATE INDEX IF NOT EXISTS idx_assistance_types_id ON assistance_types (id);
CREATE INDEX IF NOT EXISTS idx_assistance_types_name ON assistance_types (name);

-- +migrate Down
DROP TABLE "assistance_types";
