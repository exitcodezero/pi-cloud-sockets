CREATE FUNCTION sp_api_users_delete
(
    userId      INTEGER
)

RETURNS BOOLEAN AS

$$
BEGIN
    DELETE FROM     api_users
    WHERE           api_users.id = userId;
    RETURN          FOUND;
END;
$$

LANGUAGE plpgsql;
