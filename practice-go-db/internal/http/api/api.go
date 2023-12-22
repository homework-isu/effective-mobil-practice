package api

import (
	"github.com/gin-gonic/gin"
	"go-db/internal/http/hendler"
)

type api struct {
	server *gin.Engine
	departmentHendler *hendler.DepartmentHendler
	port string
}

func NewApi(port string, dh *hendler.DepartmentHendler) *api {
	server := gin.Default()

	api := &api{
		server: server,
		port: port,
		departmentHendler: dh,
	}

	api.bind()

	return api
}

func (api *api) bind() {
	api.server.GET("/department", api.departmentHendler.GetDepartment)
	api.server.GET("/departments", api.departmentHendler.GetDepartments)
	api.server.POST("/department", api.departmentHendler.AddDepartment)
	api.server.DELETE("/department", api.departmentHendler.DeleteDepartment)
	api.server.PUT("/department", api.departmentHendler.RenameDepartment)
}

func (api api) Run() error {
	return api.server.Run(":" + api.port)
}