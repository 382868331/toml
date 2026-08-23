package toml

import ("testing")

func TestTask008MissingFileReturnsError(t *testing.T) {
	var out map[string]any
	if _, err := DecodeFile("testdata/definitely-missing-task008.toml", &out); err == nil { t.Fatal("missing file returned nil error") }
}

func TestTask008DirectoryPathReturnsError(t *testing.T) {
	var out map[string]any
	if _, err := DecodeFile("testdata", &out); err == nil { t.Fatal("directory path returned nil error") }
}
