package handler

import (
	"log"
	"net/http"
	"test/fitur/family"
	"test/helper"
	"test/middlewares"

	"github.com/labstack/echo/v4"
)

type FamilyHandler struct {
	FamilyServices family.FamilyService
}

func (fh *FamilyHandler) AddFamily(c echo.Context) error {
	CustomerID := middlewares.ExtractTokenUserId(c)

	Inputform := FamilyRequest{}

	errbind := c.Bind(&Inputform)
	if errbind != nil {
		return c.JSON(http.StatusBadRequest, "format inputan salah")
	}
	datacore := FamilyRequestToEnitities(Inputform)
	datacore.CustomerID = uint(CustomerID)
	res, err2 := fh.FamilyServices.AddFamily(datacore)
	log.Print("ini handler", res)
	if err2 != nil {
		return c.JSON(helper.PesanGagalHelper(err2.Error()))
	}

	return c.JSON(http.StatusCreated, helper.PesanDataBerhasilHelper("Selamat Data Keluarga Berhasil Ditambahkan", ToFormResponse(res)))

}
func (fh *FamilyHandler) MyFamily(c echo.Context) error {
	CustomerID := middlewares.ExtractTokenUserId(c)
	res, _ := fh.FamilyServices.MyFamily(CustomerID)

	return c.JSON(http.StatusCreated, helper.PesanDataBerhasilHelper("Berhasil Menampilkan Data Keluarga Costumer", ListEntitiesToPostsRespon(res)))

}
