package iso20022

// Execution of a switch order.
type SwitchExecution3 struct {

	// Date and time at which the order was placed by the investor.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Unique and unambiguous identifier for an order execution, as assigned by a confirming party.
	DealReference *Max35Text `xml:"DealRef"`

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
	InvestmentAccountDetails *InvestmentAccount13 `xml:"InvstmtAcctDtls"`

	// Amount of money used to determine the quantity of investment fund units to be redeemed.
	TotalRedemptionAmount *ActiveCurrencyAndAmount `xml:"TtlRedAmt,omitempty"`

	// Amount of money used to determine the quantity of investment fund units to be subscribed.
	TotalSubscriptionAmount *ActiveCurrencyAndAmount `xml:"TtlSbcptAmt,omitempty"`

	// Additional amount of money paid by the investor in addition to the switch redemption amount.
	AdditionalCashIn *ActiveCurrencyAndAmount `xml:"AddtlCshIn,omitempty"`

	// Amount of money that results from a switch-out, that is not reinvested in another investment fund, and is repaid to the investor.
	ResultingCashOut *ActiveCurrencyAndAmount `xml:"RsltgCshOut,omitempty"`

	// Redemption leg of a switch order execution.
	RedemptionLegDetails []*SwitchRedemptionLegExecution2 `xml:"RedLegDtls"`

	// Subscription leg of a switch order execution.
	SubscriptionLegDetails []*SwitchSubscriptionLegExecution2 `xml:"SbcptLegDtls"`

	// Payment transaction resulting from the investment fund order execution.
	CashSettlementDetails *PaymentTransaction14 `xml:"CshSttlmDtls,omitempty"`

	// Currency exchange related to the execution of an investment fund order.
	ForeignExchangeDetails []*ForeignExchangeTerms4 `xml:"FXDtls,omitempty"`
}

func (s *SwitchExecution3) SetOrderDateTime(value string) {
	s.OrderDateTime = (*ISODateTime)(&value)
}

func (s *SwitchExecution3) SetDealReference(value string) {
	s.DealReference = (*Max35Text)(&value)
}

func (s *SwitchExecution3) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SwitchExecution3) AddInvestmentAccountDetails() *InvestmentAccount13 {
	s.InvestmentAccountDetails = new(InvestmentAccount13)
	return s.InvestmentAccountDetails
}

func (s *SwitchExecution3) SetTotalRedemptionAmount(value, currency string) {
	s.TotalRedemptionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchExecution3) SetTotalSubscriptionAmount(value, currency string) {
	s.TotalSubscriptionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchExecution3) SetAdditionalCashIn(value, currency string) {
	s.AdditionalCashIn = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchExecution3) SetResultingCashOut(value, currency string) {
	s.ResultingCashOut = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchExecution3) AddRedemptionLegDetails() *SwitchRedemptionLegExecution2 {
	newValue := new(SwitchRedemptionLegExecution2)
	s.RedemptionLegDetails = append(s.RedemptionLegDetails, newValue)
	return newValue
}

func (s *SwitchExecution3) AddSubscriptionLegDetails() *SwitchSubscriptionLegExecution2 {
	newValue := new(SwitchSubscriptionLegExecution2)
	s.SubscriptionLegDetails = append(s.SubscriptionLegDetails, newValue)
	return newValue
}

func (s *SwitchExecution3) AddCashSettlementDetails() *PaymentTransaction14 {
	s.CashSettlementDetails = new(PaymentTransaction14)
	return s.CashSettlementDetails
}

func (s *SwitchExecution3) AddForeignExchangeDetails() *ForeignExchangeTerms4 {
	newValue := new(ForeignExchangeTerms4)
	s.ForeignExchangeDetails = append(s.ForeignExchangeDetails, newValue)
	return newValue
}

