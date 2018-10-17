package iso20022

// Capabilities of the POI performing the transaction.
type PointOfInteractionCapabilities1 struct {

	// Card reading capabilities of the POI performing the transaction.
	CardReadingCapabilities []*CardDataReading1Code `xml:"CardRdngCpblties,omitempty"`

	// Cardholder verification capabilities of the POI performing the transaction.
	CardholderVerificationCapabilities []*CardholderVerificationCapability1Code `xml:"CrdhldrVrfctnCpblties,omitempty"`

	// On-line and off-line capabilities of the POI.
	OnLineCapabilities *OnLineCapability1Code `xml:"OnLineCpblties,omitempty"`

	// Capabilities of the display components performing the transaction.
	DisplayCapabilities []*DisplayCapabilities1 `xml:"DispCpblties,omitempty"`

	// Number of columns of the printer component.
	PrintLineWidth *Max3NumericText `xml:"PrtLineWidth,omitempty"`
}

func (p *PointOfInteractionCapabilities1) AddCardReadingCapabilities(value string) {
	p.CardReadingCapabilities = append(p.CardReadingCapabilities, (*CardDataReading1Code)(&value))
}

func (p *PointOfInteractionCapabilities1) AddCardholderVerificationCapabilities(value string) {
	p.CardholderVerificationCapabilities = append(p.CardholderVerificationCapabilities, (*CardholderVerificationCapability1Code)(&value))
}

func (p *PointOfInteractionCapabilities1) SetOnLineCapabilities(value string) {
	p.OnLineCapabilities = (*OnLineCapability1Code)(&value)
}

func (p *PointOfInteractionCapabilities1) AddDisplayCapabilities() *DisplayCapabilities1 {
	newValue := new(DisplayCapabilities1)
	p.DisplayCapabilities = append(p.DisplayCapabilities, newValue)
	return newValue
}

func (p *PointOfInteractionCapabilities1) SetPrintLineWidth(value string) {
	p.PrintLineWidth = (*Max3NumericText)(&value)
}

// Capabilities of the POI (Point Of Interaction) performing the transaction.
type PointOfInteractionCapabilities2 struct {

	// Card reading capabilities of the POI (Point Of Interaction) performing the transaction.
	CardReadingCapabilities []*CardDataReading1Code `xml:"CardRdngCpblties,omitempty"`

	// Cardholder verification capabilities of the POI (Point Of Interaction) performing the transaction.
	CardholderVerificationCapabilities []*CardholderVerificationCapability1Code `xml:"CrdhldrVrfctnCpblties,omitempty"`

	// On-line and off-line capabilities of the POI (Point Of Interaction).
	OnLineCapabilities *OnLineCapability1Code `xml:"OnLineCpblties,omitempty"`

	// Capabilities of the display components performing the transaction.
	DisplayCapabilities []*DisplayCapabilities2 `xml:"DispCpblties,omitempty"`

	// Number of columns of the printer component.
	PrintLineWidth *Number `xml:"PrtLineWidth,omitempty"`

	// Available language in the display and printer interface.
	AvailableLanguage []*ISO2ALanguageCode `xml:"AvlblLang,omitempty"`
}

func (p *PointOfInteractionCapabilities2) AddCardReadingCapabilities(value string) {
	p.CardReadingCapabilities = append(p.CardReadingCapabilities, (*CardDataReading1Code)(&value))
}

func (p *PointOfInteractionCapabilities2) AddCardholderVerificationCapabilities(value string) {
	p.CardholderVerificationCapabilities = append(p.CardholderVerificationCapabilities, (*CardholderVerificationCapability1Code)(&value))
}

func (p *PointOfInteractionCapabilities2) SetOnLineCapabilities(value string) {
	p.OnLineCapabilities = (*OnLineCapability1Code)(&value)
}

func (p *PointOfInteractionCapabilities2) AddDisplayCapabilities() *DisplayCapabilities2 {
	newValue := new(DisplayCapabilities2)
	p.DisplayCapabilities = append(p.DisplayCapabilities, newValue)
	return newValue
}

func (p *PointOfInteractionCapabilities2) SetPrintLineWidth(value string) {
	p.PrintLineWidth = (*Number)(&value)
}

