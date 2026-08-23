package toml

import ("testing")

func TestTask007ReportsUndecodedTopLevelKey(t *testing.T) {
	var out struct{ Known int `toml:"known"` }; md, err := Decode("known = 1\nextra = 2\n", &out)
	if err != nil { t.Fatalf("decode failed: %v", err) }
	u := md.Undecoded(); if len(u) != 1 || u[0].String() != "extra" { t.Fatalf("undecoded=%v, want extra", u) }
}

func TestTask007ReportsUndecodedNestedKey(t *testing.T) {
	var out struct{ Server struct{ Port int `toml:"port"` } `toml:"server"` }; md, err := Decode("[server]\nport = 80\nhost = \"x\"\n", &out)
	if err != nil { t.Fatalf("decode failed: %v", err) }
	u := md.Undecoded(); if len(u) != 1 || u[0].String() != "server.host" { t.Fatalf("undecoded=%v, want server.host", u) }
}
