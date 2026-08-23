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
func TestTask013NegativeFloat32(t *testing.T) {
	var v struct{ Ratio float32 }
	if _, e := Decode("Ratio = -0.5", &v); e != nil {
		t.Fatal(e)
	}
	if v.Ratio != -.5 {
		t.Fatalf("Ratio=%v, want -0.5", v.Ratio)
	}
}
