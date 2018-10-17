package iso20022

// Breakdown of cash movements out of a fund as a result of investment funds transactions, eg, redemptions or switch-out.
type FundCashOutBreakdown1 struct {

	// Amount of cash flow out, expressed as an amount of money.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`

	// Amount of the cash flow out, expressed as a number of units.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb,omitempty"`

	// Indicates whether the cash flow is an item that did not appear on the previously sent report, eg, because it was received close to cut-off time.
	NewAmountIndicator *YesNoIndicator `xml:"NewAmtInd,omitempty"`

	// Breakdown of the cash movements out of a fund by transaction type, eg, redemption, switch-out.
	InvestmentFundTransactionOutTypeDetails *InvestmentFundTransactionOutType1 `xml:"InvstmtFndTxOutTpDtls,omitempty"`

	// Breakdown of the cash movements into a fund by order type, eg, order by quantity of units or amount of money.
	OriginalOrderQuantityDetails *OriginalOrderQuantityType1 `xml:"OrgnlOrdrQtyDtls,omitempty"`

	// Information related to the commission applied to an order, eg, back-end or front-end commission.
	CommissionDetails []*Commission4 `xml:"ComssnDtls,omitempty"`
}

func (f *FundCashOutBreakdown1) SetAmount(value, currency string) {
	f.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundCashOutBreakdown1) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	f.UnitsNumber = new(FinancialInstrumentQuantity1)
	return f.UnitsNumber
}

func (f *FundCashOutBreakdown1) SetNewAmountIndicator(value string) {
	f.NewAmountIndicator = (*YesNoIndicator)(&value)
}

func (f *FundCashOutBreakdown1) AddInvestmentFundTransactionOutTypeDetails() *InvestmentFundTransactionOutType1 {
	f.InvestmentFundTransactionOutTypeDetails = new(InvestmentFundTransactionOutType1)
	return f.InvestmentFundTransactionOutTypeDetails
}

func (f *FundCashOutBreakdown1) AddOriginalOrderQuantityDetails() *OriginalOrderQuantityType1 {
	f.OriginalOrderQuantityDetails = new(OriginalOrderQuantityType1)
	return f.OriginalOrderQuantityDetails
}

func (f *FundCashOutBreakdown1) AddCommissionDetails() *Commission4 {
	newValue := new(Commission4)
	f.CommissionDetails = append(f.CommissionDetails, newValue)
	return newValue
}

// Breakdown of cash movements out of a fund as a result of investment funds transactions, eg, redemptions or switch-out.
type FundCashOutBreakdown2 struct {

	// Amount of cash flow out, expressed as an amount of money.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`

	// Amount of the cash flow out, expressed as a number of units.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb,omitempty"`

	// Indicates whether the cash flow is an item that did not appear on the previously sent report, eg, because it was received close to cut-off time.
	NewAmountIndicator *YesNoIndicator `xml:"NewAmtInd,omitempty"`

	// Breakdown of the cash movements out of a fund by transaction type, eg, redemption, switch-out.
	InvestmentFundTransactionOutType *InvestmentFundTransactionOutType1Code `xml:"InvstmtFndTxOutTp"`

	// Breakdown of the cash movements out of a fund by transaction type, eg, redemption, switch-out.
	ExtendedInvestmentFundTransactionOutType *Extended350Code `xml:"XtndedInvstmtFndTxOutTp"`

	// Breakdown of the cash movements into a fund by order type, eg, order by quantity of units or amount of money.
	OriginalOrderQuantityType *OrderQuantityType2Code `xml:"OrgnlOrdrQtyTp"`

	// Breakdown of the cash movements into a fund by transaction type, eg, subscription, switch-in.
	ExtendedOriginalOrderQuantityType *Extended350Code `xml:"XtndedOrgnlOrdrQtyTp"`

	// Charge for the placement of an order.
	ChargeDetails []*Charge16 `xml:"ChrgDtls,omitempty"`

	// Information related to the commission applied to an order, eg, back-end or front-end commission.
	CommissionDetails []*Commission9 `xml:"ComssnDtls,omitempty"`
}

func (f *FundCashOutBreakdown2) SetAmount(value, currency string) {
	f.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundCashOutBreakdown2) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	f.UnitsNumber = new(FinancialInstrumentQuantity1)
	return f.UnitsNumber
}

