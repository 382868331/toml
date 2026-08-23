package toml

import ("testing")

func TestTask009DecodeIntegerSliceLength(t *testing.T) {
	var out struct{ Values []int `toml:"values"` }
	if _, err := Decode("values = [1, 2, 3]", &out); err != nil { t.Fatalf("decode failed: %v", err) }
	if len(out.Values) != 3 { t.Fatalf("length=%d, want 3", len(out.Values)) }
}
