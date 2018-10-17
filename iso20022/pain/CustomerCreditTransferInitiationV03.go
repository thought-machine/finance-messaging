package pain

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document00100103 struct {
	XMLName xml.Name                             `xml:"urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 Document"`
	Message *CustomerCreditTransferInitiationV03 `xml:"CstmrCdtTrfInitn"`
}

func (d *Document00100103) AddMessage() *CustomerCreditTransferInitiationV03 {
	d.Message = new(CustomerCreditTransferInitiationV03)
	return d.Message
}

// Scope
// The CustomerCreditTransferInitiation message is sent by the initiating party to the forwarding agent or debtor agent. It is used to request movement of funds from the debtor account to a creditor.
// Usage
// The CustomerCreditTransferInitiation message can contain one or more customer credit transfer instructions.
// The CustomerCreditTransferInitiation message is used to exchange:
// - One or more instances of a credit transfer initiation;
// - Payment transactions that result in book transfers at the debtor agent or payments to another financial institution;
// - Payment transactions that result in an electronic cash transfer to the creditor account or in the emission of a cheque.
// The message can be used in a direct or a relay scenario:
// - In a direct scenario, the message is sent directly to the debtor agent. The debtor agent is the account servicer of the debtor.
// - In a relay scenario, the message is sent to a forwarding agent. The forwarding agent acts as a concentrating financial institution. It will forward the CustomerCreditTransferInitiation message to the debtor agent.
// The message can also be used by an initiating party that has authority to send the message on behalf of the debtor. This caters for example for the scenario of a payments factory initiating all payments on behalf of a large corporate.
// The CustomerCreditTransferInitiation message can be used in domestic and cross-border scenarios.
// The CustomerCreditTransferInitiation message must not be used by the debtor agent to execute the credit transfer instruction(s). The FIToFICustomerCreditTransfer message must be used instead.
type CustomerCreditTransferInitiationV03 struct {

	// Set of characteristics shared by all individual transactions included in the message.
	GroupHeader *iso20022.GroupHeader32 `xml:"GrpHdr"`

	// Set of characteristics that applies to the debit side of the payment transactions included in the credit transfer initiation.
	PaymentInformation []*iso20022.PaymentInstructionInformation3 `xml:"PmtInf"`
}

func (c *CustomerCreditTransferInitiationV03) AddGroupHeader() *iso20022.GroupHeader32 {
	c.GroupHeader = new(iso20022.GroupHeader32)
	return c.GroupHeader
}

func (c *CustomerCreditTransferInitiationV03) AddPaymentInformation() *iso20022.PaymentInstructionInformation3 {
	newValue := new(iso20022.PaymentInstructionInformation3)
	c.PaymentInformation = append(c.PaymentInformation, newValue)
	return newValue
}
