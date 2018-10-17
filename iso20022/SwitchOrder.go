package iso20022

// Transfer from one investment fund/fund class to another investment fund or investment fund class by the investor. A switch is composed of one or several subscription legs, and one or several redemption legs.
type SwitchOrder2 struct {

	// Date and time at which the order was placed by the investor.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
	InvestmentAccountDetails *InvestmentAccount13 `xml:"InvstmtAcctDtls"`

	// Amount of money used to derive the quantity of investment fund units to be redeemed.
	TotalRedemptionAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlRedAmt,omitempty"`

	// Amount of money used to derive the quantity of investment fund units to be subscribed.
	TotalSubscriptionAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlSbcptAmt,omitempty"`

	// Date on which the order expires.
	ExpiryDateTime *ISODateTime `xml:"XpryDtTm,omitempty"`

	// Additional amount of money paid by the investor in addition to the switch redemption amount.
	AdditionalCashIn *ActiveOrHistoricCurrencyAndAmount `xml:"AddtlCshIn,omitempty"`

	// Amount of money that results from a switch-out, that is not reinvested in another investment fund, and is repaid to the investor.
	ResultingCashOut *ActiveOrHistoricCurrencyAndAmount `xml:"RsltgCshOut,omitempty"`

	// Cancellation right of an investor with respect to an investment fund order.
	CancellationRight *CancellationRight1 `xml:"CxlRght,omitempty"`

	// Part of an investment fund switch order that is a redemption.
	RedemptionLegDetails []*SwitchRedemptionLegOrder2 `xml:"RedLegDtls"`

	// Part of an investment fund switch order that is a subscription.
	SubscriptionLegDetails []*SwitchSubscriptionLegOrder2 `xml:"SbcptLegDtls"`

	// Payment transaction resulting from the investment fund order execution.
	CashSettlementDetails *PaymentTransaction20 `xml:"CshSttlmDtls,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms5 `xml:"FXDtls,omitempty"`
}

func (s *SwitchOrder2) SetOrderDateTime(value string) {
	s.OrderDateTime = (*ISODateTime)(&value)
}

func (s *SwitchOrder2) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SwitchOrder2) AddInvestmentAccountDetails() *InvestmentAccount13 {
	s.InvestmentAccountDetails = new(InvestmentAccount13)
	return s.InvestmentAccountDetails
}

func (s *SwitchOrder2) SetTotalRedemptionAmount(value, currency string) {
	s.TotalRedemptionAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SwitchOrder2) SetTotalSubscriptionAmount(value, currency string) {
	s.TotalSubscriptionAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SwitchOrder2) SetExpiryDateTime(value string) {
	s.ExpiryDateTime = (*ISODateTime)(&value)
}

func (s *SwitchOrder2) SetAdditionalCashIn(value, currency string) {
	s.AdditionalCashIn = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SwitchOrder2) SetResultingCashOut(value, currency string) {
	s.ResultingCashOut = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SwitchOrder2) AddCancellationRight() *CancellationRight1 {
	s.CancellationRight = new(CancellationRight1)
	return s.CancellationRight
}

func (s *SwitchOrder2) AddRedemptionLegDetails() *SwitchRedemptionLegOrder2 {
	newValue := new(SwitchRedemptionLegOrder2)
	s.RedemptionLegDetails = append(s.RedemptionLegDetails, newValue)
	return newValue
}

func (s *SwitchOrder2) AddSubscriptionLegDetails() *SwitchSubscriptionLegOrder2 {
	newValue := new(SwitchSubscriptionLegOrder2)
	s.SubscriptionLegDetails = append(s.SubscriptionLegDetails, newValue)
	return newValue
}

func (s *SwitchOrder2) AddCashSettlementDetails() *PaymentTransaction20 {
	s.CashSettlementDetails = new(PaymentTransaction20)
	return s.CashSettlementDetails
}

func (s *SwitchOrder2) AddForeignExchangeDetails() *ForeignExchangeTerms5 {
	s.ForeignExchangeDetails = new(ForeignExchangeTerms5)
	return s.ForeignExchangeDetails
}

// Transfer from one investment fund/fund class to another investment fund or investment fund class by the investor. A switch is composed of one or several subscription legs, and one or several redemption legs.
type SwitchOrder3 struct {

	// Unique and unambiguous identifier for a group of individual orders, as assigned by the instructing party. This identifier links the individual orders together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Date and time at which the order was placed by the investor.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of an order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for an order cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
	InvestmentAccountDetails *InvestmentAccount21 `xml:"InvstmtAcctDtls,omitempty"`

	// Amount of money used to derive the quantity of investment fund units to be redeemed.
	TotalRedemptionAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlRedAmt,omitempty"`

	// Amount of money used to derive the quantity of investment fund units to be subscribed.
	TotalSubscriptionAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlSbcptAmt,omitempty"`

	// Future date at which the investor requests the order to be executed.
	// The specification of a requested future trade date is not allowed in some markets. The date must be a date in the future.
	RequestedFutureTradeDate *ISODate `xml:"ReqdFutrTradDt,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Method by which the transaction is settled.
	SettlementMethod *DeliveryReceiptType2Code `xml:"SttlmMtd,omitempty"`

	// Date on which the order expires.
	ExpiryDateTime *DateAndDateTimeChoice `xml:"XpryDtTm,omitempty"`

	// Additional amount of money paid by the investor in addition to the switch redemption amount.
	AdditionalCashIn *ActiveOrHistoricCurrencyAndAmount `xml:"AddtlCshIn,omitempty"`

	// Amount of money that results from a switch-out, that is not reinvested in another investment fund, and is repaid to the investor.
	ResultingCashOut *ActiveOrHistoricCurrencyAndAmount `xml:"RsltgCshOut,omitempty"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*Intermediary8 `xml:"RltdPtyDtls,omitempty"`

	// Cancellation right of an investor with respect to an investment fund order.
	CancellationRight *CancellationRight1Code `xml:"CxlRght,omitempty"`

	// Cancellation right of an investor with respect to an investment fund order.
	ExtendedCancellationRight *Extended350Code `xml:"XtndedCxlRght,omitempty"`

	// Part of an investment fund switch order that is a redemption.
	RedemptionLegDetails []*SwitchRedemptionLegOrder3 `xml:"RedLegDtls"`

	// Part of an investment fund switch order that is a subscription.
	SubscriptionLegDetails []*SwitchSubscriptionLegOrder3 `xml:"SbcptLegDtls"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	CashSettlementDetails *PaymentTransaction25 `xml:"CshSttlmDtls,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms6 `xml:"FXDtls,omitempty"`

	// Specifies if advice has been received from an independent financial advisor.
	FinancialAdvice *FinancialAdvice1Code `xml:"FinAdvc,omitempty"`

	// Specifies whether the trade is negotiated.
	NegotiatedTrade *NegotiatedTrade1Code `xml:"NgtdTrad,omitempty"`
}

