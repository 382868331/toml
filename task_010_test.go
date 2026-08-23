package toml

import ("testing")

func TestTask010PositiveIntegerKeepsSign(t *testing.T) {
	var out struct{ Count int `toml:"count"` }
	if _, err := Decode("count = 42", &out); err != nil { t.Fatalf("decode failed: %v", err) }
	if out.Count != 42 { t.Fatalf("count=%d, want 42", out.Count) }
}
