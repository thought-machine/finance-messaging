package iso20022

// Identifies the details of the transaction.
type Transaction12 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securitiess.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails25 `xml:"TxDtls,omitempty"`

	// Status and reason for the transaction.
	StatusAndReason []*StatusAndReason2 `xml:"StsAndRsn,omitempty"`
}

func (t *Transaction12) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction12) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction12) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction12) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction12) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *Transaction12) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *Transaction12) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *Transaction12) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *Transaction12) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction12) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction12) AddTransactionDetails() *TransactionDetails25 {
	t.TransactionDetails = new(TransactionDetails25)
	return t.TransactionDetails
}

func (t *Transaction12) AddStatusAndReason() *StatusAndReason2 {
	newValue := new(StatusAndReason2)
	t.StatusAndReason = append(t.StatusAndReason, newValue)
	return newValue
}

// Identifies the details of the transaction.
type Transaction13 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails24 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (t *Transaction13) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction13) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction13) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction13) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction13) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *Transaction13) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *Transaction13) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *Transaction13) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *Transaction13) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction13) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction13) AddTransactionDetails() *TransactionDetails24 {
	t.TransactionDetails = new(TransactionDetails24)
	return t.TransactionDetails
}

func (t *Transaction13) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	t.SupplementaryData = append(t.SupplementaryData, newValue)
	return newValue
}

// Identifies the details of the transaction.
type Transaction14 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails26 `xml:"TxDtls,omitempty"`
}

func (t *Transaction14) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction14) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction14) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction14) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction14) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *Transaction14) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *Transaction14) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *Transaction14) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *Transaction14) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction14) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction14) AddTransactionDetails() *TransactionDetails26 {
	t.TransactionDetails = new(TransactionDetails26)
	return t.TransactionDetails
}

// Identifies the details of the transaction.
type Transaction18 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails37 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (t *Transaction18) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction18) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction18) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction18) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction18) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *Transaction18) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *Transaction18) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *Transaction18) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *Transaction18) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction18) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction18) AddTransactionDetails() *TransactionDetails37 {
	t.TransactionDetails = new(TransactionDetails37)
	return t.TransactionDetails
}

func (t *Transaction18) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	t.SupplementaryData = append(t.SupplementaryData, newValue)
	return newValue
}

// Identifies the details of the transaction.
type Transaction19 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails39 `xml:"TxDtls,omitempty"`

	// Status and reason for the transaction.
	StatusAndReason []*Status9Choice `xml:"StsAndRsn,omitempty"`
}

func (t *Transaction19) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction19) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction19) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction19) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction19) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *Transaction19) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *Transaction19) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *Transaction19) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *Transaction19) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction19) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction19) AddTransactionDetails() *TransactionDetails39 {
	t.TransactionDetails = new(TransactionDetails39)
	return t.TransactionDetails
}

func (t *Transaction19) AddStatusAndReason() *Status9Choice {
	newValue := new(Status9Choice)
	t.StatusAndReason = append(t.StatusAndReason, newValue)
	return newValue
}

// Identifies the details of the transaction.
type Transaction20 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails39 `xml:"TxDtls,omitempty"`
}

func (t *Transaction20) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction20) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction20) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction20) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction20) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *Transaction20) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *Transaction20) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *Transaction20) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *Transaction20) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction20) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction20) AddTransactionDetails() *TransactionDetails39 {
	t.TransactionDetails = new(TransactionDetails39)
	return t.TransactionDetails
}

// Identifies the details of the transaction.
type Transaction27 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

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

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails56 `xml:"TxDtls,omitempty"`

	// Status and reason for the transaction.
	StatusAndReason []*Status9Choice `xml:"StsAndRsn,omitempty"`
}

func (t *Transaction27) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction27) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction27) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction27) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction27) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *Transaction27) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *Transaction27) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *Transaction27) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *Transaction27) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction27) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction27) SetClientCollateralInstructionIdentification(value string) {
	t.ClientCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *Transaction27) SetTripartyCollateralInstructionIdentification(value string) {
	t.TripartyCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *Transaction27) AddTransactionDetails() *TransactionDetails56 {
	t.TransactionDetails = new(TransactionDetails56)
	return t.TransactionDetails
}

func (t *Transaction27) AddStatusAndReason() *Status9Choice {
	newValue := new(Status9Choice)
	t.StatusAndReason = append(t.StatusAndReason, newValue)
	return newValue
}

// Identifies the details of the transaction.
type Transaction28 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

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

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails56 `xml:"TxDtls,omitempty"`
}