func (s *SwitchOrder3) SetMasterReference(value string) {
	s.MasterReference = (*Max35Text)(&value)
}

func (s *SwitchOrder3) SetOrderDateTime(value string) {
	s.OrderDateTime = (*ISODateTime)(&value)
}

func (s *SwitchOrder3) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SwitchOrder3) SetClientReference(value string) {
	s.ClientReference = (*Max35Text)(&value)
}

func (s *SwitchOrder3) SetCancellationReference(value string) {
	s.CancellationReference = (*Max35Text)(&value)
}

func (s *SwitchOrder3) AddInvestmentAccountDetails() *InvestmentAccount21 {
	s.InvestmentAccountDetails = new(InvestmentAccount21)
	return s.InvestmentAccountDetails
}

func (s *SwitchOrder3) SetTotalRedemptionAmount(value, currency string) {
	s.TotalRedemptionAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SwitchOrder3) SetTotalSubscriptionAmount(value, currency string) {
	s.TotalSubscriptionAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SwitchOrder3) SetRequestedFutureTradeDate(value string) {
	s.RequestedFutureTradeDate = (*ISODate)(&value)
}

func (s *SwitchOrder3) SetSettlementAmount(value, currency string) {
	s.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchOrder3) SetCashSettlementDate(value string) {
	s.CashSettlementDate = (*ISODate)(&value)
}

