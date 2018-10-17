package sese

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document03100102 struct {
	XMLName xml.Name                                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.031.001.02 Document"`
	Message *SecuritiesSettlementConditionModificationStatusAdviceV02 `xml:"SctiesSttlmCondModStsAdvc"`
}

func (d *Document03100102) AddMessage() *SecuritiesSettlementConditionModificationStatusAdviceV02 {
	d.Message = new(SecuritiesSettlementConditionModificationStatusAdviceV02)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementConditionsModificationStatusAdvice to an account owner to advise the status of a modification request previously instructed by the account owner.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
// Usage
// A SecuritiesSettlementConditionsModificationRequest may contain requests on multiple transactions. However, one SecuritiesSettlementConditionsModificationStatusAdvice must be sent per transaction modified unless the SecuritiesSettlementConditionsModificationRequest is rejected as a whole.
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// using the relevant elements in the Business Application Header.
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesSettlementConditionModificationStatusAdviceV02 struct {

	// Identification of the SecuritiesSettlementConditionsModificationRequest.
	RequestReference *iso20022.Identification1 `xml:"ReqRef"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount13 `xml:"SfkpgAcct,omitempty"`

	// Details of the request.
	RequestDetails *iso20022.RequestDetails6 `xml:"ReqDtls,omitempty"`

	// Provides details on the processing status of the request.
	ProcessingStatus *iso20022.ProcessingStatus14Choice `xml:"PrcgSts"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementConditionModificationStatusAdviceV02) AddRequestReference() *iso20022.Identification1 {
	s.RequestReference = new(iso20022.Identification1)
	return s.RequestReference
}

func (s *SecuritiesSettlementConditionModificationStatusAdviceV02) AddAccountOwner() *iso20022.PartyIdentification36Choice {
	s.AccountOwner = new(iso20022.PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SecuritiesSettlementConditionModificationStatusAdviceV02) AddSafekeepingAccount() *iso20022.SecuritiesAccount13 {
	s.SafekeepingAccount = new(iso20022.SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesSettlementConditionModificationStatusAdviceV02) AddRequestDetails() *iso20022.RequestDetails6 {
	s.RequestDetails = new(iso20022.RequestDetails6)
	return s.RequestDetails
}

func (s *SecuritiesSettlementConditionModificationStatusAdviceV02) AddProcessingStatus() *iso20022.ProcessingStatus14Choice {
	s.ProcessingStatus = new(iso20022.ProcessingStatus14Choice)
	return s.ProcessingStatus
}

func (s *SecuritiesSettlementConditionModificationStatusAdviceV02) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
