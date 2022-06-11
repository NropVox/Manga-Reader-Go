package core

import (
	"Manga-Reader/cli"
	"path/filepath"
)

var DataDirectory string

var LocalDirectory string
var CategoryDBDirectory string
var MangaDBDirectory string
var TemplatesDirectory string
var StaticDirectory string

func loadConfigsFromCli() {
	DataDirectory = cli.Configuration.DataDirectory

	LocalDirectory = filepath.Join(DataDirectory, "local")
	CategoryDBDirectory = filepath.Join(DataDirectory, "categories.json")
	MangaDBDirectory = filepath.Join(DataDirectory, "mangas.json")
	TemplatesDirectory = filepath.Join(DataDirectory, "web", "templates")
	StaticDirectory = filepath.Join(DataDirectory, "web", "static")
}

//var ThumbnailsDirectory = filepath.Join(DataDirectory, "thumbnails")
