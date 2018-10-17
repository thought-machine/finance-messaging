package iso20022

// Order to invest the investor's principal in an investment fund.
type SubscriptionOrder14 struct {

	// Unique and unambiguous identifier for the order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of the order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Category of the investment fund order.
	OrderType []*FundOrderType4Choice `xml:"OrdrTp,omitempty"`

	// Investment fund class related to the order.
	FinancialInstrumentDetails *FinancialInstrument57 `xml:"FinInstrmDtls"`

	// Subdivision of the account used to segregate specific holdings.
	SubAccountForHolding *SubAccount6 `xml:"SubAcctForHldg,omitempty"`

	// Amount of money or the number of units for the subscription order.
	AmountOrUnits *FinancialInstrumentQuantity27Choice `xml:"AmtOrUnits"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Method by which the transaction is settled.
	SettlementMethod *DeliveryReceiptType2Code `xml:"SttlmMtd,omitempty"`

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
	ForeignExchangeDetails *ForeignExchangeTerms32 `xml:"FXDtls,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Reference of a letter of intent program, in which sales commissions are reduced based on the aggregate of a customer's actual purchase and anticipated purchases, over a specific period of time, and as agreed by the customer. A letter of intent program is mainly used in the US market.
	LetterIntentReference *Max35Text `xml:"LttrInttRef,omitempty"`

	// Reference of an accumulation right program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Fees (charges/commission) and tax to be applied to the net amount.
	TransactionOverhead *FeeAndTax1 `xml:"TxOvrhd,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters11 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to the physical delivery of the securities.
	PhysicalDeliveryDetails *NameAndAddress4 `xml:"PhysDlvryDtls,omitempty"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *ActiveCurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Payment process for the transfer of cash from the debtor to the creditor.
	CashSettlementDetails *PaymentTransaction70 `xml:"CshSttlmDtls,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Breakdown of the net amount per type of order.
	StaffClientBreakdown []*InvestmentFundsOrderBreakdown2 `xml:"StffClntBrkdwn,omitempty"`

	// Specifies if advice has been received from an independent financial advisor.
	FinancialAdvice *FinancialAdvice1Code `xml:"FinAdvc,omitempty"`

	// Specifies whether the trade is negotiated.
	NegotiatedTrade *NegotiatedTrade1Code `xml:"NgtdTrad,omitempty"`

	// Party related to the transaction.
	RelatedPartyDetails []*Intermediary40 `xml:"RltdPtyDtls,omitempty"`

	// Part of an investor's subscription amount that is held by the fund in order to pay incentive/performance fees at the end of the fiscal year.
	//
	//
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

func (s *SubscriptionOrder14) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder14) SetClientReference(value string) {
	s.ClientReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder14) AddOrderType() *FundOrderType4Choice {
	newValue := new(FundOrderType4Choice)
	s.OrderType = append(s.OrderType, newValue)
	return newValue
}

func (s *SubscriptionOrder14) AddFinancialInstrumentDetails() *FinancialInstrument57 {
	s.FinancialInstrumentDetails = new(FinancialInstrument57)
	return s.FinancialInstrumentDetails
}

func (s *SubscriptionOrder14) AddSubAccountForHolding() *SubAccount6 {
	s.SubAccountForHolding = new(SubAccount6)
	return s.SubAccountForHolding
}

func (s *SubscriptionOrder14) AddAmountOrUnits() *FinancialInstrumentQuantity27Choice {
	s.AmountOrUnits = new(FinancialInstrumentQuantity27Choice)
	return s.AmountOrUnits
}

func (s *SubscriptionOrder14) SetRounding(value string) {
	s.Rounding = (*RoundingDirection2Code)(&value)
}

func (s *SubscriptionOrder14) SetSettlementAmount(value, currency string) {
	s.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionOrder14) SetCashSettlementDate(value string) {
	s.CashSettlementDate = (*ISODate)(&value)
}

func (s *SubscriptionOrder14) SetSettlementMethod(value string) {
	s.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (s *SubscriptionOrder14) AddForeignExchangeDetails() *ForeignExchangeTerms32 {
	s.ForeignExchangeDetails = new(ForeignExchangeTerms32)
	return s.ForeignExchangeDetails
}

func (s *SubscriptionOrder14) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SubscriptionOrder14) SetLetterIntentReference(value string) {
	s.LetterIntentReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder14) SetAccumulationRightReference(value string) {
	s.AccumulationRightReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder14) AddTransactionOverhead() *FeeAndTax1 {
	s.TransactionOverhead = new(FeeAndTax1)
	return s.TransactionOverhead
}

func (s *SubscriptionOrder14) AddSettlementAndCustodyDetails() *FundSettlementParameters11 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters11)
	return s.SettlementAndCustodyDetails
}

func (s *SubscriptionOrder14) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionOrder14) AddPhysicalDeliveryDetails() *NameAndAddress4 {
	s.PhysicalDeliveryDetails = new(NameAndAddress4)
	return s.PhysicalDeliveryDetails
}

func (s *SubscriptionOrder14) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (s *SubscriptionOrder14) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (s *SubscriptionOrder14) AddCashSettlementDetails() *PaymentTransaction70 {
	s.CashSettlementDetails = new(PaymentTransaction70)
	return s.CashSettlementDetails
}

func (s *SubscriptionOrder14) SetNonStandardSettlementInformation(value string) {
	s.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (s *SubscriptionOrder14) AddStaffClientBreakdown() *InvestmentFundsOrderBreakdown2 {
	newValue := new(InvestmentFundsOrderBreakdown2)
	s.StaffClientBreakdown = append(s.StaffClientBreakdown, newValue)
	return newValue
}

func (s *SubscriptionOrder14) SetFinancialAdvice(value string) {
	s.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (s *SubscriptionOrder14) SetNegotiatedTrade(value string) {
	s.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (s *SubscriptionOrder14) AddRelatedPartyDetails() *Intermediary40 {
	newValue := new(Intermediary40)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SubscriptionOrder14) AddEqualisation() *Equalisation1 {
	s.Equalisation = new(Equalisation1)
	return s.Equalisation
}

func (s *SubscriptionOrder14) AddSourceOfCash() *SourceOfCash1Choice {
	newValue := new(SourceOfCash1Choice)
	s.SourceOfCash = append(s.SourceOfCash, newValue)
	return newValue
}

func (s *SubscriptionOrder14) AddCustomerConductClassification() *CustomerConductClassification1Choice {
	s.CustomerConductClassification = new(CustomerConductClassification1Choice)
	return s.CustomerConductClassification
}

func (s *SubscriptionOrder14) AddTransactionChannelType() *TransactionChannelType1Choice {
	s.TransactionChannelType = new(TransactionChannelType1Choice)
	return s.TransactionChannelType
}

func (s *SubscriptionOrder14) AddSignatureType() *SignatureType1Choice {
	s.SignatureType = new(SignatureType1Choice)
	return s.SignatureType
}

func (s *SubscriptionOrder14) AddOrderWaiverDetails() *OrderWaiver1 {
	s.OrderWaiverDetails = new(OrderWaiver1)
	return s.OrderWaiverDetails
}

// Order to invest the investor's principal in an investment fund.
type SubscriptionOrder15 struct {

	// Unique and unambiguous identifier for the order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of the order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Category of the investment fund order.
	OrderType []*FundOrderType4Choice `xml:"OrdrTp,omitempty"`

	// Amount of money or the number of units for the subscription order.
	AmountOrUnits *FinancialInstrumentQuantity27Choice `xml:"AmtOrUnits"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Method by which the transaction is settled.
	SettlementMethod *DeliveryReceiptType2Code `xml:"SttlmMtd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Account impacted by the investment fund order.
	InvestmentAccountDetails *InvestmentAccount58 `xml:"InvstmtAcctDtls"`

	// Additional information about the investor.
	BeneficiaryDetails []*IndividualPerson31 `xml:"BnfcryDtls,omitempty"`

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
	ForeignExchangeDetails *ForeignExchangeTerms32 `xml:"FXDtls,omitempty"`

	// Reference of a letter of intent program, in which sales commissions are reduced based on the aggregate of a customer's actual purchase and anticipated purchases, over a specific period of time, and as agreed by the customer. A letter of intent program is mainly used in the US market.
	LetterIntentReference *Max35Text `xml:"LttrInttRef,omitempty"`

	// Reference of an accumulation right program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Fees (charges/commission) and tax to be applied to the net amount.
	TransactionOverhead *FeeAndTax1 `xml:"TxOvrhd,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters11 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to the physical delivery of the securities.
	PhysicalDeliveryDetails *NameAndAddress4 `xml:"PhysDlvryDtls,omitempty"`

	// Payment process for the transfer of cash from the debtor to the creditor.
	CashSettlementDetails *PaymentTransaction70 `xml:"CshSttlmDtls,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Breakdown of the net amount per type of order.
	StaffClientBreakdown []*InvestmentFundsOrderBreakdown2 `xml:"StffClntBrkdwn,omitempty"`

	// Specifies if advice has been received from an independent financial advisor.
	FinancialAdvice *FinancialAdvice1Code `xml:"FinAdvc,omitempty"`

	// Specifies whether the trade is negotiated.
	NegotiatedTrade *NegotiatedTrade1Code `xml:"NgtdTrad,omitempty"`

	// Party related to the transaction.
	RelatedPartyDetails []*Intermediary40 `xml:"RltdPtyDtls,omitempty"`

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

func (s *SubscriptionOrder15) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder15) SetClientReference(value string) {
	s.ClientReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder15) AddOrderType() *FundOrderType4Choice {
	newValue := new(FundOrderType4Choice)
	s.OrderType = append(s.OrderType, newValue)
	return newValue
}

func (s *SubscriptionOrder15) AddAmountOrUnits() *FinancialInstrumentQuantity27Choice {
	s.AmountOrUnits = new(FinancialInstrumentQuantity27Choice)
	return s.AmountOrUnits
}

func (s *SubscriptionOrder15) SetRounding(value string) {
	s.Rounding = (*RoundingDirection2Code)(&value)
}

func (s *SubscriptionOrder15) SetSettlementAmount(value, currency string) {
	s.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionOrder15) SetCashSettlementDate(value string) {
	s.CashSettlementDate = (*ISODate)(&value)
}

func (s *SubscriptionOrder15) SetSettlementMethod(value string) {
	s.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (s *SubscriptionOrder15) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SubscriptionOrder15) AddInvestmentAccountDetails() *InvestmentAccount58 {
	s.InvestmentAccountDetails = new(InvestmentAccount58)
	return s.InvestmentAccountDetails
}

func (s *SubscriptionOrder15) AddBeneficiaryDetails() *IndividualPerson31 {
	newValue := new(IndividualPerson31)
	s.BeneficiaryDetails = append(s.BeneficiaryDetails, newValue)
	return newValue
}

func (s *SubscriptionOrder15) AddForeignExchangeDetails() *ForeignExchangeTerms32 {
	s.ForeignExchangeDetails = new(ForeignExchangeTerms32)
	return s.ForeignExchangeDetails
}

func (s *SubscriptionOrder15) SetLetterIntentReference(value string) {
	s.LetterIntentReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder15) SetAccumulationRightReference(value string) {
	s.AccumulationRightReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder15) AddTransactionOverhead() *FeeAndTax1 {
	s.TransactionOverhead = new(FeeAndTax1)
	return s.TransactionOverhead
}

func (s *SubscriptionOrder15) AddSettlementAndCustodyDetails() *FundSettlementParameters11 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters11)
	return s.SettlementAndCustodyDetails
}

func (s *SubscriptionOrder15) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionOrder15) AddPhysicalDeliveryDetails() *NameAndAddress4 {
	s.PhysicalDeliveryDetails = new(NameAndAddress4)
	return s.PhysicalDeliveryDetails
}

func (s *SubscriptionOrder15) AddCashSettlementDetails() *PaymentTransaction70 {
	s.CashSettlementDetails = new(PaymentTransaction70)
	return s.CashSettlementDetails
}

func (s *SubscriptionOrder15) SetNonStandardSettlementInformation(value string) {
	s.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (s *SubscriptionOrder15) AddStaffClientBreakdown() *InvestmentFundsOrderBreakdown2 {
	newValue := new(InvestmentFundsOrderBreakdown2)
	s.StaffClientBreakdown = append(s.StaffClientBreakdown, newValue)
	return newValue
}

func (s *SubscriptionOrder15) SetFinancialAdvice(value string) {
	s.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (s *SubscriptionOrder15) SetNegotiatedTrade(value string) {
	s.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (s *SubscriptionOrder15) AddRelatedPartyDetails() *Intermediary40 {
	newValue := new(Intermediary40)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SubscriptionOrder15) AddEqualisation() *Equalisation1 {
	s.Equalisation = new(Equalisation1)
	return s.Equalisation
}

func (s *SubscriptionOrder15) AddSourceOfCash() *SourceOfCash1Choice {
	newValue := new(SourceOfCash1Choice)
	s.SourceOfCash = append(s.SourceOfCash, newValue)
	return newValue
}

func (s *SubscriptionOrder15) AddCustomerConductClassification() *CustomerConductClassification1Choice {
	s.CustomerConductClassification = new(CustomerConductClassification1Choice)
	return s.CustomerConductClassification
}

func (s *SubscriptionOrder15) AddTransactionChannelType() *TransactionChannelType1Choice {
	s.TransactionChannelType = new(TransactionChannelType1Choice)
	return s.TransactionChannelType
}

func (s *SubscriptionOrder15) AddSignatureType() *SignatureType1Choice {
	s.SignatureType = new(SignatureType1Choice)
	return s.SignatureType
}

func (s *SubscriptionOrder15) AddOrderWaiverDetails() *OrderWaiver1 {
	s.OrderWaiverDetails = new(OrderWaiver1)
	return s.OrderWaiverDetails
}

// Order to invest the investor's principal in an investment fund.
type SubscriptionOrder3 struct {

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Specifies the category of the investment fund order.
	OrderType []*FundOrderType1 `xml:"OrdrTp,omitempty"`

	// Account impacted by an investment fund order.
	InvestmentAccountDetails *InvestmentAccount13 `xml:"InvstmtAcctDtls"`

	// Additional information about the investor.
	BeneficiaryDetails *IndividualPerson2 `xml:"BnfcryDtls,omitempty"`

	// Quantity of investment fund units to be subscribed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Amount of money used to determine the quantity of investment fund units to be subscribed.
	NetAmount *ActiveOrHistoricCurrencyAndAmount `xml:"NetAmt"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Amount of money used to determine the quantity of investment fund units to be subscribed, including all charges, commissions, and tax.
	GrossAmount *ActiveOrHistoricCurrencyAndAmount `xml:"GrssAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms5 `xml:"FXDtls,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Reference of a letter of intent program, in which sales commissions are reduced based on the aggregate of a customer's actual purchase and anticipated purchases, over a specific period of time, and as agreed by the customer. A letter of intent program is mainly used in the US market.
	LetterIntentReference *Max35Text `xml:"LttrInttRef,omitempty"`

	// Reference of an accumulation right program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Charge for the placement of an order.
	ChargeDetails []*Charge8 `xml:"ChrgDtls,omitempty"`

	// Commission linked to the execution of an investment fund order.
	CommissionDetails []*Commission6 `xml:"ComssnDtls,omitempty"`

	// Tax applicable to an investment fund order.
	TaxDetails []*Tax6 `xml:"TaxDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters4 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to physical delivery of the securities.
	PhysicalDeliveryDetails *NameAndAddress4 `xml:"PhysDlvryDtls,omitempty"`

	// Payment transaction resulting from the investment fund order execution.
	CashSettlementDetails *PaymentTransaction17 `xml:"CshSttlmDtls,omitempty"`
}

func (s *SubscriptionOrder3) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder3) AddOrderType() *FundOrderType1 {
	newValue := new(FundOrderType1)
	s.OrderType = append(s.OrderType, newValue)
	return newValue
}

func (s *SubscriptionOrder3) AddInvestmentAccountDetails() *InvestmentAccount13 {
	s.InvestmentAccountDetails = new(InvestmentAccount13)
	return s.InvestmentAccountDetails
}

func (s *SubscriptionOrder3) AddBeneficiaryDetails() *IndividualPerson2 {
	s.BeneficiaryDetails = new(IndividualPerson2)
	return s.BeneficiaryDetails
}

func (s *SubscriptionOrder3) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	s.UnitsNumber = new(FinancialInstrumentQuantity1)
	return s.UnitsNumber
}

func (s *SubscriptionOrder3) SetNetAmount(value, currency string) {
	s.NetAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SubscriptionOrder3) SetRounding(value string) {
	s.Rounding = (*RoundingDirection2Code)(&value)
}

func (s *SubscriptionOrder3) SetGrossAmount(value, currency string) {
	s.GrossAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SubscriptionOrder3) AddForeignExchangeDetails() *ForeignExchangeTerms5 {
	s.ForeignExchangeDetails = new(ForeignExchangeTerms5)
	return s.ForeignExchangeDetails
}

func (s *SubscriptionOrder3) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SubscriptionOrder3) SetLetterIntentReference(value string) {
	s.LetterIntentReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder3) SetAccumulationRightReference(value string) {
	s.AccumulationRightReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder3) AddChargeDetails() *Charge8 {
	newValue := new(Charge8)
	s.ChargeDetails = append(s.ChargeDetails, newValue)
	return newValue
}

func (s *SubscriptionOrder3) AddCommissionDetails() *Commission6 {
	newValue := new(Commission6)
	s.CommissionDetails = append(s.CommissionDetails, newValue)
	return newValue
}

func (s *SubscriptionOrder3) AddTaxDetails() *Tax6 {
	newValue := new(Tax6)
	s.TaxDetails = append(s.TaxDetails, newValue)
	return newValue
}

func (s *SubscriptionOrder3) AddSettlementAndCustodyDetails() *FundSettlementParameters4 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters4)
	return s.SettlementAndCustodyDetails
}

func (s *SubscriptionOrder3) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionOrder3) AddPhysicalDeliveryDetails() *NameAndAddress4 {
	s.PhysicalDeliveryDetails = new(NameAndAddress4)
	return s.PhysicalDeliveryDetails
}

func (s *SubscriptionOrder3) AddCashSettlementDetails() *PaymentTransaction17 {
	s.CashSettlementDetails = new(PaymentTransaction17)
	return s.CashSettlementDetails
}

// Order to invest the investor's principal in an investment fund.
type SubscriptionOrder4 struct {

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Specifies the category of the investment fund order.
	OrderType []*FundOrderType1 `xml:"OrdrTp,omitempty"`

	// Investment fund class related to an order.
	FinancialInstrumentDetails *FinancialInstrument6 `xml:"FinInstrmDtls"`

	// Quantity of investment fund units to be subscribed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Amount of money used to determine the quantity of investment fund units to be subscribed.
	NetAmount *ActiveOrHistoricCurrencyAndAmount `xml:"NetAmt"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Amount of money used to determine the quantity of investment fund units to be subscribed, including all charges, commissions, and tax.
	GrossAmount *ActiveOrHistoricCurrencyAndAmount `xml:"GrssAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms5 `xml:"FXDtls,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Reference of a letter of intent program, in which sales commissions are reduced based on the aggregate of a customer's actual purchase and anticipated purchases, over a specific period of time, and as agreed by the customer. A letter of intent program is mainly used in the US market.
	LetterIntentReference *Max35Text `xml:"LttrInttRef,omitempty"`

	// Reference of an accumulation right program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Charge for the placement of an order.
	ChargeDetails []*Charge8 `xml:"ChrgDtls,omitempty"`

	// Commission linked to the execution of an investment fund order.
	CommissionDetails []*Commission6 `xml:"ComssnDtls,omitempty"`

	// Tax applicable to an investment fund order.
	TaxDetails []*Tax6 `xml:"TaxDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters4 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to physical delivery of the securities.
	PhysicalDeliveryDetails *NameAndAddress4 `xml:"PhysDlvryDtls,omitempty"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *CurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *CurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Payment transaction resulting from the investment fund order execution.
	CashSettlementDetails *PaymentTransaction19 `xml:"CshSttlmDtls,omitempty"`
}

