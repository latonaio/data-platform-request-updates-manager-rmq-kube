package apiModuleRuntimesRequestsEvent

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
	"time"
)

type EventReq struct {
	Header   Header   `json:"Event"`
	APIType  string   `json:"api_type"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	Event                         *int                    `json:"Event"`
	EventType                     string                  `json:"EventType"`
	EventOwner                    int                     `json:"EventOwner"`
	EventOwnerBusinessPartnerRole string                  `json:"EventOwnerBusinessPartnerRole"`
	PersonResponsible             string                  `json:"PersonResponsible"`
	URL                           *string                 `json:"URL"`
	ValidityStartDate             string                  `json:"ValidityStartDate"`
	ValidityStartTime             string                  `json:"ValidityStartTime"`
	ValidityEndDate               string                  `json:"ValidityEndDate"`
	ValidityEndTime               string                  `json:"ValidityEndTime"`
	OperationStartDate            string                  `json:"OperationStartDate"`
	OperationStartTime            string                  `json:"OperationStartTime"`
	OperationEndDate              string                  `json:"OperationEndDate"`
	OperationEndTime              string                  `json:"OperationEndTime"`
	Description                   string                  `json:"Description"`
	LongText                      string                  `json:"LongText"`
	Introduction                  *string                 `json:"Introduction"`
	Site                          int                     `json:"Site"`
	Capacity                      int                     `json:"Capacity"`
	Shop                          *int                    `json:"Shop"`
	Project                       *int                    `json:"Project"`
	WBSElement                    *int                    `json:"WBSElement"`
	Tag1                          *string                 `json:"Tag1"`
	Tag2                          *string                 `json:"Tag2"`
	Tag3                          *string                 `json:"Tag3"`
	Tag4                          *string                 `json:"Tag4"`
	DistributionProfile           string                  `json:"DistributionProfile"`
	PointConditionType            string                  `json:"PointConditionType"`
	QuestionnaireType             *string                 `json:"QuestionnaireType"`
	QuestionnaireTemplate         *string                 `json:"QuestionnaireTemplate"`
	CreationDate                  string                  `json:"CreationDate"`
	CreationTime                  string                  `json:"CreationTime"`
	LastChangeDate                string                  `json:"LastChangeDate"`
	LastChangeTime                string                  `json:"LastChangeTime"`
	CreateUser                    int                     `json:"CreateUser"`
	LastChangeUser                int                     `json:"LastChangeUser"`
	IsReleased                    *bool                   `json:"IsReleased"`
	IsCancelled                   *bool                   `json:"IsCancelled"`
	IsMarkedForDeletion           *bool                   `json:"IsMarkedForDeletion"`
	Partner                       []Partner               `json:"Partner"`
	Address                       []Address               `json:"Address"`
	Campaign                      []Campaign              `json:"Campaign"`
	Game                          []Game                  `json:"Game"`
	Participation                 []Participation         `json:"Participation"`
	Attendance                    []Attendance            `json:"Attendance"`
	PointTransaction              []PointTransaction      `json:"PointTransaction"`
	PointConditionElement         []PointConditionElement `json:"PointConditionElement"`
	Counter                       []Counter               `json:"Counter"`
	Like                          []Like                  `json:"Like"`
}

type Partner struct {
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

type Address struct {
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

type Campaign struct {
	Event               int    `json:"Event"`
	Campaign            int    `json:"Campaign"`
	CreationDate        string `json:"CreationDate"`
	LastChangeDate      string `json:"LastChangeDate"`
	IsReleased          *bool  `json:"IsReleased"`
	IsCancelled         *bool  `json:"IsCancelled"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type Game struct {
	Event               int    `json:"Event"`
	Game                int    `json:"Game"`
	CreationDate        string `json:"CreationDate"`
	LastChangeDate      string `json:"LastChangeDate"`
	IsReleased          *bool  `json:"IsReleased"`
	IsCancelled         *bool  `json:"IsCancelled"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type Participation struct {
	Event          int    `json:"Event"`
	Participator   int    `json:"Participator"`
	Participation  int    `json:"Participation"`
	CreationDate   string `json:"CreationDate"`
	CreationTime   string `json:"CreationTime"`
	LastChangeDate string `json:"LastChangeDate"`
	LastChangeTime string `json:"LastChangeTime"`
	IsCancelled    *bool  `json:"IsCancelled"`
}

type Attendance struct {
	Event          int    `json:"Event"`
	Attender       int    `json:"Attender"`
	Attendance     int    `json:"Attendance"`
	Participation  int    `json:"Participation"`
	CreationDate   string `json:"CreationDate"`
	CreationTime   string `json:"CreationTime"`
	LastChangeDate string `json:"LastChangeDate"`
	LastChangeTime string `json:"LastChangeTime"`
	IsCancelled    *bool  `json:"IsCancelled"`
}

type PointTransaction struct {
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

type PointConditionElement struct {
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

type Counter struct {
	Event                  int    `json:"Event"`
	NumberOfLikes          int    `json:"NumberOfLikes"`
	NumberOfParticipations int    `json:"NumberOfParticipations"`
	NumberOfAttendances    int    `json:"NumberOfAttendances"`
	CreationDate           string `json:"CreationDate"`
	CreationTime           string `json:"CreationTime"`
	LastChangeDate         string `json:"LastChangeDate"`
	LastChangeTime         string `json:"LastChangeTime"`
}

type Like struct {
	Event           int    `json:"Event"`
	BusinessPartner int    `json:"BusinessPartner"`
	Like            *bool  `json:"Like"`
	CreationDate    string `json:"CreationDate"`
	CreationTime    string `json:"CreationTime"`
	LastChangeDate  string `json:"LastChangeDate"`
	LastChangeTime  string `json:"LastChangeTime"`
}

func FuncEventCreatesRequestAll(
	requestPram *types.Request,
	input EventReq,
) EventReq {
	req := EventReq{
		Header:  input.Header,
		APIType: "creates",
		Accepter: []string{
			"Header",
			"Address",
			"Counter",
			"Like",
			//			"PointConditionElement",
		},
	}
	return req
}

func EventCreatesRequestAll(
	requestPram *types.Request,
	input types.EventSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_EVENT_SRV"
	aPIType := "creates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request EventReq

	isReleased := false
	isCancelled := false
	isMarkedForDeletion := false

	//	introduction := "Test Introduction"

	like := false

	shop := 1

	request = FuncEventCreatesRequestAll(
		requestPram,
		EventReq{
			Header: Header{
				Event:                         nil,
				EventType:                     input.Header.EventType,
				EventOwner:                    input.Header.EventOwner,
				EventOwnerBusinessPartnerRole: input.Header.EventOwnerBusinessPartnerRole,
				PersonResponsible:             input.Header.PersonResponsible,
				URL:                           input.Header.URL,
				ValidityStartDate:             input.Header.ValidityStartDate,
				ValidityStartTime:             input.Header.ValidityStartTime,
				ValidityEndDate:               input.Header.ValidityEndDate,
				ValidityEndTime:               input.Header.ValidityEndTime,
				OperationStartDate:            input.Header.OperationStartDate,
				OperationStartTime:            input.Header.OperationStartTime,
				OperationEndDate:              input.Header.OperationEndDate,
				OperationEndTime:              input.Header.OperationEndTime,
				Description:                   input.Header.Description,
				LongText:                      input.Header.LongText,
				Introduction:                  &input.Header.LongText,
				Site:                          input.Header.Site,
				Capacity:                      300,
				Shop:                          &shop,
				Tag1:                          input.Header.Tag1,
				Tag2:                          input.Header.Tag2,
				Tag3:                          input.Header.Tag3,
				Tag4:                          input.Header.Tag4,
				DistributionProfile:           input.Header.DistributionProfile,
				PointConditionType:            input.Header.PointConditionType,
				QuestionnaireType:             input.Header.QuestionnaireType,
				QuestionnaireTemplate:         input.Header.QuestionnaireTemplate,
				CreationDate:                  formattedDate,
				CreationTime:                  formattedTime,
				LastChangeDate:                formattedDate,
				LastChangeTime:                formattedTime,
				CreateUser:                    input.Header.CreateUser,
				LastChangeUser:                input.Header.LastChangeUser,
				IsReleased:                    &isReleased,
				IsCancelled:                   &isCancelled,
				IsMarkedForDeletion:           &isMarkedForDeletion,
				Address: []Address{
					{
						AddressID:      input.Header.EventAddress[0].AddressID,
						PostalCode:     input.Header.EventAddress[0].PostalCode,
						LocalSubRegion: input.Header.EventAddress[0].LocalSubRegion,
						LocalRegion:    input.Header.EventAddress[0].LocalRegion,
						Country:        input.Header.EventAddress[0].Country,
						GlobalRegion:   input.Header.EventAddress[0].GlobalRegion,
						TimeZone:       input.Header.EventAddress[0].TimeZone,
						District:       input.Header.EventAddress[0].District,
						StreetName:     input.Header.EventAddress[0].StreetName,
						CityName:       input.Header.EventAddress[0].CityName,
						Building:       input.Header.EventAddress[0].Building,
						XCoordinate:    input.Header.EventAddress[0].XCoordinate,
						YCoordinate:    input.Header.EventAddress[0].YCoordinate,
						ZCoordinate:    input.Header.EventAddress[0].ZCoordinate,
					},
				},
				Counter: []Counter{
					{
						NumberOfLikes:          0,
						NumberOfParticipations: 0,
						NumberOfAttendances:    0,
						CreationDate:           formattedDate,
						CreationTime:           formattedTime,
						LastChangeDate:         formattedDate,
						LastChangeTime:         formattedTime,
					},
				},
				Like: []Like{
					{
						BusinessPartner: input.Header.EventLike[0].BusinessPartner,
						Like:            &like,
						CreationDate:    formattedDate,
						CreationTime:    formattedTime,
						LastChangeDate:  formattedDate,
						LastChangeTime:  formattedTime,
					},
				},
				//				PointConditionElement: []PointConditionElement{
				//					{
				//						PointConditionRecord:           1,
				//						PointConditionSequentialNumber: 1,
				//						PointSymbol:                    "POYPO",
				//						Sender:                         input.Header.EventPointConditionElement[0].Sender,
				//						PointTransactionType:           input.Header.EventPointConditionElement[0].PointTransactionType,
				//						PointConditionType:             input.Header.EventPointConditionElement[0].PointConditionType,
				//						PointConditionRateValue:        input.Header.EventPointConditionElement[0].PointConditionRateValue,
				//						PointConditionRatio:            input.Header.EventPointConditionElement[0].PointConditionRatio,
				//						PlusMinus:                      input.Header.EventPointConditionElement[0].PlusMinus,
				//						CreationDate:                   formattedDate,
				//						LastChangeDate:                 formattedDate,
				//						IsReleased:                     &isReleased,
				//						IsCancelled:                    &isCancelled,
				//						IsMarkedForDeletion:            &isMarkedForDeletion,
				//					},
				//				},
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

func FuncEventCreatesRequestParticipation(
	requestPram *types.Request,
	input EventReq,
) EventReq {
	req := EventReq{
		Header:  input.Header,
		APIType: "creates",
		Accepter: []string{
			"Participation",
		},
	}
	return req
}

func EventCreatesRequestParticipation(
	requestPram *types.Request,
	input types.EventSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_EVENT_SRV"
	aPIType := "creates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request EventReq

	isCancelled := false

	request = FuncEventCreatesRequestParticipation(
		requestPram,
		EventReq{
			Header: Header{
				Event: input.Header.Event,
				Participation: []Participation{
					{
						Participator:   input.Header.EventParticipation[0].Participator,
						Participation:  input.Header.EventParticipation[0].Participation,
						CreationDate:   formattedDate,
						CreationTime:   formattedTime,
						LastChangeDate: formattedDate,
						LastChangeTime: formattedTime,
						IsCancelled:    &isCancelled,
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

func FuncEventCreatesRequestAttendance(
	requestPram *types.Request,
	input EventReq,
) EventReq {
	req := EventReq{
		Header:  input.Header,
		APIType: "creates",
		Accepter: []string{
			"Attendance",
		},
	}
	return req
}

func EventCreatesRequestAttendance(
	requestPram *types.Request,
	input types.EventSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_EVENT_SRV"
	aPIType := "creates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request EventReq

	isCancelled := false

	request = FuncEventCreatesRequestAttendance(
		requestPram,
		EventReq{
			Header: Header{
				Event: input.Header.Event,
				Attendance: []Attendance{
					{
						Attender:       input.Header.EventAttendance[0].Attender,
						Attendance:     input.Header.EventAttendance[0].Attendance,
						Participation:  1,
						CreationDate:   formattedDate,
						CreationTime:   formattedTime,
						LastChangeDate: formattedDate,
						LastChangeTime: formattedTime,
						IsCancelled:    &isCancelled,
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

func FuncEventCreatesRequestPointTransaction(
	requestPram *types.Request,
	input EventReq,
) EventReq {
	req := EventReq{
		Header:  input.Header,
		APIType: "creates",
		Accepter: []string{
			"PointTransaction",
		},
	}
	return req
}

func EventCreatesRequestPointTransaction(
	requestPram *types.Request,
	input types.EventSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_EVENT_SRV"
	aPIType := "creates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request EventReq

	isCancelled := false

	request = FuncEventCreatesRequestPointTransaction(
		requestPram,
		EventReq{
			Header: Header{
				Event: input.Header.Event,
				PointTransaction: []PointTransaction{
					{
						Sender:                         input.Header.EventPointTransaction[0].Sender,
						Receiver:                       input.Header.EventPointTransaction[0].Receiver,
						PointConditionRecord:           1,
						PointConditionSequentialNumber: 1,
						PointTransaction:               input.Header.EventPointTransaction[0].PointTransaction,
						PointSymbol:                    "POYPO",
						PointTransactionType:           input.Header.EventPointTransaction[0].PointTransactionType,
						PointConditionType:             input.Header.EventPointTransaction[0].PointConditionType,
						PointConditionRateValue:        input.Header.EventPointTransaction[0].PointConditionRateValue,
						PointConditionRatio:            input.Header.EventPointTransaction[0].PointConditionRatio,
						PlusMinus:                      input.Header.EventPointTransaction[0].PlusMinus,
						CreationDate:                   formattedDate,
						CreationTime:                   formattedTime,
						LastChangeDate:                 formattedDate,
						LastChangeTime:                 formattedTime,
						IsCancelled:                    &isCancelled,
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

func FuncEventUpdatesRequestHeader(
	requestPram *types.Request,
	input EventReq,
) EventReq {
	req := EventReq{
		Header:  input.Header,
		APIType: "updates",
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func EventUpdatesRequestHeader(
	requestPram *types.Request,
	input types.EventSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_EVENT_SRV"
	aPIType := "updates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request EventReq

	request = FuncEventUpdatesRequestHeader(
		requestPram,
		EventReq{
			Header: Header{
				Event:                         input.Header.Event,
				EventType:                     input.Header.EventType,
				EventOwner:                    input.Header.EventOwner,
				EventOwnerBusinessPartnerRole: input.Header.EventOwnerBusinessPartnerRole,
				PersonResponsible:             input.Header.PersonResponsible,
				URL:                           input.Header.URL,
				ValidityStartDate:             input.Header.ValidityStartDate,
				ValidityStartTime:             input.Header.ValidityStartTime,
				ValidityEndDate:               input.Header.ValidityEndDate,
				ValidityEndTime:               input.Header.ValidityEndTime,
				OperationStartDate:            input.Header.OperationStartDate,
				OperationStartTime:            input.Header.OperationStartTime,
				OperationEndDate:              input.Header.OperationEndDate,
				OperationEndTime:              input.Header.OperationEndTime,
				Description:                   input.Header.Description,
				LongText:                      input.Header.LongText,
				Introduction:                  input.Header.Introduction,
				Site:                          input.Header.Site,
				Capacity:                      input.Header.Capacity,
				Shop:                          input.Header.Shop,
				Tag1:                          input.Header.Tag1,
				Tag2:                          input.Header.Tag2,
				Tag3:                          input.Header.Tag3,
				Tag4:                          input.Header.Tag4,
				DistributionProfile:           input.Header.DistributionProfile,
				PointConditionType:            input.Header.PointConditionType,
				QuestionnaireType:             input.Header.QuestionnaireType,
				QuestionnaireTemplate:         input.Header.QuestionnaireTemplate,
				LastChangeDate:                formattedDate,
				LastChangeTime:                formattedTime,
				LastChangeUser:                input.Header.LastChangeUser,
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

func FuncEventUpdatesRequestCounter(
	requestPram *types.Request,
	input EventReq,
) EventReq {
	req := EventReq{
		Header:  input.Header,
		APIType: "updates",
		Accepter: []string{
			"Counter",
		},
	}
	return req
}

func EventUpdatesRequestCounter(
	requestPram *types.Request,
	input types.EventSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_EVENT_SRV"
	aPIType := "updates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request EventReq

	request = FuncEventUpdatesRequestCounter(
		requestPram,
		EventReq{
			Header: Header{
				Event: input.Header.Event,
				Counter: []Counter{
					{
						NumberOfLikes:          input.Header.EventCounter[0].NumberOfLikes,
						NumberOfParticipations: input.Header.EventCounter[0].NumberOfParticipations,
						NumberOfAttendances:    input.Header.EventCounter[0].NumberOfAttendances,
						LastChangeDate:         formattedDate,
						LastChangeTime:         formattedTime,
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

func FuncEventUpdatesRequestLike(
	requestPram *types.Request,
	input EventReq,
) EventReq {
	req := EventReq{
		Header:  input.Header,
		APIType: "updates",
		Accepter: []string{
			"Like",
		},
	}
	return req
}

func EventUpdatesRequestLike(
	requestPram *types.Request,
	input types.EventSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_EVENT_SRV"
	aPIType := "updates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request EventReq

	request = FuncEventUpdatesRequestLike(
		requestPram,
		EventReq{
			Header: Header{
				Event: input.Header.Event,
				Like: []Like{
					{
						BusinessPartner: input.Header.EventLike[0].BusinessPartner,
						Like:            input.Header.EventLike[0].Like,
						LastChangeDate:  formattedDate,
						LastChangeTime:  formattedTime,
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

func FuncEventCancelsRequestHeader(
	requestPram *types.Request,
	input EventReq,
) EventReq {
	req := EventReq{
		Header:  input.Header,
		APIType: "cancels",
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func EventCancelsRequestHeader(
	requestPram *types.Request,
	input types.EventSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_EVENT_SRV"
	aPIType := "cancels"

	//	currentDateTime := time.Now()
	//	formattedDate := currentDateTime.Format("2006-01-02")
	//	formattedTime := currentDateTime.Format("15:04:05")

	var request EventReq

	isCancelled := true

	request = FuncEventCancelsRequestHeader(
		requestPram,
		EventReq{
			Header: Header{
				Event:       input.Header.Event,
				IsCancelled: &isCancelled,
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

func FuncEventCancelsRequestParticipation(
	requestPram *types.Request,
	input EventReq,
) EventReq {
	req := EventReq{
		Header:  input.Header,
		APIType: "cancels",
		Accepter: []string{
			"Participation",
		},
	}
	return req
}

func EventCancelsRequestParticipation(
	requestPram *types.Request,
	input types.EventSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_EVENT_SRV"
	aPIType := "cancels"

	//	currentDateTime := time.Now()
	//	formattedDate := currentDateTime.Format("2006-01-02")
	//	formattedTime := currentDateTime.Format("15:04:05")

	var request EventReq

	isCancelled := true

	request = FuncEventCancelsRequestParticipation(
		requestPram,
		EventReq{
			Header: Header{
				Event: input.Header.Event,
				Participation: []Participation{
					{
						Participator: input.Header.EventParticipation[0].Participator,
						IsCancelled:  &isCancelled,
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
