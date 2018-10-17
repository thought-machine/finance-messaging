package iso20022

// Information needed due to regulatory and/or statutory requirements.
type RegulatoryReporting2 struct {

	// Identifies whether the regulatory reporting information applies to the debit side, to the credit side or to both debit and credit sides of the transaction.
	DebitCreditReportingIndicator *RegulatoryReportingType1Code `xml:"DbtCdtRptgInd,omitempty"`

	// Entity requiring the regulatory reporting information.
	Authority *RegulatoryAuthority `xml:"Authrty,omitempty"`

	// Details related to the regulatory reporting information.
	RegulatoryDetails *StructuredRegulatoryReporting2 `xml:"RgltryDtls,omitempty"`
}

func (r *RegulatoryReporting2) SetDebitCreditReportingIndicator(value string) {
	r.DebitCreditReportingIndicator = (*RegulatoryReportingType1Code)(&value)
}

func (r *RegulatoryReporting2) AddAuthority() *RegulatoryAuthority {
	r.Authority = new(RegulatoryAuthority)
	return r.Authority
}

func (r *RegulatoryReporting2) AddRegulatoryDetails() *StructuredRegulatoryReporting2 {
	r.RegulatoryDetails = new(StructuredRegulatoryReporting2)
	return r.RegulatoryDetails
}

// Information needed due to regulatory and/or statutory requirements.
type RegulatoryReporting3 struct {

	// Identifies whether the regulatory reporting information applies to the debit side, to the credit side or to both debit and credit sides of the transaction.
	DebitCreditReportingIndicator *RegulatoryReportingType1Code `xml:"DbtCdtRptgInd,omitempty"`

	// Entity requiring the regulatory reporting information.
	Authority *RegulatoryAuthority2 `xml:"Authrty,omitempty"`

	// Set of elements used to provide details on the regulatory reporting information.
	Details []*StructuredRegulatoryReporting3 `xml:"Dtls,omitempty"`
}

func (r *RegulatoryReporting3) SetDebitCreditReportingIndicator(value string) {
	r.DebitCreditReportingIndicator = (*RegulatoryReportingType1Code)(&value)
}

func (r *RegulatoryReporting3) AddAuthority() *RegulatoryAuthority2 {
	r.Authority = new(RegulatoryAuthority2)
	return r.Authority
}

func (r *RegulatoryReporting3) AddDetails() *StructuredRegulatoryReporting3 {
	newValue := new(StructuredRegulatoryReporting3)
	r.Details = append(r.Details, newValue)
	return newValue
}

