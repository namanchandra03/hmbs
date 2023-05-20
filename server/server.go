package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Server struct {
	*gin.Engine
	server *http.Server
}

const (
	ReadTimeOut       = 5 * time.Minute
	ReadHeaderTimeOut = 30 * time.Second
	WriteTimeout      = 5 * time.Minute
)

func SetUpRoutes() *Server {
	r := gin.Default()
	test := r.Group("/home")
	test.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"response": "hello",
		})
	})

	return &Server{
		Engine: r,
	}
}

func (svc *Server) Start(port string) error {
	svc.server = &http.Server{
		Addr:              port,
		Handler:           svc.Engine,
		ReadTimeout:       ReadTimeOut,
		ReadHeaderTimeout: ReadHeaderTimeOut,
		WriteTimeout:      WriteTimeout,
	}
	return svc.server.ListenAndServe()
}
