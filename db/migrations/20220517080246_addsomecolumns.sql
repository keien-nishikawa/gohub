-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE gohub_user (
  id int,
  last_name varchar(30),
  first_name varchar(30),
  age int
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE t_user
-- +goose StatementEnd
