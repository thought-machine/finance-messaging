package sese

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document02000103 struct {
	XMLName xml.Name                                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.020.001.03 Document"`
	Message *SecuritiesTransactionCancellationRequestV03 `xml:"SctiesTxCxlReq"`
}

func (d *Document02000103) AddMessage() *SecuritiesTransactionCancellationRequestV03 {
	d.Message = new(SecuritiesTransactionCancellationRequestV03)
	return d.Message
}

// Scope
// This version is identical to the previous version. It was created automatically during the 2011/2012 maintenance cycle, at the same time as new versions of other Settlement and Reconciliation messages that were truly impacted by change requests. This should not have been the case. In future releases, SWIFT will ensure that a new version of a message is not created if identical to the previous version.
// An account owner sends a SecuritiesTransactionCancellationRequest to an account servicer to request the cancellation of a securities transaction.
// The account owner/servicer relationship may be:
// - a global custodian which has an account with a local custodian, or
// - an investment management institution which manage a fund account opened at a custodian, or - a broker which has an account with a custodian, or
// - a central securities depository participant which has an account with a central securities depository, or
// - a central securities depository which has an account with a custodian, another central securities depository or another settlement market infrastructure, or
// - a central counterparty or a stock exchange or a trade matching utility which need to instruct to a central securities depository ot another settlement market infrastructure.
//
// Usage
// The transaction may be:
// - a securities settlement transaction
// - an intra-position movement
// - a securities financing transaction
// The instruction cannot be:
// - a securities settlement conditions modification (another transaction processing command should be sent to reverse a processing change previously requested).
// - a securities financing modification
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
//
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesTransactionCancellationRequestV03 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *iso20022.References2Choice `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *iso20022.Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *iso20022.Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Message Reference identifying the Processor of the transaction.
	ProcessorTransactionIdentification *iso20022.Max35Text `xml:"PrcrTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Identifies the details of the transaction.
	TransactionDetails *iso20022.TransactionDetails29 `xml:"TxDtls,omitempty"`

	// Specifies whether an associated FX should be cancelled.
	FXCancellation *iso20022.FXCancellation1Choice `xml:"FxCxl,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesTransactionCancellationRequestV03) AddAccountOwnerTransactionIdentification() *iso20022.References2Choice {
	s.AccountOwnerTransactionIdentification = new(iso20022.References2Choice)
	return s.AccountOwnerTransactionIdentification
}

func (s *SecuritiesTransactionCancellationRequestV03) SetAccountServicerTransactionIdentification(value string) {
	s.AccountServicerTransactionIdentification = (*iso20022.Max35Text)(&value)
}

func (s *SecuritiesTransactionCancellationRequestV03) SetMarketInfrastructureTransactionIdentification(value string) {
	s.MarketInfrastructureTransactionIdentification = (*iso20022.Max35Text)(&value)
}

func (s *SecuritiesTransactionCancellationRequestV03) SetProcessorTransactionIdentification(value string) {
	s.ProcessorTransactionIdentification = (*iso20022.Max35Text)(&value)
}

func (s *SecuritiesTransactionCancellationRequestV03) AddAccountOwner() *iso20022.PartyIdentification36Choice {
	s.AccountOwner = new(iso20022.PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SecuritiesTransactionCancellationRequestV03) AddSafekeepingAccount() *iso20022.SecuritiesAccount13 {
	s.SafekeepingAccount = new(iso20022.SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesTransactionCancellationRequestV03) AddTransactionDetails() *iso20022.TransactionDetails29 {
	s.TransactionDetails = new(iso20022.TransactionDetails29)
	return s.TransactionDetails
}

func (s *SecuritiesTransactionCancellationRequestV03) AddFXCancellation() *iso20022.FXCancellation1Choice {
	s.FXCancellation = new(iso20022.FXCancellation1Choice)
	return s.FXCancellation
}

func (s *SecuritiesTransactionCancellationRequestV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
