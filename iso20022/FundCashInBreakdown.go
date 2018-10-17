package iso20022

// Breakdown of cash movements into a fund as a result of investment funds transactions, eg, subscriptions or switch-in.
type FundCashInBreakdown1 struct {

	// Amount of cash flow in, expressed as an amount of money.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`

	// Amount of the cash flow in,  expressed as a number of units.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb,omitempty"`

	// Indicates whether the cash flow is an item that did not appear on the previously sent report, eg, because it was received close to cut-off time.
	NewAmountIndicator *YesNoIndicator `xml:"NewAmtInd,omitempty"`

	// Breakdown of the cash movements into a fund by transaction type, eg, subscription, switch-in.
	InvestmentFundTransactionInTypeDetails *InvestmentFundTransactionInType1 `xml:"InvstmtFndTxInTpDtls,omitempty"`

	// Breakdown of the cash movements into a fund by order type, eg, order by quantity of units or amount of money.
	OriginalOrderQuantityDetails *OriginalOrderQuantityType1 `xml:"OrgnlOrdrQtyDtls,omitempty"`

	// Information related to the commission applied to an order, eg, back-end or front-end commission.
	CommissionDetails []*Commission4 `xml:"ComssnDtls,omitempty"`
}

func (f *FundCashInBreakdown1) SetAmount(value, currency string) {
	f.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundCashInBreakdown1) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	f.UnitsNumber = new(FinancialInstrumentQuantity1)
	return f.UnitsNumber
}

func (f *FundCashInBreakdown1) SetNewAmountIndicator(value string) {
	f.NewAmountIndicator = (*YesNoIndicator)(&value)
}

func (f *FundCashInBreakdown1) AddInvestmentFundTransactionInTypeDetails() *InvestmentFundTransactionInType1 {
	f.InvestmentFundTransactionInTypeDetails = new(InvestmentFundTransactionInType1)
	return f.InvestmentFundTransactionInTypeDetails
}

func (f *FundCashInBreakdown1) AddOriginalOrderQuantityDetails() *OriginalOrderQuantityType1 {
	f.OriginalOrderQuantityDetails = new(OriginalOrderQuantityType1)
	return f.OriginalOrderQuantityDetails
}

func (f *FundCashInBreakdown1) AddCommissionDetails() *Commission4 {
	newValue := new(Commission4)
	f.CommissionDetails = append(f.CommissionDetails, newValue)
	return newValue
}

// Breakdown of cash movements into a fund as a result of investment funds transactions, eg, subscriptions or switch-in.
type FundCashInBreakdown2 struct {

	// Amount of cash flow in, expressed as an amount of money.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`

	// Amount of the cash flow in,  expressed as a number of units.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb,omitempty"`

	// Indicates whether the cash flow is an item that did not appear on the previously sent report, eg, because it was received close to cut-off time.
	NewAmountIndicator *YesNoIndicator `xml:"NewAmtInd,omitempty"`

	// Breakdown of the cash movements into a fund by transaction type, eg, subscription, switch-in.
	InvestmentFundTransactionInType *InvestmentFundTransactionInType1Code `xml:"InvstmtFndTxInTp"`

	// Breakdown of the cash movements into a fund by transaction type, eg, subscription, switch-in.
	ExtendedInvestmentFundTransactionInType *Extended350Code `xml:"XtndedInvstmtFndTxInTp"`

	// Breakdown of the cash movements into a fund by order type, eg, order by quantity of units or amount of money.
	OriginalOrderQuantityType *OrderQuantityType2Code `xml:"OrgnlOrdrQtyTp"`

	// Breakdown of the cash movements into a fund by transaction type, eg, subscription, switch-in.
	ExtendedOriginalOrderQuantityType *Extended350Code `xml:"XtndedOrgnlOrdrQtyTp"`

	// Charge for the placement of an order.
	ChargeDetails []*Charge16 `xml:"ChrgDtls,omitempty"`

	// Information related to the commission applied to an order, eg, back-end or front-end commission.
	CommissionDetails []*Commission9 `xml:"ComssnDtls,omitempty"`
}

func (f *FundCashInBreakdown2) SetAmount(value, currency string) {
	f.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundCashInBreakdown2) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	f.UnitsNumber = new(FinancialInstrumentQuantity1)
	return f.UnitsNumber
}

