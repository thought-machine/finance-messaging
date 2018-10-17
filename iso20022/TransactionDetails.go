package iso20022

// Identifies the details of the transaction.
type TransactionDetails10 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification11 `xml:"FinInstrmId"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection8 `xml:"SttlmAmt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties2 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties2 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification14Choice `xml:"Invstr,omitempty"`
}

func (t *TransactionDetails10) AddFinancialInstrumentIdentification() *SecurityIdentification11 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification11)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails10) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails10) AddSettlementDate() *SettlementDate1Choice {
	t.SettlementDate = new(SettlementDate1Choice)
	return t.SettlementDate
}

func (t *TransactionDetails10) AddSettlementQuantity() *Quantity6Choice {
	t.SettlementQuantity = new(Quantity6Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails10) AddSettlementAmount() *AmountAndDirection8 {
	t.SettlementAmount = new(AmountAndDirection8)
	return t.SettlementAmount
}

func (t *TransactionDetails10) AddDeliveringSettlementParties() *SettlementParties2 {
	t.DeliveringSettlementParties = new(SettlementParties2)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails10) AddReceivingSettlementParties() *SettlementParties2 {
	t.ReceivingSettlementParties = new(SettlementParties2)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails10) AddInvestor() *PartyIdentification14Choice {
	t.Investor = new(PartyIdentification14Choice)
	return t.Investor
}

// Specifies the details of the transaction.
type TransactionDetails100 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*RestrictedFINXMax16Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId,omitempty"`

	// Unique identification identifying the triparty collateral management transaction from the triparty-agent's/service-provider's point of view.
	TripartyAgentServiceProviderCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"TrptyAgtSvcPrvdrCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Unique identification assigned to the instruction by the client.
	ClientCollateralInstructionIdentification *RestrictedFINXMax16Text `xml:"ClntCollInstrId,omitempty"`

	// Unique identification assigned to the instruction by the triparty-agent/service-provider.
	TripartyAgentServiceProviderCollateralInstructionIdentification *RestrictedFINXMax16Text `xml:"TrptyAgtSvcPrvdrCollInstrId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification119 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount27 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification2 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity10Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection67 `xml:"SttlmAmt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Date/time at which the sender expects settlement.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// For against payment transactions, the value date/time at which the Sender expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate15Choice `xml:"SttlmDt"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Time stamp on when the transaction is acknowledged.
	AcknowledgedStatusTimeStamp *ISODateTime `xml:"AckdStsTmStmp,omitempty"`

	// Time stamp on when the transaction is matched.
	MatchedStatusTimeStamp *ISODateTime `xml:"MtchdStsTmStmp,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails133 `xml:"SttlmParams"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties49 `xml:"RcvgSttlmPties,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties49 `xml:"DlvrgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification110 `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentification111 `xml:"QlfdFrgnIntrmy,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *RestrictedFINXMax350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (t *TransactionDetails100) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (t *TransactionDetails100) SetPoolIdentification(value string) {
	t.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionDetails100) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionDetails100) SetTripartyAgentServiceProviderCollateralTransactionIdentification(value string) {
	t.TripartyAgentServiceProviderCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionDetails100) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionDetails100) SetClientCollateralInstructionIdentification(value string) {
	t.ClientCollateralInstructionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionDetails100) SetTripartyAgentServiceProviderCollateralInstructionIdentification(value string) {
	t.TripartyAgentServiceProviderCollateralInstructionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionDetails100) AddAccountOwner() *PartyIdentification119 {
	t.AccountOwner = new(PartyIdentification119)
	return t.AccountOwner
}

func (t *TransactionDetails100) AddSafekeepingAccount() *SecuritiesAccount27 {
	t.SafekeepingAccount = new(SecuritiesAccount27)
	return t.SafekeepingAccount
}

func (t *TransactionDetails100) AddSafekeepingPlace() *SafeKeepingPlace2 {
	t.SafekeepingPlace = new(SafeKeepingPlace2)
	return t.SafekeepingPlace
}

func (t *TransactionDetails100) AddPlaceOfTrade() *PlaceOfTradeIdentification2 {
	t.PlaceOfTrade = new(PlaceOfTradeIdentification2)
	return t.PlaceOfTrade
}

func (t *TransactionDetails100) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	t.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return t.PlaceOfClearing
}

func (t *TransactionDetails100) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails100) AddSettlementQuantity() *Quantity10Choice {
	t.SettlementQuantity = new(Quantity10Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails100) AddSettlementAmount() *AmountAndDirection67 {
	t.SettlementAmount = new(AmountAndDirection67)
	return t.SettlementAmount
}

func (t *TransactionDetails100) AddLateDeliveryDate() *DateAndDateTimeChoice {
	t.LateDeliveryDate = new(DateAndDateTimeChoice)
	return t.LateDeliveryDate
}

func (t *TransactionDetails100) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	t.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return t.ExpectedSettlementDate
}

func (t *TransactionDetails100) AddExpectedValueDate() *DateAndDateTimeChoice {
	t.ExpectedValueDate = new(DateAndDateTimeChoice)
	return t.ExpectedValueDate
}

func (t *TransactionDetails100) AddSettlementDate() *SettlementDate15Choice {
	t.SettlementDate = new(SettlementDate15Choice)
	return t.SettlementDate
}

func (t *TransactionDetails100) AddTradeDate() *TradeDate6Choice {
	t.TradeDate = new(TradeDate6Choice)
	return t.TradeDate
}

func (t *TransactionDetails100) SetAcknowledgedStatusTimeStamp(value string) {
	t.AcknowledgedStatusTimeStamp = (*ISODateTime)(&value)
}

func (t *TransactionDetails100) SetMatchedStatusTimeStamp(value string) {
	t.MatchedStatusTimeStamp = (*ISODateTime)(&value)
}

func (t *TransactionDetails100) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails100) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails100) AddSettlementParameters() *SettlementDetails133 {
	t.SettlementParameters = new(SettlementDetails133)
	return t.SettlementParameters
}

func (t *TransactionDetails100) AddReceivingSettlementParties() *SettlementParties49 {
	t.ReceivingSettlementParties = new(SettlementParties49)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails100) AddDeliveringSettlementParties() *SettlementParties49 {
	t.DeliveringSettlementParties = new(SettlementParties49)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails100) AddInvestor() *PartyIdentification110 {
	t.Investor = new(PartyIdentification110)
	return t.Investor
}

func (t *TransactionDetails100) AddQualifiedForeignIntermediary() *PartyIdentification111 {
	t.QualifiedForeignIntermediary = new(PartyIdentification111)
	return t.QualifiedForeignIntermediary
}

func (t *TransactionDetails100) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	t.SettlementInstructionProcessingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

// Identifies the details of the transaction.
type TransactionDetails101 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *SettlementTypeAndIdentification22 `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *RestrictedFINXMax16Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *RestrictedFINXMax16Text `xml:"PrcrTxId,omitempty"`

	// Identification of a transaction that cannot be identified using a standard reference element present in the message.
	OtherTransactionIdentification *RestrictedFINXMax16Text `xml:"OthrTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification119 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount30 `xml:"SfkpgAcct"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionDetails *TransactionDetails83 `xml:"TxDtls,omitempty"`
}

func (t *TransactionDetails101) AddAccountOwnerTransactionIdentification() *SettlementTypeAndIdentification22 {
	t.AccountOwnerTransactionIdentification = new(SettlementTypeAndIdentification22)
	return t.AccountOwnerTransactionIdentification
}

func (t *TransactionDetails101) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionDetails101) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionDetails101) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionDetails101) SetOtherTransactionIdentification(value string) {
	t.OtherTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionDetails101) AddAccountOwner() *PartyIdentification119 {
	t.AccountOwner = new(PartyIdentification119)
	return t.AccountOwner
}

func (t *TransactionDetails101) AddSafekeepingAccount() *SecuritiesAccount30 {
	t.SafekeepingAccount = new(SecuritiesAccount30)
	return t.SafekeepingAccount
}

func (t *TransactionDetails101) AddTransactionDetails() *TransactionDetails83 {
	t.TransactionDetails = new(TransactionDetails83)
	return t.TransactionDetails
}

// Identifies the details of the transaction.
type TransactionDetails11 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *References2Choice `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails10 `xml:"TxDtls,omitempty"`

	// Specifies whether an associated FX should be cancelled.
	FXCancellation *FXCancellation1Choice `xml:"FxCxl,omitempty"`
}

func (t *TransactionDetails11) AddAccountOwnerTransactionIdentification() *References2Choice {
	t.AccountOwnerTransactionIdentification = new(References2Choice)
	return t.AccountOwnerTransactionIdentification
}

func (t *TransactionDetails11) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails11) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails11) AddAccountOwner() *PartyIdentification13Choice {
	t.AccountOwner = new(PartyIdentification13Choice)
	return t.AccountOwner
}

func (t *TransactionDetails11) AddSafekeepingAccount() *SecuritiesAccount13 {
	t.SafekeepingAccount = new(SecuritiesAccount13)
	return t.SafekeepingAccount
}

func (t *TransactionDetails11) AddTransactionDetails() *TransactionDetails10 {
	t.TransactionDetails = new(TransactionDetails10)
	return t.TransactionDetails
}

func (t *TransactionDetails11) AddFXCancellation() *FXCancellation1Choice {
	t.FXCancellation = new(FXCancellation1Choice)
	return t.FXCancellation
}

// Identifies the details of the transaction.
type TransactionDetails12 struct {

	// Reference to the message advised to be cancelled by the account servicer
	Reference *References3Choice `xml:"Ref"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`
}

func (t *TransactionDetails12) AddReference() *References3Choice {
	t.Reference = new(References3Choice)
	return t.Reference
}

func (t *TransactionDetails12) AddAccountOwner() *PartyIdentification13Choice {
	t.AccountOwner = new(PartyIdentification13Choice)
	return t.AccountOwner
}

func (t *TransactionDetails12) AddSafekeepingAccount() *SecuritiesAccount13 {
	t.SafekeepingAccount = new(SecuritiesAccount13)
	return t.SafekeepingAccount
}

// Identifies the details of the transaction.
type TransactionDetails13 struct {

	// Provides transaction type and identification information.
	AccountServicerTransactionIdentification *SettlementTypeAndIdentification3 `xml:"AcctSvcrTxId"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Identification1 `xml:"MktInfrstrctrTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails10 `xml:"TxDtls,omitempty"`
}

func (t *TransactionDetails13) AddAccountServicerTransactionIdentification() *SettlementTypeAndIdentification3 {
	t.AccountServicerTransactionIdentification = new(SettlementTypeAndIdentification3)
	return t.AccountServicerTransactionIdentification
}

func (t *TransactionDetails13) AddMarketInfrastructureTransactionIdentification() *Identification1 {
	t.MarketInfrastructureTransactionIdentification = new(Identification1)
	return t.MarketInfrastructureTransactionIdentification
}

func (t *TransactionDetails13) AddAccountOwner() *PartyIdentification13Choice {
	t.AccountOwner = new(PartyIdentification13Choice)
	return t.AccountOwner
}

func (t *TransactionDetails13) AddSafekeepingAccount() *SecuritiesAccount13 {
	t.SafekeepingAccount = new(SecuritiesAccount13)
	return t.SafekeepingAccount
}

func (t *TransactionDetails13) AddTransactionDetails() *TransactionDetails10 {
	t.TransactionDetails = new(TransactionDetails10)
	return t.TransactionDetails
}

// Reference of the trade transaction.
type TransactionDetails2 struct {

	// Unique identification assigned to a trade once it is received or matched by an executing system.
	TradeReference *Max70Text `xml:"TradRef"`
}

func (t *TransactionDetails2) SetTradeReference(value string) {
	t.TradeReference = (*Max70Text)(&value)
}

// Identifies the details of the transaction.
type TransactionDetails22 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection8 `xml:"SttlmAmt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// For against payment transactions, the value date/time at which the Sender expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate2Choice `xml:"SttlmDt"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails23 `xml:"SttlmParams"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties13 `xml:"RcvgSttlmPties,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties13 `xml:"DlvrgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification37Choice `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentification45Choice `xml:"QlfdFrgnIntrmy,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (t *TransactionDetails22) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *TransactionDetails22) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails22) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails22) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails22) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails22) AddAccountOwner() *PartyIdentification36Choice {
	t.AccountOwner = new(PartyIdentification36Choice)
	return t.AccountOwner
}

func (t *TransactionDetails22) AddSafekeepingAccount() *SecuritiesAccount13 {
	t.SafekeepingAccount = new(SecuritiesAccount13)
	return t.SafekeepingAccount
}

func (t *TransactionDetails22) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	t.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return t.SafekeepingPlace
}

func (t *TransactionDetails22) AddPlaceOfTrade() *MarketIdentification4 {
	t.PlaceOfTrade = new(MarketIdentification4)
	return t.PlaceOfTrade
}

func (t *TransactionDetails22) SetPlaceOfClearing(value string) {
	t.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (t *TransactionDetails22) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails22) AddSettlementQuantity() *Quantity6Choice {
	t.SettlementQuantity = new(Quantity6Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails22) AddSettlementAmount() *AmountAndDirection8 {
	t.SettlementAmount = new(AmountAndDirection8)
	return t.SettlementAmount
}

func (t *TransactionDetails22) AddLateDeliveryDate() *DateAndDateTimeChoice {
	t.LateDeliveryDate = new(DateAndDateTimeChoice)
	return t.LateDeliveryDate
}

func (t *TransactionDetails22) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	t.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return t.ExpectedSettlementDate
}

func (t *TransactionDetails22) AddExpectedValueDate() *DateAndDateTimeChoice {
	t.ExpectedValueDate = new(DateAndDateTimeChoice)
	return t.ExpectedValueDate
}

func (t *TransactionDetails22) AddSettlementDate() *SettlementDate2Choice {
	t.SettlementDate = new(SettlementDate2Choice)
	return t.SettlementDate
}

func (t *TransactionDetails22) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails22) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails22) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails22) AddSettlementParameters() *SettlementDetails23 {
	t.SettlementParameters = new(SettlementDetails23)
	return t.SettlementParameters
}

func (t *TransactionDetails22) AddReceivingSettlementParties() *SettlementParties13 {
	t.ReceivingSettlementParties = new(SettlementParties13)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails22) AddDeliveringSettlementParties() *SettlementParties13 {
	t.DeliveringSettlementParties = new(SettlementParties13)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails22) AddInvestor() *PartyIdentification37Choice {
	t.Investor = new(PartyIdentification37Choice)
	return t.Investor
}

func (t *TransactionDetails22) AddQualifiedForeignIntermediary() *PartyIdentification45Choice {
	t.QualifiedForeignIntermediary = new(PartyIdentification45Choice)
	return t.QualifiedForeignIntermediary
}

func (t *TransactionDetails22) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	t.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Identifies the details of the transaction.
type TransactionDetails24 struct {

	// Specifies the type of activity to which this instruction relates.
	TransactionActivity *TransactionActivity1Choice `xml:"TxActvty"`

	// Choice of type for the transaction reported.
	SettlementTransactionOrCorporateActionEventType *SettlementOrCorporateActionEvent1Choice `xml:"SttlmTxOrCorpActnEvtTp,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails27 `xml:"SttlmParams,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Quantity of financial instrument (to be) posted to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *AmountAndDirection3 `xml:"PstngAmt,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection4 `xml:"AcrdIntrstAmt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, that is, payment is effected and securities are delivered.
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt,omitempty"`

	// Date and time assets become available to the account owner (in a credit entry), or cease to be available to the account owner (in a debit entry).
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties13 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties13 `xml:"RcvgSttlmPties,omitempty"`

	// Specifies whether it is the reversal of a previously reported movement.
	ReversalIndicator *YesNoIndicator `xml:"RvslInd,omitempty"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionAdditionalDetails *Max350Text `xml:"TxAddtlDtls,omitempty"`
}

func (t *TransactionDetails24) AddTransactionActivity() *TransactionActivity1Choice {
	t.TransactionActivity = new(TransactionActivity1Choice)
	return t.TransactionActivity
}

func (t *TransactionDetails24) AddSettlementTransactionOrCorporateActionEventType() *SettlementOrCorporateActionEvent1Choice {
	t.SettlementTransactionOrCorporateActionEventType = new(SettlementOrCorporateActionEvent1Choice)
	return t.SettlementTransactionOrCorporateActionEventType
}

func (t *TransactionDetails24) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails24) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails24) AddSettlementParameters() *SettlementDetails27 {
	t.SettlementParameters = new(SettlementDetails27)
	return t.SettlementParameters
}

func (t *TransactionDetails24) AddPlaceOfTrade() *MarketIdentification4 {
	t.PlaceOfTrade = new(MarketIdentification4)
	return t.PlaceOfTrade
}

func (t *TransactionDetails24) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	t.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return t.SafekeepingPlace
}

func (t *TransactionDetails24) SetPlaceOfClearing(value string) {
	t.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (t *TransactionDetails24) AddPostingQuantity() *Quantity6Choice {
	t.PostingQuantity = new(Quantity6Choice)
	return t.PostingQuantity
}

func (t *TransactionDetails24) SetNumberOfDaysAccrued(value string) {
	t.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (t *TransactionDetails24) AddPostingAmount() *AmountAndDirection3 {
	t.PostingAmount = new(AmountAndDirection3)
	return t.PostingAmount
}

func (t *TransactionDetails24) AddAccruedInterestAmount() *AmountAndDirection4 {
	t.AccruedInterestAmount = new(AmountAndDirection4)
	return t.AccruedInterestAmount
}

func (t *TransactionDetails24) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails24) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	t.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return t.EffectiveSettlementDate
}

func (t *TransactionDetails24) AddSettlementDate() *SettlementDate1Choice {
	t.SettlementDate = new(SettlementDate1Choice)
	return t.SettlementDate
}

func (t *TransactionDetails24) AddValueDate() *DateAndDateTimeChoice {
	t.ValueDate = new(DateAndDateTimeChoice)
	return t.ValueDate
}

func (t *TransactionDetails24) AddDeliveringSettlementParties() *SettlementParties13 {
	t.DeliveringSettlementParties = new(SettlementParties13)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails24) AddReceivingSettlementParties() *SettlementParties13 {
	t.ReceivingSettlementParties = new(SettlementParties13)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails24) SetReversalIndicator(value string) {
	t.ReversalIndicator = (*YesNoIndicator)(&value)
}

func (t *TransactionDetails24) SetTransactionAdditionalDetails(value string) {
	t.TransactionAdditionalDetails = (*Max350Text)(&value)
}

// Identifies the details of the transaction.
type TransactionDetails25 struct {

	// Specifies the type of activity to which this instruction relates.
	TransactionActivity *TransactionActivity1Choice `xml:"TxActvty"`

	// Choice of type for the transaction reported.
	SettlementTransactionOrCorporateActionEventType *SettlementOrCorporateActionEvent2Choice `xml:"SttlmTxOrCorpActnEvtTp,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails28 `xml:"SttlmParams,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Quantity of financial instrument (to be) posted to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *AmountAndDirection8 `xml:"PstngAmt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate2Choice `xml:"SttlmDt"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// For against payment transactions, the value date/time at which the account servicer expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties13 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties13 `xml:"RcvgSttlmPties,omitempty"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionAdditionalDetails *Max350Text `xml:"TxAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (t *TransactionDetails25) AddTransactionActivity() *TransactionActivity1Choice {
	t.TransactionActivity = new(TransactionActivity1Choice)
	return t.TransactionActivity
}

