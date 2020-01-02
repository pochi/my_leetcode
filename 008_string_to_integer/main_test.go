package leetcode008

import "testing"

type Case struct {
	String string
	Expected int
}

func Test_Examples(t *testing.T) {
	cases := []Case{
		Case{ String: "42", Expected: 42 },
		Case{ String: "-42", Expected: -42 },
		Case{ String: "4193 with words", Expected: 4193 },
		Case{ String: "words and 987", Expected: 0 },
		Case{ String: "-91283472332", Expected: -2147483648 },
		Case{ String: "+1", Expected: 1 },
		Case{ String: "   +0 123", Expected: 0 },
		Case{ String: "    -42", Expected: -42 },		
		Case{ String: "2147483648", Expected: 2147483647 },	
		Case{ String: "-2147483648", Expected: -2147483648 },
		Case{ String: "20000000000000000000", Expected: 2147483647 },
		Case{ String: "+", Expected: 0 },
		Case{ String: "+-2", Expected: 0 },
		Case{ String: "3.14159", Expected: 3 },
	}

	for _, e := range cases {
		actual := myAtoi(e.String)
		if actual != e.Expected {
			t.Errorf("Given str: %v, got integer: %v, actual: %v", e.String, actual, e.Expected)
		}

	}
}