func (s *SubscriptionOrder4) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder4) AddOrderType() *FundOrderType1 {
	newValue := new(FundOrderType1)
	s.OrderType = append(s.OrderType, newValue)
	return newValue
}

func (s *SubscriptionOrder4) AddFinancialInstrumentDetails() *FinancialInstrument6 {
	s.FinancialInstrumentDetails = new(FinancialInstrument6)
	return s.FinancialInstrumentDetails
}

func (s *SubscriptionOrder4) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	s.UnitsNumber = new(FinancialInstrumentQuantity1)
	return s.UnitsNumber
}

func (s *SubscriptionOrder4) SetNetAmount(value, currency string) {
	s.NetAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SubscriptionOrder4) SetRounding(value string) {
	s.Rounding = (*RoundingDirection2Code)(&value)
}

func (s *SubscriptionOrder4) SetGrossAmount(value, currency string) {
	s.GrossAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SubscriptionOrder4) AddForeignExchangeDetails() *ForeignExchangeTerms5 {
	s.ForeignExchangeDetails = new(ForeignExchangeTerms5)
	return s.ForeignExchangeDetails
}

func (s *SubscriptionOrder4) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SubscriptionOrder4) SetLetterIntentReference(value string) {
	s.LetterIntentReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder4) SetAccumulationRightReference(value string) {
	s.AccumulationRightReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder4) AddChargeDetails() *Charge8 {
	newValue := new(Charge8)
	s.ChargeDetails = append(s.ChargeDetails, newValue)
	return newValue
}

