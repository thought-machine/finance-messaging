package iso20022

// Conversion between the currency of a card acceptor and the currency of a card issuer, provided by a dedicated service provider. The currency conversion has to be accepted by the cardholder.
type CurrencyConversion1 struct {

	// Identification of the currency conversion operation for the service provider.
	CurrencyConversionIdentification *Max35Text `xml:"CcyConvsId,omitempty"`

	// Result of a requested currency conversion.
	Result *CurrencyConversionResponse1Code `xml:"Rslt"`

	// Plain text explaining the result of the  currency conversion request.
	ResponseReason *Max35Text `xml:"RspnRsn,omitempty"`

	// Currency into which the amount is converted (ISO 4217, 3 alphanumeric characters).
	TargetCurrency *CurrencyCode `xml:"TrgtCcy"`

	// Currency into which the amount is converted (ISO 4217, 3 numeric characters).
	TargetCurrencyNumeric *Exact3NumericText `xml:"TrgtCcyNmrc"`

	// Maximal number of digits after the decimal separator for target currency.
	TargetCurrencyDecimal *Number `xml:"TrgtCcyDcml"`

	// Full name of the target currency.
	TargetCurrencyName *Max35Text `xml:"TrgtCcyNm,omitempty"`

	// Amount converted in the target currency, including additional charges.
	ResultingAmount *ImpliedCurrencyAndAmount `xml:"RsltgAmt"`

	// Exchange rate, expressed as a percentage, applied to convert the original amount into the resulting amount.
	ExchangeRate *PercentageRate `xml:"XchgRate"`

	// Exchange rate, expressed as a percentage, applied to convert the resulting amount into the original amount.
	InvertedExchangeRate *PercentageRate `xml:"NvrtdXchgRate,omitempty"`

	// Date and time at which the exchange rate has been quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`

	// Validity limit of the exchange rate.
	ValidUntil *ISODateTime `xml:"VldUntil,omitempty"`

	// Currency from which the amount is converted (ISO 4217, 3 alphanumeric characters).
	SourceCurrency *CurrencyCode `xml:"SrcCcy"`

	// Currency from which the amount is converted (ISO 4217, 3 numeric characters).
	SourceCurrencyNumeric *CurrencyCode `xml:"SrcCcyNmrc,omitempty"`

	// Maximal number of digits after the decimal separator for source currency.
	SourceCurrencyDecimal *Number `xml:"SrcCcyDcml"`

	// Full name of the source currency.
	SourceCurrencyName *Max35Text `xml:"SrcCcyNm,omitempty"`

	// Original amount in the source currency.
	OriginalAmount *ImpliedCurrencyAndAmount `xml:"OrgnlAmt"`

	// Commission or additional charges made as part of a currency conversion.
	CommissionDetails []*Commission19 `xml:"ComssnDtls,omitempty"`

	// Markup made as part of a currency conversion.
	MarkUpDetails []*Commission18 `xml:"MrkUpDtls,omitempty"`

	// Card scheme declaration (disclaimer) to present to the cardholder.
	DeclarationDetails *Max2048Text `xml:"DclrtnDtls,omitempty"`
}

func (c *CurrencyConversion1) SetCurrencyConversionIdentification(value string) {
	c.CurrencyConversionIdentification = (*Max35Text)(&value)
}

func (c *CurrencyConversion1) SetResult(value string) {
	c.Result = (*CurrencyConversionResponse1Code)(&value)
}

func (c *CurrencyConversion1) SetResponseReason(value string) {
	c.ResponseReason = (*Max35Text)(&value)
}

func (c *CurrencyConversion1) SetTargetCurrency(value string) {
	c.TargetCurrency = (*CurrencyCode)(&value)
}

func (c *CurrencyConversion1) SetTargetCurrencyNumeric(value string) {
	c.TargetCurrencyNumeric = (*Exact3NumericText)(&value)
}

