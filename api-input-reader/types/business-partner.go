package types

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type BusinessPartnerSDC struct {
	ConnectionKey		string	`json:"connection_key"`
	Result				bool	`json:"result"`
	RedisKey			string	`json:"redis_key"`
	Filepath			string	`json:"filepath"`
	APIStatusCode		int		`json:"api_status_code"`
	RuntimeSessionID	string	`json:"runtime_session_id"`
	BusinessPartner		*int	`json:"business_partner"`
	ServiceLabel		string	`json:"service_label"`
	APIType				string	`json:"api_type"`
	General   			BusinessPartnerGeneral `json:"BusinessPartner"`
	APISchema			string	`json:"api_schema"`
	Accepter			[]string	`json:"accepter"`
	Deleted				bool	`json:"deleted"`
}

type BusinessPartnerGeneral struct {
	BusinessPartner                        *int                                     `json:"BusinessPartner"`
	BusinessPartnerType                    string                                   `json:"BusinessPartnerType"`
	BusinessPartnerFullName                *string                                  `json:"BusinessPartnerFullName"`
	BusinessPartnerName                    string                                   `json:"BusinessPartnerName"`
	Industry                               *string                                  `json:"Industry"`
	LegalEntityRegistration                *string                                  `json:"LegalEntityRegistration"`
	Country                                string                                   `json:"Country"`
	Language                               string                                   `json:"Language"`
	Currency                               *string                                  `json:"Currency"`
	Representative           	  		   *string									`json:"Representative"`
	PhoneNumber           		  		   *string 									`json:"PhoneNumber"`
	OrganizationBPName1                    *string                                  `json:"OrganizationBPName1"`
	OrganizationBPName2                    *string                                  `json:"OrganizationBPName2"`
	OrganizationBPName3                    *string                                  `json:"OrganizationBPName3"`
	OrganizationBPName4                    *string                                  `json:"OrganizationBPName4"`
	Tag1                                   *string                                  `json:"Tag1"`
	Tag2                                   *string                                  `json:"Tag2"`
	Tag3                                   *string                                  `json:"Tag3"`
	Tag4                                   *string                                  `json:"Tag4"`
	OrganizationFoundationDate             *string                                  `json:"OrganizationFoundationDate"`
	OrganizationLiquidationDate            *string                                  `json:"OrganizationLiquidationDate"`
	BusinessPartnerBirthplaceName          *string                                  `json:"BusinessPartnerBirthplaceName"`
	BusinessPartnerDeathDate               *string                                  `json:"BusinessPartnerDeathDate"`
	GroupBusinessPartnerName1              *string                                  `json:"GroupBusinessPartnerName1"`
	GroupBusinessPartnerName2              *string                                  `json:"GroupBusinessPartnerName2"`
	AddressID                              *int                                     `json:"AddressID"`
	BusinessPartnerIDByExtSystem           *string                                  `json:"BusinessPartnerIDByExtSystem"`
	BusinessPartnerIsBlocked               *bool                                    `json:"BusinessPartnerIsBlocked"`
	CertificateAuthorityChain              *string                                  `json:"CertificateAuthorityChain"`
	UsageControlChain                      *string                                  `json:"UsageControlChain"`
	Withdrawal           		  		   *bool									`json:"Withdrawal"`
	CreationDate                           string                                   `json:"CreationDate"`
	LastChangeDate                         string                                   `json:"LastChangeDate"`
	IsReleased           		  		   *bool									`json:"IsReleased"`
	IsMarkedForDeletion                    *bool                                    `json:"IsMarkedForDeletion"`
	BusinessPartnerRole                    []BusinessPartnerRole                    `json:"Role"`
	BusinessPartnerPerson                  []BusinessPartnerPerson                  `json:"Person"`
	BusinessPartnerAddress                 []BusinessPartnerAddress                 `json:"Address"`
	BusinessPartnerSNS                     []BusinessPartnerSNS                     `json:"SNS"`
	BusinessPartnerGPS                     []BusinessPartnerGPS                     `json:"GPS"`
	BusinessPartnerRank                    []BusinessPartnerRank                    `json:"Rank"`
	BusinessPartnerPersonMobilePhoneAuth   []BusinessPartnerPersonMobilePhoneAuth   `json:"PersonMobilePhoneAuth"`
	BusinessPartnerPersonGoogleAccountAuth []BusinessPartnerPersonGoogleAccountAuth `json:"PersonGoogleAccountAuth"`
	BusinessPartnerPersonInstagramAuth	   []BusinessPartnerPersonInstagramAuth		`json:"PersonInstagramAuth"`
}

