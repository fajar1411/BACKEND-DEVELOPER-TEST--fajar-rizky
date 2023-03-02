package data

import (
	"errors"
	"log"
	"strings"
	"test/fitur/family"

	"gorm.io/gorm"
)

type familyData struct {
	db *gorm.DB
}

func NewFamily(db *gorm.DB) family.FamilyData {
	return &familyData{
		db: db,
	}
}

// Add implements family.FamilyData
func (fd *familyData) AddFamily(newFamily family.FamilyEntities) (family.FamilyEntities, error) {
	data := Todata(newFamily)
	err := fd.db.Create(&data)
	if err.Error != nil {
		log.Println("add Family query error", err.Error.Error())
		msg := ""
		if strings.Contains(err.Error.Error(), "not valid") {
			msg = "wrong input"

		} else {
			msg = "server error"
		}
		return family.FamilyEntities{}, errors.New(msg)
	}

	return newFamily, nil
}

// MyFamily implements family.FamilyData
func (fd *familyData) MyFamily(userID int) ([]family.FamilyEntities, error) {
	myfamily := []FamilyUser{}
	err := fd.db.Raw("SELECT families.id, families.relation,  families.name_relation, families.dob, customers.name FROM families JOIN customers ON customers.id = families.customer_id WHERE families.customer_id = ?", userID).Find(&myfamily).Error

	if err != nil {
		return nil, err
	}

	var dataCore = ListModelTOCore(myfamily)

	return dataCore, nil
}

// UpdateFamily implements family.FamilyData
func (*familyData) UpdateFamily(updata family.FamilyEntities) (family.FamilyEntities, error) {
	panic("unimplemented")
}
