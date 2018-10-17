package iso20022

// Cash movements from or to a fund as a result of investment funds transactions, eg, subscriptions or redemptions.
type FundCashForecast1 struct {

	// Date and, if required, the time, at which the price has been applied.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Previous date and time at which a price was applied.
	PreviousTradeDateTime *DateAndDateTimeChoice `xml:"PrvsTradDtTm,omitempty"`

	// Investment fund class to which a cash flow is related.
	FinancialInstrumentDetails *FinancialInstrument5 `xml:"FinInstrmDtls"`

	// Total value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	TotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"TtlNAV,omitempty"`

	// Previous value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	PreviousTotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"PrvsTtlNAV,omitempty"`

	// Total number of investment fund class units that have been issued.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb,omitempty"`

	// Previous total number of investment fund class units that have been issued.
	PreviousTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"PrvsTtlUnitsNb,omitempty"`

	// Rate of change of the net asset value.
	TotalNAVChangeRate *PercentageRate `xml:"TtlNAVChngRate,omitempty"`

	// Currency of the investment fund class.
	InvestmentCurrency []*ActiveOrHistoricCurrencyCode `xml:"InvstmtCcy,omitempty"`

	// Indicates whether the net cash flow is exceptional.
	ExceptionalNetCashFlowIndicator *YesNoIndicator `xml:"XcptnlNetCshFlowInd"`

	// Cash movements into a fund as a result of investment funds transactions, eg, subscriptions or switch-in.
	CashInForecastDetails []*CashInForecast2 `xml:"CshInFcstDtls,omitempty"`

	// Cash movements out of a fund as a result of investment funds transactions, eg, redemptions or switch-out.
	CashOutForecastDetails []*CashOutForecast2 `xml:"CshOutFcstDtls,omitempty"`

	// Cash movements from or to a fund as a result of investment funds transactions.
	NetCashForecastDetails []*NetCashForecast1 `xml:"NetCshFcstDtls,omitempty"`
}

func (f *FundCashForecast1) AddTradeDateTime() *DateAndDateTimeChoice {
	f.TradeDateTime = new(DateAndDateTimeChoice)
	return f.TradeDateTime
}

func (f *FundCashForecast1) AddPreviousTradeDateTime() *DateAndDateTimeChoice {
	f.PreviousTradeDateTime = new(DateAndDateTimeChoice)
	return f.PreviousTradeDateTime
}

func (f *FundCashForecast1) AddFinancialInstrumentDetails() *FinancialInstrument5 {
	f.FinancialInstrumentDetails = new(FinancialInstrument5)
	return f.FinancialInstrumentDetails
}

func (f *FundCashForecast1) SetTotalNAV(value, currency string) {
	f.TotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundCashForecast1) SetPreviousTotalNAV(value, currency string) {
	f.PreviousTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundCashForecast1) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.TotalUnitsNumber
}

func (f *FundCashForecast1) AddPreviousTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.PreviousTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.PreviousTotalUnitsNumber
}

func (f *FundCashForecast1) SetTotalNAVChangeRate(value string) {
	f.TotalNAVChangeRate = (*PercentageRate)(&value)
}

func (f *FundCashForecast1) AddInvestmentCurrency(value string) {
	f.InvestmentCurrency = append(f.InvestmentCurrency, (*ActiveOrHistoricCurrencyCode)(&value))
}

func (f *FundCashForecast1) SetExceptionalNetCashFlowIndicator(value string) {
	f.ExceptionalNetCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (f *FundCashForecast1) AddCashInForecastDetails() *CashInForecast2 {
	newValue := new(CashInForecast2)
	f.CashInForecastDetails = append(f.CashInForecastDetails, newValue)
	return newValue
}

func (f *FundCashForecast1) AddCashOutForecastDetails() *CashOutForecast2 {
	newValue := new(CashOutForecast2)
	f.CashOutForecastDetails = append(f.CashOutForecastDetails, newValue)
	return newValue
}

func (f *FundCashForecast1) AddNetCashForecastDetails() *NetCashForecast1 {
	newValue := new(NetCashForecast1)
	f.NetCashForecastDetails = append(f.NetCashForecastDetails, newValue)
	return newValue
}

// Cash movements from or to a fund as a result of investment funds transactions, eg, subscriptions or redemptions.
type FundCashForecast2 struct {

	// Date and, if required, the time, at which the price has been applied.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Previous date and time at which a price was applied.
	PreviousTradeDateTime *DateAndDateTimeChoice `xml:"PrvsTradDtTm,omitempty"`

	// Investment fund class to which a cash flow is related.
	FinancialInstrumentDetails *FinancialInstrument5 `xml:"FinInstrmDtls"`

	// Total value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	TotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"TtlNAV,omitempty"`

	// Previous value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	PreviousTotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"PrvsTtlNAV,omitempty"`

	// Total number of investment fund class units that have been issued.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb,omitempty"`

	// Previous total number of investment fund class units that have been issued.
	PreviousTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"PrvsTtlUnitsNb,omitempty"`

	// Rate of change of the net asset value.
	TotalNAVChangeRate *PercentageRate `xml:"TtlNAVChngRate,omitempty"`

	// Currency of the investment fund class.
	InvestmentCurrency []*ActiveOrHistoricCurrencyCode `xml:"InvstmtCcy,omitempty"`

	// Indicates whether the net cash flow is exceptional.
	ExceptionalNetCashFlowIndicator *YesNoIndicator `xml:"XcptnlNetCshFlowInd"`

	// Information related to the cash movements reported by pre-defined or user defined criteria.
	SortingCriteriaDetails []*CashSortingCriterion2 `xml:"SrtgCritDtls"`

	// Net cash movements per financial instrument.
	NetCashForecastDetails []*NetCashForecast1 `xml:"NetCshFcstDtls,omitempty"`
}

func (f *FundCashForecast2) AddTradeDateTime() *DateAndDateTimeChoice {
	f.TradeDateTime = new(DateAndDateTimeChoice)
	return f.TradeDateTime
}

func (f *FundCashForecast2) AddPreviousTradeDateTime() *DateAndDateTimeChoice {
	f.PreviousTradeDateTime = new(DateAndDateTimeChoice)
	return f.PreviousTradeDateTime
}

func (f *FundCashForecast2) AddFinancialInstrumentDetails() *FinancialInstrument5 {
	f.FinancialInstrumentDetails = new(FinancialInstrument5)
	return f.FinancialInstrumentDetails
}

func (f *FundCashForecast2) SetTotalNAV(value, currency string) {
	f.TotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundCashForecast2) SetPreviousTotalNAV(value, currency string) {
	f.PreviousTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundCashForecast2) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.TotalUnitsNumber
}

func (f *FundCashForecast2) AddPreviousTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.PreviousTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.PreviousTotalUnitsNumber
}

func (f *FundCashForecast2) SetTotalNAVChangeRate(value string) {
	f.TotalNAVChangeRate = (*PercentageRate)(&value)
}

func (f *FundCashForecast2) AddInvestmentCurrency(value string) {
	f.InvestmentCurrency = append(f.InvestmentCurrency, (*ActiveOrHistoricCurrencyCode)(&value))
}

func (f *FundCashForecast2) SetExceptionalNetCashFlowIndicator(value string) {
	f.ExceptionalNetCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (f *FundCashForecast2) AddSortingCriteriaDetails() *CashSortingCriterion2 {
	newValue := new(CashSortingCriterion2)
	f.SortingCriteriaDetails = append(f.SortingCriteriaDetails, newValue)
	return newValue
}

func (f *FundCashForecast2) AddNetCashForecastDetails() *NetCashForecast1 {
	newValue := new(NetCashForecast1)
	f.NetCashForecastDetails = append(f.NetCashForecastDetails, newValue)
	return newValue
}

// Cash movements from or to a fund as a result of investment funds transactions, eg, subscriptions or redemptions.
type FundCashForecast3 struct {

	// Unique technical identifier for an instance of a fund cash forecast within a fund cash forecast report as assigned by the issuer of the report.
	Identification *Max35Text `xml:"Id"`

	// Date and, if required, the time, at which the price has been applied.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Previous date and time at which a price was applied.
	PreviousTradeDateTime *DateAndDateTimeChoice `xml:"PrvsTradDtTm,omitempty"`

	// Investment fund class to which a cash flow is related.
	FinancialInstrumentDetails *FinancialInstrument9 `xml:"FinInstrmDtls"`

	// Total value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	TotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"TtlNAV,omitempty"`

	// Previous value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	PreviousTotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"PrvsTtlNAV,omitempty"`

	// Total number of investment fund class units that have been issued.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb,omitempty"`

	// Previous total number of investment fund class units that have been issued.
	PreviousTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"PrvsTtlUnitsNb,omitempty"`

	// Rate of change of the net asset value.
	TotalNAVChangeRate *PercentageRate `xml:"TtlNAVChngRate,omitempty"`

	// Currency of the investment fund class.
	InvestmentCurrency []*ActiveOrHistoricCurrencyCode `xml:"InvstmtCcy,omitempty"`

	// Indicates whether the net cash flow is exceptional.
	ExceptionalNetCashFlowIndicator *YesNoIndicator `xml:"XcptnlNetCshFlowInd"`

	// Cash movements into the fund as a result of investment funds transactions, eg, subscriptions or switch-in.
	CashInForecastDetails []*CashInForecast4 `xml:"CshInFcstDtls,omitempty"`

	// Cash movements out of the fund as a result of investment funds transactions, eg, redemptions or switch-out.
	CashOutForecastDetails []*CashOutForecast4 `xml:"CshOutFcstDtls,omitempty"`

	// Cash movements from or to a fund as a result of investment funds transactions.
	NetCashForecastDetails []*NetCashForecast2 `xml:"NetCshFcstDtls,omitempty"`
}

func (f *FundCashForecast3) SetIdentification(value string) {
	f.Identification = (*Max35Text)(&value)
}

func (f *FundCashForecast3) AddTradeDateTime() *DateAndDateTimeChoice {
	f.TradeDateTime = new(DateAndDateTimeChoice)
	return f.TradeDateTime
}

func (f *FundCashForecast3) AddPreviousTradeDateTime() *DateAndDateTimeChoice {
	f.PreviousTradeDateTime = new(DateAndDateTimeChoice)
	return f.PreviousTradeDateTime
}

func (f *FundCashForecast3) AddFinancialInstrumentDetails() *FinancialInstrument9 {
	f.FinancialInstrumentDetails = new(FinancialInstrument9)
	return f.FinancialInstrumentDetails
}

func (f *FundCashForecast3) SetTotalNAV(value, currency string) {
	f.TotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundCashForecast3) SetPreviousTotalNAV(value, currency string) {
	f.PreviousTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundCashForecast3) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.TotalUnitsNumber
}

func (f *FundCashForecast3) AddPreviousTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.PreviousTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.PreviousTotalUnitsNumber
}

func (f *FundCashForecast3) SetTotalNAVChangeRate(value string) {
	f.TotalNAVChangeRate = (*PercentageRate)(&value)
}

func (f *FundCashForecast3) AddInvestmentCurrency(value string) {
	f.InvestmentCurrency = append(f.InvestmentCurrency, (*ActiveOrHistoricCurrencyCode)(&value))
}

func (f *FundCashForecast3) SetExceptionalNetCashFlowIndicator(value string) {
	f.ExceptionalNetCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (f *FundCashForecast3) AddCashInForecastDetails() *CashInForecast4 {
	newValue := new(CashInForecast4)
	f.CashInForecastDetails = append(f.CashInForecastDetails, newValue)
	return newValue
}

func (f *FundCashForecast3) AddCashOutForecastDetails() *CashOutForecast4 {
	newValue := new(CashOutForecast4)
	f.CashOutForecastDetails = append(f.CashOutForecastDetails, newValue)
	return newValue
}

func (f *FundCashForecast3) AddNetCashForecastDetails() *NetCashForecast2 {
	newValue := new(NetCashForecast2)
	f.NetCashForecastDetails = append(f.NetCashForecastDetails, newValue)
	return newValue
}

// Cash movements from or to a fund as a result of investment funds transactions, eg, subscriptions or redemptions.
type FundCashForecast4 struct {

	// Unique technical identifier for an instance of a fund cash forecast within a fund cash forecast report as assigned by the issuer of the report.
	Identification *Max35Text `xml:"Id"`

	// Date and, if required, the time, at which the price has been applied.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Previous date and time at which a price was applied.
	PreviousTradeDateTime *DateAndDateTimeChoice `xml:"PrvsTradDtTm,omitempty"`

	// Investment fund class to which the cash flow is related.
	FinancialInstrumentDetails *FinancialInstrument9 `xml:"FinInstrmDtls"`

	// Total value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	TotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"TtlNAV,omitempty"`

	// Previous value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	PreviousTotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"PrvsTtlNAV,omitempty"`

	// Total number of investment fund class units that have been issued.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb,omitempty"`

	// Previous total number of investment fund class units that have been issued.
	PreviousTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"PrvsTtlUnitsNb,omitempty"`

	// Rate of change of the net asset value.
	TotalNAVChangeRate *PercentageRate `xml:"TtlNAVChngRate,omitempty"`

	// Currency of the investment fund class.
	InvestmentCurrency []*ActiveOrHistoricCurrencyCode `xml:"InvstmtCcy,omitempty"`

	// Net cash movements per financial instrument.
	NetCashForecastDetails []*NetCashForecast2 `xml:"NetCshFcstDtls,omitempty"`

	// Indicates whether the net cash flow is exceptional.
	ExceptionalNetCashFlowIndicator *YesNoIndicator `xml:"XcptnlNetCshFlowInd"`

	// Cash flow by country.
	BreakdownByCountry []*BreakdownByCountry1 `xml:"BrkdwnByCtry,omitempty"`

	// Cash flow by currency.
	BreakdownByCurrency []*BreakdownByCurrency1 `xml:"BrkdwnByCcy,omitempty"`

	// Cash flow by party.
	BreakdownByParty []*BreakdownByParty1 `xml:"BrkdwnByPty,omitempty"`

	// Cash flow by a user defined parameter/s.
	BreakdownByUserDefinedParameter []*BreakdownByUserDefinedParameter1 `xml:"BrkdwnByUsrDfndParam,omitempty"`
}

func (f *FundCashForecast4) SetIdentification(value string) {
	f.Identification = (*Max35Text)(&value)
}

func (f *FundCashForecast4) AddTradeDateTime() *DateAndDateTimeChoice {
	f.TradeDateTime = new(DateAndDateTimeChoice)
	return f.TradeDateTime
}

func (f *FundCashForecast4) AddPreviousTradeDateTime() *DateAndDateTimeChoice {
	f.PreviousTradeDateTime = new(DateAndDateTimeChoice)
	return f.PreviousTradeDateTime
}

func (f *FundCashForecast4) AddFinancialInstrumentDetails() *FinancialInstrument9 {
	f.FinancialInstrumentDetails = new(FinancialInstrument9)
	return f.FinancialInstrumentDetails
}

func (f *FundCashForecast4) SetTotalNAV(value, currency string) {
	f.TotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundCashForecast4) SetPreviousTotalNAV(value, currency string) {
	f.PreviousTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundCashForecast4) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.TotalUnitsNumber
}

func (f *FundCashForecast4) AddPreviousTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.PreviousTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.PreviousTotalUnitsNumber
}

func (f *FundCashForecast4) SetTotalNAVChangeRate(value string) {
	f.TotalNAVChangeRate = (*PercentageRate)(&value)
}

func (f *FundCashForecast4) AddInvestmentCurrency(value string) {
	f.InvestmentCurrency = append(f.InvestmentCurrency, (*ActiveOrHistoricCurrencyCode)(&value))
}

func (f *FundCashForecast4) AddNetCashForecastDetails() *NetCashForecast2 {
	newValue := new(NetCashForecast2)
	f.NetCashForecastDetails = append(f.NetCashForecastDetails, newValue)
	return newValue
}

func (f *FundCashForecast4) SetExceptionalNetCashFlowIndicator(value string) {
	f.ExceptionalNetCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (f *FundCashForecast4) AddBreakdownByCountry() *BreakdownByCountry1 {
	newValue := new(BreakdownByCountry1)
	f.BreakdownByCountry = append(f.BreakdownByCountry, newValue)
	return newValue
}

func (f *FundCashForecast4) AddBreakdownByCurrency() *BreakdownByCurrency1 {
	newValue := new(BreakdownByCurrency1)
	f.BreakdownByCurrency = append(f.BreakdownByCurrency, newValue)
	return newValue
}

func (f *FundCashForecast4) AddBreakdownByParty() *BreakdownByParty1 {
	newValue := new(BreakdownByParty1)
	f.BreakdownByParty = append(f.BreakdownByParty, newValue)
	return newValue
}

func (f *FundCashForecast4) AddBreakdownByUserDefinedParameter() *BreakdownByUserDefinedParameter1 {
	newValue := new(BreakdownByUserDefinedParameter1)
	f.BreakdownByUserDefinedParameter = append(f.BreakdownByUserDefinedParameter, newValue)
	return newValue
}

// Cash movements from or to a fund as a result of investment funds transactions, eg, subscriptions or redemptions.
type FundCashForecast6 struct {

	// Unique technical identifier for an instance of a fund cash forecast within a fund cash forecast report as assigned by the issuer of the report.
	Identification *Max35Text `xml:"Id"`

	// Date and, if required, the time, at which the price has been applied.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Previous date and time at which the price was applied.
	PreviousTradeDateTime *DateAndDateTimeChoice `xml:"PrvsTradDtTm,omitempty"`

	// Investment fund class to which the cash flow is related.
	FinancialInstrumentDetails *FinancialInstrument9 `xml:"FinInstrmDtls"`

	// Total value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	TotalNAV []*ActiveOrHistoricCurrencyAndAmount `xml:"TtlNAV,omitempty"`

	// Previous value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	PreviousTotalNAV []*ActiveOrHistoricCurrencyAndAmount `xml:"PrvsTtlNAV,omitempty"`

	// Total number of investment fund class units that have been issued.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb,omitempty"`

	// Previous total number of investment fund class units that have been issued.
	PreviousTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"PrvsTtlUnitsNb,omitempty"`

	// Rate of change of the net asset value.
	TotalNAVChangeRate *PercentageRate `xml:"TtlNAVChngRate,omitempty"`

	// Currency of the investment fund class.
	InvestmentCurrency []*ActiveOrHistoricCurrencyCode `xml:"InvstmtCcy,omitempty"`

	// Information about the designation of the share class currency, that is, whether it is for onshore or offshore purposes and other information that may be required. This is typically only required for CNY funds.
	CurrencyStatus *CurrencyDesignation1 `xml:"CcySts,omitempty"`

	// Indicates whether the net cash flow is exceptional.
	ExceptionalNetCashFlowIndicator *YesNoIndicator `xml:"XcptnlNetCshFlowInd"`

	// Price per unit of the trade date.
	Price *UnitPrice19 `xml:"Pric,omitempty"`

	// Foreign exchange rate.
	ForeignExchangeRate *ForeignExchangeTerms19 `xml:"FXRate,omitempty"`

	// Net cash flow expressed as a percentage of the total NAV for the share class.
	PercentageOfShareClassTotalNAV *PercentageRate `xml:"PctgOfShrClssTtlNAV,omitempty"`

	// Cash flow by party.
	BreakdownByParty []*BreakdownByParty3 `xml:"BrkdwnByPty,omitempty"`

	// Cash flow by country.
	BreakdownByCountry []*BreakdownByCountry2 `xml:"BrkdwnByCtry,omitempty"`

	// Cash flow by currency.
	BreakdownByCurrency []*BreakdownByCurrency2 `xml:"BrkdwnByCcy,omitempty"`

	// Cash flow by a user defined parameter/s.
	BreakdownByUserDefinedParameter []*BreakdownByUserDefinedParameter3 `xml:"BrkdwnByUsrDfndParam,omitempty"`

	// Net cash movements per financial instrument.
	NetCashForecastDetails []*NetCashForecast4 `xml:"NetCshFcstDtls,omitempty"`
}

func (f *FundCashForecast6) SetIdentification(value string) {
	f.Identification = (*Max35Text)(&value)
}

func (f *FundCashForecast6) AddTradeDateTime() *DateAndDateTimeChoice {
	f.TradeDateTime = new(DateAndDateTimeChoice)
	return f.TradeDateTime
}

func (f *FundCashForecast6) AddPreviousTradeDateTime() *DateAndDateTimeChoice {
	f.PreviousTradeDateTime = new(DateAndDateTimeChoice)
	return f.PreviousTradeDateTime
}

func (f *FundCashForecast6) AddFinancialInstrumentDetails() *FinancialInstrument9 {
	f.FinancialInstrumentDetails = new(FinancialInstrument9)
	return f.FinancialInstrumentDetails
}

func (f *FundCashForecast6) AddTotalNAV(value, currency string) {
	f.TotalNAV = append(f.TotalNAV, NewActiveOrHistoricCurrencyAndAmount(value, currency))
}

func (f *FundCashForecast6) AddPreviousTotalNAV(value, currency string) {
	f.PreviousTotalNAV = append(f.PreviousTotalNAV, NewActiveOrHistoricCurrencyAndAmount(value, currency))
}

func (f *FundCashForecast6) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.TotalUnitsNumber
}

func (f *FundCashForecast6) AddPreviousTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.PreviousTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.PreviousTotalUnitsNumber
}

func (f *FundCashForecast6) SetTotalNAVChangeRate(value string) {
	f.TotalNAVChangeRate = (*PercentageRate)(&value)
}

func (f *FundCashForecast6) AddInvestmentCurrency(value string) {
	f.InvestmentCurrency = append(f.InvestmentCurrency, (*ActiveOrHistoricCurrencyCode)(&value))
}

func (f *FundCashForecast6) AddCurrencyStatus() *CurrencyDesignation1 {
	f.CurrencyStatus = new(CurrencyDesignation1)
	return f.CurrencyStatus
}

func (f *FundCashForecast6) SetExceptionalNetCashFlowIndicator(value string) {
	f.ExceptionalNetCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (f *FundCashForecast6) AddPrice() *UnitPrice19 {
	f.Price = new(UnitPrice19)
	return f.Price
}

func (f *FundCashForecast6) AddForeignExchangeRate() *ForeignExchangeTerms19 {
	f.ForeignExchangeRate = new(ForeignExchangeTerms19)
	return f.ForeignExchangeRate
}

func (f *FundCashForecast6) SetPercentageOfShareClassTotalNAV(value string) {
	f.PercentageOfShareClassTotalNAV = (*PercentageRate)(&value)
}

func (f *FundCashForecast6) AddBreakdownByParty() *BreakdownByParty3 {
	newValue := new(BreakdownByParty3)
	f.BreakdownByParty = append(f.BreakdownByParty, newValue)
	return newValue
}

func (f *FundCashForecast6) AddBreakdownByCountry() *BreakdownByCountry2 {
	newValue := new(BreakdownByCountry2)
	f.BreakdownByCountry = append(f.BreakdownByCountry, newValue)
	return newValue
}

func (f *FundCashForecast6) AddBreakdownByCurrency() *BreakdownByCurrency2 {
	newValue := new(BreakdownByCurrency2)
	f.BreakdownByCurrency = append(f.BreakdownByCurrency, newValue)
	return newValue
}

func (f *FundCashForecast6) AddBreakdownByUserDefinedParameter() *BreakdownByUserDefinedParameter3 {
	newValue := new(BreakdownByUserDefinedParameter3)
	f.BreakdownByUserDefinedParameter = append(f.BreakdownByUserDefinedParameter, newValue)
	return newValue
}

func (f *FundCashForecast6) AddNetCashForecastDetails() *NetCashForecast4 {
	newValue := new(NetCashForecast4)
	f.NetCashForecastDetails = append(f.NetCashForecastDetails, newValue)
	return newValue
}

// Cash movements from or to a fund as a result of investment funds transactions, eg, subscriptions or redemptions.
type FundCashForecast7 struct {

	// Unique technical identifier for an instance of a fund cash forecast within a fund cash forecast report as assigned by the issuer of the report.
	Identification *Max35Text `xml:"Id"`

	// Date and, if required, the time, at which the price has been applied.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Previous date and time at which the price was applied.
	PreviousTradeDateTime *DateAndDateTimeChoice `xml:"PrvsTradDtTm,omitempty"`

	// Investment fund class to which a cash flow is related.
	FinancialInstrumentDetails *FinancialInstrument9 `xml:"FinInstrmDtls"`

	// Total value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	TotalNAV []*ActiveOrHistoricCurrencyAndAmount `xml:"TtlNAV,omitempty"`

	// Previous value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	PreviousTotalNAV []*ActiveOrHistoricCurrencyAndAmount `xml:"PrvsTtlNAV,omitempty"`

	// Total number of investment fund class units that have been issued.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb,omitempty"`

	// Previous total number of investment fund class units that have been issued.
	PreviousTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"PrvsTtlUnitsNb,omitempty"`

	// Rate of change of the net asset value.
	TotalNAVChangeRate *PercentageRate `xml:"TtlNAVChngRate,omitempty"`

	// Currency of the investment fund class.
	InvestmentCurrency []*ActiveOrHistoricCurrencyCode `xml:"InvstmtCcy,omitempty"`

	// Information about the designation of the share class currency, that is, whether it is for onshore or offshore purposes and other information that may be required. This is typically only required for CNY funds.
	CurrencyStatus *CurrencyDesignation1 `xml:"CcySts,omitempty"`

	// Indicates whether the net cash flow is exceptional.
	ExceptionalNetCashFlowIndicator *YesNoIndicator `xml:"XcptnlNetCshFlowInd"`

	// Price per unit of the trade date.
	Price *UnitPrice19 `xml:"Pric,omitempty"`

	// Foreign exchange rate.
	ForeignExchangeRate *ForeignExchangeTerms19 `xml:"FXRate,omitempty"`

	// Net cash flow expressed as a percentage of the total NAV for the share class.
	PercentageOfShareClassTotalNAV *PercentageRate `xml:"PctgOfShrClssTtlNAV,omitempty"`

	// Cash movements into the fund as a result of transactions in shares in an investment fund, for example, subscriptions or switch-ins.
	CashInForecastDetails []*CashInForecast6 `xml:"CshInFcstDtls,omitempty"`

	// Cash movements out of the fund as a result of transactions in shares in an investment fund, for example, redemptions or switch-outs.
	CashOutForecastDetails []*CashOutForecast6 `xml:"CshOutFcstDtls,omitempty"`

	// Net cash as a result of the cash-in and cash-out flows.
	NetCashForecastDetails []*NetCashForecast4 `xml:"NetCshFcstDtls,omitempty"`
}

func (f *FundCashForecast7) SetIdentification(value string) {
	f.Identification = (*Max35Text)(&value)
}

func (f *FundCashForecast7) AddTradeDateTime() *DateAndDateTimeChoice {
	f.TradeDateTime = new(DateAndDateTimeChoice)
	return f.TradeDateTime
}

func (f *FundCashForecast7) AddPreviousTradeDateTime() *DateAndDateTimeChoice {
	f.PreviousTradeDateTime = new(DateAndDateTimeChoice)
	return f.PreviousTradeDateTime
}

func (f *FundCashForecast7) AddFinancialInstrumentDetails() *FinancialInstrument9 {
	f.FinancialInstrumentDetails = new(FinancialInstrument9)
	return f.FinancialInstrumentDetails
}

func (f *FundCashForecast7) AddTotalNAV(value, currency string) {
	f.TotalNAV = append(f.TotalNAV, NewActiveOrHistoricCurrencyAndAmount(value, currency))
}

func (f *FundCashForecast7) AddPreviousTotalNAV(value, currency string) {
	f.PreviousTotalNAV = append(f.PreviousTotalNAV, NewActiveOrHistoricCurrencyAndAmount(value, currency))
}

func (f *FundCashForecast7) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.TotalUnitsNumber
}

func (f *FundCashForecast7) AddPreviousTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	f.PreviousTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return f.PreviousTotalUnitsNumber
}

func (f *FundCashForecast7) SetTotalNAVChangeRate(value string) {
	f.TotalNAVChangeRate = (*PercentageRate)(&value)
}

func (f *FundCashForecast7) AddInvestmentCurrency(value string) {
	f.InvestmentCurrency = append(f.InvestmentCurrency, (*ActiveOrHistoricCurrencyCode)(&value))
}

func (f *FundCashForecast7) AddCurrencyStatus() *CurrencyDesignation1 {
	f.CurrencyStatus = new(CurrencyDesignation1)
	return f.CurrencyStatus
}

func (f *FundCashForecast7) SetExceptionalNetCashFlowIndicator(value string) {
	f.ExceptionalNetCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (f *FundCashForecast7) AddPrice() *UnitPrice19 {
	f.Price = new(UnitPrice19)
	return f.Price
}

func (f *FundCashForecast7) AddForeignExchangeRate() *ForeignExchangeTerms19 {
	f.ForeignExchangeRate = new(ForeignExchangeTerms19)
	return f.ForeignExchangeRate
}

func (f *FundCashForecast7) SetPercentageOfShareClassTotalNAV(value string) {
	f.PercentageOfShareClassTotalNAV = (*PercentageRate)(&value)
}

func (f *FundCashForecast7) AddCashInForecastDetails() *CashInForecast6 {
	newValue := new(CashInForecast6)
	f.CashInForecastDetails = append(f.CashInForecastDetails, newValue)
	return newValue
}

func (f *FundCashForecast7) AddCashOutForecastDetails() *CashOutForecast6 {
	newValue := new(CashOutForecast6)
	f.CashOutForecastDetails = append(f.CashOutForecastDetails, newValue)
	return newValue
}

func (f *FundCashForecast7) AddNetCashForecastDetails() *NetCashForecast4 {
	newValue := new(NetCashForecast4)
	f.NetCashForecastDetails = append(f.NetCashForecastDetails, newValue)
	return newValue
}
