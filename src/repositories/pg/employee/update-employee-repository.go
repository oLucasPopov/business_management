package pg_employee_repositories

import (
	"fmt"
	"pontos_funcionario/src/models"
	pg_helpers "pontos_funcionario/src/repositories/pg/helpers"
)

type UpdateEmployee struct {
}

const (
	updateEmployeeSQL string = `
	update funcionarios
	   set nome         = coalesce($1, nome)
		    ,cpf_cnpj     = coalesce($2, cpf_cnpj)
		    ,telefone     = coalesce($3, telefone)
		    ,celular      = coalesce($4, celular)
		    ,email        = coalesce($5, email)
		    ,observacoes  = coalesce($6, observacoes)
		    ,id_estado    = coalesce($7, id_estado)
		    ,id_cidade    = coalesce($8, id_cidade)
		    ,cep          = coalesce($9, cep)
		    ,rua          = coalesce($10, rua)
		    ,numero       = coalesce($11, numero)
		    ,bairro       = coalesce($12, bairro)
		    ,referencia   = coalesce($13, referencia)
		    ,tipo_salario = coalesce($14, tipo_salario)
		    ,salario      = coalesce($15, salario)
		where id = $16
		returning *`
)

func (e *UpdateEmployee) Handle(employee models.Employee) (models.Employee, error) {
	db, err := pg_helpers.PostgresConnect()
	updatedEmployee := models.Employee{}
	if err != nil {
		return updatedEmployee, err
	}
	defer db.Close()

	stm, err := db.Prepare(updateEmployeeSQL)

	if err != nil {
		return updatedEmployee, err
	}
	defer stm.Close()

	fmt.Println(employee)

	res, err := stm.Query(
		employee.Name,
		employee.CPF_CNPJ,
		employee.Phone,
		employee.Cellphone,
		employee.Email,
		employee.Obs,
		employee.State_id,
		employee.Town_id,
		employee.Zipcode,
		employee.Address,
		employee.Number,
		employee.Neighbourhood,
		employee.Reference,
		employee.SalaryType,
		employee.Salary,
		employee.Id,
	)

	if err != nil {
		return updatedEmployee, err
	}

	defer res.Close()

	if res.Next() {
		err := res.Scan(
			&updatedEmployee.Id,
			&updatedEmployee.Name,
			&updatedEmployee.CPF_CNPJ,
			&updatedEmployee.Phone,
			&updatedEmployee.Cellphone,
			&updatedEmployee.Email,
			&updatedEmployee.Obs,
			&updatedEmployee.State_id,
			&updatedEmployee.Town_id,
			&updatedEmployee.Zipcode,
			&updatedEmployee.Address,
			&updatedEmployee.Number,
			&updatedEmployee.Neighbourhood,
			&updatedEmployee.Reference,
			&updatedEmployee.SalaryType,
			&updatedEmployee.Salary,
		)
		if err != nil {
			return updatedEmployee, err
		}
	}

	return updatedEmployee, nil
}
