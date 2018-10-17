package iso20022

// Parameters applied to the settlement of a security transfer.
type Transfer1 struct {

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Date and time at which the securities are to be delivered or received.
	RequestedTransferDate *DateFormat1Choice `xml:"ReqdTrfDt"`

	// Total quantity of securities to be settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Total quantity of securities to be settled.
	PortfolioTransferOutRate *PercentageRate `xml:"PrtflTrfOutRate"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit1 `xml:"UnitsDtls,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`
}

func (t *Transfer1) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer1) AddRequestedTransferDate() *DateFormat1Choice {
	t.RequestedTransferDate = new(DateFormat1Choice)
	return t.RequestedTransferDate
}

func (t *Transfer1) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer1) SetPortfolioTransferOutRate(value string) {
	t.PortfolioTransferOutRate = (*PercentageRate)(&value)
}

func (t *Transfer1) AddUnitsDetails() *Unit1 {
	newValue := new(Unit1)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer1) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer1) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

// Completion of a securities settlement instruction, wherein securities are delivered/debited from a securities account and received/credited to the designated securities account.
type Transfer10 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identifier for a transfer execution, as assigned by a confirming party.
	TransferConfirmationReference *Max35Text `xml:"TrfConfRef"`

	// Reference that identifies the whole transfer out transaction.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Date and time at which the transfer was received and processed.
	EffectiveTransferDate *DateAndDateTimeChoice `xml:"FctvTrfDt"`

	// Date and time at which a transaction is completed and cleared, ie, securities are delivered.
	TradeDate *DateAndDateTimeChoice `xml:"TradDt"`

	// Total quantity of securities settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit3 `xml:"UnitsDtls,omitempty"`

	// Total quantity of securities settled.
	PortfolioTransferOutRate *PercentageRate `xml:"PrtflTrfOutRate,omitempty"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`
}

func (t *Transfer10) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *Transfer10) SetTransferConfirmationReference(value string) {
	t.TransferConfirmationReference = (*Max35Text)(&value)
}

func (t *Transfer10) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer10) SetClientReference(value string) {
	t.ClientReference = (*Max35Text)(&value)
}

func (t *Transfer10) AddEffectiveTransferDate() *DateAndDateTimeChoice {
	t.EffectiveTransferDate = new(DateAndDateTimeChoice)
	return t.EffectiveTransferDate
}

func (t *Transfer10) AddTradeDate() *DateAndDateTimeChoice {
	t.TradeDate = new(DateAndDateTimeChoice)
	return t.TradeDate
}

func (t *Transfer10) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer10) AddUnitsDetails() *Unit3 {
	newValue := new(Unit3)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer10) SetPortfolioTransferOutRate(value string) {
	t.PortfolioTransferOutRate = (*PercentageRate)(&value)
}

func (t *Transfer10) SetRounding(value string) {
	t.Rounding = (*RoundingDirection2Code)(&value)
}

func (t *Transfer10) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer10) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer10) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

// Parameters applied to the settlement of a security transfer.
type Transfer11 struct {

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *AdditionalReference2 `xml:"CtrPtyRef,omitempty"`

	// Identifies the business process in which the actors are involved. This is important to trigger the right business process, according to the market business model, which may require matching instructions in a CSD environment (double leg process) or not (single leg process).
	BusinessFlowType *BusinessFlowType1Code `xml:"BizFlowTp,omitempty"`

	// Identifies in which date the investor signed the transfer order form.
	TransferOrderDateForm *ISODate `xml:"TrfOrdrDtForm,omitempty"`

	// Identifies the transfer reason.
	TransferReason *TransferReason1 `xml:"TrfRsn,omitempty"`

	// Identifies whether or not saving plan or withdrawal or switch plan are included in the holdings.
	HoldingsPlanType []*HoldingsPlanType1Code `xml:"HldgsPlanTp,omitempty"`

	// Information related to the financial instrument to be withdrawn.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Total quantity of securities to be transferred, expressed in a number of units or a percentage rate.
	Quantity *Quantity13Choice `xml:"Qty"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit3 `xml:"UnitsDtls,omitempty"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`

	// Identifies the currency to be used to transfer the holdings.
	TransferCurrency *CurrencyCode `xml:"TrfCcy,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`
}

func (t *Transfer11) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer11) SetClientReference(value string) {
	t.ClientReference = (*Max35Text)(&value)
}

func (t *Transfer11) AddCounterpartyReference() *AdditionalReference2 {
	t.CounterpartyReference = new(AdditionalReference2)
	return t.CounterpartyReference
}

func (t *Transfer11) SetBusinessFlowType(value string) {
	t.BusinessFlowType = (*BusinessFlowType1Code)(&value)
}

func (t *Transfer11) SetTransferOrderDateForm(value string) {
	t.TransferOrderDateForm = (*ISODate)(&value)
}

func (t *Transfer11) AddTransferReason() *TransferReason1 {
	t.TransferReason = new(TransferReason1)
	return t.TransferReason
}

func (t *Transfer11) AddHoldingsPlanType(value string) {
	t.HoldingsPlanType = append(t.HoldingsPlanType, (*HoldingsPlanType1Code)(&value))
}

func (t *Transfer11) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *Transfer11) AddQuantity() *Quantity13Choice {
	t.Quantity = new(Quantity13Choice)
	return t.Quantity
}

func (t *Transfer11) AddUnitsDetails() *Unit3 {
	newValue := new(Unit3)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer11) SetRounding(value string) {
	t.Rounding = (*RoundingDirection2Code)(&value)
}

func (t *Transfer11) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer11) SetTransferCurrency(value string) {
	t.TransferCurrency = (*CurrencyCode)(&value)
}

func (t *Transfer11) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer11) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

// Parameters applied to the settlement of a security transfer.
type Transfer12 struct {

	// Identifies whether or not saving plan or withdrawal or switch plan are included in the holdings.
	HoldingsPlanType []*HoldingsPlanType1Code `xml:"HldgsPlanTp,omitempty"`

	// Information related to the financial instrument to be withdrawn.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Total quantity of securities to be transferred, expressed in a number of units or a percentage rate.
	Quantity *Quantity13Choice `xml:"Qty"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit3 `xml:"UnitsDtls,omitempty"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`
}

func (t *Transfer12) AddHoldingsPlanType(value string) {
	t.HoldingsPlanType = append(t.HoldingsPlanType, (*HoldingsPlanType1Code)(&value))
}

func (t *Transfer12) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *Transfer12) AddQuantity() *Quantity13Choice {
	t.Quantity = new(Quantity13Choice)
	return t.Quantity
}

func (t *Transfer12) AddUnitsDetails() *Unit3 {
	newValue := new(Unit3)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer12) SetRounding(value string) {
	t.Rounding = (*RoundingDirection2Code)(&value)
}

func (t *Transfer12) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer12) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer12) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

// Completion of a securities settlement instruction, wherein securities are delivered/debited from a securities account and received/credited to the designated securities account.
type Transfer13 struct {

	// Unique and unambiguous identifier for a transfer execution, as assigned by a confirming party.
	TransferConfirmationReference *Max35Text `xml:"TrfConfRef"`

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *AdditionalReference2 `xml:"CtrPtyRef,omitempty"`

	// Date and time at which the transfer was received and processed.
	EffectiveTransferDate *DateAndDateTimeChoice `xml:"FctvTrfDt"`

	// Date and time at which a transaction is completed and cleared, ie, securities are delivered.
	TradeDate *DateAndDateTimeChoice `xml:"TradDt,omitempty"`

	// Identifies in which date the investor signed the transfer order form.
	TransferOrderDateForm *ISODate `xml:"TrfOrdrDtForm,omitempty"`

	// Identifies the transfer reason.
	TransferReason *TransferReason1 `xml:"TrfRsn,omitempty"`

	// Identifies whether or not saving plan or withdrawal or switch plan are included in the holdings.
	HoldingsPlanType []*HoldingsPlanType1Code `xml:"HldgsPlanTp,omitempty"`

	// Information related to the financial instrument withdrawn.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Total quantity of securities settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit3 `xml:"UnitsDtls,omitempty"`

	// Total quantity of securities settled.
	PortfolioTransferOutRate *PercentageRate `xml:"PrtflTrfOutRate,omitempty"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`

	// Identifies the currency to be used to transfer the holdings.
	TransferCurrency *CurrencyCode `xml:"TrfCcy,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`
}

func (t *Transfer13) SetTransferConfirmationReference(value string) {
	t.TransferConfirmationReference = (*Max35Text)(&value)
}

func (t *Transfer13) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer13) SetClientReference(value string) {
	t.ClientReference = (*Max35Text)(&value)
}

func (t *Transfer13) AddCounterpartyReference() *AdditionalReference2 {
	t.CounterpartyReference = new(AdditionalReference2)
	return t.CounterpartyReference
}

func (t *Transfer13) AddEffectiveTransferDate() *DateAndDateTimeChoice {
	t.EffectiveTransferDate = new(DateAndDateTimeChoice)
	return t.EffectiveTransferDate
}

func (t *Transfer13) AddTradeDate() *DateAndDateTimeChoice {
	t.TradeDate = new(DateAndDateTimeChoice)
	return t.TradeDate
}

func (t *Transfer13) SetTransferOrderDateForm(value string) {
	t.TransferOrderDateForm = (*ISODate)(&value)
}

func (t *Transfer13) AddTransferReason() *TransferReason1 {
	t.TransferReason = new(TransferReason1)
	return t.TransferReason
}

func (t *Transfer13) AddHoldingsPlanType(value string) {
	t.HoldingsPlanType = append(t.HoldingsPlanType, (*HoldingsPlanType1Code)(&value))
}

func (t *Transfer13) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *Transfer13) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer13) AddUnitsDetails() *Unit3 {
	newValue := new(Unit3)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer13) SetPortfolioTransferOutRate(value string) {
	t.PortfolioTransferOutRate = (*PercentageRate)(&value)
}

func (t *Transfer13) SetRounding(value string) {
	t.Rounding = (*RoundingDirection2Code)(&value)
}

func (t *Transfer13) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer13) SetTransferCurrency(value string) {
	t.TransferCurrency = (*CurrencyCode)(&value)
}

func (t *Transfer13) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer13) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

