package iso20022

// Sort criteria.
type AdditionalParameters1 struct {

	// Specifies the country.
	Country *CountryCode `xml:"Ctry,omitempty"`

	// Specifies the currency.
	Currency *ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty"`

	// Specifies the geographical area, eg, Asia-Pacific, Europe, Middle-East.
	GeographicalArea *Max35Text `xml:"GeoArea,omitempty"`
}

func (a *AdditionalParameters1) SetCountry(value string) {
	a.Country = (*CountryCode)(&value)
}

func (a *AdditionalParameters1) SetCurrency(value string) {
	a.Currency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (a *AdditionalParameters1) SetGeographicalArea(value string) {
	a.GeographicalArea = (*Max35Text)(&value)
}

// Specifies additional parameters to the message or transaction.
type AdditionalParameters10 struct {

	// Specifies partial settlement information.
	PartialSettlement *PartialSettlement2Code `xml:"PrtlSttlm,omitempty"`

	// Identification of the confirmation previously sent to confirm the partial settlement of a transaction.
	PreviousPartialConfirmationIdentification *Max35Text `xml:"PrvsPrtlConfId,omitempty"`

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`
}

func (a *AdditionalParameters10) SetPartialSettlement(value string) {
	a.PartialSettlement = (*PartialSettlement2Code)(&value)
}

func (a *AdditionalParameters10) SetPreviousPartialConfirmationIdentification(value string) {
	a.PreviousPartialConfirmationIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters10) SetAccountOwnerTransactionIdentification(value string) {
	a.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters10) SetAccountServicerTransactionIdentification(value string) {
	a.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters10) SetPoolIdentification(value string) {
	a.PoolIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters10) SetCorporateActionEventIdentification(value string) {
	a.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters10) SetMarketInfrastructureTransactionIdentification(value string) {
	a.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters10) SetProcessorTransactionIdentification(value string) {
	a.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

// Specifies additional parameters to the message or transaction.
type AdditionalParameters12 struct {

	// Specifies whether there exists a pre-confirmation.
	PreConfirmation *PreConfirmation1Code `xml:"PreConf,omitempty"`

	// Specifies partial settlement information.
	PartialSettlement *PartialSettlement2Code `xml:"PrtlSttlm,omitempty"`

	// Identification of the confirmation previously sent to confirm the partial settlement of a transaction.
	PreviousPartialConfirmationIdentification *Max35Text `xml:"PrvsPrtlConfId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`
}

func (a *AdditionalParameters12) SetPreConfirmation(value string) {
	a.PreConfirmation = (*PreConfirmation1Code)(&value)
}

func (a *AdditionalParameters12) SetPartialSettlement(value string) {
	a.PartialSettlement = (*PartialSettlement2Code)(&value)
}

func (a *AdditionalParameters12) SetPreviousPartialConfirmationIdentification(value string) {
	a.PreviousPartialConfirmationIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters12) SetTripartyAgentCollateralTransactionIdentification(value string) {
	a.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters12) SetClientTripartyCollateralTransactionIdentification(value string) {
	a.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

// Specifies additional parameters to the message or transaction.
type AdditionalParameters14 struct {

	// Specifies whether there exists a pre-confirmation.
	PreConfirmation *PreConfirmation1Code `xml:"PreConf,omitempty"`

	// Specifies partial settlement information.
	PartialSettlement *PartialSettlement2Code `xml:"PrtlSttlm,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`
}

func (a *AdditionalParameters14) SetPreConfirmation(value string) {
	a.PreConfirmation = (*PreConfirmation1Code)(&value)
}

func (a *AdditionalParameters14) SetPartialSettlement(value string) {
	a.PartialSettlement = (*PartialSettlement2Code)(&value)
}

func (a *AdditionalParameters14) SetTripartyAgentCollateralTransactionIdentification(value string) {
	a.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters14) SetClientTripartyCollateralTransactionIdentification(value string) {
	a.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

// Specifies additional parameters to the message or transaction.
type AdditionalParameters17 struct {

	// Specifies whether there exists a pre-confirmation.
	PreConfirmation *PreConfirmation1Code `xml:"PreConf,omitempty"`

	// Specifies partial settlement information.
	PartialSettlement *PartialSettlement2Code `xml:"PrtlSttlm,omitempty"`

	// Identification of the confirmation previously sent to confirm the partial settlement of a transaction.
	PreviousPartialConfirmationIdentification *Max35Text `xml:"PrvsPrtlConfId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Unique identification assigned to the instruction by the client.
	ClientCollateralInstructionIdentification *Max35Text `xml:"ClntCollInstrId,omitempty"`

	// Unique identification assigned to the instruction by the triparty agent.
	TripartyCollateralInstructionIdentification *Max35Text `xml:"TrptyCollInstrId,omitempty"`
}

func (a *AdditionalParameters17) SetPreConfirmation(value string) {
	a.PreConfirmation = (*PreConfirmation1Code)(&value)
}

func (a *AdditionalParameters17) SetPartialSettlement(value string) {
	a.PartialSettlement = (*PartialSettlement2Code)(&value)
}

func (a *AdditionalParameters17) SetPreviousPartialConfirmationIdentification(value string) {
	a.PreviousPartialConfirmationIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters17) SetTripartyAgentCollateralTransactionIdentification(value string) {
	a.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters17) SetClientTripartyCollateralTransactionIdentification(value string) {
	a.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters17) SetClientCollateralInstructionIdentification(value string) {
	a.ClientCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters17) SetTripartyCollateralInstructionIdentification(value string) {
	a.TripartyCollateralInstructionIdentification = (*Max35Text)(&value)
}

// Specifies additional parameters to the message or transaction.
type AdditionalParameters18 struct {

	// Specifies whether there exists a pre-confirmation.
	PreConfirmation *PreConfirmation1Code `xml:"PreConf,omitempty"`

	// Specifies partial settlement information.
	PartialSettlement *PartialSettlement2Code `xml:"PrtlSttlm,omitempty"`

	// Identification of the confirmation previously sent to confirm the partial settlement of a transaction.
	PreviousPartialConfirmationIdentification *Max35Text `xml:"PrvsPrtlConfId,omitempty"`
}

func (a *AdditionalParameters18) SetPreConfirmation(value string) {
	a.PreConfirmation = (*PreConfirmation1Code)(&value)
}

func (a *AdditionalParameters18) SetPartialSettlement(value string) {
	a.PartialSettlement = (*PartialSettlement2Code)(&value)
}

func (a *AdditionalParameters18) SetPreviousPartialConfirmationIdentification(value string) {
	a.PreviousPartialConfirmationIdentification = (*Max35Text)(&value)
}

// Specifies additional parameters to the message or transaction.
type AdditionalParameters2 struct {

	// Specifies whether there exists a pre-confirmation.
	PreConfirmation *PreConfirmation1Code `xml:"PreConf,omitempty"`

	// Specifies partial settlement information.
	PartialSettlement *PartialSettlement1Code `xml:"PrtlSttlm,omitempty"`

	// Identification of the confirmation previously sent to confirm the partial settlement of a transaction.
	PreviousPartialConfirmationIdentification *Max35Text `xml:"PrvsPrtlConfId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`
}

func (a *AdditionalParameters2) SetPreConfirmation(value string) {
	a.PreConfirmation = (*PreConfirmation1Code)(&value)
}

func (a *AdditionalParameters2) SetPartialSettlement(value string) {
	a.PartialSettlement = (*PartialSettlement1Code)(&value)
}

func (a *AdditionalParameters2) SetPreviousPartialConfirmationIdentification(value string) {
	a.PreviousPartialConfirmationIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters2) SetTripartyAgentCollateralTransactionIdentification(value string) {
	a.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters2) SetClientTripartyCollateralTransactionIdentification(value string) {
	a.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

// Specifies additional parameters to the message or transaction.
type AdditionalParameters21 struct {

	// Specifies partial settlement information.
	PartialSettlement *PartialSettlement2Code `xml:"PrtlSttlm,omitempty"`

	// Identification of the confirmation previously sent to confirm the partial settlement of a transaction.
	PreviousPartialConfirmationIdentification *Max35Text `xml:"PrvsPrtlConfId,omitempty"`

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`
}

func (a *AdditionalParameters21) SetPartialSettlement(value string) {
	a.PartialSettlement = (*PartialSettlement2Code)(&value)
}

func (a *AdditionalParameters21) SetPreviousPartialConfirmationIdentification(value string) {
	a.PreviousPartialConfirmationIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters21) SetAccountOwnerTransactionIdentification(value string) {
	a.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters21) SetAccountServicerTransactionIdentification(value string) {
	a.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters21) SetPoolIdentification(value string) {
	a.PoolIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters21) SetCorporateActionEventIdentification(value string) {
	a.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters21) SetMarketInfrastructureTransactionIdentification(value string) {
	a.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters21) SetProcessorTransactionIdentification(value string) {
	a.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

// Specifies additional parameters to the message or transaction.
type AdditionalParameters22 struct {

	// Specifies whether there exists a pre-confirmation.
	PreConfirmation *PreConfirmation1Code `xml:"PreConf,omitempty"`

	// Specifies partial settlement information.
	PartialSettlement *PartialSettlement2Code `xml:"PrtlSttlm,omitempty"`

	// Identification of the confirmation previously sent to confirm the partial settlement of a transaction.
	PreviousPartialConfirmationIdentification *Max35Text `xml:"PrvsPrtlConfId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Unique identification assigned to the instruction by the client.
	ClientCollateralInstructionIdentification *Max35Text `xml:"ClntCollInstrId,omitempty"`

	// Unique identification assigned to the instruction by the triparty agent.
	TripartyCollateralInstructionIdentification *Max35Text `xml:"TrptyCollInstrId,omitempty"`
}

func (a *AdditionalParameters22) SetPreConfirmation(value string) {
	a.PreConfirmation = (*PreConfirmation1Code)(&value)
}

func (a *AdditionalParameters22) SetPartialSettlement(value string) {
	a.PartialSettlement = (*PartialSettlement2Code)(&value)
}

func (a *AdditionalParameters22) SetPreviousPartialConfirmationIdentification(value string) {
	a.PreviousPartialConfirmationIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters22) SetTripartyAgentCollateralTransactionIdentification(value string) {
	a.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters22) SetClientTripartyCollateralTransactionIdentification(value string) {
	a.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters22) SetClientCollateralInstructionIdentification(value string) {
	a.ClientCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters22) SetTripartyCollateralInstructionIdentification(value string) {
	a.TripartyCollateralInstructionIdentification = (*Max35Text)(&value)
}

// Specifies additional parameters to the message or transaction.
type AdditionalParameters23 struct {

	// Specifies whether there exists a pre-confirmation.
	PreConfirmation *PreConfirmation1Code `xml:"PreConf,omitempty"`

	// Specifies partial settlement information.
	PartialSettlement *PartialSettlement2Code `xml:"PrtlSttlm,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`
}

func (a *AdditionalParameters23) SetPreConfirmation(value string) {
	a.PreConfirmation = (*PreConfirmation1Code)(&value)
}

func (a *AdditionalParameters23) SetPartialSettlement(value string) {
	a.PartialSettlement = (*PartialSettlement2Code)(&value)
}

func (a *AdditionalParameters23) SetTripartyAgentCollateralTransactionIdentification(value string) {
	a.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters23) SetClientTripartyCollateralTransactionIdentification(value string) {
	a.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

// Specifies additional parameters to the message or transaction.
type AdditionalParameters24 struct {

	// Specifies whether there exists a pre-confirmation.
	PreConfirmation *PreConfirmation1Code `xml:"PreConf,omitempty"`

	// Specifies partial settlement information.
	PartialSettlement *PartialSettlement2Code `xml:"PrtlSttlm,omitempty"`

	// Identification of the confirmation previously sent to confirm the partial settlement of a transaction.
	PreviousPartialConfirmationIdentification *Max35Text `xml:"PrvsPrtlConfId,omitempty"`
}

func (a *AdditionalParameters24) SetPreConfirmation(value string) {
	a.PreConfirmation = (*PreConfirmation1Code)(&value)
}

func (a *AdditionalParameters24) SetPartialSettlement(value string) {
	a.PartialSettlement = (*PartialSettlement2Code)(&value)
}

func (a *AdditionalParameters24) SetPreviousPartialConfirmationIdentification(value string) {
	a.PreviousPartialConfirmationIdentification = (*Max35Text)(&value)
}

// Specifies additional parameters to the message or transaction.
type AdditionalParameters25 struct {

	// Specifies partial settlement information.
	PartialSettlement *PartialSettlement2Code `xml:"PrtlSttlm,omitempty"`

	// Identification of the confirmation previously sent to confirm the partial settlement of a transaction.
	PreviousPartialConfirmationIdentification *RestrictedFINXMax16Text `xml:"PrvsPrtlConfId,omitempty"`

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctOwnrTxId,omitempty"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctSvcrTxId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *RestrictedFINXMax16Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *RestrictedFINXMax16Text `xml:"PrcrTxId,omitempty"`
}

func (a *AdditionalParameters25) SetPartialSettlement(value string) {
	a.PartialSettlement = (*PartialSettlement2Code)(&value)
}

func (a *AdditionalParameters25) SetPreviousPartialConfirmationIdentification(value string) {
	a.PreviousPartialConfirmationIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (a *AdditionalParameters25) SetAccountOwnerTransactionIdentification(value string) {
	a.AccountOwnerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (a *AdditionalParameters25) SetAccountServicerTransactionIdentification(value string) {
	a.AccountServicerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (a *AdditionalParameters25) SetPoolIdentification(value string) {
	a.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (a *AdditionalParameters25) SetCorporateActionEventIdentification(value string) {
	a.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (a *AdditionalParameters25) SetMarketInfrastructureTransactionIdentification(value string) {
	a.MarketInfrastructureTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (a *AdditionalParameters25) SetProcessorTransactionIdentification(value string) {
	a.ProcessorTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

// Specifies additional parameters to the message or transaction.
type AdditionalParameters26 struct {

	// Specifies whether there exists a pre-confirmation.
	PreConfirmation *PreConfirmation1Code `xml:"PreConf,omitempty"`

	// Specifies partial settlement information.
	PartialSettlement *PartialSettlement2Code `xml:"PrtlSttlm,omitempty"`

	// Identification of the confirmation previously sent to confirm the partial settlement of a transaction.
	PreviousPartialConfirmationIdentification *RestrictedFINXMax16Text `xml:"PrvsPrtlConfId,omitempty"`
}

func (a *AdditionalParameters26) SetPreConfirmation(value string) {
	a.PreConfirmation = (*PreConfirmation1Code)(&value)
}

func (a *AdditionalParameters26) SetPartialSettlement(value string) {
	a.PartialSettlement = (*PartialSettlement2Code)(&value)
}

func (a *AdditionalParameters26) SetPreviousPartialConfirmationIdentification(value string) {
	a.PreviousPartialConfirmationIdentification = (*RestrictedFINXMax16Text)(&value)
}

// Specifies additional parameters to the message or transaction.
type AdditionalParameters27 struct {

	// Specifies whether there exists a pre-confirmation.
	PreConfirmation *PreConfirmation1Code `xml:"PreConf,omitempty"`

	// Specifies partial settlement information.
	PartialSettlement *PartialSettlement2Code `xml:"PrtlSttlm,omitempty"`

	// Identification of the confirmation previously sent to confirm the partial settlement of a transaction.
	PreviousPartialConfirmationIdentification *RestrictedFINXMax16Text `xml:"PrvsPrtlConfId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Unique identification assigned to the instruction by the client.
	ClientCollateralInstructionIdentification *RestrictedFINXMax16Text `xml:"ClntCollInstrId,omitempty"`

	// Unique identification assigned to the instruction by the triparty agent.
	TripartyCollateralInstructionIdentification *RestrictedFINXMax16Text `xml:"TrptyCollInstrId,omitempty"`
}

func (a *AdditionalParameters27) SetPreConfirmation(value string) {
	a.PreConfirmation = (*PreConfirmation1Code)(&value)
}

func (a *AdditionalParameters27) SetPartialSettlement(value string) {
	a.PartialSettlement = (*PartialSettlement2Code)(&value)
}

func (a *AdditionalParameters27) SetPreviousPartialConfirmationIdentification(value string) {
	a.PreviousPartialConfirmationIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (a *AdditionalParameters27) SetTripartyAgentCollateralTransactionIdentification(value string) {
	a.TripartyAgentCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (a *AdditionalParameters27) SetClientTripartyCollateralTransactionIdentification(value string) {
	a.ClientTripartyCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (a *AdditionalParameters27) SetClientCollateralInstructionIdentification(value string) {
	a.ClientCollateralInstructionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (a *AdditionalParameters27) SetTripartyCollateralInstructionIdentification(value string) {
	a.TripartyCollateralInstructionIdentification = (*RestrictedFINXMax16Text)(&value)
}

// Specifies additional parameters to the message or transaction.
type AdditionalParameters28 struct {

	// Specifies whether there exists a pre-confirmation.
	PreConfirmation *PreConfirmation1Code `xml:"PreConf,omitempty"`

	// Specifies partial settlement information.
	PartialSettlement *PartialSettlement2Code `xml:"PrtlSttlm,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"ClntTrptyCollTxId,omitempty"`
}

func (a *AdditionalParameters28) SetPreConfirmation(value string) {
	a.PreConfirmation = (*PreConfirmation1Code)(&value)
}

func (a *AdditionalParameters28) SetPartialSettlement(value string) {
	a.PartialSettlement = (*PartialSettlement2Code)(&value)
}

func (a *AdditionalParameters28) SetTripartyAgentCollateralTransactionIdentification(value string) {
	a.TripartyAgentCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (a *AdditionalParameters28) SetClientTripartyCollateralTransactionIdentification(value string) {
	a.ClientTripartyCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

// Specifies additional parameters to the message or transaction.
type AdditionalParameters29 struct {

	// Specifies whether there exists a pre-confirmation.
	PreConfirmation *PreConfirmation1Code `xml:"PreConf,omitempty"`

	// Specifies partial settlement information.
	PartialSettlement *PartialSettlement2Code `xml:"PrtlSttlm,omitempty"`

	// Identification of the confirmation previously sent to confirm the partial settlement of a transaction.
	PreviousPartialConfirmationIdentification *Max35Text `xml:"PrvsPrtlConfId,omitempty"`

	// Unique identification identifying the triparty collateral management transaction from the triparty-agent's/service-provider's point of view.
	TripartyAgentServiceProviderCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtSvcPrvdrCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Unique identification assigned to the instruction by the client.
	ClientCollateralInstructionIdentification *Max35Text `xml:"ClntCollInstrId,omitempty"`

	// Unique identification assigned to the instruction by the triparty-agent/service-provider.
	TripartyAgentServiceProviderCollateralInstructionIdentification *Max35Text `xml:"TrptyAgtSvcPrvdrCollInstrId,omitempty"`
}

func (a *AdditionalParameters29) SetPreConfirmation(value string) {
	a.PreConfirmation = (*PreConfirmation1Code)(&value)
}

func (a *AdditionalParameters29) SetPartialSettlement(value string) {
	a.PartialSettlement = (*PartialSettlement2Code)(&value)
}

func (a *AdditionalParameters29) SetPreviousPartialConfirmationIdentification(value string) {
	a.PreviousPartialConfirmationIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters29) SetTripartyAgentServiceProviderCollateralTransactionIdentification(value string) {
	a.TripartyAgentServiceProviderCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters29) SetClientTripartyCollateralTransactionIdentification(value string) {
	a.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters29) SetClientCollateralInstructionIdentification(value string) {
	a.ClientCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters29) SetTripartyAgentServiceProviderCollateralInstructionIdentification(value string) {
	a.TripartyAgentServiceProviderCollateralInstructionIdentification = (*Max35Text)(&value)
}

// Specifies additional parameters to the message or transaction.
type AdditionalParameters3 struct {

	// Specifies partial settlement information.
	PartialSettlement *PartialSettlement1Code `xml:"PrtlSttlm,omitempty"`

	// Identification of the confirmation previously sent to confirm the partial settlement of a transaction.
	PreviousPartialConfirmationIdentification *Max35Text `xml:"PrvsPrtlConfId,omitempty"`

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`
}

func (a *AdditionalParameters3) SetPartialSettlement(value string) {
	a.PartialSettlement = (*PartialSettlement1Code)(&value)
}

func (a *AdditionalParameters3) SetPreviousPartialConfirmationIdentification(value string) {
	a.PreviousPartialConfirmationIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters3) SetAccountOwnerTransactionIdentification(value string) {
	a.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters3) SetAccountServicerTransactionIdentification(value string) {
	a.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters3) SetPoolIdentification(value string) {
	a.PoolIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters3) SetCorporateActionEventIdentification(value string) {
	a.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters3) SetMarketInfrastructureTransactionIdentification(value string) {
	a.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

// Specifies additional parameters to the message or transaction.
type AdditionalParameters30 struct {

	// Specifies whether there exists a pre-confirmation.
	PreConfirmation *PreConfirmation1Code `xml:"PreConf,omitempty"`

	// Specifies partial settlement information.
	PartialSettlement *PartialSettlement2Code `xml:"PrtlSttlm,omitempty"`

	// Unique identification identifying the triparty collateral management transaction from the triparty-agent's/service-provider's point of view.
	TripartyAgentServiceProviderCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtSvcPrvdrCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`
}

func (a *AdditionalParameters30) SetPreConfirmation(value string) {
	a.PreConfirmation = (*PreConfirmation1Code)(&value)
}

func (a *AdditionalParameters30) SetPartialSettlement(value string) {
	a.PartialSettlement = (*PartialSettlement2Code)(&value)
}

func (a *AdditionalParameters30) SetTripartyAgentServiceProviderCollateralTransactionIdentification(value string) {
	a.TripartyAgentServiceProviderCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters30) SetClientTripartyCollateralTransactionIdentification(value string) {
	a.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

// Specifies additional parameters to the message or transaction.
type AdditionalParameters31 struct {

	// Specifies whether there exists a pre-confirmation.
	PreConfirmation *PreConfirmation1Code `xml:"PreConf,omitempty"`

	// Specifies partial settlement information.
	PartialSettlement *PartialSettlement2Code `xml:"PrtlSttlm,omitempty"`

	// Identification of the confirmation previously sent to confirm the partial settlement of a transaction.
	PreviousPartialConfirmationIdentification *RestrictedFINXMax16Text `xml:"PrvsPrtlConfId,omitempty"`

	// Unique identification identifying the triparty collateral management transaction from the triparty-agent's/service-provider's point of view.
	TripartyAgentServiceProviderCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"TrptyAgtSvcPrvdrCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Unique identification assigned to the instruction by the client.
	ClientCollateralInstructionIdentification *RestrictedFINXMax16Text `xml:"ClntCollInstrId,omitempty"`

	// Unique identification assigned to the instruction by the triparty-agent/service-provider.
	TripartyAgentServiceProviderCollateralInstructionIdentification *RestrictedFINXMax16Text `xml:"TrptyAgtSvcPrvdrCollInstrId,omitempty"`
}

func (a *AdditionalParameters31) SetPreConfirmation(value string) {
	a.PreConfirmation = (*PreConfirmation1Code)(&value)
}

func (a *AdditionalParameters31) SetPartialSettlement(value string) {
	a.PartialSettlement = (*PartialSettlement2Code)(&value)
}

func (a *AdditionalParameters31) SetPreviousPartialConfirmationIdentification(value string) {
	a.PreviousPartialConfirmationIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (a *AdditionalParameters31) SetTripartyAgentServiceProviderCollateralTransactionIdentification(value string) {
	a.TripartyAgentServiceProviderCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (a *AdditionalParameters31) SetClientTripartyCollateralTransactionIdentification(value string) {
	a.ClientTripartyCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (a *AdditionalParameters31) SetClientCollateralInstructionIdentification(value string) {
	a.ClientCollateralInstructionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (a *AdditionalParameters31) SetTripartyAgentServiceProviderCollateralInstructionIdentification(value string) {
	a.TripartyAgentServiceProviderCollateralInstructionIdentification = (*RestrictedFINXMax16Text)(&value)
}

// Specifies additional parameters to the message or transaction.
type AdditionalParameters32 struct {

	// Specifies whether there exists a pre-confirmation.
	PreConfirmation *PreConfirmation1Code `xml:"PreConf,omitempty"`

	// Specifies partial settlement information.
	PartialSettlement *PartialSettlement2Code `xml:"PrtlSttlm,omitempty"`

	// Unique identification identifying the triparty collateral management transaction from the triparty-agent's/service-provider's point of view.
	TripartyAgentServiceProviderCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"TrptyAgtSvcPrvdrCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"ClntTrptyCollTxId,omitempty"`
}

func (a *AdditionalParameters32) SetPreConfirmation(value string) {
	a.PreConfirmation = (*PreConfirmation1Code)(&value)
}

func (a *AdditionalParameters32) SetPartialSettlement(value string) {
	a.PartialSettlement = (*PartialSettlement2Code)(&value)
}

func (a *AdditionalParameters32) SetTripartyAgentServiceProviderCollateralTransactionIdentification(value string) {
	a.TripartyAgentServiceProviderCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (a *AdditionalParameters32) SetClientTripartyCollateralTransactionIdentification(value string) {
	a.ClientTripartyCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

// Specifies additional parameters to the message or transaction.
type AdditionalParameters4 struct {

	// Specifies whether there exists a pre-confirmation.
	PreConfirmation *PreConfirmation1Code `xml:"PreConf,omitempty"`

	// Specifies partial settlement information.
	PartialSettlement *PartialSettlement1Code `xml:"PrtlSttlm,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`
}

func (a *AdditionalParameters4) SetPreConfirmation(value string) {
	a.PreConfirmation = (*PreConfirmation1Code)(&value)
}

func (a *AdditionalParameters4) SetPartialSettlement(value string) {
	a.PartialSettlement = (*PartialSettlement1Code)(&value)
}

func (a *AdditionalParameters4) SetTripartyAgentCollateralTransactionIdentification(value string) {
	a.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters4) SetClientTripartyCollateralTransactionIdentification(value string) {
	a.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

// Specifies additional parameters to the message or transaction.
type AdditionalParameters8 struct {

	// Specifies partial settlement information.
	PartialSettlement *PartialSettlement1Code `xml:"PrtlSttlm,omitempty"`

	// Identification of the confirmation previously sent to confirm the partial settlement of a transaction.
	PreviousPartialConfirmationIdentification *Max35Text `xml:"PrvsPrtlConfId,omitempty"`

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`
}

func (a *AdditionalParameters8) SetPartialSettlement(value string) {
	a.PartialSettlement = (*PartialSettlement1Code)(&value)
}

func (a *AdditionalParameters8) SetPreviousPartialConfirmationIdentification(value string) {
	a.PreviousPartialConfirmationIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters8) SetAccountOwnerTransactionIdentification(value string) {
	a.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters8) SetAccountServicerTransactionIdentification(value string) {
	a.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters8) SetPoolIdentification(value string) {
	a.PoolIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters8) SetCorporateActionEventIdentification(value string) {
	a.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters8) SetMarketInfrastructureTransactionIdentification(value string) {
	a.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (a *AdditionalParameters8) SetProcessorTransactionIdentification(value string) {
	a.ProcessorTransactionIdentification = (*Max35Text)(&value)
}
