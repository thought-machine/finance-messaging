package iso20022

// Message to be displayed to the cardholder or the cashier.
type ActionMessage1 struct {

	// Destination of the message to be displayed or printed.
	MessageDestination *UserInterface1Code `xml:"MsgDstn"`

	// Text or graphic data to be display or printed to the cardholder or the cashier.
	MessageContent *Max256Text `xml:"MsgCntt"`

	// Electronic signature of the message to display or print.
	MessageContentSignature *Max70Text `xml:"MsgCnttSgntr,omitempty"`
}

func (a *ActionMessage1) SetMessageDestination(value string) {
	a.MessageDestination = (*UserInterface1Code)(&value)
}

func (a *ActionMessage1) SetMessageContent(value string) {
	a.MessageContent = (*Max256Text)(&value)
}

func (a *ActionMessage1) SetMessageContentSignature(value string) {
	a.MessageContentSignature = (*Max70Text)(&value)
}

// Information to display, print or store.
type ActionMessage2 struct {

	// Destination of the message.
	MessageDestination *UserInterface4Code `xml:"MsgDstn"`

	// Message format.
	Format *OutputFormat1Code `xml:"Frmt,omitempty"`

	// Content or reference of the message.
	MessageContent *Max20000Text `xml:"MsgCntt"`

	// Digital signature of the message.
	MessageContentSignature *Max140Binary `xml:"MsgCnttSgntr,omitempty"`
}

func (a *ActionMessage2) SetMessageDestination(value string) {
	a.MessageDestination = (*UserInterface4Code)(&value)
}

func (a *ActionMessage2) SetFormat(value string) {
	a.Format = (*OutputFormat1Code)(&value)
}

func (a *ActionMessage2) SetMessageContent(value string) {
	a.MessageContent = (*Max20000Text)(&value)
}

func (a *ActionMessage2) SetMessageContentSignature(value string) {
	a.MessageContentSignature = (*Max140Binary)(&value)
}

// Information to log.
type ActionMessage3 struct {

	// Destination of the information.
	Destination *UserInterface3Code `xml:"Dstn"`

	// Format of the content.
	Format *OutputFormat1Code `xml:"Frmt,omitempty"`

	// Content of the information.
	Content *Max20000Text `xml:"Cntt"`
}

func (a *ActionMessage3) SetDestination(value string) {
	a.Destination = (*UserInterface3Code)(&value)
}

func (a *ActionMessage3) SetFormat(value string) {
	a.Format = (*OutputFormat1Code)(&value)
}

func (a *ActionMessage3) SetContent(value string) {
	a.Content = (*Max20000Text)(&value)
}

// Information to display, print or log.
type ActionMessage4 struct {

	// Information format.
	Format *OutputFormat2Code `xml:"Frmt,omitempty"`

	// Content of the message.
	Message *Max20000Text `xml:"Msg,omitempty"`

	// Message content if this is a message reference or screen reference.
	Reference *Max35Text `xml:"Ref,omitempty"`

	// Device to be used.
	Device *ATMDevice1Code `xml:"Dvc,omitempty"`

	// Electronic signature of the message to display or print.
	MessageContentSignature *Max35Binary `xml:"MsgCnttSgntr,omitempty"`
}

func (a *ActionMessage4) SetFormat(value string) {
	a.Format = (*OutputFormat2Code)(&value)
}

func (a *ActionMessage4) SetMessage(value string) {
	a.Message = (*Max20000Text)(&value)
}

func (a *ActionMessage4) SetReference(value string) {
	a.Reference = (*Max35Text)(&value)
}

func (a *ActionMessage4) SetDevice(value string) {
	a.Device = (*ATMDevice1Code)(&value)
}

func (a *ActionMessage4) SetMessageContentSignature(value string) {
	a.MessageContentSignature = (*Max35Binary)(&value)
}

// Message to be displayed to the cardholder or the cashier.
type ActionMessage5 struct {

	// Format of the content.
	Format *OutputFormat1Code `xml:"Frmt,omitempty"`

	// Text or graphic data to be display or printed to the cardholder or the cashier.
	MessageContent *Max20000Text `xml:"MsgCntt"`
}

func (a *ActionMessage5) SetFormat(value string) {
	a.Format = (*OutputFormat1Code)(&value)
}

func (a *ActionMessage5) SetMessageContent(value string) {
	a.MessageContent = (*Max20000Text)(&value)
}
