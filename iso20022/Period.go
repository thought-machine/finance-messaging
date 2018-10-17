package iso20022

// Time span defined by a start date and time, and an end date and time.
type Period1 struct {

	// Date and time at which the range starts.
	StartDate *DateFormat4Choice `xml:"StartDt"`

	// Date and time at which the range ends.
	EndDate *DateFormat4Choice `xml:"EndDt"`
}

func (p *Period1) AddStartDate() *DateFormat4Choice {
	p.StartDate = new(DateFormat4Choice)
	return p.StartDate
}

func (p *Period1) AddEndDate() *DateFormat4Choice {
	p.EndDate = new(DateFormat4Choice)
	return p.EndDate
}

// Time span defined by a start date and time, and an end date and time.
type Period2 struct {

	// Date and time at which the range starts.
	FromDate *ISODate `xml:"FrDt"`

	// Date and time at which the range ends.
	ToDate *ISODate `xml:"ToDt"`
}

func (p *Period2) SetFromDate(value string) {
	p.FromDate = (*ISODate)(&value)
}

func (p *Period2) SetToDate(value string) {
	p.ToDate = (*ISODate)(&value)
}

// Time span defined by a start date and time, and an end date and time.
type Period3 struct {

	// Date and time at which the range starts.
	StartDate *DateFormat12Choice `xml:"StartDt"`

	// Date and time at which the range ends.
	EndDate *DateFormat12Choice `xml:"EndDt"`
}

func (p *Period3) AddStartDate() *DateFormat12Choice {
	p.StartDate = new(DateFormat12Choice)
	return p.StartDate
}

func (p *Period3) AddEndDate() *DateFormat12Choice {
	p.EndDate = new(DateFormat12Choice)
	return p.EndDate
}

// Time span defined by a start date and time, and an end date and time.
type Period4 struct {

	// Date and time at which the range starts.
	StartDate *DateFormat18Choice `xml:"StartDt"`

	// Date and time at which the range ends.
	EndDate *DateFormat18Choice `xml:"EndDt"`
}

func (p *Period4) AddStartDate() *DateFormat18Choice {
	p.StartDate = new(DateFormat18Choice)
	return p.StartDate
}

func (p *Period4) AddEndDate() *DateFormat18Choice {
	p.EndDate = new(DateFormat18Choice)
	return p.EndDate
}

// Time span defined by a start date and time, and an end date and time.
type Period5 struct {

	// Date and time at which the range starts.
	StartDate *DateFormat21Choice `xml:"StartDt"`

	// Date and time at which the range ends.
	EndDate *DateFormat21Choice `xml:"EndDt"`
}

func (p *Period5) AddStartDate() *DateFormat21Choice {
	p.StartDate = new(DateFormat21Choice)
	return p.StartDate
}

func (p *Period5) AddEndDate() *DateFormat21Choice {
	p.EndDate = new(DateFormat21Choice)
	return p.EndDate
}