func (s *SubscriptionOrder4) AddCommissionDetails() *Commission6 {
	newValue := new(Commission6)
	s.CommissionDetails = append(s.CommissionDetails, newValue)
	return newValue
}

func (s *SubscriptionOrder4) AddTaxDetails() *Tax6 {
	newValue := new(Tax6)
	s.TaxDetails = append(s.TaxDetails, newValue)
	return newValue
}

func (s *SubscriptionOrder4) AddSettlementAndCustodyDetails() *FundSettlementParameters4 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters4)
	return s.SettlementAndCustodyDetails
}

func (s *SubscriptionOrder4) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionOrder4) AddPhysicalDeliveryDetails() *NameAndAddress4 {
	s.PhysicalDeliveryDetails = new(NameAndAddress4)
	return s.PhysicalDeliveryDetails
}

func (s *SubscriptionOrder4) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*CurrencyCode)(&value)
}

func (s *SubscriptionOrder4) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*CurrencyCode)(&value)
}

func (s *SubscriptionOrder4) AddCashSettlementDetails() *PaymentTransaction19 {
	s.CashSettlementDetails = new(PaymentTransaction19)
	return s.CashSettlementDetails
}

// Order to invest the investor's principal in an investment fund.
type SubscriptionOrder5 struct {

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of an order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for an order cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Specifies the category of the investment fund order.
	OrderType []*FundOrderType2 `xml:"OrdrTp,omitempty"`

	// Additional information about the investor.
	BeneficiaryDetails *IndividualPerson9 `xml:"BnfcryDtls,omitempty"`

	// Quantity of investment fund units to be subscribed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Amount of money to be invested in a specific financial instrument by an investor before deduction of charges, commissions and taxes and used to determine the quantity of investment fund units to be subscribed.
	// [(Quantity * Price) + (Charges + Commissions +Taxes)]
	GrossAmount *ActiveOrHistoricCurrencyAndAmount `xml:"GrssAmt"`

	// Amount of money remaining after deduction of charges, commissions and taxes and  used to determine the quantity of investment fund units to be subscribed.
	// [Quantity * Price]
	NetAmount *ActiveOrHistoricCurrencyAndAmount `xml:"NetAmt"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Method by which the transaction is settled.
	SettlementMethod *DeliveryReceiptType2Code `xml:"SttlmMtd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Account impacted by an investment fund order.
	InvestmentAccountDetails *InvestmentAccount21 `xml:"InvstmtAcctDtls"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms6 `xml:"FXDtls,omitempty"`

	// Reference of a letter of intent program, in which sales commissions are reduced based on the aggregate of a customer's actual purchase and anticipated purchases, over a specific period of time, and as agreed by the customer. A letter of intent program is mainly used in the US market.
	LetterIntentReference *Max35Text `xml:"LttrInttRef,omitempty"`

	// Reference of an accumulation right program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Charge for the placement of an order.
	ChargeDetails []*Charge17 `xml:"ChrgDtls,omitempty"`

	// Amount of money due to a party as compensation for a service.
	CommissionDetails []*Commission10 `xml:"ComssnDtls,omitempty"`

	// Tax applicable to an investment fund order.
	TaxDetails []*Tax16 `xml:"TaxDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters4 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to physical delivery of the securities.
	PhysicalDeliveryDetails *NameAndAddress4 `xml:"PhysDlvryDtls,omitempty"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	CashSettlementDetails *PaymentTransaction23 `xml:"CshSttlmDtls,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Breakdown of the net amount per type of order.
	StaffClientBreakdown []*InvestmentFundsOrderBreakdown1 `xml:"StffClntBrkdwn,omitempty"`

	// Specifies if advice has been received from an independent financial advisor.
	FinancialAdvice *FinancialAdvice1Code `xml:"FinAdvc,omitempty"`

	// Specifies whether the trade is negotiated.
	NegotiatedTrade *NegotiatedTrade1Code `xml:"NgtdTrad,omitempty"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*Intermediary8 `xml:"RltdPtyDtls,omitempty"`

	// Part of an investor's subscription amount that is held by the fund in order to pay incentive / performance fees at the end of the fiscal year.
	Equalisation *Equalisation1 `xml:"Equlstn,omitempty"`
}

func (s *SubscriptionOrder5) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder5) SetClientReference(value string) {
	s.ClientReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder5) SetCancellationReference(value string) {
	s.CancellationReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder5) AddOrderType() *FundOrderType2 {
	newValue := new(FundOrderType2)
	s.OrderType = append(s.OrderType, newValue)
	return newValue
}

func (s *SubscriptionOrder5) AddBeneficiaryDetails() *IndividualPerson9 {
	s.BeneficiaryDetails = new(IndividualPerson9)
	return s.BeneficiaryDetails
}

func (s *SubscriptionOrder5) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	s.UnitsNumber = new(FinancialInstrumentQuantity1)
	return s.UnitsNumber
}

func (s *SubscriptionOrder5) SetGrossAmount(value, currency string) {
	s.GrossAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SubscriptionOrder5) SetNetAmount(value, currency string) {
	s.NetAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SubscriptionOrder5) SetRounding(value string) {
	s.Rounding = (*RoundingDirection2Code)(&value)
}

func (s *SubscriptionOrder5) SetSettlementAmount(value, currency string) {
	s.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionOrder5) SetCashSettlementDate(value string) {
	s.CashSettlementDate = (*ISODate)(&value)
}

func (s *SubscriptionOrder5) SetSettlementMethod(value string) {
	s.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (s *SubscriptionOrder5) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SubscriptionOrder5) AddInvestmentAccountDetails() *InvestmentAccount21 {
	s.InvestmentAccountDetails = new(InvestmentAccount21)
	return s.InvestmentAccountDetails
}

func (s *SubscriptionOrder5) AddForeignExchangeDetails() *ForeignExchangeTerms6 {
	s.ForeignExchangeDetails = new(ForeignExchangeTerms6)
	return s.ForeignExchangeDetails
}

func (s *SubscriptionOrder5) SetLetterIntentReference(value string) {
	s.LetterIntentReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder5) SetAccumulationRightReference(value string) {
	s.AccumulationRightReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder5) AddChargeDetails() *Charge17 {
	newValue := new(Charge17)
	s.ChargeDetails = append(s.ChargeDetails, newValue)
	return newValue
}

func (s *SubscriptionOrder5) AddCommissionDetails() *Commission10 {
	newValue := new(Commission10)
	s.CommissionDetails = append(s.CommissionDetails, newValue)
	return newValue
}

func (s *SubscriptionOrder5) AddTaxDetails() *Tax16 {
	newValue := new(Tax16)
	s.TaxDetails = append(s.TaxDetails, newValue)
	return newValue
}

func (s *SubscriptionOrder5) AddSettlementAndCustodyDetails() *FundSettlementParameters4 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters4)
	return s.SettlementAndCustodyDetails
}

func (s *SubscriptionOrder5) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionOrder5) AddPhysicalDeliveryDetails() *NameAndAddress4 {
	s.PhysicalDeliveryDetails = new(NameAndAddress4)
	return s.PhysicalDeliveryDetails
}

func (s *SubscriptionOrder5) AddCashSettlementDetails() *PaymentTransaction23 {
	s.CashSettlementDetails = new(PaymentTransaction23)
	return s.CashSettlementDetails
}

func (s *SubscriptionOrder5) SetNonStandardSettlementInformation(value string) {
	s.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (s *SubscriptionOrder5) AddStaffClientBreakdown() *InvestmentFundsOrderBreakdown1 {
	newValue := new(InvestmentFundsOrderBreakdown1)
	s.StaffClientBreakdown = append(s.StaffClientBreakdown, newValue)
	return newValue
}

func (s *SubscriptionOrder5) SetFinancialAdvice(value string) {
	s.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (s *SubscriptionOrder5) SetNegotiatedTrade(value string) {
	s.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (s *SubscriptionOrder5) AddRelatedPartyDetails() *Intermediary8 {
	newValue := new(Intermediary8)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SubscriptionOrder5) AddEqualisation() *Equalisation1 {
	s.Equalisation = new(Equalisation1)
	return s.Equalisation
}

// Order to invest the investor's principal in an investment fund.
type SubscriptionOrder6 struct {

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of an order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for an order cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Specifies the category of the investment fund order.
	OrderType []*FundOrderType2 `xml:"OrdrTp,omitempty"`

	// Investment fund class related to an order.
	FinancialInstrumentDetails *FinancialInstrument10 `xml:"FinInstrmDtls"`

	// Subdivision of an account used to segregate specific holdings.
	SubAccountForHolding *SubAccount1 `xml:"SubAcctForHldg,omitempty"`

	// Quantity of investment fund units to be subscribed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Amount of money to be invested in a specific financial instrument by an investor before deduction of charges, commissions and taxes and used to determine the quantity of investment fund units to be subscribed.
	// [(Quantity * Price) + (Charges + Commissions +Taxes)]
	GrossAmount *ActiveOrHistoricCurrencyAndAmount `xml:"GrssAmt"`

	// Amount of money remaining after deduction of charges, commissions and taxes and  used to determine the quantity of investment fund units to be subscribed.
	// [Quantity * Price]
	NetAmount *ActiveOrHistoricCurrencyAndAmount `xml:"NetAmt"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Method by which the transaction is settled.
	SettlementMethod *DeliveryReceiptType2Code `xml:"SttlmMtd,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms6 `xml:"FXDtls,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Reference of a letter of intent program, in which sales commissions are reduced based on the aggregate of a customer's actual purchase and anticipated purchases, over a specific period of time, and as agreed by the customer. A letter of intent program is mainly used in the US market.
	LetterIntentReference *Max35Text `xml:"LttrInttRef,omitempty"`

	// Reference of an accumulation right program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Charge for the placement of an order.
	ChargeDetails []*Charge17 `xml:"ChrgDtls,omitempty"`

	// Amount of money due to a party as compensation for a service.
	CommissionDetails []*Commission10 `xml:"ComssnDtls,omitempty"`

	// Tax applicable to an investment fund order.
	TaxDetails []*Tax16 `xml:"TaxDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters4 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to physical delivery of the securities.
	PhysicalDeliveryDetails *NameAndAddress4 `xml:"PhysDlvryDtls,omitempty"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *ActiveCurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	CashSettlementDetails *PaymentTransaction23 `xml:"CshSttlmDtls,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Breakdown of the net amount per type of order.
	StaffClientBreakdown []*InvestmentFundsOrderBreakdown1 `xml:"StffClntBrkdwn,omitempty"`

	// Specifies if advice has been received from an independent financial advisor.
	FinancialAdvice *FinancialAdvice1Code `xml:"FinAdvc,omitempty"`

	// Specifies whether the trade is negotiated.
	NegotiatedTrade *NegotiatedTrade1Code `xml:"NgtdTrad,omitempty"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*Intermediary8 `xml:"RltdPtyDtls,omitempty"`

	// Part of an investor's subscription amount that is held by the fund in order to pay incentive / performance fees at the end of the fiscal year.
	Equalisation *Equalisation1 `xml:"Equlstn,omitempty"`
}

