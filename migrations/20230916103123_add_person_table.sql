-- +goose Up
-- +goose StatementBegin
CREATE TABLE persons (
    id int NOT NULL,
    first_name text,
    last_name text,
    email text UNIQUE,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE persons;
-- +goose StatementEnd
