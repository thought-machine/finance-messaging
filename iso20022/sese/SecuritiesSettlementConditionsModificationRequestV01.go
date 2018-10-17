package sese

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document03000101 struct {
	XMLName xml.Name                                              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.030.001.01 Document"`
	Message *SecuritiesSettlementConditionsModificationRequestV01 `xml:"SctiesSttlmCondsModReq"`
}

func (d *Document03000101) AddMessage() *SecuritiesSettlementConditionsModificationRequestV01 {
	d.Message = new(SecuritiesSettlementConditionsModificationRequestV01)
	return d.Message
}

// Scope
// An account owner sends a SecuritiesSettlementConditionsModificationRequest to an account servicer to request the modification of a processing indicator or another non-matching information.
// The account owner/servicer relationship may be:
// - a central securities depository participant which has an account with a central securities depository.
// It could also be, if agreed in a service level agreement:
// - a global custodian which has an account with its local agent (sub-custodian), or
// - an investment management institution which manage a fund account opened at a custodian, or
// - a broker which has an account with a custodian, or
// - a central securities depository which has an account with a custodian, another central securities depository or another settlement market infrastructure.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// This message cannot be used to request the modification of trade or event details.
// The use of AdditionalInformation and its fields must be pre-agreed between account servicer and account owner. The fields in that sequence cannot be used to amend a trade or event detail unless authorised by country market practice.
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesSettlementConditionsModificationRequestV01 struct {

	// Information that unambiguously identifies a SettlementConditionModification transaction and a SecuritiesSettlementConditionsModificationRequest message as known by the account owner (or the instructing party acting on its behalf).
	Identification *iso20022.DocumentIdentification11 `xml:"Id"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Details of the request.
	RequestDetails []*iso20022.RequestDetails1 `xml:"ReqDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	AdditionalInformation []*iso20022.AdditionalInformation3 `xml:"AddtlInf,omitempty"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *iso20022.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *iso20022.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension2 `xml:"Xtnsn,omitempty"`
}

func (s *SecuritiesSettlementConditionsModificationRequestV01) AddIdentification() *iso20022.DocumentIdentification11 {
	s.Identification = new(iso20022.DocumentIdentification11)
	return s.Identification
}

func (s *SecuritiesSettlementConditionsModificationRequestV01) AddAccountOwner() *iso20022.PartyIdentification13Choice {
	s.AccountOwner = new(iso20022.PartyIdentification13Choice)
	return s.AccountOwner
}

func (s *SecuritiesSettlementConditionsModificationRequestV01) AddSafekeepingAccount() *iso20022.SecuritiesAccount13 {
	s.SafekeepingAccount = new(iso20022.SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesSettlementConditionsModificationRequestV01) AddRequestDetails() *iso20022.RequestDetails1 {
	newValue := new(iso20022.RequestDetails1)
	s.RequestDetails = append(s.RequestDetails, newValue)
	return newValue
}

func (s *SecuritiesSettlementConditionsModificationRequestV01) AddAdditionalInformation() *iso20022.AdditionalInformation3 {
	newValue := new(iso20022.AdditionalInformation3)
	s.AdditionalInformation = append(s.AdditionalInformation, newValue)
	return newValue
}

func (s *SecuritiesSettlementConditionsModificationRequestV01) AddMessageOriginator() *iso20022.PartyIdentification10Choice {
	s.MessageOriginator = new(iso20022.PartyIdentification10Choice)
	return s.MessageOriginator
}

func (s *SecuritiesSettlementConditionsModificationRequestV01) AddMessageRecipient() *iso20022.PartyIdentification10Choice {
	s.MessageRecipient = new(iso20022.PartyIdentification10Choice)
	return s.MessageRecipient
}

func (s *SecuritiesSettlementConditionsModificationRequestV01) AddExtension() *iso20022.Extension2 {
	newValue := new(iso20022.Extension2)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
