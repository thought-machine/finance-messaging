package tsmt

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document05300101 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.053.001.01 Document"`
	Message *InvoicePaymentReconciliationAdviceV01 `xml:"InvcPmtRcncltnAdvc"`
}

func (d *Document05300101) AddMessage() *InvoicePaymentReconciliationAdviceV01 {
	d.Message = new(InvoicePaymentReconciliationAdviceV01)
	return d.Message
}

// The message InvoicePaymentReconciliationAdvice is sent by a payer to a payee to indicate attribution of payments to instalment of payment obligations in order to simplify the account netting or clearing when a lot of invoices are paid with a unique payment (for instance an SCT or an SDD).
// The message contains references to payment instructions, may reference other messages and may include referenced data.
// The message can carry digital signatures if required by context.
type InvoicePaymentReconciliationAdviceV01 struct {

	// Set of characteristics that unambiguously identify the letter, common parameters, documents and identifications.
	Header *iso20022.BusinessLetter1 `xml:"Hdr"`

	// List of payment reconciliation information.
	ReconciliationList []*iso20022.ReconciliationList1 `xml:"RcncltnList"`

	// Number of reconciliation lists as control value.
	ReconciliationCount *iso20022.Max15NumericText `xml:"RcncltnCnt,omitempty"`

	// Total number of individual items in all lists.
	ItemCount *iso20022.Max15NumericText `xml:"ItmCnt,omitempty"`

	// Total of all individual amounts included in all lists, irrespective of currencies or direction.
	ControlSum *iso20022.DecimalNumber `xml:"CtrlSum,omitempty"`

	// Referenced or related business message.
	AttachedMessage []*iso20022.EncapsulatedBusinessMessage1 `xml:"AttchdMsg,omitempty"`
}

func (i *InvoicePaymentReconciliationAdviceV01) AddHeader() *iso20022.BusinessLetter1 {
	i.Header = new(iso20022.BusinessLetter1)
	return i.Header
}

func (i *InvoicePaymentReconciliationAdviceV01) AddReconciliationList() *iso20022.ReconciliationList1 {
	newValue := new(iso20022.ReconciliationList1)
	i.ReconciliationList = append(i.ReconciliationList, newValue)
	return newValue
}

func (i *InvoicePaymentReconciliationAdviceV01) SetReconciliationCount(value string) {
	i.ReconciliationCount = (*iso20022.Max15NumericText)(&value)
}

func (i *InvoicePaymentReconciliationAdviceV01) SetItemCount(value string) {
	i.ItemCount = (*iso20022.Max15NumericText)(&value)
}

func (i *InvoicePaymentReconciliationAdviceV01) SetControlSum(value string) {
	i.ControlSum = (*iso20022.DecimalNumber)(&value)
}

func (i *InvoicePaymentReconciliationAdviceV01) AddAttachedMessage() *iso20022.EncapsulatedBusinessMessage1 {
	newValue := new(iso20022.EncapsulatedBusinessMessage1)
	i.AttachedMessage = append(i.AttachedMessage, newValue)
	return newValue
}
