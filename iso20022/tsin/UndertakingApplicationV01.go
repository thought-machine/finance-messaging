package tsin

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document00500101 struct {
	XMLName xml.Name                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.005.001.01 Document"`
	Message *UndertakingApplicationV01 `xml:"UdrtkgAppl"`
}

func (d *Document00500101) AddMessage() *UndertakingApplicationV01 {
	d.Message = new(UndertakingApplicationV01)
	return d.Message
}

// The UndertakingApplication message is sent by the party requesting issuance of the undertaking (applicant or obligor) to the party issuing the undertaking. It is used to request the issuance of an undertaking (demand guarantee or standby letter of credit or suretyship) or counter-undertaking (counter-guarantee or counter-standby or suretyship), and provides details on the applicable rules, terms, conditions and content of the undertaking to be issued.
type UndertakingApplicationV01 struct {

	// Details of the application for an independent undertaking, such as a demand guarantee or standby letter of credit, that provides financial assurance, to be collected on the presentation of documents that comply with its terms and conditions.
	UndertakingApplicationDetails *iso20022.Undertaking1 `xml:"UdrtkgApplDtls"`

	// Instructions specific to the bank receiving the message.
	InstructionsToBank []*iso20022.Max2000Text `xml:"InstrsToBk,omitempty"`

	// Digital signature of the undertaking application.
	DigitalSignature *iso20022.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`
}

func (u *UndertakingApplicationV01) AddUndertakingApplicationDetails() *iso20022.Undertaking1 {
	u.UndertakingApplicationDetails = new(iso20022.Undertaking1)
	return u.UndertakingApplicationDetails
}

func (u *UndertakingApplicationV01) AddInstructionsToBank(value string) {
	u.InstructionsToBank = append(u.InstructionsToBank, (*iso20022.Max2000Text)(&value))
}

func (u *UndertakingApplicationV01) AddDigitalSignature() *iso20022.PartyAndSignature2 {
	u.DigitalSignature = new(iso20022.PartyAndSignature2)
	return u.DigitalSignature
}
