package controller

import (
	"Manga-Reader/core"
	"Manga-Reader/core/models"
	. "Manga-Reader/global"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

var NewCategoryController = newCategoryController{}

type newCategoryController struct{}

// SendNewCategoryForm Sends a form to the client to create a new category
func (n *newCategoryController) SendNewCategoryForm(c *gin.Context) {
	var mangaList []string
	for _, manga := range core.Controller.AllMangas {
		mangaList = append(mangaList, manga.Title)
	}

	c.HTML(http.StatusOK, "newCategory.html", gin.H{
		"mangaList": mangaList,
	})
}

// CreateCategory Creates a new category
func (n *newCategoryController) CreateCategory(c *gin.Context) {
	category := &models.CategoryApiModel{}
	err := c.ShouldBindJSON(category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	databaseCategory := &models.CategoryDatabaseModel{}
	databaseCategory.CategoryModel.Name = category.ApiName
	databaseCategory.Mangas = category.Mangas
	databaseCategory.Id = len(core.Controller.Categories) + 1
	databaseCategory.CategoryModel.Order = databaseCategory.Id
	databaseCategory.Default = false

	categoriesJson, err := ioutil.ReadFile(CategoryDBDirectory)
	if err != nil {
		create, err := os.Create(CategoryDBDirectory)
		if err != nil {
			panic(err)
		}
		defer func(create *os.File) {
			err := create.Close()
			if err != nil {
				panic(err)
			}
		}(create)

		categoryJson, err := json.Marshal([]*models.CategoryDatabaseModel{databaseCategory})
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(CategoryDBDirectory, categoryJson, 0644)
		if err != nil {
			panic(err)
		}
	} else {
		var categories []models.CategoryDatabaseModel
		err = json.Unmarshal(categoriesJson, &categories)
		if err != nil {
			newCategoryDatabase := []*models.CategoryDatabaseModel{
				databaseCategory,
			}

			file, err := json.Marshal(newCategoryDatabase)
			err = ioutil.WriteFile(CategoryDBDirectory, file, 0644)
			if err != nil {
				panic(err)
			}
		} else {
			categories = append(categories, *databaseCategory)

			file, err := json.Marshal(categories)
			err = ioutil.WriteFile(CategoryDBDirectory, file, 0644)
			if err != nil {
				panic(err)
			}
		}
	}

	core.Controller.UpdateCategories()
}

// EditCategories Sends a form to the client to edit a category
func (n *newCategoryController) EditCategories(c *gin.Context) {
	var categories []*models.CategoryApiModel
	for _, category := range core.Controller.Categories {
		var mangas []string
		for _, manga := range category.Mangas {
			mangas = append(mangas, manga.Title)
		}
		categories = append(categories, &models.CategoryApiModel{
			ApiName: category.Name,
			Mangas:  mangas,
		})
	}

	c.HTML(http.StatusOK, "editCategory.html", gin.H{
		"category": categories,
	})
}