func (t *Transaction28) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction28) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction28) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction28) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction28) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *Transaction28) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *Transaction28) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *Transaction28) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *Transaction28) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction28) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction28) SetClientCollateralInstructionIdentification(value string) {
	t.ClientCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *Transaction28) SetTripartyCollateralInstructionIdentification(value string) {
	t.TripartyCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *Transaction28) AddTransactionDetails() *TransactionDetails56 {
	t.TransactionDetails = new(TransactionDetails56)
	return t.TransactionDetails
}

// Identifies the details of the transaction.
type Transaction29 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

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

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails57 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (t *Transaction29) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction29) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction29) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction29) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction29) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *Transaction29) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *Transaction29) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *Transaction29) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *Transaction29) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction29) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction29) SetClientCollateralInstructionIdentification(value string) {
	t.ClientCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *Transaction29) SetTripartyCollateralInstructionIdentification(value string) {
	t.TripartyCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *Transaction29) AddTransactionDetails() *TransactionDetails57 {
	t.TransactionDetails = new(TransactionDetails57)
	return t.TransactionDetails
}

func (t *Transaction29) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	t.SupplementaryData = append(t.SupplementaryData, newValue)
	return newValue
}

// Identifies the details of the transaction.
type Transaction34 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

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

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails62 `xml:"TxDtls,omitempty"`

	// Status and reason for the transaction.
	StatusAndReason []*Status15Choice `xml:"StsAndRsn,omitempty"`
}

func (t *Transaction34) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction34) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction34) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction34) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction34) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *Transaction34) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *Transaction34) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *Transaction34) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *Transaction34) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction34) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction34) SetClientCollateralInstructionIdentification(value string) {
	t.ClientCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *Transaction34) SetTripartyCollateralInstructionIdentification(value string) {
	t.TripartyCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *Transaction34) AddTransactionDetails() *TransactionDetails62 {
	t.TransactionDetails = new(TransactionDetails62)
	return t.TransactionDetails
}

func (t *Transaction34) AddStatusAndReason() *Status15Choice {
	newValue := new(Status15Choice)
	t.StatusAndReason = append(t.StatusAndReason, newValue)
	return newValue
}

// Identifies the details of the transaction.
type Transaction35 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

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

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails62 `xml:"TxDtls,omitempty"`
}

func (t *Transaction35) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction35) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction35) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction35) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction35) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *Transaction35) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *Transaction35) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *Transaction35) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *Transaction35) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction35) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction35) SetClientCollateralInstructionIdentification(value string) {
	t.ClientCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *Transaction35) SetTripartyCollateralInstructionIdentification(value string) {
	t.TripartyCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *Transaction35) AddTransactionDetails() *TransactionDetails62 {
	t.TransactionDetails = new(TransactionDetails62)
	return t.TransactionDetails
}

// Identifies the details of the transaction.
type Transaction36 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

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

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails63 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (t *Transaction36) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction36) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction36) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction36) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction36) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *Transaction36) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *Transaction36) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *Transaction36) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *Transaction36) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction36) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction36) SetClientCollateralInstructionIdentification(value string) {
	t.ClientCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *Transaction36) SetTripartyCollateralInstructionIdentification(value string) {
	t.TripartyCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *Transaction36) AddTransactionDetails() *TransactionDetails63 {
	t.TransactionDetails = new(TransactionDetails63)
	return t.TransactionDetails
}

func (t *Transaction36) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	t.SupplementaryData = append(t.SupplementaryData, newValue)
	return newValue
}

// Identifies the details of the transaction.
type Transaction40 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

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

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails69 `xml:"TxDtls,omitempty"`
}

func (t *Transaction40) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction40) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction40) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction40) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction40) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *Transaction40) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *Transaction40) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *Transaction40) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *Transaction40) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction40) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction40) SetClientCollateralInstructionIdentification(value string) {
	t.ClientCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *Transaction40) SetTripartyCollateralInstructionIdentification(value string) {
	t.TripartyCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *Transaction40) AddTransactionDetails() *TransactionDetails69 {
	t.TransactionDetails = new(TransactionDetails69)
	return t.TransactionDetails
}

// Identifies the details of the transaction.
type Transaction41 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

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

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails69 `xml:"TxDtls,omitempty"`

	// Status and reason for the transaction.
	StatusAndReason []*Status15Choice `xml:"StsAndRsn,omitempty"`
}

