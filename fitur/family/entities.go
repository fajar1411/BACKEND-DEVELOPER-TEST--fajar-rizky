package family

type FamilyEntities struct {
	ID           uint
	Relation     string `validate:"required,min=5,required"`
	Dob_date     string
	NameRelation string `validate:"required,min=5,required"`
	Name         string
	CustomerID   uint
}

type FamilyService interface {
	AddFamily(newFamily FamilyEntities) (FamilyEntities, error)
	MyFamily(userID int) ([]FamilyEntities, error)
	UpdateFamily(id int, updata FamilyEntities) (FamilyEntities, error)
	DeleteFamily(iduser int, id int) error
}

type FamilyData interface {
	AddFamily(newFamily FamilyEntities) (FamilyEntities, error)
	MyFamily(userID int) ([]FamilyEntities, error)
	UpdateFamily(id int, updata FamilyEntities) (FamilyEntities, error)
	DeleteFamily(iduser int, id int) error
}
