package iso20022

// Details of the securities trade.
type SecuritiesTradeDetails1 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*Max35Text `xml:"CollTxId,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Specifies additional information relative to the processing of the trade.
	OpeningClosing *OpeningClosing1Choice `xml:"OpngClsg,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting2Choice `xml:"Rptg,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition1Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity1Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator1Choice `xml:"TradOrgtrRole,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice3Choice `xml:"TpOfPric,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds or to sell the indicated currency in order to obtain the necessary currency to fund the transaction.
	CurrencyToBuyOrSell *CurrencyToBuyOrSell1Choice `xml:"CcyToBuyOrSell,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus1Choice `xml:"MtchgSts,omitempty"`

	// Status of affirmation of a trade.
	AffirmationStatus *AffirmationStatus1Choice `xml:"AffirmSts,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *Max350Text `xml:"FxAddtlDtls,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails1) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails1) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails1) AddPlaceOfTrade() *MarketIdentification4 {
	s.PlaceOfTrade = new(MarketIdentification4)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails1) SetPlaceOfClearing(value string) {
	s.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (s *SecuritiesTradeDetails1) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails1) AddSettlementDate() *SettlementDate1Choice {
	s.SettlementDate = new(SettlementDate1Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails1) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesTradeDetails1) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails1) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails1) AddOpeningClosing() *OpeningClosing1Choice {
	s.OpeningClosing = new(OpeningClosing1Choice)
	return s.OpeningClosing
}

func (s *SecuritiesTradeDetails1) AddReporting() *Reporting2Choice {
	newValue := new(Reporting2Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails1) AddTradeTransactionCondition() *TradeTransactionCondition1Choice {
	newValue := new(TradeTransactionCondition1Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails1) AddInvestorCapacity() *InvestorCapacity1Choice {
	s.InvestorCapacity = new(InvestorCapacity1Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails1) AddTradeOriginatorRole() *TradeOriginator1Choice {
	s.TradeOriginatorRole = new(TradeOriginator1Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails1) AddTypeOfPrice() *TypeOfPrice3Choice {
	s.TypeOfPrice = new(TypeOfPrice3Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails1) AddCurrencyToBuyOrSell() *CurrencyToBuyOrSell1Choice {
	s.CurrencyToBuyOrSell = new(CurrencyToBuyOrSell1Choice)
	return s.CurrencyToBuyOrSell
}

func (s *SecuritiesTradeDetails1) AddMatchingStatus() *MatchingStatus1Choice {
	s.MatchingStatus = new(MatchingStatus1Choice)
	return s.MatchingStatus
}

func (s *SecuritiesTradeDetails1) AddAffirmationStatus() *AffirmationStatus1Choice {
	s.AffirmationStatus = new(AffirmationStatus1Choice)
	return s.AffirmationStatus
}

func (s *SecuritiesTradeDetails1) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*Max350Text)(&value)
}

func (s *SecuritiesTradeDetails1) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the securities trade.
type SecuritiesTradeDetails12 struct {

	// Specifies the date/time on which the trade was executed.
	TradeDate *DateAndDateTimeChoice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	OpeningSettlementDate *DateAndDateTimeChoice `xml:"OpngSttlmDt"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	InstructionProcessingAdditionalDetails *RestrictedFINXMax350Text `xml:"InstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails12) AddTradeDate() *DateAndDateTimeChoice {
	s.TradeDate = new(DateAndDateTimeChoice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails12) AddOpeningSettlementDate() *DateAndDateTimeChoice {
	s.OpeningSettlementDate = new(DateAndDateTimeChoice)
	return s.OpeningSettlementDate
}

func (s *SecuritiesTradeDetails12) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails12) SetInstructionProcessingAdditionalDetails(value string) {
	s.InstructionProcessingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

// Details of the securities trade.
type SecuritiesTradeDetails17 struct {

	// Identification of an account owner transaction that could potentially match with the allegement notified.
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*Max35Text `xml:"CollTxId,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Status of the allegement.
	Status *AllegementStatus1Choice `xml:"Sts,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes15 `xml:"FinInstrmAttrbts,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition1Choice `xml:"TradTxCond,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice3Choice `xml:"TpOfPric,omitempty"`

	// Details about the financial instrument quantity and the account involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount14 `xml:"QtyAndAcctDtls"`

	// Details of the closing of the securities financing transaction.
	SecuritiesFinancingDetails *SecuritiesFinancingTransactionDetails7 `xml:"SctiesFincgDtls,omitempty"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails25 `xml:"SttlmParams"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties11 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties11 `xml:"RcvgSttlmPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection22 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts8 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties11 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesTradeDetails17) SetAccountOwnerTransactionIdentification(value string) {
	s.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails17) SetAccountServicerTransactionIdentification(value string) {
	s.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails17) SetMarketInfrastructureTransactionIdentification(value string) {
	s.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails17) SetProcessorTransactionIdentification(value string) {
	s.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails17) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails17) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails17) SetPoolIdentification(value string) {
	s.PoolIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails17) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails17) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesTradeDetails17) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesTradeDetails17) AddStatus() *AllegementStatus1Choice {
	s.Status = new(AllegementStatus1Choice)
	return s.Status
}

func (s *SecuritiesTradeDetails17) AddPlaceOfTrade() *MarketIdentification4 {
	s.PlaceOfTrade = new(MarketIdentification4)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails17) SetPlaceOfClearing(value string) {
	s.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (s *SecuritiesTradeDetails17) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails17) AddSettlementDate() *SettlementDate1Choice {
	s.SettlementDate = new(SettlementDate1Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails17) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails17) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails17) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesTradeDetails17) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes15 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes15)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesTradeDetails17) AddTradeTransactionCondition() *TradeTransactionCondition1Choice {
	newValue := new(TradeTransactionCondition1Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails17) AddTypeOfPrice() *TypeOfPrice3Choice {
	s.TypeOfPrice = new(TypeOfPrice3Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails17) AddQuantityAndAccountDetails() *QuantityAndAccount14 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount14)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesTradeDetails17) AddSecuritiesFinancingDetails() *SecuritiesFinancingTransactionDetails7 {
	s.SecuritiesFinancingDetails = new(SecuritiesFinancingTransactionDetails7)
	return s.SecuritiesFinancingDetails
}

func (s *SecuritiesTradeDetails17) AddSettlementParameters() *SettlementDetails25 {
	s.SettlementParameters = new(SettlementDetails25)
	return s.SettlementParameters
}

func (s *SecuritiesTradeDetails17) AddDeliveringSettlementParties() *SettlementParties11 {
	s.DeliveringSettlementParties = new(SettlementParties11)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesTradeDetails17) AddReceivingSettlementParties() *SettlementParties11 {
	s.ReceivingSettlementParties = new(SettlementParties11)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesTradeDetails17) AddSettlementAmount() *AmountAndDirection22 {
	s.SettlementAmount = new(AmountAndDirection22)
	return s.SettlementAmount
}

func (s *SecuritiesTradeDetails17) AddOtherAmounts() *OtherAmounts8 {
	s.OtherAmounts = new(OtherAmounts8)
	return s.OtherAmounts
}

func (s *SecuritiesTradeDetails17) AddOtherBusinessParties() *OtherParties11 {
	s.OtherBusinessParties = new(OtherParties11)
	return s.OtherBusinessParties
}

func (s *SecuritiesTradeDetails17) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Details of the securities trade.
type SecuritiesTradeDetails18 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*Max35Text `xml:"CollTxId,omitempty"`

	// Identification of an account owner transaction that could potentially match with the allegement notified.
	AccountOwnerTransactionIdentification []*Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition1Choice `xml:"TradTxCond,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice3Choice `xml:"TpOfPric,omitempty"`
}

func (s *SecuritiesTradeDetails18) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails18) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails18) AddAccountOwnerTransactionIdentification(value string) {
	s.AccountOwnerTransactionIdentification = append(s.AccountOwnerTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails18) SetProcessorTransactionIdentification(value string) {
	s.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails18) AddPlaceOfTrade() *MarketIdentification4 {
	s.PlaceOfTrade = new(MarketIdentification4)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails18) SetPlaceOfClearing(value string) {
	s.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (s *SecuritiesTradeDetails18) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails18) AddSettlementDate() *SettlementDate1Choice {
	s.SettlementDate = new(SettlementDate1Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails18) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails18) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails18) AddTradeTransactionCondition() *TradeTransactionCondition1Choice {
	newValue := new(TradeTransactionCondition1Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails18) AddTypeOfPrice() *TypeOfPrice3Choice {
	s.TypeOfPrice = new(TypeOfPrice3Choice)
	return s.TypeOfPrice
}

// Details of the securities trade.
type SecuritiesTradeDetails19 struct {

	// Unambiguous identification of the transaction as known by the notification receiver.
	NotificationSenderTransactionIdentification *Max35Text `xml:"NtfctnSndrTxId,omitempty"`

	// Unambiguous identification of the transaction as known by the notification receiver.
	NotificationReceiverTransactionIdentification *Max35Text `xml:"NtfctnRcvrTxId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes20 `xml:"FinInstrmAttrbts,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting1Choice `xml:"Rptg,omitempty"`

	// Details about the financial instrument quantity involved in the transfer.
	QuantityDetails *Quantity5 `xml:"QtyDtls"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails4 `xml:"SttlmParams,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties10 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties10 `xml:"RcvgSttlmPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection7 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts2 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties10 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesTradeDetails19) SetNotificationSenderTransactionIdentification(value string) {
	s.NotificationSenderTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails19) SetNotificationReceiverTransactionIdentification(value string) {
	s.NotificationReceiverTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails19) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails19) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesTradeDetails19) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesTradeDetails19) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails19) AddSettlementDate() *SettlementDate1Choice {
	s.SettlementDate = new(SettlementDate1Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails19) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails19) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesTradeDetails19) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes20 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes20)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesTradeDetails19) AddReporting() *Reporting1Choice {
	newValue := new(Reporting1Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails19) AddQuantityDetails() *Quantity5 {
	s.QuantityDetails = new(Quantity5)
	return s.QuantityDetails
}

func (s *SecuritiesTradeDetails19) AddSettlementParameters() *SettlementDetails4 {
	s.SettlementParameters = new(SettlementDetails4)
	return s.SettlementParameters
}

func (s *SecuritiesTradeDetails19) AddDeliveringSettlementParties() *SettlementParties10 {
	s.DeliveringSettlementParties = new(SettlementParties10)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesTradeDetails19) AddReceivingSettlementParties() *SettlementParties10 {
	s.ReceivingSettlementParties = new(SettlementParties10)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesTradeDetails19) AddSettlementAmount() *AmountAndDirection7 {
	s.SettlementAmount = new(AmountAndDirection7)
	return s.SettlementAmount
}

func (s *SecuritiesTradeDetails19) AddOtherAmounts() *OtherAmounts2 {
	s.OtherAmounts = new(OtherAmounts2)
	return s.OtherAmounts
}

func (s *SecuritiesTradeDetails19) AddOtherBusinessParties() *OtherParties10 {
	s.OtherBusinessParties = new(OtherParties10)
	return s.OtherBusinessParties
}

func (s *SecuritiesTradeDetails19) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Details of the securities trade.
type SecuritiesTradeDetails2 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*Max35Text `xml:"CollTxId,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, ie, payment is effected and securities are delivered.
	EffectiveSettlementDate *SettlementDate3Choice `xml:"FctvSttlmDt"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Specifies additional information relative to the processing of the trade.
	OpeningClosing *OpeningClosing1Choice `xml:"OpngClsg,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting2Choice `xml:"Rptg,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition1Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity1Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator1Choice `xml:"TradOrgtrRole,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice3Choice `xml:"TpOfPric,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *Max350Text `xml:"FxAddtlDtls,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails2) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails2) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails2) AddPlaceOfTrade() *MarketIdentification4 {
	s.PlaceOfTrade = new(MarketIdentification4)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails2) SetPlaceOfClearing(value string) {
	s.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (s *SecuritiesTradeDetails2) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails2) AddSettlementDate() *SettlementDate1Choice {
	s.SettlementDate = new(SettlementDate1Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails2) AddEffectiveSettlementDate() *SettlementDate3Choice {
	s.EffectiveSettlementDate = new(SettlementDate3Choice)
	return s.EffectiveSettlementDate
}

func (s *SecuritiesTradeDetails2) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails2) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails2) AddOpeningClosing() *OpeningClosing1Choice {
	s.OpeningClosing = new(OpeningClosing1Choice)
	return s.OpeningClosing
}

func (s *SecuritiesTradeDetails2) AddReporting() *Reporting2Choice {
	newValue := new(Reporting2Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails2) AddTradeTransactionCondition() *TradeTransactionCondition1Choice {
	newValue := new(TradeTransactionCondition1Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails2) AddInvestorCapacity() *InvestorCapacity1Choice {
	s.InvestorCapacity = new(InvestorCapacity1Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails2) AddTradeOriginatorRole() *TradeOriginator1Choice {
	s.TradeOriginatorRole = new(TradeOriginator1Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails2) AddTypeOfPrice() *TypeOfPrice3Choice {
	s.TypeOfPrice = new(TypeOfPrice3Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails2) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*Max350Text)(&value)
}

func (s *SecuritiesTradeDetails2) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the securities trade.
type SecuritiesTradeDetails23 struct {

	// Identification of an account owner transaction that could potentially match with the allegement notified.
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*Max35Text `xml:"CollTxId,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Status of the allegement.
	Status *AllegementStatus1Choice `xml:"Sts,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes35 `xml:"FinInstrmAttrbts,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition1Choice `xml:"TradTxCond,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice3Choice `xml:"TpOfPric,omitempty"`

	// Details about the financial instrument quantity and the account involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount26 `xml:"QtyAndAcctDtls"`

	// Details of the closing of the securities financing transaction.
	SecuritiesFinancingDetails *SecuritiesFinancingTransactionDetails7 `xml:"SctiesFincgDtls,omitempty"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails49 `xml:"SttlmParams"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties11 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties11 `xml:"RcvgSttlmPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection22 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts8 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties11 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesTradeDetails23) SetAccountOwnerTransactionIdentification(value string) {
	s.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails23) SetAccountServicerTransactionIdentification(value string) {
	s.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails23) SetMarketInfrastructureTransactionIdentification(value string) {
	s.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails23) SetProcessorTransactionIdentification(value string) {
	s.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails23) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails23) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails23) SetPoolIdentification(value string) {
	s.PoolIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails23) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails23) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesTradeDetails23) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesTradeDetails23) AddStatus() *AllegementStatus1Choice {
	s.Status = new(AllegementStatus1Choice)
	return s.Status
}

