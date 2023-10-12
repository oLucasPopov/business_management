package models

type SalaryType string

const (
	Monthly SalaryType = "M"
	Hourly  SalaryType = "H"
)

type AddEmployee struct {
	Name          *string     `json:"name"`
	CPF_CNPJ      *string     `json:"cpf_cnpj,omitempty"`
	Phone         *string     `json:"phone,omitempty"`
	Cellphone     *string     `json:"cellphone,omitempty"`
	Email         *string     `json:"email,omitempty"`
	Obs           *string     `json:"obs,omitempty"`
	State_id      *int64      `json:"state_id,omitempty"`
	Town_id       *int64      `json:"town_id,omitempty"`
	Zipcode       *string     `json:"zipcode,omitempty"`
	Address       *string     `json:"address,omitempty"`
	Number        *string     `json:"number,omitempty"`
	Neighbourhood *string     `json:"neighbourhood,omitempty"`
	Reference     *string     `json:"reference,omitempty"`
	SalaryType    *SalaryType `json:"salary_type"`
	Salary        *float32    `json:"salary"`
}

type Employee struct {
	Id int64
	AddEmployee
}
