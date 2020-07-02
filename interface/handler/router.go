package handler

import (
	"github.com/labstack/echo"
)

// InitRouting routesの初期化
func InitRouting(e *echo.Echo, taskHandler TaskHandler) {

	e.POST("/api/task", taskHandler.Post())
	e.GET("/api/tasks", taskHandler.GetAll())
	e.GET("/api/task/:id", taskHandler.Get())
	e.PUT("/api/task/:id", taskHandler.Put())
	e.DELETE("/api/task/:id", taskHandler.Delete())

}
