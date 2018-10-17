package iso20022

// Information about an amount.
type UndertakingAmount1 struct {

	// Amount and currency of the undertaking.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Percentage by which the amount claimed under the undertaking may be more than the undertaking amount.
	PlusTolerance *PercentageRate `xml:"PlusTlrnce,omitempty"`

	// Additional information concerning the undertaking amount.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (u *UndertakingAmount1) SetAmount(value, currency string) {
	u.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (u *UndertakingAmount1) SetPlusTolerance(value string) {
	u.PlusTolerance = (*PercentageRate)(&value)
}

func (u *UndertakingAmount1) AddAdditionalInformation(value string) {
	u.AdditionalInformation = append(u.AdditionalInformation, (*Max2000Text)(&value))
}

// Information about an amount.
type UndertakingAmount2 struct {

	// Choice of amounts.
	AmountChoice *Amount1Choice `xml:"AmtChc"`

	// Additional information concerning the amended amount.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (u *UndertakingAmount2) AddAmountChoice() *Amount1Choice {
	u.AmountChoice = new(Amount1Choice)
	return u.AmountChoice
}

func (u *UndertakingAmount2) AddAdditionalInformation(value string) {
	u.AdditionalInformation = append(u.AdditionalInformation, (*Max2000Text)(&value))
}

// Information about an amount.
type UndertakingAmount3 struct {

	// Amount and currency of the demand.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Additional information concerning the demand amount.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (u *UndertakingAmount3) SetAmount(value, currency string) {
	u.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (u *UndertakingAmount3) AddAdditionalInformation(value string) {
	u.AdditionalInformation = append(u.AdditionalInformation, (*Max2000Text)(&value))
}

// Defined variation amount and balance.
type UndertakingAmount4 struct {

	// Variation amount and currency.
	VariationAmount *ActiveCurrencyAndAmount `xml:"VartnAmt"`

	// Calculated undertaking available balance amount resulting from the application of the variation amount.
	BalanceAmount *ActiveCurrencyAndAmount `xml:"BalAmt,omitempty"`
}

func (u *UndertakingAmount4) SetVariationAmount(value, currency string) {
	u.VariationAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (u *UndertakingAmount4) SetBalanceAmount(value, currency string) {
	u.BalanceAmount = NewActiveCurrencyAndAmount(value, currency)
}