func (c *CurrencyConversion1) SetTargetCurrencyDecimal(value string) {
	c.TargetCurrencyDecimal = (*Number)(&value)
}

func (c *CurrencyConversion1) SetTargetCurrencyName(value string) {
	c.TargetCurrencyName = (*Max35Text)(&value)
}

func (c *CurrencyConversion1) SetResultingAmount(value, currency string) {
	c.ResultingAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CurrencyConversion1) SetExchangeRate(value string) {
	c.ExchangeRate = (*PercentageRate)(&value)
}

func (c *CurrencyConversion1) SetInvertedExchangeRate(value string) {
	c.InvertedExchangeRate = (*PercentageRate)(&value)
}

func (c *CurrencyConversion1) SetQuotationDate(value string) {
	c.QuotationDate = (*ISODateTime)(&value)
}

func (c *CurrencyConversion1) SetValidUntil(value string) {
	c.ValidUntil = (*ISODateTime)(&value)
}

func (c *CurrencyConversion1) SetSourceCurrency(value string) {
	c.SourceCurrency = (*CurrencyCode)(&value)
}

func (c *CurrencyConversion1) SetSourceCurrencyNumeric(value string) {
	c.SourceCurrencyNumeric = (*CurrencyCode)(&value)
}

func (c *CurrencyConversion1) SetSourceCurrencyDecimal(value string) {
	c.SourceCurrencyDecimal = (*Number)(&value)
}

func (c *CurrencyConversion1) SetSourceCurrencyName(value string) {
	c.SourceCurrencyName = (*Max35Text)(&value)
}

func (c *CurrencyConversion1) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CurrencyConversion1) AddCommissionDetails() *Commission19 {
	newValue := new(Commission19)
	c.CommissionDetails = append(c.CommissionDetails, newValue)
	return newValue
}

func (c *CurrencyConversion1) AddMarkUpDetails() *Commission18 {
	newValue := new(Commission18)
	c.MarkUpDetails = append(c.MarkUpDetails, newValue)
	return newValue
}

func (c *CurrencyConversion1) SetDeclarationDetails(value string) {
	c.DeclarationDetails = (*Max2048Text)(&value)
}

// Conversion between the currency of a card acceptor and the currency of a card issuer, provided by a dedicated service provider.
type CurrencyConversion10 struct {

	// Result of a requested currency conversion.
	Result *CurrencyConversionResponse2Code `xml:"Rslt"`

	// Plain text explaining the result of the  currency conversion request.
	ResultReason *Max35Text `xml:"RsltRsn,omitempty"`

	// Information about the conversion of currency.
	Conversion *CurrencyConversion9 `xml:"Convs,omitempty"`
}

func (c *CurrencyConversion10) SetResult(value string) {
	c.Result = (*CurrencyConversionResponse2Code)(&value)
}

func (c *CurrencyConversion10) SetResultReason(value string) {
	c.ResultReason = (*Max35Text)(&value)
}

func (c *CurrencyConversion10) AddConversion() *CurrencyConversion9 {
	c.Conversion = new(CurrencyConversion9)
	return c.Conversion
}

