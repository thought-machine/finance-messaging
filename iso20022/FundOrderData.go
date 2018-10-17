package iso20022

// Extract of trade data for an investment fund order.
type FundOrderData1 struct {

	// Account information of the individual order instruction for which the status is given.
	InvestmentAccountDetails *InvestmentAccount13 `xml:"InvstmtAcctDtls,omitempty"`

	// Financial instrument information of the individual order instruction for which the status is given.
	FinancialInstrumentDetails *FinancialInstrument10 `xml:"FinInstrmDtls,omitempty"`

	// Quantity of investment fund units subscribed or redeemed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb,omitempty"`

	// Amount of money used to derive the quantity of investment fund units to be sold or subscribed to.
	NetAmount *ActiveOrHistoricCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Amount of money used to derive the quantity of investment fund units to be sold or subscribed to, including all charges, commissions, and tax.
	GrossAmount *ActiveOrHistoricCurrencyAndAmount `xml:"GrssAmt,omitempty"`

	// Portion of the investor's holdings, in a specific investment fund/ fund class, that is redeemed.
	HoldingsRedemptionRate *PercentageRate `xml:"HldgsRedRate,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveOrHistoricCurrencyCode `xml:"UnitCcy,omitempty"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *ActiveOrHistoricCurrencyCode `xml:"QtdCcy,omitempty"`
}

func (f *FundOrderData1) AddInvestmentAccountDetails() *InvestmentAccount13 {
	f.InvestmentAccountDetails = new(InvestmentAccount13)
	return f.InvestmentAccountDetails
}

func (f *FundOrderData1) AddFinancialInstrumentDetails() *FinancialInstrument10 {
	f.FinancialInstrumentDetails = new(FinancialInstrument10)
	return f.FinancialInstrumentDetails
}

func (f *FundOrderData1) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	f.UnitsNumber = new(FinancialInstrumentQuantity1)
	return f.UnitsNumber
}

func (f *FundOrderData1) SetNetAmount(value, currency string) {
	f.NetAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundOrderData1) SetGrossAmount(value, currency string) {
	f.GrossAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundOrderData1) SetHoldingsRedemptionRate(value string) {
	f.HoldingsRedemptionRate = (*PercentageRate)(&value)
}

func (f *FundOrderData1) SetSettlementAmount(value, currency string) {
	f.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *FundOrderData1) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FundOrderData1) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

// Extract of trade data for an investment fund switch order.
type FundOrderData2 struct {

	// Amount of money used to derive the quantity of investment fund units to be redeemed.
	TotalRedemptionAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlRedAmt,omitempty"`

	// Amount of money used to derive the quantity of investment fund units to be subscribed.
	TotalSubscriptionAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlSbcptAmt,omitempty"`

	// Amount of money to be transferred between the debtor and creditor before bank transaction charges.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Method by which the transaction is settled.
	SettlementMethod *DeliveryReceiptType2Code `xml:"SttlmMtd,omitempty"`

	// Additional amount of money paid by the investor in addition to the switch redemption amount.
	AdditionalCashIn *ActiveOrHistoricCurrencyAndAmount `xml:"AddtlCshIn,omitempty"`

	// Amount of money that results from a switch-out, that is not reinvested in another investment fund, and is repaid to the investor.
	ResultingCashOut *ActiveOrHistoricCurrencyAndAmount `xml:"RsltgCshOut,omitempty"`

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveOrHistoricCurrencyCode `xml:"UnitCcy,omitempty"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *ActiveOrHistoricCurrencyCode `xml:"QtdCcy,omitempty"`
}

func (f *FundOrderData2) SetTotalRedemptionAmount(value, currency string) {
	f.TotalRedemptionAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundOrderData2) SetTotalSubscriptionAmount(value, currency string) {
	f.TotalSubscriptionAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundOrderData2) SetSettlementAmount(value, currency string) {
	f.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *FundOrderData2) SetSettlementMethod(value string) {
	f.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (f *FundOrderData2) SetAdditionalCashIn(value, currency string) {
	f.AdditionalCashIn = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundOrderData2) SetResultingCashOut(value, currency string) {
	f.ResultingCashOut = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundOrderData2) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FundOrderData2) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

// Extract of trade data for an investment fund order.
type FundOrderData5 struct {

	// Account information of the individual order instruction for which the status is given.
	InvestmentAccountDetails *InvestmentAccount58 `xml:"InvstmtAcctDtls,omitempty"`

	// Financial instrument information of the individual order instruction for which the status is given.
	FinancialInstrumentDetails *FinancialInstrument57 `xml:"FinInstrmDtls,omitempty"`

	// Number of investment fund units subscribed or redeemed.
	UnitsNumber *DecimalNumber `xml:"UnitsNb,omitempty"`

	// When the status message is used for a subscription, this is the amount of money to be invested in the fund.
	// Net Amount = Quantity * Price.
	// When the status message is used for a redemption, this is the amount of money to be received following redemption of fund units.
	// Net Amount = (Quantity * Price) - (Fees + Taxes).
	NetAmount *ActiveOrHistoricCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// When the status message is used for a subscription, this is the amount of money to be paid by the investor when subscribing to fund units.
	// Gross amount = (Quantity * Price) + (Fees + Taxes).
	// When the status message is used for a redemption, this is the amount of money to be redeemed from the fund.
	// Gross Amount = Quantity * Price.
	GrossAmount *ActiveOrHistoricCurrencyAndAmount `xml:"GrssAmt,omitempty"`

	// Portion of the investor's holdings, in a specific investment fund/fund class, that is redeemed.
	HoldingsRedemptionRate *PercentageRate `xml:"HldgsRedRate,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Currency from which the quoted currency is converted in an exchange rate calculation.
	// 1 x <UnitCcy> = <XchgRate> x <QtdCcy>.
	UnitCurrency *ActiveCurrencyCode `xml:"UnitCcy,omitempty"`

	// Currency into which the unit currency is converted in an exchange rate calculation.
	// 1 x <UnitCcy> = <XchgRate> x <QtdCcy>.
	QuotedCurrency *ActiveCurrencyCode `xml:"QtdCcy,omitempty"`
}

