package models

import "time"

type AddClockInEmployee struct {
	EmployeeId      *int64     `json:"employee_id"`
	ClockInDateTime *time.Time `json:"clock_in"`
	EmployeeSalary  *float32   `json:"salary"`
}

type AddClockOutEmployee struct {
	Id               *int64     `json:"id"`
	ClockOutDateTime *time.Time `json:"clock_out"`
}

type TimeKeepingEmployee struct {
	Id               *int64     `json:"id"`
	ClockOutDateTime *time.Time `json:"clock_out"`
	AddClockInEmployee
}
