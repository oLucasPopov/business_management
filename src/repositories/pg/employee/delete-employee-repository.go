package pg_employee_repositories

import pg_helpers "pontos_funcionario/src/repositories/pg/helpers"

type DeleteEmployee struct {
}

const deleteEmployeeSQL string = "DELETE FROM FUNCIONARIOS WHERE ID = $1"

func (c *DeleteEmployee) Handle(id int64) error {
	db, err := pg_helpers.PostgresConnect()
	if err != nil {
		return err
	}
	defer db.Close()

	smt, err := db.Prepare(deleteEmployeeSQL)
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
