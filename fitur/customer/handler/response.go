package handler

import "test/fitur/customer"

type FormResponse struct {
	Name     string `json:"nama"`
	Dob_date string `json:"tanggal lahir"`
	Phonenum string `json:"phone"`
	Email    string `json:"email"`
}
type ProfileResponse struct {
	Name             string `json:"nama"`
	Dob_date         string `json:"tanggal lahir"`
	Phonenum         string `json:"phone"`
	Email            string `json:"email"`
	Nationality_name string `json:"national"`
}
type LoginResponse struct {
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func ToFormResponse(data customer.CustomerEntites) FormResponse {
	return FormResponse{
		Name:     data.Name,
		Email:    data.Email,
		Dob_date: data.Dob_date,
		Phonenum: data.Phonenum,
	}
}
func ToProfileResponse(data customer.CustomerEntites) ProfileResponse {
	return ProfileResponse{
		Name:             data.Name,
		Email:            data.Email,
		Dob_date:         data.Dob_date,
		Phonenum:         data.Phonenum,
		Nationality_name: data.NationalName,
	}
}
func ToLoginRespon(data customer.CustomerEntites, token string) LoginResponse {
	return LoginResponse{

		Nama:  data.Name,
		Email: data.Email,
		Token: token,
	}
}
