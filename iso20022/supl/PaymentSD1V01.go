package supl

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document01700101 struct {
	XMLName xml.Name       `xml:"urn:iso:std:iso:20022:tech:xsd:supl.017.001.01 Document"`
	Message *PaymentSD1V01 `xml:"PmtSD1"`
}

func (d *Document01700101) AddMessage() *PaymentSD1V01 {
	d.Message = new(PaymentSD1V01)
	return d.Message
}

// Supplementary data for payment message definitions.
type PaymentSD1V01 struct {

	// Structured card remittance information supplied in a payment.
	CardRemittanceInformation *iso20022.TransactionData1 `xml:"CardRmtInf"`
}

func (p *PaymentSD1V01) AddCardRemittanceInformation() *iso20022.TransactionData1 {
	p.CardRemittanceInformation = new(iso20022.TransactionData1)
	return p.CardRemittanceInformation
}