// Completion of a securities settlement instruction, wherein securities are delivered/debited from a securities account and received/credited to the designated securities account.
type Transfer14 struct {

	// Date and time at which the transfer was received and processed.
	EffectiveTransferDate *DateAndDateTimeChoice `xml:"FctvTrfDt"`

	// Date and time at which a transaction is completed and cleared, ie, securities are delivered.
	TradeDate *DateAndDateTimeChoice `xml:"TradDt,omitempty"`

	// Identifies whether or not saving plan or withdrawal or switch plan are included in the holdings.
	HoldingsPlanType []*HoldingsPlanType1Code `xml:"HldgsPlanTp,omitempty"`

	// Information related to the financial instrument withdrawn.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Total quantity of securities settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit3 `xml:"UnitsDtls,omitempty"`

	// Total quantity of securities settled.
	PortfolioTransferOutRate *PercentageRate `xml:"PrtflTrfOutRate,omitempty"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`
}

func (t *Transfer14) AddEffectiveTransferDate() *DateAndDateTimeChoice {
	t.EffectiveTransferDate = new(DateAndDateTimeChoice)
	return t.EffectiveTransferDate
}

func (t *Transfer14) AddTradeDate() *DateAndDateTimeChoice {
	t.TradeDate = new(DateAndDateTimeChoice)
	return t.TradeDate
}

func (t *Transfer14) AddHoldingsPlanType(value string) {
	t.HoldingsPlanType = append(t.HoldingsPlanType, (*HoldingsPlanType1Code)(&value))
}

func (t *Transfer14) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *Transfer14) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer14) AddUnitsDetails() *Unit3 {
	newValue := new(Unit3)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer14) SetPortfolioTransferOutRate(value string) {
	t.PortfolioTransferOutRate = (*PercentageRate)(&value)
}

func (t *Transfer14) SetRounding(value string) {
	t.Rounding = (*RoundingDirection2Code)(&value)
}

func (t *Transfer14) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer14) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer14) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

// Parameters applied to the settlement of a security transfer.
type Transfer15 struct {

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *AdditionalReference2 `xml:"CtrPtyRef,omitempty"`

	// Identifies the business process in which the actors are involved. This is important to trigger the right business process, according to the market business model, which may require matching instructions in a CSD environment (double leg process) or not (single leg process).
	BusinessFlowType *BusinessFlowType1Code `xml:"BizFlowTp,omitempty"`

	// Identifies the transfer reason.
	TransferReason *TransferReason1 `xml:"TrfRsn,omitempty"`

	// Date at which the instructing party places the transfer instruction.
	TransferDate *DateFormat1Choice `xml:"TrfDt,omitempty"`

	// Identifies whether or not saving plan or withdrawal or switch plan are included in the holdings.
	HoldingsPlanType []*HoldingsPlanType1Code `xml:"HldgsPlanTp,omitempty"`

	// Information related to the financial instrument to be received.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Total quantity of securities to be settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`
}

func (t *Transfer15) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer15) SetClientReference(value string) {
	t.ClientReference = (*Max35Text)(&value)
}

func (t *Transfer15) AddCounterpartyReference() *AdditionalReference2 {
	t.CounterpartyReference = new(AdditionalReference2)
	return t.CounterpartyReference
}

func (t *Transfer15) SetBusinessFlowType(value string) {
	t.BusinessFlowType = (*BusinessFlowType1Code)(&value)
}

func (t *Transfer15) AddTransferReason() *TransferReason1 {
	t.TransferReason = new(TransferReason1)
	return t.TransferReason
}

func (t *Transfer15) AddTransferDate() *DateFormat1Choice {
	t.TransferDate = new(DateFormat1Choice)
	return t.TransferDate
}

func (t *Transfer15) AddHoldingsPlanType(value string) {
	t.HoldingsPlanType = append(t.HoldingsPlanType, (*HoldingsPlanType1Code)(&value))
}

func (t *Transfer15) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *Transfer15) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer15) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer15) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

// Parameters applied to the settlement of a security transfer.
type Transfer16 struct {

	// Date at which the instructing party places the transfer instruction.
	TransferDate *DateFormat1Choice `xml:"TrfDt,omitempty"`

	// Identifies whether or not saving plan or withdrawal or switch plan are included in the holdings.
	HoldingsPlanType []*HoldingsPlanType1Code `xml:"HldgsPlanTp,omitempty"`

	// Information related to the financial instrument to be received.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Total quantity of securities to be settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`
}

func (t *Transfer16) AddTransferDate() *DateFormat1Choice {
	t.TransferDate = new(DateFormat1Choice)
	return t.TransferDate
}

func (t *Transfer16) AddHoldingsPlanType(value string) {
	t.HoldingsPlanType = append(t.HoldingsPlanType, (*HoldingsPlanType1Code)(&value))
}

func (t *Transfer16) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *Transfer16) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer16) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer16) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

// Completion of a securities settlement instruction, wherein securities are delivered/debited from a securities account and received/credited to the designated securities account.
type Transfer17 struct {

	// Date and time at which the transfer was executed.
	EffectiveTransferDate *DateAndDateTimeChoice `xml:"FctvTrfDt"`

	// Date and time at which a transaction is completed and cleared, ie, securities are delivered.
	TradeDate *DateAndDateTimeChoice `xml:"TradDt,omitempty"`

	// Identifies whether or not saving plan or withdrawal or switch plan are included in the holdings.
	HoldingsPlanType []*HoldingsPlanType1Code `xml:"HldgsPlanTp,omitempty"`

	// Information related to the financial instrument received.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Total quantity of securities settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit3 `xml:"UnitsDtls,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`
}

func (t *Transfer17) AddEffectiveTransferDate() *DateAndDateTimeChoice {
	t.EffectiveTransferDate = new(DateAndDateTimeChoice)
	return t.EffectiveTransferDate
}

func (t *Transfer17) AddTradeDate() *DateAndDateTimeChoice {
	t.TradeDate = new(DateAndDateTimeChoice)
	return t.TradeDate
}

func (t *Transfer17) AddHoldingsPlanType(value string) {
	t.HoldingsPlanType = append(t.HoldingsPlanType, (*HoldingsPlanType1Code)(&value))
}

func (t *Transfer17) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *Transfer17) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer17) AddUnitsDetails() *Unit3 {
	newValue := new(Unit3)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer17) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer17) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer17) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

// Completion of a securities settlement instruction, wherein securities are delivered/debited from a securities account and received/credited to the designated securities account.
type Transfer18 struct {

	// Unique and unambiguous identifier for a transfer execution, as assigned by a confirming party.
	TransferConfirmationReference *Max35Text `xml:"TrfConfRef"`

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *AdditionalReference2 `xml:"CtrPtyRef,omitempty"`

	// Date and time at which the transfer was executed.
	EffectiveTransferDate *DateAndDateTimeChoice `xml:"FctvTrfDt"`

	// Date and time at which a transaction is completed and cleared, ie, securities are delivered.
	TradeDate *DateAndDateTimeChoice `xml:"TradDt,omitempty"`

	// Identifies in which date the investor signed the transfer order form.
	TransferOrderDateForm *ISODate `xml:"TrfOrdrDtForm,omitempty"`

	// Specifies the reason for the assets transfer.
	TransferReason *TransferReason1 `xml:"TrfRsn,omitempty"`

	// Identifies whether or not saving plan or withdrawal or switch plan are included in the holdings.
	HoldingsPlanType []*HoldingsPlanType1Code `xml:"HldgsPlanTp,omitempty"`

	// Information related to the financial instrument received.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Total quantity of securities settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit3 `xml:"UnitsDtls,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`

	// Identifies the currency to be used to transfer the holdings.
	TransferCurrency *CurrencyCode `xml:"TrfCcy,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`
}

func (t *Transfer18) SetTransferConfirmationReference(value string) {
	t.TransferConfirmationReference = (*Max35Text)(&value)
}

func (t *Transfer18) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer18) SetClientReference(value string) {
	t.ClientReference = (*Max35Text)(&value)
}

func (t *Transfer18) AddCounterpartyReference() *AdditionalReference2 {
	t.CounterpartyReference = new(AdditionalReference2)
	return t.CounterpartyReference
}

func (t *Transfer18) AddEffectiveTransferDate() *DateAndDateTimeChoice {
	t.EffectiveTransferDate = new(DateAndDateTimeChoice)
	return t.EffectiveTransferDate
}

func (t *Transfer18) AddTradeDate() *DateAndDateTimeChoice {
	t.TradeDate = new(DateAndDateTimeChoice)
	return t.TradeDate
}

func (t *Transfer18) SetTransferOrderDateForm(value string) {
	t.TransferOrderDateForm = (*ISODate)(&value)
}

func (t *Transfer18) AddTransferReason() *TransferReason1 {
	t.TransferReason = new(TransferReason1)
	return t.TransferReason
}

func (t *Transfer18) AddHoldingsPlanType(value string) {
	t.HoldingsPlanType = append(t.HoldingsPlanType, (*HoldingsPlanType1Code)(&value))
}

func (t *Transfer18) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *Transfer18) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer18) AddUnitsDetails() *Unit3 {
	newValue := new(Unit3)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer18) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer18) SetTransferCurrency(value string) {
	t.TransferCurrency = (*CurrencyCode)(&value)
}

func (t *Transfer18) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer18) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

// Parameters applied to the settlement of a security transfer.
type Transfer19 struct {

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *AdditionalReference2 `xml:"CtrPtyRef,omitempty"`

	// Identifies the business process in which the actors are involved. This is important to trigger the right business process, according to the market business model, which may require matching instructions in a CSD environment (double leg process) or not (single leg process).
	BusinessFlowType *BusinessFlowType1Code `xml:"BizFlowTp,omitempty"`

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Identifies in which date the investor signed the transfer order form.
	TransferOrderDateForm *ISODate `xml:"TrfOrdrDtForm,omitempty"`

	// Identifies the transfer reason.
	TransferReason *TransferReason1 `xml:"TrfRsn,omitempty"`

	// Identifies whether or not saving plan or withdrawal or switch plan are included in the holdings.
	HoldingsPlanType []*HoldingsPlanType1Code `xml:"HldgsPlanTp,omitempty"`

	// Information related to the financial instrument to be withdrawn.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Total quantity of securities to be transferred, expressed in a number of units or a percentage rate.
	Quantity *Quantity13Choice `xml:"Qty"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit3 `xml:"UnitsDtls,omitempty"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`

	// Identifies the currency to be used to transfer the holdings.
	TransferCurrency *CurrencyCode `xml:"TrfCcy,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount93 `xml:"RcvgAgtDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount93 `xml:"DlvrgAgtDtls,omitempty"`
}

func (t *Transfer19) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer19) SetClientReference(value string) {
	t.ClientReference = (*Max35Text)(&value)
}

