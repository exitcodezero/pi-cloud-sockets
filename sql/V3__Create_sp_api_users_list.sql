CREATE FUNCTION sp_api_users_list
(
    skip INTEGER,
    take INTEGER
)

RETURNS TABLE
(
    record_count INTEGER,
    jdoc JSON
) AS

$$
DECLARE     recordCount INTEGER;

BEGIN
    SELECT      COUNT(*) INTO recordCount
    FROM        api_users;

    RETURN      QUERY

    WITH result AS (
        SELECT          api_users.id,
                        api_users.email,
                        api_users.is_active,
                        api_users.is_admin,
                        api_users.created_at
        FROM            api_users
        ORDER BY        api_users.id
        OFFSET          skip
        LIMIT           take
    )

    SELECT      recordCount,
                JSON_AGG(result.*)
    FROM        result;
END;
$$

LANGUAGE plpgsql;
