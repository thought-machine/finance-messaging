package iso20022

// Provides transaction type and identification information.
type SettlementTypeAndIdentification1 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`
}

func (s *SettlementTypeAndIdentification1) SetAccountOwnerTransactionIdentification(value string) {
	s.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification1) SetAccountServicerTransactionIdentification(value string) {
	s.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification1) SetMarketInfrastructureTransactionIdentification(value string) {
	s.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification1) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndIdentification1) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementTypeAndIdentification1) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification1) SetPoolIdentification(value string) {
	s.PoolIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification1) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*Max35Text)(&value)
}

// Provides transaction type and identification information.
type SettlementTypeAndIdentification10 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`
}

func (s *SettlementTypeAndIdentification10) SetAccountOwnerTransactionIdentification(value string) {
	s.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification10) SetAccountServicerTransactionIdentification(value string) {
	s.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification10) SetMarketInfrastructureTransactionIdentification(value string) {
	s.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification10) SetProcessorTransactionIdentification(value string) {
	s.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification10) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndIdentification10) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementTypeAndIdentification10) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification10) SetPoolIdentification(value string) {
	s.PoolIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification10) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*Max35Text)(&value)
}

// Provides transaction type and identification information.
type SettlementTypeAndIdentification13 struct {

	// Provides unambiguous transaction identification information.
	TransactionIdentification *Max35Text `xml:"TxId"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`
}

func (s *SettlementTypeAndIdentification13) SetTransactionIdentification(value string) {
	s.TransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification13) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndIdentification13) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

// Provides transaction type and identification information.
type SettlementTypeAndIdentification15 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`
}

func (s *SettlementTypeAndIdentification15) SetAccountOwnerTransactionIdentification(value string) {
	s.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification15) SetAccountServicerTransactionIdentification(value string) {
	s.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification15) SetMarketInfrastructureTransactionIdentification(value string) {
	s.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification15) SetProcessorTransactionIdentification(value string) {
	s.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification15) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndIdentification15) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementTypeAndIdentification15) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification15) SetPoolIdentification(value string) {
	s.PoolIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification15) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*Max35Text)(&value)
}

// Provides transaction type and identification information.
type SettlementTypeAndIdentification18 struct {

	// Provides unambiguous transaction identification information.
	TransactionIdentification *Max35Text `xml:"TxId"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`
}

func (s *SettlementTypeAndIdentification18) SetTransactionIdentification(value string) {
	s.TransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification18) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndIdentification18) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

// Provides transaction type and identification information.
type SettlementTypeAndIdentification19 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`
}

func (s *SettlementTypeAndIdentification19) SetAccountOwnerTransactionIdentification(value string) {
	s.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification19) SetAccountServicerTransactionIdentification(value string) {
	s.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification19) SetMarketInfrastructureTransactionIdentification(value string) {
	s.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification19) SetProcessorTransactionIdentification(value string) {
	s.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification19) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndIdentification19) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementTypeAndIdentification19) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification19) SetPoolIdentification(value string) {
	s.PoolIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification19) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*Max35Text)(&value)
}

// Overall process covering the trade and settlement transactions of financial instruments.
type SettlementTypeAndIdentification2 struct {

	// Specifies how the transaction is to be settled.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Reference of the transaction.
	TransactionIdentification *Max35Text `xml:"TxId"`

	// Identifies the intended settlement date.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt,omitempty"`
}

func (s *SettlementTypeAndIdentification2) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementTypeAndIdentification2) SetTransactionIdentification(value string) {
	s.TransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification2) AddSettlementDate() *DateAndDateTimeChoice {
	s.SettlementDate = new(DateAndDateTimeChoice)
	return s.SettlementDate
}

// Overall process covering the trade and settlement transactions of financial instruments.
type SettlementTypeAndIdentification20 struct {

	// Specifies how the transaction is to be settled.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Reference of the transaction.
	TransactionIdentification *Max35Text `xml:"TxId"`

	// Identifies the intended settlement date.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt,omitempty"`
}

func (s *SettlementTypeAndIdentification20) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementTypeAndIdentification20) SetTransactionIdentification(value string) {
	s.TransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification20) AddSettlementDate() *DateAndDateTimeChoice {
	s.SettlementDate = new(DateAndDateTimeChoice)
	return s.SettlementDate
}

// Overall process covering the trade and settlement transactions of financial instruments.
type SettlementTypeAndIdentification21 struct {

	// Specifies how the transaction is to be settled.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Reference of the transaction.
	TransactionIdentification *RestrictedFINXMax16Text `xml:"TxId"`

	// Identifies the intended settlement date.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt,omitempty"`
}

func (s *SettlementTypeAndIdentification21) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementTypeAndIdentification21) SetTransactionIdentification(value string) {
	s.TransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndIdentification21) AddSettlementDate() *DateAndDateTimeChoice {
	s.SettlementDate = new(DateAndDateTimeChoice)
	return s.SettlementDate
}

// Provides transaction type and identification information.
type SettlementTypeAndIdentification22 struct {

	// Provides unambiguous transaction identification information.
	TransactionIdentification *RestrictedFINXMax16Text `xml:"TxId"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`
}

func (s *SettlementTypeAndIdentification22) SetTransactionIdentification(value string) {
	s.TransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndIdentification22) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndIdentification22) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

// Provides transaction type and identification information.
type SettlementTypeAndIdentification24 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *RestrictedFINXMax16Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *RestrictedFINXMax16Text `xml:"PrcrTxId,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId,omitempty"`
}

func (s *SettlementTypeAndIdentification24) SetAccountOwnerTransactionIdentification(value string) {
	s.AccountOwnerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndIdentification24) SetAccountServicerTransactionIdentification(value string) {
	s.AccountServicerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndIdentification24) SetMarketInfrastructureTransactionIdentification(value string) {
	s.MarketInfrastructureTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndIdentification24) SetProcessorTransactionIdentification(value string) {
	s.ProcessorTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndIdentification24) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndIdentification24) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementTypeAndIdentification24) SetCommonIdentification(value string) {
	s.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndIdentification24) SetPoolIdentification(value string) {
	s.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndIdentification24) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

// Provides transaction type and identification information.
type SettlementTypeAndIdentification3 struct {

	// Provides unambiguous transaction identification information.
	TransactionIdentification *Max35Text `xml:"TxId"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`
}

func (s *SettlementTypeAndIdentification3) SetTransactionIdentification(value string) {
	s.TransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification3) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndIdentification3) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

// Provides transaction type and identification information.
type SettlementTypeAndIdentification4 struct {

	// Provides unambiguous transaction identification information.
	TransactionIdentification *Max35Text `xml:"TxId"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp,omitempty"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt,omitempty"`
}

func (s *SettlementTypeAndIdentification4) SetTransactionIdentification(value string) {
	s.TransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification4) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndIdentification4) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

// Provides transaction type and identification information.
type SettlementTypeAndIdentification9 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`
}

func (s *SettlementTypeAndIdentification9) SetAccountOwnerTransactionIdentification(value string) {
	s.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification9) SetAccountServicerTransactionIdentification(value string) {
	s.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification9) SetMarketInfrastructureTransactionIdentification(value string) {
	s.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification9) SetProcessorTransactionIdentification(value string) {
	s.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification9) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndIdentification9) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementTypeAndIdentification9) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification9) SetPoolIdentification(value string) {
	s.PoolIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndIdentification9) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*Max35Text)(&value)
}
