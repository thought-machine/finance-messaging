package tsin

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document01100101 struct {
	XMLName xml.Name                                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsin.011.001.01 Document"`
	Message *PartyRegistrationAndGuaranteeNotificationV01 `xml:"PtyRegnAndGrntNtfctn"`
}

func (d *Document01100101) AddMessage() *PartyRegistrationAndGuaranteeNotificationV01 {
	d.Message = new(PartyRegistrationAndGuaranteeNotificationV01)
	return d.Message
}

// The PartyRegistrationAndGuaranteeNotification message is sent by a factoring client or a financial service to a trade partner and, optionally, to an interested party in order to notify the status of a requested financial service agreement. The trade partner is given information to explain the consequences of a financial service agreement, for instance, the trade partner must pay the financial institution and must use the electronic address to inform it and pay it using the bank account given.
// The message may reference related messages and may include referenced data.
// The message can carry digital signatures if required by context.
type PartyRegistrationAndGuaranteeNotificationV01 struct {

	// Set of characteristics that unambiguously identify the notification, common parameters, documents and identifications.
	Header *iso20022.BusinessLetter1 `xml:"Hdr"`

	// List of otifications.
	NotificationList []*iso20022.FinancingAgreementList1 `xml:"NtfctnList"`

	// Number of notification lists as control value.
	NotificationCount *iso20022.Max15NumericText `xml:"NtfctnCnt,omitempty"`

	// Total number of individual items in all lists.
	ItemCount *iso20022.Max15NumericText `xml:"ItmCnt,omitempty"`

	// Total of all individual amounts included in all lists, irrespective of currencies or direction.
	ControlSum *iso20022.DecimalNumber `xml:"CtrlSum,omitempty"`

	// Referenced or related business message.
	AttachedMessage []*iso20022.EncapsulatedBusinessMessage1 `xml:"AttchdMsg,omitempty"`
}

func (p *PartyRegistrationAndGuaranteeNotificationV01) AddHeader() *iso20022.BusinessLetter1 {
	p.Header = new(iso20022.BusinessLetter1)
	return p.Header
}

func (p *PartyRegistrationAndGuaranteeNotificationV01) AddNotificationList() *iso20022.FinancingAgreementList1 {
	newValue := new(iso20022.FinancingAgreementList1)
	p.NotificationList = append(p.NotificationList, newValue)
	return newValue
}

func (p *PartyRegistrationAndGuaranteeNotificationV01) SetNotificationCount(value string) {
	p.NotificationCount = (*iso20022.Max15NumericText)(&value)
}

func (p *PartyRegistrationAndGuaranteeNotificationV01) SetItemCount(value string) {
	p.ItemCount = (*iso20022.Max15NumericText)(&value)
}

func (p *PartyRegistrationAndGuaranteeNotificationV01) SetControlSum(value string) {
	p.ControlSum = (*iso20022.DecimalNumber)(&value)
}

func (p *PartyRegistrationAndGuaranteeNotificationV01) AddAttachedMessage() *iso20022.EncapsulatedBusinessMessage1 {
	newValue := new(iso20022.EncapsulatedBusinessMessage1)
	p.AttachedMessage = append(p.AttachedMessage, newValue)
	return newValue
}
