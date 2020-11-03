package cov

import "testing"

type Case struct {
	in, out string
}

var cases = []Case{
	{"aaaaa", "1"},
	{"go go go go go fdfafafaf", "8ika"},
}

func TestWords(t *testing.T) {
	for i, c := range cases {
		w := Words(c.in)
		if w != c.out {
			t.Errorf("#%d: Word(%s) got %s; want %s", i, c.in, w, c.out)
		}
	}
}
