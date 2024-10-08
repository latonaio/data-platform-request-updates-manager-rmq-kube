package types

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type AuthenticatorSDC struct {
	ConnectionKey    	string	`json:"connection_key"`
	Result           	bool	`json:"result"`
	RedisKey         	string	`json:"redis_key"`
	Filepath        	string	`json:"filepath"`
	APIStatusCode  		int		`json:"api_status_code"`
	RuntimeSessionID	string	`json:"runtime_session_id"`
	BusinessPartner  	*int	`json:"business_partner"`
	ServiceLabel     	string	`json:"service_label"`
	APIType				string	`json:"api_type"`
	User				AuthenticatorUser `json:"Authenticator"`
	APISchema			string	`json:"api_schema"`
	Accepter         	[]string `json:"accepter"`
	Deleted          	bool	`json:"deleted"`
}

type AuthenticatorUser struct {
	UserID				string	`json:"UserID"`
	BusinessPartner		int		`json:"BusinessPartner"`
	Password			string	`json:"Password"`
	Qos					string	`json:"Qos"`
	IsEncrypt			bool	`json:"IsEncrypt"`
	Language			string	`json:"Language"`
	LastLoginDate		*string	`json:"LastLoginDate"`
	LastLoginTime		*string	`json:"LastLoginTime"`
	CreationDate		string	`json:"CreationDate"`
	CreationTime		string	`json:"CreationTime"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
	AuthenticatorInitialEmailAuth			AuthenticatorInitialEmailAuth			`json:"InitialEmailAuth"`
	AuthenticatorInitialSMSAuth				AuthenticatorInitialSMSAuth				`json:"InitialSMSAuth"`
	AuthenticatorSMSAuth					AuthenticatorSMSAuth					`json:"SMSAuth"`
	AuthenticatorInitialGoogleAccountAuth	AuthenticatorInitialGoogleAccountAuth	`json:"InitialGoogleAccountAuth"`
	AuthenticatorGoogleAccountAuth			AuthenticatorGoogleAccountAuth			`json:"GoogleAccountAuth"`
	AuthenticatorInstagramAuth				AuthenticatorInstagramAuth				`json:"InstagramAuth"`
}

type AuthenticatorInitialEmailAuth struct {
	EmailAddress		string	`json:"EmailAddress"`
	CreationDate		string	`json:"CreationDate"`
	CreationTime		string	`json:"CreationTime"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

type AuthenticatorInitialSMSAuth struct {
	MobilePhoneNumber	string	`json:"MobilePhoneNumber"`
	AuthenticationCode	int		`json:"AuthenticationCode"`
	CreationDate		string	`json:"CreationDate"`
	CreationTime		string	`json:"CreationTime"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

type AuthenticatorSMSAuth struct {
	UserID				string	`json:"UserID"`
	MobilePhoneNumber	string	`json:"MobilePhoneNumber"`
	AuthenticationCode	int		`json:"AuthenticationCode"`
	CreationDate		string	`json:"CreationDate"`
	CreationTime		string	`json:"CreationTime"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

type AuthenticatorInitialGoogleAccountAuth struct {
	EmailAddress		string	`json:"EmailAddress"`
	GoogleID			string	`json:"GoogleID"`
	AccessToken			string	`json:"AccessToken"`
	CreationDate		string	`json:"CreationDate"`
	CreationTime		string	`json:"CreationTime"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

type AuthenticatorGoogleAccountAuth struct {
	UserID				string	`json:"UserID"`
	EmailAddress		string	`json:"EmailAddress"`
	GoogleID			string	`json:"GoogleID"`
	AccessToken			string	`json:"AccessToken"`
	CreationDate		string	`json:"CreationDate"`
	CreationTime		string	`json:"CreationTime"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}

type AuthenticatorInstagramAuth struct {
	UserID              string `json:"UserID"`
	InstagramID         string `json:"InstagramID"`
	AccessToken         string `json:"AccessToken"`
	CreationDate        string `json:"CreationDate"`
	CreationTime        string `json:"CreationTime"`
	LastChangeDate      string `json:"LastChangeDate"`
	LastChangeTime      string `json:"LastChangeTime"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

func AuthenticatorInputRead(
	controller *beego.Controller,
) AuthenticatorSDC {
	var requestBody AuthenticatorSDC
	err := json.Unmarshal(controller.Ctx.Input.RequestBody, &requestBody)

	if err != nil {
		controller.Data["json"] = map[string]interface{}{
			"error": fmt.Sprintf("json decode error: %v", err),
		}
		controller.ServeJSON()
	}

	return requestBody
}
