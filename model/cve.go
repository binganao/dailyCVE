package model

import "github.com/jinzhu/gorm"

type CVE struct {
	gorm.Model
	Name        string
	Url         string
	Description string
	Date        string
}
