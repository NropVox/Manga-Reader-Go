package global

import "Manga-Reader/core/models"

var defaultConfigFile = models.ConfigurationModel{
	Address: "0.0.0.0",
	Port:    1234,
}
