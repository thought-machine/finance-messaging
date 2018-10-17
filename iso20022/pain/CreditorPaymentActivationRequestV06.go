package pain

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document01300106 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.013.001.06 Document"`
	Message *CreditorPaymentActivationRequestV06 `xml:"CdtrPmtActvtnReq"`
}

func (d *Document01300106) AddMessage() *CreditorPaymentActivationRequestV06 {
	d.Message = new(CreditorPaymentActivationRequestV06)
	return d.Message
}

// The CreditorPaymentActivationRequest message is sent by the Creditor sending party to the Debtor receiving party, directly or through agents. It is used by a Creditor to request movement of funds from the debtor account to a creditor.
type CreditorPaymentActivationRequestV06 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader45 `xml:"GrpHdr"`

	// Set of characteristics that applies to the debit side of the payment transactions included in the creditor payment initiation.
	PaymentInformation []*iso20022.PaymentInstruction23 `xml:"PmtInf"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CreditorPaymentActivationRequestV06) AddGroupHeader() *iso20022.GroupHeader45 {
	c.GroupHeader = new(iso20022.GroupHeader45)
	return c.GroupHeader
}

func (c *CreditorPaymentActivationRequestV06) AddPaymentInformation() *iso20022.PaymentInstruction23 {
	newValue := new(iso20022.PaymentInstruction23)
	c.PaymentInformation = append(c.PaymentInformation, newValue)
	return newValue
}

func (c *CreditorPaymentActivationRequestV06) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
