package iso20022

// Range of time defined by a start date and an end date.
type DatePeriodDetails struct {

	// Start date of the range.
	FromDate *ISODate `xml:"FrDt"`

	// End date of the range.
	ToDate *ISODate `xml:"ToDt"`
}

func (d *DatePeriodDetails) SetFromDate(value string) {
	d.FromDate = (*ISODate)(&value)
}

func (d *DatePeriodDetails) SetToDate(value string) {
	d.ToDate = (*ISODate)(&value)
}

// Range of time defined by a start date and an end date.
type DatePeriodDetails1 struct {

	// Start date of the range.
	FromDate *ISODate `xml:"FrDt"`

	// End date of the range.
	ToDate *ISODate `xml:"ToDt,omitempty"`
}

func (d *DatePeriodDetails1) SetFromDate(value string) {
	d.FromDate = (*ISODate)(&value)
}

func (d *DatePeriodDetails1) SetToDate(value string) {
	d.ToDate = (*ISODate)(&value)
}
