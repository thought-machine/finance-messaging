package iso20022

// Execution of a subscription order.
type SubscriptionExecution12 struct {

	// Unique and unambiguous identifier for the order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of the order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for the order execution, as assigned by the confirming party.
	DealReference *Max35Text `xml:"DealRef"`

	// Category of the investment fund order.
	OrderType []*FundOrderType4Choice `xml:"OrdrTp,omitempty"`

	// Account impacted by the investment fund order execution.
	InvestmentAccountDetails *InvestmentAccount58 `xml:"InvstmtAcctDtls"`

	// Additional information about the investor.
	BeneficiaryDetails []*IndividualPerson31 `xml:"BnfcryDtls,omitempty"`

	// Number of investment fund units subscribed.
	UnitsNumber *DecimalNumber `xml:"UnitsNb"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Amount of money invested in the fund.
	// Net Amount = Quantity * Price.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Amount of money to be paid by the investor when subscribing to fund units.
	// Gross amount = (Quantity * Price) + (Fees + Taxes).
	GrossAmount *ActiveCurrencyAndAmount `xml:"GrssAmt,omitempty"`

	// Date and time at which a price is applied, according to the terms stated in the prospectus.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Price at which the order was executed.
	DealingPriceDetails *UnitPrice22 `xml:"DealgPricDtls"`

	// Other quoted price than the one at which the order was executed.
	InformativePriceDetails []*UnitPrice22 `xml:"InftvPricDtls,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
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

	// Reference of a letter of intent program, in which sales commissions are reduced based on the aggregate of a customer's actual purchase and anticipated purchases, over a specific period of time, and as agreed by the customer. A letter of intent program is mainly used in the US market.
	LetterIntentReference *Max35Text `xml:"LttrInttRef,omitempty"`

	// Reference of an accumulation right program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

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

	// Breakdown of the net amount per type of order.
	StaffClientBreakdown []*InvestmentFundsOrderBreakdown2 `xml:"StffClntBrkdwn,omitempty"`

	// Return of cash that has been overpaid for a subscription.
	Refund *ActiveCurrencyAndAmount `xml:"Rfnd,omitempty"`

	// Interest received when a subscription amount is paid in advance and then invested by the fund.
	SubscriptionInterest *ActiveCurrencyAndAmount `xml:"SbcptIntrst,omitempty"`

	// Payment process for the transfer of cash from the debtor to the creditor.
	CashSettlementDetails *PaymentTransaction70 `xml:"CshSttlmDtls,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Percentage of units partially settled.
	PartialSettlementOfUnits *PercentageRate `xml:"PrtlSttlmOfUnits,omitempty"`

	// Specifies if advice has been received from an independent financial advisor.
	FinancialAdvice *FinancialAdvice1Code `xml:"FinAdvc,omitempty"`

	// Specifies whether the trade is negotiated.
	NegotiatedTrade *NegotiatedTrade1Code `xml:"NgtdTrad,omitempty"`

	// Specifies whether the order execution confirmation is late.
	LateReport *LateReport1Code `xml:"LateRpt,omitempty"`

	// Percentage of cash partially settled.
	PartialSettlementOfCash *PercentageRate `xml:"PrtlSttlmOfCsh,omitempty"`

	// Party related to the transaction.
	RelatedPartyDetails []*Intermediary39 `xml:"RltdPtyDtls,omitempty"`

	// Part of an investor's subscription amount that is held by the fund in order to pay incentive / performance fees at the end of the fiscal year.
	Equalisation *Equalisation1 `xml:"Equlstn,omitempty"`

	// Source of cash used for the settlement of the subscription.
	SourceOfCash []*SourceOfCash1Choice `xml:"SrcOfCsh,omitempty"`

	// Assessment of the customer’s behaviour at the time of the account opening application.
	CustomerConductClassification *CustomerConductClassification1Choice `xml:"CstmrCndctClssfctn,omitempty"`

	// Means by which the investor or account owner submits the open account form.
	TransactionChannelType *TransactionChannelType1Choice `xml:"TxChanlTp,omitempty"`

	// Type of signature.
	SignatureType *SignatureType1Choice `xml:"SgntrTp,omitempty"`

	// Information about a non-standard order.
	OrderWaiverDetails *OrderWaiver1 `xml:"OrdrWvrDtls,omitempty"`
}

func (s *SubscriptionExecution12) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution12) SetClientReference(value string) {
	s.ClientReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution12) SetDealReference(value string) {
	s.DealReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution12) AddOrderType() *FundOrderType4Choice {
	newValue := new(FundOrderType4Choice)
	s.OrderType = append(s.OrderType, newValue)
	return newValue
}

func (s *SubscriptionExecution12) AddInvestmentAccountDetails() *InvestmentAccount58 {
	s.InvestmentAccountDetails = new(InvestmentAccount58)
	return s.InvestmentAccountDetails
}

func (s *SubscriptionExecution12) AddBeneficiaryDetails() *IndividualPerson31 {
	newValue := new(IndividualPerson31)
	s.BeneficiaryDetails = append(s.BeneficiaryDetails, newValue)
	return newValue
}

func (s *SubscriptionExecution12) SetUnitsNumber(value string) {
	s.UnitsNumber = (*DecimalNumber)(&value)
}

func (s *SubscriptionExecution12) SetRounding(value string) {
	s.Rounding = (*RoundingDirection2Code)(&value)
}

func (s *SubscriptionExecution12) SetNetAmount(value, currency string) {
	s.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution12) SetGrossAmount(value, currency string) {
	s.GrossAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution12) AddTradeDateTime() *DateAndDateTimeChoice {
	s.TradeDateTime = new(DateAndDateTimeChoice)
	return s.TradeDateTime
}

func (s *SubscriptionExecution12) AddDealingPriceDetails() *UnitPrice22 {
	s.DealingPriceDetails = new(UnitPrice22)
	return s.DealingPriceDetails
}

func (s *SubscriptionExecution12) AddInformativePriceDetails() *UnitPrice22 {
	newValue := new(UnitPrice22)
	s.InformativePriceDetails = append(s.InformativePriceDetails, newValue)
	return newValue
}

func (s *SubscriptionExecution12) SetSettlementAmount(value, currency string) {
	s.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution12) SetCashSettlementDate(value string) {
	s.CashSettlementDate = (*ISODate)(&value)
}

func (s *SubscriptionExecution12) SetSettlementMethod(value string) {
	s.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (s *SubscriptionExecution12) SetPartiallyExecutedIndicator(value string) {
	s.PartiallyExecutedIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionExecution12) SetBestExecution(value string) {
	s.BestExecution = (*BestExecution1Code)(&value)
}

func (s *SubscriptionExecution12) SetCumDividendIndicator(value string) {
	s.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionExecution12) AddInterimProfitAmount() *ProfitAndLoss2Choice {
	s.InterimProfitAmount = new(ProfitAndLoss2Choice)
	return s.InterimProfitAmount
}

func (s *SubscriptionExecution12) AddForeignExchangeDetails() *ForeignExchangeTerms33 {
	newValue := new(ForeignExchangeTerms33)
	s.ForeignExchangeDetails = append(s.ForeignExchangeDetails, newValue)
	return newValue
}

func (s *SubscriptionExecution12) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SubscriptionExecution12) SetLetterIntentReference(value string) {
	s.LetterIntentReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution12) SetAccumulationRightReference(value string) {
	s.AccumulationRightReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution12) AddTransactionOverhead() *TotalFeesAndTaxes40 {
	s.TransactionOverhead = new(TotalFeesAndTaxes40)
	return s.TransactionOverhead
}

func (s *SubscriptionExecution12) AddInformativeTaxDetails() *InformativeTax1 {
	s.InformativeTaxDetails = new(InformativeTax1)
	return s.InformativeTaxDetails
}

func (s *SubscriptionExecution12) AddSettlementAndCustodyDetails() *FundSettlementParameters12 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters12)
	return s.SettlementAndCustodyDetails
}

func (s *SubscriptionExecution12) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionExecution12) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	s.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return s.PhysicalDeliveryDetails
}

func (s *SubscriptionExecution12) AddStaffClientBreakdown() *InvestmentFundsOrderBreakdown2 {
	newValue := new(InvestmentFundsOrderBreakdown2)
	s.StaffClientBreakdown = append(s.StaffClientBreakdown, newValue)
	return newValue
}

func (s *SubscriptionExecution12) SetRefund(value, currency string) {
	s.Refund = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution12) SetSubscriptionInterest(value, currency string) {
	s.SubscriptionInterest = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution12) AddCashSettlementDetails() *PaymentTransaction70 {
	s.CashSettlementDetails = new(PaymentTransaction70)
	return s.CashSettlementDetails
}

func (s *SubscriptionExecution12) SetNonStandardSettlementInformation(value string) {
	s.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (s *SubscriptionExecution12) SetPartialSettlementOfUnits(value string) {
	s.PartialSettlementOfUnits = (*PercentageRate)(&value)
}

func (s *SubscriptionExecution12) SetFinancialAdvice(value string) {
	s.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (s *SubscriptionExecution12) SetNegotiatedTrade(value string) {
	s.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (s *SubscriptionExecution12) SetLateReport(value string) {
	s.LateReport = (*LateReport1Code)(&value)
}

func (s *SubscriptionExecution12) SetPartialSettlementOfCash(value string) {
	s.PartialSettlementOfCash = (*PercentageRate)(&value)
}

func (s *SubscriptionExecution12) AddRelatedPartyDetails() *Intermediary39 {
	newValue := new(Intermediary39)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SubscriptionExecution12) AddEqualisation() *Equalisation1 {
	s.Equalisation = new(Equalisation1)
	return s.Equalisation
}

func (s *SubscriptionExecution12) AddSourceOfCash() *SourceOfCash1Choice {
	newValue := new(SourceOfCash1Choice)
	s.SourceOfCash = append(s.SourceOfCash, newValue)
	return newValue
}

func (s *SubscriptionExecution12) AddCustomerConductClassification() *CustomerConductClassification1Choice {
	s.CustomerConductClassification = new(CustomerConductClassification1Choice)
	return s.CustomerConductClassification
}

func (s *SubscriptionExecution12) AddTransactionChannelType() *TransactionChannelType1Choice {
	s.TransactionChannelType = new(TransactionChannelType1Choice)
	return s.TransactionChannelType
}

func (s *SubscriptionExecution12) AddSignatureType() *SignatureType1Choice {
	s.SignatureType = new(SignatureType1Choice)
	return s.SignatureType
}

func (s *SubscriptionExecution12) AddOrderWaiverDetails() *OrderWaiver1 {
	s.OrderWaiverDetails = new(OrderWaiver1)
	return s.OrderWaiverDetails
}

// Execution of a subscription order.
type SubscriptionExecution13 struct {

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

	// Number of investment fund units subscribed.
	UnitsNumber *DecimalNumber `xml:"UnitsNb"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Amount of money invested in the fund.
	// Net Amount = Quantity * Price.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Amount of money to be paid by the investor when subscribing to fund units.
	// Gross amount = (Quantity * Price) + (Fees + Taxes).
	GrossAmount *ActiveCurrencyAndAmount `xml:"GrssAmt,omitempty"`

	// Date and time at which a price is applied, according to the terms stated in the prospectus.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Price at which the order was executed.
	DealingPriceDetails *UnitPrice22 `xml:"DealgPricDtls"`

	// Other quoted price than the one at which the order was executed.
	InformativePriceDetails []*UnitPrice22 `xml:"InftvPricDtls,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
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

	// Reference of a letter of intent program, in which sales commissions are reduced based on the aggregate of a customer's actual purchase and anticipated purchases, over a specific period of time, and as agreed by the customer. A letter of intent program is mainly used in the US market.
	LetterIntentReference *Max35Text `xml:"LttrInttRef,omitempty"`

	// Reference of an accumulation right program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Fees (charges/commission) and taxes that are taken into consideration for the transaction, so that the total difference between the net amount and gross amount is known, without taking into account equalisation.
	TransactionOverhead *TotalFeesAndTaxes40 `xml:"TxOvrhd,omitempty"`

	// Information about tax that does not have an impact on the transaction overhead.
	InformativeTaxDetails *InformativeTax1 `xml:"InftvTaxDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters12 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to the physical delivery of the securities.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *ActiveCurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Return of cash that has been overpaid for a subscription.
	Refund *ActiveCurrencyAndAmount `xml:"Rfnd,omitempty"`

	// Interest received when a subscription amount is paid in advance and then invested by the fund.
	SubscriptionInterest *ActiveCurrencyAndAmount `xml:"SbcptIntrst,omitempty"`

	// Payment process for the transfer of cash from the debtor to the creditor.
	CashSettlementDetails *PaymentTransaction70 `xml:"CshSttlmDtls,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Percentage of units partially settled.
	PartialSettlementOfUnits *PercentageRate `xml:"PrtlSttlmOfUnits,omitempty"`

	// Percentage of cash partially settled.
	PartialSettlementOfCash *PercentageRate `xml:"PrtlSttlmOfCsh,omitempty"`

	// Breakdown of the net amount per type of order.
	StaffClientBreakdown []*InvestmentFundsOrderBreakdown2 `xml:"StffClntBrkdwn,omitempty"`

	// Specifies if advice has been received from an independent financial advisor.
	FinancialAdvice *FinancialAdvice1Code `xml:"FinAdvc,omitempty"`

	// Specifies whether the trade is negotiated.
	NegotiatedTrade *NegotiatedTrade1Code `xml:"NgtdTrad,omitempty"`

	// Specifies whether the order execution confirmation is late.
	LateReport *LateReport1Code `xml:"LateRpt,omitempty"`

	// Party related to the transaction.
	RelatedPartyDetails []*Intermediary39 `xml:"RltdPtyDtls,omitempty"`

	// Part of an investor's subscription amount that is held by the fund in order to pay incentive/performance fees at the end of the fiscal year.
	Equalisation *Equalisation1 `xml:"Equlstn,omitempty"`

	// Source of cash used for the settlement of the subscription.
	SourceOfCash []*SourceOfCash1Choice `xml:"SrcOfCsh,omitempty"`

	// Assessment of the customer’s behaviour at the time of the account opening application.
	CustomerConductClassification *CustomerConductClassification1Choice `xml:"CstmrCndctClssfctn,omitempty"`

	// Means by which the investor or account owner submits the open account form.
	TransactionChannelType *TransactionChannelType1Choice `xml:"TxChanlTp,omitempty"`

	// Type of signature.
	SignatureType *SignatureType1Choice `xml:"SgntrTp,omitempty"`

	// Information about a non-standard order.
	OrderWaiverDetails *OrderWaiver1 `xml:"OrdrWvrDtls,omitempty"`
}

func (s *SubscriptionExecution13) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution13) SetClientReference(value string) {
	s.ClientReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution13) SetDealReference(value string) {
	s.DealReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution13) AddOrderType() *FundOrderType4Choice {
	newValue := new(FundOrderType4Choice)
	s.OrderType = append(s.OrderType, newValue)
	return newValue
}

func (s *SubscriptionExecution13) AddFinancialInstrumentDetails() *FinancialInstrument57 {
	s.FinancialInstrumentDetails = new(FinancialInstrument57)
	return s.FinancialInstrumentDetails
}

func (s *SubscriptionExecution13) AddSubAccountForHolding() *SubAccount6 {
	s.SubAccountForHolding = new(SubAccount6)
	return s.SubAccountForHolding
}

func (s *SubscriptionExecution13) SetUnitsNumber(value string) {
	s.UnitsNumber = (*DecimalNumber)(&value)
}

func (s *SubscriptionExecution13) SetRounding(value string) {
	s.Rounding = (*RoundingDirection2Code)(&value)
}

func (s *SubscriptionExecution13) SetNetAmount(value, currency string) {
	s.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution13) SetGrossAmount(value, currency string) {
	s.GrossAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution13) AddTradeDateTime() *DateAndDateTimeChoice {
	s.TradeDateTime = new(DateAndDateTimeChoice)
	return s.TradeDateTime
}

func (s *SubscriptionExecution13) AddDealingPriceDetails() *UnitPrice22 {
	s.DealingPriceDetails = new(UnitPrice22)
	return s.DealingPriceDetails
}

func (s *SubscriptionExecution13) AddInformativePriceDetails() *UnitPrice22 {
	newValue := new(UnitPrice22)
	s.InformativePriceDetails = append(s.InformativePriceDetails, newValue)
	return newValue
}

func (s *SubscriptionExecution13) SetSettlementAmount(value, currency string) {
	s.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution13) SetCashSettlementDate(value string) {
	s.CashSettlementDate = (*ISODate)(&value)
}

func (s *SubscriptionExecution13) SetSettlementMethod(value string) {
	s.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (s *SubscriptionExecution13) SetPartiallyExecutedIndicator(value string) {
	s.PartiallyExecutedIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionExecution13) SetBestExecution(value string) {
	s.BestExecution = (*BestExecution1Code)(&value)
}

func (s *SubscriptionExecution13) SetCumDividendIndicator(value string) {
	s.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionExecution13) AddInterimProfitAmount() *ProfitAndLoss2Choice {
	s.InterimProfitAmount = new(ProfitAndLoss2Choice)
	return s.InterimProfitAmount
}

func (s *SubscriptionExecution13) AddForeignExchangeDetails() *ForeignExchangeTerms33 {
	newValue := new(ForeignExchangeTerms33)
	s.ForeignExchangeDetails = append(s.ForeignExchangeDetails, newValue)
	return newValue
}

func (s *SubscriptionExecution13) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SubscriptionExecution13) SetLetterIntentReference(value string) {
	s.LetterIntentReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution13) SetAccumulationRightReference(value string) {
	s.AccumulationRightReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution13) AddTransactionOverhead() *TotalFeesAndTaxes40 {
	s.TransactionOverhead = new(TotalFeesAndTaxes40)
	return s.TransactionOverhead
}

func (s *SubscriptionExecution13) AddInformativeTaxDetails() *InformativeTax1 {
	s.InformativeTaxDetails = new(InformativeTax1)
	return s.InformativeTaxDetails
}

func (s *SubscriptionExecution13) AddSettlementAndCustodyDetails() *FundSettlementParameters12 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters12)
	return s.SettlementAndCustodyDetails
}

func (s *SubscriptionExecution13) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionExecution13) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	s.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return s.PhysicalDeliveryDetails
}

func (s *SubscriptionExecution13) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (s *SubscriptionExecution13) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (s *SubscriptionExecution13) SetRefund(value, currency string) {
	s.Refund = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution13) SetSubscriptionInterest(value, currency string) {
	s.SubscriptionInterest = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution13) AddCashSettlementDetails() *PaymentTransaction70 {
	s.CashSettlementDetails = new(PaymentTransaction70)
	return s.CashSettlementDetails
}

func (s *SubscriptionExecution13) SetNonStandardSettlementInformation(value string) {
	s.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (s *SubscriptionExecution13) SetPartialSettlementOfUnits(value string) {
	s.PartialSettlementOfUnits = (*PercentageRate)(&value)
}

func (s *SubscriptionExecution13) SetPartialSettlementOfCash(value string) {
	s.PartialSettlementOfCash = (*PercentageRate)(&value)
}

func (s *SubscriptionExecution13) AddStaffClientBreakdown() *InvestmentFundsOrderBreakdown2 {
	newValue := new(InvestmentFundsOrderBreakdown2)
	s.StaffClientBreakdown = append(s.StaffClientBreakdown, newValue)
	return newValue
}

func (s *SubscriptionExecution13) SetFinancialAdvice(value string) {
	s.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (s *SubscriptionExecution13) SetNegotiatedTrade(value string) {
	s.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (s *SubscriptionExecution13) SetLateReport(value string) {
	s.LateReport = (*LateReport1Code)(&value)
}

func (s *SubscriptionExecution13) AddRelatedPartyDetails() *Intermediary39 {
	newValue := new(Intermediary39)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SubscriptionExecution13) AddEqualisation() *Equalisation1 {
	s.Equalisation = new(Equalisation1)
	return s.Equalisation
}

func (s *SubscriptionExecution13) AddSourceOfCash() *SourceOfCash1Choice {
	newValue := new(SourceOfCash1Choice)
	s.SourceOfCash = append(s.SourceOfCash, newValue)
	return newValue
}

func (s *SubscriptionExecution13) AddCustomerConductClassification() *CustomerConductClassification1Choice {
	s.CustomerConductClassification = new(CustomerConductClassification1Choice)
	return s.CustomerConductClassification
}

func (s *SubscriptionExecution13) AddTransactionChannelType() *TransactionChannelType1Choice {
	s.TransactionChannelType = new(TransactionChannelType1Choice)
	return s.TransactionChannelType
}

func (s *SubscriptionExecution13) AddSignatureType() *SignatureType1Choice {
	s.SignatureType = new(SignatureType1Choice)
	return s.SignatureType
}

func (s *SubscriptionExecution13) AddOrderWaiverDetails() *OrderWaiver1 {
	s.OrderWaiverDetails = new(OrderWaiver1)
	return s.OrderWaiverDetails
}

// Execution of a subscription order.
type SubscriptionExecution3 struct {

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

	// Number of investment fund units subscribed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Net amount of money invested in a specific financial instrument by an investor, expressed in the currency requested by the investor.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt"`

	// Amount of money invested in a specific financial instrument by an investor, including all charges, commissions, and tax, expressed in the currency requested by the investor.
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

	// Reference of a letter of intent program, in which sales commissions are reduced based on the aggregate of a customer's actual purchase and anticipated purchases, over a specific period of time, and as agreed by the customer. A letter of intent program is mainly used in the US market.
	LetterIntentReference *Max35Text `xml:"LttrInttRef,omitempty"`

	// Reference of an accumulation right program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Amount of money associated with a service.
	ChargeGeneralDetails *TotalCharges2 `xml:"ChrgGnlDtls,omitempty"`

	// Amount of money due to a party as compensation for a service.
	CommissionGeneralDetails *TotalCommissions2 `xml:"ComssnGnlDtls,omitempty"`

	// Tax related to an investment fund order.
	TaxGeneralDetails *TotalTaxes2 `xml:"TaxGnlDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters3 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the securities are to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to physical delivery of the securities.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`

	// Return of cash that has been overpaid for a subscription.
	Refund *ActiveCurrencyAndAmount `xml:"Rfnd,omitempty"`

	// Interest received when a subscription amount is paid in advance and then invested by the fund.
	SubscriptionInterest *ActiveCurrencyAndAmount `xml:"SbcptIntrst,omitempty"`

	// Payment transaction resulting from the investment fund order execution.
	CashSettlementDetails *PaymentTransaction16 `xml:"CshSttlmDtls,omitempty"`
}

