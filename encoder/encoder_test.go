package encoder

import "testing"

func TestXxx(t *testing.T) {
	tests := []struct {
		text string
		want Table
	}{
		/*
			https://en.wikipedia.org/wiki/Byte_pair_encoding

			count the frequency of each token in aaabdaaabac:

			iteration 1:
			a=6
			b=2
			d=1
			c=1

			ZabdZabac
			Z=aa

			iteration 2:
			ab=2
			d=1
			c=1

			ZYdZYac
			Y=ab
			Z=aa

			iteration 3:
			ZY=2

			XdXac
			X=ZY
			Y=ab
			Z=aa
			this data cannot be compressed further by byte pair encoding because there are no repeated pairs of characters
		*/
		{"aaabdaaabac",
			Table{
				t: map[string]string{"Z": "aa", "Y": "ab", "X": "ZY"},
				// frequency of each token, used to determine which token can be encoded
				// the most frequently used token is encoded first
				// we should delete the token from the frequency map after it is encoded and replace it with the encoded token
				freq:    map[string]int{"Z": 2, "Y": 2, "d": 1, "c": 1},
				encoded: "XdXac",
			},
		},
	}
	for _, test := range tests {
		t.Log(test.text, test.want)
	}
}