func (s *SwitchOrder3) SetSettlementMethod(value string) {
	s.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (s *SwitchOrder3) AddExpiryDateTime() *DateAndDateTimeChoice {
	s.ExpiryDateTime = new(DateAndDateTimeChoice)
	return s.ExpiryDateTime
}

func (s *SwitchOrder3) SetAdditionalCashIn(value, currency string) {
	s.AdditionalCashIn = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SwitchOrder3) SetResultingCashOut(value, currency string) {
	s.ResultingCashOut = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SwitchOrder3) AddRelatedPartyDetails() *Intermediary8 {
	newValue := new(Intermediary8)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SwitchOrder3) SetCancellationRight(value string) {
	s.CancellationRight = (*CancellationRight1Code)(&value)
}

func (s *SwitchOrder3) SetExtendedCancellationRight(value string) {
	s.ExtendedCancellationRight = (*Extended350Code)(&value)
}

func (s *SwitchOrder3) AddRedemptionLegDetails() *SwitchRedemptionLegOrder3 {
	newValue := new(SwitchRedemptionLegOrder3)
	s.RedemptionLegDetails = append(s.RedemptionLegDetails, newValue)
	return newValue
}

func (s *SwitchOrder3) AddSubscriptionLegDetails() *SwitchSubscriptionLegOrder3 {
	newValue := new(SwitchSubscriptionLegOrder3)
	s.SubscriptionLegDetails = append(s.SubscriptionLegDetails, newValue)
	return newValue
}

func (s *SwitchOrder3) AddCashSettlementDetails() *PaymentTransaction25 {
	s.CashSettlementDetails = new(PaymentTransaction25)
	return s.CashSettlementDetails
}

func (s *SwitchOrder3) AddForeignExchangeDetails() *ForeignExchangeTerms6 {
	s.ForeignExchangeDetails = new(ForeignExchangeTerms6)
	return s.ForeignExchangeDetails
}

func (s *SwitchOrder3) SetFinancialAdvice(value string) {
	s.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (s *SwitchOrder3) SetNegotiatedTrade(value string) {
	s.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

// Transfer from one investment fund/fund class to another investment fund or investment fund class by the investor. A switch is composed of one or several subscription legs, and one or several redemption legs.
type SwitchOrder4 struct {

	// Unique and unambiguous identifier for a group of individual orders, as assigned by the instructing party. This identifier links the individual orders together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Date and time at which the order was placed by the investor.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of an order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
	InvestmentAccountDetails *InvestmentAccount21 `xml:"InvstmtAcctDtls,omitempty"`

	// Amount of money used to derive the quantity of investment fund units to be redeemed.
	TotalRedemptionAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlRedAmt,omitempty"`

	// Amount of money used to derive the quantity of investment fund units to be subscribed.
	TotalSubscriptionAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlSbcptAmt,omitempty"`

	// Future date at which the investor requests the order to be executed.
	// The specification of a requested future trade date is not allowed in some markets. The date must be a date in the future.
	RequestedFutureTradeDate *ISODate `xml:"ReqdFutrTradDt,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Method by which the transaction is settled.
	SettlementMethod *DeliveryReceiptType2Code `xml:"SttlmMtd,omitempty"`

	// Date on which the order expires.
	ExpiryDateTime *DateAndDateTimeChoice `xml:"XpryDtTm,omitempty"`

	// Additional amount of money paid by the investor in addition to the switch redemption amount.
	AdditionalCashIn *ActiveOrHistoricCurrencyAndAmount `xml:"AddtlCshIn,omitempty"`

	// Amount of money that results from a switch-out, that is not reinvested in another investment fund, and is repaid to the investor.
	ResultingCashOut *ActiveOrHistoricCurrencyAndAmount `xml:"RsltgCshOut,omitempty"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*Intermediary8 `xml:"RltdPtyDtls,omitempty"`

	// Cancellation right of an investor with respect to an investment fund order.
	CancellationRight *CancellationRight1Code `xml:"CxlRght,omitempty"`

	// Cancellation right of an investor with respect to an investment fund order.
	ExtendedCancellationRight *Extended350Code `xml:"XtndedCxlRght,omitempty"`

	// Part of an investment fund switch order that is a redemption.
	RedemptionLegDetails []*SwitchRedemptionLegOrder3 `xml:"RedLegDtls"`

	// Part of an investment fund switch order that is a subscription.
	SubscriptionLegDetails []*SwitchSubscriptionLegOrder3 `xml:"SbcptLegDtls"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	CashSettlementDetails *PaymentTransaction25 `xml:"CshSttlmDtls,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms6 `xml:"FXDtls,omitempty"`

	// Specifies if advice has been received from an independent financial advisor.
	FinancialAdvice *FinancialAdvice1Code `xml:"FinAdvc,omitempty"`

	// Specifies whether the trade is negotiated.
	NegotiatedTrade *NegotiatedTrade1Code `xml:"NgtdTrad,omitempty"`
}

func (s *SwitchOrder4) SetMasterReference(value string) {
	s.MasterReference = (*Max35Text)(&value)
}

func (s *SwitchOrder4) SetOrderDateTime(value string) {
	s.OrderDateTime = (*ISODateTime)(&value)
}

func (s *SwitchOrder4) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SwitchOrder4) SetClientReference(value string) {
	s.ClientReference = (*Max35Text)(&value)
}

func (s *SwitchOrder4) AddInvestmentAccountDetails() *InvestmentAccount21 {
	s.InvestmentAccountDetails = new(InvestmentAccount21)
	return s.InvestmentAccountDetails
}

func (s *SwitchOrder4) SetTotalRedemptionAmount(value, currency string) {
	s.TotalRedemptionAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SwitchOrder4) SetTotalSubscriptionAmount(value, currency string) {
	s.TotalSubscriptionAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SwitchOrder4) SetRequestedFutureTradeDate(value string) {
	s.RequestedFutureTradeDate = (*ISODate)(&value)
}

func (s *SwitchOrder4) SetSettlementAmount(value, currency string) {
	s.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchOrder4) SetCashSettlementDate(value string) {
	s.CashSettlementDate = (*ISODate)(&value)
}

func (s *SwitchOrder4) SetSettlementMethod(value string) {
	s.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (s *SwitchOrder4) AddExpiryDateTime() *DateAndDateTimeChoice {
	s.ExpiryDateTime = new(DateAndDateTimeChoice)
	return s.ExpiryDateTime
}

func (s *SwitchOrder4) SetAdditionalCashIn(value, currency string) {
	s.AdditionalCashIn = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SwitchOrder4) SetResultingCashOut(value, currency string) {
	s.ResultingCashOut = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SwitchOrder4) AddRelatedPartyDetails() *Intermediary8 {
	newValue := new(Intermediary8)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SwitchOrder4) SetCancellationRight(value string) {
	s.CancellationRight = (*CancellationRight1Code)(&value)
}

func (s *SwitchOrder4) SetExtendedCancellationRight(value string) {
	s.ExtendedCancellationRight = (*Extended350Code)(&value)
}

func (s *SwitchOrder4) AddRedemptionLegDetails() *SwitchRedemptionLegOrder3 {
	newValue := new(SwitchRedemptionLegOrder3)
	s.RedemptionLegDetails = append(s.RedemptionLegDetails, newValue)
	return newValue
}

func (s *SwitchOrder4) AddSubscriptionLegDetails() *SwitchSubscriptionLegOrder3 {
	newValue := new(SwitchSubscriptionLegOrder3)
	s.SubscriptionLegDetails = append(s.SubscriptionLegDetails, newValue)
	return newValue
}

func (s *SwitchOrder4) AddCashSettlementDetails() *PaymentTransaction25 {
	s.CashSettlementDetails = new(PaymentTransaction25)
	return s.CashSettlementDetails
}

func (s *SwitchOrder4) AddForeignExchangeDetails() *ForeignExchangeTerms6 {
	s.ForeignExchangeDetails = new(ForeignExchangeTerms6)
	return s.ForeignExchangeDetails
}

func (s *SwitchOrder4) SetFinancialAdvice(value string) {
	s.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (s *SwitchOrder4) SetNegotiatedTrade(value string) {
	s.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

// Transfer from one investment fund/fund class to another investment fund or investment fund class by the investor. A switch is composed of one or several subscription legs, and one or several redemption legs.
type SwitchOrder7 struct {

	// Unique and unambiguous identifier for a group of individual orders, as assigned by the instructing party. This identifier links the individual orders together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Date and time the order is placed by the investor or its agent.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Market in which the advised trade transaction was executed.
	PlaceOfTrade *PlaceOfTradeIdentification1Choice `xml:"PlcOfTrad,omitempty"`

	// Unique and unambiguous identifier for the order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of the order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Account impacted by the investment fund order.
	InvestmentAccountDetails *InvestmentAccount58 `xml:"InvstmtAcctDtls,omitempty"`

	// Party related to the transaction.
	RelatedPartyDetails []*Intermediary40 `xml:"RltdPtyDtls,omitempty"`

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

	// Date on which the order expires.
	ExpiryDateTime *DateAndDateTimeChoice `xml:"XpryDtTm,omitempty"`

	// Cancellation right of the investor with respect to the investment fund order.
	CancellationRight *CancellationRight1Choice `xml:"CxlRght,omitempty"`

	// Part of the investment fund switch order that is a redemption.
	RedemptionLegDetails []*SwitchRedemptionLegOrder6 `xml:"RedLegDtls"`

	// Part of the investment fund switch order that is a subscription.
	SubscriptionLegDetails []*SwitchSubscriptionLegOrder6 `xml:"SbcptLegDtls"`

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
	ForeignExchangeDetails *ForeignExchangeTerms32 `xml:"FXDtls,omitempty"`

	// Specifies if advice has been received from an independent financial advisor.
	FinancialAdvice *FinancialAdvice1Code `xml:"FinAdvc,omitempty"`

	// Specifies whether the trade is negotiated.
	NegotiatedTrade *NegotiatedTrade1Code `xml:"NgtdTrad,omitempty"`

	// Assessment of the customer’s behaviour at the time of the account opening application.
	CustomerConductClassification *CustomerConductClassification1Choice `xml:"CstmrCndctClssfctn,omitempty"`

	// Means by which the investor or account owner submits the open account form.
	TransactionChannelType *TransactionChannelType1Choice `xml:"TxChanlTp,omitempty"`

	// Type of signature.
	SignatureType *SignatureType1Choice `xml:"SgntrTp,omitempty"`

	// Information about a non-standard order.
	OrderWaiverDetails *OrderWaiver1 `xml:"OrdrWvrDtls,omitempty"`
}

func (s *SwitchOrder7) SetMasterReference(value string) {
	s.MasterReference = (*Max35Text)(&value)
}

func (s *SwitchOrder7) SetOrderDateTime(value string) {
	s.OrderDateTime = (*ISODateTime)(&value)
}

func (s *SwitchOrder7) AddPlaceOfTrade() *PlaceOfTradeIdentification1Choice {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification1Choice)
	return s.PlaceOfTrade
}

func (s *SwitchOrder7) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SwitchOrder7) SetClientReference(value string) {
	s.ClientReference = (*Max35Text)(&value)
}

func (s *SwitchOrder7) AddInvestmentAccountDetails() *InvestmentAccount58 {
	s.InvestmentAccountDetails = new(InvestmentAccount58)
	return s.InvestmentAccountDetails
}

func (s *SwitchOrder7) AddRelatedPartyDetails() *Intermediary40 {
	newValue := new(Intermediary40)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SwitchOrder7) SetRequestedFutureTradeDate(value string) {
	s.RequestedFutureTradeDate = (*ISODate)(&value)
}

func (s *SwitchOrder7) SetSettlementAmount(value, currency string) {
	s.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchOrder7) SetCashSettlementDate(value string) {
	s.CashSettlementDate = (*ISODate)(&value)
}

func (s *SwitchOrder7) SetSettlementMethod(value string) {
	s.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (s *SwitchOrder7) AddAdditionalAmount() *AdditionalAmount1Choice {
	s.AdditionalAmount = new(AdditionalAmount1Choice)
	return s.AdditionalAmount
}

func (s *SwitchOrder7) AddExpiryDateTime() *DateAndDateTimeChoice {
	s.ExpiryDateTime = new(DateAndDateTimeChoice)
	return s.ExpiryDateTime
}

func (s *SwitchOrder7) AddCancellationRight() *CancellationRight1Choice {
	s.CancellationRight = new(CancellationRight1Choice)
	return s.CancellationRight
}

func (s *SwitchOrder7) AddRedemptionLegDetails() *SwitchRedemptionLegOrder6 {
	newValue := new(SwitchRedemptionLegOrder6)
	s.RedemptionLegDetails = append(s.RedemptionLegDetails, newValue)
	return newValue
}

func (s *SwitchOrder7) AddSubscriptionLegDetails() *SwitchSubscriptionLegOrder6 {
	newValue := new(SwitchSubscriptionLegOrder6)
	s.SubscriptionLegDetails = append(s.SubscriptionLegDetails, newValue)
	return newValue
}

func (s *SwitchOrder7) AddCashSettlementDetails() *PaymentTransaction71 {
	s.CashSettlementDetails = new(PaymentTransaction71)
	return s.CashSettlementDetails
}

func (s *SwitchOrder7) AddForeignExchangeDetails() *ForeignExchangeTerms32 {
	s.ForeignExchangeDetails = new(ForeignExchangeTerms32)
	return s.ForeignExchangeDetails
}

func (s *SwitchOrder7) SetFinancialAdvice(value string) {
	s.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (s *SwitchOrder7) SetNegotiatedTrade(value string) {
	s.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (s *SwitchOrder7) AddCustomerConductClassification() *CustomerConductClassification1Choice {
	s.CustomerConductClassification = new(CustomerConductClassification1Choice)
	return s.CustomerConductClassification
}

func (s *SwitchOrder7) AddTransactionChannelType() *TransactionChannelType1Choice {
	s.TransactionChannelType = new(TransactionChannelType1Choice)
	return s.TransactionChannelType
}

func (s *SwitchOrder7) AddSignatureType() *SignatureType1Choice {
	s.SignatureType = new(SignatureType1Choice)
	return s.SignatureType
}

func (s *SwitchOrder7) AddOrderWaiverDetails() *OrderWaiver1 {
	s.OrderWaiverDetails = new(OrderWaiver1)
	return s.OrderWaiverDetails
}
