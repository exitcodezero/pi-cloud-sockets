CREATE TABLE IF NOT EXISTS api_password_reset_requests
(
    id              SERIAL PRIMARY KEY,
    user_id         INTEGER REFERENCES api_users(id) ON DELETE CASCADE,
    code            TEXT NOT NULL,
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT(NOW() AT TIME ZONE 'UTC'),
    expires_at      TIMESTAMP WITH TIME ZONE DEFAULT(NOW() AT TIME ZONE 'UTC' + INTERVAL '1 DAY'),
    UNIQUE (user_id, code)
);
