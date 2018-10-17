package iso20022

// Cash movements from or to a fund as a result of investment funds transactions, eg, subscriptions or redemptions.
type EstimatedFundCashForecast1 struct {

	// Date and, if required, the time, at which the price has been applied.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Previous date and time at which a price was applied.
	PreviousTradeDateTime *DateAndDateTimeChoice `xml:"PrvsTradDtTm"`

	// Investment fund class to which a cash flow is related.
	FinancialInstrumentDetails *FinancialInstrument5 `xml:"FinInstrmDtls"`

	// Estimated total value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	EstimatedTotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"EstmtdTtlNAV,omitempty"`

	// Previous estimated value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	PreviousEstimatedTotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"PrvsEstmtdTtlNAV,omitempty"`

	// Estimated total number of investment fund class units that have been issued.
	EstimatedTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"EstmtdTtlUnitsNb,omitempty"`

	// Previous estimated value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	PreviousEstimatedTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"PrvsEstmtdTtlUnitsNb,omitempty"`

	// Rate of change of the net asset value.
	EstimatedTotalNAVChangeRate *PercentageRate `xml:"EstmtdTtlNAVChngRate,omitempty"`

	// Currency of the investment fund class.
	InvestmentCurrency []*ActiveOrHistoricCurrencyCode `xml:"InvstmtCcy,omitempty"`

	// Indicates whether the estimated net cash flow is exceptional.
	ExceptionalNetCashFlowIndicator *YesNoIndicator `xml:"XcptnlNetCshFlowInd"`

	// Cash movements into a fund as a result of investment funds transactions, eg, subscriptions or switch-in.
	EstimatedCashInForecastDetails []*CashInForecast2 `xml:"EstmtdCshInFcstDtls,omitempty"`

	// Cash movements out of a fund as a result of investment funds transactions, eg, redemptions or switch-out.
	EstimatedCashOutForecastDetails []*CashOutForecast2 `xml:"EstmtdCshOutFcstDtls,omitempty"`

	// Net cash movements to a fund as a result of investment funds transactions.
	EstimatedNetCashForecastDetails []*NetCashForecast1 `xml:"EstmtdNetCshFcstDtls,omitempty"`
}

func (e *EstimatedFundCashForecast1) AddTradeDateTime() *DateAndDateTimeChoice {
	e.TradeDateTime = new(DateAndDateTimeChoice)
	return e.TradeDateTime
}

func (e *EstimatedFundCashForecast1) AddPreviousTradeDateTime() *DateAndDateTimeChoice {
	e.PreviousTradeDateTime = new(DateAndDateTimeChoice)
	return e.PreviousTradeDateTime
}

func (e *EstimatedFundCashForecast1) AddFinancialInstrumentDetails() *FinancialInstrument5 {
	e.FinancialInstrumentDetails = new(FinancialInstrument5)
	return e.FinancialInstrumentDetails
}

func (e *EstimatedFundCashForecast1) SetEstimatedTotalNAV(value, currency string) {
	e.EstimatedTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (e *EstimatedFundCashForecast1) SetPreviousEstimatedTotalNAV(value, currency string) {
	e.PreviousEstimatedTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (e *EstimatedFundCashForecast1) AddEstimatedTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	e.EstimatedTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return e.EstimatedTotalUnitsNumber
}

func (e *EstimatedFundCashForecast1) AddPreviousEstimatedTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	e.PreviousEstimatedTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return e.PreviousEstimatedTotalUnitsNumber
}

func (e *EstimatedFundCashForecast1) SetEstimatedTotalNAVChangeRate(value string) {
	e.EstimatedTotalNAVChangeRate = (*PercentageRate)(&value)
}

func (e *EstimatedFundCashForecast1) AddInvestmentCurrency(value string) {
	e.InvestmentCurrency = append(e.InvestmentCurrency, (*ActiveOrHistoricCurrencyCode)(&value))
}

