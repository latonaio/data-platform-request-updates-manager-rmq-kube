package apiModuleRuntimesResponsesAddress

type AddressRes struct {
	Message Address `json:"message,omitempty"`
}

type Address struct {
	Address *[]Address `json:"Address,omitempty"`
}

type Address struct {
	AddressID			int     `json:"AddressID"`
	ValidityStartDate	string  `json:"ValidityStartDate"`
	ValidityEndDate		string	`json:"ValidityEndDate"`
	PostalCode			*string	`json:"PostalCode"`
	LocalSubRegion		string	`json:"LocalSubRegion"`
	LocalRegion			string	`json:"LocalRegion"`
	Country				string	`json:"Country"`
	GlobalRegion		string	`json:"GlobalRegion"`
	TimeZone			string	`json:"TimeZone"`
	District			*string	`json:"District"`
	StreetName			*string	`json:"StreetName"`
	CityName			*string	`json:"CityName"`
	Building			*string	`json:"Building"`
	Floor				*int    `json:"Floor"`
	Room				*int    `json:"Room"`
	XCoordinate 		*float32 `json:"XCoordinate"`
	YCoordinate 		*float32 `json:"YCoordinate"`
	ZCoordinate 		*float32 `json:"ZCoordinate"`
	Site				*int	 `json:"Site"`
	CreationDate        string  `json:"CreationDate"`
	CreationTime        string  `json:"CreationTime"`
	LastChangeDate      string  `json:"LastChangeDate"`
	LastChangeTime      string  `json:"LastChangeTime"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}