// Execution of a switch order.
type SwitchExecution4 struct {

	// Unique and unambiguous identifier for a group of individual orders, as assigned by the instructing party. This identifier links the individual orders together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Date and time at which the order was placed by the investor.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
	InvestmentAccountDetails *InvestmentAccount21 `xml:"InvstmtAcctDtls,omitempty"`

	// Unique and unambiguous identifier for an order execution, as assigned by a confirming party.
	DealReference *Max35Text `xml:"DealRef"`

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of an order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Amount of money used to determine the quantity of investment fund units to be redeemed.
	TotalRedemptionAmount *ActiveCurrencyAndAmount `xml:"TtlRedAmt,omitempty"`

	// Amount of money used to determine the quantity of investment fund units to be subscribed.
	TotalSubscriptionAmount *ActiveCurrencyAndAmount `xml:"TtlSbcptAmt,omitempty"`

	// Confirmation of the information about parties related to the transaction.
	RelatedPartyDetails []*Intermediary9 `xml:"RltdPtyDtls,omitempty"`

	// Future date at which the investor requests the order to be executed.
	// The specification of a requested future trade date is not allowed in some markets. The date must be a date in the future.
	RequestedFutureTradeDate *ISODate `xml:"ReqdFutrTradDt,omitempty"`

	// Additional amount of money paid by the investor in addition to the switch redemption amount.
	AdditionalCashIn *ActiveCurrencyAndAmount `xml:"AddtlCshIn,omitempty"`

	// Amount of money that results from a switch-out, that is not reinvested in another investment fund, and is repaid to the investor.
	ResultingCashOut *ActiveCurrencyAndAmount `xml:"RsltgCshOut,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Method by which the transaction is settled.
	SettlementMethod *DeliveryReceiptType2Code `xml:"SttlmMtd,omitempty"`

	// Specifies that the execution was subject to best execution rules as defined by MiFID.
	BestExecution *BestExecution1Code `xml:"BestExctn,omitempty"`

	// Redemption leg of a switch order execution.
	RedemptionLegDetails []*SwitchRedemptionLegExecution3 `xml:"RedLegDtls"`

	// Subscription leg of a switch order execution.
	SubscriptionLegDetails []*SwitchSubscriptionLegExecution3 `xml:"SbcptLegDtls"`

	// Payment transaction resulting from the investment fund order execution.
	CashSettlementDetails *PaymentTransaction26 `xml:"CshSttlmDtls,omitempty"`

	// Currency exchange related to the execution of an investment fund order.
	ForeignExchangeDetails []*ForeignExchangeTerms7 `xml:"FXDtls,omitempty"`

	// Specifies if advice has been received from an independent financial advisor.
	FinancialAdvice *FinancialAdvice1Code `xml:"FinAdvc,omitempty"`

	// Specifies whether the trade is negotiated.
	NegotiatedTrade *NegotiatedTrade1Code `xml:"NgtdTrad,omitempty"`

	// Specifies whether the order execution confirmation is late.
	LateReport *LateReport1Code `xml:"LateRpt,omitempty"`
}

func (s *SwitchExecution4) SetMasterReference(value string) {
	s.MasterReference = (*Max35Text)(&value)
}

func (s *SwitchExecution4) SetOrderDateTime(value string) {
	s.OrderDateTime = (*ISODateTime)(&value)
}

func (s *SwitchExecution4) AddInvestmentAccountDetails() *InvestmentAccount21 {
	s.InvestmentAccountDetails = new(InvestmentAccount21)
	return s.InvestmentAccountDetails
}

func (s *SwitchExecution4) SetDealReference(value string) {
	s.DealReference = (*Max35Text)(&value)
}

func (s *SwitchExecution4) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SwitchExecution4) SetClientReference(value string) {
	s.ClientReference = (*Max35Text)(&value)
}

func (s *SwitchExecution4) SetTotalRedemptionAmount(value, currency string) {
	s.TotalRedemptionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchExecution4) SetTotalSubscriptionAmount(value, currency string) {
	s.TotalSubscriptionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchExecution4) AddRelatedPartyDetails() *Intermediary9 {
	newValue := new(Intermediary9)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SwitchExecution4) SetRequestedFutureTradeDate(value string) {
	s.RequestedFutureTradeDate = (*ISODate)(&value)
}

func (s *SwitchExecution4) SetAdditionalCashIn(value, currency string) {
	s.AdditionalCashIn = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchExecution4) SetResultingCashOut(value, currency string) {
	s.ResultingCashOut = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchExecution4) SetSettlementAmount(value, currency string) {
	s.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchExecution4) SetCashSettlementDate(value string) {
	s.CashSettlementDate = (*ISODate)(&value)
}

func (s *SwitchExecution4) SetSettlementMethod(value string) {
	s.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (s *SwitchExecution4) SetBestExecution(value string) {
	s.BestExecution = (*BestExecution1Code)(&value)
}

func (s *SwitchExecution4) AddRedemptionLegDetails() *SwitchRedemptionLegExecution3 {
	newValue := new(SwitchRedemptionLegExecution3)
	s.RedemptionLegDetails = append(s.RedemptionLegDetails, newValue)
	return newValue
}

func (s *SwitchExecution4) AddSubscriptionLegDetails() *SwitchSubscriptionLegExecution3 {
	newValue := new(SwitchSubscriptionLegExecution3)
	s.SubscriptionLegDetails = append(s.SubscriptionLegDetails, newValue)
	return newValue
}

func (s *SwitchExecution4) AddCashSettlementDetails() *PaymentTransaction26 {
	s.CashSettlementDetails = new(PaymentTransaction26)
	return s.CashSettlementDetails
}

func (s *SwitchExecution4) AddForeignExchangeDetails() *ForeignExchangeTerms7 {
	newValue := new(ForeignExchangeTerms7)
	s.ForeignExchangeDetails = append(s.ForeignExchangeDetails, newValue)
	return newValue
}

func (s *SwitchExecution4) SetFinancialAdvice(value string) {
	s.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (s *SwitchExecution4) SetNegotiatedTrade(value string) {
	s.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (s *SwitchExecution4) SetLateReport(value string) {
	s.LateReport = (*LateReport1Code)(&value)
}

// Execution of a switch order.
type SwitchExecution7 struct {

	// Indicates whether the confirmation is an amendment of a previous confirmation.
	AmendmentIndicator *YesNoIndicator `xml:"AmdmntInd,omitempty"`

	// Unique and unambiguous identifier for a group of individual orders, as assigned by the instructing party. This identifier links the individual orders together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Market in which the advised trade transaction was executed.
	PlaceOfTrade *PlaceOfTradeIdentification1Choice `xml:"PlcOfTrad,omitempty"`

	// Date and time at which the order was placed by the investor or its agent.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Date and time the order was received by the executing party, for example, the transfer agent.
	ReceivedDateTime *ISODateTime `xml:"RcvdDtTm,omitempty"`

	// Unique and unambiguous identifier for the order execution, as assigned by the confirming party.
	DealReference *Max35Text `xml:"DealRef"`

	// Unique and unambiguous identifier for the order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of the order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Account impacted by the investment fund order execution.
	InvestmentAccountDetails *InvestmentAccount58 `xml:"InvstmtAcctDtls,omitempty"`

	// Party related to the transaction.
	RelatedPartyDetails []*Intermediary39 `xml:"RltdPtyDtls,omitempty"`

	// Cancellation right of the investor with respect to the investment fund order.
	CancellationRight *CancellationRight1Choice `xml:"CxlRght,omitempty"`

	// Future date at which the investor requests the order to be executed.
	// The specification of a requested future trade date is not allowed in some markets. The date must be a date in the future.
	RequestedFutureTradeDate *ISODate `xml:"ReqdFutrTradDt,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Method by which the transaction is settled.
	SettlementMethod *DeliveryReceiptType2Code `xml:"SttlmMtd,omitempty"`

	// Choice between additional cash in or resulting cash out.
	AdditionalAmount *AdditionalAmount1Choice `xml:"AddtlAmt,omitempty"`

	// Specifies that the execution was subject to best execution rules as defined by MiFID.
	BestExecution *BestExecution1Code `xml:"BestExctn,omitempty"`

	// Redemption leg of a switch order execution.
	RedemptionLegDetails []*SwitchRedemptionLegExecution4 `xml:"RedLegDtls"`

	// Subscription leg of a switch order execution.
	SubscriptionLegDetails []*SwitchSubscriptionLegExecution4 `xml:"SbcptLegDtls"`

	// Payment process for the transfer of cash from the debtor to the creditor.
	CashSettlementDetails *PaymentTransaction71 `xml:"CshSttlmDtls,omitempty"`

	// Information needed to process a currency exchange or conversion.
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

	// Specifies if advice has been received from an independent financial advisor.
	FinancialAdvice *FinancialAdvice1Code `xml:"FinAdvc,omitempty"`

	// Specifies whether the trade is negotiated.
	NegotiatedTrade *NegotiatedTrade1Code `xml:"NgtdTrad,omitempty"`

	// Specifies whether the order execution confirmation is late.
	LateReport *LateReport1Code `xml:"LateRpt,omitempty"`

	// Assessment of the customer’s behaviour at the time of the account opening application.
	CustomerConductClassification *CustomerConductClassification1Choice `xml:"CstmrCndctClssfctn,omitempty"`

	// Means by which the investor or account owner submits the open account form.
	TransactionChannelType *TransactionChannelType1Choice `xml:"TxChanlTp,omitempty"`

	// Type of signature.
	SignatureType *SignatureType1Choice `xml:"SgntrTp,omitempty"`

	// Information about a non-standard order.
	OrderWaiverDetails *OrderWaiver1 `xml:"OrdrWvrDtls,omitempty"`
}