func (e *EstimatedFundCashForecast1) SetExceptionalNetCashFlowIndicator(value string) {
	e.ExceptionalNetCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (e *EstimatedFundCashForecast1) AddEstimatedCashInForecastDetails() *CashInForecast2 {
	newValue := new(CashInForecast2)
	e.EstimatedCashInForecastDetails = append(e.EstimatedCashInForecastDetails, newValue)
	return newValue
}

func (e *EstimatedFundCashForecast1) AddEstimatedCashOutForecastDetails() *CashOutForecast2 {
	newValue := new(CashOutForecast2)
	e.EstimatedCashOutForecastDetails = append(e.EstimatedCashOutForecastDetails, newValue)
	return newValue
}

func (e *EstimatedFundCashForecast1) AddEstimatedNetCashForecastDetails() *NetCashForecast1 {
	newValue := new(NetCashForecast1)
	e.EstimatedNetCashForecastDetails = append(e.EstimatedNetCashForecastDetails, newValue)
	return newValue
}

// Cash movements from or to a fund as a result of investment funds transactions, eg, subscriptions or redemptions.
type EstimatedFundCashForecast2 struct {

	// Date and, if required, the time, at which the price has been applied.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Previous date and time at which a price was applied.
	PreviousTradeDateTime *DateAndDateTimeChoice `xml:"PrvsTradDtTm"`

	// Investment fund class to which a cash flow is related.
	FinancialInstrumentDetails *FinancialInstrument5 `xml:"FinInstrmDtls"`

	// Estimated total value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	EstimatedTotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"EstmtdTtlNAV,omitempty"`

	// Previous estimated value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	PreviousEstimatedTotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"PrvsEstmtdTtlNAV,omitempty"`

	// Estimated total number of investment fund class units that have been issued.
	EstimatedTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"EstmtdTtlUnitsNb,omitempty"`

	// Previous estimated value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	PreviousEstimatedTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"PrvsEstmtdTtlUnitsNb,omitempty"`

	// Rate of change of the net asset value.
	EstimatedTotalNAVChangeRate *PercentageRate `xml:"EstmtdTtlNAVChngRate,omitempty"`

	// Currency of the investment fund class.
	InvestmentCurrency []*ActiveOrHistoricCurrencyCode `xml:"InvstmtCcy,omitempty"`

	// Indicates whether the estimated net cash flow is exceptional.
	ExceptionalNetCashFlowIndicator *YesNoIndicator `xml:"XcptnlNetCshFlowInd"`

	// Information related to the estimated cash movements reported by pre-defined or user defined criteria.
	SortingCriteriaDetails []*CashSortingCriterion1 `xml:"SrtgCritDtls"`

	// Net cash movements per financial instrument.
	EstimatedNetCashForecastDetails []*NetCashForecast1 `xml:"EstmtdNetCshFcstDtls,omitempty"`
}

func (e *EstimatedFundCashForecast2) AddTradeDateTime() *DateAndDateTimeChoice {
	e.TradeDateTime = new(DateAndDateTimeChoice)
	return e.TradeDateTime
}

func (e *EstimatedFundCashForecast2) AddPreviousTradeDateTime() *DateAndDateTimeChoice {
	e.PreviousTradeDateTime = new(DateAndDateTimeChoice)
	return e.PreviousTradeDateTime
}

func (e *EstimatedFundCashForecast2) AddFinancialInstrumentDetails() *FinancialInstrument5 {
	e.FinancialInstrumentDetails = new(FinancialInstrument5)
	return e.FinancialInstrumentDetails
}

func (e *EstimatedFundCashForecast2) SetEstimatedTotalNAV(value, currency string) {
	e.EstimatedTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (e *EstimatedFundCashForecast2) SetPreviousEstimatedTotalNAV(value, currency string) {
	e.PreviousEstimatedTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (e *EstimatedFundCashForecast2) AddEstimatedTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	e.EstimatedTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return e.EstimatedTotalUnitsNumber
}

func (e *EstimatedFundCashForecast2) AddPreviousEstimatedTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	e.PreviousEstimatedTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return e.PreviousEstimatedTotalUnitsNumber
}

func (e *EstimatedFundCashForecast2) SetEstimatedTotalNAVChangeRate(value string) {
	e.EstimatedTotalNAVChangeRate = (*PercentageRate)(&value)
}

func (e *EstimatedFundCashForecast2) AddInvestmentCurrency(value string) {
	e.InvestmentCurrency = append(e.InvestmentCurrency, (*ActiveOrHistoricCurrencyCode)(&value))
}

func (e *EstimatedFundCashForecast2) SetExceptionalNetCashFlowIndicator(value string) {
	e.ExceptionalNetCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (e *EstimatedFundCashForecast2) AddSortingCriteriaDetails() *CashSortingCriterion1 {
	newValue := new(CashSortingCriterion1)
	e.SortingCriteriaDetails = append(e.SortingCriteriaDetails, newValue)
	return newValue
}

func (e *EstimatedFundCashForecast2) AddEstimatedNetCashForecastDetails() *NetCashForecast1 {
	newValue := new(NetCashForecast1)
	e.EstimatedNetCashForecastDetails = append(e.EstimatedNetCashForecastDetails, newValue)
	return newValue
}

// Cash movements from or to a fund as a result of investment funds transactions, eg, subscriptions or redemptions.
type EstimatedFundCashForecast3 struct {

	// Unique technical identifier for an instance of a fund cash forecast within a fund cash forecast report as assigned by the issuer of the report.
	Identification *Max35Text `xml:"Id"`

	// Date and, if required, the time, at which the price has been applied.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Previous date and time at which a price was applied.
	PreviousTradeDateTime *DateAndDateTimeChoice `xml:"PrvsTradDtTm"`

	// Investment fund class to which the cash flow is related.
	FinancialInstrumentDetails *FinancialInstrument9 `xml:"FinInstrmDtls"`

	// Estimated total value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	EstimatedTotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"EstmtdTtlNAV,omitempty"`

	// Previous value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	PreviousTotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"PrvsTtlNAV,omitempty"`

	// Estimated total number of investment fund class units that have been issued.
	EstimatedTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"EstmtdTtlUnitsNb,omitempty"`

	// Previous value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	PreviousTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"PrvsTtlUnitsNb,omitempty"`

	// Rate of change of the net asset value.
	EstimatedTotalNAVChangeRate *PercentageRate `xml:"EstmtdTtlNAVChngRate,omitempty"`

	// Currency of the investment fund class.
	InvestmentCurrency []*ActiveOrHistoricCurrencyCode `xml:"InvstmtCcy,omitempty"`

	// Indicates whether the estimated net cash flow is exceptional.
	ExceptionalNetCashFlowIndicator *YesNoIndicator `xml:"XcptnlNetCshFlowInd"`

	// Estimated cash movements into the fund as a result of investment funds transactions, eg, subscriptions or switch-in.
	EstimatedCashInForecastDetails []*CashInForecast4 `xml:"EstmtdCshInFcstDtls,omitempty"`

	// Estimated cash movements out of the fund as a result of investment funds transactions, eg, redemptions or switch-out.
	EstimatedCashOutForecastDetails []*CashOutForecast4 `xml:"EstmtdCshOutFcstDtls,omitempty"`

	// Net cash movements to a fund as a result of investment funds transactions.
	EstimatedNetCashForecastDetails []*NetCashForecast2 `xml:"EstmtdNetCshFcstDtls,omitempty"`
}

func (e *EstimatedFundCashForecast3) SetIdentification(value string) {
	e.Identification = (*Max35Text)(&value)
}

func (e *EstimatedFundCashForecast3) AddTradeDateTime() *DateAndDateTimeChoice {
	e.TradeDateTime = new(DateAndDateTimeChoice)
	return e.TradeDateTime
}

func (e *EstimatedFundCashForecast3) AddPreviousTradeDateTime() *DateAndDateTimeChoice {
	e.PreviousTradeDateTime = new(DateAndDateTimeChoice)
	return e.PreviousTradeDateTime
}

func (e *EstimatedFundCashForecast3) AddFinancialInstrumentDetails() *FinancialInstrument9 {
	e.FinancialInstrumentDetails = new(FinancialInstrument9)
	return e.FinancialInstrumentDetails
}

func (e *EstimatedFundCashForecast3) SetEstimatedTotalNAV(value, currency string) {
	e.EstimatedTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (e *EstimatedFundCashForecast3) SetPreviousTotalNAV(value, currency string) {
	e.PreviousTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (e *EstimatedFundCashForecast3) AddEstimatedTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	e.EstimatedTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return e.EstimatedTotalUnitsNumber
}

func (e *EstimatedFundCashForecast3) AddPreviousTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	e.PreviousTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return e.PreviousTotalUnitsNumber
}

func (e *EstimatedFundCashForecast3) SetEstimatedTotalNAVChangeRate(value string) {
	e.EstimatedTotalNAVChangeRate = (*PercentageRate)(&value)
}

func (e *EstimatedFundCashForecast3) AddInvestmentCurrency(value string) {
	e.InvestmentCurrency = append(e.InvestmentCurrency, (*ActiveOrHistoricCurrencyCode)(&value))
}

func (e *EstimatedFundCashForecast3) SetExceptionalNetCashFlowIndicator(value string) {
	e.ExceptionalNetCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (e *EstimatedFundCashForecast3) AddEstimatedCashInForecastDetails() *CashInForecast4 {
	newValue := new(CashInForecast4)
	e.EstimatedCashInForecastDetails = append(e.EstimatedCashInForecastDetails, newValue)
	return newValue
}

func (e *EstimatedFundCashForecast3) AddEstimatedCashOutForecastDetails() *CashOutForecast4 {
	newValue := new(CashOutForecast4)
	e.EstimatedCashOutForecastDetails = append(e.EstimatedCashOutForecastDetails, newValue)
	return newValue
}

func (e *EstimatedFundCashForecast3) AddEstimatedNetCashForecastDetails() *NetCashForecast2 {
	newValue := new(NetCashForecast2)
	e.EstimatedNetCashForecastDetails = append(e.EstimatedNetCashForecastDetails, newValue)
	return newValue
}

// Cash movements from or to a fund as a result of investment funds transactions, eg, subscriptions or redemptions.
type EstimatedFundCashForecast4 struct {

	// Unique technical identifier for an instance of a fund cash forecast within a fund cash forecast report as assigned by the issuer of the report.
	Identification *Max35Text `xml:"Id"`

	// Date and, if required, the time, at which the price has been applied.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Previous date and time at which a price was applied.
	PreviousTradeDateTime *DateAndDateTimeChoice `xml:"PrvsTradDtTm"`

	// Investment fund class to which the cash flow is related.
	FinancialInstrumentDetails *FinancialInstrument9 `xml:"FinInstrmDtls"`

	// Estimated total value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	EstimatedTotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"EstmtdTtlNAV,omitempty"`

	// Previous value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	PreviousTotalNAV *ActiveOrHistoricCurrencyAndAmount `xml:"PrvsTtlNAV,omitempty"`

	// Estimated total number of investment fund class units that have been issued.
	EstimatedTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"EstmtdTtlUnitsNb,omitempty"`

	// Previous value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	PreviousTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"PrvsTtlUnitsNb,omitempty"`

	// Estimated cash flow by party.
	BreakdownByParty []*BreakdownByParty1 `xml:"BrkdwnByPty,omitempty"`

	// Estimated cash flow by country.
	BreakdownByCountry []*BreakdownByCountry1 `xml:"BrkdwnByCtry,omitempty"`

	// Estimated cash flow by currency.
	BreakdownByCurrency []*BreakdownByCurrency1 `xml:"BrkdwnByCcy,omitempty"`

	// Estimated cash flow by a user defined parameter/s.
	BreakdownByUserDefinedParameter []*BreakdownByUserDefinedParameter1 `xml:"BrkdwnByUsrDfndParam,omitempty"`

	// Rate of change of the net asset value.
	EstimatedTotalNAVChangeRate *PercentageRate `xml:"EstmtdTtlNAVChngRate,omitempty"`

	// Currency of the investment fund class.
	InvestmentCurrency []*ActiveOrHistoricCurrencyCode `xml:"InvstmtCcy,omitempty"`

	// Indicates whether the estimated net cash flow is exceptional.
	ExceptionalNetCashFlowIndicator *YesNoIndicator `xml:"XcptnlNetCshFlowInd"`

	// Estimated net cash movements per financial instrument.
	EstimatedNetCashForecastDetails []*NetCashForecast2 `xml:"EstmtdNetCshFcstDtls,omitempty"`
}

func (e *EstimatedFundCashForecast4) SetIdentification(value string) {
	e.Identification = (*Max35Text)(&value)
}

func (e *EstimatedFundCashForecast4) AddTradeDateTime() *DateAndDateTimeChoice {
	e.TradeDateTime = new(DateAndDateTimeChoice)
	return e.TradeDateTime
}

func (e *EstimatedFundCashForecast4) AddPreviousTradeDateTime() *DateAndDateTimeChoice {
	e.PreviousTradeDateTime = new(DateAndDateTimeChoice)
	return e.PreviousTradeDateTime
}

func (e *EstimatedFundCashForecast4) AddFinancialInstrumentDetails() *FinancialInstrument9 {
	e.FinancialInstrumentDetails = new(FinancialInstrument9)
	return e.FinancialInstrumentDetails
}

func (e *EstimatedFundCashForecast4) SetEstimatedTotalNAV(value, currency string) {
	e.EstimatedTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (e *EstimatedFundCashForecast4) SetPreviousTotalNAV(value, currency string) {
	e.PreviousTotalNAV = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (e *EstimatedFundCashForecast4) AddEstimatedTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	e.EstimatedTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return e.EstimatedTotalUnitsNumber
}

func (e *EstimatedFundCashForecast4) AddPreviousTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	e.PreviousTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return e.PreviousTotalUnitsNumber
}

func (e *EstimatedFundCashForecast4) AddBreakdownByParty() *BreakdownByParty1 {
	newValue := new(BreakdownByParty1)
	e.BreakdownByParty = append(e.BreakdownByParty, newValue)
	return newValue
}

func (e *EstimatedFundCashForecast4) AddBreakdownByCountry() *BreakdownByCountry1 {
	newValue := new(BreakdownByCountry1)
	e.BreakdownByCountry = append(e.BreakdownByCountry, newValue)
	return newValue
}

func (e *EstimatedFundCashForecast4) AddBreakdownByCurrency() *BreakdownByCurrency1 {
	newValue := new(BreakdownByCurrency1)
	e.BreakdownByCurrency = append(e.BreakdownByCurrency, newValue)
	return newValue
}

func (e *EstimatedFundCashForecast4) AddBreakdownByUserDefinedParameter() *BreakdownByUserDefinedParameter1 {
	newValue := new(BreakdownByUserDefinedParameter1)
	e.BreakdownByUserDefinedParameter = append(e.BreakdownByUserDefinedParameter, newValue)
	return newValue
}

func (e *EstimatedFundCashForecast4) SetEstimatedTotalNAVChangeRate(value string) {
	e.EstimatedTotalNAVChangeRate = (*PercentageRate)(&value)
}

func (e *EstimatedFundCashForecast4) AddInvestmentCurrency(value string) {
	e.InvestmentCurrency = append(e.InvestmentCurrency, (*ActiveOrHistoricCurrencyCode)(&value))
}

func (e *EstimatedFundCashForecast4) SetExceptionalNetCashFlowIndicator(value string) {
	e.ExceptionalNetCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (e *EstimatedFundCashForecast4) AddEstimatedNetCashForecastDetails() *NetCashForecast2 {
	newValue := new(NetCashForecast2)
	e.EstimatedNetCashForecastDetails = append(e.EstimatedNetCashForecastDetails, newValue)
	return newValue
}

// Cash movements from or to a fund as a result of investment funds transactions, eg, subscriptions or redemptions.
type EstimatedFundCashForecast5 struct {

	// Unique technical identifier for an instance of a fund cash forecast within a fund cash forecast report as assigned by the issuer of the report.
	Identification *Max35Text `xml:"Id"`

	// Date and, if required, the time, at which the price will be applied.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Previous date and time at which the price was applied.
	PreviousTradeDateTime *DateAndDateTimeChoice `xml:"PrvsTradDtTm,omitempty"`

	// Investment fund class to which the cash flow is related.
	FinancialInstrumentDetails *FinancialInstrument9 `xml:"FinInstrmDtls"`

	// Estimated total value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	EstimatedTotalNAV []*ActiveOrHistoricCurrencyAndAmount `xml:"EstmtdTtlNAV,omitempty"`

	// Previous value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	PreviousTotalNAV []*ActiveOrHistoricCurrencyAndAmount `xml:"PrvsTtlNAV,omitempty"`

	// Estimated total number of investment fund class units that have been issued.
	EstimatedTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"EstmtdTtlUnitsNb,omitempty"`

	// Previous value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	PreviousTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"PrvsTtlUnitsNb,omitempty"`

	// Rate of change of the net asset value.
	EstimatedTotalNAVChangeRate *PercentageRate `xml:"EstmtdTtlNAVChngRate,omitempty"`

	// Currency of the investment fund class.
	InvestmentCurrency []*ActiveOrHistoricCurrencyCode `xml:"InvstmtCcy,omitempty"`

	// Information about the designation of the share class currency, that is, whether it is for onshore or offshore purposes and other information that may be required. This is typically only required for CNY funds.
	CurrencyStatus *CurrencyDesignation1 `xml:"CcySts,omitempty"`

	// Indicates whether the estimated net cash flow is exceptional.
	ExceptionalNetCashFlowIndicator *YesNoIndicator `xml:"XcptnlNetCshFlowInd"`

	// Price per unit of the previous trade date.
	Price *UnitPrice19 `xml:"Pric,omitempty"`

	// Foreign exchange rate.
	ForeignExchangeRate *ForeignExchangeTerms19 `xml:"FXRate,omitempty"`

	// Estimated net cash flow expressed as a percentage of the previous total NAV for the share class.
	EstimatedPercentageOfShareClassTotalNAV *PercentageRate `xml:"EstmtdPctgOfShrClssTtlNAV,omitempty"`

	// Estimated cash flow by party.
	BreakdownByParty []*BreakdownByParty3 `xml:"BrkdwnByPty,omitempty"`

	// Estimated cash flow by country.
	BreakdownByCountry []*BreakdownByCountry2 `xml:"BrkdwnByCtry,omitempty"`

	// Estimated cash flow by currency.
	BreakdownByCurrency []*BreakdownByCurrency2 `xml:"BrkdwnByCcy,omitempty"`

	// Estimated cash flow by a user defined parameter/s.
	BreakdownByUserDefinedParameter []*BreakdownByUserDefinedParameter3 `xml:"BrkdwnByUsrDfndParam,omitempty"`

	// Estimated net cash movements per financial instrument.
	EstimatedNetCashForecastDetails []*NetCashForecast4 `xml:"EstmtdNetCshFcstDtls,omitempty"`
}

func (e *EstimatedFundCashForecast5) SetIdentification(value string) {
	e.Identification = (*Max35Text)(&value)
}

func (e *EstimatedFundCashForecast5) AddTradeDateTime() *DateAndDateTimeChoice {
	e.TradeDateTime = new(DateAndDateTimeChoice)
	return e.TradeDateTime
}

func (e *EstimatedFundCashForecast5) AddPreviousTradeDateTime() *DateAndDateTimeChoice {
	e.PreviousTradeDateTime = new(DateAndDateTimeChoice)
	return e.PreviousTradeDateTime
}

func (e *EstimatedFundCashForecast5) AddFinancialInstrumentDetails() *FinancialInstrument9 {
	e.FinancialInstrumentDetails = new(FinancialInstrument9)
	return e.FinancialInstrumentDetails
}

func (e *EstimatedFundCashForecast5) AddEstimatedTotalNAV(value, currency string) {
	e.EstimatedTotalNAV = append(e.EstimatedTotalNAV, NewActiveOrHistoricCurrencyAndAmount(value, currency))
}

func (e *EstimatedFundCashForecast5) AddPreviousTotalNAV(value, currency string) {
	e.PreviousTotalNAV = append(e.PreviousTotalNAV, NewActiveOrHistoricCurrencyAndAmount(value, currency))
}

func (e *EstimatedFundCashForecast5) AddEstimatedTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	e.EstimatedTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return e.EstimatedTotalUnitsNumber
}

func (e *EstimatedFundCashForecast5) AddPreviousTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	e.PreviousTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return e.PreviousTotalUnitsNumber
}

func (e *EstimatedFundCashForecast5) SetEstimatedTotalNAVChangeRate(value string) {
	e.EstimatedTotalNAVChangeRate = (*PercentageRate)(&value)
}

func (e *EstimatedFundCashForecast5) AddInvestmentCurrency(value string) {
	e.InvestmentCurrency = append(e.InvestmentCurrency, (*ActiveOrHistoricCurrencyCode)(&value))
}

func (e *EstimatedFundCashForecast5) AddCurrencyStatus() *CurrencyDesignation1 {
	e.CurrencyStatus = new(CurrencyDesignation1)
	return e.CurrencyStatus
}

func (e *EstimatedFundCashForecast5) SetExceptionalNetCashFlowIndicator(value string) {
	e.ExceptionalNetCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (e *EstimatedFundCashForecast5) AddPrice() *UnitPrice19 {
	e.Price = new(UnitPrice19)
	return e.Price
}

func (e *EstimatedFundCashForecast5) AddForeignExchangeRate() *ForeignExchangeTerms19 {
	e.ForeignExchangeRate = new(ForeignExchangeTerms19)
	return e.ForeignExchangeRate
}

func (e *EstimatedFundCashForecast5) SetEstimatedPercentageOfShareClassTotalNAV(value string) {
	e.EstimatedPercentageOfShareClassTotalNAV = (*PercentageRate)(&value)
}

func (e *EstimatedFundCashForecast5) AddBreakdownByParty() *BreakdownByParty3 {
	newValue := new(BreakdownByParty3)
	e.BreakdownByParty = append(e.BreakdownByParty, newValue)
	return newValue
}

func (e *EstimatedFundCashForecast5) AddBreakdownByCountry() *BreakdownByCountry2 {
	newValue := new(BreakdownByCountry2)
	e.BreakdownByCountry = append(e.BreakdownByCountry, newValue)
	return newValue
}

func (e *EstimatedFundCashForecast5) AddBreakdownByCurrency() *BreakdownByCurrency2 {
	newValue := new(BreakdownByCurrency2)
	e.BreakdownByCurrency = append(e.BreakdownByCurrency, newValue)
	return newValue
}

func (e *EstimatedFundCashForecast5) AddBreakdownByUserDefinedParameter() *BreakdownByUserDefinedParameter3 {
	newValue := new(BreakdownByUserDefinedParameter3)
	e.BreakdownByUserDefinedParameter = append(e.BreakdownByUserDefinedParameter, newValue)
	return newValue
}

func (e *EstimatedFundCashForecast5) AddEstimatedNetCashForecastDetails() *NetCashForecast4 {
	newValue := new(NetCashForecast4)
	e.EstimatedNetCashForecastDetails = append(e.EstimatedNetCashForecastDetails, newValue)
	return newValue
}

// Cash movements from or to a fund as a result of investment funds transactions, eg, subscriptions or redemptions.
type EstimatedFundCashForecast6 struct {

	// Unique technical identifier for an instance of a fund cash forecast within a fund cash forecast report as assigned by the issuer of the report.
	Identification *Max35Text `xml:"Id"`

	// Date and, if required, the time, at which the price will be applied.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Previous date and time at which the price was applied.
	PreviousTradeDateTime *DateAndDateTimeChoice `xml:"PrvsTradDtTm,omitempty"`

	// Investment fund class to which the cash flow is related.
	FinancialInstrumentDetails *FinancialInstrument9 `xml:"FinInstrmDtls"`

	// Estimated total value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	EstimatedTotalNAV []*ActiveOrHistoricCurrencyAndAmount `xml:"EstmtdTtlNAV,omitempty"`

	// Previous value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	PreviousTotalNAV []*ActiveOrHistoricCurrencyAndAmount `xml:"PrvsTtlNAV,omitempty"`

	// Estimated total number of investment fund class units that have been issued.
	EstimatedTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"EstmtdTtlUnitsNb,omitempty"`

	// Previous value of all the holdings, less the fund's liabilities, attributable to a specific investment fund class.
	PreviousTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"PrvsTtlUnitsNb,omitempty"`

	// Rate of change of the net asset value.
	EstimatedTotalNAVChangeRate *PercentageRate `xml:"EstmtdTtlNAVChngRate,omitempty"`

	// Currency of the investment fund class.
	InvestmentCurrency []*ActiveOrHistoricCurrencyCode `xml:"InvstmtCcy,omitempty"`

	// Information about the designation of the share class currency, that is, whether it is for onshore or offshore purposes and other information that may be required. This is typically only required for CNY funds.
	CurrencyStatus *CurrencyDesignation1 `xml:"CcySts,omitempty"`

	// Indicates whether the estimated net cash flow is exceptional.
	ExceptionalNetCashFlowIndicator *YesNoIndicator `xml:"XcptnlNetCshFlowInd"`

	// Price per unit of the previous trade date.
	Price *UnitPrice19 `xml:"Pric,omitempty"`

	// Foreign exchange rate.
	ForeignExchangeRate *ForeignExchangeTerms19 `xml:"FXRate,omitempty"`

	// Estimated net cash flow expressed as a percentage of the previous total NAV for the share class.
	EstimatedPercentageOfShareClassTotalNAV *PercentageRate `xml:"EstmtdPctgOfShrClssTtlNAV,omitempty"`

	// Estimated cash movements into the fund as a result of transactions in shares in an investment fund, for example, subscriptions or switch-ins.
	EstimatedCashInForecastDetails []*CashInForecast6 `xml:"EstmtdCshInFcstDtls,omitempty"`

	// Estimated cash movements out of the fund as a result of transactions in shares in an investment fund, for example, redemptions or switch-outs.
	EstimatedCashOutForecastDetails []*CashOutForecast6 `xml:"EstmtdCshOutFcstDtls,omitempty"`

	// Net cash as a result of the cash-in and cash-out flows.
	EstimatedNetCashForecastDetails []*NetCashForecast4 `xml:"EstmtdNetCshFcstDtls,omitempty"`
}

func (e *EstimatedFundCashForecast6) SetIdentification(value string) {
	e.Identification = (*Max35Text)(&value)
}

func (e *EstimatedFundCashForecast6) AddTradeDateTime() *DateAndDateTimeChoice {
	e.TradeDateTime = new(DateAndDateTimeChoice)
	return e.TradeDateTime
}

func (e *EstimatedFundCashForecast6) AddPreviousTradeDateTime() *DateAndDateTimeChoice {
	e.PreviousTradeDateTime = new(DateAndDateTimeChoice)
	return e.PreviousTradeDateTime
}

func (e *EstimatedFundCashForecast6) AddFinancialInstrumentDetails() *FinancialInstrument9 {
	e.FinancialInstrumentDetails = new(FinancialInstrument9)
	return e.FinancialInstrumentDetails
}

func (e *EstimatedFundCashForecast6) AddEstimatedTotalNAV(value, currency string) {
	e.EstimatedTotalNAV = append(e.EstimatedTotalNAV, NewActiveOrHistoricCurrencyAndAmount(value, currency))
}

func (e *EstimatedFundCashForecast6) AddPreviousTotalNAV(value, currency string) {
	e.PreviousTotalNAV = append(e.PreviousTotalNAV, NewActiveOrHistoricCurrencyAndAmount(value, currency))
}

func (e *EstimatedFundCashForecast6) AddEstimatedTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	e.EstimatedTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return e.EstimatedTotalUnitsNumber
}

func (e *EstimatedFundCashForecast6) AddPreviousTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	e.PreviousTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return e.PreviousTotalUnitsNumber
}

func (e *EstimatedFundCashForecast6) SetEstimatedTotalNAVChangeRate(value string) {
	e.EstimatedTotalNAVChangeRate = (*PercentageRate)(&value)
}

func (e *EstimatedFundCashForecast6) AddInvestmentCurrency(value string) {
	e.InvestmentCurrency = append(e.InvestmentCurrency, (*ActiveOrHistoricCurrencyCode)(&value))
}

func (e *EstimatedFundCashForecast6) AddCurrencyStatus() *CurrencyDesignation1 {
	e.CurrencyStatus = new(CurrencyDesignation1)
	return e.CurrencyStatus
}

func (e *EstimatedFundCashForecast6) SetExceptionalNetCashFlowIndicator(value string) {
	e.ExceptionalNetCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (e *EstimatedFundCashForecast6) AddPrice() *UnitPrice19 {
	e.Price = new(UnitPrice19)
	return e.Price
}

func (e *EstimatedFundCashForecast6) AddForeignExchangeRate() *ForeignExchangeTerms19 {
	e.ForeignExchangeRate = new(ForeignExchangeTerms19)
	return e.ForeignExchangeRate
}

func (e *EstimatedFundCashForecast6) SetEstimatedPercentageOfShareClassTotalNAV(value string) {
	e.EstimatedPercentageOfShareClassTotalNAV = (*PercentageRate)(&value)
}

func (e *EstimatedFundCashForecast6) AddEstimatedCashInForecastDetails() *CashInForecast6 {
	newValue := new(CashInForecast6)
	e.EstimatedCashInForecastDetails = append(e.EstimatedCashInForecastDetails, newValue)
	return newValue
}

func (e *EstimatedFundCashForecast6) AddEstimatedCashOutForecastDetails() *CashOutForecast6 {
	newValue := new(CashOutForecast6)
	e.EstimatedCashOutForecastDetails = append(e.EstimatedCashOutForecastDetails, newValue)
	return newValue
}

func (e *EstimatedFundCashForecast6) AddEstimatedNetCashForecastDetails() *NetCashForecast4 {
	newValue := new(NetCashForecast4)
	e.EstimatedNetCashForecastDetails = append(e.EstimatedNetCashForecastDetails, newValue)
	return newValue
}
