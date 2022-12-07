package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Frutas() {
	for i := 0; i <= 100; i++ {
		fmt.Println("Usted tiene ", i, "Frutas")
	}
}

func main() {
	// Inicializando echo
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world!")
	})
	go Frutas()
	e.Logger.Fatal(e.Start(":8080"))

}
