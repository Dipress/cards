CREATE TABLE IF NOT EXISTS cards (
  id            SERIAL PRIMARY KEY,
  user_id       INT NOT NULL,
  word          VARCHAR(255) NOT NULL,
  transcription VARCHAR(255) NOT NULL,
  translation   VARCHAR(255) NOT NULL,

  /* timestamp */
  created_at	  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at	  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
