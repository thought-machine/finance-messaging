package tsrv

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document01100101 struct {
	XMLName xml.Name                                `xml:"urn:iso:std:iso:20022:tech:xsd:tsrv.011.001.01 Document"`
	Message *UndertakingNonExtensionNotificationV01 `xml:"UdrtkgNonXtnsnNtfctn"`
}

func (d *Document01100101) AddMessage() *UndertakingNonExtensionNotificationV01 {
	d.Message = new(UndertakingNonExtensionNotificationV01)
	return d.Message
}

// The UndertakingNonExtensionNotification message is sent by the party that issued the undertaking to the beneficiary, either directly or via one or more advising parties, to notify it of the non-extension of the referenced undertaking beyond the current expiry date.
type UndertakingNonExtensionNotificationV01 struct {

	// Details of the non-extension notification.
	UndertakingNonExtensionNotificationDetails *iso20022.UndertakingNonExtensionStatusAdvice1 `xml:"UdrtkgNonXtnsnNtfctnDtls"`

	// Digital signature of the notification.
	DigitalSignature *iso20022.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`
}

func (u *UndertakingNonExtensionNotificationV01) AddUndertakingNonExtensionNotificationDetails() *iso20022.UndertakingNonExtensionStatusAdvice1 {
	u.UndertakingNonExtensionNotificationDetails = new(iso20022.UndertakingNonExtensionStatusAdvice1)
	return u.UndertakingNonExtensionNotificationDetails
}

func (u *UndertakingNonExtensionNotificationV01) AddDigitalSignature() *iso20022.PartyAndSignature2 {
	u.DigitalSignature = new(iso20022.PartyAndSignature2)
	return u.DigitalSignature
}
