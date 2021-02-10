package string_test

import "testing"

// string 其实是不可变的byte切片
func TestString(t *testing.T) {
	s := "\xE4\xB8\xA5"
	t.Logf("string val %s", s)
	t.Logf("long han %d", len(s))
	// len求string的长度求的是指其byte的长度，并不是其字符的长度
	s1 := "\xE4\xA8\xA5\xFF"
	t.Logf("string  is %s", s1)
	t.Logf("long han %d", len(s1))
}

func TestUnicode(t *testing.T) {
	s := "中"
	t.Logf("han len %d", len(s))
	c := []rune(s)
	t.Logf("rune %v and len %d", c, len(c))
	t.Logf("unicode %x", c[0])
	t.Logf("utf8 %x", s)

}
