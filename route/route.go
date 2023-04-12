package route

import (
	"net/http"
	"project_prakerja2023/controller"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RouteInit(e *echo.Echo) {
	//route agar saat path kosong ke file docs
	e.GET("/", func(e echo.Context) error {
		return e.Redirect(http.StatusMovedPermanently, "/docs/")
	})
	//Saat sudah di route doc, bakal memanggil handler default swagger
	e.GET("/docs/*", echoSwagger.WrapHandler)

	e.POST("/books", controller.CreateBook)
	e.GET("/books", controller.GetAllBook)
	e.GET("/books/:id", controller.GetBookById)
	e.PUT("/books/:id", controller.UpdateBook)
	e.DELETE("/books/:id", controller.DeleteBook)
}
