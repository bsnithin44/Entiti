
CREATE TABLE IF NOT EXISTS fpa.entities(
  plural_name text ,
  singular_name text,
  id text,
  unique_code text,
  created_at timestamp,
  updated_at timestamp,

  is_active boolean,

  PRIMARY KEY (id, created_at));

CREATE INDEX ON fpa.entities(plural_name);