func (f *FundOrderData5) AddInvestmentAccountDetails() *InvestmentAccount58 {
	f.InvestmentAccountDetails = new(InvestmentAccount58)
	return f.InvestmentAccountDetails
}

func (f *FundOrderData5) AddFinancialInstrumentDetails() *FinancialInstrument57 {
	f.FinancialInstrumentDetails = new(FinancialInstrument57)
	return f.FinancialInstrumentDetails
}

func (f *FundOrderData5) SetUnitsNumber(value string) {
	f.UnitsNumber = (*DecimalNumber)(&value)
}

func (f *FundOrderData5) SetNetAmount(value, currency string) {
	f.NetAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundOrderData5) SetGrossAmount(value, currency string) {
	f.GrossAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundOrderData5) SetHoldingsRedemptionRate(value string) {
	f.HoldingsRedemptionRate = (*PercentageRate)(&value)
}

func (f *FundOrderData5) SetSettlementAmount(value, currency string) {
	f.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *FundOrderData5) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *FundOrderData5) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveCurrencyCode)(&value)
}

// Extract of trade data for an investment fund switch order.
type FundOrderData6 struct {

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Method by which the transaction is settled.
	SettlementMethod *DeliveryReceiptType2Code `xml:"SttlmMtd,omitempty"`

	// Choice between additional cash in or resulting cash out.
	AdditionalAmount *AdditionalAmount1Choice `xml:"AddtlAmt,omitempty"`

	// Currency from which the quoted currency is converted in an exchange rate calculation.
	// 1 x <UnitCcy> = <XchgRate> x <QtdCcy>.
	UnitCurrency *ActiveCurrencyCode `xml:"UnitCcy,omitempty"`

	// Currency into which the unit currency is converted in an exchange rate calculation.
	// 1 x <UnitCcy> = <XchgRate> x <QtdCcy>.
	QuotedCurrency *ActiveCurrencyCode `xml:"QtdCcy,omitempty"`
}

func (f *FundOrderData6) SetSettlementAmount(value, currency string) {
	f.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *FundOrderData6) SetSettlementMethod(value string) {
	f.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (f *FundOrderData6) AddAdditionalAmount() *AdditionalAmount1Choice {
	f.AdditionalAmount = new(AdditionalAmount1Choice)
	return f.AdditionalAmount
}

func (f *FundOrderData6) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *FundOrderData6) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveCurrencyCode)(&value)
}
