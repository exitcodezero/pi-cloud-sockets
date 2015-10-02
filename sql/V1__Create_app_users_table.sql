CREATE TABLE IF NOT EXISTS app_users
(
    id          SERIAL PRIMARY KEY,
    email       TEXT NOT NULL,
    password    TEXT,
    is_active   boolean DEFAULT TRUE,
    is_admin    boolean DEFAULT FALSE,
    UNIQUE (email)
);
