package camt

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document03400103 struct {
	XMLName xml.Name      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.034.001.03 Document"`
	Message *DuplicateV03 `xml:"Dplct"`
}

func (d *Document03400103) AddMessage() *DuplicateV03 {
	d.Message = new(DuplicateV03)
	return d.Message
}

// Scope
// The Duplicate message is used by financial institutions, with their own offices, and/or with other financial institutions with which they have established bilateral agreements. It allows to exchange duplicate payment instructions.
// Usage
// This message must be sent in response to a Request For Duplicate message.
// The Duplicate Data element must contain a well formed XML document. This means XML special characters such as '<' must be used in a way that is consistent with XML well-formedness criteria.
type DuplicateV03 struct {

	// Identifies the assignment of an investigation case from an assigner to an assignee.
	// Usage: The Assigner must be the sender of this confirmation and the Assignee must be the receiver.
	Assignment *iso20022.CaseAssignment2 `xml:"Assgnmt"`

	// Identifies the investigation case.
	Case *iso20022.Case2 `xml:"Case"`

	// Duplicate of a previously sent message.
	Duplicate *iso20022.ProprietaryData4 `xml:"Dplct"`
}

func (d *DuplicateV03) AddAssignment() *iso20022.CaseAssignment2 {
	d.Assignment = new(iso20022.CaseAssignment2)
	return d.Assignment
}

func (d *DuplicateV03) AddCase() *iso20022.Case2 {
	d.Case = new(iso20022.Case2)
	return d.Case
}

func (d *DuplicateV03) AddDuplicate() *iso20022.ProprietaryData4 {
	d.Duplicate = new(iso20022.ProprietaryData4)
	return d.Duplicate
}
