-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_verify_link (
    user_id UUID NOT NULL,
    token VARCHAR(1000) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    PRIMARY KEY (user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_verify_link;
-- +goose StatementEnd
