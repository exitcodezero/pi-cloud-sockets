CREATE FUNCTION sp_api_users_select
(
    userId      INTEGER
)

RETURNS TABLE
(
    jdoc JSON
) AS

$$
BEGIN
    RETURN      QUERY
    WITH result AS (
        SELECT          api_users.id,
                        api_users.email,
                        api_users.is_active,
                        api_users.is_admin,
                        api_users.created_at
        FROM            api_users
        WHERE           api_users.id = userId
    )
    SELECT      ROW_TO_JSON(result.*)
    FROM        result;

END;
$$

LANGUAGE plpgsql;
