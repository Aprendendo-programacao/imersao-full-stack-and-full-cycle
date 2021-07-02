package webserver

import (
	"github.com/Aprendendo-programacao/imersao-full-stack-and-full-cycle/data"
	"github.com/labstack/echo"
	"net/http"
)

type WebServer struct {

}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	e := echo.New()
	e.GET("/products", w.getAll)
	e.Logger.Fatal(e.Start(":8585"))
}

func (w WebServer) getAll(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, data.Products)
}