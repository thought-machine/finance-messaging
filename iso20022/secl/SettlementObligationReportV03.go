package secl

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document01000103 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:secl.010.001.03 Document"`
	Message *SettlementObligationReportV03 `xml:"SttlmOblgtnRpt"`
}

func (d *Document01000103) AddMessage() *SettlementObligationReportV03 {
	d.Message = new(SettlementObligationReportV03)
	return d.Message
}

// Scope
// The SettlementObligationReport message is sent by the central counterparty (CCP) to a clearing member to report on the settlement obligation that will be submitted for settlement.
//
// The message definition is intended for use with the ISO20022 Business Application Header.
//
// Usage
// The SettlementObligationReport message may also be sent to a third party processing the settlement obligation(s) on behalf of more than one clearing member.
// The Settlement Obligation Report message is provided per delivery account and per instrument. The report can be provided for one specific delivering party or one specific receiving party. It can also be generated per non clearing member.
type SettlementObligationReportV03 struct {

	// Provides details about the report such as the report identification, the creation date and time.
	ReportParameters *iso20022.ReportParameters4 `xml:"RptParams"`

	// Page number of the message (within a report) and continuation indicator to indicate that the report is to continue or that the message is the last page of the report.
	Pagination *iso20022.Pagination `xml:"Pgntn"`

	// Provides the identification of the clearing member (individual clearing member or general clearing member).
	ClearingMember *iso20022.PartyIdentification35Choice `xml:"ClrMmb,omitempty"`

	// Clearing organisation that will clear the trade.
	// Note: This field allows Clearing Member Firm to segregate flows coming from clearing counterparty's clearing system. Indeed, Clearing Member Firms receive messages from the same system (same sender) and this field allows them to know if the message is related to equities or derivatives.
	ClearingSegment *iso20022.PartyIdentification35Choice `xml:"ClrSgmt,omitempty"`

	// Provides the identification of the account used for netting. This is an account opened by the central counterparty in the name of the clearing member or its settlement agent within the account structure, for settlement purposes (gives information about the clearing member/its settlement agent account at the central securities depository).
	DeliveryAccount *iso20022.SecuritiesAccount19 `xml:"DlvryAcct,omitempty"`

	// Provides details on the settlement obligation report.
	ReportDetails []*iso20022.Report5 `xml:"RptDtls"`

	// Provides details about the receiving parties involved in the settlement chain.
	SettlementParties *iso20022.SettlementParties2Choice `xml:"SttlmPties,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SettlementObligationReportV03) AddReportParameters() *iso20022.ReportParameters4 {
	s.ReportParameters = new(iso20022.ReportParameters4)
	return s.ReportParameters
}

func (s *SettlementObligationReportV03) AddPagination() *iso20022.Pagination {
	s.Pagination = new(iso20022.Pagination)
	return s.Pagination
}

func (s *SettlementObligationReportV03) AddClearingMember() *iso20022.PartyIdentification35Choice {
	s.ClearingMember = new(iso20022.PartyIdentification35Choice)
	return s.ClearingMember
}

func (s *SettlementObligationReportV03) AddClearingSegment() *iso20022.PartyIdentification35Choice {
	s.ClearingSegment = new(iso20022.PartyIdentification35Choice)
	return s.ClearingSegment
}

func (s *SettlementObligationReportV03) AddDeliveryAccount() *iso20022.SecuritiesAccount19 {
	s.DeliveryAccount = new(iso20022.SecuritiesAccount19)
	return s.DeliveryAccount
}

func (s *SettlementObligationReportV03) AddReportDetails() *iso20022.Report5 {
	newValue := new(iso20022.Report5)
	s.ReportDetails = append(s.ReportDetails, newValue)
	return newValue
}

func (s *SettlementObligationReportV03) AddSettlementParties() *iso20022.SettlementParties2Choice {
	s.SettlementParties = new(iso20022.SettlementParties2Choice)
	return s.SettlementParties
}

func (s *SettlementObligationReportV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
