package iso20022

// Execution of a redemption order.
type RedemptionExecution15 struct {

	// Unique and unambiguous identifier for the order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of the order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for the order execution, as assigned by the confirming party.
	DealReference *Max35Text `xml:"DealRef"`

	// Category of the investment fund order.
	OrderType []*FundOrderType4Choice `xml:"OrdrTp,omitempty"`

	// Investment fund class to which the investment fund order execution is related.
	FinancialInstrumentDetails *FinancialInstrument57 `xml:"FinInstrmDtls"`

	// Subdivision of the account used to segregate specific holdings.
	SubAccountForHolding *SubAccount6 `xml:"SubAcctForHldg,omitempty"`

	// Number of investment funds units redeemed.
	UnitsNumber *DecimalNumber `xml:"UnitsNb"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Amount of money paid to the investor when redeeming fund units.
	// Net amount = (Quantity * Price) – (Fees + Taxes).
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Amount of money redeemed from the fund.
	// Gross Amount = Quantity * Price.
	GrossAmount *ActiveCurrencyAndAmount `xml:"GrssAmt,omitempty"`

	// Portion of the investor's holdings redeemed.
	//
	HoldingsRedemptionRate *PercentageRate `xml:"HldgsRedRate,omitempty"`

	// Date and time at which a price is applied, according to the terms stated in the prospectus.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Price at which the order was executed.
	DealingPriceDetails *UnitPrice22 `xml:"DealgPricDtls"`

	// Other quoted price than the one at which the order was executed.
	InformativePriceDetails []*UnitPrice22 `xml:"InftvPricDtls,omitempty"`

	// Total amount of money paid/to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Method by which the transaction is settled.
	SettlementMethod *DeliveryReceiptType2Code `xml:"SttlmMtd,omitempty"`

	// Indicates whether the order has been partially executed, that is, the confirmed quantity does not match the ordered quantity for a given financial instrument.
	PartiallyExecutedIndicator *YesNoIndicator `xml:"PrtlyExctdInd"`

	// Specifies that the execution was subject to best execution rules as defined by MiFID.
	BestExecution *BestExecution1Code `xml:"BestExctn,omitempty"`

	// Indicates whether the dividend is included, that is, cum-dividend, in the executed price. When the dividend is not included, the price will be ex-dividend.
	CumDividendIndicator *YesNoIndicator `xml:"CumDvddInd"`

	// Part of the price deemed as accrued income or profit rather than capital. The interim profit amount is used for tax purposes.
	InterimProfitAmount *ProfitAndLoss2Choice `xml:"IntrmPrftAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	// How the exchange rate is expressed determines which currency is the Unit Currency and Quoted Currency. If the amounts concerned are EUR 1000 and USD 1300, the exchange rate may be expressed as per either of the following examples:
	// EXAMPLE 1
	// UnitCurrency  EUR
	// QuotedCurrency  USD
	// ExchangeRate  1.300
	// EXAMPLE 2
	// UnitCurrency  USD
	// QuotedCurrency  EUR
	// ExchangeRate  0.769
	ForeignExchangeDetails []*ForeignExchangeTerms33 `xml:"FXDtls,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnit1Code `xml:"Grp1Or2Units,omitempty"`

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

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *ActiveCurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Payment process for the transfer of cash from the debtor to the creditor.
	CashSettlementDetails *PaymentTransaction72 `xml:"CshSttlmDtls,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Percentage of units partially settled.
	PartialSettlementOfUnits *PercentageRate `xml:"PrtlSttlmOfUnits,omitempty"`

	// Percentage of cash partially settled.
	PartialSettlementOfCash *PercentageRate `xml:"PrtlSttlmOfCsh,omitempty"`

	// Breakdown of the net amount per type of order.
	StaffClientBreakdown []*InvestmentFundsOrderBreakdown2 `xml:"StffClntBrkdwn,omitempty"`

	// Amount retained by the fund and paid out later at a time decided by the fund.
	PartialRedemptionWithholdingAmount *ActiveCurrencyAndAmount `xml:"PrtlRedWhldgAmt,omitempty"`

	// Specifies if advice has been received from an independent financial advisor.
	FinancialAdvice *FinancialAdvice1Code `xml:"FinAdvc,omitempty"`

	// Specifies whether the trade is negotiated.
	NegotiatedTrade *NegotiatedTrade1Code `xml:"NgtdTrad,omitempty"`

	// Specifies whether the order execution confirmation is late.
	LateReport *LateReport1Code `xml:"LateRpt,omitempty"`

	// Party related to the transaction.
	RelatedPartyDetails []*Intermediary39 `xml:"RltdPtyDtls,omitempty"`

	// Part of an investor's subscription amount that was held by the fund in order to pay incentive/performance fees at the end of the fiscal year, and is returned due to the redemption.
	Equalisation *Equalisation1 `xml:"Equlstn,omitempty"`

	// Assessment of the customer’s behaviour at the time of the account opening application.
	CustomerConductClassification *CustomerConductClassification1Choice `xml:"CstmrCndctClssfctn,omitempty"`

	// Means by which the investor or account owner submits the open account form.
	TransactionChannelType *TransactionChannelType1Choice `xml:"TxChanlTp,omitempty"`

	// Type of signature.
	SignatureType *SignatureType1Choice `xml:"SgntrTp,omitempty"`

	// Information about a non-standard order.
	OrderWaiverDetails *OrderWaiver1 `xml:"OrdrWvrDtls,omitempty"`

	// Information about gating and hold back of redemption proceeds.
	GatingOrHoldBackDetails *HoldBackInformation2 `xml:"GtgOrHldBckDtls,omitempty"`
}

func (r *RedemptionExecution15) SetOrderReference(value string) {
	r.OrderReference = (*Max35Text)(&value)
}

func (r *RedemptionExecution15) SetClientReference(value string) {
	r.ClientReference = (*Max35Text)(&value)
}

func (r *RedemptionExecution15) SetDealReference(value string) {
	r.DealReference = (*Max35Text)(&value)
}

func (r *RedemptionExecution15) AddOrderType() *FundOrderType4Choice {
	newValue := new(FundOrderType4Choice)
	r.OrderType = append(r.OrderType, newValue)
	return newValue
}

func (r *RedemptionExecution15) AddFinancialInstrumentDetails() *FinancialInstrument57 {
	r.FinancialInstrumentDetails = new(FinancialInstrument57)
	return r.FinancialInstrumentDetails
}

func (r *RedemptionExecution15) AddSubAccountForHolding() *SubAccount6 {
	r.SubAccountForHolding = new(SubAccount6)
	return r.SubAccountForHolding
}

func (r *RedemptionExecution15) SetUnitsNumber(value string) {
	r.UnitsNumber = (*DecimalNumber)(&value)
}

func (r *RedemptionExecution15) SetRounding(value string) {
	r.Rounding = (*RoundingDirection2Code)(&value)
}

func (r *RedemptionExecution15) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution15) SetGrossAmount(value, currency string) {
	r.GrossAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution15) SetHoldingsRedemptionRate(value string) {
	r.HoldingsRedemptionRate = (*PercentageRate)(&value)
}

func (r *RedemptionExecution15) AddTradeDateTime() *DateAndDateTimeChoice {
	r.TradeDateTime = new(DateAndDateTimeChoice)
	return r.TradeDateTime
}

func (r *RedemptionExecution15) AddDealingPriceDetails() *UnitPrice22 {
	r.DealingPriceDetails = new(UnitPrice22)
	return r.DealingPriceDetails
}

func (r *RedemptionExecution15) AddInformativePriceDetails() *UnitPrice22 {
	newValue := new(UnitPrice22)
	r.InformativePriceDetails = append(r.InformativePriceDetails, newValue)
	return newValue
}

