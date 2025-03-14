package encoder

import "testing"

func TestXxx(t *testing.T) {
	tests := []struct {
		text string
		want Table
	}{
		/*
			https://en.wikipedia.org/wiki/Byte_pair_encoding

			count the frequency of each token pair in "aaabdaaabac":

			iteration 1:
			aaabdaaabac
			a=2
			b=0
			d=0
			c=0

			ZabdZabac
			Z=aa

			iteration 2:
			ZabdZabac
			Z=0
			ab=2
			d=0
			a=0
			c=0

			ZYdZYac
			Z=aa
			Y=ab

			iteration 3:
			ZYdZYac
			ZY=2
			d=0
			a=0
			c=0

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
				// freq at iteration 3:
				freq:    map[string]int{"ZY": 2, "d": 0, "a": 0, "c": 0},
				encoded: "XdXac",
			},
		},
	}
	for _, test := range tests {
		t.Log(test.text, test.want)
	}
}
