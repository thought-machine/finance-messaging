package iso20022

// Information about an accounting statement of holdings.
type AccountingStatementOfHoldings1 struct {

	// Pagination of the message.
	MessagePagination *Pagination `xml:"MsgPgntn"`

	// General information related to the custody statement of holdings that is being cancelled.
	StatementGeneralDetails *Statement4 `xml:"StmtGnlDtls,omitempty"`

	// The safekeeping or investment account of the statement that is being cancelled.
	AccountDetails *SafekeepingAccount1 `xml:"AcctDtls,omitempty"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	BalanceForAccount []*AggregateBalanceInformation2 `xml:"BalForAcct,omitempty"`

	// The sub-account of the safekeeping or investment account.
	SubAccountDetails []*SubAccountIdentification2 `xml:"SubAcctDtls,omitempty"`

	// Value of total holdings reported.
	TotalValues *TotalValueInPageAndStatement `xml:"TtlVals,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (a *AccountingStatementOfHoldings1) AddMessagePagination() *Pagination {
	a.MessagePagination = new(Pagination)
	return a.MessagePagination
}

func (a *AccountingStatementOfHoldings1) AddStatementGeneralDetails() *Statement4 {
	a.StatementGeneralDetails = new(Statement4)
	return a.StatementGeneralDetails
}

func (a *AccountingStatementOfHoldings1) AddAccountDetails() *SafekeepingAccount1 {
	a.AccountDetails = new(SafekeepingAccount1)
	return a.AccountDetails
}

func (a *AccountingStatementOfHoldings1) AddBalanceForAccount() *AggregateBalanceInformation2 {
	newValue := new(AggregateBalanceInformation2)
	a.BalanceForAccount = append(a.BalanceForAccount, newValue)
	return newValue
}

func (a *AccountingStatementOfHoldings1) AddSubAccountDetails() *SubAccountIdentification2 {
	newValue := new(SubAccountIdentification2)
	a.SubAccountDetails = append(a.SubAccountDetails, newValue)
	return newValue
}

func (a *AccountingStatementOfHoldings1) AddTotalValues() *TotalValueInPageAndStatement {
	a.TotalValues = new(TotalValueInPageAndStatement)
	return a.TotalValues
}

func (a *AccountingStatementOfHoldings1) AddExtension() *Extension1 {
	newValue := new(Extension1)
	a.Extension = append(a.Extension, newValue)
	return newValue
}

// Information about an accounting statement of holdings.
type AccountingStatementOfHoldings2 struct {

	// General information related to the custody statement of holdings that is being cancelled.
	StatementGeneralDetails *Statement6 `xml:"StmtGnlDtls,omitempty"`

	// The safekeeping or investment account of the statement that is being cancelled.
	AccountDetails *SafekeepingAccount2 `xml:"AcctDtls,omitempty"`

	// Net position of a segregated holding, in a single security, within the overall position held in a securities account.
	BalanceForAccount []*AggregateBalanceInformation3 `xml:"BalForAcct,omitempty"`

	// The sub-account of the safekeeping or investment account.
	SubAccountDetails []*SubAccountIdentification3 `xml:"SubAcctDtls,omitempty"`

	// Value of total holdings reported.
	TotalValues *TotalValueInPageAndStatement `xml:"TtlVals,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (a *AccountingStatementOfHoldings2) AddStatementGeneralDetails() *Statement6 {
	a.StatementGeneralDetails = new(Statement6)
	return a.StatementGeneralDetails
}

func (a *AccountingStatementOfHoldings2) AddAccountDetails() *SafekeepingAccount2 {
	a.AccountDetails = new(SafekeepingAccount2)
	return a.AccountDetails
}

func (a *AccountingStatementOfHoldings2) AddBalanceForAccount() *AggregateBalanceInformation3 {
	newValue := new(AggregateBalanceInformation3)
	a.BalanceForAccount = append(a.BalanceForAccount, newValue)
	return newValue
}

func (a *AccountingStatementOfHoldings2) AddSubAccountDetails() *SubAccountIdentification3 {
	newValue := new(SubAccountIdentification3)
	a.SubAccountDetails = append(a.SubAccountDetails, newValue)
	return newValue
}

func (a *AccountingStatementOfHoldings2) AddTotalValues() *TotalValueInPageAndStatement {
	a.TotalValues = new(TotalValueInPageAndStatement)
	return a.TotalValues
}

func (a *AccountingStatementOfHoldings2) AddExtension() *Extension1 {
	newValue := new(Extension1)
	a.Extension = append(a.Extension, newValue)
	return newValue
}
