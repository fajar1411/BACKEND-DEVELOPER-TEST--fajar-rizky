package handler

import (
	"net/http"
	"test/fitur/national"
	"test/helper"

	"github.com/labstack/echo/v4"
)

type NationalHandler struct {
	NationalServices national.NationalService
}

func (nd *NationalHandler) AddNational(c echo.Context) error {

	Inputform := NationalRequest{}

	errbind := c.Bind(&Inputform)
	if errbind != nil {
		return c.JSON(http.StatusBadRequest, "format inputan salah")
	}
	dataCore := NationalRequestToEnitities(Inputform)
	res, err := nd.NationalServices.AddNation(dataCore)
	if err != nil {
		return c.JSON(helper.PesanGagalHelper(err.Error()))
	}
	dataResp := ToFormResponse(res)
	return c.JSON(http.StatusCreated, helper.PesanDataBerhasilHelper("Mengisi Form Berhasil", dataResp))

}
