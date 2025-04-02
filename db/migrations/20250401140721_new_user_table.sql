-- +goose Up
-- +goose StatementBegin
CREATE TABLE Users (
ID serial PRIMARY KEY,
Name VARCHAR(255) NOT NULL,
Surname VARCHAR(255) NOT NULL,
Partonimyc VARCHAR(255),
Nation VARCHAR(255) NOT NULL,
Gender VARCHAR(255) UNIQUE NOT NULL,
Age INTEGER
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Users;
-- +goose StatementEnd
