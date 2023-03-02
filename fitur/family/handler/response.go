package handler

import (
	"test/fitur/family"
)

type FamilyResponse struct {
	Name         string `json:"customer name"`
	Dob_date     string `json:"tanggal"`
	Relation     string `json:"relasi"`
	NameRelation string `json:"name relasi"`
}

func ToFormResponse(data family.FamilyEntities) FamilyResponse {
	return FamilyResponse{
		Name:         data.Name,
		Relation:     data.Relation,
		Dob_date:     data.Dob_date,
		NameRelation: data.NameRelation,
	}
}

func ListEntitiesToPostsRespon(dataCore []family.FamilyEntities) []FamilyResponse {
	var ResponData []FamilyResponse

	for _, value := range dataCore {
		ResponData = append(ResponData, ToFormResponse(value))
	}
	return ResponData
}
