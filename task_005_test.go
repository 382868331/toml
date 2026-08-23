package toml

import ("testing")

func TestTask005IntegerTypeMetadata(t *testing.T) {
	var out map[string]any; md, err := Decode("count = 3", &out)
	if err != nil { t.Fatalf("decode failed: %v", err) }
	if got := md.Type("count"); got != "Integer" { t.Fatalf("type=%q, want Integer", got) }
}

func TestTask005ArrayTypeMetadata(t *testing.T) {
	var out map[string]any; md, err := Decode("names = [\"a\", \"b\"]", &out)
	if err != nil { t.Fatalf("decode failed: %v", err) }
	if got := md.Type("names"); got != "Array" { t.Fatalf("type=%q, want Array", got) }
}