func (t *TransactionDetails25) AddSettlementTransactionOrCorporateActionEventType() *SettlementOrCorporateActionEvent2Choice {
	t.SettlementTransactionOrCorporateActionEventType = new(SettlementOrCorporateActionEvent2Choice)
	return t.SettlementTransactionOrCorporateActionEventType
}

func (t *TransactionDetails25) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails25) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails25) AddSettlementParameters() *SettlementDetails28 {
	t.SettlementParameters = new(SettlementDetails28)
	return t.SettlementParameters
}

func (t *TransactionDetails25) AddPlaceOfTrade() *MarketIdentification4 {
	t.PlaceOfTrade = new(MarketIdentification4)
	return t.PlaceOfTrade
}

func (t *TransactionDetails25) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	t.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return t.SafekeepingPlace
}

func (t *TransactionDetails25) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails25) AddPostingQuantity() *Quantity6Choice {
	t.PostingQuantity = new(Quantity6Choice)
	return t.PostingQuantity
}

func (t *TransactionDetails25) AddPostingAmount() *AmountAndDirection8 {
	t.PostingAmount = new(AmountAndDirection8)
	return t.PostingAmount
}

func (t *TransactionDetails25) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails25) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	t.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return t.ExpectedSettlementDate
}

func (t *TransactionDetails25) AddSettlementDate() *SettlementDate2Choice {
	t.SettlementDate = new(SettlementDate2Choice)
	return t.SettlementDate
}

func (t *TransactionDetails25) AddLateDeliveryDate() *DateAndDateTimeChoice {
	t.LateDeliveryDate = new(DateAndDateTimeChoice)
	return t.LateDeliveryDate
}

func (t *TransactionDetails25) AddExpectedValueDate() *DateAndDateTimeChoice {
	t.ExpectedValueDate = new(DateAndDateTimeChoice)
	return t.ExpectedValueDate
}

func (t *TransactionDetails25) AddDeliveringSettlementParties() *SettlementParties13 {
	t.DeliveringSettlementParties = new(SettlementParties13)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails25) AddReceivingSettlementParties() *SettlementParties13 {
	t.ReceivingSettlementParties = new(SettlementParties13)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails25) SetTransactionAdditionalDetails(value string) {
	t.TransactionAdditionalDetails = (*Max350Text)(&value)
}

func (t *TransactionDetails25) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	t.SupplementaryData = append(t.SupplementaryData, newValue)
	return newValue
}

// Identifies the details of the transaction.
type TransactionDetails26 struct {

	// Specifies the type of activity to which this instruction relates.
	TransactionActivity *TransactionActivity1Choice `xml:"TxActvty"`

	// Choice of type for the transaction reported.
	SettlementTransactionOrCorporateActionEventType *SettlementOrCorporateActionEvent2Choice `xml:"SttlmTxOrCorpActnEvtTp,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails28 `xml:"SttlmParams,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Quantity of financial instrument (to be) posted to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *AmountAndDirection8 `xml:"PstngAmt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate2Choice `xml:"SttlmDt"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// For against payment transactions, the value date/time at which the account servicer expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties13 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties13 `xml:"RcvgSttlmPties,omitempty"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionAdditionalDetails *Max350Text `xml:"TxAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (t *TransactionDetails26) AddTransactionActivity() *TransactionActivity1Choice {
	t.TransactionActivity = new(TransactionActivity1Choice)
	return t.TransactionActivity
}

func (t *TransactionDetails26) AddSettlementTransactionOrCorporateActionEventType() *SettlementOrCorporateActionEvent2Choice {
	t.SettlementTransactionOrCorporateActionEventType = new(SettlementOrCorporateActionEvent2Choice)
	return t.SettlementTransactionOrCorporateActionEventType
}

func (t *TransactionDetails26) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails26) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails26) AddSettlementParameters() *SettlementDetails28 {
	t.SettlementParameters = new(SettlementDetails28)
	return t.SettlementParameters
}

func (t *TransactionDetails26) AddPlaceOfTrade() *MarketIdentification4 {
	t.PlaceOfTrade = new(MarketIdentification4)
	return t.PlaceOfTrade
}

func (t *TransactionDetails26) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	t.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return t.SafekeepingPlace
}

func (t *TransactionDetails26) SetPlaceOfClearing(value string) {
	t.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (t *TransactionDetails26) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails26) AddPostingQuantity() *Quantity6Choice {
	t.PostingQuantity = new(Quantity6Choice)
	return t.PostingQuantity
}

func (t *TransactionDetails26) AddPostingAmount() *AmountAndDirection8 {
	t.PostingAmount = new(AmountAndDirection8)
	return t.PostingAmount
}

func (t *TransactionDetails26) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails26) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	t.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return t.ExpectedSettlementDate
}

func (t *TransactionDetails26) AddSettlementDate() *SettlementDate2Choice {
	t.SettlementDate = new(SettlementDate2Choice)
	return t.SettlementDate
}

func (t *TransactionDetails26) AddLateDeliveryDate() *DateAndDateTimeChoice {
	t.LateDeliveryDate = new(DateAndDateTimeChoice)
	return t.LateDeliveryDate
}

func (t *TransactionDetails26) AddExpectedValueDate() *DateAndDateTimeChoice {
	t.ExpectedValueDate = new(DateAndDateTimeChoice)
	return t.ExpectedValueDate
}

func (t *TransactionDetails26) AddDeliveringSettlementParties() *SettlementParties13 {
	t.DeliveringSettlementParties = new(SettlementParties13)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails26) AddReceivingSettlementParties() *SettlementParties13 {
	t.ReceivingSettlementParties = new(SettlementParties13)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails26) SetTransactionAdditionalDetails(value string) {
	t.TransactionAdditionalDetails = (*Max350Text)(&value)
}

func (t *TransactionDetails26) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	t.SupplementaryData = append(t.SupplementaryData, newValue)
	return newValue
}

// Identifies the details of the transaction.
type TransactionDetails28 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection8 `xml:"SttlmAmt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties12 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties12 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification37Choice `xml:"Invstr,omitempty"`
}

func (t *TransactionDetails28) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails28) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails28) AddSettlementDate() *SettlementDate1Choice {
	t.SettlementDate = new(SettlementDate1Choice)
	return t.SettlementDate
}

func (t *TransactionDetails28) AddSettlementQuantity() *Quantity6Choice {
	t.SettlementQuantity = new(Quantity6Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails28) AddSettlementAmount() *AmountAndDirection8 {
	t.SettlementAmount = new(AmountAndDirection8)
	return t.SettlementAmount
}

func (t *TransactionDetails28) AddDeliveringSettlementParties() *SettlementParties12 {
	t.DeliveringSettlementParties = new(SettlementParties12)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails28) AddReceivingSettlementParties() *SettlementParties12 {
	t.ReceivingSettlementParties = new(SettlementParties12)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails28) AddInvestor() *PartyIdentification37Choice {
	t.Investor = new(PartyIdentification37Choice)
	return t.Investor
}

// Identifies the details of the transaction.
type TransactionDetails29 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection8 `xml:"SttlmAmt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties13 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties13 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification37Choice `xml:"Invstr,omitempty"`
}

func (t *TransactionDetails29) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails29) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails29) AddSettlementDate() *SettlementDate1Choice {
	t.SettlementDate = new(SettlementDate1Choice)
	return t.SettlementDate
}

func (t *TransactionDetails29) AddSettlementQuantity() *Quantity6Choice {
	t.SettlementQuantity = new(Quantity6Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails29) AddSettlementAmount() *AmountAndDirection8 {
	t.SettlementAmount = new(AmountAndDirection8)
	return t.SettlementAmount
}

func (t *TransactionDetails29) AddDeliveringSettlementParties() *SettlementParties13 {
	t.DeliveringSettlementParties = new(SettlementParties13)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails29) AddReceivingSettlementParties() *SettlementParties13 {
	t.ReceivingSettlementParties = new(SettlementParties13)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails29) AddInvestor() *PartyIdentification37Choice {
	t.Investor = new(PartyIdentification37Choice)
	return t.Investor
}

// Details of the transaction.
type TransactionDetails3 struct {

	// Unique identification assigned to a trade. This is the reference generated by a firm or the reference allocated by the executing system if the trade was executed automatically.
	TradeReference *Max70Text `xml:"TradRef"`

	// Reference that links to other trades that are/will be sent, eg for straddles where put and call legs need to be reported together.
	AssociatedTradeReference []*Max70Text `xml:"AssoctdTradRef,omitempty"`

	// Identifies the execution venue. In the case of an exchange or a Multilateral Trading Facility (MTF), this should be identified using a MIC code. In the case of a systematic internaliser, place of trade should be identified using a BIC code.
	PlaceOfTrade *PlaceOfTradeIdentification2Choice `xml:"PlcOfTrad"`

	// Specifies the date/time on which the trade was executed.
	TradeDateTime *ISODateTime `xml:"TradDtTm"`

	// Provides details of the financial instrument for which the transaction report is being sent.
	FinancialInstrumentDetails *FinancialInstrument15 `xml:"FinInstrmDtls"`

	// Identifies whether the transaction was a buy or a sell from the perspective of the reporting firm.
	Side *OrderDriverCode `xml:"Sd"`

	// Identifies the regulator(s) to whom the transaction report must be sent.
	TransactionReportMarker []*PartyIdentification24Choice `xml:"TxRptMrkr,omitempty"`

	// Provides details of the counterparty.
	Counterparty *PartyIdentification11Choice `xml:"CtrPty"`

	// Provides details of the client.
	Client *PartyIdentification23 `xml:"Clnt,omitempty"`

	// Identifies the trading capacity of the firm reporting the transaction, eg  Agent or Principal.
	Capacity *TradingCapacity3Code `xml:"Cpcty"`

	// Specifies the currency and price at which the trade has been executed, excluding commission or accrued interest.
	ExecutedTradePrice *PriceRateOrAmountChoice `xml:"ExctdTradPric"`

	// Quantity of financial instrument executed by the trading party.
	ExecutedTradeQuantity *UnitOrFaceAmountChoice `xml:"ExctdTradQty"`

	// The total consideration or value.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Identifies the intended settlement date.
	SettlementDate *ISODateTime `xml:"SttlmDt,omitempty"`

	// Provides details of the person/organisation that has the power of attorney.
	ProxyHolder *PartyIdentification2Choice `xml:"PrxyHldr,omitempty"`

	// Additional domestic regulatory transaction information.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (t *TransactionDetails3) SetTradeReference(value string) {
	t.TradeReference = (*Max70Text)(&value)
}

func (t *TransactionDetails3) AddAssociatedTradeReference(value string) {
	t.AssociatedTradeReference = append(t.AssociatedTradeReference, (*Max70Text)(&value))
}

func (t *TransactionDetails3) AddPlaceOfTrade() *PlaceOfTradeIdentification2Choice {
	t.PlaceOfTrade = new(PlaceOfTradeIdentification2Choice)
	return t.PlaceOfTrade
}

func (t *TransactionDetails3) SetTradeDateTime(value string) {
	t.TradeDateTime = (*ISODateTime)(&value)
}

func (t *TransactionDetails3) AddFinancialInstrumentDetails() *FinancialInstrument15 {
	t.FinancialInstrumentDetails = new(FinancialInstrument15)
	return t.FinancialInstrumentDetails
}

func (t *TransactionDetails3) SetSide(value string) {
	t.Side = (*OrderDriverCode)(&value)
}

func (t *TransactionDetails3) AddTransactionReportMarker() *PartyIdentification24Choice {
	newValue := new(PartyIdentification24Choice)
	t.TransactionReportMarker = append(t.TransactionReportMarker, newValue)
	return newValue
}

func (t *TransactionDetails3) AddCounterparty() *PartyIdentification11Choice {
	t.Counterparty = new(PartyIdentification11Choice)
	return t.Counterparty
}

func (t *TransactionDetails3) AddClient() *PartyIdentification23 {
	t.Client = new(PartyIdentification23)
	return t.Client
}

func (t *TransactionDetails3) SetCapacity(value string) {
	t.Capacity = (*TradingCapacity3Code)(&value)
}

func (t *TransactionDetails3) AddExecutedTradePrice() *PriceRateOrAmountChoice {
	t.ExecutedTradePrice = new(PriceRateOrAmountChoice)
	return t.ExecutedTradePrice
}

func (t *TransactionDetails3) AddExecutedTradeQuantity() *UnitOrFaceAmountChoice {
	t.ExecutedTradeQuantity = new(UnitOrFaceAmountChoice)
	return t.ExecutedTradeQuantity
}

func (t *TransactionDetails3) SetSettlementAmount(value, currency string) {
	t.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (t *TransactionDetails3) SetSettlementDate(value string) {
	t.SettlementDate = (*ISODateTime)(&value)
}

func (t *TransactionDetails3) AddProxyHolder() *PartyIdentification2Choice {
	t.ProxyHolder = new(PartyIdentification2Choice)
	return t.ProxyHolder
}

func (t *TransactionDetails3) SetAdditionalInformation(value string) {
	t.AdditionalInformation = (*Max350Text)(&value)
}

// Identifies the details of the transaction.
type TransactionDetails30 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection8 `xml:"SttlmAmt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate2Choice `xml:"SttlmDt"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties13 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties13 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification37Choice `xml:"Invstr,omitempty"`
}

func (t *TransactionDetails30) AddAccountOwner() *PartyIdentification36Choice {
	t.AccountOwner = new(PartyIdentification36Choice)
	return t.AccountOwner
}

func (t *TransactionDetails30) AddSafekeepingAccount() *SecuritiesAccount13 {
	t.SafekeepingAccount = new(SecuritiesAccount13)
	return t.SafekeepingAccount
}

func (t *TransactionDetails30) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails30) AddSettlementQuantity() *Quantity6Choice {
	t.SettlementQuantity = new(Quantity6Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails30) AddSettlementAmount() *AmountAndDirection8 {
	t.SettlementAmount = new(AmountAndDirection8)
	return t.SettlementAmount
}

func (t *TransactionDetails30) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails30) AddSettlementDate() *SettlementDate2Choice {
	t.SettlementDate = new(SettlementDate2Choice)
	return t.SettlementDate
}

func (t *TransactionDetails30) AddDeliveringSettlementParties() *SettlementParties13 {
	t.DeliveringSettlementParties = new(SettlementParties13)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails30) AddReceivingSettlementParties() *SettlementParties13 {
	t.ReceivingSettlementParties = new(SettlementParties13)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails30) AddInvestor() *PartyIdentification37Choice {
	t.Investor = new(PartyIdentification37Choice)
	return t.Investor
}

// Identifies the details of the transaction.
type TransactionDetails36 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection8 `xml:"SttlmAmt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// For against payment transactions, the value date/time at which the Sender expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate2Choice `xml:"SttlmDt"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails44 `xml:"SttlmParams"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties13 `xml:"RcvgSttlmPties,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties13 `xml:"DlvrgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification37Choice `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentification45Choice `xml:"QlfdFrgnIntrmy,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (t *TransactionDetails36) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *TransactionDetails36) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails36) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails36) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails36) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails36) AddAccountOwner() *PartyIdentification36Choice {
	t.AccountOwner = new(PartyIdentification36Choice)
	return t.AccountOwner
}

func (t *TransactionDetails36) AddSafekeepingAccount() *SecuritiesAccount13 {
	t.SafekeepingAccount = new(SecuritiesAccount13)
	return t.SafekeepingAccount
}

func (t *TransactionDetails36) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	t.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return t.SafekeepingPlace
}

func (t *TransactionDetails36) AddPlaceOfTrade() *MarketIdentification4 {
	t.PlaceOfTrade = new(MarketIdentification4)
	return t.PlaceOfTrade
}

func (t *TransactionDetails36) SetPlaceOfClearing(value string) {
	t.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (t *TransactionDetails36) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails36) AddSettlementQuantity() *Quantity6Choice {
	t.SettlementQuantity = new(Quantity6Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails36) AddSettlementAmount() *AmountAndDirection8 {
	t.SettlementAmount = new(AmountAndDirection8)
	return t.SettlementAmount
}

func (t *TransactionDetails36) AddLateDeliveryDate() *DateAndDateTimeChoice {
	t.LateDeliveryDate = new(DateAndDateTimeChoice)
	return t.LateDeliveryDate
}

func (t *TransactionDetails36) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	t.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return t.ExpectedSettlementDate
}

func (t *TransactionDetails36) AddExpectedValueDate() *DateAndDateTimeChoice {
	t.ExpectedValueDate = new(DateAndDateTimeChoice)
	return t.ExpectedValueDate
}

func (t *TransactionDetails36) AddSettlementDate() *SettlementDate2Choice {
	t.SettlementDate = new(SettlementDate2Choice)
	return t.SettlementDate
}

func (t *TransactionDetails36) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails36) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails36) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails36) AddSettlementParameters() *SettlementDetails44 {
	t.SettlementParameters = new(SettlementDetails44)
	return t.SettlementParameters
}

func (t *TransactionDetails36) AddReceivingSettlementParties() *SettlementParties13 {
	t.ReceivingSettlementParties = new(SettlementParties13)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails36) AddDeliveringSettlementParties() *SettlementParties13 {
	t.DeliveringSettlementParties = new(SettlementParties13)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails36) AddInvestor() *PartyIdentification37Choice {
	t.Investor = new(PartyIdentification37Choice)
	return t.Investor
}

func (t *TransactionDetails36) AddQualifiedForeignIntermediary() *PartyIdentification45Choice {
	t.QualifiedForeignIntermediary = new(PartyIdentification45Choice)
	return t.QualifiedForeignIntermediary
}

