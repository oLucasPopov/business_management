package clock_controller

import (
	"encoding/json"
	"net/http"
	controller_protocols "pontos_funcionario/src/controllers/protocols"
	pg_timekeeping_repositories "pontos_funcionario/src/repositories/pg/employee_timekeeping"
	"strconv"
	"time"
)

type ListTimekeeping struct {
	ListTimekeepingRepository pg_timekeeping_repositories.ListTimekeeping
}

func (c *ListTimekeeping) Handle(request *controller_protocols.ControllerRequest) controller_protocols.ControllerResponse {
	const dateLayout = "2006-01-02"
	pageStr := request.Query.Get("page")
	beginDateStr := request.Query.Get("beginDate")
	endDateStr := request.Query.Get("endDate")

	var page int64 = 1
	var beginDate *time.Time = nil
	var endDate *time.Time = nil

	if pageStr != "" {
		page, _ = strconv.ParseInt(pageStr, 10, 32)
	}

	if beginDateStr != "" {
		begin, _ := time.Parse(dateLayout, beginDateStr)
		beginDate = &begin
	}

	if endDateStr != "" {
		end, _ := time.Parse(dateLayout, endDateStr)
		end = end.Add(time.Hour * 23).Add(time.Minute * 59).Add(time.Second * 59)
		endDate = &end
	}

	timekeeping, err := c.ListTimekeepingRepository.Handle(int32(page), beginDate, endDate)
	if err != nil {
		return controller_protocols.ControllerResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}
	}

	jsonTimekeeping, err := json.Marshal(timekeeping)
	if err != nil {
		return controller_protocols.ControllerResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}
	}

	return controller_protocols.ControllerResponse{
		StatusCode: http.StatusOK,
		Body:       jsonTimekeeping,
	}
}
