package iso20022

// Description of the element that creates the mismatch.
type ElementIdentification1 struct {

	// Refers to the identification number of  the document in which the mismatch was found
	DocumentIndex *Max3NumericText `xml:"DocIndx"`

	// Specifies from the root of the message the complete path of the element that violated a rule.
	ElementPath *Max350Text `xml:"ElmtPth"`

	// Name of the element.
	ElementName *Max35Text `xml:"ElmtNm"`

	// Content of the element
	ElementValue *Max140Text `xml:"ElmtVal,omitempty"`
}

func (e *ElementIdentification1) SetDocumentIndex(value string) {
	e.DocumentIndex = (*Max3NumericText)(&value)
}

func (e *ElementIdentification1) SetElementPath(value string) {
	e.ElementPath = (*Max350Text)(&value)
}

func (e *ElementIdentification1) SetElementName(value string) {
	e.ElementName = (*Max35Text)(&value)
}

func (e *ElementIdentification1) SetElementValue(value string) {
	e.ElementValue = (*Max140Text)(&value)
}

// Description of the elements that violated a rule.
type ElementIdentification3 struct {

	// Specifies from the root of the message the complete path of the element that violated a rule.
	ElementPath *Max350Text `xml:"ElmtPth"`

	// Name of the element.
	ElementName *Max35Text `xml:"ElmtNm"`

	// Contents of the element.
	ElementValue *Max140Text `xml:"ElmtVal,omitempty"`
}

func (e *ElementIdentification3) SetElementPath(value string) {
	e.ElementPath = (*Max350Text)(&value)
}

func (e *ElementIdentification3) SetElementName(value string) {
	e.ElementName = (*Max35Text)(&value)
}

func (e *ElementIdentification3) SetElementValue(value string) {
	e.ElementValue = (*Max140Text)(&value)
}
