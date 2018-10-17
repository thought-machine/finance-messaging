package iso20022

// Specification of the cash account.
type CashAccountType1 struct {

	// Structured format.
	Structured *FundCashAccount1Code `xml:"Strd"`

	// Additional information about the type of tax.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CashAccountType1) SetStructured(value string) {
	c.Structured = (*FundCashAccount1Code)(&value)
}

func (c *CashAccountType1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}

// Nature or use of the account.
type CashAccountType2 struct {

	// Account type, in a coded form.
	Code *CashAccountType4Code `xml:"Cd"`

	// Nature or use of the account in a proprietary form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (c *CashAccountType2) SetCode(value string) {
	c.Code = (*CashAccountType4Code)(&value)
}

func (c *CashAccountType2) SetProprietary(value string) {
	c.Proprietary = (*Max35Text)(&value)
}
