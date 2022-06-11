package server

import (
	"Manga-Reader/core"
	. "Manga-Reader/global"
	"Manga-Reader/manga"
	. "Manga-Reader/manga/controller"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

// NewMangaServer Creates new manga server
func NewMangaServer(controller *core.MainController) *Server {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	router.LoadHTMLGlob(filepath.Join(TemplatesDirectory, "*.html"))
	router.Static("/static", StaticDirectory)

	mangaAPI.Router(router.Group("/api/v1"))

	newCategory := router.Group("/new-category")
	{
		newCategory.GET("/", NewCategoryController.SendNewCategoryForm)
		newCategory.POST("/", NewCategoryController.CreateCategory)
	}

	println("Server Started...")

	return &Server{
		Router:     router,
		Controller: controller,
	}
}

type Server struct {
	Router     *gin.Engine
	Controller *core.MainController
}
