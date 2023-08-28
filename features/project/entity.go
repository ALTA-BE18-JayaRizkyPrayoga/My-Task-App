package project

type Core struct {
	ID          uint
	Name        string
	UserID      uint
	Description string
	User        UserCore
}

type UserCore struct {
	ID    uint
	Name  string
	Email string
}

type ProjectDataInterface interface {
	SelectAll() ([]Core, error)
	Insert(input Core) error
	Delete(id uint) error
	GetByID(id uint) (Core, error)
	Update(id uint, UpdateProjectByID Core) error
}

type ProjectServiceInterface interface {
	GetAll() ([]Core, error)
	Create(input Core) error
	GetByID(id uint) (Core, error)
	Delete(id uint) error
	Update(id uint, UpdateProjectByID Core) error
}