func (p *PointOfInteractionCapabilities2) AddAvailableLanguage(value string) {
	p.AvailableLanguage = append(p.AvailableLanguage, (*ISO2ALanguageCode)(&value))
}

// Capabilities of the POI (Point Of Interaction) performing the transaction.
type PointOfInteractionCapabilities3 struct {

	// Card reading capabilities of the POI (Point Of Interaction) performing the transaction.
	CardReadingCapabilities []*CardDataReading1Code `xml:"CardRdngCpblties,omitempty"`

	// Cardholder verification capabilities of the POI (Point Of Interaction) performing the transaction.
	CardholderVerificationCapabilities []*CardholderVerificationCapability1Code `xml:"CrdhldrVrfctnCpblties,omitempty"`

	// Maximum number of digits the POI is able to accept when the cardholder enters its PIN.
	PINLengthCapabilities *Number `xml:"PINLngthCpblties,omitempty"`

	// Maximum number of characters of the approval code the POI is able to manage.
	ApprovalCodeLength *Number `xml:"ApprvlCdLngth,omitempty"`

	// True if the POI is able to capture card.
	CardCaptureCapable *TrueFalseIndicator `xml:"CardCaptrCpbl,omitempty"`

	// On-line and off-line capabilities of the POI (Point Of Interaction).
	OnLineCapabilities *OnLineCapability1Code `xml:"OnLineCpblties,omitempty"`

	// Capabilities of the display components performing the transaction.
	DisplayCapabilities []*DisplayCapabilities2 `xml:"DispCpblties,omitempty"`

	// Number of columns of the printer component.
	PrintLineWidth *Number `xml:"PrtLineWidth,omitempty"`

	// Available language in the display and printer interface.
	// Reference ISO 639-1 (alpha-2) et ISO 639-2 (alpha-3).
	AvailableLanguage []*LanguageCode `xml:"AvlblLang,omitempty"`
}

func (p *PointOfInteractionCapabilities3) AddCardReadingCapabilities(value string) {
	p.CardReadingCapabilities = append(p.CardReadingCapabilities, (*CardDataReading1Code)(&value))
}

func (p *PointOfInteractionCapabilities3) AddCardholderVerificationCapabilities(value string) {
	p.CardholderVerificationCapabilities = append(p.CardholderVerificationCapabilities, (*CardholderVerificationCapability1Code)(&value))
}

func (p *PointOfInteractionCapabilities3) SetPINLengthCapabilities(value string) {
	p.PINLengthCapabilities = (*Number)(&value)
}

func (p *PointOfInteractionCapabilities3) SetApprovalCodeLength(value string) {
	p.ApprovalCodeLength = (*Number)(&value)
}

func (p *PointOfInteractionCapabilities3) SetCardCaptureCapable(value string) {
	p.CardCaptureCapable = (*TrueFalseIndicator)(&value)
}

func (p *PointOfInteractionCapabilities3) SetOnLineCapabilities(value string) {
	p.OnLineCapabilities = (*OnLineCapability1Code)(&value)
}

func (p *PointOfInteractionCapabilities3) AddDisplayCapabilities() *DisplayCapabilities2 {
	newValue := new(DisplayCapabilities2)
	p.DisplayCapabilities = append(p.DisplayCapabilities, newValue)
	return newValue
}

func (p *PointOfInteractionCapabilities3) SetPrintLineWidth(value string) {
	p.PrintLineWidth = (*Number)(&value)
}

func (p *PointOfInteractionCapabilities3) AddAvailableLanguage(value string) {
	p.AvailableLanguage = append(p.AvailableLanguage, (*LanguageCode)(&value))
}

