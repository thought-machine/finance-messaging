package iso20022

// Status report of the individual orders of a bulk or multiple order that was previously received.
type IndividualOrderStatusAndReason1 struct {

	// Status of the order is accepted or already executed or sent to next party or received. There is no reason attached.
	Status *OrderStatus2Code `xml:"Sts"`

	// Status of the individual order details is cancelled.
	Cancelled *CancelledStatus1 `xml:"Canc"`

	// Status of the individual order details is conditionally accepted.
	ConditionallyAccepted *ConditionallyAcceptedStatus1 `xml:"CondlyAccptd"`

	// Status of the individual order details is in repair.
	InRepair *InRepairStatus1 `xml:"InRpr"`

	// Status of the individual order details is rejected.
	Rejected *RejectedStatus3 `xml:"Rjctd"`

	// Status of the individual order details is suspended.
	Suspended *SuspendedStatus1 `xml:"Sspd"`

	// Elements from the original individual order details that have been repaired so that the order can be accepted.
	RepairedConditions *RepairedConditions2 `xml:"RprdConds"`

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Party that initiates the status of the order.
	StatusInitiator *PartyIdentification2Choice `xml:"StsInitr,omitempty"`

	// Choice between the investment account and the financial instrument.
	InvestmentAccountOrFinancialInstrument *InvestmentAccountOrFinancialInstrument1Choice `xml:"InvstmtAcctOrFinInstrm,omitempty"`

	// Information that has been added to the original order.
	NewDetails *ExpectedExecutionDetails1 `xml:"NewDtls,omitempty"`
}

func (i *IndividualOrderStatusAndReason1) SetStatus(value string) {
	i.Status = (*OrderStatus2Code)(&value)
}

func (i *IndividualOrderStatusAndReason1) AddCancelled() *CancelledStatus1 {
	i.Cancelled = new(CancelledStatus1)
	return i.Cancelled
}

func (i *IndividualOrderStatusAndReason1) AddConditionallyAccepted() *ConditionallyAcceptedStatus1 {
	i.ConditionallyAccepted = new(ConditionallyAcceptedStatus1)
	return i.ConditionallyAccepted
}

func (i *IndividualOrderStatusAndReason1) AddInRepair() *InRepairStatus1 {
	i.InRepair = new(InRepairStatus1)
	return i.InRepair
}

func (i *IndividualOrderStatusAndReason1) AddRejected() *RejectedStatus3 {
	i.Rejected = new(RejectedStatus3)
	return i.Rejected
}

func (i *IndividualOrderStatusAndReason1) AddSuspended() *SuspendedStatus1 {
	i.Suspended = new(SuspendedStatus1)
	return i.Suspended
}

func (i *IndividualOrderStatusAndReason1) AddRepairedConditions() *RepairedConditions2 {
	i.RepairedConditions = new(RepairedConditions2)
	return i.RepairedConditions
}

func (i *IndividualOrderStatusAndReason1) SetOrderReference(value string) {
	i.OrderReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason1) AddStatusInitiator() *PartyIdentification2Choice {
	i.StatusInitiator = new(PartyIdentification2Choice)
	return i.StatusInitiator
}

func (i *IndividualOrderStatusAndReason1) AddInvestmentAccountOrFinancialInstrument() *InvestmentAccountOrFinancialInstrument1Choice {
	i.InvestmentAccountOrFinancialInstrument = new(InvestmentAccountOrFinancialInstrument1Choice)
	return i.InvestmentAccountOrFinancialInstrument
}

func (i *IndividualOrderStatusAndReason1) AddNewDetails() *ExpectedExecutionDetails1 {
	i.NewDetails = new(ExpectedExecutionDetails1)
	return i.NewDetails
}

// Status report of the individual orders of a bulk or multiple order that was previously received.
type IndividualOrderStatusAndReason2 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of an order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for an order execution, as assigned by a confirming party.
	DealReference *Max35Text `xml:"DealRef,omitempty"`

	// Unique and unambiguous identifier for an order cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Status of the individual order is accepted or already executed or sent to next party or received. There is no reason attached.
	Status *OrderStatus4Code `xml:"Sts"`

	// Status of the individual order is cancelled. This status is used for an order that has been accepted or that has been entered in an order book but that can not be executed.
	Cancelled *CancelledStatus2 `xml:"Canc"`

	// Status of the individual order is conditionally accepted.
	ConditionallyAccepted *ConditionallyAcceptedStatus2 `xml:"CondlyAccptd"`

	// Status of the individual order is rejected. This status is used for an order that has not been accepted or entered in an order book.
	Rejected []*RejectedStatus6 `xml:"Rjctd"`

	// Status of the individual order is suspended.
	Suspended *SuspendedStatus2 `xml:"Sspd"`

	// Status of the individual order is in repair.
	InRepair *InRepairStatus2 `xml:"InRpr"`

	// Status of the individual order is partially settled.
	PartiallySettled *PartiallySettledStatus1 `xml:"PrtlySttld"`

	// Elements from the original individual order that have been repaired so that the order can be accepted.
	RepairedConditions *RepairedConditions3 `xml:"RprdConds,omitempty"`

	// Party that initiates the status of the order cancellation.
	StatusInitiator *PartyIdentification2Choice `xml:"StsInitr,omitempty"`

	// Order data.
	OrderData *FundOrderData1 `xml:"OrdrData,omitempty"`

	// Information that has been added to the original order.
	NewDetails *ExpectedExecutionDetails2 `xml:"NewDtls,omitempty"`
}

