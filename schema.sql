CREATE TABLE address (
  id   BIGSERIAL PRIMARY KEY,
  type text      NOT NULL,
  name text      NOT NULL,
  parent INTEGER NULL,
  FOREIGN KEY (parent) REFERENCES address(id)
);