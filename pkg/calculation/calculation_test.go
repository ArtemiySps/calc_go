package calculation

import "testing"

type Test struct {
	in  string
	out float64
	err error
}

var tests = []Test{
	{"3*2+(56+6*(34+5/(4-4)-23))*2", 0, divisionByZero},
	{"2+4*a+5-3", 0, unexpectedSymbol},
	{"2*3--2+4", 0, twoOperationsInARow},
	{"2+4+(34-45)*", 0, operationAtTheEnd},
	{"2+(34-(78-3)+4", 0, bracketsProblem},
	{"(2+2)*2", 8, nil},
	{"1+1", 2, nil},
	{"(2+2*2)", 6, nil},
	{"1/2", 0.5, nil},
	{"((4+5)-2)*2", 14, nil},
}

func TestCalc(t *testing.T) {
	for i, test := range tests {
		got_value, got_error := Calc(test.in)
		if got_value != test.out || got_error != test.err {
			t.Errorf("Test №%d:\n--expected %g and %q\n--got %g and %q", i, test.out, test.err, got_value, got_error)
		}
	}
}