func (t *TransactionDetails36) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	t.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Identifies the details of the transaction.
type TransactionDetails37 struct {

	// Specifies the type of activity to which this instruction relates.
	TransactionActivity *TransactionActivity1Choice `xml:"TxActvty"`

	// Choice of type for the transaction reported.
	SettlementTransactionOrCorporateActionEventType *SettlementOrCorporateActionEvent5Choice `xml:"SttlmTxOrCorpActnEvtTp,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails48 `xml:"SttlmParams,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Quantity of financial instrument (to be) posted to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *AmountAndDirection3 `xml:"PstngAmt,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection4 `xml:"AcrdIntrstAmt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, that is, payment is effected and securities are delivered.
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt,omitempty"`

	// Date and time assets become available to the account owner (in a credit entry), or cease to be available to the account owner (in a debit entry).
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties13 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties13 `xml:"RcvgSttlmPties,omitempty"`

	// Specifies whether it is the reversal of a previously reported movement.
	ReversalIndicator *YesNoIndicator `xml:"RvslInd,omitempty"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionAdditionalDetails *Max350Text `xml:"TxAddtlDtls,omitempty"`
}

func (t *TransactionDetails37) AddTransactionActivity() *TransactionActivity1Choice {
	t.TransactionActivity = new(TransactionActivity1Choice)
	return t.TransactionActivity
}

func (t *TransactionDetails37) AddSettlementTransactionOrCorporateActionEventType() *SettlementOrCorporateActionEvent5Choice {
	t.SettlementTransactionOrCorporateActionEventType = new(SettlementOrCorporateActionEvent5Choice)
	return t.SettlementTransactionOrCorporateActionEventType
}

func (t *TransactionDetails37) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails37) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails37) AddSettlementParameters() *SettlementDetails48 {
	t.SettlementParameters = new(SettlementDetails48)
	return t.SettlementParameters
}

func (t *TransactionDetails37) AddPlaceOfTrade() *MarketIdentification4 {
	t.PlaceOfTrade = new(MarketIdentification4)
	return t.PlaceOfTrade
}

func (t *TransactionDetails37) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	t.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return t.SafekeepingPlace
}

func (t *TransactionDetails37) SetPlaceOfClearing(value string) {
	t.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (t *TransactionDetails37) AddPostingQuantity() *Quantity6Choice {
	t.PostingQuantity = new(Quantity6Choice)
	return t.PostingQuantity
}

func (t *TransactionDetails37) SetNumberOfDaysAccrued(value string) {
	t.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (t *TransactionDetails37) AddPostingAmount() *AmountAndDirection3 {
	t.PostingAmount = new(AmountAndDirection3)
	return t.PostingAmount
}

func (t *TransactionDetails37) AddAccruedInterestAmount() *AmountAndDirection4 {
	t.AccruedInterestAmount = new(AmountAndDirection4)
	return t.AccruedInterestAmount
}

func (t *TransactionDetails37) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails37) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	t.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return t.EffectiveSettlementDate
}

func (t *TransactionDetails37) AddSettlementDate() *SettlementDate1Choice {
	t.SettlementDate = new(SettlementDate1Choice)
	return t.SettlementDate
}

func (t *TransactionDetails37) AddValueDate() *DateAndDateTimeChoice {
	t.ValueDate = new(DateAndDateTimeChoice)
	return t.ValueDate
}

func (t *TransactionDetails37) AddDeliveringSettlementParties() *SettlementParties13 {
	t.DeliveringSettlementParties = new(SettlementParties13)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails37) AddReceivingSettlementParties() *SettlementParties13 {
	t.ReceivingSettlementParties = new(SettlementParties13)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails37) SetReversalIndicator(value string) {
	t.ReversalIndicator = (*YesNoIndicator)(&value)
}

func (t *TransactionDetails37) SetTransactionAdditionalDetails(value string) {
	t.TransactionAdditionalDetails = (*Max350Text)(&value)
}

// Identifies the details of the transaction.
type TransactionDetails39 struct {

	// Specifies the type of activity to which this instruction relates.
	TransactionActivity *TransactionActivity1Choice `xml:"TxActvty"`

	// Choice of type for the transaction reported.
	SettlementTransactionOrCorporateActionEventType *SettlementOrCorporateActionEvent6Choice `xml:"SttlmTxOrCorpActnEvtTp,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails47 `xml:"SttlmParams,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Quantity of financial instrument (to be) posted to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *AmountAndDirection8 `xml:"PstngAmt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate2Choice `xml:"SttlmDt"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// For against payment transactions, the value date/time at which the account servicer expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties13 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties13 `xml:"RcvgSttlmPties,omitempty"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionAdditionalDetails *Max350Text `xml:"TxAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (t *TransactionDetails39) AddTransactionActivity() *TransactionActivity1Choice {
	t.TransactionActivity = new(TransactionActivity1Choice)
	return t.TransactionActivity
}

func (t *TransactionDetails39) AddSettlementTransactionOrCorporateActionEventType() *SettlementOrCorporateActionEvent6Choice {
	t.SettlementTransactionOrCorporateActionEventType = new(SettlementOrCorporateActionEvent6Choice)
	return t.SettlementTransactionOrCorporateActionEventType
}

func (t *TransactionDetails39) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails39) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails39) AddSettlementParameters() *SettlementDetails47 {
	t.SettlementParameters = new(SettlementDetails47)
	return t.SettlementParameters
}

func (t *TransactionDetails39) AddPlaceOfTrade() *MarketIdentification4 {
	t.PlaceOfTrade = new(MarketIdentification4)
	return t.PlaceOfTrade
}

func (t *TransactionDetails39) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	t.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return t.SafekeepingPlace
}

func (t *TransactionDetails39) SetPlaceOfClearing(value string) {
	t.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (t *TransactionDetails39) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails39) AddPostingQuantity() *Quantity6Choice {
	t.PostingQuantity = new(Quantity6Choice)
	return t.PostingQuantity
}

func (t *TransactionDetails39) AddPostingAmount() *AmountAndDirection8 {
	t.PostingAmount = new(AmountAndDirection8)
	return t.PostingAmount
}

func (t *TransactionDetails39) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails39) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	t.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return t.ExpectedSettlementDate
}

func (t *TransactionDetails39) AddSettlementDate() *SettlementDate2Choice {
	t.SettlementDate = new(SettlementDate2Choice)
	return t.SettlementDate
}

func (t *TransactionDetails39) AddLateDeliveryDate() *DateAndDateTimeChoice {
	t.LateDeliveryDate = new(DateAndDateTimeChoice)
	return t.LateDeliveryDate
}

func (t *TransactionDetails39) AddExpectedValueDate() *DateAndDateTimeChoice {
	t.ExpectedValueDate = new(DateAndDateTimeChoice)
	return t.ExpectedValueDate
}

func (t *TransactionDetails39) AddDeliveringSettlementParties() *SettlementParties13 {
	t.DeliveringSettlementParties = new(SettlementParties13)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails39) AddReceivingSettlementParties() *SettlementParties13 {
	t.ReceivingSettlementParties = new(SettlementParties13)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails39) SetTransactionAdditionalDetails(value string) {
	t.TransactionAdditionalDetails = (*Max350Text)(&value)
}

func (t *TransactionDetails39) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	t.SupplementaryData = append(t.SupplementaryData, newValue)
	return newValue
}

// Identifies the details of the transaction.
type TransactionDetails4 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification11 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection8 `xml:"SttlmAmt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate2Choice `xml:"SttlmDt"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties2 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties2 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification14Choice `xml:"Invstr,omitempty"`
}

func (t *TransactionDetails4) AddAccountOwner() *PartyIdentification13Choice {
	t.AccountOwner = new(PartyIdentification13Choice)
	return t.AccountOwner
}

func (t *TransactionDetails4) AddSafekeepingAccount() *SecuritiesAccount13 {
	t.SafekeepingAccount = new(SecuritiesAccount13)
	return t.SafekeepingAccount
}

func (t *TransactionDetails4) AddFinancialInstrumentIdentification() *SecurityIdentification11 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification11)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails4) AddSettlementQuantity() *Quantity6Choice {
	t.SettlementQuantity = new(Quantity6Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails4) AddSettlementAmount() *AmountAndDirection8 {
	t.SettlementAmount = new(AmountAndDirection8)
	return t.SettlementAmount
}

func (t *TransactionDetails4) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails4) AddSettlementDate() *SettlementDate2Choice {
	t.SettlementDate = new(SettlementDate2Choice)
	return t.SettlementDate
}

func (t *TransactionDetails4) AddDeliveringSettlementParties() *SettlementParties2 {
	t.DeliveringSettlementParties = new(SettlementParties2)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails4) AddReceivingSettlementParties() *SettlementParties2 {
	t.ReceivingSettlementParties = new(SettlementParties2)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails4) AddInvestor() *PartyIdentification14Choice {
	t.Investor = new(PartyIdentification14Choice)
	return t.Investor
}

// Identifies the details of the transaction.
type TransactionDetails40 struct {

	// Way(s) of identifying the security.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, eg, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection8 `xml:"SttlmAmt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate2Choice `xml:"SttlmDt"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties13 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties13 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification37Choice `xml:"Invstr,omitempty"`
}

func (t *TransactionDetails40) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails40) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails40) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails40) AddSettlementQuantity() *Quantity6Choice {
	t.SettlementQuantity = new(Quantity6Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails40) AddSafekeepingAccount() *SecuritiesAccount13 {
	t.SafekeepingAccount = new(SecuritiesAccount13)
	return t.SafekeepingAccount
}

func (t *TransactionDetails40) AddSettlementAmount() *AmountAndDirection8 {
	t.SettlementAmount = new(AmountAndDirection8)
	return t.SettlementAmount
}

func (t *TransactionDetails40) AddSettlementDate() *SettlementDate2Choice {
	t.SettlementDate = new(SettlementDate2Choice)
	return t.SettlementDate
}

func (t *TransactionDetails40) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails40) AddDeliveringSettlementParties() *SettlementParties13 {
	t.DeliveringSettlementParties = new(SettlementParties13)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails40) AddReceivingSettlementParties() *SettlementParties13 {
	t.ReceivingSettlementParties = new(SettlementParties13)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails40) AddInvestor() *PartyIdentification37Choice {
	t.Investor = new(PartyIdentification37Choice)
	return t.Investor
}

// Identifies the details of the transaction.
type TransactionDetails41 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *SettlementTypeAndIdentification3 `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Identification of a transaction that cannot be identified using a standard reference element present in the message.
	OtherTransactionIdentification *Max35Text `xml:"OthrTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionDetails *TransactionDetails29 `xml:"TxDtls,omitempty"`
}

func (t *TransactionDetails41) AddAccountOwnerTransactionIdentification() *SettlementTypeAndIdentification3 {
	t.AccountOwnerTransactionIdentification = new(SettlementTypeAndIdentification3)
	return t.AccountOwnerTransactionIdentification
}

func (t *TransactionDetails41) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails41) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails41) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails41) SetOtherTransactionIdentification(value string) {
	t.OtherTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails41) AddAccountOwner() *PartyIdentification36Choice {
	t.AccountOwner = new(PartyIdentification36Choice)
	return t.AccountOwner
}

func (t *TransactionDetails41) AddSafekeepingAccount() *SecuritiesAccount13 {
	t.SafekeepingAccount = new(SecuritiesAccount13)
	return t.SafekeepingAccount
}

func (t *TransactionDetails41) AddTransactionDetails() *TransactionDetails29 {
	t.TransactionDetails = new(TransactionDetails29)
	return t.TransactionDetails
}

// Identifies the details of the transaction.
type TransactionDetails45 struct {

	// Way(s) of identifying the security.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, eg, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection8 `xml:"SttlmAmt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate2Choice `xml:"SttlmDt"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties13 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties13 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification37Choice `xml:"Invstr,omitempty"`
}

func (t *TransactionDetails45) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails45) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails45) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails45) AddSettlementQuantity() *Quantity6Choice {
	t.SettlementQuantity = new(Quantity6Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails45) AddSettlementAmount() *AmountAndDirection8 {
	t.SettlementAmount = new(AmountAndDirection8)
	return t.SettlementAmount
}

func (t *TransactionDetails45) AddSettlementDate() *SettlementDate2Choice {
	t.SettlementDate = new(SettlementDate2Choice)
	return t.SettlementDate
}

func (t *TransactionDetails45) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails45) AddDeliveringSettlementParties() *SettlementParties13 {
	t.DeliveringSettlementParties = new(SettlementParties13)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails45) AddReceivingSettlementParties() *SettlementParties13 {
	t.ReceivingSettlementParties = new(SettlementParties13)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails45) AddInvestor() *PartyIdentification37Choice {
	t.Investor = new(PartyIdentification37Choice)
	return t.Investor
}

// Identifies the details of the transaction.
type TransactionDetails5 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification11 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection8 `xml:"SttlmAmt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// For against payment transactions, the value date/time at which the Sender expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate2Choice `xml:"SttlmDt"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails8 `xml:"SttlmParams"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties2 `xml:"RcvgSttlmPties,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties2 `xml:"DlvrgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification14Choice `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentification10Choice `xml:"QlfdFrgnIntrmy,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (t *TransactionDetails5) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *TransactionDetails5) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails5) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails5) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails5) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails5) AddAccountOwner() *PartyIdentification13Choice {
	t.AccountOwner = new(PartyIdentification13Choice)
	return t.AccountOwner
}

func (t *TransactionDetails5) AddSafekeepingAccount() *SecuritiesAccount13 {
	t.SafekeepingAccount = new(SecuritiesAccount13)
	return t.SafekeepingAccount
}

func (t *TransactionDetails5) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	t.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return t.SafekeepingPlace
}

func (t *TransactionDetails5) AddPlaceOfTrade() *MarketIdentification4 {
	t.PlaceOfTrade = new(MarketIdentification4)
	return t.PlaceOfTrade
}

func (t *TransactionDetails5) SetPlaceOfClearing(value string) {
	t.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (t *TransactionDetails5) AddFinancialInstrumentIdentification() *SecurityIdentification11 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification11)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails5) AddSettlementQuantity() *Quantity6Choice {
	t.SettlementQuantity = new(Quantity6Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails5) AddSettlementAmount() *AmountAndDirection8 {
	t.SettlementAmount = new(AmountAndDirection8)
	return t.SettlementAmount
}

func (t *TransactionDetails5) AddLateDeliveryDate() *DateAndDateTimeChoice {
	t.LateDeliveryDate = new(DateAndDateTimeChoice)
	return t.LateDeliveryDate
}

func (t *TransactionDetails5) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	t.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return t.ExpectedSettlementDate
}

func (t *TransactionDetails5) AddExpectedValueDate() *DateAndDateTimeChoice {
	t.ExpectedValueDate = new(DateAndDateTimeChoice)
	return t.ExpectedValueDate
}

func (t *TransactionDetails5) AddSettlementDate() *SettlementDate2Choice {
	t.SettlementDate = new(SettlementDate2Choice)
	return t.SettlementDate
}

func (t *TransactionDetails5) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails5) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails5) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails5) AddSettlementParameters() *SettlementDetails8 {
	t.SettlementParameters = new(SettlementDetails8)
	return t.SettlementParameters
}

func (t *TransactionDetails5) AddReceivingSettlementParties() *SettlementParties2 {
	t.ReceivingSettlementParties = new(SettlementParties2)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails5) AddDeliveringSettlementParties() *SettlementParties2 {
	t.DeliveringSettlementParties = new(SettlementParties2)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails5) AddInvestor() *PartyIdentification14Choice {
	t.Investor = new(PartyIdentification14Choice)
	return t.Investor
}

func (t *TransactionDetails5) AddQualifiedForeignIntermediary() *PartyIdentification10Choice {
	t.QualifiedForeignIntermediary = new(PartyIdentification10Choice)
	return t.QualifiedForeignIntermediary
}

func (t *TransactionDetails5) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	t.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Identifies the details of the transaction.
type TransactionDetails53 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Unique identification assigned to the instruction by the client.
	ClientCollateralInstructionIdentification *Max35Text `xml:"ClntCollInstrId,omitempty"`

	// Unique identification assigned to the instruction by the triparty agent.
	TripartyCollateralInstructionIdentification *Max35Text `xml:"TrptyCollInstrId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection8 `xml:"SttlmAmt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// For against payment transactions, the value date/time at which the Sender expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate2Choice `xml:"SttlmDt"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails44 `xml:"SttlmParams"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties13 `xml:"RcvgSttlmPties,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties13 `xml:"DlvrgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification37Choice `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentification45Choice `xml:"QlfdFrgnIntrmy,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (t *TransactionDetails53) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *TransactionDetails53) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails53) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails53) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails53) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails53) SetClientCollateralInstructionIdentification(value string) {
	t.ClientCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails53) SetTripartyCollateralInstructionIdentification(value string) {
	t.TripartyCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails53) AddAccountOwner() *PartyIdentification36Choice {
	t.AccountOwner = new(PartyIdentification36Choice)
	return t.AccountOwner
}

func (t *TransactionDetails53) AddSafekeepingAccount() *SecuritiesAccount13 {
	t.SafekeepingAccount = new(SecuritiesAccount13)
	return t.SafekeepingAccount
}

func (t *TransactionDetails53) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	t.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return t.SafekeepingPlace
}

func (t *TransactionDetails53) AddPlaceOfTrade() *MarketIdentification4 {
	t.PlaceOfTrade = new(MarketIdentification4)
	return t.PlaceOfTrade
}

func (t *TransactionDetails53) SetPlaceOfClearing(value string) {
	t.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (t *TransactionDetails53) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails53) AddSettlementQuantity() *Quantity6Choice {
	t.SettlementQuantity = new(Quantity6Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails53) AddSettlementAmount() *AmountAndDirection8 {
	t.SettlementAmount = new(AmountAndDirection8)
	return t.SettlementAmount
}

func (t *TransactionDetails53) AddLateDeliveryDate() *DateAndDateTimeChoice {
	t.LateDeliveryDate = new(DateAndDateTimeChoice)
	return t.LateDeliveryDate
}

func (t *TransactionDetails53) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	t.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return t.ExpectedSettlementDate
}

func (t *TransactionDetails53) AddExpectedValueDate() *DateAndDateTimeChoice {
	t.ExpectedValueDate = new(DateAndDateTimeChoice)
	return t.ExpectedValueDate
}

func (t *TransactionDetails53) AddSettlementDate() *SettlementDate2Choice {
	t.SettlementDate = new(SettlementDate2Choice)
	return t.SettlementDate
}

func (t *TransactionDetails53) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails53) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails53) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails53) AddSettlementParameters() *SettlementDetails44 {
	t.SettlementParameters = new(SettlementDetails44)
	return t.SettlementParameters
}

func (t *TransactionDetails53) AddReceivingSettlementParties() *SettlementParties13 {
	t.ReceivingSettlementParties = new(SettlementParties13)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails53) AddDeliveringSettlementParties() *SettlementParties13 {
	t.DeliveringSettlementParties = new(SettlementParties13)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails53) AddInvestor() *PartyIdentification37Choice {
	t.Investor = new(PartyIdentification37Choice)
	return t.Investor
}

func (t *TransactionDetails53) AddQualifiedForeignIntermediary() *PartyIdentification45Choice {
	t.QualifiedForeignIntermediary = new(PartyIdentification45Choice)
	return t.QualifiedForeignIntermediary
}

func (t *TransactionDetails53) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	t.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Identifies the details of the transaction.
type TransactionDetails56 struct {

	// Specifies the type of activity to which this instruction relates.
	TransactionActivity *TransactionActivity1Choice `xml:"TxActvty"`

	// Choice of type for the transaction reported.
	SettlementTransactionOrCorporateActionEventType *SettlementOrCorporateActionEvent10Choice `xml:"SttlmTxOrCorpActnEvtTp,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails47 `xml:"SttlmParams,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Quantity of financial instrument (to be) posted to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *AmountAndDirection8 `xml:"PstngAmt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate2Choice `xml:"SttlmDt"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// For against payment transactions, the value date/time at which the account servicer expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties13 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties13 `xml:"RcvgSttlmPties,omitempty"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionAdditionalDetails *Max350Text `xml:"TxAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (t *TransactionDetails56) AddTransactionActivity() *TransactionActivity1Choice {
	t.TransactionActivity = new(TransactionActivity1Choice)
	return t.TransactionActivity
}

func (t *TransactionDetails56) AddSettlementTransactionOrCorporateActionEventType() *SettlementOrCorporateActionEvent10Choice {
	t.SettlementTransactionOrCorporateActionEventType = new(SettlementOrCorporateActionEvent10Choice)
	return t.SettlementTransactionOrCorporateActionEventType
}

func (t *TransactionDetails56) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails56) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails56) AddSettlementParameters() *SettlementDetails47 {
	t.SettlementParameters = new(SettlementDetails47)
	return t.SettlementParameters
}

func (t *TransactionDetails56) AddPlaceOfTrade() *MarketIdentification4 {
	t.PlaceOfTrade = new(MarketIdentification4)
	return t.PlaceOfTrade
}

func (t *TransactionDetails56) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	t.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return t.SafekeepingPlace
}

func (t *TransactionDetails56) SetPlaceOfClearing(value string) {
	t.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (t *TransactionDetails56) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails56) AddPostingQuantity() *Quantity6Choice {
	t.PostingQuantity = new(Quantity6Choice)
	return t.PostingQuantity
}