func (i *IndividualOrderStatusAndReason2) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason2) SetOrderReference(value string) {
	i.OrderReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason2) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason2) SetDealReference(value string) {
	i.DealReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason2) SetCancellationReference(value string) {
	i.CancellationReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason2) SetStatus(value string) {
	i.Status = (*OrderStatus4Code)(&value)
}

func (i *IndividualOrderStatusAndReason2) AddCancelled() *CancelledStatus2 {
	i.Cancelled = new(CancelledStatus2)
	return i.Cancelled
}

func (i *IndividualOrderStatusAndReason2) AddConditionallyAccepted() *ConditionallyAcceptedStatus2 {
	i.ConditionallyAccepted = new(ConditionallyAcceptedStatus2)
	return i.ConditionallyAccepted
}

func (i *IndividualOrderStatusAndReason2) AddRejected() *RejectedStatus6 {
	newValue := new(RejectedStatus6)
	i.Rejected = append(i.Rejected, newValue)
	return newValue
}

func (i *IndividualOrderStatusAndReason2) AddSuspended() *SuspendedStatus2 {
	i.Suspended = new(SuspendedStatus2)
	return i.Suspended
}

func (i *IndividualOrderStatusAndReason2) AddInRepair() *InRepairStatus2 {
	i.InRepair = new(InRepairStatus2)
	return i.InRepair
}

func (i *IndividualOrderStatusAndReason2) AddPartiallySettled() *PartiallySettledStatus1 {
	i.PartiallySettled = new(PartiallySettledStatus1)
	return i.PartiallySettled
}

func (i *IndividualOrderStatusAndReason2) AddRepairedConditions() *RepairedConditions3 {
	i.RepairedConditions = new(RepairedConditions3)
	return i.RepairedConditions
}

func (i *IndividualOrderStatusAndReason2) AddStatusInitiator() *PartyIdentification2Choice {
	i.StatusInitiator = new(PartyIdentification2Choice)
	return i.StatusInitiator
}

func (i *IndividualOrderStatusAndReason2) AddOrderData() *FundOrderData1 {
	i.OrderData = new(FundOrderData1)
	return i.OrderData
}

func (i *IndividualOrderStatusAndReason2) AddNewDetails() *ExpectedExecutionDetails2 {
	i.NewDetails = new(ExpectedExecutionDetails2)
	return i.NewDetails
}

// Status report of an individual order of a bulk or multiple or switch order cancellation instruction that was previously received.
type IndividualOrderStatusAndReason4 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of an order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for an order cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Cancellation status of the order.
	Status *OrderCancellationStatus1Code `xml:"Sts"`

	// Status of the individual order cancellation request is rejected.
	Rejected *RejectedStatus7 `xml:"Rjctd"`

	// Party that initiates the status of the individual order cancellation.
	StatusInitiator *PartyIdentification2Choice `xml:"StsInitr,omitempty"`

	// Account information of the individual order cancellation for which the status is given.
	InvestmentAccountDetails *InvestmentAccount13 `xml:"InvstmtAcctDtls,omitempty"`

	// Financial instrument information of the individual order cancellation for which the status is given.
	FinancialInstrumentDetails *FinancialInstrument10 `xml:"FinInstrmDtls,omitempty"`
}

func (i *IndividualOrderStatusAndReason4) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason4) SetOrderReference(value string) {
	i.OrderReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason4) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason4) SetCancellationReference(value string) {
	i.CancellationReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason4) SetStatus(value string) {
	i.Status = (*OrderCancellationStatus1Code)(&value)
}

func (i *IndividualOrderStatusAndReason4) AddRejected() *RejectedStatus7 {
	i.Rejected = new(RejectedStatus7)
	return i.Rejected
}

func (i *IndividualOrderStatusAndReason4) AddStatusInitiator() *PartyIdentification2Choice {
	i.StatusInitiator = new(PartyIdentification2Choice)
	return i.StatusInitiator
}

func (i *IndividualOrderStatusAndReason4) AddInvestmentAccountDetails() *InvestmentAccount13 {
	i.InvestmentAccountDetails = new(InvestmentAccount13)
	return i.InvestmentAccountDetails
}

