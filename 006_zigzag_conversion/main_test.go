package leetcode006

import "testing"

type Case struct {
	String string
	NumRows int
	Expected string
}

func Test_Examples(t *testing.T) {
	cases := []Case{
		Case{ String: "PAYPALISHIRING", NumRows: 3, Expected: "PAHNAPLSIIGYIR", },
	}

	for _, e := range cases {
		actual := convert(e.String, e.NumRows)
		if actual != e.Expected {
			t.Errorf("got string: %v, got num rows: %v, actual: %v", actual, e.NumRows, e.Expected)
		}

	}
}