func (f *FundCashOutBreakdown2) SetNewAmountIndicator(value string) {
	f.NewAmountIndicator = (*YesNoIndicator)(&value)
}

func (f *FundCashOutBreakdown2) SetInvestmentFundTransactionOutType(value string) {
	f.InvestmentFundTransactionOutType = (*InvestmentFundTransactionOutType1Code)(&value)
}

func (f *FundCashOutBreakdown2) SetExtendedInvestmentFundTransactionOutType(value string) {
	f.ExtendedInvestmentFundTransactionOutType = (*Extended350Code)(&value)
}

func (f *FundCashOutBreakdown2) SetOriginalOrderQuantityType(value string) {
	f.OriginalOrderQuantityType = (*OrderQuantityType2Code)(&value)
}

func (f *FundCashOutBreakdown2) SetExtendedOriginalOrderQuantityType(value string) {
	f.ExtendedOriginalOrderQuantityType = (*Extended350Code)(&value)
}

func (f *FundCashOutBreakdown2) AddChargeDetails() *Charge16 {
	newValue := new(Charge16)
	f.ChargeDetails = append(f.ChargeDetails, newValue)
	return newValue
}

func (f *FundCashOutBreakdown2) AddCommissionDetails() *Commission9 {
	newValue := new(Commission9)
	f.CommissionDetails = append(f.CommissionDetails, newValue)
	return newValue
}

// Breakdown of cash movements out of a fund as a result of investment funds transactions, eg, redemptions or switch-out.
type FundCashOutBreakdown3 struct {

	// Amount of cash flow out, expressed as an amount of money.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`

	// Amount of the cash flow out, expressed as a number of units.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb,omitempty"`

	// Indicates whether the cash flow is an item that did not appear on the previously sent report, for example, because it was received close to cut-off time.
	NewAmountIndicator *YesNoIndicator `xml:"NewAmtInd,omitempty"`

	// Type of transaction that resulted in the cash-out movement, for example, redemption, switch-out.
	InvestmentFundTransactionOutType *InvestmentFundTransactionOutType1Choice `xml:"InvstmtFndTxOutTp"`

	// Specifies how the original order was expressed that resulted in the cash-out movement, that is cash or units.
	OriginalOrderQuantityType *QuantityType1Choice `xml:"OrgnlOrdrQtyTp"`

	// Charge for the placement of an order.
	ChargeDetails []*Charge26 `xml:"ChrgDtls,omitempty"`

	// Information related to the commission applied to an order, for example, back-end or front-end commission.
	CommissionDetails []*Commission21 `xml:"ComssnDtls,omitempty"`

	// Settlement currency for the transaction.
	SettlementCurrency *ActiveCurrencyCode `xml:"SttlmCcy,omitempty"`
}

func (f *FundCashOutBreakdown3) SetAmount(value, currency string) {
	f.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundCashOutBreakdown3) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	f.UnitsNumber = new(FinancialInstrumentQuantity1)
	return f.UnitsNumber
}

func (f *FundCashOutBreakdown3) SetNewAmountIndicator(value string) {
	f.NewAmountIndicator = (*YesNoIndicator)(&value)
}

func (f *FundCashOutBreakdown3) AddInvestmentFundTransactionOutType() *InvestmentFundTransactionOutType1Choice {
	f.InvestmentFundTransactionOutType = new(InvestmentFundTransactionOutType1Choice)
	return f.InvestmentFundTransactionOutType
}

func (f *FundCashOutBreakdown3) AddOriginalOrderQuantityType() *QuantityType1Choice {
	f.OriginalOrderQuantityType = new(QuantityType1Choice)
	return f.OriginalOrderQuantityType
}

func (f *FundCashOutBreakdown3) AddChargeDetails() *Charge26 {
	newValue := new(Charge26)
	f.ChargeDetails = append(f.ChargeDetails, newValue)
	return newValue
}

func (f *FundCashOutBreakdown3) AddCommissionDetails() *Commission21 {
	newValue := new(Commission21)
	f.CommissionDetails = append(f.CommissionDetails, newValue)
	return newValue
}

func (f *FundCashOutBreakdown3) SetSettlementCurrency(value string) {
	f.SettlementCurrency = (*ActiveCurrencyCode)(&value)
}
