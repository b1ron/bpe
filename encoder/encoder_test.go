package encoder

import "testing"

func TestXxx(t *testing.T) {
	tests := []struct {
		text string
		want Table
	}{
		{"aaabdaaabac", Table{m: map[string]string{"Z": "aa", "Y": "ab", "X": "ZY"}, encoded: "XdXac"}},
	}
	for _, test := range tests {
		t.Log(test.text, test.want)
	}
}