func (t *Transfer19) AddCounterpartyReference() *AdditionalReference2 {
	t.CounterpartyReference = new(AdditionalReference2)
	return t.CounterpartyReference
}

func (t *Transfer19) SetBusinessFlowType(value string) {
	t.BusinessFlowType = (*BusinessFlowType1Code)(&value)
}

func (t *Transfer19) SetRequestedSettlementDate(value string) {
	t.RequestedSettlementDate = (*ISODate)(&value)
}

func (t *Transfer19) SetTransferOrderDateForm(value string) {
	t.TransferOrderDateForm = (*ISODate)(&value)
}

func (t *Transfer19) AddTransferReason() *TransferReason1 {
	t.TransferReason = new(TransferReason1)
	return t.TransferReason
}

func (t *Transfer19) AddHoldingsPlanType(value string) {
	t.HoldingsPlanType = append(t.HoldingsPlanType, (*HoldingsPlanType1Code)(&value))
}

func (t *Transfer19) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *Transfer19) AddQuantity() *Quantity13Choice {
	t.Quantity = new(Quantity13Choice)
	return t.Quantity
}

func (t *Transfer19) AddUnitsDetails() *Unit3 {
	newValue := new(Unit3)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer19) SetRounding(value string) {
	t.Rounding = (*RoundingDirection2Code)(&value)
}

func (t *Transfer19) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer19) SetTransferCurrency(value string) {
	t.TransferCurrency = (*CurrencyCode)(&value)
}

func (t *Transfer19) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer19) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (t *Transfer19) AddReceivingAgentDetails() *PartyIdentificationAndAccount93 {
	t.ReceivingAgentDetails = new(PartyIdentificationAndAccount93)
	return t.ReceivingAgentDetails
}

func (t *Transfer19) AddDeliveringAgentDetails() *PartyIdentificationAndAccount93 {
	t.DeliveringAgentDetails = new(PartyIdentificationAndAccount93)
	return t.DeliveringAgentDetails
}

// Completion of a securities settlement instruction, wherein securities are delivered/debited from a securities account and received/credited to the designated securities account.
type Transfer2 struct {

	// Unique and unambiguous identifier for a transfer execution, as assigned by a confirming party.
	TransferConfirmationReference *Max35Text `xml:"TrfConfRef"`

	// Reference that identifies the whole transfer out transaction.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Date and optionally  time at which a transaction is completed and cleared, ie, securities are delivered.
	EffectiveTransferDate *DateAndDateTimeChoice `xml:"FctvTrfDt"`

	// Date when the transfer was received and processed.
	TradeDate *ISODate `xml:"TradDt"`

	// Total quantity of securities settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit1 `xml:"UnitsDtls,omitempty"`

	// Total quantity of securities settled.
	PortfolioTransferOutRate *PercentageRate `xml:"PrtflTrfOutRate,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`
}

func (t *Transfer2) SetTransferConfirmationReference(value string) {
	t.TransferConfirmationReference = (*Max35Text)(&value)
}

func (t *Transfer2) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer2) AddEffectiveTransferDate() *DateAndDateTimeChoice {
	t.EffectiveTransferDate = new(DateAndDateTimeChoice)
	return t.EffectiveTransferDate
}

func (t *Transfer2) SetTradeDate(value string) {
	t.TradeDate = (*ISODate)(&value)
}

func (t *Transfer2) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer2) AddUnitsDetails() *Unit1 {
	newValue := new(Unit1)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer2) SetPortfolioTransferOutRate(value string) {
	t.PortfolioTransferOutRate = (*PercentageRate)(&value)
}

func (t *Transfer2) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer2) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

// Parameters applied to the settlement of a security transfer.
type Transfer20 struct {

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Identifies whether or not saving plan or withdrawal or switch plan are included in the holdings.
	HoldingsPlanType []*HoldingsPlanType1Code `xml:"HldgsPlanTp,omitempty"`

	// Information related to the financial instrument to be withdrawn.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Total quantity of securities to be transferred, expressed in a number of units or a percentage rate.
	Quantity *Quantity13Choice `xml:"Qty"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit3 `xml:"UnitsDtls,omitempty"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount93 `xml:"RcvgAgtDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount93 `xml:"DlvrgAgtDtls,omitempty"`
}

func (t *Transfer20) SetRequestedSettlementDate(value string) {
	t.RequestedSettlementDate = (*ISODate)(&value)
}

func (t *Transfer20) AddHoldingsPlanType(value string) {
	t.HoldingsPlanType = append(t.HoldingsPlanType, (*HoldingsPlanType1Code)(&value))
}

func (t *Transfer20) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *Transfer20) AddQuantity() *Quantity13Choice {
	t.Quantity = new(Quantity13Choice)
	return t.Quantity
}

func (t *Transfer20) AddUnitsDetails() *Unit3 {
	newValue := new(Unit3)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer20) SetRounding(value string) {
	t.Rounding = (*RoundingDirection2Code)(&value)
}

func (t *Transfer20) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer20) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer20) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (t *Transfer20) AddReceivingAgentDetails() *PartyIdentificationAndAccount93 {
	t.ReceivingAgentDetails = new(PartyIdentificationAndAccount93)
	return t.ReceivingAgentDetails
}

func (t *Transfer20) AddDeliveringAgentDetails() *PartyIdentificationAndAccount93 {
	t.DeliveringAgentDetails = new(PartyIdentificationAndAccount93)
	return t.DeliveringAgentDetails
}

// Parameters applied to the settlement of a security transfer.
type Transfer21 struct {

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *AdditionalReference2 `xml:"CtrPtyRef,omitempty"`

	// Identifies the business process in which the actors are involved. This is important to trigger the right business process, according to the market business model, which may require matching instructions in a CSD environment (double leg process) or not (single leg process).
	BusinessFlowType *BusinessFlowType1Code `xml:"BizFlowTp,omitempty"`

	// Identifies the transfer reason.
	TransferReason *TransferReason1 `xml:"TrfRsn,omitempty"`

	// Date at which the instructing party places the transfer instruction.
	TransferDate *DateFormat1Choice `xml:"TrfDt,omitempty"`

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Identifies whether or not saving plan or withdrawal or switch plan are included in the holdings.
	HoldingsPlanType []*HoldingsPlanType1Code `xml:"HldgsPlanTp,omitempty"`

	// Information related to the financial instrument to be received.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Total quantity of securities to be settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount93 `xml:"RcvgAgtDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount93 `xml:"DlvrgAgtDtls,omitempty"`
}

func (t *Transfer21) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer21) SetClientReference(value string) {
	t.ClientReference = (*Max35Text)(&value)
}

func (t *Transfer21) AddCounterpartyReference() *AdditionalReference2 {
	t.CounterpartyReference = new(AdditionalReference2)
	return t.CounterpartyReference
}

func (t *Transfer21) SetBusinessFlowType(value string) {
	t.BusinessFlowType = (*BusinessFlowType1Code)(&value)
}

func (t *Transfer21) AddTransferReason() *TransferReason1 {
	t.TransferReason = new(TransferReason1)
	return t.TransferReason
}

func (t *Transfer21) AddTransferDate() *DateFormat1Choice {
	t.TransferDate = new(DateFormat1Choice)
	return t.TransferDate
}

func (t *Transfer21) SetRequestedSettlementDate(value string) {
	t.RequestedSettlementDate = (*ISODate)(&value)
}

func (t *Transfer21) AddHoldingsPlanType(value string) {
	t.HoldingsPlanType = append(t.HoldingsPlanType, (*HoldingsPlanType1Code)(&value))
}

func (t *Transfer21) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *Transfer21) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer21) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer21) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (t *Transfer21) AddReceivingAgentDetails() *PartyIdentificationAndAccount93 {
	t.ReceivingAgentDetails = new(PartyIdentificationAndAccount93)
	return t.ReceivingAgentDetails
}

func (t *Transfer21) AddDeliveringAgentDetails() *PartyIdentificationAndAccount93 {
	t.DeliveringAgentDetails = new(PartyIdentificationAndAccount93)
	return t.DeliveringAgentDetails
}

// Parameters applied to the settlement of a security transfer.
type Transfer22 struct {

	// Date at which the instructing party places the transfer instruction.
	TransferDate *DateFormat1Choice `xml:"TrfDt,omitempty"`

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Identifies whether or not saving plan or withdrawal or switch plan are included in the holdings.
	HoldingsPlanType []*HoldingsPlanType1Code `xml:"HldgsPlanTp,omitempty"`

	// Information related to the financial instrument to be received.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Total quantity of securities to be settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount93 `xml:"RcvgAgtDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount93 `xml:"DlvrgAgtDtls,omitempty"`
}

func (t *Transfer22) AddTransferDate() *DateFormat1Choice {
	t.TransferDate = new(DateFormat1Choice)
	return t.TransferDate
}

func (t *Transfer22) SetRequestedSettlementDate(value string) {
	t.RequestedSettlementDate = (*ISODate)(&value)
}

func (t *Transfer22) AddHoldingsPlanType(value string) {
	t.HoldingsPlanType = append(t.HoldingsPlanType, (*HoldingsPlanType1Code)(&value))
}

func (t *Transfer22) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *Transfer22) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer22) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer22) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (t *Transfer22) AddReceivingAgentDetails() *PartyIdentificationAndAccount93 {
	t.ReceivingAgentDetails = new(PartyIdentificationAndAccount93)
	return t.ReceivingAgentDetails
}

func (t *Transfer22) AddDeliveringAgentDetails() *PartyIdentificationAndAccount93 {
	t.DeliveringAgentDetails = new(PartyIdentificationAndAccount93)
	return t.DeliveringAgentDetails
}

