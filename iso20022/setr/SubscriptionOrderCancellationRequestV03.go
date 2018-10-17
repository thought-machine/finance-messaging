package setr

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document01100103 struct {
	XMLName xml.Name                                 `xml:"urn:iso:std:iso:20022:tech:xsd:setr.011.001.03 Document"`
	Message *SubscriptionOrderCancellationRequestV03 `xml:"SbcptOrdrCxlReqV03"`
}

func (d *Document01100103) AddMessage() *SubscriptionOrderCancellationRequestV03 {
	d.Message = new(SubscriptionOrderCancellationRequestV03)
	return d.Message
}

// Scope
// An instructing party, for example, an investment manager or its authorised representative, sends the SubscriptionOrderCancellationRequest message to the executing party, for example, a transfer agent, to request the cancellation of a previously sent SubscriptionOrder.
// Usage
// The SubscriptionOrderCancellationRequest message is used to either:
// - request the cancellation of an entire SubscriptionOrder message, that is, all the individual orders that it contained, or,
// - request the cancellation of one or more individual orders.
// There is no amendment, but a cancellation and re-instruct policy.
// There are two ways to use the message:
// (1) When the SubscriptionOrderCancellationRequest message is used to request the cancellation of an entire SubscriptionOrder message, this can be done by either:
// - quoting the order references of all the individual orders listed in the SubscriptionOrder message, or,
// - quoting the details of all the individual orders (this includes the OrderReference) listed in SubscriptionOrder message, but this is not recommended.
// The message identification of the SubscriptionOrder message may also be quoted in PreviousReference.
// It is also possible to request the cancellation of an entire SubscriptionOrder message by quoting its message identification in PreviousReference, but this is not recommended.
// (2) When the SubscriptionOrderCancellationRequest message is used to request the cancellation of one or more individual orders, this can be done by either:
// - quoting the OrderReference of each individual order listed in the SubscriptionOrder message, or,
// - quoting the details of each individual order (including the OrderReference) listed in SubscriptionOrder message, but this is not recommended.
// The message identification of the SubscriptionOrder message in which the individual order was conveyed may also be quoted in PreviousReference.
// The deadline and acceptance of a cancellation request is subject to a service level agreement (SLA). This cancellation message is a cancellation request. There is no automatic acceptance of the cancellation.
// The rejection or acceptance of a SubscriptionOrderCancellationRequest is made using an OrderCancellationStatusReport message.
type SubscriptionOrderCancellationRequestV03 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// References of the orders to be cancelled.
	CancellationByReference *iso20022.InvestmentFundOrder1 `xml:"CxlByRef,omitempty"`

	// Common information related to all the orders to be cancelled.
	CancellationByOrderDetails *iso20022.SubscriptionMultipleOrderInstruction2 `xml:"CxlByOrdrDtls,omitempty"`

	// Message is a copy.
	CopyDetails *iso20022.CopyInformation2 `xml:"CpyDtls,omitempty"`
}

func (s *SubscriptionOrderCancellationRequestV03) AddMessageIdentification() *iso20022.MessageIdentification1 {
	s.MessageIdentification = new(iso20022.MessageIdentification1)
	return s.MessageIdentification
}

func (s *SubscriptionOrderCancellationRequestV03) AddPoolReference() *iso20022.AdditionalReference3 {
	s.PoolReference = new(iso20022.AdditionalReference3)
	return s.PoolReference
}

func (s *SubscriptionOrderCancellationRequestV03) AddPreviousReference() *iso20022.AdditionalReference3 {
	s.PreviousReference = new(iso20022.AdditionalReference3)
	return s.PreviousReference
}

func (s *SubscriptionOrderCancellationRequestV03) AddCancellationByReference() *iso20022.InvestmentFundOrder1 {
	s.CancellationByReference = new(iso20022.InvestmentFundOrder1)
	return s.CancellationByReference
}

func (s *SubscriptionOrderCancellationRequestV03) AddCancellationByOrderDetails() *iso20022.SubscriptionMultipleOrderInstruction2 {
	s.CancellationByOrderDetails = new(iso20022.SubscriptionMultipleOrderInstruction2)
	return s.CancellationByOrderDetails
}

func (s *SubscriptionOrderCancellationRequestV03) AddCopyDetails() *iso20022.CopyInformation2 {
	s.CopyDetails = new(iso20022.CopyInformation2)
	return s.CopyDetails
}
