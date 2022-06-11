package core

import (
	"Manga-Reader/core/models"
	. "Manga-Reader/global"
	"Manga-Reader/utils"
	"archive/zip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var err error

var Controller = &MainController{}

func NewMainController() *MainController {
	controller := &MainController{}

	if err = os.MkdirAll(LocalDirectory, 0755); err != nil {
		panic(err)
	}
	if err = os.MkdirAll(filepath.Join(DataDirectory, "thumbnails"), 0755); err != nil {
		panic(err)
	}

	detailsJson, err := ioutil.ReadFile(MangaDBDirectory)
	if err == nil {
		err = json.Unmarshal(detailsJson, &controller.AllMangas)
		if err != nil {
			panic(err)
		}
	} else {
		controller.Update(models.UpdateConfigModel{})
	}

	controller.Categories = []models.CategoryMangaModel{
		{
			CategoryModel: &models.CategoryModel{
				Id:      1,
				Order:   1,
				Name:    "All Mangas",
				Default: true,
			},
		},
	}

	var mangas []*models.MangaDataModel
	for _, manga := range controller.AllMangas {
		mangas = append(mangas, manga.MangaDataModel)
	}
	controller.Categories[0].Mangas = mangas

	controller.UpdateCategories()

	Controller = controller

	return controller
}

type MainController struct {
	AllMangas []models.MangaModel

	Categories []models.CategoryMangaModel
}

// FindCategoryWithId returns the category with the given id
func (m *MainController) FindCategoryWithId(id int) *models.CategoryMangaModel {
	for _, category := range m.Categories {
		if category.Id == id {
			return &category
		}
	}
	return nil
}

// FindMangaWithId returns the manga with the given id
func (m *MainController) FindMangaWithId(id int) *models.MangaModel {
	for _, manga := range m.AllMangas {
		if manga.Id == id {
			return &manga
		}
	}
	return nil
}

// FindMangaWithName returns the manga with the given name
func (m *MainController) FindMangaWithName(name string) *models.MangaModel {
	for _, manga := range m.AllMangas {
		if manga.Title == name {
			return &manga
		}
	}
	return nil
}

// FindChapterWithId returns the chapter with the given manga ID and chapter ID
func (m *MainController) FindChapterWithId(mangaId int, chapterId int) *models.ChapterModel {
	manga := m.FindMangaWithId(mangaId)
	if manga == nil {
		return nil
	}

	for _, chapter := range manga.Chapters {
		if chapter.Index == chapterId {
			return &chapter
		}
	}

	return nil
}

// FindPageWithId returns the page with the given manga ID, chapter ID and page ID
func (m *MainController) FindPageWithId(mangaId int, chapterId int, pageId int) *models.PageModel {
	chapter := m.FindChapterWithId(mangaId, chapterId)
	if chapter.Archived {
		return &models.PageModel{
			Name: "Archived",
		}
	} else {
		for _, page := range chapter.Pages {
			if page.Id == pageId {
				return &page
			}
		}
	}
	return nil
}

// UpdateCategories updates the categories
func (m *MainController) UpdateCategories() {
	var categoriesRaw []*models.CategoryDatabaseModel
	categoriesJson, err := ioutil.ReadFile(CategoryDBDirectory)
	if err == nil {
		err = json.Unmarshal(categoriesJson, &categoriesRaw)
		if err != nil {
			err := ioutil.WriteFile(CategoryDBDirectory, []byte(""), 0644)
			if err != nil {
				panic(err)
			}
		}
	} else {
		return
	}

	var categories []models.CategoryMangaModel
	for _, category := range categoriesRaw {
		categoryData := models.CategoryMangaModel{
			CategoryModel: &models.CategoryModel{},
		}
		for _, manga := range category.Mangas {
			searchedManga := m.FindMangaWithName(manga)
			if searchedManga != nil {
				categoryData.Name = category.Name
				categoryData.Id = category.Id
				categoryData.Order = category.Order
				categoryData.Default = category.Default
				categoryData.Mangas = append(categoryData.Mangas, searchedManga.MangaDataModel)
			}
		}
		categories = append(categories, categoryData)
	}

	m.Categories = append(m.Categories, categories...)
}

// Update updates the list of mangas from the disk
func (m *MainController) Update(conf models.UpdateConfigModel) {
	var mangaList []models.MangaModel

	mangas, err := ioutil.ReadDir(LocalDirectory)
	if err != nil {
		panic(err)
	}

	for mangaId, manga := range mangas {
		if utils.Contains(m.AllMangas, manga.Name()) && conf.New {
			continue
		}
		fileMode := manga.Mode().String()[0]
		if fileMode == 'd' || fileMode == 'L' {
			mangaData := models.MangaModel{
				MangaDataModel: &models.MangaDataModel{
					Id:           mangaId + 1,
					Title:        manga.Name(),
					SourceId:     "0",
					Genre:        []string{},
					Url:          manga.Name(),
					ThumbnailUrl: fmt.Sprintf("/api/v1/manga/%d/thumbnail", mangaId+1),
				},
			}

			contents, err := ioutil.ReadDir(filepath.Join(LocalDirectory, manga.Name()))
			if err != nil {
				panic(err)
			}

			utils.SortFiles(contents)

			for chapterId, content := range contents {
				ext := filepath.Ext(content.Name())
				if content.IsDir() || ext == ".cbz" {
					chapter := models.ChapterModel{
						ChapterDataModel: &models.ChapterDataModel{
							Url:        filepath.Join(manga.Name(), content.Name()),
							Name:       strings.ReplaceAll(content.Name(), ext, ""),
							UploadDate: content.ModTime().UnixMilli(),
							Index:      chapterId + 1,
							MangaId:    mangaId + 1,
							Scanlator:  "",
						},
						Archived: false,
					}

					var pageCount int

					if content.IsDir() {
						pages, err := ioutil.ReadDir(filepath.Join(LocalDirectory, manga.Name(), content.Name()))
						if err != nil {
							panic(err)
						}

						pageCount = len(pages)

						var pageId int
						for _, page := range pages {
							var pageItem models.PageModel
							if page.Name() != ".nomedia" {
								pageItem.Name = page.Name()
								pageItem.Id = pageId
								chapter.Pages = append(chapter.Pages, pageItem)
								pageId++
							} else {
								pageCount--
							}

						}
					} else if ext == ".cbz" || ext == ".cbr" || ext == ".zip" {
						chapter.Archived = true

						if ext == ".cbz" || ext == ".zip" {
							reader, err := zip.OpenReader(filepath.Join(LocalDirectory, chapter.Url))
							if err != nil {
								panic(err)
							}

							pageCount = len(reader.File)
							err = reader.Close()
							if err != nil {
								panic(err)
							}
						} else {
							panic("CBR files are not supported yet")
						}

					}

					chapter.PageCount = pageCount
					mangaData.Chapters = append(mangaData.Chapters, chapter)
				} else if content.Name() == "details.json" {
					detailsJson, err := ioutil.ReadFile(filepath.Join(LocalDirectory, manga.Name(), content.Name()))
					if err != nil {
						panic(err)
					}

					details := &models.DetailsJson{}
					err = json.Unmarshal(detailsJson, details)
					if err != nil {
						panic(err)
					}

					var status string
					switch details.StatusCode {
					case 1:
						status = "ONGOING"
					case 2:
						status = "COMPLETED"
					case 3:
						status = "LICENCED"
					default:
						status = "UNKNOWN"
					}

					mangaData.MangaDataModel.Title = details.Title
					mangaData.MangaDataModel.Genre = details.Genre
					mangaData.MangaDataModel.Author = details.Author
					mangaData.MangaDataModel.Artist = details.Artist
					mangaData.MangaDataModel.Status = status
					mangaData.MangaDataModel.Description = details.Description

				} else if strings.Contains(content.Name(), "cover") {
					mangaData.CoverPath = filepath.Join(LocalDirectory, manga.Name(), content.Name())
				}
			}

			mangaList = append(mangaList, mangaData)
		}
	}

	m.AllMangas = append(m.AllMangas, mangaList...)

	// Save mangaList to file
	mangaListJson, err := json.Marshal(m.AllMangas)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(MangaDBDirectory, mangaListJson, 0644)
	if err != nil {
		panic(err)
	}
}

func (m *MainController) UpdateNew() {

}
