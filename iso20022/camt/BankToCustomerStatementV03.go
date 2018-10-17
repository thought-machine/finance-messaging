package camt

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document05300103 struct {
	XMLName xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.053.001.03 Document"`
	Message *BankToCustomerStatementV03 `xml:"BkToCstmrStmt"`
}

func (d *Document05300103) AddMessage() *BankToCustomerStatementV03 {
	d.Message = new(BankToCustomerStatementV03)
	return d.Message
}

// Scope
// The BankToCustomerStatement message is sent by the account servicer to an account owner or to a party authorised by the account owner to receive the message. It is used to inform the account owner, or authorised party, of the entries booked to the account, and to provide the owner with balance information on the account at a given point in time.
// Usage
// The BankToCustomerStatement message can contain reports for more than one account. It provides information for cash management and/or reconciliation.
// It contains information on booked entries only.
// It can include underlying details of transactions that have been included in the entry.
// The message is exchanged as defined between the account servicer and the account owner. It provides information on items that have been booked to the account and also balance information. Depending on services and schedule agreed between banks and their customers, statements may be generated and exchanged accordingly, for example for intraday or prior day periods.
// It is possible that the receiver of the message is not the account owner, but a party entitled through arrangement with the account owner to receive the account information (also known as recipient).
type BankToCustomerStatementV03 struct {

	// Common information for the message.
	GroupHeader *iso20022.GroupHeader58 `xml:"GrpHdr"`

	// Reports on booked entries and balances for a cash account.
	Statement []*iso20022.AccountStatement3 `xml:"Stmt"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (b *BankToCustomerStatementV03) AddGroupHeader() *iso20022.GroupHeader58 {
	b.GroupHeader = new(iso20022.GroupHeader58)
	return b.GroupHeader
}

func (b *BankToCustomerStatementV03) AddStatement() *iso20022.AccountStatement3 {
	newValue := new(iso20022.AccountStatement3)
	b.Statement = append(b.Statement, newValue)
	return newValue
}

func (b *BankToCustomerStatementV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	b.SupplementaryData = append(b.SupplementaryData, newValue)
	return newValue
}
