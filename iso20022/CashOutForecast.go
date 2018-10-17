package iso20022

// Cash movements out of a fund as a result of investment funds transactions, eg, redemptions or switch-out.
type CashOutForecast1 struct {

	// Date on which cash is available.
	SettlementDate *ISODate `xml:"SttlmDt"`

	// Sub-total amount of the cash flow out, expressed as an amount of money.
	SubTotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"SubTtlAmt,omitempty"`

	// Sub-total amount of the cash flow out, expressed as a number of units.
	SubTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"SubTtlUnitsNb,omitempty"`

	// Indicates whether the estimated cash flow out is exceptional.
	ExceptionalCashFlowIndicator *YesNoIndicator `xml:"XcptnlCshFlowInd,omitempty"`

	// Breakdown of cash out amounts, eg,  by transaction or order type.
	CashOutBreakdownDetails []*FundCashOutBreakdown1 `xml:"CshOutBrkdwnDtls,omitempty"`
}

func (c *CashOutForecast1) SetSettlementDate(value string) {
	c.SettlementDate = (*ISODate)(&value)
}

func (c *CashOutForecast1) SetSubTotalAmount(value, currency string) {
	c.SubTotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *CashOutForecast1) AddSubTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	c.SubTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return c.SubTotalUnitsNumber
}

func (c *CashOutForecast1) SetExceptionalCashFlowIndicator(value string) {
	c.ExceptionalCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (c *CashOutForecast1) AddCashOutBreakdownDetails() *FundCashOutBreakdown1 {
	newValue := new(FundCashOutBreakdown1)
	c.CashOutBreakdownDetails = append(c.CashOutBreakdownDetails, newValue)
	return newValue
}

// Cash movements out of a fund as a result of investment funds transactions, eg, redemptions or switch-out.
type CashOutForecast2 struct {

	// Date on which cash is available.
	SettlementDate *ISODate `xml:"SttlmDt"`

	// Sub-total amount of the cash flow out, expressed as an amount of money.
	SubTotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"SubTtlAmt,omitempty"`

	// Sub-total amount of the cash flow out, expressed as a number of units.
	SubTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"SubTtlUnitsNb,omitempty"`

	// Indicates whether the cash flow out is exceptional.
	ExceptionalCashFlowIndicator *YesNoIndicator `xml:"XcptnlCshFlowInd,omitempty"`
}

func (c *CashOutForecast2) SetSettlementDate(value string) {
	c.SettlementDate = (*ISODate)(&value)
}

func (c *CashOutForecast2) SetSubTotalAmount(value, currency string) {
	c.SubTotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *CashOutForecast2) AddSubTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	c.SubTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return c.SubTotalUnitsNumber
}

func (c *CashOutForecast2) SetExceptionalCashFlowIndicator(value string) {
	c.ExceptionalCashFlowIndicator = (*YesNoIndicator)(&value)
}

// Cash movements out of a fund as a result of investment funds transactions, eg, redemptions or switch-out.
type CashOutForecast3 struct {

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt"`

	// Sub-total amount of the cash flow out, expressed as an amount of money.
	SubTotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"SubTtlAmt,omitempty"`

	// Sub-total amount of the cash flow out, expressed as a number of units.
	SubTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"SubTtlUnitsNb,omitempty"`

	// Indicates whether the estimated cash flow out is exceptional.
	ExceptionalCashFlowIndicator *YesNoIndicator `xml:"XcptnlCshFlowInd,omitempty"`

	// Breakdown of cash out amounts, eg,  by transaction or order type.
	CashOutBreakdownDetails []*FundCashOutBreakdown2 `xml:"CshOutBrkdwnDtls,omitempty"`
}

func (c *CashOutForecast3) SetCashSettlementDate(value string) {
	c.CashSettlementDate = (*ISODate)(&value)
}

func (c *CashOutForecast3) SetSubTotalAmount(value, currency string) {
	c.SubTotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *CashOutForecast3) AddSubTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	c.SubTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return c.SubTotalUnitsNumber
}

