package slice

import "testing"

func TestPractice(t *testing.T) {
	s1 := []byte{'p', 'o', 'e', 'm'}
	s2 := s1[2:]
	t.Log(string(s1), string(s2))
	s2[1] = 't'
	t.Log(string(s1), string(s2))
}