func (t *TransactionDetails56) AddPostingAmount() *AmountAndDirection8 {
	t.PostingAmount = new(AmountAndDirection8)
	return t.PostingAmount
}

func (t *TransactionDetails56) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails56) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	t.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return t.ExpectedSettlementDate
}

func (t *TransactionDetails56) AddSettlementDate() *SettlementDate2Choice {
	t.SettlementDate = new(SettlementDate2Choice)
	return t.SettlementDate
}

func (t *TransactionDetails56) AddLateDeliveryDate() *DateAndDateTimeChoice {
	t.LateDeliveryDate = new(DateAndDateTimeChoice)
	return t.LateDeliveryDate
}

func (t *TransactionDetails56) AddExpectedValueDate() *DateAndDateTimeChoice {
	t.ExpectedValueDate = new(DateAndDateTimeChoice)
	return t.ExpectedValueDate
}

func (t *TransactionDetails56) AddDeliveringSettlementParties() *SettlementParties13 {
	t.DeliveringSettlementParties = new(SettlementParties13)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails56) AddReceivingSettlementParties() *SettlementParties13 {
	t.ReceivingSettlementParties = new(SettlementParties13)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails56) SetTransactionAdditionalDetails(value string) {
	t.TransactionAdditionalDetails = (*Max350Text)(&value)
}

func (t *TransactionDetails56) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	t.SupplementaryData = append(t.SupplementaryData, newValue)
	return newValue
}

// Identifies the details of the transaction.
type TransactionDetails57 struct {

	// Specifies the type of activity to which this instruction relates.
	TransactionActivity *TransactionActivity1Choice `xml:"TxActvty"`

	// Choice of type for the transaction reported.
	SettlementTransactionOrCorporateActionEventType *SettlementOrCorporateActionEvent9Choice `xml:"SttlmTxOrCorpActnEvtTp,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails48 `xml:"SttlmParams,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Quantity of financial instrument (to be) posted to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *AmountAndDirection3 `xml:"PstngAmt,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection4 `xml:"AcrdIntrstAmt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, that is, payment is effected and securities are delivered.
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt,omitempty"`

	// Date and time assets become available to the account owner (in a credit entry), or cease to be available to the account owner (in a debit entry).
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties13 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties13 `xml:"RcvgSttlmPties,omitempty"`

	// Specifies whether it is the reversal of a previously reported movement.
	ReversalIndicator *YesNoIndicator `xml:"RvslInd,omitempty"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionAdditionalDetails *Max350Text `xml:"TxAddtlDtls,omitempty"`
}

func (t *TransactionDetails57) AddTransactionActivity() *TransactionActivity1Choice {
	t.TransactionActivity = new(TransactionActivity1Choice)
	return t.TransactionActivity
}

func (t *TransactionDetails57) AddSettlementTransactionOrCorporateActionEventType() *SettlementOrCorporateActionEvent9Choice {
	t.SettlementTransactionOrCorporateActionEventType = new(SettlementOrCorporateActionEvent9Choice)
	return t.SettlementTransactionOrCorporateActionEventType
}

func (t *TransactionDetails57) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails57) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails57) AddSettlementParameters() *SettlementDetails48 {
	t.SettlementParameters = new(SettlementDetails48)
	return t.SettlementParameters
}

func (t *TransactionDetails57) AddPlaceOfTrade() *MarketIdentification4 {
	t.PlaceOfTrade = new(MarketIdentification4)
	return t.PlaceOfTrade
}

func (t *TransactionDetails57) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	t.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return t.SafekeepingPlace
}

func (t *TransactionDetails57) SetPlaceOfClearing(value string) {
	t.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (t *TransactionDetails57) AddPostingQuantity() *Quantity6Choice {
	t.PostingQuantity = new(Quantity6Choice)
	return t.PostingQuantity
}

func (t *TransactionDetails57) SetNumberOfDaysAccrued(value string) {
	t.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (t *TransactionDetails57) AddPostingAmount() *AmountAndDirection3 {
	t.PostingAmount = new(AmountAndDirection3)
	return t.PostingAmount
}

func (t *TransactionDetails57) AddAccruedInterestAmount() *AmountAndDirection4 {
	t.AccruedInterestAmount = new(AmountAndDirection4)
	return t.AccruedInterestAmount
}

func (t *TransactionDetails57) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails57) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	t.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return t.EffectiveSettlementDate
}

func (t *TransactionDetails57) AddSettlementDate() *SettlementDate1Choice {
	t.SettlementDate = new(SettlementDate1Choice)
	return t.SettlementDate
}

func (t *TransactionDetails57) AddValueDate() *DateAndDateTimeChoice {
	t.ValueDate = new(DateAndDateTimeChoice)
	return t.ValueDate
}

func (t *TransactionDetails57) AddDeliveringSettlementParties() *SettlementParties13 {
	t.DeliveringSettlementParties = new(SettlementParties13)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails57) AddReceivingSettlementParties() *SettlementParties13 {
	t.ReceivingSettlementParties = new(SettlementParties13)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails57) SetReversalIndicator(value string) {
	t.ReversalIndicator = (*YesNoIndicator)(&value)
}

func (t *TransactionDetails57) SetTransactionAdditionalDetails(value string) {
	t.TransactionAdditionalDetails = (*Max350Text)(&value)
}

// Identifies the details of the transaction.
type TransactionDetails6 struct {

	// Specifies the type of activity to which this instruction relates.
	TransactionActivity *TransactionActivity1Choice `xml:"TxActvty"`

	// Choice of type for the transaction reported.
	SettlementTransactionOrCorporateActionEventType *SettlementOrCorporateActionEvent1Choice `xml:"SttlmTxOrCorpActnEvtTp,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails2 `xml:"SttlmParams,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Quantity of financial instrument (to be) posted to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *AmountAndDirection3 `xml:"PstngAmt,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection4 `xml:"AcrdIntrstAmt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, ie, payment is effected and securities are delivered.
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt,omitempty"`

	// Date and time assets become available to the account owner (in a credit entry), or cease to be available to the account owner (in a debit entry).
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties2 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties2 `xml:"RcvgSttlmPties,omitempty"`

	// Specifies whether it is the reversal of a previously reported movement.
	ReversalIndicator *YesNoIndicator `xml:"RvslInd,omitempty"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionAdditionalDetails *Max350Text `xml:"TxAddtlDtls,omitempty"`
}

func (t *TransactionDetails6) AddTransactionActivity() *TransactionActivity1Choice {
	t.TransactionActivity = new(TransactionActivity1Choice)
	return t.TransactionActivity
}

func (t *TransactionDetails6) AddSettlementTransactionOrCorporateActionEventType() *SettlementOrCorporateActionEvent1Choice {
	t.SettlementTransactionOrCorporateActionEventType = new(SettlementOrCorporateActionEvent1Choice)
	return t.SettlementTransactionOrCorporateActionEventType
}

func (t *TransactionDetails6) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails6) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails6) AddSettlementParameters() *SettlementDetails2 {
	t.SettlementParameters = new(SettlementDetails2)
	return t.SettlementParameters
}

func (t *TransactionDetails6) AddPlaceOfTrade() *MarketIdentification4 {
	t.PlaceOfTrade = new(MarketIdentification4)
	return t.PlaceOfTrade
}

func (t *TransactionDetails6) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	t.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return t.SafekeepingPlace
}

func (t *TransactionDetails6) SetPlaceOfClearing(value string) {
	t.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (t *TransactionDetails6) AddPostingQuantity() *Quantity6Choice {
	t.PostingQuantity = new(Quantity6Choice)
	return t.PostingQuantity
}

func (t *TransactionDetails6) SetNumberOfDaysAccrued(value string) {
	t.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (t *TransactionDetails6) AddPostingAmount() *AmountAndDirection3 {
	t.PostingAmount = new(AmountAndDirection3)
	return t.PostingAmount
}

func (t *TransactionDetails6) AddAccruedInterestAmount() *AmountAndDirection4 {
	t.AccruedInterestAmount = new(AmountAndDirection4)
	return t.AccruedInterestAmount
}

func (t *TransactionDetails6) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails6) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	t.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return t.EffectiveSettlementDate
}

func (t *TransactionDetails6) AddSettlementDate() *SettlementDate1Choice {
	t.SettlementDate = new(SettlementDate1Choice)
	return t.SettlementDate
}

func (t *TransactionDetails6) AddValueDate() *DateAndDateTimeChoice {
	t.ValueDate = new(DateAndDateTimeChoice)
	return t.ValueDate
}

func (t *TransactionDetails6) AddDeliveringSettlementParties() *SettlementParties2 {
	t.DeliveringSettlementParties = new(SettlementParties2)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails6) AddReceivingSettlementParties() *SettlementParties2 {
	t.ReceivingSettlementParties = new(SettlementParties2)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails6) SetReversalIndicator(value string) {
	t.ReversalIndicator = (*YesNoIndicator)(&value)
}

func (t *TransactionDetails6) SetTransactionAdditionalDetails(value string) {
	t.TransactionAdditionalDetails = (*Max350Text)(&value)
}

// Identifies the details of the transaction.
type TransactionDetails61 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Unique identification assigned to the instruction by the client.
	ClientCollateralInstructionIdentification *Max35Text `xml:"ClntCollInstrId,omitempty"`

	// Unique identification assigned to the instruction by the triparty agent.
	TripartyCollateralInstructionIdentification *Max35Text `xml:"TrptyCollInstrId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification78 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection8 `xml:"SttlmAmt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// For against payment transactions, the value date/time at which the Sender expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate2Choice `xml:"SttlmDt"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails69 `xml:"SttlmParams"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties13 `xml:"RcvgSttlmPties,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties13 `xml:"DlvrgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification37Choice `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentification45Choice `xml:"QlfdFrgnIntrmy,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (t *TransactionDetails61) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *TransactionDetails61) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails61) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails61) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails61) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails61) SetClientCollateralInstructionIdentification(value string) {
	t.ClientCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails61) SetTripartyCollateralInstructionIdentification(value string) {
	t.TripartyCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails61) AddAccountOwner() *PartyIdentification36Choice {
	t.AccountOwner = new(PartyIdentification36Choice)
	return t.AccountOwner
}

func (t *TransactionDetails61) AddSafekeepingAccount() *SecuritiesAccount13 {
	t.SafekeepingAccount = new(SecuritiesAccount13)
	return t.SafekeepingAccount
}

func (t *TransactionDetails61) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	t.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return t.SafekeepingPlace
}

func (t *TransactionDetails61) AddPlaceOfTrade() *MarketIdentification78 {
	t.PlaceOfTrade = new(MarketIdentification78)
	return t.PlaceOfTrade
}

func (t *TransactionDetails61) SetPlaceOfClearing(value string) {
	t.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (t *TransactionDetails61) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails61) AddSettlementQuantity() *Quantity6Choice {
	t.SettlementQuantity = new(Quantity6Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails61) AddSettlementAmount() *AmountAndDirection8 {
	t.SettlementAmount = new(AmountAndDirection8)
	return t.SettlementAmount
}

func (t *TransactionDetails61) AddLateDeliveryDate() *DateAndDateTimeChoice {
	t.LateDeliveryDate = new(DateAndDateTimeChoice)
	return t.LateDeliveryDate
}

func (t *TransactionDetails61) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	t.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return t.ExpectedSettlementDate
}

func (t *TransactionDetails61) AddExpectedValueDate() *DateAndDateTimeChoice {
	t.ExpectedValueDate = new(DateAndDateTimeChoice)
	return t.ExpectedValueDate
}

func (t *TransactionDetails61) AddSettlementDate() *SettlementDate2Choice {
	t.SettlementDate = new(SettlementDate2Choice)
	return t.SettlementDate
}

func (t *TransactionDetails61) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails61) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails61) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails61) AddSettlementParameters() *SettlementDetails69 {
	t.SettlementParameters = new(SettlementDetails69)
	return t.SettlementParameters
}

func (t *TransactionDetails61) AddReceivingSettlementParties() *SettlementParties13 {
	t.ReceivingSettlementParties = new(SettlementParties13)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails61) AddDeliveringSettlementParties() *SettlementParties13 {
	t.DeliveringSettlementParties = new(SettlementParties13)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails61) AddInvestor() *PartyIdentification37Choice {
	t.Investor = new(PartyIdentification37Choice)
	return t.Investor
}

func (t *TransactionDetails61) AddQualifiedForeignIntermediary() *PartyIdentification45Choice {
	t.QualifiedForeignIntermediary = new(PartyIdentification45Choice)
	return t.QualifiedForeignIntermediary
}

func (t *TransactionDetails61) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	t.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Identifies the details of the transaction.
type TransactionDetails62 struct {

	// Specifies the type of activity to which this instruction relates.
	TransactionActivity *TransactionActivity1Choice `xml:"TxActvty"`

	// Choice of type for the transaction reported.
	SettlementTransactionOrCorporateActionEventType *SettlementOrCorporateActionEvent10Choice `xml:"SttlmTxOrCorpActnEvtTp,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails74 `xml:"SttlmParams,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification78 `xml:"PlcOfTrad,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Quantity of financial instrument (to be) posted to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *AmountAndDirection8 `xml:"PstngAmt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate2Choice `xml:"SttlmDt"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// For against payment transactions, the value date/time at which the account servicer expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties13 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties13 `xml:"RcvgSttlmPties,omitempty"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionAdditionalDetails *Max350Text `xml:"TxAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (t *TransactionDetails62) AddTransactionActivity() *TransactionActivity1Choice {
	t.TransactionActivity = new(TransactionActivity1Choice)
	return t.TransactionActivity
}

func (t *TransactionDetails62) AddSettlementTransactionOrCorporateActionEventType() *SettlementOrCorporateActionEvent10Choice {
	t.SettlementTransactionOrCorporateActionEventType = new(SettlementOrCorporateActionEvent10Choice)
	return t.SettlementTransactionOrCorporateActionEventType
}

func (t *TransactionDetails62) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails62) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails62) AddSettlementParameters() *SettlementDetails74 {
	t.SettlementParameters = new(SettlementDetails74)
	return t.SettlementParameters
}

func (t *TransactionDetails62) AddPlaceOfTrade() *MarketIdentification78 {
	t.PlaceOfTrade = new(MarketIdentification78)
	return t.PlaceOfTrade
}

func (t *TransactionDetails62) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	t.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return t.SafekeepingPlace
}

func (t *TransactionDetails62) SetPlaceOfClearing(value string) {
	t.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (t *TransactionDetails62) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails62) AddPostingQuantity() *Quantity6Choice {
	t.PostingQuantity = new(Quantity6Choice)
	return t.PostingQuantity
}

func (t *TransactionDetails62) AddPostingAmount() *AmountAndDirection8 {
	t.PostingAmount = new(AmountAndDirection8)
	return t.PostingAmount
}

func (t *TransactionDetails62) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails62) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	t.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return t.ExpectedSettlementDate
}

func (t *TransactionDetails62) AddSettlementDate() *SettlementDate2Choice {
	t.SettlementDate = new(SettlementDate2Choice)
	return t.SettlementDate
}

func (t *TransactionDetails62) AddLateDeliveryDate() *DateAndDateTimeChoice {
	t.LateDeliveryDate = new(DateAndDateTimeChoice)
	return t.LateDeliveryDate
}

func (t *TransactionDetails62) AddExpectedValueDate() *DateAndDateTimeChoice {
	t.ExpectedValueDate = new(DateAndDateTimeChoice)
	return t.ExpectedValueDate
}

func (t *TransactionDetails62) AddDeliveringSettlementParties() *SettlementParties13 {
	t.DeliveringSettlementParties = new(SettlementParties13)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails62) AddReceivingSettlementParties() *SettlementParties13 {
	t.ReceivingSettlementParties = new(SettlementParties13)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails62) SetTransactionAdditionalDetails(value string) {
	t.TransactionAdditionalDetails = (*Max350Text)(&value)
}

func (t *TransactionDetails62) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	t.SupplementaryData = append(t.SupplementaryData, newValue)
	return newValue
}

// Identifies the details of the transaction.
type TransactionDetails63 struct {

	// Specifies the type of activity to which this instruction relates.
	TransactionActivity *TransactionActivity1Choice `xml:"TxActvty"`

	// Choice of type for the transaction reported.
	SettlementTransactionOrCorporateActionEventType *SettlementOrCorporateActionEvent9Choice `xml:"SttlmTxOrCorpActnEvtTp,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails75 `xml:"SttlmParams,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification78 `xml:"PlcOfTrad,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Quantity of financial instrument (to be) posted to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *AmountAndDirection3 `xml:"PstngAmt,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection4 `xml:"AcrdIntrstAmt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, that is, payment is effected and securities are delivered.
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt,omitempty"`

	// Date and time assets become available to the account owner (in a credit entry), or cease to be available to the account owner (in a debit entry).
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties13 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties13 `xml:"RcvgSttlmPties,omitempty"`

	// Specifies whether it is the reversal of a previously reported movement.
	ReversalIndicator *YesNoIndicator `xml:"RvslInd,omitempty"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionAdditionalDetails *Max350Text `xml:"TxAddtlDtls,omitempty"`
}

func (t *TransactionDetails63) AddTransactionActivity() *TransactionActivity1Choice {
	t.TransactionActivity = new(TransactionActivity1Choice)
	return t.TransactionActivity
}

func (t *TransactionDetails63) AddSettlementTransactionOrCorporateActionEventType() *SettlementOrCorporateActionEvent9Choice {
	t.SettlementTransactionOrCorporateActionEventType = new(SettlementOrCorporateActionEvent9Choice)
	return t.SettlementTransactionOrCorporateActionEventType
}

func (t *TransactionDetails63) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails63) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails63) AddSettlementParameters() *SettlementDetails75 {
	t.SettlementParameters = new(SettlementDetails75)
	return t.SettlementParameters
}

func (t *TransactionDetails63) AddPlaceOfTrade() *MarketIdentification78 {
	t.PlaceOfTrade = new(MarketIdentification78)
	return t.PlaceOfTrade
}

func (t *TransactionDetails63) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	t.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return t.SafekeepingPlace
}

func (t *TransactionDetails63) SetPlaceOfClearing(value string) {
	t.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (t *TransactionDetails63) AddPostingQuantity() *Quantity6Choice {
	t.PostingQuantity = new(Quantity6Choice)
	return t.PostingQuantity
}

func (t *TransactionDetails63) SetNumberOfDaysAccrued(value string) {
	t.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (t *TransactionDetails63) AddPostingAmount() *AmountAndDirection3 {
	t.PostingAmount = new(AmountAndDirection3)
	return t.PostingAmount
}

func (t *TransactionDetails63) AddAccruedInterestAmount() *AmountAndDirection4 {
	t.AccruedInterestAmount = new(AmountAndDirection4)
	return t.AccruedInterestAmount
}

func (t *TransactionDetails63) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails63) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	t.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return t.EffectiveSettlementDate
}

func (t *TransactionDetails63) AddSettlementDate() *SettlementDate1Choice {
	t.SettlementDate = new(SettlementDate1Choice)
	return t.SettlementDate
}

func (t *TransactionDetails63) AddValueDate() *DateAndDateTimeChoice {
	t.ValueDate = new(DateAndDateTimeChoice)
	return t.ValueDate
}

func (t *TransactionDetails63) AddDeliveringSettlementParties() *SettlementParties13 {
	t.DeliveringSettlementParties = new(SettlementParties13)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails63) AddReceivingSettlementParties() *SettlementParties13 {
	t.ReceivingSettlementParties = new(SettlementParties13)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails63) SetReversalIndicator(value string) {
	t.ReversalIndicator = (*YesNoIndicator)(&value)
}

func (t *TransactionDetails63) SetTransactionAdditionalDetails(value string) {
	t.TransactionAdditionalDetails = (*Max350Text)(&value)
}

// Identifies the details of the transaction.
type TransactionDetails64 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *SettlementTypeAndIdentification3 `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Identification of a transaction that cannot be identified using a standard reference element present in the message.
	OtherTransactionIdentification *Max35Text `xml:"OthrTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionDetails *TransactionDetails28 `xml:"TxDtls,omitempty"`
}

func (t *TransactionDetails64) AddAccountOwnerTransactionIdentification() *SettlementTypeAndIdentification3 {
	t.AccountOwnerTransactionIdentification = new(SettlementTypeAndIdentification3)
	return t.AccountOwnerTransactionIdentification
}

func (t *TransactionDetails64) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails64) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails64) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails64) SetOtherTransactionIdentification(value string) {
	t.OtherTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails64) AddAccountOwner() *PartyIdentification36Choice {
	t.AccountOwner = new(PartyIdentification36Choice)
	return t.AccountOwner
}

func (t *TransactionDetails64) AddSafekeepingAccount() *SecuritiesAccount13 {
	t.SafekeepingAccount = new(SecuritiesAccount13)
	return t.SafekeepingAccount
}

func (t *TransactionDetails64) AddTransactionDetails() *TransactionDetails28 {
	t.TransactionDetails = new(TransactionDetails28)
	return t.TransactionDetails
}

// Identifies the details of the transaction.
type TransactionDetails69 struct {

	// Specifies the type of activity to which this instruction relates.
	TransactionActivity *TransactionActivity1Choice `xml:"TxActvty"`

	// Choice of type for the transaction reported.
	SettlementTransactionOrCorporateActionEventType *SettlementOrCorporateActionEvent10Choice `xml:"SttlmTxOrCorpActnEvtTp,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails74 `xml:"SttlmParams,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification78 `xml:"PlcOfTrad,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Quantity of financial instrument (to be) posted to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *AmountAndDirection8 `xml:"PstngAmt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date/time at which the sender expects settlement.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate2Choice `xml:"SttlmDt"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// For against payment transactions, the value date/time at which the account servicer expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties13 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties13 `xml:"RcvgSttlmPties,omitempty"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionAdditionalDetails *Max350Text `xml:"TxAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (t *TransactionDetails69) AddTransactionActivity() *TransactionActivity1Choice {
	t.TransactionActivity = new(TransactionActivity1Choice)
	return t.TransactionActivity
}

func (t *TransactionDetails69) AddSettlementTransactionOrCorporateActionEventType() *SettlementOrCorporateActionEvent10Choice {
	t.SettlementTransactionOrCorporateActionEventType = new(SettlementOrCorporateActionEvent10Choice)
	return t.SettlementTransactionOrCorporateActionEventType
}

func (t *TransactionDetails69) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails69) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails69) AddSettlementParameters() *SettlementDetails74 {
	t.SettlementParameters = new(SettlementDetails74)
	return t.SettlementParameters
}

