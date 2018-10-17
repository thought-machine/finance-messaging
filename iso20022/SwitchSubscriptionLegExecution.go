package iso20022

// Execution of the subscription part, in a switch between investment funds or investment fund classes.
type SwitchSubscriptionLegExecution2 struct {

	// Unique technical identifier for an instance of a leg within a switch.
	LegIdentification *Max35Text `xml:"LegId,omitempty"`

	// Unique identifier for an instance of a leg execution within a switch confirmation.
	LegExecutionIdentification *Max35Text `xml:"LegExctnId,omitempty"`

	// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
	FinancialInstrumentDetails *FinancialInstrument6 `xml:"FinInstrmDtls"`

	// Number of investment fund units subscribed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Net amount of money invested in a specific financial instrument by an investor, expressed in the currency requested by the investor.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt"`

	// Gross amount of money invested in a specific financial instrument by an investor, expressed in the currency requested by the investor.
	GrossAmount *ActiveCurrencyAndAmount `xml:"GrssAmt,omitempty"`

	// Date and time at which a price is applied, according to the terms stated in the prospectus.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Price at which the order was executed.
	PriceDetails *UnitPrice5 `xml:"PricDtls"`

	// Indicates whether the dividend is included, ie, cum-dividend, in the executed price. When the dividend is not included, the price will be ex-dividend.
	CumDividendIndicator *YesNoIndicator `xml:"CumDvddInd"`

	// Part of the price deemed as accrued income or profit rather than capital. The interim profit amount is used for tax purposes.
	InterimProfitAmount *ProfitAndLoss1Choice `xml:"IntrmPrftAmt,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *CurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *CurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Charge for the execution of an order.
	ChargeGeneralDetails *TotalCharges2 `xml:"ChrgGnlDtls,omitempty"`

	// Commission for the execution of an investment fund order.
	CommissionGeneralDetails *TotalCommissions2 `xml:"ComssnGnlDtls,omitempty"`

	// Tax applicable to execution of an investment fund order.
	TaxGeneralDetails *TotalTaxes2 `xml:"TaxGnlDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters3 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to physical delivery of the securities.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`
}

func (s *SwitchSubscriptionLegExecution2) SetLegIdentification(value string) {
	s.LegIdentification = (*Max35Text)(&value)
}

func (s *SwitchSubscriptionLegExecution2) SetLegExecutionIdentification(value string) {
	s.LegExecutionIdentification = (*Max35Text)(&value)
}

func (s *SwitchSubscriptionLegExecution2) AddFinancialInstrumentDetails() *FinancialInstrument6 {
	s.FinancialInstrumentDetails = new(FinancialInstrument6)
	return s.FinancialInstrumentDetails
}

func (s *SwitchSubscriptionLegExecution2) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	s.UnitsNumber = new(FinancialInstrumentQuantity1)
	return s.UnitsNumber
}

func (s *SwitchSubscriptionLegExecution2) SetNetAmount(value, currency string) {
	s.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchSubscriptionLegExecution2) SetGrossAmount(value, currency string) {
	s.GrossAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchSubscriptionLegExecution2) AddTradeDateTime() *DateAndDateTimeChoice {
	s.TradeDateTime = new(DateAndDateTimeChoice)
	return s.TradeDateTime
}

func (s *SwitchSubscriptionLegExecution2) AddPriceDetails() *UnitPrice5 {
	s.PriceDetails = new(UnitPrice5)
	return s.PriceDetails
}

func (s *SwitchSubscriptionLegExecution2) SetCumDividendIndicator(value string) {
	s.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (s *SwitchSubscriptionLegExecution2) AddInterimProfitAmount() *ProfitAndLoss1Choice {
	s.InterimProfitAmount = new(ProfitAndLoss1Choice)
	return s.InterimProfitAmount
}

func (s *SwitchSubscriptionLegExecution2) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SwitchSubscriptionLegExecution2) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*CurrencyCode)(&value)
}

func (s *SwitchSubscriptionLegExecution2) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*CurrencyCode)(&value)
}

func (s *SwitchSubscriptionLegExecution2) AddChargeGeneralDetails() *TotalCharges2 {
	s.ChargeGeneralDetails = new(TotalCharges2)
	return s.ChargeGeneralDetails
}

