package apiModuleRuntimesRequestsQuestionnaire

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
	"time"
)

type QuestionnaireReq struct {
	Header   Header   `json:"Questionnaire"`
	APIType  string   `json:"api_type"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	Questionnaire           *int   `json:"Questionnaire"`
	QuestionnaireOwner      string `json:"QuestionnaireOwner"`
	QuestionnaireType       string `json:"QuestionnaireType"`
	QuestionnaireTemplate   string `json:"QuestionnaireTemplate"`
	QuestionnaireDate       string `json:"QuestionnaireDate"`
	QuestionnaireTime       string `json:"QuestionnaireTime"`
	Respondent              int    `json:"Respondent"`
	QuestionnaireObjectType string `json:"QuestionnaireObjectType"`
	QuestionnaireObject     int    `json:"QuestionnaireObject"`
	CreationDate            string `json:"CreationDate"`
	CreationTime            string `json:"CreationTime"`
	IsMarkedForDeletion     *bool  `json:"IsMarkedForDeletion"`
	Item                    []Item `json:"Item"`
}

type Item struct {
	Questionnaire                  *int    `json:"Questionnaire"`
	QuestionnaireItem              int     `json:"QuestionnaireItem"`
	QuestionnaireItemDescription   string  `json:"QuestionnaireItemDescription"`
	QuestionnaireItemFormType      string  `json:"QuestionnaireItemFormType"`
	QuestionnaireItemReplyType     string  `json:"QuestionnaireItemReplyType"`
	QuestionnaireItemReplyByYesNo  *bool   `json:"QuestionnaireItemReplyByYesNo"`
	QuestionnaireItemReplyByNumber *int    `json:"QuestionnaireItemReplyByNumber"`
	QuestionnaireItemReplyByText   *string `json:"QuestionnaireItemReplyByText"`
	CreationDate                   string  `json:"CreationDate"`
	CreationTime                   string  `json:"CreationTime"`
	IsMarkedForDeletion            *bool   `json:"IsMarkedForDeletion"`
}

func FuncQuestionnaireCreatesRequestAll(
	requestPram *types.Request,
	input QuestionnaireReq,
) QuestionnaireReq {
	req := QuestionnaireReq{
		Header:  input.Header,
		APIType: "creates",
		Accepter: []string{
			"Header",
			"Item",
		},
	}
	return req
}

func QuestionnaireCreatesRequestAll(
	requestPram *types.Request,
	input types.QuestionnaireSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_QUESTIONNAIRE_SRV"
	aPIType := "creates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request QuestionnaireReq

	isMarkedForDeletion := false

	request = FuncQuestionnaireCreatesRequestAll(
		requestPram,
		QuestionnaireReq{
			Header: Header{
				Questionnaire:           nil,
				QuestionnaireOwner:      input.Header.QuestionnaireOwner,
				QuestionnaireType:       "1000",
				QuestionnaireTemplate:   "1000",
				QuestionnaireDate:       formattedDate,
				QuestionnaireTime:       formattedTime,
				Respondent:              input.Header.Respondent,
				QuestionnaireObjectType: "EVENT",
				QuestionnaireObject:     input.Header.QuestionnaireObject,
				CreationDate:            formattedDate,
				CreationTime:            formattedTime,
				IsMarkedForDeletion:     &isMarkedForDeletion,
				Item: []Item{
					{
						QuestionnaireItem:              input.Header.QuestionnaireItem[0].QuestionnaireItem,
						QuestionnaireItemDescription:   input.Header.QuestionnaireItem[0].QuestionnaireItemDescription,
						QuestionnaireItemFormType:      input.Header.QuestionnaireItem[0].QuestionnaireItemFormType,
						QuestionnaireItemReplyType:     input.Header.QuestionnaireItem[0].QuestionnaireItemReplyType,
						QuestionnaireItemReplyByYesNo:  input.Header.QuestionnaireItem[0].QuestionnaireItemReplyByYesNo,
						QuestionnaireItemReplyByNumber: input.Header.QuestionnaireItem[0].QuestionnaireItemReplyByNumber,
						QuestionnaireItemReplyByText:   input.Header.QuestionnaireItem[0].QuestionnaireItemReplyByText,
						CreationDate:                   formattedDate,
						CreationTime:                   formattedTime,
						IsMarkedForDeletion:            &isMarkedForDeletion,
					},
				},
			},
		},
	)

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
	)

	return responseBody
}
