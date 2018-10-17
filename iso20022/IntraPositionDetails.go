package iso20022

// Details of the intra-position movement.
type IntraPositionDetails1 struct {

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Date and time at which the securities are to be moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Balance from which the securities are moving.
	BalanceFrom *SecuritiesBalanceType2Choice `xml:"BalFr"`

	// Balance to which the securities are moving.
	BalanceTo *SecuritiesBalanceType2Choice `xml:"BalTo"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	InstructionProcessingAdditionalDetails *Max350Text `xml:"InstrPrcgAddtlDtls,omitempty"`
}

func (i *IntraPositionDetails1) AddPriority() *PriorityNumeric1Choice {
	i.Priority = new(PriorityNumeric1Choice)
	return i.Priority
}

func (i *IntraPositionDetails1) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	i.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.SettlementQuantity
}

func (i *IntraPositionDetails1) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails1) AddBalanceFrom() *SecuritiesBalanceType2Choice {
	i.BalanceFrom = new(SecuritiesBalanceType2Choice)
	return i.BalanceFrom
}

func (i *IntraPositionDetails1) AddBalanceTo() *SecuritiesBalanceType2Choice {
	i.BalanceTo = new(SecuritiesBalanceType2Choice)
	return i.BalanceTo
}

func (i *IntraPositionDetails1) SetInstructionProcessingAdditionalDetails(value string) {
	i.InstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the intra-position movement.
type IntraPositionDetails11 struct {

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Number identifying a lot constituting the financial instrument.
	LotNumber *Number2Choice `xml:"LotNb,omitempty"`

	// Date and time at which the securities are to be moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Balance from which the securities are moving.
	BalanceFrom *SecuritiesBalanceType2Choice `xml:"BalFr"`

	// Balance to which the securities are moving.
	BalanceTo *SecuritiesBalanceType2Choice `xml:"BalTo"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	InstructionProcessingAdditionalDetails *Max350Text `xml:"InstrPrcgAddtlDtls,omitempty"`
}

func (i *IntraPositionDetails11) AddPriority() *PriorityNumeric1Choice {
	i.Priority = new(PriorityNumeric1Choice)
	return i.Priority
}

func (i *IntraPositionDetails11) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	i.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.SettlementQuantity
}

func (i *IntraPositionDetails11) AddLotNumber() *Number2Choice {
	i.LotNumber = new(Number2Choice)
	return i.LotNumber
}

func (i *IntraPositionDetails11) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails11) AddBalanceFrom() *SecuritiesBalanceType2Choice {
	i.BalanceFrom = new(SecuritiesBalanceType2Choice)
	return i.BalanceFrom
}

func (i *IntraPositionDetails11) AddBalanceTo() *SecuritiesBalanceType2Choice {
	i.BalanceTo = new(SecuritiesBalanceType2Choice)
	return i.BalanceTo
}

func (i *IntraPositionDetails11) SetInstructionProcessingAdditionalDetails(value string) {
	i.InstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the intra-position movement.
type IntraPositionDetails12 struct {

	// Quantity of financial instrument effectively settled.
	SettledQuantity *FinancialInstrumentQuantity1Choice `xml:"SttldQty"`

	// Number identifying a lot constituting the financial instrument.
	LotNumber *Number2Choice `xml:"LotNb,omitempty"`

	// Value of the collateral available for the delivery settlement process at the account level.
	CollateralMonitorAmount *AmountAndDirection9 `xml:"CollMntrAmt,omitempty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity1Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity1Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Date and time at which the securities were moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Date/time securities become available for sale (if securities become unavailable, this specifies the date/time at which they will become available again).
	AvailableDate *DateAndDateTimeChoice `xml:"AvlblDt,omitempty"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType3Choice `xml:"CorpActnEvtTp,omitempty"`

	// Balance from which the securities were moved.
	BalanceFrom *SecuritiesBalanceType3Choice `xml:"BalFr"`

	// Balance to which the securities were moved.
	BalanceTo *SecuritiesBalanceType3Choice `xml:"BalTo"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	InstructionProcessingAdditionalDetails *Max350Text `xml:"InstrPrcgAddtlDtls,omitempty"`
}

func (i *IntraPositionDetails12) AddSettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.SettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.SettledQuantity
}

func (i *IntraPositionDetails12) AddLotNumber() *Number2Choice {
	i.LotNumber = new(Number2Choice)
	return i.LotNumber
}

func (i *IntraPositionDetails12) AddCollateralMonitorAmount() *AmountAndDirection9 {
	i.CollateralMonitorAmount = new(AmountAndDirection9)
	return i.CollateralMonitorAmount
}

func (i *IntraPositionDetails12) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.PreviouslySettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.PreviouslySettledQuantity
}

func (i *IntraPositionDetails12) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.RemainingToBeSettledQuantity
}

func (i *IntraPositionDetails12) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails12) AddAvailableDate() *DateAndDateTimeChoice {
	i.AvailableDate = new(DateAndDateTimeChoice)
	return i.AvailableDate
}

func (i *IntraPositionDetails12) AddCorporateActionEventType() *CorporateActionEventType3Choice {
	i.CorporateActionEventType = new(CorporateActionEventType3Choice)
	return i.CorporateActionEventType
}

func (i *IntraPositionDetails12) AddBalanceFrom() *SecuritiesBalanceType3Choice {
	i.BalanceFrom = new(SecuritiesBalanceType3Choice)
	return i.BalanceFrom
}

func (i *IntraPositionDetails12) AddBalanceTo() *SecuritiesBalanceType3Choice {
	i.BalanceTo = new(SecuritiesBalanceType3Choice)
	return i.BalanceTo
}

func (i *IntraPositionDetails12) SetInstructionProcessingAdditionalDetails(value string) {
	i.InstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the intra-position movement.
type IntraPositionDetails17 struct {

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Balance from which the securities were moved.
	BalanceFrom *SecuritiesBalanceType3Choice `xml:"BalFr"`

	// Intra-position movement(s) having been performed.
	IntraPositionMovement []*IntraPositionMovementDetails5 `xml:"IntraPosMvmnt"`
}

func (i *IntraPositionDetails17) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	i.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return i.SafekeepingPlace
}

func (i *IntraPositionDetails17) AddBalanceFrom() *SecuritiesBalanceType3Choice {
	i.BalanceFrom = new(SecuritiesBalanceType3Choice)
	return i.BalanceFrom
}

func (i *IntraPositionDetails17) AddIntraPositionMovement() *IntraPositionMovementDetails5 {
	newValue := new(IntraPositionMovementDetails5)
	i.IntraPositionMovement = append(i.IntraPositionMovement, newValue)
	return newValue
}

// Details of the intra-position movement.
type IntraPositionDetails19 struct {

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Number identifying a lot constituting the financial instrument.
	LotNumber *GenericIdentification37 `xml:"LotNb,omitempty"`

	// Date and time at which the securities are to be moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Balance from which the securities are moving.
	BalanceFrom *SecuritiesBalanceType2Choice `xml:"BalFr,omitempty"`

	// Balance to which the securities are moving.
	BalanceTo *SecuritiesBalanceType2Choice `xml:"BalTo,omitempty"`
}

func (i *IntraPositionDetails19) SetPoolIdentification(value string) {
	i.PoolIdentification = (*Max35Text)(&value)
}

func (i *IntraPositionDetails19) AddAccountOwner() *PartyIdentification36Choice {
	i.AccountOwner = new(PartyIdentification36Choice)
	return i.AccountOwner
}

func (i *IntraPositionDetails19) AddSafekeepingAccount() *SecuritiesAccount13 {
	i.SafekeepingAccount = new(SecuritiesAccount13)
	return i.SafekeepingAccount
}

func (i *IntraPositionDetails19) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	i.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return i.FinancialInstrumentIdentification
}

func (i *IntraPositionDetails19) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	i.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.SettlementQuantity
}

func (i *IntraPositionDetails19) AddLotNumber() *GenericIdentification37 {
	i.LotNumber = new(GenericIdentification37)
	return i.LotNumber
}

func (i *IntraPositionDetails19) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails19) AddBalanceFrom() *SecuritiesBalanceType2Choice {
	i.BalanceFrom = new(SecuritiesBalanceType2Choice)
	return i.BalanceFrom
}

func (i *IntraPositionDetails19) AddBalanceTo() *SecuritiesBalanceType2Choice {
	i.BalanceTo = new(SecuritiesBalanceType2Choice)
	return i.BalanceTo
}

// Details of the intra-position movement.
type IntraPositionDetails2 struct {

	// Quantity of financial instrument effectively settled.
	SettledQuantity *FinancialInstrumentQuantity1Choice `xml:"SttldQty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity1Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity1Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Date and time at which the securities were moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Date/time securities become available for sale (if securities become unavailable, this specifies the date/time at which they will become available again).
	AvailableDate *DateAndDateTimeChoice `xml:"AvlblDt,omitempty"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType3Choice `xml:"CorpActnEvtTp,omitempty"`

	// Balance from which the securities were moved.
	BalanceFrom *SecuritiesBalanceType3Choice `xml:"BalFr"`

	// Balance to which the securities were moved.
	BalanceTo *SecuritiesBalanceType3Choice `xml:"BalTo"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	InstructionProcessingAdditionalDetails *Max350Text `xml:"InstrPrcgAddtlDtls,omitempty"`
}

func (i *IntraPositionDetails2) AddSettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.SettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.SettledQuantity
}

