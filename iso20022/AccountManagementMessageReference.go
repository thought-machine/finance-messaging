package iso20022

// Information about the message reference of the account management instruction message for which the status is requested .
type AccountManagementMessageReference1 struct {

	// Reference to a linked message sent in a proprietary way or reference of a system.
	OtherReference *AdditionalReference3 `xml:"OthrRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Specifies if the status request refers to an earlier account opening or modification instruction message.
	StatusRequestType *AccountManagementType1Code `xml:"StsReqTp"`

	// Unique and unambiguous identifier of the account opening or account modification instruction at application level.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`

	// Account information for which the status of an account management instruction is requested.
	InvestmentAccount *InvestmentAccount14 `xml:"InvstmtAcct,omitempty"`
}

func (a *AccountManagementMessageReference1) AddOtherReference() *AdditionalReference3 {
	a.OtherReference = new(AdditionalReference3)
	return a.OtherReference
}

func (a *AccountManagementMessageReference1) AddPreviousReference() *AdditionalReference3 {
	a.PreviousReference = new(AdditionalReference3)
	return a.PreviousReference
}

func (a *AccountManagementMessageReference1) SetStatusRequestType(value string) {
	a.StatusRequestType = (*AccountManagementType1Code)(&value)
}

func (a *AccountManagementMessageReference1) SetAccountApplicationIdentification(value string) {
	a.AccountApplicationIdentification = (*Max35Text)(&value)
}

func (a *AccountManagementMessageReference1) AddInvestmentAccount() *InvestmentAccount14 {
	a.InvestmentAccount = new(InvestmentAccount14)
	return a.InvestmentAccount
}

// Information about the message reference of the account management instruction message for which the status is requested .
type AccountManagementMessageReference2 struct {

	// Reference to a linked message.
	LinkedReference *LinkedMessage2Choice `xml:"LkdRef,omitempty"`

	// Specifies if the status request refers to an earlier account opening or modification instruction message.
	StatusRequestType *AccountManagementType1Code `xml:"StsReqTp"`

	// Unique and unambiguous identifier of the account opening or account modification instruction at application level.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`

	// Account information for which the status of an account management instruction is requested.
	InvestmentAccount *InvestmentAccount45 `xml:"InvstmtAcct,omitempty"`
}

func (a *AccountManagementMessageReference2) AddLinkedReference() *LinkedMessage2Choice {
	a.LinkedReference = new(LinkedMessage2Choice)
	return a.LinkedReference
}

func (a *AccountManagementMessageReference2) SetStatusRequestType(value string) {
	a.StatusRequestType = (*AccountManagementType1Code)(&value)
}

func (a *AccountManagementMessageReference2) SetAccountApplicationIdentification(value string) {
	a.AccountApplicationIdentification = (*Max35Text)(&value)
}

func (a *AccountManagementMessageReference2) AddInvestmentAccount() *InvestmentAccount45 {
	a.InvestmentAccount = new(InvestmentAccount45)
	return a.InvestmentAccount
}

// Information about the references of an account management instruction message.
type AccountManagementMessageReference3 struct {

	// Reference to a linked message.
	LinkedReference *LinkedMessage3Choice `xml:"LkdRef,omitempty"`

	// Specifies if the status request refers to an earlier account opening or modification instruction message.
	StatusRequestType *AccountManagementType1Code `xml:"StsReqTp"`

	// Unique and unambiguous identifier of the account opening or account modification instruction at application level.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`

	// Account to which the account opening is related.
	ExistingAccountIdentification *Max35Text `xml:"ExstgAcctId,omitempty"`

	// Account information for which the status of an account management instruction is requested.
	InvestmentAccount *InvestmentAccount53 `xml:"InvstmtAcct,omitempty"`
}

func (a *AccountManagementMessageReference3) AddLinkedReference() *LinkedMessage3Choice {
	a.LinkedReference = new(LinkedMessage3Choice)
	return a.LinkedReference
}

func (a *AccountManagementMessageReference3) SetStatusRequestType(value string) {
	a.StatusRequestType = (*AccountManagementType1Code)(&value)
}

func (a *AccountManagementMessageReference3) SetAccountApplicationIdentification(value string) {
	a.AccountApplicationIdentification = (*Max35Text)(&value)
}

func (a *AccountManagementMessageReference3) SetExistingAccountIdentification(value string) {
	a.ExistingAccountIdentification = (*Max35Text)(&value)
}

func (a *AccountManagementMessageReference3) AddInvestmentAccount() *InvestmentAccount53 {
	a.InvestmentAccount = new(InvestmentAccount53)
	return a.InvestmentAccount
}

// Information about the references of an account management instruction message.
type AccountManagementMessageReference4 struct {

	// Reference to a linked message.
	LinkedReference *LinkedMessage4Choice `xml:"LkdRef,omitempty"`

	// Type of account management instruction for which the status  is requested or a request to know the status of the account.
	StatusRequestType *AccountManagementType3Code `xml:"StsReqTp"`

	// Unique and unambiguous identifier of the account opening or account modification instruction at application level.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`

	// Account to which the account opening is related.
	ExistingAccountIdentification *Account23 `xml:"ExstgAcctId,omitempty"`

	// Account information for which the status of an account management instruction is requested.
	InvestmentAccount *InvestmentAccount53 `xml:"InvstmtAcct,omitempty"`
}

func (a *AccountManagementMessageReference4) AddLinkedReference() *LinkedMessage4Choice {
	a.LinkedReference = new(LinkedMessage4Choice)
	return a.LinkedReference
}

func (a *AccountManagementMessageReference4) SetStatusRequestType(value string) {
	a.StatusRequestType = (*AccountManagementType3Code)(&value)
}

func (a *AccountManagementMessageReference4) SetAccountApplicationIdentification(value string) {
	a.AccountApplicationIdentification = (*Max35Text)(&value)
}

func (a *AccountManagementMessageReference4) AddExistingAccountIdentification() *Account23 {
	a.ExistingAccountIdentification = new(Account23)
	return a.ExistingAccountIdentification
}

func (a *AccountManagementMessageReference4) AddInvestmentAccount() *InvestmentAccount53 {
	a.InvestmentAccount = new(InvestmentAccount53)
	return a.InvestmentAccount
}
