package controller

import (
	"Manga-Reader/core"
	"Manga-Reader/core/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var LibraryUpdateController = libraryUpdateController{}

type libraryUpdateController struct{}

func (l *libraryUpdateController) UpdateAll(c *gin.Context) {
	core.Controller.Update(models.UpdateConfigModel{})
	c.JSON(http.StatusOK, gin.H{"message": "Library updated"})
}

func (l *libraryUpdateController) UpdateNew(c *gin.Context) {
	core.Controller.Update(models.UpdateConfigModel{New: true})
	c.JSON(http.StatusOK, gin.H{"message": "Library updated"})
}
