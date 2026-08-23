package toml

import (
	"strings"
	"testing"
)

func TestTask020ArrayNoTrailingComma(t *testing.T) {
	b, e := Marshal(struct{ Values []int }{[]int{1, 2}})
	if e != nil {
		t.Fatal(e)
	}
	if strings.Contains(string(b), "2, ]") {
		t.Fatalf("trailing comma: %s", b)
	}
}
func TestTask020SingletonNoTrailingComma(t *testing.T) {
	b, e := Marshal(struct{ Values []int }{[]int{7}})
	if e != nil {
		t.Fatal(e)
	}
	if strings.Contains(string(b), "7, ]") {
		t.Fatalf("trailing comma: %s", b)
	}
}
