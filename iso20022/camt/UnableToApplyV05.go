package camt

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document02600105 struct {
	XMLName xml.Name          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.05 Document"`
	Message *UnableToApplyV05 `xml:"UblToApply"`
}

func (d *Document02600105) AddMessage() *UnableToApplyV05 {
	d.Message = new(UnableToApplyV05)
	return d.Message
}

// Scope
// The Unable To Apply message is sent by a case creator or a case assigner to a case assignee. This message is used to initiate an investigation of a payment instruction that cannot be executed or reconciled.
// Usage
// The Unable To Apply case occurs in two situations:
// - an agent cannot execute the payment instruction due to missing or incorrect information
// - a creditor cannot reconcile the payment entry as it is received unexpectedly, or missing or incorrect information prevents reconciliation
// The Unable To Apply message
// - always follows the reverse route of the original payment instruction
// - must be forwarded to the preceding agents in the payment processing chain, where appropriate
// - covers one and only one payment instruction (or payment entry) at a time; if several payment instructions cannot be executed or several payment entries cannot be reconciled, then multiple Unable To Apply messages must be sent.
// Depending on what stage the payment is and what has been done to it, for example incorrect routing, errors/omissions when processing the instruction or even the absence of any error, the unable to apply case may lead to a:
// - Additional Payment Information message, sent to the case creator/case assigner, if a truncation or omission in a payment instruction prevented reconciliation.
// - Request To Cancel Payment message, sent to the subsequent agent in the payment processing chain, if the original payment instruction has been incorrectly routed through the chain of agents (this also entails a new corrected payment instruction being issued). Prior to sending the payment cancellation request, the agent should first send a resolution indicating that a cancellation will follow (CWFW).
// - Request To Modify Payment message, sent to the subsequent agent in the payment processing chain, if a truncation or omission has occurred during the processing of the original payment instruction. Prior to sending the modify payment request, the agent should first send a resolution indicating that a modification will follow (MWFW).
// - Debit Authorisation Request message, sent to the case creator/case assigner, if the payment instruction has reached an incorrect creditor.
// The UnableToApply message has the following main characteristics:
// The case creator (the instructed party/creditor of a payment instruction) assigns a unique case identification and optionally the reason code for the Unable To Apply message. This information will be passed unchanged to all subsequent case assignees.
// The case creator specifies the identification of the underlying payment instruction. This identification needs to be updated by the subsequent case assigner(s) in order to match the one used with their case assignee(s).
// The Unable To Apply Justification element allows the assigner to indicate whether a specific element causes the unable to apply situation (incorrect element and/or mismatched element can be listed) or whether any supplementary information needs to be forwarded.
type UnableToApplyV05 struct {

	// Identifies the assignment of an investigation case from an assigner to an assignee.
	// Usage: The Assigner must be the sender of this confirmation and the Assignee must be the receiver.
	Assignment *iso20022.CaseAssignment3 `xml:"Assgnmt"`

	// Identifies the investigation case.
	Case *iso20022.Case3 `xml:"Case"`

	// References the payment instruction or statement entry that a party is unable to execute or unable to reconcile.
	Underlying *iso20022.UnderlyingTransaction3Choice `xml:"Undrlyg"`

	// Explains the reason why the case creator is unable to apply the instruction.
	Justification *iso20022.UnableToApplyJustification3Choice `xml:"Justfn"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (u *UnableToApplyV05) AddAssignment() *iso20022.CaseAssignment3 {
	u.Assignment = new(iso20022.CaseAssignment3)
	return u.Assignment
}

func (u *UnableToApplyV05) AddCase() *iso20022.Case3 {
	u.Case = new(iso20022.Case3)
	return u.Case
}

func (u *UnableToApplyV05) AddUnderlying() *iso20022.UnderlyingTransaction3Choice {
	u.Underlying = new(iso20022.UnderlyingTransaction3Choice)
	return u.Underlying
}

func (u *UnableToApplyV05) AddJustification() *iso20022.UnableToApplyJustification3Choice {
	u.Justification = new(iso20022.UnableToApplyJustification3Choice)
	return u.Justification
}

func (u *UnableToApplyV05) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	u.SupplementaryData = append(u.SupplementaryData, newValue)
	return newValue
}
