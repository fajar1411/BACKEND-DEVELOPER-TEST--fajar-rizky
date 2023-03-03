package data

import (
	"test/fitur/national"

	"gorm.io/gorm"
)

type Nationality struct {
	gorm.Model
	Nationality_Name string `gorm:"type:varchar(50);not null"`
	Nationality_Code string `gorm:"type:char(2);not null"`
}

func FromEntities(data national.NationalityEntities) Nationality {
	return Nationality{
		Model:            gorm.Model{ID: data.ID},
		Nationality_Name: data.Nationality_Name,
		Nationality_Code: data.Nationality_Code,
	}
}
