package national

type NationalityEntities struct {
	ID               uint
	Nationality_Name string `validate:"required,min=5,required"`
	Nationality_Code string `validate:"required,min=2,required"`
}

type NationalService interface {
	AddNation(Addnation NationalityEntities) (NationalityEntities, error)
}

type NationalData interface {
	AddNation(Addnation NationalityEntities) (NationalityEntities, error)
}