func (c *CashOutForecast3) SetExceptionalCashFlowIndicator(value string) {
	c.ExceptionalCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (c *CashOutForecast3) AddCashOutBreakdownDetails() *FundCashOutBreakdown2 {
	newValue := new(FundCashOutBreakdown2)
	c.CashOutBreakdownDetails = append(c.CashOutBreakdownDetails, newValue)
	return newValue
}

// Cash movements out of a fund as a result of investment funds transactions, eg, redemptions or switch-out.
type CashOutForecast4 struct {

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt"`

	// Sub-total amount of the cash flow out, expressed as an amount of money.
	SubTotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"SubTtlAmt,omitempty"`

	// Sub-total amount of the cash flow out, expressed as a number of units.
	SubTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"SubTtlUnitsNb,omitempty"`

	// Indicates whether the cash flow out is exceptional.
	ExceptionalCashFlowIndicator *YesNoIndicator `xml:"XcptnlCshFlowInd,omitempty"`
}

func (c *CashOutForecast4) SetCashSettlementDate(value string) {
	c.CashSettlementDate = (*ISODate)(&value)
}

func (c *CashOutForecast4) SetSubTotalAmount(value, currency string) {
	c.SubTotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *CashOutForecast4) AddSubTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	c.SubTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return c.SubTotalUnitsNumber
}

func (c *CashOutForecast4) SetExceptionalCashFlowIndicator(value string) {
	c.ExceptionalCashFlowIndicator = (*YesNoIndicator)(&value)
}

// Cash movements out of a fund as a result of investment funds transactions, eg, redemptions or switch-out.
type CashOutForecast5 struct {

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt"`

	// Sub-total amount of the cash flow out, expressed as an amount of money.
	SubTotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"SubTtlAmt,omitempty"`

	// Sub-total amount of the cash flow out, expressed as a number of units.
	SubTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"SubTtlUnitsNb,omitempty"`

	// Indicates whether the estimated cash flow out is exceptional.
	ExceptionalCashFlowIndicator *YesNoIndicator `xml:"XcptnlCshFlowInd,omitempty"`

	// Breakdown of cash out amounts by transaction and order type.
	CashOutBreakdownDetails []*FundCashOutBreakdown3 `xml:"CshOutBrkdwnDtls,omitempty"`

	// Additional balances for cash amounts and number of units.
	// In an estimated report, the total cash derived from orders placed as a number of units is an estimated cash amount and the total number of units derived from orders placed as a cash amount is an estimated number of units.
	AdditionalBalance *FundBalance1 `xml:"AddtlBal,omitempty"`
}

func (c *CashOutForecast5) SetCashSettlementDate(value string) {
	c.CashSettlementDate = (*ISODate)(&value)
}

func (c *CashOutForecast5) SetSubTotalAmount(value, currency string) {
	c.SubTotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *CashOutForecast5) AddSubTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	c.SubTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return c.SubTotalUnitsNumber
}

func (c *CashOutForecast5) SetExceptionalCashFlowIndicator(value string) {
	c.ExceptionalCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (c *CashOutForecast5) AddCashOutBreakdownDetails() *FundCashOutBreakdown3 {
	newValue := new(FundCashOutBreakdown3)
	c.CashOutBreakdownDetails = append(c.CashOutBreakdownDetails, newValue)
	return newValue
}

func (c *CashOutForecast5) AddAdditionalBalance() *FundBalance1 {
	c.AdditionalBalance = new(FundBalance1)
	return c.AdditionalBalance
}

// Cash movements out of a fund as a result of investment funds transactions, eg, redemptions or switch-out.
type CashOutForecast6 struct {

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt"`

	// Sub-total amount of the cash flow out, expressed as an amount of money.
	SubTotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"SubTtlAmt,omitempty"`

	// Sub-total amount of the cash flow out, expressed as a number of units.
	SubTotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"SubTtlUnitsNb,omitempty"`

	// Indicates whether the cash flow out is exceptional.
	ExceptionalCashFlowIndicator *YesNoIndicator `xml:"XcptnlCshFlowInd,omitempty"`

	// Additional balances for cash amounts and number of units.
	// In an estimated report, the total cash derived from orders placed as a number of units is an estimated cash amount and the total number of units derived from orders placed as a cash amount is an estimated number of units.
	AdditionalBalance *FundBalance1 `xml:"AddtlBal,omitempty"`
}

func (c *CashOutForecast6) SetCashSettlementDate(value string) {
	c.CashSettlementDate = (*ISODate)(&value)
}

func (c *CashOutForecast6) SetSubTotalAmount(value, currency string) {
	c.SubTotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *CashOutForecast6) AddSubTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	c.SubTotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return c.SubTotalUnitsNumber
}

func (c *CashOutForecast6) SetExceptionalCashFlowIndicator(value string) {
	c.ExceptionalCashFlowIndicator = (*YesNoIndicator)(&value)
}

func (c *CashOutForecast6) AddAdditionalBalance() *FundBalance1 {
	c.AdditionalBalance = new(FundBalance1)
	return c.AdditionalBalance
}
