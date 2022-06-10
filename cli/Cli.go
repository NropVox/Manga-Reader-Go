package cli

import (
	"Manga-Reader/core/models"
	"flag"
	"path/filepath"
)

//TODO: Add cli commands
var Configuration = models.ConfigurationModel{}

func ParseCommandLineArguments() {
	host := flag.String("host", "1234", "Port to run the server on")
	dataDirectory := flag.String("dir", "D:\\NropVox-Manga", "Data directory")

	flag.Parse()

	Configuration.Host = *host
	Configuration.DataDirectory = *dataDirectory
	Configuration.MangaDBDirectory = filepath.Join(*dataDirectory, "mangas.json")
	Configuration.CategoryDBDirectory = filepath.Join(*dataDirectory, "categories.json")
	Configuration.TemplatesDirectory = filepath.Join(*dataDirectory, "web", "templates")
	Configuration.MangaDBDirectory = filepath.Join(*dataDirectory, "web", "static")
}
