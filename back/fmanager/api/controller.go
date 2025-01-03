package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/icehardstone/fmanager/serror"
	"github.com/icehardstone/fmanager/store"
)

type Controller struct {
	Store *store.Store
}

func NewController() *Controller {
	return &Controller{
		Store: store.NewStore(),
	}
}
func (c *Controller) NotFound(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, &serror.APIError{
		ErrorCode:    http.StatusBadRequest,
		ErrorMessage: "Not found data",
	})
}

func (c *Controller) BadRequest(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, &serror.APIError{
		ErrorCode:    http.StatusBadRequest,
		ErrorMessage: "Bad Request",
	})
}
