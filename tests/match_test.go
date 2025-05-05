package tests

import (
	"regexengine/engine"
	"testing"
)

func TestMatch(t *testing.T) {
	tests := []struct {
		pattern string
		input   string
		want    bool
	}{
		{"^abc", "abcdef", true},
		{"^abc", "aabcdef", false},
		{"abc$", "xxabc", true},
		{"abc$", "abcxx", false},
		{"^abc$", "abc", true},
		{"^abc$", "abcc", false},
	}

	for _, tt := range tests {
		data := engine.Data{Input: tt.input, Pattern: tt.pattern}
		result := data.Match()
		if result != tt.want {
			t.Errorf("Match(%q, %q) = %v; want %v", tt.pattern, tt.input, result, tt.want)
		}
	}
}
