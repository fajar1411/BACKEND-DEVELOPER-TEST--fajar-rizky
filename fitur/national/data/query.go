package data

import (
	"errors"
	"log"
	"strings"
	"test/fitur/national"

	"gorm.io/gorm"
)

type nationalData struct {
	db *gorm.DB
}

func NewNational(db *gorm.DB) national.NationalData {
	return &nationalData{
		db: db,
	}
}

// AddNation implements national.NationalData
func (nd *nationalData) AddNation(Addnation national.NationalityEntities) (national.NationalityEntities, error) {
	National := FromEntities(Addnation)

	tx := nd.db.Create(&National) // proses insert data

	if tx.Error != nil {
		log.Println("register query error", tx.Error.Error())
		msg := ""
		if strings.Contains(tx.Error.Error(), "Duplicate") {
			msg = "data is duplicated"
		} else {
			msg = "server error"
		}
		return national.NationalityEntities{}, errors.New(msg)
	}
	return Addnation, nil
}
