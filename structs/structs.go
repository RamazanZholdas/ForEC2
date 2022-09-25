package structs

import (
	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	Name     string `json:"name"`
	Duration string `json:"duration"`
	Genre    string `json:"genre"`
	Artist   string `json:"artist"`
}

type SimpleSong struct {
	Name     string `json:"name"`
	Duration string `json:"duration"`
	Genre    string `json:"genre"`
	Artist   string `json:"artist"`
}