func (r *RedemptionExecution15) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution15) SetCashSettlementDate(value string) {
	r.CashSettlementDate = (*ISODate)(&value)
}

func (r *RedemptionExecution15) SetSettlementMethod(value string) {
	r.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (r *RedemptionExecution15) SetPartiallyExecutedIndicator(value string) {
	r.PartiallyExecutedIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionExecution15) SetBestExecution(value string) {
	r.BestExecution = (*BestExecution1Code)(&value)
}

func (r *RedemptionExecution15) SetCumDividendIndicator(value string) {
	r.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionExecution15) AddInterimProfitAmount() *ProfitAndLoss2Choice {
	r.InterimProfitAmount = new(ProfitAndLoss2Choice)
	return r.InterimProfitAmount
}

func (r *RedemptionExecution15) AddForeignExchangeDetails() *ForeignExchangeTerms33 {
	newValue := new(ForeignExchangeTerms33)
	r.ForeignExchangeDetails = append(r.ForeignExchangeDetails, newValue)
	return newValue
}

func (r *RedemptionExecution15) SetIncomePreference(value string) {
	r.IncomePreference = (*IncomePreference1Code)(&value)
}

func (r *RedemptionExecution15) SetGroup1Or2Units(value string) {
	r.Group1Or2Units = (*UKTaxGroupUnit1Code)(&value)
}

func (r *RedemptionExecution15) AddTransactionOverhead() *TotalFeesAndTaxes40 {
	r.TransactionOverhead = new(TotalFeesAndTaxes40)
	return r.TransactionOverhead
}

func (r *RedemptionExecution15) AddInformativeTaxDetails() *InformativeTax1 {
	r.InformativeTaxDetails = new(InformativeTax1)
	return r.InformativeTaxDetails
}

func (r *RedemptionExecution15) AddSettlementAndCustodyDetails() *FundSettlementParameters11 {
	r.SettlementAndCustodyDetails = new(FundSettlementParameters11)
	return r.SettlementAndCustodyDetails
}

func (r *RedemptionExecution15) SetPhysicalDeliveryIndicator(value string) {
	r.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionExecution15) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	r.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return r.PhysicalDeliveryDetails
}

