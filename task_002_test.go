package toml

import "testing"

func TestTask002MetadataKeepsKeyCase(t *testing.T) {
	var dst map[string]any
	md, err := Decode("Name = 1\n", &dst)
	if err != nil { t.Fatalf("decode failed: %v", err) }
	if !md.IsDefined("Name") { t.Fatal("exact-case key Name was not found") }
	if md.IsDefined("name") { t.Fatal("lowercase name unexpectedly matched Name") }
}

func TestTask002NestedMetadataKeepsEachSegmentCase(t *testing.T) {
	var dst map[string]any
	md, err := Decode("[Server]\nPort = 8080\n", &dst)
	if err != nil { t.Fatalf("decode failed: %v", err) }
	if !md.IsDefined("Server", "Port") { t.Fatal("exact nested key was not found") }
	if md.IsDefined("server", "Port") || md.IsDefined("Server", "port") { t.Fatal("nested key lookup ignored case") }
}
