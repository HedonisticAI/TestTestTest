-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Users (
ID serial PRIMARY KEY,
Name VARCHAR(255) NOT NULL,
Surname VARCHAR(255) NOT NULL,
Patronymic VARCHAR(255),
Nation VARCHAR(255) NOT NULL,
Gender VARCHAR(255) NOT NULL,
Age INTEGER
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Users;
-- +goose StatementEnd