// Includes data elements that can be used for reporting to trade repositories, it is not to be used on regular trade confirmations. Although some fields, for example, unique transaction identifier and prior unique transaction identifier, might be used on regular trade confirmations.
type RegulatoryReporting4 struct {

	// Regulatory transaction reporting information from the Trading Side party.
	TradingSideTransactionReporting []*TradingSideTransactionReporting1 `xml:"TradgSdTxRptg,omitempty"`

	// Regulatory transaction reporting information from the Counterparty Side party.
	CounterpartySideTransactionReporting []*CounterpartySideTransactionReporting1 `xml:"CtrPtySdTxRptg,omitempty"`

	// Identifies an agency or separate corporation of a futures exchange responsible for settling and
	// clearing trades, collecting and maintaining margins, regulating delivery and reporting trade data. This can also be known as a Central Counterparty (CCP).
	CentralCounterpartyClearingHouse *PartyIdentification73Choice `xml:"CntrlCtrPtyClrHs,omitempty"`

	// Identifies the party that is a member of the clearing house (CCP) and that acts as a liaison between the investor and the Cntral Counterparty (CCP).
	ClearingBroker *PartyIdentification73Choice `xml:"ClrBrkr,omitempty"`

	// Identifies the party that is exempt from a clearing obligation.
	ClearingExceptionParty *PartyIdentification73Choice `xml:"ClrXcptnPty,omitempty"`

	// Specifies the reference number assigned by the clearing broker. A distinction can be made between the reference for the Central Counterparty (CCP) leg and the reference for the client leg of the transaction.
	ClearingBrokerIdentification *ClearingBrokerIdentification1 `xml:"ClrBrkrId,omitempty"`

	// Specifies whether the contract is above or below the clearing threshold. Where No indicates the contract is below the clearing threshold and Yes indicates the contract is above the clearing threshold.
	ClearingThresholdIndicator *YesNoIndicator `xml:"ClrThrshldInd,omitempty"`

	// Specifies the reference number assigned by the Central Counterparty (CCP).
	ClearedProductIdentification *Max35Text `xml:"ClrdPdctId,omitempty"`

	// Specifies the underlying product type.
	UnderlyingProductIdentifier *UnderlyingProductIdentifier1Code `xml:"UndrlygPdctIdr,omitempty"`

	// Specifies whether the trade is a pre-allocation or a post-allocation trade, or whether the trade is unallocated.
	AllocationIndicator *AllocationIndicator1Code `xml:"AllcnInd,omitempty"`

	// Specifies whether the transaction is collateralised.
	CollateralisationIndicator *CollateralisationIndicator1Code `xml:"CollstnInd,omitempty"`

	// Specifies the trading venue of the transaction.
	ExecutionVenue *Max35Text `xml:"ExctnVn,omitempty"`

	// Specifies the date and time of the execution of the transaction in Coordinated Universal Time (UTC).
	ExecutionTimestamp *DateAndDateTimeChoice `xml:"ExctnTmstmp,omitempty"`

	// Specifies whether the reportable transaction has one or more additional terms or provisions, other than those listed in the required real-time data fields, that materially affects the price of the reportable transaction.
	NonStandardFlag *YesNoIndicator `xml:"NonStdFlg,omitempty"`

	// Specifies the common reference or correlation identification for a swap transaction where the near and far leg are confirmed separately.
	LinkSwapIdentification *Exact42Text `xml:"LkSwpId,omitempty"`

	// Specifies the financial nature of the reporting counterparty.
	FinancialNatureOfTheCounterpartyIndicator *YesNoIndicator `xml:"FinNtrOfTheCtrPtyInd,omitempty"`

	// Specifies if the collateral is posted on a portfolio basis.
	CollateralPortfolioIndicator *YesNoIndicator `xml:"CollPrtflInd,omitempty"`

	// Identifies the portfolio code to which the trade belongs if the collateral is posted on a portfolio basis (and not trade by trade).
	CollateralPortfolioCode *Max10Text `xml:"CollPrtflCd,omitempty"`

	// Indicates if the trade results from portfolio compression.
	PortfolioCompressionIndicator *YesNoIndicator `xml:"PrtflCmprssnInd,omitempty"`

	// Specifies the corporate sector of the counterparty.
	CorporateSectorIndicator *CorporateSectorIdentifier1Code `xml:"CorpSctrInd,omitempty"`

	// Specifies whether the counterparty has entered into a trade with a non-European Economic Area (EEA) counterparty that is not subject to the reporting obligation.
	TradeWithNonEEACounterpartyIndicator *YesNoIndicator `xml:"TradWthNonEEACtrPtyInd,omitempty"`

	// To indicate if a reported trade falls under the definition of intragroup transaction, as defined by European Securities and Markets Authority (ESMA) in the Technical Standards.
	IntragroupTradeIndicator *YesNoIndicator `xml:"NtrgrpTradInd,omitempty"`

	// Specifies whether the contract is objectively measurable as directly linked to the non-financial counterparty's commercial or treasury financing activity.
	CommercialOrTreasuryFinancingIndicator *YesNoIndicator `xml:"ComrclOrTrsrFincgInd,omitempty"`

	// Specifies additional information that might be required by the regulator.
	AdditionalReportingInformation *Max210Text `xml:"AddtlRptgInf,omitempty"`
}

