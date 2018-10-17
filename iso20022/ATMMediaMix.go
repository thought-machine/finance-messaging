package iso20022

// Media mix selected.
type ATMMediaMix1 struct {

	// Logical unit number of the cash dispenser.
	CashUnitNumber *Number `xml:"CshUnitNb,omitempty"`

	// Number of notes or coins.
	Number *Number `xml:"Nb"`

	// Unit value.
	UnitValue *ImpliedCurrencyAndAmount `xml:"UnitVal"`
}

func (a *ATMMediaMix1) SetCashUnitNumber(value string) {
	a.CashUnitNumber = (*Number)(&value)
}

func (a *ATMMediaMix1) SetNumber(value string) {
	a.Number = (*Number)(&value)
}

func (a *ATMMediaMix1) SetUnitValue(value, currency string) {
	a.UnitValue = NewImpliedCurrencyAndAmount(value, currency)
}

// Media mix selected.
type ATMMediaMix2 struct {

	// Number of notes or coins.
	Number *Number `xml:"Nb"`

	// Unit value.
	UnitValue *ImpliedCurrencyAndAmount `xml:"UnitVal"`
}

func (a *ATMMediaMix2) SetNumber(value string) {
	a.Number = (*Number)(&value)
}

func (a *ATMMediaMix2) SetUnitValue(value, currency string) {
	a.UnitValue = NewImpliedCurrencyAndAmount(value, currency)
}
