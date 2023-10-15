package pg_employee_repositories

import (
	"pontos_funcionario/src/models"
	pg_helpers "pontos_funcionario/src/repositories/pg/helpers"
)

type GetEmployee struct {
}

const getEmployeeSQL = "SELECT * FROM FUNCIONARIOS WHERE ID = $1"

func (c *GetEmployee) Handle(id int64) (*models.Employee, error) {
	employee := models.Employee{}
	db, err := pg_helpers.PostgresConnect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	smt, err := db.Prepare(getEmployeeSQL)
	if err != nil {
		return nil, err
	}
	defer smt.Close()

	rows, err := smt.Query(id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
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
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}

	return &employee, nil
}
