CREATE FUNCTION sp_api_password_reset_requests_insert
(
    userEmail       TEXT,
    resetCode       TEXT
)

RETURNS TABLE
(
    jdoc JSON
) AS

$$
BEGIN
    DELETE FROM     api_password_reset_requests
    WHERE           api_password_reset_requests.user_id = api_users.id AND
                    api_users.email = userEmail;

    RETURN      QUERY
    WITH inserted AS (
        INSERT INTO     api_password_reset_requests
                        (
                            user_id,
                            code
                        )
        SELECT          api_users.id,
                        resetCode
        FROM            api_users
        WHERE           api_users.email = userEmail
        RETURNING       api_password_reset_requests.id,
                        api_password_reset_requests.user_id,
                        api_password_reset_requests.code,
                        api_password_reset_requests.created_at,
                        api_password_reset_requests.expires_at
    )
    SELECT      ROW_TO_JSON(inserted.*)
    FROM        inserted;

END;
$$

LANGUAGE plpgsql;
