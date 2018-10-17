package sese

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document02900103 struct {
	XMLName xml.Name                                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.029.001.03 Document"`
	Message *SecuritiesSettlementAllegementRemovalAdviceV03 `xml:"SctiesSttlmAllgmtRmvlAdvc"`
}

func (d *Document02900103) AddMessage() *SecuritiesSettlementAllegementRemovalAdviceV03 {
	d.Message = new(SecuritiesSettlementAllegementRemovalAdviceV03)
	return d.Message
}

// This version is identical to the previous version. It was created automatically during the 2011/2012 maintenance cycle, at the same time as new versions of other Settlement and Reconciliation messages that were truly impacted by change requests. This should not have been the case. In future releases, SWIFT will ensure that a new version of a message is not created if identical to the previous version.
// Scope
// An account servicer sends a SecuritiesSettlementAllegementRemovalAdvice to an account owner to acknowledge that a previously sent allegement is no longer outstanding, because the alleged party sent its instruction.
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
type SecuritiesSettlementAllegementRemovalAdviceV03 struct {

	// Provides transaction type and identification information.
	AccountServicerTransactionIdentification *iso20022.SettlementTypeAndIdentification3 `xml:"AcctSvcrTxId"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *iso20022.Identification1 `xml:"MktInfrstrctrTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Identifies the details of the transaction.
	TransactionDetails *iso20022.TransactionDetails29 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementAllegementRemovalAdviceV03) AddAccountServicerTransactionIdentification() *iso20022.SettlementTypeAndIdentification3 {
	s.AccountServicerTransactionIdentification = new(iso20022.SettlementTypeAndIdentification3)
	return s.AccountServicerTransactionIdentification
}

func (s *SecuritiesSettlementAllegementRemovalAdviceV03) AddMarketInfrastructureTransactionIdentification() *iso20022.Identification1 {
	s.MarketInfrastructureTransactionIdentification = new(iso20022.Identification1)
	return s.MarketInfrastructureTransactionIdentification
}

func (s *SecuritiesSettlementAllegementRemovalAdviceV03) AddAccountOwner() *iso20022.PartyIdentification36Choice {
	s.AccountOwner = new(iso20022.PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SecuritiesSettlementAllegementRemovalAdviceV03) AddSafekeepingAccount() *iso20022.SecuritiesAccount13 {
	s.SafekeepingAccount = new(iso20022.SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesSettlementAllegementRemovalAdviceV03) AddTransactionDetails() *iso20022.TransactionDetails29 {
	s.TransactionDetails = new(iso20022.TransactionDetails29)
	return s.TransactionDetails
}

func (s *SecuritiesSettlementAllegementRemovalAdviceV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
