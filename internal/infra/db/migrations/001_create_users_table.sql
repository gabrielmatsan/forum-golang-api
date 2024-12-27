-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS students (
  id UUID PRIMARY KEY NOT NULL,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
---- create above / drop below ----
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.


DROP TABLE IF EXISTS students;