func (t *Transaction41) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction41) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction41) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction41) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction41) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *Transaction41) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *Transaction41) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *Transaction41) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *Transaction41) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction41) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction41) SetClientCollateralInstructionIdentification(value string) {
	t.ClientCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *Transaction41) SetTripartyCollateralInstructionIdentification(value string) {
	t.TripartyCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *Transaction41) AddTransactionDetails() *TransactionDetails69 {
	t.TransactionDetails = new(TransactionDetails69)
	return t.TransactionDetails
}

func (t *Transaction41) AddStatusAndReason() *Status15Choice {
	newValue := new(Status15Choice)
	t.StatusAndReason = append(t.StatusAndReason, newValue)
	return newValue
}

// Identifies the details of the transaction.
type Transaction45 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

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

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails75 `xml:"TxDtls,omitempty"`
}

func (t *Transaction45) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction45) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction45) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction45) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction45) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *Transaction45) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *Transaction45) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *Transaction45) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *Transaction45) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction45) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction45) SetClientCollateralInstructionIdentification(value string) {
	t.ClientCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *Transaction45) SetTripartyCollateralInstructionIdentification(value string) {
	t.TripartyCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *Transaction45) AddTransactionDetails() *TransactionDetails75 {
	t.TransactionDetails = new(TransactionDetails75)
	return t.TransactionDetails
}

// Identifies the details of the transaction.
type Transaction46 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

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

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails78 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (t *Transaction46) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction46) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction46) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction46) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction46) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *Transaction46) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *Transaction46) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *Transaction46) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *Transaction46) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction46) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction46) SetClientCollateralInstructionIdentification(value string) {
	t.ClientCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *Transaction46) SetTripartyCollateralInstructionIdentification(value string) {
	t.TripartyCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *Transaction46) AddTransactionDetails() *TransactionDetails78 {
	t.TransactionDetails = new(TransactionDetails78)
	return t.TransactionDetails
}

func (t *Transaction46) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	t.SupplementaryData = append(t.SupplementaryData, newValue)
	return newValue
}

// Identifies the details of the transaction.
type Transaction47 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

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

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails75 `xml:"TxDtls,omitempty"`

	// Status and reason for the transaction.
	StatusAndReason []*Status18Choice `xml:"StsAndRsn,omitempty"`
}

func (t *Transaction47) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction47) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction47) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction47) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction47) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *Transaction47) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *Transaction47) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *Transaction47) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *Transaction47) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction47) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction47) SetClientCollateralInstructionIdentification(value string) {
	t.ClientCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *Transaction47) SetTripartyCollateralInstructionIdentification(value string) {
	t.TripartyCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *Transaction47) AddTransactionDetails() *TransactionDetails75 {
	t.TransactionDetails = new(TransactionDetails75)
	return t.TransactionDetails
}

func (t *Transaction47) AddStatusAndReason() *Status18Choice {
	newValue := new(Status18Choice)
	t.StatusAndReason = append(t.StatusAndReason, newValue)
	return newValue
}

// Identifies the details of the transaction.
type Transaction48 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *RestrictedFINXMax16Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *RestrictedFINXMax16Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*RestrictedFINXMax16Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

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

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails90 `xml:"TxDtls,omitempty"`
}

func (t *Transaction48) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction48) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction48) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction48) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction48) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (t *Transaction48) SetPoolIdentification(value string) {
	t.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction48) SetCommonIdentification(value string) {
	t.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction48) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction48) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction48) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction48) SetClientCollateralInstructionIdentification(value string) {
	t.ClientCollateralInstructionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction48) SetTripartyCollateralInstructionIdentification(value string) {
	t.TripartyCollateralInstructionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction48) AddTransactionDetails() *TransactionDetails90 {
	t.TransactionDetails = new(TransactionDetails90)
	return t.TransactionDetails
}

// Identifies the details of the transaction.
type Transaction49 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *RestrictedFINXMax16Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *RestrictedFINXMax16Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*RestrictedFINXMax16Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

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

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails90 `xml:"TxDtls,omitempty"`

	// Status and reason for the transaction.
	StatusAndReason []*Status23Choice `xml:"StsAndRsn,omitempty"`
}

func (t *Transaction49) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction49) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction49) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction49) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction49) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (t *Transaction49) SetPoolIdentification(value string) {
	t.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction49) SetCommonIdentification(value string) {
	t.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction49) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction49) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction49) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction49) SetClientCollateralInstructionIdentification(value string) {
	t.ClientCollateralInstructionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction49) SetTripartyCollateralInstructionIdentification(value string) {
	t.TripartyCollateralInstructionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction49) AddTransactionDetails() *TransactionDetails90 {
	t.TransactionDetails = new(TransactionDetails90)
	return t.TransactionDetails
}

