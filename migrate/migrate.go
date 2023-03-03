package migrate

import (
	customer "test/fitur/customer/data"
	family "test/fitur/family/data"
	national "test/fitur/national/data"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&customer.Customer{})
	db.AutoMigrate(&family.Family{})
	db.AutoMigrate(&national.Nationality{})

}
