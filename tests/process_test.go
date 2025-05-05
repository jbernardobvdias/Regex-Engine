package tests

import (
	"regexengine/engine"
	"testing"
)

func TestProcess(t *testing.T) {
	tests := []struct {
		pattern string
		input   string
		want    []string
	}{
		{`(\d+)`, "abc123def456", []string{"123", "456"}},
		{`a(.)c`, "abc adc aec", []string{"b", "d", "e"}},
		{`(\w+)@(\w+)\.(\w+)`, "contact me at test@example.com", []string{"test@example.com"}},
		{`(\d{4})-(\d{2})-(\d{2})`, "Date: 2024-12-25", []string{"2024-12-25"}},
	}

	for _, tt := range tests {
		data := engine.Data{Input: tt.input, Pattern: tt.pattern}
		result := data.Process()
		if !slicesEqual(result, tt.want) {
			t.Errorf("Process(%q, %q) = %v; want %v", tt.pattern, tt.input, result, tt.want)
		}
	}
}

func slicesEqual(sli1 []string, sli2 []string) bool {
	len1 := len(sli1)
	if len1 != len(sli2) {
		return false
	}

	for i := 0; i < len1; i++ {
		if sli1[i] != sli2[i] {
			return false
		}
	}

	return true
}
