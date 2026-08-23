package toml

import ("strings"
"testing")

func TestTask004ArrayDepthLimit(t *testing.T) {
	d := NewDecoder(strings.NewReader("v = [[[1]]]")); d.MaxDepth(2)
	var out map[string]any
	if _, err := d.Decode(&out); err == nil { t.Fatal("array nesting above MaxDepth was accepted") }
}

func TestTask004TableDepthLimit(t *testing.T) {
	d := NewDecoder(strings.NewReader("[a.b.c]\nv = 1\n")); d.MaxDepth(2)
	var out map[string]any
	if _, err := d.Decode(&out); err == nil { t.Fatal("table nesting above MaxDepth was accepted") }
}
