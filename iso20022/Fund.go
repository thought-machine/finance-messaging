package iso20022

// Information about an investment fund.
type Fund1 struct {

	// Name of the fund/sub fund.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Identification of the fund/sub fund with a Legal Entity Identifier. This is a code allocated to a party as described in ISO 17442 "Financial Services - Legal Entity Identifier (LEI)".
	LegalEntityIdentifier *LEIIdentifier `xml:"LglNttyIdr,omitempty"`

	// Unique and unambiguous identifier for the fund/sub fund, assigned under a formal or proprietary identification scheme.
	Identification *OtherIdentification4 `xml:"Id,omitempty"`

	// Currency of the fund/sub fund.
	Currency *ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty"`

	// Date and, if required, the time, at which the price will be applied.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm,omitempty"`

	// Previous date and time at which a price was applied.
	PreviousTradeDateTime *DateAndDateTimeChoice `xml:"PrvsTradDtTm,omitempty"`

	// Estimated total value of all the holdings, less the fund's liabilities, of the fund/sub fund.
	EstimatedTotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"EstmtdTtlNAV,omitempty"`

	// Previous total value of all the holdings, less the fund's liabilities, of the fund/sub fund.
	PreviousTotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"PrvsTtlNAV,omitempty"`

	// Estimated total number of units of the fund/sub fund.
	EstimatedTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"EstmtdTtlUnitsNb,omitempty"`

	// Previous total number of units of the fund/sub fund.
	PreviousTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"PrvsTtlUnitsNb,omitempty"`

	// Estimated consolidated net cash flow expressed as a percentage of the previous total NAV for the fund/sub fund.
	EstimatedPercentageOfFundTotalNAV *PercentageRate `xml:"EstmtdPctgOfFndTtlNAV,omitempty"`

	// Estimated cash movement into the fund/sub fund.
	EstimatedCashInForecastDetails []*CashInOutForecast7 `xml:"EstmtdCshInFcstDtls,omitempty"`

	// Estimated cash movement out of the fund/sub fund.
	EstimatedCashOutForecastDetails []*CashInOutForecast7 `xml:"EstmtdCshOutFcstDtls,omitempty"`

	// Net cash as a result of the cash-in and cash-out flows.
	EstimatedNetCashForecastDetails []*NetCashForecast5 `xml:"EstmtdNetCshFcstDtls,omitempty"`
}

func (f *Fund1) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *Fund1) SetLegalEntityIdentifier(value string) {
	f.LegalEntityIdentifier = (*LEIIdentifier)(&value)
}

func (f *Fund1) AddIdentification() *OtherIdentification4 {
	f.Identification = new(OtherIdentification4)
	return f.Identification
}

func (f *Fund1) SetCurrency(value string) {
	f.Currency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *Fund1) AddTradeDateTime() *DateAndDateTimeChoice {
	f.TradeDateTime = new(DateAndDateTimeChoice)
	return f.TradeDateTime
}

func (f *Fund1) AddPreviousTradeDateTime() *DateAndDateTimeChoice {
	f.PreviousTradeDateTime = new(DateAndDateTimeChoice)
	return f.PreviousTradeDateTime
}

func (f *Fund1) SetEstimatedTotalNAV(value, currency string) {
	f.EstimatedTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *Fund1) SetPreviousTotalNAV(value, currency string) {
	f.PreviousTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *Fund1) AddEstimatedTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.EstimatedTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.EstimatedTotalUnitsNumber
}

func (f *Fund1) AddPreviousTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.PreviousTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.PreviousTotalUnitsNumber
}

func (f *Fund1) SetEstimatedPercentageOfFundTotalNAV(value string) {
	f.EstimatedPercentageOfFundTotalNAV = (*PercentageRate)(&value)
}

func (f *Fund1) AddEstimatedCashInForecastDetails() *CashInOutForecast7 {
	newValue := new(CashInOutForecast7)
	f.EstimatedCashInForecastDetails = append(f.EstimatedCashInForecastDetails, newValue)
	return newValue
}

func (f *Fund1) AddEstimatedCashOutForecastDetails() *CashInOutForecast7 {
	newValue := new(CashInOutForecast7)
	f.EstimatedCashOutForecastDetails = append(f.EstimatedCashOutForecastDetails, newValue)
	return newValue
}

