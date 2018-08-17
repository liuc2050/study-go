package comma

import "testing"

func TestCommaFloat(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"", ""},
		{"12345", "12,345"},
		{"123456", "123,456"},
		{"1", "1"},
		{"12", "12"},
		{"123", "123"},
		{"1234", "1,234"},
		{"-123456", "-123,456"},
		{"-123456.54312", "-123,456.54312"},
		{"-1.23456e-1234", "-1.23456e-1,234"},
		{"-1.234E123", "-1.234E123"},
		{"+1.234E123", "+1.234E123"},
	}
	for _, test := range tests {
		ret := commaFloat(test.input)
		if test.want != ret {
			t.Errorf("commaFloat(%q) = %q, want %q\n", test.input, ret, test.want)
		}
	}
}