func (r *RegulatoryReporting4) AddTradingSideTransactionReporting() *TradingSideTransactionReporting1 {
	newValue := new(TradingSideTransactionReporting1)
	r.TradingSideTransactionReporting = append(r.TradingSideTransactionReporting, newValue)
	return newValue
}

func (r *RegulatoryReporting4) AddCounterpartySideTransactionReporting() *CounterpartySideTransactionReporting1 {
	newValue := new(CounterpartySideTransactionReporting1)
	r.CounterpartySideTransactionReporting = append(r.CounterpartySideTransactionReporting, newValue)
	return newValue
}

func (r *RegulatoryReporting4) AddCentralCounterpartyClearingHouse() *PartyIdentification73Choice {
	r.CentralCounterpartyClearingHouse = new(PartyIdentification73Choice)
	return r.CentralCounterpartyClearingHouse
}

func (r *RegulatoryReporting4) AddClearingBroker() *PartyIdentification73Choice {
	r.ClearingBroker = new(PartyIdentification73Choice)
	return r.ClearingBroker
}

func (r *RegulatoryReporting4) AddClearingExceptionParty() *PartyIdentification73Choice {
	r.ClearingExceptionParty = new(PartyIdentification73Choice)
	return r.ClearingExceptionParty
}

func (r *RegulatoryReporting4) AddClearingBrokerIdentification() *ClearingBrokerIdentification1 {
	r.ClearingBrokerIdentification = new(ClearingBrokerIdentification1)
	return r.ClearingBrokerIdentification
}

func (r *RegulatoryReporting4) SetClearingThresholdIndicator(value string) {
	r.ClearingThresholdIndicator = (*YesNoIndicator)(&value)
}

func (r *RegulatoryReporting4) SetClearedProductIdentification(value string) {
	r.ClearedProductIdentification = (*Max35Text)(&value)
}

func (r *RegulatoryReporting4) SetUnderlyingProductIdentifier(value string) {
	r.UnderlyingProductIdentifier = (*UnderlyingProductIdentifier1Code)(&value)
}

func (r *RegulatoryReporting4) SetAllocationIndicator(value string) {
	r.AllocationIndicator = (*AllocationIndicator1Code)(&value)
}

func (r *RegulatoryReporting4) SetCollateralisationIndicator(value string) {
	r.CollateralisationIndicator = (*CollateralisationIndicator1Code)(&value)
}

func (r *RegulatoryReporting4) SetExecutionVenue(value string) {
	r.ExecutionVenue = (*Max35Text)(&value)
}

func (r *RegulatoryReporting4) AddExecutionTimestamp() *DateAndDateTimeChoice {
	r.ExecutionTimestamp = new(DateAndDateTimeChoice)
	return r.ExecutionTimestamp
}

func (r *RegulatoryReporting4) SetNonStandardFlag(value string) {
	r.NonStandardFlag = (*YesNoIndicator)(&value)
}

func (r *RegulatoryReporting4) SetLinkSwapIdentification(value string) {
	r.LinkSwapIdentification = (*Exact42Text)(&value)
}

func (r *RegulatoryReporting4) SetFinancialNatureOfTheCounterpartyIndicator(value string) {
	r.FinancialNatureOfTheCounterpartyIndicator = (*YesNoIndicator)(&value)
}

func (r *RegulatoryReporting4) SetCollateralPortfolioIndicator(value string) {
	r.CollateralPortfolioIndicator = (*YesNoIndicator)(&value)
}

func (r *RegulatoryReporting4) SetCollateralPortfolioCode(value string) {
	r.CollateralPortfolioCode = (*Max10Text)(&value)
}

func (r *RegulatoryReporting4) SetPortfolioCompressionIndicator(value string) {
	r.PortfolioCompressionIndicator = (*YesNoIndicator)(&value)
}

func (r *RegulatoryReporting4) SetCorporateSectorIndicator(value string) {
	r.CorporateSectorIndicator = (*CorporateSectorIdentifier1Code)(&value)
}