func (f *Fund1) AddEstimatedNetCashForecastDetails() *NetCashForecast5 {
	newValue := new(NetCashForecast5)
	f.EstimatedNetCashForecastDetails = append(f.EstimatedNetCashForecastDetails, newValue)
	return newValue
}

// Information about an investment fund.
type Fund2 struct {

	// Name of the fund/sub fund.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Identification of the fund/sub fund with a Legal Entity Identifier. This is a code allocated to a party as described in ISO 17442 "Financial Services - Legal Entity Identifier (LEI)".
	LegalEntityIdentifier *LEIIdentifier `xml:"LglNttyIdr,omitempty"`

	// Unique and unambiguous identifier for the fund/sub fund, assigned under a formal or proprietary identification scheme.
	Identification *OtherIdentification4 `xml:"Id,omitempty"`

	// Currency of the fund/sub fund.
	Currency *ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty"`

	// Date and, if required, the time, at which the price will be applied.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm,omitempty"`

	// Previous date and time at which the price was applied.
	PreviousTradeDateTime *DateAndDateTimeChoice `xml:"PrvsTradDtTm,omitempty"`

	// Total value of all the holdings, less the fund's liabilities, of the fund/sub fund.
	TotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"TtlNAV,omitempty"`

	// Previous total value of all the holdings, less the fund's liabilities, of the fund/sub fund.
	PreviousTotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"PrvsTtlNAV,omitempty"`

	// Total number of units of the fund/sub fund.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb,omitempty"`

	// Previous total number of units of the fund/sub fund.
	PreviousTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"PrvsTtlUnitsNb,omitempty"`

	// Consolidated net cash flow expressed as a percentage of the total NAV for the fund/sub fund.
	PercentageOfFundTotalNAV *PercentageRate `xml:"PctgOfFndTtlNAV,omitempty"`

	// Cash movement into the fund/sub fund.
	CashInForecastDetails []*CashInOutForecast7 `xml:"CshInFcstDtls,omitempty"`

	// Cash movement out of the fund/sub fund.
	CashOutForecastDetails []*CashInOutForecast7 `xml:"CshOutFcstDtls,omitempty"`

	// Net cash as a result of the cash-in and cash-out flows.
	NetCashForecastDetails []*NetCashForecast5 `xml:"NetCshFcstDtls,omitempty"`
}

func (f *Fund2) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *Fund2) SetLegalEntityIdentifier(value string) {
	f.LegalEntityIdentifier = (*LEIIdentifier)(&value)
}

func (f *Fund2) AddIdentification() *OtherIdentification4 {
	f.Identification = new(OtherIdentification4)
	return f.Identification
}

func (f *Fund2) SetCurrency(value string) {
	f.Currency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *Fund2) AddTradeDateTime() *DateAndDateTimeChoice {
	f.TradeDateTime = new(DateAndDateTimeChoice)
	return f.TradeDateTime
}

func (f *Fund2) AddPreviousTradeDateTime() *DateAndDateTimeChoice {
	f.PreviousTradeDateTime = new(DateAndDateTimeChoice)
	return f.PreviousTradeDateTime
}

func (f *Fund2) SetTotalNAV(value, currency string) {
	f.TotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *Fund2) SetPreviousTotalNAV(value, currency string) {
	f.PreviousTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *Fund2) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.TotalUnitsNumber
}

func (f *Fund2) AddPreviousTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.PreviousTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.PreviousTotalUnitsNumber
}

func (f *Fund2) SetPercentageOfFundTotalNAV(value string) {
	f.PercentageOfFundTotalNAV = (*PercentageRate)(&value)
}

func (f *Fund2) AddCashInForecastDetails() *CashInOutForecast7 {
	newValue := new(CashInOutForecast7)
	f.CashInForecastDetails = append(f.CashInForecastDetails, newValue)
	return newValue
}

func (f *Fund2) AddCashOutForecastDetails() *CashInOutForecast7 {
	newValue := new(CashInOutForecast7)
	f.CashOutForecastDetails = append(f.CashOutForecastDetails, newValue)
	return newValue
}

func (f *Fund2) AddNetCashForecastDetails() *NetCashForecast5 {
	newValue := new(NetCashForecast5)
	f.NetCashForecastDetails = append(f.NetCashForecastDetails, newValue)
	return newValue
}

// Information about an investment fund.
type Fund3 struct {

	// Name of the fund/sub fund.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Identification of the fund/sub fund with a Legal Entity Identifier. This is a code allocated to a party as described in ISO 17442 "Financial Services - Legal Entity Identifier (LEI)".
	LegalEntityIdentifier *LEIIdentifier `xml:"LglNttyIdr,omitempty"`

	// Unique and unambiguous identifier for the fund/sub fund, assigned under a formal or proprietary identification scheme.
	Identification *OtherIdentification4 `xml:"Id,omitempty"`

	// Currency of the fund/sub fund.
	Currency *ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty"`

	// Estimated total value of all the holdings, less the fund's liabilities, of the fund/sub fund.
	EstimatedTotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"EstmtdTtlNAV,omitempty"`

	// Previous total value of all the holdings, less the fund's liabilities, of the fund/sub fund
	PreviousTotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"PrvsTtlNAV,omitempty"`

	// Estimated total number of units of the fund/sub fund.
	EstimatedTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"EstmtdTtlUnitsNb,omitempty"`

	// Previous total number of units of the fund/sub fund.
	PreviousTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"PrvsTtlUnitsNb,omitempty"`

	// Estimated consolidated net cash flow expressed as a percentage of the previous total NAV for the fund/sub fund.
	EstimatedPercentageOfFundTotalNAV *PercentageRate `xml:"EstmtdPctgOfFndTtlNAV,omitempty"`
}

