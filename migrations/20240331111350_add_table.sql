-- +goose Up
-- +goose StatementBegin
CREATE TABLE address (
  id   BIGSERIAL PRIMARY KEY,
  type text      NOT NULL,
  name text      NOT NULL,
  parent INTEGER NULL,
  FOREIGN KEY (parent) REFERENCES address(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE address
-- +goose StatementEnd
