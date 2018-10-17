package iso20022

// Specifies the number of occurrences of a particular event and the maximum number of times this event may occur.
type Limit1 struct {

	// Number of occurrences of a particular event.
	Current *Max3NumericText `xml:"Cur"`

	// Specifies the maximum number of times an event may occur.
	Limit *Max3NumericText `xml:"Lmt"`
}

func (l *Limit1) SetCurrent(value string) {
	l.Current = (*Max3NumericText)(&value)
}

func (l *Limit1) SetLimit(value string) {
	l.Limit = (*Max3NumericText)(&value)
}

// Specifies the minimum value of entries to be reported in the requested message.
type Limit2 struct {

	// Minimum transaction amount to be reported in the requested message.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether the floor limit applies to credit, to debit or to both credit and debit entries.
	CreditDebitIndicator *FloorLimitType1Code `xml:"CdtDbtInd"`
}

func (l *Limit2) SetAmount(value, currency string) {
	l.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (l *Limit2) SetCreditDebitIndicator(value string) {
	l.CreditDebitIndicator = (*FloorLimitType1Code)(&value)
}
