package customer

type CustomerEntites struct {
	ID       uint
	Name     string `validate:"required,min=5,required"`
	Dob_date string
	Phonenum string `validate:"required,min=5,required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=5,required"`
	Nasional uint
}

type CostumerService interface {
	FormData(newUser CustomerEntites) (CustomerEntites, error)
	Login(email, password string) (string, CustomerEntites, error)
	Profile(id int) (CustomerEntites, error)
	Update(id int, Updata CustomerEntites) (CustomerEntites, error)
	Delete(id int) error
}

type CostumerData interface {
	FormData(newUser CustomerEntites) (CustomerEntites, error)
	Login(email string) (CustomerEntites, error)
	Profile(id int) (CustomerEntites, error)
	Update(id int, Updata CustomerEntites) (CustomerEntites, error)
	Delete(id int) error
}
