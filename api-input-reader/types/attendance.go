package types

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type AttendanceSDC struct {
	ConnectionKey    string    `json:"connection_key"`
	Result           bool      `json:"result"`
	RedisKey         string    `json:"redis_key"`
	Filepath         string    `json:"filepath"`
	APIStatusCode    int       `json:"api_status_code"`
	RuntimeSessionID string    `json:"runtime_session_id"`
	BusinessPartner  *int      `json:"business_partner"`
	ServiceLabel     string    `json:"service_label"`
	APIType   		 string    `json:"api_type"`
	Header    		 AttendanceHeader `json:"Attendance"`
	APISchema 		 string    `json:"api_schema"`
	Accepter         []string  `json:"accepter"`
	Deleted          bool      `json:"deleted"`
}

type AttendanceHeader struct {
	Attendance				*int	`json:"Attendance"`
	AttendanceDate			string	`json:"AttendanceDate"`
	AttendanceTime			string	`json:"AttendanceTime"`
	Attender				int		`json:"Attender"`
	AttendanceObjectType	string	`json:"AttendanceObjectType"`
	AttendanceObject		int		`json:"AttendanceObject"`
	Participation			*int	`json:"Participation"`
	CreationDate			string	`json:"CreationDate"`
	CreationTime			string	`json:"CreationTime"`
	IsCancelled				*bool	`json:"IsCancelled"`
}

func AttendanceInputRead(
	controller *beego.Controller,
) AttendanceSDC {
	var requestBody AttendanceSDC
	err := json.Unmarshal(controller.Ctx.Input.RequestBody, &requestBody)

	if err != nil {
		controller.Data["json"] = map[string]interface{}{
			"error": fmt.Sprintf("json decode error: %v", err),
		}
		controller.ServeJSON()
	}

	return requestBody
}
