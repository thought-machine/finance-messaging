package iso20022

// Information about a price report.
type PriceReport1 struct {

	// Information related to the price valuation of a financial instrument.
	PriceValuationDetails []*PriceValuation2 `xml:"PricValtnDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (p *PriceReport1) AddPriceValuationDetails() *PriceValuation2 {
	newValue := new(PriceValuation2)
	p.PriceValuationDetails = append(p.PriceValuationDetails, newValue)
	return newValue
}

func (p *PriceReport1) AddExtension() *Extension1 {
	newValue := new(Extension1)
	p.Extension = append(p.Extension, newValue)
	return newValue
}

// Information about a price report.
type PriceReport2 struct {

	// Information related to the price valuation of a financial instrument.
	PriceValuationDetails []*PriceValuation3 `xml:"PricValtnDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (p *PriceReport2) AddPriceValuationDetails() *PriceValuation3 {
	newValue := new(PriceValuation3)
	p.PriceValuationDetails = append(p.PriceValuationDetails, newValue)
	return newValue
}

func (p *PriceReport2) AddExtension() *Extension1 {
	newValue := new(Extension1)
	p.Extension = append(p.Extension, newValue)
	return newValue
}

// Information about a price report.
type PriceReport3 struct {

	// Information related to the price valuation of a financial instrument.
	PriceValuationDetails []*PriceValuation4 `xml:"PricValtnDtls"`
}

func (p *PriceReport3) AddPriceValuationDetails() *PriceValuation4 {
	newValue := new(PriceValuation4)
	p.PriceValuationDetails = append(p.PriceValuationDetails, newValue)
	return newValue
}