// Completion of a securities settlement instruction, wherein securities are delivered/debited from a securities account and received/credited to the designated securities account.
type Transfer23 struct {

	// Unique and unambiguous identifier for a transfer execution, as assigned by a confirming party.
	TransferConfirmationReference *Max35Text `xml:"TrfConfRef"`

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *AdditionalReference2 `xml:"CtrPtyRef,omitempty"`

	// Date and time at which the transfer was executed.
	EffectiveTransferDate *DateAndDateTimeChoice `xml:"FctvTrfDt"`

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Date and time at which the securities were exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, ie, securities are delivered.
	TradeDate *DateAndDateTimeChoice `xml:"TradDt,omitempty"`

	// Identifies in which date the investor signed the transfer order form.
	TransferOrderDateForm *ISODate `xml:"TrfOrdrDtForm,omitempty"`

	// Identifies the transfer reason.
	TransferReason *TransferReason1 `xml:"TrfRsn,omitempty"`

	// Identifies whether or not saving plan or withdrawal or switch plan are included in the holdings.
	HoldingsPlanType []*HoldingsPlanType1Code `xml:"HldgsPlanTp,omitempty"`

	// Information related to the financial instrument withdrawn.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Total quantity of securities settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit3 `xml:"UnitsDtls,omitempty"`

	// Total quantity of securities settled.
	PortfolioTransferOutRate *PercentageRate `xml:"PrtflTrfOutRate,omitempty"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`

	// Identifies the currency to be used to transfer the holdings.
	TransferCurrency *CurrencyCode `xml:"TrfCcy,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount93 `xml:"RcvgAgtDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount93 `xml:"DlvrgAgtDtls,omitempty"`
}

func (t *Transfer23) SetTransferConfirmationReference(value string) {
	t.TransferConfirmationReference = (*Max35Text)(&value)
}

func (t *Transfer23) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer23) SetClientReference(value string) {
	t.ClientReference = (*Max35Text)(&value)
}

func (t *Transfer23) AddCounterpartyReference() *AdditionalReference2 {
	t.CounterpartyReference = new(AdditionalReference2)
	return t.CounterpartyReference
}

func (t *Transfer23) AddEffectiveTransferDate() *DateAndDateTimeChoice {
	t.EffectiveTransferDate = new(DateAndDateTimeChoice)
	return t.EffectiveTransferDate
}

func (t *Transfer23) SetRequestedSettlementDate(value string) {
	t.RequestedSettlementDate = (*ISODate)(&value)
}

func (t *Transfer23) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	t.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return t.EffectiveSettlementDate
}

func (t *Transfer23) AddTradeDate() *DateAndDateTimeChoice {
	t.TradeDate = new(DateAndDateTimeChoice)
	return t.TradeDate
}

func (t *Transfer23) SetTransferOrderDateForm(value string) {
	t.TransferOrderDateForm = (*ISODate)(&value)
}

func (t *Transfer23) AddTransferReason() *TransferReason1 {
	t.TransferReason = new(TransferReason1)
	return t.TransferReason
}

func (t *Transfer23) AddHoldingsPlanType(value string) {
	t.HoldingsPlanType = append(t.HoldingsPlanType, (*HoldingsPlanType1Code)(&value))
}

func (t *Transfer23) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *Transfer23) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer23) AddUnitsDetails() *Unit3 {
	newValue := new(Unit3)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer23) SetPortfolioTransferOutRate(value string) {
	t.PortfolioTransferOutRate = (*PercentageRate)(&value)
}

func (t *Transfer23) SetRounding(value string) {
	t.Rounding = (*RoundingDirection2Code)(&value)
}

func (t *Transfer23) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer23) SetTransferCurrency(value string) {
	t.TransferCurrency = (*CurrencyCode)(&value)
}

func (t *Transfer23) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer23) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (t *Transfer23) AddReceivingAgentDetails() *PartyIdentificationAndAccount93 {
	t.ReceivingAgentDetails = new(PartyIdentificationAndAccount93)
	return t.ReceivingAgentDetails
}

func (t *Transfer23) AddDeliveringAgentDetails() *PartyIdentificationAndAccount93 {
	t.DeliveringAgentDetails = new(PartyIdentificationAndAccount93)
	return t.DeliveringAgentDetails
}

// Completion of a securities settlement instruction, wherein securities are delivered/debited from a securities account and received/credited to the designated securities account.
type Transfer24 struct {

	// Date and time at which the transfer was executed.
	EffectiveTransferDate *DateAndDateTimeChoice `xml:"FctvTrfDt"`

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Date and time at which the securities were exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, ie, securities are delivered.
	TradeDate *DateAndDateTimeChoice `xml:"TradDt,omitempty"`

	// Identifies whether or not saving plan or withdrawal or switch plan are included in the holdings.
	HoldingsPlanType []*HoldingsPlanType1Code `xml:"HldgsPlanTp,omitempty"`

	// Information related to the financial instrument withdrawn.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Total quantity of securities settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit3 `xml:"UnitsDtls,omitempty"`

	// Total quantity of securities settled.
	PortfolioTransferOutRate *PercentageRate `xml:"PrtflTrfOutRate,omitempty"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount93 `xml:"RcvgAgtDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount93 `xml:"DlvrgAgtDtls,omitempty"`
}

func (t *Transfer24) AddEffectiveTransferDate() *DateAndDateTimeChoice {
	t.EffectiveTransferDate = new(DateAndDateTimeChoice)
	return t.EffectiveTransferDate
}

func (t *Transfer24) SetRequestedSettlementDate(value string) {
	t.RequestedSettlementDate = (*ISODate)(&value)
}

func (t *Transfer24) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	t.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return t.EffectiveSettlementDate
}

func (t *Transfer24) AddTradeDate() *DateAndDateTimeChoice {
	t.TradeDate = new(DateAndDateTimeChoice)
	return t.TradeDate
}

func (t *Transfer24) AddHoldingsPlanType(value string) {
	t.HoldingsPlanType = append(t.HoldingsPlanType, (*HoldingsPlanType1Code)(&value))
}

func (t *Transfer24) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *Transfer24) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer24) AddUnitsDetails() *Unit3 {
	newValue := new(Unit3)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer24) SetPortfolioTransferOutRate(value string) {
	t.PortfolioTransferOutRate = (*PercentageRate)(&value)
}

func (t *Transfer24) SetRounding(value string) {
	t.Rounding = (*RoundingDirection2Code)(&value)
}

func (t *Transfer24) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer24) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer24) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (t *Transfer24) AddReceivingAgentDetails() *PartyIdentificationAndAccount93 {
	t.ReceivingAgentDetails = new(PartyIdentificationAndAccount93)
	return t.ReceivingAgentDetails
}

func (t *Transfer24) AddDeliveringAgentDetails() *PartyIdentificationAndAccount93 {
	t.DeliveringAgentDetails = new(PartyIdentificationAndAccount93)
	return t.DeliveringAgentDetails
}

// Completion of a securities settlement instruction, wherein securities are delivered/debited from a securities account and received/credited to the designated securities account.
type Transfer25 struct {

	// Date and time at which the transfer was executed.
	EffectiveTransferDate *DateAndDateTimeChoice `xml:"FctvTrfDt"`

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Date and time at which the securities were exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, ie, securities are delivered.
	TradeDate *DateAndDateTimeChoice `xml:"TradDt,omitempty"`

	// Identifies whether or not saving plan or withdrawal or switch plan are included in the holdings.
	HoldingsPlanType []*HoldingsPlanType1Code `xml:"HldgsPlanTp,omitempty"`

	// Information related to the financial instrument received.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Total quantity of securities settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit3 `xml:"UnitsDtls,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount93 `xml:"RcvgAgtDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount93 `xml:"DlvrgAgtDtls,omitempty"`
}

func (t *Transfer25) AddEffectiveTransferDate() *DateAndDateTimeChoice {
	t.EffectiveTransferDate = new(DateAndDateTimeChoice)
	return t.EffectiveTransferDate
}

func (t *Transfer25) SetRequestedSettlementDate(value string) {
	t.RequestedSettlementDate = (*ISODate)(&value)
}

func (t *Transfer25) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	t.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return t.EffectiveSettlementDate
}

func (t *Transfer25) AddTradeDate() *DateAndDateTimeChoice {
	t.TradeDate = new(DateAndDateTimeChoice)
	return t.TradeDate
}

func (t *Transfer25) AddHoldingsPlanType(value string) {
	t.HoldingsPlanType = append(t.HoldingsPlanType, (*HoldingsPlanType1Code)(&value))
}

func (t *Transfer25) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *Transfer25) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer25) AddUnitsDetails() *Unit3 {
	newValue := new(Unit3)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer25) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer25) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer25) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (t *Transfer25) AddReceivingAgentDetails() *PartyIdentificationAndAccount93 {
	t.ReceivingAgentDetails = new(PartyIdentificationAndAccount93)
	return t.ReceivingAgentDetails
}

func (t *Transfer25) AddDeliveringAgentDetails() *PartyIdentificationAndAccount93 {
	t.DeliveringAgentDetails = new(PartyIdentificationAndAccount93)
	return t.DeliveringAgentDetails
}

// Completion of a securities settlement instruction, wherein securities are delivered/debited from a securities account and received/credited to the designated securities account.
type Transfer26 struct {

	// Unique and unambiguous identifier for a transfer execution, as assigned by a confirming party.
	TransferConfirmationReference *Max35Text `xml:"TrfConfRef"`

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *AdditionalReference2 `xml:"CtrPtyRef,omitempty"`

	// Date and time at which the transfer was executed.
	EffectiveTransferDate *DateAndDateTimeChoice `xml:"FctvTrfDt"`

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Date and time at which the securities were exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, ie, securities are delivered.
	TradeDate *DateAndDateTimeChoice `xml:"TradDt,omitempty"`

	// Identifies in which date the investor signed the transfer order form.
	TransferOrderDateForm *ISODate `xml:"TrfOrdrDtForm,omitempty"`

	// Specifies the reason for the assets transfer.
	TransferReason *TransferReason1 `xml:"TrfRsn,omitempty"`

	// Identifies whether or not saving plan or withdrawal or switch plan are included in the holdings.
	HoldingsPlanType []*HoldingsPlanType1Code `xml:"HldgsPlanTp,omitempty"`

	// Information related to the financial instrument received.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Total quantity of securities settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit3 `xml:"UnitsDtls,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`

	// Identifies the currency to be used to transfer the holdings.
	TransferCurrency *CurrencyCode `xml:"TrfCcy,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount93 `xml:"RcvgAgtDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount93 `xml:"DlvrgAgtDtls,omitempty"`
}

func (t *Transfer26) SetTransferConfirmationReference(value string) {
	t.TransferConfirmationReference = (*Max35Text)(&value)
}

func (t *Transfer26) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer26) SetClientReference(value string) {
	t.ClientReference = (*Max35Text)(&value)
}

func (t *Transfer26) AddCounterpartyReference() *AdditionalReference2 {
	t.CounterpartyReference = new(AdditionalReference2)
	return t.CounterpartyReference
}

func (t *Transfer26) AddEffectiveTransferDate() *DateAndDateTimeChoice {
	t.EffectiveTransferDate = new(DateAndDateTimeChoice)
	return t.EffectiveTransferDate
}

