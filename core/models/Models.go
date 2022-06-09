package models

type CategoryMangaModel struct {
	//Id      int
	//Name    string
	//Default bool
	*CategoryModel

	Mangas []*MangaDataModel
}

type MangaModel struct {
	*MangaDataModel

	CoverPath string
	Chapters  []ChapterModel
}

type ChapterModel struct {
	*ChapterDataModel

	Archived bool
	Pages    []PageModel
}

type PageModel struct {
	Name string
	Id   int
}

type CategoryModel struct {
	Id      int    `json:"id"`
	Order   int    `json:"order"`
	Name    string `json:"name"`
	Default bool   `json:"default"`
}

type MangaDataModel struct {
	Id           int      `json:"id"`
	Url          string   `json:"url"`
	SourceId     string   `json:"sourceId"`
	ThumbnailUrl string   `json:"thumbnailUrl"`
	Title        string   `json:"title"`
	Artist       string   `json:"artist"`
	Author       string   `json:"author"`
	Description  string   `json:"description"`
	Genre        []string `json:"genre"`
	Status       string   `json:"status"`
}

type ChapterDataModel struct {
	Url           string  `json:"url"`
	Name          string  `json:"name"`
	UploadDate    int64   `json:"uploadDate"`
	Index         int     `json:"index"`
	ChapterNumber float32 `json:"chapterNumber"`
	MangaId       int     `json:"mangaId"`
	Scanlator     string  `json:"scanlator"`
	PageCount     int     `json:"pageCount"`
	Read          bool    `json:"read"`
	Bookmarked    bool    `json:"bookmarked"`
	LastPageRead  int     `json:"lastPageRead"`
	LastReadAt    int     `json:"lastReadAt"`
	Downloaded    bool    `json:"downloaded"`
}

type DetailsJson struct {
	Title       string   `json:"title"`
	Artist      string   `json:"artist"`
	Author      string   `json:"author"`
	Description string   `json:"description"`
	Genre       []string `json:"genre"`
	StatusCode  int      `json:"status"`
}

type UpdateConfigModel struct {
	New bool
}