func (s *SubscriptionOrder6) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder6) SetClientReference(value string) {
	s.ClientReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder6) SetCancellationReference(value string) {
	s.CancellationReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder6) AddOrderType() *FundOrderType2 {
	newValue := new(FundOrderType2)
	s.OrderType = append(s.OrderType, newValue)
	return newValue
}

func (s *SubscriptionOrder6) AddFinancialInstrumentDetails() *FinancialInstrument10 {
	s.FinancialInstrumentDetails = new(FinancialInstrument10)
	return s.FinancialInstrumentDetails
}

func (s *SubscriptionOrder6) AddSubAccountForHolding() *SubAccount1 {
	s.SubAccountForHolding = new(SubAccount1)
	return s.SubAccountForHolding
}

func (s *SubscriptionOrder6) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	s.UnitsNumber = new(FinancialInstrumentQuantity1)
	return s.UnitsNumber
}

func (s *SubscriptionOrder6) SetGrossAmount(value, currency string) {
	s.GrossAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SubscriptionOrder6) SetNetAmount(value, currency string) {
	s.NetAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SubscriptionOrder6) SetRounding(value string) {
	s.Rounding = (*RoundingDirection2Code)(&value)
}

func (s *SubscriptionOrder6) SetSettlementAmount(value, currency string) {
	s.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionOrder6) SetCashSettlementDate(value string) {
	s.CashSettlementDate = (*ISODate)(&value)
}