func (t *TransactionDetails69) AddPlaceOfTrade() *MarketIdentification78 {
	t.PlaceOfTrade = new(MarketIdentification78)
	return t.PlaceOfTrade
}

func (t *TransactionDetails69) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	t.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return t.SafekeepingPlace
}

func (t *TransactionDetails69) SetPlaceOfClearing(value string) {
	t.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (t *TransactionDetails69) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails69) AddPostingQuantity() *Quantity6Choice {
	t.PostingQuantity = new(Quantity6Choice)
	return t.PostingQuantity
}

func (t *TransactionDetails69) AddPostingAmount() *AmountAndDirection8 {
	t.PostingAmount = new(AmountAndDirection8)
	return t.PostingAmount
}

func (t *TransactionDetails69) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails69) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	t.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return t.ExpectedSettlementDate
}

func (t *TransactionDetails69) AddSettlementDate() *SettlementDate2Choice {
	t.SettlementDate = new(SettlementDate2Choice)
	return t.SettlementDate
}

func (t *TransactionDetails69) AddLateDeliveryDate() *DateAndDateTimeChoice {
	t.LateDeliveryDate = new(DateAndDateTimeChoice)
	return t.LateDeliveryDate
}

func (t *TransactionDetails69) AddExpectedValueDate() *DateAndDateTimeChoice {
	t.ExpectedValueDate = new(DateAndDateTimeChoice)
	return t.ExpectedValueDate
}

func (t *TransactionDetails69) AddDeliveringSettlementParties() *SettlementParties13 {
	t.DeliveringSettlementParties = new(SettlementParties13)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails69) AddReceivingSettlementParties() *SettlementParties13 {
	t.ReceivingSettlementParties = new(SettlementParties13)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails69) SetTransactionAdditionalDetails(value string) {
	t.TransactionAdditionalDetails = (*Max350Text)(&value)
}

func (t *TransactionDetails69) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	t.SupplementaryData = append(t.SupplementaryData, newValue)
	return newValue
}

// Identifies the details of the transaction.
type TransactionDetails7 struct {

	// Specifies the type of activity to which this instruction relates.
	TransactionActivity *TransactionActivity1Choice `xml:"TxActvty"`

	// Choice of type for the transaction reported.
	SettlementTransactionOrCorporateActionEventType *SettlementOrCorporateActionEvent2Choice `xml:"SttlmTxOrCorpActnEvtTp,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails7 `xml:"SttlmParams,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification11 `xml:"FinInstrmId"`

	// Quantity of financial instrument (to be) posted to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *AmountAndDirection8 `xml:"PstngAmt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate2Choice `xml:"SttlmDt"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// For against payment transactions, the value date/time at which the account servicer expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties2 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties2 `xml:"RcvgSttlmPties,omitempty"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionAdditionalDetails *Max350Text `xml:"TxAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension2 `xml:"Xtnsn,omitempty"`
}

func (t *TransactionDetails7) AddTransactionActivity() *TransactionActivity1Choice {
	t.TransactionActivity = new(TransactionActivity1Choice)
	return t.TransactionActivity
}

func (t *TransactionDetails7) AddSettlementTransactionOrCorporateActionEventType() *SettlementOrCorporateActionEvent2Choice {
	t.SettlementTransactionOrCorporateActionEventType = new(SettlementOrCorporateActionEvent2Choice)
	return t.SettlementTransactionOrCorporateActionEventType
}

func (t *TransactionDetails7) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails7) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails7) AddSettlementParameters() *SettlementDetails7 {
	t.SettlementParameters = new(SettlementDetails7)
	return t.SettlementParameters
}

func (t *TransactionDetails7) AddPlaceOfTrade() *MarketIdentification4 {
	t.PlaceOfTrade = new(MarketIdentification4)
	return t.PlaceOfTrade
}

func (t *TransactionDetails7) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	t.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return t.SafekeepingPlace
}

func (t *TransactionDetails7) SetPlaceOfClearing(value string) {
	t.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (t *TransactionDetails7) AddFinancialInstrumentIdentification() *SecurityIdentification11 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification11)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails7) AddPostingQuantity() *Quantity6Choice {
	t.PostingQuantity = new(Quantity6Choice)
	return t.PostingQuantity
}

func (t *TransactionDetails7) AddPostingAmount() *AmountAndDirection8 {
	t.PostingAmount = new(AmountAndDirection8)
	return t.PostingAmount
}

func (t *TransactionDetails7) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails7) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	t.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return t.ExpectedSettlementDate
}

func (t *TransactionDetails7) AddSettlementDate() *SettlementDate2Choice {
	t.SettlementDate = new(SettlementDate2Choice)
	return t.SettlementDate
}

func (t *TransactionDetails7) AddLateDeliveryDate() *DateAndDateTimeChoice {
	t.LateDeliveryDate = new(DateAndDateTimeChoice)
	return t.LateDeliveryDate
}

func (t *TransactionDetails7) AddExpectedValueDate() *DateAndDateTimeChoice {
	t.ExpectedValueDate = new(DateAndDateTimeChoice)
	return t.ExpectedValueDate
}

func (t *TransactionDetails7) AddDeliveringSettlementParties() *SettlementParties2 {
	t.DeliveringSettlementParties = new(SettlementParties2)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails7) AddReceivingSettlementParties() *SettlementParties2 {
	t.ReceivingSettlementParties = new(SettlementParties2)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails7) SetTransactionAdditionalDetails(value string) {
	t.TransactionAdditionalDetails = (*Max350Text)(&value)
}

func (t *TransactionDetails7) AddExtension() *Extension2 {
	newValue := new(Extension2)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

// Identifies the details of the transaction.
type TransactionDetails70 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Unique identification assigned to the instruction by the client.
	ClientCollateralInstructionIdentification *Max35Text `xml:"ClntCollInstrId,omitempty"`

	// Unique identification assigned to the instruction by the triparty agent.
	TripartyCollateralInstructionIdentification *Max35Text `xml:"TrptyCollInstrId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification78 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection8 `xml:"SttlmAmt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Date/time at which the sender expects settlement.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// For against payment transactions, the value date/time at which the Sender expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate2Choice `xml:"SttlmDt"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails69 `xml:"SttlmParams"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties13 `xml:"RcvgSttlmPties,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties13 `xml:"DlvrgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification37Choice `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentification45Choice `xml:"QlfdFrgnIntrmy,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (t *TransactionDetails70) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *TransactionDetails70) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails70) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails70) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails70) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails70) SetClientCollateralInstructionIdentification(value string) {
	t.ClientCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails70) SetTripartyCollateralInstructionIdentification(value string) {
	t.TripartyCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails70) AddAccountOwner() *PartyIdentification36Choice {
	t.AccountOwner = new(PartyIdentification36Choice)
	return t.AccountOwner
}

func (t *TransactionDetails70) AddSafekeepingAccount() *SecuritiesAccount13 {
	t.SafekeepingAccount = new(SecuritiesAccount13)
	return t.SafekeepingAccount
}

func (t *TransactionDetails70) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	t.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return t.SafekeepingPlace
}

func (t *TransactionDetails70) AddPlaceOfTrade() *MarketIdentification78 {
	t.PlaceOfTrade = new(MarketIdentification78)
	return t.PlaceOfTrade
}

func (t *TransactionDetails70) SetPlaceOfClearing(value string) {
	t.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (t *TransactionDetails70) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails70) AddSettlementQuantity() *Quantity6Choice {
	t.SettlementQuantity = new(Quantity6Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails70) AddSettlementAmount() *AmountAndDirection8 {
	t.SettlementAmount = new(AmountAndDirection8)
	return t.SettlementAmount
}

func (t *TransactionDetails70) AddLateDeliveryDate() *DateAndDateTimeChoice {
	t.LateDeliveryDate = new(DateAndDateTimeChoice)
	return t.LateDeliveryDate
}

func (t *TransactionDetails70) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	t.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return t.ExpectedSettlementDate
}

func (t *TransactionDetails70) AddExpectedValueDate() *DateAndDateTimeChoice {
	t.ExpectedValueDate = new(DateAndDateTimeChoice)
	return t.ExpectedValueDate
}

func (t *TransactionDetails70) AddSettlementDate() *SettlementDate2Choice {
	t.SettlementDate = new(SettlementDate2Choice)
	return t.SettlementDate
}

func (t *TransactionDetails70) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails70) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails70) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails70) AddSettlementParameters() *SettlementDetails69 {
	t.SettlementParameters = new(SettlementDetails69)
	return t.SettlementParameters
}

func (t *TransactionDetails70) AddReceivingSettlementParties() *SettlementParties13 {
	t.ReceivingSettlementParties = new(SettlementParties13)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails70) AddDeliveringSettlementParties() *SettlementParties13 {
	t.DeliveringSettlementParties = new(SettlementParties13)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails70) AddInvestor() *PartyIdentification37Choice {
	t.Investor = new(PartyIdentification37Choice)
	return t.Investor
}

func (t *TransactionDetails70) AddQualifiedForeignIntermediary() *PartyIdentification45Choice {
	t.QualifiedForeignIntermediary = new(PartyIdentification45Choice)
	return t.QualifiedForeignIntermediary
}

func (t *TransactionDetails70) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	t.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Identifies the details of the transaction.
type TransactionDetails74 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate9Choice `xml:"SttlmDt"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection51 `xml:"SttlmAmt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties40 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties40 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification99 `xml:"Invstr,omitempty"`
}

func (t *TransactionDetails74) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails74) AddTradeDate() *TradeDate5Choice {
	t.TradeDate = new(TradeDate5Choice)
	return t.TradeDate
}

func (t *TransactionDetails74) AddSettlementDate() *SettlementDate9Choice {
	t.SettlementDate = new(SettlementDate9Choice)
	return t.SettlementDate
}

func (t *TransactionDetails74) AddSettlementQuantity() *Quantity6Choice {
	t.SettlementQuantity = new(Quantity6Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails74) AddSettlementAmount() *AmountAndDirection51 {
	t.SettlementAmount = new(AmountAndDirection51)
	return t.SettlementAmount
}

func (t *TransactionDetails74) AddDeliveringSettlementParties() *SettlementParties40 {
	t.DeliveringSettlementParties = new(SettlementParties40)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails74) AddReceivingSettlementParties() *SettlementParties40 {
	t.ReceivingSettlementParties = new(SettlementParties40)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails74) AddInvestor() *PartyIdentification99 {
	t.Investor = new(PartyIdentification99)
	return t.Investor
}

// Identifies the details of the transaction.
type TransactionDetails75 struct {

	// Specifies the type of activity to which this instruction relates.
	TransactionActivity *TransactionActivity3Choice `xml:"TxActvty"`

	// Choice of type for the transaction reported.
	SettlementTransactionOrCorporateActionEventType *SettlementOrCorporateActionEvent14Choice `xml:"SttlmTxOrCorpActnEvtTp,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails92 `xml:"SttlmParams,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification1 `xml:"PlcOfTrad,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Quantity of financial instrument (to be) posted to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *AmountAndDirection51 `xml:"PstngAmt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Date/time at which the sender expects settlement.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate10Choice `xml:"SttlmDt"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// For against payment transactions, the value date/time at which the account servicer expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties40 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties40 `xml:"RcvgSttlmPties,omitempty"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionAdditionalDetails *Max350Text `xml:"TxAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (t *TransactionDetails75) AddTransactionActivity() *TransactionActivity3Choice {
	t.TransactionActivity = new(TransactionActivity3Choice)
	return t.TransactionActivity
}

func (t *TransactionDetails75) AddSettlementTransactionOrCorporateActionEventType() *SettlementOrCorporateActionEvent14Choice {
	t.SettlementTransactionOrCorporateActionEventType = new(SettlementOrCorporateActionEvent14Choice)
	return t.SettlementTransactionOrCorporateActionEventType
}

func (t *TransactionDetails75) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails75) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails75) AddSettlementParameters() *SettlementDetails92 {
	t.SettlementParameters = new(SettlementDetails92)
	return t.SettlementParameters
}

func (t *TransactionDetails75) AddPlaceOfTrade() *PlaceOfTradeIdentification1 {
	t.PlaceOfTrade = new(PlaceOfTradeIdentification1)
	return t.PlaceOfTrade
}

func (t *TransactionDetails75) AddSafekeepingPlace() *SafeKeepingPlace1 {
	t.SafekeepingPlace = new(SafeKeepingPlace1)
	return t.SafekeepingPlace
}

func (t *TransactionDetails75) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	t.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return t.PlaceOfClearing
}

func (t *TransactionDetails75) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails75) AddPostingQuantity() *Quantity6Choice {
	t.PostingQuantity = new(Quantity6Choice)
	return t.PostingQuantity
}

func (t *TransactionDetails75) AddPostingAmount() *AmountAndDirection51 {
	t.PostingAmount = new(AmountAndDirection51)
	return t.PostingAmount
}

func (t *TransactionDetails75) AddTradeDate() *TradeDate5Choice {
	t.TradeDate = new(TradeDate5Choice)
	return t.TradeDate
}

func (t *TransactionDetails75) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	t.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return t.ExpectedSettlementDate
}

func (t *TransactionDetails75) AddSettlementDate() *SettlementDate10Choice {
	t.SettlementDate = new(SettlementDate10Choice)
	return t.SettlementDate
}

func (t *TransactionDetails75) AddLateDeliveryDate() *DateAndDateTimeChoice {
	t.LateDeliveryDate = new(DateAndDateTimeChoice)
	return t.LateDeliveryDate
}

func (t *TransactionDetails75) AddExpectedValueDate() *DateAndDateTimeChoice {
	t.ExpectedValueDate = new(DateAndDateTimeChoice)
	return t.ExpectedValueDate
}

func (t *TransactionDetails75) AddDeliveringSettlementParties() *SettlementParties40 {
	t.DeliveringSettlementParties = new(SettlementParties40)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails75) AddReceivingSettlementParties() *SettlementParties40 {
	t.ReceivingSettlementParties = new(SettlementParties40)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails75) SetTransactionAdditionalDetails(value string) {
	t.TransactionAdditionalDetails = (*Max350Text)(&value)
}

func (t *TransactionDetails75) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	t.SupplementaryData = append(t.SupplementaryData, newValue)
	return newValue
}

// Identifies the details of the transaction.
type TransactionDetails76 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *SettlementTypeAndIdentification18 `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Identification of a transaction that cannot be identified using a standard reference element present in the message.
	OtherTransactionIdentification *Max35Text `xml:"OthrTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionDetails *TransactionDetails74 `xml:"TxDtls,omitempty"`
}

func (t *TransactionDetails76) AddAccountOwnerTransactionIdentification() *SettlementTypeAndIdentification18 {
	t.AccountOwnerTransactionIdentification = new(SettlementTypeAndIdentification18)
	return t.AccountOwnerTransactionIdentification
}

func (t *TransactionDetails76) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails76) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails76) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails76) SetOtherTransactionIdentification(value string) {
	t.OtherTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails76) AddAccountOwner() *PartyIdentification98 {
	t.AccountOwner = new(PartyIdentification98)
	return t.AccountOwner
}

func (t *TransactionDetails76) AddSafekeepingAccount() *SecuritiesAccount19 {
	t.SafekeepingAccount = new(SecuritiesAccount19)
	return t.SafekeepingAccount
}

func (t *TransactionDetails76) AddTransactionDetails() *TransactionDetails74 {
	t.TransactionDetails = new(TransactionDetails74)
	return t.TransactionDetails
}

