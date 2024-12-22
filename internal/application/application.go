package application

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/ArtemiySps/calc_go/pkg/calculation"
)

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	if len(os.Args) < 2 {
		config.Addr = "8080"
	} else {
		config.Addr = os.Args[1]
	}
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

type Request struct {
	Expression string `json:"expression"`
}

type Answer struct {
	Result float64 `json:"result"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	request := new(Request)

	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		fmt.Fprint(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := calculation.Calc(request.Expression)
	if err != nil {
		if errors.Is(err, calculation.ErrServerError) {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "ERROR: ", err.Error(), "\nStatusCode: ", http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusUnprocessableEntity)
			fmt.Fprint(w, err.Error(), "\nStatusCode: ", http.StatusUnprocessableEntity)
		}
	} else {
		ans := Answer{Result: result}
		var buf bytes.Buffer
		encoder := json.NewEncoder(&buf)
		err := encoder.Encode(ans)
		if err != nil {
			fmt.Fprint(w, err.Error(), "\n", http.StatusInternalServerError)
		}
		fmt.Fprint(w, buf.String())
	}
}

func (a *Application) RunServer() error {
	http.HandleFunc("/", CalcHandler)
	return http.ListenAndServe(":"+a.config.Addr, nil)
}
