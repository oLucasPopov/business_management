package timekeeping_repositories_protocols

import (
	"pontos_funcionario/src/models"
	"time"
)

type ClockIn interface {
	Handle(addClockIn models.AddClockInEmployee) (models.TimeKeepingEmployee, error)
}

type ClockOut interface {
	Handle(addClockIn models.AddClockOutEmployee) (models.TimeKeepingEmployee, error)
}

type DeleteTimekeeping interface {
	Handle(id int64) error
}

type ListTimekeeping interface {
	Handle(page int32, beginDate *time.Time, endDate *time.Time) (models.Timekeepings, error)
}

type GetOpenClockIn interface {
	Handle(employeeId int64) (*int64, error)
}

type GetClosedClockIn interface {
	Handle(clockinId int64) (*bool, error)
}
