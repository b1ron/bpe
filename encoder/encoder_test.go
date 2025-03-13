package encoder

import "testing"

func TestXxx(t *testing.T) {
	tests := []struct {
		text string
		want Table
	}{
		/*
			XdXac
			X=ZY
			Y=ab
			Z=aa
			this data cannot be compressed further by byte pair encoding because there are no repeated pairs of characters
			https://en.wikipedia.org/wiki/Byte_pair_encoding
		*/
		{"aaabdaaabac", Table{m: map[string]string{"Z": "aa", "Y": "ab", "X": "ZY"}, encoded: "XdXac"}},
	}
	for _, test := range tests {
		t.Log(test.text, test.want)
	}
}
