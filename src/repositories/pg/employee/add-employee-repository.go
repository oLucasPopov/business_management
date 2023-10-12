package pg_employee_repositories

import (
	"pontos_funcionario/src/models"
	pg_helpers "pontos_funcionario/src/repositories/pg/helpers"
)

type Employee struct {
}

const (
	insertSQL string = `
	insert into funcionarios(
		nome,
		cpf_cnpj,
		telefone,
		celular,
		email,
		observacoes,
		id_estado,
		id_cidade,
		cep,
		rua,
		numero,
		bairro,
		referencia,
		tipo_salario,
		salario)
	values(
		 $1
		,$2
		,$3
		,$4
		,$5
		,$6
		,$7
		,$8
		,$9
		,$10
		,$11
		,$12
		,$13
		,$14
		,$15) returning *`
)

func (e *Employee) Add(addEmployee models.AddEmployee) (models.Employee, error) {
	db, err := pg_helpers.PostgresConnect()
	createdEmployee := models.Employee{}
	if err != nil {
		return createdEmployee, err
	}
	defer db.Close()

	stm, err := db.Prepare(insertSQL)

	if err != nil {
		return createdEmployee, err
	}
	defer stm.Close()

	res, err := stm.Query(
		addEmployee.Name,
		addEmployee.CPF_CNPJ,
		addEmployee.Phone,
		addEmployee.Cellphone,
		addEmployee.Email,
		addEmployee.Obs,
		addEmployee.State_id,
		addEmployee.Town_id,
		addEmployee.Zipcode,
		addEmployee.Address,
		addEmployee.Number,
		addEmployee.Neighbourhood,
		addEmployee.Reference,
		addEmployee.SalaryType,
		addEmployee.Salary,
	)
	if err != nil {
		return createdEmployee, err
	}
	defer res.Close()

	if res.Next() {
		err := res.Scan(
			&createdEmployee.Id,
			&createdEmployee.Name,
			&createdEmployee.CPF_CNPJ,
			&createdEmployee.Phone,
			&createdEmployee.Cellphone,
			&createdEmployee.Email,
			&createdEmployee.Obs,
			&createdEmployee.State_id,
			&createdEmployee.Town_id,
			&createdEmployee.Zipcode,
			&createdEmployee.Address,
			&createdEmployee.Number,
			&createdEmployee.Neighbourhood,
			&createdEmployee.Reference,
			&createdEmployee.SalaryType,
			&createdEmployee.Salary)
		if err != nil {
			return createdEmployee, err
		}
	}

	return createdEmployee, nil
}
