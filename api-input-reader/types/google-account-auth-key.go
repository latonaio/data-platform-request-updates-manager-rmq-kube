package types

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type GoogleAccountAuthKeySDC struct {
	ConnectionKey    	string	`json:"connection_key"`
	Result           	bool	`json:"result"`
	RedisKey         	string	`json:"redis_key"`
	Filepath        	string	`json:"filepath"`
	APIStatusCode  		int		`json:"api_status_code"`
	RuntimeSessionID	string	`json:"runtime_session_id"`
	BusinessPartner  	*int	`json:"business_partner"`
	ServiceLabel     	string	`json:"service_label"`
	APIType				string	`json:"api_type"`
	GoogleAccountAuthKey    GoogleAccountAuthKey    `json:"GoogleAccountAuthKey"`
	APISchema			string	`json:"api_schema"`
	Accepter         	[]string `json:"accepter"`
	Deleted          	bool	`json:"deleted"`
}

type GoogleAccountAuthKey struct {
	URL string `json:"URL"`
}

func GoogleAccountAuthKeyInputRead(
	controller *beego.Controller,
) GoogleAccountAuthKeySDC {
	var requestBody GoogleAccountAuthKeySDC
	err := json.Unmarshal(controller.Ctx.Input.RequestBody, &requestBody)

	if err != nil {
		controller.Data["json"] = map[string]interface{}{
			"error": fmt.Sprintf("json decode error: %v", err),
		}
		controller.ServeJSON()
	}

	return requestBody
}
