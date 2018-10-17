package iso20022

// Provides details about the settlement obligation.
type SettlementObligation5 struct {

	// Provides the identification of an existing obligation that is linked to the new obligation.
	RelatedSettlementObligationIdentification *Max35Text `xml:"RltdSttlmOblgtnId,omitempty"`

	// Indicates the type of the obligation.
	ObligationType *ObligationType1Choice `xml:"OblgtnTp,omitempty"`

	// Provides additional information related to the linked obligation.
	Description *Max35Text `xml:"Desc,omitempty"`

	// Provides the original trade date.
	TradeDate *ISODate `xml:"TradDt,omitempty"`

	// Specifies the quantity related to the settlement obligation.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty"`

	// Provides the price applied to that net position.
	NetPositionPrice *Price4 `xml:"NetPosPric,omitempty"`

	// Specifies the ISO code of the trade currency.
	TradingCurrency *CurrencyCode `xml:"TradgCcy,omitempty"`

	// Provides the total amount to be settled.
	SettlementAmount *AmountAndDirection27 `xml:"SttlmAmt"`

	// Provides the contractual settlement date.
	SettlementDate *ISODate `xml:"SttlmDt"`

	// Indicates if the obligation will result in a receive or a delivery of securities.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Provides the references of the underlying trade leg(s) and/or the reference to the related net position message.
	References *Reference19 `xml:"Refs,omitempty"`
}

func (s *SettlementObligation5) SetRelatedSettlementObligationIdentification(value string) {
	s.RelatedSettlementObligationIdentification = (*Max35Text)(&value)
}

func (s *SettlementObligation5) AddObligationType() *ObligationType1Choice {
	s.ObligationType = new(ObligationType1Choice)
	return s.ObligationType
}

func (s *SettlementObligation5) SetDescription(value string) {
	s.Description = (*Max35Text)(&value)
}

func (s *SettlementObligation5) SetTradeDate(value string) {
	s.TradeDate = (*ISODate)(&value)
}

func (s *SettlementObligation5) AddQuantity() *FinancialInstrumentQuantity1Choice {
	s.Quantity = new(FinancialInstrumentQuantity1Choice)
	return s.Quantity
}

func (s *SettlementObligation5) AddNetPositionPrice() *Price4 {
	s.NetPositionPrice = new(Price4)
	return s.NetPositionPrice
}

func (s *SettlementObligation5) SetTradingCurrency(value string) {
	s.TradingCurrency = (*CurrencyCode)(&value)
}

func (s *SettlementObligation5) AddSettlementAmount() *AmountAndDirection27 {
	s.SettlementAmount = new(AmountAndDirection27)
	return s.SettlementAmount
}

func (s *SettlementObligation5) SetSettlementDate(value string) {
	s.SettlementDate = (*ISODate)(&value)
}

func (s *SettlementObligation5) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementObligation5) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementObligation5) AddReferences() *Reference19 {
	s.References = new(Reference19)
	return s.References
}

