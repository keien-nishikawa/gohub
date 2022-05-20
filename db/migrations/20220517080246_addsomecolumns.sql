-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE gohub_user (
  id int,
  last_name varchar(30),
  first_name varchar(30),
  age int
)
INSERT INTO gohub_user (id, last_name, first_name, age) VALUES (1, 'Keien', 'Nishikawa', 26);
INSERT INTO gohub_user (id, last_name, first_name, age) VALUES (2, 'Shunichi', 'Sawada', 25);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE t_user
-- +goose StatementEnd

INSERT INTO gohub_user (id, last_name, first_name, age) VALUES (1, 'Keien', 'Nishikawa', 26);
