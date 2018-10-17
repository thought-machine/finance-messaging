package iso20022

// Specifies the payment terms by means of a code and a period.
type PaymentPeriod1 struct {

	// Code for the payment.
	Code *PaymentTime1Code `xml:"Cd"`

	// Number of days after which the payment must be effected.
	NumberOfDays *Number `xml:"NbOfDays,omitempty"`
}

func (p *PaymentPeriod1) SetCode(value string) {
	p.Code = (*PaymentTime1Code)(&value)
}

func (p *PaymentPeriod1) SetNumberOfDays(value string) {
	p.NumberOfDays = (*Number)(&value)
}

// Specifies the payment terms by means of a code and a period.
type PaymentPeriod2 struct {

	// Code for the payment.
	Code *PaymentTime2Code `xml:"Cd"`

	// Number of days after which the payment must be effected.
	NumberOfDays *Number `xml:"NbOfDays,omitempty"`
}

func (p *PaymentPeriod2) SetCode(value string) {
	p.Code = (*PaymentTime2Code)(&value)
}

func (p *PaymentPeriod2) SetNumberOfDays(value string) {
	p.NumberOfDays = (*Number)(&value)
}

// Specifies the payment terms by means of a code and a period.
type PaymentPeriod3 struct {

	// Code for the payment.
	Code *PaymentTime3Code `xml:"Cd"`

	// Number of days after which the payment must be effected.
	NumberOfDays *Number `xml:"NbOfDays,omitempty"`
}

func (p *PaymentPeriod3) SetCode(value string) {
	p.Code = (*PaymentTime3Code)(&value)
}

func (p *PaymentPeriod3) SetNumberOfDays(value string) {
	p.NumberOfDays = (*Number)(&value)
}

// Specifies the payment terms by means of a code and a period.
type PaymentPeriod4 struct {

	// Code for the payment.
	Code *PaymentTime4Code `xml:"Cd"`

	// Number of days after which the payment must be effected.
	NumberOfDays *Number `xml:"NbOfDays,omitempty"`
}

func (p *PaymentPeriod4) SetCode(value string) {
	p.Code = (*PaymentTime4Code)(&value)
}

func (p *PaymentPeriod4) SetNumberOfDays(value string) {
	p.NumberOfDays = (*Number)(&value)
}
