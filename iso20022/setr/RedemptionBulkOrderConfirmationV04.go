package setr

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document00300104 struct {
	XMLName xml.Name                            `xml:"urn:iso:std:iso:20022:tech:xsd:setr.003.001.04 Document"`
	Message *RedemptionBulkOrderConfirmationV04 `xml:"RedBlkOrdrConf"`
}

func (d *Document00300104) AddMessage() *RedemptionBulkOrderConfirmationV04 {
	d.Message = new(RedemptionBulkOrderConfirmationV04)
	return d.Message
}

// Scope
// The RedemptionBulkOrderConfirmation message is sent by an executing party, for example, a transfer agent, to the instructing party, for example, an investment manager or its authorised representative, to confirm the details of execution for a previously received RedemptionBulkOrder message.
// Usage
// The RedemptionBulkOrderConfirmation message is used to confirm the execution of all individual orders included in a previously sent RedemptionBulkOrder message.
// There is usually one bulk confirmation message for one bulk order message.
// Each individual order confirmation specified is identified in DealReference. The reference of the original individual order is specified in OrderReference. The message identification of the RedemptionBulkOrder message in which the individual order was conveyed may also be quoted in RelatedReference, but this is not recommended.
// A RedemptionBulkOrder must in all cases be responded to by a RedemptionBulkOrderConfirmation and in no circumstances by a RedemptionOrderConfirmation.
// If the executing party needs to confirm a RedemptionOrder instruction, then the RedemptionOrderConfirmation must be used.
// When the message is used to convey a confirmation amendment/s, the AmendmentIndicator must be present with the value ‘true’ or ‘1’. When this is the case, the message must only contain a confirmation amendment/s and not contain both a confirmation amendment/s and a ‘new’ confirmation/s.
type RedemptionBulkOrderConfirmationV04 struct {

	// Reference that uniquely identifies the message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference9 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*iso20022.AdditionalReference8 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference8 `xml:"RltdRef,omitempty"`

	// General information related to the execution of the orders.
	BulkExecutionDetails *iso20022.RedemptionBulkExecution5 `xml:"BlkExctnDtls"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation4 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (r *RedemptionBulkOrderConfirmationV04) AddMessageIdentification() *iso20022.MessageIdentification1 {
	r.MessageIdentification = new(iso20022.MessageIdentification1)
	return r.MessageIdentification
}

func (r *RedemptionBulkOrderConfirmationV04) AddPoolReference() *iso20022.AdditionalReference9 {
	r.PoolReference = new(iso20022.AdditionalReference9)
	return r.PoolReference
}

func (r *RedemptionBulkOrderConfirmationV04) AddPreviousReference() *iso20022.AdditionalReference8 {
	newValue := new(iso20022.AdditionalReference8)
	r.PreviousReference = append(r.PreviousReference, newValue)
	return newValue
}

func (r *RedemptionBulkOrderConfirmationV04) AddRelatedReference() *iso20022.AdditionalReference8 {
	r.RelatedReference = new(iso20022.AdditionalReference8)
	return r.RelatedReference
}

func (r *RedemptionBulkOrderConfirmationV04) AddBulkExecutionDetails() *iso20022.RedemptionBulkExecution5 {
	r.BulkExecutionDetails = new(iso20022.RedemptionBulkExecution5)
	return r.BulkExecutionDetails
}

func (r *RedemptionBulkOrderConfirmationV04) AddCopyDetails() *iso20022.CopyInformation4 {
	r.CopyDetails = new(iso20022.CopyInformation4)
	return r.CopyDetails
}

func (r *RedemptionBulkOrderConfirmationV04) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	r.Extension = append(r.Extension, newValue)
	return newValue
}