func (t *Transfer26) SetRequestedSettlementDate(value string) {
	t.RequestedSettlementDate = (*ISODate)(&value)
}

func (t *Transfer26) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	t.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return t.EffectiveSettlementDate
}

func (t *Transfer26) AddTradeDate() *DateAndDateTimeChoice {
	t.TradeDate = new(DateAndDateTimeChoice)
	return t.TradeDate
}

func (t *Transfer26) SetTransferOrderDateForm(value string) {
	t.TransferOrderDateForm = (*ISODate)(&value)
}

func (t *Transfer26) AddTransferReason() *TransferReason1 {
	t.TransferReason = new(TransferReason1)
	return t.TransferReason
}

func (t *Transfer26) AddHoldingsPlanType(value string) {
	t.HoldingsPlanType = append(t.HoldingsPlanType, (*HoldingsPlanType1Code)(&value))
}

func (t *Transfer26) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *Transfer26) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer26) AddUnitsDetails() *Unit3 {
	newValue := new(Unit3)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer26) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer26) SetTransferCurrency(value string) {
	t.TransferCurrency = (*CurrencyCode)(&value)
}

func (t *Transfer26) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer26) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (t *Transfer26) AddReceivingAgentDetails() *PartyIdentificationAndAccount93 {
	t.ReceivingAgentDetails = new(PartyIdentificationAndAccount93)
	return t.ReceivingAgentDetails
}

func (t *Transfer26) AddDeliveringAgentDetails() *PartyIdentificationAndAccount93 {
	t.DeliveringAgentDetails = new(PartyIdentificationAndAccount93)
	return t.DeliveringAgentDetails
}

// Parameters applied to the settlement of a security transfer.
type Transfer27 struct {

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *AdditionalReference2 `xml:"CtrPtyRef,omitempty"`

	// Identifies the business process in which the actors are involved. This is important to trigger the right business process, according to the market business model, which may require matching instructions in a CSD environment (double leg process) or not (single leg process).
	BusinessFlowType *BusinessFlowType1Code `xml:"BizFlowTp,omitempty"`

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Identifies in which date the investor signed the transfer order form.
	TransferOrderDateForm *ISODate `xml:"TrfOrdrDtForm,omitempty"`

	// Identifies the transfer reason.
	TransferReason *TransferReason1 `xml:"TrfRsn,omitempty"`

	// Identifies whether or not saving plan or withdrawal or switch plan are included in the holdings.
	HoldingsPlanType []*HoldingsPlanType1Code `xml:"HldgsPlanTp,omitempty"`

	// Information related to the financial instrument to be withdrawn.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Total quantity of securities to be transferred, expressed in a number of units or a percentage rate.
	Quantity *Quantity13Choice `xml:"Qty"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit3 `xml:"UnitsDtls,omitempty"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`

	// Identifies the currency to be used to transfer the holdings.
	TransferCurrency *CurrencyCode `xml:"TrfCcy,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount93 `xml:"RcvgAgtDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount93 `xml:"DlvrgAgtDtls,omitempty"`

	// Specifies how the payment of charges, taxes and commissions as a result of the transfer is covered, that is, whether by cash or the redemption of units.
	TransferExpensesPaymentType *ChargePaymentMethod1Choice `xml:"TrfExpnssPmtTp,omitempty"`
}

func (t *Transfer27) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer27) SetClientReference(value string) {
	t.ClientReference = (*Max35Text)(&value)
}

func (t *Transfer27) AddCounterpartyReference() *AdditionalReference2 {
	t.CounterpartyReference = new(AdditionalReference2)
	return t.CounterpartyReference
}

func (t *Transfer27) SetBusinessFlowType(value string) {
	t.BusinessFlowType = (*BusinessFlowType1Code)(&value)
}

func (t *Transfer27) SetRequestedSettlementDate(value string) {
	t.RequestedSettlementDate = (*ISODate)(&value)
}

func (t *Transfer27) SetTransferOrderDateForm(value string) {
	t.TransferOrderDateForm = (*ISODate)(&value)
}

func (t *Transfer27) AddTransferReason() *TransferReason1 {
	t.TransferReason = new(TransferReason1)
	return t.TransferReason
}

func (t *Transfer27) AddHoldingsPlanType(value string) {
	t.HoldingsPlanType = append(t.HoldingsPlanType, (*HoldingsPlanType1Code)(&value))
}

func (t *Transfer27) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *Transfer27) AddQuantity() *Quantity13Choice {
	t.Quantity = new(Quantity13Choice)
	return t.Quantity
}

func (t *Transfer27) AddUnitsDetails() *Unit3 {
	newValue := new(Unit3)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer27) SetRounding(value string) {
	t.Rounding = (*RoundingDirection2Code)(&value)
}

func (t *Transfer27) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer27) SetTransferCurrency(value string) {
	t.TransferCurrency = (*CurrencyCode)(&value)
}

func (t *Transfer27) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer27) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (t *Transfer27) AddReceivingAgentDetails() *PartyIdentificationAndAccount93 {
	t.ReceivingAgentDetails = new(PartyIdentificationAndAccount93)
	return t.ReceivingAgentDetails
}

func (t *Transfer27) AddDeliveringAgentDetails() *PartyIdentificationAndAccount93 {
	t.DeliveringAgentDetails = new(PartyIdentificationAndAccount93)
	return t.DeliveringAgentDetails
}

func (t *Transfer27) AddTransferExpensesPaymentType() *ChargePaymentMethod1Choice {
	t.TransferExpensesPaymentType = new(ChargePaymentMethod1Choice)
	return t.TransferExpensesPaymentType
}

// Parameters applied to the settlement of a security transfer.
type Transfer28 struct {

	// Unique and unambiguous identifier for a transfer execution, as assigned by a confirming party.
	TransferConfirmationReference *Max35Text `xml:"TrfConfRef"`

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *AdditionalReference2 `xml:"CtrPtyRef,omitempty"`

	// Identifies the business process in which the actors are involved. This is important to trigger the right business process, according to the market business model, which may require matching instructions in a CSD environment (double leg process) or not (single leg process).
	BusinessFlowType *BusinessFlowType1Code `xml:"BizFlowTp,omitempty"`

	// Date and time at which the transfer was executed.
	EffectiveTransferDate *DateAndDateTimeChoice `xml:"FctvTrfDt"`

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Date and time at which the securities were exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, ie, securities are delivered.
	TradeDate *DateAndDateTimeChoice `xml:"TradDt,omitempty"`

	// Identifies in which date the investor signed the transfer order form.
	TransferOrderDateForm *ISODate `xml:"TrfOrdrDtForm,omitempty"`

	// Identifies the transfer reason.
	TransferReason *TransferReason1 `xml:"TrfRsn,omitempty"`

	// Identifies whether or not saving plan or withdrawal or switch plan are included in the holdings.
	HoldingsPlanType []*HoldingsPlanType1Code `xml:"HldgsPlanTp,omitempty"`

	// Information related to the financial instrument withdrawn.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Total quantity of securities settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit3 `xml:"UnitsDtls,omitempty"`

	// Total quantity of securities settled.
	PortfolioTransferOutRate *PercentageRate `xml:"PrtflTrfOutRate,omitempty"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Average price of the units in the account before the transfer was executed.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`

	// Average price of the units in the account after the transfer was executed.
	NewAveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"NewAvrgPric,omitempty"`

	// Identifies the currency to be used to transfer the holdings.
	TransferCurrency *CurrencyCode `xml:"TrfCcy,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount93 `xml:"RcvgAgtDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount93 `xml:"DlvrgAgtDtls,omitempty"`

	// Specifies how the payment of charges, taxes and commissions as a result of the transfer is covered, that is, whether by cash or the redemption of units.
	TransferExpensesPaymentType *ChargePaymentMethod1Choice `xml:"TrfExpnssPmtTp,omitempty"`
}

func (t *Transfer28) SetTransferConfirmationReference(value string) {
	t.TransferConfirmationReference = (*Max35Text)(&value)
}

func (t *Transfer28) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer28) SetClientReference(value string) {
	t.ClientReference = (*Max35Text)(&value)
}

func (t *Transfer28) AddCounterpartyReference() *AdditionalReference2 {
	t.CounterpartyReference = new(AdditionalReference2)
	return t.CounterpartyReference
}

func (t *Transfer28) SetBusinessFlowType(value string) {
	t.BusinessFlowType = (*BusinessFlowType1Code)(&value)
}

func (t *Transfer28) AddEffectiveTransferDate() *DateAndDateTimeChoice {
	t.EffectiveTransferDate = new(DateAndDateTimeChoice)
	return t.EffectiveTransferDate
}

func (t *Transfer28) SetRequestedSettlementDate(value string) {
	t.RequestedSettlementDate = (*ISODate)(&value)
}

func (t *Transfer28) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	t.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return t.EffectiveSettlementDate
}

func (t *Transfer28) AddTradeDate() *DateAndDateTimeChoice {
	t.TradeDate = new(DateAndDateTimeChoice)
	return t.TradeDate
}

func (t *Transfer28) SetTransferOrderDateForm(value string) {
	t.TransferOrderDateForm = (*ISODate)(&value)
}

func (t *Transfer28) AddTransferReason() *TransferReason1 {
	t.TransferReason = new(TransferReason1)
	return t.TransferReason
}

func (t *Transfer28) AddHoldingsPlanType(value string) {
	t.HoldingsPlanType = append(t.HoldingsPlanType, (*HoldingsPlanType1Code)(&value))
}

func (t *Transfer28) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *Transfer28) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer28) AddUnitsDetails() *Unit3 {
	newValue := new(Unit3)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer28) SetPortfolioTransferOutRate(value string) {
	t.PortfolioTransferOutRate = (*PercentageRate)(&value)
}

func (t *Transfer28) SetRounding(value string) {
	t.Rounding = (*RoundingDirection2Code)(&value)
}

func (t *Transfer28) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer28) SetNewAveragePrice(value, currency string) {
	t.NewAveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer28) SetTransferCurrency(value string) {
	t.TransferCurrency = (*CurrencyCode)(&value)
}

func (t *Transfer28) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer28) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (t *Transfer28) AddReceivingAgentDetails() *PartyIdentificationAndAccount93 {
	t.ReceivingAgentDetails = new(PartyIdentificationAndAccount93)
	return t.ReceivingAgentDetails
}

func (t *Transfer28) AddDeliveringAgentDetails() *PartyIdentificationAndAccount93 {
	t.DeliveringAgentDetails = new(PartyIdentificationAndAccount93)
	return t.DeliveringAgentDetails
}