// Provides details about the settlement obligation.
type SettlementObligation7 struct {

	// Last reference given by the settlement platform (this is the central securities depository)  to the transaction (non settled instruction).
	CSDTransactionIdentification *Max35Text `xml:"CSDTxId,omitempty"`

	// Reference of the transaction (non settled instruction) given by the central counterparty.
	CentralCounterpartyTransactionIdentification *Max35Text `xml:"CntrlCtrPtyTxId,omitempty"`

	// Original buy-in identification number in case an event causes a generation of a new buy-in identification.
	PreviousBuyInIdentification *Max35Text `xml:"PrvsBuyInId,omitempty"`

	// An account opened by the central counterparty in the name of the clearing member within the account structure, for settlement purposes (gives information about the clearing member account at central counterparty).
	DeliveryAccount *SecuritiesAccount19 `xml:"DlvryAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally. This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat7Choice `xml:"SfkpgPlc,omitempty"`

	// Clearing member account at the central securities depository.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct,omitempty"`

	// Clearing organisation that will clear the trade.
	//
	// Note: This field allows clearing member firm to segregate flows coming from clearing counterparty's clearing system. Indeed, clearing member firms receive messages from the same system (same sender) and this field allows them to know if the message is related to equities or derivatives.
	ClearingSegment *PartyIdentification35Choice `xml:"ClrSgmt,omitempty"`

	// Provides the identification for the non-clearing member and account.
	NonClearingMember *PartyIdentificationAndAccount31 `xml:"NonClrMmb,omitempty"`

	// Provides the intended settlement date of the position.
	IntendedSettlementDate *ISODate `xml:"IntnddSttlmDt,omitempty"`

	// Provides details about the security identification.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Provides the trade date.
	TradeDate *ISODate `xml:"TradDt,omitempty"`

	// Provides the price of the trade.
	DealPrice *Price4 `xml:"DealPric,omitempty"`

	// Provides the quantity of the trade.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty"`

	// Place where settlement of the securities takes place.
	Depository *PartyIdentification34Choice `xml:"Dpstry,omitempty"`

	// Provides the remaining quantity to be settled.
	RemainingQuantityToBeSettled *FinancialInstrumentQuantity1Choice `xml:"RmngQtyToBeSttld,omitempty"`

	// Provides the amount to be settled.
	SettlementAmount *AmountAndDirection27 `xml:"SttlmAmt"`

	// Provides the remaining amount to be settled.
	RemainingAmountToBeSettled *AmountAndDirection27 `xml:"RmngAmtToBeSttld,omitempty"`
}

func (s *SettlementObligation7) SetCSDTransactionIdentification(value string) {
	s.CSDTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementObligation7) SetCentralCounterpartyTransactionIdentification(value string) {
	s.CentralCounterpartyTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementObligation7) SetPreviousBuyInIdentification(value string) {
	s.PreviousBuyInIdentification = (*Max35Text)(&value)
}

func (s *SettlementObligation7) AddDeliveryAccount() *SecuritiesAccount19 {
	s.DeliveryAccount = new(SecuritiesAccount19)
	return s.DeliveryAccount
}

func (s *SettlementObligation7) AddSafekeepingPlace() *SafekeepingPlaceFormat7Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat7Choice)
	return s.SafekeepingPlace
}

func (s *SettlementObligation7) AddSafekeepingAccount() *SecuritiesAccount19 {
	s.SafekeepingAccount = new(SecuritiesAccount19)
	return s.SafekeepingAccount
}

func (s *SettlementObligation7) AddClearingSegment() *PartyIdentification35Choice {
	s.ClearingSegment = new(PartyIdentification35Choice)
	return s.ClearingSegment
}

func (s *SettlementObligation7) AddNonClearingMember() *PartyIdentificationAndAccount31 {
	s.NonClearingMember = new(PartyIdentificationAndAccount31)
	return s.NonClearingMember
}

func (s *SettlementObligation7) SetIntendedSettlementDate(value string) {
	s.IntendedSettlementDate = (*ISODate)(&value)
}

func (s *SettlementObligation7) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SettlementObligation7) SetTradeDate(value string) {
	s.TradeDate = (*ISODate)(&value)
}

func (s *SettlementObligation7) AddDealPrice() *Price4 {
	s.DealPrice = new(Price4)
	return s.DealPrice
}

func (s *SettlementObligation7) AddQuantity() *FinancialInstrumentQuantity1Choice {
	s.Quantity = new(FinancialInstrumentQuantity1Choice)
	return s.Quantity
}

func (s *SettlementObligation7) AddDepository() *PartyIdentification34Choice {
	s.Depository = new(PartyIdentification34Choice)
	return s.Depository
}

func (s *SettlementObligation7) AddRemainingQuantityToBeSettled() *FinancialInstrumentQuantity1Choice {
	s.RemainingQuantityToBeSettled = new(FinancialInstrumentQuantity1Choice)
	return s.RemainingQuantityToBeSettled
}

func (s *SettlementObligation7) AddSettlementAmount() *AmountAndDirection27 {
	s.SettlementAmount = new(AmountAndDirection27)
	return s.SettlementAmount
}

func (s *SettlementObligation7) AddRemainingAmountToBeSettled() *AmountAndDirection27 {
	s.RemainingAmountToBeSettled = new(AmountAndDirection27)
	return s.RemainingAmountToBeSettled
}

