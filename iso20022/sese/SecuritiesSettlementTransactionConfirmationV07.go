package sese

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document02500107 struct {
	XMLName xml.Name                                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.025.001.07 Document"`
	Message *SecuritiesSettlementTransactionConfirmationV07 `xml:"SctiesSttlmTxConf"`
}

func (d *Document02500107) AddMessage() *SecuritiesSettlementTransactionConfirmationV07 {
	d.Message = new(SecuritiesSettlementTransactionConfirmationV07)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementTransactionConfirmation to an account owner to confirm the partial or full delivery or receipt of financial instruments, free or against of payment, physically or by book-entry.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
//
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information.
// using the relevant elements in the Business Application Header.
type SecuritiesSettlementTransactionConfirmationV07 struct {

	// Provides transaction type and identification information.
	TransactionIdentificationDetails *iso20022.SettlementTypeAndIdentification19 `xml:"TxIdDtls"`

	// Link to another transaction - provided for information only.
	Linkages *iso20022.Linkages41 `xml:"Lnkgs,omitempty"`

	// Additional parameters to the transaction.
	AdditionalParameters *iso20022.AdditionalParameters29 `xml:"AddtlParams,omitempty"`

	// Details of the trade.
	TradeDetails *iso20022.SecuritiesTradeDetails53 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *iso20022.SecurityIdentification19 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *iso20022.FinancialInstrumentAttributes64 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *iso20022.QuantityAndAccount41 `xml:"QtyAndAcctDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *iso20022.SettlementDetails128 `xml:"SttlmParams"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *iso20022.StandingSettlementInstruction11 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *iso20022.SettlementParties36 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *iso20022.SettlementParties36 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *iso20022.CashParties26 `xml:"CshPties,omitempty"`

	// Amount effectively settled and which will be credited to/debited from the account owner's cash account. It may differ from the instructed settlement amount based on market tolerance level.
	SettledAmount *iso20022.AmountAndDirection46 `xml:"SttldAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *iso20022.OtherAmounts30 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *iso20022.OtherParties27 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *iso20022.RegistrationParameters4 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddTransactionIdentificationDetails() *iso20022.SettlementTypeAndIdentification19 {
	s.TransactionIdentificationDetails = new(iso20022.SettlementTypeAndIdentification19)
	return s.TransactionIdentificationDetails
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddLinkages() *iso20022.Linkages41 {
	s.Linkages = new(iso20022.Linkages41)
	return s.Linkages
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddAdditionalParameters() *iso20022.AdditionalParameters29 {
	s.AdditionalParameters = new(iso20022.AdditionalParameters29)
	return s.AdditionalParameters
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddTradeDetails() *iso20022.SecuritiesTradeDetails53 {
	s.TradeDetails = new(iso20022.SecuritiesTradeDetails53)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddFinancialInstrumentIdentification() *iso20022.SecurityIdentification19 {
	s.FinancialInstrumentIdentification = new(iso20022.SecurityIdentification19)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddFinancialInstrumentAttributes() *iso20022.FinancialInstrumentAttributes64 {
	s.FinancialInstrumentAttributes = new(iso20022.FinancialInstrumentAttributes64)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddQuantityAndAccountDetails() *iso20022.QuantityAndAccount41 {
	s.QuantityAndAccountDetails = new(iso20022.QuantityAndAccount41)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddSettlementParameters() *iso20022.SettlementDetails128 {
	s.SettlementParameters = new(iso20022.SettlementDetails128)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddStandingSettlementInstructionDetails() *iso20022.StandingSettlementInstruction11 {
	s.StandingSettlementInstructionDetails = new(iso20022.StandingSettlementInstruction11)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddDeliveringSettlementParties() *iso20022.SettlementParties36 {
	s.DeliveringSettlementParties = new(iso20022.SettlementParties36)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddReceivingSettlementParties() *iso20022.SettlementParties36 {
	s.ReceivingSettlementParties = new(iso20022.SettlementParties36)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddCashParties() *iso20022.CashParties26 {
	s.CashParties = new(iso20022.CashParties26)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddSettledAmount() *iso20022.AmountAndDirection46 {
	s.SettledAmount = new(iso20022.AmountAndDirection46)
	return s.SettledAmount
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddOtherAmounts() *iso20022.OtherAmounts30 {
	s.OtherAmounts = new(iso20022.OtherAmounts30)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddOtherBusinessParties() *iso20022.OtherParties27 {
	s.OtherBusinessParties = new(iso20022.OtherParties27)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddAdditionalPhysicalOrRegistrationDetails() *iso20022.RegistrationParameters4 {
	s.AdditionalPhysicalOrRegistrationDetails = new(iso20022.RegistrationParameters4)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionConfirmationV07) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
