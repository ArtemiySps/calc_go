package calculation

import (
	"fmt"
	"os"
	"testing"
	"text/tabwriter"
)

func TestCalc(t *testing.T) {
	TestSucsses := []struct {
		name           string
		expression     string
		expectedResult float64
	}{
		{
			name:           "simple",
			expression:     "2+2",
			expectedResult: 4,
		},
		{
			name:           "priority 1",
			expression:     "2+2*3+5",
			expectedResult: 13,
		},
		{
			name:           "priority 2",
			expression:     "(5+5)*2+(3*2)",
			expectedResult: 26,
		},
		{
			name:           "hard",
			expression:     "(2-4+(3*5+4)+60-20*(4/2*7)+200)/2",
			expectedResult: -1.5,
		},
	}

	TestError := []struct {
		name        string
		expression  string
		expectedErr error
	}{
		{
			name:        "division by zero",
			expression:  "2+2/0",
			expectedErr: ErrDivisionByZero,
		},
		{
			name:        "unexpected symbol 1",
			expression:  "3+4+a",
			expectedErr: ErrUnexpectedSymbol,
		},
		{
			name:        "unexpected symbol 2",
			expression:  "55*$+56",
			expectedErr: ErrUnexpectedSymbol,
		},
		{
			name:        "two operations in a row",
			expression:  "3++5",
			expectedErr: ErrTwoOperationsInARow,
		},
		{
			name:        "operation at the end",
			expression:  "3+5*",
			expectedErr: ErrOperationAtTheEnd,
		},
		{
			name:        "brackets problem",
			expression:  "3+(5+6",
			expectedErr: ErrBracketsProblem,
		},
		{
			name:        "no numbers",
			expression:  "",
			expectedErr: ErrNoNumbers,
		},
	}

	w := tabwriter.NewWriter(os.Stdout, 25, 1, 1, ' ', 0)
	//w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, '.', tabwriter.AlignRight|tabwriter.Debug)

	fmt.Fprintln(w, "-------------------------------------------------------------------")
	for _, test := range TestSucsses {
		got_value, got_error := Calc(test.expression)
		if got_value != test.expectedResult || got_error != nil {
			// t.Errorf("Test %s:\n  --expected result: %g and error: nil\n  --got result: %g and error: %s", test.name, test.expectedResult, got_value, got_error)

			fmt.Fprintf(w, "%s\tERROR!\n", test.name)
			fmt.Fprintf(w, " \t\n")
			fmt.Fprintf(w, " \tresult\terror\n")
			fmt.Fprintf(w, "Expected:\t%g\tnil\n", test.expectedResult)
			fmt.Fprintf(w, "Got:\t%g\t%s\n", got_value, got_error)
			//w.Flush()
			fmt.Fprintln(w, "-------------------------------------------------------------------")
		} else {
			fmt.Fprintf(w, "%s\tsucsses\n", test.name)
			//w.Flush()
			fmt.Fprintln(w, "-------------------------------------------------------------------")
		}
	}

	for _, test := range TestError {
		got_value, got_error := Calc(test.expression)
		if got_value != 0 || got_error != test.expectedErr {
			// t.Errorf("Test %s:\n  --expected result: %g and error: nil\n  --got result: %g and error: %s", test.name, test.expectedResult, got_value, got_error)

			fmt.Fprintf(w, "%s\tERROR!\n", test.name)
			fmt.Fprintf(w, " \t\n")
			fmt.Fprintf(w, " \tresult\terror\n")
			fmt.Fprintf(w, "Expected:\t0\t%s\n", test.expectedErr)
			fmt.Fprintf(w, "Got:\t%g\t%s\n", got_value, got_error)
			//w.Flush()
			fmt.Fprintln(w, "-------------------------------------------------------------------")
		} else {
			fmt.Fprintf(w, "%s\tsucsses\n", test.name)
			//w.Flush()
			fmt.Fprintln(w, "-------------------------------------------------------------------")
		}
	}
	w.Flush()
}
