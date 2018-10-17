package iso20022

// Current totals of the ATM.
type ATMTotals1 struct {

	// Type of media inside the cassette.
	MediaType *ATMMediaType1Code `xml:"MdiaTp,omitempty"`

	// Currency of the totals.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Total balance of the ATM including reject cassette, but excluding the retract cassette.
	ATMBalance *ImpliedCurrencyAndAmount `xml:"ATMBal,omitempty"`

	// Available amount for dispense in the ATM.
	ATMCurrent *ImpliedCurrencyAndAmount `xml:"ATMCur,omitempty"`

	// Total number of units for non-valued media, including reject cassette.
	ATMBalanceNumber *Number `xml:"ATMBalNb,omitempty"`

	// Total number of units for non-valued media, excluding reject cassette.
	ATMCurrentNumber *Number `xml:"ATMCurNb,omitempty"`
}

func (a *ATMTotals1) SetMediaType(value string) {
	a.MediaType = (*ATMMediaType1Code)(&value)
}

func (a *ATMTotals1) SetCurrency(value string) {
	a.Currency = (*ActiveCurrencyCode)(&value)
}

func (a *ATMTotals1) SetATMBalance(value, currency string) {
	a.ATMBalance = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTotals1) SetATMCurrent(value, currency string) {
	a.ATMCurrent = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTotals1) SetATMBalanceNumber(value string) {
	a.ATMBalanceNumber = (*Number)(&value)
}

func (a *ATMTotals1) SetATMCurrentNumber(value string) {
	a.ATMCurrentNumber = (*Number)(&value)
}

// Totals of the ATM.
type ATMTotals3 struct {

	// Identification of the type of transaction. The following values are predefined: Withdrawal, QuickWithdrawal, WithdrawalForDisabledPerson, CashDeposit, Transfer, InternationalTransfer, PeriodicTransfer, CheckCommand, MobileTopUp, Moneo.
	Identification *Max70Text `xml:"Id"`

	// Additional identification of the type of transaction. The following values are predefined:  Vodaphone, TMobile, Verizon.
	AdditionalIdentification *Max70Text `xml:"AddtlId,omitempty"`

	// Period of computation for the counters.
	Period *ATMCounterType2Code `xml:"Prd"`

	// Currency of the totals.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Number of transaction with the defined currency.
	Count *Number `xml:"Cnt"`

	// Amount of transaction with the defined currency.
	Amount *ImpliedCurrencyAndAmount `xml:"Amt,omitempty"`
}

func (a *ATMTotals3) SetIdentification(value string) {
	a.Identification = (*Max70Text)(&value)
}

func (a *ATMTotals3) SetAdditionalIdentification(value string) {
	a.AdditionalIdentification = (*Max70Text)(&value)
}

func (a *ATMTotals3) SetPeriod(value string) {
	a.Period = (*ATMCounterType2Code)(&value)
}

func (a *ATMTotals3) SetCurrency(value string) {
	a.Currency = (*ActiveCurrencyCode)(&value)
}

func (a *ATMTotals3) SetCount(value string) {
	a.Count = (*Number)(&value)
}

func (a *ATMTotals3) SetAmount(value, currency string) {
	a.Amount = NewImpliedCurrencyAndAmount(value, currency)
}
