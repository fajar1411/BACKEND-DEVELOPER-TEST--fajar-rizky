package data

import (
	"test/fitur/customer"
	"test/fitur/family/data"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name       string `gorm:"type:char(50);not null"`
	Dob_date   string `gorm:"not null"`
	Phonenum   string `gorm:"type:varchar(20);not null"`
	Email      string `gorm:"type:varchar(50);unique;not null"`
	Password   string `gorm:"not null"`
	NasionalID uint
	Families   []data.Family
}

func FromEntities(dataCore customer.CustomerEntites) Customer { //fungsi yang mengambil data dari entities usercore dan merubah data ke user gorm(model.go)
	return Customer{

		Email:    dataCore.Email,
		Dob_date: dataCore.Dob_date,
		Phonenum: dataCore.Phonenum,
		Name:     dataCore.Name,
		Password: dataCore.Password,

		// Role:      dataCore.Role,
		// Image_url: dataCore.Image_url,
	}

}
func ToCore(data Customer) customer.CustomerEntites {
	return customer.CustomerEntites{
		ID:       data.ID,
		Email:    data.Email,
		Dob_date: data.Dob_date,
		Phonenum: data.Phonenum,
		Name:     data.Name,
		Password: data.Password,
	}
}
func (dataModel *Customer) ModelsToCore() customer.CustomerEntites { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return customer.CustomerEntites{
		Name:     dataModel.Name,
		Email:    dataModel.Email, //mapping data core ke data gorm model
		Password: dataModel.Password,
		Dob_date: dataModel.Dob_date,
		Phonenum: dataModel.Phonenum,
	}
}
