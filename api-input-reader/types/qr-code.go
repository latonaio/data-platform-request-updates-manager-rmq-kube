package types

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type QRCodeSDC struct {
	ConnectionKey       string   `json:"connection_key"`
	Result              bool     `json:"result"`
	RedisKey            string   `json:"redis_key"`
	Filepath            string   `json:"filepath"`
	APIStatusCode       int      `json:"api_status_code"`
	RuntimeSessionID    string   `json:"runtime_session_id"`
	BusinessPartnerID   *int     `json:"business_partner"`
	ServiceLabel        string   `json:"service_label"`
	QRCodeGlobal        QRCode   `json:"QRCode"`
	APISchema           string   `json:"api_schema"`
	Accepter            []string `json:"accepter"`
	Deleted             bool     `json:"deleted"`
	DocData             string   `json:"doc_data"`
	APIProcessingResult *bool    `json:"api_processing_result"`
	APIProcessingError  string   `json:"api_processing_error"`
}

type QRCode struct {
	APIServiceName            string `json:"APIServiceName"`
	ServiceLabel              string `json:"ServiceLabel"`
	APIType                   string `json:"APIType"`
	URLFormatBeforeConversion string `json:"URLFormatBeforeConversion"`
	BusinessPartner           *int   `json:"BusinessPartner"`
	Event                     *int   `json:"Event"`
	Article                   *int   `json:"Article"`
	Shop                      *int   `json:"Shop"`
	Site                      *int   `json:"Site"`
	Station                   *int   `json:"Station"`
	BusStop                   *int   `json:"BusStop"`
	Participation             *int   `json:"Participation"`
}

func QRCodeInputRead(
	controller *beego.Controller,
) QRCodeSDC {
	var requestBody QRCodeSDC
	err := json.Unmarshal(controller.Ctx.Input.RequestBody, &requestBody)

	if err != nil {
		controller.Data["json"] = map[string]interface{}{
			"error": fmt.Sprintf("json decode error: %v", err),
		}
		controller.ServeJSON()
	}

	return requestBody
}
