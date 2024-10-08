package apiModuleRuntimesResponsesQuestionnaire

type QuestionnaireRes struct {
	Message Questionnaire `json:"message,omitempty"`
}

type Questionnaire struct {
	Header *[]Header `json:"Header,omitempty"`
	Item   *[]Item   `json:"Item,omitempty"`
}

type Header struct {
	Questionnaire			int		`json:"Questionnaire"`
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
}

type Item struct {
	Questionnaire					int		`json:"Questionnaire"`
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
