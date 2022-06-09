package main

import (
	"Manga-Reader/core"
	"Manga-Reader/server"
)

func main() {
	err := server.NewMangaServer(core.Controller).Router.Run(":1234")
	if err != nil {
		panic(err)
	}
}
