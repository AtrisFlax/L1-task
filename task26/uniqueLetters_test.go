package task26

import "testing"

func TestUnique(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "unique letters", args: struct{ input string }{input: "⌘こんにちは"}, want: true},
		{name: "unique letters", args: struct{ input string }{input: "⌘⌘こんにちは"}, want: false},
		{name: "unique letters", args: struct{ input string }{input: "abcd"}, want: true},
		{name: "unique letters", args: struct{ input string }{input: "abCdefAaf"}, want: false},
		{name: "unique letters", args: struct{ input string }{input: "aabcd"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqueLetters(tt.args.input); got != tt.want {
				t.Errorf("uniqueLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
