package token

import "testing"

func TestTypeKeywords(t *testing.T) {
	tests := []struct {
		input    string
		expected TokenType
	}{
		{"পূর্ণসংখ্যা", TYPE_INT},
		{"লেখা", TYPE_STRING},
		{"বুলিয়ান", TYPE_BOOL},
		{"তালিকা", TYPE_ARRAY},
		{"হ্যাশ", TYPE_HASH},
		{"ফাংশন_টাইপ", TYPE_FUNC},
	}

	for _, tt := range tests {
		tok := LookupIdent(tt.input)
		if tok != tt.expected {
			t.Errorf("LookupIdent(%q) = %q, want %q", tt.input, tok, tt.expected)
		}
	}
}