func (i *IntraPositionDetails2) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.PreviouslySettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.PreviouslySettledQuantity
}

func (i *IntraPositionDetails2) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.RemainingToBeSettledQuantity
}

func (i *IntraPositionDetails2) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails2) AddAvailableDate() *DateAndDateTimeChoice {
	i.AvailableDate = new(DateAndDateTimeChoice)
	return i.AvailableDate
}

func (i *IntraPositionDetails2) AddCorporateActionEventType() *CorporateActionEventType3Choice {
	i.CorporateActionEventType = new(CorporateActionEventType3Choice)
	return i.CorporateActionEventType
}

func (i *IntraPositionDetails2) AddBalanceFrom() *SecuritiesBalanceType3Choice {
	i.BalanceFrom = new(SecuritiesBalanceType3Choice)
	return i.BalanceFrom
}

func (i *IntraPositionDetails2) AddBalanceTo() *SecuritiesBalanceType3Choice {
	i.BalanceTo = new(SecuritiesBalanceType3Choice)
	return i.BalanceTo
}

func (i *IntraPositionDetails2) SetInstructionProcessingAdditionalDetails(value string) {
	i.InstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the intra-position movement.
type IntraPositionDetails20 struct {

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Balance from which the securities were moved.
	BalanceFrom *SecuritiesBalanceType3Choice `xml:"BalFr"`

	// Intra-position movement(s) having been performed.
	IntraPositionMovement []*IntraPositionMovementDetails7 `xml:"IntraPosMvmnt"`
}

func (i *IntraPositionDetails20) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	i.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return i.SafekeepingPlace
}

func (i *IntraPositionDetails20) AddBalanceFrom() *SecuritiesBalanceType3Choice {
	i.BalanceFrom = new(SecuritiesBalanceType3Choice)
	return i.BalanceFrom
}

func (i *IntraPositionDetails20) AddIntraPositionMovement() *IntraPositionMovementDetails7 {
	newValue := new(IntraPositionMovementDetails7)
	i.IntraPositionMovement = append(i.IntraPositionMovement, newValue)
	return newValue
}

// Details of the intra-position movement.
type IntraPositionDetails21 struct {

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Number identifying a Securities Sub balance Type (e.g. restriction identification etc…).
	SecuritiesSubBalanceIdentification *GenericIdentification37 `xml:"SctiesSubBalId,omitempty"`

	// Date and time at which the securities are to be moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Balance from which the securities are moving.
	BalanceFrom *SecuritiesSubBalanceTypeAndQuantityBreakdown1 `xml:"BalFr"`

	// Balance to which the securities are moving.
	BalanceTo *SecuritiesSubBalanceTypeAndQuantityBreakdown1 `xml:"BalTo"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	InstructionProcessingAdditionalDetails *Max350Text `xml:"InstrPrcgAddtlDtls,omitempty"`
}

func (i *IntraPositionDetails21) AddPriority() *PriorityNumeric1Choice {
	i.Priority = new(PriorityNumeric1Choice)
	return i.Priority
}

func (i *IntraPositionDetails21) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	i.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.SettlementQuantity
}

func (i *IntraPositionDetails21) AddSecuritiesSubBalanceIdentification() *GenericIdentification37 {
	i.SecuritiesSubBalanceIdentification = new(GenericIdentification37)
	return i.SecuritiesSubBalanceIdentification
}

func (i *IntraPositionDetails21) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails21) AddBalanceFrom() *SecuritiesSubBalanceTypeAndQuantityBreakdown1 {
	i.BalanceFrom = new(SecuritiesSubBalanceTypeAndQuantityBreakdown1)
	return i.BalanceFrom
}

func (i *IntraPositionDetails21) AddBalanceTo() *SecuritiesSubBalanceTypeAndQuantityBreakdown1 {
	i.BalanceTo = new(SecuritiesSubBalanceTypeAndQuantityBreakdown1)
	return i.BalanceTo
}

func (i *IntraPositionDetails21) SetInstructionProcessingAdditionalDetails(value string) {
	i.InstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the intra-position movement.
type IntraPositionDetails22 struct {

	// Quantity of financial instrument effectively settled.
	SettledQuantity *FinancialInstrumentQuantity1Choice `xml:"SttldQty"`

	// Number identifying a Securities Sub balance Type (e.g. restriction identification etc…).
	SecuritiesSubBalanceIdentification *GenericIdentification37 `xml:"SctiesSubBalId,omitempty"`

	// Value of the collateral available for the delivery settlement process at the account level.
	CollateralMonitorAmount *AmountAndDirection9 `xml:"CollMntrAmt,omitempty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity1Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity1Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Date and time at which the securities were moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Date/time securities become available for sale (if securities become unavailable, this specifies the date/time at which they will become available again).
	AvailableDate *DateAndDateTimeChoice `xml:"AvlblDt,omitempty"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType3Choice `xml:"CorpActnEvtTp,omitempty"`

	// Balance from which the securities are moving.
	BalanceFrom *SecuritiesSubBalanceTypeAndQuantityBreakdown1 `xml:"BalFr"`

	// Balance to which the securities are moving.
	BalanceTo *SecuritiesSubBalanceTypeAndQuantityBreakdown1 `xml:"BalTo"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	InstructionProcessingAdditionalDetails *Max350Text `xml:"InstrPrcgAddtlDtls,omitempty"`
}