type BusinessPartnerRole struct {
	BusinessPartner     int    `json:"BusinessPartner"`
	BusinessPartnerRole string `json:"BusinessPartnerRole"`
	ValidityStartDate   string `json:"ValidityStartDate"`
	ValidityEndDate     string `json:"ValidityEndDate"`
	CreationDate        string `json:"CreationDate"`
	LastChangeDate      string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type BusinessPartnerPerson struct {
	BusinessPartner        		int     `json:"BusinessPartner"`
	BusinessPartnerType    		string  `json:"BusinessPartnerType"`
	FirstName              		string	`json:"FirstName"`
	LastName               		string	`json:"LastName"`
	FullName               		string	`json:"FullName"`
	MiddleName             		*string `json:"MiddleName"`
	NickName               		string  `json:"NickName"`
	Gender                 		string  `json:"Gender"`
	Language               		string  `json:"Language"`
	CorrespondenceLanguage 		*string `json:"CorrespondenceLanguage"`
	BirthDate              		string	`json:"BirthDate"`
	Nationality		            string  `json:"Nationality"`
	EmailAddress           		*string `json:"EmailAddress"`
	MobilePhoneNumber      		*string `json:"MobilePhoneNumber"`
	ProfileComment         		*string `json:"ProfileComment"`
	PreferableLocalSubRegion	string  `json:"PreferableLocalSubRegion"`
	PreferableLocalRegion		string  `json:"PreferableLocalRegion"`
	PreferableCountry			string  `json:"PreferableCountry"`
	ActPurpose					string  `json:"ActPurpose"`
	TermsOfUseIsConfirmed		*bool   `json:"TermsOfUseIsConfirmed"`
	CreationDate           		string  `json:"CreationDate"`
	LastChangeDate         		string  `json:"LastChangeDate"`
	IsMarkedForDeletion    		*bool   `json:"IsMarkedForDeletion"`
}

type BusinessPartnerAddress struct {
	BusinessPartner int      `json:"BusinessPartner"`
	AddressID       int      `json:"AddressID"`
	PostalCode      string   `json:"PostalCode"`
	LocalSubRegion  string   `json:"LocalSubRegion"`
	LocalRegion     string   `json:"LocalRegion"`
	Country         string   `json:"Country"`
	GlobalRegion    string   `json:"GlobalRegion"`
	TimeZone        string   `json:"TimeZone"`
	District        *string  `json:"District"`
	StreetName      *string  `json:"StreetName"`
	CityName        *string  `json:"CityName"`
	Building        *string  `json:"Building"`
	Floor           *int     `json:"Floor"`
	Room            *int     `json:"Room"`
	XCoordinate     *float32 `json:"XCoordinate"`
	YCoordinate     *float32 `json:"YCoordinate"`
	ZCoordinate     *float32 `json:"ZCoordinate"`
	Site            *int     `json:"Site"`
}

type BusinessPartnerSNS struct {
	BusinessPartner     int     `json:"BusinessPartner"`
	BusinessPartnerType string  `json:"BusinessPartnerType"`
	XURL                *string `json:"XURL"`
	InstagramURL        *string `json:"InstagramURL"`
	TikTokURL           *string `json:"TikTokURL"`
	PointAppsURL        string	`json:"PointAppsURL"`
	CreationDate        string  `json:"CreationDate"`
	LastChangeDate      string  `json:"LastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

type BusinessPartnerGPS struct {
	BusinessPartner     int     `json:"BusinessPartner"`
	BusinessPartnerType string  `json:"BusinessPartnerType"`
	XCoordinate         float32 `json:"XCoordinate"`
	YCoordinate         float32 `json:"YCoordinate"`
	ZCoordinate         float32 `json:"ZCoordinate"`
	LocalSubRegion      string  `json:"LocalSubRegion"`
	LocalRegion         string  `json:"LocalRegion"`
	Country             string  `json:"Country"`
	CreationDate        string  `json:"CreationDate"`
	CreationTime        string  `json:"CreationTime"`
	LastChangeDate      string  `json:"LastChangeDate"`
	LastChangeTime      string  `json:"LastChangeTime"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

type BusinessPartnerRank struct {
	BusinessPartner     int    `json:"BusinessPartner"`
	RankType            string `json:"RankType"`
	Rank                int    `json:"Rank"`
	ValidityStartDate   string `json:"ValidityStartDate"`
	ValidityEndDate     string `json:"ValidityEndDate"`
	CreationDate        string `json:"CreationDate"`
	LastChangeDate      string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type BusinessPartnerPersonMobilePhoneAuth struct {
	BusinessPartner            int    `json:"BusinessPartner"`
	BusinessPartnerType        string `json:"BusinessPartnerType"`
	MobilePhoneNumber          string `json:"MobilePhoneNumber"`
	MobilePhoneIsAuthenticated bool   `json:"MobilePhoneIsAuthenticated"`
	CreationDate               string `json:"CreationDate"`
	LastChangeDate             string `json:"LastChangeDate"`
	IsMarkedForDeletion        *bool  `json:"IsMarkedForDeletion"`
}

type BusinessPartnerPersonGoogleAccountAuth struct {
	BusinessPartner              int    `json:"BusinessPartner"`
	BusinessPartnerType          string `json:"BusinessPartnerType"`
	EmailAddress                 string `json:"EmailAddress"`
	GoogleID					 string	`json:"GoogleID"`
	GoogleAccountIsAuthenticated bool   `json:"GoogleAccountIsAuthenticated"`
	CreationDate                 string `json:"CreationDate"`
	LastChangeDate               string `json:"LastChangeDate"`
	IsMarkedForDeletion          *bool  `json:"IsMarkedForDeletion"`
}

type BusinessPartnerPersonInstagramAuth struct {
	BusinessPartner				int		`json:"BusinessPartner"`
	BusinessPartnerType			string	`json:"BusinessPartnerType"`
	InstagramID					string	`json:"InstagramID"`
	InstagramUserName        	string  `json:"InstagramUserName"`
	InstagramHasProfilePicture 	bool    `json:"InstagramHasProfilePicture"`
	InstagramProfilePictureURI	*string `json:"InstagramProfilePictureURI"`
	InstagramIsPrivate 			bool    `json:"InstagramIsPrivate"`
	InstagramIsPublished 		bool    `json:"InstagramIsPublished"`
	InstagramIsAuthenticated	bool	`json:"InstagramIsAuthenticated"`
	CreationDate				string	`json:"CreationDate"`
	LastChangeDate				string	`json:"LastChangeDate"`
	IsMarkedForDeletion			*bool	`json:"IsMarkedForDeletion"`
}

func BusinessPartnerInputRead(
	controller *beego.Controller,
) BusinessPartnerSDC {
	var requestBody BusinessPartnerSDC
	err := json.Unmarshal(controller.Ctx.Input.RequestBody, &requestBody)

	if err != nil {
		controller.Data["json"] = map[string]interface{}{
			"error": fmt.Sprintf("json decode error: %v", err),
		}
		controller.ServeJSON()
	}

	return requestBody
}