func (s *SubscriptionOrder6) SetSettlementMethod(value string) {
	s.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (s *SubscriptionOrder6) AddForeignExchangeDetails() *ForeignExchangeTerms6 {
	s.ForeignExchangeDetails = new(ForeignExchangeTerms6)
	return s.ForeignExchangeDetails
}

func (s *SubscriptionOrder6) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SubscriptionOrder6) SetLetterIntentReference(value string) {
	s.LetterIntentReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder6) SetAccumulationRightReference(value string) {
	s.AccumulationRightReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder6) AddChargeDetails() *Charge17 {
	newValue := new(Charge17)
	s.ChargeDetails = append(s.ChargeDetails, newValue)
	return newValue
}

func (s *SubscriptionOrder6) AddCommissionDetails() *Commission10 {
	newValue := new(Commission10)
	s.CommissionDetails = append(s.CommissionDetails, newValue)
	return newValue
}

func (s *SubscriptionOrder6) AddTaxDetails() *Tax16 {
	newValue := new(Tax16)
	s.TaxDetails = append(s.TaxDetails, newValue)
	return newValue
}

func (s *SubscriptionOrder6) AddSettlementAndCustodyDetails() *FundSettlementParameters4 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters4)
	return s.SettlementAndCustodyDetails
}

