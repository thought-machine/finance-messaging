package iso20022

// Instruction from an investor to sell investment fund units back to the fund.
type RedemptionOrder14 struct {

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

	// Amount of money or the number of units or percentage to be redeemed for the redemption order.
	AmountOrUnitsOrPercentage *FinancialInstrumentQuantity28Choice `xml:"AmtOrUnitsOrPctg"`

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

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnit1Code `xml:"Grp1Or2Units,omitempty"`

	// Fees (charges/commission) and tax to be applied to the gross amount.
	TransactionOverhead *FeeAndTax1 `xml:"TxOvrhd,omitempty"`

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

	// Payment process for the transfer of cash from the debtor to the creditor.
	CashSettlementDetails *PaymentTransaction72 `xml:"CshSttlmDtls,omitempty"`

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

	// Part of an investor's retained subscription amount that is returned by the fund in order to reimburse preliminary incentive/performance fees.
	Equalisation *Equalisation1 `xml:"Equlstn,omitempty"`

	// Assessment of the customer’s behaviour at the time of the account opening application.
	CustomerConductClassification *CustomerConductClassification1Choice `xml:"CstmrCndctClssfctn,omitempty"`

	// Means by which the investor or account owner submits the open account form.
	TransactionChannelType *TransactionChannelType1Choice `xml:"TxChanlTp,omitempty"`

	// Type of signature.
	SignatureType *SignatureType1Choice `xml:"SgntrTp,omitempty"`

	// Information about a non-standard order.
	OrderWaiverDetails *OrderWaiver1 `xml:"OrdrWvrDtls,omitempty"`
}

func (r *RedemptionOrder14) SetOrderReference(value string) {
	r.OrderReference = (*Max35Text)(&value)
}

func (r *RedemptionOrder14) SetClientReference(value string) {
	r.ClientReference = (*Max35Text)(&value)
}

func (r *RedemptionOrder14) AddOrderType() *FundOrderType4Choice {
	newValue := new(FundOrderType4Choice)
	r.OrderType = append(r.OrderType, newValue)
	return newValue
}

func (r *RedemptionOrder14) AddFinancialInstrumentDetails() *FinancialInstrument57 {
	r.FinancialInstrumentDetails = new(FinancialInstrument57)
	return r.FinancialInstrumentDetails
}

func (r *RedemptionOrder14) AddSubAccountForHolding() *SubAccount6 {
	r.SubAccountForHolding = new(SubAccount6)
	return r.SubAccountForHolding
}

func (r *RedemptionOrder14) AddAmountOrUnitsOrPercentage() *FinancialInstrumentQuantity28Choice {
	r.AmountOrUnitsOrPercentage = new(FinancialInstrumentQuantity28Choice)
	return r.AmountOrUnitsOrPercentage
}

func (r *RedemptionOrder14) SetRounding(value string) {
	r.Rounding = (*RoundingDirection2Code)(&value)
}

func (r *RedemptionOrder14) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionOrder14) SetCashSettlementDate(value string) {
	r.CashSettlementDate = (*ISODate)(&value)
}

func (r *RedemptionOrder14) SetSettlementMethod(value string) {
	r.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (r *RedemptionOrder14) AddForeignExchangeDetails() *ForeignExchangeTerms32 {
	r.ForeignExchangeDetails = new(ForeignExchangeTerms32)
	return r.ForeignExchangeDetails
}

func (r *RedemptionOrder14) SetIncomePreference(value string) {
	r.IncomePreference = (*IncomePreference1Code)(&value)
}

func (r *RedemptionOrder14) SetGroup1Or2Units(value string) {
	r.Group1Or2Units = (*UKTaxGroupUnit1Code)(&value)
}

func (r *RedemptionOrder14) AddTransactionOverhead() *FeeAndTax1 {
	r.TransactionOverhead = new(FeeAndTax1)
	return r.TransactionOverhead
}

func (r *RedemptionOrder14) AddSettlementAndCustodyDetails() *FundSettlementParameters12 {
	r.SettlementAndCustodyDetails = new(FundSettlementParameters12)
	return r.SettlementAndCustodyDetails
}

func (r *RedemptionOrder14) SetPhysicalDeliveryIndicator(value string) {
	r.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionOrder14) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	r.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return r.PhysicalDeliveryDetails
}

