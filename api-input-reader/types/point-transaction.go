package types

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type PointTransactionSDC struct {
	ConnectionKey    string                 `json:"connection_key"`
	Result           bool                   `json:"result"`
	RedisKey         string                 `json:"redis_key"`
	Filepath         string                 `json:"filepath"`
	APIStatusCode    int                    `json:"api_status_code"`
	RuntimeSessionID string                 `json:"runtime_session_id"`
	BusinessPartner  *int                   `json:"business_partner"`
	ServiceLabel     string                 `json:"service_label"`
	APIType          string                 `json:"api_type"`
	HeaderFromEvent  PointTransactionHeader `json:"PointTransaction"`
	APISchema        string                 `json:"api_schema"`
	Accepter         []string               `json:"accepter"`
	Deleted          bool                   `json:"deleted"`
}

type PointTransactionHeader struct {
	PointTransaction                      *int    `json:"PointTransaction"`
	PointTransactionType                  string  `json:"PointTransactionType"`
	PointTransactionDate                  string  `json:"PointTransactionDate"`
	PointTransactionTime                  string  `json:"PointTransactionTime"`
	Sender                                int     `json:"Sender"`
	Receiver                              int     `json:"Receiver"`
	PointSymbol                           string  `json:"PointSymbol"`
	PlusMinus                             string  `json:"PlusMinus"`
	PointTransactionAmount                float32 `json:"PointTransactionAmount"`
	PointTransactionObjectType            string  `json:"PointTransactionObjectType"`
	PointTransactionObject                int     `json:"PointTransactionObject"`
	SenderPointBalanceBeforeTransaction   float32 `json:"SenderPointBalanceBeforeTransaction"`
	SenderPointBalanceAfterTransaction    float32 `json:"SenderPointBalanceAfterTransaction"`
	ReceiverPointBalanceBeforeTransaction float32 `json:"ReceiverPointBalanceBeforeTransaction"`
	ReceiverPointBalanceAfterTransaction  float32 `json:"ReceiverPointBalanceAfterTransaction"`
	Attendance							  *int	  `json:"Attendance"`
	Participation						  *int	  `json:"Participation"`
	CreationDate                          string  `json:"CreationDate"`
	CreationTime                          string  `json:"CreationTime"`
	IsCancelled                           *bool   `json:"IsCancelled"`
	BusinessPartner                       int     `json:"BusinessPartner"`
	Event                                 *int    `json:"Event"`
	EventOwner                            *int    `json:"EventOwner"`
	PointConditionRateValue               float32 `json:"PointConditionRateValue"`
}

func PointTransactionInputRead(
	controller *beego.Controller,
) PointTransactionSDC {
	var requestBody PointTransactionSDC
	err := json.Unmarshal(controller.Ctx.Input.RequestBody, &requestBody)

	if err != nil {
		controller.Data["json"] = map[string]interface{}{
			"error": fmt.Sprintf("json decode error: %v", err),
		}
		controller.ServeJSON()
	}

	return requestBody
}
