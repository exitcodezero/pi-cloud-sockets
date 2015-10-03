CREATE FUNCTION sp_api_users_reset_password
(
    resetCode       TEXT,
    newPassword     TEXT
)

RETURNS BOOLEAN AS

$$
BEGIN
    UPDATE      api_users
    SET         password = newPassword
    FROM        api_password_reset_requests
    WHERE       api_users.id = api_password_reset_requests.user_id AND
                api_password_reset_requests.code = resetCode AND
                api_password_reset_requests.expires_at > NOW() AT TIME ZONE 'UTC';
    RETURN      FOUND;
END;
$$

LANGUAGE plpgsql;