func (r *RedemptionOrder14) SetRequestedSettlementCurrency(value string) {
	r.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (r *RedemptionOrder14) SetRequestedNAVCurrency(value string) {
	r.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (r *RedemptionOrder14) AddCashSettlementDetails() *PaymentTransaction72 {
	r.CashSettlementDetails = new(PaymentTransaction72)
	return r.CashSettlementDetails
}

func (r *RedemptionOrder14) SetNonStandardSettlementInformation(value string) {
	r.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (r *RedemptionOrder14) AddStaffClientBreakdown() *InvestmentFundsOrderBreakdown2 {
	newValue := new(InvestmentFundsOrderBreakdown2)
	r.StaffClientBreakdown = append(r.StaffClientBreakdown, newValue)
	return newValue
}

func (r *RedemptionOrder14) SetFinancialAdvice(value string) {
	r.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (r *RedemptionOrder14) SetNegotiatedTrade(value string) {
	r.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (r *RedemptionOrder14) AddRelatedPartyDetails() *Intermediary40 {
	newValue := new(Intermediary40)
	r.RelatedPartyDetails = append(r.RelatedPartyDetails, newValue)
	return newValue
}

func (r *RedemptionOrder14) AddEqualisation() *Equalisation1 {
	r.Equalisation = new(Equalisation1)
	return r.Equalisation
}

func (r *RedemptionOrder14) AddCustomerConductClassification() *CustomerConductClassification1Choice {
	r.CustomerConductClassification = new(CustomerConductClassification1Choice)
	return r.CustomerConductClassification
}

func (r *RedemptionOrder14) AddTransactionChannelType() *TransactionChannelType1Choice {
	r.TransactionChannelType = new(TransactionChannelType1Choice)
	return r.TransactionChannelType
}

func (r *RedemptionOrder14) AddSignatureType() *SignatureType1Choice {
	r.SignatureType = new(SignatureType1Choice)
	return r.SignatureType
}

func (r *RedemptionOrder14) AddOrderWaiverDetails() *OrderWaiver1 {
	r.OrderWaiverDetails = new(OrderWaiver1)
	return r.OrderWaiverDetails
}

// Instruction from an investor to sell investment fund units back to the fund.
type RedemptionOrder15 struct {

	// Unique and unambiguous identifier for the order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of the order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Account impacted by the investment fund order.
	InvestmentAccountDetails *InvestmentAccount58 `xml:"InvstmtAcctDtls"`

	// Category of the investment fund order.
	OrderType []*FundOrderType4Choice `xml:"OrdrTp,omitempty"`

	// Additional information about the investor.
	BeneficiaryDetails []*IndividualPerson32 `xml:"BnfcryDtls,omitempty"`

	// Amount of money or the number of units or percentage to be redeemed for the redemption order.
	AmountOrUnitsOrPercentage *FinancialInstrumentQuantity28Choice `xml:"AmtOrUnitsOrPctg"`

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

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnit1Code `xml:"Grp1Or2Units,omitempty"`

	// Fees (charges/commission) and tax to be applied to the gross amount.
	TransactionOverhead *FeeAndTax1 `xml:"TxOvrhd,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters12 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to the physical delivery of the securities.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`

	// Payment process for the transfer of cash from the debtor to the creditor.
	CashSettlementDetails *PaymentTransaction72 `xml:"CshSttlmDtls,omitempty"`

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

	// Part of an investor's retained subscription amount that is returned by the fund in order to reimburse preliminary incentive/performance fees.
	Equalisation *Equalisation1 `xml:"Equlstn,omitempty"`

	// Assessment of the customer’s behaviour at the time of the account opening application.
	CustomerConductClassification *CustomerConductClassification1Choice `xml:"CstmrCndctClssfctn,omitempty"`

	// Means by which the investor or account owner submits the open account form.
	TransactionChannelType *TransactionChannelType1Choice `xml:"TxChanlTp,omitempty"`

	// Type of signature.
	SignatureType *SignatureType1Choice `xml:"SgntrTp,omitempty"`

	// Information about a non-standard order.
	OrderWaiverDetails *OrderWaiver1 `xml:"OrdrWvrDtls,omitempty"`
}

func (r *RedemptionOrder15) SetOrderReference(value string) {
	r.OrderReference = (*Max35Text)(&value)
}

func (r *RedemptionOrder15) SetClientReference(value string) {
	r.ClientReference = (*Max35Text)(&value)
}

func (r *RedemptionOrder15) AddInvestmentAccountDetails() *InvestmentAccount58 {
	r.InvestmentAccountDetails = new(InvestmentAccount58)
	return r.InvestmentAccountDetails
}

func (r *RedemptionOrder15) AddOrderType() *FundOrderType4Choice {
	newValue := new(FundOrderType4Choice)
	r.OrderType = append(r.OrderType, newValue)
	return newValue
}

func (r *RedemptionOrder15) AddBeneficiaryDetails() *IndividualPerson32 {
	newValue := new(IndividualPerson32)
	r.BeneficiaryDetails = append(r.BeneficiaryDetails, newValue)
	return newValue
}

func (r *RedemptionOrder15) AddAmountOrUnitsOrPercentage() *FinancialInstrumentQuantity28Choice {
	r.AmountOrUnitsOrPercentage = new(FinancialInstrumentQuantity28Choice)
	return r.AmountOrUnitsOrPercentage
}

func (r *RedemptionOrder15) SetRounding(value string) {
	r.Rounding = (*RoundingDirection2Code)(&value)
}

func (r *RedemptionOrder15) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionOrder15) SetCashSettlementDate(value string) {
	r.CashSettlementDate = (*ISODate)(&value)
}

func (r *RedemptionOrder15) SetSettlementMethod(value string) {
	r.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (r *RedemptionOrder15) AddForeignExchangeDetails() *ForeignExchangeTerms32 {
	r.ForeignExchangeDetails = new(ForeignExchangeTerms32)
	return r.ForeignExchangeDetails
}

func (r *RedemptionOrder15) SetIncomePreference(value string) {
	r.IncomePreference = (*IncomePreference1Code)(&value)
}

func (r *RedemptionOrder15) SetGroup1Or2Units(value string) {
	r.Group1Or2Units = (*UKTaxGroupUnit1Code)(&value)
}

func (r *RedemptionOrder15) AddTransactionOverhead() *FeeAndTax1 {
	r.TransactionOverhead = new(FeeAndTax1)
	return r.TransactionOverhead
}

func (r *RedemptionOrder15) AddSettlementAndCustodyDetails() *FundSettlementParameters12 {
	r.SettlementAndCustodyDetails = new(FundSettlementParameters12)
	return r.SettlementAndCustodyDetails
}

func (r *RedemptionOrder15) SetPhysicalDeliveryIndicator(value string) {
	r.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionOrder15) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	r.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return r.PhysicalDeliveryDetails
}

func (r *RedemptionOrder15) AddCashSettlementDetails() *PaymentTransaction72 {
	r.CashSettlementDetails = new(PaymentTransaction72)
	return r.CashSettlementDetails
}

func (r *RedemptionOrder15) SetNonStandardSettlementInformation(value string) {
	r.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (r *RedemptionOrder15) AddStaffClientBreakdown() *InvestmentFundsOrderBreakdown2 {
	newValue := new(InvestmentFundsOrderBreakdown2)
	r.StaffClientBreakdown = append(r.StaffClientBreakdown, newValue)
	return newValue
}

func (r *RedemptionOrder15) SetFinancialAdvice(value string) {
	r.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (r *RedemptionOrder15) SetNegotiatedTrade(value string) {
	r.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (r *RedemptionOrder15) AddRelatedPartyDetails() *Intermediary40 {
	newValue := new(Intermediary40)
	r.RelatedPartyDetails = append(r.RelatedPartyDetails, newValue)
	return newValue
}

func (r *RedemptionOrder15) AddEqualisation() *Equalisation1 {
	r.Equalisation = new(Equalisation1)
	return r.Equalisation
}

func (r *RedemptionOrder15) AddCustomerConductClassification() *CustomerConductClassification1Choice {
	r.CustomerConductClassification = new(CustomerConductClassification1Choice)
	return r.CustomerConductClassification
}

func (r *RedemptionOrder15) AddTransactionChannelType() *TransactionChannelType1Choice {
	r.TransactionChannelType = new(TransactionChannelType1Choice)
	return r.TransactionChannelType
}

func (r *RedemptionOrder15) AddSignatureType() *SignatureType1Choice {
	r.SignatureType = new(SignatureType1Choice)
	return r.SignatureType
}

func (r *RedemptionOrder15) AddOrderWaiverDetails() *OrderWaiver1 {
	r.OrderWaiverDetails = new(OrderWaiver1)
	return r.OrderWaiverDetails
}

// Instruction from an investor to sell investment fund units back to the fund.
type RedemptionOrder3 struct {

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Specifies the category of the investment fund order.
	OrderType []*FundOrderType1 `xml:"OrdrTp,omitempty"`

	// Account impacted by an investment fund order.
	InvestmentAccountDetails *InvestmentAccount13 `xml:"InvstmtAcctDtls"`

	// Additional information about the investor.
	BeneficiaryDetails *IndividualPerson2 `xml:"BnfcryDtls,omitempty"`

	// Quantity of investment fund units redeemed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Net amount of money used to derive the quantity of investment fund units to be sold.
	NetAmount *ActiveOrHistoricCurrencyAndAmount `xml:"NetAmt"`

	// Portion of the investor's holdings, in a specific investment fund/ fund class, that is redeemed.
	HoldingsRedemptionRate *PercentageRate `xml:"HldgsRedRate"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Amount of money used to derive the quantity of investment fund units to be sold, including all charges, commissions, and tax.
	GrossAmount *ActiveOrHistoricCurrencyAndAmount `xml:"GrssAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms5 `xml:"FXDtls,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnitCode `xml:"Grp1Or2Units,omitempty"`

	// Charge for the placement of an order.
	ChargeDetails []*Charge8 `xml:"ChrgDtls,omitempty"`

	// Commission linked to the execution of an investment fund order.
	CommissionDetails []*Commission6 `xml:"ComssnDtls,omitempty"`

	// Tax applicable to an investment fund order.
	TaxDetails []*Tax6 `xml:"TaxDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters3 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to physical delivery of the securities.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`

	// Payment transaction resulting from the investment fund order execution.
	CashSettlementDetails *PaymentTransaction18 `xml:"CshSttlmDtls,omitempty"`
}

func (r *RedemptionOrder3) SetOrderReference(value string) {
	r.OrderReference = (*Max35Text)(&value)
}

func (r *RedemptionOrder3) AddOrderType() *FundOrderType1 {
	newValue := new(FundOrderType1)
	r.OrderType = append(r.OrderType, newValue)
	return newValue
}

func (r *RedemptionOrder3) AddInvestmentAccountDetails() *InvestmentAccount13 {
	r.InvestmentAccountDetails = new(InvestmentAccount13)
	return r.InvestmentAccountDetails
}

func (r *RedemptionOrder3) AddBeneficiaryDetails() *IndividualPerson2 {
	r.BeneficiaryDetails = new(IndividualPerson2)
	return r.BeneficiaryDetails
}

func (r *RedemptionOrder3) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	r.UnitsNumber = new(FinancialInstrumentQuantity1)
	return r.UnitsNumber
}

func (r *RedemptionOrder3) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RedemptionOrder3) SetHoldingsRedemptionRate(value string) {
	r.HoldingsRedemptionRate = (*PercentageRate)(&value)
}

func (r *RedemptionOrder3) SetRounding(value string) {
	r.Rounding = (*RoundingDirection2Code)(&value)
}

func (r *RedemptionOrder3) SetGrossAmount(value, currency string) {
	r.GrossAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RedemptionOrder3) AddForeignExchangeDetails() *ForeignExchangeTerms5 {
	r.ForeignExchangeDetails = new(ForeignExchangeTerms5)
	return r.ForeignExchangeDetails
}

func (r *RedemptionOrder3) SetIncomePreference(value string) {
	r.IncomePreference = (*IncomePreference1Code)(&value)
}

func (r *RedemptionOrder3) SetGroup1Or2Units(value string) {
	r.Group1Or2Units = (*UKTaxGroupUnitCode)(&value)
}

func (r *RedemptionOrder3) AddChargeDetails() *Charge8 {
	newValue := new(Charge8)
	r.ChargeDetails = append(r.ChargeDetails, newValue)
	return newValue
}

func (r *RedemptionOrder3) AddCommissionDetails() *Commission6 {
	newValue := new(Commission6)
	r.CommissionDetails = append(r.CommissionDetails, newValue)
	return newValue
}

func (r *RedemptionOrder3) AddTaxDetails() *Tax6 {
	newValue := new(Tax6)
	r.TaxDetails = append(r.TaxDetails, newValue)
	return newValue
}

func (r *RedemptionOrder3) AddSettlementAndCustodyDetails() *FundSettlementParameters3 {
	r.SettlementAndCustodyDetails = new(FundSettlementParameters3)
	return r.SettlementAndCustodyDetails
}

func (r *RedemptionOrder3) SetPhysicalDeliveryIndicator(value string) {
	r.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionOrder3) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	r.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return r.PhysicalDeliveryDetails
}

func (r *RedemptionOrder3) AddCashSettlementDetails() *PaymentTransaction18 {
	r.CashSettlementDetails = new(PaymentTransaction18)
	return r.CashSettlementDetails
}

// Instruction from an investor to sell investment fund units back to the fund.
type RedemptionOrder4 struct {

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Specifies the category of the investment fund order.
	OrderType []*FundOrderType1 `xml:"OrdrTp,omitempty"`

	// Investment fund class related to an order.
	FinancialInstrumentDetails *FinancialInstrument6 `xml:"FinInstrmDtls"`

	// Quantity of investment fund units redeemed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Amount of money used to derive the quantity of investment fund units to be sold.
	NetAmount *ActiveOrHistoricCurrencyAndAmount `xml:"NetAmt"`

	// Portion of the investor's holdings, in a specific investment fund/ fund class, that is redeemed.
	HoldingsRedemptionRate *PercentageRate `xml:"HldgsRedRate"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Amount of money used to derive the quantity of investment fund units to be sold, including all charges, commissions, and tax.
	GrossAmount *ActiveOrHistoricCurrencyAndAmount `xml:"GrssAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms5 `xml:"FXDtls,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnitCode `xml:"Grp1Or2Units,omitempty"`

	// Charge for the placement of an order.
	ChargeDetails []*Charge8 `xml:"ChrgDtls,omitempty"`

	// Commission linked to the execution of an investment fund order.
	CommissionDetails []*Commission6 `xml:"ComssnDtls,omitempty"`

	// Tax applicable to an investment fund order.
	TaxDetails []*Tax6 `xml:"TaxDtls,omitempty"`

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

	// Payment transaction resulting from the investment fund order execution.
	CashSettlementDetails *PaymentTransaction15 `xml:"CshSttlmDtls,omitempty"`
}

func (r *RedemptionOrder4) SetOrderReference(value string) {
	r.OrderReference = (*Max35Text)(&value)
}

func (r *RedemptionOrder4) AddOrderType() *FundOrderType1 {
	newValue := new(FundOrderType1)
	r.OrderType = append(r.OrderType, newValue)
	return newValue
}

func (r *RedemptionOrder4) AddFinancialInstrumentDetails() *FinancialInstrument6 {
	r.FinancialInstrumentDetails = new(FinancialInstrument6)
	return r.FinancialInstrumentDetails
}

func (r *RedemptionOrder4) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	r.UnitsNumber = new(FinancialInstrumentQuantity1)
	return r.UnitsNumber
}

func (r *RedemptionOrder4) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RedemptionOrder4) SetHoldingsRedemptionRate(value string) {
	r.HoldingsRedemptionRate = (*PercentageRate)(&value)
}

func (r *RedemptionOrder4) SetRounding(value string) {
	r.Rounding = (*RoundingDirection2Code)(&value)
}

func (r *RedemptionOrder4) SetGrossAmount(value, currency string) {
	r.GrossAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RedemptionOrder4) AddForeignExchangeDetails() *ForeignExchangeTerms5 {
	r.ForeignExchangeDetails = new(ForeignExchangeTerms5)
	return r.ForeignExchangeDetails
}

func (r *RedemptionOrder4) SetIncomePreference(value string) {
	r.IncomePreference = (*IncomePreference1Code)(&value)
}

func (r *RedemptionOrder4) SetGroup1Or2Units(value string) {
	r.Group1Or2Units = (*UKTaxGroupUnitCode)(&value)
}

func (r *RedemptionOrder4) AddChargeDetails() *Charge8 {
	newValue := new(Charge8)
	r.ChargeDetails = append(r.ChargeDetails, newValue)
	return newValue
}

func (r *RedemptionOrder4) AddCommissionDetails() *Commission6 {
	newValue := new(Commission6)
	r.CommissionDetails = append(r.CommissionDetails, newValue)
	return newValue
}

func (r *RedemptionOrder4) AddTaxDetails() *Tax6 {
	newValue := new(Tax6)
	r.TaxDetails = append(r.TaxDetails, newValue)
	return newValue
}

func (r *RedemptionOrder4) AddSettlementAndCustodyDetails() *FundSettlementParameters3 {
	r.SettlementAndCustodyDetails = new(FundSettlementParameters3)
	return r.SettlementAndCustodyDetails
}

func (r *RedemptionOrder4) SetPhysicalDeliveryIndicator(value string) {
	r.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionOrder4) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	r.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return r.PhysicalDeliveryDetails
}

func (r *RedemptionOrder4) SetRequestedSettlementCurrency(value string) {
	r.RequestedSettlementCurrency = (*CurrencyCode)(&value)
}

func (r *RedemptionOrder4) SetRequestedNAVCurrency(value string) {
	r.RequestedNAVCurrency = (*CurrencyCode)(&value)
}

func (r *RedemptionOrder4) AddCashSettlementDetails() *PaymentTransaction15 {
	r.CashSettlementDetails = new(PaymentTransaction15)
	return r.CashSettlementDetails
}

// Instruction from an investor to sell investment fund units back to the fund.
type RedemptionOrder5 struct {

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of an order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for an order cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Account impacted by an investment fund order.
	InvestmentAccountDetails *InvestmentAccount21 `xml:"InvstmtAcctDtls"`

	// Specifies the category of the investment fund order.
	OrderType []*FundOrderType2 `xml:"OrdrTp,omitempty"`

	// Additional information about the investor.
	BeneficiaryDetails *IndividualPerson12 `xml:"BnfcryDtls,omitempty"`

	// Quantity of investment fund units redeemed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Amount of money used to derive the quantity of investment fund units to be sold, before deduction of charges, commissions, and taxes.
	// [Quantity * Price]
	GrossAmount *ActiveOrHistoricCurrencyAndAmount `xml:"GrssAmt"`

	// Amount of money to be received following redemption after deduction of charges, commissions and taxes and used to derive the quantity of investment fund units to be sold.
	// [(Quantity * Price) - (Charges + Commissions +Taxes)]
	NetAmount *ActiveOrHistoricCurrencyAndAmount `xml:"NetAmt"`

	// Portion of the investor's holdings, in a specific investment fund/ fund class, that is redeemed.
	HoldingsRedemptionRate *PercentageRate `xml:"HldgsRedRate"`

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

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnitCode `xml:"Grp1Or2Units,omitempty"`

	// Charge for the placement of an order.
	ChargeDetails []*Charge17 `xml:"ChrgDtls,omitempty"`

	// Amount of money due to a party as compensation for a service.
	CommissionDetails []*Commission10 `xml:"ComssnDtls,omitempty"`

	// Tax applicable to an investment fund order.
	TaxDetails []*Tax16 `xml:"TaxDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters3 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to physical delivery of the securities.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	CashSettlementDetails *PaymentTransaction21 `xml:"CshSttlmDtls,omitempty"`

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

func (r *RedemptionOrder5) SetOrderReference(value string) {
	r.OrderReference = (*Max35Text)(&value)
}

func (r *RedemptionOrder5) SetClientReference(value string) {
	r.ClientReference = (*Max35Text)(&value)
}

func (r *RedemptionOrder5) SetCancellationReference(value string) {
	r.CancellationReference = (*Max35Text)(&value)
}

func (r *RedemptionOrder5) AddInvestmentAccountDetails() *InvestmentAccount21 {
	r.InvestmentAccountDetails = new(InvestmentAccount21)
	return r.InvestmentAccountDetails
}

func (r *RedemptionOrder5) AddOrderType() *FundOrderType2 {
	newValue := new(FundOrderType2)
	r.OrderType = append(r.OrderType, newValue)
	return newValue
}

func (r *RedemptionOrder5) AddBeneficiaryDetails() *IndividualPerson12 {
	r.BeneficiaryDetails = new(IndividualPerson12)
	return r.BeneficiaryDetails
}

func (r *RedemptionOrder5) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	r.UnitsNumber = new(FinancialInstrumentQuantity1)
	return r.UnitsNumber
}

func (r *RedemptionOrder5) SetGrossAmount(value, currency string) {
	r.GrossAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RedemptionOrder5) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RedemptionOrder5) SetHoldingsRedemptionRate(value string) {
	r.HoldingsRedemptionRate = (*PercentageRate)(&value)
}

func (r *RedemptionOrder5) SetRounding(value string) {
	r.Rounding = (*RoundingDirection2Code)(&value)
}

func (r *RedemptionOrder5) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionOrder5) SetCashSettlementDate(value string) {
	r.CashSettlementDate = (*ISODate)(&value)
}

func (r *RedemptionOrder5) SetSettlementMethod(value string) {
	r.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (r *RedemptionOrder5) AddForeignExchangeDetails() *ForeignExchangeTerms6 {
	r.ForeignExchangeDetails = new(ForeignExchangeTerms6)
	return r.ForeignExchangeDetails
}

func (r *RedemptionOrder5) SetIncomePreference(value string) {
	r.IncomePreference = (*IncomePreference1Code)(&value)
}

func (r *RedemptionOrder5) SetGroup1Or2Units(value string) {
	r.Group1Or2Units = (*UKTaxGroupUnitCode)(&value)
}

func (r *RedemptionOrder5) AddChargeDetails() *Charge17 {
	newValue := new(Charge17)
	r.ChargeDetails = append(r.ChargeDetails, newValue)
	return newValue
}

func (r *RedemptionOrder5) AddCommissionDetails() *Commission10 {
	newValue := new(Commission10)
	r.CommissionDetails = append(r.CommissionDetails, newValue)
	return newValue
}

func (r *RedemptionOrder5) AddTaxDetails() *Tax16 {
	newValue := new(Tax16)
	r.TaxDetails = append(r.TaxDetails, newValue)
	return newValue
}

func (r *RedemptionOrder5) AddSettlementAndCustodyDetails() *FundSettlementParameters3 {
	r.SettlementAndCustodyDetails = new(FundSettlementParameters3)
	return r.SettlementAndCustodyDetails
}

func (r *RedemptionOrder5) SetPhysicalDeliveryIndicator(value string) {
	r.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionOrder5) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	r.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return r.PhysicalDeliveryDetails
}

func (r *RedemptionOrder5) AddCashSettlementDetails() *PaymentTransaction21 {
	r.CashSettlementDetails = new(PaymentTransaction21)
	return r.CashSettlementDetails
}

func (r *RedemptionOrder5) SetNonStandardSettlementInformation(value string) {
	r.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (r *RedemptionOrder5) AddStaffClientBreakdown() *InvestmentFundsOrderBreakdown1 {
	newValue := new(InvestmentFundsOrderBreakdown1)
	r.StaffClientBreakdown = append(r.StaffClientBreakdown, newValue)
	return newValue
}

func (r *RedemptionOrder5) SetFinancialAdvice(value string) {
	r.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (r *RedemptionOrder5) SetNegotiatedTrade(value string) {
	r.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (r *RedemptionOrder5) AddRelatedPartyDetails() *Intermediary8 {
	newValue := new(Intermediary8)
	r.RelatedPartyDetails = append(r.RelatedPartyDetails, newValue)
	return newValue
}

func (r *RedemptionOrder5) AddEqualisation() *Equalisation1 {
	r.Equalisation = new(Equalisation1)
	return r.Equalisation
}

// Instruction from an investor to sell investment fund units back to the fund.
type RedemptionOrder6 struct {

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

	// Quantity of investment fund units redeemed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Amount of money used to derive the quantity of investment fund units to be sold, before deduction of charges, commissions, and taxes.
	// [Quantity * Price]
	GrossAmount *ActiveOrHistoricCurrencyAndAmount `xml:"GrssAmt"`

	// Amount of money to be received following redemption after deduction of charges, commissions and taxes and used to derive the quantity of investment fund units to be sold.
	// [(Quantity * Price) - (Charges + Commissions +Taxes)]
	NetAmount *ActiveOrHistoricCurrencyAndAmount `xml:"NetAmt"`

	// Portion of the investor's holdings, in a specific investment fund/ fund class, that is redeemed.
	HoldingsRedemptionRate *PercentageRate `xml:"HldgsRedRate"`

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

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnitCode `xml:"Grp1Or2Units,omitempty"`

	// Charge for the placement of an order.
	ChargeDetails []*Charge17 `xml:"ChrgDtls,omitempty"`

	// Amount of money due to a party as compensation for a service.
	CommissionDetails []*Commission10 `xml:"ComssnDtls,omitempty"`

	// Tax applicable to an investment fund order.
	TaxDetails []*Tax16 `xml:"TaxDtls,omitempty"`

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

	// Payment processes required to transfer cash from the debtor to the creditor.
	CashSettlementDetails *PaymentTransaction21 `xml:"CshSttlmDtls,omitempty"`

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

func (r *RedemptionOrder6) SetOrderReference(value string) {
	r.OrderReference = (*Max35Text)(&value)
}

func (r *RedemptionOrder6) SetClientReference(value string) {
	r.ClientReference = (*Max35Text)(&value)
}

func (r *RedemptionOrder6) SetCancellationReference(value string) {
	r.CancellationReference = (*Max35Text)(&value)
}

func (r *RedemptionOrder6) AddOrderType() *FundOrderType2 {
	newValue := new(FundOrderType2)
	r.OrderType = append(r.OrderType, newValue)
	return newValue
}

func (r *RedemptionOrder6) AddFinancialInstrumentDetails() *FinancialInstrument10 {
	r.FinancialInstrumentDetails = new(FinancialInstrument10)
	return r.FinancialInstrumentDetails
}

func (r *RedemptionOrder6) AddSubAccountForHolding() *SubAccount1 {
	r.SubAccountForHolding = new(SubAccount1)
	return r.SubAccountForHolding
}

func (r *RedemptionOrder6) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	r.UnitsNumber = new(FinancialInstrumentQuantity1)
	return r.UnitsNumber
}

func (r *RedemptionOrder6) SetGrossAmount(value, currency string) {
	r.GrossAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RedemptionOrder6) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RedemptionOrder6) SetHoldingsRedemptionRate(value string) {
	r.HoldingsRedemptionRate = (*PercentageRate)(&value)
}

func (r *RedemptionOrder6) SetRounding(value string) {
	r.Rounding = (*RoundingDirection2Code)(&value)
}

func (r *RedemptionOrder6) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionOrder6) SetCashSettlementDate(value string) {
	r.CashSettlementDate = (*ISODate)(&value)
}

func (r *RedemptionOrder6) SetSettlementMethod(value string) {
	r.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (r *RedemptionOrder6) AddForeignExchangeDetails() *ForeignExchangeTerms6 {
	r.ForeignExchangeDetails = new(ForeignExchangeTerms6)
	return r.ForeignExchangeDetails
}

func (r *RedemptionOrder6) SetIncomePreference(value string) {
	r.IncomePreference = (*IncomePreference1Code)(&value)
}

func (r *RedemptionOrder6) SetGroup1Or2Units(value string) {
	r.Group1Or2Units = (*UKTaxGroupUnitCode)(&value)
}

func (r *RedemptionOrder6) AddChargeDetails() *Charge17 {
	newValue := new(Charge17)
	r.ChargeDetails = append(r.ChargeDetails, newValue)
	return newValue
}

func (r *RedemptionOrder6) AddCommissionDetails() *Commission10 {
	newValue := new(Commission10)
	r.CommissionDetails = append(r.CommissionDetails, newValue)
	return newValue
}

func (r *RedemptionOrder6) AddTaxDetails() *Tax16 {
	newValue := new(Tax16)
	r.TaxDetails = append(r.TaxDetails, newValue)
	return newValue
}

func (r *RedemptionOrder6) AddSettlementAndCustodyDetails() *FundSettlementParameters3 {
	r.SettlementAndCustodyDetails = new(FundSettlementParameters3)
	return r.SettlementAndCustodyDetails
}

func (r *RedemptionOrder6) SetPhysicalDeliveryIndicator(value string) {
	r.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionOrder6) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	r.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return r.PhysicalDeliveryDetails
}

func (r *RedemptionOrder6) SetRequestedSettlementCurrency(value string) {
	r.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (r *RedemptionOrder6) SetRequestedNAVCurrency(value string) {
	r.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (r *RedemptionOrder6) AddCashSettlementDetails() *PaymentTransaction21 {
	r.CashSettlementDetails = new(PaymentTransaction21)
	return r.CashSettlementDetails
}

func (r *RedemptionOrder6) SetNonStandardSettlementInformation(value string) {
	r.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (r *RedemptionOrder6) AddStaffClientBreakdown() *InvestmentFundsOrderBreakdown1 {
	newValue := new(InvestmentFundsOrderBreakdown1)
	r.StaffClientBreakdown = append(r.StaffClientBreakdown, newValue)
	return newValue
}

func (r *RedemptionOrder6) SetFinancialAdvice(value string) {
	r.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (r *RedemptionOrder6) SetNegotiatedTrade(value string) {
	r.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (r *RedemptionOrder6) AddRelatedPartyDetails() *Intermediary8 {
	newValue := new(Intermediary8)
	r.RelatedPartyDetails = append(r.RelatedPartyDetails, newValue)
	return newValue
}

func (r *RedemptionOrder6) AddEqualisation() *Equalisation1 {
	r.Equalisation = new(Equalisation1)
	return r.Equalisation
}

// Instruction from an investor to sell investment fund units back to the fund.
type RedemptionOrder7 struct {

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of an order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Account impacted by an investment fund order.
	InvestmentAccountDetails *InvestmentAccount21 `xml:"InvstmtAcctDtls"`

	// Specifies the category of the investment fund order.
	OrderType []*FundOrderType2 `xml:"OrdrTp,omitempty"`

	// Additional information about the investor.
	BeneficiaryDetails *IndividualPerson12 `xml:"BnfcryDtls,omitempty"`

	// Quantity of investment fund units redeemed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Amount of money used to derive the quantity of investment fund units to be sold, before deduction of charges, commissions, and taxes.
	// [Quantity * Price]
	GrossAmount *ActiveOrHistoricCurrencyAndAmount `xml:"GrssAmt"`

	// Amount of money to be received following redemption after deduction of charges, commissions and taxes and used to derive the quantity of investment fund units to be sold.
	// [(Quantity * Price) - (Charges + Commissions +Taxes)]
	NetAmount *ActiveOrHistoricCurrencyAndAmount `xml:"NetAmt"`

	// Portion of the investor's holdings, in a specific investment fund/ fund class, that is redeemed.
	HoldingsRedemptionRate *PercentageRate `xml:"HldgsRedRate"`

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

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnitCode `xml:"Grp1Or2Units,omitempty"`

	// Charge for the placement of an order.
	ChargeDetails []*Charge17 `xml:"ChrgDtls,omitempty"`

	// Amount of money due to a party as compensation for a service.
	CommissionDetails []*Commission10 `xml:"ComssnDtls,omitempty"`

	// Tax applicable to an investment fund order.
	TaxDetails []*Tax16 `xml:"TaxDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters3 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to physical delivery of the securities.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	CashSettlementDetails *PaymentTransaction21 `xml:"CshSttlmDtls,omitempty"`

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

func (r *RedemptionOrder7) SetOrderReference(value string) {
	r.OrderReference = (*Max35Text)(&value)
}

func (r *RedemptionOrder7) SetClientReference(value string) {
	r.ClientReference = (*Max35Text)(&value)
}

func (r *RedemptionOrder7) AddInvestmentAccountDetails() *InvestmentAccount21 {
	r.InvestmentAccountDetails = new(InvestmentAccount21)
	return r.InvestmentAccountDetails
}

func (r *RedemptionOrder7) AddOrderType() *FundOrderType2 {
	newValue := new(FundOrderType2)
	r.OrderType = append(r.OrderType, newValue)
	return newValue
}

func (r *RedemptionOrder7) AddBeneficiaryDetails() *IndividualPerson12 {
	r.BeneficiaryDetails = new(IndividualPerson12)
	return r.BeneficiaryDetails
}

func (r *RedemptionOrder7) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	r.UnitsNumber = new(FinancialInstrumentQuantity1)
	return r.UnitsNumber
}

func (r *RedemptionOrder7) SetGrossAmount(value, currency string) {
	r.GrossAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RedemptionOrder7) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RedemptionOrder7) SetHoldingsRedemptionRate(value string) {
	r.HoldingsRedemptionRate = (*PercentageRate)(&value)
}