func (s *SwitchExecution7) SetAmendmentIndicator(value string) {
	s.AmendmentIndicator = (*YesNoIndicator)(&value)
}

func (s *SwitchExecution7) SetMasterReference(value string) {
	s.MasterReference = (*Max35Text)(&value)
}

func (s *SwitchExecution7) AddPlaceOfTrade() *PlaceOfTradeIdentification1Choice {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification1Choice)
	return s.PlaceOfTrade
}

func (s *SwitchExecution7) SetOrderDateTime(value string) {
	s.OrderDateTime = (*ISODateTime)(&value)
}

func (s *SwitchExecution7) SetReceivedDateTime(value string) {
	s.ReceivedDateTime = (*ISODateTime)(&value)
}

func (s *SwitchExecution7) SetDealReference(value string) {
	s.DealReference = (*Max35Text)(&value)
}

func (s *SwitchExecution7) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SwitchExecution7) SetClientReference(value string) {
	s.ClientReference = (*Max35Text)(&value)
}

func (s *SwitchExecution7) AddInvestmentAccountDetails() *InvestmentAccount58 {
	s.InvestmentAccountDetails = new(InvestmentAccount58)
	return s.InvestmentAccountDetails
}

func (s *SwitchExecution7) AddRelatedPartyDetails() *Intermediary39 {
	newValue := new(Intermediary39)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SwitchExecution7) AddCancellationRight() *CancellationRight1Choice {
	s.CancellationRight = new(CancellationRight1Choice)
	return s.CancellationRight
}

