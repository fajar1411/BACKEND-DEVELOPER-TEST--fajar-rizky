package service

import (
	"errors"
	"log"
	"strings"
	"test/fitur/national"
	"test/validasi"

	"github.com/go-playground/validator/v10"
)

type nationalCase struct {
	qry national.NationalData
	vld *validator.Validate
}

func NewService(fd national.NationalData, vld *validator.Validate) national.NationalService {
	return &nationalCase{
		qry: fd,
		vld: vld,
	}
}

// AddNation implements national.NationalService
func (nc *nationalCase) AddNation(Addnation national.NationalityEntities) (national.NationalityEntities, error) {
	valerr := nc.vld.Struct(&Addnation)
	if valerr != nil {
		log.Println("validation error", valerr)
		msg := validasi.ValidationErrorHandle(valerr)
		return national.NationalityEntities{}, errors.New(msg)
	}

	res, err := nc.qry.AddNation(Addnation)
	if err != nil {
		msg2 := ""
		if strings.Contains(err.Error(), "not found") {
			msg2 = "content not found"
		} else {
			msg2 = "server error"
		}
		return national.NationalityEntities{}, errors.New(msg2)
	}

	return res, nil
}
