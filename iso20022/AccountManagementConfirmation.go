package iso20022

// Provide information about the type of request or instruction which triggered this confirmation.
type AccountManagementConfirmation1 struct {

	// Specifies if the confirmation message applies to an account opening,  an account modification request or to a get account details.
	ConfirmationType *AccountManagementType2Code `xml:"ConfTp"`

	// Unique and unambiguous identifier of the account opening or modification instruction at application level.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`
}

func (a *AccountManagementConfirmation1) SetConfirmationType(value string) {
	a.ConfirmationType = (*AccountManagementType2Code)(&value)
}

func (a *AccountManagementConfirmation1) SetAccountApplicationIdentification(value string) {
	a.AccountApplicationIdentification = (*Max35Text)(&value)
}

// Information about the type of request or instruction.
type AccountManagementConfirmation2 struct {

	// Specifies if the confirmation message applies to an account opening,  an account modification request or to a get account details.
	ConfirmationType *AccountManagementType2Code `xml:"ConfTp"`

	// Unique and unambiguous identifier of the account opening or modification instruction at application level.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`

	// Unique and unambiguous investor's identification of a transfer.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer as allocated by the counterparty.
	CounterpartyReference *AdditionalReference2 `xml:"CtrPtyRef,omitempty"`
}

func (a *AccountManagementConfirmation2) SetConfirmationType(value string) {
	a.ConfirmationType = (*AccountManagementType2Code)(&value)
}

func (a *AccountManagementConfirmation2) SetAccountApplicationIdentification(value string) {
	a.AccountApplicationIdentification = (*Max35Text)(&value)
}

func (a *AccountManagementConfirmation2) SetClientReference(value string) {
	a.ClientReference = (*Max35Text)(&value)
}

func (a *AccountManagementConfirmation2) AddCounterpartyReference() *AdditionalReference2 {
	a.CounterpartyReference = new(AdditionalReference2)
	return a.CounterpartyReference
}

// Information about the type of request or instruction.
type AccountManagementConfirmation3 struct {

	// Specifies if the confirmation message applies to an account opening,  an account modification request or to a get account details.
	ConfirmationType *AccountManagementType2Code `xml:"ConfTp"`

	// Unique and unambiguous identifier of the account opening or modification instruction at application level.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`

	// Unique and unambiguous identification of a transaction, for example, a transfer, as assigned by the investor or account owner.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transaction, for example, a transfer, as allocated by the counterparty.
	CounterpartyReference *AdditionalReference6 `xml:"CtrPtyRef,omitempty"`

	// Account to which the account opening is related.
	ExistingAccountIdentification *Max35Text `xml:"ExstgAcctId,omitempty"`
}

func (a *AccountManagementConfirmation3) SetConfirmationType(value string) {
	a.ConfirmationType = (*AccountManagementType2Code)(&value)
}

func (a *AccountManagementConfirmation3) SetAccountApplicationIdentification(value string) {
	a.AccountApplicationIdentification = (*Max35Text)(&value)
}

func (a *AccountManagementConfirmation3) SetClientReference(value string) {
	a.ClientReference = (*Max35Text)(&value)
}

func (a *AccountManagementConfirmation3) AddCounterpartyReference() *AdditionalReference6 {
	a.CounterpartyReference = new(AdditionalReference6)
	return a.CounterpartyReference
}

func (a *AccountManagementConfirmation3) SetExistingAccountIdentification(value string) {
	a.ExistingAccountIdentification = (*Max35Text)(&value)
}

// Information about the type of request or instruction.
type AccountManagementConfirmation4 struct {

	// Specifies the confirmation type.
	ConfirmationType *ConfirmationType1Choice `xml:"ConfTp"`

	// Unique and unambiguous identifier of the account opening or modification instruction at application level.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`

	// Unique and unambiguous identification of a transaction, for example, a transfer, as assigned by the investor or account owner.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transaction, for example, a transfer, as allocated by the counterparty.
	CounterpartyReference *AdditionalReference6 `xml:"CtrPtyRef,omitempty"`

	// Account to which the account opening is related.
	ExistingAccountIdentification []*Account23 `xml:"ExstgAcctId,omitempty"`
}

func (a *AccountManagementConfirmation4) AddConfirmationType() *ConfirmationType1Choice {
	a.ConfirmationType = new(ConfirmationType1Choice)
	return a.ConfirmationType
}

func (a *AccountManagementConfirmation4) SetAccountApplicationIdentification(value string) {
	a.AccountApplicationIdentification = (*Max35Text)(&value)
}

func (a *AccountManagementConfirmation4) SetClientReference(value string) {
	a.ClientReference = (*Max35Text)(&value)
}

func (a *AccountManagementConfirmation4) AddCounterpartyReference() *AdditionalReference6 {
	a.CounterpartyReference = new(AdditionalReference6)
	return a.CounterpartyReference
}

func (a *AccountManagementConfirmation4) AddExistingAccountIdentification() *Account23 {
	newValue := new(Account23)
	a.ExistingAccountIdentification = append(a.ExistingAccountIdentification, newValue)
	return newValue
}
