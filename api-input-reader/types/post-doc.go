package types

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type PostDocSDC struct {
	ConnectionKey       string    `json:"connection_key"`
	Result              bool      `json:"result"`
	RedisKey            string    `json:"redis_key"`
	Filepath            string    `json:"filepath"`
	APIStatusCode       int       `json:"api_status_code"`
	RuntimeSessionID    string    `json:"runtime_session_id"`
	BusinessPartnerID   *int      `json:"business_partner"`
	ServiceLabel        string    `json:"service_label"`
	Post                PostDoc   `json:"Post"`
	Posts               []PostDoc `json:"Posts"`
	APISchema           string    `json:"api_schema"`
	Accepter            []string  `json:"accepter"`
	Deleted             bool      `json:"deleted"`
	DocData             string    `json:"doc_data"`
	APIProcessingResult *bool     `json:"api_processing_result"`
	APIProcessingError  string    `json:"api_processing_error"`
}

type PostDoc struct {
	Post      int           `json:"Post"`
	HeaderDoc PostHeaderDoc `json:"HeaderDoc"`
}

type PostHeaderDoc struct {
	Post                     int     `json:"Post"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	FileExtension            string  `json:"FileExtension"`
	FileName                 string  `json:"FileName"`
	FilePath                 string  `json:"FilePath"`
	DocIssuerBusinessPartner int     `json:"DocIssuerBusinessPartner"`
}

func PostDocInputRead(
	controller *beego.Controller,
) PostDocSDC {
	var requestBody PostDocSDC
	err := json.Unmarshal(controller.Ctx.Input.RequestBody, &requestBody)

	if err != nil {
		controller.Data["json"] = map[string]interface{}{
			"error": fmt.Sprintf("json decode error: %v", err),
		}
		controller.ServeJSON()
	}

	return requestBody
}