func (s *SecuritiesTradeDetails23) AddPlaceOfTrade() *MarketIdentification4 {
	s.PlaceOfTrade = new(MarketIdentification4)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails23) SetPlaceOfClearing(value string) {
	s.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (s *SecuritiesTradeDetails23) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails23) AddSettlementDate() *SettlementDate1Choice {
	s.SettlementDate = new(SettlementDate1Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails23) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails23) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails23) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesTradeDetails23) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes35 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes35)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesTradeDetails23) AddTradeTransactionCondition() *TradeTransactionCondition1Choice {
	newValue := new(TradeTransactionCondition1Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails23) AddTypeOfPrice() *TypeOfPrice3Choice {
	s.TypeOfPrice = new(TypeOfPrice3Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails23) AddQuantityAndAccountDetails() *QuantityAndAccount26 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount26)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesTradeDetails23) AddSecuritiesFinancingDetails() *SecuritiesFinancingTransactionDetails7 {
	s.SecuritiesFinancingDetails = new(SecuritiesFinancingTransactionDetails7)
	return s.SecuritiesFinancingDetails
}

func (s *SecuritiesTradeDetails23) AddSettlementParameters() *SettlementDetails49 {
	s.SettlementParameters = new(SettlementDetails49)
	return s.SettlementParameters
}

func (s *SecuritiesTradeDetails23) AddDeliveringSettlementParties() *SettlementParties11 {
	s.DeliveringSettlementParties = new(SettlementParties11)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesTradeDetails23) AddReceivingSettlementParties() *SettlementParties11 {
	s.ReceivingSettlementParties = new(SettlementParties11)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesTradeDetails23) AddSettlementAmount() *AmountAndDirection22 {
	s.SettlementAmount = new(AmountAndDirection22)
	return s.SettlementAmount
}

func (s *SecuritiesTradeDetails23) AddOtherAmounts() *OtherAmounts8 {
	s.OtherAmounts = new(OtherAmounts8)
	return s.OtherAmounts
}

func (s *SecuritiesTradeDetails23) AddOtherBusinessParties() *OtherParties11 {
	s.OtherBusinessParties = new(OtherParties11)
	return s.OtherBusinessParties
}

func (s *SecuritiesTradeDetails23) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Details of the securities trade.
type SecuritiesTradeDetails24 struct {

	// Unambiguous identification of the transaction as known by the notification receiver.
	NotificationSenderTransactionIdentification *Max35Text `xml:"NtfctnSndrTxId,omitempty"`

	// Unambiguous identification of the transaction as known by the notification receiver.
	NotificationReceiverTransactionIdentification *Max35Text `xml:"NtfctnRcvrTxId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes35 `xml:"FinInstrmAttrbts,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting1Choice `xml:"Rptg,omitempty"`

	// Details about the financial instrument quantity involved in the transfer.
	QuantityDetails *Quantity7 `xml:"QtyDtls"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails4 `xml:"SttlmParams,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties10 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties10 `xml:"RcvgSttlmPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection7 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts15 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties10 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesTradeDetails24) SetNotificationSenderTransactionIdentification(value string) {
	s.NotificationSenderTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails24) SetNotificationReceiverTransactionIdentification(value string) {
	s.NotificationReceiverTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails24) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails24) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesTradeDetails24) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesTradeDetails24) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails24) AddSettlementDate() *SettlementDate1Choice {
	s.SettlementDate = new(SettlementDate1Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails24) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails24) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesTradeDetails24) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes35 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes35)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesTradeDetails24) AddReporting() *Reporting1Choice {
	newValue := new(Reporting1Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails24) AddQuantityDetails() *Quantity7 {
	s.QuantityDetails = new(Quantity7)
	return s.QuantityDetails
}

func (s *SecuritiesTradeDetails24) AddSettlementParameters() *SettlementDetails4 {
	s.SettlementParameters = new(SettlementDetails4)
	return s.SettlementParameters
}

func (s *SecuritiesTradeDetails24) AddDeliveringSettlementParties() *SettlementParties10 {
	s.DeliveringSettlementParties = new(SettlementParties10)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesTradeDetails24) AddReceivingSettlementParties() *SettlementParties10 {
	s.ReceivingSettlementParties = new(SettlementParties10)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesTradeDetails24) AddSettlementAmount() *AmountAndDirection7 {
	s.SettlementAmount = new(AmountAndDirection7)
	return s.SettlementAmount
}

func (s *SecuritiesTradeDetails24) AddOtherAmounts() *OtherAmounts15 {
	s.OtherAmounts = new(OtherAmounts15)
	return s.OtherAmounts
}

func (s *SecuritiesTradeDetails24) AddOtherBusinessParties() *OtherParties10 {
	s.OtherBusinessParties = new(OtherParties10)
	return s.OtherBusinessParties
}

func (s *SecuritiesTradeDetails24) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Details of the securities trade.
type SecuritiesTradeDetails25 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*Max35Text `xml:"CollTxId,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Specifies additional information relative to the processing of the trade.
	OpeningClosing *OpeningClosing1Choice `xml:"OpngClsg,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting2Choice `xml:"Rptg,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition1Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity1Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator1Choice `xml:"TradOrgtrRole,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice3Choice `xml:"TpOfPric,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds or to sell the indicated currency in order to obtain the necessary currency to fund the transaction.
	CurrencyToBuyOrSell *CurrencyToBuyOrSell1Choice `xml:"CcyToBuyOrSell,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus1Choice `xml:"MtchgSts,omitempty"`

	// Status of affirmation of a trade.
	AffirmationStatus *AffirmationStatus1Choice `xml:"AffirmSts,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *Max350Text `xml:"FxAddtlDtls,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails25) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails25) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails25) AddPlaceOfTrade() *MarketIdentification4 {
	s.PlaceOfTrade = new(MarketIdentification4)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails25) SetPlaceOfClearing(value string) {
	s.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (s *SecuritiesTradeDetails25) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails25) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesTradeDetails25) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails25) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails25) AddOpeningClosing() *OpeningClosing1Choice {
	s.OpeningClosing = new(OpeningClosing1Choice)
	return s.OpeningClosing
}

func (s *SecuritiesTradeDetails25) AddReporting() *Reporting2Choice {
	newValue := new(Reporting2Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails25) AddTradeTransactionCondition() *TradeTransactionCondition1Choice {
	newValue := new(TradeTransactionCondition1Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails25) AddInvestorCapacity() *InvestorCapacity1Choice {
	s.InvestorCapacity = new(InvestorCapacity1Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails25) AddTradeOriginatorRole() *TradeOriginator1Choice {
	s.TradeOriginatorRole = new(TradeOriginator1Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails25) AddTypeOfPrice() *TypeOfPrice3Choice {
	s.TypeOfPrice = new(TypeOfPrice3Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails25) AddCurrencyToBuyOrSell() *CurrencyToBuyOrSell1Choice {
	s.CurrencyToBuyOrSell = new(CurrencyToBuyOrSell1Choice)
	return s.CurrencyToBuyOrSell
}

func (s *SecuritiesTradeDetails25) AddMatchingStatus() *MatchingStatus1Choice {
	s.MatchingStatus = new(MatchingStatus1Choice)
	return s.MatchingStatus
}

func (s *SecuritiesTradeDetails25) AddAffirmationStatus() *AffirmationStatus1Choice {
	s.AffirmationStatus = new(AffirmationStatus1Choice)
	return s.AffirmationStatus
}

func (s *SecuritiesTradeDetails25) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*Max350Text)(&value)
}

func (s *SecuritiesTradeDetails25) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the securities trade.
type SecuritiesTradeDetails26 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*Max35Text `xml:"CollTxId,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Specifies additional information relative to the processing of the trade.
	OpeningClosing *OpeningClosing1Choice `xml:"OpngClsg,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting2Choice `xml:"Rptg,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition1Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity1Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator1Choice `xml:"TradOrgtrRole,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice3Choice `xml:"TpOfPric,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds or to sell the indicated currency in order to obtain the necessary currency to fund the transaction.
	CurrencyToBuyOrSell *CurrencyToBuyOrSell1Choice `xml:"CcyToBuyOrSell,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus1Choice `xml:"MtchgSts,omitempty"`

	// Status of affirmation of a trade.
	AffirmationStatus *AffirmationStatus1Choice `xml:"AffirmSts,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *Max350Text `xml:"FxAddtlDtls,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails26) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails26) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails26) AddPlaceOfTrade() *MarketIdentification4 {
	s.PlaceOfTrade = new(MarketIdentification4)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails26) SetPlaceOfClearing(value string) {
	s.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (s *SecuritiesTradeDetails26) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails26) AddSettlementDate() *SettlementDate1Choice {
	s.SettlementDate = new(SettlementDate1Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails26) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesTradeDetails26) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails26) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails26) AddOpeningClosing() *OpeningClosing1Choice {
	s.OpeningClosing = new(OpeningClosing1Choice)
	return s.OpeningClosing
}

func (s *SecuritiesTradeDetails26) AddReporting() *Reporting2Choice {
	newValue := new(Reporting2Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails26) AddTradeTransactionCondition() *TradeTransactionCondition1Choice {
	newValue := new(TradeTransactionCondition1Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails26) AddInvestorCapacity() *InvestorCapacity1Choice {
	s.InvestorCapacity = new(InvestorCapacity1Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails26) AddTradeOriginatorRole() *TradeOriginator1Choice {
	s.TradeOriginatorRole = new(TradeOriginator1Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails26) AddTypeOfPrice() *TypeOfPrice3Choice {
	s.TypeOfPrice = new(TypeOfPrice3Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails26) AddCurrencyToBuyOrSell() *CurrencyToBuyOrSell1Choice {
	s.CurrencyToBuyOrSell = new(CurrencyToBuyOrSell1Choice)
	return s.CurrencyToBuyOrSell
}

func (s *SecuritiesTradeDetails26) AddMatchingStatus() *MatchingStatus1Choice {
	s.MatchingStatus = new(MatchingStatus1Choice)
	return s.MatchingStatus
}

func (s *SecuritiesTradeDetails26) AddAffirmationStatus() *AffirmationStatus1Choice {
	s.AffirmationStatus = new(AffirmationStatus1Choice)
	return s.AffirmationStatus
}

func (s *SecuritiesTradeDetails26) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*Max350Text)(&value)
}

func (s *SecuritiesTradeDetails26) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the securities trade.
type SecuritiesTradeDetails3 struct {

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	OpeningSettlementDate *DateAndDateTimeChoice `xml:"OpngSttlmDt"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting2Choice `xml:"Rptg,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition1Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity1Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator1Choice `xml:"TradOrgtrRole,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds or to sell the indicated currency in order to obtain the necessary currency to fund the transaction.
	CurrencyToBuyOrSell *CurrencyToBuyOrSell1Choice `xml:"CcyToBuyOrSell,omitempty"`

	// Status of affirmation of a trade.
	AffirmationStatus *AffirmationStatus1Choice `xml:"AffirmSts,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus1Choice `xml:"MtchgSts,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *Max350Text `xml:"FxAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails3) AddPlaceOfTrade() *MarketIdentification4 {
	s.PlaceOfTrade = new(MarketIdentification4)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails3) SetPlaceOfClearing(value string) {
	s.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (s *SecuritiesTradeDetails3) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails3) AddOpeningSettlementDate() *DateAndDateTimeChoice {
	s.OpeningSettlementDate = new(DateAndDateTimeChoice)
	return s.OpeningSettlementDate
}

func (s *SecuritiesTradeDetails3) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails3) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails3) AddReporting() *Reporting2Choice {
	newValue := new(Reporting2Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails3) AddTradeTransactionCondition() *TradeTransactionCondition1Choice {
	newValue := new(TradeTransactionCondition1Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails3) AddInvestorCapacity() *InvestorCapacity1Choice {
	s.InvestorCapacity = new(InvestorCapacity1Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails3) AddTradeOriginatorRole() *TradeOriginator1Choice {
	s.TradeOriginatorRole = new(TradeOriginator1Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails3) AddCurrencyToBuyOrSell() *CurrencyToBuyOrSell1Choice {
	s.CurrencyToBuyOrSell = new(CurrencyToBuyOrSell1Choice)
	return s.CurrencyToBuyOrSell
}

func (s *SecuritiesTradeDetails3) AddAffirmationStatus() *AffirmationStatus1Choice {
	s.AffirmationStatus = new(AffirmationStatus1Choice)
	return s.AffirmationStatus
}

func (s *SecuritiesTradeDetails3) AddMatchingStatus() *MatchingStatus1Choice {
	s.MatchingStatus = new(MatchingStatus1Choice)
	return s.MatchingStatus
}

func (s *SecuritiesTradeDetails3) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

func (s *SecuritiesTradeDetails3) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*Max350Text)(&value)
}

// Details of the securities trade.
type SecuritiesTradeDetails31 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*Max35Text `xml:"CollTxId,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification78 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, ie, payment is effected and securities are delivered.
	EffectiveSettlementDate *SettlementDate3Choice `xml:"FctvSttlmDt"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Specifies additional information relative to the processing of the trade.
	OpeningClosing *OpeningClosing1Choice `xml:"OpngClsg,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting2Choice `xml:"Rptg,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition1Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity1Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator1Choice `xml:"TradOrgtrRole,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice3Choice `xml:"TpOfPric,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *Max350Text `xml:"FxAddtlDtls,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails31) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails31) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails31) AddPlaceOfTrade() *MarketIdentification78 {
	s.PlaceOfTrade = new(MarketIdentification78)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails31) SetPlaceOfClearing(value string) {
	s.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (s *SecuritiesTradeDetails31) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails31) AddSettlementDate() *SettlementDate1Choice {
	s.SettlementDate = new(SettlementDate1Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails31) AddEffectiveSettlementDate() *SettlementDate3Choice {
	s.EffectiveSettlementDate = new(SettlementDate3Choice)
	return s.EffectiveSettlementDate
}

func (s *SecuritiesTradeDetails31) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails31) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails31) AddOpeningClosing() *OpeningClosing1Choice {
	s.OpeningClosing = new(OpeningClosing1Choice)
	return s.OpeningClosing
}

func (s *SecuritiesTradeDetails31) AddReporting() *Reporting2Choice {
	newValue := new(Reporting2Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails31) AddTradeTransactionCondition() *TradeTransactionCondition1Choice {
	newValue := new(TradeTransactionCondition1Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails31) AddInvestorCapacity() *InvestorCapacity1Choice {
	s.InvestorCapacity = new(InvestorCapacity1Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails31) AddTradeOriginatorRole() *TradeOriginator1Choice {
	s.TradeOriginatorRole = new(TradeOriginator1Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails31) AddTypeOfPrice() *TypeOfPrice3Choice {
	s.TypeOfPrice = new(TypeOfPrice3Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails31) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*Max350Text)(&value)
}

func (s *SecuritiesTradeDetails31) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the securities trade.
type SecuritiesTradeDetails32 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*Max35Text `xml:"CollTxId,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification78 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Specifies additional information relative to the processing of the trade.
	OpeningClosing *OpeningClosing1Choice `xml:"OpngClsg,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting2Choice `xml:"Rptg,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition1Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity1Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator1Choice `xml:"TradOrgtrRole,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice3Choice `xml:"TpOfPric,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds or to sell the indicated currency in order to obtain the necessary currency to fund the transaction.
	CurrencyToBuyOrSell *CurrencyToBuyOrSell1Choice `xml:"CcyToBuyOrSell,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus1Choice `xml:"MtchgSts,omitempty"`

	// Status of affirmation of a trade.
	AffirmationStatus *AffirmationStatus1Choice `xml:"AffirmSts,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *Max350Text `xml:"FxAddtlDtls,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails32) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails32) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails32) AddPlaceOfTrade() *MarketIdentification78 {
	s.PlaceOfTrade = new(MarketIdentification78)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails32) SetPlaceOfClearing(value string) {
	s.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (s *SecuritiesTradeDetails32) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails32) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesTradeDetails32) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails32) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails32) AddOpeningClosing() *OpeningClosing1Choice {
	s.OpeningClosing = new(OpeningClosing1Choice)
	return s.OpeningClosing
}

