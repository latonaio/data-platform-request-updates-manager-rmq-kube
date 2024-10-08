package apiModuleRuntimesResponsesPointBalance

type PointBalanceRes struct {
	Message PointBalance `json:"message,omitempty"`
}

type PointBalance struct {
	PointBalance *[]PointBalance `json:"PointBalance,omitempty"`
	ByShop		 *[]ByShop		 `json:"ByShop,omitempty"`
}

type PointBalance struct {
	BusinessPartner		int			`json:"BusinessPartner"`
	PointSymbol			string		`json:"PointSymbol"`
	CurrentBalance		float32		`json:"CurrentBalance"`
	LimitBalance		*float32	`json:"LimitBalance"`
	CreationDate		string		`json:"CreationDate"`
	CreationTime		string		`json:"CreationTime"`
	LastChangeDate		string		`json:"LastChangeDate"`
	LastChangeTime		string		`json:"LastChangeTime"`
}

type ByShop struct {
	BusinessPartner		int			`json:"BusinessPartner"`
	PointSymbol			string		`json:"PointSymbol"`
	Shop				int			`json:"Shop"`
	CurrentBalance		float32		`json:"CurrentBalance"`
	LimitBalance		*float32	`json:"LimitBalance"`
	CreationDate		string		`json:"CreationDate"`
	CreationTime		string		`json:"CreationTime"`
	LastChangeDate		string		`json:"LastChangeDate"`
	LastChangeTime		string		`json:"LastChangeTime"`
}
