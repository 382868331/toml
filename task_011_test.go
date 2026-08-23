package toml

import "testing"

func TestTask011TrueBoolean(t *testing.T) {
	var v struct{ Enabled bool }
	if _, e := Decode("Enabled = true", &v); e != nil {
		t.Fatal(e)
	}
	if !v.Enabled {
		t.Fatal("Enabled=false, want true")
	}
}
func TestTask011FalseBoolean(t *testing.T) {
	var v struct{ Enabled bool }
	if _, e := Decode("Enabled = false", &v); e != nil {
		t.Fatal(e)
	}
	if v.Enabled {
		t.Fatal("Enabled=true, want false")
	}
}
