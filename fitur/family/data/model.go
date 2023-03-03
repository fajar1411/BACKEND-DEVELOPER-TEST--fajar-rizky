package data

import (
	"test/fitur/family"

	"gorm.io/gorm"
)

type Family struct {
	gorm.Model
	Relation     string `gorm:"type:varchar(50);not null"`
	NameRelation string `gorm:"type:varchar(50);uniqueIndex;not null"`
	Dob          string `gorm:"type:varchar(50);not null"`
	CustomerID   uint
}
type FamilyUser struct {
	ID           uint
	Relation     string
	NameRelation string
	Dob          string
	Name         string
}

func Todata(data family.FamilyEntities) Family {
	return Family{
		Relation:     data.Relation,
		NameRelation: data.NameRelation,
		Dob:          data.Dob_date,
		CustomerID:   data.CustomerID,
	}
}
func (dataModel *FamilyUser) ModelsToCore() family.FamilyEntities { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return family.FamilyEntities{
		ID:           dataModel.ID,
		Relation:     dataModel.Relation,
		Dob_date:     dataModel.Dob,
		NameRelation: dataModel.NameRelation,
		Name:         dataModel.Name,
	}
}
func ToCore(data Family) family.FamilyEntities {
	return family.FamilyEntities{
		ID:           data.ID,
		Relation:     data.Relation,
		NameRelation: data.NameRelation,
		Dob_date:     data.Dob,
		CustomerID:   data.CustomerID,
	}
}
func ListModelTOCore(dataModel []FamilyUser) []family.FamilyEntities { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	var dataCore []family.FamilyEntities
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ModelsToCore())
	}
	return dataCore //  untuk menampilkan data ke controller
}
