package main

import (
	"github.com/ArtemiySps/calc_go/internal/application"
)

func main() {
	app := application.New()
	//app.Run()
	app.RunServer()
}
