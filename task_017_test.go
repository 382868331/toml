package toml

import "testing"

func TestTask017QuoteInKey(t *testing.T) {
	got := Key{`a"b`}.String()
	want := `"a\"b"`
	if got != want {
		t.Fatalf("Key.String=%q, want %q", got, want)
	}
}