// Conversion between the currency of a card acceptor and the currency of a card issuer, provided by a dedicated service provider. The currency conversion has to be accepted by the cardholder.
type CurrencyConversion2 struct {

	// Identification of the currency conversion operation for the service provider.
	CurrencyConversionIdentification *Max35Text `xml:"CcyConvsId,omitempty"`

	// Currency into which the amount is converted (ISO 4217, 3 alphanumeric characters).
	TargetCurrency *CurrencyDetails1 `xml:"TrgtCcy"`

	// Amount converted in the target currency, including additional charges.
	ResultingAmount *ImpliedCurrencyAndAmount `xml:"RsltgAmt"`

	// Exchange rate, expressed as a percentage, applied to convert the original amount into the resulting amount.
	ExchangeRate *PercentageRate `xml:"XchgRate"`

	// Exchange rate expressed as a decimal, for example 0.7 is 7/10 and 70%.
	ExchangeRateDecimal *BaseOneRate `xml:"XchgRateDcml,omitempty"`

	// Exchange rate, expressed as a percentage, applied to convert the resulting amount into the original amount.
	InvertedExchangeRate *PercentageRate `xml:"NvrtdXchgRate,omitempty"`

	// Date and time at which the exchange rate has been quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`

	// Validity limit of the exchange rate.
	ValidUntil *ISODateTime `xml:"VldUntil,omitempty"`

	// Currency from which the amount is converted (ISO 4217, 3 alphanumeric characters).
	SourceCurrency *CurrencyDetails1 `xml:"SrcCcy"`

	// Original amount in the source currency.
	OriginalAmount *ImpliedCurrencyAndAmount `xml:"OrgnlAmt"`

	// Commission or additional charges made as part of a currency conversion.
	CommissionDetails []*Commission19 `xml:"ComssnDtls,omitempty"`

	// Markup made as part of a currency conversion.
	MarkUpDetails []*Commission18 `xml:"MrkUpDtls,omitempty"`

	// Card scheme declaration (disclaimer) to present to the cardholder.
	DeclarationDetails *Max2048Text `xml:"DclrtnDtls,omitempty"`
}

func (c *CurrencyConversion2) SetCurrencyConversionIdentification(value string) {
	c.CurrencyConversionIdentification = (*Max35Text)(&value)
}

func (c *CurrencyConversion2) AddTargetCurrency() *CurrencyDetails1 {
	c.TargetCurrency = new(CurrencyDetails1)
	return c.TargetCurrency
}

func (c *CurrencyConversion2) SetResultingAmount(value, currency string) {
	c.ResultingAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CurrencyConversion2) SetExchangeRate(value string) {
	c.ExchangeRate = (*PercentageRate)(&value)
}

func (c *CurrencyConversion2) SetExchangeRateDecimal(value string) {
	c.ExchangeRateDecimal = (*BaseOneRate)(&value)
}

func (c *CurrencyConversion2) SetInvertedExchangeRate(value string) {
	c.InvertedExchangeRate = (*PercentageRate)(&value)
}

func (c *CurrencyConversion2) SetQuotationDate(value string) {
	c.QuotationDate = (*ISODateTime)(&value)
}

func (c *CurrencyConversion2) SetValidUntil(value string) {
	c.ValidUntil = (*ISODateTime)(&value)
}

func (c *CurrencyConversion2) AddSourceCurrency() *CurrencyDetails1 {
	c.SourceCurrency = new(CurrencyDetails1)
	return c.SourceCurrency
}

func (c *CurrencyConversion2) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CurrencyConversion2) AddCommissionDetails() *Commission19 {
	newValue := new(Commission19)
	c.CommissionDetails = append(c.CommissionDetails, newValue)
	return newValue
}

func (c *CurrencyConversion2) AddMarkUpDetails() *Commission18 {
	newValue := new(Commission18)
	c.MarkUpDetails = append(c.MarkUpDetails, newValue)
	return newValue
}

func (c *CurrencyConversion2) SetDeclarationDetails(value string) {
	c.DeclarationDetails = (*Max2048Text)(&value)
}

// Conversion between the currency of a card acceptor and the currency of a card issuer, provided by a dedicated service provider.
type CurrencyConversion3 struct {

	// Result of a requested currency conversion.
	Result *CurrencyConversionResponse1Code `xml:"Rslt"`

	// Plain text explaining the result of the  currency conversion request.
	ResultReason *Max35Text `xml:"RsltRsn,omitempty"`

	// Information about the conversion of currency.
	Conversion *CurrencyConversion2 `xml:"Convs,omitempty"`
}

func (c *CurrencyConversion3) SetResult(value string) {
	c.Result = (*CurrencyConversionResponse1Code)(&value)
}

func (c *CurrencyConversion3) SetResultReason(value string) {
	c.ResultReason = (*Max35Text)(&value)
}

