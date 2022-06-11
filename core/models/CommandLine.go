package models

type ConfigurationModel struct {
	Host string

	DataDirectory string `json:"data_directory"`
}
