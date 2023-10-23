package employee_repositories_protocols

import "pontos_funcionario/src/models"

type AddEmployee interface {
	Add(addEmployee models.AddEmployee) (models.Employee, error)
}

type DeleteEmployee interface {
	Handle(id int64) error
}

type GetEmployee interface {
	Handle(id int64) (*models.Employee, error)
}

type ListEmployees interface {
	Handle(page int32) ([]*models.Employee, error)
}

type UpdateEmployee interface {
	Handle(employee models.Employee) (models.Employee, error)
}
