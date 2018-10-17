package reda

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document05600101 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.056.001.01 Document"`
	Message *StandingSettlementInstructionV01 `xml:"StgSttlmInstr"`
}

func (d *Document05600101) AddMessage() *StandingSettlementInstructionV01 {
	d.Message = new(StandingSettlementInstructionV01)
	return d.Message
}

// Scope
// An instructing party sends the StandingSettlementInstruction (SSI) message to the receiver to create or update a standing cash or securities settlement instruction. The message can also be used to notify a counterparty of an SSI.
//
// Usage
// The instructing party (initiator) is:
// •	An account servicer, for example, a global custodian or prime broker
// •	A counterparty in a transaction, for example:
// -	an investment manager (executing broker),
// -	a global custodian (executing broker, prime broker)
// •	A vendor’s application communicating on behalf of an account servicer or counterparty
// The receiver is:
// •	An account owner, for example, an investment manager, hedge fund administrator or a party to which SSI operations have been outsourced
// •	A counterparty, for example:
// -	an investment manager (executing broker)
// -	a global custodian (executing broker, prime broker)
// •	A vendor’s application communicating on behalf of the account owner or counterparty
type StandingSettlementInstructionV01 struct {

	// Reference of this message.
	MessageReferenceIdentification *iso20022.Max35Text `xml:"MsgRefId"`

	// Date on which the SSI is effective.
	EffectiveDateDetails *iso20022.EffectiveDate1 `xml:"FctvDtDtls,omitempty"`

	// Unique and unambiguous master identification known to the sender (or its authorised agent) and receiver (or its authorised agent), below which the SSI will be lodged. This may be an account number or reference to a fund.
	// If no account or reference is available then “NONREF” must be specified.
	AccountIdentification []*iso20022.AccountIdentification26 `xml:"AcctId"`

	// Identifies the market for the standing settlement instruction.
	MarketIdentification *iso20022.MarketIdentificationOrCashPurpose1Choice `xml:"MktId"`

	// Currency for which the SSI is specified.
	SettlementCurrency *iso20022.ActiveCurrencyCode `xml:"SttlmCcy,omitempty"`

	// Settlement chain parties, accounts and other details.
	SettlementDetails *iso20022.SecuritiesOrCash1Choice `xml:"SttlmDtls"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *StandingSettlementInstructionV01) SetMessageReferenceIdentification(value string) {
	s.MessageReferenceIdentification = (*iso20022.Max35Text)(&value)
}

func (s *StandingSettlementInstructionV01) AddEffectiveDateDetails() *iso20022.EffectiveDate1 {
	s.EffectiveDateDetails = new(iso20022.EffectiveDate1)
	return s.EffectiveDateDetails
}

func (s *StandingSettlementInstructionV01) AddAccountIdentification() *iso20022.AccountIdentification26 {
	newValue := new(iso20022.AccountIdentification26)
	s.AccountIdentification = append(s.AccountIdentification, newValue)
	return newValue
}

func (s *StandingSettlementInstructionV01) AddMarketIdentification() *iso20022.MarketIdentificationOrCashPurpose1Choice {
	s.MarketIdentification = new(iso20022.MarketIdentificationOrCashPurpose1Choice)
	return s.MarketIdentification
}

func (s *StandingSettlementInstructionV01) SetSettlementCurrency(value string) {
	s.SettlementCurrency = (*iso20022.ActiveCurrencyCode)(&value)
}

func (s *StandingSettlementInstructionV01) AddSettlementDetails() *iso20022.SecuritiesOrCash1Choice {
	s.SettlementDetails = new(iso20022.SecuritiesOrCash1Choice)
	return s.SettlementDetails
}

func (s *StandingSettlementInstructionV01) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
