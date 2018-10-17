package caaa

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document00700103 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.007.001.03 Document"`
	Message *AcceptorCancellationAdviceV03 `xml:"AccptrCxlAdvc"`
}

func (d *Document00700103) AddMessage() *AcceptorCancellationAdviceV03 {
	d.Message = new(AcceptorCancellationAdviceV03)
	return d.Message
}

// The AcceptorCancellationAdvice message is sent by an acceptor (or its agent) to notify the acquirer (or its agent) of the cancellation of a successfully completed transaction. The transaction has been completed without financial transfer, or the acceptor is aware that the transaction was not cleared by the acquirer.
type AcceptorCancellationAdviceV03 struct {

	// Cancellation advice message management information.
	Header *iso20022.Header8 `xml:"Hdr"`

	// Information related to the cancellation advice.
	CancellationAdvice *iso20022.AcceptorCancellationAdvice3 `xml:"CxlAdvc"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType8 `xml:"SctyTrlr"`
}

func (a *AcceptorCancellationAdviceV03) AddHeader() *iso20022.Header8 {
	a.Header = new(iso20022.Header8)
	return a.Header
}

func (a *AcceptorCancellationAdviceV03) AddCancellationAdvice() *iso20022.AcceptorCancellationAdvice3 {
	a.CancellationAdvice = new(iso20022.AcceptorCancellationAdvice3)
	return a.CancellationAdvice
}

func (a *AcceptorCancellationAdviceV03) AddSecurityTrailer() *iso20022.ContentInformationType8 {
	a.SecurityTrailer = new(iso20022.ContentInformationType8)
	return a.SecurityTrailer
}