func (c *CurrencyConversion3) AddConversion() *CurrencyConversion2 {
	c.Conversion = new(CurrencyConversion2)
	return c.Conversion
}

// Currency conversion accepted by the customer, either to convert the amount to dispense in the base currency of the ATM, or to convert the total requested amount in the currency of the customer (so called dynamic currency conversion).
type CurrencyConversion4 struct {

	// Identification of the currency conversion operation.
	CurrencyConversionIdentification *Max35Text `xml:"CcyConvsId,omitempty"`

	// Currency into which the amount is converted (ISO 4217, 3 alphanumeric characters).
	TargetCurrency *CurrencyDetails2 `xml:"TrgtCcy"`

	// Amount converted in the target currency, including commission and mark-up.
	ResultingAmount *ImpliedCurrencyAndAmount `xml:"RsltgAmt"`

	// Exchange rate, expressed as a percentage, applied to convert the original amount into the resulting amount.
	ExchangeRate *PercentageRate `xml:"XchgRate"`

	// Exchange rate expressed as a decimal, for example 0.7 is 7/10 and 70%.
	ExchangeRateDecimal *BaseOneRate `xml:"XchgRateDcml,omitempty"`

	// Exchange rate, expressed as a percentage, applied to convert the resulting amount into the original amount.
	InvertedExchangeRate *PercentageRate `xml:"NvrtdXchgRate,omitempty"`

	// Date and time at which the exchange rate has been quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`

	// Validity limit of the exchange rate.
	ValidUntil *ISODateTime `xml:"VldUntil,omitempty"`

	// Currency from which the amount is converted (ISO 4217, 3 alphanumeric characters).
	SourceCurrency *CurrencyDetails2 `xml:"SrcCcy"`

	// Original amount in the source currency.
	OriginalAmount *ImpliedCurrencyAndAmount `xml:"OrgnlAmt"`

	// Commission or additional charges made as part of a currency conversion.
	CommissionDetails []*Commission19 `xml:"ComssnDtls,omitempty"`

	// Mark-up made as part of a currency conversion.
	MarkUpDetails []*Commission18 `xml:"MrkUpDtls,omitempty"`

	// Card scheme declaration (disclaimer) to present to the cardholder.
	DeclarationDetails *Max2048Text `xml:"DclrtnDtls,omitempty"`
}

func (c *CurrencyConversion4) SetCurrencyConversionIdentification(value string) {
	c.CurrencyConversionIdentification = (*Max35Text)(&value)
}

func (c *CurrencyConversion4) AddTargetCurrency() *CurrencyDetails2 {
	c.TargetCurrency = new(CurrencyDetails2)
	return c.TargetCurrency
}

func (c *CurrencyConversion4) SetResultingAmount(value, currency string) {
	c.ResultingAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CurrencyConversion4) SetExchangeRate(value string) {
	c.ExchangeRate = (*PercentageRate)(&value)
}

func (c *CurrencyConversion4) SetExchangeRateDecimal(value string) {
	c.ExchangeRateDecimal = (*BaseOneRate)(&value)
}

func (c *CurrencyConversion4) SetInvertedExchangeRate(value string) {
	c.InvertedExchangeRate = (*PercentageRate)(&value)
}

func (c *CurrencyConversion4) SetQuotationDate(value string) {
	c.QuotationDate = (*ISODateTime)(&value)
}

func (c *CurrencyConversion4) SetValidUntil(value string) {
	c.ValidUntil = (*ISODateTime)(&value)
}

func (c *CurrencyConversion4) AddSourceCurrency() *CurrencyDetails2 {
	c.SourceCurrency = new(CurrencyDetails2)
	return c.SourceCurrency
}

func (c *CurrencyConversion4) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CurrencyConversion4) AddCommissionDetails() *Commission19 {
	newValue := new(Commission19)
	c.CommissionDetails = append(c.CommissionDetails, newValue)
	return newValue
}

