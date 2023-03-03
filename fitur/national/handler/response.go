package handler

import "test/fitur/national"

type NationalityResponse struct {
	Nationality_Name string `json:"nasional_name" `
	Nationality_Code string `json:"code"`
}

func ToFormResponse(data national.NationalityEntities) NationalityResponse {
	return NationalityResponse{
		Nationality_Name: data.Nationality_Name,
		Nationality_Code: data.Nationality_Code,
	}
}