// Capabilities of the terminal performing the transaction.
type PointOfInteractionCapabilities4 struct {

	// Card reading capabilities of the terminal performing the transaction.
	// It correspond to the ISO 8583 field number 22-2 for the version 93, and field number 27-1 for the version 2003.
	CardReadingCapabilities []*CardDataReading2Code `xml:"CardRdngCpblties"`

	// Card writting capabilities of the terminal performing the transaction.
	// It correspond to the ISO 8583 field number 22-10 for the version 93, and field number 27-8_9 for the version 2003.
	CardWrittingCapabilities []*CardDataReading3Code `xml:"CardWrttgCpblties,omitempty"`

	// Cardholder verification capabilities by the terminal.
	// It correspond to the ISO 8583 field number 22-2 for the versions 87 and 93, and field number 27-2 for the version 2003.
	CardholderVerificationCapabilities []*CardholderVerificationCapability2Code `xml:"CrdhldrVrfctnCpblties,omitempty"`

	// Maximum number of digits the POI is able to accept when the cardholder enters its PIN.
	// It correspond to the ISO 8583, field number 25 for the version 87, 22-12 for the version 93, and field number 27-11 for the version 2003.
	PINLengthCapabilities *Number `xml:"PINLngthCpblties,omitempty"`

	// Maximum number of characters of the approval code the POI is able to manage.
	ApprovalCodeLength *Number `xml:"ApprvlCdLngth,omitempty"`

	// Maximum data length in bytes that a card issuer can return to the ICC at the terminal.
	MaxScriptLength *Number `xml:"MxScrptLngth,omitempty"`

	// True if the POI is able to capture card.
	CardCaptureCapable *TrueFalseIndicator `xml:"CardCaptrCpbl,omitempty"`

	// On-line and off-line capabilities of the POI (Point Of Interaction).
	OnLineCapabilities *OnLineCapability1Code `xml:"OnLineCpblties,omitempty"`

	// Capabilities of the terminal to display or print message to the cardholder and the merchant.
	// It correspond to the ISO 8583 field number 22-11 for the version 93, and field number 27-6 for the version 2003.
	MessageCapabilities []*DisplayCapabilities3 `xml:"MsgCpblties,omitempty"`
}

func (p *PointOfInteractionCapabilities4) AddCardReadingCapabilities(value string) {
	p.CardReadingCapabilities = append(p.CardReadingCapabilities, (*CardDataReading2Code)(&value))
}

func (p *PointOfInteractionCapabilities4) AddCardWrittingCapabilities(value string) {
	p.CardWrittingCapabilities = append(p.CardWrittingCapabilities, (*CardDataReading3Code)(&value))
}

func (p *PointOfInteractionCapabilities4) AddCardholderVerificationCapabilities(value string) {
	p.CardholderVerificationCapabilities = append(p.CardholderVerificationCapabilities, (*CardholderVerificationCapability2Code)(&value))
}

func (p *PointOfInteractionCapabilities4) SetPINLengthCapabilities(value string) {
	p.PINLengthCapabilities = (*Number)(&value)
}

func (p *PointOfInteractionCapabilities4) SetApprovalCodeLength(value string) {
	p.ApprovalCodeLength = (*Number)(&value)
}

func (p *PointOfInteractionCapabilities4) SetMaxScriptLength(value string) {
	p.MaxScriptLength = (*Number)(&value)
}

func (p *PointOfInteractionCapabilities4) SetCardCaptureCapable(value string) {
	p.CardCaptureCapable = (*TrueFalseIndicator)(&value)
}

func (p *PointOfInteractionCapabilities4) SetOnLineCapabilities(value string) {
	p.OnLineCapabilities = (*OnLineCapability1Code)(&value)
}

func (p *PointOfInteractionCapabilities4) AddMessageCapabilities() *DisplayCapabilities3 {
	newValue := new(DisplayCapabilities3)
	p.MessageCapabilities = append(p.MessageCapabilities, newValue)
	return newValue
}

// Capabilities of the ATM terminal.
type PointOfInteractionCapabilities5 struct {

	// Card reading capabilities of the ATM performing the transaction.
	CardReadData []*CardDataReading4Code `xml:"CardRdData,omitempty"`

	// Card writing capabilities of the terminal performing the transaction.
	CardWriteData []*CardDataReading4Code `xml:"CardWrtData,omitempty"`

	// Customer and card authentication capabilities available at the ATM.
	Authentication []*CardholderVerificationCapability3Code `xml:"Authntcn,omitempty"`

	// Maximum number of digits the ATM is able to accept when the cardholder enters its PIN.
	PINLengthCapabilities *Number `xml:"PINLngthCpblties,omitempty"`

	// Maximum number of characters of the approval code the ATM is able to manage.
	ApprovalCodeLength *Number `xml:"ApprvlCdLngth,omitempty"`

	// Maximum data length in bytes that a card issuer can return to the ICC at the terminal.
	MaxScriptLength *Number `xml:"MxScrptLngth,omitempty"`

	// True if the ATM is able to capture card.
	CardCaptureCapable *TrueFalseIndicator `xml:"CardCaptrCpbl,omitempty"`
}

