package types

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type ParticipationSDC struct {
	ConnectionKey    string    `json:"connection_key"`
	Result           bool      `json:"result"`
	RedisKey         string    `json:"redis_key"`
	Filepath         string    `json:"filepath"`
	APIStatusCode    int       `json:"api_status_code"`
	RuntimeSessionID string    `json:"runtime_session_id"`
	BusinessPartner  *int      `json:"business_partner"`
	ServiceLabel     string    `json:"service_label"`
	APIType   		 string    `json:"api_type"`
	Header    		 ParticipationHeader `json:"Participation"`
	APISchema 		 string    `json:"api_schema"`
	Accepter         []string  `json:"accepter"`
	Deleted          bool      `json:"deleted"`
}

type ParticipationHeader struct {
	Participation				*int	`json:"Participation"`
	ParticipationDate			string	`json:"ParticipationDate"`
	ParticipationTime			string	`json:"ParticipationTime"`
	Participator				int		`json:"Participator"`
	ParticipationObjectType		string	`json:"ParticipationObjectType"`
	ParticipationObject			int		`json:"ParticipationObject"`
	Attendance					*int	`json:"Attendance"`
	CreationDate				string	`json:"CreationDate"`
	CreationTime				string	`json:"CreationTime"`
	IsCancelled					*bool	`json:"IsCancelled"`
}

func ParticipationInputRead(
	controller *beego.Controller,
) ParticipationSDC {
	var requestBody ParticipationSDC
	err := json.Unmarshal(controller.Ctx.Input.RequestBody, &requestBody)

	if err != nil {
		controller.Data["json"] = map[string]interface{}{
			"error": fmt.Sprintf("json decode error: %v", err),
		}
		controller.ServeJSON()
	}

	return requestBody
}
