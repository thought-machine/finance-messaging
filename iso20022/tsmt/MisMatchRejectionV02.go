package tsmt

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document02200102 struct {
	XMLName xml.Name              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.022.001.02 Document"`
	Message *MisMatchRejectionV02 `xml:"MisMtchRjctn"`
}

func (d *Document02200102) AddMessage() *MisMatchRejectionV02 {
	d.Message = new(MisMatchRejectionV02)
	return d.Message
}

// Scope
// The MisMatchRejection message is sent by the party requested to accept or reject data set mis-matches to the matching application.
// This message is used to reject mis-matches between data sets and the related baseline.
// Usage
// The MisMatchRejection message can be sent by the party requested to accept or reject data set mis-match to inform that it rejects the data set(s).
// The message can be sent in response to a DataSetMatchReport message conveying mis-matches.
// The information about the rejection of the mis-matched data sets will be forwarded by the matching application to the submitter of the data sets by a MisMatchRejectionNotification message.
// The acceptance of mis-matched data sets can be achieved by sending a MisMatchAcceptance message.
type MisMatchRejectionV02 struct {

	// Identifies the rejection message.
	RejectionIdentification *iso20022.MessageIdentification1 `xml:"RjctnId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *iso20022.SimpleIdentificationInformation `xml:"TxId"`

	// Reference to the transaction for the requesting financial institution.
	SubmitterTransactionReference *iso20022.SimpleIdentificationInformation `xml:"SubmitrTxRef,omitempty"`

	// Reference to the identification of the report that contained the difference.
	DataSetMatchReportReference *iso20022.MessageIdentification1 `xml:"DataSetMtchRptRef"`

	// Reason why the user cannot accept the request.
	RejectionReason *iso20022.RejectionReason1Choice `xml:"RjctnRsn"`
}

func (m *MisMatchRejectionV02) AddRejectionIdentification() *iso20022.MessageIdentification1 {
	m.RejectionIdentification = new(iso20022.MessageIdentification1)
	return m.RejectionIdentification
}

func (m *MisMatchRejectionV02) AddTransactionIdentification() *iso20022.SimpleIdentificationInformation {
	m.TransactionIdentification = new(iso20022.SimpleIdentificationInformation)
	return m.TransactionIdentification
}

func (m *MisMatchRejectionV02) AddSubmitterTransactionReference() *iso20022.SimpleIdentificationInformation {
	m.SubmitterTransactionReference = new(iso20022.SimpleIdentificationInformation)
	return m.SubmitterTransactionReference
}

func (m *MisMatchRejectionV02) AddDataSetMatchReportReference() *iso20022.MessageIdentification1 {
	m.DataSetMatchReportReference = new(iso20022.MessageIdentification1)
	return m.DataSetMatchReportReference
}

func (m *MisMatchRejectionV02) AddRejectionReason() *iso20022.RejectionReason1Choice {
	m.RejectionReason = new(iso20022.RejectionReason1Choice)
	return m.RejectionReason
}