func (s *SubscriptionExecution3) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution3) SetDealReference(value string) {
	s.DealReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution3) AddOrderType() *FundOrderType1 {
	newValue := new(FundOrderType1)
	s.OrderType = append(s.OrderType, newValue)
	return newValue
}

func (s *SubscriptionExecution3) AddInvestmentAccountDetails() *InvestmentAccount13 {
	s.InvestmentAccountDetails = new(InvestmentAccount13)
	return s.InvestmentAccountDetails
}

func (s *SubscriptionExecution3) AddBeneficiaryDetails() *IndividualPerson2 {
	s.BeneficiaryDetails = new(IndividualPerson2)
	return s.BeneficiaryDetails
}

func (s *SubscriptionExecution3) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	s.UnitsNumber = new(FinancialInstrumentQuantity1)
	return s.UnitsNumber
}

func (s *SubscriptionExecution3) SetRounding(value string) {
	s.Rounding = (*RoundingDirection2Code)(&value)
}

func (s *SubscriptionExecution3) SetNetAmount(value, currency string) {
	s.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution3) SetGrossAmount(value, currency string) {
	s.GrossAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution3) AddTradeDateTime() *DateAndDateTimeChoice {
	s.TradeDateTime = new(DateAndDateTimeChoice)
	return s.TradeDateTime
}

