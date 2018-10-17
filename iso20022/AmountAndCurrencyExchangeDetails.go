package iso20022

// Set of elements providing information on the original amount and currency information.
type AmountAndCurrencyExchangeDetails1 struct {

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Reports on currency exchange information.
	CurrencyExchange *CurrencyExchange3 `xml:"CcyXchg,omitempty"`
}

func (a *AmountAndCurrencyExchangeDetails1) SetAmount(value, currency string) {
	a.Amount = NewCurrencyAndAmount(value, currency)
}

func (a *AmountAndCurrencyExchangeDetails1) AddCurrencyExchange() *CurrencyExchange3 {
	a.CurrencyExchange = new(CurrencyExchange3)
	return a.CurrencyExchange
}

// Set of elements providing information on the original amount and currency information.
type AmountAndCurrencyExchangeDetails2 struct {

	// Identifies the type of amount.
	Type *Max35Text `xml:"Tp"`

	// Identifies the proprietary amount.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Reports on currency exchange information.
	CurrencyExchange *CurrencyExchange3 `xml:"CcyXchg,omitempty"`
}

func (a *AmountAndCurrencyExchangeDetails2) SetType(value string) {
	a.Type = (*Max35Text)(&value)
}

func (a *AmountAndCurrencyExchangeDetails2) SetAmount(value, currency string) {
	a.Amount = NewCurrencyAndAmount(value, currency)
}

func (a *AmountAndCurrencyExchangeDetails2) AddCurrencyExchange() *CurrencyExchange3 {
	a.CurrencyExchange = new(CurrencyExchange3)
	return a.CurrencyExchange
}

// Set of elements used to provide information on the original amount and currency exchange.
type AmountAndCurrencyExchangeDetails3 struct {

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Set of elements used to provide details on the currency exchange.
	CurrencyExchange *CurrencyExchange5 `xml:"CcyXchg,omitempty"`
}

func (a *AmountAndCurrencyExchangeDetails3) SetAmount(value, currency string) {
	a.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndCurrencyExchangeDetails3) AddCurrencyExchange() *CurrencyExchange5 {
	a.CurrencyExchange = new(CurrencyExchange5)
	return a.CurrencyExchange
}

// Set of elements used to provide information on the original amount and currency exchange.
type AmountAndCurrencyExchangeDetails4 struct {

	// Specifies the type of amount.
	Type *Max35Text `xml:"Tp"`

	// Amount of money to be exchanged against another amount of money in the counter currency.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Set of elements used to provide details on the currency exchange.
	CurrencyExchange *CurrencyExchange5 `xml:"CcyXchg,omitempty"`
}

func (a *AmountAndCurrencyExchangeDetails4) SetType(value string) {
	a.Type = (*Max35Text)(&value)
}

func (a *AmountAndCurrencyExchangeDetails4) SetAmount(value, currency string) {
	a.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndCurrencyExchangeDetails4) AddCurrencyExchange() *CurrencyExchange5 {
	a.CurrencyExchange = new(CurrencyExchange5)
	return a.CurrencyExchange
}
