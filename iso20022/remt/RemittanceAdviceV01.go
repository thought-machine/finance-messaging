package remt

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document00100101 struct {
	XMLName xml.Name             `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.01 Document"`
	Message *RemittanceAdviceV01 `xml:"RmtAdvc"`
}

func (d *Document00100101) AddMessage() *RemittanceAdviceV01 {
	d.Message = new(RemittanceAdviceV01)
	return d.Message
}

// The RemittanceAdvice message allows the originator to provide remittance details that can be associated with a payment.
type RemittanceAdviceV01 struct {

	// Set of characteristics shared by all remittance information included in the message.
	GroupHeader *iso20022.GroupHeader62 `xml:"GrpHdr"`

	// Provides information to enable the matching of an entry with the items that the associated payment is intended to settle, such as commercial invoices in an accounts' receivable system, tax obligations, or garnishment orders.
	RemittanceInformation []*iso20022.RemittanceInformation8 `xml:"RmtInf"`

	// Additional information that cannot be  captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (r *RemittanceAdviceV01) AddGroupHeader() *iso20022.GroupHeader62 {
	r.GroupHeader = new(iso20022.GroupHeader62)
	return r.GroupHeader
}

func (r *RemittanceAdviceV01) AddRemittanceInformation() *iso20022.RemittanceInformation8 {
	newValue := new(iso20022.RemittanceInformation8)
	r.RemittanceInformation = append(r.RemittanceInformation, newValue)
	return newValue
}

func (r *RemittanceAdviceV01) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	r.SupplementaryData = append(r.SupplementaryData, newValue)
	return newValue
}