func (t *Transfer28) AddTransferExpensesPaymentType() *ChargePaymentMethod1Choice {
	t.TransferExpensesPaymentType = new(ChargePaymentMethod1Choice)
	return t.TransferExpensesPaymentType
}

// Parameters applied to the settlement of a security transfer.
type Transfer29 struct {

	// Unique and unambiguous identifier for a transfer execution, as assigned by a confirming party.
	TransferConfirmationReference *Max35Text `xml:"TrfConfRef"`

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *AdditionalReference2 `xml:"CtrPtyRef,omitempty"`

	// Identifies the business process in which the actors are involved. This is important to trigger the right business process, according to the market business model, which may require matching instructions in a CSD environment (double leg process) or not (single leg process).
	BusinessFlowType *BusinessFlowType1Code `xml:"BizFlowTp,omitempty"`

	// Date and time at which the transfer was executed.
	EffectiveTransferDate *DateAndDateTimeChoice `xml:"FctvTrfDt"`

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Date and time at which the securities were exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, ie, securities are delivered.
	TradeDate *DateAndDateTimeChoice `xml:"TradDt,omitempty"`

	// Identifies in which date the investor signed the transfer order form.
	TransferOrderDateForm *ISODate `xml:"TrfOrdrDtForm,omitempty"`

	// Specifies the reason for the assets transfer.
	TransferReason *TransferReason1 `xml:"TrfRsn,omitempty"`

	// Identifies whether or not saving plan or withdrawal or switch plan are included in the holdings.
	HoldingsPlanType []*HoldingsPlanType1Code `xml:"HldgsPlanTp,omitempty"`

	// Information related to the financial instrument received.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Total quantity of securities settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit3 `xml:"UnitsDtls,omitempty"`

	// Average price of the units in the account before the transfer was executed.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`

	// Average price of the units in the account after the transfer was executed.
	NewAveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"NewAvrgPric,omitempty"`

	// Identifies the currency to be used to transfer the holdings.
	TransferCurrency *CurrencyCode `xml:"TrfCcy,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount93 `xml:"RcvgAgtDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount93 `xml:"DlvrgAgtDtls,omitempty"`

	// Specifies how the payment of charges, taxes and commissions as a result of the transfer is covered, that is, whether by cash or the redemption of units.
	TransferExpensesPaymentType *ChargePaymentMethod1Choice `xml:"TrfExpnssPmtTp,omitempty"`
}

func (t *Transfer29) SetTransferConfirmationReference(value string) {
	t.TransferConfirmationReference = (*Max35Text)(&value)
}

func (t *Transfer29) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer29) SetClientReference(value string) {
	t.ClientReference = (*Max35Text)(&value)
}

func (t *Transfer29) AddCounterpartyReference() *AdditionalReference2 {
	t.CounterpartyReference = new(AdditionalReference2)
	return t.CounterpartyReference
}

func (t *Transfer29) SetBusinessFlowType(value string) {
	t.BusinessFlowType = (*BusinessFlowType1Code)(&value)
}

func (t *Transfer29) AddEffectiveTransferDate() *DateAndDateTimeChoice {
	t.EffectiveTransferDate = new(DateAndDateTimeChoice)
	return t.EffectiveTransferDate
}

func (t *Transfer29) SetRequestedSettlementDate(value string) {
	t.RequestedSettlementDate = (*ISODate)(&value)
}

func (t *Transfer29) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	t.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return t.EffectiveSettlementDate
}

func (t *Transfer29) AddTradeDate() *DateAndDateTimeChoice {
	t.TradeDate = new(DateAndDateTimeChoice)
	return t.TradeDate
}

func (t *Transfer29) SetTransferOrderDateForm(value string) {
	t.TransferOrderDateForm = (*ISODate)(&value)
}

func (t *Transfer29) AddTransferReason() *TransferReason1 {
	t.TransferReason = new(TransferReason1)
	return t.TransferReason
}

func (t *Transfer29) AddHoldingsPlanType(value string) {
	t.HoldingsPlanType = append(t.HoldingsPlanType, (*HoldingsPlanType1Code)(&value))
}

func (t *Transfer29) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *Transfer29) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer29) AddUnitsDetails() *Unit3 {
	newValue := new(Unit3)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer29) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer29) SetNewAveragePrice(value, currency string) {
	t.NewAveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer29) SetTransferCurrency(value string) {
	t.TransferCurrency = (*CurrencyCode)(&value)
}

func (t *Transfer29) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer29) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (t *Transfer29) AddReceivingAgentDetails() *PartyIdentificationAndAccount93 {
	t.ReceivingAgentDetails = new(PartyIdentificationAndAccount93)
	return t.ReceivingAgentDetails
}

func (t *Transfer29) AddDeliveringAgentDetails() *PartyIdentificationAndAccount93 {
	t.DeliveringAgentDetails = new(PartyIdentificationAndAccount93)
	return t.DeliveringAgentDetails
}

func (t *Transfer29) AddTransferExpensesPaymentType() *ChargePaymentMethod1Choice {
	t.TransferExpensesPaymentType = new(ChargePaymentMethod1Choice)
	return t.TransferExpensesPaymentType
}

// Parameters applied to the settlement of a security transfer.
type Transfer3 struct {

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Date and time at which the securities are to be delivered or received.
	TransferDate *DateFormat1Choice `xml:"TrfDt,omitempty"`

	// Total quantity of securities to be settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd"`
}

func (t *Transfer3) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer3) AddTransferDate() *DateFormat1Choice {
	t.TransferDate = new(DateFormat1Choice)
	return t.TransferDate
}

func (t *Transfer3) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer3) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

// Parameters applied to the settlement of a security transfer.
type Transfer30 struct {

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *AdditionalReference7 `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *AdditionalReference7 `xml:"CtrPtyRef,omitempty"`

	// Identifies the business process in which the actors are involved. This is important to trigger the right business process, according to the market business model, which may require matching instructions in a CSD environment (double leg process) or not (single leg process).
	BusinessFlowType *BusinessFlowType1Code `xml:"BizFlowTp,omitempty"`

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Identifies in which date the investor signed the transfer order form.
	TransferOrderDateForm *ISODate `xml:"TrfOrdrDtForm,omitempty"`

	// Identifies the transfer reason.
	TransferReason *TransferReason1 `xml:"TrfRsn,omitempty"`

	// Identifies whether or not saving plan or withdrawal or switch plan are included in the holdings.
	HoldingsPlanType []*HoldingsPlanType1Code `xml:"HldgsPlanTp,omitempty"`

	// Information related to the financial instrument to be withdrawn.
	FinancialInstrumentDetails *FinancialInstrument49 `xml:"FinInstrmDtls"`

	// Total quantity of securities to be transferred, expressed in a number of units or a percentage rate.
	Quantity *Quantity13Choice `xml:"Qty"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit6 `xml:"UnitsDtls,omitempty"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`

	// Identifies the currency to be used to transfer the holdings.
	TransferCurrency *ActiveOrHistoricCurrencyCode `xml:"TrfCcy,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount125 `xml:"RcvgAgtDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount125 `xml:"DlvrgAgtDtls,omitempty"`

	// Specifies how the payment of charges, taxes and commissions as a result of the transfer is covered, that is, whether by cash or the redemption of units.
	TransferExpensesPaymentType *ChargePaymentMethod1Choice `xml:"TrfExpnssPmtTp,omitempty"`
}

func (t *Transfer30) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer30) AddClientReference() *AdditionalReference7 {
	t.ClientReference = new(AdditionalReference7)
	return t.ClientReference
}

func (t *Transfer30) AddCounterpartyReference() *AdditionalReference7 {
	t.CounterpartyReference = new(AdditionalReference7)
	return t.CounterpartyReference
}

func (t *Transfer30) SetBusinessFlowType(value string) {
	t.BusinessFlowType = (*BusinessFlowType1Code)(&value)
}

func (t *Transfer30) SetRequestedSettlementDate(value string) {
	t.RequestedSettlementDate = (*ISODate)(&value)
}

func (t *Transfer30) SetTransferOrderDateForm(value string) {
	t.TransferOrderDateForm = (*ISODate)(&value)
}

func (t *Transfer30) AddTransferReason() *TransferReason1 {
	t.TransferReason = new(TransferReason1)
	return t.TransferReason
}

func (t *Transfer30) AddHoldingsPlanType(value string) {
	t.HoldingsPlanType = append(t.HoldingsPlanType, (*HoldingsPlanType1Code)(&value))
}

func (t *Transfer30) AddFinancialInstrumentDetails() *FinancialInstrument49 {
	t.FinancialInstrumentDetails = new(FinancialInstrument49)
	return t.FinancialInstrumentDetails
}

func (t *Transfer30) AddQuantity() *Quantity13Choice {
	t.Quantity = new(Quantity13Choice)
	return t.Quantity
}

func (t *Transfer30) AddUnitsDetails() *Unit6 {
	newValue := new(Unit6)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer30) SetRounding(value string) {
	t.Rounding = (*RoundingDirection2Code)(&value)
}

func (t *Transfer30) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer30) SetTransferCurrency(value string) {
	t.TransferCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (t *Transfer30) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer30) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (t *Transfer30) AddReceivingAgentDetails() *PartyIdentificationAndAccount125 {
	t.ReceivingAgentDetails = new(PartyIdentificationAndAccount125)
	return t.ReceivingAgentDetails
}

func (t *Transfer30) AddDeliveringAgentDetails() *PartyIdentificationAndAccount125 {
	t.DeliveringAgentDetails = new(PartyIdentificationAndAccount125)
	return t.DeliveringAgentDetails
}

func (t *Transfer30) AddTransferExpensesPaymentType() *ChargePaymentMethod1Choice {
	t.TransferExpensesPaymentType = new(ChargePaymentMethod1Choice)
	return t.TransferExpensesPaymentType
}

