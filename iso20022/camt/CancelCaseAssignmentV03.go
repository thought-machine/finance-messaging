package camt

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document03200103 struct {
	XMLName xml.Name                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.032.001.03 Document"`
	Message *CancelCaseAssignmentV03 `xml:"CclCaseAssgnmt"`
}

func (d *Document03200103) AddMessage() *CancelCaseAssignmentV03 {
	d.Message = new(CancelCaseAssignmentV03)
	return d.Message
}

// Scope
// The Cancel Case Assignment message is sent by a case creator or case assigner to a case assignee. This message is used to request the cancellation of a case.
// Usage
// The Cancel Case Assignment message is used to stop the processing of a case at a case assignee when a case assignment is incorrect or when the root cause for the case disappears (such as the account owner was able to reconcile after sending a Claim Non Receipt message).
// The Cancel Case Assignment message can be used to stop the processing of a:
// - request to cancel payment case
// - request to modify payment case
// - unable to apply case
// - claim non receipt case
// The Cancel Case Assignment message covers one and only one case at a time. If several case assignments need to be cancelled, then multiple Cancel Case Assignment messages must be sent.
// The Cancel Case Assignment message must be forwarded by all subsequent case assignee(s) in the case processing chain until it reaches the end point.
// When an agent re-assigns the Cancel Case Assignment to a subsequent case assignee, this agent must send a Notification Of Case Assignment message to its assigner.
// When the Cancel Case Assignment instruction has been acted upon by the relevant case assignee, a Resolution Of Investigation message is sent to the case assigner or case creator, in reply.
// The Cancel Case Assignment message must not be used for other purposes. If, for example, a request to modify payment fails, and the case creator requests the cancellation of the payment, then a Customer or FIToFI Payment Cancellation Request message must be used, with the case identification of the original Request To Modify Payment message. In this context it is incorrect to use the Cancel Case Assignment message.
type CancelCaseAssignmentV03 struct {

	// Identifies the assignment of an investigation case from an assigner to an assignee.
	// Usage: The Assigner must be the sender of this confirmation and the Assignee must be the receiver.
	Assignment *iso20022.CaseAssignment3 `xml:"Assgnmt"`

	// Identifies the investigation case.
	Case *iso20022.Case3 `xml:"Case"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CancelCaseAssignmentV03) AddAssignment() *iso20022.CaseAssignment3 {
	c.Assignment = new(iso20022.CaseAssignment3)
	return c.Assignment
}

func (c *CancelCaseAssignmentV03) AddCase() *iso20022.Case3 {
	c.Case = new(iso20022.Case3)
	return c.Case
}

func (c *CancelCaseAssignmentV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