func (s *SubscriptionExecution3) AddPriceDetails() *UnitPrice5 {
	s.PriceDetails = new(UnitPrice5)
	return s.PriceDetails
}

func (s *SubscriptionExecution3) SetPartiallyExecutedIndicator(value string) {
	s.PartiallyExecutedIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionExecution3) SetCumDividendIndicator(value string) {
	s.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionExecution3) AddInterimProfitAmount() *ProfitAndLoss1Choice {
	s.InterimProfitAmount = new(ProfitAndLoss1Choice)
	return s.InterimProfitAmount
}

func (s *SubscriptionExecution3) AddForeignExchangeDetails() *ForeignExchangeTerms4 {
	newValue := new(ForeignExchangeTerms4)
	s.ForeignExchangeDetails = append(s.ForeignExchangeDetails, newValue)
	return newValue
}

func (s *SubscriptionExecution3) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SubscriptionExecution3) SetLetterIntentReference(value string) {
	s.LetterIntentReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution3) SetAccumulationRightReference(value string) {
	s.AccumulationRightReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution3) AddChargeGeneralDetails() *TotalCharges2 {
	s.ChargeGeneralDetails = new(TotalCharges2)
	return s.ChargeGeneralDetails
}

func (s *SubscriptionExecution3) AddCommissionGeneralDetails() *TotalCommissions2 {
	s.CommissionGeneralDetails = new(TotalCommissions2)
	return s.CommissionGeneralDetails
}

