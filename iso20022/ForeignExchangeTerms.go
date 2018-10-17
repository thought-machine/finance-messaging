package iso20022

// Information needed to process a currency exchange or conversion.
type ForeignExchangeTerms1 struct {

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveOrHistoricCurrencyCode `xml:"UnitCcy"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *ActiveOrHistoricCurrencyCode `xml:"QtdCcy"`

	// The value of one currency expressed in relation to another currency. ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Date and time at which an exchange rate is quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`

	// Party that proposes a  foreign exchange rate.
	QuotingInstitution *PartyIdentification10Choice `xml:"QtgInstn,omitempty"`
}

func (f *ForeignExchangeTerms1) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms1) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms1) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}

func (f *ForeignExchangeTerms1) SetQuotationDate(value string) {
	f.QuotationDate = (*ISODateTime)(&value)
}

func (f *ForeignExchangeTerms1) AddQuotingInstitution() *PartyIdentification10Choice {
	f.QuotingInstitution = new(PartyIdentification10Choice)
	return f.QuotingInstitution
}

// Provides information about the terms of the foreign exchange transaction.
type ForeignExchangeTerms11 struct {

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveCurrencyCode `xml:"UnitCcy"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *ActiveCurrencyCode `xml:"QtdCcy"`

	// Factor used for the conversion of an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	// Usage: ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Counter value of a foreign exchange conversion.
	ResultingAmount *ActiveCurrencyAndAmount `xml:"RsltgAmt"`
}

func (f *ForeignExchangeTerms11) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms11) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms11) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}

func (f *ForeignExchangeTerms11) SetResultingAmount(value, currency string) {
	f.ResultingAmount = NewActiveCurrencyAndAmount(value, currency)
}

// Provides information about the terms of the foreign exchange transaction.
type ForeignExchangeTerms13 struct {

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveCurrencyCode `xml:"UnitCcy"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *ActiveCurrencyCode `xml:"QtdCcy"`

	// Factor used for the conversion of an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	// Usage: ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Counter value of a foreign exchange conversion.
	ResultingAmount *ActiveCurrencyAndAmount `xml:"RsltgAmt,omitempty"`
}

func (f *ForeignExchangeTerms13) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms13) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms13) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}

func (f *ForeignExchangeTerms13) SetResultingAmount(value, currency string) {
	f.ResultingAmount = NewActiveCurrencyAndAmount(value, currency)
}

// Information needed to process a currency exchange or conversion.
type ForeignExchangeTerms14 struct {

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveOrHistoricCurrencyCode `xml:"UnitCcy"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *ActiveOrHistoricCurrencyCode `xml:"QtdCcy"`

	// The value of one currency expressed in relation to another currency. ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Date and time at which an exchange rate is quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`

	// Party that proposes a  foreign exchange rate.
	QuotingInstitution *PartyIdentification49Choice `xml:"QtgInstn,omitempty"`
}

func (f *ForeignExchangeTerms14) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms14) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms14) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}

func (f *ForeignExchangeTerms14) SetQuotationDate(value string) {
	f.QuotationDate = (*ISODateTime)(&value)
}

func (f *ForeignExchangeTerms14) AddQuotingInstitution() *PartyIdentification49Choice {
	f.QuotingInstitution = new(PartyIdentification49Choice)
	return f.QuotingInstitution
}

// Provides information about the terms of the foreign exchange transaction.
type ForeignExchangeTerms17 struct {

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxUSD, the unit currency is GBP.
	UnitCurrency *ActiveCurrencyCode `xml:"UnitCcy"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *ActiveCurrencyCode `xml:"QtdCcy"`

	// Factor used for the conversion of an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	// Usage: ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Counter value of a foreign exchange conversion.
	ResultingAmount *ActiveCurrencyAndAmount `xml:"RsltgAmt"`
}

func (f *ForeignExchangeTerms17) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms17) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms17) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}

func (f *ForeignExchangeTerms17) SetResultingAmount(value, currency string) {
	f.ResultingAmount = NewActiveCurrencyAndAmount(value, currency)
}

// Provides information about the terms of the foreign exchange transaction.
type ForeignExchangeTerms18 struct {

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveCurrencyCode `xml:"UnitCcy"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *ActiveCurrencyCode `xml:"QtdCcy"`

	// Factor used for the conversion of an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	// Usage: ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Amount following a foreign exchange conversion.
	ConvertedAmount *ActiveCurrencyAndAmount `xml:"ConvtdAmt"`
}

