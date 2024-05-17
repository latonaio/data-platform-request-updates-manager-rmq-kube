package types

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type EventSDC struct {
	ConnectionKey    string      `json:"connection_key"`
	Result           bool        `json:"result"`
	RedisKey         string      `json:"redis_key"`
	Filepath         string      `json:"filepath"`
	APIStatusCode    int         `json:"api_status_code"`
	RuntimeSessionID string      `json:"runtime_session_id"`
	BusinessPartner  *int        `json:"business_partner"`
	ServiceLabel     string      `json:"service_label"`
	APIType          string      `json:"api_type"`
	Header           EventHeader `json:"Event"`
	APISchema        string      `json:"api_schema"`
	Accepter         []string    `json:"accepter"`
	Deleted          bool        `json:"deleted"`
}

type EventHeader struct {
	Event                         *int                         `json:"Event"`
	EventType                     string                       `json:"EventType"`
	EventOwner                    int                          `json:"EventOwner"`
	EventOwnerBusinessPartnerRole string                       `json:"EventOwnerBusinessPartnerRole"`
	PersonResponsible             string                       `json:"PersonResponsible"`
	ValidityStartDate             string                       `json:"ValidityStartDate"`
	ValidityStartTime             string                       `json:"ValidityStartTime"`
	ValidityEndDate               string                       `json:"ValidityEndDate"`
	ValidityEndTime               string                       `json:"ValidityEndTime"`
	OperationStartDate			  string					   `json:"OperationStartDate"`
	OperationStartTime			  string					   `json:"OperationStartTime"`
	OperationEndDate			  string					   `json:"OperationEndDate"`
	OperationEndTime			  string					   `json:"OperationEndTime"`
	Description                   string                       `json:"Description"`
	LongText                      string                       `json:"LongText"`
	Introduction                  *string                      `json:"Introduction"`
	Site                          int                          `json:"Site"`
	Project                       *int                         `json:"Project"`
	WBSElement                    *int                         `json:"WBSElement"`
	Tag1                          *string                      `json:"Tag1"`
	Tag2                          *string                      `json:"Tag2"`
	Tag3                          *string                      `json:"Tag3"`
	Tag4                          *string                      `json:"Tag4"`
	DistributionProfile           string                       `json:"DistributionProfile"`
	PointConditionType            string                       `json:"PointConditionType"`
	QuestionnaireType			  *string					   `json:"QuestionnaireType"`
	QuestionnaireTemplate		  *string					   `json:"QuestionnaireTemplate"`
	CreationDate                  string                       `json:"CreationDate"`
	CreationTime                  string                       `json:"CreationTime"`
	LastChangeDate                string                       `json:"LastChangeDate"`
	LastChangeTime                string                       `json:"LastChangeTime"`
	CreateUser					  int	   				       `json:"CreateUser"`
	LastChangeUser				  int	   				       `json:"LastChangeUser"`
	IsReleased                    *bool                        `json:"IsReleased"`
	IsCancelled                   *bool                        `json:"IsCancelled"`
	IsMarkedForDeletion           *bool                        `json:"IsMarkedForDeletion"`
	EventPartner                  []EventPartner               `json:"Partner"`
	EventAddress                  []EventAddress               `json:"Address"`
	EventCampaign                 []EventCampaign              `json:"Campaign"`
	EventGame                     []EventGame                  `json:"Game"`
	EventParticipation	          []EventParticipation         `json:"EventParticipation"`
	EventAttendance		          []EventAttendance            `json:"EventAttendance"`
	EventPointTransaction         []EventPointTransaction      `json:"PointTransaction"`
	EventPointConditionElement    []EventPointConditionElement `json:"PointConditionElement"`
}

