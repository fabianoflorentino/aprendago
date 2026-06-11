package trim

import "testing"

func TestNew(t *testing.T) {
	s := New()
	if s == nil {
		t.Fatal("expected non-nil *Space from New()")
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "trims leading and trailing whitespace",
			input: "  hello world  ",
			want:  "hello world",
		},
		{
			name:  "trims each line",
			input: "  line1  \n  line2  ",
			want:  "line1\nline2",
		},
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
		{
			name:  "only whitespace",
			input: "  \n  \n  ",
			want:  "",
		},
		{
			name:  "no whitespace to trim",
			input: "hello\nworld",
			want:  "hello\nworld",
		},
	}

	s := New()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := s.String(tt.input)
			if got != tt.want {
				t.Errorf("String() = %q, want %q", got, tt.want)
			}
		})
	}
}
