package main

import (
	"project_prakerja2023/configuration"

	"github.com/labstack/echo/v4"
)

func main() {
	configuration.DatabaseInit()
	e := echo.New()
	e.Logger.Fatal(e.Start(":3002"))
}
