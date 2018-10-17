package sese

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document02600105 struct {
	XMLName xml.Name                                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.026.001.05 Document"`
	Message *SecuritiesSettlementTransactionReversalAdviceV05 `xml:"SctiesSttlmTxRvslAdvc"`
}

func (d *Document02600105) AddMessage() *SecuritiesSettlementTransactionReversalAdviceV05 {
	d.Message = new(SecuritiesSettlementTransactionReversalAdviceV05)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementTransactionReversalAdvice to an account owner to reverse the confirmation of a partial or full delivery or receipt of financial instruments, free or against of payment, physically or by book-entry.
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
type SecuritiesSettlementTransactionReversalAdviceV05 struct {

	// Provides transaction type and identification information.
	TransactionIdentificationDetails *iso20022.SettlementTypeAndIdentification19 `xml:"TxIdDtls"`

	// Reference to the unambiguous identification of the confirmation as per the account servicer.
	ConfirmationReference *iso20022.Identification14 `xml:"ConfRef"`

	// Additional parameters to the transaction.
	AdditionalParameters *iso20022.AdditionalParameters23 `xml:"AddtlParams,omitempty"`

	// Details of the trade.
	TradeDetails *iso20022.SecuritiesTradeDetails53 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *iso20022.SecurityIdentification19 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *iso20022.FinancialInstrumentAttributes64 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *iso20022.QuantityAndAccount41 `xml:"QtyAndAcctDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *iso20022.SettlementDetails90 `xml:"SttlmParams"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *iso20022.StandingSettlementInstruction11 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *iso20022.SettlementParties36 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *iso20022.SettlementParties36 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *iso20022.CashParties25 `xml:"CshPties,omitempty"`

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

func (s *SecuritiesSettlementTransactionReversalAdviceV05) AddTransactionIdentificationDetails() *iso20022.SettlementTypeAndIdentification19 {
	s.TransactionIdentificationDetails = new(iso20022.SettlementTypeAndIdentification19)
	return s.TransactionIdentificationDetails
}

func (s *SecuritiesSettlementTransactionReversalAdviceV05) AddConfirmationReference() *iso20022.Identification14 {
	s.ConfirmationReference = new(iso20022.Identification14)
	return s.ConfirmationReference
}

func (s *SecuritiesSettlementTransactionReversalAdviceV05) AddAdditionalParameters() *iso20022.AdditionalParameters23 {
	s.AdditionalParameters = new(iso20022.AdditionalParameters23)
	return s.AdditionalParameters
}

func (s *SecuritiesSettlementTransactionReversalAdviceV05) AddTradeDetails() *iso20022.SecuritiesTradeDetails53 {
	s.TradeDetails = new(iso20022.SecuritiesTradeDetails53)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionReversalAdviceV05) AddFinancialInstrumentIdentification() *iso20022.SecurityIdentification19 {
	s.FinancialInstrumentIdentification = new(iso20022.SecurityIdentification19)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionReversalAdviceV05) AddFinancialInstrumentAttributes() *iso20022.FinancialInstrumentAttributes64 {
	s.FinancialInstrumentAttributes = new(iso20022.FinancialInstrumentAttributes64)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionReversalAdviceV05) AddQuantityAndAccountDetails() *iso20022.QuantityAndAccount41 {
	s.QuantityAndAccountDetails = new(iso20022.QuantityAndAccount41)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionReversalAdviceV05) AddSettlementParameters() *iso20022.SettlementDetails90 {
	s.SettlementParameters = new(iso20022.SettlementDetails90)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionReversalAdviceV05) AddStandingSettlementInstructionDetails() *iso20022.StandingSettlementInstruction11 {
	s.StandingSettlementInstructionDetails = new(iso20022.StandingSettlementInstruction11)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionReversalAdviceV05) AddDeliveringSettlementParties() *iso20022.SettlementParties36 {
	s.DeliveringSettlementParties = new(iso20022.SettlementParties36)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionReversalAdviceV05) AddReceivingSettlementParties() *iso20022.SettlementParties36 {
	s.ReceivingSettlementParties = new(iso20022.SettlementParties36)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionReversalAdviceV05) AddCashParties() *iso20022.CashParties25 {
	s.CashParties = new(iso20022.CashParties25)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionReversalAdviceV05) AddSettledAmount() *iso20022.AmountAndDirection46 {
	s.SettledAmount = new(iso20022.AmountAndDirection46)
	return s.SettledAmount
}

func (s *SecuritiesSettlementTransactionReversalAdviceV05) AddOtherAmounts() *iso20022.OtherAmounts30 {
	s.OtherAmounts = new(iso20022.OtherAmounts30)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionReversalAdviceV05) AddOtherBusinessParties() *iso20022.OtherParties27 {
	s.OtherBusinessParties = new(iso20022.OtherParties27)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionReversalAdviceV05) AddAdditionalPhysicalOrRegistrationDetails() *iso20022.RegistrationParameters4 {
	s.AdditionalPhysicalOrRegistrationDetails = new(iso20022.RegistrationParameters4)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionReversalAdviceV05) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