func (s *SecuritiesTradeDetails32) AddReporting() *Reporting2Choice {
	newValue := new(Reporting2Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails32) AddTradeTransactionCondition() *TradeTransactionCondition1Choice {
	newValue := new(TradeTransactionCondition1Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails32) AddInvestorCapacity() *InvestorCapacity1Choice {
	s.InvestorCapacity = new(InvestorCapacity1Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails32) AddTradeOriginatorRole() *TradeOriginator1Choice {
	s.TradeOriginatorRole = new(TradeOriginator1Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails32) AddTypeOfPrice() *TypeOfPrice3Choice {
	s.TypeOfPrice = new(TypeOfPrice3Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails32) AddCurrencyToBuyOrSell() *CurrencyToBuyOrSell1Choice {
	s.CurrencyToBuyOrSell = new(CurrencyToBuyOrSell1Choice)
	return s.CurrencyToBuyOrSell
}

func (s *SecuritiesTradeDetails32) AddMatchingStatus() *MatchingStatus1Choice {
	s.MatchingStatus = new(MatchingStatus1Choice)
	return s.MatchingStatus
}

func (s *SecuritiesTradeDetails32) AddAffirmationStatus() *AffirmationStatus1Choice {
	s.AffirmationStatus = new(AffirmationStatus1Choice)
	return s.AffirmationStatus
}

func (s *SecuritiesTradeDetails32) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*Max350Text)(&value)
}

func (s *SecuritiesTradeDetails32) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the securities trade.
type SecuritiesTradeDetails33 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*Max35Text `xml:"CollTxId,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification78 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Specifies additional information relative to the processing of the trade.
	OpeningClosing *OpeningClosing1Choice `xml:"OpngClsg,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting2Choice `xml:"Rptg,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition1Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity1Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator1Choice `xml:"TradOrgtrRole,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice3Choice `xml:"TpOfPric,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds or to sell the indicated currency in order to obtain the necessary currency to fund the transaction.
	CurrencyToBuyOrSell *CurrencyToBuyOrSell1Choice `xml:"CcyToBuyOrSell,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus1Choice `xml:"MtchgSts,omitempty"`

	// Status of affirmation of a trade.
	AffirmationStatus *AffirmationStatus1Choice `xml:"AffirmSts,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *Max350Text `xml:"FxAddtlDtls,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails33) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails33) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails33) AddPlaceOfTrade() *MarketIdentification78 {
	s.PlaceOfTrade = new(MarketIdentification78)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails33) SetPlaceOfClearing(value string) {
	s.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (s *SecuritiesTradeDetails33) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails33) AddSettlementDate() *SettlementDate1Choice {
	s.SettlementDate = new(SettlementDate1Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails33) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesTradeDetails33) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails33) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails33) AddOpeningClosing() *OpeningClosing1Choice {
	s.OpeningClosing = new(OpeningClosing1Choice)
	return s.OpeningClosing
}

func (s *SecuritiesTradeDetails33) AddReporting() *Reporting2Choice {
	newValue := new(Reporting2Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails33) AddTradeTransactionCondition() *TradeTransactionCondition1Choice {
	newValue := new(TradeTransactionCondition1Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails33) AddInvestorCapacity() *InvestorCapacity1Choice {
	s.InvestorCapacity = new(InvestorCapacity1Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails33) AddTradeOriginatorRole() *TradeOriginator1Choice {
	s.TradeOriginatorRole = new(TradeOriginator1Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails33) AddTypeOfPrice() *TypeOfPrice3Choice {
	s.TypeOfPrice = new(TypeOfPrice3Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails33) AddCurrencyToBuyOrSell() *CurrencyToBuyOrSell1Choice {
	s.CurrencyToBuyOrSell = new(CurrencyToBuyOrSell1Choice)
	return s.CurrencyToBuyOrSell
}

func (s *SecuritiesTradeDetails33) AddMatchingStatus() *MatchingStatus1Choice {
	s.MatchingStatus = new(MatchingStatus1Choice)
	return s.MatchingStatus
}

func (s *SecuritiesTradeDetails33) AddAffirmationStatus() *AffirmationStatus1Choice {
	s.AffirmationStatus = new(AffirmationStatus1Choice)
	return s.AffirmationStatus
}

func (s *SecuritiesTradeDetails33) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*Max350Text)(&value)
}

func (s *SecuritiesTradeDetails33) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the securities trade.
type SecuritiesTradeDetails34 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*Max35Text `xml:"CollTxId,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification78 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Specifies additional information relative to the processing of the trade.
	OpeningClosing *OpeningClosing1Choice `xml:"OpngClsg,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting2Choice `xml:"Rptg,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition1Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity1Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator1Choice `xml:"TradOrgtrRole,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice3Choice `xml:"TpOfPric,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds or to sell the indicated currency in order to obtain the necessary currency to fund the transaction.
	CurrencyToBuyOrSell *CurrencyToBuyOrSell1Choice `xml:"CcyToBuyOrSell,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus1Choice `xml:"MtchgSts,omitempty"`

	// Status of affirmation of a trade.
	AffirmationStatus *AffirmationStatus1Choice `xml:"AffirmSts,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *Max350Text `xml:"FxAddtlDtls,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails34) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails34) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails34) AddPlaceOfTrade() *MarketIdentification78 {
	s.PlaceOfTrade = new(MarketIdentification78)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails34) SetPlaceOfClearing(value string) {
	s.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (s *SecuritiesTradeDetails34) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails34) AddSettlementDate() *SettlementDate1Choice {
	s.SettlementDate = new(SettlementDate1Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails34) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesTradeDetails34) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails34) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails34) AddOpeningClosing() *OpeningClosing1Choice {
	s.OpeningClosing = new(OpeningClosing1Choice)
	return s.OpeningClosing
}

func (s *SecuritiesTradeDetails34) AddReporting() *Reporting2Choice {
	newValue := new(Reporting2Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails34) AddTradeTransactionCondition() *TradeTransactionCondition1Choice {
	newValue := new(TradeTransactionCondition1Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails34) AddInvestorCapacity() *InvestorCapacity1Choice {
	s.InvestorCapacity = new(InvestorCapacity1Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails34) AddTradeOriginatorRole() *TradeOriginator1Choice {
	s.TradeOriginatorRole = new(TradeOriginator1Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails34) AddTypeOfPrice() *TypeOfPrice3Choice {
	s.TypeOfPrice = new(TypeOfPrice3Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails34) AddCurrencyToBuyOrSell() *CurrencyToBuyOrSell1Choice {
	s.CurrencyToBuyOrSell = new(CurrencyToBuyOrSell1Choice)
	return s.CurrencyToBuyOrSell
}

func (s *SecuritiesTradeDetails34) AddMatchingStatus() *MatchingStatus1Choice {
	s.MatchingStatus = new(MatchingStatus1Choice)
	return s.MatchingStatus
}

func (s *SecuritiesTradeDetails34) AddAffirmationStatus() *AffirmationStatus1Choice {
	s.AffirmationStatus = new(AffirmationStatus1Choice)
	return s.AffirmationStatus
}

func (s *SecuritiesTradeDetails34) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*Max350Text)(&value)
}

func (s *SecuritiesTradeDetails34) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the securities trade.
type SecuritiesTradeDetails35 struct {

	// Identification of an account owner transaction that could potentially match with the allegement notified.
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*Max35Text `xml:"CollTxId,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Status of the allegement.
	Status *AllegementStatus1Choice `xml:"Sts,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification78 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes35 `xml:"FinInstrmAttrbts,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition1Choice `xml:"TradTxCond,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice3Choice `xml:"TpOfPric,omitempty"`

	// Details about the financial instrument quantity and the account involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount26 `xml:"QtyAndAcctDtls"`

	// Details of the closing of the securities financing transaction.
	SecuritiesFinancingDetails *SecuritiesFinancingTransactionDetails7 `xml:"SctiesFincgDtls,omitempty"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails49 `xml:"SttlmParams"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties11 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties11 `xml:"RcvgSttlmPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection22 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts8 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties11 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesTradeDetails35) SetAccountOwnerTransactionIdentification(value string) {
	s.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails35) SetAccountServicerTransactionIdentification(value string) {
	s.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails35) SetMarketInfrastructureTransactionIdentification(value string) {
	s.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails35) SetProcessorTransactionIdentification(value string) {
	s.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails35) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails35) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails35) SetPoolIdentification(value string) {
	s.PoolIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails35) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails35) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesTradeDetails35) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesTradeDetails35) AddStatus() *AllegementStatus1Choice {
	s.Status = new(AllegementStatus1Choice)
	return s.Status
}

func (s *SecuritiesTradeDetails35) AddPlaceOfTrade() *MarketIdentification78 {
	s.PlaceOfTrade = new(MarketIdentification78)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails35) SetPlaceOfClearing(value string) {
	s.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (s *SecuritiesTradeDetails35) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails35) AddSettlementDate() *SettlementDate1Choice {
	s.SettlementDate = new(SettlementDate1Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails35) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails35) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails35) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesTradeDetails35) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes35 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes35)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesTradeDetails35) AddTradeTransactionCondition() *TradeTransactionCondition1Choice {
	newValue := new(TradeTransactionCondition1Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails35) AddTypeOfPrice() *TypeOfPrice3Choice {
	s.TypeOfPrice = new(TypeOfPrice3Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails35) AddQuantityAndAccountDetails() *QuantityAndAccount26 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount26)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesTradeDetails35) AddSecuritiesFinancingDetails() *SecuritiesFinancingTransactionDetails7 {
	s.SecuritiesFinancingDetails = new(SecuritiesFinancingTransactionDetails7)
	return s.SecuritiesFinancingDetails
}

func (s *SecuritiesTradeDetails35) AddSettlementParameters() *SettlementDetails49 {
	s.SettlementParameters = new(SettlementDetails49)
	return s.SettlementParameters
}

func (s *SecuritiesTradeDetails35) AddDeliveringSettlementParties() *SettlementParties11 {
	s.DeliveringSettlementParties = new(SettlementParties11)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesTradeDetails35) AddReceivingSettlementParties() *SettlementParties11 {
	s.ReceivingSettlementParties = new(SettlementParties11)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesTradeDetails35) AddSettlementAmount() *AmountAndDirection22 {
	s.SettlementAmount = new(AmountAndDirection22)
	return s.SettlementAmount
}

func (s *SecuritiesTradeDetails35) AddOtherAmounts() *OtherAmounts8 {
	s.OtherAmounts = new(OtherAmounts8)
	return s.OtherAmounts
}

func (s *SecuritiesTradeDetails35) AddOtherBusinessParties() *OtherParties11 {
	s.OtherBusinessParties = new(OtherParties11)
	return s.OtherBusinessParties
}

func (s *SecuritiesTradeDetails35) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Details of the securities trade.
type SecuritiesTradeDetails36 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*Max35Text `xml:"CollTxId,omitempty"`

	// Identification of an account owner transaction that could potentially match with the allegement notified.
	AccountOwnerTransactionIdentification []*Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification78 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition1Choice `xml:"TradTxCond,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice3Choice `xml:"TpOfPric,omitempty"`
}

func (s *SecuritiesTradeDetails36) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails36) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails36) AddAccountOwnerTransactionIdentification(value string) {
	s.AccountOwnerTransactionIdentification = append(s.AccountOwnerTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails36) SetProcessorTransactionIdentification(value string) {
	s.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails36) AddPlaceOfTrade() *MarketIdentification78 {
	s.PlaceOfTrade = new(MarketIdentification78)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails36) SetPlaceOfClearing(value string) {
	s.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (s *SecuritiesTradeDetails36) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails36) AddSettlementDate() *SettlementDate1Choice {
	s.SettlementDate = new(SettlementDate1Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails36) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails36) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails36) AddTradeTransactionCondition() *TradeTransactionCondition1Choice {
	newValue := new(TradeTransactionCondition1Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails36) AddTypeOfPrice() *TypeOfPrice3Choice {
	s.TypeOfPrice = new(TypeOfPrice3Choice)
	return s.TypeOfPrice
}

// Details of the securities trade.
type SecuritiesTradeDetails37 struct {

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification78 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	OpeningSettlementDate *DateAndDateTimeChoice `xml:"OpngSttlmDt"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting2Choice `xml:"Rptg,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition1Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity1Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator1Choice `xml:"TradOrgtrRole,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds or to sell the indicated currency in order to obtain the necessary currency to fund the transaction.
	CurrencyToBuyOrSell *CurrencyToBuyOrSell1Choice `xml:"CcyToBuyOrSell,omitempty"`

	// Status of affirmation of a trade.
	AffirmationStatus *AffirmationStatus1Choice `xml:"AffirmSts,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus1Choice `xml:"MtchgSts,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *Max350Text `xml:"FxAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails37) AddPlaceOfTrade() *MarketIdentification78 {
	s.PlaceOfTrade = new(MarketIdentification78)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails37) SetPlaceOfClearing(value string) {
	s.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (s *SecuritiesTradeDetails37) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails37) AddOpeningSettlementDate() *DateAndDateTimeChoice {
	s.OpeningSettlementDate = new(DateAndDateTimeChoice)
	return s.OpeningSettlementDate
}