func (c *CurrencyConversion4) AddMarkUpDetails() *Commission18 {
	newValue := new(Commission18)
	c.MarkUpDetails = append(c.MarkUpDetails, newValue)
	return newValue
}

func (c *CurrencyConversion4) SetDeclarationDetails(value string) {
	c.DeclarationDetails = (*Max2048Text)(&value)
}

// Exchange rate and calculated amount to be presented to the customer when the dispense currency or the deposit currency (target currency) is different to account currency (source currency).
type CurrencyConversion5 struct {

	// Currency from which the amount is converted (ISO 4217, 3 alphanumeric characters).
	SourceCurrency *ActiveCurrencyCode `xml:"SrcCcy"`

	// Currency from which the amount is converted (ISO 4217, 3 numeric characters).
	SourceCurrencyNumeric *ActiveCurrencyCode `xml:"SrcCcyNmrc"`

	// Currency into which the amount is converted (ISO 4217, 3 alphanumeric characters).
	TargetCurrency *ActiveCurrencyCode `xml:"TrgtCcy"`

	// Currency into which the amount is converted (ISO 4217, 3 numeric characters).
	TargetCurrencyNumeric *Exact3NumericText `xml:"TrgtCcyNmrc"`

	// Currency exchange rate.
	Rate *BaseOneRate `xml:"Rate"`

	// Resulting calculated amount is in target currency.
	CalculatedAmount *ImpliedCurrencyAndAmount `xml:"ClctdAmt"`
}

func (c *CurrencyConversion5) SetSourceCurrency(value string) {
	c.SourceCurrency = (*ActiveCurrencyCode)(&value)
}

func (c *CurrencyConversion5) SetSourceCurrencyNumeric(value string) {
	c.SourceCurrencyNumeric = (*ActiveCurrencyCode)(&value)
}

func (c *CurrencyConversion5) SetTargetCurrency(value string) {
	c.TargetCurrency = (*ActiveCurrencyCode)(&value)
}

func (c *CurrencyConversion5) SetTargetCurrencyNumeric(value string) {
	c.TargetCurrencyNumeric = (*Exact3NumericText)(&value)
}

func (c *CurrencyConversion5) SetRate(value string) {
	c.Rate = (*BaseOneRate)(&value)
}