type EventPartner struct {
	Event                   int     `json:"Event"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
	EmailAddress            *string `json:"EmailAddress"`
}

type EventAddress struct {
	Event          int      `json:"Event"`
	AddressID      int      `json:"AddressID"`
	PostalCode     string   `json:"PostalCode"`
	LocalSubRegion string   `json:"LocalSubRegion"`
	LocalRegion    string   `json:"LocalRegion"`
	Country        string   `json:"Country"`
	GlobalRegion   string   `json:"GlobalRegion"`
	TimeZone       string   `json:"TimeZone"`
	District       *string  `json:"District"`
	StreetName     *string  `json:"StreetName"`
	CityName       *string  `json:"CityName"`
	Building       *string  `json:"Building"`
	Floor          *int     `json:"Floor"`
	Room           *int     `json:"Room"`
	XCoordinate    *float32 `json:"XCoordinate"`
	YCoordinate    *float32 `json:"YCoordinate"`
	ZCoordinate    *float32 `json:"ZCoordinate"`
	Site           *int     `json:"Site"`
}

type EventCampaign struct {
	Event               int    `json:"Event"`
	Campaign            int    `json:"Campaign"`
	CreationDate        string `json:"CreationDate"`
	LastChangeDate      string `json:"LastChangeDate"`
	IsReleased          *bool  `json:"IsReleased"`
	IsCancelled         *bool  `json:"IsCancelled"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type EventGame struct {
	Event               int    `json:"Event"`
	Game                int    `json:"Game"`
	CreationDate        string `json:"CreationDate"`
	LastChangeDate      string `json:"LastChangeDate"`
	IsReleased          *bool  `json:"IsReleased"`
	IsCancelled         *bool  `json:"IsCancelled"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type EventParticipation struct {
	Event							int		`json:"Event"`
	Participator					int		`json:"Participator"`
	Participation					int		`json:"Participation"`
	CreationDate					string	`json:"CreationDate"`
	CreationTime					string	`json:"CreationTime"`
	LastChangeDate					string	`json:"LastChangeDate"`
	LastChangeTime					string	`json:"LastChangeTime"`
	IsCancelled						*bool	`json:"IsCancelled"`
}

type EventAttendance struct {
	Event							int		`json:"Event"`
	Attender						int		`json:"Attender"`
	Attendance						int		`json:"Attendance"`
	Participation					int		`json:"Participation"`
	CreationDate					string	`json:"CreationDate"`
	CreationTime					string	`json:"CreationTime"`
	LastChangeDate					string	`json:"LastChangeDate"`
	LastChangeTime					string	`json:"LastChangeTime"`
	IsCancelled						*bool	`json:"IsCancelled"`
}

type EventPointTransaction struct {
	Event                          int     `json:"Event"`
	Sender                         int     `json:"Sender"`
	Receiver                       int     `json:"Receiver"`
	PointConditionRecord           int     `json:"PointConditionRecord"`
	PointConditionSequentialNumber int     `json:"PointConditionSequentialNumber"`
	PointTransaction               int     `json:"PointTransaction"`
	PointSymbol                    string  `json:"PointSymbol"`
	PointTransactionType           string  `json:"PointTransactionType"`
	PointConditionType             string  `json:"PointConditionType"`
	PointConditionRateValue        float32 `json:"PointConditionRateValue"`
	PointConditionRatio            float32 `json:"PointConditionRatio"`
	PlusMinus                      string  `json:"PlusMinus"`
	CreationDate                   string  `json:"CreationDate"`
	CreationTime                   string  `json:"CreationTime"`
	LastChangeDate                 string  `json:"LastChangeDate"`
	LastChangeTime                 string  `json:"LastChangeTime"`
	IsCancelled                    *bool   `json:"IsCancelled"`
}

type EventPointConditionElement struct {
	Event                          int     `json:"Event"`
	PointConditionRecord           int     `json:"PointConditionRecord"`
	PointConditionSequentialNumber int     `json:"PointConditionSequentialNumber"`
	PointSymbol                    string  `json:"PointSymbol"`
	Sender                         int     `json:"Sender"`
	PointTransactionType           string  `json:"PointTransactionType"`
	PointConditionType             string  `json:"PointConditionType"`
	PointConditionRateValue        float32 `json:"PointConditionRateValue"`
	PointConditionRatio            float32 `json:"PointConditionRatio"`
	PlusMinus                      string  `json:"PlusMinus"`
	CreationDate                   string  `json:"CreationDate"`
	LastChangeDate                 string  `json:"LastChangeDate"`
	IsReleased                     *bool   `json:"IsReleased"`
	IsCancelled                    *bool   `json:"IsCancelled"`
	IsMarkedForDeletion            *bool   `json:"IsMarkedForDeletion"`
}

func EventInputRead(
	controller *beego.Controller,
) EventSDC {
	var requestBody EventSDC
	err := json.Unmarshal(controller.Ctx.Input.RequestBody, &requestBody)

	if err != nil {
		controller.Data["json"] = map[string]interface{}{
			"error": fmt.Sprintf("json decode error: %v", err),
		}
		controller.ServeJSON()
	}

	return requestBody
}
