package pg_timekeeping_repositories

import (
	"pontos_funcionario/src/models"
	pg_helpers "pontos_funcionario/src/repositories/pg/helpers"
)

type ClockIn struct {
}

const (
	insertSQL string = `
	insert into PONTOS_FUNCIONARIOS(
		ID_FUNCIONARIO
	 ,DATA_HORA_ENTRADA
	 ,SALARIO_FUNCIONARIO
	)
	values(
		$1
	 ,$2
	 ,$3
	) returning *`
)

func (c *ClockIn) Handle(addClockIn models.AddClockInEmployee) (models.TimeKeepingEmployee, error) {
	db, err := pg_helpers.PostgresConnect()
	employeeTimeKeeping := models.TimeKeepingEmployee{}
	if err != nil {
		return employeeTimeKeeping, err
	}
	defer db.Close()

	stm, err := db.Prepare(insertSQL)

	if err != nil {
		return employeeTimeKeeping, err
	}
	defer stm.Close()

	res, err := stm.Query(
		addClockIn.EmployeeId,
		addClockIn.ClockInDateTime,
		addClockIn.EmployeeSalary,
	)

	if err != nil {
		return employeeTimeKeeping, err
	}
	defer res.Close()

	if res.Next() {
		err := res.Scan(
			&employeeTimeKeeping.Id,
			&employeeTimeKeeping.EmployeeId,
			&employeeTimeKeeping.ClockInDateTime,
			&employeeTimeKeeping.ClockOutDateTime,
			&employeeTimeKeeping.EmployeeSalary,
		)

		if err != nil {
			return employeeTimeKeeping, err
		}
	}

	return employeeTimeKeeping, nil
}
