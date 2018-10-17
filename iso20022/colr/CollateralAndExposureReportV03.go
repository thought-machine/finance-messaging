package colr

import (
	"encoding/xml"

	"github.com/thought-machine/finance-messaging/iso20022"
)

type Document01600103 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:colr.016.001.03 Document"`
	Message *CollateralAndExposureReportV03 `xml:"CollAndXpsrRpt"`
}

func (d *Document01600103) AddMessage() *CollateralAndExposureReportV03 {
	d.Message = new(CollateralAndExposureReportV03)
	return d.Message
}

// Scope
// The CollateralAndExposureReport message is sent:
// - either by the collateral giver, or its collateral manager, to the collateral taker, or its collateral manager, or
// - or by the collateral taker, or its collateral manager to the collateral giver, or its collateral manager
//
// The message definition is intended for use with the ISO20022 Business Application Header.
//
// Usage
// The CollateralAndExposureReport is used to provide the details of the valuation of the collateral, that is, the valuation of securities collateral, cash collateral or other type of collateral, posted at a specific calculation date.
type CollateralAndExposureReportV03 struct {

	// Provides information about the report such as the report identification, the report date and time or the report frequency.
	ReportParameters *iso20022.ReportParameters5 `xml:"RptParams"`

	// Specifies the page number and an indicator of whether it is the only or last page, or if there are additional pages.
	//
	Pagination *iso20022.Pagination `xml:"Pgntn,omitempty"`

	// Provides information like the identification of the party or parties associated with the collateral agreement, the exposure type and the valuation date.
	Obligation *iso20022.Obligation4 `xml:"Oblgtn"`

	// Agreement details for the over the counter market.
	Agreement *iso20022.Agreement4 `xml:"Agrmt,omitempty"`

	// Provides details on the collateral report.
	CollateralReport []*iso20022.Collateral13 `xml:"CollRpt"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CollateralAndExposureReportV03) AddReportParameters() *iso20022.ReportParameters5 {
	c.ReportParameters = new(iso20022.ReportParameters5)
	return c.ReportParameters
}

func (c *CollateralAndExposureReportV03) AddPagination() *iso20022.Pagination {
	c.Pagination = new(iso20022.Pagination)
	return c.Pagination
}

func (c *CollateralAndExposureReportV03) AddObligation() *iso20022.Obligation4 {
	c.Obligation = new(iso20022.Obligation4)
	return c.Obligation
}

func (c *CollateralAndExposureReportV03) AddAgreement() *iso20022.Agreement4 {
	c.Agreement = new(iso20022.Agreement4)
	return c.Agreement
}

func (c *CollateralAndExposureReportV03) AddCollateralReport() *iso20022.Collateral13 {
	newValue := new(iso20022.Collateral13)
	c.CollateralReport = append(c.CollateralReport, newValue)
	return newValue
}

func (c *CollateralAndExposureReportV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
