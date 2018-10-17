package camt

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document03600101 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.036.001.01 Document"`
	Message *DebitAuthorisationResponse `xml:"camt.036.001.01"`
}

func (d *Document03600101) AddMessage() *DebitAuthorisationResponse {
	d.Message = new(DebitAuthorisationResponse)
	return d.Message
}

// Scope
// The Debit Authorisation Response message is sent by an account owner to its account servicing institution. This message is used to approve or reject a debit authorisation request.
// Usage
// The Debit Authorisation Response message is used to reply to a Debit Authorisation Request message.
// The Debit Authorisation Response message covers one and only one payment instruction at a time. If an account owner needs to reply to several Debit Authorisation Request messages, then multiple Debit Authorisation Response messages must be sent.
// The Debit Authorisation Response message indicates whether the account owner agrees with the request by means of a code. It also allows further details to be given about the debit authorisation, such as acceptable amount and value date for the debit.
// The Debit Authorisation Response message must be used exclusively between the account owner and the account servicing institution. It must not be used in place of a Resolution Of Investigation message between subsequent agents.
type DebitAuthorisationResponse struct {

	// Identifies an assignment.
	Assignment *iso20022.CaseAssignment `xml:"Assgnmt"`

	// Identifies a case.
	Case *iso20022.Case `xml:"Case"`

	// Indicates if the debit authorisation is granted or not.
	Confirmation *iso20022.DebitAuthorisationConfirmation `xml:"Conf"`
}

func (d *DebitAuthorisationResponse) AddAssignment() *iso20022.CaseAssignment {
	d.Assignment = new(iso20022.CaseAssignment)
	return d.Assignment
}

func (d *DebitAuthorisationResponse) AddCase() *iso20022.Case {
	d.Case = new(iso20022.Case)
	return d.Case
}

func (d *DebitAuthorisationResponse) AddConfirmation() *iso20022.DebitAuthorisationConfirmation {
	d.Confirmation = new(iso20022.DebitAuthorisationConfirmation)
	return d.Confirmation
}
