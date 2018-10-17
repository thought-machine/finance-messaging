package iso20022

// Information needed to process a currency exchange or conversion.
type CurrencyExchange3 struct {

	// Currency of the amount to be converted in a currency conversion.
	SourceCurrency *CurrencyCode `xml:"SrcCcy"`

	// Currency into which an amount is to be converted in a currency conversion.
	TargetCurrency *CurrencyCode `xml:"TrgtCcy,omitempty"`

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *CurrencyCode `xml:"UnitCcy,omitempty"`

	// Factor used for the conversion of an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	//
	// Usage: ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Unique and unambiguous identifier of the foreign exchange contract.
	ContractIdentification *Max35Text `xml:"CtrctId,omitempty"`

	// Date and time at which an exchange rate is quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`
}

func (c *CurrencyExchange3) SetSourceCurrency(value string) {
	c.SourceCurrency = (*CurrencyCode)(&value)
}

func (c *CurrencyExchange3) SetTargetCurrency(value string) {
	c.TargetCurrency = (*CurrencyCode)(&value)
}

func (c *CurrencyExchange3) SetUnitCurrency(value string) {
	c.UnitCurrency = (*CurrencyCode)(&value)
}

func (c *CurrencyExchange3) SetExchangeRate(value string) {
	c.ExchangeRate = (*BaseOneRate)(&value)
}

func (c *CurrencyExchange3) SetContractIdentification(value string) {
	c.ContractIdentification = (*Max35Text)(&value)
}

func (c *CurrencyExchange3) SetQuotationDate(value string) {
	c.QuotationDate = (*ISODateTime)(&value)
}

// Set of elements used to provide details of the currency exchange.
type CurrencyExchange5 struct {

	// Currency from which an amount is to be converted in a currency conversion.
	SourceCurrency *ActiveOrHistoricCurrencyCode `xml:"SrcCcy"`

	// Currency into which an amount is to be converted in a currency conversion.
	TargetCurrency *ActiveOrHistoricCurrencyCode `xml:"TrgtCcy,omitempty"`

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveOrHistoricCurrencyCode `xml:"UnitCcy,omitempty"`

	// Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	//
	// Usage: ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Unique identification to unambiguously identify the foreign exchange contract.
	ContractIdentification *Max35Text `xml:"CtrctId,omitempty"`

	// Date and time at which an exchange rate is quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`
}

func (c *CurrencyExchange5) SetSourceCurrency(value string) {
	c.SourceCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (c *CurrencyExchange5) SetTargetCurrency(value string) {
	c.TargetCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (c *CurrencyExchange5) SetUnitCurrency(value string) {
	c.UnitCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (c *CurrencyExchange5) SetExchangeRate(value string) {
	c.ExchangeRate = (*BaseOneRate)(&value)
}

func (c *CurrencyExchange5) SetContractIdentification(value string) {
	c.ContractIdentification = (*Max35Text)(&value)
}

func (c *CurrencyExchange5) SetQuotationDate(value string) {
	c.QuotationDate = (*ISODateTime)(&value)
}

// Describes the details of the currency exchange.
type CurrencyExchange6 struct {

	// Currency from which an amount is to be converted in a currency conversion.
	SourceCurrency *ActiveOrHistoricCurrencyCode `xml:"SrcCcy"`

	// Currency into which an amount is to be converted in a currency conversion.
	TargetCurrency *ActiveOrHistoricCurrencyCode `xml:"TrgtCcy"`

	// Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Unique identification to unambiguously identify the foreign exchange contract.
	Description *Max40Text `xml:"Desc,omitempty"`

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveOrHistoricCurrencyCode `xml:"UnitCcy,omitempty"`

	// Further information on the exchange rate.
	Comments *Max70Text `xml:"Cmnts,omitempty"`

	// Date and time at which an exchange rate is quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`
}

func (c *CurrencyExchange6) SetSourceCurrency(value string) {
	c.SourceCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (c *CurrencyExchange6) SetTargetCurrency(value string) {
	c.TargetCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (c *CurrencyExchange6) SetExchangeRate(value string) {
	c.ExchangeRate = (*BaseOneRate)(&value)
}

func (c *CurrencyExchange6) SetDescription(value string) {
	c.Description = (*Max40Text)(&value)
}

func (c *CurrencyExchange6) SetUnitCurrency(value string) {
	c.UnitCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (c *CurrencyExchange6) SetComments(value string) {
	c.Comments = (*Max70Text)(&value)
}

func (c *CurrencyExchange6) SetQuotationDate(value string) {
	c.QuotationDate = (*ISODateTime)(&value)
}