func (f *ForeignExchangeTerms18) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms18) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms18) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}

func (f *ForeignExchangeTerms18) SetConvertedAmount(value, currency string) {
	f.ConvertedAmount = NewActiveCurrencyAndAmount(value, currency)
}

// Provides information about the terms of the foreign exchange transaction.
type ForeignExchangeTerms19 struct {

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveCurrencyCode `xml:"UnitCcy"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *ActiveCurrencyCode `xml:"QtdCcy"`

	// Factor used for the conversion of an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	// Usage: ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`
}

func (f *ForeignExchangeTerms19) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms19) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms19) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}

// Information needed to process a currency exchange or conversion.
type ForeignExchangeTerms22 struct {

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveOrHistoricCurrencyCode `xml:"UnitCcy"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *ActiveOrHistoricCurrencyCode `xml:"QtdCcy"`

	// The value of one currency expressed in relation to another currency. ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Date and time at which an exchange rate is quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`

	// Party that proposes a  foreign exchange rate.
	QuotingInstitution *PartyIdentification71Choice `xml:"QtgInstn,omitempty"`
}

func (f *ForeignExchangeTerms22) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms22) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms22) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}

func (f *ForeignExchangeTerms22) SetQuotationDate(value string) {
	f.QuotationDate = (*ISODateTime)(&value)
}

func (f *ForeignExchangeTerms22) AddQuotingInstitution() *PartyIdentification71Choice {
	f.QuotingInstitution = new(PartyIdentification71Choice)
	return f.QuotingInstitution
}

// Provides information about the terms of the foreign exchange transaction.
type ForeignExchangeTerms23 struct {

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveCurrencyCode `xml:"UnitCcy"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *ActiveCurrencyCode `xml:"QtdCcy"`

	// Factor used for the conversion of an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	// Usage: ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Counter value of a foreign exchange conversion.
	ResultingAmount *ActiveCurrencyAndAmount `xml:"RsltgAmt"`
}

func (f *ForeignExchangeTerms23) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms23) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms23) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}

func (f *ForeignExchangeTerms23) SetResultingAmount(value, currency string) {
	f.ResultingAmount = NewActiveCurrencyAndAmount(value, currency)
}

// Provides information about the terms of the foreign exchange transaction.
type ForeignExchangeTerms24 struct {

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveCurrencyCode `xml:"UnitCcy"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *ActiveCurrencyCode `xml:"QtdCcy"`

	// Factor used for the conversion of an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	// Usage: ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Counter value of a foreign exchange conversion.
	ResultingAmount *ActiveCurrencyAndAmount `xml:"RsltgAmt,omitempty"`
}

func (f *ForeignExchangeTerms24) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms24) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms24) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}

func (f *ForeignExchangeTerms24) SetResultingAmount(value, currency string) {
	f.ResultingAmount = NewActiveCurrencyAndAmount(value, currency)
}

// Information needed to process a currency exchange or conversion.
type ForeignExchangeTerms26 struct {

	// Cash amount resulting from a foreign exchange trade.
	ToAmount *ActiveCurrencyAnd13DecimalAmount `xml:"ToAmt,omitempty"`

	// Cash amount for which a foreign exchange is required.
	FromAmount *ActiveCurrencyAndAmount `xml:"FrAmt,omitempty"`

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveOrHistoricCurrencyCode `xml:"UnitCcy"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *ActiveOrHistoricCurrencyCode `xml:"QtdCcy"`

	// The value of one currency expressed in relation to another currency. ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Date and time at which an exchange rate is quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`

	// Party that proposes a  foreign exchange rate.
	QuotingInstitution *PartyIdentification70Choice `xml:"QtgInstn,omitempty"`
}

func (f *ForeignExchangeTerms26) SetToAmount(value, currency string) {
	f.ToAmount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (f *ForeignExchangeTerms26) SetFromAmount(value, currency string) {
	f.FromAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *ForeignExchangeTerms26) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms26) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms26) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}

func (f *ForeignExchangeTerms26) SetQuotationDate(value string) {
	f.QuotationDate = (*ISODateTime)(&value)
}

func (f *ForeignExchangeTerms26) AddQuotingInstitution() *PartyIdentification70Choice {
	f.QuotingInstitution = new(PartyIdentification70Choice)
	return f.QuotingInstitution
}

