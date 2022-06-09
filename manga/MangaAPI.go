package mangaAPI

import (
	"Manga-Reader/core"
	. "Manga-Reader/manga/controller"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Router handles routes in "/manga"
func Router(r *gin.RouterGroup) {
	manga := r.Group("/manga")
	{
		manga.Use(idValidator("mangaID"))
		manga.GET("/:mangaID", MangaController.SendMangaDetails)
		manga.GET("/:mangaID/", MangaController.SendMangaDetails)

		manga.GET("/:mangaID/thumbnail", MangaController.SendCoverPhoto)
		manga.GET("/:mangaID/chapters", MangaController.SendChaptersList)
		manga.GET("/:mangaID/chapter/:chapterID", idValidator("chapterID"), MangaController.SendChapterDetails)
		manga.GET("/:mangaID/chapter/:chapterID/page/:pageID", idValidator("chapterID"), idValidator("pageID"), MangaController.SendPage)
	}

	category := r.Group("/category")
	{
		category.GET("/", CategoryController.SendCategoriesList)
		category.GET("/:categoryID", idValidator("categoryID"), CategoryController.SendMangasList)
	}

	update := r.Group("/update")
	{
		update.GET("/all", LibraryUpdateController.UpdateAll)
		update.GET("/new", LibraryUpdateController.UpdateNew)
		update.GET("/manga/:mangaID", LibraryUpdateController.UpdateAll)
	}
}

func idValidator(id string) gin.HandlerFunc {
	return func(c *gin.Context) {
		paramId, err := strconv.Atoi(c.Param(id))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
			c.Abort()
			return
		}
		c.Set(id, paramId)

		if id == "mangaID" {
			manga := core.Controller.FindMangaWithId(paramId)
			if manga == nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": "Manga not found"})
				c.Abort()
				return
			}
			c.Set("manga", manga)
		} else if id == "categoryID" {
			category := core.Controller.FindCategoryWithId(paramId)
			if category == nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": "Category not found"})
				c.Abort()
				return
			}
			c.Set("category", category)
		} else if id == "chapterID" {
			mangaId := c.MustGet("mangaID").(int)
			chapter := core.Controller.FindChapterWithId(mangaId, paramId)
			if chapter == nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": "Chapter not found"})
				c.Abort()
				return
			}
			c.Set("chapter", chapter)
		} else if id == "pageID" {
			mangaId := c.MustGet("mangaID").(int)
			chapterId := c.MustGet("chapterID").(int)
			page := core.Controller.FindPageWithId(mangaId, chapterId, paramId)
			if page == nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": "Page not found"})
				c.Abort()
				return
			}
			c.Set("page", page)
		}

		c.Next()
	}
}