func (s *SubscriptionExecution3) AddTaxGeneralDetails() *TotalTaxes2 {
	s.TaxGeneralDetails = new(TotalTaxes2)
	return s.TaxGeneralDetails
}

func (s *SubscriptionExecution3) AddSettlementAndCustodyDetails() *FundSettlementParameters3 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters3)
	return s.SettlementAndCustodyDetails
}

func (s *SubscriptionExecution3) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionExecution3) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	s.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return s.PhysicalDeliveryDetails
}

func (s *SubscriptionExecution3) SetRefund(value, currency string) {
	s.Refund = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution3) SetSubscriptionInterest(value, currency string) {
	s.SubscriptionInterest = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution3) AddCashSettlementDetails() *PaymentTransaction16 {
	s.CashSettlementDetails = new(PaymentTransaction16)
	return s.CashSettlementDetails
}

// Execution of a subscription order.
type SubscriptionExecution4 struct {

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous identifier for an order execution, as assigned by a confirming party.
	DealReference *Max35Text `xml:"DealRef"`

	// Specifies the category of the investment fund order.
	OrderType []*FundOrderType1 `xml:"OrdrTp,omitempty"`

	// Investment fund class to which an investment fund order execution is related.
	FinancialInstrumentDetails *FinancialInstrument6 `xml:"FinInstrmDtls"`

	// Number of investment fund units subscribed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Net amount of money invested in a specific financial instrument by an investor, expressed in the currency requested by the investor.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt"`

	// Amount of money invested in a specific financial instrument by an investor, including all charges, commissions, and tax, expressed in the currency requested by the investor.
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

	// Reference of a letter of intent program, in which sales commissions are reduced based on the aggregate of a customer's actual purchase and anticipated purchases, over a specific period of time, and as agreed by the customer. A letter of intent program is mainly used in the US market.
	LetterIntentReference *Max35Text `xml:"LttrInttRef,omitempty"`

	// Reference of an accumulation right program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

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

	// Information related to physical delivery of the securities.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *CurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *CurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Return of cash that has been overpaid for a subscription.
	Refund *ActiveCurrencyAndAmount `xml:"Rfnd,omitempty"`

	// Interest received when a subscription amount is paid in advance and then invested by the fund.
	SubscriptionInterest *ActiveCurrencyAndAmount `xml:"SbcptIntrst,omitempty"`

	// Payment transaction resulting from the investment fund order execution.
	CashSettlementDetails *PaymentTransaction13 `xml:"CshSttlmDtls,omitempty"`
}

