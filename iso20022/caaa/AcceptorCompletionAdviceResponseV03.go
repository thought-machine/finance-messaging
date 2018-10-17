package caaa

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document00400103 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.004.001.03 Document"`
	Message *AcceptorCompletionAdviceResponseV03 `xml:"AccptrCmpltnAdvcRspn"`
}

func (d *Document00400103) AddMessage() *AcceptorCompletionAdviceResponseV03 {
	d.Message = new(AcceptorCompletionAdviceResponseV03)
	return d.Message
}

// The AcceptorCompletionAdviceResponse message is sent by the acquirer (or its agent) to acknowledge the acceptor (or its agent) of the outcome of the payment transaction, and the transfer the  financial data of the transaction contained in the completion advice.
type AcceptorCompletionAdviceResponseV03 struct {

	// Completion advice response message management information.
	Header *iso20022.Header8 `xml:"Hdr"`

	// Information related to the completion advice response.
	CompletionAdviceResponse *iso20022.AcceptorCompletionAdviceResponse3 `xml:"CmpltnAdvcRspn"`

	// Trailer of the message containing a MAC.
	SecurityTrailer *iso20022.ContentInformationType8 `xml:"SctyTrlr"`
}

func (a *AcceptorCompletionAdviceResponseV03) AddHeader() *iso20022.Header8 {
	a.Header = new(iso20022.Header8)
	return a.Header
}

func (a *AcceptorCompletionAdviceResponseV03) AddCompletionAdviceResponse() *iso20022.AcceptorCompletionAdviceResponse3 {
	a.CompletionAdviceResponse = new(iso20022.AcceptorCompletionAdviceResponse3)
	return a.CompletionAdviceResponse
}

func (a *AcceptorCompletionAdviceResponseV03) AddSecurityTrailer() *iso20022.ContentInformationType8 {
	a.SecurityTrailer = new(iso20022.ContentInformationType8)
	return a.SecurityTrailer
}