func (c *CurrencyConversion5) SetCalculatedAmount(value, currency string) {
	c.CalculatedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

// Conversion between the currency of a card acceptor and the currency of a card issuer, provided by a dedicated service provider. The currency conversion has to be accepted by the cardholder.
type CurrencyConversion6 struct {

	// Identification of the currency conversion operation for the service provider.
	CurrencyConversionIdentification *Max35Text `xml:"CcyConvsId,omitempty"`

	// Currency into which the amount is converted (ISO 4217, 3 alphanumeric characters).
	TargetCurrency *CurrencyDetails1 `xml:"TrgtCcy"`

	// Amount converted in the target currency, including additional charges.
	ResultingAmount *ImpliedCurrencyAndAmount `xml:"RsltgAmt"`

	// Exchange rate, expressed as a percentage, applied to convert the original amount into the resulting amount.
	ExchangeRate *PercentageRate `xml:"XchgRate"`

	// Exchange rate, expressed as a percentage, applied to convert the resulting amount into the original amount.
	InvertedExchangeRate *PercentageRate `xml:"NvrtdXchgRate,omitempty"`

	// Date and time at which the exchange rate has been quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`

	// Validity limit of the exchange rate.
	ValidUntil *ISODateTime `xml:"VldUntil,omitempty"`

	// Currency from which the amount is converted (ISO 4217, 3 alphanumeric characters).
	SourceCurrency *CurrencyDetails1 `xml:"SrcCcy"`

	// Original amount in the source currency.
	OriginalAmount *ImpliedCurrencyAndAmount `xml:"OrgnlAmt"`

	// Commission or additional charges made as part of a currency conversion.
	CommissionDetails []*Commission19 `xml:"ComssnDtls,omitempty"`

	// Markup made as part of a currency conversion.
	MarkUpDetails []*Commission18 `xml:"MrkUpDtls,omitempty"`

	// Card scheme declaration (disclaimer) to present to the cardholder.
	DeclarationDetails *ActionMessage5 `xml:"DclrtnDtls,omitempty"`
}

func (c *CurrencyConversion6) SetCurrencyConversionIdentification(value string) {
	c.CurrencyConversionIdentification = (*Max35Text)(&value)
}

func (c *CurrencyConversion6) AddTargetCurrency() *CurrencyDetails1 {
	c.TargetCurrency = new(CurrencyDetails1)
	return c.TargetCurrency
}

func (c *CurrencyConversion6) SetResultingAmount(value, currency string) {
	c.ResultingAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CurrencyConversion6) SetExchangeRate(value string) {
	c.ExchangeRate = (*PercentageRate)(&value)
}

func (c *CurrencyConversion6) SetInvertedExchangeRate(value string) {
	c.InvertedExchangeRate = (*PercentageRate)(&value)
}

func (c *CurrencyConversion6) SetQuotationDate(value string) {
	c.QuotationDate = (*ISODateTime)(&value)
}

func (c *CurrencyConversion6) SetValidUntil(value string) {
	c.ValidUntil = (*ISODateTime)(&value)
}

func (c *CurrencyConversion6) AddSourceCurrency() *CurrencyDetails1 {
	c.SourceCurrency = new(CurrencyDetails1)
	return c.SourceCurrency
}

func (c *CurrencyConversion6) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CurrencyConversion6) AddCommissionDetails() *Commission19 {
	newValue := new(Commission19)
	c.CommissionDetails = append(c.CommissionDetails, newValue)
	return newValue
}

func (c *CurrencyConversion6) AddMarkUpDetails() *Commission18 {
	newValue := new(Commission18)
	c.MarkUpDetails = append(c.MarkUpDetails, newValue)
	return newValue
}

func (c *CurrencyConversion6) AddDeclarationDetails() *ActionMessage5 {
	c.DeclarationDetails = new(ActionMessage5)
	return c.DeclarationDetails
}

// Conversion between the currency of a card acceptor and the currency of a card issuer, provided by a dedicated service provider.
type CurrencyConversion7 struct {

	// Result of a requested currency conversion.
	Result *CurrencyConversionResponse1Code `xml:"Rslt"`

	// Plain text explaining the result of the  currency conversion request.
	ResultReason *Max35Text `xml:"RsltRsn,omitempty"`

	// Information about the conversion of currency.
	ConversionDetails *CurrencyConversion6 `xml:"ConvsDtls,omitempty"`
}

func (c *CurrencyConversion7) SetResult(value string) {
	c.Result = (*CurrencyConversionResponse1Code)(&value)
}

func (c *CurrencyConversion7) SetResultReason(value string) {
	c.ResultReason = (*Max35Text)(&value)
}

func (c *CurrencyConversion7) AddConversionDetails() *CurrencyConversion6 {
	c.ConversionDetails = new(CurrencyConversion6)
	return c.ConversionDetails
}

// Conversion between the currency of a card acceptor and the currency of a card issuer, provided by a dedicated service provider.
type CurrencyConversion8 struct {

	// True if the cardholder has accepted the currency conversion that the acquirer has proposed.
	AcceptedByCardholder *TrueFalseIndicator `xml:"AccptdByCrdhldr,omitempty"`

	// Conversion between the currency of a card acceptor and the currency of a cardholder, provided by a dedicated service provider.
	Conversion *CurrencyConversion6 `xml:"Convs,omitempty"`
}

func (c *CurrencyConversion8) SetAcceptedByCardholder(value string) {
	c.AcceptedByCardholder = (*TrueFalseIndicator)(&value)
}

