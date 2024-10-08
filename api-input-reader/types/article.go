package types

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type ArticleSDC struct {
	ConnectionKey    string      `json:"connection_key"`
	Result           bool        `json:"result"`
	RedisKey         string      `json:"redis_key"`
	Filepath         string      `json:"filepath"`
	APIStatusCode    int         `json:"api_status_code"`
	RuntimeSessionID string      `json:"runtime_session_id"`
	BusinessPartner  *int        `json:"business_partner"`
	ServiceLabel     string      `json:"service_label"`
	APIType          string      `json:"api_type"`
	Header           ArticleHeader `json:"Article"`
	APISchema        string      `json:"api_schema"`
	Accepter         []string    `json:"accepter"`
	Deleted          bool        `json:"deleted"`
}

type ArticleHeader struct {
	Article                         *int		`json:"Article"`
	ArticleType                     string		`json:"ArticleType"`
	ArticleOwner                    int			`json:"ArticleOwner"`
	ArticleOwnerBusinessPartnerRole string		`json:"ArticleOwnerBusinessPartnerRole"`
	PersonResponsible             	string		`json:"PersonResponsible"`
	ValidityStartDate             	string		`json:"ValidityStartDate"`
	ValidityStartTime             	string		`json:"ValidityStartTime"`
	ValidityEndDate               	string		`json:"ValidityEndDate"`
	ValidityEndTime               	string		`json:"ValidityEndTime"`
	Description                   	string		`json:"Description"`
	LongText                      	string		`json:"LongText"`
	Introduction                  	*string		`json:"Introduction"`
	Site                          	int			`json:"Site"`
	Shop						  	*int		`json:"Shop"`
	Project                       	*int		`json:"Project"`
	WBSElement                    	*int		`json:"WBSElement"`
	Tag1                          	*string		`json:"Tag1"`
	Tag2                          	*string		`json:"Tag2"`
	Tag3                          	*string		`json:"Tag3"`
	Tag4                          	*string		`json:"Tag4"`
	DistributionProfile           	string		`json:"DistributionProfile"`
	QuestionnaireType			  	*string		`json:"QuestionnaireType"`
	QuestionnaireTemplate		  	*string		`json:"QuestionnaireTemplate"`
	CreationDate                  	string		`json:"CreationDate"`
	CreationTime                  	string		`json:"CreationTime"`
	LastChangeDate                	string		`json:"LastChangeDate"`
	LastChangeTime                	string		`json:"LastChangeTime"`
	CreateUser					  	int			`json:"CreateUser"`
	LastChangeUser				  	int			`json:"LastChangeUser"`
	IsReleased                    	*bool		`json:"IsReleased"`
	IsMarkedForDeletion           	*bool		`json:"IsMarkedForDeletion"`
	ArticlePartner                  []ArticlePartner	`json:"Partner"`
	ArticleAddress                  []ArticleAddress	`json:"Address"`
	ArticleCounter                  []ArticleCounter	`json:"Counter"`
	ArticleLike                     []ArticleLike		`json:"Like"`
}

type ArticlePartner struct {
	Article                   int     `json:"Article"`
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

type ArticleAddress struct {
	Article          int      `json:"Article"`
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

type ArticleCounter struct {
	Article					int		`json:"Article"`
	NumberOfLikes			int		`json:"NumberOfLikes"`
	CreationDate			string	`json:"CreationDate"`
	CreationTime			string	`json:"CreationTime"`
	LastChangeDate			string	`json:"LastChangeDate"`
	LastChangeTime			string	`json:"LastChangeTime"`
}

type ArticleLike struct {
	Article					int		`json:"Article"`
	BusinessPartner			int		`json:"BusinessPartner"`
	Like					*bool	`json:"Like"`
	CreationDate			string	`json:"CreationDate"`
	CreationTime			string	`json:"CreationTime"`
	LastChangeDate			string	`json:"LastChangeDate"`
	LastChangeTime			string	`json:"LastChangeTime"`
}

func ArticleInputRead(
	controller *beego.Controller,
) ArticleSDC {
	var requestBody ArticleSDC
	err := json.Unmarshal(controller.Ctx.Input.RequestBody, &requestBody)

	if err != nil {
		controller.Data["json"] = map[string]interface{}{
			"error": fmt.Sprintf("json decode error: %v", err),
		}
		controller.ServeJSON()
	}

	return requestBody
}