// Parameters applied to the settlement of a security transfer.
type Transfer31 struct {

	// Unique and unambiguous identifier for a transfer execution, as assigned by a confirming party.
	TransferConfirmationReference *Max35Text `xml:"TrfConfRef"`

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *AdditionalReference7 `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *AdditionalReference7 `xml:"CtrPtyRef,omitempty"`

	// Identifies the business process in which the actors are involved. This is important to trigger the right business process, according to the market business model, which may require matching instructions in a CSD environment (double leg process) or not (single leg process).
	BusinessFlowType *BusinessFlowType1Code `xml:"BizFlowTp,omitempty"`

	// Date and time at which the transfer was executed.
	EffectiveTransferDate *DateAndDateTimeChoice `xml:"FctvTrfDt"`

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Date and time at which the securities were exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, that is, securities are delivered.
	TradeDate *DateAndDateTimeChoice `xml:"TradDt,omitempty"`

	// Identifies in which date the investor signed the transfer order form.
	TransferOrderDateForm *ISODate `xml:"TrfOrdrDtForm,omitempty"`

	// Identifies the transfer reason.
	TransferReason *TransferReason1 `xml:"TrfRsn,omitempty"`

	// Identifies whether or not saving plan or withdrawal or switch plan are included in the holdings.
	HoldingsPlanType []*HoldingsPlanType1Code `xml:"HldgsPlanTp,omitempty"`

	// Information related to the financial instrument withdrawn.
	FinancialInstrumentDetails *FinancialInstrument49 `xml:"FinInstrmDtls"`

	// Total quantity of securities settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit6 `xml:"UnitsDtls,omitempty"`

	// Total quantity of securities settled.
	PortfolioTransferOutRate *PercentageRate `xml:"PrtflTrfOutRate,omitempty"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Weighted average price of the units in the account before the transfer was executed.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`

	// Weighted average price of the units in the account after the transfer was executed.
	NewAveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"NewAvrgPric,omitempty"`

	// Trade date of the average weighted data of units in the account before the transfer was executed.
	AverageDate *ISODate `xml:"AvrgDt,omitempty"`

	// Trade date of the average weighted data of units in the account after the transfer was executed.
	NewAverageDate *ISODate `xml:"NewAvrgDt,omitempty"`

	// Identifies the currency to be used to transfer the holdings.
	TransferCurrency *ActiveOrHistoricCurrencyCode `xml:"TrfCcy,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount125 `xml:"RcvgAgtDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount125 `xml:"DlvrgAgtDtls,omitempty"`

	// Specifies how the payment of charges, taxes and commissions as a result of the transfer is covered, that is, whether by cash or the redemption of units.
	TransferExpensesPaymentType *ChargePaymentMethod1Choice `xml:"TrfExpnssPmtTp,omitempty"`
}

func (t *Transfer31) SetTransferConfirmationReference(value string) {
	t.TransferConfirmationReference = (*Max35Text)(&value)
}

func (t *Transfer31) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer31) AddClientReference() *AdditionalReference7 {
	t.ClientReference = new(AdditionalReference7)
	return t.ClientReference
}

func (t *Transfer31) AddCounterpartyReference() *AdditionalReference7 {
	t.CounterpartyReference = new(AdditionalReference7)
	return t.CounterpartyReference
}

func (t *Transfer31) SetBusinessFlowType(value string) {
	t.BusinessFlowType = (*BusinessFlowType1Code)(&value)
}

func (t *Transfer31) AddEffectiveTransferDate() *DateAndDateTimeChoice {
	t.EffectiveTransferDate = new(DateAndDateTimeChoice)
	return t.EffectiveTransferDate
}

func (t *Transfer31) SetRequestedSettlementDate(value string) {
	t.RequestedSettlementDate = (*ISODate)(&value)
}

func (t *Transfer31) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	t.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return t.EffectiveSettlementDate
}

func (t *Transfer31) AddTradeDate() *DateAndDateTimeChoice {
	t.TradeDate = new(DateAndDateTimeChoice)
	return t.TradeDate
}

func (t *Transfer31) SetTransferOrderDateForm(value string) {
	t.TransferOrderDateForm = (*ISODate)(&value)
}

func (t *Transfer31) AddTransferReason() *TransferReason1 {
	t.TransferReason = new(TransferReason1)
	return t.TransferReason
}

func (t *Transfer31) AddHoldingsPlanType(value string) {
	t.HoldingsPlanType = append(t.HoldingsPlanType, (*HoldingsPlanType1Code)(&value))
}

func (t *Transfer31) AddFinancialInstrumentDetails() *FinancialInstrument49 {
	t.FinancialInstrumentDetails = new(FinancialInstrument49)
	return t.FinancialInstrumentDetails
}

func (t *Transfer31) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer31) AddUnitsDetails() *Unit6 {
	newValue := new(Unit6)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer31) SetPortfolioTransferOutRate(value string) {
	t.PortfolioTransferOutRate = (*PercentageRate)(&value)
}

func (t *Transfer31) SetRounding(value string) {
	t.Rounding = (*RoundingDirection2Code)(&value)
}

func (t *Transfer31) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer31) SetNewAveragePrice(value, currency string) {
	t.NewAveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer31) SetAverageDate(value string) {
	t.AverageDate = (*ISODate)(&value)
}

func (t *Transfer31) SetNewAverageDate(value string) {
	t.NewAverageDate = (*ISODate)(&value)
}

func (t *Transfer31) SetTransferCurrency(value string) {
	t.TransferCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (t *Transfer31) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer31) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (t *Transfer31) AddReceivingAgentDetails() *PartyIdentificationAndAccount125 {
	t.ReceivingAgentDetails = new(PartyIdentificationAndAccount125)
	return t.ReceivingAgentDetails
}

func (t *Transfer31) AddDeliveringAgentDetails() *PartyIdentificationAndAccount125 {
	t.DeliveringAgentDetails = new(PartyIdentificationAndAccount125)
	return t.DeliveringAgentDetails
}

func (t *Transfer31) AddTransferExpensesPaymentType() *ChargePaymentMethod1Choice {
	t.TransferExpensesPaymentType = new(ChargePaymentMethod1Choice)
	return t.TransferExpensesPaymentType
}

// Parameters applied to the settlement of a security transfer.
type Transfer32 struct {

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *AdditionalReference7 `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *AdditionalReference7 `xml:"CtrPtyRef,omitempty"`

	// Identifies the business process in which the actors are involved. This is important to trigger the right business process, according to the market business model, which may require matching instructions in a CSD environment (double leg process) or not (single leg process).
	BusinessFlowType *BusinessFlowType1Code `xml:"BizFlowTp,omitempty"`

	// Identifies the transfer reason.
	TransferReason *TransferReason1 `xml:"TrfRsn,omitempty"`

	// Date at which the instructing party places the transfer instruction.
	TransferDate *DateFormat1Choice `xml:"TrfDt,omitempty"`

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Identifies whether or not saving plan or withdrawal or switch plan are included in the holdings.
	HoldingsPlanType []*HoldingsPlanType1Code `xml:"HldgsPlanTp,omitempty"`

	// Information related to the financial instrument to be received.
	FinancialInstrumentDetails *FinancialInstrument49 `xml:"FinInstrmDtls"`

	// Total quantity of securities to be settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount125 `xml:"RcvgAgtDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount125 `xml:"DlvrgAgtDtls,omitempty"`
}

func (t *Transfer32) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer32) AddClientReference() *AdditionalReference7 {
	t.ClientReference = new(AdditionalReference7)
	return t.ClientReference
}

func (t *Transfer32) AddCounterpartyReference() *AdditionalReference7 {
	t.CounterpartyReference = new(AdditionalReference7)
	return t.CounterpartyReference
}

func (t *Transfer32) SetBusinessFlowType(value string) {
	t.BusinessFlowType = (*BusinessFlowType1Code)(&value)
}

func (t *Transfer32) AddTransferReason() *TransferReason1 {
	t.TransferReason = new(TransferReason1)
	return t.TransferReason
}

func (t *Transfer32) AddTransferDate() *DateFormat1Choice {
	t.TransferDate = new(DateFormat1Choice)
	return t.TransferDate
}

func (t *Transfer32) SetRequestedSettlementDate(value string) {
	t.RequestedSettlementDate = (*ISODate)(&value)
}

func (t *Transfer32) AddHoldingsPlanType(value string) {
	t.HoldingsPlanType = append(t.HoldingsPlanType, (*HoldingsPlanType1Code)(&value))
}

func (t *Transfer32) AddFinancialInstrumentDetails() *FinancialInstrument49 {
	t.FinancialInstrumentDetails = new(FinancialInstrument49)
	return t.FinancialInstrumentDetails
}

func (t *Transfer32) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer32) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer32) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (t *Transfer32) AddReceivingAgentDetails() *PartyIdentificationAndAccount125 {
	t.ReceivingAgentDetails = new(PartyIdentificationAndAccount125)
	return t.ReceivingAgentDetails
}

func (t *Transfer32) AddDeliveringAgentDetails() *PartyIdentificationAndAccount125 {
	t.DeliveringAgentDetails = new(PartyIdentificationAndAccount125)
	return t.DeliveringAgentDetails
}

// Parameters applied to the settlement of a security transfer.
type Transfer33 struct {

	// Unique and unambiguous identifier for a transfer execution, as assigned by a confirming party.
	TransferConfirmationReference *Max35Text `xml:"TrfConfRef"`

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *AdditionalReference7 `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *AdditionalReference7 `xml:"CtrPtyRef,omitempty"`

	// Identifies the business process in which the actors are involved. This is important to trigger the right business process, according to the market business model, which may require matching instructions in a CSD environment (double leg process) or not (single leg process).
	BusinessFlowType *BusinessFlowType1Code `xml:"BizFlowTp,omitempty"`

	// Date and time at which the transfer was executed.
	EffectiveTransferDate *DateAndDateTimeChoice `xml:"FctvTrfDt,omitempty"`

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Date and time at which the securities were exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, that is, securities are delivered.
	TradeDate *DateAndDateTimeChoice `xml:"TradDt,omitempty"`

	// Identifies in which date the investor signed the transfer order form.
	TransferOrderDateForm *ISODate `xml:"TrfOrdrDtForm,omitempty"`

	// Specifies the reason for the assets transfer.
	TransferReason *TransferReason1 `xml:"TrfRsn,omitempty"`

	// Identifies whether or not saving plan or withdrawal or switch plan are included in the holdings.
	HoldingsPlanType []*HoldingsPlanType1Code `xml:"HldgsPlanTp,omitempty"`

	// Information related to the financial instrument received.
	FinancialInstrumentDetails *FinancialInstrument49 `xml:"FinInstrmDtls"`

	// Total quantity of securities settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit6 `xml:"UnitsDtls,omitempty"`

	// Weighted average price of the units in the account before the transfer was executed.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`

	// Weighted average price of the units in the account after the transfer was executed.
	NewAveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"NewAvrgPric,omitempty"`

	// Trade date of the average weighted data of units in the account before the transfer was executed.
	AverageDate *ISODate `xml:"AvrgDt,omitempty"`

	// Trade date of the average weighted data of units in the account after the transfer was executed.
	NewAverageDate *ISODate `xml:"NewAvrgDt,omitempty"`

	// Identifies the currency to be used to transfer the holdings.
	TransferCurrency *ActiveOrHistoricCurrencyCode `xml:"TrfCcy,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount125 `xml:"RcvgAgtDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount125 `xml:"DlvrgAgtDtls,omitempty"`

	// Specifies how the payment of charges, taxes and commissions as a result of the transfer is covered, that is, whether by cash or the redemption of units.
	TransferExpensesPaymentType *ChargePaymentMethod1Choice `xml:"TrfExpnssPmtTp,omitempty"`
}

