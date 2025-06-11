package runner

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteToHostFile(t *testing.T) {
	dir := t.TempDir()
	r := &Runner{
		options: &Options{
			OutputDir:  dir,
			JsonOutput: false,
		},
		rootDomains: map[string]bool{"example.com": true},
	}

	if err := r.writeToHostFile("www.example.com", "data"); err != nil {
		t.Fatalf("writeToHostFile returned error: %v", err)
	}

	path := filepath.Join(dir, "example.com.txt")
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read output file: %v", err)
	}
	if got := strings.TrimSpace(string(b)); got != "Hostname: www.example.com" {
		t.Errorf("unexpected file contents: %q", got)
	}

	// Non matching domain should not create a file
	if err := r.writeToHostFile("other.com", "data"); err != nil {
		t.Fatalf("writeToHostFile returned error: %v", err)
	}

	if _, err := os.Stat(filepath.Join(dir, "other.com.txt")); !os.IsNotExist(err) {
		t.Errorf("file created for unmatched domain")
	}
}
