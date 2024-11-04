package router

import (
	"blog/config"
	"blog/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewServer() *http.Server {
	r := gin.Default()
	config.SetCORS(r)
	config.InitSession(r)
	InitRouter(r)
	// r = r.Delims("left", "right")
	s := &http.Server{
		Addr:    "0.0.0.0:8088",
		Handler: r,
	}
	return s

}

var ctr = controller.New()
