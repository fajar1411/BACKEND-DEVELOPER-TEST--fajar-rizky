package handler

import (
	"log"
	"net/http"
	"strconv"
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

	return c.JSON(http.StatusOK, helper.PesanDataBerhasilHelper("Berhasil Menampilkan Data Keluarga Costumer", ListEntitiesToPostsRespon(res)))

}
func (fh *FamilyHandler) UpdateFamily(c echo.Context) error {
	CustomerID := middlewares.ExtractTokenUserId(c)
	id, _ := strconv.Atoi(c.Param("id"))
	Inputform := FamilyRequest{}

	errbind := c.Bind(&Inputform)
	if errbind != nil {
		return c.JSON(http.StatusBadRequest, "format inputan salah")
	}
	datacore := FamilyRequestToEnitities(Inputform)
	datacore.CustomerID = uint(CustomerID)
	res, err2 := fh.FamilyServices.UpdateFamily(id, datacore)

	if err2 != nil {
		return c.JSON(helper.PesanGagalHelper(err2.Error()))
	}

	return c.JSON(http.StatusCreated, helper.PesanDataBerhasilHelper("Selamat Data Keluarga Berhasil DiUpdate", ToFormResponse(res)))
}

func (fh *FamilyHandler) DeleteFamily(c echo.Context) error {
	CustomerID := middlewares.ExtractTokenUserId(c)
	famID, _ := strconv.Atoi(c.Param("id"))
	del := fh.FamilyServices.DeleteFamily(CustomerID, famID)
	if del != nil {
		return c.JSON(helper.PesanGagalHelper(del.Error()))
	}
	return c.JSON(http.StatusOK, helper.PesanSuksesHelper("Berhasil Menghapus Data Family"))
}