func (f *FundCashInBreakdown2) SetNewAmountIndicator(value string) {
	f.NewAmountIndicator = (*YesNoIndicator)(&value)
}

func (f *FundCashInBreakdown2) SetInvestmentFundTransactionInType(value string) {
	f.InvestmentFundTransactionInType = (*InvestmentFundTransactionInType1Code)(&value)
}

func (f *FundCashInBreakdown2) SetExtendedInvestmentFundTransactionInType(value string) {
	f.ExtendedInvestmentFundTransactionInType = (*Extended350Code)(&value)
}

func (f *FundCashInBreakdown2) SetOriginalOrderQuantityType(value string) {
	f.OriginalOrderQuantityType = (*OrderQuantityType2Code)(&value)
}

func (f *FundCashInBreakdown2) SetExtendedOriginalOrderQuantityType(value string) {
	f.ExtendedOriginalOrderQuantityType = (*Extended350Code)(&value)
}

func (f *FundCashInBreakdown2) AddChargeDetails() *Charge16 {
	newValue := new(Charge16)
	f.ChargeDetails = append(f.ChargeDetails, newValue)
	return newValue
}

func (f *FundCashInBreakdown2) AddCommissionDetails() *Commission9 {
	newValue := new(Commission9)
	f.CommissionDetails = append(f.CommissionDetails, newValue)
	return newValue
}

// Breakdown of cash movements into a fund as a result of investment funds transactions, eg, subscriptions or switch-in.
type FundCashInBreakdown3 struct {

	// Amount of cash flow in, expressed as an amount of money.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`

	// Amount of the cash flow in,  expressed as a number of units.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb,omitempty"`

	// Indicates whether the cash flow is an item that did not appear on the previously sent report, for example, because it was received close to cut-off time.
	NewAmountIndicator *YesNoIndicator `xml:"NewAmtInd,omitempty"`

	// Type of transaction that resulted in the cash-in movement, for example, subscription, switch-in.
	InvestmentFundTransactionInType *InvestmentFundTransactionInType1Choice `xml:"InvstmtFndTxInTp"`

	// Specifies how the original order was expressed that resulted in the cash-in movement, that is cash or units.
	OriginalOrderQuantityType *QuantityType1Choice `xml:"OrgnlOrdrQtyTp"`

	// Charge for the placement of an order.
	ChargeDetails []*Charge26 `xml:"ChrgDtls,omitempty"`

	// Information related to the commission applied to an order, for example, back-end or front-end commission.
	CommissionDetails []*Commission21 `xml:"ComssnDtls,omitempty"`

	// Settlement currency for the transaction.
	SettlementCurrency *ActiveCurrencyCode `xml:"SttlmCcy,omitempty"`
}

func (f *FundCashInBreakdown3) SetAmount(value, currency string) {
	f.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundCashInBreakdown3) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	f.UnitsNumber = new(FinancialInstrumentQuantity1)
	return f.UnitsNumber
}

func (f *FundCashInBreakdown3) SetNewAmountIndicator(value string) {
	f.NewAmountIndicator = (*YesNoIndicator)(&value)
}

func (f *FundCashInBreakdown3) AddInvestmentFundTransactionInType() *InvestmentFundTransactionInType1Choice {
	f.InvestmentFundTransactionInType = new(InvestmentFundTransactionInType1Choice)
	return f.InvestmentFundTransactionInType
}

func (f *FundCashInBreakdown3) AddOriginalOrderQuantityType() *QuantityType1Choice {
	f.OriginalOrderQuantityType = new(QuantityType1Choice)
	return f.OriginalOrderQuantityType
}

func (f *FundCashInBreakdown3) AddChargeDetails() *Charge26 {
	newValue := new(Charge26)
	f.ChargeDetails = append(f.ChargeDetails, newValue)
	return newValue
}

func (f *FundCashInBreakdown3) AddCommissionDetails() *Commission21 {
	newValue := new(Commission21)
	f.CommissionDetails = append(f.CommissionDetails, newValue)
	return newValue
}

func (f *FundCashInBreakdown3) SetSettlementCurrency(value string) {
	f.SettlementCurrency = (*ActiveCurrencyCode)(&value)
}
