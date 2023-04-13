package main

import (
	"project_prakerja2023/configuration"
	"project_prakerja2023/docs"
	"project_prakerja2023/route"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @title                      API Pengolahan Data Buku Perpustakaan
// @version                    Development 1.0
// @description                This is a dedicated API Pengolahan Data Buku Perpustakaan.
// @host                       projectprakerja2023-production.up.railway.app
// @BasePath                   /
// @schemes 				   http

func main() {
	//Inisialisasi database dan migration
	configuration.DatabaseInit()

	e := echo.New()
	//Set cors
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: configuration.AllowOrigins(),
		AllowMethods: configuration.AllowMethods(),
	}))

	//Inisialisasi route
	route.RouteInit(e)

	//set swagger host
	docs.SwaggerInfo.Host = configuration.Get_env("APP_HOST")

	e.Logger.Fatal(e.Start(":" + configuration.GetPort()))
}