func (r *RedemptionOrder7) SetRounding(value string) {
	r.Rounding = (*RoundingDirection2Code)(&value)
}

func (r *RedemptionOrder7) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionOrder7) SetCashSettlementDate(value string) {
	r.CashSettlementDate = (*ISODate)(&value)
}

func (r *RedemptionOrder7) SetSettlementMethod(value string) {
	r.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (r *RedemptionOrder7) AddForeignExchangeDetails() *ForeignExchangeTerms6 {
	r.ForeignExchangeDetails = new(ForeignExchangeTerms6)
	return r.ForeignExchangeDetails
}

func (r *RedemptionOrder7) SetIncomePreference(value string) {
	r.IncomePreference = (*IncomePreference1Code)(&value)
}

func (r *RedemptionOrder7) SetGroup1Or2Units(value string) {
	r.Group1Or2Units = (*UKTaxGroupUnitCode)(&value)
}

func (r *RedemptionOrder7) AddChargeDetails() *Charge17 {
	newValue := new(Charge17)
	r.ChargeDetails = append(r.ChargeDetails, newValue)
	return newValue
}

func (r *RedemptionOrder7) AddCommissionDetails() *Commission10 {
	newValue := new(Commission10)
	r.CommissionDetails = append(r.CommissionDetails, newValue)
	return newValue
}

func (r *RedemptionOrder7) AddTaxDetails() *Tax16 {
	newValue := new(Tax16)
	r.TaxDetails = append(r.TaxDetails, newValue)
	return newValue
}

func (r *RedemptionOrder7) AddSettlementAndCustodyDetails() *FundSettlementParameters3 {
	r.SettlementAndCustodyDetails = new(FundSettlementParameters3)
	return r.SettlementAndCustodyDetails
}

func (r *RedemptionOrder7) SetPhysicalDeliveryIndicator(value string) {
	r.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionOrder7) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	r.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return r.PhysicalDeliveryDetails
}

