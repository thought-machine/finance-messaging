package semt

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document01400102 struct {
	XMLName xml.Name                              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.02 Document"`
	Message *IntraPositionMovementStatusAdviceV02 `xml:"IntraPosMvmntStsAdvc"`
}

func (d *Document01400102) AddMessage() *IntraPositionMovementStatusAdviceV02 {
	d.Message = new(IntraPositionMovementStatusAdviceV02)
	return d.Message
}

// Scope
// An account servicer sends a IntraPositionMovementStatusAdvice to an account owner to advise the status of a intra-position movement instruction previously sent by the account owner.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information.
// using the relevant elements in the Business Application Header.
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type IntraPositionMovementStatusAdviceV02 struct {

	// Unambiguous identification of a transaction as per the account owner (or the instructing party managing the account).
	TransactionIdentification *iso20022.TransactionIdentifications10 `xml:"TxId"`

	// Provides details on the processing status of the transaction.
	ProcessingStatus *iso20022.IntraPositionProcessingStatus1Choice `xml:"PrcgSts,omitempty"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *iso20022.SettlementStatus2Choice `xml:"SttlmSts,omitempty"`

	// Identifies the details of the transaction.
	TransactionDetails *iso20022.IntraPositionDetails9 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *IntraPositionMovementStatusAdviceV02) AddTransactionIdentification() *iso20022.TransactionIdentifications10 {
	i.TransactionIdentification = new(iso20022.TransactionIdentifications10)
	return i.TransactionIdentification
}

func (i *IntraPositionMovementStatusAdviceV02) AddProcessingStatus() *iso20022.IntraPositionProcessingStatus1Choice {
	i.ProcessingStatus = new(iso20022.IntraPositionProcessingStatus1Choice)
	return i.ProcessingStatus
}

func (i *IntraPositionMovementStatusAdviceV02) AddSettlementStatus() *iso20022.SettlementStatus2Choice {
	i.SettlementStatus = new(iso20022.SettlementStatus2Choice)
	return i.SettlementStatus
}

func (i *IntraPositionMovementStatusAdviceV02) AddTransactionDetails() *iso20022.IntraPositionDetails9 {
	i.TransactionDetails = new(iso20022.IntraPositionDetails9)
	return i.TransactionDetails
}

func (i *IntraPositionMovementStatusAdviceV02) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
