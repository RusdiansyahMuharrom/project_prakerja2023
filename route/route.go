package route

import (
	"project_prakerja2023/controller"

	"github.com/labstack/echo/v4"
)

func RouteInit(e *echo.Echo) {
	e.POST("/books", controller.CreateBook)
	e.GET("/books", controller.GetAllBook)
	e.GET("/books/:id", controller.GetBookById)
	e.PUT("/books/:id", controller.UpdateBook)
	e.DELETE("/books/:id", controller.DeleteBook)
}
