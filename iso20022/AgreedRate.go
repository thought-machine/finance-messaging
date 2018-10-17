package iso20022

// Information needed to process a currency exchange or conversion.
type AgreedRate1 struct {

	// The value of one currency expressed in relation to another currency. ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *CurrencyCode `xml:"UnitCcy,omitempty"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *CurrencyCode `xml:"QtdCcy,omitempty"`
}

func (a *AgreedRate1) SetExchangeRate(value string) {
	a.ExchangeRate = (*BaseOneRate)(&value)
}

func (a *AgreedRate1) SetUnitCurrency(value string) {
	a.UnitCurrency = (*CurrencyCode)(&value)
}

func (a *AgreedRate1) SetQuotedCurrency(value string) {
	a.QuotedCurrency = (*CurrencyCode)(&value)
}

// Information needed to process a currency exchange or conversion.
type AgreedRate2 struct {

	// The value of one currency expressed in relation to another currency. ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *CurrencyCode `xml:"UnitCcy"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *CurrencyCode `xml:"QtdCcy"`
}

func (a *AgreedRate2) SetExchangeRate(value string) {
	a.ExchangeRate = (*BaseOneRate)(&value)
}

func (a *AgreedRate2) SetUnitCurrency(value string) {
	a.UnitCurrency = (*CurrencyCode)(&value)
}

func (a *AgreedRate2) SetQuotedCurrency(value string) {
	a.QuotedCurrency = (*CurrencyCode)(&value)
}

// Information needed to process a currency exchange or conversion.
type AgreedRate3 struct {

	// The value of one currency expressed in relation to another currency. ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveCurrencyCode `xml:"UnitCcy,omitempty"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *ActiveCurrencyCode `xml:"QtdCcy,omitempty"`
}

func (a *AgreedRate3) SetExchangeRate(value string) {
	a.ExchangeRate = (*BaseOneRate)(&value)
}

func (a *AgreedRate3) SetUnitCurrency(value string) {
	a.UnitCurrency = (*ActiveCurrencyCode)(&value)
}

func (a *AgreedRate3) SetQuotedCurrency(value string) {
	a.QuotedCurrency = (*ActiveCurrencyCode)(&value)
}
