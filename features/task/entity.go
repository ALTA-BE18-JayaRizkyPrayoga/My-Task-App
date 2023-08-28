package task

type Core struct {
	ID        uint
	Status    string
	ProjectID uint
	Project   ProjectCore
}

type ProjectCore struct {
	ID          uint
	Name        string
	Description string
}

type TaskDataInterface interface {
	SelectAll() ([]Core, error)
	Insert(input Core) error
	Delete(id uint) error
	GetByID(id uint) (Core, error)
	Update(id uint, UpdateTaskByID Core) error
}

type TaskServiceInterface interface {
	GetAll() ([]Core, error)
	Create(input Core) error
	GetByID(id uint) (Core, error)
	Delete(id uint) error
	Update(id uint, UpdateTaskByID Core) error
}
