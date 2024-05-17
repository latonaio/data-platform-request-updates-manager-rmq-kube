package types

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type PointBalanceSDC struct {
	ConnectionKey    	string	`json:"connection_key"`
	Result           	bool	`json:"result"`
	RedisKey         	string	`json:"redis_key"`
	Filepath        	string	`json:"filepath"`
	APIStatusCode  		int		`json:"api_status_code"`
	RuntimeSessionID	string	`json:"runtime_session_id"`
	BusinessPartner  	*int	`json:"business_partner"`
	ServiceLabel     	string	`json:"service_label"`
	APIType      		string	`json:"api_type"`
	PointBalance 		PointBalancePointBalance `json:"PointBalance"`
	APISchema    		string	`json:"api_schema"`
	Accepter         	[]string `json:"accepter"`
	Deleted          	bool	 `json:"deleted"`
}

type PointBalancePointBalance struct {
	BusinessPartner		int			`json:"BusinessPartner"`
	PointSymbol			string		`json:"PointSymbol"`
	CurrentBalance		float32		`json:"CurrentBalance"`
	LimitBalance		*float32	`json:"LimitBalance"`
	CreationDate		string		`json:"CreationDate"`
	CreationTime		string		`json:"CreationTime"`
	LastChangeDate		string		`json:"LastChangeDate"`
	LastChangeTime		string		`json:"LastChangeTime"`
}

func PointBalanceInputRead(
	controller *beego.Controller,
) PointBalanceSDC {
	var requestBody PointBalanceSDC
	err := json.Unmarshal(controller.Ctx.Input.RequestBody, &requestBody)

	if err != nil {
		controller.Data["json"] = map[string]interface{}{
			"error": fmt.Sprintf("json decode error: %v", err),
		}
		controller.ServeJSON()
	}

	return requestBody
}
