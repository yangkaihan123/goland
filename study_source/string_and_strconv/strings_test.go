package string_conversion

import (
	"strings"
	"testing"
)

func TestStringsHasPrefixHasSuffix(t *testing.T) {
	var str string = "the example string"
	t.Logf("T/F, Does the string \"%s\" have prefix %s?", str, "th")
	t.Logf("%t\n", strings.HasPrefix(str, "th"))
	t.Logf("T/F, Does the string \"%s\" have suffix %s?", str, "ng")
	t.Logf("%t\n", strings.HasSuffix(str, "ng"))
}

func TestContains(t *testing.T) {
	var a = "example string test"
	t.Logf("%t\n", strings.Contains(a, "test"))
	t.Logf("%t\n", strings.Contains(a, "qw"))
}

func TestIndex(t *testing.T) {
	var a = "Hi i am Marc"
	t.Logf("the position of \"Marc\" is : %d ", strings.Index(a, "Marc"))
	t.Logf("the position of the first instance of \"Hi\" is : %d", strings.Index(a, "Hi"))
	t.Logf("the position of the last instance os \"am\" is %d", strings.LastIndex(a, "am"))
	t.Logf("the position of \"www\" is %d", strings.Index(a, "www"))
}

func TestReplace(t *testing.T) {
	var a string = "old old old"
	t.Logf("the example string is %s, the new string is  %s ", a, strings.Replace(a, "old", "new", 1))
	t.Logf("the example string is %s, the new string is  %s ", a, strings.Replace(a, "old", "new", -1))
}

func TestCount(t *testing.T) {
	var a = "hello, how is it going, Hugo!"
	var b = "ggggg"
	t.Logf("Number of h's in %s is %d", a, strings.Count(a, "h"))
	t.Logf("Number of g's in %s is %d", b, strings.Count(b, "g"))

}

func TestRepeat(t *testing.T) {
	var oldString = "hi there!"
	var newString = strings.Repeat(oldString, 3)
	t.Logf("the old is %s , the new is %s ", oldString, newString)
}

func TestLowerAndUpper(t *testing.T) {
	var ex = "abcd , ABCD"
	t.Logf("the ex strings are %s", ex)
	t.Logf("the Lower strings are %s", strings.ToLower(ex))
	t.Logf("the Upper strings are %s", strings.ToUpper(ex))
}

func TestTrim(t *testing.T) {
	var ex = "abcdefa"
	t.Logf("the ex is %s", ex)
	t.Logf("cut a is %s", strings.Trim(ex, "a"))
	t.Logf("cut left a is %s", strings.TrimLeft(ex, "a"))
	t.Logf("cut right a is %s", strings.TrimRight(ex, "a"))
	var ex2 = "  asds  "
	t.Logf("the ex2 is %s,", ex2)
	t.Logf("cut space is %s,", strings.TrimSpace(ex2))
}
