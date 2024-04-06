-- +goose Up
-- +goose StatementBegin
CREATE TABLE address (
  id   BIGSERIAL PRIMARY KEY,
  type text      NOT NULL,
  name text      NOT NULL,
  parent_id INTEGER NULL,
  FOREIGN KEY (parent_id) REFERENCES address(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE address;
-- +goose StatementEnd
