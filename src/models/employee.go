package models

type SalaryType string

const (
	Monthly SalaryType = "M"
	Hourly  SalaryType = "H"
)

type AddEmployee struct {
	Name          *string     `json:"name"`
	CPF_CNPJ      *string     `json:"cpf_cnpj"`
	Phone         *string     `json:"phone"`
	Cellphone     *string     `json:"cellphone"`
	Email         *string     `json:"email"`
	Obs           *string     `json:"obs"`
	State_id      *int64      `json:"state_id"`
	Town_id       *int64      `json:"town_id"`
	Zipcode       *string     `json:"zipcode"`
	Address       *string     `json:"address"`
	Number        *string     `json:"number"`
	Neighbourhood *string     `json:"neighbourhood"`
	Reference     *string     `json:"reference"`
	SalaryType    *SalaryType `json:"salary_type"`
	Salary        *float32    `json:"salary"`
}

type Employee struct {
	Id int64
	AddEmployee
}

type Employees []Employee
