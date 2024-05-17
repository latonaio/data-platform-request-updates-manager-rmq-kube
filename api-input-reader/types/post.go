package types

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type PostSDC struct {
	ConnectionKey    string    `json:"connection_key"`
	Result           bool      `json:"result"`
	RedisKey         string    `json:"redis_key"`
	Filepath         string    `json:"filepath"`
	APIStatusCode    int       `json:"api_status_code"`
	RuntimeSessionID string    `json:"runtime_session_id"`
	BusinessPartner  *int      `json:"business_partner"`
	ServiceLabel     string    `json:"service_label"`
	APIType   		 string    `json:"api_type"`
	Header    		 PostHeader `json:"Post"`
	APISchema 		 string    `json:"api_schema"`
	Accepter         []string  `json:"accepter"`
	Deleted          bool      `json:"deleted"`
}

type PostHeader struct {
	Post							*int	`json:"Post"`
	PostType						string	`json:"PostType"`
	PostOwner						int		`json:"PostOwner"`
	PostOwnerBusinessPartnerRole	string	`json:"PostOwnerBusinessPartnerRole"`
	Description						string	`json:"Description"`
	LongText						string	`json:"LongText"`
	Tag1							*string	`json:"Tag1"`
	Tag2							*string	`json:"Tag2"`
	Tag3							*string	`json:"Tag3"`
	Tag4							*string	`json:"Tag4"`
	CreationDate					string	`json:"CreationDate"`
	CreationTime					string	`json:"CreationTime"`
	LastChangeDate					string	`json:"LastChangeDate"`
	LastChangeTime					string	`json:"LastChangeTime"`
	IsMarkedForDeletion				*bool	`json:"IsMarkedForDeletion"`
}

func PostInputRead(
	controller *beego.Controller,
) PostSDC {
	var requestBody PostSDC
	err := json.Unmarshal(controller.Ctx.Input.RequestBody, &requestBody)

	if err != nil {
		controller.Data["json"] = map[string]interface{}{
			"error": fmt.Sprintf("json decode error: %v", err),
		}
		controller.ServeJSON()
	}

	return requestBody
}
