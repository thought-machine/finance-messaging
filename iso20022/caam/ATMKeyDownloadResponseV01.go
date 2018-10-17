package caam

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document00400101 struct {
	XMLName xml.Name                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.004.001.01 Document"`
	Message *ATMKeyDownloadResponseV01 `xml:"ATMKeyDwnldRspn"`
}

func (d *Document00400101) AddMessage() *ATMKeyDownloadResponseV01 {
	d.Message = new(ATMKeyDownloadResponseV01)
	return d.Message
}

// The ATMKeyDownloadResponse message is sent from an acquirer to an ATM in response to an ATMKeyDownloadRequest message, to download of one or several cryptographic keys.
type ATMKeyDownloadResponseV01 struct {

	// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
	Header *iso20022.Header20 `xml:"Hdr"`

	// Encrypted body of the message.
	ProtectedATMKeyDownloadResponse *iso20022.ContentInformationType10 `xml:"PrtctdATMKeyDwnldRspn,omitempty"`

	// Information related to the response of an ATM key download from an ATM manager.
	ATMKeyDownloadResponse *iso20022.ATMKeyDownloadResponse1 `xml:"ATMKeyDwnldRspn,omitempty"`

	// Trailer of the message containing a MAC or a digital signature.
	SecurityTrailer *iso20022.ContentInformationType13 `xml:"SctyTrlr,omitempty"`
}

func (a *ATMKeyDownloadResponseV01) AddHeader() *iso20022.Header20 {
	a.Header = new(iso20022.Header20)
	return a.Header
}

func (a *ATMKeyDownloadResponseV01) AddProtectedATMKeyDownloadResponse() *iso20022.ContentInformationType10 {
	a.ProtectedATMKeyDownloadResponse = new(iso20022.ContentInformationType10)
	return a.ProtectedATMKeyDownloadResponse
}

func (a *ATMKeyDownloadResponseV01) AddATMKeyDownloadResponse() *iso20022.ATMKeyDownloadResponse1 {
	a.ATMKeyDownloadResponse = new(iso20022.ATMKeyDownloadResponse1)
	return a.ATMKeyDownloadResponse
}

func (a *ATMKeyDownloadResponseV01) AddSecurityTrailer() *iso20022.ContentInformationType13 {
	a.SecurityTrailer = new(iso20022.ContentInformationType13)
	return a.SecurityTrailer
}
