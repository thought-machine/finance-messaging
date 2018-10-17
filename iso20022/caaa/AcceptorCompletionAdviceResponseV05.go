package caaa

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document00400105 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.004.001.05 Document"`
	Message *AcceptorCompletionAdviceResponseV05 `xml:"AccptrCmpltnAdvcRspn"`
}

func (d *Document00400105) AddMessage() *AcceptorCompletionAdviceResponseV05 {
	d.Message = new(AcceptorCompletionAdviceResponseV05)
	return d.Message
}

// The AcceptorCompletionAdviceResponse message is sent by the acquirer (or its agent) to acknowledge the acceptor (or its agent) of the outcome of the payment transaction, and the transfer the  financial data of the transaction contained in the completion advice.
type AcceptorCompletionAdviceResponseV05 struct {

	// Completion advice response message management information.
	Header *iso20022.Header24 `xml:"Hdr"`

	// Information related to the completion advice response.
	CompletionAdviceResponse *iso20022.AcceptorCompletionAdviceResponse5 `xml:"CmpltnAdvcRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType15 `xml:"SctyTrlr,omitempty"`
}

func (a *AcceptorCompletionAdviceResponseV05) AddHeader() *iso20022.Header24 {
	a.Header = new(iso20022.Header24)
	return a.Header
}

func (a *AcceptorCompletionAdviceResponseV05) AddCompletionAdviceResponse() *iso20022.AcceptorCompletionAdviceResponse5 {
	a.CompletionAdviceResponse = new(iso20022.AcceptorCompletionAdviceResponse5)
	return a.CompletionAdviceResponse
}

func (a *AcceptorCompletionAdviceResponseV05) AddSecurityTrailer() *iso20022.ContentInformationType15 {
	a.SecurityTrailer = new(iso20022.ContentInformationType15)
	return a.SecurityTrailer
}