func (i *IntraPositionDetails22) AddSettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.SettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.SettledQuantity
}

func (i *IntraPositionDetails22) AddSecuritiesSubBalanceIdentification() *GenericIdentification37 {
	i.SecuritiesSubBalanceIdentification = new(GenericIdentification37)
	return i.SecuritiesSubBalanceIdentification
}

func (i *IntraPositionDetails22) AddCollateralMonitorAmount() *AmountAndDirection9 {
	i.CollateralMonitorAmount = new(AmountAndDirection9)
	return i.CollateralMonitorAmount
}

func (i *IntraPositionDetails22) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.PreviouslySettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.PreviouslySettledQuantity
}

func (i *IntraPositionDetails22) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.RemainingToBeSettledQuantity
}

func (i *IntraPositionDetails22) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails22) AddAvailableDate() *DateAndDateTimeChoice {
	i.AvailableDate = new(DateAndDateTimeChoice)
	return i.AvailableDate
}

func (i *IntraPositionDetails22) AddCorporateActionEventType() *CorporateActionEventType3Choice {
	i.CorporateActionEventType = new(CorporateActionEventType3Choice)
	return i.CorporateActionEventType
}

func (i *IntraPositionDetails22) AddBalanceFrom() *SecuritiesSubBalanceTypeAndQuantityBreakdown1 {
	i.BalanceFrom = new(SecuritiesSubBalanceTypeAndQuantityBreakdown1)
	return i.BalanceFrom
}

func (i *IntraPositionDetails22) AddBalanceTo() *SecuritiesSubBalanceTypeAndQuantityBreakdown1 {
	i.BalanceTo = new(SecuritiesSubBalanceTypeAndQuantityBreakdown1)
	return i.BalanceTo
}

func (i *IntraPositionDetails22) SetInstructionProcessingAdditionalDetails(value string) {
	i.InstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the intra-position movement.
type IntraPositionDetails27 struct {

	// Quantity of financial instrument effectively settled.
	SettledQuantity *FinancialInstrumentQuantity1Choice `xml:"SttldQty"`

	// Number identifying a Securities Sub balance Type (example restriction identification etc…).
	SecuritiesSubBalanceIdentification *GenericIdentification37 `xml:"SctiesSubBalId,omitempty"`

	// Value of the collateral available for the delivery settlement process at the account level.
	CollateralMonitorAmount *AmountAndDirection9 `xml:"CollMntrAmt,omitempty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity1Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity1Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Date and time at which the securities were moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Date/time securities become available for sale (if securities become unavailable, this specifies the date/time at which they will become available again).
	AvailableDate *DateAndDateTimeChoice `xml:"AvlblDt,omitempty"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType16Choice `xml:"CorpActnEvtTp,omitempty"`

	// Balance from which the securities are moving.
	BalanceFrom *SecuritiesSubBalanceTypeAndQuantityBreakdown1 `xml:"BalFr"`

	// Balance to which the securities are moving.
	BalanceTo *SecuritiesSubBalanceTypeAndQuantityBreakdown1 `xml:"BalTo"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	InstructionProcessingAdditionalDetails *Max350Text `xml:"InstrPrcgAddtlDtls,omitempty"`
}

func (i *IntraPositionDetails27) AddSettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.SettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.SettledQuantity
}

func (i *IntraPositionDetails27) AddSecuritiesSubBalanceIdentification() *GenericIdentification37 {
	i.SecuritiesSubBalanceIdentification = new(GenericIdentification37)
	return i.SecuritiesSubBalanceIdentification
}

func (i *IntraPositionDetails27) AddCollateralMonitorAmount() *AmountAndDirection9 {
	i.CollateralMonitorAmount = new(AmountAndDirection9)
	return i.CollateralMonitorAmount
}

func (i *IntraPositionDetails27) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.PreviouslySettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.PreviouslySettledQuantity
}

func (i *IntraPositionDetails27) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.RemainingToBeSettledQuantity
}

func (i *IntraPositionDetails27) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails27) AddAvailableDate() *DateAndDateTimeChoice {
	i.AvailableDate = new(DateAndDateTimeChoice)
	return i.AvailableDate
}

func (i *IntraPositionDetails27) AddCorporateActionEventType() *CorporateActionEventType16Choice {
	i.CorporateActionEventType = new(CorporateActionEventType16Choice)
	return i.CorporateActionEventType
}

func (i *IntraPositionDetails27) AddBalanceFrom() *SecuritiesSubBalanceTypeAndQuantityBreakdown1 {
	i.BalanceFrom = new(SecuritiesSubBalanceTypeAndQuantityBreakdown1)
	return i.BalanceFrom
}

func (i *IntraPositionDetails27) AddBalanceTo() *SecuritiesSubBalanceTypeAndQuantityBreakdown1 {
	i.BalanceTo = new(SecuritiesSubBalanceTypeAndQuantityBreakdown1)
	return i.BalanceTo
}

func (i *IntraPositionDetails27) SetInstructionProcessingAdditionalDetails(value string) {
	i.InstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the intra-position movement.
type IntraPositionDetails28 struct {

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Balance from which the securities were moved.
	BalanceFrom *SecuritiesBalanceType3Choice `xml:"BalFr"`

	// Intra-position movement(s) having been performed.
	IntraPositionMovement []*IntraPositionMovementDetails9 `xml:"IntraPosMvmnt"`
}

func (i *IntraPositionDetails28) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	i.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return i.SafekeepingPlace
}

func (i *IntraPositionDetails28) AddBalanceFrom() *SecuritiesBalanceType3Choice {
	i.BalanceFrom = new(SecuritiesBalanceType3Choice)
	return i.BalanceFrom
}

func (i *IntraPositionDetails28) AddIntraPositionMovement() *IntraPositionMovementDetails9 {
	newValue := new(IntraPositionMovementDetails9)
	i.IntraPositionMovement = append(i.IntraPositionMovement, newValue)
	return newValue
}

// Details of the intra-position movement.
type IntraPositionDetails3 struct {

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Balance from which the securities were moved.
	BalanceFrom *SecuritiesBalanceType3Choice `xml:"BalFr"`

	// Intra-position movement(s) having been performed.
	IntraPositionMovement []*IntraPositionMovementDetails1 `xml:"IntraPosMvmnt"`
}

func (i *IntraPositionDetails3) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	i.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return i.SafekeepingPlace
}

func (i *IntraPositionDetails3) AddBalanceFrom() *SecuritiesBalanceType3Choice {
	i.BalanceFrom = new(SecuritiesBalanceType3Choice)
	return i.BalanceFrom
}

func (i *IntraPositionDetails3) AddIntraPositionMovement() *IntraPositionMovementDetails1 {
	newValue := new(IntraPositionMovementDetails1)
	i.IntraPositionMovement = append(i.IntraPositionMovement, newValue)
	return newValue
}

// Details of the intra-position movement.
type IntraPositionDetails31 struct {

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification92Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Number identifying a lot constituting the financial instrument.
	LotNumber *GenericIdentification37 `xml:"LotNb,omitempty"`

	// Date and time at which the securities are to be moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Balance from which the securities are moving.
	BalanceFrom *SecuritiesBalanceType7Choice `xml:"BalFr,omitempty"`

	// Balance to which the securities are moving.
	BalanceTo *SecuritiesBalanceType7Choice `xml:"BalTo,omitempty"`
}

func (i *IntraPositionDetails31) SetPoolIdentification(value string) {
	i.PoolIdentification = (*Max35Text)(&value)
}

func (i *IntraPositionDetails31) AddAccountOwner() *PartyIdentification92Choice {
	i.AccountOwner = new(PartyIdentification92Choice)
	return i.AccountOwner
}

func (i *IntraPositionDetails31) AddSafekeepingAccount() *SecuritiesAccount19 {
	i.SafekeepingAccount = new(SecuritiesAccount19)
	return i.SafekeepingAccount
}

func (i *IntraPositionDetails31) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	i.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return i.FinancialInstrumentIdentification
}

func (i *IntraPositionDetails31) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	i.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.SettlementQuantity
}

func (i *IntraPositionDetails31) AddLotNumber() *GenericIdentification37 {
	i.LotNumber = new(GenericIdentification37)
	return i.LotNumber
}

func (i *IntraPositionDetails31) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails31) AddBalanceFrom() *SecuritiesBalanceType7Choice {
	i.BalanceFrom = new(SecuritiesBalanceType7Choice)
	return i.BalanceFrom
}

func (i *IntraPositionDetails31) AddBalanceTo() *SecuritiesBalanceType7Choice {
	i.BalanceTo = new(SecuritiesBalanceType7Choice)
	return i.BalanceTo
}

// Details of the intra-position movement.
type IntraPositionDetails32 struct {

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat10Choice `xml:"SfkpgPlc,omitempty"`

	// Balance from which the securities were moved.
	BalanceFrom *SecuritiesBalanceType6Choice `xml:"BalFr"`

	// Intra-position movement(s) having been performed.
	IntraPositionMovement []*IntraPositionMovementDetails11 `xml:"IntraPosMvmnt"`
}

func (i *IntraPositionDetails32) AddSafekeepingPlace() *SafekeepingPlaceFormat10Choice {
	i.SafekeepingPlace = new(SafekeepingPlaceFormat10Choice)
	return i.SafekeepingPlace
}

func (i *IntraPositionDetails32) AddBalanceFrom() *SecuritiesBalanceType6Choice {
	i.BalanceFrom = new(SecuritiesBalanceType6Choice)
	return i.BalanceFrom
}

func (i *IntraPositionDetails32) AddIntraPositionMovement() *IntraPositionMovementDetails11 {
	newValue := new(IntraPositionMovementDetails11)
	i.IntraPositionMovement = append(i.IntraPositionMovement, newValue)
	return newValue
}

// Details of the intra-position movement.
type IntraPositionDetails33 struct {

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric4Choice `xml:"Prty,omitempty"`

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Number identifying a Securities Sub balance Type (e.g. restriction identification etc…).
	SecuritiesSubBalanceIdentification *GenericIdentification37 `xml:"SctiesSubBalId,omitempty"`

	// Date and time at which the securities are to be moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Balance from which the securities are moving.
	BalanceFrom *SecuritiesSubBalanceTypeAndQuantityBreakdown3 `xml:"BalFr"`

	// Balance to which the securities are moving.
	BalanceTo *SecuritiesSubBalanceTypeAndQuantityBreakdown3 `xml:"BalTo"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	InstructionProcessingAdditionalDetails *Max350Text `xml:"InstrPrcgAddtlDtls,omitempty"`
}

func (i *IntraPositionDetails33) AddPriority() *PriorityNumeric4Choice {
	i.Priority = new(PriorityNumeric4Choice)
	return i.Priority
}

func (i *IntraPositionDetails33) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	i.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.SettlementQuantity
}

