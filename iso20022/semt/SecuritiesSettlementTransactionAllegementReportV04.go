package semt

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document01900104 struct {
	XMLName xml.Name                                            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.019.001.04 Document"`
	Message *SecuritiesSettlementTransactionAllegementReportV04 `xml:"SctiesSttlmTxAllgmtRpt"`
}

func (d *Document01900104) AddMessage() *SecuritiesSettlementTransactionAllegementReportV04 {
	d.Message = new(SecuritiesSettlementTransactionAllegementReportV04)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementTransactionAllegementReport to an account owner to provide, at a specified time, the status and details of pending settlement allegements, for all or selected securities in a specified safekeeping account.
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
type SecuritiesSettlementTransactionAllegementReportV04 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *iso20022.Pagination `xml:"Pgntn"`

	// General information related to report.
	StatementGeneralDetails *iso20022.Statement17 `xml:"StmtGnlDtls"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Details of the allegement.
	AllegementDetails []*iso20022.SecuritiesTradeDetails35 `xml:"AllgmtDtls,omitempty"`
}

func (s *SecuritiesSettlementTransactionAllegementReportV04) AddPagination() *iso20022.Pagination {
	s.Pagination = new(iso20022.Pagination)
	return s.Pagination
}

func (s *SecuritiesSettlementTransactionAllegementReportV04) AddStatementGeneralDetails() *iso20022.Statement17 {
	s.StatementGeneralDetails = new(iso20022.Statement17)
	return s.StatementGeneralDetails
}

func (s *SecuritiesSettlementTransactionAllegementReportV04) AddAccountOwner() *iso20022.PartyIdentification36Choice {
	s.AccountOwner = new(iso20022.PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SecuritiesSettlementTransactionAllegementReportV04) AddSafekeepingAccount() *iso20022.SecuritiesAccount13 {
	s.SafekeepingAccount = new(iso20022.SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesSettlementTransactionAllegementReportV04) AddAllegementDetails() *iso20022.SecuritiesTradeDetails35 {
	newValue := new(iso20022.SecuritiesTradeDetails35)
	s.AllegementDetails = append(s.AllegementDetails, newValue)
	return newValue
}
