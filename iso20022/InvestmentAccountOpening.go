package iso20022

// Provides information about the type of opening instruction and identification of the application request.
type InvestmentAccountOpening1 struct {

	// Specifies if the account opening instruction is about a newly created account or a supplementary account.
	OpeningType *AccountOpeningType1Code `xml:"OpngTp"`

	// Unique and unambiguous identifier of the account opening request at application level.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`

	// Unique and unambiguous investor's identification of a transfer.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer as allocated by the counterparty.
	CounterpartyReference *AdditionalReference2 `xml:"CtrPtyRef,omitempty"`
}

func (i *InvestmentAccountOpening1) SetOpeningType(value string) {
	i.OpeningType = (*AccountOpeningType1Code)(&value)
}

func (i *InvestmentAccountOpening1) SetAccountApplicationIdentification(value string) {
	i.AccountApplicationIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccountOpening1) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *InvestmentAccountOpening1) AddCounterpartyReference() *AdditionalReference2 {
	i.CounterpartyReference = new(AdditionalReference2)
	return i.CounterpartyReference
}

// Information about the type of opening instruction and identification of the application request.
type InvestmentAccountOpening2 struct {

	// Specifies if the account opening instruction is about a newly created account or a supplementary account.
	OpeningType *AccountOpeningType1Code `xml:"OpngTp"`

	// Unique and unambiguous identifier of the account opening request at application level.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`

	// Unique and unambiguous identification of a transaction, for example, a transfer, as assigned by the investor or account owner.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transaction, for example, a transfer, as allocated by the counterparty.
	CounterpartyReference *AdditionalReference6 `xml:"CtrPtyRef,omitempty"`

	// Account to which the account opening is related.
	ExistingAccountIdentification *Max35Text `xml:"ExstgAcctId,omitempty"`
}

func (i *InvestmentAccountOpening2) SetOpeningType(value string) {
	i.OpeningType = (*AccountOpeningType1Code)(&value)
}

func (i *InvestmentAccountOpening2) SetAccountApplicationIdentification(value string) {
	i.AccountApplicationIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccountOpening2) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *InvestmentAccountOpening2) AddCounterpartyReference() *AdditionalReference6 {
	i.CounterpartyReference = new(AdditionalReference6)
	return i.CounterpartyReference
}

func (i *InvestmentAccountOpening2) SetExistingAccountIdentification(value string) {
	i.ExistingAccountIdentification = (*Max35Text)(&value)
}

// Information about the type of opening instruction and identification of the application request.
type InvestmentAccountOpening3 struct {

	// Specifies if the account opening instruction is about a newly created account or a supplementary account.
	OpeningType *AccountOpeningType1Choice `xml:"OpngTp"`

	// Unique and unambiguous identifier of the account opening request at application level.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`

	// Unique and unambiguous identification of a transaction, for example, a transfer, as assigned by the investor or account owner.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transaction, for example, a transfer, as allocated by the counterparty.
	CounterpartyReference *AdditionalReference6 `xml:"CtrPtyRef,omitempty"`

	// Account to which the account opening is related.
	ExistingAccountIdentification []*Account23 `xml:"ExstgAcctId,omitempty"`
}

func (i *InvestmentAccountOpening3) AddOpeningType() *AccountOpeningType1Choice {
	i.OpeningType = new(AccountOpeningType1Choice)
	return i.OpeningType
}

func (i *InvestmentAccountOpening3) SetAccountApplicationIdentification(value string) {
	i.AccountApplicationIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccountOpening3) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *InvestmentAccountOpening3) AddCounterpartyReference() *AdditionalReference6 {
	i.CounterpartyReference = new(AdditionalReference6)
	return i.CounterpartyReference
}

func (i *InvestmentAccountOpening3) AddExistingAccountIdentification() *Account23 {
	newValue := new(Account23)
	i.ExistingAccountIdentification = append(i.ExistingAccountIdentification, newValue)
	return newValue
}
