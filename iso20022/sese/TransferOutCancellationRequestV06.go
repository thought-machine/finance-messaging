package sese

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document00200106 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.002.001.06 Document"`
	Message *TransferOutCancellationRequestV06 `xml:"TrfOutCxlReq"`
}

func (d *Document00200106) AddMessage() *TransferOutCancellationRequestV06 {
	d.Message = new(TransferOutCancellationRequestV06)
	return d.Message
}

// Scope
// An instructing party, for example, an investment manager or its authorised representative, sends the TransferOutCancellationRequest message to the executing party, for example, a transfer agent, to request the cancellation of a previously sent TransferOutInstruction.
// Usage
// The TransferOutCancellationRequest message is used to request cancellation of a previously sent TransferOutInstruction. There are two ways to specify the transfer cancellation request. Either:
// - the transfer reference of the original transfer is quoted, or,
// - all the details of the original transfer (this includes TransferReference) are quoted but this is not recommended.
// The message identification of the TransferOutInstruction message in which the original transfer was conveyed may also be quoted in PreviousReference. It is also possible to request the cancellation of a TransferOutInstruction message by quoting its message identification in PreviousReference.
type TransferOutCancellationRequestV06 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Reference to the transaction identifier issued by the counterparty. Building block may also be used to reference a previous transaction, or tie a set of messages together.
	References []*iso20022.References15 `xml:"Refs,omitempty"`

	// Choice between cancellation by reference or by transfer details.
	Cancellation []*iso20022.Cancellation8Choice `xml:"Cxl"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *iso20022.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation2 `xml:"CpyDtls,omitempty"`
}

func (t *TransferOutCancellationRequestV06) AddMessageIdentification() *iso20022.MessageIdentification1 {
	t.MessageIdentification = new(iso20022.MessageIdentification1)
	return t.MessageIdentification
}

func (t *TransferOutCancellationRequestV06) AddReferences() *iso20022.References15 {
	newValue := new(iso20022.References15)
	t.References = append(t.References, newValue)
	return newValue
}

func (t *TransferOutCancellationRequestV06) AddCancellation() *iso20022.Cancellation8Choice {
	newValue := new(iso20022.Cancellation8Choice)
	t.Cancellation = append(t.Cancellation, newValue)
	return newValue
}

func (t *TransferOutCancellationRequestV06) AddMarketPracticeVersion() *iso20022.MarketPracticeVersion1 {
	t.MarketPracticeVersion = new(iso20022.MarketPracticeVersion1)
	return t.MarketPracticeVersion
}

func (t *TransferOutCancellationRequestV06) AddCopyDetails() *iso20022.CopyInformation2 {
	t.CopyDetails = new(iso20022.CopyInformation2)
	return t.CopyDetails
}
