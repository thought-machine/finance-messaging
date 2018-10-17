package iso20022

// Data to authenticate.
type EncapsulatedContent1 struct {

	// Type of data which have been authenticated.
	ContentType *ContentType1Code `xml:"CnttTp"`

	// Actual data to authenticate.
	Content *Max10000Binary `xml:"Cntt,omitempty"`
}

func (e *EncapsulatedContent1) SetContentType(value string) {
	e.ContentType = (*ContentType1Code)(&value)
}

func (e *EncapsulatedContent1) SetContent(value string) {
	e.Content = (*Max10000Binary)(&value)
}

// Data to authenticate.
type EncapsulatedContent2 struct {

	// Type of data which have been authenticated.
	ContentType *ContentType1Code `xml:"CnttTp"`

	// Actual data to authenticate.
	Content *Max100KBinary `xml:"Cntt,omitempty"`
}

func (e *EncapsulatedContent2) SetContentType(value string) {
	e.ContentType = (*ContentType1Code)(&value)
}

func (e *EncapsulatedContent2) SetContent(value string) {
	e.Content = (*Max100KBinary)(&value)
}

// Data to authenticate.
type EncapsulatedContent3 struct {

	// Type of data which have been authenticated.
	ContentType *ContentType2Code `xml:"CnttTp"`

	// Actual data to authenticate.
	Content *Max100KBinary `xml:"Cntt,omitempty"`
}

func (e *EncapsulatedContent3) SetContentType(value string) {
	e.ContentType = (*ContentType2Code)(&value)
}

func (e *EncapsulatedContent3) SetContent(value string) {
	e.Content = (*Max100KBinary)(&value)
}
