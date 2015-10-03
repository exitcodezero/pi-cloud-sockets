CREATE FUNCTION sp_api_users_insert
(
    userEmail       TEXT,
    userPassword    TEXT
)

RETURNS TABLE
(
    jdoc    JSON
) AS

$$
BEGIN
    RETURN      QUERY
    WITH inserted AS (
        INSERT INTO     api_users
                        (
                            email,
                            password
                        )
        VALUES          (
                            userEmail,
                            userPassword
                        )
        RETURNING       api_users.id,
                        api_users.email,
                        api_users.is_active,
                        api_users.is_admin,
                        api_users.created_at
    )
    SELECT      ROW_TO_JSON(inserted.*)
    FROM        inserted;
END;
$$

LANGUAGE plpgsql;
