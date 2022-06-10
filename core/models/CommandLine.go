package models

type ConfigurationModel struct {
	Host string

	DataDirectory       string `json:"data_directory"`
	LocalDirectory      string `json:"local_directory"`
	CategoryDBDirectory string `json:"category_db_directory"`
	MangaDBDirectory    string `json:"manga_db_directory"`
	TemplatesDirectory  string `json:"templates_directory"`
	StaticDirectory     string `json:"static_directory"`
}
