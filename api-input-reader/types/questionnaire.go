package types

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type QuestionnaireSDC struct {
	ConnectionKey    string                 `json:"connection_key"`
	Result           bool                   `json:"result"`
	RedisKey         string                 `json:"redis_key"`
	Filepath         string                 `json:"filepath"`
	APIStatusCode    int                    `json:"api_status_code"`
	RuntimeSessionID string                 `json:"runtime_session_id"`
	BusinessPartner  *int                   `json:"business_partner"`
	ServiceLabel     string                 `json:"service_label"`
	APIType          string                 `json:"api_type"`
	Header           QuestionnaireHeader    `json:"Questionnaire"`
	APISchema        string                 `json:"api_schema"`
	Accepter         []string               `json:"accepter"`
	Deleted          bool                   `json:"deleted"`
}

type QuestionnaireHeader struct {
	Questionnaire			*int	`json:"Questionnaire"`
	QuestionnaireOwner		string	`json:"QuestionnaireOwner"`
	QuestionnaireType		string	`json:"QuestionnaireType"`
	QuestionnaireTemplate	string	`json:"QuestionnaireTemplate"`
	QuestionnaireDate		string	`json:"QuestionnaireDate"`
	QuestionnaireTime		string	`json:"QuestionnaireTime"`
	Respondent				int		`json:"Respondent"`
	QuestionnaireObjectType	string	`json:"QuestionnaireObjectType"`
	QuestionnaireObject		int		`json:"QuestionnaireObject"`
	CreationDate			string	`json:"CreationDate"`
	CreationTime			string	`json:"CreationTime"`
	IsMarkedForDeletion		*bool	`json:"IsMarkedForDeletion"`
	QuestionnaireItem		[]QuestionnaireItem	`json:"Item"`
}

type QuestionnaireItem struct {
	Questionnaire					*int	`json:"Questionnaire"`
	QuestionnaireItem				int		`json:"QuestionnaireItem"`
	QuestionnaireItemDescription	string	`json:"QuestionnaireItemDescription"`
	QuestionnaireItemFormType		string	`json:"QuestionnaireItemFormType"`
	QuestionnaireItemReplyType		string	`json:"QuestionnaireItemReplyType"`
	QuestionnaireItemReplyByYesNo	*bool	`json:"QuestionnaireItemReplyByYesNo"`
	QuestionnaireItemReplyByNumber	*int	`json:"QuestionnaireItemReplyByNumber"`
	QuestionnaireItemReplyByText	*string	`json:"QuestionnaireItemReplyByText"`
	CreationDate					string	`json:"CreationDate"`
	CreationTime					string	`json:"CreationTime"`
	IsMarkedForDeletion				*bool	`json:"IsMarkedForDeletion"`
}

func QuestionnaireInputRead(
	controller *beego.Controller,
) QuestionnaireSDC {
	var requestBody QuestionnaireSDC
	err := json.Unmarshal(controller.Ctx.Input.RequestBody, &requestBody)

	if err != nil {
		controller.Data["json"] = map[string]interface{}{
			"error": fmt.Sprintf("json decode error: %v", err),
		}
		controller.ServeJSON()
	}

	return requestBody
}
