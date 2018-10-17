package iso20022

// Proprietary date information.
type ProprietaryDate1 struct {

	// Identifies the type of date reported.
	Type *Max35Text `xml:"Tp"`

	// Date in ISO format.
	Date *ISODate `xml:"Dt"`

	// Date and time in ISO format.
	DateTime *ISODateTime `xml:"DtTm"`
}

func (p *ProprietaryDate1) SetType(value string) {
	p.Type = (*Max35Text)(&value)
}

func (p *ProprietaryDate1) SetDate(value string) {
	p.Date = (*ISODate)(&value)
}

func (p *ProprietaryDate1) SetDateTime(value string) {
	p.DateTime = (*ISODateTime)(&value)
}

// Set of elements used to identify a proprietary date.
type ProprietaryDate2 struct {

	// Specifies the type of date.
	Type *Max35Text `xml:"Tp"`

	// Date in ISO format.
	Date *DateAndDateTimeChoice `xml:"Dt"`
}

func (p *ProprietaryDate2) SetType(value string) {
	p.Type = (*Max35Text)(&value)
}

func (p *ProprietaryDate2) AddDate() *DateAndDateTimeChoice {
	p.Date = new(DateAndDateTimeChoice)
	return p.Date
}