// Provides details about the settlement obligation.
type SettlementObligation8 struct {

	// Provides the identification of the settlement obligation.
	SettlementObligationIdentification *Max35Text `xml:"SttlmOblgtnId"`

	// Provides details about the security identification.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Intended settlement date of the position.
	IntendedSettlementDate *DateFormat11Choice `xml:"IntnddSttlmDt"`

	// Specifies the quantity related to the settlement obligation.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty"`

	// Provides the total amount to be settled.
	SettlementAmount *AmountAndDirection27 `xml:"SttlmAmt"`

	// Place at which the security is traded.
	PlaceOfTrade *MarketIdentification84 `xml:"PlcOfTrad"`

	// Provides the trade date.
	TradeDate *TradeDate3Choice `xml:"TradDt,omitempty"`

	// Identifies the trading capacity of seller.
	TradingCapacity *TradingCapacity5Code `xml:"TradgCpcty,omitempty"`

	// Specifies the clearing account type.
	ClearingAccountType *ClearingAccountType1Code `xml:"ClrAcctTp,omitempty"`

	// Place where the securities are safe-kept, physically or notionally. This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat7Choice `xml:"SfkpgPlc,omitempty"`

	// Clearing member account at the central securities depository.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct,omitempty"`

	// Indicates if the obligation will result in a receive or a delivery of securities.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp,omitempty"`

	// Specifies the amount to be paid under a specific obligation.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Provides details on the parties that are part of the settlement chain.
	SettlementParties *SettlementParties4Choice `xml:"SttlmPties,omitempty"`

	// Provides additional information about the settlement obligation details.
	AdditionalSettlementObligationDetails []*SettlementObligation5 `xml:"AddtlSttlmOblgtnDtls,omitempty"`
}

func (s *SettlementObligation8) SetSettlementObligationIdentification(value string) {
	s.SettlementObligationIdentification = (*Max35Text)(&value)
}

func (s *SettlementObligation8) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SettlementObligation8) AddIntendedSettlementDate() *DateFormat11Choice {
	s.IntendedSettlementDate = new(DateFormat11Choice)
	return s.IntendedSettlementDate
}

func (s *SettlementObligation8) AddQuantity() *FinancialInstrumentQuantity1Choice {
	s.Quantity = new(FinancialInstrumentQuantity1Choice)
	return s.Quantity
}

func (s *SettlementObligation8) AddSettlementAmount() *AmountAndDirection27 {
	s.SettlementAmount = new(AmountAndDirection27)
	return s.SettlementAmount
}

func (s *SettlementObligation8) AddPlaceOfTrade() *MarketIdentification84 {
	s.PlaceOfTrade = new(MarketIdentification84)
	return s.PlaceOfTrade
}

func (s *SettlementObligation8) AddTradeDate() *TradeDate3Choice {
	s.TradeDate = new(TradeDate3Choice)
	return s.TradeDate
}

func (s *SettlementObligation8) SetTradingCapacity(value string) {
	s.TradingCapacity = (*TradingCapacity5Code)(&value)
}

func (s *SettlementObligation8) SetClearingAccountType(value string) {
	s.ClearingAccountType = (*ClearingAccountType1Code)(&value)
}

func (s *SettlementObligation8) AddSafekeepingPlace() *SafekeepingPlaceFormat7Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat7Choice)
	return s.SafekeepingPlace
}

func (s *SettlementObligation8) AddSafekeepingAccount() *SecuritiesAccount19 {
	s.SafekeepingAccount = new(SecuritiesAccount19)
	return s.SafekeepingAccount
}

func (s *SettlementObligation8) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementObligation8) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementObligation8) AddSettlementParties() *SettlementParties4Choice {
	s.SettlementParties = new(SettlementParties4Choice)
	return s.SettlementParties
}

func (s *SettlementObligation8) AddAdditionalSettlementObligationDetails() *SettlementObligation5 {
	newValue := new(SettlementObligation5)
	s.AdditionalSettlementObligationDetails = append(s.AdditionalSettlementObligationDetails, newValue)
	return newValue
}
