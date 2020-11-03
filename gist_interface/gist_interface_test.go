package gist_interface

import (
	"io"
	"strings"
	"testing"
)

type dummyDoer struct{}

func (d *dummyDoer) doGistsRequest(user string) (io.Reader, error) {
	return strings.NewReader(`
	[
		{"html_url": "https://gist.github.com/example1"},
		{"html_url": "https://gist.github.com/example2"}
	]
`), nil
}

func TestListGists(t *testing.T) {
	c := Client{&dummyDoer{}}

	urls, err := c.ListGists("test")
	if err != nil {
		t.Fatalf("list gists caused error: %s", err)
	}
	if expected := 2; len(urls) != expected {
		t.Fatalf("want %d, go %d", expected, len(urls))
	}
}
