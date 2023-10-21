package pg_timekeeping_repositories

import pg_helpers "pontos_funcionario/src/repositories/pg/helpers"

type OpenClockIn struct {
}

const openClockInSQL string = `
	SELECT COUNT(*)
	  FROM PONTOS_FUNCIONARIOS
	 WHERE ID_FUNCIONARIO = $1
	   AND DATA_HORA_SAIDA IS NULL
`

func (c *OpenClockIn) Handle(employeeId int64) (*int64, error) {
	db, err := pg_helpers.PostgresConnect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	smt, err := db.Prepare(openClockInSQL)
	if err != nil {
		return nil, err
	}
	defer smt.Close()

	rows, err := smt.Query(employeeId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var count int64

	if rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return nil, err
		}
	}

	return &count, nil
}