func (s *SecuritiesTradeDetails37) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails37) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails37) AddReporting() *Reporting2Choice {
	newValue := new(Reporting2Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails37) AddTradeTransactionCondition() *TradeTransactionCondition1Choice {
	newValue := new(TradeTransactionCondition1Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails37) AddInvestorCapacity() *InvestorCapacity1Choice {
	s.InvestorCapacity = new(InvestorCapacity1Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails37) AddTradeOriginatorRole() *TradeOriginator1Choice {
	s.TradeOriginatorRole = new(TradeOriginator1Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails37) AddCurrencyToBuyOrSell() *CurrencyToBuyOrSell1Choice {
	s.CurrencyToBuyOrSell = new(CurrencyToBuyOrSell1Choice)
	return s.CurrencyToBuyOrSell
}

func (s *SecuritiesTradeDetails37) AddAffirmationStatus() *AffirmationStatus1Choice {
	s.AffirmationStatus = new(AffirmationStatus1Choice)
	return s.AffirmationStatus
}

func (s *SecuritiesTradeDetails37) AddMatchingStatus() *MatchingStatus1Choice {
	s.MatchingStatus = new(MatchingStatus1Choice)
	return s.MatchingStatus
}

func (s *SecuritiesTradeDetails37) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

func (s *SecuritiesTradeDetails37) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*Max350Text)(&value)
}

// Details of the securities trade.
type SecuritiesTradeDetails38 struct {

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification78 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, ie, payment is effected and securities are delivered.
	EffectiveSettlementDate *SettlementDate3Choice `xml:"FctvSttlmDt"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting2Choice `xml:"Rptg,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition1Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity1Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator1Choice `xml:"TradOrgtrRole,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *Max350Text `xml:"FxAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails38) AddPlaceOfTrade() *MarketIdentification78 {
	s.PlaceOfTrade = new(MarketIdentification78)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails38) SetPlaceOfClearing(value string) {
	s.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (s *SecuritiesTradeDetails38) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails38) AddSettlementDate() *SettlementDate1Choice {
	s.SettlementDate = new(SettlementDate1Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails38) AddEffectiveSettlementDate() *SettlementDate3Choice {
	s.EffectiveSettlementDate = new(SettlementDate3Choice)
	return s.EffectiveSettlementDate
}

func (s *SecuritiesTradeDetails38) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails38) AddReporting() *Reporting2Choice {
	newValue := new(Reporting2Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails38) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails38) AddTradeTransactionCondition() *TradeTransactionCondition1Choice {
	newValue := new(TradeTransactionCondition1Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails38) AddInvestorCapacity() *InvestorCapacity1Choice {
	s.InvestorCapacity = new(InvestorCapacity1Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails38) AddTradeOriginatorRole() *TradeOriginator1Choice {
	s.TradeOriginatorRole = new(TradeOriginator1Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails38) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

func (s *SecuritiesTradeDetails38) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*Max350Text)(&value)
}

// Details of the securities trade.
type SecuritiesTradeDetails4 struct {

	// Identification of an account owner transaction that could potentially match with the allegement notified.
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*Max35Text `xml:"CollTxId,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Status of the allegement.
	Status *AllegementStatus1Choice `xml:"Sts,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification11 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes8 `xml:"FinInstrmAttrbts,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition1Choice `xml:"TradTxCond,omitempty"`

	// Specifies additional information relative to the processing of the trade.
	OpeningClosing *OpeningClosing1Choice `xml:"OpngClsg,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice3Choice `xml:"TpOfPric,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting1Choice `xml:"Rptg,omitempty"`

	// Details about the financial instrument quantity and the account involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount8 `xml:"QtyAndAcctDtls"`

	// Details of the closing of the securities financing transaction.
	SecuritiesFinancingDetails *SecuritiesFinancingTransactionDetails1 `xml:"SctiesFincgDtls,omitempty"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails5 `xml:"SttlmParams"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties5 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties5 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *CashParties3 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection2 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts3 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties3 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension2 `xml:"Xtnsn,omitempty"`
}

func (s *SecuritiesTradeDetails4) SetAccountOwnerTransactionIdentification(value string) {
	s.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails4) SetAccountServicerTransactionIdentification(value string) {
	s.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails4) SetMarketInfrastructureTransactionIdentification(value string) {
	s.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails4) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails4) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails4) SetPoolIdentification(value string) {
	s.PoolIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails4) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails4) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesTradeDetails4) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesTradeDetails4) AddStatus() *AllegementStatus1Choice {
	s.Status = new(AllegementStatus1Choice)
	return s.Status
}

func (s *SecuritiesTradeDetails4) AddPlaceOfTrade() *MarketIdentification4 {
	s.PlaceOfTrade = new(MarketIdentification4)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails4) SetPlaceOfClearing(value string) {
	s.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (s *SecuritiesTradeDetails4) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails4) AddSettlementDate() *SettlementDate1Choice {
	s.SettlementDate = new(SettlementDate1Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails4) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesTradeDetails4) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails4) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails4) AddFinancialInstrumentIdentification() *SecurityIdentification11 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification11)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesTradeDetails4) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes8 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes8)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesTradeDetails4) AddTradeTransactionCondition() *TradeTransactionCondition1Choice {
	newValue := new(TradeTransactionCondition1Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails4) AddOpeningClosing() *OpeningClosing1Choice {
	s.OpeningClosing = new(OpeningClosing1Choice)
	return s.OpeningClosing
}

func (s *SecuritiesTradeDetails4) AddTypeOfPrice() *TypeOfPrice3Choice {
	s.TypeOfPrice = new(TypeOfPrice3Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails4) AddReporting() *Reporting1Choice {
	newValue := new(Reporting1Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails4) AddQuantityAndAccountDetails() *QuantityAndAccount8 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount8)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesTradeDetails4) AddSecuritiesFinancingDetails() *SecuritiesFinancingTransactionDetails1 {
	s.SecuritiesFinancingDetails = new(SecuritiesFinancingTransactionDetails1)
	return s.SecuritiesFinancingDetails
}

func (s *SecuritiesTradeDetails4) AddSettlementParameters() *SettlementDetails5 {
	s.SettlementParameters = new(SettlementDetails5)
	return s.SettlementParameters
}

func (s *SecuritiesTradeDetails4) AddDeliveringSettlementParties() *SettlementParties5 {
	s.DeliveringSettlementParties = new(SettlementParties5)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesTradeDetails4) AddReceivingSettlementParties() *SettlementParties5 {
	s.ReceivingSettlementParties = new(SettlementParties5)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesTradeDetails4) AddCashParties() *CashParties3 {
	s.CashParties = new(CashParties3)
	return s.CashParties
}

func (s *SecuritiesTradeDetails4) AddSettlementAmount() *AmountAndDirection2 {
	s.SettlementAmount = new(AmountAndDirection2)
	return s.SettlementAmount
}

func (s *SecuritiesTradeDetails4) AddOtherAmounts() *OtherAmounts3 {
	s.OtherAmounts = new(OtherAmounts3)
	return s.OtherAmounts
}

func (s *SecuritiesTradeDetails4) AddOtherBusinessParties() *OtherParties3 {
	s.OtherBusinessParties = new(OtherParties3)
	return s.OtherBusinessParties
}

func (s *SecuritiesTradeDetails4) AddExtension() *Extension2 {
	newValue := new(Extension2)
	s.Extension = append(s.Extension, newValue)
	return newValue
}

// Details of the securities trade.
type SecuritiesTradeDetails48 struct {

	// Unambiguous identification of the transaction as known by the notification receiver.
	NotificationSenderTransactionIdentification *Max35Text `xml:"NtfctnSndrTxId,omitempty"`

	// Unambiguous identification of the transaction as known by the notification receiver.
	NotificationReceiverTransactionIdentification *Max35Text `xml:"NtfctnRcvrTxId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate9Choice `xml:"SttlmDt"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes64 `xml:"FinInstrmAttrbts,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting7Choice `xml:"Rptg,omitempty"`

	// Details about the financial instrument quantity involved in the transfer.
	QuantityDetails *Quantity11 `xml:"QtyDtls"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails100 `xml:"SttlmParams,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties36 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties36 `xml:"RcvgSttlmPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection52 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts29 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties26 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesTradeDetails48) SetNotificationSenderTransactionIdentification(value string) {
	s.NotificationSenderTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails48) SetNotificationReceiverTransactionIdentification(value string) {
	s.NotificationReceiverTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails48) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails48) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesTradeDetails48) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesTradeDetails48) AddTradeDate() *TradeDate5Choice {
	s.TradeDate = new(TradeDate5Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails48) AddSettlementDate() *SettlementDate9Choice {
	s.SettlementDate = new(SettlementDate9Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails48) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails48) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesTradeDetails48) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes64 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes64)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesTradeDetails48) AddReporting() *Reporting7Choice {
	newValue := new(Reporting7Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails48) AddQuantityDetails() *Quantity11 {
	s.QuantityDetails = new(Quantity11)
	return s.QuantityDetails
}

func (s *SecuritiesTradeDetails48) AddSettlementParameters() *SettlementDetails100 {
	s.SettlementParameters = new(SettlementDetails100)
	return s.SettlementParameters
}

func (s *SecuritiesTradeDetails48) AddDeliveringSettlementParties() *SettlementParties36 {
	s.DeliveringSettlementParties = new(SettlementParties36)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesTradeDetails48) AddReceivingSettlementParties() *SettlementParties36 {
	s.ReceivingSettlementParties = new(SettlementParties36)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesTradeDetails48) AddSettlementAmount() *AmountAndDirection52 {
	s.SettlementAmount = new(AmountAndDirection52)
	return s.SettlementAmount
}

func (s *SecuritiesTradeDetails48) AddOtherAmounts() *OtherAmounts29 {
	s.OtherAmounts = new(OtherAmounts29)
	return s.OtherAmounts
}

func (s *SecuritiesTradeDetails48) AddOtherBusinessParties() *OtherParties26 {
	s.OtherBusinessParties = new(OtherParties26)
	return s.OtherBusinessParties
}

func (s *SecuritiesTradeDetails48) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Details of the securities trade.
type SecuritiesTradeDetails49 struct {

	// Identification of an account owner transaction that could potentially match with the allegement notified.
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*Max35Text `xml:"CollTxId,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Status of the allegement.
	Status *AllegementStatus3Choice `xml:"Sts,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification1 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate9Choice `xml:"SttlmDt"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes64 `xml:"FinInstrmAttrbts,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition5Choice `xml:"TradTxCond,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice29Choice `xml:"TpOfPric,omitempty"`

	// Details about the financial instrument quantity and the account involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount45 `xml:"QtyAndAcctDtls"`

	// Details of the closing of the securities financing transaction.
	SecuritiesFinancingDetails *SecuritiesFinancingTransactionDetails29 `xml:"SctiesFincgDtls,omitempty"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails99 `xml:"SttlmParams"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties36 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties36 `xml:"RcvgSttlmPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection48 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts32 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties28 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesTradeDetails49) SetAccountOwnerTransactionIdentification(value string) {
	s.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails49) SetAccountServicerTransactionIdentification(value string) {
	s.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails49) SetMarketInfrastructureTransactionIdentification(value string) {
	s.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails49) SetProcessorTransactionIdentification(value string) {
	s.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails49) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails49) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails49) SetPoolIdentification(value string) {
	s.PoolIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails49) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails49) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesTradeDetails49) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesTradeDetails49) AddStatus() *AllegementStatus3Choice {
	s.Status = new(AllegementStatus3Choice)
	return s.Status
}

func (s *SecuritiesTradeDetails49) AddPlaceOfTrade() *PlaceOfTradeIdentification1 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification1)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails49) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	s.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return s.PlaceOfClearing
}

func (s *SecuritiesTradeDetails49) AddTradeDate() *TradeDate5Choice {
	s.TradeDate = new(TradeDate5Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails49) AddSettlementDate() *SettlementDate9Choice {
	s.SettlementDate = new(SettlementDate9Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails49) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails49) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails49) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesTradeDetails49) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes64 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes64)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesTradeDetails49) AddTradeTransactionCondition() *TradeTransactionCondition5Choice {
	newValue := new(TradeTransactionCondition5Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails49) AddTypeOfPrice() *TypeOfPrice29Choice {
	s.TypeOfPrice = new(TypeOfPrice29Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails49) AddQuantityAndAccountDetails() *QuantityAndAccount45 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount45)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesTradeDetails49) AddSecuritiesFinancingDetails() *SecuritiesFinancingTransactionDetails29 {
	s.SecuritiesFinancingDetails = new(SecuritiesFinancingTransactionDetails29)
	return s.SecuritiesFinancingDetails
}

func (s *SecuritiesTradeDetails49) AddSettlementParameters() *SettlementDetails99 {
	s.SettlementParameters = new(SettlementDetails99)
	return s.SettlementParameters
}

func (s *SecuritiesTradeDetails49) AddDeliveringSettlementParties() *SettlementParties36 {
	s.DeliveringSettlementParties = new(SettlementParties36)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesTradeDetails49) AddReceivingSettlementParties() *SettlementParties36 {
	s.ReceivingSettlementParties = new(SettlementParties36)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesTradeDetails49) AddSettlementAmount() *AmountAndDirection48 {
	s.SettlementAmount = new(AmountAndDirection48)
	return s.SettlementAmount
}

func (s *SecuritiesTradeDetails49) AddOtherAmounts() *OtherAmounts32 {
	s.OtherAmounts = new(OtherAmounts32)
	return s.OtherAmounts
}

func (s *SecuritiesTradeDetails49) AddOtherBusinessParties() *OtherParties28 {
	s.OtherBusinessParties = new(OtherParties28)
	return s.OtherBusinessParties
}

func (s *SecuritiesTradeDetails49) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Details of the securities trade.
type SecuritiesTradeDetails5 struct {

	// Specifies the date/time on which the trade was executed.
	TradeDate *DateAndDateTimeChoice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	OpeningSettlementDate *DateAndDateTimeChoice `xml:"OpngSttlmDt"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	InstructionProcessingAdditionalDetails *Max350Text `xml:"InstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails5) AddTradeDate() *DateAndDateTimeChoice {
	s.TradeDate = new(DateAndDateTimeChoice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails5) AddOpeningSettlementDate() *DateAndDateTimeChoice {
	s.OpeningSettlementDate = new(DateAndDateTimeChoice)
	return s.OpeningSettlementDate
}

func (s *SecuritiesTradeDetails5) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails5) SetInstructionProcessingAdditionalDetails(value string) {
	s.InstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the securities trade.
type SecuritiesTradeDetails50 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*Max35Text `xml:"CollTxId,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification1 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Specifies additional information relative to the processing of the trade.
	OpeningClosing *OpeningClosing3Choice `xml:"OpngClsg,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting6Choice `xml:"Rptg,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition5Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity4Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator3Choice `xml:"TradOrgtrRole,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice29Choice `xml:"TpOfPric,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds or to sell the indicated currency in order to obtain the necessary currency to fund the transaction.
	CurrencyToBuyOrSell *CurrencyToBuyOrSell1Choice `xml:"CcyToBuyOrSell,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus27Choice `xml:"MtchgSts,omitempty"`

	// Status of affirmation of a trade.
	AffirmationStatus *AffirmationStatus8Choice `xml:"AffirmSts,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *Max350Text `xml:"FxAddtlDtls,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails50) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails50) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails50) AddPlaceOfTrade() *PlaceOfTradeIdentification1 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification1)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails50) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	s.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return s.PlaceOfClearing
}

func (s *SecuritiesTradeDetails50) AddTradeDate() *TradeDate5Choice {
	s.TradeDate = new(TradeDate5Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails50) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesTradeDetails50) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails50) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails50) AddOpeningClosing() *OpeningClosing3Choice {
	s.OpeningClosing = new(OpeningClosing3Choice)
	return s.OpeningClosing
}

