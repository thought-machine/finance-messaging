package iso20022

// Execution of the redemption part, in a switch between investment funds or investment fund classes.
type SwitchRedemptionLegExecution2 struct {

	// Unique technical identifier for an instance of a leg within a switch.
	LegIdentification *Max35Text `xml:"LegId,omitempty"`

	// Unique identifier for an instance of a leg execution within a switch confirmation.
	LegExecutionIdentification *Max35Text `xml:"LegExctnId,omitempty"`

	// Investment fund class to which an investment fund order execution is related.
	FinancialInstrumentDetails *FinancialInstrument6 `xml:"FinInstrmDtls"`

	// Number of investment funds units redeemed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Portion of the investor's holdings, in a specific investment fund/ fund class, that is redeemed.
	HoldingsRedemptionRate *PercentageRate `xml:"HldgsRedRate,omitempty"`

	// Net amount of money paid to the investor as a result of the redemption.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt"`

	// Gross amount of money as a result of the redemption.
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

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnitCode `xml:"Grp1Or2Units,omitempty"`

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
	SettlementAndCustodyDetails *FundSettlementParameters4 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to physical delivery of the securities.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`
}

func (s *SwitchRedemptionLegExecution2) SetLegIdentification(value string) {
	s.LegIdentification = (*Max35Text)(&value)
}

func (s *SwitchRedemptionLegExecution2) SetLegExecutionIdentification(value string) {
	s.LegExecutionIdentification = (*Max35Text)(&value)
}

func (s *SwitchRedemptionLegExecution2) AddFinancialInstrumentDetails() *FinancialInstrument6 {
	s.FinancialInstrumentDetails = new(FinancialInstrument6)
	return s.FinancialInstrumentDetails
}

func (s *SwitchRedemptionLegExecution2) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	s.UnitsNumber = new(FinancialInstrumentQuantity1)
	return s.UnitsNumber
}

func (s *SwitchRedemptionLegExecution2) SetHoldingsRedemptionRate(value string) {
	s.HoldingsRedemptionRate = (*PercentageRate)(&value)
}

func (s *SwitchRedemptionLegExecution2) SetNetAmount(value, currency string) {
	s.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchRedemptionLegExecution2) SetGrossAmount(value, currency string) {
	s.GrossAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchRedemptionLegExecution2) AddTradeDateTime() *DateAndDateTimeChoice {
	s.TradeDateTime = new(DateAndDateTimeChoice)
	return s.TradeDateTime
}

func (s *SwitchRedemptionLegExecution2) AddPriceDetails() *UnitPrice5 {
	s.PriceDetails = new(UnitPrice5)
	return s.PriceDetails
}

func (s *SwitchRedemptionLegExecution2) SetCumDividendIndicator(value string) {
	s.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (s *SwitchRedemptionLegExecution2) AddInterimProfitAmount() *ProfitAndLoss1Choice {
	s.InterimProfitAmount = new(ProfitAndLoss1Choice)
	return s.InterimProfitAmount
}

func (s *SwitchRedemptionLegExecution2) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SwitchRedemptionLegExecution2) SetGroup1Or2Units(value string) {
	s.Group1Or2Units = (*UKTaxGroupUnitCode)(&value)
}

func (s *SwitchRedemptionLegExecution2) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*CurrencyCode)(&value)
}

func (s *SwitchRedemptionLegExecution2) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*CurrencyCode)(&value)
}

func (s *SwitchRedemptionLegExecution2) AddChargeGeneralDetails() *TotalCharges2 {
	s.ChargeGeneralDetails = new(TotalCharges2)
	return s.ChargeGeneralDetails
}

func (s *SwitchRedemptionLegExecution2) AddCommissionGeneralDetails() *TotalCommissions2 {
	s.CommissionGeneralDetails = new(TotalCommissions2)
	return s.CommissionGeneralDetails
}

func (s *SwitchRedemptionLegExecution2) AddTaxGeneralDetails() *TotalTaxes2 {
	s.TaxGeneralDetails = new(TotalTaxes2)
	return s.TaxGeneralDetails
}

func (s *SwitchRedemptionLegExecution2) AddSettlementAndCustodyDetails() *FundSettlementParameters4 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters4)
	return s.SettlementAndCustodyDetails
}

func (s *SwitchRedemptionLegExecution2) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SwitchRedemptionLegExecution2) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	s.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return s.PhysicalDeliveryDetails
}

// Execution of the redemption part, in a switch between investment funds or investment fund classes.
type SwitchRedemptionLegExecution3 struct {

	// Unique technical identifier for an instance of a leg within a switch.
	LegIdentification *Max35Text `xml:"LegId,omitempty"`

	// Unique identifier for an instance of a leg execution within a switch confirmation.
	LegExecutionIdentification *Max35Text `xml:"LegExctnId,omitempty"`

	// Investment fund class to which an investment fund order execution is related.
	FinancialInstrumentDetails *FinancialInstrument10 `xml:"FinInstrmDtls"`

	// Number of investment funds units redeemed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
	InvestmentAccountDetails *InvestmentAccount21 `xml:"InvstmtAcctDtls,omitempty"`

	// Portion of the investor's holdings, in a specific investment fund/ fund class, that is redeemed.
	HoldingsRedemptionRate *PercentageRate `xml:"HldgsRedRate,omitempty"`

	// Amount of money paid to the investor as a result of the redemption after deduction of charges, commissions and taxes.
	// [(Quantity * Price) - (Charges + Commissions +Taxes)]
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Amount of money resulting from the redemption before deduction of  charges, commissions and taxes.
	// [Quantity * Price]
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

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnitCode `xml:"Grp1Or2Units,omitempty"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *ActiveCurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Amount of money associated with a service.
	ChargeGeneralDetails *TotalCharges3 `xml:"ChrgGnlDtls,omitempty"`

	// Commission for the execution of an investment fund order.
	CommissionGeneralDetails *TotalCommissions3 `xml:"ComssnGnlDtls,omitempty"`

	// Tax applicable to execution of an investment fund order.
	TaxGeneralDetails *TotalTaxes3 `xml:"TaxGnlDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters4 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to physical delivery of the securities.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Part of an investor's subscription amount that is held by the fund in order to pay incentive / performance fees at the end of the fiscal year.
	Equalisation *Equalisation1 `xml:"Equlstn,omitempty"`
}

func (s *SwitchRedemptionLegExecution3) SetLegIdentification(value string) {
	s.LegIdentification = (*Max35Text)(&value)
}

func (s *SwitchRedemptionLegExecution3) SetLegExecutionIdentification(value string) {
	s.LegExecutionIdentification = (*Max35Text)(&value)
}

func (s *SwitchRedemptionLegExecution3) AddFinancialInstrumentDetails() *FinancialInstrument10 {
	s.FinancialInstrumentDetails = new(FinancialInstrument10)
	return s.FinancialInstrumentDetails
}

func (s *SwitchRedemptionLegExecution3) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	s.UnitsNumber = new(FinancialInstrumentQuantity1)
	return s.UnitsNumber
}

func (s *SwitchRedemptionLegExecution3) AddInvestmentAccountDetails() *InvestmentAccount21 {
	s.InvestmentAccountDetails = new(InvestmentAccount21)
	return s.InvestmentAccountDetails
}

func (s *SwitchRedemptionLegExecution3) SetHoldingsRedemptionRate(value string) {
	s.HoldingsRedemptionRate = (*PercentageRate)(&value)
}

func (s *SwitchRedemptionLegExecution3) SetNetAmount(value, currency string) {
	s.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchRedemptionLegExecution3) SetGrossAmount(value, currency string) {
	s.GrossAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchRedemptionLegExecution3) AddTradeDateTime() *DateAndDateTimeChoice {
	s.TradeDateTime = new(DateAndDateTimeChoice)
	return s.TradeDateTime
}

func (s *SwitchRedemptionLegExecution3) AddPriceDetails() *UnitPrice10 {
	s.PriceDetails = new(UnitPrice10)
	return s.PriceDetails
}

func (s *SwitchRedemptionLegExecution3) SetCumDividendIndicator(value string) {
	s.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (s *SwitchRedemptionLegExecution3) AddInterimProfitAmount() *ProfitAndLoss1Choice {
	s.InterimProfitAmount = new(ProfitAndLoss1Choice)
	return s.InterimProfitAmount
}

func (s *SwitchRedemptionLegExecution3) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SwitchRedemptionLegExecution3) SetGroup1Or2Units(value string) {
	s.Group1Or2Units = (*UKTaxGroupUnitCode)(&value)
}

func (s *SwitchRedemptionLegExecution3) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (s *SwitchRedemptionLegExecution3) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (s *SwitchRedemptionLegExecution3) AddChargeGeneralDetails() *TotalCharges3 {
	s.ChargeGeneralDetails = new(TotalCharges3)
	return s.ChargeGeneralDetails
}

func (s *SwitchRedemptionLegExecution3) AddCommissionGeneralDetails() *TotalCommissions3 {
	s.CommissionGeneralDetails = new(TotalCommissions3)
	return s.CommissionGeneralDetails
}

func (s *SwitchRedemptionLegExecution3) AddTaxGeneralDetails() *TotalTaxes3 {
	s.TaxGeneralDetails = new(TotalTaxes3)
	return s.TaxGeneralDetails
}

func (s *SwitchRedemptionLegExecution3) AddSettlementAndCustodyDetails() *FundSettlementParameters4 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters4)
	return s.SettlementAndCustodyDetails
}

func (s *SwitchRedemptionLegExecution3) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SwitchRedemptionLegExecution3) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	s.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return s.PhysicalDeliveryDetails
}

func (s *SwitchRedemptionLegExecution3) SetNonStandardSettlementInformation(value string) {
	s.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (s *SwitchRedemptionLegExecution3) AddEqualisation() *Equalisation1 {
	s.Equalisation = new(Equalisation1)
	return s.Equalisation
}

// Execution of the redemption part, in a switch between investment funds or investment fund classes.
type SwitchRedemptionLegExecution4 struct {

	// Unique technical identifier for the instance of the leg within a switch.
	LegIdentification *Max35Text `xml:"LegId,omitempty"`

	// Unique identifier for the instance of the leg execution within a switch confirmation.
	LegExecutionIdentification *Max35Text `xml:"LegExctnId,omitempty"`

	// Investment fund class to which the redemption leg of the investment fund order execution is related.
	FinancialInstrumentDetails *FinancialInstrument57 `xml:"FinInstrmDtls"`

	// Number of investment funds units redeemed.
	UnitsNumber *DecimalNumber `xml:"UnitsNb"`

	// Portion of the investor's holdings redeemed.
	HoldingsRedemptionRate *PercentageRate `xml:"HldgsRedRate,omitempty"`

	// Amount of money paid to the investor when redeeming fund units.
	// Net amount = (Quantity * Price) – (Fees + Taxes).
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Amount of money redeemed from the fund.
	// Gross Amount = Quantity * Price.
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

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnit1Code `xml:"Grp1Or2Units,omitempty"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *ActiveCurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Fees (charges/commission) and taxes that are taken into consideration for the transaction, so that the total difference between the net amount and gross amount is known, without taking into account equalisation.
	TransactionOverhead *TotalFeesAndTaxes40 `xml:"TxOvrhd,omitempty"`

	// Additional information about tax that does not have an impact on the transaction overhead.
	InformativeTaxDetails *InformativeTax1 `xml:"InftvTaxDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters11 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to the physical delivery of the securities.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Part of an investor's subscription amount that was held by the fund in order to pay incentive/performance fees at the end of the fiscal year, and is returned due to the redemption.
	Equalisation *Equalisation1 `xml:"Equlstn,omitempty"`

	// Information about gating and hold back of redemption proceeds.
	GatingOrHoldBackDetails *HoldBackInformation2 `xml:"GtgOrHldBckDtls,omitempty"`
}

func (s *SwitchRedemptionLegExecution4) SetLegIdentification(value string) {
	s.LegIdentification = (*Max35Text)(&value)
}

func (s *SwitchRedemptionLegExecution4) SetLegExecutionIdentification(value string) {
	s.LegExecutionIdentification = (*Max35Text)(&value)
}

func (s *SwitchRedemptionLegExecution4) AddFinancialInstrumentDetails() *FinancialInstrument57 {
	s.FinancialInstrumentDetails = new(FinancialInstrument57)
	return s.FinancialInstrumentDetails
}

func (s *SwitchRedemptionLegExecution4) SetUnitsNumber(value string) {
	s.UnitsNumber = (*DecimalNumber)(&value)
}

func (s *SwitchRedemptionLegExecution4) SetHoldingsRedemptionRate(value string) {
	s.HoldingsRedemptionRate = (*PercentageRate)(&value)
}

func (s *SwitchRedemptionLegExecution4) SetNetAmount(value, currency string) {
	s.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchRedemptionLegExecution4) SetGrossAmount(value, currency string) {
	s.GrossAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchRedemptionLegExecution4) AddInvestmentAccountDetails() *InvestmentAccount58 {
	s.InvestmentAccountDetails = new(InvestmentAccount58)
	return s.InvestmentAccountDetails
}

func (s *SwitchRedemptionLegExecution4) AddTradeDateTime() *DateAndDateTimeChoice {
	s.TradeDateTime = new(DateAndDateTimeChoice)
	return s.TradeDateTime
}

func (s *SwitchRedemptionLegExecution4) AddPriceDetails() *UnitPrice22 {
	s.PriceDetails = new(UnitPrice22)
	return s.PriceDetails
}

func (s *SwitchRedemptionLegExecution4) AddInformativePriceDetails() *UnitPrice22 {
	newValue := new(UnitPrice22)
	s.InformativePriceDetails = append(s.InformativePriceDetails, newValue)
	return newValue
}

func (s *SwitchRedemptionLegExecution4) SetCumDividendIndicator(value string) {
	s.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (s *SwitchRedemptionLegExecution4) AddInterimProfitAmount() *ProfitAndLoss2Choice {
	s.InterimProfitAmount = new(ProfitAndLoss2Choice)
	return s.InterimProfitAmount
}

func (s *SwitchRedemptionLegExecution4) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SwitchRedemptionLegExecution4) SetGroup1Or2Units(value string) {
	s.Group1Or2Units = (*UKTaxGroupUnit1Code)(&value)
}

func (s *SwitchRedemptionLegExecution4) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (s *SwitchRedemptionLegExecution4) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (s *SwitchRedemptionLegExecution4) AddTransactionOverhead() *TotalFeesAndTaxes40 {
	s.TransactionOverhead = new(TotalFeesAndTaxes40)
	return s.TransactionOverhead
}

func (s *SwitchRedemptionLegExecution4) AddInformativeTaxDetails() *InformativeTax1 {
	s.InformativeTaxDetails = new(InformativeTax1)
	return s.InformativeTaxDetails
}

func (s *SwitchRedemptionLegExecution4) AddSettlementAndCustodyDetails() *FundSettlementParameters11 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters11)
	return s.SettlementAndCustodyDetails
}

func (s *SwitchRedemptionLegExecution4) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SwitchRedemptionLegExecution4) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	s.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return s.PhysicalDeliveryDetails
}

func (s *SwitchRedemptionLegExecution4) SetNonStandardSettlementInformation(value string) {
	s.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (s *SwitchRedemptionLegExecution4) AddEqualisation() *Equalisation1 {
	s.Equalisation = new(Equalisation1)
	return s.Equalisation
}

func (s *SwitchRedemptionLegExecution4) AddGatingOrHoldBackDetails() *HoldBackInformation2 {
	s.GatingOrHoldBackDetails = new(HoldBackInformation2)
	return s.GatingOrHoldBackDetails
}
