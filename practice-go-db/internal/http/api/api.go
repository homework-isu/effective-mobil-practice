package api

import (
	"github.com/gin-gonic/gin"
	"go-db/internal/http/handler"
)

type api struct {
	server *gin.Engine
	departmentHandler *handler.DepartmentHandler
	port string
}

func NewApi(port string, dh *handler.DepartmentHandler) *api {
	server := gin.Default()

	api := &api{
		server: server,
		port: port,
		departmentHandler: dh,
	}

	api.bind()

	return api
}

func (api *api) bind() {
	api.server.GET("/department", api.departmentHandler.GetDepartment)
	api.server.GET("/departments", api.departmentHandler.GetDepartments)
	api.server.POST("/department", api.departmentHandler.AddDepartment)
	api.server.DELETE("/department", api.departmentHandler.DeleteDepartment)
	api.server.PUT("/department", api.departmentHandler.RenameDepartment)
}

func (api api) Run() error {
	return api.server.Run(":" + api.port)
}