func (s *SecuritiesTradeDetails50) AddReporting() *Reporting6Choice {
	newValue := new(Reporting6Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails50) AddTradeTransactionCondition() *TradeTransactionCondition5Choice {
	newValue := new(TradeTransactionCondition5Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails50) AddInvestorCapacity() *InvestorCapacity4Choice {
	s.InvestorCapacity = new(InvestorCapacity4Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails50) AddTradeOriginatorRole() *TradeOriginator3Choice {
	s.TradeOriginatorRole = new(TradeOriginator3Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails50) AddTypeOfPrice() *TypeOfPrice29Choice {
	s.TypeOfPrice = new(TypeOfPrice29Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails50) AddCurrencyToBuyOrSell() *CurrencyToBuyOrSell1Choice {
	s.CurrencyToBuyOrSell = new(CurrencyToBuyOrSell1Choice)
	return s.CurrencyToBuyOrSell
}

func (s *SecuritiesTradeDetails50) AddMatchingStatus() *MatchingStatus27Choice {
	s.MatchingStatus = new(MatchingStatus27Choice)
	return s.MatchingStatus
}

func (s *SecuritiesTradeDetails50) AddAffirmationStatus() *AffirmationStatus8Choice {
	s.AffirmationStatus = new(AffirmationStatus8Choice)
	return s.AffirmationStatus
}

func (s *SecuritiesTradeDetails50) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*Max350Text)(&value)
}

func (s *SecuritiesTradeDetails50) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the securities trade.
type SecuritiesTradeDetails51 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*Max35Text `xml:"CollTxId,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification1 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate9Choice `xml:"SttlmDt"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Specifies additional information relative to the processing of the trade.
	OpeningClosing *OpeningClosing3Choice `xml:"OpngClsg,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting6Choice `xml:"Rptg,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition5Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity4Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator3Choice `xml:"TradOrgtrRole,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice29Choice `xml:"TpOfPric,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds or to sell the indicated currency in order to obtain the necessary currency to fund the transaction.
	CurrencyToBuyOrSell *CurrencyToBuyOrSell1Choice `xml:"CcyToBuyOrSell,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus27Choice `xml:"MtchgSts,omitempty"`

	// Status of affirmation of a trade.
	AffirmationStatus *AffirmationStatus8Choice `xml:"AffirmSts,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *Max350Text `xml:"FxAddtlDtls,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails51) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails51) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails51) AddPlaceOfTrade() *PlaceOfTradeIdentification1 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification1)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails51) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	s.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return s.PlaceOfClearing
}

func (s *SecuritiesTradeDetails51) AddTradeDate() *TradeDate5Choice {
	s.TradeDate = new(TradeDate5Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails51) AddSettlementDate() *SettlementDate9Choice {
	s.SettlementDate = new(SettlementDate9Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails51) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesTradeDetails51) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails51) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails51) AddOpeningClosing() *OpeningClosing3Choice {
	s.OpeningClosing = new(OpeningClosing3Choice)
	return s.OpeningClosing
}

func (s *SecuritiesTradeDetails51) AddReporting() *Reporting6Choice {
	newValue := new(Reporting6Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails51) AddTradeTransactionCondition() *TradeTransactionCondition5Choice {
	newValue := new(TradeTransactionCondition5Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails51) AddInvestorCapacity() *InvestorCapacity4Choice {
	s.InvestorCapacity = new(InvestorCapacity4Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails51) AddTradeOriginatorRole() *TradeOriginator3Choice {
	s.TradeOriginatorRole = new(TradeOriginator3Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails51) AddTypeOfPrice() *TypeOfPrice29Choice {
	s.TypeOfPrice = new(TypeOfPrice29Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails51) AddCurrencyToBuyOrSell() *CurrencyToBuyOrSell1Choice {
	s.CurrencyToBuyOrSell = new(CurrencyToBuyOrSell1Choice)
	return s.CurrencyToBuyOrSell
}

func (s *SecuritiesTradeDetails51) AddMatchingStatus() *MatchingStatus27Choice {
	s.MatchingStatus = new(MatchingStatus27Choice)
	return s.MatchingStatus
}

func (s *SecuritiesTradeDetails51) AddAffirmationStatus() *AffirmationStatus8Choice {
	s.AffirmationStatus = new(AffirmationStatus8Choice)
	return s.AffirmationStatus
}

func (s *SecuritiesTradeDetails51) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*Max350Text)(&value)
}

func (s *SecuritiesTradeDetails51) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the securities trade.
type SecuritiesTradeDetails52 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*Max35Text `xml:"CollTxId,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification1 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate9Choice `xml:"SttlmDt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Specifies additional information relative to the processing of the trade.
	OpeningClosing *OpeningClosing3Choice `xml:"OpngClsg,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting6Choice `xml:"Rptg,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition5Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity4Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator3Choice `xml:"TradOrgtrRole,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice29Choice `xml:"TpOfPric,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds or to sell the indicated currency in order to obtain the necessary currency to fund the transaction.
	CurrencyToBuyOrSell *CurrencyToBuyOrSell1Choice `xml:"CcyToBuyOrSell,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus27Choice `xml:"MtchgSts,omitempty"`

	// Status of affirmation of a trade.
	AffirmationStatus *AffirmationStatus8Choice `xml:"AffirmSts,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *Max350Text `xml:"FxAddtlDtls,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails52) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails52) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails52) AddPlaceOfTrade() *PlaceOfTradeIdentification1 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification1)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails52) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	s.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return s.PlaceOfClearing
}

func (s *SecuritiesTradeDetails52) AddTradeDate() *TradeDate5Choice {
	s.TradeDate = new(TradeDate5Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails52) AddSettlementDate() *SettlementDate9Choice {
	s.SettlementDate = new(SettlementDate9Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails52) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesTradeDetails52) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails52) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails52) AddOpeningClosing() *OpeningClosing3Choice {
	s.OpeningClosing = new(OpeningClosing3Choice)
	return s.OpeningClosing
}

func (s *SecuritiesTradeDetails52) AddReporting() *Reporting6Choice {
	newValue := new(Reporting6Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails52) AddTradeTransactionCondition() *TradeTransactionCondition5Choice {
	newValue := new(TradeTransactionCondition5Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails52) AddInvestorCapacity() *InvestorCapacity4Choice {
	s.InvestorCapacity = new(InvestorCapacity4Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails52) AddTradeOriginatorRole() *TradeOriginator3Choice {
	s.TradeOriginatorRole = new(TradeOriginator3Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails52) AddTypeOfPrice() *TypeOfPrice29Choice {
	s.TypeOfPrice = new(TypeOfPrice29Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails52) AddCurrencyToBuyOrSell() *CurrencyToBuyOrSell1Choice {
	s.CurrencyToBuyOrSell = new(CurrencyToBuyOrSell1Choice)
	return s.CurrencyToBuyOrSell
}

func (s *SecuritiesTradeDetails52) AddMatchingStatus() *MatchingStatus27Choice {
	s.MatchingStatus = new(MatchingStatus27Choice)
	return s.MatchingStatus
}

func (s *SecuritiesTradeDetails52) AddAffirmationStatus() *AffirmationStatus8Choice {
	s.AffirmationStatus = new(AffirmationStatus8Choice)
	return s.AffirmationStatus
}

func (s *SecuritiesTradeDetails52) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*Max350Text)(&value)
}

func (s *SecuritiesTradeDetails52) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the securities trade.
type SecuritiesTradeDetails53 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*Max35Text `xml:"CollTxId,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification1 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate9Choice `xml:"SttlmDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, ie, payment is effected and securities are delivered.
	EffectiveSettlementDate *SettlementDate11Choice `xml:"FctvSttlmDt"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Specifies additional information relative to the processing of the trade.
	OpeningClosing *OpeningClosing3Choice `xml:"OpngClsg,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting6Choice `xml:"Rptg,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition5Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity4Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator3Choice `xml:"TradOrgtrRole,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice29Choice `xml:"TpOfPric,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *Max350Text `xml:"FxAddtlDtls,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails53) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails53) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails53) AddPlaceOfTrade() *PlaceOfTradeIdentification1 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification1)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails53) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	s.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return s.PlaceOfClearing
}

func (s *SecuritiesTradeDetails53) AddTradeDate() *TradeDate5Choice {
	s.TradeDate = new(TradeDate5Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails53) AddSettlementDate() *SettlementDate9Choice {
	s.SettlementDate = new(SettlementDate9Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails53) AddEffectiveSettlementDate() *SettlementDate11Choice {
	s.EffectiveSettlementDate = new(SettlementDate11Choice)
	return s.EffectiveSettlementDate
}

func (s *SecuritiesTradeDetails53) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails53) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails53) AddOpeningClosing() *OpeningClosing3Choice {
	s.OpeningClosing = new(OpeningClosing3Choice)
	return s.OpeningClosing
}

func (s *SecuritiesTradeDetails53) AddReporting() *Reporting6Choice {
	newValue := new(Reporting6Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails53) AddTradeTransactionCondition() *TradeTransactionCondition5Choice {
	newValue := new(TradeTransactionCondition5Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails53) AddInvestorCapacity() *InvestorCapacity4Choice {
	s.InvestorCapacity = new(InvestorCapacity4Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails53) AddTradeOriginatorRole() *TradeOriginator3Choice {
	s.TradeOriginatorRole = new(TradeOriginator3Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails53) AddTypeOfPrice() *TypeOfPrice29Choice {
	s.TypeOfPrice = new(TypeOfPrice29Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails53) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*Max350Text)(&value)
}

func (s *SecuritiesTradeDetails53) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the securities trade.
type SecuritiesTradeDetails54 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*Max35Text `xml:"CollTxId,omitempty"`

	// Identification of an account owner transaction that could potentially match with the allegement notified.
	AccountOwnerTransactionIdentification []*Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification1 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate9Choice `xml:"SttlmDt"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition5Choice `xml:"TradTxCond,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice29Choice `xml:"TpOfPric,omitempty"`
}

func (s *SecuritiesTradeDetails54) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails54) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails54) AddAccountOwnerTransactionIdentification(value string) {
	s.AccountOwnerTransactionIdentification = append(s.AccountOwnerTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails54) SetProcessorTransactionIdentification(value string) {
	s.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails54) AddPlaceOfTrade() *PlaceOfTradeIdentification1 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification1)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails54) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	s.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return s.PlaceOfClearing
}

func (s *SecuritiesTradeDetails54) AddTradeDate() *TradeDate5Choice {
	s.TradeDate = new(TradeDate5Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails54) AddSettlementDate() *SettlementDate9Choice {
	s.SettlementDate = new(SettlementDate9Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails54) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails54) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails54) AddTradeTransactionCondition() *TradeTransactionCondition5Choice {
	newValue := new(TradeTransactionCondition5Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails54) AddTypeOfPrice() *TypeOfPrice29Choice {
	s.TypeOfPrice = new(TypeOfPrice29Choice)
	return s.TypeOfPrice
}

// Details of the securities trade.
type SecuritiesTradeDetails55 struct {

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification1 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate9Choice `xml:"SttlmDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, ie, payment is effected and securities are delivered.
	EffectiveSettlementDate *SettlementDate11Choice `xml:"FctvSttlmDt"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting6Choice `xml:"Rptg,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition5Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity4Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator3Choice `xml:"TradOrgtrRole,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *Max350Text `xml:"FxAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails55) AddPlaceOfTrade() *PlaceOfTradeIdentification1 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification1)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails55) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	s.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return s.PlaceOfClearing
}

func (s *SecuritiesTradeDetails55) AddTradeDate() *TradeDate5Choice {
	s.TradeDate = new(TradeDate5Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails55) AddSettlementDate() *SettlementDate9Choice {
	s.SettlementDate = new(SettlementDate9Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails55) AddEffectiveSettlementDate() *SettlementDate11Choice {
	s.EffectiveSettlementDate = new(SettlementDate11Choice)
	return s.EffectiveSettlementDate
}

func (s *SecuritiesTradeDetails55) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails55) AddReporting() *Reporting6Choice {
	newValue := new(Reporting6Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails55) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails55) AddTradeTransactionCondition() *TradeTransactionCondition5Choice {
	newValue := new(TradeTransactionCondition5Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails55) AddInvestorCapacity() *InvestorCapacity4Choice {
	s.InvestorCapacity = new(InvestorCapacity4Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails55) AddTradeOriginatorRole() *TradeOriginator3Choice {
	s.TradeOriginatorRole = new(TradeOriginator3Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails55) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

func (s *SecuritiesTradeDetails55) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*Max350Text)(&value)
}

// Details of the securities trade.
type SecuritiesTradeDetails56 struct {

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification1 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	OpeningSettlementDate *DateAndDateTimeChoice `xml:"OpngSttlmDt"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting6Choice `xml:"Rptg,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition5Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity4Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator3Choice `xml:"TradOrgtrRole,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds or to sell the indicated currency in order to obtain the necessary currency to fund the transaction.
	CurrencyToBuyOrSell *CurrencyToBuyOrSell1Choice `xml:"CcyToBuyOrSell,omitempty"`

	// Status of affirmation of a trade.
	AffirmationStatus *AffirmationStatus8Choice `xml:"AffirmSts,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus27Choice `xml:"MtchgSts,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *Max350Text `xml:"FxAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails56) AddPlaceOfTrade() *PlaceOfTradeIdentification1 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification1)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails56) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	s.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return s.PlaceOfClearing
}

