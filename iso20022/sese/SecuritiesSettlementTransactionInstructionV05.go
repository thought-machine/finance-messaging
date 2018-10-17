package sese

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document02300105 struct {
	XMLName xml.Name                                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.023.001.05 Document"`
	Message *SecuritiesSettlementTransactionInstructionV05 `xml:"SctiesSttlmTxInstr"`
}

func (d *Document02300105) AddMessage() *SecuritiesSettlementTransactionInstructionV05 {
	d.Message = new(SecuritiesSettlementTransactionInstructionV05)
	return d.Message
}

// Scope
// An account owner sends a SecuritiesSettlementTransactionInstruction to an account servicer to instruct the receipt or delivery of financial instruments with or without payment, physically or by book-entry.
// The account owner/servicer relationship may be:
// - a global custodian which has an account with a local custodian, or
// - an investment management institution which manages a fund account opened at a custodian, or
// - a broker which has an account with a custodian, or
// - a central securities depository participant which has an account with a central securities depository, or
// - a central securities depository which has an account with a custodian, another central securities depository or another settlement market infrastructure, or
// - a central counterparty or a stock exchange or a trade matching utility which need to instruct the settlement of transactions to a central securities depository or another settlement market infrastructure.
//
// Usage
// The instruction may be linked to other settlement instructions, for example, for a turnaround or back-to-back, or other transactions, for example, foreign exchange deal, using the linkage functionality.
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
//
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesSettlementTransactionInstructionV05 struct {

	// Unambiguous identification of the transaction as know by the instructing party.
	TransactionIdentification *iso20022.Max35Text `xml:"TxId"`

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *iso20022.SettlementTypeAndAdditionalParameters9 `xml:"SttlmTpAndAddtlParams"`

	// Count of the number of transactions linked.
	NumberCounts *iso20022.NumberCount1Choice `xml:"NbCounts,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*iso20022.Linkages17 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *iso20022.SecuritiesTradeDetails33 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *iso20022.SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *iso20022.FinancialInstrumentAttributes35 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *iso20022.QuantityAndAccount25 `xml:"QtyAndAcctDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *iso20022.SettlementDetails67 `xml:"SttlmParams"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *iso20022.StandingSettlementInstruction4 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *iso20022.SettlementParties11 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *iso20022.SettlementParties11 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *iso20022.CashParties17 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *iso20022.AmountAndDirection36 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *iso20022.OtherAmounts14 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *iso20022.OtherParties19 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *iso20022.RegistrationParameters1 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionInstructionV05) SetTransactionIdentification(value string) {
	s.TransactionIdentification = (*iso20022.Max35Text)(&value)
}

func (s *SecuritiesSettlementTransactionInstructionV05) AddSettlementTypeAndAdditionalParameters() *iso20022.SettlementTypeAndAdditionalParameters9 {
	s.SettlementTypeAndAdditionalParameters = new(iso20022.SettlementTypeAndAdditionalParameters9)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionInstructionV05) AddNumberCounts() *iso20022.NumberCount1Choice {
	s.NumberCounts = new(iso20022.NumberCount1Choice)
	return s.NumberCounts
}

func (s *SecuritiesSettlementTransactionInstructionV05) AddLinkages() *iso20022.Linkages17 {
	newValue := new(iso20022.Linkages17)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionInstructionV05) AddTradeDetails() *iso20022.SecuritiesTradeDetails33 {
	s.TradeDetails = new(iso20022.SecuritiesTradeDetails33)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionInstructionV05) AddFinancialInstrumentIdentification() *iso20022.SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(iso20022.SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionInstructionV05) AddFinancialInstrumentAttributes() *iso20022.FinancialInstrumentAttributes35 {
	s.FinancialInstrumentAttributes = new(iso20022.FinancialInstrumentAttributes35)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionInstructionV05) AddQuantityAndAccountDetails() *iso20022.QuantityAndAccount25 {
	s.QuantityAndAccountDetails = new(iso20022.QuantityAndAccount25)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionInstructionV05) AddSettlementParameters() *iso20022.SettlementDetails67 {
	s.SettlementParameters = new(iso20022.SettlementDetails67)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionInstructionV05) AddStandingSettlementInstructionDetails() *iso20022.StandingSettlementInstruction4 {
	s.StandingSettlementInstructionDetails = new(iso20022.StandingSettlementInstruction4)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionInstructionV05) AddDeliveringSettlementParties() *iso20022.SettlementParties11 {
	s.DeliveringSettlementParties = new(iso20022.SettlementParties11)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionInstructionV05) AddReceivingSettlementParties() *iso20022.SettlementParties11 {
	s.ReceivingSettlementParties = new(iso20022.SettlementParties11)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionInstructionV05) AddCashParties() *iso20022.CashParties17 {
	s.CashParties = new(iso20022.CashParties17)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionInstructionV05) AddSettlementAmount() *iso20022.AmountAndDirection36 {
	s.SettlementAmount = new(iso20022.AmountAndDirection36)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionInstructionV05) AddOtherAmounts() *iso20022.OtherAmounts14 {
	s.OtherAmounts = new(iso20022.OtherAmounts14)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionInstructionV05) AddOtherBusinessParties() *iso20022.OtherParties19 {
	s.OtherBusinessParties = new(iso20022.OtherParties19)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionInstructionV05) AddAdditionalPhysicalOrRegistrationDetails() *iso20022.RegistrationParameters1 {
	s.AdditionalPhysicalOrRegistrationDetails = new(iso20022.RegistrationParameters1)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionInstructionV05) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
