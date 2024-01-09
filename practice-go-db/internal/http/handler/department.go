package handler

import (
	"context"
	"go-db/internal/core/dto"
	core_errors "go-db/internal/core/errors"
	"go-db/internal/http/response"
	"go-db/internal/http/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type DepartmentHandler struct {
	departmentService service.DepartmentService
	timeout           time.Duration
}

func NewDepartmentHandler(service service.DepartmentService, timeout time.Duration) *DepartmentHandler {
	return &DepartmentHandler{
		departmentService: service,
		timeout:           timeout,
	}
}

func (h *DepartmentHandler) AddDepartment(c *gin.Context) {
	request := &dto.AddDepartmentDTO{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.BadResponse{Error: err.Error()})
		return
	}
	ctx := context.TODO()
	// defer cansel()

	result, err := h.departmentService.AddDepartment(ctx, request)
	if err != nil {

		c.JSON(http.StatusBadRequest, response.NewBadResponse(err))
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *DepartmentHandler) GetDepartment(c *gin.Context) {
	request := &dto.IdDerartmentDTO{}
	limitStr := c.Query("id")
	if limitStr != "" {
		idVal, err := strconv.Atoi(limitStr)
		if err != nil || idVal <= 0 {
			c.JSON(http.StatusBadRequest, response.NewBadResponse(core_errors.ErrorInvalidId))
			return
		}
		request.Id = int64(idVal)
	}

	ctx, cansel := context.WithTimeout(context.Background(), h.timeout)
	defer cansel()

	result, err := h.departmentService.GetDepartmentById(ctx, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewBadResponse(err))
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *DepartmentHandler) RenameDepartment(c *gin.Context) {
	request := &dto.RenameDerartmentDTO{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.NewBadResponse(err))
		return
	}

	ctx, cansel := context.WithTimeout(context.Background(), h.timeout)
	defer cansel()

	result, err := h.departmentService.RenameDepartment(ctx, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewBadResponse(err))
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *DepartmentHandler) GetDepartments(c *gin.Context) {
	request := &dto.LimitOffsetDTO{}
	limitStr := c.Query("limit")
	if limitStr != "" {
		limitVal, err := strconv.Atoi(limitStr)
		if err != nil || limitVal <= 0 {
			c.JSON(http.StatusBadRequest, response.NewBadResponse(core_errors.ErrorInvalidLinit))
			return
		}
		request.Limit = uint(limitVal)
	}
	offsetStr := c.Query("offset")
	if offsetStr != "" {
		offsetVal, err := strconv.Atoi(offsetStr)
		if err != nil || offsetVal <= 0 {
			c.JSON(http.StatusBadRequest, response.NewBadResponse(core_errors.ErrorInvalidOffset))
			return
		}
		request.Offset = uint(offsetVal)
	}

	ctx, cansel := context.WithTimeout(context.Background(), h.timeout)
	defer cansel()

	result, err := h.departmentService.GetDepartments(ctx, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewBadResponse(err))
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *DepartmentHandler) DeleteDepartment(c *gin.Context) {
	request := &dto.IdDerartmentDTO{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.NewBadResponse(err))
		return
	}

	ctx, cansel := context.WithTimeout(context.Background(), h.timeout)
	defer cansel()

	err := h.departmentService.DeleteDepartment(ctx, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewBadResponse(err))
		return
	}

	c.JSON(http.StatusOK, response.NewGoogResponse())
}
