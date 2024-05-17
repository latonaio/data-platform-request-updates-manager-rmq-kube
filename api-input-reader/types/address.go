package types

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type AddressSDC struct {
	ConnectionKey    	string	`json:"connection_key"`
	Result           	bool	`json:"result"`
	RedisKey         	string	`json:"redis_key"`
	Filepath        	string	`json:"filepath"`
	APIStatusCode  		int		`json:"api_status_code"`
	RuntimeSessionID	string	`json:"runtime_session_id"`
	BusinessPartner  	*int	`json:"business_partner"`
	ServiceLabel     	string	`json:"service_label"`
	APIType      		string	`json:"api_type"`
	Address 		    AddressAddress `json:"Address"`
	APISchema    		string	`json:"api_schema"`
	Accepter         	[]string `json:"accepter"`
	Deleted          	bool	 `json:"deleted"`
}

type AddressAddress struct {
	AddressID			*int    `json:"AddressID"`
	ValidityStartDate	string  `json:"ValidityStartDate"`
	ValidityEndDate		string	`json:"ValidityEndDate"`
	PostalCode			*string	`json:"PostalCode"`
	LocalSubRegion		string	`json:"LocalSubRegion"`
	LocalRegion			string	`json:"LocalRegion"`
	Country				string	`json:"Country"`
	GlobalRegion		string	`json:"GlobalRegion"`
	TimeZone			string	`json:"TimeZone"`
	District			*string	`json:"District"`
	StreetName			*string	`json:"StreetName"`
	CityName			*string	`json:"CityName"`
	Building			*string	`json:"Building"`
	Floor				*int    `json:"Floor"`
	Room				*int    `json:"Room"`
	XCoordinate 		*float32 `json:"XCoordinate"`
	YCoordinate 		*float32 `json:"YCoordinate"`
	ZCoordinate 		*float32 `json:"ZCoordinate"`
	Site				*int	 `json:"Site"`
	CreationDate        string  `json:"CreationDate"`
	CreationTime        string  `json:"CreationTime"`
	LastChangeDate      string  `json:"LastChangeDate"`
	LastChangeTime      string  `json:"LastChangeTime"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

func AddressInputRead(
	controller *beego.Controller,
) AddressSDC {
	var requestBody AddressSDC
	err := json.Unmarshal(controller.Ctx.Input.RequestBody, &requestBody)

	if err != nil {
		controller.Data["json"] = map[string]interface{}{
			"error": fmt.Sprintf("json decode error: %v", err),
		}
		controller.ServeJSON()
	}

	return requestBody
}
