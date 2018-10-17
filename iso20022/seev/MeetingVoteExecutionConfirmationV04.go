package seev

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document00700104 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.007.001.04 Document"`
	Message *MeetingVoteExecutionConfirmationV04 `xml:"MtgVoteExctnConf"`
}

func (d *Document00700104) AddMessage() *MeetingVoteExecutionConfirmationV04 {
	d.Message = new(MeetingVoteExecutionConfirmationV04)
	return d.Message
}

// Scope
// An issuer, its agent or an intermediary sends the MeetingVoteExecutionConfirmation message to confirm to the Sender of the MeetingInstruction message, the execution of their voting instruction.
// Usage
// This message is sent after the shareholders meeting has taken place. The Sender of this message confirms the execution of the vote at the meeting and confirms that the vote has been processed as instructed via the MeetingInstruction message.
// This messages is sent if the Sender of the MeetingInstruction message has requested such a confirmation or if market practice or regulation stipulates the need for a full audit trail.
type MeetingVoteExecutionConfirmationV04 struct {

	// Identifies the vote execution confirmation message.
	Identification *iso20022.MessageIdentification1 `xml:"Id"`

	// Identifies the meeting instruction message.
	RelatedReference *iso20022.MessageIdentification `xml:"RltdRef"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *iso20022.MeetingReference4 `xml:"MtgRef"`

	// Party confirming the votes.
	ReportingParty *iso20022.PartyIdentification9Choice `xml:"RptgPty"`

	// Identifies the securities for which the meeting is organised.
	SecurityIdentification *iso20022.SecurityIdentification11 `xml:"SctyId"`

	// Specifies how a party has voted for each agenda item.
	VoteInstructions []*iso20022.DetailedInstructionStatus9 `xml:"VoteInstrs"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*iso20022.Extension2 `xml:"Xtnsn,omitempty"`
}

func (m *MeetingVoteExecutionConfirmationV04) AddIdentification() *iso20022.MessageIdentification1 {
	m.Identification = new(iso20022.MessageIdentification1)
	return m.Identification
}

func (m *MeetingVoteExecutionConfirmationV04) AddRelatedReference() *iso20022.MessageIdentification {
	m.RelatedReference = new(iso20022.MessageIdentification)
	return m.RelatedReference
}

func (m *MeetingVoteExecutionConfirmationV04) AddMeetingReference() *iso20022.MeetingReference4 {
	m.MeetingReference = new(iso20022.MeetingReference4)
	return m.MeetingReference
}

func (m *MeetingVoteExecutionConfirmationV04) AddReportingParty() *iso20022.PartyIdentification9Choice {
	m.ReportingParty = new(iso20022.PartyIdentification9Choice)
	return m.ReportingParty
}

func (m *MeetingVoteExecutionConfirmationV04) AddSecurityIdentification() *iso20022.SecurityIdentification11 {
	m.SecurityIdentification = new(iso20022.SecurityIdentification11)
	return m.SecurityIdentification
}

func (m *MeetingVoteExecutionConfirmationV04) AddVoteInstructions() *iso20022.DetailedInstructionStatus9 {
	newValue := new(iso20022.DetailedInstructionStatus9)
	m.VoteInstructions = append(m.VoteInstructions, newValue)
	return newValue
}

func (m *MeetingVoteExecutionConfirmationV04) AddExtension() *iso20022.Extension2 {
	newValue := new(iso20022.Extension2)
	m.Extension = append(m.Extension, newValue)
	return newValue
}