func (c *CurrencyConversion8) AddConversion() *CurrencyConversion6 {
	c.Conversion = new(CurrencyConversion6)
	return c.Conversion
}

// Currency conversion accepted by the customer, either to convert the amount to dispense in the base currency of the ATM, or to convert the total requested amount in the currency of the customer (so called dynamic currency conversion).
type CurrencyConversion9 struct {

	// Identification of the currency conversion operation.
	CurrencyConversionIdentification *Max35Text `xml:"CcyConvsId,omitempty"`

	// Currency into which the amount is converted (ISO 4217, 3 alphanumeric characters).
	TargetCurrency *CurrencyDetails2 `xml:"TrgtCcy"`

	// Amount converted in the target currency, including commission and mark-up.
	ResultingAmount *ImpliedCurrencyAndAmount `xml:"RsltgAmt"`

	// Exchange rate, expressed as a percentage, applied to convert the original amount into the resulting amount.
	ExchangeRate *PercentageRate `xml:"XchgRate"`

	// Exchange rate, expressed as a percentage, applied to convert the resulting amount into the original amount.
	InvertedExchangeRate *PercentageRate `xml:"NvrtdXchgRate,omitempty"`

	// Date and time at which the exchange rate has been quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`

	// Validity limit of the exchange rate.
	ValidUntil *ISODateTime `xml:"VldUntil,omitempty"`

	// Currency from which the amount is converted (ISO 4217, 3 alphanumeric characters).
	SourceCurrency *CurrencyDetails2 `xml:"SrcCcy"`

	// Original amount in the source currency.
	OriginalAmount *ImpliedCurrencyAndAmount `xml:"OrgnlAmt"`

	// Commission or additional charges made as part of a currency conversion.
	CommissionDetails []*Commission19 `xml:"ComssnDtls,omitempty"`

	// Mark-up made as part of a currency conversion.
	MarkUpDetails []*Commission18 `xml:"MrkUpDtls,omitempty"`

	// Card scheme declaration (disclaimer) to present to the cardholder.
	DeclarationDetails *ActionMessage5 `xml:"DclrtnDtls,omitempty"`
}

func (c *CurrencyConversion9) SetCurrencyConversionIdentification(value string) {
	c.CurrencyConversionIdentification = (*Max35Text)(&value)
}

func (c *CurrencyConversion9) AddTargetCurrency() *CurrencyDetails2 {
	c.TargetCurrency = new(CurrencyDetails2)
	return c.TargetCurrency
}

func (c *CurrencyConversion9) SetResultingAmount(value, currency string) {
	c.ResultingAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CurrencyConversion9) SetExchangeRate(value string) {
	c.ExchangeRate = (*PercentageRate)(&value)
}

func (c *CurrencyConversion9) SetInvertedExchangeRate(value string) {
	c.InvertedExchangeRate = (*PercentageRate)(&value)
}

func (c *CurrencyConversion9) SetQuotationDate(value string) {
	c.QuotationDate = (*ISODateTime)(&value)
}

func (c *CurrencyConversion9) SetValidUntil(value string) {
	c.ValidUntil = (*ISODateTime)(&value)
}

func (c *CurrencyConversion9) AddSourceCurrency() *CurrencyDetails2 {
	c.SourceCurrency = new(CurrencyDetails2)
	return c.SourceCurrency
}

func (c *CurrencyConversion9) SetOriginalAmount(value, currency string) {
	c.OriginalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CurrencyConversion9) AddCommissionDetails() *Commission19 {
	newValue := new(Commission19)
	c.CommissionDetails = append(c.CommissionDetails, newValue)
	return newValue
}

func (c *CurrencyConversion9) AddMarkUpDetails() *Commission18 {
	newValue := new(Commission18)
	c.MarkUpDetails = append(c.MarkUpDetails, newValue)
	return newValue
}

func (c *CurrencyConversion9) AddDeclarationDetails() *ActionMessage5 {
	c.DeclarationDetails = new(ActionMessage5)
	return c.DeclarationDetails
}
