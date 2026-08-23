package toml

import (
	"strings"
	"testing"
)

func TestTask016EscapesNewline(t *testing.T) {
	b, e := Marshal(struct{ Text string }{"a\nb"})
	if e != nil {
		t.Fatal(e)
	}
	if strings.Contains(string(b), "a\nb") {
		t.Fatalf("raw newline in %q", b)
	}
	if !strings.Contains(string(b), `a\nb`) {
		t.Fatalf("missing escaped newline in %q", b)
	}
}
func TestTask016EscapesQuote(t *testing.T) {
	b, e := Marshal(struct{ Text string }{`a"b`})
	if e != nil {
		t.Fatal(e)
	}
	if !strings.Contains(string(b), `a\"b`) {
		t.Fatalf("missing escaped quote in %q", b)
	}
}
