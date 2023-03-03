package faktory

import (
	customerdata "test/fitur/customer/data"
	customerservice "test/fitur/customer/service"
	familydata "test/fitur/family/data"
	familyservice "test/fitur/family/service"
	nationaldata "test/fitur/national/data"
	nationalservice "test/fitur/national/service"
	customerhandler "test/routes"
	familyhandler "test/routes"
	nationalhandler "test/routes"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	v := validator.New()

	userRepofaktory := customerdata.NewCustomer(db)
	userServiceFaktory := customerservice.NewService(userRepofaktory, v)
	customerhandler.NewHandlerCostumer(userServiceFaktory, e)

	familyRepofaktory := familydata.NewFamily(db)
	familyServiceFaktory := familyservice.NewService(familyRepofaktory, v)
	familyhandler.NewHandlerFamily(familyServiceFaktory, e)

	nationaldatafaktory := nationaldata.NewNational(db)
	nationalserviceFaktory := nationalservice.NewService(nationaldatafaktory, v)
	nationalhandler.NewHandlerNational(nationalserviceFaktory, e)
}
