package iso20022

// Choice between selected investment plans issued during previous years or the entirety of the investment plans.
type PreviousYear1 struct {

	// Selection ot the entirety of the investment plans.
	AllPreviousYears *PreviousAll `xml:"AllPrvsYrs"`

	// Selection of investment plans issued during previous years.
	SpecificPreviousYears []*ISOYear `xml:"SpcfcPrvsYrs"`

	// Indicates whether the ISA contains a cash component asset for transfer.
	CashComponentIndicator *YesNoIndicator `xml:"CshCmpntInd"`
}

func (p *PreviousYear1) SetAllPreviousYears(value string) {
	p.AllPreviousYears = (*PreviousAll)(&value)
}

func (p *PreviousYear1) AddSpecificPreviousYears(value string) {
	p.SpecificPreviousYears = append(p.SpecificPreviousYears, (*ISOYear)(&value))
}

func (p *PreviousYear1) SetCashComponentIndicator(value string) {
	p.CashComponentIndicator = (*YesNoIndicator)(&value)
}

// Choice between selected investment plans issued during previous years or the entirety of the investment plans.
type PreviousYear2 struct {

	// Selection of investment plans issued during previous years.
	PreviousYears *PreviousYear1Choice `xml:"PrvsYrs"`

	// Indicates whether the ISA contains a cash component asset for transfer.
	CashComponentIndicator *YesNoIndicator `xml:"CshCmpntInd"`
}

func (p *PreviousYear2) AddPreviousYears() *PreviousYear1Choice {
	p.PreviousYears = new(PreviousYear1Choice)
	return p.PreviousYears
}

func (p *PreviousYear2) SetCashComponentIndicator(value string) {
	p.CashComponentIndicator = (*YesNoIndicator)(&value)
}

// Choice between selected investment plans issued during previous years or the entirety of the investment plans.
type PreviousYear3 struct {

	// Choice between selected investment plans issued during previous years or the entirety of the investment plans.
	PreviousYear *PreviousYear1Choice `xml:"PrvsYr"`

	// Indicates whether the ISA contains a cash component asset for transfer.
	CashComponentIndicator *YesNoIndicator `xml:"CshCmpntInd"`
}

func (p *PreviousYear3) AddPreviousYear() *PreviousYear1Choice {
	p.PreviousYear = new(PreviousYear1Choice)
	return p.PreviousYear
}

func (p *PreviousYear3) SetCashComponentIndicator(value string) {
	p.CashComponentIndicator = (*YesNoIndicator)(&value)
}