func (i *IntraPositionDetails33) AddSecuritiesSubBalanceIdentification() *GenericIdentification37 {
	i.SecuritiesSubBalanceIdentification = new(GenericIdentification37)
	return i.SecuritiesSubBalanceIdentification
}

func (i *IntraPositionDetails33) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails33) AddBalanceFrom() *SecuritiesSubBalanceTypeAndQuantityBreakdown3 {
	i.BalanceFrom = new(SecuritiesSubBalanceTypeAndQuantityBreakdown3)
	return i.BalanceFrom
}

func (i *IntraPositionDetails33) AddBalanceTo() *SecuritiesSubBalanceTypeAndQuantityBreakdown3 {
	i.BalanceTo = new(SecuritiesSubBalanceTypeAndQuantityBreakdown3)
	return i.BalanceTo
}

func (i *IntraPositionDetails33) SetInstructionProcessingAdditionalDetails(value string) {
	i.InstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the intra-position movement.
type IntraPositionDetails34 struct {

	// Quantity of financial instrument effectively settled.
	SettledQuantity *FinancialInstrumentQuantity1Choice `xml:"SttldQty"`

	// Number identifying a Securities Sub balance Type (example restriction identification etc…).
	SecuritiesSubBalanceIdentification *GenericIdentification37 `xml:"SctiesSubBalId,omitempty"`

	// Value of the collateral available for the delivery settlement process at the account level.
	CollateralMonitorAmount *AmountAndDirection44 `xml:"CollMntrAmt,omitempty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity1Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity1Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Date and time at which the securities were moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Date/time securities become available for sale (if securities become unavailable, this specifies the date/time at which they will become available again).
	AvailableDate *DateAndDateTimeChoice `xml:"AvlblDt,omitempty"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType29Choice `xml:"CorpActnEvtTp,omitempty"`

	// Balance from which the securities are moving.
	BalanceFrom *SecuritiesSubBalanceTypeAndQuantityBreakdown3 `xml:"BalFr"`

	// Balance to which the securities are moving.
	BalanceTo *SecuritiesSubBalanceTypeAndQuantityBreakdown3 `xml:"BalTo"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	InstructionProcessingAdditionalDetails *Max350Text `xml:"InstrPrcgAddtlDtls,omitempty"`
}

func (i *IntraPositionDetails34) AddSettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.SettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.SettledQuantity
}

func (i *IntraPositionDetails34) AddSecuritiesSubBalanceIdentification() *GenericIdentification37 {
	i.SecuritiesSubBalanceIdentification = new(GenericIdentification37)
	return i.SecuritiesSubBalanceIdentification
}

func (i *IntraPositionDetails34) AddCollateralMonitorAmount() *AmountAndDirection44 {
	i.CollateralMonitorAmount = new(AmountAndDirection44)
	return i.CollateralMonitorAmount
}

func (i *IntraPositionDetails34) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.PreviouslySettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.PreviouslySettledQuantity
}

func (i *IntraPositionDetails34) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.RemainingToBeSettledQuantity
}

func (i *IntraPositionDetails34) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails34) AddAvailableDate() *DateAndDateTimeChoice {
	i.AvailableDate = new(DateAndDateTimeChoice)
	return i.AvailableDate
}

func (i *IntraPositionDetails34) AddCorporateActionEventType() *CorporateActionEventType29Choice {
	i.CorporateActionEventType = new(CorporateActionEventType29Choice)
	return i.CorporateActionEventType
}

func (i *IntraPositionDetails34) AddBalanceFrom() *SecuritiesSubBalanceTypeAndQuantityBreakdown3 {
	i.BalanceFrom = new(SecuritiesSubBalanceTypeAndQuantityBreakdown3)
	return i.BalanceFrom
}

func (i *IntraPositionDetails34) AddBalanceTo() *SecuritiesSubBalanceTypeAndQuantityBreakdown3 {
	i.BalanceTo = new(SecuritiesSubBalanceTypeAndQuantityBreakdown3)
	return i.BalanceTo
}

func (i *IntraPositionDetails34) SetInstructionProcessingAdditionalDetails(value string) {
	i.InstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the intra-position movement.
type IntraPositionDetails35 struct {

	// Quantity of financial instrument effectively settled.
	SettledQuantity *FinancialInstrumentQuantity15Choice `xml:"SttldQty"`

	// Number identifying a Securities Sub balance Type (example restriction identification etc…).
	SecuritiesSubBalanceIdentification *GenericIdentification39 `xml:"SctiesSubBalId,omitempty"`

	// Value of the collateral available for the delivery settlement process at the account level.
	CollateralMonitorAmount *AmountAndDirection55 `xml:"CollMntrAmt,omitempty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity15Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity15Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Date and time at which the securities were moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Date/time securities become available for sale (if securities become unavailable, this specifies the date/time at which they will become available again).
	AvailableDate *DateAndDateTimeChoice `xml:"AvlblDt,omitempty"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType46Choice `xml:"CorpActnEvtTp,omitempty"`

	// Balance from which the securities are moving.
	BalanceFrom *SecuritiesSubBalanceTypeAndQuantityBreakdown4 `xml:"BalFr"`

	// Balance to which the securities are moving.
	BalanceTo *SecuritiesSubBalanceTypeAndQuantityBreakdown4 `xml:"BalTo"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	InstructionProcessingAdditionalDetails *RestrictedFINXMax350Text `xml:"InstrPrcgAddtlDtls,omitempty"`
}

func (i *IntraPositionDetails35) AddSettledQuantity() *FinancialInstrumentQuantity15Choice {
	i.SettledQuantity = new(FinancialInstrumentQuantity15Choice)
	return i.SettledQuantity
}

func (i *IntraPositionDetails35) AddSecuritiesSubBalanceIdentification() *GenericIdentification39 {
	i.SecuritiesSubBalanceIdentification = new(GenericIdentification39)
	return i.SecuritiesSubBalanceIdentification
}

func (i *IntraPositionDetails35) AddCollateralMonitorAmount() *AmountAndDirection55 {
	i.CollateralMonitorAmount = new(AmountAndDirection55)
	return i.CollateralMonitorAmount
}

func (i *IntraPositionDetails35) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity15Choice {
	i.PreviouslySettledQuantity = new(FinancialInstrumentQuantity15Choice)
	return i.PreviouslySettledQuantity
}

func (i *IntraPositionDetails35) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity15Choice {
	i.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity15Choice)
	return i.RemainingToBeSettledQuantity
}

