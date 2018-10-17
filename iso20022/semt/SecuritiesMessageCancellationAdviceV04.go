package semt

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document02000104 struct {
	XMLName xml.Name                                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.04 Document"`
	Message *SecuritiesMessageCancellationAdviceV04 `xml:"SctiesMsgCxlAdvc"`
}

func (d *Document02000104) AddMessage() *SecuritiesMessageCancellationAdviceV04 {
	d.Message = new(SecuritiesMessageCancellationAdviceV04)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesMessageCancellationAdvice to an account owner to inform of the cancellation of a securities message previously sent by an account servicer.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
//
// Usage
// The previously sent message may be:
// - a securities settlement transaction confirmation
// - a report (transactions, pending transactions, allegements, accounting and custody securities balance)
// - a allegement notification (when sent by mistake or because the counterparty cancelled its instruction)
// - a portfolio transfer notification
// - an intra-position movement confirmation
// - a transaction generation notification
// The previously sent message cannot be a status advice message (any). If a status advice should not have been sent, a new status advice with the correct status should be sent, not a cancellation advice.
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
//
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesMessageCancellationAdviceV04 struct {

	// Reference to the message advised to be cancelled by the account servicer
	Reference *iso20022.References37Choice `xml:"Ref"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesMessageCancellationAdviceV04) AddReference() *iso20022.References37Choice {
	s.Reference = new(iso20022.References37Choice)
	return s.Reference
}

func (s *SecuritiesMessageCancellationAdviceV04) AddAccountOwner() *iso20022.PartyIdentification36Choice {
	s.AccountOwner = new(iso20022.PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SecuritiesMessageCancellationAdviceV04) AddSafekeepingAccount() *iso20022.SecuritiesAccount13 {
	s.SafekeepingAccount = new(iso20022.SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesMessageCancellationAdviceV04) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