func (r *RedemptionOrder7) AddCashSettlementDetails() *PaymentTransaction21 {
	r.CashSettlementDetails = new(PaymentTransaction21)
	return r.CashSettlementDetails
}

func (r *RedemptionOrder7) SetNonStandardSettlementInformation(value string) {
	r.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (r *RedemptionOrder7) AddStaffClientBreakdown() *InvestmentFundsOrderBreakdown1 {
	newValue := new(InvestmentFundsOrderBreakdown1)
	r.StaffClientBreakdown = append(r.StaffClientBreakdown, newValue)
	return newValue
}

func (r *RedemptionOrder7) SetFinancialAdvice(value string) {
	r.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (r *RedemptionOrder7) SetNegotiatedTrade(value string) {
	r.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (r *RedemptionOrder7) AddRelatedPartyDetails() *Intermediary8 {
	newValue := new(Intermediary8)
	r.RelatedPartyDetails = append(r.RelatedPartyDetails, newValue)
	return newValue
}

func (r *RedemptionOrder7) AddEqualisation() *Equalisation1 {
	r.Equalisation = new(Equalisation1)
	return r.Equalisation
}

// Instruction from an investor to sell investment fund units back to the fund.
type RedemptionOrder8 struct {

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

	// Quantity of investment fund units redeemed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Amount of money used to derive the quantity of investment fund units to be sold, before deduction of charges, commissions, and taxes.
	// [Quantity * Price]
	GrossAmount *ActiveOrHistoricCurrencyAndAmount `xml:"GrssAmt"`

	// Amount of money to be received following redemption after deduction of charges, commissions and taxes and used to derive the quantity of investment fund units to be sold.
	// [(Quantity * Price) - (Charges + Commissions +Taxes)]
	NetAmount *ActiveOrHistoricCurrencyAndAmount `xml:"NetAmt"`

	// Portion of the investor's holdings, in a specific investment fund/ fund class, that is redeemed.
	HoldingsRedemptionRate *PercentageRate `xml:"HldgsRedRate"`

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

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnitCode `xml:"Grp1Or2Units,omitempty"`

	// Charge for the placement of an order.
	ChargeDetails []*Charge17 `xml:"ChrgDtls,omitempty"`

	// Amount of money due to a party as compensation for a service.
	CommissionDetails []*Commission10 `xml:"ComssnDtls,omitempty"`

	// Tax applicable to an investment fund order.
	TaxDetails []*Tax16 `xml:"TaxDtls,omitempty"`

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

	// Payment processes required to transfer cash from the debtor to the creditor.
	CashSettlementDetails *PaymentTransaction21 `xml:"CshSttlmDtls,omitempty"`

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

func (r *RedemptionOrder8) SetOrderReference(value string) {
	r.OrderReference = (*Max35Text)(&value)
}

func (r *RedemptionOrder8) SetClientReference(value string) {
	r.ClientReference = (*Max35Text)(&value)
}

func (r *RedemptionOrder8) AddOrderType() *FundOrderType2 {
	newValue := new(FundOrderType2)
	r.OrderType = append(r.OrderType, newValue)
	return newValue
}

func (r *RedemptionOrder8) AddFinancialInstrumentDetails() *FinancialInstrument10 {
	r.FinancialInstrumentDetails = new(FinancialInstrument10)
	return r.FinancialInstrumentDetails
}

func (r *RedemptionOrder8) AddSubAccountForHolding() *SubAccount1 {
	r.SubAccountForHolding = new(SubAccount1)
	return r.SubAccountForHolding
}

func (r *RedemptionOrder8) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	r.UnitsNumber = new(FinancialInstrumentQuantity1)
	return r.UnitsNumber
}

func (r *RedemptionOrder8) SetGrossAmount(value, currency string) {
	r.GrossAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RedemptionOrder8) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RedemptionOrder8) SetHoldingsRedemptionRate(value string) {
	r.HoldingsRedemptionRate = (*PercentageRate)(&value)
}

func (r *RedemptionOrder8) SetRounding(value string) {
	r.Rounding = (*RoundingDirection2Code)(&value)
}

func (r *RedemptionOrder8) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionOrder8) SetCashSettlementDate(value string) {
	r.CashSettlementDate = (*ISODate)(&value)
}

func (r *RedemptionOrder8) SetSettlementMethod(value string) {
	r.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (r *RedemptionOrder8) AddForeignExchangeDetails() *ForeignExchangeTerms6 {
	r.ForeignExchangeDetails = new(ForeignExchangeTerms6)
	return r.ForeignExchangeDetails
}

func (r *RedemptionOrder8) SetIncomePreference(value string) {
	r.IncomePreference = (*IncomePreference1Code)(&value)
}

func (r *RedemptionOrder8) SetGroup1Or2Units(value string) {
	r.Group1Or2Units = (*UKTaxGroupUnitCode)(&value)
}

func (r *RedemptionOrder8) AddChargeDetails() *Charge17 {
	newValue := new(Charge17)
	r.ChargeDetails = append(r.ChargeDetails, newValue)
	return newValue
}

func (r *RedemptionOrder8) AddCommissionDetails() *Commission10 {
	newValue := new(Commission10)
	r.CommissionDetails = append(r.CommissionDetails, newValue)
	return newValue
}

func (r *RedemptionOrder8) AddTaxDetails() *Tax16 {
	newValue := new(Tax16)
	r.TaxDetails = append(r.TaxDetails, newValue)
	return newValue
}

func (r *RedemptionOrder8) AddSettlementAndCustodyDetails() *FundSettlementParameters3 {
	r.SettlementAndCustodyDetails = new(FundSettlementParameters3)
	return r.SettlementAndCustodyDetails
}

func (r *RedemptionOrder8) SetPhysicalDeliveryIndicator(value string) {
	r.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionOrder8) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	r.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return r.PhysicalDeliveryDetails
}

func (r *RedemptionOrder8) SetRequestedSettlementCurrency(value string) {
	r.RequestedSettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (r *RedemptionOrder8) SetRequestedNAVCurrency(value string) {
	r.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (r *RedemptionOrder8) AddCashSettlementDetails() *PaymentTransaction21 {
	r.CashSettlementDetails = new(PaymentTransaction21)
	return r.CashSettlementDetails
}

func (r *RedemptionOrder8) SetNonStandardSettlementInformation(value string) {
	r.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (r *RedemptionOrder8) AddStaffClientBreakdown() *InvestmentFundsOrderBreakdown1 {
	newValue := new(InvestmentFundsOrderBreakdown1)
	r.StaffClientBreakdown = append(r.StaffClientBreakdown, newValue)
	return newValue
}

func (r *RedemptionOrder8) SetFinancialAdvice(value string) {
	r.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (r *RedemptionOrder8) SetNegotiatedTrade(value string) {
	r.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (r *RedemptionOrder8) AddRelatedPartyDetails() *Intermediary8 {
	newValue := new(Intermediary8)
	r.RelatedPartyDetails = append(r.RelatedPartyDetails, newValue)
	return newValue
}

func (r *RedemptionOrder8) AddEqualisation() *Equalisation1 {
	r.Equalisation = new(Equalisation1)
	return r.Equalisation
}
