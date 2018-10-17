package setr

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document04800101 struct {
	XMLName xml.Name                                   `xml:"urn:iso:std:iso:20022:tech:xsd:setr.048.001.01 Document"`
	Message *SubscriptionOrderConfirmationAmendmentV01 `xml:"SbcptOrdrConfAmdmntV01"`
}

func (d *Document04800101) AddMessage() *SubscriptionOrderConfirmationAmendmentV01 {
	d.Message = new(SubscriptionOrderConfirmationAmendmentV01)
	return d.Message
}

// Scope
// An executing party, for example, a transfer agent, sends the SubscriptionOrderConfirmationAmendment message to the instructing party, for example, an investment manager or its authorised representative to amend a previously sent SubscriptionOrderConfirmation.
// Usage
// The SubscriptionOrderConfirmationAmendment message is used to amend one or more previously sent subscription order confirmations.
// Each individual order confirmation amendment specified is identified in DealReference. The reference of the original individual order is specified in OrderReference.
// The message identification of the SubscriptionOrder message in which the individual orders were conveyed may also be quoted in RelatedReference. The message identification of the SubscriptionOrderConfirmation message in which the original order confirmations were conveyed may also be quoted in PreviousReference.
type SubscriptionOrderConfirmationAmendmentV01 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// General information related to the execution of investment fund order.
	MultipleExecutionDetails *iso20022.SubscriptionMultipleExecution3 `xml:"MltplExctnDtls"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*iso20022.Intermediary9 `xml:"RltdPtyDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation2 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SubscriptionOrderConfirmationAmendmentV01) AddMessageIdentification() *iso20022.MessageIdentification1 {
	s.MessageIdentification = new(iso20022.MessageIdentification1)
	return s.MessageIdentification
}

func (s *SubscriptionOrderConfirmationAmendmentV01) AddPoolReference() *iso20022.AdditionalReference3 {
	s.PoolReference = new(iso20022.AdditionalReference3)
	return s.PoolReference
}

func (s *SubscriptionOrderConfirmationAmendmentV01) AddPreviousReference() *iso20022.AdditionalReference3 {
	newValue := new(iso20022.AdditionalReference3)
	s.PreviousReference = append(s.PreviousReference, newValue)
	return newValue
}

func (s *SubscriptionOrderConfirmationAmendmentV01) AddRelatedReference() *iso20022.AdditionalReference3 {
	s.RelatedReference = new(iso20022.AdditionalReference3)
	return s.RelatedReference
}

func (s *SubscriptionOrderConfirmationAmendmentV01) AddMultipleExecutionDetails() *iso20022.SubscriptionMultipleExecution3 {
	s.MultipleExecutionDetails = new(iso20022.SubscriptionMultipleExecution3)
	return s.MultipleExecutionDetails
}

func (s *SubscriptionOrderConfirmationAmendmentV01) AddRelatedPartyDetails() *iso20022.Intermediary9 {
	newValue := new(iso20022.Intermediary9)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SubscriptionOrderConfirmationAmendmentV01) AddCopyDetails() *iso20022.CopyInformation2 {
	s.CopyDetails = new(iso20022.CopyInformation2)
	return s.CopyDetails
}

func (s *SubscriptionOrderConfirmationAmendmentV01) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
