package server

import (
	"Manga-Reader/core"
	"Manga-Reader/manga"
	"github.com/gin-gonic/gin"
)

// NewMangaServer Creates new manga server
func NewMangaServer(controller *core.MainController) *Server {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "static")

	mangaAPI.Router(router.Group("/api/v1"))

	return &Server{
		Router:     router,
		Controller: controller,
	}
}

type Server struct {
	Router     *gin.Engine
	Controller *core.MainController
}
