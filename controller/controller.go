package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Admin
	FileManagement
}

func (con Controller) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}

func New() *Controller {
	Controller := &Controller{}
	return Controller
}
