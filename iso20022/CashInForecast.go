package iso20022

// Cash movements into a fund as a result of investment funds transactions, eg, subscriptions or switch-in.
type CashInForecast1 struct {

	// Date on which cash is available.
	SettlementDate *ISODate `xml:"SttlmDt"`

	// Sub-total amount of the cash flow in, expressed as an amount of money.
	SubTotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"SubTtlAmt,omitempty"`

	// Sub-total amount of the cash flow in, expressed as a number of units.
	SubTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"SubTtlUnitsNb,omitempty"`

	// Indicates whether the estimated cash flow in is exceptional.
	ExceptionalCashFlowIndicator *YesNoIndicator `xml:"XcptnlCshFlowInd,omitempty"`

	// Breakdown of cash in amounts, eg, by transaction and order type.
	CashInBreakdownDetails []*FundCashInBreakdown1 `xml:"CshInBrkdwnDtls,omitempty"`
}

func (c *CashInForecast1) SetSettlementDate(value string) {
	c.SettlementDate = (*ISODate)(&value)
}

func (c *CashInForecast1) SetSubTotalAmount(value, currency string) {
	c.SubTotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *CashInForecast1) AddSubTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	c.SubTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return c.SubTotalUnitsNumber
}

func (c *CashInForecast1) SetExceptionalCashFlowIndicator(value string) {
	c.ExceptionalCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (c *CashInForecast1) AddCashInBreakdownDetails() *FundCashInBreakdown1 {
	newValue := new(FundCashInBreakdown1)
	c.CashInBreakdownDetails = append(c.CashInBreakdownDetails, newValue)
	return newValue
}

// Cash movements into a fund as a result of investment funds transactions, eg, subscriptions or switch-in.
type CashInForecast2 struct {

	// Date on which cash is available.
	SettlementDate *ISODate `xml:"SttlmDt"`

	// Sub-total amount of the cash flow in, expressed as an amount of money.
	SubTotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"SubTtlAmt,omitempty"`

	// Sub-total amount of the cash flow in, expressed as a number of units.
	SubTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"SubTtlUnitsNb,omitempty"`

	// Indicates whether the cash flow in is exceptional.
	ExceptionalCashFlowIndicator *YesNoIndicator `xml:"XcptnlCshFlowInd,omitempty"`
}

func (c *CashInForecast2) SetSettlementDate(value string) {
	c.SettlementDate = (*ISODate)(&value)
}

func (c *CashInForecast2) SetSubTotalAmount(value, currency string) {
	c.SubTotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *CashInForecast2) AddSubTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	c.SubTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return c.SubTotalUnitsNumber
}

func (c *CashInForecast2) SetExceptionalCashFlowIndicator(value string) {
	c.ExceptionalCashFlowIndicator = (*YesNoIndicator)(&value)
}

// Cash movements into a fund as a result of investment funds transactions, eg, subscriptions or switch-in.
type CashInForecast3 struct {

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt"`

	// Sub-total amount of the cash flow in, expressed as an amount of money.
	SubTotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"SubTtlAmt,omitempty"`

	// Sub-total amount of the cash flow in, expressed as a number of units.
	SubTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"SubTtlUnitsNb,omitempty"`

	// Indicates whether the estimated cash flow in is exceptional.
	ExceptionalCashFlowIndicator *YesNoIndicator `xml:"XcptnlCshFlowInd,omitempty"`

	// Breakdown of cash in amounts, eg, by transaction and order type.
	CashInBreakdownDetails []*FundCashInBreakdown2 `xml:"CshInBrkdwnDtls,omitempty"`
}

func (c *CashInForecast3) SetCashSettlementDate(value string) {
	c.CashSettlementDate = (*ISODate)(&value)
}

func (c *CashInForecast3) SetSubTotalAmount(value, currency string) {
	c.SubTotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *CashInForecast3) AddSubTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	c.SubTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return c.SubTotalUnitsNumber
}

