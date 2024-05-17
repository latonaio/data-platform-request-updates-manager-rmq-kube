package types

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type GoogleAccountUserInfoSDC struct {
	ConnectionKey    	string	`json:"connection_key"`
	Result           	bool	`json:"result"`
	RedisKey         	string	`json:"redis_key"`
	Filepath        	string	`json:"filepath"`
	APIStatusCode  		int		`json:"api_status_code"`
	RuntimeSessionID	string	`json:"runtime_session_id"`
	BusinessPartner  	*int	`json:"business_partner"`
	ServiceLabel     	string	`json:"service_label"`
	APIType				string	`json:"api_type"`
	GoogleAccountUserInfo    GoogleAccountUserInfo    `json:"GoogleAccountUserInfo"`
	APISchema			string	`json:"api_schema"`
	Accepter         	[]string `json:"accepter"`
	Deleted          	bool	`json:"deleted"`
}

type GoogleAccountUserInfo struct {
	AccessToken string `json:"AccessToken"`
}

func GoogleAccountUserInfoInputRead(
	controller *beego.Controller,
) GoogleAccountUserInfoSDC {
	var requestBody GoogleAccountUserInfoSDC
	err := json.Unmarshal(controller.Ctx.Input.RequestBody, &requestBody)

	if err != nil {
		controller.Data["json"] = map[string]interface{}{
			"error": fmt.Sprintf("json decode error: %v", err),
		}
		controller.ServeJSON()
	}

	return requestBody
}
