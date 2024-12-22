package application

import (
	"fmt"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"text/tabwriter"
)

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
			name:       "st422-#1",
			expression: Expression{`{"expression":"2+2+"}`},
			want: want{
				code: 422,
			},
		},
		{
			name:       "st422-#2",
			expression: Expression{`{"expression":"2+2++"}`},
			want: want{
				code: 422,
			},
		},
		{
			name:       "st422-#3",
			expression: Expression{`{"expression":"2+(3+4"}`},
			want: want{
				code: 422,
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
