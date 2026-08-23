package toml

import ("strings"
"testing")

func TestTask004ArrayDepthLimit(t *testing.T) {
	d := NewDecoder(strings.NewReader("v = [[[1]]]")); d.MaxDepth(2)
	var out map[string]any
	if _, err := d.Decode(&out); err == nil { t.Fatal("array nesting above MaxDepth was accepted") }
}
