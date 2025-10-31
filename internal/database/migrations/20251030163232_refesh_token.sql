-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS refresh_token (
    user_id uuid NOT NULL,
    token varchar(255) NOT NULL,
    expires_at timestamp NOT NULL,
    PRIMARY KEY (user_id),
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE refresh_token;
-- +goose StatementEnd
