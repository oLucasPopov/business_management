package pg_timekeeping_repositories

import (
	"pontos_funcionario/src/models"
	pg_helpers "pontos_funcionario/src/repositories/pg/helpers"
)

type ClockOut struct {
}

const (
	clockOutSQL string = `
	UPDATE PONTOS_FUNCIONARIOS
	   SET DATA_HORA_SAIDA = $1
	 WHERE ID = $2
	RETURNING *
	`
)

func (c *ClockOut) Handle(addClockIn models.AddClockOutEmployee) (models.TimeKeepingEmployee, error) {
	db, err := pg_helpers.PostgresConnect()
	employeeTimeKeeping := models.TimeKeepingEmployee{}
	if err != nil {
		return employeeTimeKeeping, err
	}
	defer db.Close()

	stm, err := db.Prepare(clockOutSQL)

	if err != nil {
		return employeeTimeKeeping, err
	}
	defer stm.Close()

	res, err := stm.Query(
		addClockIn.ClockOutDateTime,
		addClockIn.Id,
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