func (t *Transaction49) AddStatusAndReason() *Status23Choice {
	newValue := new(Status23Choice)
	t.StatusAndReason = append(t.StatusAndReason, newValue)
	return newValue
}

// Identifies the details of the transaction.
type Transaction50 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *RestrictedFINXMax16Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *RestrictedFINXMax16Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*RestrictedFINXMax16Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

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

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails91 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (t *Transaction50) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction50) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction50) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction50) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction50) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (t *Transaction50) SetPoolIdentification(value string) {
	t.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction50) SetCommonIdentification(value string) {
	t.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction50) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction50) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction50) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction50) SetClientCollateralInstructionIdentification(value string) {
	t.ClientCollateralInstructionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction50) SetTripartyCollateralInstructionIdentification(value string) {
	t.TripartyCollateralInstructionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction50) AddTransactionDetails() *TransactionDetails91 {
	t.TransactionDetails = new(TransactionDetails91)
	return t.TransactionDetails
}

func (t *Transaction50) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	t.SupplementaryData = append(t.SupplementaryData, newValue)
	return newValue
}

// Specifies the details of the transaction.
type Transaction52 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

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

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails95 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (t *Transaction52) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction52) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction52) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction52) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction52) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *Transaction52) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *Transaction52) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *Transaction52) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *Transaction52) SetTripartyAgentServiceProviderCollateralTransactionIdentification(value string) {
	t.TripartyAgentServiceProviderCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction52) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction52) SetClientCollateralInstructionIdentification(value string) {
	t.ClientCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *Transaction52) SetTripartyAgentServiceProviderCollateralInstructionIdentification(value string) {
	t.TripartyAgentServiceProviderCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *Transaction52) AddTransactionDetails() *TransactionDetails95 {
	t.TransactionDetails = new(TransactionDetails95)
	return t.TransactionDetails
}

func (t *Transaction52) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	t.SupplementaryData = append(t.SupplementaryData, newValue)
	return newValue
}

// Specifies the details of the transaction.
type Transaction53 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

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

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails96 `xml:"TxDtls,omitempty"`

	// Status and reason for the transaction.
	StatusAndReason []*Status18Choice `xml:"StsAndRsn,omitempty"`
}

func (t *Transaction53) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction53) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction53) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction53) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction53) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *Transaction53) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *Transaction53) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *Transaction53) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *Transaction53) SetTripartyAgentServiceProviderCollateralTransactionIdentification(value string) {
	t.TripartyAgentServiceProviderCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction53) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction53) SetClientCollateralInstructionIdentification(value string) {
	t.ClientCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *Transaction53) SetTripartyAgentServiceProviderCollateralInstructionIdentification(value string) {
	t.TripartyAgentServiceProviderCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *Transaction53) AddTransactionDetails() *TransactionDetails96 {
	t.TransactionDetails = new(TransactionDetails96)
	return t.TransactionDetails
}

func (t *Transaction53) AddStatusAndReason() *Status18Choice {
	newValue := new(Status18Choice)
	t.StatusAndReason = append(t.StatusAndReason, newValue)
	return newValue
}

// Specifies the details of the transaction.
type Transaction54 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

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

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails96 `xml:"TxDtls,omitempty"`
}

func (t *Transaction54) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction54) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction54) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction54) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction54) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *Transaction54) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *Transaction54) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *Transaction54) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *Transaction54) SetTripartyAgentServiceProviderCollateralTransactionIdentification(value string) {
	t.TripartyAgentServiceProviderCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction54) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction54) SetClientCollateralInstructionIdentification(value string) {
	t.ClientCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *Transaction54) SetTripartyAgentServiceProviderCollateralInstructionIdentification(value string) {
	t.TripartyAgentServiceProviderCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *Transaction54) AddTransactionDetails() *TransactionDetails96 {
	t.TransactionDetails = new(TransactionDetails96)
	return t.TransactionDetails
}

// Specifies the details of the transaction.
type Transaction55 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *RestrictedFINXMax16Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *RestrictedFINXMax16Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*RestrictedFINXMax16Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

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

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails98 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (t *Transaction55) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction55) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction55) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction55) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction55) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (t *Transaction55) SetPoolIdentification(value string) {
	t.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction55) SetCommonIdentification(value string) {
	t.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction55) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction55) SetTripartyAgentServiceProviderCollateralTransactionIdentification(value string) {
	t.TripartyAgentServiceProviderCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction55) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction55) SetClientCollateralInstructionIdentification(value string) {
	t.ClientCollateralInstructionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction55) SetTripartyAgentServiceProviderCollateralInstructionIdentification(value string) {
	t.TripartyAgentServiceProviderCollateralInstructionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction55) AddTransactionDetails() *TransactionDetails98 {
	t.TransactionDetails = new(TransactionDetails98)
	return t.TransactionDetails
}

