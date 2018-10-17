package iso20022

// The capabilities of the display components performing the transaction.
type DisplayCapabilities1 struct {

	// Type of display (for example merchant or cardholder).
	DisplayType *UserInterface2Code `xml:"DispTp"`

	// Number of lines of the display component.
	NumberOfLines *Max3NumericText `xml:"NbOfLines"`

	// Number of columns of the display component.
	LineWidth *Max3NumericText `xml:"LineWidth"`
}

func (d *DisplayCapabilities1) SetDisplayType(value string) {
	d.DisplayType = (*UserInterface2Code)(&value)
}

func (d *DisplayCapabilities1) SetNumberOfLines(value string) {
	d.NumberOfLines = (*Max3NumericText)(&value)
}

func (d *DisplayCapabilities1) SetLineWidth(value string) {
	d.LineWidth = (*Max3NumericText)(&value)
}

// The capabilities of the display components performing the transaction.
type DisplayCapabilities2 struct {

	// Type of display (for example merchant or cardholder).
	DisplayType *UserInterface2Code `xml:"DispTp"`

	// Number of lines of the display component.
	NumberOfLines *Number `xml:"NbOfLines"`

	// Number of columns of the display component.
	LineWidth *Number `xml:"LineWidth"`
}

func (d *DisplayCapabilities2) SetDisplayType(value string) {
	d.DisplayType = (*UserInterface2Code)(&value)
}

func (d *DisplayCapabilities2) SetNumberOfLines(value string) {
	d.NumberOfLines = (*Number)(&value)
}

func (d *DisplayCapabilities2) SetLineWidth(value string) {
	d.LineWidth = (*Number)(&value)
}

// Capabilities of the display components performing the transaction.
type DisplayCapabilities3 struct {

	// Destination of the message to present.
	Destination *UserInterface1Code `xml:"Dstn"`

	// Available message format.
	AvailableFormat []*OutputFormat1Code `xml:"AvlblFrmt"`

	// Number of lines of the display.
	NumberOfLines *Number `xml:"NbOfLines,omitempty"`

	// Number of columns of the display or printer.
	LineWidth *Number `xml:"LineWidth,omitempty"`

	// Available language for the message. Reference ISO 639-1 (alpha-2) et ISO 639-2 (alpha-3).
	AvailableLanguage []*LanguageCode `xml:"AvlblLang,omitempty"`
}

func (d *DisplayCapabilities3) SetDestination(value string) {
	d.Destination = (*UserInterface1Code)(&value)
}

func (d *DisplayCapabilities3) AddAvailableFormat(value string) {
	d.AvailableFormat = append(d.AvailableFormat, (*OutputFormat1Code)(&value))
}

func (d *DisplayCapabilities3) SetNumberOfLines(value string) {
	d.NumberOfLines = (*Number)(&value)
}

func (d *DisplayCapabilities3) SetLineWidth(value string) {
	d.LineWidth = (*Number)(&value)
}

func (d *DisplayCapabilities3) AddAvailableLanguage(value string) {
	d.AvailableLanguage = append(d.AvailableLanguage, (*LanguageCode)(&value))
}

// Capabilities of the display components performing the transaction.
type DisplayCapabilities4 struct {

	// Destination of the message to present.
	Destination []*UserInterface4Code `xml:"Dstn"`

	// Available message format.
	AvailableFormat []*OutputFormat1Code `xml:"AvlblFrmt,omitempty"`

	// Number of lines of the display.
	NumberOfLines *Number `xml:"NbOfLines,omitempty"`

	// Number of columns of the display or printer.
	LineWidth *Number `xml:"LineWidth,omitempty"`

	// Available language for the message. Reference ISO 639-1 (alpha-2) et ISO 639-2 (alpha-3).
	AvailableLanguage []*LanguageCode `xml:"AvlblLang,omitempty"`
}

func (d *DisplayCapabilities4) AddDestination(value string) {
	d.Destination = append(d.Destination, (*UserInterface4Code)(&value))
}

func (d *DisplayCapabilities4) AddAvailableFormat(value string) {
	d.AvailableFormat = append(d.AvailableFormat, (*OutputFormat1Code)(&value))
}

func (d *DisplayCapabilities4) SetNumberOfLines(value string) {
	d.NumberOfLines = (*Number)(&value)
}

func (d *DisplayCapabilities4) SetLineWidth(value string) {
	d.LineWidth = (*Number)(&value)
}

func (d *DisplayCapabilities4) AddAvailableLanguage(value string) {
	d.AvailableLanguage = append(d.AvailableLanguage, (*LanguageCode)(&value))
}

// Capabilities of the display components performing the transaction.
type DisplayCapabilities5 struct {

	// Destination of the message to present.
	Destination []*UserInterface5Code `xml:"Dstn"`

	// Available message format.
	AvailableFormat []*OutputFormat1Code `xml:"AvlblFrmt,omitempty"`

	// Number of lines of the display.
	NumberOfLines *Number `xml:"NbOfLines,omitempty"`

	// Number of columns of the display or printer.
	LineWidth *Number `xml:"LineWidth,omitempty"`

	// Available language for the message. Reference ISO 639-1 (alpha-2) et ISO 639-2 (alpha-3).
	AvailableLanguage []*LanguageCode `xml:"AvlblLang,omitempty"`
}

func (d *DisplayCapabilities5) AddDestination(value string) {
	d.Destination = append(d.Destination, (*UserInterface5Code)(&value))
}

func (d *DisplayCapabilities5) AddAvailableFormat(value string) {
	d.AvailableFormat = append(d.AvailableFormat, (*OutputFormat1Code)(&value))
}

func (d *DisplayCapabilities5) SetNumberOfLines(value string) {
	d.NumberOfLines = (*Number)(&value)
}

func (d *DisplayCapabilities5) SetLineWidth(value string) {
	d.LineWidth = (*Number)(&value)
}

func (d *DisplayCapabilities5) AddAvailableLanguage(value string) {
	d.AvailableLanguage = append(d.AvailableLanguage, (*LanguageCode)(&value))
}
