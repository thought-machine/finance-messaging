package iso20022

// Amount of money associated with a service.
type ChargesDetails1 struct {

	// Type of service for which a charge is asked or paid.
	Type *ChargeType8Code `xml:"Tp"`

	// Specifies the type of charge by means of a code.
	OtherChargesType *Max35Text `xml:"OthrChrgsTp"`

	// Amount of money asked or paid for the charge.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Rate used to calculate the amount of the charge or fee.
	Rate *PercentageRate `xml:"Rate"`
}

func (c *ChargesDetails1) SetType(value string) {
	c.Type = (*ChargeType8Code)(&value)
}

func (c *ChargesDetails1) SetOtherChargesType(value string) {
	c.OtherChargesType = (*Max35Text)(&value)
}

func (c *ChargesDetails1) SetAmount(value, currency string) {
	c.Amount = NewCurrencyAndAmount(value, currency)
}

func (c *ChargesDetails1) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

// Amount of money associated with a service.
type ChargesDetails2 struct {

	// Type of service for which a charge is asked or paid.
	Type *ChargeType8Code `xml:"Tp"`

	// Specifies the type of charge by means of a code.
	OtherChargesType *Max35Text `xml:"OthrChrgsTp"`

	// Amount of money asked or paid for the charge.
	Amount *CurrencyAndAmount `xml:"Amt"`
}

func (c *ChargesDetails2) SetType(value string) {
	c.Type = (*ChargeType8Code)(&value)
}

func (c *ChargesDetails2) SetOtherChargesType(value string) {
	c.OtherChargesType = (*Max35Text)(&value)
}

func (c *ChargesDetails2) SetAmount(value, currency string) {
	c.Amount = NewCurrencyAndAmount(value, currency)
}

// Amount of money associated with a service.
type ChargesDetails3 struct {

	// Specifies the type of charges as a code or free text.
	Type *ChargesType1Choice `xml:"Tp"`

	// Specifies if it is a fixed amount or a percentage.
	AmountOrPercentage *AmountOrPercentage2Choice `xml:"AmtOrPctg"`
}

func (c *ChargesDetails3) AddType() *ChargesType1Choice {
	c.Type = new(ChargesType1Choice)
	return c.Type
}

func (c *ChargesDetails3) AddAmountOrPercentage() *AmountOrPercentage2Choice {
	c.AmountOrPercentage = new(AmountOrPercentage2Choice)
	return c.AmountOrPercentage
}

// Amount of money associated with a service.
type ChargesDetails4 struct {

	// Specifies the type of charges as a code or free text.
	ChargesType *ChargesType1Choice `xml:"ChrgsTp"`

	// Amount of money asked or paid for the charge.
	Amount *CurrencyAndAmount `xml:"Amt"`
}

func (c *ChargesDetails4) AddChargesType() *ChargesType1Choice {
	c.ChargesType = new(ChargesType1Choice)
	return c.ChargesType
}

func (c *ChargesDetails4) SetAmount(value, currency string) {
	c.Amount = NewCurrencyAndAmount(value, currency)
}
