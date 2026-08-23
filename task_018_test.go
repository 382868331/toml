package toml

import "testing"

func TestTask018UnicodeString(t *testing.T) {
	var v struct{ Name string }
	if _, e := Decode(`Name = "世界"`, &v); e != nil {
		t.Fatal(e)
	}
	if v.Name != "世界" {
		t.Fatalf("Name=%q", v.Name)
	}
}
