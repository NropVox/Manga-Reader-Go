package main

import (
	"Manga-Reader/core"
	. "Manga-Reader/global"
	"Manga-Reader/server"
	"fmt"
)

func main() {
	ParseCommandLineArguments()
	core.NewMainController()

	host := fmt.Sprintf("%s:%d", Address, Port)

	err := server.NewMangaServer(core.Controller).Router.Run(host)
	if err != nil {
		panic(err)
	}
}
