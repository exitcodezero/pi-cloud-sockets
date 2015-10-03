CREATE TABLE IF NOT EXISTS api_users
(
    id              SERIAL PRIMARY KEY,
    email           TEXT NOT NULL,
    password        TEXT NOT NULL,
    is_active       BOOLEAN DEFAULT TRUE,
    is_admin        BOOLEAN DEFAULT FALSE,
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT(NOW() AT TIME ZONE 'UTC'),
    UNIQUE (email)
);