func (p *PointOfInteractionCapabilities5) AddCardReadData(value string) {
	p.CardReadData = append(p.CardReadData, (*CardDataReading4Code)(&value))
}

func (p *PointOfInteractionCapabilities5) AddCardWriteData(value string) {
	p.CardWriteData = append(p.CardWriteData, (*CardDataReading4Code)(&value))
}

func (p *PointOfInteractionCapabilities5) AddAuthentication(value string) {
	p.Authentication = append(p.Authentication, (*CardholderVerificationCapability3Code)(&value))
}

func (p *PointOfInteractionCapabilities5) SetPINLengthCapabilities(value string) {
	p.PINLengthCapabilities = (*Number)(&value)
}

func (p *PointOfInteractionCapabilities5) SetApprovalCodeLength(value string) {
	p.ApprovalCodeLength = (*Number)(&value)
}

func (p *PointOfInteractionCapabilities5) SetMaxScriptLength(value string) {
	p.MaxScriptLength = (*Number)(&value)
}

func (p *PointOfInteractionCapabilities5) SetCardCaptureCapable(value string) {
	p.CardCaptureCapable = (*TrueFalseIndicator)(&value)
}

// Capabilities of the POI (Point Of Interaction) performing the transaction.
type PointOfInteractionCapabilities6 struct {

	// Card reading capabilities of the POI (Point Of Interaction) performing the transaction.
	CardReadingCapabilities []*CardDataReading5Code `xml:"CardRdngCpblties,omitempty"`

	// Cardholder verification capabilities of the POI (Point Of Interaction) performing the transaction.
	CardholderVerificationCapabilities []*CardholderVerificationCapability4Code `xml:"CrdhldrVrfctnCpblties,omitempty"`

	// Maximum number of digits the POI is able to accept when the cardholder enters its PIN.
	PINLengthCapabilities *Number `xml:"PINLngthCpblties,omitempty"`

	// Maximum number of characters of the approval code the POI is able to manage.
	ApprovalCodeLength *Number `xml:"ApprvlCdLngth,omitempty"`

	// Maximum data length in bytes that a card issuer can return to the ICC at the terminal.
	MaxScriptLength *Number `xml:"MxScrptLngth,omitempty"`

	// True if the POI is able to capture card.
	CardCaptureCapable *TrueFalseIndicator `xml:"CardCaptrCpbl,omitempty"`

	// On-line and off-line capabilities of the POI (Point Of Interaction).
	OnLineCapabilities *OnLineCapability1Code `xml:"OnLineCpblties,omitempty"`

	// Capabilities of the terminal to display or print message to the cardholder and the merchant.
	MessageCapabilities []*DisplayCapabilities4 `xml:"MsgCpblties,omitempty"`
}

func (p *PointOfInteractionCapabilities6) AddCardReadingCapabilities(value string) {
	p.CardReadingCapabilities = append(p.CardReadingCapabilities, (*CardDataReading5Code)(&value))
}

func (p *PointOfInteractionCapabilities6) AddCardholderVerificationCapabilities(value string) {
	p.CardholderVerificationCapabilities = append(p.CardholderVerificationCapabilities, (*CardholderVerificationCapability4Code)(&value))
}

func (p *PointOfInteractionCapabilities6) SetPINLengthCapabilities(value string) {
	p.PINLengthCapabilities = (*Number)(&value)
}

func (p *PointOfInteractionCapabilities6) SetApprovalCodeLength(value string) {
	p.ApprovalCodeLength = (*Number)(&value)
}

func (p *PointOfInteractionCapabilities6) SetMaxScriptLength(value string) {
	p.MaxScriptLength = (*Number)(&value)
}

