package toml

import "testing"

func TestTask012UnsignedZero(t *testing.T) {
	var v struct{ Count uint }
	if _, e := Decode("Count = 0", &v); e != nil {
		t.Fatal(e)
	}
	if v.Count != 0 {
		t.Fatalf("Count=%d, want 0", v.Count)
	}
}