func (s *SubscriptionOrder6) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionOrder6) AddPhysicalDeliveryDetails() *NameAndAddress4 {
	s.PhysicalDeliveryDetails = new(NameAndAddress4)
	return s.PhysicalDeliveryDetails
}

func (s *SubscriptionOrder6) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (s *SubscriptionOrder6) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (s *SubscriptionOrder6) AddCashSettlementDetails() *PaymentTransaction23 {
	s.CashSettlementDetails = new(PaymentTransaction23)
	return s.CashSettlementDetails
}

func (s *SubscriptionOrder6) SetNonStandardSettlementInformation(value string) {
	s.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (s *SubscriptionOrder6) AddStaffClientBreakdown() *InvestmentFundsOrderBreakdown1 {
	newValue := new(InvestmentFundsOrderBreakdown1)
	s.StaffClientBreakdown = append(s.StaffClientBreakdown, newValue)
	return newValue
}

func (s *SubscriptionOrder6) SetFinancialAdvice(value string) {
	s.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (s *SubscriptionOrder6) SetNegotiatedTrade(value string) {
	s.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (s *SubscriptionOrder6) AddRelatedPartyDetails() *Intermediary8 {
	newValue := new(Intermediary8)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SubscriptionOrder6) AddEqualisation() *Equalisation1 {
	s.Equalisation = new(Equalisation1)
	return s.Equalisation
}

// Order to invest the investor's principal in an investment fund.
type SubscriptionOrder7 struct {

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of an order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Specifies the category of the investment fund order.
	OrderType []*FundOrderType2 `xml:"OrdrTp,omitempty"`

	// Quantity of investment fund units to be subscribed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Amount of money to be invested in a specific financial instrument by an investor before deduction of charges, commissions and taxes and used to determine the quantity of investment fund units to be subscribed.
	// [(Quantity * Price) + (Charges + Commissions +Taxes)]
	GrossAmount *ActiveOrHistoricCurrencyAndAmount `xml:"GrssAmt"`

	// Amount of money remaining after deduction of charges, commissions and taxes and  used to determine the quantity of investment fund units to be subscribed.
	// [Quantity * Price]
	NetAmount *ActiveOrHistoricCurrencyAndAmount `xml:"NetAmt"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Method by which the transaction is settled.
	SettlementMethod *DeliveryReceiptType2Code `xml:"SttlmMtd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Account impacted by an investment fund order.
	InvestmentAccountDetails *InvestmentAccount21 `xml:"InvstmtAcctDtls"`

	// Additional information about the investor.
	BeneficiaryDetails *IndividualPerson9 `xml:"BnfcryDtls,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms6 `xml:"FXDtls,omitempty"`

	// Reference of a letter of intent program, in which sales commissions are reduced based on the aggregate of a customer's actual purchase and anticipated purchases, over a specific period of time, and as agreed by the customer. A letter of intent program is mainly used in the US market.
	LetterIntentReference *Max35Text `xml:"LttrInttRef,omitempty"`

	// Reference of an accumulation right program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Charge for the placement of an order.
	ChargeDetails []*Charge17 `xml:"ChrgDtls,omitempty"`

	// Amount of money due to a party as compensation for a service.
	CommissionDetails []*Commission10 `xml:"ComssnDtls,omitempty"`

	// Tax applicable to an investment fund order.
	TaxDetails []*Tax16 `xml:"TaxDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters4 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to physical delivery of the securities.
	PhysicalDeliveryDetails *NameAndAddress4 `xml:"PhysDlvryDtls,omitempty"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	CashSettlementDetails *PaymentTransaction23 `xml:"CshSttlmDtls,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Breakdown of the net amount per type of order.
	StaffClientBreakdown []*InvestmentFundsOrderBreakdown1 `xml:"StffClntBrkdwn,omitempty"`

	// Specifies if advice has been received from an independent financial advisor.
	FinancialAdvice *FinancialAdvice1Code `xml:"FinAdvc,omitempty"`

	// Specifies whether the trade is negotiated.
	NegotiatedTrade *NegotiatedTrade1Code `xml:"NgtdTrad,omitempty"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*Intermediary8 `xml:"RltdPtyDtls,omitempty"`

	// Part of an investor's subscription amount that is held by the fund in order to pay incentive / performance fees at the end of the fiscal year.
	Equalisation *Equalisation1 `xml:"Equlstn,omitempty"`
}

func (s *SubscriptionOrder7) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder7) SetClientReference(value string) {
	s.ClientReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder7) AddOrderType() *FundOrderType2 {
	newValue := new(FundOrderType2)
	s.OrderType = append(s.OrderType, newValue)
	return newValue
}

func (s *SubscriptionOrder7) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	s.UnitsNumber = new(FinancialInstrumentQuantity1)
	return s.UnitsNumber
}

func (s *SubscriptionOrder7) SetGrossAmount(value, currency string) {
	s.GrossAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SubscriptionOrder7) SetNetAmount(value, currency string) {
	s.NetAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SubscriptionOrder7) SetRounding(value string) {
	s.Rounding = (*RoundingDirection2Code)(&value)
}

func (s *SubscriptionOrder7) SetSettlementAmount(value, currency string) {
	s.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionOrder7) SetCashSettlementDate(value string) {
	s.CashSettlementDate = (*ISODate)(&value)
}

func (s *SubscriptionOrder7) SetSettlementMethod(value string) {
	s.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (s *SubscriptionOrder7) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SubscriptionOrder7) AddInvestmentAccountDetails() *InvestmentAccount21 {
	s.InvestmentAccountDetails = new(InvestmentAccount21)
	return s.InvestmentAccountDetails
}

func (s *SubscriptionOrder7) AddBeneficiaryDetails() *IndividualPerson9 {
	s.BeneficiaryDetails = new(IndividualPerson9)
	return s.BeneficiaryDetails
}

func (s *SubscriptionOrder7) AddForeignExchangeDetails() *ForeignExchangeTerms6 {
	s.ForeignExchangeDetails = new(ForeignExchangeTerms6)
	return s.ForeignExchangeDetails
}

func (s *SubscriptionOrder7) SetLetterIntentReference(value string) {
	s.LetterIntentReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder7) SetAccumulationRightReference(value string) {
	s.AccumulationRightReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder7) AddChargeDetails() *Charge17 {
	newValue := new(Charge17)
	s.ChargeDetails = append(s.ChargeDetails, newValue)
	return newValue
}

func (s *SubscriptionOrder7) AddCommissionDetails() *Commission10 {
	newValue := new(Commission10)
	s.CommissionDetails = append(s.CommissionDetails, newValue)
	return newValue
}

func (s *SubscriptionOrder7) AddTaxDetails() *Tax16 {
	newValue := new(Tax16)
	s.TaxDetails = append(s.TaxDetails, newValue)
	return newValue
}

func (s *SubscriptionOrder7) AddSettlementAndCustodyDetails() *FundSettlementParameters4 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters4)
	return s.SettlementAndCustodyDetails
}

func (s *SubscriptionOrder7) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionOrder7) AddPhysicalDeliveryDetails() *NameAndAddress4 {
	s.PhysicalDeliveryDetails = new(NameAndAddress4)
	return s.PhysicalDeliveryDetails
}

func (s *SubscriptionOrder7) AddCashSettlementDetails() *PaymentTransaction23 {
	s.CashSettlementDetails = new(PaymentTransaction23)
	return s.CashSettlementDetails
}

func (s *SubscriptionOrder7) SetNonStandardSettlementInformation(value string) {
	s.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (s *SubscriptionOrder7) AddStaffClientBreakdown() *InvestmentFundsOrderBreakdown1 {
	newValue := new(InvestmentFundsOrderBreakdown1)
	s.StaffClientBreakdown = append(s.StaffClientBreakdown, newValue)
	return newValue
}

func (s *SubscriptionOrder7) SetFinancialAdvice(value string) {
	s.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (s *SubscriptionOrder7) SetNegotiatedTrade(value string) {
	s.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (s *SubscriptionOrder7) AddRelatedPartyDetails() *Intermediary8 {
	newValue := new(Intermediary8)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SubscriptionOrder7) AddEqualisation() *Equalisation1 {
	s.Equalisation = new(Equalisation1)
	return s.Equalisation
}

// Order to invest the investor's principal in an investment fund.
type SubscriptionOrder8 struct {

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of an order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Specifies the category of the investment fund order.
	OrderType []*FundOrderType2 `xml:"OrdrTp,omitempty"`

	// Investment fund class related to an order.
	FinancialInstrumentDetails *FinancialInstrument10 `xml:"FinInstrmDtls"`

	// Subdivision of an account used to segregate specific holdings.
	SubAccountForHolding *SubAccount1 `xml:"SubAcctForHldg,omitempty"`

	// Quantity of investment fund units to be subscribed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Amount of money to be invested in a specific financial instrument by an investor before deduction of charges, commissions and taxes and used to determine the quantity of investment fund units to be subscribed.
	// [(Quantity * Price) + (Charges + Commissions +Taxes)]
	GrossAmount *ActiveOrHistoricCurrencyAndAmount `xml:"GrssAmt"`

	// Amount of money remaining after deduction of charges, commissions and taxes and  used to determine the quantity of investment fund units to be subscribed.
	// [Quantity * Price]
	NetAmount *ActiveOrHistoricCurrencyAndAmount `xml:"NetAmt"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Method by which the transaction is settled.
	SettlementMethod *DeliveryReceiptType2Code `xml:"SttlmMtd,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms6 `xml:"FXDtls,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Reference of a letter of intent program, in which sales commissions are reduced based on the aggregate of a customer's actual purchase and anticipated purchases, over a specific period of time, and as agreed by the customer. A letter of intent program is mainly used in the US market.
	LetterIntentReference *Max35Text `xml:"LttrInttRef,omitempty"`

	// Reference of an accumulation right program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Charge for the placement of an order.
	ChargeDetails []*Charge17 `xml:"ChrgDtls,omitempty"`

	// Amount of money due to a party as compensation for a service.
	CommissionDetails []*Commission10 `xml:"ComssnDtls,omitempty"`

	// Tax applicable to an investment fund order.
	TaxDetails []*Tax16 `xml:"TaxDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters4 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to physical delivery of the securities.
	PhysicalDeliveryDetails *NameAndAddress4 `xml:"PhysDlvryDtls,omitempty"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *ActiveCurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	CashSettlementDetails *PaymentTransaction23 `xml:"CshSttlmDtls,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Breakdown of the net amount per type of order.
	StaffClientBreakdown []*InvestmentFundsOrderBreakdown1 `xml:"StffClntBrkdwn,omitempty"`

	// Specifies if advice has been received from an independent financial advisor.
	FinancialAdvice *FinancialAdvice1Code `xml:"FinAdvc,omitempty"`

	// Specifies whether the trade is negotiated.
	NegotiatedTrade *NegotiatedTrade1Code `xml:"NgtdTrad,omitempty"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*Intermediary8 `xml:"RltdPtyDtls,omitempty"`

	// Part of an investor's subscription amount that is held by the fund in order to pay incentive / performance fees at the end of the fiscal year.
	Equalisation *Equalisation1 `xml:"Equlstn,omitempty"`
}

func (s *SubscriptionOrder8) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder8) SetClientReference(value string) {
	s.ClientReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder8) AddOrderType() *FundOrderType2 {
	newValue := new(FundOrderType2)
	s.OrderType = append(s.OrderType, newValue)
	return newValue
}

func (s *SubscriptionOrder8) AddFinancialInstrumentDetails() *FinancialInstrument10 {
	s.FinancialInstrumentDetails = new(FinancialInstrument10)
	return s.FinancialInstrumentDetails
}

func (s *SubscriptionOrder8) AddSubAccountForHolding() *SubAccount1 {
	s.SubAccountForHolding = new(SubAccount1)
	return s.SubAccountForHolding
}

func (s *SubscriptionOrder8) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	s.UnitsNumber = new(FinancialInstrumentQuantity1)
	return s.UnitsNumber
}

func (s *SubscriptionOrder8) SetGrossAmount(value, currency string) {
	s.GrossAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SubscriptionOrder8) SetNetAmount(value, currency string) {
	s.NetAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SubscriptionOrder8) SetRounding(value string) {
	s.Rounding = (*RoundingDirection2Code)(&value)
}

func (s *SubscriptionOrder8) SetSettlementAmount(value, currency string) {
	s.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionOrder8) SetCashSettlementDate(value string) {
	s.CashSettlementDate = (*ISODate)(&value)
}

func (s *SubscriptionOrder8) SetSettlementMethod(value string) {
	s.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (s *SubscriptionOrder8) AddForeignExchangeDetails() *ForeignExchangeTerms6 {
	s.ForeignExchangeDetails = new(ForeignExchangeTerms6)
	return s.ForeignExchangeDetails
}

func (s *SubscriptionOrder8) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SubscriptionOrder8) SetLetterIntentReference(value string) {
	s.LetterIntentReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder8) SetAccumulationRightReference(value string) {
	s.AccumulationRightReference = (*Max35Text)(&value)
}

func (s *SubscriptionOrder8) AddChargeDetails() *Charge17 {
	newValue := new(Charge17)
	s.ChargeDetails = append(s.ChargeDetails, newValue)
	return newValue
}

func (s *SubscriptionOrder8) AddCommissionDetails() *Commission10 {
	newValue := new(Commission10)
	s.CommissionDetails = append(s.CommissionDetails, newValue)
	return newValue
}

func (s *SubscriptionOrder8) AddTaxDetails() *Tax16 {
	newValue := new(Tax16)
	s.TaxDetails = append(s.TaxDetails, newValue)
	return newValue
}

func (s *SubscriptionOrder8) AddSettlementAndCustodyDetails() *FundSettlementParameters4 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters4)
	return s.SettlementAndCustodyDetails
}

