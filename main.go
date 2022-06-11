package main

import (
	"Manga-Reader/cli"
	"Manga-Reader/core"
	"Manga-Reader/server"
)

func main() {
	cli.ParseCommandLineArguments()
	core.NewMainController()

	err := server.NewMangaServer(core.Controller).Router.Run(cli.Configuration.Host)
	if err != nil {
		panic(err)
	}
}
