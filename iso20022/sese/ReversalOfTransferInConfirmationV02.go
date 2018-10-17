package sese

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document00800102 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.008.001.02 Document"`
	Message *ReversalOfTransferInConfirmationV02 `xml:"RvslOfTrfInConfV02"`
}

func (d *Document00800102) AddMessage() *ReversalOfTransferInConfirmationV02 {
	d.Message = new(ReversalOfTransferInConfirmationV02)
	return d.Message
}

// Scope
// An executing party, eg, a transfer agent, sends the ReversalOfTransferInConfirmation message to the instructing party, eg, an investment manager or its authorised representative, to cancel a previously sent TransferInConfirmation message.
// Usage
// The ReversalOfTransferInConfirmation message is used to reverse a previously sent TransferInConfirmation.
// There are two ways to specify the reversal of the transfer in confirmation. Either:
// - the business references, eg, TransferReference, TransferConfirmationIdentification, of the transfer confirmation are quoted, or,
// - all the details of the transfer confirmation (this includes TransferReference and TransferConfirmationIdentification) are quoted but this is not recommended.
// The message identification of the TransferInConfirmation message in which the transfer confirmation was conveyed may also be quoted in PreviousReference.
// The message reference (MessageIdentification) of the TransferInInstruction message in which the transfer instruction was conveyed may also be quoted in RelatedReference.
// It is also possible to request a reversal of a TransferInConfirmation by quoting its message reference (MessageIdentification) in PreviousReference.
type ReversalOfTransferInConfirmationV02 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference2 `xml:"PrvsRef,omitempty"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference2 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference2 `xml:"RltdRef,omitempty"`

	// Reference of the transfer in confirmation to be reversed.
	ReversalByReference *iso20022.TransferReference2 `xml:"RvslByRef,omitempty"`

	// Copy of the transfer in confirmation to reverse.
	ReversalByTransferInConfirmationDetails *iso20022.TransferIn4 `xml:"RvslByTrfInConfDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation2 `xml:"CpyDtls,omitempty"`
}

func (r *ReversalOfTransferInConfirmationV02) AddMessageIdentification() *iso20022.MessageIdentification1 {
	r.MessageIdentification = new(iso20022.MessageIdentification1)
	return r.MessageIdentification
}

func (r *ReversalOfTransferInConfirmationV02) AddPreviousReference() *iso20022.AdditionalReference2 {
	r.PreviousReference = new(iso20022.AdditionalReference2)
	return r.PreviousReference
}

func (r *ReversalOfTransferInConfirmationV02) AddPoolReference() *iso20022.AdditionalReference2 {
	r.PoolReference = new(iso20022.AdditionalReference2)
	return r.PoolReference
}

func (r *ReversalOfTransferInConfirmationV02) AddRelatedReference() *iso20022.AdditionalReference2 {
	r.RelatedReference = new(iso20022.AdditionalReference2)
	return r.RelatedReference
}

func (r *ReversalOfTransferInConfirmationV02) AddReversalByReference() *iso20022.TransferReference2 {
	r.ReversalByReference = new(iso20022.TransferReference2)
	return r.ReversalByReference
}

func (r *ReversalOfTransferInConfirmationV02) AddReversalByTransferInConfirmationDetails() *iso20022.TransferIn4 {
	r.ReversalByTransferInConfirmationDetails = new(iso20022.TransferIn4)
	return r.ReversalByTransferInConfirmationDetails
}

func (r *ReversalOfTransferInConfirmationV02) AddCopyDetails() *iso20022.CopyInformation2 {
	r.CopyDetails = new(iso20022.CopyInformation2)
	return r.CopyDetails
}