func (r *RegulatoryReporting4) SetTradeWithNonEEACounterpartyIndicator(value string) {
	r.TradeWithNonEEACounterpartyIndicator = (*YesNoIndicator)(&value)
}

func (r *RegulatoryReporting4) SetIntragroupTradeIndicator(value string) {
	r.IntragroupTradeIndicator = (*YesNoIndicator)(&value)
}

func (r *RegulatoryReporting4) SetCommercialOrTreasuryFinancingIndicator(value string) {
	r.CommercialOrTreasuryFinancingIndicator = (*YesNoIndicator)(&value)
}

func (r *RegulatoryReporting4) SetAdditionalReportingInformation(value string) {
	r.AdditionalReportingInformation = (*Max210Text)(&value)
}

// Includes data elements that can be used for reporting to trade repositories, it is not to be used on regular trade confirmations. Although some fields, for example, unique transaction identifier and prior unique transaction identifier, might be used on regular trade confirmations.
type RegulatoryReporting6 struct {

	// Regulatory transaction reporting information from the Trading Side party.
	TradingSideTransactionReporting []*TradingSideTransactionReporting1 `xml:"TradgSdTxRptg,omitempty"`

	// Regulatory transaction reporting information from the Counterparty Side party.
	CounterpartySideTransactionReporting []*CounterpartySideTransactionReporting1 `xml:"CtrPtySdTxRptg,omitempty"`

	// Identifies an agency or separate corporation of a futures exchange responsible for settling and
	// clearing trades, collecting and maintaining margins, regulating delivery and reporting trade data. This can also be known as a Central Counterparty (CCP).
	CentralCounterpartyClearingHouse *PartyIdentification73Choice `xml:"CntrlCtrPtyClrHs,omitempty"`

	// Identifies the party that is a member of the clearing house (CCP) and that acts as a liaison between the investor and the Cntral Counterparty (CCP).
	ClearingBroker *PartyIdentification73Choice `xml:"ClrBrkr,omitempty"`

	// Identifies the party that is exempt from a clearing obligation.
	ClearingExceptionParty *PartyIdentification73Choice `xml:"ClrXcptnPty,omitempty"`

	// Specifies the reference number assigned by the clearing broker. A distinction can be made between the reference for the Central Counterparty (CCP) leg and the reference for the client leg of the transaction.
	ClearingBrokerIdentification *ClearingBrokerIdentification1 `xml:"ClrBrkrId,omitempty"`

	// Specifies whether the contract is above or below the clearing threshold. Where No indicates the contract is below the clearing threshold and Yes indicates the contract is above the clearing threshold.
	ClearingThresholdIndicator *YesNoIndicator `xml:"ClrThrshldInd,omitempty"`

	// Specifies the reference number assigned by the Central Counterparty (CCP).
	ClearedProductIdentification *Max35Text `xml:"ClrdPdctId,omitempty"`

	// Specifies the underlying product type.
	UnderlyingProductIdentifier *UnderlyingProductIdentifier1Code `xml:"UndrlygPdctIdr,omitempty"`

	// Specifies whether the trade is a pre-allocation or a post-allocation trade, or whether the trade is unallocated.
	AllocationIndicator *AllocationIndicator1Code `xml:"AllcnInd,omitempty"`

	// Specifies whether the transaction is collateralised.
	CollateralisationIndicator *CollateralisationIndicator1Code `xml:"CollstnInd,omitempty"`

	// Specifies the trading venue of the transaction.
	ExecutionVenue *Max35Text `xml:"ExctnVn,omitempty"`

	// Specifies the date and time of the execution of the transaction in Coordinated Universal Time (UTC).
	ExecutionTimestamp *DateAndDateTimeChoice `xml:"ExctnTmstmp,omitempty"`

	// Specifies whether the reportable transaction has one or more additional terms or provisions, other than those listed in the required real-time data fields, that materially affects the price of the reportable transaction.
	NonStandardFlag *YesNoIndicator `xml:"NonStdFlg,omitempty"`

	// Specifies the common reference or correlation identification for a swap transaction where the near and far leg are confirmed separately.
	LinkSwapIdentification *Exact42Text `xml:"LkSwpId,omitempty"`

	// Specifies the financial nature of the reporting counterparty.
	FinancialNatureOfTheCounterpartyIndicator *YesNoIndicator `xml:"FinNtrOfTheCtrPtyInd,omitempty"`

	// Specifies if the collateral is posted on a portfolio basis.
	CollateralPortfolioIndicator *YesNoIndicator `xml:"CollPrtflInd,omitempty"`

	// Identifies the portfolio code to which the trade belongs if the collateral is posted on a portfolio basis (and not trade by trade).
	CollateralPortfolioCode *Max10Text `xml:"CollPrtflCd,omitempty"`

	// Indicates if the trade results from portfolio compression.
	PortfolioCompressionIndicator *YesNoIndicator `xml:"PrtflCmprssnInd,omitempty"`

	// Specifies the corporate sector of the counterparty.
	CorporateSectorIndicator *CorporateSectorIdentifier1Code `xml:"CorpSctrInd,omitempty"`

	// Specifies whether the counterparty has entered into a trade with a non-European Economic Area (EEA) counterparty that is not subject to the reporting obligation.
	TradeWithNonEEACounterpartyIndicator *YesNoIndicator `xml:"TradWthNonEEACtrPtyInd,omitempty"`

	// To indicate if a reported trade falls under the definition of intragroup transaction, as defined by European Securities and Markets Authority (ESMA) in the Technical Standards.
	IntragroupTradeIndicator *YesNoIndicator `xml:"NtrgrpTradInd,omitempty"`

	// Specifies whether the contract is objectively measurable as directly linked to the non-financial counterparty's commercial or treasury financing activity.
	CommercialOrTreasuryFinancingIndicator *YesNoIndicator `xml:"ComrclOrTrsrFincgInd,omitempty"`

	// Identification of a security, as assigned under a formal or proprietary identification scheme.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId,omitempty"`

	// Specifies the date and time of the confirmation of the trade in UTC.
	ConfirmationDateAndTimestamp *ISODateTime `xml:"ConfDtAndTmstmp,omitempty"`

	// Specifies the time at which the CCP has legally taken on the clearing of the trade in UTC.
	ClearingTimestamp *ISOTime `xml:"ClrTmstmp,omitempty"`

	// Specifies additional information that might be required by the regulator.
	AdditionalReportingInformation *Max210Text `xml:"AddtlRptgInf,omitempty"`
}

