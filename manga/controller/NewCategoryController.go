package controller

//
//import (
//	"Manga-Reader/core"
//	"Manga-Reader/core/models"
//	"encoding/json"
//	"github.com/gin-gonic/gin"
//	"io/ioutil"
//	"net/http"
//	"os"
//	"path/filepath"
//)
//
//var NewCategoryController = newCategoryController{}
//
//type newCategoryController struct{}
//
//// SendNewCategoryForm Sends a form to the client to create a new category
//func (n *newCategoryController) SendNewCategoryForm(c *gin.Context) {
//	var mangaList []string
//	for _, manga := range core.Controller.AllMangas {
//		mangaList = append(mangaList, manga.Title)
//	}
//
//	c.HTML(http.StatusOK, "newCategory.html", gin.H{
//		"mangaList": mangaList,
//	})
//}
//
//// CreateCategory Creates a new category
//func (n *newCategoryController) CreateCategory(c *gin.Context) {
//	category := &models.CategoryApiModel{}
//	err := c.ShouldBindJSON(category)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	databaseCategory := &models.CategoryDatabaseModel{}
//	databaseCategory.CategoryModel.Name = category.Name
//	databaseCategory.Mangas = category.Mangas
//	databaseCategory.Id = len(core.Controller.Categories) + 1
//	databaseCategory.CategoryModel.Order = databaseCategory.Id
//	databaseCategory.Default = false
//
//	categoryJson, err := json.Marshal(databaseCategory)
//	if err != nil {
//		panic(err)
//	}
//
//	var categories []models.CategoryDatabaseModel
//	categoriesJson, err := ioutil.ReadFile(filepath.Join(core.DataDirectory, "categories.json"))
//	if err != nil {
//		create, err := os.Create(filepath.Join(core.DataDirectory, "categories.json"))
//		if err != nil {
//			panic(err)
//		}
//
//		_, err = create.Write(categoryJson)
//		if err != nil {
//			panic(err)
//		}
//
//		err = create.Close()
//		if err != nil {
//			panic(err)
//		}
//	}
//
//	err = json.Unmarshal(categoriesJson, &categories)
//	if err != nil {
//		file, err := json.Marshal(databaseCategory)
//		err = ioutil.WriteFile(filepath.Join(core.DataDirectory, "categories.json"), file, 0644)
//		if err != nil {
//			panic(err)
//		}
//	}
//
//	categories = append(categories, *databaseCategory)
//	core.Controller.CategoryUpdate()
//}
//
//// EditCategories Sends a form to the client to edit a category
//func (n *newCategoryController) EditCategories(c *gin.Context) {
//	var categories []*models.CategoryApiModel
//	for _, category := range core.Controller.Categories {
//		var mangas []string
//		for _, manga := range category.Mangas {
//			mangas = append(mangas, manga.Title)
//		}
//		categories = append(categories, &models.CategoryApiModel{
//			Name:   category.Name,
//			Mangas: mangas,
//		})
//	}
//
//	c.HTML(http.StatusOK, "editCategory.html", gin.H{
//		"category": categories,
//	})
//}