func (f *Fund3) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *Fund3) SetLegalEntityIdentifier(value string) {
	f.LegalEntityIdentifier = (*LEIIdentifier)(&value)
}

func (f *Fund3) AddIdentification() *OtherIdentification4 {
	f.Identification = new(OtherIdentification4)
	return f.Identification
}

func (f *Fund3) SetCurrency(value string) {
	f.Currency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *Fund3) SetEstimatedTotalNAV(value, currency string) {
	f.EstimatedTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *Fund3) SetPreviousTotalNAV(value, currency string) {
	f.PreviousTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *Fund3) AddEstimatedTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.EstimatedTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.EstimatedTotalUnitsNumber
}

func (f *Fund3) AddPreviousTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.PreviousTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.PreviousTotalUnitsNumber
}

func (f *Fund3) SetEstimatedPercentageOfFundTotalNAV(value string) {
	f.EstimatedPercentageOfFundTotalNAV = (*PercentageRate)(&value)
}

// Information about an investment fund.
type Fund4 struct {

	// Name of the fund/sub fund.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Identification of the fund/sub fund with a Legal Entity Identifier. This is a code allocated to a party as described in ISO 17442 "Financial Services - Legal Entity Identifier (LEI)".
	LegalEntityIdentifier *LEIIdentifier `xml:"LglNttyIdr,omitempty"`

	// Unique and unambiguous identifier for the fund/sub fund, assigned under a formal or proprietary identification scheme.
	Identification *OtherIdentification4 `xml:"Id,omitempty"`

	// Currency of the fund/sub fund.
	Currency *ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty"`

	// Total value of all the holdings, less the fund's liabilities, of the fund/sub fund.
	TotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"TtlNAV,omitempty"`

	// Previous total value of all the holdings, less the fund's liabilities, of the fund/sub fund.
	PreviousTotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"PrvsTtlNAV,omitempty"`

	// Total number of units of the fund/sub fund.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb,omitempty"`

	// Previous total number of units of the fund/sub fund.
	PreviousTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"PrvsTtlUnitsNb,omitempty"`

	// Consolidated net cash flow expressed as a percentage of the total NAV for the fund/sub fund.
	PercentageOfFundTotalNAV *PercentageRate `xml:"PctgOfFndTtlNAV,omitempty"`
}

func (f *Fund4) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *Fund4) SetLegalEntityIdentifier(value string) {
	f.LegalEntityIdentifier = (*LEIIdentifier)(&value)
}

func (f *Fund4) AddIdentification() *OtherIdentification4 {
	f.Identification = new(OtherIdentification4)
	return f.Identification
}

func (f *Fund4) SetCurrency(value string) {
	f.Currency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *Fund4) SetTotalNAV(value, currency string) {
	f.TotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *Fund4) SetPreviousTotalNAV(value, currency string) {
	f.PreviousTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *Fund4) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.TotalUnitsNumber
}

func (f *Fund4) AddPreviousTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.PreviousTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.PreviousTotalUnitsNumber
}

func (f *Fund4) SetPercentageOfFundTotalNAV(value string) {
	f.PercentageOfFundTotalNAV = (*PercentageRate)(&value)
}
