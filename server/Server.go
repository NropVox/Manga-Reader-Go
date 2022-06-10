package server

import (
	"Manga-Reader/core"
	"Manga-Reader/manga"
	. "Manga-Reader/manga/controller"
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

	newCategory := router.Group("/new-category")
	{
		newCategory.GET("/", NewCategoryController.SendNewCategoryForm)
		newCategory.POST("/", NewCategoryController.CreateCategory)
	}

	return &Server{
		Router:     router,
		Controller: controller,
	}
}

type Server struct {
	Router     *gin.Engine
	Controller *core.MainController
}
