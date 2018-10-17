package semt

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document01800102 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.018.001.02 Document"`
	Message *SecuritiesTransactionPendingReportV02 `xml:"SctiesTxPdgRpt"`
}

func (d *Document01800102) AddMessage() *SecuritiesTransactionPendingReportV02 {
	d.Message = new(SecuritiesTransactionPendingReportV02)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesTransactionPendingReport to an account owner to provide, as at a specified time, the details of pending increases and decreases of holdings, for all or selected securities in a specified safekeeping account, for all or selected reasons why the transaction is pending.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
// Usage
// The statement may also include future settlement or forward transactions which have become binding on the account owner.
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// using the relevant elements in the Business Application Header.
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesTransactionPendingReportV02 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *iso20022.Pagination `xml:"Pgntn"`

	// Provides general information on the report.
	StatementGeneralDetails *iso20022.Statement14 `xml:"StmtGnlDtls"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Status information.
	Status []*iso20022.StatusAndReason7 `xml:"Sts,omitempty"`

	// Details of the transactions reported.
	Transactions []*iso20022.Transaction12 `xml:"Txs,omitempty"`
}

func (s *SecuritiesTransactionPendingReportV02) AddPagination() *iso20022.Pagination {
	s.Pagination = new(iso20022.Pagination)
	return s.Pagination
}

func (s *SecuritiesTransactionPendingReportV02) AddStatementGeneralDetails() *iso20022.Statement14 {
	s.StatementGeneralDetails = new(iso20022.Statement14)
	return s.StatementGeneralDetails
}

func (s *SecuritiesTransactionPendingReportV02) AddAccountOwner() *iso20022.PartyIdentification36Choice {
	s.AccountOwner = new(iso20022.PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SecuritiesTransactionPendingReportV02) AddSafekeepingAccount() *iso20022.SecuritiesAccount13 {
	s.SafekeepingAccount = new(iso20022.SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesTransactionPendingReportV02) AddStatus() *iso20022.StatusAndReason7 {
	newValue := new(iso20022.StatusAndReason7)
	s.Status = append(s.Status, newValue)
	return newValue
}

func (s *SecuritiesTransactionPendingReportV02) AddTransactions() *iso20022.Transaction12 {
	newValue := new(iso20022.Transaction12)
	s.Transactions = append(s.Transactions, newValue)
	return newValue
}
