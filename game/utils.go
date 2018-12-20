package game

// BoolToIntString converts true to "1" and false to "0"
func BoolToIntString(b bool) string {
	if b {
		return "1"
	}
	return "_"
}