func (s *SwitchSubscriptionLegExecution2) AddCommissionGeneralDetails() *TotalCommissions2 {
	s.CommissionGeneralDetails = new(TotalCommissions2)
	return s.CommissionGeneralDetails
}

func (s *SwitchSubscriptionLegExecution2) AddTaxGeneralDetails() *TotalTaxes2 {
	s.TaxGeneralDetails = new(TotalTaxes2)
	return s.TaxGeneralDetails
}

func (s *SwitchSubscriptionLegExecution2) AddSettlementAndCustodyDetails() *FundSettlementParameters3 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters3)
	return s.SettlementAndCustodyDetails
}

func (s *SwitchSubscriptionLegExecution2) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SwitchSubscriptionLegExecution2) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	s.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return s.PhysicalDeliveryDetails
}

// Execution of the subscription part, in a switch between investment funds or investment fund classes.
type SwitchSubscriptionLegExecution3 struct {

	// Unique technical identifier for an instance of a leg within a switch.
	LegIdentification *Max35Text `xml:"LegId,omitempty"`

	// Unique identifier for an instance of a leg execution within a switch confirmation.
	LegExecutionIdentification *Max35Text `xml:"LegExctnId,omitempty"`

	// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
	FinancialInstrumentDetails *FinancialInstrument10 `xml:"FinInstrmDtls"`

	// Number of investment fund units subscribed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
	InvestmentAccountDetails *InvestmentAccount21 `xml:"InvstmtAcctDtls,omitempty"`

	// Amount of money invested in a specific financial instrument by an investor, after deduction of charges, commissions and taxes.
	// [Quantity * Price]
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Amount of money invested in a specific financial instrument by an investor, before deduction of  charges, commissions, and taxes.
	// [(Quantity * Price) + (Charges + Commissions +Taxes)]
	GrossAmount *ActiveCurrencyAndAmount `xml:"GrssAmt,omitempty"`

	// Date and time at which a price is applied, according to the terms stated in the prospectus.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Price at which the order was executed.
	PriceDetails *UnitPrice10 `xml:"PricDtls"`

	// Indicates whether the dividend is included, ie, cum-dividend, in the executed price. When the dividend is not included, the price will be ex-dividend.
	CumDividendIndicator *YesNoIndicator `xml:"CumDvddInd"`

	// Part of the price deemed as accrued income or profit rather than capital. The interim profit amount is used for tax purposes.
	InterimProfitAmount *ProfitAndLoss1Choice `xml:"IntrmPrftAmt,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *ActiveCurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Charge for the execution of an order.
	ChargeGeneralDetails *TotalCharges3 `xml:"ChrgGnlDtls,omitempty"`

	// Commission for the execution of an investment fund order.
	CommissionGeneralDetails *TotalCommissions3 `xml:"ComssnGnlDtls,omitempty"`

	// Tax applicable to execution of an investment fund order.
	TaxGeneralDetails *TotalTaxes3 `xml:"TaxGnlDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters3 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to physical delivery of the securities.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Part of an investor's subscription amount that is held by the fund in order to pay incentive / performance fees at the end of the fiscal year.
	Equalisation *Equalisation1 `xml:"Equlstn,omitempty"`
}

func (s *SwitchSubscriptionLegExecution3) SetLegIdentification(value string) {
	s.LegIdentification = (*Max35Text)(&value)
}

func (s *SwitchSubscriptionLegExecution3) SetLegExecutionIdentification(value string) {
	s.LegExecutionIdentification = (*Max35Text)(&value)
}

func (s *SwitchSubscriptionLegExecution3) AddFinancialInstrumentDetails() *FinancialInstrument10 {
	s.FinancialInstrumentDetails = new(FinancialInstrument10)
	return s.FinancialInstrumentDetails
}

func (s *SwitchSubscriptionLegExecution3) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	s.UnitsNumber = new(FinancialInstrumentQuantity1)
	return s.UnitsNumber
}

func (s *SwitchSubscriptionLegExecution3) AddInvestmentAccountDetails() *InvestmentAccount21 {
	s.InvestmentAccountDetails = new(InvestmentAccount21)
	return s.InvestmentAccountDetails
}

