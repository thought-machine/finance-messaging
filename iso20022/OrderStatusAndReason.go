package iso20022

// Status report of a bulk or multiple or switch order that was previously received.
type OrderStatusAndReason10 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Status of a 'bulk' of orders. Can be used if all the individual orders conveyed in a bulk or multiple order message have the same status.
	OrderStatus *OrderStatus3Choice `xml:"OrdrSts"`

	// Party that initiates the status of the order.
	StatusInitiator *PartyIdentification113 `xml:"StsInitr,omitempty"`
}

func (o *OrderStatusAndReason10) SetMasterReference(value string) {
	o.MasterReference = (*Max35Text)(&value)
}

func (o *OrderStatusAndReason10) AddOrderStatus() *OrderStatus3Choice {
	o.OrderStatus = new(OrderStatus3Choice)
	return o.OrderStatus
}

func (o *OrderStatusAndReason10) AddStatusInitiator() *PartyIdentification113 {
	o.StatusInitiator = new(PartyIdentification113)
	return o.StatusInitiator
}

// Status report of a bulk or multiple or switch order that was previously received.
type OrderStatusAndReason3 struct {

	// Status of the order is accepted or already executed or sent to next party or received. There is no reason attached.
	Status *OrderStatus2Code `xml:"Sts"`

	// Status of the order details is cancelled. This status is used for orders that have been accepted or that have been entered in an order book but that can not be executed.
	Cancelled *CancelledStatus1 `xml:"Canc"`

	// Status of the order details is conditionally accepted.
	ConditionallyAccepted *ConditionallyAcceptedStatus1 `xml:"CondlyAccptd"`

	// Status of the order details is rejected. This status is used for orders that have not been accepted or entered in an order book.
	Rejected *RejectedStatus3 `xml:"Rjctd"`

	// Status of the order details is suspended.
	Suspended *SuspendedStatus1 `xml:"Sspd"`

	// Party that initiates the status of the order.
	StatusInitiator *PartyIdentification2Choice `xml:"StsInitr,omitempty"`

	// Unique and unambiguous technical identification of an instance of a leg within a switch.
	SwitchOrderLegIdentification *Max35Text `xml:"SwtchOrdrLegId,omitempty"`
}

func (o *OrderStatusAndReason3) SetStatus(value string) {
	o.Status = (*OrderStatus2Code)(&value)
}

func (o *OrderStatusAndReason3) AddCancelled() *CancelledStatus1 {
	o.Cancelled = new(CancelledStatus1)
	return o.Cancelled
}

func (o *OrderStatusAndReason3) AddConditionallyAccepted() *ConditionallyAcceptedStatus1 {
	o.ConditionallyAccepted = new(ConditionallyAcceptedStatus1)
	return o.ConditionallyAccepted
}

func (o *OrderStatusAndReason3) AddRejected() *RejectedStatus3 {
	o.Rejected = new(RejectedStatus3)
	return o.Rejected
}

func (o *OrderStatusAndReason3) AddSuspended() *SuspendedStatus1 {
	o.Suspended = new(SuspendedStatus1)
	return o.Suspended
}

func (o *OrderStatusAndReason3) AddStatusInitiator() *PartyIdentification2Choice {
	o.StatusInitiator = new(PartyIdentification2Choice)
	return o.StatusInitiator
}

func (o *OrderStatusAndReason3) SetSwitchOrderLegIdentification(value string) {
	o.SwitchOrderLegIdentification = (*Max35Text)(&value)
}

// Status report of a bulk or multiple or switch order cancellation instruction that was previously received.
type OrderStatusAndReason4 struct {

	// Status of the order.
	Status *OrderStatus3Code `xml:"Sts"`

	// Status of the order is rejected.
	Rejected *RejectedStatus4 `xml:"Rjctd"`

	// Party that initiates the status of the order.
	StatusInitiator *PartyIdentification2Choice `xml:"StsInitr,omitempty"`
}

func (o *OrderStatusAndReason4) SetStatus(value string) {
	o.Status = (*OrderStatus3Code)(&value)
}

func (o *OrderStatusAndReason4) AddRejected() *RejectedStatus4 {
	o.Rejected = new(RejectedStatus4)
	return o.Rejected
}

func (o *OrderStatusAndReason4) AddStatusInitiator() *PartyIdentification2Choice {
	o.StatusInitiator = new(PartyIdentification2Choice)
	return o.StatusInitiator
}

