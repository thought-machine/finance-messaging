package supl

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document02700101 struct {
	XMLName xml.Name                   `xml:"urn:iso:std:iso:20022:tech:xsd:supl.027.001.01 Document"`
	Message *InformationResponseSD1V01 `xml:"InfRspnSD1"`
}

func (d *Document02700101) AddMessage() *InformationResponseSD1V01 {
	d.Message = new(InformationResponseSD1V01)
	return d.Message
}

// This extends the message InformationRequestResponse.
type InformationResponseSD1V01 struct {

	// Information used to identify the request.
	InvestigationIdentification *iso20022.Max35Text `xml:"InvstgtnId"`

	// Date and time of creation of the extension.
	CreationDateTime *iso20022.ISODateTime `xml:"CreDtTm"`

	// Identifies the account servicing institution.
	AccountServicerIdentification *iso20022.BranchAndFinancialInstitutionIdentification4 `xml:"AcctSvcrId"`

	// Requested account and its owners.
	AccountAndParties []*iso20022.AccountAndParties2 `xml:"AcctAndPties"`
}

func (i *InformationResponseSD1V01) SetInvestigationIdentification(value string) {
	i.InvestigationIdentification = (*iso20022.Max35Text)(&value)
}

func (i *InformationResponseSD1V01) SetCreationDateTime(value string) {
	i.CreationDateTime = (*iso20022.ISODateTime)(&value)
}

func (i *InformationResponseSD1V01) AddAccountServicerIdentification() *iso20022.BranchAndFinancialInstitutionIdentification4 {
	i.AccountServicerIdentification = new(iso20022.BranchAndFinancialInstitutionIdentification4)
	return i.AccountServicerIdentification
}

func (i *InformationResponseSD1V01) AddAccountAndParties() *iso20022.AccountAndParties2 {
	newValue := new(iso20022.AccountAndParties2)
	i.AccountAndParties = append(i.AccountAndParties, newValue)
	return newValue
}
