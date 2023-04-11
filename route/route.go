package route

import (
	"project_prakerja2023/controller"

	"github.com/labstack/echo/v4"
)

func RouteInit(e *echo.Echo) {
	e.POST("/books", controller.CreateBook)
}