func (p *PointOfInteractionCapabilities6) SetCardCaptureCapable(value string) {
	p.CardCaptureCapable = (*TrueFalseIndicator)(&value)
}

func (p *PointOfInteractionCapabilities6) SetOnLineCapabilities(value string) {
	p.OnLineCapabilities = (*OnLineCapability1Code)(&value)
}

func (p *PointOfInteractionCapabilities6) AddMessageCapabilities() *DisplayCapabilities4 {
	newValue := new(DisplayCapabilities4)
	p.MessageCapabilities = append(p.MessageCapabilities, newValue)
	return newValue
}

// Capabilities of the ATM terminal.
type PointOfInteractionCapabilities7 struct {

	// Card reading capabilities of the ATM performing the transaction.
	CardReadData []*CardDataReading4Code `xml:"CardRdData,omitempty"`

	// Card writing capabilities of the terminal performing the transaction.
	CardWriteData []*CardDataReading4Code `xml:"CardWrtData,omitempty"`

	// Customer and card authentication capabilities available at the ATM.
	Authentication []*CardholderVerificationCapability3Code `xml:"Authntcn,omitempty"`

	// Maximum number of digits the ATM is able to accept when the cardholder enters its PIN.
	PINLengthCapabilities *Number `xml:"PINLngthCpblties,omitempty"`

	// Maximum number of characters of the approval code the ATM is able to manage.
	ApprovalCodeLength *Number `xml:"ApprvlCdLngth,omitempty"`

	// Maximum data length in bytes that a card issuer can return to the ICC at the terminal.
	MaxScriptLength *Number `xml:"MxScrptLngth,omitempty"`

	// True if the ATM is able to capture card.
	CardCaptureCapable *TrueFalseIndicator `xml:"CardCaptrCpbl,omitempty"`

	// Type of media the ATM is able to dispense.
	WithdrawalMedia []*ATMMediaType1Code `xml:"WdrwlMdia,omitempty"`

	// Type of media the customer is able to deposit in the ATM.
	DepositedMedia []*ATMMediaType2Code `xml:"DpstdMdia,omitempty"`

	// Capabilities of the terminal to display or print message to the cardholder and the merchant.
	MessageCapabilities []*DisplayCapabilities5 `xml:"MsgCpblties,omitempty"`
}

func (p *PointOfInteractionCapabilities7) AddCardReadData(value string) {
	p.CardReadData = append(p.CardReadData, (*CardDataReading4Code)(&value))
}

func (p *PointOfInteractionCapabilities7) AddCardWriteData(value string) {
	p.CardWriteData = append(p.CardWriteData, (*CardDataReading4Code)(&value))
}

func (p *PointOfInteractionCapabilities7) AddAuthentication(value string) {
	p.Authentication = append(p.Authentication, (*CardholderVerificationCapability3Code)(&value))
}

func (p *PointOfInteractionCapabilities7) SetPINLengthCapabilities(value string) {
	p.PINLengthCapabilities = (*Number)(&value)
}

func (p *PointOfInteractionCapabilities7) SetApprovalCodeLength(value string) {
	p.ApprovalCodeLength = (*Number)(&value)
}

func (p *PointOfInteractionCapabilities7) SetMaxScriptLength(value string) {
	p.MaxScriptLength = (*Number)(&value)
}

func (p *PointOfInteractionCapabilities7) SetCardCaptureCapable(value string) {
	p.CardCaptureCapable = (*TrueFalseIndicator)(&value)
}

func (p *PointOfInteractionCapabilities7) AddWithdrawalMedia(value string) {
	p.WithdrawalMedia = append(p.WithdrawalMedia, (*ATMMediaType1Code)(&value))
}

func (p *PointOfInteractionCapabilities7) AddDepositedMedia(value string) {
	p.DepositedMedia = append(p.DepositedMedia, (*ATMMediaType2Code)(&value))
}

func (p *PointOfInteractionCapabilities7) AddMessageCapabilities() *DisplayCapabilities5 {
	newValue := new(DisplayCapabilities5)
	p.MessageCapabilities = append(p.MessageCapabilities, newValue)
	return newValue
}