// Identifies the details of the transaction.
type TransactionDetails78 struct {

	// Specifies the type of activity to which this instruction relates.
	TransactionActivity *TransactionActivity3Choice `xml:"TxActvty"`

	// Choice of type for the transaction reported.
	SettlementTransactionOrCorporateActionEventType *SettlementOrCorporateActionEvent13Choice `xml:"SttlmTxOrCorpActnEvtTp,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails91 `xml:"SttlmParams,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification1 `xml:"PlcOfTrad,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Quantity of financial instrument (to be) posted to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *AmountAndDirection3 `xml:"PstngAmt,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection21 `xml:"AcrdIntrstAmt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, that is, payment is effected and securities are delivered.
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate9Choice `xml:"SttlmDt,omitempty"`

	// Date and time assets become available to the account owner (in a credit entry), or cease to be available to the account owner (in a debit entry).
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties40 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties40 `xml:"RcvgSttlmPties,omitempty"`

	// Specifies whether it is the reversal of a previously reported movement.
	ReversalIndicator *YesNoIndicator `xml:"RvslInd,omitempty"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionAdditionalDetails *Max350Text `xml:"TxAddtlDtls,omitempty"`
}

func (t *TransactionDetails78) AddTransactionActivity() *TransactionActivity3Choice {
	t.TransactionActivity = new(TransactionActivity3Choice)
	return t.TransactionActivity
}

func (t *TransactionDetails78) AddSettlementTransactionOrCorporateActionEventType() *SettlementOrCorporateActionEvent13Choice {
	t.SettlementTransactionOrCorporateActionEventType = new(SettlementOrCorporateActionEvent13Choice)
	return t.SettlementTransactionOrCorporateActionEventType
}

func (t *TransactionDetails78) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails78) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails78) AddSettlementParameters() *SettlementDetails91 {
	t.SettlementParameters = new(SettlementDetails91)
	return t.SettlementParameters
}

func (t *TransactionDetails78) AddPlaceOfTrade() *PlaceOfTradeIdentification1 {
	t.PlaceOfTrade = new(PlaceOfTradeIdentification1)
	return t.PlaceOfTrade
}

func (t *TransactionDetails78) AddSafekeepingPlace() *SafeKeepingPlace1 {
	t.SafekeepingPlace = new(SafeKeepingPlace1)
	return t.SafekeepingPlace
}

func (t *TransactionDetails78) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	t.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return t.PlaceOfClearing
}

func (t *TransactionDetails78) AddPostingQuantity() *Quantity6Choice {
	t.PostingQuantity = new(Quantity6Choice)
	return t.PostingQuantity
}

func (t *TransactionDetails78) SetNumberOfDaysAccrued(value string) {
	t.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (t *TransactionDetails78) AddPostingAmount() *AmountAndDirection3 {
	t.PostingAmount = new(AmountAndDirection3)
	return t.PostingAmount
}

func (t *TransactionDetails78) AddAccruedInterestAmount() *AmountAndDirection21 {
	t.AccruedInterestAmount = new(AmountAndDirection21)
	return t.AccruedInterestAmount
}

func (t *TransactionDetails78) AddTradeDate() *TradeDate5Choice {
	t.TradeDate = new(TradeDate5Choice)
	return t.TradeDate
}

func (t *TransactionDetails78) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	t.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return t.EffectiveSettlementDate
}

func (t *TransactionDetails78) AddSettlementDate() *SettlementDate9Choice {
	t.SettlementDate = new(SettlementDate9Choice)
	return t.SettlementDate
}

func (t *TransactionDetails78) AddValueDate() *DateAndDateTimeChoice {
	t.ValueDate = new(DateAndDateTimeChoice)
	return t.ValueDate
}

func (t *TransactionDetails78) AddDeliveringSettlementParties() *SettlementParties40 {
	t.DeliveringSettlementParties = new(SettlementParties40)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails78) AddReceivingSettlementParties() *SettlementParties40 {
	t.ReceivingSettlementParties = new(SettlementParties40)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails78) SetReversalIndicator(value string) {
	t.ReversalIndicator = (*YesNoIndicator)(&value)
}

func (t *TransactionDetails78) SetTransactionAdditionalDetails(value string) {
	t.TransactionAdditionalDetails = (*Max350Text)(&value)
}

// Identifies the details of the transaction.
type TransactionDetails79 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Unique identification assigned to the instruction by the client.
	ClientCollateralInstructionIdentification *Max35Text `xml:"ClntCollInstrId,omitempty"`

	// Unique identification assigned to the instruction by the triparty agent.
	TripartyCollateralInstructionIdentification *Max35Text `xml:"TrptyCollInstrId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount24 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification1 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection51 `xml:"SttlmAmt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Date/time at which the sender expects settlement.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// For against payment transactions, the value date/time at which the Sender expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate10Choice `xml:"SttlmDt"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails95 `xml:"SttlmParams"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties40 `xml:"RcvgSttlmPties,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties40 `xml:"DlvrgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification99 `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentification100 `xml:"QlfdFrgnIntrmy,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (t *TransactionDetails79) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *TransactionDetails79) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails79) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails79) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails79) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails79) SetClientCollateralInstructionIdentification(value string) {
	t.ClientCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails79) SetTripartyCollateralInstructionIdentification(value string) {
	t.TripartyCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails79) AddAccountOwner() *PartyIdentification98 {
	t.AccountOwner = new(PartyIdentification98)
	return t.AccountOwner
}

func (t *TransactionDetails79) AddSafekeepingAccount() *SecuritiesAccount24 {
	t.SafekeepingAccount = new(SecuritiesAccount24)
	return t.SafekeepingAccount
}

func (t *TransactionDetails79) AddSafekeepingPlace() *SafeKeepingPlace1 {
	t.SafekeepingPlace = new(SafeKeepingPlace1)
	return t.SafekeepingPlace
}

func (t *TransactionDetails79) AddPlaceOfTrade() *PlaceOfTradeIdentification1 {
	t.PlaceOfTrade = new(PlaceOfTradeIdentification1)
	return t.PlaceOfTrade
}

func (t *TransactionDetails79) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	t.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return t.PlaceOfClearing
}

func (t *TransactionDetails79) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails79) AddSettlementQuantity() *Quantity6Choice {
	t.SettlementQuantity = new(Quantity6Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails79) AddSettlementAmount() *AmountAndDirection51 {
	t.SettlementAmount = new(AmountAndDirection51)
	return t.SettlementAmount
}

func (t *TransactionDetails79) AddLateDeliveryDate() *DateAndDateTimeChoice {
	t.LateDeliveryDate = new(DateAndDateTimeChoice)
	return t.LateDeliveryDate
}

func (t *TransactionDetails79) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	t.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return t.ExpectedSettlementDate
}

func (t *TransactionDetails79) AddExpectedValueDate() *DateAndDateTimeChoice {
	t.ExpectedValueDate = new(DateAndDateTimeChoice)
	return t.ExpectedValueDate
}

func (t *TransactionDetails79) AddSettlementDate() *SettlementDate10Choice {
	t.SettlementDate = new(SettlementDate10Choice)
	return t.SettlementDate
}

func (t *TransactionDetails79) AddTradeDate() *TradeDate5Choice {
	t.TradeDate = new(TradeDate5Choice)
	return t.TradeDate
}

func (t *TransactionDetails79) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails79) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails79) AddSettlementParameters() *SettlementDetails95 {
	t.SettlementParameters = new(SettlementDetails95)
	return t.SettlementParameters
}

func (t *TransactionDetails79) AddReceivingSettlementParties() *SettlementParties40 {
	t.ReceivingSettlementParties = new(SettlementParties40)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails79) AddDeliveringSettlementParties() *SettlementParties40 {
	t.DeliveringSettlementParties = new(SettlementParties40)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails79) AddInvestor() *PartyIdentification99 {
	t.Investor = new(PartyIdentification99)
	return t.Investor
}

func (t *TransactionDetails79) AddQualifiedForeignIntermediary() *PartyIdentification100 {
	t.QualifiedForeignIntermediary = new(PartyIdentification100)
	return t.QualifiedForeignIntermediary
}

func (t *TransactionDetails79) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	t.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Identifies the details of the transaction.
type TransactionDetails8 struct {

	// Specifies the type of activity to which this instruction relates.
	TransactionActivity *TransactionActivity1Choice `xml:"TxActvty"`

	// Choice of type for the transaction reported.
	SettlementTransactionOrCorporateActionEventType *SettlementOrCorporateActionEvent2Choice `xml:"SttlmTxOrCorpActnEvtTp,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails7 `xml:"SttlmParams,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification11 `xml:"FinInstrmId"`

	// Quantity of financial instrument (to be) posted to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *AmountAndDirection8 `xml:"PstngAmt"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate2Choice `xml:"SttlmDt"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// For against payment transactions, the value date/time at which the account servicer expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties2 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties2 `xml:"RcvgSttlmPties,omitempty"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionAdditionalDetails *Max350Text `xml:"TxAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension2 `xml:"Xtnsn,omitempty"`
}

func (t *TransactionDetails8) AddTransactionActivity() *TransactionActivity1Choice {
	t.TransactionActivity = new(TransactionActivity1Choice)
	return t.TransactionActivity
}

func (t *TransactionDetails8) AddSettlementTransactionOrCorporateActionEventType() *SettlementOrCorporateActionEvent2Choice {
	t.SettlementTransactionOrCorporateActionEventType = new(SettlementOrCorporateActionEvent2Choice)
	return t.SettlementTransactionOrCorporateActionEventType
}

func (t *TransactionDetails8) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails8) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails8) AddSettlementParameters() *SettlementDetails7 {
	t.SettlementParameters = new(SettlementDetails7)
	return t.SettlementParameters
}

func (t *TransactionDetails8) AddPlaceOfTrade() *MarketIdentification4 {
	t.PlaceOfTrade = new(MarketIdentification4)
	return t.PlaceOfTrade
}

func (t *TransactionDetails8) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	t.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return t.SafekeepingPlace
}

func (t *TransactionDetails8) AddFinancialInstrumentIdentification() *SecurityIdentification11 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification11)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails8) AddPostingQuantity() *Quantity6Choice {
	t.PostingQuantity = new(Quantity6Choice)
	return t.PostingQuantity
}

func (t *TransactionDetails8) AddPostingAmount() *AmountAndDirection8 {
	t.PostingAmount = new(AmountAndDirection8)
	return t.PostingAmount
}

func (t *TransactionDetails8) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails8) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	t.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return t.ExpectedSettlementDate
}

func (t *TransactionDetails8) AddSettlementDate() *SettlementDate2Choice {
	t.SettlementDate = new(SettlementDate2Choice)
	return t.SettlementDate
}

func (t *TransactionDetails8) AddLateDeliveryDate() *DateAndDateTimeChoice {
	t.LateDeliveryDate = new(DateAndDateTimeChoice)
	return t.LateDeliveryDate
}

func (t *TransactionDetails8) AddExpectedValueDate() *DateAndDateTimeChoice {
	t.ExpectedValueDate = new(DateAndDateTimeChoice)
	return t.ExpectedValueDate
}

func (t *TransactionDetails8) AddDeliveringSettlementParties() *SettlementParties2 {
	t.DeliveringSettlementParties = new(SettlementParties2)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails8) AddReceivingSettlementParties() *SettlementParties2 {
	t.ReceivingSettlementParties = new(SettlementParties2)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails8) SetTransactionAdditionalDetails(value string) {
	t.TransactionAdditionalDetails = (*Max350Text)(&value)
}

func (t *TransactionDetails8) AddExtension() *Extension2 {
	newValue := new(Extension2)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

// Identifies the details of the transaction.
type TransactionDetails80 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount24 `xml:"SfkpgAcct"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection51 `xml:"SttlmAmt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate10Choice `xml:"SttlmDt"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties40 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties40 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification99 `xml:"Invstr,omitempty"`
}

func (t *TransactionDetails80) AddAccountOwner() *PartyIdentification98 {
	t.AccountOwner = new(PartyIdentification98)
	return t.AccountOwner
}

func (t *TransactionDetails80) AddSafekeepingAccount() *SecuritiesAccount24 {
	t.SafekeepingAccount = new(SecuritiesAccount24)
	return t.SafekeepingAccount
}

func (t *TransactionDetails80) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails80) AddSettlementQuantity() *Quantity6Choice {
	t.SettlementQuantity = new(Quantity6Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails80) AddSettlementAmount() *AmountAndDirection51 {
	t.SettlementAmount = new(AmountAndDirection51)
	return t.SettlementAmount
}

func (t *TransactionDetails80) AddTradeDate() *TradeDate5Choice {
	t.TradeDate = new(TradeDate5Choice)
	return t.TradeDate
}

func (t *TransactionDetails80) AddSettlementDate() *SettlementDate10Choice {
	t.SettlementDate = new(SettlementDate10Choice)
	return t.SettlementDate
}

func (t *TransactionDetails80) AddDeliveringSettlementParties() *SettlementParties40 {
	t.DeliveringSettlementParties = new(SettlementParties40)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails80) AddReceivingSettlementParties() *SettlementParties40 {
	t.ReceivingSettlementParties = new(SettlementParties40)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails80) AddInvestor() *PartyIdentification99 {
	t.Investor = new(PartyIdentification99)
	return t.Investor
}

// Identifies the details of the transaction.
type TransactionDetails81 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, eg, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection51 `xml:"SttlmAmt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate10Choice `xml:"SttlmDt"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties40 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties40 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification99 `xml:"Invstr,omitempty"`
}

func (t *TransactionDetails81) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails81) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails81) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails81) AddSettlementQuantity() *Quantity6Choice {
	t.SettlementQuantity = new(Quantity6Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails81) AddSettlementAmount() *AmountAndDirection51 {
	t.SettlementAmount = new(AmountAndDirection51)
	return t.SettlementAmount
}

func (t *TransactionDetails81) AddSettlementDate() *SettlementDate10Choice {
	t.SettlementDate = new(SettlementDate10Choice)
	return t.SettlementDate
}

func (t *TransactionDetails81) AddTradeDate() *TradeDate5Choice {
	t.TradeDate = new(TradeDate5Choice)
	return t.TradeDate
}

func (t *TransactionDetails81) AddDeliveringSettlementParties() *SettlementParties40 {
	t.DeliveringSettlementParties = new(SettlementParties40)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails81) AddReceivingSettlementParties() *SettlementParties40 {
	t.ReceivingSettlementParties = new(SettlementParties40)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails81) AddInvestor() *PartyIdentification99 {
	t.Investor = new(PartyIdentification99)
	return t.Investor
}

// Identifies the details of the transaction.
type TransactionDetails82 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, eg, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection51 `xml:"SttlmAmt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate10Choice `xml:"SttlmDt"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties40 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties40 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification99 `xml:"Invstr,omitempty"`
}

func (t *TransactionDetails82) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails82) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails82) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails82) AddSettlementQuantity() *Quantity6Choice {
	t.SettlementQuantity = new(Quantity6Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails82) AddSafekeepingAccount() *SecuritiesAccount19 {
	t.SafekeepingAccount = new(SecuritiesAccount19)
	return t.SafekeepingAccount
}

func (t *TransactionDetails82) AddSettlementAmount() *AmountAndDirection51 {
	t.SettlementAmount = new(AmountAndDirection51)
	return t.SettlementAmount
}

func (t *TransactionDetails82) AddSettlementDate() *SettlementDate10Choice {
	t.SettlementDate = new(SettlementDate10Choice)
	return t.SettlementDate
}

func (t *TransactionDetails82) AddTradeDate() *TradeDate5Choice {
	t.TradeDate = new(TradeDate5Choice)
	return t.TradeDate
}

func (t *TransactionDetails82) AddDeliveringSettlementParties() *SettlementParties40 {
	t.DeliveringSettlementParties = new(SettlementParties40)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails82) AddReceivingSettlementParties() *SettlementParties40 {
	t.ReceivingSettlementParties = new(SettlementParties40)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails82) AddInvestor() *PartyIdentification99 {
	t.Investor = new(PartyIdentification99)
	return t.Investor
}

// Identifies the details of the transaction.
type TransactionDetails83 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate12Choice `xml:"SttlmDt"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity10Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection67 `xml:"SttlmAmt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties49 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties49 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification110 `xml:"Invstr,omitempty"`
}

func (t *TransactionDetails83) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails83) AddTradeDate() *TradeDate6Choice {
	t.TradeDate = new(TradeDate6Choice)
	return t.TradeDate
}

func (t *TransactionDetails83) AddSettlementDate() *SettlementDate12Choice {
	t.SettlementDate = new(SettlementDate12Choice)
	return t.SettlementDate
}

func (t *TransactionDetails83) AddSettlementQuantity() *Quantity10Choice {
	t.SettlementQuantity = new(Quantity10Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails83) AddSettlementAmount() *AmountAndDirection67 {
	t.SettlementAmount = new(AmountAndDirection67)
	return t.SettlementAmount
}

func (t *TransactionDetails83) AddDeliveringSettlementParties() *SettlementParties49 {
	t.DeliveringSettlementParties = new(SettlementParties49)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails83) AddReceivingSettlementParties() *SettlementParties49 {
	t.ReceivingSettlementParties = new(SettlementParties49)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails83) AddInvestor() *PartyIdentification110 {
	t.Investor = new(PartyIdentification110)
	return t.Investor
}

// Identifies the details of the transaction.
type TransactionDetails84 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, eg, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity10Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection67 `xml:"SttlmAmt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate15Choice `xml:"SttlmDt"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties49 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties49 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification110 `xml:"Invstr,omitempty"`
}

func (t *TransactionDetails84) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails84) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails84) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails84) AddSettlementQuantity() *Quantity10Choice {
	t.SettlementQuantity = new(Quantity10Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails84) AddSettlementAmount() *AmountAndDirection67 {
	t.SettlementAmount = new(AmountAndDirection67)
	return t.SettlementAmount
}

func (t *TransactionDetails84) AddSettlementDate() *SettlementDate15Choice {
	t.SettlementDate = new(SettlementDate15Choice)
	return t.SettlementDate
}

func (t *TransactionDetails84) AddTradeDate() *TradeDate6Choice {
	t.TradeDate = new(TradeDate6Choice)
	return t.TradeDate
}

func (t *TransactionDetails84) AddDeliveringSettlementParties() *SettlementParties49 {
	t.DeliveringSettlementParties = new(SettlementParties49)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails84) AddReceivingSettlementParties() *SettlementParties49 {
	t.ReceivingSettlementParties = new(SettlementParties49)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails84) AddInvestor() *PartyIdentification110 {
	t.Investor = new(PartyIdentification110)
	return t.Investor
}

// Identifies the details of the transaction.
type TransactionDetails85 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *SettlementTypeAndIdentification22 `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *RestrictedFINXMax16Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *RestrictedFINXMax16Text `xml:"PrcrTxId,omitempty"`

	// Identification of a transaction that cannot be identified using a standard reference element present in the message.
	OtherTransactionIdentification *RestrictedFINXMax16Text `xml:"OthrTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount30 `xml:"SfkpgAcct"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionDetails *TransactionDetails83 `xml:"TxDtls,omitempty"`
}

func (t *TransactionDetails85) AddAccountOwnerTransactionIdentification() *SettlementTypeAndIdentification22 {
	t.AccountOwnerTransactionIdentification = new(SettlementTypeAndIdentification22)
	return t.AccountOwnerTransactionIdentification
}

func (t *TransactionDetails85) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionDetails85) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionDetails85) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionDetails85) SetOtherTransactionIdentification(value string) {
	t.OtherTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionDetails85) AddAccountOwner() *PartyIdentification109 {
	t.AccountOwner = new(PartyIdentification109)
	return t.AccountOwner
}

func (t *TransactionDetails85) AddSafekeepingAccount() *SecuritiesAccount30 {
	t.SafekeepingAccount = new(SecuritiesAccount30)
	return t.SafekeepingAccount
}

func (t *TransactionDetails85) AddTransactionDetails() *TransactionDetails83 {
	t.TransactionDetails = new(TransactionDetails83)
	return t.TransactionDetails
}

