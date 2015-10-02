CREATE FUNCTION sp_app_user_insert
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
        INSERT INTO     app_users
                        (
                            email,
                            password
                        )
        VALUES          (
                            userEmail,
                            userPassword
                        )
        RETURNING       app_users.id,
                        app_users.email,
                        app_users.is_active,
                        app_users.is_admin
    )
    SELECT      ROW_TO_JSON(inserted.*)
    FROM        inserted;
END;
$$

LANGUAGE plpgsql;
