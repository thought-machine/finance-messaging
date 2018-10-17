package seev

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document00300103 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.003.001.03 Document"`
	Message *MeetingEntitlementNotificationV03 `xml:"MtgEntitlmntNtfctn"`
}

func (d *Document00300103) AddMessage() *MeetingEntitlementNotificationV03 {
	d.Message = new(MeetingEntitlementNotificationV03)
	return d.Message
}

// Scope
// An account servicer sends the MeetingEntitlementNotification to an issuer, its agent, an intermediary or an account owner to advise the entitlement in relation to a shareholders meeting.
// Usage
// This message is sent to advise the quantity of securities held by an account owner. The balance is specified for the securities for which the meeting is taking place.
// This entitlement message is sent by the account servicer or the registrar to an intermediary, the issuer's agent or the issuer. It is also sent between the account servicer and the account owner or the party holding the right to vote.
// The message is also used to amend a previously sent MeetingEntitlementNotification. To notify an update, the RelatedReference must be included in the message.
type MeetingEntitlementNotificationV03 struct {

	// Identifies the notification of entitlement instruction.
	Identification *iso20022.MessageIdentification1 `xml:"Id"`

	// Identifies the meeting entitlement message to be modified.
	RelatedReference *iso20022.MessageIdentification `xml:"RltdRef,omitempty"`

	// Series of elements which allow to identify a meeting.
	MeetingReference *iso20022.MeetingReference4 `xml:"MtgRef"`

	// Party notifying the entitlement.
	NotifyingParty *iso20022.PartyIdentification9Choice `xml:"NtifngPty"`

	// Identifies the security for which the meeting is organised, the account and the positions of the security holder.
	Security []*iso20022.SecurityPosition6 `xml:"Scty"`

	// Defines the dates determining eligibility.
	Eligibility *iso20022.EligibilityDates1 `xml:"Elgblty"`
}

func (m *MeetingEntitlementNotificationV03) AddIdentification() *iso20022.MessageIdentification1 {
	m.Identification = new(iso20022.MessageIdentification1)
	return m.Identification
}

func (m *MeetingEntitlementNotificationV03) AddRelatedReference() *iso20022.MessageIdentification {
	m.RelatedReference = new(iso20022.MessageIdentification)
	return m.RelatedReference
}

func (m *MeetingEntitlementNotificationV03) AddMeetingReference() *iso20022.MeetingReference4 {
	m.MeetingReference = new(iso20022.MeetingReference4)
	return m.MeetingReference
}

func (m *MeetingEntitlementNotificationV03) AddNotifyingParty() *iso20022.PartyIdentification9Choice {
	m.NotifyingParty = new(iso20022.PartyIdentification9Choice)
	return m.NotifyingParty
}

func (m *MeetingEntitlementNotificationV03) AddSecurity() *iso20022.SecurityPosition6 {
	newValue := new(iso20022.SecurityPosition6)
	m.Security = append(m.Security, newValue)
	return newValue
}

func (m *MeetingEntitlementNotificationV03) AddEligibility() *iso20022.EligibilityDates1 {
	m.Eligibility = new(iso20022.EligibilityDates1)
	return m.Eligibility
}
