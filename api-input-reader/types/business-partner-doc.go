package types

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type BusinessPartnerDocSDC struct {
	ConnectionKey       string   `json:"connection_key"`
	Result              bool     `json:"result"`
	RedisKey            string   `json:"redis_key"`
	Filepath            string   `json:"filepath"`
	APIStatusCode       int      `json:"api_status_code"`
	RuntimeSessionID    string   `json:"runtime_session_id"`
	BusinessPartnerID   *int     `json:"business_partner"`
	ServiceLabel        string   `json:"service_label"`
	BusinessPartner     BusinessPartnerDoc  `json:"BusinessPartner"`
	APISchema           string   `json:"api_schema"`
	Accepter            []string `json:"accepter"`
	Deleted             bool     `json:"deleted"`
	DocData             string   `json:"doc_data"`
	APIProcessingResult *bool    `json:"api_processing_result"`
	APIProcessingError  string   `json:"api_processing_error"`
}

type BusinessPartnerDoc struct {
	BusinessPartner      int                        `json:"BusinessPartner"`
	GeneralDoc           BusinessPartnerGeneralDoc  `json:"GeneralDoc"`
}

type BusinessPartnerGeneralDoc struct {
	BusinessPartner          int     `json:"BusinessPartner"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	FileExtension            string  `json:"FileExtension"`
	FileName                 string  `json:"FileName"`
	FilePath                 string  `json:"FilePath"`
	DocIssuerBusinessPartner int     `json:"DocIssuerBusinessPartner"`
}

func BusinessPartnerDocInputRead(
	controller *beego.Controller,
) BusinessPartnerDocSDC {
	var requestBody BusinessPartnerDocSDC
	err := json.Unmarshal(controller.Ctx.Input.RequestBody, &requestBody)

	if err != nil {
		controller.Data["json"] = map[string]interface{}{
			"error": fmt.Sprintf("json decode error: %v", err),
		}
		controller.ServeJSON()
	}

	return requestBody
}