// Provides information about the terms of the foreign exchange transaction.
type ForeignExchangeTerms27 struct {

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveCurrencyCode `xml:"UnitCcy"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *ActiveCurrencyCode `xml:"QtdCcy"`

	// Factor used for the conversion of an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	// Usage: ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Counter value of a foreign exchange conversion.
	ResultingAmount *RestrictedFINActiveCurrencyAndAmount `xml:"RsltgAmt"`
}

func (f *ForeignExchangeTerms27) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms27) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms27) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}

func (f *ForeignExchangeTerms27) SetResultingAmount(value, currency string) {
	f.ResultingAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

// Provides information about the terms of the foreign exchange transaction.
type ForeignExchangeTerms28 struct {

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveCurrencyCode `xml:"UnitCcy"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *ActiveCurrencyCode `xml:"QtdCcy"`

	// Factor used for the conversion of an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	// Usage: ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Counter value of a foreign exchange conversion.
	ResultingAmount *RestrictedFINActiveCurrencyAndAmount `xml:"RsltgAmt,omitempty"`
}

func (f *ForeignExchangeTerms28) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms28) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms28) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}

func (f *ForeignExchangeTerms28) SetResultingAmount(value, currency string) {
	f.ResultingAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

// Information needed to process a currency exchange or conversion.
type ForeignExchangeTerms3 struct {

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *CurrencyCode `xml:"UnitCcy"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *CurrencyCode `xml:"QtdCcy"`

	// The value of one currency expressed in relation to another currency. ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Date and time at which an exchange rate is quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`

	// Party that proposes a  foreign exchange rate.
	QuotingInstitution *PartyIdentification1Choice `xml:"QtgInstn,omitempty"`
}

func (f *ForeignExchangeTerms3) SetUnitCurrency(value string) {
	f.UnitCurrency = (*CurrencyCode)(&value)
}

func (f *ForeignExchangeTerms3) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*CurrencyCode)(&value)
}

func (f *ForeignExchangeTerms3) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}

func (f *ForeignExchangeTerms3) SetQuotationDate(value string) {
	f.QuotationDate = (*ISODateTime)(&value)
}

func (f *ForeignExchangeTerms3) AddQuotingInstitution() *PartyIdentification1Choice {
	f.QuotingInstitution = new(PartyIdentification1Choice)
	return f.QuotingInstitution
}

// Information needed to process a currency exchange or conversion.
type ForeignExchangeTerms31 struct {

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveOrHistoricCurrencyCode `xml:"UnitCcy"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *ActiveOrHistoricCurrencyCode `xml:"QtdCcy"`

	// The value of one currency expressed in relation to another currency. ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Date and time at which an exchange rate is quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`

	// Party that proposes a  foreign exchange rate.
	QuotingInstitution *PartyIdentification104Choice `xml:"QtgInstn,omitempty"`
}

func (f *ForeignExchangeTerms31) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms31) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms31) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}

func (f *ForeignExchangeTerms31) SetQuotationDate(value string) {
	f.QuotationDate = (*ISODateTime)(&value)
}

func (f *ForeignExchangeTerms31) AddQuotingInstitution() *PartyIdentification104Choice {
	f.QuotingInstitution = new(PartyIdentification104Choice)
	return f.QuotingInstitution
}