func (i *IntraPositionDetails35) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails35) AddAvailableDate() *DateAndDateTimeChoice {
	i.AvailableDate = new(DateAndDateTimeChoice)
	return i.AvailableDate
}

func (i *IntraPositionDetails35) AddCorporateActionEventType() *CorporateActionEventType46Choice {
	i.CorporateActionEventType = new(CorporateActionEventType46Choice)
	return i.CorporateActionEventType
}

func (i *IntraPositionDetails35) AddBalanceFrom() *SecuritiesSubBalanceTypeAndQuantityBreakdown4 {
	i.BalanceFrom = new(SecuritiesSubBalanceTypeAndQuantityBreakdown4)
	return i.BalanceFrom
}

func (i *IntraPositionDetails35) AddBalanceTo() *SecuritiesSubBalanceTypeAndQuantityBreakdown4 {
	i.BalanceTo = new(SecuritiesSubBalanceTypeAndQuantityBreakdown4)
	return i.BalanceTo
}

func (i *IntraPositionDetails35) SetInstructionProcessingAdditionalDetails(value string) {
	i.InstructionProcessingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

// Details of the intra-position movement.
type IntraPositionDetails36 struct {

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric5Choice `xml:"Prty,omitempty"`

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity15Choice `xml:"SttlmQty"`

	// Number identifying a Securities Sub balance Type (e.g. restriction identification etc…).
	SecuritiesSubBalanceIdentification *GenericIdentification39 `xml:"SctiesSubBalId,omitempty"`

	// Date and time at which the securities are to be moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Balance from which the securities are moving.
	BalanceFrom *SecuritiesSubBalanceTypeAndQuantityBreakdown4 `xml:"BalFr"`

	// Balance to which the securities are moving.
	BalanceTo *SecuritiesSubBalanceTypeAndQuantityBreakdown4 `xml:"BalTo"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	InstructionProcessingAdditionalDetails *RestrictedFINXMax350Text `xml:"InstrPrcgAddtlDtls,omitempty"`
}

func (i *IntraPositionDetails36) AddPriority() *PriorityNumeric5Choice {
	i.Priority = new(PriorityNumeric5Choice)
	return i.Priority
}

func (i *IntraPositionDetails36) AddSettlementQuantity() *FinancialInstrumentQuantity15Choice {
	i.SettlementQuantity = new(FinancialInstrumentQuantity15Choice)
	return i.SettlementQuantity
}

func (i *IntraPositionDetails36) AddSecuritiesSubBalanceIdentification() *GenericIdentification39 {
	i.SecuritiesSubBalanceIdentification = new(GenericIdentification39)
	return i.SecuritiesSubBalanceIdentification
}

func (i *IntraPositionDetails36) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails36) AddBalanceFrom() *SecuritiesSubBalanceTypeAndQuantityBreakdown4 {
	i.BalanceFrom = new(SecuritiesSubBalanceTypeAndQuantityBreakdown4)
	return i.BalanceFrom
}

func (i *IntraPositionDetails36) AddBalanceTo() *SecuritiesSubBalanceTypeAndQuantityBreakdown4 {
	i.BalanceTo = new(SecuritiesSubBalanceTypeAndQuantityBreakdown4)
	return i.BalanceTo
}

func (i *IntraPositionDetails36) SetInstructionProcessingAdditionalDetails(value string) {
	i.InstructionProcessingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

// Details of the intra-position movement.
type IntraPositionDetails37 struct {

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat17Choice `xml:"SfkpgPlc,omitempty"`

	// Balance from which the securities were moved.
	BalanceFrom *SecuritiesBalanceType8Choice `xml:"BalFr"`

	// Intra-position movement(s) having been performed.
	IntraPositionMovement []*IntraPositionMovementDetails12 `xml:"IntraPosMvmnt"`
}

func (i *IntraPositionDetails37) AddSafekeepingPlace() *SafekeepingPlaceFormat17Choice {
	i.SafekeepingPlace = new(SafekeepingPlaceFormat17Choice)
	return i.SafekeepingPlace
}

func (i *IntraPositionDetails37) AddBalanceFrom() *SecuritiesBalanceType8Choice {
	i.BalanceFrom = new(SecuritiesBalanceType8Choice)
	return i.BalanceFrom
}

func (i *IntraPositionDetails37) AddIntraPositionMovement() *IntraPositionMovementDetails12 {
	newValue := new(IntraPositionMovementDetails12)
	i.IntraPositionMovement = append(i.IntraPositionMovement, newValue)
	return newValue
}

// Details of the intra-position movement.
type IntraPositionDetails38 struct {

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification103Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount30 `xml:"SfkpgAcct"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Number identifying a lot constituting the financial instrument.
	LotNumber *GenericIdentification39 `xml:"LotNb,omitempty"`

	// Date and time at which the securities are to be moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Balance from which the securities are moving.
	BalanceFrom *SecuritiesBalanceType11Choice `xml:"BalFr,omitempty"`

	// Balance to which the securities are moving.
	BalanceTo *SecuritiesBalanceType11Choice `xml:"BalTo,omitempty"`
}

func (i *IntraPositionDetails38) SetPoolIdentification(value string) {
	i.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (i *IntraPositionDetails38) AddAccountOwner() *PartyIdentification103Choice {
	i.AccountOwner = new(PartyIdentification103Choice)
	return i.AccountOwner
}

func (i *IntraPositionDetails38) AddSafekeepingAccount() *SecuritiesAccount30 {
	i.SafekeepingAccount = new(SecuritiesAccount30)
	return i.SafekeepingAccount
}

func (i *IntraPositionDetails38) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	i.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return i.FinancialInstrumentIdentification
}

func (i *IntraPositionDetails38) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	i.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.SettlementQuantity
}

