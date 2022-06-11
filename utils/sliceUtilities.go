package utils

import (
	"Manga-Reader/core/models"
	"io/fs"
	"regexp"
	"sort"
	"strconv"
)

func SortFiles(s []fs.FileInfo) {
	sort.SliceStable(s, func(i, j int) bool {
		re, err := regexp.Compile("(\\d+(\\.\\d+)?)|(\\.\\d+)")
		if err != nil {
			panic(err)
		}
		i1 := re.FindString(s[i].Name())
		i2 := re.FindString(s[j].Name())

		if i1 == "" || i2 == "" {
			return false
		}

		//convert to int
		i1Int, err := strconv.ParseFloat(i1, 32)
		if err != nil {
			panic(err)
		}

		i2Int, err := strconv.ParseFloat(i2, 32)
		if err != nil {
			panic(err)
		}
		//firstNum, err
		return i1Int < i2Int
	})
}

// ReverseChapters returns the reversed copy of the given chapters
//func ReverseChapters(chapters []models.ChapterModel) []models.ChapterModel {
//
//}

func Reverse[S ~[]E, E any](s S) S {
	newSlice := make([]E, len(s))
	copy(newSlice, s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		newSlice[i], newSlice[j] = newSlice[j], newSlice[i]
	}
	return newSlice
}

func Contains(s []models.MangaModel, e string) bool {
	for _, a := range s {
		if a.Title == e {
			return true
		}
	}
	return false
}