func (s *SwitchExecution7) SetRequestedFutureTradeDate(value string) {
	s.RequestedFutureTradeDate = (*ISODate)(&value)
}

func (s *SwitchExecution7) SetSettlementAmount(value, currency string) {
	s.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchExecution7) SetCashSettlementDate(value string) {
	s.CashSettlementDate = (*ISODate)(&value)
}

func (s *SwitchExecution7) SetSettlementMethod(value string) {
	s.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (s *SwitchExecution7) AddAdditionalAmount() *AdditionalAmount1Choice {
	s.AdditionalAmount = new(AdditionalAmount1Choice)
	return s.AdditionalAmount
}

func (s *SwitchExecution7) SetBestExecution(value string) {
	s.BestExecution = (*BestExecution1Code)(&value)
}

func (s *SwitchExecution7) AddRedemptionLegDetails() *SwitchRedemptionLegExecution4 {
	newValue := new(SwitchRedemptionLegExecution4)
	s.RedemptionLegDetails = append(s.RedemptionLegDetails, newValue)
	return newValue
}

func (s *SwitchExecution7) AddSubscriptionLegDetails() *SwitchSubscriptionLegExecution4 {
	newValue := new(SwitchSubscriptionLegExecution4)
	s.SubscriptionLegDetails = append(s.SubscriptionLegDetails, newValue)
	return newValue
}

func (s *SwitchExecution7) AddCashSettlementDetails() *PaymentTransaction71 {
	s.CashSettlementDetails = new(PaymentTransaction71)
	return s.CashSettlementDetails
}

func (s *SwitchExecution7) AddForeignExchangeDetails() *ForeignExchangeTerms33 {
	newValue := new(ForeignExchangeTerms33)
	s.ForeignExchangeDetails = append(s.ForeignExchangeDetails, newValue)
	return newValue
}

func (s *SwitchExecution7) SetFinancialAdvice(value string) {
	s.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (s *SwitchExecution7) SetNegotiatedTrade(value string) {
	s.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (s *SwitchExecution7) SetLateReport(value string) {
	s.LateReport = (*LateReport1Code)(&value)
}

func (s *SwitchExecution7) AddCustomerConductClassification() *CustomerConductClassification1Choice {
	s.CustomerConductClassification = new(CustomerConductClassification1Choice)
	return s.CustomerConductClassification
}

func (s *SwitchExecution7) AddTransactionChannelType() *TransactionChannelType1Choice {
	s.TransactionChannelType = new(TransactionChannelType1Choice)
	return s.TransactionChannelType
}

func (s *SwitchExecution7) AddSignatureType() *SignatureType1Choice {
	s.SignatureType = new(SignatureType1Choice)
	return s.SignatureType
}

func (s *SwitchExecution7) AddOrderWaiverDetails() *OrderWaiver1 {
	s.OrderWaiverDetails = new(OrderWaiver1)
	return s.OrderWaiverDetails
}
