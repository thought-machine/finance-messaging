package iso20022

// Provide information about the reason for the modification and about the application request which triggered this modification.
type InvestmentAccountModification1 struct {

	// Reason for the modification brought to the investment account information.
	ModificationReason *Max350Text `xml:"ModRsn,omitempty"`

	// Unique and unambiguous identifier of the account modification request.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`

	// Unique and unambiguous investor's identification of a transfer.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer as allocated by the counterparty.
	CounterpartyReference *AdditionalReference2 `xml:"CtrPtyRef,omitempty"`
}

func (i *InvestmentAccountModification1) SetModificationReason(value string) {
	i.ModificationReason = (*Max350Text)(&value)
}

func (i *InvestmentAccountModification1) SetAccountApplicationIdentification(value string) {
	i.AccountApplicationIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccountModification1) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *InvestmentAccountModification1) AddCounterpartyReference() *AdditionalReference2 {
	i.CounterpartyReference = new(AdditionalReference2)
	return i.CounterpartyReference
}

// Information about the modification of an account.
type InvestmentAccountModification2 struct {

	// Reason for the modification to the investment account information.
	ModificationReason *Max350Text `xml:"ModRsn,omitempty"`

	// Unique and unambiguous identifier of the account modification request.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`

	// Unique and unambiguous investor's identification of a transfer.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transaction, for example, a transfer, as allocated by the counterparty.
	CounterpartyReference *AdditionalReference6 `xml:"CtrPtyRef,omitempty"`

	// Account to which the account opening is related.
	ExistingAccountIdentification *Max35Text `xml:"ExstgAcctId,omitempty"`
}

func (i *InvestmentAccountModification2) SetModificationReason(value string) {
	i.ModificationReason = (*Max350Text)(&value)
}

func (i *InvestmentAccountModification2) SetAccountApplicationIdentification(value string) {
	i.AccountApplicationIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccountModification2) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *InvestmentAccountModification2) AddCounterpartyReference() *AdditionalReference6 {
	i.CounterpartyReference = new(AdditionalReference6)
	return i.CounterpartyReference
}

func (i *InvestmentAccountModification2) SetExistingAccountIdentification(value string) {
	i.ExistingAccountIdentification = (*Max35Text)(&value)
}

// Information about the modification of an account.
type InvestmentAccountModification3 struct {

	// Reason for the modification to the investment account information.
	ModificationReason *Max350Text `xml:"ModRsn,omitempty"`

	// Unique and unambiguous identifier of the account modification request.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`

	// Unique and unambiguous investor's identification of a transfer.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transaction, for example, a transfer, as allocated by the counterparty.
	CounterpartyReference *AdditionalReference6 `xml:"CtrPtyRef,omitempty"`

	// Account to which the account opening is related.
	ExistingAccountIdentification []*Account23 `xml:"ExstgAcctId,omitempty"`
}

func (i *InvestmentAccountModification3) SetModificationReason(value string) {
	i.ModificationReason = (*Max350Text)(&value)
}

func (i *InvestmentAccountModification3) SetAccountApplicationIdentification(value string) {
	i.AccountApplicationIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccountModification3) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *InvestmentAccountModification3) AddCounterpartyReference() *AdditionalReference6 {
	i.CounterpartyReference = new(AdditionalReference6)
	return i.CounterpartyReference
}

func (i *InvestmentAccountModification3) AddExistingAccountIdentification() *Account23 {
	newValue := new(Account23)
	i.ExistingAccountIdentification = append(i.ExistingAccountIdentification, newValue)
	return newValue
}
