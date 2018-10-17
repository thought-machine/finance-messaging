package caaa

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document00100102 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.001.001.02 Document"`
	Message *AcceptorAuthorisationRequestV02 `xml:"AccptrAuthstnReq"`
}

func (d *Document00100102) AddMessage() *AcceptorAuthorisationRequestV02 {
	d.Message = new(AcceptorAuthorisationRequestV02)
	return d.Message
}

// The AcceptorAuthorisationRequest message is sent by an acceptor (or its agent) to the acquirer (or its agent) , to check with the issuer (or its agent) that the account associated to the card has the resources to fund the payment. This checking will include validation of the card data and any additional transaction data provided.
type AcceptorAuthorisationRequestV02 struct {

	// Authorisation request message management information.
	Header *iso20022.Header1 `xml:"Hdr"`

	// Information related to the authorisation request.
	AuthorisationRequest *iso20022.AcceptorAuthorisationRequest2 `xml:"AuthstnReq"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType6 `xml:"SctyTrlr"`
}

func (a *AcceptorAuthorisationRequestV02) AddHeader() *iso20022.Header1 {
	a.Header = new(iso20022.Header1)
	return a.Header
}

func (a *AcceptorAuthorisationRequestV02) AddAuthorisationRequest() *iso20022.AcceptorAuthorisationRequest2 {
	a.AuthorisationRequest = new(iso20022.AcceptorAuthorisationRequest2)
	return a.AuthorisationRequest
}

func (a *AcceptorAuthorisationRequestV02) AddSecurityTrailer() *iso20022.ContentInformationType6 {
	a.SecurityTrailer = new(iso20022.ContentInformationType6)
	return a.SecurityTrailer
}
