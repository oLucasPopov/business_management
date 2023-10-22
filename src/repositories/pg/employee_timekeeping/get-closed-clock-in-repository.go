package pg_timekeeping_repositories

import pg_helpers "pontos_funcionario/src/repositories/pg/helpers"

type ClosedClockIn struct {
}

const closedClockInSQL string = `
	SELECT TRUE
	  FROM PONTOS_FUNCIONARIOS
	 WHERE ID = $1
	   AND DATA_HORA_SAIDA IS NOT NULL
`

func (c *ClosedClockIn) Handle(clockinId int64) (*bool, error) {
	db, err := pg_helpers.PostgresConnect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	smt, err := db.Prepare(closedClockInSQL)
	if err != nil {
		return nil, err
	}
	defer smt.Close()

	rows, err := smt.Query(clockinId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var closed bool = false

	if rows.Next() {
		err := rows.Scan(&closed)
		if err != nil {
			return nil, err
		}
	}

	return &closed, nil
}