func (t *Transaction55) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	t.SupplementaryData = append(t.SupplementaryData, newValue)
	return newValue
}

// Specifies the details of the transaction.
type Transaction56 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *RestrictedFINXMax16Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *RestrictedFINXMax16Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*RestrictedFINXMax16Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

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

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails99 `xml:"TxDtls,omitempty"`
}

func (t *Transaction56) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction56) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction56) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction56) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction56) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (t *Transaction56) SetPoolIdentification(value string) {
	t.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction56) SetCommonIdentification(value string) {
	t.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction56) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction56) SetTripartyAgentServiceProviderCollateralTransactionIdentification(value string) {
	t.TripartyAgentServiceProviderCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction56) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction56) SetClientCollateralInstructionIdentification(value string) {
	t.ClientCollateralInstructionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction56) SetTripartyAgentServiceProviderCollateralInstructionIdentification(value string) {
	t.TripartyAgentServiceProviderCollateralInstructionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction56) AddTransactionDetails() *TransactionDetails99 {
	t.TransactionDetails = new(TransactionDetails99)
	return t.TransactionDetails
}

// Specifies the details of the transaction.
type Transaction57 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *RestrictedFINXMax16Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *RestrictedFINXMax16Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*RestrictedFINXMax16Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

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

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails99 `xml:"TxDtls,omitempty"`

	// Status and reason for the transaction.
	StatusAndReason []*Status23Choice `xml:"StsAndRsn,omitempty"`
}

func (t *Transaction57) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction57) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction57) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction57) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction57) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (t *Transaction57) SetPoolIdentification(value string) {
	t.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction57) SetCommonIdentification(value string) {
	t.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction57) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction57) SetTripartyAgentServiceProviderCollateralTransactionIdentification(value string) {
	t.TripartyAgentServiceProviderCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction57) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction57) SetClientCollateralInstructionIdentification(value string) {
	t.ClientCollateralInstructionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction57) SetTripartyAgentServiceProviderCollateralInstructionIdentification(value string) {
	t.TripartyAgentServiceProviderCollateralInstructionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction57) AddTransactionDetails() *TransactionDetails99 {
	t.TransactionDetails = new(TransactionDetails99)
	return t.TransactionDetails
}

func (t *Transaction57) AddStatusAndReason() *Status23Choice {
	newValue := new(Status23Choice)
	t.StatusAndReason = append(t.StatusAndReason, newValue)
	return newValue
}

// Identifies the details of the transaction.
type Transaction6 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails6 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension2 `xml:"Xtnsn,omitempty"`
}

func (t *Transaction6) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction6) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction6) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction6) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *Transaction6) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *Transaction6) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *Transaction6) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *Transaction6) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction6) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction6) AddTransactionDetails() *TransactionDetails6 {
	t.TransactionDetails = new(TransactionDetails6)
	return t.TransactionDetails
}

func (t *Transaction6) AddExtension() *Extension2 {
	newValue := new(Extension2)
	t.Extension = append(t.Extension, newValue)
	return newValue
}

// Identifies the details of the transaction.
type Transaction7 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails7 `xml:"TxDtls,omitempty"`
}

func (t *Transaction7) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction7) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction7) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction7) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *Transaction7) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *Transaction7) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *Transaction7) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *Transaction7) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction7) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction7) AddTransactionDetails() *TransactionDetails7 {
	t.TransactionDetails = new(TransactionDetails7)
	return t.TransactionDetails
}

// Identifies the details of the transaction.
type Transaction8 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails8 `xml:"TxDtls,omitempty"`

	// Status and reason for the transaction.
	StatusAndReason []*StatusAndReason2 `xml:"StsAndRsn,omitempty"`
}

func (t *Transaction8) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction8) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction8) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction8) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *Transaction8) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *Transaction8) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *Transaction8) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *Transaction8) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction8) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *Transaction8) AddTransactionDetails() *TransactionDetails8 {
	t.TransactionDetails = new(TransactionDetails8)
	return t.TransactionDetails
}

func (t *Transaction8) AddStatusAndReason() *StatusAndReason2 {
	newValue := new(StatusAndReason2)
	t.StatusAndReason = append(t.StatusAndReason, newValue)
	return newValue
}