func (r *RedemptionExecution15) SetRequestedSettlementCurrency(value string) {
	r.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (r *RedemptionExecution15) SetRequestedNAVCurrency(value string) {
	r.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (r *RedemptionExecution15) AddCashSettlementDetails() *PaymentTransaction72 {
	r.CashSettlementDetails = new(PaymentTransaction72)
	return r.CashSettlementDetails
}

func (r *RedemptionExecution15) SetNonStandardSettlementInformation(value string) {
	r.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (r *RedemptionExecution15) SetPartialSettlementOfUnits(value string) {
	r.PartialSettlementOfUnits = (*PercentageRate)(&value)
}

func (r *RedemptionExecution15) SetPartialSettlementOfCash(value string) {
	r.PartialSettlementOfCash = (*PercentageRate)(&value)
}

func (r *RedemptionExecution15) AddStaffClientBreakdown() *InvestmentFundsOrderBreakdown2 {
	newValue := new(InvestmentFundsOrderBreakdown2)
	r.StaffClientBreakdown = append(r.StaffClientBreakdown, newValue)
	return newValue
}

func (r *RedemptionExecution15) SetPartialRedemptionWithholdingAmount(value, currency string) {
	r.PartialRedemptionWithholdingAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution15) SetFinancialAdvice(value string) {
	r.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (r *RedemptionExecution15) SetNegotiatedTrade(value string) {
	r.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (r *RedemptionExecution15) SetLateReport(value string) {
	r.LateReport = (*LateReport1Code)(&value)
}

func (r *RedemptionExecution15) AddRelatedPartyDetails() *Intermediary39 {
	newValue := new(Intermediary39)
	r.RelatedPartyDetails = append(r.RelatedPartyDetails, newValue)
	return newValue
}

func (r *RedemptionExecution15) AddEqualisation() *Equalisation1 {
	r.Equalisation = new(Equalisation1)
	return r.Equalisation
}

func (r *RedemptionExecution15) AddCustomerConductClassification() *CustomerConductClassification1Choice {
	r.CustomerConductClassification = new(CustomerConductClassification1Choice)
	return r.CustomerConductClassification
}

func (r *RedemptionExecution15) AddTransactionChannelType() *TransactionChannelType1Choice {
	r.TransactionChannelType = new(TransactionChannelType1Choice)
	return r.TransactionChannelType
}

func (r *RedemptionExecution15) AddSignatureType() *SignatureType1Choice {
	r.SignatureType = new(SignatureType1Choice)
	return r.SignatureType
}

func (r *RedemptionExecution15) AddOrderWaiverDetails() *OrderWaiver1 {
	r.OrderWaiverDetails = new(OrderWaiver1)
	return r.OrderWaiverDetails
}

func (r *RedemptionExecution15) AddGatingOrHoldBackDetails() *HoldBackInformation2 {
	r.GatingOrHoldBackDetails = new(HoldBackInformation2)
	return r.GatingOrHoldBackDetails
}

// Execution of a redemption order.
type RedemptionExecution16 struct {

	// Unique and unambiguous identifier for the order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of the order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for the order execution, as assigned by the confirming party.
	DealReference *Max35Text `xml:"DealRef"`

	// Category of the investment fund order.
	OrderType []*FundOrderType4Choice `xml:"OrdrTp,omitempty"`

	// Additional information about the investor.
	BeneficiaryDetails []*IndividualPerson32 `xml:"BnfcryDtls,omitempty"`

	// Number of investment funds units redeemed.
	UnitsNumber *DecimalNumber `xml:"UnitsNb"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Amount of money paid to the investor when redeeming fund units.
	// Net amount = (Quantity * Price) – (Fees + Taxes).
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Account impacted by the investment fund order execution.
	InvestmentAccountDetails *InvestmentAccount58 `xml:"InvstmtAcctDtls"`

	// Portion of the investor's holdings redeemed.
	HoldingsRedemptionRate *PercentageRate `xml:"HldgsRedRate,omitempty"`

	// Amount of money redeemed from the fund.
	// Gross Amount = Quantity * Price.
	GrossAmount *ActiveCurrencyAndAmount `xml:"GrssAmt,omitempty"`

	// Date and time at which a price is applied, according to the terms stated in the prospectus.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Method by which the transaction is settled.
	SettlementMethod *DeliveryReceiptType2Code `xml:"SttlmMtd,omitempty"`

	// Price at which the order was executed.
	DealingPriceDetails *UnitPrice22 `xml:"DealgPricDtls"`

	// Other quoted price than the one at which the order was executed.
	InformativePriceDetails []*UnitPrice22 `xml:"InftvPricDtls,omitempty"`

	// Indicates whether the order has been partially executed, that is, the confirmed quantity does not match the ordered quantity for a given financial instrument.
	PartiallyExecutedIndicator *YesNoIndicator `xml:"PrtlyExctdInd"`

	// Specifies that the execution was subject to best execution rules as defined by MiFID.
	BestExecution *BestExecution1Code `xml:"BestExctn,omitempty"`

	// Indicates whether the dividend is included, that is, cum-dividend, in the executed price. When the dividend is not included, the price will be ex-dividend.
	CumDividendIndicator *YesNoIndicator `xml:"CumDvddInd"`

	// Part of the price deemed as accrued income or profit rather than capital. The interim profit amount is used for tax purposes.
	InterimProfitAmount *ProfitAndLoss2Choice `xml:"IntrmPrftAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	// How the exchange rate is expressed determines which currency is the Unit Currency and Quoted Currency. If the amounts concerned are EUR 1000 and USD 1300, the exchange rate may be expressed as per either of the following examples:
	// EXAMPLE 1
	// UnitCurrency  EUR
	// QuotedCurrency  USD
	// ExchangeRate  1.300
	// EXAMPLE 2
	// UnitCurrency  USD
	// QuotedCurrency  EUR
	// ExchangeRate  0.769
	ForeignExchangeDetails []*ForeignExchangeTerms33 `xml:"FXDtls,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnit1Code `xml:"Grp1Or2Units,omitempty"`

	// Fees (charges/commission) and taxes that are taken into consideration for the transaction, so that the total difference between the net amount and gross amount is known, without taking into account equalisation.
	TransactionOverhead *TotalFeesAndTaxes40 `xml:"TxOvrhd,omitempty"`

	// Additional information about tax that does not have an impact on the transaction overhead.
	InformativeTaxDetails *InformativeTax1 `xml:"InftvTaxDtls,omitempty"`

	// Breakdown of the net amount per type of order.
	StaffClientBreakdown []*InvestmentFundsOrderBreakdown2 `xml:"StffClntBrkdwn,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters11 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to the physical delivery of the securities.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`

	// Payment process for the transfer of cash from the debtor to the creditor.
	CashSettlementDetails *PaymentTransaction72 `xml:"CshSttlmDtls,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Percentage of units partially settled.
	PartialSettlementOfUnits *PercentageRate `xml:"PrtlSttlmOfUnits,omitempty"`

	// Percentage of cash partially settled.
	PartialSettlementOfCash *PercentageRate `xml:"PrtlSttlmOfCsh,omitempty"`

	// Amount retained by the fund and paid out later at a time decided by the fund.
	PartialRedemptionWithholdingAmount *ActiveCurrencyAndAmount `xml:"PrtlRedWhldgAmt,omitempty"`

	// Specifies if advice has been received from an independent financial advisor.
	FinancialAdvice *FinancialAdvice1Code `xml:"FinAdvc,omitempty"`

	// Specifies whether the trade is negotiated.
	NegotiatedTrade *NegotiatedTrade1Code `xml:"NgtdTrad,omitempty"`

	// Specifies whether the order execution confirmation is late.
	LateReport *LateReport1Code `xml:"LateRpt,omitempty"`

	// Party related to the transaction.
	RelatedPartyDetails []*Intermediary39 `xml:"RltdPtyDtls,omitempty"`

	// Part of an investor's subscription amount that was held by the fund in order to pay incentive/performance fees at the end of the fiscal year, and is returned due to the redemption.
	Equalisation *Equalisation1 `xml:"Equlstn,omitempty"`

	// Assessment of the customer’s behaviour at the time of the account opening application.
	CustomerConductClassification *CustomerConductClassification1Choice `xml:"CstmrCndctClssfctn,omitempty"`

	// Means by which the investor or account owner submits the open account form.
	TransactionChannelType *TransactionChannelType1Choice `xml:"TxChanlTp,omitempty"`

	// Type of signature.
	SignatureType *SignatureType1Choice `xml:"SgntrTp,omitempty"`

	// Information about a non-standard order.
	OrderWaiverDetails *OrderWaiver1 `xml:"OrdrWvrDtls,omitempty"`

	// Information about gating and hold back of redemption proceeds.
	GatingOrHoldBackDetails *HoldBackInformation2 `xml:"GtgOrHldBckDtls,omitempty"`
}

func (r *RedemptionExecution16) SetOrderReference(value string) {
	r.OrderReference = (*Max35Text)(&value)
}

func (r *RedemptionExecution16) SetClientReference(value string) {
	r.ClientReference = (*Max35Text)(&value)
}

func (r *RedemptionExecution16) SetDealReference(value string) {
	r.DealReference = (*Max35Text)(&value)
}

func (r *RedemptionExecution16) AddOrderType() *FundOrderType4Choice {
	newValue := new(FundOrderType4Choice)
	r.OrderType = append(r.OrderType, newValue)
	return newValue
}

func (r *RedemptionExecution16) AddBeneficiaryDetails() *IndividualPerson32 {
	newValue := new(IndividualPerson32)
	r.BeneficiaryDetails = append(r.BeneficiaryDetails, newValue)
	return newValue
}

func (r *RedemptionExecution16) SetUnitsNumber(value string) {
	r.UnitsNumber = (*DecimalNumber)(&value)
}

func (r *RedemptionExecution16) SetRounding(value string) {
	r.Rounding = (*RoundingDirection2Code)(&value)
}

func (r *RedemptionExecution16) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution16) AddInvestmentAccountDetails() *InvestmentAccount58 {
	r.InvestmentAccountDetails = new(InvestmentAccount58)
	return r.InvestmentAccountDetails
}

func (r *RedemptionExecution16) SetHoldingsRedemptionRate(value string) {
	r.HoldingsRedemptionRate = (*PercentageRate)(&value)
}

func (r *RedemptionExecution16) SetGrossAmount(value, currency string) {
	r.GrossAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution16) AddTradeDateTime() *DateAndDateTimeChoice {
	r.TradeDateTime = new(DateAndDateTimeChoice)
	return r.TradeDateTime
}

func (r *RedemptionExecution16) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution16) SetCashSettlementDate(value string) {
	r.CashSettlementDate = (*ISODate)(&value)
}

func (r *RedemptionExecution16) SetSettlementMethod(value string) {
	r.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (r *RedemptionExecution16) AddDealingPriceDetails() *UnitPrice22 {
	r.DealingPriceDetails = new(UnitPrice22)
	return r.DealingPriceDetails
}

func (r *RedemptionExecution16) AddInformativePriceDetails() *UnitPrice22 {
	newValue := new(UnitPrice22)
	r.InformativePriceDetails = append(r.InformativePriceDetails, newValue)
	return newValue
}

func (r *RedemptionExecution16) SetPartiallyExecutedIndicator(value string) {
	r.PartiallyExecutedIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionExecution16) SetBestExecution(value string) {
	r.BestExecution = (*BestExecution1Code)(&value)
}

func (r *RedemptionExecution16) SetCumDividendIndicator(value string) {
	r.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionExecution16) AddInterimProfitAmount() *ProfitAndLoss2Choice {
	r.InterimProfitAmount = new(ProfitAndLoss2Choice)
	return r.InterimProfitAmount
}

func (r *RedemptionExecution16) AddForeignExchangeDetails() *ForeignExchangeTerms33 {
	newValue := new(ForeignExchangeTerms33)
	r.ForeignExchangeDetails = append(r.ForeignExchangeDetails, newValue)
	return newValue
}

func (r *RedemptionExecution16) SetIncomePreference(value string) {
	r.IncomePreference = (*IncomePreference1Code)(&value)
}

func (r *RedemptionExecution16) SetGroup1Or2Units(value string) {
	r.Group1Or2Units = (*UKTaxGroupUnit1Code)(&value)
}

func (r *RedemptionExecution16) AddTransactionOverhead() *TotalFeesAndTaxes40 {
	r.TransactionOverhead = new(TotalFeesAndTaxes40)
	return r.TransactionOverhead
}

func (r *RedemptionExecution16) AddInformativeTaxDetails() *InformativeTax1 {
	r.InformativeTaxDetails = new(InformativeTax1)
	return r.InformativeTaxDetails
}

func (r *RedemptionExecution16) AddStaffClientBreakdown() *InvestmentFundsOrderBreakdown2 {
	newValue := new(InvestmentFundsOrderBreakdown2)
	r.StaffClientBreakdown = append(r.StaffClientBreakdown, newValue)
	return newValue
}

func (r *RedemptionExecution16) AddSettlementAndCustodyDetails() *FundSettlementParameters11 {
	r.SettlementAndCustodyDetails = new(FundSettlementParameters11)
	return r.SettlementAndCustodyDetails
}

func (r *RedemptionExecution16) SetPhysicalDeliveryIndicator(value string) {
	r.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionExecution16) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	r.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return r.PhysicalDeliveryDetails
}

func (r *RedemptionExecution16) AddCashSettlementDetails() *PaymentTransaction72 {
	r.CashSettlementDetails = new(PaymentTransaction72)
	return r.CashSettlementDetails
}

func (r *RedemptionExecution16) SetNonStandardSettlementInformation(value string) {
	r.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (r *RedemptionExecution16) SetPartialSettlementOfUnits(value string) {
	r.PartialSettlementOfUnits = (*PercentageRate)(&value)
}

func (r *RedemptionExecution16) SetPartialSettlementOfCash(value string) {
	r.PartialSettlementOfCash = (*PercentageRate)(&value)
}

func (r *RedemptionExecution16) SetPartialRedemptionWithholdingAmount(value, currency string) {
	r.PartialRedemptionWithholdingAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution16) SetFinancialAdvice(value string) {
	r.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (r *RedemptionExecution16) SetNegotiatedTrade(value string) {
	r.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (r *RedemptionExecution16) SetLateReport(value string) {
	r.LateReport = (*LateReport1Code)(&value)
}

func (r *RedemptionExecution16) AddRelatedPartyDetails() *Intermediary39 {
	newValue := new(Intermediary39)
	r.RelatedPartyDetails = append(r.RelatedPartyDetails, newValue)
	return newValue
}

func (r *RedemptionExecution16) AddEqualisation() *Equalisation1 {
	r.Equalisation = new(Equalisation1)
	return r.Equalisation
}

func (r *RedemptionExecution16) AddCustomerConductClassification() *CustomerConductClassification1Choice {
	r.CustomerConductClassification = new(CustomerConductClassification1Choice)
	return r.CustomerConductClassification
}

func (r *RedemptionExecution16) AddTransactionChannelType() *TransactionChannelType1Choice {
	r.TransactionChannelType = new(TransactionChannelType1Choice)
	return r.TransactionChannelType
}

func (r *RedemptionExecution16) AddSignatureType() *SignatureType1Choice {
	r.SignatureType = new(SignatureType1Choice)
	return r.SignatureType
}

func (r *RedemptionExecution16) AddOrderWaiverDetails() *OrderWaiver1 {
	r.OrderWaiverDetails = new(OrderWaiver1)
	return r.OrderWaiverDetails
}

func (r *RedemptionExecution16) AddGatingOrHoldBackDetails() *HoldBackInformation2 {
	r.GatingOrHoldBackDetails = new(HoldBackInformation2)
	return r.GatingOrHoldBackDetails
}

// Execution of a redemption order.
type RedemptionExecution3 struct {

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous identifier for an order execution, as assigned by a confirming party.
	DealReference *Max35Text `xml:"DealRef"`

	// Specifies the category of the investment fund order.
	OrderType []*FundOrderType1 `xml:"OrdrTp,omitempty"`

	// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
	InvestmentAccountDetails *InvestmentAccount13 `xml:"InvstmtAcctDtls"`

	// Additional information about the investor.
	BeneficiaryDetails *IndividualPerson2 `xml:"BnfcryDtls,omitempty"`

	// Number of investment funds units redeemed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Net amount of money paid to the investor as a result of the redemption.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt"`

	// Portion of the investor's holdings, in a specific investment fund/ fund class, that is redeemed.
	HoldingsRedemptionRate *PercentageRate `xml:"HldgsRedRate,omitempty"`

	// Amount of money paid to the investor as a result of the redemption, including all charges, commissions, and tax.
	GrossAmount *ActiveCurrencyAndAmount `xml:"GrssAmt,omitempty"`

	// Date and time at which a price is applied, according to the terms stated in the prospectus.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Price at which the order was executed.
	PriceDetails *UnitPrice5 `xml:"PricDtls"`

	// Indicates whether the order has been partially executed, ie, the confirmed quantity does not match the ordered quantity for a given financial instrument.
	PartiallyExecutedIndicator *YesNoIndicator `xml:"PrtlyExctdInd"`

	// Indicates whether the dividend is included, ie, cum-dividend, in the executed price. When the dividend is not included, the price will be ex-dividend.
	CumDividendIndicator *YesNoIndicator `xml:"CumDvddInd"`

	// Part of the price deemed as accrued income or profit rather than capital. The interim profit amount is used for tax purposes.
	InterimProfitAmount *ProfitAndLoss1Choice `xml:"IntrmPrftAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails []*ForeignExchangeTerms4 `xml:"FXDtls,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnitCode `xml:"Grp1Or2Units,omitempty"`

	// Amount of money associated with a service.
	ChargeGeneralDetails *TotalCharges2 `xml:"ChrgGnlDtls,omitempty"`

	// Amount of money due to a party as compensation for a service.
	CommissionGeneralDetails *TotalCommissions2 `xml:"ComssnGnlDtls,omitempty"`

	// Tax related to an investment fund order.
	TaxGeneralDetails *TotalTaxes2 `xml:"TaxGnlDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters3 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Parameters of a physical delivery.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`

	// Payment transaction resulting from the investment fund order execution.
	CashSettlementDetails *PaymentTransaction18 `xml:"CshSttlmDtls,omitempty"`
}

func (r *RedemptionExecution3) SetOrderReference(value string) {
	r.OrderReference = (*Max35Text)(&value)
}

func (r *RedemptionExecution3) SetDealReference(value string) {
	r.DealReference = (*Max35Text)(&value)
}

func (r *RedemptionExecution3) AddOrderType() *FundOrderType1 {
	newValue := new(FundOrderType1)
	r.OrderType = append(r.OrderType, newValue)
	return newValue
}

func (r *RedemptionExecution3) AddInvestmentAccountDetails() *InvestmentAccount13 {
	r.InvestmentAccountDetails = new(InvestmentAccount13)
	return r.InvestmentAccountDetails
}

func (r *RedemptionExecution3) AddBeneficiaryDetails() *IndividualPerson2 {
	r.BeneficiaryDetails = new(IndividualPerson2)
	return r.BeneficiaryDetails
}

func (r *RedemptionExecution3) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	r.UnitsNumber = new(FinancialInstrumentQuantity1)
	return r.UnitsNumber
}

func (r *RedemptionExecution3) SetRounding(value string) {
	r.Rounding = (*RoundingDirection2Code)(&value)
}

func (r *RedemptionExecution3) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution3) SetHoldingsRedemptionRate(value string) {
	r.HoldingsRedemptionRate = (*PercentageRate)(&value)
}

func (r *RedemptionExecution3) SetGrossAmount(value, currency string) {
	r.GrossAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution3) AddTradeDateTime() *DateAndDateTimeChoice {
	r.TradeDateTime = new(DateAndDateTimeChoice)
	return r.TradeDateTime
}

func (r *RedemptionExecution3) AddPriceDetails() *UnitPrice5 {
	r.PriceDetails = new(UnitPrice5)
	return r.PriceDetails
}

func (r *RedemptionExecution3) SetPartiallyExecutedIndicator(value string) {
	r.PartiallyExecutedIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionExecution3) SetCumDividendIndicator(value string) {
	r.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionExecution3) AddInterimProfitAmount() *ProfitAndLoss1Choice {
	r.InterimProfitAmount = new(ProfitAndLoss1Choice)
	return r.InterimProfitAmount
}

func (r *RedemptionExecution3) AddForeignExchangeDetails() *ForeignExchangeTerms4 {
	newValue := new(ForeignExchangeTerms4)
	r.ForeignExchangeDetails = append(r.ForeignExchangeDetails, newValue)
	return newValue
}

func (r *RedemptionExecution3) SetIncomePreference(value string) {
	r.IncomePreference = (*IncomePreference1Code)(&value)
}

func (r *RedemptionExecution3) SetGroup1Or2Units(value string) {
	r.Group1Or2Units = (*UKTaxGroupUnitCode)(&value)
}

func (r *RedemptionExecution3) AddChargeGeneralDetails() *TotalCharges2 {
	r.ChargeGeneralDetails = new(TotalCharges2)
	return r.ChargeGeneralDetails
}

func (r *RedemptionExecution3) AddCommissionGeneralDetails() *TotalCommissions2 {
	r.CommissionGeneralDetails = new(TotalCommissions2)
	return r.CommissionGeneralDetails
}

func (r *RedemptionExecution3) AddTaxGeneralDetails() *TotalTaxes2 {
	r.TaxGeneralDetails = new(TotalTaxes2)
	return r.TaxGeneralDetails
}

func (r *RedemptionExecution3) AddSettlementAndCustodyDetails() *FundSettlementParameters3 {
	r.SettlementAndCustodyDetails = new(FundSettlementParameters3)
	return r.SettlementAndCustodyDetails
}

func (r *RedemptionExecution3) SetPhysicalDeliveryIndicator(value string) {
	r.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionExecution3) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	r.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return r.PhysicalDeliveryDetails
}

func (r *RedemptionExecution3) AddCashSettlementDetails() *PaymentTransaction18 {
	r.CashSettlementDetails = new(PaymentTransaction18)
	return r.CashSettlementDetails
}

// Execution of a redemption order.
type RedemptionExecution4 struct {

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous identifier for an order execution, as assigned by a confirming party.
	DealReference *Max35Text `xml:"DealRef"`

	// Specifies the category of the investment fund order.
	OrderType []*FundOrderType1 `xml:"OrdrTp,omitempty"`

	// Investment fund class to which an investment fund order execution is related.
	FinancialInstrumentDetails *FinancialInstrument6 `xml:"FinInstrmDtls"`

	// Number of investment funds units redeemed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Net amount of money paid to the investor as a result of the redemption.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt"`

	// Portion of the investor's holdings, in a specific investment fund/ fund class, that is redeemed.
	HoldingsRedemptionRate *PercentageRate `xml:"HldgsRedRate,omitempty"`

	// Amount of money paid to the investor as a result of the redemption, including all charges, commissions, and tax.
	GrossAmount *ActiveCurrencyAndAmount `xml:"GrssAmt,omitempty"`

	// Date and time at which a price is applied, according to the terms stated in the prospectus.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Price at which the order was executed.
	PriceDetails *UnitPrice5 `xml:"PricDtls"`

	// Indicates whether the order has been partially executed, ie, the confirmed quantity does not match the ordered quantity for a given financial instrument.
	PartiallyExecutedIndicator *YesNoIndicator `xml:"PrtlyExctdInd"`

	// Indicates whether the dividend is included, ie, cum-dividend, in the executed price. When the dividend is not included, the price will be ex-dividend.
	CumDividendIndicator *YesNoIndicator `xml:"CumDvddInd"`

	// Part of the price deemed as accrued income or profit rather than capital. The interim profit amount is used for tax purposes.
	InterimProfitAmount *ProfitAndLoss1Choice `xml:"IntrmPrftAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails []*ForeignExchangeTerms4 `xml:"FXDtls,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnitCode `xml:"Grp1Or2Units,omitempty"`

	// Amount of money associated with a service.
	ChargeGeneralDetails *TotalCharges2 `xml:"ChrgGnlDtls,omitempty"`

	// Amount of money due to a party as compensation for a service.
	CommissionGeneralDetails *TotalCommissions2 `xml:"ComssnGnlDtls,omitempty"`

	// Tax related to an investment fund order.
	TaxGeneralDetails *TotalTaxes2 `xml:"TaxGnlDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order
	SettlementAndCustodyDetails *FundSettlementParameters4 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Parameters of a physical delivery.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *CurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *CurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Payment transaction resulting from the investment fund order execution.
	CashSettlementDetails *PaymentTransaction15 `xml:"CshSttlmDtls,omitempty"`
}

func (r *RedemptionExecution4) SetOrderReference(value string) {
	r.OrderReference = (*Max35Text)(&value)
}

func (r *RedemptionExecution4) SetDealReference(value string) {
	r.DealReference = (*Max35Text)(&value)
}

func (r *RedemptionExecution4) AddOrderType() *FundOrderType1 {
	newValue := new(FundOrderType1)
	r.OrderType = append(r.OrderType, newValue)
	return newValue
}

func (r *RedemptionExecution4) AddFinancialInstrumentDetails() *FinancialInstrument6 {
	r.FinancialInstrumentDetails = new(FinancialInstrument6)
	return r.FinancialInstrumentDetails
}

func (r *RedemptionExecution4) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	r.UnitsNumber = new(FinancialInstrumentQuantity1)
	return r.UnitsNumber
}

func (r *RedemptionExecution4) SetRounding(value string) {
	r.Rounding = (*RoundingDirection2Code)(&value)
}

func (r *RedemptionExecution4) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution4) SetHoldingsRedemptionRate(value string) {
	r.HoldingsRedemptionRate = (*PercentageRate)(&value)
}

func (r *RedemptionExecution4) SetGrossAmount(value, currency string) {
	r.GrossAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution4) AddTradeDateTime() *DateAndDateTimeChoice {
	r.TradeDateTime = new(DateAndDateTimeChoice)
	return r.TradeDateTime
}

func (r *RedemptionExecution4) AddPriceDetails() *UnitPrice5 {
	r.PriceDetails = new(UnitPrice5)
	return r.PriceDetails
}

func (r *RedemptionExecution4) SetPartiallyExecutedIndicator(value string) {
	r.PartiallyExecutedIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionExecution4) SetCumDividendIndicator(value string) {
	r.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionExecution4) AddInterimProfitAmount() *ProfitAndLoss1Choice {
	r.InterimProfitAmount = new(ProfitAndLoss1Choice)
	return r.InterimProfitAmount
}

func (r *RedemptionExecution4) AddForeignExchangeDetails() *ForeignExchangeTerms4 {
	newValue := new(ForeignExchangeTerms4)
	r.ForeignExchangeDetails = append(r.ForeignExchangeDetails, newValue)
	return newValue
}

func (r *RedemptionExecution4) SetIncomePreference(value string) {
	r.IncomePreference = (*IncomePreference1Code)(&value)
}

func (r *RedemptionExecution4) SetGroup1Or2Units(value string) {
	r.Group1Or2Units = (*UKTaxGroupUnitCode)(&value)
}

func (r *RedemptionExecution4) AddChargeGeneralDetails() *TotalCharges2 {
	r.ChargeGeneralDetails = new(TotalCharges2)
	return r.ChargeGeneralDetails
}

func (r *RedemptionExecution4) AddCommissionGeneralDetails() *TotalCommissions2 {
	r.CommissionGeneralDetails = new(TotalCommissions2)
	return r.CommissionGeneralDetails
}

func (r *RedemptionExecution4) AddTaxGeneralDetails() *TotalTaxes2 {
	r.TaxGeneralDetails = new(TotalTaxes2)
	return r.TaxGeneralDetails
}

func (r *RedemptionExecution4) AddSettlementAndCustodyDetails() *FundSettlementParameters4 {
	r.SettlementAndCustodyDetails = new(FundSettlementParameters4)
	return r.SettlementAndCustodyDetails
}

func (r *RedemptionExecution4) SetPhysicalDeliveryIndicator(value string) {
	r.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionExecution4) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	r.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return r.PhysicalDeliveryDetails
}

func (r *RedemptionExecution4) SetRequestedSettlementCurrency(value string) {
	r.RequestedSettlementCurrency = (*CurrencyCode)(&value)
}

func (r *RedemptionExecution4) SetRequestedNAVCurrency(value string) {
	r.RequestedNAVCurrency = (*CurrencyCode)(&value)
}

func (r *RedemptionExecution4) AddCashSettlementDetails() *PaymentTransaction15 {
	r.CashSettlementDetails = new(PaymentTransaction15)
	return r.CashSettlementDetails
}

// Execution of a redemption order.
type RedemptionExecution5 struct {

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of an order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for an order execution, as assigned by a confirming party.
	DealReference *Max35Text `xml:"DealRef"`

	// Specifies the category of the investment fund order.
	OrderType []*FundOrderType3 `xml:"OrdrTp,omitempty"`

	// Additional information about the investor.
	BeneficiaryDetails *IndividualPerson12 `xml:"BnfcryDtls,omitempty"`

	// Number of investment funds units redeemed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Amount of money paid to the investor as a result of the redemption after deduction of charges, commissions and taxes.
	// [(Quantity * Price) - (Charges + Commissions +Taxes)]
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
	InvestmentAccountDetails *InvestmentAccount21 `xml:"InvstmtAcctDtls"`

	// Portion of the investor's holdings, in a specific investment fund/ fund class, that is redeemed.
	HoldingsRedemptionRate *PercentageRate `xml:"HldgsRedRate,omitempty"`

	// Amount of money resulting from the redemption before deduction of  charges, commissions and taxes.
	// [Quantity * Price]
	GrossAmount *ActiveCurrencyAndAmount `xml:"GrssAmt,omitempty"`

	// Date and time at which a price is applied, according to the terms stated in the prospectus.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Method by which the transaction is settled.
	SettlementMethod *DeliveryReceiptType2Code `xml:"SttlmMtd,omitempty"`

	// Price at which the order was executed.
	DealingPriceDetails *UnitPrice10 `xml:"DealgPricDtls"`

	// Other quoted price than the one at which the order was executed.
	InformativePriceDetails []*UnitPrice10 `xml:"InftvPricDtls,omitempty"`

	// Indicates whether the order has been partially executed, ie, the confirmed quantity does not match the ordered quantity for a given financial instrument.
	PartiallyExecutedIndicator *YesNoIndicator `xml:"PrtlyExctdInd"`

	// Specifies that the execution was subject to best execution rules as defined by MiFID.
	BestExecution *BestExecution1Code `xml:"BestExctn,omitempty"`

	// Indicates whether the dividend is included, ie, cum-dividend, in the executed price. When the dividend is not included, the price will be ex-dividend.
	CumDividendIndicator *YesNoIndicator `xml:"CumDvddInd"`

	// Part of the price deemed as accrued income or profit rather than capital. The interim profit amount is used for tax purposes.
	InterimProfitAmount *ProfitAndLoss1Choice `xml:"IntrmPrftAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails []*ForeignExchangeTerms7 `xml:"FXDtls,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnitCode `xml:"Grp1Or2Units,omitempty"`

	// Amount of money associated with a service.
	ChargeGeneralDetails *TotalCharges3 `xml:"ChrgGnlDtls,omitempty"`

	// Amount of money due to a party as compensation for a service.
	CommissionGeneralDetails *TotalCommissions3 `xml:"ComssnGnlDtls,omitempty"`

	// Breakdown of the net amount per type of order.
	StaffClientBreakdown []*InvestmentFundsOrderBreakdown1 `xml:"StffClntBrkdwn,omitempty"`

	// Tax related to an investment fund order.
	TaxGeneralDetails *TotalTaxes3 `xml:"TaxGnlDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters4 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Parameters of a physical delivery.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	CashSettlementDetails *PaymentTransaction22 `xml:"CshSttlmDtls,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Percentage of units partially settled.
	PartialSettlementOfUnits *PercentageRate `xml:"PrtlSttlmOfUnits,omitempty"`

	// Percentage of cash partially settled.
	PartialSettlementOfCash *PercentageRate `xml:"PrtlSttlmOfCsh,omitempty"`

	// Amount retained by the Fund and paid out later at a time decided by the Fund.
	PartialRedemptionWithholdingAmount *CurrencyAndAmount `xml:"PrtlRedWhldgAmt,omitempty"`

	// Specifies if advice has been received from an independent financial advisor.
	FinancialAdvice *FinancialAdvice1Code `xml:"FinAdvc,omitempty"`

	// Specifies whether the trade is negotiated.
	NegotiatedTrade *NegotiatedTrade1Code `xml:"NgtdTrad,omitempty"`

	// Specifies whether the order execution confirmation is late.
	LateReport *LateReport1Code `xml:"LateRpt,omitempty"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*Intermediary9 `xml:"RltdPtyDtls,omitempty"`

	// Part of an investor's subscription amount that is held by the fund in order to pay incentive / performance fees at the end of the fiscal year.
	Equalisation *Equalisation1 `xml:"Equlstn,omitempty"`
}

func (r *RedemptionExecution5) SetOrderReference(value string) {
	r.OrderReference = (*Max35Text)(&value)
}

func (r *RedemptionExecution5) SetClientReference(value string) {
	r.ClientReference = (*Max35Text)(&value)
}

func (r *RedemptionExecution5) SetDealReference(value string) {
	r.DealReference = (*Max35Text)(&value)
}

func (r *RedemptionExecution5) AddOrderType() *FundOrderType3 {
	newValue := new(FundOrderType3)
	r.OrderType = append(r.OrderType, newValue)
	return newValue
}

func (r *RedemptionExecution5) AddBeneficiaryDetails() *IndividualPerson12 {
	r.BeneficiaryDetails = new(IndividualPerson12)
	return r.BeneficiaryDetails
}

func (r *RedemptionExecution5) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	r.UnitsNumber = new(FinancialInstrumentQuantity1)
	return r.UnitsNumber
}

func (r *RedemptionExecution5) SetRounding(value string) {
	r.Rounding = (*RoundingDirection2Code)(&value)
}

func (r *RedemptionExecution5) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution5) AddInvestmentAccountDetails() *InvestmentAccount21 {
	r.InvestmentAccountDetails = new(InvestmentAccount21)
	return r.InvestmentAccountDetails
}

func (r *RedemptionExecution5) SetHoldingsRedemptionRate(value string) {
	r.HoldingsRedemptionRate = (*PercentageRate)(&value)
}

func (r *RedemptionExecution5) SetGrossAmount(value, currency string) {
	r.GrossAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution5) AddTradeDateTime() *DateAndDateTimeChoice {
	r.TradeDateTime = new(DateAndDateTimeChoice)
	return r.TradeDateTime
}

func (r *RedemptionExecution5) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution5) SetCashSettlementDate(value string) {
	r.CashSettlementDate = (*ISODate)(&value)
}

func (r *RedemptionExecution5) SetSettlementMethod(value string) {
	r.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (r *RedemptionExecution5) AddDealingPriceDetails() *UnitPrice10 {
	r.DealingPriceDetails = new(UnitPrice10)
	return r.DealingPriceDetails
}

func (r *RedemptionExecution5) AddInformativePriceDetails() *UnitPrice10 {
	newValue := new(UnitPrice10)
	r.InformativePriceDetails = append(r.InformativePriceDetails, newValue)
	return newValue
}

func (r *RedemptionExecution5) SetPartiallyExecutedIndicator(value string) {
	r.PartiallyExecutedIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionExecution5) SetBestExecution(value string) {
	r.BestExecution = (*BestExecution1Code)(&value)
}

func (r *RedemptionExecution5) SetCumDividendIndicator(value string) {
	r.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionExecution5) AddInterimProfitAmount() *ProfitAndLoss1Choice {
	r.InterimProfitAmount = new(ProfitAndLoss1Choice)
	return r.InterimProfitAmount
}

func (r *RedemptionExecution5) AddForeignExchangeDetails() *ForeignExchangeTerms7 {
	newValue := new(ForeignExchangeTerms7)
	r.ForeignExchangeDetails = append(r.ForeignExchangeDetails, newValue)
	return newValue
}

func (r *RedemptionExecution5) SetIncomePreference(value string) {
	r.IncomePreference = (*IncomePreference1Code)(&value)
}

func (r *RedemptionExecution5) SetGroup1Or2Units(value string) {
	r.Group1Or2Units = (*UKTaxGroupUnitCode)(&value)
}

func (r *RedemptionExecution5) AddChargeGeneralDetails() *TotalCharges3 {
	r.ChargeGeneralDetails = new(TotalCharges3)
	return r.ChargeGeneralDetails
}

func (r *RedemptionExecution5) AddCommissionGeneralDetails() *TotalCommissions3 {
	r.CommissionGeneralDetails = new(TotalCommissions3)
	return r.CommissionGeneralDetails
}

func (r *RedemptionExecution5) AddStaffClientBreakdown() *InvestmentFundsOrderBreakdown1 {
	newValue := new(InvestmentFundsOrderBreakdown1)
	r.StaffClientBreakdown = append(r.StaffClientBreakdown, newValue)
	return newValue
}

func (r *RedemptionExecution5) AddTaxGeneralDetails() *TotalTaxes3 {
	r.TaxGeneralDetails = new(TotalTaxes3)
	return r.TaxGeneralDetails
}

func (r *RedemptionExecution5) AddSettlementAndCustodyDetails() *FundSettlementParameters4 {
	r.SettlementAndCustodyDetails = new(FundSettlementParameters4)
	return r.SettlementAndCustodyDetails
}

func (r *RedemptionExecution5) SetPhysicalDeliveryIndicator(value string) {
	r.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionExecution5) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	r.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return r.PhysicalDeliveryDetails
}

func (r *RedemptionExecution5) AddCashSettlementDetails() *PaymentTransaction22 {
	r.CashSettlementDetails = new(PaymentTransaction22)
	return r.CashSettlementDetails
}

func (r *RedemptionExecution5) SetNonStandardSettlementInformation(value string) {
	r.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (r *RedemptionExecution5) SetPartialSettlementOfUnits(value string) {
	r.PartialSettlementOfUnits = (*PercentageRate)(&value)
}

func (r *RedemptionExecution5) SetPartialSettlementOfCash(value string) {
	r.PartialSettlementOfCash = (*PercentageRate)(&value)
}

func (r *RedemptionExecution5) SetPartialRedemptionWithholdingAmount(value, currency string) {
	r.PartialRedemptionWithholdingAmount = NewCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution5) SetFinancialAdvice(value string) {
	r.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (r *RedemptionExecution5) SetNegotiatedTrade(value string) {
	r.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (r *RedemptionExecution5) SetLateReport(value string) {
	r.LateReport = (*LateReport1Code)(&value)
}

func (r *RedemptionExecution5) AddRelatedPartyDetails() *Intermediary9 {
	newValue := new(Intermediary9)
	r.RelatedPartyDetails = append(r.RelatedPartyDetails, newValue)
	return newValue
}

func (r *RedemptionExecution5) AddEqualisation() *Equalisation1 {
	r.Equalisation = new(Equalisation1)
	return r.Equalisation
}

// Execution of a redemption order.
type RedemptionExecution6 struct {

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of an order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for an order execution, as assigned by a confirming party.
	DealReference *Max35Text `xml:"DealRef"`

	// Specifies the category of the investment fund order.
	OrderType []*FundOrderType3 `xml:"OrdrTp,omitempty"`

	// Investment fund class to which an investment fund order execution is related.
	FinancialInstrumentDetails *FinancialInstrument10 `xml:"FinInstrmDtls"`

	// Subdivision of an account used to segregate specific holdings.
	SubAccountForHolding *SubAccount1 `xml:"SubAcctForHldg,omitempty"`

	// Number of investment funds units redeemed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Amount of money paid to the investor as a result of the redemption after deduction of charges, commissions and taxes.
	// [(Quantity * Price) - (Charges + Commissions +Taxes)]
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Portion of the investor's holdings, in a specific investment fund/ fund class, that is redeemed.
	HoldingsRedemptionRate *PercentageRate `xml:"HldgsRedRate,omitempty"`

	// Amount of money resulting from the redemption before deduction of  charges, commissions and taxes.
	// [Quantity * Price]
	GrossAmount *ActiveCurrencyAndAmount `xml:"GrssAmt,omitempty"`

	// Date and time at which a price is applied, according to the terms stated in the prospectus.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Method by which the transaction is settled.
	SettlementMethod *DeliveryReceiptType2Code `xml:"SttlmMtd,omitempty"`

	// Indicates whether the order has been partially executed, ie, the confirmed quantity does not match the ordered quantity for a given financial instrument.
	PartiallyExecutedIndicator *YesNoIndicator `xml:"PrtlyExctdInd"`

	// Price at which the order was executed.
	DealingPriceDetails *UnitPrice10 `xml:"DealgPricDtls"`

	// Other quoted price than the one at which the order was executed.
	InformativePriceDetails []*UnitPrice10 `xml:"InftvPricDtls,omitempty"`

	// Specifies that the execution was subject to best execution rules as defined by MiFID.
	BestExecution *BestExecution1Code `xml:"BestExctn,omitempty"`

	// Indicates whether the dividend is included, ie, cum-dividend, in the executed price. When the dividend is not included, the price will be ex-dividend.
	CumDividendIndicator *YesNoIndicator `xml:"CumDvddInd"`

	// Part of the price deemed as accrued income or profit rather than capital. The interim profit amount is used for tax purposes.
	InterimProfitAmount *ProfitAndLoss1Choice `xml:"IntrmPrftAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails []*ForeignExchangeTerms7 `xml:"FXDtls,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnitCode `xml:"Grp1Or2Units,omitempty"`

	// Amount of money associated with a service.
	ChargeGeneralDetails *TotalCharges3 `xml:"ChrgGnlDtls,omitempty"`

	// Breakdown of the net amount per type of order.
	StaffClientBreakdown []*InvestmentFundsOrderBreakdown1 `xml:"StffClntBrkdwn,omitempty"`

	// Amount of money due to a party as compensation for a service.
	CommissionGeneralDetails *TotalCommissions3 `xml:"ComssnGnlDtls,omitempty"`

	// Tax related to an investment fund order.
	TaxGeneralDetails *TotalTaxes3 `xml:"TaxGnlDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order
	SettlementAndCustodyDetails *FundSettlementParameters4 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Parameters of a physical delivery.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *ActiveCurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Payment transaction resulting from the investment fund order execution.
	CashSettlementDetails *PaymentTransaction22 `xml:"CshSttlmDtls,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Percentage of units partially settled.
	PartialSettlementOfUnits *PercentageRate `xml:"PrtlSttlmOfUnits,omitempty"`

	// Percentage of cash partially settled.
	PartialSettlementOfCash *PercentageRate `xml:"PrtlSttlmOfCsh,omitempty"`

	// Amount retained by the Fund and paid out later at a time decided by the Fund.
	PartialRedemptionWithholdingAmount *CurrencyAndAmount `xml:"PrtlRedWhldgAmt,omitempty"`

	// Specifies if advice has been received from an independent financial advisor.
	FinancialAdvice *FinancialAdvice1Code `xml:"FinAdvc,omitempty"`

	// Specifies whether the trade is negotiated.
	NegotiatedTrade *NegotiatedTrade1Code `xml:"NgtdTrad,omitempty"`

	// Specifies whether the order execution confirmation is late.
	LateReport *LateReport1Code `xml:"LateRpt,omitempty"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*Intermediary9 `xml:"RltdPtyDtls,omitempty"`

	// Part of an investor's subscription amount that is held by the fund in order to pay incentive / performance fees at the end of the fiscal year.
	Equalisation *Equalisation1 `xml:"Equlstn,omitempty"`
}

func (r *RedemptionExecution6) SetOrderReference(value string) {
	r.OrderReference = (*Max35Text)(&value)
}

func (r *RedemptionExecution6) SetClientReference(value string) {
	r.ClientReference = (*Max35Text)(&value)
}

func (r *RedemptionExecution6) SetDealReference(value string) {
	r.DealReference = (*Max35Text)(&value)
}

func (r *RedemptionExecution6) AddOrderType() *FundOrderType3 {
	newValue := new(FundOrderType3)
	r.OrderType = append(r.OrderType, newValue)
	return newValue
}

func (r *RedemptionExecution6) AddFinancialInstrumentDetails() *FinancialInstrument10 {
	r.FinancialInstrumentDetails = new(FinancialInstrument10)
	return r.FinancialInstrumentDetails
}

func (r *RedemptionExecution6) AddSubAccountForHolding() *SubAccount1 {
	r.SubAccountForHolding = new(SubAccount1)
	return r.SubAccountForHolding
}

func (r *RedemptionExecution6) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	r.UnitsNumber = new(FinancialInstrumentQuantity1)
	return r.UnitsNumber
}

func (r *RedemptionExecution6) SetRounding(value string) {
	r.Rounding = (*RoundingDirection2Code)(&value)
}

func (r *RedemptionExecution6) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution6) SetHoldingsRedemptionRate(value string) {
	r.HoldingsRedemptionRate = (*PercentageRate)(&value)
}

func (r *RedemptionExecution6) SetGrossAmount(value, currency string) {
	r.GrossAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution6) AddTradeDateTime() *DateAndDateTimeChoice {
	r.TradeDateTime = new(DateAndDateTimeChoice)
	return r.TradeDateTime
}

func (r *RedemptionExecution6) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution6) SetCashSettlementDate(value string) {
	r.CashSettlementDate = (*ISODate)(&value)
}

func (r *RedemptionExecution6) SetSettlementMethod(value string) {
	r.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (r *RedemptionExecution6) SetPartiallyExecutedIndicator(value string) {
	r.PartiallyExecutedIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionExecution6) AddDealingPriceDetails() *UnitPrice10 {
	r.DealingPriceDetails = new(UnitPrice10)
	return r.DealingPriceDetails
}

func (r *RedemptionExecution6) AddInformativePriceDetails() *UnitPrice10 {
	newValue := new(UnitPrice10)
	r.InformativePriceDetails = append(r.InformativePriceDetails, newValue)
	return newValue
}

func (r *RedemptionExecution6) SetBestExecution(value string) {
	r.BestExecution = (*BestExecution1Code)(&value)
}

func (r *RedemptionExecution6) SetCumDividendIndicator(value string) {
	r.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionExecution6) AddInterimProfitAmount() *ProfitAndLoss1Choice {
	r.InterimProfitAmount = new(ProfitAndLoss1Choice)
	return r.InterimProfitAmount
}

func (r *RedemptionExecution6) AddForeignExchangeDetails() *ForeignExchangeTerms7 {
	newValue := new(ForeignExchangeTerms7)
	r.ForeignExchangeDetails = append(r.ForeignExchangeDetails, newValue)
	return newValue
}

func (r *RedemptionExecution6) SetIncomePreference(value string) {
	r.IncomePreference = (*IncomePreference1Code)(&value)
}

func (r *RedemptionExecution6) SetGroup1Or2Units(value string) {
	r.Group1Or2Units = (*UKTaxGroupUnitCode)(&value)
}

func (r *RedemptionExecution6) AddChargeGeneralDetails() *TotalCharges3 {
	r.ChargeGeneralDetails = new(TotalCharges3)
	return r.ChargeGeneralDetails
}

func (r *RedemptionExecution6) AddStaffClientBreakdown() *InvestmentFundsOrderBreakdown1 {
	newValue := new(InvestmentFundsOrderBreakdown1)
	r.StaffClientBreakdown = append(r.StaffClientBreakdown, newValue)
	return newValue
}

func (r *RedemptionExecution6) AddCommissionGeneralDetails() *TotalCommissions3 {
	r.CommissionGeneralDetails = new(TotalCommissions3)
	return r.CommissionGeneralDetails
}

func (r *RedemptionExecution6) AddTaxGeneralDetails() *TotalTaxes3 {
	r.TaxGeneralDetails = new(TotalTaxes3)
	return r.TaxGeneralDetails
}

func (r *RedemptionExecution6) AddSettlementAndCustodyDetails() *FundSettlementParameters4 {
	r.SettlementAndCustodyDetails = new(FundSettlementParameters4)
	return r.SettlementAndCustodyDetails
}

func (r *RedemptionExecution6) SetPhysicalDeliveryIndicator(value string) {
	r.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionExecution6) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	r.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return r.PhysicalDeliveryDetails
}

func (r *RedemptionExecution6) SetRequestedSettlementCurrency(value string) {
	r.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (r *RedemptionExecution6) SetRequestedNAVCurrency(value string) {
	r.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (r *RedemptionExecution6) AddCashSettlementDetails() *PaymentTransaction22 {
	r.CashSettlementDetails = new(PaymentTransaction22)
	return r.CashSettlementDetails
}

func (r *RedemptionExecution6) SetNonStandardSettlementInformation(value string) {
	r.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (r *RedemptionExecution6) SetPartialSettlementOfUnits(value string) {
	r.PartialSettlementOfUnits = (*PercentageRate)(&value)
}

func (r *RedemptionExecution6) SetPartialSettlementOfCash(value string) {
	r.PartialSettlementOfCash = (*PercentageRate)(&value)
}

func (r *RedemptionExecution6) SetPartialRedemptionWithholdingAmount(value, currency string) {
	r.PartialRedemptionWithholdingAmount = NewCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution6) SetFinancialAdvice(value string) {
	r.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (r *RedemptionExecution6) SetNegotiatedTrade(value string) {
	r.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (r *RedemptionExecution6) SetLateReport(value string) {
	r.LateReport = (*LateReport1Code)(&value)
}

func (r *RedemptionExecution6) AddRelatedPartyDetails() *Intermediary9 {
	newValue := new(Intermediary9)
	r.RelatedPartyDetails = append(r.RelatedPartyDetails, newValue)
	return newValue
}

func (r *RedemptionExecution6) AddEqualisation() *Equalisation1 {
	r.Equalisation = new(Equalisation1)
	return r.Equalisation
}
