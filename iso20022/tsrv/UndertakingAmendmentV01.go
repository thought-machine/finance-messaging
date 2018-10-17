package tsrv

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document00500101 struct {
	XMLName xml.Name                 `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.005.001.01 Document"`
	Message *UndertakingAmendmentV01 `xml:"UdrtkgAmdmnt"`
}

func (d *Document00500101) AddMessage() *UndertakingAmendmentV01 {
	d.Message = new(UndertakingAmendmentV01)
	return d.Message
}

// The UndertakingAmendment message is sent (and is thus issued) by the party that issued the undertaking. The message may be sent either directly to the beneficiary or via an advising party. The proposed undertaking amendment could be to a demand guarantee, standby letter of credit, or counter-undertaking (counter-guarantee or counter-standby). The message provides details on proposed changes to the undertaking, for example, to the expiry date, the amount, and terms and conditions of the undertaking. It may also be used to propose the termination or cancellation of the undertaking. Under practice and law, this communication binds the party issuing it. The message constitutes an operative financial instrument.
type UndertakingAmendmentV01 struct {

	// Details related to the proposed undertaking amendment.
	UndertakingAmendmentDetails *iso20022.Amendment1 `xml:"UdrtkgAmdmntDtls"`

	// Additional information specific to the bank-to-bank communication.
	BankToBankInformation []*iso20022.Max2000Text `xml:"BkToBkInf,omitempty"`

	// Digital signature of the proposed undertaking amendment.
	DigitalSignature []*iso20022.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`
}

func (u *UndertakingAmendmentV01) AddUndertakingAmendmentDetails() *iso20022.Amendment1 {
	u.UndertakingAmendmentDetails = new(iso20022.Amendment1)
	return u.UndertakingAmendmentDetails
}

func (u *UndertakingAmendmentV01) AddBankToBankInformation(value string) {
	u.BankToBankInformation = append(u.BankToBankInformation, (*iso20022.Max2000Text)(&value))
}

func (u *UndertakingAmendmentV01) AddDigitalSignature() *iso20022.PartyAndSignature2 {
	newValue := new(iso20022.PartyAndSignature2)
	u.DigitalSignature = append(u.DigitalSignature, newValue)
	return newValue
}
