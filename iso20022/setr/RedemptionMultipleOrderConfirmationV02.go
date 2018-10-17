package setr

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document00600102 struct {
	XMLName xml.Name                                `xml:"urn:iso:std:iso:20022:tech:xsd:setr.006.001.02 Document"`
	Message *RedemptionMultipleOrderConfirmationV02 `xml:"setr.006.001.02"`
}

func (d *Document00600102) AddMessage() *RedemptionMultipleOrderConfirmationV02 {
	d.Message = new(RedemptionMultipleOrderConfirmationV02)
	return d.Message
}

// Scope
// The RedemptionMultipleOrderConfirmation message is sent by an exacting party, eg, a transfer agent, to an instructing party, eg, an investment manager or its authorised representative. There may be one or more intermediary parties between the executing party and the instructing party. The intermediary party is, for example, an intermediary or a concentrator.
// This message is used to confirm the details of the execution of a RedemptionMultipleOrder message.
// Usage
// The RedemptionMultipleOrderConfirmation message is sent, after the price has been determined, to confirm the execution of all individual orders.
// A RedemptionMultipleOrder may be responded to by more than one RedemptionMultipleOrderConfirmation, as the valuation cycle of the financial instruments in each individual order may be different.
// When the executing party sends several confirmations, there is no specific indication in the message that it is an incomplete confirmation. Reconciliation must be based on the references.
// A RedemptionMultipleOrder must in all cases be responded to by a RedemptionMultipleOrderConfirmation message/s and in no circumstances by a RedemptionBulkOrderConfirmation message/s.
// If the executing party needs to confirm a RedemptionBulkOrder message, then a RedemptionBulkOrderConfirmation message must be used.
type RedemptionMultipleOrderConfirmationV02 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *iso20022.AdditionalReference3 `xml:"MstrRef,omitempty"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *iso20022.AdditionalReference3 `xml:"RltdRef"`

	// General information related to the execution of investment fund orders.
	MultipleExecutionDetails *iso20022.RedemptionMultipleExecution2 `xml:"MltplExctnDtls"`

	// Information related to an intermediary.
	IntermediaryDetails []*iso20022.Intermediary4 `xml:"IntrmyDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation1 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (r *RedemptionMultipleOrderConfirmationV02) AddMasterReference() *iso20022.AdditionalReference3 {
	r.MasterReference = new(iso20022.AdditionalReference3)
	return r.MasterReference
}

func (r *RedemptionMultipleOrderConfirmationV02) AddPoolReference() *iso20022.AdditionalReference3 {
	r.PoolReference = new(iso20022.AdditionalReference3)
	return r.PoolReference
}

func (r *RedemptionMultipleOrderConfirmationV02) AddPreviousReference() *iso20022.AdditionalReference3 {
	newValue := new(iso20022.AdditionalReference3)
	r.PreviousReference = append(r.PreviousReference, newValue)
	return newValue
}

func (r *RedemptionMultipleOrderConfirmationV02) AddRelatedReference() *iso20022.AdditionalReference3 {
	r.RelatedReference = new(iso20022.AdditionalReference3)
	return r.RelatedReference
}

func (r *RedemptionMultipleOrderConfirmationV02) AddMultipleExecutionDetails() *iso20022.RedemptionMultipleExecution2 {
	r.MultipleExecutionDetails = new(iso20022.RedemptionMultipleExecution2)
	return r.MultipleExecutionDetails
}

func (r *RedemptionMultipleOrderConfirmationV02) AddIntermediaryDetails() *iso20022.Intermediary4 {
	newValue := new(iso20022.Intermediary4)
	r.IntermediaryDetails = append(r.IntermediaryDetails, newValue)
	return newValue
}

func (r *RedemptionMultipleOrderConfirmationV02) AddCopyDetails() *iso20022.CopyInformation1 {
	r.CopyDetails = new(iso20022.CopyInformation1)
	return r.CopyDetails
}

func (r *RedemptionMultipleOrderConfirmationV02) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	r.Extension = append(r.Extension, newValue)
	return newValue
}
