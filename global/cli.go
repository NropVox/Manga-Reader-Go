package global

import (
	"Manga-Reader/core/models"
	"flag"
	"path/filepath"
)

var CLIArgs models.CommandLineArgsModel

func ParseCommandLineArguments() {
	dataDirectory := flag.String("dir", filepath.Join("d:", "NropVox-Manga"), "Data directory")

	flag.Parse()

	CLIArgs.DataDirectory = *dataDirectory
	loadArgsFromCli()
}
