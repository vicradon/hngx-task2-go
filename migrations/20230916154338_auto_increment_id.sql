-- +goose Up
-- +goose StatementBegin
CREATE TABLE new_persons (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name TEXT,
    last_name TEXT,
    email TEXT UNIQUE
);

INSERT INTO new_persons (first_name, last_name, email)
SELECT first_name, last_name, email FROM persons;

DROP TABLE persons;
ALTER TABLE new_persons RENAME TO persons;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE persons
MODIFY COLUMN id INT;
-- +goose StatementEnd
