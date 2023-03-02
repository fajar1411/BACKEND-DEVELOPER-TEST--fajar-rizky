package service

import (
	"errors"
	"log"
	"strings"
	"test/fitur/family"
	"test/validasi"

	"github.com/go-playground/validator/v10"
)

type familyCase struct {
	qry family.FamilyData
	vld *validator.Validate
}

func NewService(fd family.FamilyData, vld *validator.Validate) family.FamilyService {
	return &familyCase{
		qry: fd,
		vld: vld,
	}
}

// Add implements family.FamilyService
func (fc *familyCase) AddFamily(newFamily family.FamilyEntities) (family.FamilyEntities, error) {

	valerr := fc.vld.Struct(&newFamily)
	if valerr != nil {
		log.Println("validation error", valerr)
		msg := validasi.ValidationErrorHandle(valerr)
		return family.FamilyEntities{}, errors.New(msg)
	}
	res, err := fc.qry.AddFamily(newFamily)
	if err != nil {
		// fmt.Println(err)
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "content family not found"
		} else {
			msg = "internal server error"
		}
		return family.FamilyEntities{}, errors.New(msg)
	}

	return res, nil
}

// MyFamily implements family.FamilyService
func (fc *familyCase) MyFamily(userID int) ([]family.FamilyEntities, error) {

	if userID <= 0 {
		return nil, errors.New("user not found")
	}

	res, _ := fc.qry.MyFamily(userID)

	return res, nil
}

// UpdateFamily implements family.FamilyService
func (*familyCase) UpdateFamily(updata family.FamilyEntities) (family.FamilyEntities, error) {
	panic("unimplemented")
}