func (i *IntraPositionDetails38) AddLotNumber() *GenericIdentification39 {
	i.LotNumber = new(GenericIdentification39)
	return i.LotNumber
}

func (i *IntraPositionDetails38) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails38) AddBalanceFrom() *SecuritiesBalanceType11Choice {
	i.BalanceFrom = new(SecuritiesBalanceType11Choice)
	return i.BalanceFrom
}

func (i *IntraPositionDetails38) AddBalanceTo() *SecuritiesBalanceType11Choice {
	i.BalanceTo = new(SecuritiesBalanceType11Choice)
	return i.BalanceTo
}

// Details of the intra-position movement.
type IntraPositionDetails39 struct {

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification92Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Number identifying a lot constituting the financial instrument.
	LotNumber *GenericIdentification37 `xml:"LotNb,omitempty"`

	// Date and time at which the securities are to be moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Time stamp on when the transaction is acknowledged.
	AcknowledgedStatusTimeStamp *ISODateTime `xml:"AckdStsTmStmp,omitempty"`

	// Balance from which the securities are moving.
	BalanceFrom *SecuritiesBalanceType7Choice `xml:"BalFr,omitempty"`

	// Balance to which the securities are moving.
	BalanceTo *SecuritiesBalanceType7Choice `xml:"BalTo,omitempty"`
}

func (i *IntraPositionDetails39) SetPoolIdentification(value string) {
	i.PoolIdentification = (*Max35Text)(&value)
}

func (i *IntraPositionDetails39) AddAccountOwner() *PartyIdentification92Choice {
	i.AccountOwner = new(PartyIdentification92Choice)
	return i.AccountOwner
}

func (i *IntraPositionDetails39) AddSafekeepingAccount() *SecuritiesAccount19 {
	i.SafekeepingAccount = new(SecuritiesAccount19)
	return i.SafekeepingAccount
}

func (i *IntraPositionDetails39) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	i.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return i.FinancialInstrumentIdentification
}

func (i *IntraPositionDetails39) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	i.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.SettlementQuantity
}

func (i *IntraPositionDetails39) AddLotNumber() *GenericIdentification37 {
	i.LotNumber = new(GenericIdentification37)
	return i.LotNumber
}

func (i *IntraPositionDetails39) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails39) SetAcknowledgedStatusTimeStamp(value string) {
	i.AcknowledgedStatusTimeStamp = (*ISODateTime)(&value)
}

func (i *IntraPositionDetails39) AddBalanceFrom() *SecuritiesBalanceType7Choice {
	i.BalanceFrom = new(SecuritiesBalanceType7Choice)
	return i.BalanceFrom
}

func (i *IntraPositionDetails39) AddBalanceTo() *SecuritiesBalanceType7Choice {
	i.BalanceTo = new(SecuritiesBalanceType7Choice)
	return i.BalanceTo
}

// Details of the intra-position movement.
type IntraPositionDetails4 struct {

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification11 `xml:"FinInstrmId"`

	// Quantity of financial instrument effectively settled.
	SettledQuantity *FinancialInstrumentQuantity1Choice `xml:"SttldQty"`

	// Date and time at which the securities are to be moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Balance from which the securities are moving.
	BalanceFrom *SecuritiesBalanceType2Choice `xml:"BalFr,omitempty"`

	// Balance to which the securities are moving.
	BalanceTo *SecuritiesBalanceType2Choice `xml:"BalTo,omitempty"`
}

func (i *IntraPositionDetails4) SetPoolIdentification(value string) {
	i.PoolIdentification = (*Max35Text)(&value)
}

func (i *IntraPositionDetails4) AddAccountOwner() *PartyIdentification13Choice {
	i.AccountOwner = new(PartyIdentification13Choice)
	return i.AccountOwner
}

func (i *IntraPositionDetails4) AddSafekeepingAccount() *SecuritiesAccount13 {
	i.SafekeepingAccount = new(SecuritiesAccount13)
	return i.SafekeepingAccount
}

func (i *IntraPositionDetails4) AddFinancialInstrumentIdentification() *SecurityIdentification11 {
	i.FinancialInstrumentIdentification = new(SecurityIdentification11)
	return i.FinancialInstrumentIdentification
}

func (i *IntraPositionDetails4) AddSettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.SettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.SettledQuantity
}

func (i *IntraPositionDetails4) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails4) AddBalanceFrom() *SecuritiesBalanceType2Choice {
	i.BalanceFrom = new(SecuritiesBalanceType2Choice)
	return i.BalanceFrom
}

