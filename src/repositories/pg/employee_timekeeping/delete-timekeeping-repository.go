package pg_timekeeping_repositories

import pg_helpers "pontos_funcionario/src/repositories/pg/helpers"

type DeleteTimekeeping struct {
}

const deleteTimekeepingSQL string = `
  DELETE 
	  FROM PONTOS_FUNCIONARIOS
	 WHERE ID = $1
`

func (c *DeleteTimekeeping) Handle(id int64) error {
	db, err := pg_helpers.PostgresConnect()
	if err != nil {
		return err
	}
	defer db.Close()

	smt, err := db.Prepare(deleteTimekeepingSQL)
	if err != nil {
		return err
	}
	defer smt.Close()

	_, err = smt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
