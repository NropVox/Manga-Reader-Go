package core

import (
	"Manga-Reader/cli"
)

var DataDirectory = cli.Configuration.DataDirectory

var LocalDirectory = cli.Configuration.LocalDirectory
var CategoryDBDirectory = cli.Configuration.CategoryDBDirectory
var MangaDBDirectory = cli.Configuration.MangaDBDirectory
var TemplatesDirectory = cli.Configuration.TemplatesDirectory
var StaticDirectory = cli.Configuration.StaticDirectory

//var ThumbnailsDirectory = filepath.Join(DataDirectory, "thumbnails")
