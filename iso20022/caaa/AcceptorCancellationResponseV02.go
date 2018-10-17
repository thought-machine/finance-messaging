package caaa

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document00600102 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.006.001.02 Document"`
	Message *AcceptorCancellationResponseV02 `xml:"AccptrCxlRspn"`
}

func (d *Document00600102) AddMessage() *AcceptorCancellationResponseV02 {
	d.Message = new(AcceptorCancellationResponseV02)
	return d.Message
}

// The AcceptorCancellationResponse message is sent by the acquirer (or its agent) to an acceptor (or its agent), to return the outcome of the cancellation request. If the response is positive, the acquirer has voided the financial data from the captured transaction.
type AcceptorCancellationResponseV02 struct {

	// Cancellation response message management information.
	Header *iso20022.Header1 `xml:"Hdr"`

	// Information related to the cancellation response.
	CancellationResponse *iso20022.AcceptorCancellationResponse2 `xml:"CxlRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType6 `xml:"SctyTrlr"`
}

func (a *AcceptorCancellationResponseV02) AddHeader() *iso20022.Header1 {
	a.Header = new(iso20022.Header1)
	return a.Header
}

func (a *AcceptorCancellationResponseV02) AddCancellationResponse() *iso20022.AcceptorCancellationResponse2 {
	a.CancellationResponse = new(iso20022.AcceptorCancellationResponse2)
	return a.CancellationResponse
}

func (a *AcceptorCancellationResponseV02) AddSecurityTrailer() *iso20022.ContentInformationType6 {
	a.SecurityTrailer = new(iso20022.ContentInformationType6)
	return a.SecurityTrailer
}
