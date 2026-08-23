package toml

import "testing"

func TestTask013Fraction(t *testing.T) {
	var v struct{ Ratio float64 }
	if _, e := Decode("Ratio = 1.25", &v); e != nil {
		t.Fatal(e)
	}
	if v.Ratio != 1.25 {
		t.Fatalf("Ratio=%v, want 1.25", v.Ratio)
	}
}