func (c *CashInForecast3) SetExceptionalCashFlowIndicator(value string) {
	c.ExceptionalCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (c *CashInForecast3) AddCashInBreakdownDetails() *FundCashInBreakdown2 {
	newValue := new(FundCashInBreakdown2)
	c.CashInBreakdownDetails = append(c.CashInBreakdownDetails, newValue)
	return newValue
}

// Cash movements into a fund as a result of investment funds transactions, eg, subscriptions or switch-in.
type CashInForecast4 struct {

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt"`

	// Sub-total amount of the cash flow in, expressed as an amount of money.
	SubTotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"SubTtlAmt,omitempty"`

	// Sub-total amount of the cash flow in, expressed as a number of units.
	SubTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"SubTtlUnitsNb,omitempty"`

	// Indicates whether the cash flow in is exceptional.
	ExceptionalCashFlowIndicator *YesNoIndicator `xml:"XcptnlCshFlowInd,omitempty"`
}

func (c *CashInForecast4) SetCashSettlementDate(value string) {
	c.CashSettlementDate = (*ISODate)(&value)
}

func (c *CashInForecast4) SetSubTotalAmount(value, currency string) {
	c.SubTotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *CashInForecast4) AddSubTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	c.SubTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return c.SubTotalUnitsNumber
}

func (c *CashInForecast4) SetExceptionalCashFlowIndicator(value string) {
	c.ExceptionalCashFlowIndicator = (*YesNoIndicator)(&value)
}

// Cash movements into a fund as a result of investment funds transactions, eg, subscriptions or switch-in.
type CashInForecast5 struct {

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt"`

	// Sub-total amount of the cash flow in, expressed as an amount of money.
	SubTotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"SubTtlAmt,omitempty"`

	// Sub-total amount of the cash flow in, expressed as a number of units.
	SubTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"SubTtlUnitsNb,omitempty"`

	// Indicates whether the estimated cash flow in is exceptional.
	ExceptionalCashFlowIndicator *YesNoIndicator `xml:"XcptnlCshFlowInd,omitempty"`

	// Breakdown of cash in amounts by transaction and order type.
	CashInBreakdownDetails []*FundCashInBreakdown3 `xml:"CshInBrkdwnDtls,omitempty"`

	// Additional balances for cash amounts and number of units.
	// In an estimated report, the total cash derived from orders placed as a number of units is an estimated cash amount and the total number of units derived from orders placed as a cash amount is an estimated number of units.
	AdditionalBalance *FundBalance1 `xml:"AddtlBal,omitempty"`
}

func (c *CashInForecast5) SetCashSettlementDate(value string) {
	c.CashSettlementDate = (*ISODate)(&value)
}

func (c *CashInForecast5) SetSubTotalAmount(value, currency string) {
	c.SubTotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *CashInForecast5) AddSubTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	c.SubTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return c.SubTotalUnitsNumber
}

func (c *CashInForecast5) SetExceptionalCashFlowIndicator(value string) {
	c.ExceptionalCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (c *CashInForecast5) AddCashInBreakdownDetails() *FundCashInBreakdown3 {
	newValue := new(FundCashInBreakdown3)
	c.CashInBreakdownDetails = append(c.CashInBreakdownDetails, newValue)
	return newValue
}

func (c *CashInForecast5) AddAdditionalBalance() *FundBalance1 {
	c.AdditionalBalance = new(FundBalance1)
	return c.AdditionalBalance
}

// Cash movements into a fund as a result of investment funds transactions, eg, subscriptions or switch-in.
type CashInForecast6 struct {

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt"`

	// Sub-total amount of the cash flow in, expressed as an amount of money.
	SubTotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"SubTtlAmt,omitempty"`

	// Sub-total amount of the cash flow in, expressed as a number of units.
	SubTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"SubTtlUnitsNb,omitempty"`

	// Indicates whether the cash flow in is exceptional.
	ExceptionalCashFlowIndicator *YesNoIndicator `xml:"XcptnlCshFlowInd,omitempty"`

	// Additional balances for cash amounts and number of units.
	// In an estimated report, the total cash derived from orders placed as a number of units is an estimated cash amount and the total number of units derived from orders placed as a cash amount is an estimated number of units.
	AdditionalBalance *FundBalance1 `xml:"AddtlBal,omitempty"`
}

func (c *CashInForecast6) SetCashSettlementDate(value string) {
	c.CashSettlementDate = (*ISODate)(&value)
}

func (c *CashInForecast6) SetSubTotalAmount(value, currency string) {
	c.SubTotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *CashInForecast6) AddSubTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	c.SubTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return c.SubTotalUnitsNumber
}

func (c *CashInForecast6) SetExceptionalCashFlowIndicator(value string) {
	c.ExceptionalCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (c *CashInForecast6) AddAdditionalBalance() *FundBalance1 {
	c.AdditionalBalance = new(FundBalance1)
	return c.AdditionalBalance
}
