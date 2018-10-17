package camt

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document04100102 struct {
	XMLName xml.Name                            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.041.001.02 Document"`
	Message *FundConfirmedCashForecastReportV02 `xml:"camt.041.001.02"`
}

func (d *Document04100102) AddMessage() *FundConfirmedCashForecastReportV02 {
	d.Message = new(FundConfirmedCashForecastReportV02)
	return d.Message
}

// Scope
// The FundConfirmedCashForecastReport message is sent by a report provider, such as a transfer agent or a designated agent of the fund, to a report user, such as an investment manager, a fund accountant or any other interested party.
// This message is used to report the confirmed cash incomings and outgoings of one or more investment funds, on one or more trade dates. These cash movements may result from, for example, redemption, subscription, switch transactions or dividends.
// Usage
// The FundConfirmedCashForecastReport message is used to provide definitive cash movements, that is it is sent after the cut-off time and/or the price valuation of the fund.
type FundConfirmedCashForecastReportV02 struct {

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference []*iso20022.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Information related to the cash-in and cash-out flows for a specific trade date as a result of investment fund transactions, for example, subscriptions, redemptions or switches to/from a specified investment fund.
	//
	//
	FundCashForecastDetails []*iso20022.FundCashForecast1 `xml:"FndCshFcstDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`
}

func (f *FundConfirmedCashForecastReportV02) AddPoolReference() *iso20022.AdditionalReference3 {
	f.PoolReference = new(iso20022.AdditionalReference3)
	return f.PoolReference
}

func (f *FundConfirmedCashForecastReportV02) AddPreviousReference() *iso20022.AdditionalReference3 {
	newValue := new(iso20022.AdditionalReference3)
	f.PreviousReference = append(f.PreviousReference, newValue)
	return newValue
}

func (f *FundConfirmedCashForecastReportV02) AddRelatedReference() *iso20022.AdditionalReference3 {
	newValue := new(iso20022.AdditionalReference3)
	f.RelatedReference = append(f.RelatedReference, newValue)
	return newValue
}

func (f *FundConfirmedCashForecastReportV02) AddFundCashForecastDetails() *iso20022.FundCashForecast1 {
	newValue := new(iso20022.FundCashForecast1)
	f.FundCashForecastDetails = append(f.FundCashForecastDetails, newValue)
	return newValue
}

func (f *FundConfirmedCashForecastReportV02) AddExtension() *iso20022.Extension1 {
	newValue := new(iso20022.Extension1)
	f.Extension = append(f.Extension, newValue)
	return newValue
}
