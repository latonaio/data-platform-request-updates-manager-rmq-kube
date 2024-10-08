package types

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type MessageSDC struct {
	ConnectionKey    string        `json:"connection_key"`
	Result           bool          `json:"result"`
	RedisKey         string        `json:"redis_key"`
	Filepath         string        `json:"filepath"`
	APIStatusCode    int           `json:"api_status_code"`
	RuntimeSessionID string        `json:"runtime_session_id"`
	BusinessPartner  *int          `json:"business_partner"`
	ServiceLabel     string        `json:"service_label"`
	APIType          string        `json:"api_type"`
	Header           MessageHeader `json:"Message"`
	APISchema        string        `json:"api_schema"`
	Accepter         []string      `json:"accepter"`
	Deleted          bool          `json:"deleted"`
}

type MessageHeader struct {
	Message             int    `json:"Message"`
	MessageType         string `json:"MessageType"`
	Sender              int    `json:"Sender"`
	Receiver            int    `json:"Receiver"`
	Language            string `json:"Language"`
	Title               string `json:"Title"`
	LongText            string `json:"LongText"`
	MessageIsSent       bool   `json:"MessageIsSent"`
	MessageIsRead       bool   `json:"MessageIsRead"`
	CreationDate        string `json:"CreationDate"`
	CreationTime        string `json:"CreationTime"`
	LastChangeDate      string `json:"LastChangeDate"`
	LastChangeTime      string `json:"LastChangeTime"`
	IsCancelled         *bool  `json:"IsCancelled"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

func MessageInputRead(
	controller *beego.Controller,
) MessageSDC {
	var requestBody MessageSDC
	err := json.Unmarshal(controller.Ctx.Input.RequestBody, &requestBody)

	if err != nil {
		controller.Data["json"] = map[string]interface{}{
			"error": fmt.Sprintf("json decode error: %v", err),
		}
		controller.ServeJSON()
	}

	return requestBody
}
