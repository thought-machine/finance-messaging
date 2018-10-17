package camt

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document03300104 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.033.001.04 Document"`
	Message *RequestForDuplicateV04 `xml:"ReqForDplct"`
}

func (d *Document03300104) AddMessage() *RequestForDuplicateV04 {
	d.Message = new(RequestForDuplicateV04)
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
type RequestForDuplicateV04 struct {

	// Identifies the assignment of an investigation case from an assigner to an assignee.
	// Usage: The Assigner must be the sender of this confirmation and the Assignee must be the receiver.
	Assignment *iso20022.CaseAssignment3 `xml:"Assgnmt"`

	// Identifies the investigation case.
	Case *iso20022.Case3 `xml:"Case"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (r *RequestForDuplicateV04) AddAssignment() *iso20022.CaseAssignment3 {
	r.Assignment = new(iso20022.CaseAssignment3)
	return r.Assignment
}

func (r *RequestForDuplicateV04) AddCase() *iso20022.Case3 {
	r.Case = new(iso20022.Case3)
	return r.Case
}

func (r *RequestForDuplicateV04) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	r.SupplementaryData = append(r.SupplementaryData, newValue)
	return newValue
}