func (i *IndividualOrderStatusAndReason4) AddFinancialInstrumentDetails() *FinancialInstrument10 {
	i.FinancialInstrumentDetails = new(FinancialInstrument10)
	return i.FinancialInstrumentDetails
}

// Status report of the individual orders of a bulk or multiple order that was previously received.
type IndividualOrderStatusAndReason7 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identifier for the order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of the order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for the order execution, as assigned by the confirming party.
	DealReference *Max35Text `xml:"DealRef,omitempty"`

	// Unique and unambiguous identifier for the order cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Status of the individual order.
	OrderStatus *OrderStatus5Choice `xml:"OrdrSts"`

	// Elements from the original individual order that have been repaired so that the order can be accepted.
	RepairedFee []*Fee3 `xml:"RprdFee,omitempty"`

	// Party that initiates the status of the order.
	StatusInitiator *PartyIdentification113 `xml:"StsInitr,omitempty"`

	// Order data.
	OrderData *FundOrderData5 `xml:"OrdrData,omitempty"`

	// Expected execution information.
	NewDetails *ExpectedExecutionDetails4 `xml:"NewDtls,omitempty"`

	// Information about gating and hold back of redemption proceeds.
	GatingOrHoldBackDetails *HoldBackInformation3 `xml:"GtgOrHldBckDtls,omitempty"`
}

func (i *IndividualOrderStatusAndReason7) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason7) SetOrderReference(value string) {
	i.OrderReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason7) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason7) SetDealReference(value string) {
	i.DealReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason7) SetCancellationReference(value string) {
	i.CancellationReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason7) AddOrderStatus() *OrderStatus5Choice {
	i.OrderStatus = new(OrderStatus5Choice)
	return i.OrderStatus
}

func (i *IndividualOrderStatusAndReason7) AddRepairedFee() *Fee3 {
	newValue := new(Fee3)
	i.RepairedFee = append(i.RepairedFee, newValue)
	return newValue
}

func (i *IndividualOrderStatusAndReason7) AddStatusInitiator() *PartyIdentification113 {
	i.StatusInitiator = new(PartyIdentification113)
	return i.StatusInitiator
}

func (i *IndividualOrderStatusAndReason7) AddOrderData() *FundOrderData5 {
	i.OrderData = new(FundOrderData5)
	return i.OrderData
}

func (i *IndividualOrderStatusAndReason7) AddNewDetails() *ExpectedExecutionDetails4 {
	i.NewDetails = new(ExpectedExecutionDetails4)
	return i.NewDetails
}

func (i *IndividualOrderStatusAndReason7) AddGatingOrHoldBackDetails() *HoldBackInformation3 {
	i.GatingOrHoldBackDetails = new(HoldBackInformation3)
	return i.GatingOrHoldBackDetails
}

// Status report of an individual order of a bulk or multiple or switch order cancellation instruction that was previously received.
type IndividualOrderStatusAndReason8 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identifier for the order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of the order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for the order cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Cancellation status of the order cancellation.
	CancellationStatus *CancellationStatus22Choice `xml:"CxlSts"`

	// Party that initiates the status of the individual order cancellation.
	StatusInitiator *PartyIdentification113 `xml:"StsInitr,omitempty"`

	// Account information of the individual order cancellation for which the status is given.
	InvestmentAccountDetails *InvestmentAccount58 `xml:"InvstmtAcctDtls,omitempty"`

	// Financial instrument information of the individual order cancellation for which the status is given.
	FinancialInstrumentDetails *FinancialInstrument57 `xml:"FinInstrmDtls,omitempty"`
}

func (i *IndividualOrderStatusAndReason8) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason8) SetOrderReference(value string) {
	i.OrderReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason8) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason8) SetCancellationReference(value string) {
	i.CancellationReference = (*Max35Text)(&value)
}

func (i *IndividualOrderStatusAndReason8) AddCancellationStatus() *CancellationStatus22Choice {
	i.CancellationStatus = new(CancellationStatus22Choice)
	return i.CancellationStatus
}

func (i *IndividualOrderStatusAndReason8) AddStatusInitiator() *PartyIdentification113 {
	i.StatusInitiator = new(PartyIdentification113)
	return i.StatusInitiator
}

func (i *IndividualOrderStatusAndReason8) AddInvestmentAccountDetails() *InvestmentAccount58 {
	i.InvestmentAccountDetails = new(InvestmentAccount58)
	return i.InvestmentAccountDetails
}

func (i *IndividualOrderStatusAndReason8) AddFinancialInstrumentDetails() *FinancialInstrument57 {
	i.FinancialInstrumentDetails = new(FinancialInstrument57)
	return i.FinancialInstrumentDetails
}
