package setr

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document00200102 struct {
	XMLName xml.Name                                       `xml:"urn:iso:std:iso:20022:tech:xsd:setr.002.001.02 Document"`
	Message *RedemptionBulkOrderCancellationInstructionV02 `xml:"setr.002.001.02"`
}

func (d *Document00200102) AddMessage() *RedemptionBulkOrderCancellationInstructionV02 {
	d.Message = new(RedemptionBulkOrderCancellationInstructionV02)
	return d.Message
}

// Scope
// The RedemptionBulkOrderCancellationInstruction message is sent by an instructing party, eg, an investment manager or its authorised representative, to an executing party, eg, a transfer agent. There may be one or more intermediary parties between the instructing party and the executing party. The intermediary party is, for example, an intermediary or a concentrator.
// This message is sent to cancel a previously sent RedemptionBulkOrder instruction.
// Usage
// The RedemptionBulkOrderCancellationInstruction message is used to cancel the entire previously sent RedemptionBulkOrder message and all the individual orders that it contained. There is no amendment, but a cancellation and re-instruct policy.
// This message must contain the reference of the message to be cancelled. This message may also contain all the details of the message to be cancelled, but this is not recommended.
// The deadline and acceptance of a cancellation instruction is subject to a service level agreement (SLA). This cancellation message is a cancellation instruction. There is no automatic acceptance of the cancellation instruction.
// The rejection or acceptance of a cancellation message instruction is made using an OrderCancellationStatusReport message.
type RedemptionBulkOrderCancellationInstructionV02 struct {

	// Reference assigned to a set of orders or trades in order to link them together.
	MasterReference *iso20022.AdditionalReference3 `xml:"MstrRef,omitempty"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *iso20022.AdditionalReference3 `xml:"PrvsRef"`

	// Common information related to all the orders to be cancelled.
	OrderToBeCancelled *iso20022.RedemptionBulkOrderInstruction1 `xml:"OrdrToBeCanc,omitempty"`
}

func (r *RedemptionBulkOrderCancellationInstructionV02) AddMasterReference() *iso20022.AdditionalReference3 {
	r.MasterReference = new(iso20022.AdditionalReference3)
	return r.MasterReference
}

func (r *RedemptionBulkOrderCancellationInstructionV02) AddPoolReference() *iso20022.AdditionalReference3 {
	r.PoolReference = new(iso20022.AdditionalReference3)
	return r.PoolReference
}

func (r *RedemptionBulkOrderCancellationInstructionV02) AddPreviousReference() *iso20022.AdditionalReference3 {
	r.PreviousReference = new(iso20022.AdditionalReference3)
	return r.PreviousReference
}

func (r *RedemptionBulkOrderCancellationInstructionV02) AddOrderToBeCancelled() *iso20022.RedemptionBulkOrderInstruction1 {
	r.OrderToBeCancelled = new(iso20022.RedemptionBulkOrderInstruction1)
	return r.OrderToBeCancelled
}