func (s *SecuritiesTradeDetails56) AddTradeDate() *TradeDate5Choice {
	s.TradeDate = new(TradeDate5Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails56) AddOpeningSettlementDate() *DateAndDateTimeChoice {
	s.OpeningSettlementDate = new(DateAndDateTimeChoice)
	return s.OpeningSettlementDate
}

func (s *SecuritiesTradeDetails56) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails56) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails56) AddReporting() *Reporting6Choice {
	newValue := new(Reporting6Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails56) AddTradeTransactionCondition() *TradeTransactionCondition5Choice {
	newValue := new(TradeTransactionCondition5Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails56) AddInvestorCapacity() *InvestorCapacity4Choice {
	s.InvestorCapacity = new(InvestorCapacity4Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails56) AddTradeOriginatorRole() *TradeOriginator3Choice {
	s.TradeOriginatorRole = new(TradeOriginator3Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails56) AddCurrencyToBuyOrSell() *CurrencyToBuyOrSell1Choice {
	s.CurrencyToBuyOrSell = new(CurrencyToBuyOrSell1Choice)
	return s.CurrencyToBuyOrSell
}

func (s *SecuritiesTradeDetails56) AddAffirmationStatus() *AffirmationStatus8Choice {
	s.AffirmationStatus = new(AffirmationStatus8Choice)
	return s.AffirmationStatus
}

func (s *SecuritiesTradeDetails56) AddMatchingStatus() *MatchingStatus27Choice {
	s.MatchingStatus = new(MatchingStatus27Choice)
	return s.MatchingStatus
}

func (s *SecuritiesTradeDetails56) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

func (s *SecuritiesTradeDetails56) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*Max350Text)(&value)
}

// Details of the securities trade.
type SecuritiesTradeDetails57 struct {

	// Unambiguous identification of the transaction as known by the notification receiver.
	NotificationSenderTransactionIdentification *RestrictedFINXMax16Text `xml:"NtfctnSndrTxId,omitempty"`

	// Unambiguous identification of the transaction as known by the notification receiver.
	NotificationReceiverTransactionIdentification *RestrictedFINXMax16Text `xml:"NtfctnRcvrTxId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate12Choice `xml:"SttlmDt"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes78 `xml:"FinInstrmAttrbts,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting8Choice `xml:"Rptg,omitempty"`

	// Details about the financial instrument quantity involved in the transfer.
	QuantityDetails *Quantity12 `xml:"QtyDtls"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails103 `xml:"SttlmParams,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties44 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties44 `xml:"RcvgSttlmPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection57 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts33 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties30 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesTradeDetails57) SetNotificationSenderTransactionIdentification(value string) {
	s.NotificationSenderTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesTradeDetails57) SetNotificationReceiverTransactionIdentification(value string) {
	s.NotificationReceiverTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesTradeDetails57) SetCommonIdentification(value string) {
	s.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesTradeDetails57) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesTradeDetails57) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesTradeDetails57) AddTradeDate() *TradeDate6Choice {
	s.TradeDate = new(TradeDate6Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails57) AddSettlementDate() *SettlementDate12Choice {
	s.SettlementDate = new(SettlementDate12Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails57) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails57) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesTradeDetails57) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes78 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes78)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesTradeDetails57) AddReporting() *Reporting8Choice {
	newValue := new(Reporting8Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails57) AddQuantityDetails() *Quantity12 {
	s.QuantityDetails = new(Quantity12)
	return s.QuantityDetails
}

func (s *SecuritiesTradeDetails57) AddSettlementParameters() *SettlementDetails103 {
	s.SettlementParameters = new(SettlementDetails103)
	return s.SettlementParameters
}

func (s *SecuritiesTradeDetails57) AddDeliveringSettlementParties() *SettlementParties44 {
	s.DeliveringSettlementParties = new(SettlementParties44)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesTradeDetails57) AddReceivingSettlementParties() *SettlementParties44 {
	s.ReceivingSettlementParties = new(SettlementParties44)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesTradeDetails57) AddSettlementAmount() *AmountAndDirection57 {
	s.SettlementAmount = new(AmountAndDirection57)
	return s.SettlementAmount
}

func (s *SecuritiesTradeDetails57) AddOtherAmounts() *OtherAmounts33 {
	s.OtherAmounts = new(OtherAmounts33)
	return s.OtherAmounts
}

func (s *SecuritiesTradeDetails57) AddOtherBusinessParties() *OtherParties30 {
	s.OtherBusinessParties = new(OtherParties30)
	return s.OtherBusinessParties
}

func (s *SecuritiesTradeDetails57) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Details of the securities trade.
type SecuritiesTradeDetails58 struct {

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification2 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate12Choice `xml:"SttlmDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, ie, payment is effected and securities are delivered.
	EffectiveSettlementDate *SettlementDate14Choice `xml:"FctvSttlmDt"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price3 `xml:"DealPric,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting9Choice `xml:"Rptg,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition6Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity5Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator4Choice `xml:"TradOrgtrRole,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *RestrictedFINXMax350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *RestrictedFINXMax350Text `xml:"FxAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails58) AddPlaceOfTrade() *PlaceOfTradeIdentification2 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification2)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails58) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	s.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return s.PlaceOfClearing
}

func (s *SecuritiesTradeDetails58) AddTradeDate() *TradeDate6Choice {
	s.TradeDate = new(TradeDate6Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails58) AddSettlementDate() *SettlementDate12Choice {
	s.SettlementDate = new(SettlementDate12Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails58) AddEffectiveSettlementDate() *SettlementDate14Choice {
	s.EffectiveSettlementDate = new(SettlementDate14Choice)
	return s.EffectiveSettlementDate
}

func (s *SecuritiesTradeDetails58) AddDealPrice() *Price3 {
	s.DealPrice = new(Price3)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails58) AddReporting() *Reporting9Choice {
	newValue := new(Reporting9Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails58) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails58) AddTradeTransactionCondition() *TradeTransactionCondition6Choice {
	newValue := new(TradeTransactionCondition6Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails58) AddInvestorCapacity() *InvestorCapacity5Choice {
	s.InvestorCapacity = new(InvestorCapacity5Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails58) AddTradeOriginatorRole() *TradeOriginator4Choice {
	s.TradeOriginatorRole = new(TradeOriginator4Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails58) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

func (s *SecuritiesTradeDetails58) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

// Details of the securities trade.
type SecuritiesTradeDetails59 struct {

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification2 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	OpeningSettlementDate *DateAndDateTimeChoice `xml:"OpngSttlmDt"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price3 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting9Choice `xml:"Rptg,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition6Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity5Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator4Choice `xml:"TradOrgtrRole,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds or to sell the indicated currency in order to obtain the necessary currency to fund the transaction.
	CurrencyToBuyOrSell *CurrencyToBuyOrSell1Choice `xml:"CcyToBuyOrSell,omitempty"`

	// Status of affirmation of a trade.
	AffirmationStatus *AffirmationStatus9Choice `xml:"AffirmSts,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus28Choice `xml:"MtchgSts,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *RestrictedFINXMax350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *RestrictedFINXMax350Text `xml:"FxAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails59) AddPlaceOfTrade() *PlaceOfTradeIdentification2 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification2)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails59) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	s.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return s.PlaceOfClearing
}

