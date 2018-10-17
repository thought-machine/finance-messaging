package tsrv

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document01300101 struct {
	XMLName xml.Name              `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.013.001.01 Document"`
	Message *UndertakingDemandV01 `xml:"UdrtkgDmnd"`
}

func (d *Document01300101) AddMessage() *UndertakingDemandV01 {
	d.Message = new(UndertakingDemandV01)
	return d.Message
}

// The UndertakingDemand message and other required documents are sent by the beneficiary to the party that issued the undertaking, either directly or via a presenting or nominated party. It is a demand for payment and may include a request to extend the undertaking expiry date. The demand itself must be contained in an enclosed file within the message or must be specified as narrative text within the message. It may contain other required documents in addition to the demand.
type UndertakingDemandV01 struct {

	// Details of the demand.
	UndertakingDemandDetails *iso20022.Demand1 `xml:"UdrtkgDmndDtls"`

	// Additional information specific to the bank-to-bank communication.
	BankToBankInformation []*iso20022.Max2000Text `xml:"BkToBkInf,omitempty"`

	// Digital signature of the demand.
	DigitalSignature *iso20022.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`
}

func (u *UndertakingDemandV01) AddUndertakingDemandDetails() *iso20022.Demand1 {
	u.UndertakingDemandDetails = new(iso20022.Demand1)
	return u.UndertakingDemandDetails
}

func (u *UndertakingDemandV01) AddBankToBankInformation(value string) {
	u.BankToBankInformation = append(u.BankToBankInformation, (*iso20022.Max2000Text)(&value))
}

func (u *UndertakingDemandV01) AddDigitalSignature() *iso20022.PartyAndSignature2 {
	u.DigitalSignature = new(iso20022.PartyAndSignature2)
	return u.DigitalSignature
}
