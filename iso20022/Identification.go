package iso20022

// Unique identifier of a document, message or transaction.
type Identification1 struct {

	// Unique identifier of a document, message or transaction.
	Identification *Max35Text `xml:"Id"`
}

func (i *Identification1) SetIdentification(value string) {
	i.Identification = (*Max35Text)(&value)
}

// Unique identifier of a document, message or transaction.
type Identification11 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterIdentification *Max35Text `xml:"MstrId,omitempty"`

	// Identification of a basket trade.
	BasketIdentification *Max35Text `xml:"BsktId,omitempty"`

	// Reference identifying a index trade.
	IndexIdentification *Max35Text `xml:"IndxId,omitempty"`

	// Unique identifier for a list, as assigned by the trading party. The identifier must be unique within a single trading day.
	ListIdentification *Max35Text `xml:"ListId,omitempty"`

	// Program reference which identifies a program trade.
	ProgramIdentification *Max35Text `xml:"PrgmId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`
}

func (i *Identification11) SetAccountOwnerTransactionIdentification(value string) {
	i.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (i *Identification11) SetAccountServicerTransactionIdentification(value string) {
	i.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (i *Identification11) SetMarketInfrastructureTransactionIdentification(value string) {
	i.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (i *Identification11) SetProcessorTransactionIdentification(value string) {
	i.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (i *Identification11) SetCommonIdentification(value string) {
	i.CommonIdentification = (*Max35Text)(&value)
}

func (i *Identification11) AddTradeIdentification(value string) {
	i.TradeIdentification = append(i.TradeIdentification, (*Max35Text)(&value))
}

func (i *Identification11) SetMasterIdentification(value string) {
	i.MasterIdentification = (*Max35Text)(&value)
}

func (i *Identification11) SetBasketIdentification(value string) {
	i.BasketIdentification = (*Max35Text)(&value)
}

func (i *Identification11) SetIndexIdentification(value string) {
	i.IndexIdentification = (*Max35Text)(&value)
}

func (i *Identification11) SetListIdentification(value string) {
	i.ListIdentification = (*Max35Text)(&value)
}

func (i *Identification11) SetProgramIdentification(value string) {
	i.ProgramIdentification = (*Max35Text)(&value)
}

func (i *Identification11) SetPoolIdentification(value string) {
	i.PoolIdentification = (*Max35Text)(&value)
}

func (i *Identification11) SetCorporateActionEventIdentification(value string) {
	i.CorporateActionEventIdentification = (*Max35Text)(&value)
}

// Unique identifier of a document, message or transaction.
type Identification14 struct {

	// Unique identifier of a document, message or transaction.
	Identification *Max35Text `xml:"Id"`
}

func (i *Identification14) SetIdentification(value string) {
	i.Identification = (*Max35Text)(&value)
}

// Unique identifier of a document, message or transaction.
type Identification15 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterIdentification *Max35Text `xml:"MstrId,omitempty"`

	// Identification of a basket trade.
	BasketIdentification *Max35Text `xml:"BsktId,omitempty"`

	// Reference identifying a index trade.
	IndexIdentification *Max35Text `xml:"IndxId,omitempty"`

	// Unique identifier for a list, as assigned by the trading party. The identifier must be unique within a single trading day.
	ListIdentification *Max35Text `xml:"ListId,omitempty"`

	// Program reference which identifies a program trade.
	ProgramIdentification *Max35Text `xml:"PrgmId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`
}

func (i *Identification15) SetAccountOwnerTransactionIdentification(value string) {
	i.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (i *Identification15) SetAccountServicerTransactionIdentification(value string) {
	i.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (i *Identification15) SetMarketInfrastructureTransactionIdentification(value string) {
	i.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (i *Identification15) SetProcessorTransactionIdentification(value string) {
	i.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (i *Identification15) SetCommonIdentification(value string) {
	i.CommonIdentification = (*Max35Text)(&value)
}

func (i *Identification15) AddTradeIdentification(value string) {
	i.TradeIdentification = append(i.TradeIdentification, (*Max35Text)(&value))
}

func (i *Identification15) SetMasterIdentification(value string) {
	i.MasterIdentification = (*Max35Text)(&value)
}

func (i *Identification15) SetBasketIdentification(value string) {
	i.BasketIdentification = (*Max35Text)(&value)
}

func (i *Identification15) SetIndexIdentification(value string) {
	i.IndexIdentification = (*Max35Text)(&value)
}

func (i *Identification15) SetListIdentification(value string) {
	i.ListIdentification = (*Max35Text)(&value)
}

func (i *Identification15) SetProgramIdentification(value string) {
	i.ProgramIdentification = (*Max35Text)(&value)
}

func (i *Identification15) SetPoolIdentification(value string) {
	i.PoolIdentification = (*Max35Text)(&value)
}

func (i *Identification15) SetCorporateActionEventIdentification(value string) {
	i.CorporateActionEventIdentification = (*Max35Text)(&value)
}

// Unique identifier of a document, message or transaction.
type Identification16 struct {

	// Unique identifier of a document, message or transaction.
	Identification *RestrictedFINXMax16Text `xml:"Id"`
}

func (i *Identification16) SetIdentification(value string) {
	i.Identification = (*RestrictedFINXMax16Text)(&value)
}

// Unique identifier of a document, message or transaction.
type Identification2 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterIdentification *Max35Text `xml:"MstrId,omitempty"`

	// Identification of a basket trade.
	BasketIdentification *Max35Text `xml:"BsktId,omitempty"`

	// Reference identifying a index trade.
	IndexIdentification *Max35Text `xml:"IndxId,omitempty"`

	// Unique identifier for a list, as assigned by the trading party. The identifier must be unique within a single trading day.
	ListIdentification *Max35Text `xml:"ListId,omitempty"`

	// Program reference which identifies a program trade.
	ProgramIdentification *Max35Text `xml:"PrgmId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`
}

func (i *Identification2) SetAccountOwnerTransactionIdentification(value string) {
	i.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (i *Identification2) SetAccountServicerTransactionIdentification(value string) {
	i.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (i *Identification2) SetMarketInfrastructureTransactionIdentification(value string) {
	i.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (i *Identification2) SetCommonIdentification(value string) {
	i.CommonIdentification = (*Max35Text)(&value)
}

func (i *Identification2) AddTradeIdentification(value string) {
	i.TradeIdentification = append(i.TradeIdentification, (*Max35Text)(&value))
}

func (i *Identification2) SetMasterIdentification(value string) {
	i.MasterIdentification = (*Max35Text)(&value)
}

func (i *Identification2) SetBasketIdentification(value string) {
	i.BasketIdentification = (*Max35Text)(&value)
}

func (i *Identification2) SetIndexIdentification(value string) {
	i.IndexIdentification = (*Max35Text)(&value)
}

func (i *Identification2) SetListIdentification(value string) {
	i.ListIdentification = (*Max35Text)(&value)
}

func (i *Identification2) SetProgramIdentification(value string) {
	i.ProgramIdentification = (*Max35Text)(&value)
}

func (i *Identification2) SetPoolIdentification(value string) {
	i.PoolIdentification = (*Max35Text)(&value)
}

func (i *Identification2) SetCorporateActionEventIdentification(value string) {
	i.CorporateActionEventIdentification = (*Max35Text)(&value)
}

// Unique identifier of a document, message or transaction.
type Identification24 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *RestrictedFINXMax16Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *RestrictedFINXMax16Text `xml:"PrcrTxId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*RestrictedFINXMax16Text `xml:"TradId,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterIdentification *RestrictedFINXMax16Text `xml:"MstrId,omitempty"`

	// Identification of a basket trade.
	BasketIdentification *RestrictedFINXMax16Text `xml:"BsktId,omitempty"`

	// Reference identifying a index trade.
	IndexIdentification *RestrictedFINXMax16Text `xml:"IndxId,omitempty"`

	// Unique identifier for a list, as assigned by the trading party. The identifier must be unique within a single trading day.
	ListIdentification *RestrictedFINXMax16Text `xml:"ListId,omitempty"`

	// Program reference which identifies a program trade.
	ProgramIdentification *RestrictedFINXMax16Text `xml:"PrgmId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId,omitempty"`
}

func (i *Identification24) SetAccountOwnerTransactionIdentification(value string) {
	i.AccountOwnerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (i *Identification24) SetAccountServicerTransactionIdentification(value string) {
	i.AccountServicerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (i *Identification24) SetMarketInfrastructureTransactionIdentification(value string) {
	i.MarketInfrastructureTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (i *Identification24) SetProcessorTransactionIdentification(value string) {
	i.ProcessorTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (i *Identification24) SetCommonIdentification(value string) {
	i.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (i *Identification24) AddTradeIdentification(value string) {
	i.TradeIdentification = append(i.TradeIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (i *Identification24) SetMasterIdentification(value string) {
	i.MasterIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (i *Identification24) SetBasketIdentification(value string) {
	i.BasketIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (i *Identification24) SetIndexIdentification(value string) {
	i.IndexIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (i *Identification24) SetListIdentification(value string) {
	i.ListIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (i *Identification24) SetProgramIdentification(value string) {
	i.ProgramIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (i *Identification24) SetPoolIdentification(value string) {
	i.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (i *Identification24) SetCorporateActionEventIdentification(value string) {
	i.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

// Unique identifier of a document, message or transaction.
type Identification6 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterIdentification *Max35Text `xml:"MstrId,omitempty"`

	// Identification of a basket trade.
	BasketIdentification *Max35Text `xml:"BsktId,omitempty"`

	// Reference identifying a index trade.
	IndexIdentification *Max35Text `xml:"IndxId,omitempty"`

	// Unique identifier for a list, as assigned by the trading party. The identifier must be unique within a single trading day.
	ListIdentification *Max35Text `xml:"ListId,omitempty"`

	// Program reference which identifies a program trade.
	ProgramIdentification *Max35Text `xml:"PrgmId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`
}

func (i *Identification6) SetAccountOwnerTransactionIdentification(value string) {
	i.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (i *Identification6) SetAccountServicerTransactionIdentification(value string) {
	i.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (i *Identification6) SetMarketInfrastructureTransactionIdentification(value string) {
	i.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (i *Identification6) SetProcessorTransactionIdentification(value string) {
	i.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (i *Identification6) SetCommonIdentification(value string) {
	i.CommonIdentification = (*Max35Text)(&value)
}

func (i *Identification6) AddTradeIdentification(value string) {
	i.TradeIdentification = append(i.TradeIdentification, (*Max35Text)(&value))
}

func (i *Identification6) SetMasterIdentification(value string) {
	i.MasterIdentification = (*Max35Text)(&value)
}

func (i *Identification6) SetBasketIdentification(value string) {
	i.BasketIdentification = (*Max35Text)(&value)
}

func (i *Identification6) SetIndexIdentification(value string) {
	i.IndexIdentification = (*Max35Text)(&value)
}

func (i *Identification6) SetListIdentification(value string) {
	i.ListIdentification = (*Max35Text)(&value)
}

func (i *Identification6) SetProgramIdentification(value string) {
	i.ProgramIdentification = (*Max35Text)(&value)
}

func (i *Identification6) SetPoolIdentification(value string) {
	i.PoolIdentification = (*Max35Text)(&value)
}

func (i *Identification6) SetCorporateActionEventIdentification(value string) {
	i.CorporateActionEventIdentification = (*Max35Text)(&value)
}

// Unique identifier of a document, message or transaction.
type Identification7 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterIdentification *Max35Text `xml:"MstrId,omitempty"`

	// Identification of a basket trade.
	BasketIdentification *Max35Text `xml:"BsktId,omitempty"`

	// Reference identifying a index trade.
	IndexIdentification *Max35Text `xml:"IndxId,omitempty"`

	// Unique identifier for a list, as assigned by the trading party. The identifier must be unique within a single trading day.
	ListIdentification *Max35Text `xml:"ListId,omitempty"`

	// Program reference which identifies a program trade.
	ProgramIdentification *Max35Text `xml:"PrgmId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`
}

func (i *Identification7) SetAccountOwnerTransactionIdentification(value string) {
	i.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (i *Identification7) SetAccountServicerTransactionIdentification(value string) {
	i.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (i *Identification7) SetMarketInfrastructureTransactionIdentification(value string) {
	i.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (i *Identification7) SetProcessorTransactionIdentification(value string) {
	i.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (i *Identification7) SetCommonIdentification(value string) {
	i.CommonIdentification = (*Max35Text)(&value)
}

func (i *Identification7) AddTradeIdentification(value string) {
	i.TradeIdentification = append(i.TradeIdentification, (*Max35Text)(&value))
}

func (i *Identification7) SetMasterIdentification(value string) {
	i.MasterIdentification = (*Max35Text)(&value)
}

func (i *Identification7) SetBasketIdentification(value string) {
	i.BasketIdentification = (*Max35Text)(&value)
}

func (i *Identification7) SetIndexIdentification(value string) {
	i.IndexIdentification = (*Max35Text)(&value)
}

func (i *Identification7) SetListIdentification(value string) {
	i.ListIdentification = (*Max35Text)(&value)
}

func (i *Identification7) SetProgramIdentification(value string) {
	i.ProgramIdentification = (*Max35Text)(&value)
}

func (i *Identification7) SetPoolIdentification(value string) {
	i.PoolIdentification = (*Max35Text)(&value)
}

func (i *Identification7) SetCorporateActionEventIdentification(value string) {
	i.CorporateActionEventIdentification = (*Max35Text)(&value)
}