func (s *SubscriptionExecution4) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution4) SetDealReference(value string) {
	s.DealReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution4) AddOrderType() *FundOrderType1 {
	newValue := new(FundOrderType1)
	s.OrderType = append(s.OrderType, newValue)
	return newValue
}

func (s *SubscriptionExecution4) AddFinancialInstrumentDetails() *FinancialInstrument6 {
	s.FinancialInstrumentDetails = new(FinancialInstrument6)
	return s.FinancialInstrumentDetails
}

func (s *SubscriptionExecution4) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	s.UnitsNumber = new(FinancialInstrumentQuantity1)
	return s.UnitsNumber
}

func (s *SubscriptionExecution4) SetRounding(value string) {
	s.Rounding = (*RoundingDirection2Code)(&value)
}

func (s *SubscriptionExecution4) SetNetAmount(value, currency string) {
	s.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution4) SetGrossAmount(value, currency string) {
	s.GrossAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution4) AddTradeDateTime() *DateAndDateTimeChoice {
	s.TradeDateTime = new(DateAndDateTimeChoice)
	return s.TradeDateTime
}

func (s *SubscriptionExecution4) AddPriceDetails() *UnitPrice5 {
	s.PriceDetails = new(UnitPrice5)
	return s.PriceDetails
}

func (s *SubscriptionExecution4) SetPartiallyExecutedIndicator(value string) {
	s.PartiallyExecutedIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionExecution4) SetCumDividendIndicator(value string) {
	s.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionExecution4) AddInterimProfitAmount() *ProfitAndLoss1Choice {
	s.InterimProfitAmount = new(ProfitAndLoss1Choice)
	return s.InterimProfitAmount
}

func (s *SubscriptionExecution4) AddForeignExchangeDetails() *ForeignExchangeTerms4 {
	newValue := new(ForeignExchangeTerms4)
	s.ForeignExchangeDetails = append(s.ForeignExchangeDetails, newValue)
	return newValue
}

func (s *SubscriptionExecution4) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SubscriptionExecution4) SetLetterIntentReference(value string) {
	s.LetterIntentReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution4) SetAccumulationRightReference(value string) {
	s.AccumulationRightReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution4) AddChargeGeneralDetails() *TotalCharges2 {
	s.ChargeGeneralDetails = new(TotalCharges2)
	return s.ChargeGeneralDetails
}

func (s *SubscriptionExecution4) AddCommissionGeneralDetails() *TotalCommissions2 {
	s.CommissionGeneralDetails = new(TotalCommissions2)
	return s.CommissionGeneralDetails
}

func (s *SubscriptionExecution4) AddTaxGeneralDetails() *TotalTaxes2 {
	s.TaxGeneralDetails = new(TotalTaxes2)
	return s.TaxGeneralDetails
}

func (s *SubscriptionExecution4) AddSettlementAndCustodyDetails() *FundSettlementParameters3 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters3)
	return s.SettlementAndCustodyDetails
}

