package toml

import (
	"strings"
	"testing"
)

func TestTask003ExponentFloatRemainsValid(t *testing.T) {
	b, err := Marshal(struct{ Value float64 }{Value: 1e20})
	if err != nil { t.Fatalf("marshal failed: %v", err) }
	if strings.Contains(string(b), "e+20.0") { t.Fatalf("invalid exponent spelling: %s", b) }
	var dst struct{ Value float64 }
	if _, err := Decode(string(b), &dst); err != nil { t.Fatalf("encoded TOML does not round-trip: %v", err) }
}

func TestTask003NegativeExponentRoundTrips(t *testing.T) {
	b, err := Marshal(struct{ Value float64 }{Value: 1e-20})
	if err != nil { t.Fatalf("marshal failed: %v", err) }
	if strings.Contains(string(b), "e-20.0") { t.Fatalf("invalid negative exponent spelling: %s", b) }
	var dst struct{ Value float64 }
	if _, err := Decode(string(b), &dst); err != nil { t.Fatalf("negative exponent did not round-trip: %v", err) }
}
