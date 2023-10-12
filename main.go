package main

import (
	"fmt"
	employee_factory "pontos_funcionario/src/main/factories"
)

func main() {

	add := employee_factory.MakeEmployee()

	// a := add.Handle(`{"name":"any_name","salary":10.89,"salary_type":"H", "cpf_cnpj": "43699838862"}`)
	a := add.Handle(`{}`)

	fmt.Print(a)
}