func (s *SecuritiesTradeDetails59) AddTradeDate() *TradeDate6Choice {
	s.TradeDate = new(TradeDate6Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails59) AddOpeningSettlementDate() *DateAndDateTimeChoice {
	s.OpeningSettlementDate = new(DateAndDateTimeChoice)
	return s.OpeningSettlementDate
}

func (s *SecuritiesTradeDetails59) AddDealPrice() *Price3 {
	s.DealPrice = new(Price3)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails59) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails59) AddReporting() *Reporting9Choice {
	newValue := new(Reporting9Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails59) AddTradeTransactionCondition() *TradeTransactionCondition6Choice {
	newValue := new(TradeTransactionCondition6Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails59) AddInvestorCapacity() *InvestorCapacity5Choice {
	s.InvestorCapacity = new(InvestorCapacity5Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails59) AddTradeOriginatorRole() *TradeOriginator4Choice {
	s.TradeOriginatorRole = new(TradeOriginator4Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails59) AddCurrencyToBuyOrSell() *CurrencyToBuyOrSell1Choice {
	s.CurrencyToBuyOrSell = new(CurrencyToBuyOrSell1Choice)
	return s.CurrencyToBuyOrSell
}

func (s *SecuritiesTradeDetails59) AddAffirmationStatus() *AffirmationStatus9Choice {
	s.AffirmationStatus = new(AffirmationStatus9Choice)
	return s.AffirmationStatus
}

func (s *SecuritiesTradeDetails59) AddMatchingStatus() *MatchingStatus28Choice {
	s.MatchingStatus = new(MatchingStatus28Choice)
	return s.MatchingStatus
}

func (s *SecuritiesTradeDetails59) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

func (s *SecuritiesTradeDetails59) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

// Details of the securities trade.
type SecuritiesTradeDetails6 struct {

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, ie, payment is effected and securities are delivered.
	EffectiveSettlementDate *SettlementDate3Choice `xml:"FctvSttlmDt"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting2Choice `xml:"Rptg,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition1Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity1Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator1Choice `xml:"TradOrgtrRole,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *Max350Text `xml:"FxAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails6) AddPlaceOfTrade() *MarketIdentification4 {
	s.PlaceOfTrade = new(MarketIdentification4)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails6) SetPlaceOfClearing(value string) {
	s.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (s *SecuritiesTradeDetails6) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails6) AddSettlementDate() *SettlementDate1Choice {
	s.SettlementDate = new(SettlementDate1Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails6) AddEffectiveSettlementDate() *SettlementDate3Choice {
	s.EffectiveSettlementDate = new(SettlementDate3Choice)
	return s.EffectiveSettlementDate
}

func (s *SecuritiesTradeDetails6) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails6) AddReporting() *Reporting2Choice {
	newValue := new(Reporting2Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails6) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails6) AddTradeTransactionCondition() *TradeTransactionCondition1Choice {
	newValue := new(TradeTransactionCondition1Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails6) AddInvestorCapacity() *InvestorCapacity1Choice {
	s.InvestorCapacity = new(InvestorCapacity1Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails6) AddTradeOriginatorRole() *TradeOriginator1Choice {
	s.TradeOriginatorRole = new(TradeOriginator1Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails6) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

func (s *SecuritiesTradeDetails6) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*Max350Text)(&value)
}

// Details of the securities trade.
type SecuritiesTradeDetails60 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*RestrictedFINXMax16Text `xml:"TradId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*RestrictedFINXMax16Text `xml:"CollTxId,omitempty"`

	// Identification of an account owner transaction that could potentially match with the allegement notified.
	AccountOwnerTransactionIdentification []*RestrictedFINXMax16Text `xml:"AcctOwnrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *RestrictedFINXMax16Text `xml:"PrcrTxId,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification2 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate12Choice `xml:"SttlmDt"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price3 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition6Choice `xml:"TradTxCond,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice32Choice `xml:"TpOfPric,omitempty"`
}

func (s *SecuritiesTradeDetails60) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (s *SecuritiesTradeDetails60) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (s *SecuritiesTradeDetails60) AddAccountOwnerTransactionIdentification(value string) {
	s.AccountOwnerTransactionIdentification = append(s.AccountOwnerTransactionIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (s *SecuritiesTradeDetails60) SetProcessorTransactionIdentification(value string) {
	s.ProcessorTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesTradeDetails60) AddPlaceOfTrade() *PlaceOfTradeIdentification2 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification2)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails60) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	s.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return s.PlaceOfClearing
}

func (s *SecuritiesTradeDetails60) AddTradeDate() *TradeDate6Choice {
	s.TradeDate = new(TradeDate6Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails60) AddSettlementDate() *SettlementDate12Choice {
	s.SettlementDate = new(SettlementDate12Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails60) AddDealPrice() *Price3 {
	s.DealPrice = new(Price3)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails60) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails60) AddTradeTransactionCondition() *TradeTransactionCondition6Choice {
	newValue := new(TradeTransactionCondition6Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails60) AddTypeOfPrice() *TypeOfPrice32Choice {
	s.TypeOfPrice = new(TypeOfPrice32Choice)
	return s.TypeOfPrice
}

// Details of the securities trade.
type SecuritiesTradeDetails61 struct {

	// Identification of an account owner transaction that could potentially match with the allegement notified.
	AccountOwnerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctOwnrTxId,omitempty"`

	// Identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *RestrictedFINXMax16Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *RestrictedFINXMax16Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*RestrictedFINXMax16Text `xml:"TradId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*RestrictedFINXMax16Text `xml:"CollTxId,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Status of the allegement.
	Status *AllegementStatus4Choice `xml:"Sts,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification2 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate12Choice `xml:"SttlmDt"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price3 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes78 `xml:"FinInstrmAttrbts,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition6Choice `xml:"TradTxCond,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice32Choice `xml:"TpOfPric,omitempty"`

	// Details about the financial instrument quantity and the account involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount50 `xml:"QtyAndAcctDtls"`

	// Details of the closing of the securities financing transaction.
	SecuritiesFinancingDetails *SecuritiesFinancingTransactionDetails34 `xml:"SctiesFincgDtls,omitempty"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails108 `xml:"SttlmParams"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties44 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties44 `xml:"RcvgSttlmPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection71 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts36 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties31 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesTradeDetails61) SetAccountOwnerTransactionIdentification(value string) {
	s.AccountOwnerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesTradeDetails61) SetAccountServicerTransactionIdentification(value string) {
	s.AccountServicerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesTradeDetails61) SetMarketInfrastructureTransactionIdentification(value string) {
	s.MarketInfrastructureTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesTradeDetails61) SetProcessorTransactionIdentification(value string) {
	s.ProcessorTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesTradeDetails61) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (s *SecuritiesTradeDetails61) SetCommonIdentification(value string) {
	s.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesTradeDetails61) SetPoolIdentification(value string) {
	s.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesTradeDetails61) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (s *SecuritiesTradeDetails61) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesTradeDetails61) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesTradeDetails61) AddStatus() *AllegementStatus4Choice {
	s.Status = new(AllegementStatus4Choice)
	return s.Status
}

func (s *SecuritiesTradeDetails61) AddPlaceOfTrade() *PlaceOfTradeIdentification2 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification2)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails61) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	s.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return s.PlaceOfClearing
}

func (s *SecuritiesTradeDetails61) AddTradeDate() *TradeDate6Choice {
	s.TradeDate = new(TradeDate6Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails61) AddSettlementDate() *SettlementDate12Choice {
	s.SettlementDate = new(SettlementDate12Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails61) AddDealPrice() *Price3 {
	s.DealPrice = new(Price3)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails61) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails61) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesTradeDetails61) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes78 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes78)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesTradeDetails61) AddTradeTransactionCondition() *TradeTransactionCondition6Choice {
	newValue := new(TradeTransactionCondition6Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails61) AddTypeOfPrice() *TypeOfPrice32Choice {
	s.TypeOfPrice = new(TypeOfPrice32Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails61) AddQuantityAndAccountDetails() *QuantityAndAccount50 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount50)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesTradeDetails61) AddSecuritiesFinancingDetails() *SecuritiesFinancingTransactionDetails34 {
	s.SecuritiesFinancingDetails = new(SecuritiesFinancingTransactionDetails34)
	return s.SecuritiesFinancingDetails
}

func (s *SecuritiesTradeDetails61) AddSettlementParameters() *SettlementDetails108 {
	s.SettlementParameters = new(SettlementDetails108)
	return s.SettlementParameters
}

func (s *SecuritiesTradeDetails61) AddDeliveringSettlementParties() *SettlementParties44 {
	s.DeliveringSettlementParties = new(SettlementParties44)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesTradeDetails61) AddReceivingSettlementParties() *SettlementParties44 {
	s.ReceivingSettlementParties = new(SettlementParties44)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesTradeDetails61) AddSettlementAmount() *AmountAndDirection71 {
	s.SettlementAmount = new(AmountAndDirection71)
	return s.SettlementAmount
}

func (s *SecuritiesTradeDetails61) AddOtherAmounts() *OtherAmounts36 {
	s.OtherAmounts = new(OtherAmounts36)
	return s.OtherAmounts
}

func (s *SecuritiesTradeDetails61) AddOtherBusinessParties() *OtherParties31 {
	s.OtherBusinessParties = new(OtherParties31)
	return s.OtherBusinessParties
}

func (s *SecuritiesTradeDetails61) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Details of the securities trade.
type SecuritiesTradeDetails62 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*RestrictedFINXMax16Text `xml:"TradId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*RestrictedFINXMax16Text `xml:"CollTxId,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification2 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate12Choice `xml:"SttlmDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, ie, payment is effected and securities are delivered.
	EffectiveSettlementDate *SettlementDate14Choice `xml:"FctvSttlmDt"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price3 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Specifies additional information relative to the processing of the trade.
	OpeningClosing *OpeningClosing4Choice `xml:"OpngClsg,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting9Choice `xml:"Rptg,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition6Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity5Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator4Choice `xml:"TradOrgtrRole,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice32Choice `xml:"TpOfPric,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *RestrictedFINXMax350Text `xml:"FxAddtlDtls,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *RestrictedFINXMax350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails62) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (s *SecuritiesTradeDetails62) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (s *SecuritiesTradeDetails62) AddPlaceOfTrade() *PlaceOfTradeIdentification2 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification2)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails62) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	s.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return s.PlaceOfClearing
}

func (s *SecuritiesTradeDetails62) AddTradeDate() *TradeDate6Choice {
	s.TradeDate = new(TradeDate6Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails62) AddSettlementDate() *SettlementDate12Choice {
	s.SettlementDate = new(SettlementDate12Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails62) AddEffectiveSettlementDate() *SettlementDate14Choice {
	s.EffectiveSettlementDate = new(SettlementDate14Choice)
	return s.EffectiveSettlementDate
}

func (s *SecuritiesTradeDetails62) AddDealPrice() *Price3 {
	s.DealPrice = new(Price3)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails62) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails62) AddOpeningClosing() *OpeningClosing4Choice {
	s.OpeningClosing = new(OpeningClosing4Choice)
	return s.OpeningClosing
}

func (s *SecuritiesTradeDetails62) AddReporting() *Reporting9Choice {
	newValue := new(Reporting9Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails62) AddTradeTransactionCondition() *TradeTransactionCondition6Choice {
	newValue := new(TradeTransactionCondition6Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails62) AddInvestorCapacity() *InvestorCapacity5Choice {
	s.InvestorCapacity = new(InvestorCapacity5Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails62) AddTradeOriginatorRole() *TradeOriginator4Choice {
	s.TradeOriginatorRole = new(TradeOriginator4Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails62) AddTypeOfPrice() *TypeOfPrice32Choice {
	s.TypeOfPrice = new(TypeOfPrice32Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails62) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

func (s *SecuritiesTradeDetails62) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

// Details of the securities trade.
type SecuritiesTradeDetails63 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*RestrictedFINXMax16Text `xml:"TradId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*RestrictedFINXMax16Text `xml:"CollTxId,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification2 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate12Choice `xml:"SttlmDt"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price3 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Specifies additional information relative to the processing of the trade.
	OpeningClosing *OpeningClosing4Choice `xml:"OpngClsg,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting9Choice `xml:"Rptg,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition6Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity5Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator4Choice `xml:"TradOrgtrRole,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice32Choice `xml:"TpOfPric,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds or to sell the indicated currency in order to obtain the necessary currency to fund the transaction.
	CurrencyToBuyOrSell *CurrencyToBuyOrSell1Choice `xml:"CcyToBuyOrSell,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus28Choice `xml:"MtchgSts,omitempty"`

	// Status of affirmation of a trade.
	AffirmationStatus *AffirmationStatus9Choice `xml:"AffirmSts,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *RestrictedFINXMax350Text `xml:"FxAddtlDtls,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *RestrictedFINXMax350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails63) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (s *SecuritiesTradeDetails63) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (s *SecuritiesTradeDetails63) AddPlaceOfTrade() *PlaceOfTradeIdentification2 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification2)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails63) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	s.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return s.PlaceOfClearing
}

func (s *SecuritiesTradeDetails63) AddTradeDate() *TradeDate6Choice {
	s.TradeDate = new(TradeDate6Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails63) AddSettlementDate() *SettlementDate12Choice {
	s.SettlementDate = new(SettlementDate12Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails63) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesTradeDetails63) AddDealPrice() *Price3 {
	s.DealPrice = new(Price3)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails63) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails63) AddOpeningClosing() *OpeningClosing4Choice {
	s.OpeningClosing = new(OpeningClosing4Choice)
	return s.OpeningClosing
}

func (s *SecuritiesTradeDetails63) AddReporting() *Reporting9Choice {
	newValue := new(Reporting9Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails63) AddTradeTransactionCondition() *TradeTransactionCondition6Choice {
	newValue := new(TradeTransactionCondition6Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails63) AddInvestorCapacity() *InvestorCapacity5Choice {
	s.InvestorCapacity = new(InvestorCapacity5Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails63) AddTradeOriginatorRole() *TradeOriginator4Choice {
	s.TradeOriginatorRole = new(TradeOriginator4Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails63) AddTypeOfPrice() *TypeOfPrice32Choice {
	s.TypeOfPrice = new(TypeOfPrice32Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails63) AddCurrencyToBuyOrSell() *CurrencyToBuyOrSell1Choice {
	s.CurrencyToBuyOrSell = new(CurrencyToBuyOrSell1Choice)
	return s.CurrencyToBuyOrSell
}

func (s *SecuritiesTradeDetails63) AddMatchingStatus() *MatchingStatus28Choice {
	s.MatchingStatus = new(MatchingStatus28Choice)
	return s.MatchingStatus
}

func (s *SecuritiesTradeDetails63) AddAffirmationStatus() *AffirmationStatus9Choice {
	s.AffirmationStatus = new(AffirmationStatus9Choice)
	return s.AffirmationStatus
}

func (s *SecuritiesTradeDetails63) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

func (s *SecuritiesTradeDetails63) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

// Details of the securities trade.
type SecuritiesTradeDetails65 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*RestrictedFINXMax16Text `xml:"TradId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*RestrictedFINXMax16Text `xml:"CollTxId,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification2 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price3 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Specifies additional information relative to the processing of the trade.
	OpeningClosing *OpeningClosing4Choice `xml:"OpngClsg,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting9Choice `xml:"Rptg,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition6Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity5Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator4Choice `xml:"TradOrgtrRole,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice32Choice `xml:"TpOfPric,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds or to sell the indicated currency in order to obtain the necessary currency to fund the transaction.
	CurrencyToBuyOrSell *CurrencyToBuyOrSell1Choice `xml:"CcyToBuyOrSell,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus28Choice `xml:"MtchgSts,omitempty"`

	// Status of affirmation of a trade.
	AffirmationStatus *AffirmationStatus9Choice `xml:"AffirmSts,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *RestrictedFINXMax350Text `xml:"FxAddtlDtls,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *RestrictedFINXMax350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails65) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (s *SecuritiesTradeDetails65) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (s *SecuritiesTradeDetails65) AddPlaceOfTrade() *PlaceOfTradeIdentification2 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification2)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails65) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	s.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return s.PlaceOfClearing
}

func (s *SecuritiesTradeDetails65) AddTradeDate() *TradeDate6Choice {
	s.TradeDate = new(TradeDate6Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails65) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesTradeDetails65) AddDealPrice() *Price3 {
	s.DealPrice = new(Price3)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails65) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails65) AddOpeningClosing() *OpeningClosing4Choice {
	s.OpeningClosing = new(OpeningClosing4Choice)
	return s.OpeningClosing
}

func (s *SecuritiesTradeDetails65) AddReporting() *Reporting9Choice {
	newValue := new(Reporting9Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails65) AddTradeTransactionCondition() *TradeTransactionCondition6Choice {
	newValue := new(TradeTransactionCondition6Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails65) AddInvestorCapacity() *InvestorCapacity5Choice {
	s.InvestorCapacity = new(InvestorCapacity5Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails65) AddTradeOriginatorRole() *TradeOriginator4Choice {
	s.TradeOriginatorRole = new(TradeOriginator4Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails65) AddTypeOfPrice() *TypeOfPrice32Choice {
	s.TypeOfPrice = new(TypeOfPrice32Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails65) AddCurrencyToBuyOrSell() *CurrencyToBuyOrSell1Choice {
	s.CurrencyToBuyOrSell = new(CurrencyToBuyOrSell1Choice)
	return s.CurrencyToBuyOrSell
}

func (s *SecuritiesTradeDetails65) AddMatchingStatus() *MatchingStatus28Choice {
	s.MatchingStatus = new(MatchingStatus28Choice)
	return s.MatchingStatus
}

func (s *SecuritiesTradeDetails65) AddAffirmationStatus() *AffirmationStatus9Choice {
	s.AffirmationStatus = new(AffirmationStatus9Choice)
	return s.AffirmationStatus
}

func (s *SecuritiesTradeDetails65) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

func (s *SecuritiesTradeDetails65) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

// Details of the securities trade.
type SecuritiesTradeDetails66 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*RestrictedFINXMax16Text `xml:"TradId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*RestrictedFINXMax16Text `xml:"CollTxId,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification2 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate12Choice `xml:"SttlmDt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price3 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Specifies additional information relative to the processing of the trade.
	OpeningClosing *OpeningClosing4Choice `xml:"OpngClsg,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting9Choice `xml:"Rptg,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition6Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity5Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator4Choice `xml:"TradOrgtrRole,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice32Choice `xml:"TpOfPric,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds or to sell the indicated currency in order to obtain the necessary currency to fund the transaction.
	CurrencyToBuyOrSell *CurrencyToBuyOrSell1Choice `xml:"CcyToBuyOrSell,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus28Choice `xml:"MtchgSts,omitempty"`

	// Status of affirmation of a trade.
	AffirmationStatus *AffirmationStatus9Choice `xml:"AffirmSts,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *RestrictedFINXMax350Text `xml:"FxAddtlDtls,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *RestrictedFINXMax350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails66) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (s *SecuritiesTradeDetails66) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (s *SecuritiesTradeDetails66) AddPlaceOfTrade() *PlaceOfTradeIdentification2 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification2)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails66) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	s.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return s.PlaceOfClearing
}

func (s *SecuritiesTradeDetails66) AddTradeDate() *TradeDate6Choice {
	s.TradeDate = new(TradeDate6Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails66) AddSettlementDate() *SettlementDate12Choice {
	s.SettlementDate = new(SettlementDate12Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails66) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesTradeDetails66) AddDealPrice() *Price3 {
	s.DealPrice = new(Price3)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails66) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails66) AddOpeningClosing() *OpeningClosing4Choice {
	s.OpeningClosing = new(OpeningClosing4Choice)
	return s.OpeningClosing
}

func (s *SecuritiesTradeDetails66) AddReporting() *Reporting9Choice {
	newValue := new(Reporting9Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails66) AddTradeTransactionCondition() *TradeTransactionCondition6Choice {
	newValue := new(TradeTransactionCondition6Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails66) AddInvestorCapacity() *InvestorCapacity5Choice {
	s.InvestorCapacity = new(InvestorCapacity5Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails66) AddTradeOriginatorRole() *TradeOriginator4Choice {
	s.TradeOriginatorRole = new(TradeOriginator4Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails66) AddTypeOfPrice() *TypeOfPrice32Choice {
	s.TypeOfPrice = new(TypeOfPrice32Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails66) AddCurrencyToBuyOrSell() *CurrencyToBuyOrSell1Choice {
	s.CurrencyToBuyOrSell = new(CurrencyToBuyOrSell1Choice)
	return s.CurrencyToBuyOrSell
}

func (s *SecuritiesTradeDetails66) AddMatchingStatus() *MatchingStatus28Choice {
	s.MatchingStatus = new(MatchingStatus28Choice)
	return s.MatchingStatus
}

func (s *SecuritiesTradeDetails66) AddAffirmationStatus() *AffirmationStatus9Choice {
	s.AffirmationStatus = new(AffirmationStatus9Choice)
	return s.AffirmationStatus
}

func (s *SecuritiesTradeDetails66) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

func (s *SecuritiesTradeDetails66) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

// Details of the securities trade.
type SecuritiesTradeDetails67 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*Max35Text `xml:"CollTxId,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification1 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate9Choice `xml:"SttlmDt"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Time stamp on when the transaction is acknowledged.
	AcknowledgedStatusTimeStamp *ISODateTime `xml:"AckdStsTmStmp,omitempty"`

	// Time stamp on when the transaction is matched.
	MatchedStatusTimeStamp *ISODateTime `xml:"MtchdStsTmStmp,omitempty"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Specifies additional information relative to the processing of the trade.
	OpeningClosing *OpeningClosing3Choice `xml:"OpngClsg,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting6Choice `xml:"Rptg,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition5Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity4Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator3Choice `xml:"TradOrgtrRole,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice29Choice `xml:"TpOfPric,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds or to sell the indicated currency in order to obtain the necessary currency to fund the transaction.
	CurrencyToBuyOrSell *CurrencyToBuyOrSell1Choice `xml:"CcyToBuyOrSell,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus27Choice `xml:"MtchgSts,omitempty"`

	// Status of affirmation of a trade.
	AffirmationStatus *AffirmationStatus8Choice `xml:"AffirmSts,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *Max350Text `xml:"FxAddtlDtls,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails67) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails67) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails67) AddPlaceOfTrade() *PlaceOfTradeIdentification1 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification1)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails67) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	s.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return s.PlaceOfClearing
}

func (s *SecuritiesTradeDetails67) AddTradeDate() *TradeDate5Choice {
	s.TradeDate = new(TradeDate5Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails67) AddSettlementDate() *SettlementDate9Choice {
	s.SettlementDate = new(SettlementDate9Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails67) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesTradeDetails67) SetAcknowledgedStatusTimeStamp(value string) {
	s.AcknowledgedStatusTimeStamp = (*ISODateTime)(&value)
}

func (s *SecuritiesTradeDetails67) SetMatchedStatusTimeStamp(value string) {
	s.MatchedStatusTimeStamp = (*ISODateTime)(&value)
}

func (s *SecuritiesTradeDetails67) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails67) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails67) AddOpeningClosing() *OpeningClosing3Choice {
	s.OpeningClosing = new(OpeningClosing3Choice)
	return s.OpeningClosing
}

func (s *SecuritiesTradeDetails67) AddReporting() *Reporting6Choice {
	newValue := new(Reporting6Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails67) AddTradeTransactionCondition() *TradeTransactionCondition5Choice {
	newValue := new(TradeTransactionCondition5Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails67) AddInvestorCapacity() *InvestorCapacity4Choice {
	s.InvestorCapacity = new(InvestorCapacity4Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails67) AddTradeOriginatorRole() *TradeOriginator3Choice {
	s.TradeOriginatorRole = new(TradeOriginator3Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails67) AddTypeOfPrice() *TypeOfPrice29Choice {
	s.TypeOfPrice = new(TypeOfPrice29Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails67) AddCurrencyToBuyOrSell() *CurrencyToBuyOrSell1Choice {
	s.CurrencyToBuyOrSell = new(CurrencyToBuyOrSell1Choice)
	return s.CurrencyToBuyOrSell
}

func (s *SecuritiesTradeDetails67) AddMatchingStatus() *MatchingStatus27Choice {
	s.MatchingStatus = new(MatchingStatus27Choice)
	return s.MatchingStatus
}

func (s *SecuritiesTradeDetails67) AddAffirmationStatus() *AffirmationStatus8Choice {
	s.AffirmationStatus = new(AffirmationStatus8Choice)
	return s.AffirmationStatus
}

func (s *SecuritiesTradeDetails67) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*Max350Text)(&value)
}

func (s *SecuritiesTradeDetails67) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the securities trade.
type SecuritiesTradeDetails68 struct {

	// Identification of an account owner transaction that could potentially match with the allegement notified.
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*Max35Text `xml:"CollTxId,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Status of the allegement.
	Status *AllegementStatus3Choice `xml:"Sts,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification1 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate9Choice `xml:"SttlmDt"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes64 `xml:"FinInstrmAttrbts,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition5Choice `xml:"TradTxCond,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice29Choice `xml:"TpOfPric,omitempty"`

	// Details about the financial instrument quantity and the account involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount45 `xml:"QtyAndAcctDtls"`

	// Details of the closing of the securities financing transaction.
	SecuritiesFinancingDetails *SecuritiesFinancingTransactionDetails29 `xml:"SctiesFincgDtls,omitempty"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails125 `xml:"SttlmParams"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties36 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties36 `xml:"RcvgSttlmPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection48 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts32 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties28 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesTradeDetails68) SetAccountOwnerTransactionIdentification(value string) {
	s.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails68) SetAccountServicerTransactionIdentification(value string) {
	s.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails68) SetMarketInfrastructureTransactionIdentification(value string) {
	s.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails68) SetProcessorTransactionIdentification(value string) {
	s.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails68) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails68) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails68) SetPoolIdentification(value string) {
	s.PoolIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails68) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails68) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesTradeDetails68) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesTradeDetails68) AddStatus() *AllegementStatus3Choice {
	s.Status = new(AllegementStatus3Choice)
	return s.Status
}

func (s *SecuritiesTradeDetails68) AddPlaceOfTrade() *PlaceOfTradeIdentification1 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification1)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails68) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	s.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return s.PlaceOfClearing
}

func (s *SecuritiesTradeDetails68) AddTradeDate() *TradeDate5Choice {
	s.TradeDate = new(TradeDate5Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails68) AddSettlementDate() *SettlementDate9Choice {
	s.SettlementDate = new(SettlementDate9Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails68) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails68) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails68) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesTradeDetails68) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes64 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes64)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesTradeDetails68) AddTradeTransactionCondition() *TradeTransactionCondition5Choice {
	newValue := new(TradeTransactionCondition5Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails68) AddTypeOfPrice() *TypeOfPrice29Choice {
	s.TypeOfPrice = new(TypeOfPrice29Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails68) AddQuantityAndAccountDetails() *QuantityAndAccount45 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount45)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesTradeDetails68) AddSecuritiesFinancingDetails() *SecuritiesFinancingTransactionDetails29 {
	s.SecuritiesFinancingDetails = new(SecuritiesFinancingTransactionDetails29)
	return s.SecuritiesFinancingDetails
}

func (s *SecuritiesTradeDetails68) AddSettlementParameters() *SettlementDetails125 {
	s.SettlementParameters = new(SettlementDetails125)
	return s.SettlementParameters
}

func (s *SecuritiesTradeDetails68) AddDeliveringSettlementParties() *SettlementParties36 {
	s.DeliveringSettlementParties = new(SettlementParties36)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesTradeDetails68) AddReceivingSettlementParties() *SettlementParties36 {
	s.ReceivingSettlementParties = new(SettlementParties36)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesTradeDetails68) AddSettlementAmount() *AmountAndDirection48 {
	s.SettlementAmount = new(AmountAndDirection48)
	return s.SettlementAmount
}

func (s *SecuritiesTradeDetails68) AddOtherAmounts() *OtherAmounts32 {
	s.OtherAmounts = new(OtherAmounts32)
	return s.OtherAmounts
}

func (s *SecuritiesTradeDetails68) AddOtherBusinessParties() *OtherParties28 {
	s.OtherBusinessParties = new(OtherParties28)
	return s.OtherBusinessParties
}

func (s *SecuritiesTradeDetails68) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Details of the securities trade.
type SecuritiesTradeDetails69 struct {

	// Identification of an account owner transaction that could potentially match with the allegement notified.
	AccountOwnerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctOwnrTxId,omitempty"`

	// Identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *RestrictedFINXMax16Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *RestrictedFINXMax16Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*RestrictedFINXMax16Text `xml:"TradId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*RestrictedFINXMax16Text `xml:"CollTxId,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Status of the allegement.
	Status *AllegementStatus4Choice `xml:"Sts,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification2 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate12Choice `xml:"SttlmDt"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price3 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes78 `xml:"FinInstrmAttrbts,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition6Choice `xml:"TradTxCond,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice32Choice `xml:"TpOfPric,omitempty"`

	// Details about the financial instrument quantity and the account involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount50 `xml:"QtyAndAcctDtls"`

	// Details of the closing of the securities financing transaction.
	SecuritiesFinancingDetails *SecuritiesFinancingTransactionDetails34 `xml:"SctiesFincgDtls,omitempty"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails131 `xml:"SttlmParams"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties44 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties44 `xml:"RcvgSttlmPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection71 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts36 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties31 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesTradeDetails69) SetAccountOwnerTransactionIdentification(value string) {
	s.AccountOwnerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesTradeDetails69) SetAccountServicerTransactionIdentification(value string) {
	s.AccountServicerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesTradeDetails69) SetMarketInfrastructureTransactionIdentification(value string) {
	s.MarketInfrastructureTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesTradeDetails69) SetProcessorTransactionIdentification(value string) {
	s.ProcessorTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesTradeDetails69) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (s *SecuritiesTradeDetails69) SetCommonIdentification(value string) {
	s.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesTradeDetails69) SetPoolIdentification(value string) {
	s.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesTradeDetails69) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (s *SecuritiesTradeDetails69) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesTradeDetails69) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesTradeDetails69) AddStatus() *AllegementStatus4Choice {
	s.Status = new(AllegementStatus4Choice)
	return s.Status
}

func (s *SecuritiesTradeDetails69) AddPlaceOfTrade() *PlaceOfTradeIdentification2 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification2)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails69) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	s.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return s.PlaceOfClearing
}

func (s *SecuritiesTradeDetails69) AddTradeDate() *TradeDate6Choice {
	s.TradeDate = new(TradeDate6Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails69) AddSettlementDate() *SettlementDate12Choice {
	s.SettlementDate = new(SettlementDate12Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails69) AddDealPrice() *Price3 {
	s.DealPrice = new(Price3)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails69) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails69) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesTradeDetails69) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes78 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes78)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesTradeDetails69) AddTradeTransactionCondition() *TradeTransactionCondition6Choice {
	newValue := new(TradeTransactionCondition6Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails69) AddTypeOfPrice() *TypeOfPrice32Choice {
	s.TypeOfPrice = new(TypeOfPrice32Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails69) AddQuantityAndAccountDetails() *QuantityAndAccount50 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount50)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesTradeDetails69) AddSecuritiesFinancingDetails() *SecuritiesFinancingTransactionDetails34 {
	s.SecuritiesFinancingDetails = new(SecuritiesFinancingTransactionDetails34)
	return s.SecuritiesFinancingDetails
}

func (s *SecuritiesTradeDetails69) AddSettlementParameters() *SettlementDetails131 {
	s.SettlementParameters = new(SettlementDetails131)
	return s.SettlementParameters
}

func (s *SecuritiesTradeDetails69) AddDeliveringSettlementParties() *SettlementParties44 {
	s.DeliveringSettlementParties = new(SettlementParties44)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesTradeDetails69) AddReceivingSettlementParties() *SettlementParties44 {
	s.ReceivingSettlementParties = new(SettlementParties44)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesTradeDetails69) AddSettlementAmount() *AmountAndDirection71 {
	s.SettlementAmount = new(AmountAndDirection71)
	return s.SettlementAmount
}

func (s *SecuritiesTradeDetails69) AddOtherAmounts() *OtherAmounts36 {
	s.OtherAmounts = new(OtherAmounts36)
	return s.OtherAmounts
}

func (s *SecuritiesTradeDetails69) AddOtherBusinessParties() *OtherParties31 {
	s.OtherBusinessParties = new(OtherParties31)
	return s.OtherBusinessParties
}

func (s *SecuritiesTradeDetails69) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

// Details of the securities trade.
type SecuritiesTradeDetails7 struct {

	// Unambiguous identification of the transaction as known by the notification receiver.
	NotificationSenderTransactionIdentification *Max35Text `xml:"NtfctnSndrTxId,omitempty"`

	// Unambiguous identification of the transaction as known by the notification receiver.
	NotificationReceiverTransactionIdentification *Max35Text `xml:"NtfctnRcvrTxId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification11 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes8 `xml:"FinInstrmAttrbts,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting1Choice `xml:"Rptg,omitempty"`

	// Details about the financial instrument quantity involved in the transfer.
	QuantityDetails *Quantity5 `xml:"QtyDtls"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails4 `xml:"SttlmParams,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties5 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties5 `xml:"RcvgSttlmPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection7 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts2 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties4 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension2 `xml:"Xtnsn,omitempty"`
}

func (s *SecuritiesTradeDetails7) SetNotificationSenderTransactionIdentification(value string) {
	s.NotificationSenderTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails7) SetNotificationReceiverTransactionIdentification(value string) {
	s.NotificationReceiverTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails7) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails7) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesTradeDetails7) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesTradeDetails7) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails7) AddSettlementDate() *SettlementDate1Choice {
	s.SettlementDate = new(SettlementDate1Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails7) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails7) AddFinancialInstrumentIdentification() *SecurityIdentification11 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification11)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesTradeDetails7) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes8 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes8)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesTradeDetails7) AddReporting() *Reporting1Choice {
	newValue := new(Reporting1Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails7) AddQuantityDetails() *Quantity5 {
	s.QuantityDetails = new(Quantity5)
	return s.QuantityDetails
}

func (s *SecuritiesTradeDetails7) AddSettlementParameters() *SettlementDetails4 {
	s.SettlementParameters = new(SettlementDetails4)
	return s.SettlementParameters
}

func (s *SecuritiesTradeDetails7) AddDeliveringSettlementParties() *SettlementParties5 {
	s.DeliveringSettlementParties = new(SettlementParties5)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesTradeDetails7) AddReceivingSettlementParties() *SettlementParties5 {
	s.ReceivingSettlementParties = new(SettlementParties5)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesTradeDetails7) AddSettlementAmount() *AmountAndDirection7 {
	s.SettlementAmount = new(AmountAndDirection7)
	return s.SettlementAmount
}

func (s *SecuritiesTradeDetails7) AddOtherAmounts() *OtherAmounts2 {
	s.OtherAmounts = new(OtherAmounts2)
	return s.OtherAmounts
}

func (s *SecuritiesTradeDetails7) AddOtherBusinessParties() *OtherParties4 {
	s.OtherBusinessParties = new(OtherParties4)
	return s.OtherBusinessParties
}

func (s *SecuritiesTradeDetails7) AddExtension() *Extension2 {
	newValue := new(Extension2)
	s.Extension = append(s.Extension, newValue)
	return newValue
}

// Details of the securities trade.
type SecuritiesTradeDetails70 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*RestrictedFINXMax16Text `xml:"TradId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*RestrictedFINXMax16Text `xml:"CollTxId,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification2 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate12Choice `xml:"SttlmDt"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Time stamp on when the transaction is acknowledged.
	AcknowledgedStatusTimeStamp *ISODateTime `xml:"AckdStsTmStmp,omitempty"`

	// Time stamp on when the transaction is matched.
	MatchedStatusTimeStamp *ISODateTime `xml:"MtchdStsTmStmp,omitempty"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price3 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Specifies additional information relative to the processing of the trade.
	OpeningClosing *OpeningClosing4Choice `xml:"OpngClsg,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting9Choice `xml:"Rptg,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition6Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity5Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator4Choice `xml:"TradOrgtrRole,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice32Choice `xml:"TpOfPric,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds or to sell the indicated currency in order to obtain the necessary currency to fund the transaction.
	CurrencyToBuyOrSell *CurrencyToBuyOrSell1Choice `xml:"CcyToBuyOrSell,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus28Choice `xml:"MtchgSts,omitempty"`

	// Status of affirmation of a trade.
	AffirmationStatus *AffirmationStatus9Choice `xml:"AffirmSts,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *RestrictedFINXMax350Text `xml:"FxAddtlDtls,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *RestrictedFINXMax350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails70) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (s *SecuritiesTradeDetails70) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (s *SecuritiesTradeDetails70) AddPlaceOfTrade() *PlaceOfTradeIdentification2 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification2)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails70) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	s.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return s.PlaceOfClearing
}

