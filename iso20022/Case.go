package iso20022

// Information identifying a case.
type Case struct {

	// Unique id assigned by the case creator.
	Identification *Max35Text `xml:"Id"`

	// Party that created the case.
	Creator *AnyBICIdentifier `xml:"Cretr"`

	// Set to yes if the case was closed and needs to be re-opened.
	ReopenCaseIndication *YesNoIndicator `xml:"ReopCaseIndctn,omitempty"`
}

func (c *Case) SetIdentification(value string) {
	c.Identification = (*Max35Text)(&value)
}

func (c *Case) SetCreator(value string) {
	c.Creator = (*AnyBICIdentifier)(&value)
}

func (c *Case) SetReopenCaseIndication(value string) {
	c.ReopenCaseIndication = (*YesNoIndicator)(&value)
}

// Set of elements used to identify a case.
type Case2 struct {

	// Uniquely identifies the case.
	Identification *Max35Text `xml:"Id"`

	// Party that created the investigation case.
	Creator *Party7Choice `xml:"Cretr"`

	// Indicates whether or not the case was previously closed and is now re-opened.
	ReopenCaseIndication *YesNoIndicator `xml:"ReopCaseIndctn,omitempty"`
}

func (c *Case2) SetIdentification(value string) {
	c.Identification = (*Max35Text)(&value)
}

func (c *Case2) AddCreator() *Party7Choice {
	c.Creator = new(Party7Choice)
	return c.Creator
}

func (c *Case2) SetReopenCaseIndication(value string) {
	c.ReopenCaseIndication = (*YesNoIndicator)(&value)
}

// Provides further details to identify an investigation case.
type Case3 struct {

	// Uniquely identifies the case.
	Identification *Max35Text `xml:"Id"`

	// Party that created the investigation case.
	Creator *Party12Choice `xml:"Cretr"`

	// Indicates whether or not the case was previously closed and is now re-opened.
	ReopenCaseIndication *YesNoIndicator `xml:"ReopCaseIndctn,omitempty"`
}

func (c *Case3) SetIdentification(value string) {
	c.Identification = (*Max35Text)(&value)
}

func (c *Case3) AddCreator() *Party12Choice {
	c.Creator = new(Party12Choice)
	return c.Creator
}

func (c *Case3) SetReopenCaseIndication(value string) {
	c.ReopenCaseIndication = (*YesNoIndicator)(&value)
}