func (r *RegulatoryReporting6) AddTradingSideTransactionReporting() *TradingSideTransactionReporting1 {
	newValue := new(TradingSideTransactionReporting1)
	r.TradingSideTransactionReporting = append(r.TradingSideTransactionReporting, newValue)
	return newValue
}

func (r *RegulatoryReporting6) AddCounterpartySideTransactionReporting() *CounterpartySideTransactionReporting1 {
	newValue := new(CounterpartySideTransactionReporting1)
	r.CounterpartySideTransactionReporting = append(r.CounterpartySideTransactionReporting, newValue)
	return newValue
}

func (r *RegulatoryReporting6) AddCentralCounterpartyClearingHouse() *PartyIdentification73Choice {
	r.CentralCounterpartyClearingHouse = new(PartyIdentification73Choice)
	return r.CentralCounterpartyClearingHouse
}

func (r *RegulatoryReporting6) AddClearingBroker() *PartyIdentification73Choice {
	r.ClearingBroker = new(PartyIdentification73Choice)
	return r.ClearingBroker
}

func (r *RegulatoryReporting6) AddClearingExceptionParty() *PartyIdentification73Choice {
	r.ClearingExceptionParty = new(PartyIdentification73Choice)
	return r.ClearingExceptionParty
}

func (r *RegulatoryReporting6) AddClearingBrokerIdentification() *ClearingBrokerIdentification1 {
	r.ClearingBrokerIdentification = new(ClearingBrokerIdentification1)
	return r.ClearingBrokerIdentification
}