func (s *SubscriptionExecution4) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionExecution4) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	s.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return s.PhysicalDeliveryDetails
}

func (s *SubscriptionExecution4) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*CurrencyCode)(&value)
}

func (s *SubscriptionExecution4) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*CurrencyCode)(&value)
}

func (s *SubscriptionExecution4) SetRefund(value, currency string) {
	s.Refund = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution4) SetSubscriptionInterest(value, currency string) {
	s.SubscriptionInterest = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution4) AddCashSettlementDetails() *PaymentTransaction13 {
	s.CashSettlementDetails = new(PaymentTransaction13)
	return s.CashSettlementDetails
}

// Execution of a subscription order.
type SubscriptionExecution5 struct {

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of an order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for an order execution, as assigned by a confirming party.
	DealReference *Max35Text `xml:"DealRef"`

	// Specifies the category of the investment fund order.
	OrderType []*FundOrderType3 `xml:"OrdrTp,omitempty"`

	// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
	InvestmentAccountDetails *InvestmentAccount21 `xml:"InvstmtAcctDtls"`

	// Additional information about the investor.
	BeneficiaryDetails *IndividualPerson12 `xml:"BnfcryDtls,omitempty"`

	// Number of investment fund units subscribed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Amount of money invested in a specific financial instrument by an investor, after deduction of charges, commissions and taxes.
	// [Quantity * Price]
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Amount of money invested in a specific financial instrument by an investor, before deduction of  charges, commissions, and taxes.
	// [(Quantity * Price) + (Charges + Commissions +Taxes)]
	GrossAmount *ActiveCurrencyAndAmount `xml:"GrssAmt,omitempty"`

	// Date and time at which a price is applied, according to the terms stated in the prospectus.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Price at which the order was executed.
	DealingPriceDetails *UnitPrice10 `xml:"DealgPricDtls"`

	// Other quoted price than the one at which the order was executed.
	InformativePriceDetails []*UnitPrice10 `xml:"InftvPricDtls,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Method by which the transaction is settled.
	SettlementMethod *DeliveryReceiptType2Code `xml:"SttlmMtd,omitempty"`

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

	// Reference of a letter of intent program, in which sales commissions are reduced based on the aggregate of a customer's actual purchase and anticipated purchases, over a specific period of time, and as agreed by the customer. A letter of intent program is mainly used in the US market.
	LetterIntentReference *Max35Text `xml:"LttrInttRef,omitempty"`

	// Reference of an accumulation right program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Amount of money due to a party as compensation for a service.
	CommissionGeneralDetails *TotalCommissions3 `xml:"ComssnGnlDtls,omitempty"`

	// Amount of money associated with a service.
	ChargeGeneralDetails *TotalCharges3 `xml:"ChrgGnlDtls,omitempty"`

	// Tax related to an investment fund order.
	TaxGeneralDetails *TotalTaxes3 `xml:"TaxGnlDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters3 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the securities are to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to physical delivery of the securities.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`

	// Breakdown of the net amount per type of order.
	StaffClientBreakdown []*InvestmentFundsOrderBreakdown1 `xml:"StffClntBrkdwn,omitempty"`

	// Return of cash that has been overpaid for a subscription.
	Refund *ActiveCurrencyAndAmount `xml:"Rfnd,omitempty"`

	// Interest received when a subscription amount is paid in advance and then invested by the fund.
	SubscriptionInterest *ActiveCurrencyAndAmount `xml:"SbcptIntrst,omitempty"`

	// Payment transaction resulting from the investment fund order execution.
	CashSettlementDetails *PaymentTransaction24 `xml:"CshSttlmDtls,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Percentage of units partially settled.
	PartialSettlementOfUnits *PercentageRate `xml:"PrtlSttlmOfUnits,omitempty"`

	// Specifies if advice has been received from an independent financial advisor.
	FinancialAdvice *FinancialAdvice1Code `xml:"FinAdvc,omitempty"`

	// Specifies whether the trade is negotiated.
	NegotiatedTrade *NegotiatedTrade1Code `xml:"NgtdTrad,omitempty"`

	// Specifies whether the order execution confirmation is late.
	LateReport *LateReport1Code `xml:"LateRpt,omitempty"`

	// Percentage of cash partially settled.
	PartialSettlementOfCash *PercentageRate `xml:"PrtlSttlmOfCsh,omitempty"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*Intermediary9 `xml:"RltdPtyDtls,omitempty"`

	// Part of an investor's subscription amount that is held by the fund in order to pay incentive / performance fees at the end of the fiscal year.
	Equalisation *Equalisation1 `xml:"Equlstn,omitempty"`
}

func (s *SubscriptionExecution5) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution5) SetClientReference(value string) {
	s.ClientReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution5) SetDealReference(value string) {
	s.DealReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution5) AddOrderType() *FundOrderType3 {
	newValue := new(FundOrderType3)
	s.OrderType = append(s.OrderType, newValue)
	return newValue
}

func (s *SubscriptionExecution5) AddInvestmentAccountDetails() *InvestmentAccount21 {
	s.InvestmentAccountDetails = new(InvestmentAccount21)
	return s.InvestmentAccountDetails
}

func (s *SubscriptionExecution5) AddBeneficiaryDetails() *IndividualPerson12 {
	s.BeneficiaryDetails = new(IndividualPerson12)
	return s.BeneficiaryDetails
}

func (s *SubscriptionExecution5) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	s.UnitsNumber = new(FinancialInstrumentQuantity1)
	return s.UnitsNumber
}

func (s *SubscriptionExecution5) SetRounding(value string) {
	s.Rounding = (*RoundingDirection2Code)(&value)
}

func (s *SubscriptionExecution5) SetNetAmount(value, currency string) {
	s.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution5) SetGrossAmount(value, currency string) {
	s.GrossAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution5) AddTradeDateTime() *DateAndDateTimeChoice {
	s.TradeDateTime = new(DateAndDateTimeChoice)
	return s.TradeDateTime
}

func (s *SubscriptionExecution5) AddDealingPriceDetails() *UnitPrice10 {
	s.DealingPriceDetails = new(UnitPrice10)
	return s.DealingPriceDetails
}