func (s *SecuritiesTradeDetails70) AddTradeDate() *TradeDate6Choice {
	s.TradeDate = new(TradeDate6Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails70) AddSettlementDate() *SettlementDate12Choice {
	s.SettlementDate = new(SettlementDate12Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails70) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesTradeDetails70) SetAcknowledgedStatusTimeStamp(value string) {
	s.AcknowledgedStatusTimeStamp = (*ISODateTime)(&value)
}

func (s *SecuritiesTradeDetails70) SetMatchedStatusTimeStamp(value string) {
	s.MatchedStatusTimeStamp = (*ISODateTime)(&value)
}

func (s *SecuritiesTradeDetails70) AddDealPrice() *Price3 {
	s.DealPrice = new(Price3)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails70) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails70) AddOpeningClosing() *OpeningClosing4Choice {
	s.OpeningClosing = new(OpeningClosing4Choice)
	return s.OpeningClosing
}

func (s *SecuritiesTradeDetails70) AddReporting() *Reporting9Choice {
	newValue := new(Reporting9Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails70) AddTradeTransactionCondition() *TradeTransactionCondition6Choice {
	newValue := new(TradeTransactionCondition6Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails70) AddInvestorCapacity() *InvestorCapacity5Choice {
	s.InvestorCapacity = new(InvestorCapacity5Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails70) AddTradeOriginatorRole() *TradeOriginator4Choice {
	s.TradeOriginatorRole = new(TradeOriginator4Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails70) AddTypeOfPrice() *TypeOfPrice32Choice {
	s.TypeOfPrice = new(TypeOfPrice32Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails70) AddCurrencyToBuyOrSell() *CurrencyToBuyOrSell1Choice {
	s.CurrencyToBuyOrSell = new(CurrencyToBuyOrSell1Choice)
	return s.CurrencyToBuyOrSell
}

func (s *SecuritiesTradeDetails70) AddMatchingStatus() *MatchingStatus28Choice {
	s.MatchingStatus = new(MatchingStatus28Choice)
	return s.MatchingStatus
}

func (s *SecuritiesTradeDetails70) AddAffirmationStatus() *AffirmationStatus9Choice {
	s.AffirmationStatus = new(AffirmationStatus9Choice)
	return s.AffirmationStatus
}

func (s *SecuritiesTradeDetails70) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

func (s *SecuritiesTradeDetails70) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

// Details of the securities trade.
type SecuritiesTradeDetails8 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*Max35Text `xml:"CollTxId,omitempty"`

	// Identification of an account owner transaction that could potentially match with the allegement notified.
	AccountOwnerTransactionIdentification []*Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Specifies additional information relative to the processing of the trade.
	OpeningClosing *OpeningClosing1Choice `xml:"OpngClsg,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting1Choice `xml:"Rptg,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition1Choice `xml:"TradTxCond,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice3Choice `xml:"TpOfPric,omitempty"`
}

func (s *SecuritiesTradeDetails8) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails8) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails8) AddAccountOwnerTransactionIdentification(value string) {
	s.AccountOwnerTransactionIdentification = append(s.AccountOwnerTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails8) AddPlaceOfTrade() *MarketIdentification4 {
	s.PlaceOfTrade = new(MarketIdentification4)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails8) SetPlaceOfClearing(value string) {
	s.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (s *SecuritiesTradeDetails8) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails8) AddSettlementDate() *SettlementDate1Choice {
	s.SettlementDate = new(SettlementDate1Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails8) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesTradeDetails8) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails8) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails8) AddOpeningClosing() *OpeningClosing1Choice {
	s.OpeningClosing = new(OpeningClosing1Choice)
	return s.OpeningClosing
}

func (s *SecuritiesTradeDetails8) AddReporting() *Reporting1Choice {
	newValue := new(Reporting1Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails8) AddTradeTransactionCondition() *TradeTransactionCondition1Choice {
	newValue := new(TradeTransactionCondition1Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails8) AddTypeOfPrice() *TypeOfPrice3Choice {
	s.TypeOfPrice = new(TypeOfPrice3Choice)
	return s.TypeOfPrice
}
