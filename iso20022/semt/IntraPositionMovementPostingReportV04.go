package semt

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document01600104 struct {
	XMLName xml.Name                               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.016.001.04 Document"`
	Message *IntraPositionMovementPostingReportV04 `xml:"IntraPosMvmntPstngRpt"`
}

func (d *Document01600104) AddMessage() *IntraPositionMovementPostingReportV04 {
	d.Message = new(IntraPositionMovementPostingReportV04)
	return d.Message
}

// Scope
// An account servicer sends an IntraPositionMovementPostingReport to an account owner to provide the details of increases and decreases in securities with a given status within a holding, that is, intra-position transfers, which occurred during a specified period, for all or selected securities in a specified safekeeping account which the account servicer holds for the account owner.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information.
// using the relevant elements in the Business Application Header.
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type IntraPositionMovementPostingReportV04 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *iso20022.Pagination `xml:"Pgntn"`

	// General information related to report.
	StatementGeneralDetails *iso20022.Statement15 `xml:"StmtGnlDtls"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Reporting per financial instrument.
	FinancialInstrument []*iso20022.FinancialInstrumentDetails14 `xml:"FinInstrm,omitempty"`
}

func (i *IntraPositionMovementPostingReportV04) AddPagination() *iso20022.Pagination {
	i.Pagination = new(iso20022.Pagination)
	return i.Pagination
}

func (i *IntraPositionMovementPostingReportV04) AddStatementGeneralDetails() *iso20022.Statement15 {
	i.StatementGeneralDetails = new(iso20022.Statement15)
	return i.StatementGeneralDetails
}

func (i *IntraPositionMovementPostingReportV04) AddAccountOwner() *iso20022.PartyIdentification36Choice {
	i.AccountOwner = new(iso20022.PartyIdentification36Choice)
	return i.AccountOwner
}

func (i *IntraPositionMovementPostingReportV04) AddSafekeepingAccount() *iso20022.SecuritiesAccount13 {
	i.SafekeepingAccount = new(iso20022.SecuritiesAccount13)
	return i.SafekeepingAccount
}

func (i *IntraPositionMovementPostingReportV04) AddFinancialInstrument() *iso20022.FinancialInstrumentDetails14 {
	newValue := new(iso20022.FinancialInstrumentDetails14)
	i.FinancialInstrument = append(i.FinancialInstrument, newValue)
	return newValue
}
