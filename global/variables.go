package global

import (
	"Manga-Reader/core/models"
	json "github.com/json-iterator/go"
	"io/ioutil"
	"path/filepath"
)

var DataDirectory string

var LocalDirectory string
var CategoryDBDirectory string
var MangaDBDirectory string
var TemplatesDirectory string
var StaticDirectory string

var Address string
var Port int

func loadArgsFromCli() {
	DataDirectory = CLIArgs.DataDirectory

	go loadConfigFile()

	LocalDirectory = filepath.Join(DataDirectory, "local")
	CategoryDBDirectory = filepath.Join(DataDirectory, "categories.json")
	MangaDBDirectory = filepath.Join(DataDirectory, "mangas.json")
	TemplatesDirectory = filepath.Join(DataDirectory, "web", "templates")
	StaticDirectory = filepath.Join(DataDirectory, "web", "static")
}

func loadConfigFile() {
	configFile, err := ioutil.ReadFile(filepath.Join(DataDirectory, "config.json"))
	if err != nil {
		panic(err)
	}

	config := &models.ConfigurationModel{}

	err = json.Unmarshal(configFile, config)
	if err != nil {
		panic(err)
	}

	Address = config.Address
	Port = config.Port
}
