package sese

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document01000102 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.010.001.02 Document"`
	Message *TransferCancellationStatusReportV02 `xml:"TrfCxlStsRptV02"`
}

func (d *Document01000102) AddMessage() *TransferCancellationStatusReportV02 {
	d.Message = new(TransferCancellationStatusReportV02)
	return d.Message
}

// Scope
// An executing party, eg, a transfer agent, sends the TransferCancellationStatusReport message to the instructing party, eg, an investment manager or one of its authorised representatives to provide the status of a previously received transfer cancellation instruction.
// Usage
// The TransferCancellationStatusReport message is used to report on the status of a transfer in or transfer out cancellation request.
// The reference of the transfer instruction for which the cancellation status is reported is identified in TransferReference. The message identification of the transfer cancellation request message in which the transfer instruction was conveyed may also be quoted in RelatedReference.
// The message identification of the transfer instruction request message in which the transfer instruction was conveyed may also be quoted in RelatedReference.
// One of the following statuses can be reported:
// - the transfer cancellation is accepted, or,
// - the transfer cancellation has been sent to the next party, or,
// - the transfer cancellation is complete and the reason for the status,
// - the transfer cancellation pending and the reason for the status,
// - the transfer cancellation is rejected and the reason for the status.
type TransferCancellationStatusReportV02 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Reference to a linked message that was previously received.
	RelatedReference []*iso20022.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Reference to the linked message sent in a proprietary way or the reference of a system.
	OtherReference []*iso20022.AdditionalReference3 `xml:"OthrRef,omitempty"`

	// Status of the transfer cancellation instruction.
	StatusReport *iso20022.CancellationStatusAndReason2 `xml:"StsRpt"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferCancellationStatusReportV02) AddMessageIdentification() *iso20022.MessageIdentification1 {
	t.MessageIdentification = new(iso20022.MessageIdentification1)
	return t.MessageIdentification
}

func (t *TransferCancellationStatusReportV02) AddRelatedReference() *iso20022.AdditionalReference3 {
	newValue := new(iso20022.AdditionalReference3)
	t.RelatedReference = append(t.RelatedReference, newValue)
	return newValue
}

func (t *TransferCancellationStatusReportV02) AddOtherReference() *iso20022.AdditionalReference3 {
	newValue := new(iso20022.AdditionalReference3)
	t.OtherReference = append(t.OtherReference, newValue)
	return newValue
}

func (t *TransferCancellationStatusReportV02) AddStatusReport() *iso20022.CancellationStatusAndReason2 {
	t.StatusReport = new(iso20022.CancellationStatusAndReason2)
	return t.StatusReport
}

func (t *TransferCancellationStatusReportV02) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
