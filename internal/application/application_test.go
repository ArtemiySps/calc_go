package application

import (
	"fmt"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"text/tabwriter"
)

/*func TestCalcHandler(t *testing.T) {
	TestSucsses := []struct {
		name           string
		expression     string
		expectedStatus int
	}{
		{
			name:           "simple",
			expression:     "2+2",
			expectedStatus: 200,
		},
	}

	TestError := []struct {
		name           string
		expression     string
		expectedStatus int
	}{
		{
			name:           "division by zero",
			expression:     "2+2/0",
			expectedStatus: 400,
		},
		{
			name:           "unexpected symbol",
			expression:     "3+4+a",
			expectedStatus: 400,
		},
	}

	w := tabwriter.NewWriter(os.Stdout, 25, 1, 1, ' ', 0)

	fmt.Fprintln(w, "-------------------------------------------------------------------")

	for _, test := range TestSucsses {
		req := httptest.NewRequest("GET", "http://localhost:8080", nil)
		w := httptest.NewRecorder()
		CalcHandler(w, req)
		resp := w.Result()
		fmt.Println("asdsasda")
		defer resp.Body.Close()
		fmt.Println(resp.StatusCode)

		if resp.StatusCode != 200 {
			fmt.Println("ffffffffff")
			fmt.Fprintf(w, "%s\tERROR!\n", test.name)
			fmt.Fprintf(w, " \t\n")
			fmt.Fprintf(w, " \tresult\n")
			fmt.Fprintf(w, "Expected:\t200\n")
			fmt.Fprintf(w, "Got:\t%d\n", resp.StatusCode)
			fmt.Fprintln(w, "-------------------------------------------------------------------")
		} else {
			fmt.Fprintf(w, "%s\tsucsses\n", test.name)
			fmt.Fprintln(w, "-------------------------------------------------------------------")
		}
	}

	for _, test := range TestError {
		req := httptest.NewRequest("POST", "http://localhost:8080", nil)
		w := httptest.NewRecorder()
		CalcHandler(w, req)
		resp := w.Result()
		fmt.Println(resp.StatusCode)

		if resp.StatusCode != 400 {
			fmt.Fprintf(w, "%s\tERROR!\n", test.name)
			fmt.Fprintf(w, " \t\n")
			fmt.Fprintf(w, " \tresult\n")
			fmt.Fprintf(w, "Expected:\t200\n")
			fmt.Fprintf(w, "Got:\t%d\n", resp.StatusCode)
			fmt.Fprintln(w, "-------------------------------------------------------------------")
		} else {
			fmt.Fprintf(w, "%s\tsucsses\n", test.name)
			fmt.Fprintln(w, "-------------------------------------------------------------------")
		}
	}
	w.Flush()
}*/

type want struct {
	code int
}

type Expression struct {
	Expression string `json:"expression"`
}

func PrintResult(marker bool, name string, code int, want int) {
	w := tabwriter.NewWriter(os.Stdout, 25, 1, 1, ' ', 0)

	if marker {
		fmt.Fprintf(w, "%s\tsucsses\n", name)
		fmt.Fprintln(w, "-------------------------------------------------------------------")
	} else {
		fmt.Println("ffffffffff")
		fmt.Fprintf(w, "%s\tERROR!\n", name)
		fmt.Fprintf(w, " \t\n")
		fmt.Fprintf(w, " \tresult\n")
		fmt.Fprintf(w, "Expected:\t%d\n", want)
		fmt.Fprintf(w, "Got:\t%d\n", code)
		fmt.Fprintln(w, "-------------------------------------------------------------------")
	}
	w.Flush()
}

func TestCalcHandler(t *testing.T) {
	tests := []struct {
		name       string
		expression Expression
		want       want
	}{
		{
			name:       "st200",
			expression: Expression{`{"expression":"2+2"}`},
			want: want{
				code: 200,
			},
		},
		{
			name:       "st400",
			expression: Expression{`{"expression":"2+2+"}`},
			want: want{
				code: 400,
			},
		},
	}

	for _, tt := range tests {
		reader := strings.NewReader(tt.expression.Expression)
		req := httptest.NewRequest("POST", "http://localhost:8080", reader)
		w := httptest.NewRecorder()
		CalcHandler(w, req)
		resp := w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != tt.want.code {
			PrintResult(false, tt.name, resp.StatusCode, tt.want.code)
		} else {
			PrintResult(true, tt.name, resp.StatusCode, tt.want.code)
		}

	}
}
