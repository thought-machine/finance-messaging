package auth

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document02600101 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.026.001.01 Document"`
	Message *CurrencyControlRequestOrLetterV01 `xml:"CcyCtrlReqOrLttr"`
}

func (d *Document02600101) AddMessage() *CurrencyControlRequestOrLetterV01 {
	d.Message = new(CurrencyControlRequestOrLetterV01)
	return d.Message
}

// The CurrencyControlRequestOrLetter message is sent by the reporting party (respectively the registration agent) to the registration agent (respectively the reporting party) to send a currency control related letter or to request for supporting documents.
type CurrencyControlRequestOrLetterV01 struct {

	// Characteristics shared by all individual items included in the message.
	GroupHeader *iso20022.CurrencyControlHeader3 `xml:"GrpHdr"`

	// Supporting document request or letter details.
	RequestOrLetter []*iso20022.SupportingDocumentRequestOrLetter1 `xml:"ReqOrLttr"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CurrencyControlRequestOrLetterV01) AddGroupHeader() *iso20022.CurrencyControlHeader3 {
	c.GroupHeader = new(iso20022.CurrencyControlHeader3)
	return c.GroupHeader
}

func (c *CurrencyControlRequestOrLetterV01) AddRequestOrLetter() *iso20022.SupportingDocumentRequestOrLetter1 {
	newValue := new(iso20022.SupportingDocumentRequestOrLetter1)
	c.RequestOrLetter = append(c.RequestOrLetter, newValue)
	return newValue
}

func (c *CurrencyControlRequestOrLetterV01) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
