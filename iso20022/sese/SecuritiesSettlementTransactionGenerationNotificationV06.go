package sese

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document03200106 struct {
	XMLName xml.Name                                                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.032.001.06 Document"`
	Message *SecuritiesSettlementTransactionGenerationNotificationV06 `xml:"SctiesSttlmTxGnrtnNtfctn"`
}

func (d *Document03200106) AddMessage() *SecuritiesSettlementTransactionGenerationNotificationV06 {
	d.Message = new(SecuritiesSettlementTransactionGenerationNotificationV06)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementTransactionGenerationNotification to an account owner to advise the account owner of a securities settlement transaction that has been generated/created by the account servicer for the account owner. The reason for creation can range from the automatic transformation of pending settlement instructions following a corporate event to the recovery of an account owner transactions' database initiated by its account servicer.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer
//
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type SecuritiesSettlementTransactionGenerationNotificationV06 struct {

	// Provides transaction type and identification information.
	TransactionIdentificationDetails *iso20022.SettlementTypeAndIdentification19 `xml:"TxIdDtls"`

	// Count of the number of transactions linked.
	NumberCounts *iso20022.NumberCount1Choice `xml:"NbCounts,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*iso20022.Linkages37 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *iso20022.SecuritiesTradeDetails51 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *iso20022.SecurityIdentification19 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *iso20022.FinancialInstrumentAttributes64 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails []*iso20022.QuantityAndAccount39 `xml:"QtyAndAcctDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *iso20022.SettlementDetails101 `xml:"SttlmParams"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *iso20022.SettlementParties36 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *iso20022.SettlementParties36 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *iso20022.CashParties26 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *iso20022.AmountAndDirection46 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *iso20022.OtherAmounts28 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *iso20022.OtherParties27 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *iso20022.RegistrationParameters4 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Specifies the reason why the transaction was generated.
	GeneratedReason []*iso20022.GeneratedReason5 `xml:"GnrtdRsn,omitempty"`

	// Status and reason of the transaction.
	StatusAndReason *iso20022.StatusAndReason28 `xml:"StsAndRsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddTransactionIdentificationDetails() *iso20022.SettlementTypeAndIdentification19 {
	s.TransactionIdentificationDetails = new(iso20022.SettlementTypeAndIdentification19)
	return s.TransactionIdentificationDetails
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddNumberCounts() *iso20022.NumberCount1Choice {
	s.NumberCounts = new(iso20022.NumberCount1Choice)
	return s.NumberCounts
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddLinkages() *iso20022.Linkages37 {
	newValue := new(iso20022.Linkages37)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddTradeDetails() *iso20022.SecuritiesTradeDetails51 {
	s.TradeDetails = new(iso20022.SecuritiesTradeDetails51)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddFinancialInstrumentIdentification() *iso20022.SecurityIdentification19 {
	s.FinancialInstrumentIdentification = new(iso20022.SecurityIdentification19)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddFinancialInstrumentAttributes() *iso20022.FinancialInstrumentAttributes64 {
	s.FinancialInstrumentAttributes = new(iso20022.FinancialInstrumentAttributes64)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddQuantityAndAccountDetails() *iso20022.QuantityAndAccount39 {
	newValue := new(iso20022.QuantityAndAccount39)
	s.QuantityAndAccountDetails = append(s.QuantityAndAccountDetails, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddSettlementParameters() *iso20022.SettlementDetails101 {
	s.SettlementParameters = new(iso20022.SettlementDetails101)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddDeliveringSettlementParties() *iso20022.SettlementParties36 {
	s.DeliveringSettlementParties = new(iso20022.SettlementParties36)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddReceivingSettlementParties() *iso20022.SettlementParties36 {
	s.ReceivingSettlementParties = new(iso20022.SettlementParties36)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddCashParties() *iso20022.CashParties26 {
	s.CashParties = new(iso20022.CashParties26)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddSettlementAmount() *iso20022.AmountAndDirection46 {
	s.SettlementAmount = new(iso20022.AmountAndDirection46)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddOtherAmounts() *iso20022.OtherAmounts28 {
	s.OtherAmounts = new(iso20022.OtherAmounts28)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddOtherBusinessParties() *iso20022.OtherParties27 {
	s.OtherBusinessParties = new(iso20022.OtherParties27)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddAdditionalPhysicalOrRegistrationDetails() *iso20022.RegistrationParameters4 {
	s.AdditionalPhysicalOrRegistrationDetails = new(iso20022.RegistrationParameters4)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddGeneratedReason() *iso20022.GeneratedReason5 {
	newValue := new(iso20022.GeneratedReason5)
	s.GeneratedReason = append(s.GeneratedReason, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddStatusAndReason() *iso20022.StatusAndReason28 {
	s.StatusAndReason = new(iso20022.StatusAndReason28)
	return s.StatusAndReason
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV06) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