func (t *Transfer33) SetTransferConfirmationReference(value string) {
	t.TransferConfirmationReference = (*Max35Text)(&value)
}

func (t *Transfer33) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer33) AddClientReference() *AdditionalReference7 {
	t.ClientReference = new(AdditionalReference7)
	return t.ClientReference
}

func (t *Transfer33) AddCounterpartyReference() *AdditionalReference7 {
	t.CounterpartyReference = new(AdditionalReference7)
	return t.CounterpartyReference
}

func (t *Transfer33) SetBusinessFlowType(value string) {
	t.BusinessFlowType = (*BusinessFlowType1Code)(&value)
}

func (t *Transfer33) AddEffectiveTransferDate() *DateAndDateTimeChoice {
	t.EffectiveTransferDate = new(DateAndDateTimeChoice)
	return t.EffectiveTransferDate
}

func (t *Transfer33) SetRequestedSettlementDate(value string) {
	t.RequestedSettlementDate = (*ISODate)(&value)
}

func (t *Transfer33) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	t.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return t.EffectiveSettlementDate
}

func (t *Transfer33) AddTradeDate() *DateAndDateTimeChoice {
	t.TradeDate = new(DateAndDateTimeChoice)
	return t.TradeDate
}

func (t *Transfer33) SetTransferOrderDateForm(value string) {
	t.TransferOrderDateForm = (*ISODate)(&value)
}

func (t *Transfer33) AddTransferReason() *TransferReason1 {
	t.TransferReason = new(TransferReason1)
	return t.TransferReason
}

func (t *Transfer33) AddHoldingsPlanType(value string) {
	t.HoldingsPlanType = append(t.HoldingsPlanType, (*HoldingsPlanType1Code)(&value))
}

func (t *Transfer33) AddFinancialInstrumentDetails() *FinancialInstrument49 {
	t.FinancialInstrumentDetails = new(FinancialInstrument49)
	return t.FinancialInstrumentDetails
}

func (t *Transfer33) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer33) AddUnitsDetails() *Unit6 {
	newValue := new(Unit6)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer33) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer33) SetNewAveragePrice(value, currency string) {
	t.NewAveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer33) SetAverageDate(value string) {
	t.AverageDate = (*ISODate)(&value)
}

func (t *Transfer33) SetNewAverageDate(value string) {
	t.NewAverageDate = (*ISODate)(&value)
}

func (t *Transfer33) SetTransferCurrency(value string) {
	t.TransferCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (t *Transfer33) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer33) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (t *Transfer33) AddReceivingAgentDetails() *PartyIdentificationAndAccount125 {
	t.ReceivingAgentDetails = new(PartyIdentificationAndAccount125)
	return t.ReceivingAgentDetails
}

func (t *Transfer33) AddDeliveringAgentDetails() *PartyIdentificationAndAccount125 {
	t.DeliveringAgentDetails = new(PartyIdentificationAndAccount125)
	return t.DeliveringAgentDetails
}

func (t *Transfer33) AddTransferExpensesPaymentType() *ChargePaymentMethod1Choice {
	t.TransferExpensesPaymentType = new(ChargePaymentMethod1Choice)
	return t.TransferExpensesPaymentType
}

// Completion of a securities settlement instruction, wherein securities are delivered/debited from a securities account and received/credited to the designated securities account.
type Transfer4 struct {

	// Unique and unambiguous identifier for a transfer execution, as assigned by a confirming party.
	TransferConfirmationReference *Max35Text `xml:"TrfConfRef"`

	// Reference that identifies the transfer in transaction.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Date and optionally time at which a transaction is completed and cleared, ie, securities are delivered.
	EffectiveTransferDate *DateAndDateTimeChoice `xml:"FctvTrfDt"`

	// Date when the transfer was received and processed.
	TradeDate *ISODate `xml:"TradDt"`

	// Total quantity of securities settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit1 `xml:"UnitsDtls,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`
}

func (t *Transfer4) SetTransferConfirmationReference(value string) {
	t.TransferConfirmationReference = (*Max35Text)(&value)
}

func (t *Transfer4) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer4) AddEffectiveTransferDate() *DateAndDateTimeChoice {
	t.EffectiveTransferDate = new(DateAndDateTimeChoice)
	return t.EffectiveTransferDate
}

func (t *Transfer4) SetTradeDate(value string) {
	t.TradeDate = (*ISODate)(&value)
}

func (t *Transfer4) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer4) AddUnitsDetails() *Unit1 {
	newValue := new(Unit1)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer4) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer4) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

// Parameters applied to the settlement of a security transfer.
type Transfer5 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Date at which the instructing party places the transfer instruction.
	TransferDate *DateFormat1Choice `xml:"TrfDt,omitempty"`

	// Total quantity of securities to be settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`
}

func (t *Transfer5) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *Transfer5) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer5) SetClientReference(value string) {
	t.ClientReference = (*Max35Text)(&value)
}

func (t *Transfer5) AddTransferDate() *DateFormat1Choice {
	t.TransferDate = new(DateFormat1Choice)
	return t.TransferDate
}

func (t *Transfer5) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer5) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer5) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

// Parameters applied to the settlement of a security transfer.
type Transfer6 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Date at which the instructing party places the transfer instruction.
	TransferDate *DateFormat1Choice `xml:"TrfDt,omitempty"`

	// Total quantity of securities to be settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`
}

func (t *Transfer6) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *Transfer6) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer6) SetClientReference(value string) {
	t.ClientReference = (*Max35Text)(&value)
}

func (t *Transfer6) SetCancellationReference(value string) {
	t.CancellationReference = (*Max35Text)(&value)
}

func (t *Transfer6) AddTransferDate() *DateFormat1Choice {
	t.TransferDate = new(DateFormat1Choice)
	return t.TransferDate
}

func (t *Transfer6) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer6) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer6) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

// Completion of a securities settlement instruction, wherein securities are delivered/debited from a securities account and received/credited to the designated securities account.
type Transfer7 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identifier for a transfer execution, as assigned by a confirming party.
	TransferConfirmationReference *Max35Text `xml:"TrfConfRef"`

	// Reference that identifies the transfer in transaction.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Date and time at which the transfer was executed.
	EffectiveTransferDate *DateAndDateTimeChoice `xml:"FctvTrfDt"`

	// Date and time at which a transaction is completed and cleared, ie, securities are delivered.
	TradeDate *DateAndDateTimeChoice `xml:"TradDt"`

	// Total quantity of securities settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit3 `xml:"UnitsDtls,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`
}

func (t *Transfer7) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *Transfer7) SetTransferConfirmationReference(value string) {
	t.TransferConfirmationReference = (*Max35Text)(&value)
}

func (t *Transfer7) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer7) SetClientReference(value string) {
	t.ClientReference = (*Max35Text)(&value)
}

func (t *Transfer7) AddEffectiveTransferDate() *DateAndDateTimeChoice {
	t.EffectiveTransferDate = new(DateAndDateTimeChoice)
	return t.EffectiveTransferDate
}

func (t *Transfer7) AddTradeDate() *DateAndDateTimeChoice {
	t.TradeDate = new(DateAndDateTimeChoice)
	return t.TradeDate
}

func (t *Transfer7) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer7) AddUnitsDetails() *Unit3 {
	newValue := new(Unit3)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer7) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer7) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer7) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

// Parameters applied to the settlement of a security transfer.
type Transfer8 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Requested date at which the instructing party places the transfer instruction.
	RequestedTransferDate *DateFormat1Choice `xml:"ReqdTrfDt"`

	// Total quantity of securities to be settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Total quantity of securities to be settled.
	PortfolioTransferOutRate *PercentageRate `xml:"PrtflTrfOutRate"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit3 `xml:"UnitsDtls,omitempty"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`
}

func (t *Transfer8) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *Transfer8) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer8) SetClientReference(value string) {
	t.ClientReference = (*Max35Text)(&value)
}

func (t *Transfer8) AddRequestedTransferDate() *DateFormat1Choice {
	t.RequestedTransferDate = new(DateFormat1Choice)
	return t.RequestedTransferDate
}

func (t *Transfer8) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer8) SetPortfolioTransferOutRate(value string) {
	t.PortfolioTransferOutRate = (*PercentageRate)(&value)
}

func (t *Transfer8) AddUnitsDetails() *Unit3 {
	newValue := new(Unit3)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer8) SetRounding(value string) {
	t.Rounding = (*RoundingDirection2Code)(&value)
}

func (t *Transfer8) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer8) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer8) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

// Parameters applied to the settlement of a security transfer.
type Transfer9 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Requested date at which the instructing party places the transfer instruction.
	RequestedTransferDate *DateFormat1Choice `xml:"ReqdTrfDt"`

	// Total quantity of securities to be settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Total quantity of securities to be settled.
	PortfolioTransferOutRate *PercentageRate `xml:"PrtflTrfOutRate"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit3 `xml:"UnitsDtls,omitempty"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`
}

func (t *Transfer9) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *Transfer9) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer9) SetClientReference(value string) {
	t.ClientReference = (*Max35Text)(&value)
}

func (t *Transfer9) SetCancellationReference(value string) {
	t.CancellationReference = (*Max35Text)(&value)
}

func (t *Transfer9) AddRequestedTransferDate() *DateFormat1Choice {
	t.RequestedTransferDate = new(DateFormat1Choice)
	return t.RequestedTransferDate
}

func (t *Transfer9) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer9) SetPortfolioTransferOutRate(value string) {
	t.PortfolioTransferOutRate = (*PercentageRate)(&value)
}

func (t *Transfer9) AddUnitsDetails() *Unit3 {
	newValue := new(Unit3)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer9) SetRounding(value string) {
	t.Rounding = (*RoundingDirection2Code)(&value)
}

func (t *Transfer9) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer9) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer9) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}
