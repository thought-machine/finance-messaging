package semt

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document01400205 struct {
	XMLName xml.Name                                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.002.05 Document"`
	Message *IntraPositionMovementStatusAdvice002V05 `xml:"IntraPosMvmntStsAdvc"`
}

func (d *Document01400205) AddMessage() *IntraPositionMovementStatusAdvice002V05 {
	d.Message = new(IntraPositionMovementStatusAdvice002V05)
	return d.Message
}

// Scope
// An account servicer sends a IntraPositionMovementStatusAdvice to an account owner to advise the status of a intra-position movement instruction previously sent by the account owner.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
//
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information.
// using the relevant elements in the Business Application Header.
type IntraPositionMovementStatusAdvice002V05 struct {

	// Unambiguous identification of a transaction as per the account owner (or the instructing party managing the account).
	TransactionIdentification *iso20022.TransactionIdentifications34 `xml:"TxId"`

	// Provides details on the processing status of the transaction.
	ProcessingStatus *iso20022.IntraPositionProcessingStatus6Choice `xml:"PrcgSts,omitempty"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *iso20022.SettlementStatus20Choice `xml:"SttlmSts,omitempty"`

	// Identifies the details of the transaction.
	TransactionDetails *iso20022.IntraPositionDetails42 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *IntraPositionMovementStatusAdvice002V05) AddTransactionIdentification() *iso20022.TransactionIdentifications34 {
	i.TransactionIdentification = new(iso20022.TransactionIdentifications34)
	return i.TransactionIdentification
}

func (i *IntraPositionMovementStatusAdvice002V05) AddProcessingStatus() *iso20022.IntraPositionProcessingStatus6Choice {
	i.ProcessingStatus = new(iso20022.IntraPositionProcessingStatus6Choice)
	return i.ProcessingStatus
}

func (i *IntraPositionMovementStatusAdvice002V05) AddSettlementStatus() *iso20022.SettlementStatus20Choice {
	i.SettlementStatus = new(iso20022.SettlementStatus20Choice)
	return i.SettlementStatus
}

func (i *IntraPositionMovementStatusAdvice002V05) AddTransactionDetails() *iso20022.IntraPositionDetails42 {
	i.TransactionDetails = new(iso20022.IntraPositionDetails42)
	return i.TransactionDetails
}

func (i *IntraPositionMovementStatusAdvice002V05) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
