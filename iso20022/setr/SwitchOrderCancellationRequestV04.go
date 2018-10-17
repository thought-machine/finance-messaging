package setr

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document01400104 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:setr.014.001.04 Document"`
	Message *SwitchOrderCancellationRequestV04 `xml:"SwtchOrdrCxlReq"`
}

func (d *Document01400104) AddMessage() *SwitchOrderCancellationRequestV04 {
	d.Message = new(SwitchOrderCancellationRequestV04)
	return d.Message
}

// Scope
// The SwitchOrderCancellationRequest message is sent by an instructing party, for example, an investment manager or its authorised representative, to the executing party, for example, a transfer agent, to request the cancellation of a previously sent SwitchOrder instruction.
// Usage
// The SwitchOrderCancellationRequest is used to cancel the entire previously sent SwitchOrder instruction/s and all the individual legs that it contains. There is no amendment, but a cancellation and re-instruct policy.
// To request the cancellation of a switch order, the order reference of the original switch order is quoted in the order reference element. The message identification of the SwitchOrder message may also be quoted in PreviousReference but this is not recommended.
// The deadline and acceptance of a cancellation request is subject to a service level agreement (SLA). This cancellation message is a cancellation request. There is no automatic acceptance of the cancellation request.
// The rejection or acceptance of a SwitchOrderCancellationRequest is made using an OrderCancellationStatusReport message.
type SwitchOrderCancellationRequestV04 struct {

	// Reference that uniquely identifies the message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference9 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference8 `xml:"PrvsRef,omitempty"`

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *iso20022.Max35Text `xml:"MstrRef,omitempty"`

	// Identification of the individual order to be cancelled.
	OrderReferences []*iso20022.InvestmentFundOrder9 `xml:"OrdrRefs"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation4 `xml:"CpyDtls,omitempty"`
}

func (s *SwitchOrderCancellationRequestV04) AddMessageIdentification() *iso20022.MessageIdentification1 {
	s.MessageIdentification = new(iso20022.MessageIdentification1)
	return s.MessageIdentification
}

func (s *SwitchOrderCancellationRequestV04) AddPoolReference() *iso20022.AdditionalReference9 {
	s.PoolReference = new(iso20022.AdditionalReference9)
	return s.PoolReference
}

func (s *SwitchOrderCancellationRequestV04) AddPreviousReference() *iso20022.AdditionalReference8 {
	s.PreviousReference = new(iso20022.AdditionalReference8)
	return s.PreviousReference
}

func (s *SwitchOrderCancellationRequestV04) SetMasterReference(value string) {
	s.MasterReference = (*iso20022.Max35Text)(&value)
}

func (s *SwitchOrderCancellationRequestV04) AddOrderReferences() *iso20022.InvestmentFundOrder9 {
	newValue := new(iso20022.InvestmentFundOrder9)
	s.OrderReferences = append(s.OrderReferences, newValue)
	return newValue
}

func (s *SwitchOrderCancellationRequestV04) AddCopyDetails() *iso20022.CopyInformation4 {
	s.CopyDetails = new(iso20022.CopyInformation4)
	return s.CopyDetails
}