// Identifies the details of the transaction.
type TransactionDetails87 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*RestrictedFINXMax16Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Unique identification assigned to the instruction by the client.
	ClientCollateralInstructionIdentification *RestrictedFINXMax16Text `xml:"ClntCollInstrId,omitempty"`

	// Unique identification assigned to the instruction by the triparty agent.
	TripartyCollateralInstructionIdentification *RestrictedFINXMax16Text `xml:"TrptyCollInstrId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount27 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification2 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity10Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection67 `xml:"SttlmAmt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Date/time at which the sender expects settlement.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// For against payment transactions, the value date/time at which the Sender expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate15Choice `xml:"SttlmDt"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails115 `xml:"SttlmParams"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties49 `xml:"RcvgSttlmPties,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties49 `xml:"DlvrgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification110 `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentification111 `xml:"QlfdFrgnIntrmy,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *RestrictedFINXMax350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (t *TransactionDetails87) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (t *TransactionDetails87) SetPoolIdentification(value string) {
	t.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionDetails87) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionDetails87) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionDetails87) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionDetails87) SetClientCollateralInstructionIdentification(value string) {
	t.ClientCollateralInstructionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionDetails87) SetTripartyCollateralInstructionIdentification(value string) {
	t.TripartyCollateralInstructionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionDetails87) AddAccountOwner() *PartyIdentification109 {
	t.AccountOwner = new(PartyIdentification109)
	return t.AccountOwner
}

func (t *TransactionDetails87) AddSafekeepingAccount() *SecuritiesAccount27 {
	t.SafekeepingAccount = new(SecuritiesAccount27)
	return t.SafekeepingAccount
}

func (t *TransactionDetails87) AddSafekeepingPlace() *SafeKeepingPlace2 {
	t.SafekeepingPlace = new(SafeKeepingPlace2)
	return t.SafekeepingPlace
}

func (t *TransactionDetails87) AddPlaceOfTrade() *PlaceOfTradeIdentification2 {
	t.PlaceOfTrade = new(PlaceOfTradeIdentification2)
	return t.PlaceOfTrade
}

func (t *TransactionDetails87) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	t.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return t.PlaceOfClearing
}

func (t *TransactionDetails87) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails87) AddSettlementQuantity() *Quantity10Choice {
	t.SettlementQuantity = new(Quantity10Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails87) AddSettlementAmount() *AmountAndDirection67 {
	t.SettlementAmount = new(AmountAndDirection67)
	return t.SettlementAmount
}

func (t *TransactionDetails87) AddLateDeliveryDate() *DateAndDateTimeChoice {
	t.LateDeliveryDate = new(DateAndDateTimeChoice)
	return t.LateDeliveryDate
}

func (t *TransactionDetails87) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	t.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return t.ExpectedSettlementDate
}

func (t *TransactionDetails87) AddExpectedValueDate() *DateAndDateTimeChoice {
	t.ExpectedValueDate = new(DateAndDateTimeChoice)
	return t.ExpectedValueDate
}

func (t *TransactionDetails87) AddSettlementDate() *SettlementDate15Choice {
	t.SettlementDate = new(SettlementDate15Choice)
	return t.SettlementDate
}

func (t *TransactionDetails87) AddTradeDate() *TradeDate6Choice {
	t.TradeDate = new(TradeDate6Choice)
	return t.TradeDate
}

func (t *TransactionDetails87) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails87) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails87) AddSettlementParameters() *SettlementDetails115 {
	t.SettlementParameters = new(SettlementDetails115)
	return t.SettlementParameters
}

func (t *TransactionDetails87) AddReceivingSettlementParties() *SettlementParties49 {
	t.ReceivingSettlementParties = new(SettlementParties49)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails87) AddDeliveringSettlementParties() *SettlementParties49 {
	t.DeliveringSettlementParties = new(SettlementParties49)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails87) AddInvestor() *PartyIdentification110 {
	t.Investor = new(PartyIdentification110)
	return t.Investor
}

func (t *TransactionDetails87) AddQualifiedForeignIntermediary() *PartyIdentification111 {
	t.QualifiedForeignIntermediary = new(PartyIdentification111)
	return t.QualifiedForeignIntermediary
}

func (t *TransactionDetails87) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	t.SettlementInstructionProcessingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

// Identifies the details of the transaction.
type TransactionDetails88 struct {

	// Party that legally owns the account.
	AccountOwner *PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount27 `xml:"SfkpgAcct"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity10Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection67 `xml:"SttlmAmt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate15Choice `xml:"SttlmDt"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties49 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties49 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification110 `xml:"Invstr,omitempty"`
}

func (t *TransactionDetails88) AddAccountOwner() *PartyIdentification109 {
	t.AccountOwner = new(PartyIdentification109)
	return t.AccountOwner
}

func (t *TransactionDetails88) AddSafekeepingAccount() *SecuritiesAccount27 {
	t.SafekeepingAccount = new(SecuritiesAccount27)
	return t.SafekeepingAccount
}

func (t *TransactionDetails88) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails88) AddSettlementQuantity() *Quantity10Choice {
	t.SettlementQuantity = new(Quantity10Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails88) AddSettlementAmount() *AmountAndDirection67 {
	t.SettlementAmount = new(AmountAndDirection67)
	return t.SettlementAmount
}

func (t *TransactionDetails88) AddTradeDate() *TradeDate6Choice {
	t.TradeDate = new(TradeDate6Choice)
	return t.TradeDate
}

func (t *TransactionDetails88) AddSettlementDate() *SettlementDate15Choice {
	t.SettlementDate = new(SettlementDate15Choice)
	return t.SettlementDate
}

func (t *TransactionDetails88) AddDeliveringSettlementParties() *SettlementParties49 {
	t.DeliveringSettlementParties = new(SettlementParties49)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails88) AddReceivingSettlementParties() *SettlementParties49 {
	t.ReceivingSettlementParties = new(SettlementParties49)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails88) AddInvestor() *PartyIdentification110 {
	t.Investor = new(PartyIdentification110)
	return t.Investor
}

// Identifies the details of the transaction.
type TransactionDetails90 struct {

	// Specifies the type of activity to which this instruction relates.
	TransactionActivity *TransactionActivity4Choice `xml:"TxActvty"`

	// Choice of type for the transaction reported.
	SettlementTransactionOrCorporateActionEventType *SettlementOrCorporateActionEvent15Choice `xml:"SttlmTxOrCorpActnEvtTp,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails116 `xml:"SttlmParams,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification2 `xml:"PlcOfTrad,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Quantity of financial instrument (to be) posted to the safekeeping account.
	PostingQuantity *Quantity10Choice `xml:"PstngQty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *AmountAndDirection67 `xml:"PstngAmt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Date/time at which the sender expects settlement.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate15Choice `xml:"SttlmDt"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// For against payment transactions, the value date/time at which the account servicer expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties49 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties49 `xml:"RcvgSttlmPties,omitempty"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionAdditionalDetails *RestrictedFINXMax350Text `xml:"TxAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (t *TransactionDetails90) AddTransactionActivity() *TransactionActivity4Choice {
	t.TransactionActivity = new(TransactionActivity4Choice)
	return t.TransactionActivity
}

func (t *TransactionDetails90) AddSettlementTransactionOrCorporateActionEventType() *SettlementOrCorporateActionEvent15Choice {
	t.SettlementTransactionOrCorporateActionEventType = new(SettlementOrCorporateActionEvent15Choice)
	return t.SettlementTransactionOrCorporateActionEventType
}

func (t *TransactionDetails90) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails90) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails90) AddSettlementParameters() *SettlementDetails116 {
	t.SettlementParameters = new(SettlementDetails116)
	return t.SettlementParameters
}

func (t *TransactionDetails90) AddPlaceOfTrade() *PlaceOfTradeIdentification2 {
	t.PlaceOfTrade = new(PlaceOfTradeIdentification2)
	return t.PlaceOfTrade
}

func (t *TransactionDetails90) AddSafekeepingPlace() *SafeKeepingPlace2 {
	t.SafekeepingPlace = new(SafeKeepingPlace2)
	return t.SafekeepingPlace
}

func (t *TransactionDetails90) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	t.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return t.PlaceOfClearing
}

func (t *TransactionDetails90) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails90) AddPostingQuantity() *Quantity10Choice {
	t.PostingQuantity = new(Quantity10Choice)
	return t.PostingQuantity
}

func (t *TransactionDetails90) AddPostingAmount() *AmountAndDirection67 {
	t.PostingAmount = new(AmountAndDirection67)
	return t.PostingAmount
}

func (t *TransactionDetails90) AddTradeDate() *TradeDate6Choice {
	t.TradeDate = new(TradeDate6Choice)
	return t.TradeDate
}

func (t *TransactionDetails90) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	t.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return t.ExpectedSettlementDate
}

func (t *TransactionDetails90) AddSettlementDate() *SettlementDate15Choice {
	t.SettlementDate = new(SettlementDate15Choice)
	return t.SettlementDate
}

func (t *TransactionDetails90) AddLateDeliveryDate() *DateAndDateTimeChoice {
	t.LateDeliveryDate = new(DateAndDateTimeChoice)
	return t.LateDeliveryDate
}

func (t *TransactionDetails90) AddExpectedValueDate() *DateAndDateTimeChoice {
	t.ExpectedValueDate = new(DateAndDateTimeChoice)
	return t.ExpectedValueDate
}

func (t *TransactionDetails90) AddDeliveringSettlementParties() *SettlementParties49 {
	t.DeliveringSettlementParties = new(SettlementParties49)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails90) AddReceivingSettlementParties() *SettlementParties49 {
	t.ReceivingSettlementParties = new(SettlementParties49)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails90) SetTransactionAdditionalDetails(value string) {
	t.TransactionAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

func (t *TransactionDetails90) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	t.SupplementaryData = append(t.SupplementaryData, newValue)
	return newValue
}

// Identifies the details of the transaction.
type TransactionDetails91 struct {

	// Specifies the type of activity to which this instruction relates.
	TransactionActivity *TransactionActivity4Choice `xml:"TxActvty"`

	// Choice of type for the transaction reported.
	SettlementTransactionOrCorporateActionEventType *SettlementOrCorporateActionEvent16Choice `xml:"SttlmTxOrCorpActnEvtTp,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails117 `xml:"SttlmParams,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification2 `xml:"PlcOfTrad,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Quantity of financial instrument (to be) posted to the safekeeping account.
	PostingQuantity *Quantity10Choice `xml:"PstngQty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *AmountAndDirection18 `xml:"PstngAmt,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection59 `xml:"AcrdIntrstAmt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, that is, payment is effected and securities are delivered.
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate12Choice `xml:"SttlmDt,omitempty"`

	// Date and time assets become available to the account owner (in a credit entry), or cease to be available to the account owner (in a debit entry).
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties49 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties49 `xml:"RcvgSttlmPties,omitempty"`

	// Specifies whether it is the reversal of a previously reported movement.
	ReversalIndicator *YesNoIndicator `xml:"RvslInd,omitempty"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionAdditionalDetails *RestrictedFINXMax350Text `xml:"TxAddtlDtls,omitempty"`
}

func (t *TransactionDetails91) AddTransactionActivity() *TransactionActivity4Choice {
	t.TransactionActivity = new(TransactionActivity4Choice)
	return t.TransactionActivity
}

func (t *TransactionDetails91) AddSettlementTransactionOrCorporateActionEventType() *SettlementOrCorporateActionEvent16Choice {
	t.SettlementTransactionOrCorporateActionEventType = new(SettlementOrCorporateActionEvent16Choice)
	return t.SettlementTransactionOrCorporateActionEventType
}

func (t *TransactionDetails91) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails91) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails91) AddSettlementParameters() *SettlementDetails117 {
	t.SettlementParameters = new(SettlementDetails117)
	return t.SettlementParameters
}

func (t *TransactionDetails91) AddPlaceOfTrade() *PlaceOfTradeIdentification2 {
	t.PlaceOfTrade = new(PlaceOfTradeIdentification2)
	return t.PlaceOfTrade
}

func (t *TransactionDetails91) AddSafekeepingPlace() *SafeKeepingPlace2 {
	t.SafekeepingPlace = new(SafeKeepingPlace2)
	return t.SafekeepingPlace
}

func (t *TransactionDetails91) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	t.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return t.PlaceOfClearing
}

func (t *TransactionDetails91) AddPostingQuantity() *Quantity10Choice {
	t.PostingQuantity = new(Quantity10Choice)
	return t.PostingQuantity
}

func (t *TransactionDetails91) SetNumberOfDaysAccrued(value string) {
	t.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (t *TransactionDetails91) AddPostingAmount() *AmountAndDirection18 {
	t.PostingAmount = new(AmountAndDirection18)
	return t.PostingAmount
}

func (t *TransactionDetails91) AddAccruedInterestAmount() *AmountAndDirection59 {
	t.AccruedInterestAmount = new(AmountAndDirection59)
	return t.AccruedInterestAmount
}

func (t *TransactionDetails91) AddTradeDate() *TradeDate6Choice {
	t.TradeDate = new(TradeDate6Choice)
	return t.TradeDate
}

func (t *TransactionDetails91) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	t.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return t.EffectiveSettlementDate
}

func (t *TransactionDetails91) AddSettlementDate() *SettlementDate12Choice {
	t.SettlementDate = new(SettlementDate12Choice)
	return t.SettlementDate
}

func (t *TransactionDetails91) AddValueDate() *DateAndDateTimeChoice {
	t.ValueDate = new(DateAndDateTimeChoice)
	return t.ValueDate
}

func (t *TransactionDetails91) AddDeliveringSettlementParties() *SettlementParties49 {
	t.DeliveringSettlementParties = new(SettlementParties49)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails91) AddReceivingSettlementParties() *SettlementParties49 {
	t.ReceivingSettlementParties = new(SettlementParties49)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails91) SetReversalIndicator(value string) {
	t.ReversalIndicator = (*YesNoIndicator)(&value)
}

func (t *TransactionDetails91) SetTransactionAdditionalDetails(value string) {
	t.TransactionAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

// Identifies the details of the transaction.
type TransactionDetails92 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, eg, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity10Choice `xml:"SttlmQty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount30 `xml:"SfkpgAcct"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection67 `xml:"SttlmAmt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate15Choice `xml:"SttlmDt"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties49 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties49 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification110 `xml:"Invstr,omitempty"`
}

func (t *TransactionDetails92) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails92) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails92) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails92) AddSettlementQuantity() *Quantity10Choice {
	t.SettlementQuantity = new(Quantity10Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails92) AddSafekeepingAccount() *SecuritiesAccount30 {
	t.SafekeepingAccount = new(SecuritiesAccount30)
	return t.SafekeepingAccount
}

func (t *TransactionDetails92) AddSettlementAmount() *AmountAndDirection67 {
	t.SettlementAmount = new(AmountAndDirection67)
	return t.SettlementAmount
}

func (t *TransactionDetails92) AddSettlementDate() *SettlementDate15Choice {
	t.SettlementDate = new(SettlementDate15Choice)
	return t.SettlementDate
}

func (t *TransactionDetails92) AddTradeDate() *TradeDate6Choice {
	t.TradeDate = new(TradeDate6Choice)
	return t.TradeDate
}

func (t *TransactionDetails92) AddDeliveringSettlementParties() *SettlementParties49 {
	t.DeliveringSettlementParties = new(SettlementParties49)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails92) AddReceivingSettlementParties() *SettlementParties49 {
	t.ReceivingSettlementParties = new(SettlementParties49)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails92) AddInvestor() *PartyIdentification110 {
	t.Investor = new(PartyIdentification110)
	return t.Investor
}

// Specifies the details of the transaction.
type TransactionDetails95 struct {

	// Specifies the type of activity to which this instruction relates.
	TransactionActivity *TransactionActivity3Choice `xml:"TxActvty"`

	// Choice of type for the transaction reported.
	SettlementTransactionOrCorporateActionEventType *SettlementOrCorporateActionEvent19Choice `xml:"SttlmTxOrCorpActnEvtTp,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails126 `xml:"SttlmParams,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification1 `xml:"PlcOfTrad,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Quantity of financial instrument (to be) posted to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *AmountAndDirection3 `xml:"PstngAmt,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection21 `xml:"AcrdIntrstAmt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, that is, payment is effected and securities are delivered.
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate9Choice `xml:"SttlmDt,omitempty"`

	// Date and time assets become available to the account owner (in a credit entry), or cease to be available to the account owner (in a debit entry).
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`

	// Time stamp on when the transaction is acknowledged.
	AcknowledgedStatusTimeStamp *ISODateTime `xml:"AckdStsTmStmp,omitempty"`

	// Time stamp on when the transaction is matched.
	MatchedStatusTimeStamp *ISODateTime `xml:"MtchdStsTmStmp,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties40 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties40 `xml:"RcvgSttlmPties,omitempty"`

	// Specifies whether it is the reversal of a previously reported movement.
	ReversalIndicator *YesNoIndicator `xml:"RvslInd,omitempty"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionAdditionalDetails *Max350Text `xml:"TxAddtlDtls,omitempty"`
}

func (t *TransactionDetails95) AddTransactionActivity() *TransactionActivity3Choice {
	t.TransactionActivity = new(TransactionActivity3Choice)
	return t.TransactionActivity
}

func (t *TransactionDetails95) AddSettlementTransactionOrCorporateActionEventType() *SettlementOrCorporateActionEvent19Choice {
	t.SettlementTransactionOrCorporateActionEventType = new(SettlementOrCorporateActionEvent19Choice)
	return t.SettlementTransactionOrCorporateActionEventType
}

func (t *TransactionDetails95) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails95) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails95) AddSettlementParameters() *SettlementDetails126 {
	t.SettlementParameters = new(SettlementDetails126)
	return t.SettlementParameters
}

func (t *TransactionDetails95) AddPlaceOfTrade() *PlaceOfTradeIdentification1 {
	t.PlaceOfTrade = new(PlaceOfTradeIdentification1)
	return t.PlaceOfTrade
}

func (t *TransactionDetails95) AddSafekeepingPlace() *SafeKeepingPlace1 {
	t.SafekeepingPlace = new(SafeKeepingPlace1)
	return t.SafekeepingPlace
}

func (t *TransactionDetails95) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	t.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return t.PlaceOfClearing
}

func (t *TransactionDetails95) AddPostingQuantity() *Quantity6Choice {
	t.PostingQuantity = new(Quantity6Choice)
	return t.PostingQuantity
}

func (t *TransactionDetails95) SetNumberOfDaysAccrued(value string) {
	t.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (t *TransactionDetails95) AddPostingAmount() *AmountAndDirection3 {
	t.PostingAmount = new(AmountAndDirection3)
	return t.PostingAmount
}

func (t *TransactionDetails95) AddAccruedInterestAmount() *AmountAndDirection21 {
	t.AccruedInterestAmount = new(AmountAndDirection21)
	return t.AccruedInterestAmount
}

func (t *TransactionDetails95) AddTradeDate() *TradeDate5Choice {
	t.TradeDate = new(TradeDate5Choice)
	return t.TradeDate
}

func (t *TransactionDetails95) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	t.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return t.EffectiveSettlementDate
}

func (t *TransactionDetails95) AddSettlementDate() *SettlementDate9Choice {
	t.SettlementDate = new(SettlementDate9Choice)
	return t.SettlementDate
}

func (t *TransactionDetails95) AddValueDate() *DateAndDateTimeChoice {
	t.ValueDate = new(DateAndDateTimeChoice)
	return t.ValueDate
}

func (t *TransactionDetails95) SetAcknowledgedStatusTimeStamp(value string) {
	t.AcknowledgedStatusTimeStamp = (*ISODateTime)(&value)
}

func (t *TransactionDetails95) SetMatchedStatusTimeStamp(value string) {
	t.MatchedStatusTimeStamp = (*ISODateTime)(&value)
}

func (t *TransactionDetails95) AddDeliveringSettlementParties() *SettlementParties40 {
	t.DeliveringSettlementParties = new(SettlementParties40)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails95) AddReceivingSettlementParties() *SettlementParties40 {
	t.ReceivingSettlementParties = new(SettlementParties40)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails95) SetReversalIndicator(value string) {
	t.ReversalIndicator = (*YesNoIndicator)(&value)
}

func (t *TransactionDetails95) SetTransactionAdditionalDetails(value string) {
	t.TransactionAdditionalDetails = (*Max350Text)(&value)
}

// Specifies the details of the transaction.
type TransactionDetails96 struct {

	// Specifies the type of activity to which this instruction relates.
	TransactionActivity *TransactionActivity3Choice `xml:"TxActvty"`

	// Choice of type for the transaction reported.
	SettlementTransactionOrCorporateActionEventType *SettlementOrCorporateActionEvent17Choice `xml:"SttlmTxOrCorpActnEvtTp,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails127 `xml:"SttlmParams,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification1 `xml:"PlcOfTrad,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Quantity of financial instrument (to be) posted to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *AmountAndDirection51 `xml:"PstngAmt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Date/time at which the sender expects settlement.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate10Choice `xml:"SttlmDt"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// For against payment transactions, the value date/time at which the account servicer expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Time stamp on when the transaction is acknowledged.
	AcknowledgedStatusTimeStamp *ISODateTime `xml:"AckdStsTmStmp,omitempty"`

	// Time stamp on when the transaction is matched.
	MatchedStatusTimeStamp *ISODateTime `xml:"MtchdStsTmStmp,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties40 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties40 `xml:"RcvgSttlmPties,omitempty"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionAdditionalDetails *Max350Text `xml:"TxAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (t *TransactionDetails96) AddTransactionActivity() *TransactionActivity3Choice {
	t.TransactionActivity = new(TransactionActivity3Choice)
	return t.TransactionActivity
}

func (t *TransactionDetails96) AddSettlementTransactionOrCorporateActionEventType() *SettlementOrCorporateActionEvent17Choice {
	t.SettlementTransactionOrCorporateActionEventType = new(SettlementOrCorporateActionEvent17Choice)
	return t.SettlementTransactionOrCorporateActionEventType
}

func (t *TransactionDetails96) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails96) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails96) AddSettlementParameters() *SettlementDetails127 {
	t.SettlementParameters = new(SettlementDetails127)
	return t.SettlementParameters
}

func (t *TransactionDetails96) AddPlaceOfTrade() *PlaceOfTradeIdentification1 {
	t.PlaceOfTrade = new(PlaceOfTradeIdentification1)
	return t.PlaceOfTrade
}

func (t *TransactionDetails96) AddSafekeepingPlace() *SafeKeepingPlace1 {
	t.SafekeepingPlace = new(SafeKeepingPlace1)
	return t.SafekeepingPlace
}

func (t *TransactionDetails96) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	t.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return t.PlaceOfClearing
}

func (t *TransactionDetails96) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails96) AddPostingQuantity() *Quantity6Choice {
	t.PostingQuantity = new(Quantity6Choice)
	return t.PostingQuantity
}

func (t *TransactionDetails96) AddPostingAmount() *AmountAndDirection51 {
	t.PostingAmount = new(AmountAndDirection51)
	return t.PostingAmount
}

func (t *TransactionDetails96) AddTradeDate() *TradeDate5Choice {
	t.TradeDate = new(TradeDate5Choice)
	return t.TradeDate
}

func (t *TransactionDetails96) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	t.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return t.ExpectedSettlementDate
}

func (t *TransactionDetails96) AddSettlementDate() *SettlementDate10Choice {
	t.SettlementDate = new(SettlementDate10Choice)
	return t.SettlementDate
}

func (t *TransactionDetails96) AddLateDeliveryDate() *DateAndDateTimeChoice {
	t.LateDeliveryDate = new(DateAndDateTimeChoice)
	return t.LateDeliveryDate
}

func (t *TransactionDetails96) AddExpectedValueDate() *DateAndDateTimeChoice {
	t.ExpectedValueDate = new(DateAndDateTimeChoice)
	return t.ExpectedValueDate
}

func (t *TransactionDetails96) SetAcknowledgedStatusTimeStamp(value string) {
	t.AcknowledgedStatusTimeStamp = (*ISODateTime)(&value)
}

func (t *TransactionDetails96) SetMatchedStatusTimeStamp(value string) {
	t.MatchedStatusTimeStamp = (*ISODateTime)(&value)
}

func (t *TransactionDetails96) AddDeliveringSettlementParties() *SettlementParties40 {
	t.DeliveringSettlementParties = new(SettlementParties40)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails96) AddReceivingSettlementParties() *SettlementParties40 {
	t.ReceivingSettlementParties = new(SettlementParties40)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails96) SetTransactionAdditionalDetails(value string) {
	t.TransactionAdditionalDetails = (*Max350Text)(&value)
}

func (t *TransactionDetails96) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	t.SupplementaryData = append(t.SupplementaryData, newValue)
	return newValue
}

// Specifies the details of the transaction.
type TransactionDetails97 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Unique identification identifying the triparty collateral management transaction from the triparty-agent's/service-provider's point of view.
	TripartyAgentServiceProviderCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtSvcPrvdrCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Unique identification assigned to the instruction by the client.
	ClientCollateralInstructionIdentification *Max35Text `xml:"ClntCollInstrId,omitempty"`

	// Unique identification assigned to the instruction by the triparty-agent/service-provider.
	TripartyAgentServiceProviderCollateralInstructionIdentification *Max35Text `xml:"TrptyAgtSvcPrvdrCollInstrId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount24 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification1 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection51 `xml:"SttlmAmt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Date/time at which the sender expects settlement.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// For against payment transactions, the value date/time at which the Sender expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate10Choice `xml:"SttlmDt"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Time stamp on when the transaction is acknowledged.
	AcknowledgedStatusTimeStamp *ISODateTime `xml:"AckdStsTmStmp,omitempty"`

	// Time stamp on when the transaction is matched.
	MatchedStatusTimeStamp *ISODateTime `xml:"MtchdStsTmStmp,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails121 `xml:"SttlmParams"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties40 `xml:"RcvgSttlmPties,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties40 `xml:"DlvrgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification99 `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentification100 `xml:"QlfdFrgnIntrmy,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (t *TransactionDetails97) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *TransactionDetails97) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails97) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails97) SetTripartyAgentServiceProviderCollateralTransactionIdentification(value string) {
	t.TripartyAgentServiceProviderCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails97) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails97) SetClientCollateralInstructionIdentification(value string) {
	t.ClientCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails97) SetTripartyAgentServiceProviderCollateralInstructionIdentification(value string) {
	t.TripartyAgentServiceProviderCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails97) AddAccountOwner() *PartyIdentification98 {
	t.AccountOwner = new(PartyIdentification98)
	return t.AccountOwner
}

