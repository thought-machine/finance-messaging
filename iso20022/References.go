package iso20022

// Provides a set if identifications to reference to a securities settlement transaction.
type References1 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification *Max35Text `xml:"TradId,omitempty"`
}

func (r *References1) SetAccountOwnerTransactionIdentification(value string) {
	r.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (r *References1) SetAccountServicerTransactionIdentification(value string) {
	r.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (r *References1) SetMarketInfrastructureTransactionIdentification(value string) {
	r.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (r *References1) SetPoolIdentification(value string) {
	r.PoolIdentification = (*Max35Text)(&value)
}

func (r *References1) SetCommonIdentification(value string) {
	r.CommonIdentification = (*Max35Text)(&value)
}

func (r *References1) SetTradeIdentification(value string) {
	r.TradeIdentification = (*Max35Text)(&value)
}

// Reference to the transaction identifier issued by the counterparty. Building block may also be used to reference a previous transaction, or tie a set of messages together.
type References11 struct {

	// Collective reference identifying a set of messages.
	PoolReference *AdditionalReference2 `xml:"PoolRef,omitempty"`

	// Reference of the linked message that was previously sent.
	PreviousReference *AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *AdditionalReference2 `xml:"RltdRef,omitempty"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *AdditionalReference2 `xml:"CtrPtyRef,omitempty"`
}

func (r *References11) AddPoolReference() *AdditionalReference2 {
	r.PoolReference = new(AdditionalReference2)
	return r.PoolReference
}

func (r *References11) AddPreviousReference() *AdditionalReference2 {
	r.PreviousReference = new(AdditionalReference2)
	return r.PreviousReference
}

func (r *References11) AddRelatedReference() *AdditionalReference2 {
	r.RelatedReference = new(AdditionalReference2)
	return r.RelatedReference
}

func (r *References11) AddCounterpartyReference() *AdditionalReference2 {
	r.CounterpartyReference = new(AdditionalReference2)
	return r.CounterpartyReference
}

// Reference to the transaction identifier issued by the counterparty. Building block may also be used to reference a previous transaction, or tie a set of messages together.
type References15 struct {

	// Collective reference identifying a set of messages.
	PoolReference *AdditionalReference2 `xml:"PoolRef,omitempty"`

	// Reference of the linked message that was previously sent.
	PreviousReference *AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *AdditionalReference2 `xml:"RltdRef,omitempty"`
}

func (r *References15) AddPoolReference() *AdditionalReference2 {
	r.PoolReference = new(AdditionalReference2)
	return r.PoolReference
}

func (r *References15) AddPreviousReference() *AdditionalReference2 {
	r.PreviousReference = new(AdditionalReference2)
	return r.PreviousReference
}

func (r *References15) AddRelatedReference() *AdditionalReference2 {
	r.RelatedReference = new(AdditionalReference2)
	return r.RelatedReference
}

// Provides a set if identifications to reference to a securities settlement transaction.
type References16 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification *Max35Text `xml:"TradId,omitempty"`
}

func (r *References16) SetAccountOwnerTransactionIdentification(value string) {
	r.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (r *References16) SetAccountServicerTransactionIdentification(value string) {
	r.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (r *References16) SetMarketInfrastructureTransactionIdentification(value string) {
	r.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (r *References16) SetProcessorTransactionIdentification(value string) {
	r.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (r *References16) SetPoolIdentification(value string) {
	r.PoolIdentification = (*Max35Text)(&value)
}

func (r *References16) SetCommonIdentification(value string) {
	r.CommonIdentification = (*Max35Text)(&value)
}

func (r *References16) SetTradeIdentification(value string) {
	r.TradeIdentification = (*Max35Text)(&value)
}

// Provides a set if identifications to reference to a securities settlement transaction.
type References18 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification *Max35Text `xml:"TradId,omitempty"`
}

func (r *References18) SetAccountOwnerTransactionIdentification(value string) {
	r.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (r *References18) SetAccountServicerTransactionIdentification(value string) {
	r.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (r *References18) SetMarketInfrastructureTransactionIdentification(value string) {
	r.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (r *References18) SetProcessorTransactionIdentification(value string) {
	r.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (r *References18) SetPoolIdentification(value string) {
	r.PoolIdentification = (*Max35Text)(&value)
}

func (r *References18) SetCommonIdentification(value string) {
	r.CommonIdentification = (*Max35Text)(&value)
}

func (r *References18) SetTradeIdentification(value string) {
	r.TradeIdentification = (*Max35Text)(&value)
}

// Reference to the transaction identifier issued by the counterparty. Building block may also be used to reference a previous transaction, or tie a set of messages together.
type References20 struct {

	// Collective reference identifying a set of messages.
	PoolReference *AdditionalReference6 `xml:"PoolRef,omitempty"`

	// Reference of the linked message that was previously sent.
	PreviousReference *AdditionalReference6 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *AdditionalReference6 `xml:"RltdRef,omitempty"`
}

func (r *References20) AddPoolReference() *AdditionalReference6 {
	r.PoolReference = new(AdditionalReference6)
	return r.PoolReference
}

func (r *References20) AddPreviousReference() *AdditionalReference6 {
	r.PreviousReference = new(AdditionalReference6)
	return r.PreviousReference
}

func (r *References20) AddRelatedReference() *AdditionalReference6 {
	r.RelatedReference = new(AdditionalReference6)
	return r.RelatedReference
}

// Provides a set if identifications to reference to a securities settlement transaction.
type References21 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctOwnrTxId,omitempty"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *RestrictedFINXMax16Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *RestrictedFINXMax16Text `xml:"PrcrTxId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification *RestrictedFINXMax16Text `xml:"TradId,omitempty"`
}

func (r *References21) SetAccountOwnerTransactionIdentification(value string) {
	r.AccountOwnerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References21) SetAccountServicerTransactionIdentification(value string) {
	r.AccountServicerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References21) SetMarketInfrastructureTransactionIdentification(value string) {
	r.MarketInfrastructureTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References21) SetProcessorTransactionIdentification(value string) {
	r.ProcessorTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References21) SetPoolIdentification(value string) {
	r.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References21) SetCommonIdentification(value string) {
	r.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References21) SetTradeIdentification(value string) {
	r.TradeIdentification = (*RestrictedFINXMax16Text)(&value)
}

// Set of elements for the identification of the message and related references.
type References3 struct {

	// Identifies a message by a unique identifier and the date and time when the message was created by the sender.
	MessageIdentification *MessageIdentification1 `xml:"MsgId"`

	// Identification of the request message that has to be completed.
	RequestToBeCompletedIdentification *MessageIdentification1 `xml:"ReqToBeCmpltdId"`

	// Identifies a process by a unique identifier and the date and time when the first message belonging to the process was created by the sender. The process identification remains the same in all messages belonging to the same process, from the initial request message to the final account report closing the process.
	ProcessIdentification *MessageIdentification1 `xml:"PrcId"`

	// Reason of the request.
	RequestReason []*Max35Text `xml:"ReqRsn"`

	// File name of a document logically related to the request.
	AttachedDocumentName []*Max70Text `xml:"AttchdDocNm,omitempty"`
}

func (r *References3) AddMessageIdentification() *MessageIdentification1 {
	r.MessageIdentification = new(MessageIdentification1)
	return r.MessageIdentification
}

func (r *References3) AddRequestToBeCompletedIdentification() *MessageIdentification1 {
	r.RequestToBeCompletedIdentification = new(MessageIdentification1)
	return r.RequestToBeCompletedIdentification
}

func (r *References3) AddProcessIdentification() *MessageIdentification1 {
	r.ProcessIdentification = new(MessageIdentification1)
	return r.ProcessIdentification
}

func (r *References3) AddRequestReason(value string) {
	r.RequestReason = append(r.RequestReason, (*Max35Text)(&value))
}

func (r *References3) AddAttachedDocumentName(value string) {
	r.AttachedDocumentName = append(r.AttachedDocumentName, (*Max70Text)(&value))
}

// Set of elements for the identification of the message and related references.
type References4 struct {

	// Identifies a message by a unique identifier and the date and time when the message was created by the sender.
	MessageIdentification *MessageIdentification1 `xml:"MsgId"`

	// Identifies a process by a unique identifier and the date and time when the first message belonging to the process was created by the sender. The process identification remains the same in all messages belonging to the same process, from the initial request message to the final account report closing the process.
	ProcessIdentification *MessageIdentification1 `xml:"PrcId"`

	// File name of a document logically related to the request.
	AttachedDocumentName []*Max70Text `xml:"AttchdDocNm,omitempty"`
}

func (r *References4) AddMessageIdentification() *MessageIdentification1 {
	r.MessageIdentification = new(MessageIdentification1)
	return r.MessageIdentification
}

func (r *References4) AddProcessIdentification() *MessageIdentification1 {
	r.ProcessIdentification = new(MessageIdentification1)
	return r.ProcessIdentification
}

func (r *References4) AddAttachedDocumentName(value string) {
	r.AttachedDocumentName = append(r.AttachedDocumentName, (*Max70Text)(&value))
}

// Set of elements for the identification of the message and related references.
type References5 struct {

	// Identifies the type of acknowledged request.
	RequestType *UseCases1Code `xml:"ReqTp"`

	// Identifies a message by a unique identifier and the date and time when the message was created by the sender.
	MessageIdentification *MessageIdentification1 `xml:"MsgId"`

	// Identifies a process by a unique identifier and the date and time when the first message belonging to the process was created by the sender. The process identification remains the same in all messages belonging to the same process, from the initial request message to the final account report closing the process.
	ProcessIdentification *MessageIdentification1 `xml:"PrcId"`

	// Reference to the message that is acknowledged.
	AcknowledgedMessageIdentification []*MessageIdentification1 `xml:"AckdMsgId,omitempty"`

	// Status of the request.
	Status *Max35Text `xml:"Sts,omitempty"`

	// File name of a document logically related to the request.
	AttachedDocumentName []*Max70Text `xml:"AttchdDocNm,omitempty"`
}

func (r *References5) SetRequestType(value string) {
	r.RequestType = (*UseCases1Code)(&value)
}

func (r *References5) AddMessageIdentification() *MessageIdentification1 {
	r.MessageIdentification = new(MessageIdentification1)
	return r.MessageIdentification
}

func (r *References5) AddProcessIdentification() *MessageIdentification1 {
	r.ProcessIdentification = new(MessageIdentification1)
	return r.ProcessIdentification
}

func (r *References5) AddAcknowledgedMessageIdentification() *MessageIdentification1 {
	newValue := new(MessageIdentification1)
	r.AcknowledgedMessageIdentification = append(r.AcknowledgedMessageIdentification, newValue)
	return newValue
}

func (r *References5) SetStatus(value string) {
	r.Status = (*Max35Text)(&value)
}

func (r *References5) AddAttachedDocumentName(value string) {
	r.AttachedDocumentName = append(r.AttachedDocumentName, (*Max70Text)(&value))
}

// Set of elements for the identification of the message and related references.
type References6 struct {

	// Identify the type of rejected request.
	RejectedRequestType *UseCases1Code `xml:"RjctdReqTp"`

	// Reason of the message rejection.
	RejectionReason []*Max350Text `xml:"RjctnRsn"`

	// Identification of the rejected request message.
	RejectedRequestIdentification *MessageIdentification1 `xml:"RjctdReqId"`

	// Identifies a message by a unique identifier and the date and time when the message was created by the sender.
	MessageIdentification *MessageIdentification1 `xml:"MsgId"`

	// Identifies a process by a unique identifier and the date and time when the first message belonging to the process was created by the sender. The process identification remains the same in all messages belonging to the same process, from the initial request message to the final account report closing the process.
	ProcessIdentification *MessageIdentification1 `xml:"PrcId"`

	// File name of a document logically related to the request.
	AttachedDocumentName []*Max70Text `xml:"AttchdDocNm,omitempty"`
}

func (r *References6) SetRejectedRequestType(value string) {
	r.RejectedRequestType = (*UseCases1Code)(&value)
}

func (r *References6) AddRejectionReason(value string) {
	r.RejectionReason = append(r.RejectionReason, (*Max350Text)(&value))
}

func (r *References6) AddRejectedRequestIdentification() *MessageIdentification1 {
	r.RejectedRequestIdentification = new(MessageIdentification1)
	return r.RejectedRequestIdentification
}

func (r *References6) AddMessageIdentification() *MessageIdentification1 {
	r.MessageIdentification = new(MessageIdentification1)
	return r.MessageIdentification
}

func (r *References6) AddProcessIdentification() *MessageIdentification1 {
	r.ProcessIdentification = new(MessageIdentification1)
	return r.ProcessIdentification
}

func (r *References6) AddAttachedDocumentName(value string) {
	r.AttachedDocumentName = append(r.AttachedDocumentName, (*Max70Text)(&value))
}

// Provides a set if identifications to reference to a securities settlement transaction.
type References7 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification *Max35Text `xml:"TradId,omitempty"`
}

func (r *References7) SetAccountOwnerTransactionIdentification(value string) {
	r.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (r *References7) SetAccountServicerTransactionIdentification(value string) {
	r.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (r *References7) SetMarketInfrastructureTransactionIdentification(value string) {
	r.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (r *References7) SetProcessorTransactionIdentification(value string) {
	r.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (r *References7) SetPoolIdentification(value string) {
	r.PoolIdentification = (*Max35Text)(&value)
}

func (r *References7) SetCommonIdentification(value string) {
	r.CommonIdentification = (*Max35Text)(&value)
}

func (r *References7) SetTradeIdentification(value string) {
	r.TradeIdentification = (*Max35Text)(&value)
}

// Provides a set if identifications to reference to a securities settlement transaction.
type References9 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification *Max35Text `xml:"TradId,omitempty"`
}

func (r *References9) SetAccountOwnerTransactionIdentification(value string) {
	r.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (r *References9) SetAccountServicerTransactionIdentification(value string) {
	r.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (r *References9) SetMarketInfrastructureTransactionIdentification(value string) {
	r.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (r *References9) SetProcessorTransactionIdentification(value string) {
	r.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (r *References9) SetPoolIdentification(value string) {
	r.PoolIdentification = (*Max35Text)(&value)
}

func (r *References9) SetCommonIdentification(value string) {
	r.CommonIdentification = (*Max35Text)(&value)
}

func (r *References9) SetTradeIdentification(value string) {
	r.TradeIdentification = (*Max35Text)(&value)
}
