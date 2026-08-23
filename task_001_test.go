package toml

import "testing"

func TestTask001MalformedArrayReturnsError(t *testing.T) {
	var dst map[string]any
	if err := Unmarshal([]byte("items = [1,"), &dst); err == nil {
		t.Fatal("malformed array was accepted; want a syntax error")
	}
}