func (t *TransactionDetails97) AddSafekeepingAccount() *SecuritiesAccount24 {
	t.SafekeepingAccount = new(SecuritiesAccount24)
	return t.SafekeepingAccount
}

func (t *TransactionDetails97) AddSafekeepingPlace() *SafeKeepingPlace1 {
	t.SafekeepingPlace = new(SafeKeepingPlace1)
	return t.SafekeepingPlace
}

func (t *TransactionDetails97) AddPlaceOfTrade() *PlaceOfTradeIdentification1 {
	t.PlaceOfTrade = new(PlaceOfTradeIdentification1)
	return t.PlaceOfTrade
}

func (t *TransactionDetails97) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	t.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return t.PlaceOfClearing
}

func (t *TransactionDetails97) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails97) AddSettlementQuantity() *Quantity6Choice {
	t.SettlementQuantity = new(Quantity6Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails97) AddSettlementAmount() *AmountAndDirection51 {
	t.SettlementAmount = new(AmountAndDirection51)
	return t.SettlementAmount
}

func (t *TransactionDetails97) AddLateDeliveryDate() *DateAndDateTimeChoice {
	t.LateDeliveryDate = new(DateAndDateTimeChoice)
	return t.LateDeliveryDate
}

func (t *TransactionDetails97) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	t.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return t.ExpectedSettlementDate
}

func (t *TransactionDetails97) AddExpectedValueDate() *DateAndDateTimeChoice {
	t.ExpectedValueDate = new(DateAndDateTimeChoice)
	return t.ExpectedValueDate
}

func (t *TransactionDetails97) AddSettlementDate() *SettlementDate10Choice {
	t.SettlementDate = new(SettlementDate10Choice)
	return t.SettlementDate
}

func (t *TransactionDetails97) AddTradeDate() *TradeDate5Choice {
	t.TradeDate = new(TradeDate5Choice)
	return t.TradeDate
}

func (t *TransactionDetails97) SetAcknowledgedStatusTimeStamp(value string) {
	t.AcknowledgedStatusTimeStamp = (*ISODateTime)(&value)
}

func (t *TransactionDetails97) SetMatchedStatusTimeStamp(value string) {
	t.MatchedStatusTimeStamp = (*ISODateTime)(&value)
}

func (t *TransactionDetails97) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails97) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails97) AddSettlementParameters() *SettlementDetails121 {
	t.SettlementParameters = new(SettlementDetails121)
	return t.SettlementParameters
}

func (t *TransactionDetails97) AddReceivingSettlementParties() *SettlementParties40 {
	t.ReceivingSettlementParties = new(SettlementParties40)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails97) AddDeliveringSettlementParties() *SettlementParties40 {
	t.DeliveringSettlementParties = new(SettlementParties40)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails97) AddInvestor() *PartyIdentification99 {
	t.Investor = new(PartyIdentification99)
	return t.Investor
}

func (t *TransactionDetails97) AddQualifiedForeignIntermediary() *PartyIdentification100 {
	t.QualifiedForeignIntermediary = new(PartyIdentification100)
	return t.QualifiedForeignIntermediary
}

func (t *TransactionDetails97) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	t.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Specifies the details of the transaction.
type TransactionDetails98 struct {

	// Specifies the type of activity to which this instruction relates.
	TransactionActivity *TransactionActivity4Choice `xml:"TxActvty"`

	// Choice of type for the transaction reported.
	SettlementTransactionOrCorporateActionEventType *SettlementOrCorporateActionEvent20Choice `xml:"SttlmTxOrCorpActnEvtTp,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails139 `xml:"SttlmParams,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification2 `xml:"PlcOfTrad,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Quantity of financial instrument (to be) posted to the safekeeping account.
	PostingQuantity *Quantity10Choice `xml:"PstngQty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *AmountAndDirection18 `xml:"PstngAmt,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection59 `xml:"AcrdIntrstAmt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, that is, payment is effected and securities are delivered.
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate12Choice `xml:"SttlmDt,omitempty"`

	// Date and time assets become available to the account owner (in a credit entry), or cease to be available to the account owner (in a debit entry).
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`

	// Time stamp on when the transaction is acknowledged.
	AcknowledgedStatusTimeStamp *ISODateTime `xml:"AckdStsTmStmp,omitempty"`

	// Time stamp on when the transaction is matched.
	MatchedStatusTimeStamp *ISODateTime `xml:"MtchdStsTmStmp,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties49 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties49 `xml:"RcvgSttlmPties,omitempty"`

	// Specifies whether it is the reversal of a previously reported movement.
	ReversalIndicator *YesNoIndicator `xml:"RvslInd,omitempty"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionAdditionalDetails *RestrictedFINXMax350Text `xml:"TxAddtlDtls,omitempty"`
}

func (t *TransactionDetails98) AddTransactionActivity() *TransactionActivity4Choice {
	t.TransactionActivity = new(TransactionActivity4Choice)
	return t.TransactionActivity
}

func (t *TransactionDetails98) AddSettlementTransactionOrCorporateActionEventType() *SettlementOrCorporateActionEvent20Choice {
	t.SettlementTransactionOrCorporateActionEventType = new(SettlementOrCorporateActionEvent20Choice)
	return t.SettlementTransactionOrCorporateActionEventType
}

func (t *TransactionDetails98) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails98) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails98) AddSettlementParameters() *SettlementDetails139 {
	t.SettlementParameters = new(SettlementDetails139)
	return t.SettlementParameters
}

func (t *TransactionDetails98) AddPlaceOfTrade() *PlaceOfTradeIdentification2 {
	t.PlaceOfTrade = new(PlaceOfTradeIdentification2)
	return t.PlaceOfTrade
}

func (t *TransactionDetails98) AddSafekeepingPlace() *SafeKeepingPlace2 {
	t.SafekeepingPlace = new(SafeKeepingPlace2)
	return t.SafekeepingPlace
}

func (t *TransactionDetails98) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	t.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return t.PlaceOfClearing
}

func (t *TransactionDetails98) AddPostingQuantity() *Quantity10Choice {
	t.PostingQuantity = new(Quantity10Choice)
	return t.PostingQuantity
}

func (t *TransactionDetails98) SetNumberOfDaysAccrued(value string) {
	t.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (t *TransactionDetails98) AddPostingAmount() *AmountAndDirection18 {
	t.PostingAmount = new(AmountAndDirection18)
	return t.PostingAmount
}

func (t *TransactionDetails98) AddAccruedInterestAmount() *AmountAndDirection59 {
	t.AccruedInterestAmount = new(AmountAndDirection59)
	return t.AccruedInterestAmount
}

func (t *TransactionDetails98) AddTradeDate() *TradeDate6Choice {
	t.TradeDate = new(TradeDate6Choice)
	return t.TradeDate
}

func (t *TransactionDetails98) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	t.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return t.EffectiveSettlementDate
}

func (t *TransactionDetails98) AddSettlementDate() *SettlementDate12Choice {
	t.SettlementDate = new(SettlementDate12Choice)
	return t.SettlementDate
}

func (t *TransactionDetails98) AddValueDate() *DateAndDateTimeChoice {
	t.ValueDate = new(DateAndDateTimeChoice)
	return t.ValueDate
}

func (t *TransactionDetails98) SetAcknowledgedStatusTimeStamp(value string) {
	t.AcknowledgedStatusTimeStamp = (*ISODateTime)(&value)
}

func (t *TransactionDetails98) SetMatchedStatusTimeStamp(value string) {
	t.MatchedStatusTimeStamp = (*ISODateTime)(&value)
}

func (t *TransactionDetails98) AddDeliveringSettlementParties() *SettlementParties49 {
	t.DeliveringSettlementParties = new(SettlementParties49)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails98) AddReceivingSettlementParties() *SettlementParties49 {
	t.ReceivingSettlementParties = new(SettlementParties49)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails98) SetReversalIndicator(value string) {
	t.ReversalIndicator = (*YesNoIndicator)(&value)
}

func (t *TransactionDetails98) SetTransactionAdditionalDetails(value string) {
	t.TransactionAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

// Specifies the details of the transaction.
type TransactionDetails99 struct {

	// Specifies the type of activity to which this instruction relates.
	TransactionActivity *TransactionActivity4Choice `xml:"TxActvty"`

	// Choice of type for the transaction reported.
	SettlementTransactionOrCorporateActionEventType *SettlementOrCorporateActionEvent21Choice `xml:"SttlmTxOrCorpActnEvtTp,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails130 `xml:"SttlmParams,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification2 `xml:"PlcOfTrad,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Quantity of financial instrument (to be) posted to the safekeeping account.
	PostingQuantity *Quantity10Choice `xml:"PstngQty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *AmountAndDirection67 `xml:"PstngAmt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Date/time at which the sender expects settlement.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate15Choice `xml:"SttlmDt"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// For against payment transactions, the value date/time at which the account servicer expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Time stamp on when the transaction is acknowledged.
	AcknowledgedStatusTimeStamp *ISODateTime `xml:"AckdStsTmStmp,omitempty"`

	// Time stamp on when the transaction is matched.
	MatchedStatusTimeStamp *ISODateTime `xml:"MtchdStsTmStmp,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties49 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties49 `xml:"RcvgSttlmPties,omitempty"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionAdditionalDetails *RestrictedFINXMax350Text `xml:"TxAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (t *TransactionDetails99) AddTransactionActivity() *TransactionActivity4Choice {
	t.TransactionActivity = new(TransactionActivity4Choice)
	return t.TransactionActivity
}

func (t *TransactionDetails99) AddSettlementTransactionOrCorporateActionEventType() *SettlementOrCorporateActionEvent21Choice {
	t.SettlementTransactionOrCorporateActionEventType = new(SettlementOrCorporateActionEvent21Choice)
	return t.SettlementTransactionOrCorporateActionEventType
}

func (t *TransactionDetails99) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails99) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails99) AddSettlementParameters() *SettlementDetails130 {
	t.SettlementParameters = new(SettlementDetails130)
	return t.SettlementParameters
}

func (t *TransactionDetails99) AddPlaceOfTrade() *PlaceOfTradeIdentification2 {
	t.PlaceOfTrade = new(PlaceOfTradeIdentification2)
	return t.PlaceOfTrade
}

func (t *TransactionDetails99) AddSafekeepingPlace() *SafeKeepingPlace2 {
	t.SafekeepingPlace = new(SafeKeepingPlace2)
	return t.SafekeepingPlace
}

func (t *TransactionDetails99) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	t.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return t.PlaceOfClearing
}

func (t *TransactionDetails99) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails99) AddPostingQuantity() *Quantity10Choice {
	t.PostingQuantity = new(Quantity10Choice)
	return t.PostingQuantity
}

func (t *TransactionDetails99) AddPostingAmount() *AmountAndDirection67 {
	t.PostingAmount = new(AmountAndDirection67)
	return t.PostingAmount
}

func (t *TransactionDetails99) AddTradeDate() *TradeDate6Choice {
	t.TradeDate = new(TradeDate6Choice)
	return t.TradeDate
}

func (t *TransactionDetails99) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	t.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return t.ExpectedSettlementDate
}

func (t *TransactionDetails99) AddSettlementDate() *SettlementDate15Choice {
	t.SettlementDate = new(SettlementDate15Choice)
	return t.SettlementDate
}

func (t *TransactionDetails99) AddLateDeliveryDate() *DateAndDateTimeChoice {
	t.LateDeliveryDate = new(DateAndDateTimeChoice)
	return t.LateDeliveryDate
}

func (t *TransactionDetails99) AddExpectedValueDate() *DateAndDateTimeChoice {
	t.ExpectedValueDate = new(DateAndDateTimeChoice)
	return t.ExpectedValueDate
}

func (t *TransactionDetails99) SetAcknowledgedStatusTimeStamp(value string) {
	t.AcknowledgedStatusTimeStamp = (*ISODateTime)(&value)
}

func (t *TransactionDetails99) SetMatchedStatusTimeStamp(value string) {
	t.MatchedStatusTimeStamp = (*ISODateTime)(&value)
}

func (t *TransactionDetails99) AddDeliveringSettlementParties() *SettlementParties49 {
	t.DeliveringSettlementParties = new(SettlementParties49)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails99) AddReceivingSettlementParties() *SettlementParties49 {
	t.ReceivingSettlementParties = new(SettlementParties49)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails99) SetTransactionAdditionalDetails(value string) {
	t.TransactionAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

func (t *TransactionDetails99) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	t.SupplementaryData = append(t.SupplementaryData, newValue)
	return newValue
}