func (s *SubscriptionExecution5) AddInformativePriceDetails() *UnitPrice10 {
	newValue := new(UnitPrice10)
	s.InformativePriceDetails = append(s.InformativePriceDetails, newValue)
	return newValue
}

func (s *SubscriptionExecution5) SetSettlementAmount(value, currency string) {
	s.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution5) SetCashSettlementDate(value string) {
	s.CashSettlementDate = (*ISODate)(&value)
}

func (s *SubscriptionExecution5) SetSettlementMethod(value string) {
	s.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (s *SubscriptionExecution5) SetPartiallyExecutedIndicator(value string) {
	s.PartiallyExecutedIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionExecution5) SetBestExecution(value string) {
	s.BestExecution = (*BestExecution1Code)(&value)
}

func (s *SubscriptionExecution5) SetCumDividendIndicator(value string) {
	s.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionExecution5) AddInterimProfitAmount() *ProfitAndLoss1Choice {
	s.InterimProfitAmount = new(ProfitAndLoss1Choice)
	return s.InterimProfitAmount
}

func (s *SubscriptionExecution5) AddForeignExchangeDetails() *ForeignExchangeTerms7 {
	newValue := new(ForeignExchangeTerms7)
	s.ForeignExchangeDetails = append(s.ForeignExchangeDetails, newValue)
	return newValue
}

func (s *SubscriptionExecution5) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SubscriptionExecution5) SetLetterIntentReference(value string) {
	s.LetterIntentReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution5) SetAccumulationRightReference(value string) {
	s.AccumulationRightReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution5) AddCommissionGeneralDetails() *TotalCommissions3 {
	s.CommissionGeneralDetails = new(TotalCommissions3)
	return s.CommissionGeneralDetails
}

func (s *SubscriptionExecution5) AddChargeGeneralDetails() *TotalCharges3 {
	s.ChargeGeneralDetails = new(TotalCharges3)
	return s.ChargeGeneralDetails
}

func (s *SubscriptionExecution5) AddTaxGeneralDetails() *TotalTaxes3 {
	s.TaxGeneralDetails = new(TotalTaxes3)
	return s.TaxGeneralDetails
}

func (s *SubscriptionExecution5) AddSettlementAndCustodyDetails() *FundSettlementParameters3 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters3)
	return s.SettlementAndCustodyDetails
}

func (s *SubscriptionExecution5) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionExecution5) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	s.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return s.PhysicalDeliveryDetails
}

func (s *SubscriptionExecution5) AddStaffClientBreakdown() *InvestmentFundsOrderBreakdown1 {
	newValue := new(InvestmentFundsOrderBreakdown1)
	s.StaffClientBreakdown = append(s.StaffClientBreakdown, newValue)
	return newValue
}

func (s *SubscriptionExecution5) SetRefund(value, currency string) {
	s.Refund = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution5) SetSubscriptionInterest(value, currency string) {
	s.SubscriptionInterest = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution5) AddCashSettlementDetails() *PaymentTransaction24 {
	s.CashSettlementDetails = new(PaymentTransaction24)
	return s.CashSettlementDetails
}

