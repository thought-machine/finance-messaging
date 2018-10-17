package iso20022

// Provides information about the cash movement.
type CashMovement1 struct {

	// Identification of the movement.
	MovementIdentification *Max35Text `xml:"MvmntId,omitempty"`

	// Cash amount.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Amount of taxes.
	TaxAmount *ActiveCurrencyAndAmount `xml:"TaxAmt,omitempty"`

	// Specifies the charges.
	Charges []*Charges1 `xml:"Chrgs,omitempty"`

	// Provides information about the account which is debited/credited.
	AccountDetails []*CashAccount18 `xml:"AcctDtls"`
}

func (c *CashMovement1) SetMovementIdentification(value string) {
	c.MovementIdentification = (*Max35Text)(&value)
}

func (c *CashMovement1) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CashMovement1) SetTaxAmount(value, currency string) {
	c.TaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CashMovement1) AddCharges() *Charges1 {
	newValue := new(Charges1)
	c.Charges = append(c.Charges, newValue)
	return newValue
}

func (c *CashMovement1) AddAccountDetails() *CashAccount18 {
	newValue := new(CashAccount18)
	c.AccountDetails = append(c.AccountDetails, newValue)
	return newValue
}

// Provides information about the cash movement.
type CashMovement2 struct {

	// Cash amount.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Provides information about the account which is debited/credited.
	AccountDetails []*CashAccount19 `xml:"AcctDtls"`
}

func (c *CashMovement2) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CashMovement2) AddAccountDetails() *CashAccount19 {
	newValue := new(CashAccount19)
	c.AccountDetails = append(c.AccountDetails, newValue)
	return newValue
}

// Provides information about the cash movement.
type CashMovement3 struct {

	// Date and time of the posting.
	PostingDateTime *DateAndDateTimeChoice `xml:"PstngDtTm,omitempty"`

	// Value date.
	ValueDate *ISODate `xml:"ValDt"`

	// Cash amount that is posted.
	PostingAmount *ActiveCurrencyAndAmount `xml:"PstngAmt"`

	// Provides information about the account which is debited/credited.
	AccountDetails []*CashAccount18 `xml:"AcctDtls"`
}

func (c *CashMovement3) AddPostingDateTime() *DateAndDateTimeChoice {
	c.PostingDateTime = new(DateAndDateTimeChoice)
	return c.PostingDateTime
}

func (c *CashMovement3) SetValueDate(value string) {
	c.ValueDate = (*ISODate)(&value)
}

func (c *CashMovement3) SetPostingAmount(value, currency string) {
	c.PostingAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CashMovement3) AddAccountDetails() *CashAccount18 {
	newValue := new(CashAccount18)
	c.AccountDetails = append(c.AccountDetails, newValue)
	return newValue
}
