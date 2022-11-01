package MajorProject

import (
	"ecommerce/dbc"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Dbhand() Handler {
	return Handler{
		DB: dbc.Dbinit(),
	}
}
