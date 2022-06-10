package core

import "path/filepath"

var DataDirectory = filepath.Join("d:", "NropVox-Manga")

var LocalDirectory = filepath.Join(DataDirectory, "local")
var CategoryDBDirectory = filepath.Join(DataDirectory, "categories.json")
var MangaDBDirectory = filepath.Join(DataDirectory, "mangas.json")
var TemplatesDirectory = filepath.Join(DataDirectory, "web", "templates")
var StaticDirectory = filepath.Join(DataDirectory, "web", "static")

//var ThumbnailsDirectory = filepath.Join(DataDirectory, "thumbnails")
