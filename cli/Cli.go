package cli

import (
	"Manga-Reader/core/models"
	"flag"
	"path/filepath"
)

var Configuration models.ConfigurationModel

func ParseCommandLineArguments() {
	host := flag.String("host", ":1234", "Port to run the server on")
	dataDirectory := flag.String("dir", filepath.Join("d:", "NropVox-Manga"), "Data directory")

	flag.Parse()

	Configuration.Host = *host
	Configuration.DataDirectory = *dataDirectory
}
