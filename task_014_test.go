package toml

import (
	"testing"
	"time"
)

func TestTask014PositiveDuration(t *testing.T) {
	var v struct{ Timeout time.Duration }
	if _, e := Decode(`Timeout = "1500ms"`, &v); e != nil {
		t.Fatal(e)
	}
	if v.Timeout != 1500*time.Millisecond {
		t.Fatalf("Timeout=%v", v.Timeout)
	}
}
func TestTask014NegativeDuration(t *testing.T) {
	var v struct{ Offset time.Duration }
	if _, e := Decode(`Offset = "-2s"`, &v); e != nil {
		t.Fatal(e)
	}
	if v.Offset != -2*time.Second {
		t.Fatalf("Offset=%v", v.Offset)
	}
}
