package toml

import "testing"

type task015Text string

func (v *task015Text) UnmarshalText(b []byte) error { *v = task015Text(string(b)); return nil }
func TestTask015UnicodeText(t *testing.T) {
	var v struct{ Label task015Text }
	if _, e := Decode(`Label = "你好"`, &v); e != nil {
		t.Fatal(e)
	}
	if v.Label != "你好" {
		t.Fatalf("Label=%q", v.Label)
	}
}
func TestTask015BooleanText(t *testing.T) {
	var v struct{ Flag task015Text }
	if _, e := Decode(`Flag = true`, &v); e != nil {
		t.Fatal(e)
	}
	if v.Flag != "true" {
		t.Fatalf("Flag=%q", v.Flag)
	}
}
