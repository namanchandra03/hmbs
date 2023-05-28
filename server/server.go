package server

import (
	"github.com/gin-gonic/gin"
	"hms/handlers"
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
	complainAPI := r.Group("/complain")
	{
		complainAPI.POST("", handlers.AddComplain)
		complainAPI.PUT("/complainID", handlers.EditComplain)
		complainAPI.DELETE("/complainID", handlers.DeleteComplain)
		complainAPI.GET("", handlers.GetComplain)
	}

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