func (i *IntraPositionDetails4) AddBalanceTo() *SecuritiesBalanceType2Choice {
	i.BalanceTo = new(SecuritiesBalanceType2Choice)
	return i.BalanceTo
}

// Details of the intra-position movement.
type IntraPositionDetails40 struct {

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat10Choice `xml:"SfkpgPlc,omitempty"`

	// Balance from which the securities were moved.
	BalanceFrom *SecuritiesBalanceType6Choice `xml:"BalFr"`

	// Intra-position movement(s) having been performed.
	IntraPositionMovement []*IntraPositionMovementDetails13 `xml:"IntraPosMvmnt"`
}

func (i *IntraPositionDetails40) AddSafekeepingPlace() *SafekeepingPlaceFormat10Choice {
	i.SafekeepingPlace = new(SafekeepingPlaceFormat10Choice)
	return i.SafekeepingPlace
}

func (i *IntraPositionDetails40) AddBalanceFrom() *SecuritiesBalanceType6Choice {
	i.BalanceFrom = new(SecuritiesBalanceType6Choice)
	return i.BalanceFrom
}

func (i *IntraPositionDetails40) AddIntraPositionMovement() *IntraPositionMovementDetails13 {
	newValue := new(IntraPositionMovementDetails13)
	i.IntraPositionMovement = append(i.IntraPositionMovement, newValue)
	return newValue
}

// Details of the intra-position movement.
type IntraPositionDetails41 struct {

	// Quantity of financial instrument effectively settled.
	SettledQuantity *FinancialInstrumentQuantity1Choice `xml:"SttldQty"`

	// Number identifying a Securities Sub balance Type (example restriction identification etc…).
	SecuritiesSubBalanceIdentification *GenericIdentification37 `xml:"SctiesSubBalId,omitempty"`

	// Value of the collateral available for the delivery settlement process at the account level.
	CollateralMonitorAmount *AmountAndDirection44 `xml:"CollMntrAmt,omitempty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity1Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity1Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Date and time at which the securities were moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Date/time securities become available for sale (if securities become unavailable, this specifies the date/time at which they will become available again).
	AvailableDate *DateAndDateTimeChoice `xml:"AvlblDt,omitempty"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType56Choice `xml:"CorpActnEvtTp,omitempty"`

	// Balance from which the securities are moving.
	BalanceFrom *SecuritiesSubBalanceTypeAndQuantityBreakdown3 `xml:"BalFr"`

	// Balance to which the securities are moving.
	BalanceTo *SecuritiesSubBalanceTypeAndQuantityBreakdown3 `xml:"BalTo"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	InstructionProcessingAdditionalDetails *Max350Text `xml:"InstrPrcgAddtlDtls,omitempty"`
}

func (i *IntraPositionDetails41) AddSettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.SettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.SettledQuantity
}

func (i *IntraPositionDetails41) AddSecuritiesSubBalanceIdentification() *GenericIdentification37 {
	i.SecuritiesSubBalanceIdentification = new(GenericIdentification37)
	return i.SecuritiesSubBalanceIdentification
}

func (i *IntraPositionDetails41) AddCollateralMonitorAmount() *AmountAndDirection44 {
	i.CollateralMonitorAmount = new(AmountAndDirection44)
	return i.CollateralMonitorAmount
}

func (i *IntraPositionDetails41) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.PreviouslySettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.PreviouslySettledQuantity
}

func (i *IntraPositionDetails41) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.RemainingToBeSettledQuantity
}

func (i *IntraPositionDetails41) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails41) AddAvailableDate() *DateAndDateTimeChoice {
	i.AvailableDate = new(DateAndDateTimeChoice)
	return i.AvailableDate
}

func (i *IntraPositionDetails41) AddCorporateActionEventType() *CorporateActionEventType56Choice {
	i.CorporateActionEventType = new(CorporateActionEventType56Choice)
	return i.CorporateActionEventType
}

func (i *IntraPositionDetails41) AddBalanceFrom() *SecuritiesSubBalanceTypeAndQuantityBreakdown3 {
	i.BalanceFrom = new(SecuritiesSubBalanceTypeAndQuantityBreakdown3)
	return i.BalanceFrom
}

func (i *IntraPositionDetails41) AddBalanceTo() *SecuritiesSubBalanceTypeAndQuantityBreakdown3 {
	i.BalanceTo = new(SecuritiesSubBalanceTypeAndQuantityBreakdown3)
	return i.BalanceTo
}

func (i *IntraPositionDetails41) SetInstructionProcessingAdditionalDetails(value string) {
	i.InstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

// Details of the intra-position movement.
type IntraPositionDetails42 struct {

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification103Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount30 `xml:"SfkpgAcct"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *FinancialInstrumentQuantity1Choice `xml:"SttlmQty"`

	// Number identifying a lot constituting the financial instrument.
	LotNumber *GenericIdentification39 `xml:"LotNb,omitempty"`

	// Date and time at which the securities are to be moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Time stamp on when the transaction is acknowledged.
	AcknowledgedStatusTimeStamp *ISODateTime `xml:"AckdStsTmStmp,omitempty"`

	// Balance from which the securities are moving.
	BalanceFrom *SecuritiesBalanceType11Choice `xml:"BalFr,omitempty"`

	// Balance to which the securities are moving.
	BalanceTo *SecuritiesBalanceType11Choice `xml:"BalTo,omitempty"`
}

func (i *IntraPositionDetails42) SetPoolIdentification(value string) {
	i.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (i *IntraPositionDetails42) AddAccountOwner() *PartyIdentification103Choice {
	i.AccountOwner = new(PartyIdentification103Choice)
	return i.AccountOwner
}

func (i *IntraPositionDetails42) AddSafekeepingAccount() *SecuritiesAccount30 {
	i.SafekeepingAccount = new(SecuritiesAccount30)
	return i.SafekeepingAccount
}

func (i *IntraPositionDetails42) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	i.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return i.FinancialInstrumentIdentification
}

func (i *IntraPositionDetails42) AddSettlementQuantity() *FinancialInstrumentQuantity1Choice {
	i.SettlementQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.SettlementQuantity
}

func (i *IntraPositionDetails42) AddLotNumber() *GenericIdentification39 {
	i.LotNumber = new(GenericIdentification39)
	return i.LotNumber
}

func (i *IntraPositionDetails42) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails42) SetAcknowledgedStatusTimeStamp(value string) {
	i.AcknowledgedStatusTimeStamp = (*ISODateTime)(&value)
}

func (i *IntraPositionDetails42) AddBalanceFrom() *SecuritiesBalanceType11Choice {
	i.BalanceFrom = new(SecuritiesBalanceType11Choice)
	return i.BalanceFrom
}

func (i *IntraPositionDetails42) AddBalanceTo() *SecuritiesBalanceType11Choice {
	i.BalanceTo = new(SecuritiesBalanceType11Choice)
	return i.BalanceTo
}

