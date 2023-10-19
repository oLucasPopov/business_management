package pg_employee_repositories

import (
	"pontos_funcionario/src/models"
	pg_helpers "pontos_funcionario/src/repositories/pg/helpers"
)

type ListEmployees struct {
}

const listEmployeeSQL = "SELECT * FROM FUNCIONARIOS OFFSET ($1 - 1) * 20 LIMIT 20"

func (c *ListEmployees) Handle(page int32) ([]*models.Employee, error) {
	employees := []*models.Employee{}
	db, err := pg_helpers.PostgresConnect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	smt, err := db.Prepare(listEmployeeSQL)
	if err != nil {
		return nil, err
	}
	defer smt.Close()

	rows, err := smt.Query(page)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		employee := &models.Employee{}
		err := rows.Scan(
			&employee.Id,
			&employee.Name,
			&employee.CPF_CNPJ,
			&employee.Phone,
			&employee.Cellphone,
			&employee.Email,
			&employee.Obs,
			&employee.State_id,
			&employee.Town_id,
			&employee.Zipcode,
			&employee.Address,
			&employee.Number,
			&employee.Neighbourhood,
			&employee.Reference,
			&employee.SalaryType,
			&employee.Salary)

		employees = append(employees, employee)
		if err != nil {
			return nil, err
		}
	}

	return employees, nil
}
