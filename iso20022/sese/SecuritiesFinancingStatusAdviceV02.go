package sese

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document03400102 struct {
	XMLName xml.Name                            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.034.001.02 Document"`
	Message *SecuritiesFinancingStatusAdviceV02 `xml:"SctiesFincgStsAdvc"`
}

func (d *Document03400102) AddMessage() *SecuritiesFinancingStatusAdviceV02 {
	d.Message = new(SecuritiesFinancingStatusAdviceV02)
	return d.Message
}

// Scope
// An securities financing transaction account servicer sends a SecuritiesFinancingStatusAdvice to an account owner to advise the status of a securities financing transaction previously instructed by the account owner.
// The status advice may be sent as a response to the request of the account owner or not.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure managing securities financing transactions on behalf of their participants
// - an agent (sub-custodian) managing securities financing transactions on behalf of their global custodian customer, or
// - a custodian managing securities financing transactions on behalf of an investment management institution or a broker/dealer.
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information.
// using the relevant elements in the Business Application Header.
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesFinancingStatusAdviceV02 struct {

	// Provides unambiguous transaction identification information.
	TransactionIdentification *iso20022.TransactionIdentifications1 `xml:"TxId"`

	// Processing status of the transaction.
	ProcessingStatus *iso20022.ProcessingStatus3Choice `xml:"PrcgSts,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *iso20022.MatchingStatus3Choice `xml:"MtchgSts,omitempty"`

	// Provides the matching status of an instruction as per the account servicer based on an allegement. At this time no matching took place on the market (at the CSD/ICSD).
	InferredMatchingStatus *iso20022.MatchingStatus3Choice `xml:"IfrrdMtchgSts,omitempty"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *iso20022.SettlementStatus2Choice `xml:"SttlmSts,omitempty"`

	// Provides the status of the repurchase agreement call request.
	RepoCallRequestStatus *iso20022.RepoCallRequestStatus2Choice `xml:"RepoCallReqSts,omitempty"`

	// Identifies the details of the transaction.
	TransactionDetails *iso20022.SecuritiesFinancingTransactionDetails8 `xml:"TxDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesFinancingStatusAdviceV02) AddTransactionIdentification() *iso20022.TransactionIdentifications1 {
	s.TransactionIdentification = new(iso20022.TransactionIdentifications1)
	return s.TransactionIdentification
}

func (s *SecuritiesFinancingStatusAdviceV02) AddProcessingStatus() *iso20022.ProcessingStatus3Choice {
	s.ProcessingStatus = new(iso20022.ProcessingStatus3Choice)
	return s.ProcessingStatus
}

func (s *SecuritiesFinancingStatusAdviceV02) AddMatchingStatus() *iso20022.MatchingStatus3Choice {
	s.MatchingStatus = new(iso20022.MatchingStatus3Choice)
	return s.MatchingStatus
}

func (s *SecuritiesFinancingStatusAdviceV02) AddInferredMatchingStatus() *iso20022.MatchingStatus3Choice {
	s.InferredMatchingStatus = new(iso20022.MatchingStatus3Choice)
	return s.InferredMatchingStatus
}

func (s *SecuritiesFinancingStatusAdviceV02) AddSettlementStatus() *iso20022.SettlementStatus2Choice {
	s.SettlementStatus = new(iso20022.SettlementStatus2Choice)
	return s.SettlementStatus
}

func (s *SecuritiesFinancingStatusAdviceV02) AddRepoCallRequestStatus() *iso20022.RepoCallRequestStatus2Choice {
	s.RepoCallRequestStatus = new(iso20022.RepoCallRequestStatus2Choice)
	return s.RepoCallRequestStatus
}

func (s *SecuritiesFinancingStatusAdviceV02) AddTransactionDetails() *iso20022.SecuritiesFinancingTransactionDetails8 {
	s.TransactionDetails = new(iso20022.SecuritiesFinancingTransactionDetails8)
	return s.TransactionDetails
}

func (s *SecuritiesFinancingStatusAdviceV02) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
