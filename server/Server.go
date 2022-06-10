package server

import (
	"Manga-Reader/core"
	"Manga-Reader/manga"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

// NewMangaServer Creates new manga server
func NewMangaServer(controller *core.MainController) *Server {
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)

	router.LoadHTMLGlob(filepath.Join(core.TemplatesDirectory, "*.html"))
	router.Static("/static", core.StaticDirectory)

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
