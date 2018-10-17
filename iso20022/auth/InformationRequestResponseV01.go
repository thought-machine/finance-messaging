package auth

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document00200101 struct {
	XMLName xml.Name                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.002.001.01 Document"`
	Message *InformationRequestResponseV01 `xml:"InfReqRspn"`
}

func (d *Document00200101) AddMessage() *InformationRequestResponseV01 {
	d.Message = new(InformationRequestResponseV01)
	return d.Message
}

// This message is sent by the financial institution to the authorities (police, customs, tax authorities, enforcement authorities) to provide a part or all of the requested information.
// The financial institution previously received a request for financial information in the scope of a financial investigation.
//
// Depending on whether the response can be provided STP within the authorities financial investigations messages, the requested information may be
// •	provided in part or in full within the response message itself, or
// •	only referred to in the response message
type InformationRequestResponseV01 struct {

	// Unique identification for the specific investigation as know by the responding party.
	ResponseIdentification *iso20022.Max35Text `xml:"RspnId"`

	// Unique identification for the specific investigation as known by the requesting party.
	InvestigationIdentification *iso20022.Max35Text `xml:"InvstgtnId"`

	// Provides the status of the response.
	ResponseStatus *iso20022.StatusResponse1Code `xml:"RspnSts"`

	// Specifies the the search criteria for the financial institution to perform the search on. The search criteria can be an account, a customer identification or a payment instrument type.
	SearchCriteria *iso20022.SearchCriteria1Choice `xml:"SchCrit"`

	// Provides the return indicators and the investigation result.
	ReturnIndicator []*iso20022.ReturnIndicator1 `xml:"RtrInd"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *InformationRequestResponseV01) SetResponseIdentification(value string) {
	i.ResponseIdentification = (*iso20022.Max35Text)(&value)
}

func (i *InformationRequestResponseV01) SetInvestigationIdentification(value string) {
	i.InvestigationIdentification = (*iso20022.Max35Text)(&value)
}

func (i *InformationRequestResponseV01) SetResponseStatus(value string) {
	i.ResponseStatus = (*iso20022.StatusResponse1Code)(&value)
}

func (i *InformationRequestResponseV01) AddSearchCriteria() *iso20022.SearchCriteria1Choice {
	i.SearchCriteria = new(iso20022.SearchCriteria1Choice)
	return i.SearchCriteria
}

func (i *InformationRequestResponseV01) AddReturnIndicator() *iso20022.ReturnIndicator1 {
	newValue := new(iso20022.ReturnIndicator1)
	i.ReturnIndicator = append(i.ReturnIndicator, newValue)
	return newValue
}

func (i *InformationRequestResponseV01) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
