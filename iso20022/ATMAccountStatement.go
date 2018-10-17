package iso20022

// Statement information of an account.
type ATMAccountStatement1 struct {

	// Unique identifier of the account, as assigned by the account servicer.
	AccountIdentifier *AccountIdentification31Choice `xml:"AcctIdr"`

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	AccountName *Max70Text `xml:"AcctNm,omitempty"`

	// Statement information.
	AccountStatement []*ATMAccountStatement2 `xml:"AcctStmt,omitempty"`
}

func (a *ATMAccountStatement1) AddAccountIdentifier() *AccountIdentification31Choice {
	a.AccountIdentifier = new(AccountIdentification31Choice)
	return a.AccountIdentifier
}

func (a *ATMAccountStatement1) SetAccountName(value string) {
	a.AccountName = (*Max70Text)(&value)
}

func (a *ATMAccountStatement1) AddAccountStatement() *ATMAccountStatement2 {
	newValue := new(ATMAccountStatement2)
	a.AccountStatement = append(a.AccountStatement, newValue)
	return newValue
}

// Statement information of an account.
type ATMAccountStatement2 struct {

	// Date of the transaction.
	TransactionDate *ISODate `xml:"TxDt,omitempty"`

	// Value date of the transaction.
	ValueDate *ISODate `xml:"ValDt,omitempty"`

	// Short text to display or print for the statement.
	ShortText *Max70Text `xml:"ShrtTxt,omitempty"`

	// True if credit transaction.
	CreditTransaction *TrueFalseIndicator `xml:"CdtTx,omitempty"`

	// Amount of the transaction.
	Amount *ImpliedCurrencyAndAmount `xml:"Amt"`

	// Currency of the amount.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Alternative text of the statement to print or display.
	LongText *Max256Text `xml:"LngTxt,omitempty"`
}

func (a *ATMAccountStatement2) SetTransactionDate(value string) {
	a.TransactionDate = (*ISODate)(&value)
}

func (a *ATMAccountStatement2) SetValueDate(value string) {
	a.ValueDate = (*ISODate)(&value)
}

func (a *ATMAccountStatement2) SetShortText(value string) {
	a.ShortText = (*Max70Text)(&value)
}

func (a *ATMAccountStatement2) SetCreditTransaction(value string) {
	a.CreditTransaction = (*TrueFalseIndicator)(&value)
}

func (a *ATMAccountStatement2) SetAmount(value, currency string) {
	a.Amount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMAccountStatement2) SetCurrency(value string) {
	a.Currency = (*ActiveCurrencyCode)(&value)
}

func (a *ATMAccountStatement2) SetLongText(value string) {
	a.LongText = (*Max256Text)(&value)
}
