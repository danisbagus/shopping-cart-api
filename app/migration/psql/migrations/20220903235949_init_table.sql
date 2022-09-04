-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id          SERIAL NOT NULL,
    name        VARCHAR(50) NOT NULL,
    email       VARCHAR(100) NOT NULL UNIQUE,
    password    VARCHAR(255) NOT NULL,
    role_id     INT NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    PRIMARY KEY (id)    
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX idx_email ON "users" ("email");
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE roles (
    id          SERIAL NOT NULL,
    name        VARCHAR(50) NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS roles;
-- +goose StatementEnd
