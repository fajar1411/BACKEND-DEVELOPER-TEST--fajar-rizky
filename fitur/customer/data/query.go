package data

import (
	"errors"
	"log"
	"strings"
	"test/fitur/customer"

	"gorm.io/gorm"
)

type costumerData struct {
	db *gorm.DB
}

// Profile implements user.UserData

func NewCustomer(db *gorm.DB) customer.CostumerData {
	return &costumerData{
		db: db,
	}
}

// FormData implements customer.CostumerData
func (cd *costumerData) FormData(newUser customer.CustomerEntites) (customer.CustomerEntites, error) {
	userGorm := FromEntities(newUser)

	tx := cd.db.Create(&userGorm) // proses insert data

	if tx.Error != nil {
		log.Println("register query error", tx.Error.Error())
		msg := ""
		if strings.Contains(tx.Error.Error(), "Duplicate") {
			msg = "data is duplicated"
		} else {
			msg = "server error"
		}
		return customer.CustomerEntites{}, errors.New(msg)
	}
	return newUser, nil
}

// Login implements customer.CostumerData
func (cd *costumerData) Login(email string) (customer.CustomerEntites, error) {
	res := Customer{}

	if err := cd.db.Where("email = ?", email).First(&res).Error; err != nil {
		log.Println("login query error", err.Error())
		return customer.CustomerEntites{}, errors.New("data not found")
	}

	return ToCore(res), nil
}

// Profile implements customer.CostumerData
func (cd *costumerData) Profile(id int) (customer.CustomerEntites, error) {
	users := CustomerName{}
	err := cd.db.Raw("SELECT customers.dob_date, customers.email, customers.phonenum, nationalities.nationality_name, customers.name FROM customers JOIN nationalities ON nationalities.id = customers.nationalities_id WHERE customers.id = ?", id).Find(&users).Error

	if err != nil {
		return customer.CustomerEntites{}, err
	}
	gorms := users.ModelsToCore()
	return gorms, nil

}

// Delete implements customer.CostumerData
func (cd *costumerData) Delete(id int) error {
	users := Customer{}
	qry := cd.db.Delete(&users, id)

	rowAffect := qry.RowsAffected
	if rowAffect <= 0 {
		log.Println("no data processed")
		return errors.New("no user has delete")
	}

	err := qry.Error
	if err != nil {
		log.Println("delete query error", err.Error())
		return errors.New("delete account fail")
	}

	return nil
}

// Update implements customer.CostumerData
func (cd *costumerData) Update(id int, Updata customer.CustomerEntites) (customer.CustomerEntites, error) {
	users := Customer{}
	userGorm := FromEntities(Updata)
	qry := cd.db.Model(&users).Where("id = ?", id).Updates(&userGorm)
	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return customer.CustomerEntites{}, errors.New("no data updated")
	}
	err := qry.Error
	if err != nil {
		log.Println("update user query error", err.Error())
		return customer.CustomerEntites{}, err
	}

	return ToCore(userGorm), nil
}
