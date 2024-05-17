package apiModuleRuntimesResponsesBusinessPartner

type BusinessPartnerRes struct {
	Message BusinessPartner `json:"message,omitempty"`
}

type BusinessPartner struct {
	General                 *[]General                 `json:"General"`
	Role                    *[]Role                    `json:"Role"`
	Person                  *[]Person                  `json:"Person"`
	Address                 *[]Address                 `json:"Address"`
	Rank                    *[]Rank                    `json:"Rank"`
	PersonMobilePhoneAuth   *[]PersonMobilePhoneAuth   `json:"PersonMobilePhoneAuth"`
	PersonGoogleAccountAuth *[]PersonGoogleAccountAuth `json:"PersonGoogleAccountAuth"`
	PersonInstagramAuth		*[]PersonInstagramAuth	   `json:"PersonInstagramAuth"`
}

type General struct {
	BusinessPartner               int     `json:"BusinessPartner"`
	BusinessPartnerType           string  `json:"BusinessPartnerType"`
	BusinessPartnerFullName       *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName           string  `json:"BusinessPartnerName"`
	Industry                      *string `json:"Industry"`
	LegalEntityRegistration       *string `json:"LegalEntityRegistration"`
	Country                       string  `json:"Country"`
	Language                      string  `json:"Language"`
	Currency                      *string `json:"Currency"`
	Representative           	  *string `json:"Representative"`
	PhoneNumber           		  *string `json:"PhoneNumber"`
	OrganizationBPName1           *string `json:"OrganizationBPName1"`
	OrganizationBPName2           *string `json:"OrganizationBPName2"`
	OrganizationBPName3           *string `json:"OrganizationBPName3"`
	OrganizationBPName4           *string `json:"OrganizationBPName4"`
	Tag1                          *string `json:"Tag1"`
	Tag2                          *string `json:"Tag2"`
	Tag3                          *string `json:"Tag3"`
	Tag4                          *string `json:"Tag4"`
	OrganizationFoundationDate    *string `json:"OrganizationFoundationDate"`
	OrganizationLiquidationDate   *string `json:"OrganizationLiquidationDate"`
	BusinessPartnerBirthplaceName *string `json:"BusinessPartnerBirthplaceName"`
	BusinessPartnerDeathDate      *string `json:"BusinessPartnerDeathDate"`
	GroupBusinessPartnerName1     *string `json:"GroupBusinessPartnerName1"`
	GroupBusinessPartnerName2     *string `json:"GroupBusinessPartnerName2"`
	AddressID                     *int    `json:"AddressID"`
	BusinessPartnerIDByExtSystem  *string `json:"BusinessPartnerIDByExtSystem"`
	BusinessPartnerIsBlocked      *bool   `json:"BusinessPartnerIsBlocked"`
	CertificateAuthorityChain     *string `json:"CertificateAuthorityChain"`
	UsageControlChain             *string `json:"UsageControlChain"`
	CreationDate                  string  `json:"CreationDate"`
	LastChangeDate                string  `json:"LastChangeDate"`
	IsMarkedForDeletion           *bool   `json:"IsMarkedForDeletion"`
}

type Role struct {
	BusinessPartner     int    `json:"BusinessPartner"`
	BusinessPartnerRole string `json:"BusinessPartnerRole"`
	ValidityStartDate   string `json:"ValidityStartDate"`
	ValidityEndDate     string `json:"ValidityEndDate"`
	CreationDate        string `json:"CreationDate"`
	LastChangeDate      string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type Person struct {
	BusinessPartner        		int     `json:"BusinessPartner"`
	BusinessPartnerType    		string  `json:"BusinessPartnerType"`
	FirstName              		*string `json:"FirstName"`
	LastName               		*string `json:"LastName"`
	FullName               		*string `json:"FullName"`
	MiddleName             		*string `json:"MiddleName"`
	NickName               		string  `json:"NickName"`
	Gender                 		string  `json:"Gender"`
	Language               		string  `json:"Language"`
	CorrespondenceLanguage 		*string `json:"CorrespondenceLanguage"`
	BirthDate              		*string `json:"BirthDate"`
	Nationality		            string  `json:"Nationality"`
	EmailAddress           		*string `json:"EmailAddress"`
	MobilePhoneNumber      		*string `json:"MobilePhoneNumber"`
	ProfileComment         		*string `json:"ProfileComment"`
	PreferableLocalSubRegion	string `json:"PreferableLocalSubRegion"`
	PreferableLocalRegion		string `json:"PreferableLocalRegion"`
	PreferableCountry			string `json:"PreferableCountry"`
	ActPurpose					string  `json:"ActPurpose"`
	CreationDate           		string  `json:"CreationDate"`
	LastChangeDate         		string  `json:"LastChangeDate"`
	IsMarkedForDeletion    		*bool   `json:"IsMarkedForDeletion"`
}

type Address struct {
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

type Rank struct {
	BusinessPartner     int    `json:"BusinessPartner"`
	RankType            string `json:"RankType"`
	Rank                int    `json:"Rank"`
	ValidityStartDate   string `json:"ValidityStartDate"`
	ValidityEndDate     string `json:"ValidityEndDate"`
	CreationDate        string `json:"CreationDate"`
	LastChangeDate      string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type PersonMobilePhoneAuth struct {
	BusinessPartner            int    `json:"BusinessPartner"`
	BusinessPartnerType        string `json:"BusinessPartnerType"`
	MobilePhoneNumber          string `json:"MobilePhoneNumber"`
	MobilePhoneIsAuthenticated bool   `json:"MobilePhoneIsAuthenticated"`
	CreationDate               string `json:"CreationDate"`
	LastChangeDate             string `json:"LastChangeDate"`
	IsMarkedForDeletion        *bool  `json:"IsMarkedForDeletion"`
}

type PersonGoogleAccountAuth struct {
	BusinessPartner              int    `json:"BusinessPartner"`
	BusinessPartnerType          string `json:"BusinessPartnerType"`
	EmailAddress                 string `json:"EmailAddress"`
	GoogleID					 string	`json:"GoogleID"`
	GoogleAccountIsAuthenticated bool   `json:"GoogleAccountIsAuthenticated"`
	CreationDate                 string `json:"CreationDate"`
	LastChangeDate               string `json:"LastChangeDate"`
	IsMarkedForDeletion          *bool  `json:"IsMarkedForDeletion"`
}

type PersonInstagramAuth struct {
	BusinessPartner				int		`json:"BusinessPartner"`
	BusinessPartnerType			string	`json:"BusinessPartnerType"`
	InstagramID					string	`json:"InstagramID"`
	InstagramUserName        	string  `json:"InstagramUserName"`
	InstagramHasProfilePicture  bool    `json:"InstagramHasProfilePicture"`
	InstagramProfilePictureURI	*string `json:"InstagramProfilePictureURI"`
	InstagramIsPrivate 			bool    `json:"InstagramIsPrivate"`
	InstagramIsPublished 		bool    `json:"InstagramIsPublished"`
	InstagramIsAuthenticated	bool	`json:"InstagramIsAuthenticated"`
	CreationDate				string	`json:"CreationDate"`
	LastChangeDate				string	`json:"LastChangeDate"`
	IsMarkedForDeletion			*bool	`json:"IsMarkedForDeletion"`
}
