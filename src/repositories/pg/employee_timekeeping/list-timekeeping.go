package pg_timekeeping_repositories

import (
	"pontos_funcionario/src/models"
	pg_helpers "pontos_funcionario/src/repositories/pg/helpers"
	"time"
)

type ListTimekeeping struct {
}

const listTimekeepingSQL = `
SELECT * 
  FROM PONTOS_FUNCIONARIOS 
 WHERE DATA_HORA_ENTRADA >= COALESCE($1, DATA_HORA_ENTRADA)
   AND DATA_HORA_ENTRADA <= COALESCE($2, DATA_HORA_ENTRADA)
OFFSET ($3 - 1) * 20 
 LIMIT 20
`

func (c *ListTimekeeping) Handle(page int32, beginDate *time.Time, endDate *time.Time) (models.Timekeepings, error) {
	timekeepings := models.Timekeepings{}

	db, err := pg_helpers.PostgresConnect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	smt, err := db.Prepare(listTimekeepingSQL)
	if err != nil {
		return nil, err
	}
	defer smt.Close()
	rows, err := smt.Query(beginDate, endDate, page)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		timekeeping := &models.TimeKeepingEmployee{}
		err := rows.Scan(
			&timekeeping.Id,
			&timekeeping.EmployeeId,
			&timekeeping.ClockInDateTime,
			&timekeeping.ClockOutDateTime,
			&timekeeping.EmployeeSalary,
		)

		timekeepings = append(timekeepings, timekeeping)
		if err != nil {
			return nil, err
		}
	}

	return timekeepings, nil
}
