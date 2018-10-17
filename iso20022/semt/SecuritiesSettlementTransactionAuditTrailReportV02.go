package semt

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document02200102 struct {
	XMLName xml.Name                                            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.022.001.02 Document"`
	Message *SecuritiesSettlementTransactionAuditTrailReportV02 `xml:"SctiesSttlmTxAudtTrlRpt"`
}

func (d *Document02200102) AddMessage() *SecuritiesSettlementTransactionAuditTrailReportV02 {
	d.Message = new(SecuritiesSettlementTransactionAuditTrailReportV02)
	return d.Message
}

// Scope
//
// This message is sent by the Market Infrastructure to the CSD to advise of the history of all the statuses, modifications, replacement and cancellation of a specific transaction during its whole life cycle when the instructing party is a direct participant to the Settlement Infrastructure.
//
//
// Usage
//
// The message may also be used to:
//
// - re-send a message sent by the market infrastructure to the direct participant,
//
// - provide a third party with a copy of a message being sent by the market infrastructure for information,
//
// - re-send to a third party a copy of a message being sent by the market infrastructure for information
// using the relevant elements in the Business Application Header.
//
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesSettlementTransactionAuditTrailReportV02 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *iso20022.Pagination `xml:"Pgntn"`

	// Identification of the SecuritiesStatusQuery message sent to request this report.
	QueryReference *iso20022.Identification1 `xml:"QryRef,omitempty"`

	// Provides unambiguous transaction identification information.
	TransactionIdentification *iso20022.TransactionIdentifications15 `xml:"TxId,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	//  Provides the history of status and reasons for a pending, posted or cancelled transaction.
	StatusTrail []*iso20022.StatusTrail4 `xml:"StsTrl,omitempty"`
}

func (s *SecuritiesSettlementTransactionAuditTrailReportV02) AddPagination() *iso20022.Pagination {
	s.Pagination = new(iso20022.Pagination)
	return s.Pagination
}

func (s *SecuritiesSettlementTransactionAuditTrailReportV02) AddQueryReference() *iso20022.Identification1 {
	s.QueryReference = new(iso20022.Identification1)
	return s.QueryReference
}

func (s *SecuritiesSettlementTransactionAuditTrailReportV02) AddTransactionIdentification() *iso20022.TransactionIdentifications15 {
	s.TransactionIdentification = new(iso20022.TransactionIdentifications15)
	return s.TransactionIdentification
}

func (s *SecuritiesSettlementTransactionAuditTrailReportV02) AddSafekeepingAccount() *iso20022.SecuritiesAccount13 {
	s.SafekeepingAccount = new(iso20022.SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesSettlementTransactionAuditTrailReportV02) AddAccountOwner() *iso20022.PartyIdentification36Choice {
	s.AccountOwner = new(iso20022.PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SecuritiesSettlementTransactionAuditTrailReportV02) AddStatusTrail() *iso20022.StatusTrail4 {
	newValue := new(iso20022.StatusTrail4)
	s.StatusTrail = append(s.StatusTrail, newValue)
	return newValue
}
