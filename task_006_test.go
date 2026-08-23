package toml

import ("testing")

func TestTask006KeysIncludeFirstDeclaration(t *testing.T) {
	var out map[string]any; md, err := Decode("first = 1\nsecond = 2\n", &out)
	if err != nil { t.Fatalf("decode failed: %v", err) }
	keys := md.Keys(); if len(keys) != 2 || keys[0].String() != "first" { t.Fatalf("keys=%v, want first then second", keys) }
}

func TestTask006KeysPreserveTableHeader(t *testing.T) {
	var out map[string]any; md, err := Decode("[server]\nport = 80\n", &out)
	if err != nil { t.Fatalf("decode failed: %v", err) }
	keys := md.Keys(); if len(keys) < 1 || keys[0].String() != "server" { t.Fatalf("keys=%v, want table header first", keys) }
}
