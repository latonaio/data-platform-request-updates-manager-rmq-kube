package types

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type FriendSDC struct {
	ConnectionKey    	string	`json:"connection_key"`
	Result           	bool	`json:"result"`
	RedisKey         	string	`json:"redis_key"`
	Filepath        	string	`json:"filepath"`
	APIStatusCode  		int		`json:"api_status_code"`
	RuntimeSessionID	string	`json:"runtime_session_id"`
	BusinessPartner  	*int	`json:"business_partner"`
	ServiceLabel     	string	`json:"service_label"`
	APIType				string			`json:"api_type"`
	General				FriendGeneral	`json:"Friend"`
	APISchema			string		`json:"api_schema"`
	Accepter         	[]string	`json:"accepter"`
	Deleted          	bool	`json:"deleted"`
}

type FriendGeneral struct {
	BusinessPartner				int		`json:"BusinessPartner"`
	Friend						int		`json:"Friend"`
	BPBusinessPartnerType		string	`json:"BPBusinessPartnerType"`
	FriendBusinessPartnerType	string	`json:"FriendBusinessPartnerType"`
	RankType					string	`json:"RankType"`
	Rank						int		`json:"Rank"`
	FriendIsBlocked				bool	`json:"FriendIsBlocked"`
	CreationDate				string	`json:"CreationDate"`
	CreationTime				string	`json:"CreationTime"`
	LastChangeDate				string	`json:"LastChangeDate"`
	LastChangeTime				string	`json:"LastChangeTime"`
	IsMarkedForDeletion			*bool	`json:"IsMarkedForDeletion"`
}

func FriendInputRead(
	controller *beego.Controller,
) FriendSDC {
	var requestBody FriendSDC
	err := json.Unmarshal(controller.Ctx.Input.RequestBody, &requestBody)

	if err != nil {
		controller.Data["json"] = map[string]interface{}{
			"error": fmt.Sprintf("json decode error: %v", err),
		}
		controller.ServeJSON()
	}

	return requestBody
}
