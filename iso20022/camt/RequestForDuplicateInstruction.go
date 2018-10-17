package camt

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document03300101 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.033.001.01 Document"`
	Message *RequestForDuplicateInstruction `xml:"camt.033.001.01"`
}

func (d *Document03300101) AddMessage() *RequestForDuplicateInstruction {
	d.Message = new(RequestForDuplicateInstruction)
	return d.Message
}

// Scope
// The Request For Duplicate message is sent by the case assignee to the case creator or case assigner.
// This message is used to request a copy of the original payment instruction considered in the case.
// Usage
// The Request For Duplicate message:
// - must be answered with a Duplicate message
// - must be used when a case assignee requests a copy of the original payment instruction. This occurs, for example, when the case assignee cannot trace the payment instruction based on the elements mentioned in the case assignment message
// - covers one and only one instruction at a time. If several payment instruction copies are needed by the case assignee, then multiple Request For Duplicate messages must be sent
// - must be used exclusively between the case assignee and its case creator/case assigner
type RequestForDuplicateInstruction struct {

	//
	Assignment *iso20022.CaseAssignment `xml:"Assgnmt"`

	//
	Case *iso20022.Case `xml:"Case"`
}

func (r *RequestForDuplicateInstruction) AddAssignment() *iso20022.CaseAssignment {
	r.Assignment = new(iso20022.CaseAssignment)
	return r.Assignment
}

func (r *RequestForDuplicateInstruction) AddCase() *iso20022.Case {
	r.Case = new(iso20022.Case)
	return r.Case
}
