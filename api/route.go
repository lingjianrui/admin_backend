package api

import (
	"backend/middleware"
	"net/http"
)

func (s *Server) initializeRoutes() {
	s.Router.Use(middleware.Cors())
	s.Router.Use(middleware.LoggerToFile())

	s.Router.StaticFS("/upload", http.Dir("upload"))
	s.Router.GET("/ping", s.Ping)
	s.Router.POST("/login", s.Login)
	s.Router.POST("/post", s.ImagePost)
	v1 := s.Router.Group("/api/v1")
	v1.Use(middleware.JWT())
	{
		v1.GET("/vendor", s.GetVendorList)
		v1.POST("/vendor", s.CreateVendor)
		v1.DELETE("/vendor", s.DeleteVendor)
		v1.GET("/user", s.GetUserList)
		v1.GET("/user/info", s.GetUserInfo)
		v1.POST("/user", s.CreateUser)

	}
}
