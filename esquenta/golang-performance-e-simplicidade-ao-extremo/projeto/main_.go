package main

import (
	"github.com/Aprendendo-programacao/imersao-full-stack-and-full-cycle/data"
	webserver2 "github.com/Aprendendo-programacao/imersao-full-stack-and-full-cycle/webserver"
)

func main() {
	data.LoadData()
	webserver := webserver2.NewWebServer()
	webserver.Serve()
}