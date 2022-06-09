package controller

import (
	"Manga-Reader/core"
	"Manga-Reader/core/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

var CategoryController = categoryController{}

type categoryController struct {
}

func (c *categoryController) SendCategoriesList(g *gin.Context) {
	var categories []*models.CategoryModel
	for _, categoryManga := range core.Controller.Categories {
		categories = append(categories, &models.CategoryModel{
			Id:      categoryManga.Id,
			Name:    categoryManga.Name,
			Default: categoryManga.Default,
		})
	}

	g.JSON(http.StatusOK, categories)
}

func (c *categoryController) SendMangasList(g *gin.Context) {
	categoryId := g.GetInt("categoryID")
	category := core.Controller.FindCategoryWithId(categoryId)
	if category == nil {
		g.JSON(http.StatusBadRequest, gin.H{"message": "Category not found"})
		return
	}

	var mangas []*models.MangaDataModel
	for _, manga := range category.Mangas {
		mangas = append(mangas, manga)
	}

	jsonMangas, err := json.Marshal(mangas)
	if err != nil {
		panic(err)
	}

	g.Data(http.StatusOK, "application/json", jsonMangas)
}
