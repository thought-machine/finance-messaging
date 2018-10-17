package sese

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document03800103 struct {
	XMLName xml.Name                                               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.038.001.03 Document"`
	Message *SecuritiesSettlementTransactionModificationRequestV03 `xml:"SctiesSttlmTxModReq"`
}

func (d *Document03800103) AddMessage() *SecuritiesSettlementTransactionModificationRequestV03 {
	d.Message = new(SecuritiesSettlementTransactionModificationRequestV03)
	return d.Message
}

// Scope
// This message is sent by an account owner to an account servicer.
//
// The account owner will generally be:
// - a central securities depository participant which has an account with a central securities depository or a market infrastructure
// - an investment manager which has an account with a custodian acting as accounting and/or settlement agent.
//
// It is used to request the modification of non core business data (matching or non-matching) information in a pending or settled instruction. It can also be used for the enrichment of an incomplete transaction.
//
// Usage
// The modification must only contain the data to be modified.
// The message may also be used to:
// - re-send a message sent by the account owner to the account servicer,
// - provide a third party with a copy of a message being sent by the account owner for information,
// - re-send to a third party a copy of a message being sent by the account owner for information
// using the relevant elements in the Business Application Header.
//
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesSettlementTransactionModificationRequestV03 struct {

	// Identifies the details of the transaction that is being modified.
	ModifiedTransactionDetails *iso20022.TransactionDetails64 `xml:"ModfdTxDtls"`

	// Specifies the type of update requested.
	UpdateType []*iso20022.UpdateType11Choice `xml:"UpdTp"`
}

func (s *SecuritiesSettlementTransactionModificationRequestV03) AddModifiedTransactionDetails() *iso20022.TransactionDetails64 {
	s.ModifiedTransactionDetails = new(iso20022.TransactionDetails64)
	return s.ModifiedTransactionDetails
}

func (s *SecuritiesSettlementTransactionModificationRequestV03) AddUpdateType() *iso20022.UpdateType11Choice {
	newValue := new(iso20022.UpdateType11Choice)
	s.UpdateType = append(s.UpdateType, newValue)
	return newValue
}