// Details of the intra-position movement.
type IntraPositionDetails43 struct {

	// Quantity of financial instrument effectively settled.
	SettledQuantity *FinancialInstrumentQuantity15Choice `xml:"SttldQty"`

	// Number identifying a Securities Sub balance Type (example restriction identification etc…).
	SecuritiesSubBalanceIdentification *GenericIdentification39 `xml:"SctiesSubBalId,omitempty"`

	// Value of the collateral available for the delivery settlement process at the account level.
	CollateralMonitorAmount *AmountAndDirection55 `xml:"CollMntrAmt,omitempty"`

	// Quantity of financial instrument previously settled.
	PreviouslySettledQuantity *FinancialInstrumentQuantity15Choice `xml:"PrevslySttldQty,omitempty"`

	// Quantity of financial instrument remaining to be settled.
	RemainingToBeSettledQuantity *FinancialInstrumentQuantity15Choice `xml:"RmngToBeSttldQty,omitempty"`

	// Date and time at which the securities were moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Date/time securities become available for sale (if securities become unavailable, this specifies the date/time at which they will become available again).
	AvailableDate *DateAndDateTimeChoice `xml:"AvlblDt,omitempty"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType69Choice `xml:"CorpActnEvtTp,omitempty"`

	// Balance from which the securities are moving.
	BalanceFrom *SecuritiesSubBalanceTypeAndQuantityBreakdown4 `xml:"BalFr"`

	// Balance to which the securities are moving.
	BalanceTo *SecuritiesSubBalanceTypeAndQuantityBreakdown4 `xml:"BalTo"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	InstructionProcessingAdditionalDetails *RestrictedFINXMax350Text `xml:"InstrPrcgAddtlDtls,omitempty"`
}

func (i *IntraPositionDetails43) AddSettledQuantity() *FinancialInstrumentQuantity15Choice {
	i.SettledQuantity = new(FinancialInstrumentQuantity15Choice)
	return i.SettledQuantity
}

func (i *IntraPositionDetails43) AddSecuritiesSubBalanceIdentification() *GenericIdentification39 {
	i.SecuritiesSubBalanceIdentification = new(GenericIdentification39)
	return i.SecuritiesSubBalanceIdentification
}

func (i *IntraPositionDetails43) AddCollateralMonitorAmount() *AmountAndDirection55 {
	i.CollateralMonitorAmount = new(AmountAndDirection55)
	return i.CollateralMonitorAmount
}

func (i *IntraPositionDetails43) AddPreviouslySettledQuantity() *FinancialInstrumentQuantity15Choice {
	i.PreviouslySettledQuantity = new(FinancialInstrumentQuantity15Choice)
	return i.PreviouslySettledQuantity
}

func (i *IntraPositionDetails43) AddRemainingToBeSettledQuantity() *FinancialInstrumentQuantity15Choice {
	i.RemainingToBeSettledQuantity = new(FinancialInstrumentQuantity15Choice)
	return i.RemainingToBeSettledQuantity
}

func (i *IntraPositionDetails43) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails43) AddAvailableDate() *DateAndDateTimeChoice {
	i.AvailableDate = new(DateAndDateTimeChoice)
	return i.AvailableDate
}

func (i *IntraPositionDetails43) AddCorporateActionEventType() *CorporateActionEventType69Choice {
	i.CorporateActionEventType = new(CorporateActionEventType69Choice)
	return i.CorporateActionEventType
}

func (i *IntraPositionDetails43) AddBalanceFrom() *SecuritiesSubBalanceTypeAndQuantityBreakdown4 {
	i.BalanceFrom = new(SecuritiesSubBalanceTypeAndQuantityBreakdown4)
	return i.BalanceFrom
}

func (i *IntraPositionDetails43) AddBalanceTo() *SecuritiesSubBalanceTypeAndQuantityBreakdown4 {
	i.BalanceTo = new(SecuritiesSubBalanceTypeAndQuantityBreakdown4)
	return i.BalanceTo
}

func (i *IntraPositionDetails43) SetInstructionProcessingAdditionalDetails(value string) {
	i.InstructionProcessingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

// Details of the intra-position movement.
type IntraPositionDetails44 struct {

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat17Choice `xml:"SfkpgPlc,omitempty"`

	// Balance from which the securities were moved.
	BalanceFrom *SecuritiesBalanceType8Choice `xml:"BalFr"`

	// Intra-position movement(s) having been performed.
	IntraPositionMovement []*IntraPositionMovementDetails14 `xml:"IntraPosMvmnt"`
}

func (i *IntraPositionDetails44) AddSafekeepingPlace() *SafekeepingPlaceFormat17Choice {
	i.SafekeepingPlace = new(SafekeepingPlaceFormat17Choice)
	return i.SafekeepingPlace
}

func (i *IntraPositionDetails44) AddBalanceFrom() *SecuritiesBalanceType8Choice {
	i.BalanceFrom = new(SecuritiesBalanceType8Choice)
	return i.BalanceFrom
}

func (i *IntraPositionDetails44) AddIntraPositionMovement() *IntraPositionMovementDetails14 {
	newValue := new(IntraPositionMovementDetails14)
	i.IntraPositionMovement = append(i.IntraPositionMovement, newValue)
	return newValue
}

// Details of the intra-position movement.
type IntraPositionDetails9 struct {

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Quantity of financial instrument effectively settled.
	SettledQuantity *FinancialInstrumentQuantity1Choice `xml:"SttldQty"`

	// Date and time at which the securities are to be moved.
	SettlementDate *DateAndDateTimeChoice `xml:"SttlmDt"`

	// Balance from which the securities are moving.
	BalanceFrom *SecuritiesBalanceType2Choice `xml:"BalFr,omitempty"`

	// Balance to which the securities are moving.
	BalanceTo *SecuritiesBalanceType2Choice `xml:"BalTo,omitempty"`
}

func (i *IntraPositionDetails9) SetPoolIdentification(value string) {
	i.PoolIdentification = (*Max35Text)(&value)
}

func (i *IntraPositionDetails9) AddAccountOwner() *PartyIdentification36Choice {
	i.AccountOwner = new(PartyIdentification36Choice)
	return i.AccountOwner
}

func (i *IntraPositionDetails9) AddSafekeepingAccount() *SecuritiesAccount13 {
	i.SafekeepingAccount = new(SecuritiesAccount13)
	return i.SafekeepingAccount
}

func (i *IntraPositionDetails9) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	i.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return i.FinancialInstrumentIdentification
}

func (i *IntraPositionDetails9) AddSettledQuantity() *FinancialInstrumentQuantity1Choice {
	i.SettledQuantity = new(FinancialInstrumentQuantity1Choice)
	return i.SettledQuantity
}

func (i *IntraPositionDetails9) AddSettlementDate() *DateAndDateTimeChoice {
	i.SettlementDate = new(DateAndDateTimeChoice)
	return i.SettlementDate
}

func (i *IntraPositionDetails9) AddBalanceFrom() *SecuritiesBalanceType2Choice {
	i.BalanceFrom = new(SecuritiesBalanceType2Choice)
	return i.BalanceFrom
}

func (i *IntraPositionDetails9) AddBalanceTo() *SecuritiesBalanceType2Choice {
	i.BalanceTo = new(SecuritiesBalanceType2Choice)
	return i.BalanceTo
}
