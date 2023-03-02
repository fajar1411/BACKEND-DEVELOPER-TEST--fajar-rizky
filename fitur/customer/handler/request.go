package handler

import "test/fitur/customer"

type CostumerRequest struct {
	Name     string `json:"name" form:"nama"`
	Dob_date string `json:"tanggal" form:"tanggal"`
	Phonenum string `json:"phone" form:"phone"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func CostumerRequestToUserCore(data CostumerRequest) customer.CustomerEntites {
	return customer.CustomerEntites{
		Name:     data.Name,
		Dob_date: data.Dob_date,
		Phonenum: data.Phonenum,
		Email:    data.Email,
		Password: data.Password,
	}
}
