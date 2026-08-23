package toml

import ("testing")

func TestTask009DecodeIntegerSliceLength(t *testing.T) {
	var out struct{ Values []int `toml:"values"` }
	if _, err := Decode("values = [1, 2, 3]", &out); err != nil { t.Fatalf("decode failed: %v", err) }
	if len(out.Values) != 3 { t.Fatalf("length=%d, want 3", len(out.Values)) }
}

func TestTask009ReuseSliceAdjustsLength(t *testing.T) {
	out := struct{ Values []int `toml:"values"` }{Values: make([]int, 5, 8)}
	if _, err := Decode("values = [7, 8]", &out); err != nil { t.Fatalf("decode failed: %v", err) }
	if len(out.Values) != 2 || out.Values[0] != 7 || out.Values[1] != 8 { t.Fatalf("values=%v, want [7 8]", out.Values) }
}
