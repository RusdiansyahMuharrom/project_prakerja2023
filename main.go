package main

import (
	"project_prakerja2023/configuration"
	"project_prakerja2023/route"

	"github.com/labstack/echo/v4"
)

func main() {
	//Inisialisasi database dan migration
	configuration.DatabaseInit()

	e := echo.New()
	//Inisialisasi route
	route.RouteInit(e)

	e.Logger.Fatal(e.Start(":3002"))
}
