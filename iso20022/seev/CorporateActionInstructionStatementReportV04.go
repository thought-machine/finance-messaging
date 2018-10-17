package seev

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document04200104 struct {
	XMLName xml.Name                                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.042.001.04 Document"`
	Message *CorporateActionInstructionStatementReportV04 `xml:"CorpActnInstrStmtRpt"`
}

func (d *Document04200104) AddMessage() *CorporateActionInstructionStatementReportV04 {
	d.Message = new(CorporateActionInstructionStatementReportV04)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionInstructionStatementReport message to an account owner or its designated agent to report balances at the safekeeping account level for one or more corporate action events or at the corporate action event level for one or more safekeeping accounts.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
// ISO 15022 - 20022 COEXISTENCE
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type CorporateActionInstructionStatementReportV04 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *iso20022.Pagination `xml:"Pgntn"`

	// General characteristics related to a statement which reports information.
	StatementGeneralDetails *iso20022.Statement12 `xml:"StmtGnlDtls"`

	// Account information and detailed account holdings information report for corporate action events.
	AccountAndStatementDetails []*iso20022.AccountIdentification25 `xml:"AcctAndStmtDtls"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionInstructionStatementReportV04) AddPagination() *iso20022.Pagination {
	c.Pagination = new(iso20022.Pagination)
	return c.Pagination
}

func (c *CorporateActionInstructionStatementReportV04) AddStatementGeneralDetails() *iso20022.Statement12 {
	c.StatementGeneralDetails = new(iso20022.Statement12)
	return c.StatementGeneralDetails
}

func (c *CorporateActionInstructionStatementReportV04) AddAccountAndStatementDetails() *iso20022.AccountIdentification25 {
	newValue := new(iso20022.AccountIdentification25)
	c.AccountAndStatementDetails = append(c.AccountAndStatementDetails, newValue)
	return newValue
}

func (c *CorporateActionInstructionStatementReportV04) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