// Information needed to process a currency exchange or conversion.
type ForeignExchangeTerms32 struct {

	// Currency from which the quoted currency is converted in an exchange rate calculation.
	// 1 x <UnitCcy> = <XchgRate> x <QtdCcy>.
	UnitCurrency *ActiveCurrencyCode `xml:"UnitCcy"`

	// Currency into which the unit currency is converted in an exchange rate calculation.
	// 1 x <UnitCcy> = <XchgRate> x <QtdCcy>.
	QuotedCurrency *ActiveCurrencyCode `xml:"QtdCcy"`

	// Factor used for the conversion of an amount from one currency into another. This reflects that amount of the quoted currency that can be purchased with one unit of the unit currency , as follows:
	// 1 x CUR1 = nnn x CUR2,
	// where:
	// CUR1 is the unit currency
	// CUR2 is the quoted currency
	// nnn is the exchange rate.
	// 1 x <UnitCcy> = <XchgRate> x <QtdCcy>.
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Date and time at which an exchange rate is quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`

	// Party that proposes a foreign exchange rate.
	QuotingInstitution *PartyIdentification113 `xml:"QtgInstn,omitempty"`
}

func (f *ForeignExchangeTerms32) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms32) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms32) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}

func (f *ForeignExchangeTerms32) SetQuotationDate(value string) {
	f.QuotationDate = (*ISODateTime)(&value)
}

func (f *ForeignExchangeTerms32) AddQuotingInstitution() *PartyIdentification113 {
	f.QuotingInstitution = new(PartyIdentification113)
	return f.QuotingInstitution
}

// Information needed to process a currency exchange or conversion.
type ForeignExchangeTerms33 struct {

	// Cash amount resulting from a foreign exchange trade.
	ToAmount *ActiveCurrencyAndAmount `xml:"ToAmt,omitempty"`

	// Cash amount for which a foreign exchange is required.
	FromAmount *ActiveCurrencyAndAmount `xml:"FrAmt,omitempty"`

	// Currency from which the quoted currency is converted in an exchange rate calculation.
	// 1 x <UnitCcy> = <XchgRate> x <QtdCcy>.
	UnitCurrency *ActiveCurrencyCode `xml:"UnitCcy"`

	// Currency into which the unit currency is converted in an exchange rate calculation.
	// 1 x <UnitCcy> = <XchgRate> x <QtdCcy>.
	QuotedCurrency *ActiveCurrencyCode `xml:"QtdCcy"`

	// Factor used for the conversion of an amount from one currency into another. This reflects that amount of the quoted currency that can be purchased with one unit of the unit currency , as follows:
	// 1 x CUR1 = nnn x CUR2,
	// where:
	// CUR1 is the unit currency
	// CUR2 is the quoted currency
	// nnn is the exchange rate.
	// 1 x <UnitCcy> = <XchgRate> x <QtdCcy>.
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Date and time at which an exchange rate is quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`

	// Party that proposes the foreign exchange rate.
	QuotingInstitution *PartyIdentification113 `xml:"QtgInstn,omitempty"`
}

func (f *ForeignExchangeTerms33) SetToAmount(value, currency string) {
	f.ToAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *ForeignExchangeTerms33) SetFromAmount(value, currency string) {
	f.FromAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *ForeignExchangeTerms33) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms33) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms33) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}

func (f *ForeignExchangeTerms33) SetQuotationDate(value string) {
	f.QuotationDate = (*ISODateTime)(&value)
}

func (f *ForeignExchangeTerms33) AddQuotingInstitution() *PartyIdentification113 {
	f.QuotingInstitution = new(PartyIdentification113)
	return f.QuotingInstitution
}

// Information needed to process a currency exchange or conversion.
type ForeignExchangeTerms4 struct {

	// Currency and amount bought in a foreign exchange trade. The buy amount is received by the buyer.
	BuyAmount *ActiveCurrencyAnd13DecimalAmount `xml:"BuyAmt,omitempty"`

	// Currency and amount sold in a foreign exchange trade. The sold amount is delivered by the buyer.
	SellAmount *ActiveCurrencyAndAmount `xml:"SellAmt,omitempty"`

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *CurrencyCode `xml:"UnitCcy"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *CurrencyCode `xml:"QtdCcy"`

	// The value of one currency expressed in relation to another currency. ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Date and time at which an exchange rate is quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`

	// Party that proposes a  foreign exchange rate.
	QuotingInstitution *PartyIdentification2Choice `xml:"QtgInstn,omitempty"`
}

func (f *ForeignExchangeTerms4) SetBuyAmount(value, currency string) {
	f.BuyAmount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (f *ForeignExchangeTerms4) SetSellAmount(value, currency string) {
	f.SellAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *ForeignExchangeTerms4) SetUnitCurrency(value string) {
	f.UnitCurrency = (*CurrencyCode)(&value)
}

func (f *ForeignExchangeTerms4) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*CurrencyCode)(&value)
}

func (f *ForeignExchangeTerms4) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}

func (f *ForeignExchangeTerms4) SetQuotationDate(value string) {
	f.QuotationDate = (*ISODateTime)(&value)
}

func (f *ForeignExchangeTerms4) AddQuotingInstitution() *PartyIdentification2Choice {
	f.QuotingInstitution = new(PartyIdentification2Choice)
	return f.QuotingInstitution
}

// Information needed to process a currency exchange or conversion.
type ForeignExchangeTerms5 struct {

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *CurrencyCode `xml:"UnitCcy"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *CurrencyCode `xml:"QtdCcy"`

	// The value of one currency expressed in relation to another currency. ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Date and time at which an exchange rate is quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`

	// Party that proposes a  foreign exchange rate.
	QuotingInstitution *PartyIdentification2Choice `xml:"QtgInstn,omitempty"`
}

func (f *ForeignExchangeTerms5) SetUnitCurrency(value string) {
	f.UnitCurrency = (*CurrencyCode)(&value)
}