func (r *RegulatoryReporting6) SetClearingThresholdIndicator(value string) {
	r.ClearingThresholdIndicator = (*YesNoIndicator)(&value)
}

func (r *RegulatoryReporting6) SetClearedProductIdentification(value string) {
	r.ClearedProductIdentification = (*Max35Text)(&value)
}

func (r *RegulatoryReporting6) SetUnderlyingProductIdentifier(value string) {
	r.UnderlyingProductIdentifier = (*UnderlyingProductIdentifier1Code)(&value)
}

func (r *RegulatoryReporting6) SetAllocationIndicator(value string) {
	r.AllocationIndicator = (*AllocationIndicator1Code)(&value)
}

func (r *RegulatoryReporting6) SetCollateralisationIndicator(value string) {
	r.CollateralisationIndicator = (*CollateralisationIndicator1Code)(&value)
}

func (r *RegulatoryReporting6) SetExecutionVenue(value string) {
	r.ExecutionVenue = (*Max35Text)(&value)
}

func (r *RegulatoryReporting6) AddExecutionTimestamp() *DateAndDateTimeChoice {
	r.ExecutionTimestamp = new(DateAndDateTimeChoice)
	return r.ExecutionTimestamp
}

func (r *RegulatoryReporting6) SetNonStandardFlag(value string) {
	r.NonStandardFlag = (*YesNoIndicator)(&value)
}

func (r *RegulatoryReporting6) SetLinkSwapIdentification(value string) {
	r.LinkSwapIdentification = (*Exact42Text)(&value)
}

func (r *RegulatoryReporting6) SetFinancialNatureOfTheCounterpartyIndicator(value string) {
	r.FinancialNatureOfTheCounterpartyIndicator = (*YesNoIndicator)(&value)
}

func (r *RegulatoryReporting6) SetCollateralPortfolioIndicator(value string) {
	r.CollateralPortfolioIndicator = (*YesNoIndicator)(&value)
}

func (r *RegulatoryReporting6) SetCollateralPortfolioCode(value string) {
	r.CollateralPortfolioCode = (*Max10Text)(&value)
}

func (r *RegulatoryReporting6) SetPortfolioCompressionIndicator(value string) {
	r.PortfolioCompressionIndicator = (*YesNoIndicator)(&value)
}

func (r *RegulatoryReporting6) SetCorporateSectorIndicator(value string) {
	r.CorporateSectorIndicator = (*CorporateSectorIdentifier1Code)(&value)
}

func (r *RegulatoryReporting6) SetTradeWithNonEEACounterpartyIndicator(value string) {
	r.TradeWithNonEEACounterpartyIndicator = (*YesNoIndicator)(&value)
}

func (r *RegulatoryReporting6) SetIntragroupTradeIndicator(value string) {
	r.IntragroupTradeIndicator = (*YesNoIndicator)(&value)
}

func (r *RegulatoryReporting6) SetCommercialOrTreasuryFinancingIndicator(value string) {
	r.CommercialOrTreasuryFinancingIndicator = (*YesNoIndicator)(&value)
}

func (r *RegulatoryReporting6) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	r.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return r.FinancialInstrumentIdentification
}

func (r *RegulatoryReporting6) SetConfirmationDateAndTimestamp(value string) {
	r.ConfirmationDateAndTimestamp = (*ISODateTime)(&value)
}

func (r *RegulatoryReporting6) SetClearingTimestamp(value string) {
	r.ClearingTimestamp = (*ISOTime)(&value)
}

func (r *RegulatoryReporting6) SetAdditionalReportingInformation(value string) {
	r.AdditionalReportingInformation = (*Max210Text)(&value)
}
