package handler

import (
	"test/fitur/national"
)

type NationalRequest struct {
	Nationality_Name string `json:"name" form:"name"`
	Nationality_Code string `json:"code" form:"code"`
}

func NationalRequestToEnitities(data NationalRequest) national.NationalityEntities {
	return national.NationalityEntities{
		Nationality_Name: data.Nationality_Name,
		Nationality_Code: data.Nationality_Code,
	}
}
