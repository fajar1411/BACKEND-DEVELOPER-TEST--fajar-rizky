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
func (fd *familyData) UpdateFamily(id int, updata family.FamilyEntities) (family.FamilyEntities, error) {

	data := Todata(updata)
	qry := fd.db.Model(&data).Where("id = ?", id).Updates(&data)
	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return family.FamilyEntities{}, errors.New("no data updated")
	}
	err := qry.Error
	if err != nil {
		log.Println("update user query error", err.Error())
		return family.FamilyEntities{}, err
	}

	return ToCore(data), nil
}

// DeleteFamily implements family.FamilyData
func (fd *familyData) DeleteFamily(iduser int, id int) error {
	var Family Family
	tx := fd.db.Where("id = ? AND families.customer_id = ?", id, iduser).Delete(&Family, id)
	rowAffect := tx.RowsAffected
	if rowAffect <= 0 {
		log.Println("no data processed")
		return errors.New("no user has delete")
	}

	err := tx.Error
	if err != nil {
		log.Println("delete query error", err.Error())
		return errors.New("delete Family fail")
	}

	return nil
}