func (s *SubscriptionOrder8) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionOrder8) AddPhysicalDeliveryDetails() *NameAndAddress4 {
	s.PhysicalDeliveryDetails = new(NameAndAddress4)
	return s.PhysicalDeliveryDetails
}

func (s *SubscriptionOrder8) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (s *SubscriptionOrder8) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (s *SubscriptionOrder8) AddCashSettlementDetails() *PaymentTransaction23 {
	s.CashSettlementDetails = new(PaymentTransaction23)
	return s.CashSettlementDetails
}

func (s *SubscriptionOrder8) SetNonStandardSettlementInformation(value string) {
	s.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (s *SubscriptionOrder8) AddStaffClientBreakdown() *InvestmentFundsOrderBreakdown1 {
	newValue := new(InvestmentFundsOrderBreakdown1)
	s.StaffClientBreakdown = append(s.StaffClientBreakdown, newValue)
	return newValue
}

func (s *SubscriptionOrder8) SetFinancialAdvice(value string) {
	s.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (s *SubscriptionOrder8) SetNegotiatedTrade(value string) {
	s.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (s *SubscriptionOrder8) AddRelatedPartyDetails() *Intermediary8 {
	newValue := new(Intermediary8)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SubscriptionOrder8) AddEqualisation() *Equalisation1 {
	s.Equalisation = new(Equalisation1)
	return s.Equalisation
}
