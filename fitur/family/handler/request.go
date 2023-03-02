package handler

import "test/fitur/family"

type FamilyRequest struct {
	NameRelation string `json:"name" form:"name"`
	Dob_date     string `json:"tanggal" form:"tanggal"`
	Relation     string `json:"relasi" form:"relasi"`
}

func FamilyRequestToEnitities(data FamilyRequest) family.FamilyEntities {
	return family.FamilyEntities{
		NameRelation: data.NameRelation,
		Dob_date:     data.Dob_date,
		Relation:     data.Relation,
	}
}