func (f *ForeignExchangeTerms5) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*CurrencyCode)(&value)
}

func (f *ForeignExchangeTerms5) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}

func (f *ForeignExchangeTerms5) SetQuotationDate(value string) {
	f.QuotationDate = (*ISODateTime)(&value)
}

func (f *ForeignExchangeTerms5) AddQuotingInstitution() *PartyIdentification2Choice {
	f.QuotingInstitution = new(PartyIdentification2Choice)
	return f.QuotingInstitution
}

// Information needed to process a currency exchange or conversion.
type ForeignExchangeTerms6 struct {

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveOrHistoricCurrencyCode `xml:"UnitCcy"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *ActiveOrHistoricCurrencyCode `xml:"QtdCcy"`

	// The value of one currency expressed in relation to another currency. ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Date and time at which an exchange rate is quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`

	// Party that proposes a  foreign exchange rate.
	QuotingInstitution *PartyIdentification2Choice `xml:"QtgInstn,omitempty"`
}

func (f *ForeignExchangeTerms6) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms6) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms6) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}

func (f *ForeignExchangeTerms6) SetQuotationDate(value string) {
	f.QuotationDate = (*ISODateTime)(&value)
}

func (f *ForeignExchangeTerms6) AddQuotingInstitution() *PartyIdentification2Choice {
	f.QuotingInstitution = new(PartyIdentification2Choice)
	return f.QuotingInstitution
}

// Information needed to process a currency exchange or conversion.
type ForeignExchangeTerms7 struct {

	// Cash amount resulting from a foreign exchange trade.
	ToAmount *ActiveCurrencyAnd13DecimalAmount `xml:"ToAmt,omitempty"`

	// Cash amount for which a foreign exchange is required.
	FromAmount *ActiveCurrencyAndAmount `xml:"FrAmt,omitempty"`

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveOrHistoricCurrencyCode `xml:"UnitCcy"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *ActiveOrHistoricCurrencyCode `xml:"QtdCcy"`

	// The value of one currency expressed in relation to another currency. ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Date and time at which an exchange rate is quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`

	// Party that proposes a  foreign exchange rate.
	QuotingInstitution *PartyIdentification2Choice `xml:"QtgInstn,omitempty"`
}

func (f *ForeignExchangeTerms7) SetToAmount(value, currency string) {
	f.ToAmount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (f *ForeignExchangeTerms7) SetFromAmount(value, currency string) {
	f.FromAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *ForeignExchangeTerms7) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms7) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms7) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}

func (f *ForeignExchangeTerms7) SetQuotationDate(value string) {
	f.QuotationDate = (*ISODateTime)(&value)
}

func (f *ForeignExchangeTerms7) AddQuotingInstitution() *PartyIdentification2Choice {
	f.QuotingInstitution = new(PartyIdentification2Choice)
	return f.QuotingInstitution
}

// Provides information about the terms of the foreign exchange transaction.
type ForeignExchangeTerms8 struct {

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveCurrencyCode `xml:"UnitCcy"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *ActiveCurrencyCode `xml:"QtdCcy"`

	// Factor used for the conversion of an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	//
	// Usage: ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Amount of money resulting from a foreign exchange transaction.
	ResultingAmount *ActiveCurrencyAndAmount `xml:"RsltgAmt"`
}

func (f *ForeignExchangeTerms8) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms8) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms8) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}

func (f *ForeignExchangeTerms8) SetResultingAmount(value, currency string) {
	f.ResultingAmount = NewActiveCurrencyAndAmount(value, currency)
}

// Provides information about the terms of the foreign exchange transaction.
type ForeignExchangeTerms9 struct {

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveCurrencyCode `xml:"UnitCcy"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *ActiveCurrencyCode `xml:"QtdCcy"`

	// Factor used for the conversion of an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	//
	// Usage: ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Amount of money resulting from a foreign exchange transaction.
	ResultingAmount *ActiveCurrencyAndAmount `xml:"RsltgAmt"`

	// Amount in its original currency when conversion from/into another currency has occurred.
	OriginalAmount *ActiveCurrencyAndAmount `xml:"OrgnlAmt,omitempty"`
}

func (f *ForeignExchangeTerms9) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms9) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *ForeignExchangeTerms9) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}

func (f *ForeignExchangeTerms9) SetResultingAmount(value, currency string) {
	f.ResultingAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *ForeignExchangeTerms9) SetOriginalAmount(value, currency string) {
	f.OriginalAmount = NewActiveCurrencyAndAmount(value, currency)
}