// Status report of a bulk or multiple or switch order that was previously received.
type OrderStatusAndReason7 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Status of all the orders in the order message. There is no reason attached.
	Status *OrderStatus4Code `xml:"Sts"`

	// Status of all the orders in the order message is cancelled. This status is used for orders that have been accepted or that have been entered in an order book but that can not be executed.
	Cancelled *CancelledStatus2 `xml:"Canc"`

	// Status of all the orders in the order message is conditionally accepted.
	ConditionallyAccepted *ConditionallyAcceptedStatus2 `xml:"CondlyAccptd"`

	// Status of all the orders in the order message is rejected. This status is used for orders that have not been accepted or entered in an order book.
	Rejected []*RejectedStatus6 `xml:"Rjctd"`

	// Status of all the orders in the order message is suspended.
	Suspended *SuspendedStatus2 `xml:"Sspd"`

	// Status of all the orders in the order message is partially settled.
	PartiallySettled *PartiallySettledStatus1 `xml:"PrtlySttld"`

	// Party that initiates the status of the order.
	StatusInitiator *PartyIdentification2Choice `xml:"StsInitr,omitempty"`
}

func (o *OrderStatusAndReason7) SetMasterReference(value string) {
	o.MasterReference = (*Max35Text)(&value)
}

func (o *OrderStatusAndReason7) SetStatus(value string) {
	o.Status = (*OrderStatus4Code)(&value)
}

func (o *OrderStatusAndReason7) AddCancelled() *CancelledStatus2 {
	o.Cancelled = new(CancelledStatus2)
	return o.Cancelled
}

func (o *OrderStatusAndReason7) AddConditionallyAccepted() *ConditionallyAcceptedStatus2 {
	o.ConditionallyAccepted = new(ConditionallyAcceptedStatus2)
	return o.ConditionallyAccepted
}

func (o *OrderStatusAndReason7) AddRejected() *RejectedStatus6 {
	newValue := new(RejectedStatus6)
	o.Rejected = append(o.Rejected, newValue)
	return newValue
}

func (o *OrderStatusAndReason7) AddSuspended() *SuspendedStatus2 {
	o.Suspended = new(SuspendedStatus2)
	return o.Suspended
}

func (o *OrderStatusAndReason7) AddPartiallySettled() *PartiallySettledStatus1 {
	o.PartiallySettled = new(PartiallySettledStatus1)
	return o.PartiallySettled
}

func (o *OrderStatusAndReason7) AddStatusInitiator() *PartyIdentification2Choice {
	o.StatusInitiator = new(PartyIdentification2Choice)
	return o.StatusInitiator
}

// Status report of a bulk or multiple or switch order cancellation instruction that was previously received.
type OrderStatusAndReason8 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Cancellation status of the order.
	Status *OrderCancellationStatus1Code `xml:"Sts"`

	// Status of the order cancellation request is rejected.
	Rejected *RejectedStatus7 `xml:"Rjctd"`

	// Party that initiates the status of the order cancellation.
	StatusInitiator *PartyIdentification2Choice `xml:"StsInitr,omitempty"`
}

func (o *OrderStatusAndReason8) SetMasterReference(value string) {
	o.MasterReference = (*Max35Text)(&value)
}

func (o *OrderStatusAndReason8) SetStatus(value string) {
	o.Status = (*OrderCancellationStatus1Code)(&value)
}

func (o *OrderStatusAndReason8) AddRejected() *RejectedStatus7 {
	o.Rejected = new(RejectedStatus7)
	return o.Rejected
}

func (o *OrderStatusAndReason8) AddStatusInitiator() *PartyIdentification2Choice {
	o.StatusInitiator = new(PartyIdentification2Choice)
	return o.StatusInitiator
}

// Status report of a bulk or multiple or switch order cancellation instruction that was previously received.
type OrderStatusAndReason9 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Cancellation status of the order cancellation.
	CancellationStatus *CancellationStatus22Choice `xml:"CxlSts"`

	// Party that initiates the status of the order cancellation.
	StatusInitiator *PartyIdentification113 `xml:"StsInitr,omitempty"`
}

func (o *OrderStatusAndReason9) SetMasterReference(value string) {
	o.MasterReference = (*Max35Text)(&value)
}

func (o *OrderStatusAndReason9) AddCancellationStatus() *CancellationStatus22Choice {
	o.CancellationStatus = new(CancellationStatus22Choice)
	return o.CancellationStatus
}

func (o *OrderStatusAndReason9) AddStatusInitiator() *PartyIdentification113 {
	o.StatusInitiator = new(PartyIdentification113)
	return o.StatusInitiator
}
