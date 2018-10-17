package iso20022

// Identification of a document.
type DocumentNumber1 struct {

	// Number used to identify a message or document.
	Number *DocumentNumber1Choice `xml:"Nb"`
}

func (d *DocumentNumber1) AddNumber() *DocumentNumber1Choice {
	d.Number = new(DocumentNumber1Choice)
	return d.Number
}

// Identification of the status being requested.
type DocumentNumber12 struct {

	// Number used to identify a message or document.
	Number *DocumentNumber5Choice `xml:"Nb"`

	// References of transaction for which the status is requested.
	References []*Identification15 `xml:"Refs"`
}

func (d *DocumentNumber12) AddNumber() *DocumentNumber5Choice {
	d.Number = new(DocumentNumber5Choice)
	return d.Number
}

func (d *DocumentNumber12) AddReferences() *Identification15 {
	newValue := new(Identification15)
	d.References = append(d.References, newValue)
	return newValue
}

// Identification of a document.
type DocumentNumber13 struct {

	// Number used to identify a message or document.
	Number *DocumentNumber5Choice `xml:"Nb"`
}

func (d *DocumentNumber13) AddNumber() *DocumentNumber5Choice {
	d.Number = new(DocumentNumber5Choice)
	return d.Number
}

// Identification of a document.
type DocumentNumber14 struct {

	// Number used to identify a message or document.
	Number *DocumentNumber6Choice `xml:"Nb"`
}

func (d *DocumentNumber14) AddNumber() *DocumentNumber6Choice {
	d.Number = new(DocumentNumber6Choice)
	return d.Number
}

// Identification of the status being requested.
type DocumentNumber15 struct {

	// Number used to identify a message or document.
	Number *DocumentNumber6Choice `xml:"Nb"`

	// References of transaction for which the status is requested.
	References []*Identification24 `xml:"Refs"`
}

func (d *DocumentNumber15) AddNumber() *DocumentNumber6Choice {
	d.Number = new(DocumentNumber6Choice)
	return d.Number
}

func (d *DocumentNumber15) AddReferences() *Identification24 {
	newValue := new(Identification24)
	d.References = append(d.References, newValue)
	return newValue
}

// Identification of the status being requested.
type DocumentNumber2 struct {

	// Number used to identify a message or document.
	Number *DocumentNumber1Choice `xml:"Nb"`

	// References of transaction for which the status is requested.
	References []*Identification2 `xml:"Refs"`
}

func (d *DocumentNumber2) AddNumber() *DocumentNumber1Choice {
	d.Number = new(DocumentNumber1Choice)
	return d.Number
}

func (d *DocumentNumber2) AddReferences() *Identification2 {
	newValue := new(Identification2)
	d.References = append(d.References, newValue)
	return newValue
}

// Identification of the status being requested.
type DocumentNumber5 struct {

	// Number used to identify a message or document.
	Number *DocumentNumber1Choice `xml:"Nb"`

	// References of transaction for which the status is requested.
	References []*Identification6 `xml:"Refs"`
}

func (d *DocumentNumber5) AddNumber() *DocumentNumber1Choice {
	d.Number = new(DocumentNumber1Choice)
	return d.Number
}

func (d *DocumentNumber5) AddReferences() *Identification6 {
	newValue := new(Identification6)
	d.References = append(d.References, newValue)
	return newValue
}

// Identification of the status being requested.
type DocumentNumber6 struct {

	// Number used to identify a message or document.
	Number *DocumentNumber1Choice `xml:"Nb"`

	// References of transaction for which the status is requested.
	References []*Identification7 `xml:"Refs"`
}

func (d *DocumentNumber6) AddNumber() *DocumentNumber1Choice {
	d.Number = new(DocumentNumber1Choice)
	return d.Number
}

func (d *DocumentNumber6) AddReferences() *Identification7 {
	newValue := new(Identification7)
	d.References = append(d.References, newValue)
	return newValue
}

// Identification of the status being requested.
type DocumentNumber9 struct {

	// Number used to identify a message or document.
	Number *DocumentNumber1Choice `xml:"Nb"`

	// References of transaction for which the status is requested.
	References []*Identification11 `xml:"Refs"`
}

func (d *DocumentNumber9) AddNumber() *DocumentNumber1Choice {
	d.Number = new(DocumentNumber1Choice)
	return d.Number
}

func (d *DocumentNumber9) AddReferences() *Identification11 {
	newValue := new(Identification11)
	d.References = append(d.References, newValue)
	return newValue
}
