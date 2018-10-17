package iso20022

// Specifies  a date code and a time.
type DateCodeAndTimeFormat1 struct {

	// Specifies the type of date.
	DateCode *DateCode4Choice `xml:"DtCd"`

	// Specifies the time.
	Time *ISOTime `xml:"Tm"`
}

func (d *DateCodeAndTimeFormat1) AddDateCode() *DateCode4Choice {
	d.DateCode = new(DateCode4Choice)
	return d.DateCode
}

func (d *DateCodeAndTimeFormat1) SetTime(value string) {
	d.Time = (*ISOTime)(&value)
}

// Specifies  a date code and a time.
type DateCodeAndTimeFormat3 struct {

	// Specifies the type of date.
	DateCode *DateCode21Choice `xml:"DtCd"`

	// Specifies the time.
	Time *ISOTime `xml:"Tm"`
}

func (d *DateCodeAndTimeFormat3) AddDateCode() *DateCode21Choice {
	d.DateCode = new(DateCode21Choice)
	return d.DateCode
}

func (d *DateCodeAndTimeFormat3) SetTime(value string) {
	d.Time = (*ISOTime)(&value)
}

// Specifies  a date code and a time.
type DateCodeAndTimeFormat4 struct {

	// Specifies the type of date.
	DateCode *DateCode26Choice `xml:"DtCd"`

	// Specifies the time.
	Time *ISOTime `xml:"Tm"`
}

func (d *DateCodeAndTimeFormat4) AddDateCode() *DateCode26Choice {
	d.DateCode = new(DateCode26Choice)
	return d.DateCode
}

func (d *DateCodeAndTimeFormat4) SetTime(value string) {
	d.Time = (*ISOTime)(&value)
}
