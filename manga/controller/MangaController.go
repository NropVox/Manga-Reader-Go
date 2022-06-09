package controller

import (
	"Manga-Reader/core"
	"Manga-Reader/core/models"
	"Manga-Reader/utils"
	"archive/zip"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

var MangaController = mangaController{}

type mangaController struct{}

//func (m *mangaController) SendMangasList(c *gin.Context) {
//
//}

func (m *mangaController) SendMangaDetails(c *gin.Context) {
	manga := c.MustGet("manga").(*models.MangaModel)
	c.JSON(http.StatusOK, manga.MangaDataModel)
}

// SendCoverPhoto sends the cover photo of the manga
func (m *mangaController) SendCoverPhoto(c *gin.Context) {
	manga := c.MustGet("manga").(*models.MangaModel)
	cover, err := ioutil.ReadFile(manga.CoverPath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cover not found"})
		return
	}

	c.Data(http.StatusOK, "image/jpeg", cover)
}

func (m *mangaController) SendChaptersList(c *gin.Context) {
	manga := c.MustGet("manga").(*models.MangaModel)

	c.JSON(http.StatusOK, utils.Reverse(manga.Chapters))
}

func (m *mangaController) SendChapterDetails(c *gin.Context) {
	chapter := c.MustGet("chapter").(*models.ChapterModel)

	c.JSON(http.StatusOK, chapter.ChapterDataModel)
}

func (m *mangaController) SendPage(c *gin.Context) {
	chapter := c.MustGet("chapter").(*models.ChapterModel)
	page := c.MustGet("page").(*models.PageModel)
	pageId := c.GetInt("pageID")

	var err error
	var pageData []byte
	var contentType string
	var ext string

	chapterPath := filepath.Join(core.LocalDirectory, chapter.Url)

	if page.Name == "Archived" {
		chapterExt := filepath.Ext(chapter.Url)

		if chapterExt == ".zip" || chapterExt == ".cbz" {
			reader, err := zip.OpenReader(chapterPath)
			if err != nil {
				panic(err)
			}

			pageReader := reader.File[pageId]
			pageOpen, err := pageReader.Open()
			defer func(pageOpen io.ReadCloser) {
				err := pageOpen.Close()
				if err != nil {
					panic(err)
				}
			}(pageOpen)

			if err != nil {
				panic(err)
			}

			pageData, err = io.ReadAll(pageOpen)
			if err != nil {
				panic(err)
			}

			ext = filepath.Ext(pageReader.Name)
		} else {
			panic("Archive not supported")
		}
	} else {
		pagePath := filepath.Join(chapterPath, page.Name)
		pageData, err = ioutil.ReadFile(pagePath)
		if err != nil {
			panic(err)
		}

		ext = filepath.Ext(pagePath)
	}

	switch ext {
	case ".jpg":
		contentType = "image/jpeg"
	case ".png":
		contentType = "image/png"
	case ".webp":
		contentType = "image/webp"
	default:
		contentType = "image"
	}

	c.Data(http.StatusOK, contentType, pageData)
}
