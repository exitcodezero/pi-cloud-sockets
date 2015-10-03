CREATE FUNCTION sp_api_users_update
(
    userId      INTEGER,
    userDoc     JSON
)

RETURNS TABLE
(
    jdoc JSON
) AS

$$
BEGIN
    RETURN     QUERY
    WITH updated AS (
        UPDATE          api_users
        SET             api_users.email = CAST(userDoc->>'email' AS TEXT),
                        api_users.password = CAST(userDoc->>'password' AS TEXT),
                        api_users.is_active = CAST(userDoc->>'is_active' AS BOOLEAN),
                        api_users.is_admin = CAST(userDoc->>'is_admin' AS BOOLEAN),
                        api_users.created_at = CAST(userDoc->>'created_at' AS TIMESTAMP)
        WHERE           api_users.id = userId
        RETURNING       api_users.id,
                        api_users.email,
                        api_users.is_active,
                        api_users.is_admin,
                        api_users.created_at
    )
    SELECT      ROW_TO_JSON(updated.*)
    FROM        updated;
END;
$$

LANGUAGE plpgsql;