func (s *SubscriptionExecution5) SetNonStandardSettlementInformation(value string) {
	s.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (s *SubscriptionExecution5) SetPartialSettlementOfUnits(value string) {
	s.PartialSettlementOfUnits = (*PercentageRate)(&value)
}

func (s *SubscriptionExecution5) SetFinancialAdvice(value string) {
	s.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (s *SubscriptionExecution5) SetNegotiatedTrade(value string) {
	s.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (s *SubscriptionExecution5) SetLateReport(value string) {
	s.LateReport = (*LateReport1Code)(&value)
}

func (s *SubscriptionExecution5) SetPartialSettlementOfCash(value string) {
	s.PartialSettlementOfCash = (*PercentageRate)(&value)
}

func (s *SubscriptionExecution5) AddRelatedPartyDetails() *Intermediary9 {
	newValue := new(Intermediary9)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SubscriptionExecution5) AddEqualisation() *Equalisation1 {
	s.Equalisation = new(Equalisation1)
	return s.Equalisation
}

// Execution of a subscription order.
type SubscriptionExecution6 struct {

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

	// Number of investment fund units subscribed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Amount of money invested in a specific financial instrument by an investor, after deduction of charges, commissions and taxes.
	// [Quantity * Price]
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Amount of money invested in a specific financial instrument by an investor, before deduction of  charges, commissions, and taxes.
	// [(Quantity * Price) + (Charges + Commissions +Taxes)]
	GrossAmount *ActiveCurrencyAndAmount `xml:"GrssAmt,omitempty"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Date and time at which a price is applied, according to the terms stated in the prospectus.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Price at which the order was executed.
	DealingPriceDetails *UnitPrice10 `xml:"DealgPricDtls"`

	// Other quoted price than the one at which the order was executed.
	InformativePriceDetails []*UnitPrice10 `xml:"InftvPricDtls,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Method by which the transaction is settled.
	SettlementMethod *DeliveryReceiptType2Code `xml:"SttlmMtd,omitempty"`

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

	// Reference of a letter of intent program, in which sales commissions are reduced based on the aggregate of a customer's actual purchase and anticipated purchases, over a specific period of time, and as agreed by the customer. A letter of intent program is mainly used in the US market.
	LetterIntentReference *Max35Text `xml:"LttrInttRef,omitempty"`

	// Reference of an accumulation right program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Amount of money associated with a service.
	ChargeGeneralDetails *TotalCharges3 `xml:"ChrgGnlDtls,omitempty"`

	// Amount of money due to a party as compensation for a service.
	CommissionGeneralDetails *TotalCommissions3 `xml:"ComssnGnlDtls,omitempty"`

	// Tax related to an investment fund order.
	TaxGeneralDetails *TotalTaxes3 `xml:"TaxGnlDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters3 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to physical delivery of the securities.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *ActiveCurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Return of cash that has been overpaid for a subscription.
	Refund *ActiveCurrencyAndAmount `xml:"Rfnd,omitempty"`

	// Interest received when a subscription amount is paid in advance and then invested by the fund.
	SubscriptionInterest *ActiveCurrencyAndAmount `xml:"SbcptIntrst,omitempty"`

	// Payment transaction resulting from the investment fund order execution.
	CashSettlementDetails *PaymentTransaction24 `xml:"CshSttlmDtls,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Percentage of units partially settled.
	PartialSettlementOfUnits *PercentageRate `xml:"PrtlSttlmOfUnits,omitempty"`

	// Percentage of cash partially settled.
	PartialSettlementOfCash *PercentageRate `xml:"PrtlSttlmOfCsh,omitempty"`

	// Breakdown of the net amount per type of order.
	StaffClientBreakdown []*InvestmentFundsOrderBreakdown1 `xml:"StffClntBrkdwn,omitempty"`

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

func (s *SubscriptionExecution6) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution6) SetClientReference(value string) {
	s.ClientReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution6) SetDealReference(value string) {
	s.DealReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution6) AddOrderType() *FundOrderType3 {
	newValue := new(FundOrderType3)
	s.OrderType = append(s.OrderType, newValue)
	return newValue
}

func (s *SubscriptionExecution6) AddFinancialInstrumentDetails() *FinancialInstrument10 {
	s.FinancialInstrumentDetails = new(FinancialInstrument10)
	return s.FinancialInstrumentDetails
}

func (s *SubscriptionExecution6) AddSubAccountForHolding() *SubAccount1 {
	s.SubAccountForHolding = new(SubAccount1)
	return s.SubAccountForHolding
}

func (s *SubscriptionExecution6) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	s.UnitsNumber = new(FinancialInstrumentQuantity1)
	return s.UnitsNumber
}

func (s *SubscriptionExecution6) SetNetAmount(value, currency string) {
	s.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution6) SetGrossAmount(value, currency string) {
	s.GrossAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution6) SetRounding(value string) {
	s.Rounding = (*RoundingDirection2Code)(&value)
}

func (s *SubscriptionExecution6) AddTradeDateTime() *DateAndDateTimeChoice {
	s.TradeDateTime = new(DateAndDateTimeChoice)
	return s.TradeDateTime
}

func (s *SubscriptionExecution6) AddDealingPriceDetails() *UnitPrice10 {
	s.DealingPriceDetails = new(UnitPrice10)
	return s.DealingPriceDetails
}

func (s *SubscriptionExecution6) AddInformativePriceDetails() *UnitPrice10 {
	newValue := new(UnitPrice10)
	s.InformativePriceDetails = append(s.InformativePriceDetails, newValue)
	return newValue
}

func (s *SubscriptionExecution6) SetSettlementAmount(value, currency string) {
	s.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution6) SetCashSettlementDate(value string) {
	s.CashSettlementDate = (*ISODate)(&value)
}

func (s *SubscriptionExecution6) SetSettlementMethod(value string) {
	s.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (s *SubscriptionExecution6) SetPartiallyExecutedIndicator(value string) {
	s.PartiallyExecutedIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionExecution6) SetBestExecution(value string) {
	s.BestExecution = (*BestExecution1Code)(&value)
}

func (s *SubscriptionExecution6) SetCumDividendIndicator(value string) {
	s.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionExecution6) AddInterimProfitAmount() *ProfitAndLoss1Choice {
	s.InterimProfitAmount = new(ProfitAndLoss1Choice)
	return s.InterimProfitAmount
}

func (s *SubscriptionExecution6) AddForeignExchangeDetails() *ForeignExchangeTerms7 {
	newValue := new(ForeignExchangeTerms7)
	s.ForeignExchangeDetails = append(s.ForeignExchangeDetails, newValue)
	return newValue
}

func (s *SubscriptionExecution6) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SubscriptionExecution6) SetLetterIntentReference(value string) {
	s.LetterIntentReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution6) SetAccumulationRightReference(value string) {
	s.AccumulationRightReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution6) AddChargeGeneralDetails() *TotalCharges3 {
	s.ChargeGeneralDetails = new(TotalCharges3)
	return s.ChargeGeneralDetails
}

func (s *SubscriptionExecution6) AddCommissionGeneralDetails() *TotalCommissions3 {
	s.CommissionGeneralDetails = new(TotalCommissions3)
	return s.CommissionGeneralDetails
}

func (s *SubscriptionExecution6) AddTaxGeneralDetails() *TotalTaxes3 {
	s.TaxGeneralDetails = new(TotalTaxes3)
	return s.TaxGeneralDetails
}

func (s *SubscriptionExecution6) AddSettlementAndCustodyDetails() *FundSettlementParameters3 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters3)
	return s.SettlementAndCustodyDetails
}

func (s *SubscriptionExecution6) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionExecution6) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	s.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return s.PhysicalDeliveryDetails
}

func (s *SubscriptionExecution6) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (s *SubscriptionExecution6) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (s *SubscriptionExecution6) SetRefund(value, currency string) {
	s.Refund = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution6) SetSubscriptionInterest(value, currency string) {
	s.SubscriptionInterest = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution6) AddCashSettlementDetails() *PaymentTransaction24 {
	s.CashSettlementDetails = new(PaymentTransaction24)
	return s.CashSettlementDetails
}

func (s *SubscriptionExecution6) SetNonStandardSettlementInformation(value string) {
	s.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (s *SubscriptionExecution6) SetPartialSettlementOfUnits(value string) {
	s.PartialSettlementOfUnits = (*PercentageRate)(&value)
}

func (s *SubscriptionExecution6) SetPartialSettlementOfCash(value string) {
	s.PartialSettlementOfCash = (*PercentageRate)(&value)
}

func (s *SubscriptionExecution6) AddStaffClientBreakdown() *InvestmentFundsOrderBreakdown1 {
	newValue := new(InvestmentFundsOrderBreakdown1)
	s.StaffClientBreakdown = append(s.StaffClientBreakdown, newValue)
	return newValue
}

func (s *SubscriptionExecution6) SetFinancialAdvice(value string) {
	s.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (s *SubscriptionExecution6) SetNegotiatedTrade(value string) {
	s.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (s *SubscriptionExecution6) SetLateReport(value string) {
	s.LateReport = (*LateReport1Code)(&value)
}

func (s *SubscriptionExecution6) AddRelatedPartyDetails() *Intermediary9 {
	newValue := new(Intermediary9)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SubscriptionExecution6) AddEqualisation() *Equalisation1 {
	s.Equalisation = new(Equalisation1)
	return s.Equalisation
}
