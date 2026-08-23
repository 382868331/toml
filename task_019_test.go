package toml

import (
	"encoding/json"
	"testing"
)

func TestTask019JSONNumberDecimal(t *testing.T) {
	var v struct{ N json.Number }
	if _, e := Decode("N = 42", &v); e != nil {
		t.Fatal(e)
	}
	if string(v.N) != "42" {
		t.Fatalf("N=%q", v.N)
	}
}
func TestTask019JSONNumberNegative(t *testing.T) {
	var v struct{ N json.Number }
	if _, e := Decode("N = -31", &v); e != nil {
		t.Fatal(e)
	}
	if string(v.N) != "-31" {
		t.Fatalf("N=%q", v.N)
	}
}
