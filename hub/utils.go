package hub

func findConnectionIndex(conn *Connection, cs []*Connection) int {
	for i, c := range cs {
		if c.ID == conn.ID {
			return i
		}
	}
	return -1
}
