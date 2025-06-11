package utils

import (
	"strings"
	"testing"
	"testing/quick"
)

func TestIsSubdomain(t *testing.T) {
	roots := map[string]bool{
		"example.com": true,
		"test.org":    true,
	}

	tests := []struct {
		domain string
		want   bool
	}{
		{"example.com", true},
		{"www.example.com", true},
		{"foo.bar.example.com", true},
		{"test.org", true},
		{"sub.test.org", true},
		{"other.com", false},
		{"test.org.example.com", false},
	}

	for _, tt := range tests {
		got := IsSubdomain(tt.domain, roots)
		if got != tt.want {
			t.Errorf("IsSubdomain(%q) = %v, want %v", tt.domain, got, tt.want)
		}
	}
}

func sanitize(s string) string {
	var b strings.Builder
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z':
			b.WriteRune(r)
		case r >= 'A' && r <= 'Z':
			b.WriteRune(r + ('a' - 'A'))
		case r >= '0' && r <= '9':
			b.WriteRune(r)
		}
	}
	if b.Len() == 0 {
		b.WriteRune('a')
	}
	return b.String()
}

func TestIsSubdomainProperties(t *testing.T) {
	f := func(root, a, b string) bool {
		root = sanitize(root) + ".com"
		domain := sanitize(a) + "." + sanitize(b) + "." + root
		return IsSubdomain(domain, map[string]bool{root: true})
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