func (s *SwitchSubscriptionLegExecution3) SetNetAmount(value, currency string) {
	s.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchSubscriptionLegExecution3) SetGrossAmount(value, currency string) {
	s.GrossAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchSubscriptionLegExecution3) AddTradeDateTime() *DateAndDateTimeChoice {
	s.TradeDateTime = new(DateAndDateTimeChoice)
	return s.TradeDateTime
}

func (s *SwitchSubscriptionLegExecution3) AddPriceDetails() *UnitPrice10 {
	s.PriceDetails = new(UnitPrice10)
	return s.PriceDetails
}

func (s *SwitchSubscriptionLegExecution3) SetCumDividendIndicator(value string) {
	s.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (s *SwitchSubscriptionLegExecution3) AddInterimProfitAmount() *ProfitAndLoss1Choice {
	s.InterimProfitAmount = new(ProfitAndLoss1Choice)
	return s.InterimProfitAmount
}

func (s *SwitchSubscriptionLegExecution3) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SwitchSubscriptionLegExecution3) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (s *SwitchSubscriptionLegExecution3) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (s *SwitchSubscriptionLegExecution3) AddChargeGeneralDetails() *TotalCharges3 {
	s.ChargeGeneralDetails = new(TotalCharges3)
	return s.ChargeGeneralDetails
}

func (s *SwitchSubscriptionLegExecution3) AddCommissionGeneralDetails() *TotalCommissions3 {
	s.CommissionGeneralDetails = new(TotalCommissions3)
	return s.CommissionGeneralDetails
}

func (s *SwitchSubscriptionLegExecution3) AddTaxGeneralDetails() *TotalTaxes3 {
	s.TaxGeneralDetails = new(TotalTaxes3)
	return s.TaxGeneralDetails
}

func (s *SwitchSubscriptionLegExecution3) AddSettlementAndCustodyDetails() *FundSettlementParameters3 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters3)
	return s.SettlementAndCustodyDetails
}

func (s *SwitchSubscriptionLegExecution3) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SwitchSubscriptionLegExecution3) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	s.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return s.PhysicalDeliveryDetails
}

