package slice_03

import "testing"

func TestRemove(t *testing.T) {
	s := []string{"M", "N", "O", "P", "Q", "R"}
	res := RemoveStringSlice(s, 2, 4)
	t.Log(res)
}

func RemoveStringSlice(s []string, start int, end int) []string {
	result := make([]string, len(s)-(end-start))
	at := copy(result, s[:start])
	copy(result[at:], s[end:])
	return result
}
