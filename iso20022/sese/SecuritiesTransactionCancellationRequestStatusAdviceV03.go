package sese

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document02700103 struct {
	XMLName xml.Name                                                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.027.001.03 Document"`
	Message *SecuritiesTransactionCancellationRequestStatusAdviceV03 `xml:"SctiesTxCxlReqStsAdvc"`
}

func (d *Document02700103) AddMessage() *SecuritiesTransactionCancellationRequestStatusAdviceV03 {
	d.Message = new(SecuritiesTransactionCancellationRequestStatusAdviceV03)
	return d.Message
}

// Scope
// An account servicer sends an SecuritiesTransactionCancellationRequestStatusAdvice to an account owner to advise the status of a securities transaction cancellation request previously sent by the account owner.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
//
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
//
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesTransactionCancellationRequestStatusAdviceV03 struct {

	// Reference to the unambiguous identification of the cancellation request as per the account owner.
	CancellationRequestReference *iso20022.Identification1 `xml:"CxlReqRef"`

	// Unambiguous identification of the transaction as known by the account servicer.
	TransactionIdentification *iso20022.TransactionIdentifications17 `xml:"TxId,omitempty"`

	// Provides details on the processing status of the request.
	ProcessingStatus *iso20022.ProcessingStatus20Choice `xml:"PrcgSts"`

	// Identifies the details of the transaction.
	TransactionDetails *iso20022.TransactionDetails30 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesTransactionCancellationRequestStatusAdviceV03) AddCancellationRequestReference() *iso20022.Identification1 {
	s.CancellationRequestReference = new(iso20022.Identification1)
	return s.CancellationRequestReference
}

func (s *SecuritiesTransactionCancellationRequestStatusAdviceV03) AddTransactionIdentification() *iso20022.TransactionIdentifications17 {
	s.TransactionIdentification = new(iso20022.TransactionIdentifications17)
	return s.TransactionIdentification
}

func (s *SecuritiesTransactionCancellationRequestStatusAdviceV03) AddProcessingStatus() *iso20022.ProcessingStatus20Choice {
	s.ProcessingStatus = new(iso20022.ProcessingStatus20Choice)
	return s.ProcessingStatus
}

func (s *SecuritiesTransactionCancellationRequestStatusAdviceV03) AddTransactionDetails() *iso20022.TransactionDetails30 {
	s.TransactionDetails = new(iso20022.TransactionDetails30)
	return s.TransactionDetails
}

func (s *SecuritiesTransactionCancellationRequestStatusAdviceV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