func (s *SwitchSubscriptionLegExecution3) SetNonStandardSettlementInformation(value string) {
	s.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (s *SwitchSubscriptionLegExecution3) AddEqualisation() *Equalisation1 {
	s.Equalisation = new(Equalisation1)
	return s.Equalisation
}

// Execution of the subscription part, in a switch between investment funds or investment fund classes.
type SwitchSubscriptionLegExecution4 struct {

	// Unique technical identifier for the instance of the leg within a switch.
	LegIdentification *Max35Text `xml:"LegId,omitempty"`

	// Unique identifier for the instance of the leg execution within a switch confirmation.
	LegExecutionIdentification *Max35Text `xml:"LegExctnId,omitempty"`

	// Investment fund class to which the subscription leg of the investment fund order execution is related.
	FinancialInstrumentDetails *FinancialInstrument57 `xml:"FinInstrmDtls"`

	// Number of investment fund units subscribed.
	UnitsNumber *DecimalNumber `xml:"UnitsNb"`

	// Amount of money invested in the fund.
	// Net Amount = Quantity * Price.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Amount of money to be paid by the investor when subscribing to fund units.
	// Gross amount = (Quantity * Price) + (Fees + Taxes).
	GrossAmount *ActiveCurrencyAndAmount `xml:"GrssAmt,omitempty"`

	// Account impacted by the investment fund order execution.
	InvestmentAccountDetails *InvestmentAccount58 `xml:"InvstmtAcctDtls,omitempty"`

	// Date and time at which a price is applied, according to the terms stated in the prospectus.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Price at which the order was executed.
	PriceDetails *UnitPrice22 `xml:"PricDtls"`

	// Other quoted price than the one at which the order was executed.
	InformativePriceDetails []*UnitPrice22 `xml:"InftvPricDtls,omitempty"`

	// Indicates whether the dividend is included, that is, cum-dividend, in the executed price. When the dividend is not included, the price will be ex-dividend.
	CumDividendIndicator *YesNoIndicator `xml:"CumDvddInd"`

	// Part of the price deemed as accrued income or profit rather than capital. The interim profit amount is used for tax purposes.
	InterimProfitAmount *ProfitAndLoss2Choice `xml:"IntrmPrftAmt,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *ActiveCurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Fees (charges/commission) and taxes that are taken into consideration for the transaction, so that the total difference between the net amount and gross amount is known, without taking into account equalisation.
	TransactionOverhead *TotalFeesAndTaxes40 `xml:"TxOvrhd,omitempty"`

	// Additional information about tax that does not have an impact on the transaction overhead.
	InformativeTaxDetails *InformativeTax1 `xml:"InftvTaxDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters12 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to the physical delivery of the securities.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Part of an investor's subscription amount that is held by the fund in order to pay incentive/performance fees at the end of the fiscal year.
	Equalisation *Equalisation1 `xml:"Equlstn,omitempty"`
}

func (s *SwitchSubscriptionLegExecution4) SetLegIdentification(value string) {
	s.LegIdentification = (*Max35Text)(&value)
}

func (s *SwitchSubscriptionLegExecution4) SetLegExecutionIdentification(value string) {
	s.LegExecutionIdentification = (*Max35Text)(&value)
}

func (s *SwitchSubscriptionLegExecution4) AddFinancialInstrumentDetails() *FinancialInstrument57 {
	s.FinancialInstrumentDetails = new(FinancialInstrument57)
	return s.FinancialInstrumentDetails
}

func (s *SwitchSubscriptionLegExecution4) SetUnitsNumber(value string) {
	s.UnitsNumber = (*DecimalNumber)(&value)
}

func (s *SwitchSubscriptionLegExecution4) SetNetAmount(value, currency string) {
	s.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchSubscriptionLegExecution4) SetGrossAmount(value, currency string) {
	s.GrossAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchSubscriptionLegExecution4) AddInvestmentAccountDetails() *InvestmentAccount58 {
	s.InvestmentAccountDetails = new(InvestmentAccount58)
	return s.InvestmentAccountDetails
}

func (s *SwitchSubscriptionLegExecution4) AddTradeDateTime() *DateAndDateTimeChoice {
	s.TradeDateTime = new(DateAndDateTimeChoice)
	return s.TradeDateTime
}

func (s *SwitchSubscriptionLegExecution4) AddPriceDetails() *UnitPrice22 {
	s.PriceDetails = new(UnitPrice22)
	return s.PriceDetails
}

func (s *SwitchSubscriptionLegExecution4) AddInformativePriceDetails() *UnitPrice22 {
	newValue := new(UnitPrice22)
	s.InformativePriceDetails = append(s.InformativePriceDetails, newValue)
	return newValue
}

func (s *SwitchSubscriptionLegExecution4) SetCumDividendIndicator(value string) {
	s.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (s *SwitchSubscriptionLegExecution4) AddInterimProfitAmount() *ProfitAndLoss2Choice {
	s.InterimProfitAmount = new(ProfitAndLoss2Choice)
	return s.InterimProfitAmount
}

func (s *SwitchSubscriptionLegExecution4) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SwitchSubscriptionLegExecution4) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (s *SwitchSubscriptionLegExecution4) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (s *SwitchSubscriptionLegExecution4) AddTransactionOverhead() *TotalFeesAndTaxes40 {
	s.TransactionOverhead = new(TotalFeesAndTaxes40)
	return s.TransactionOverhead
}

func (s *SwitchSubscriptionLegExecution4) AddInformativeTaxDetails() *InformativeTax1 {
	s.InformativeTaxDetails = new(InformativeTax1)
	return s.InformativeTaxDetails
}

func (s *SwitchSubscriptionLegExecution4) AddSettlementAndCustodyDetails() *FundSettlementParameters12 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters12)
	return s.SettlementAndCustodyDetails
}

func (s *SwitchSubscriptionLegExecution4) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SwitchSubscriptionLegExecution4) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	s.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return s.PhysicalDeliveryDetails
}

func (s *SwitchSubscriptionLegExecution4) SetNonStandardSettlementInformation(value string) {
	s.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (s *SwitchSubscriptionLegExecution4) AddEqualisation() *Equalisation1 {
	s.Equalisation = new(Equalisation1)
	return s.Equalisation
}
