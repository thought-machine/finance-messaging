package iso20022

// Encrypted personal identification number (PIN) and related information.
type OnLinePIN1 struct {

	// Encrypted PIN (Personal Identification Number).
	EncryptedPINBlock *ContentInformationType2 `xml:"NcrptdPINBlck"`

	// PIN format before encryption.
	PINFormat *PINFormat1Code `xml:"PINFrmt"`

	// Additional information required to verify the PIN.
	AdditionalInput *Max35Text `xml:"AddtlInpt,omitempty"`
}

func (o *OnLinePIN1) AddEncryptedPINBlock() *ContentInformationType2 {
	o.EncryptedPINBlock = new(ContentInformationType2)
	return o.EncryptedPINBlock
}

func (o *OnLinePIN1) SetPINFormat(value string) {
	o.PINFormat = (*PINFormat1Code)(&value)
}

func (o *OnLinePIN1) SetAdditionalInput(value string) {
	o.AdditionalInput = (*Max35Text)(&value)
}

// Encrypted personal identification number (PIN) and related information.
type OnLinePIN2 struct {

	// Encrypted PIN (Personal Identification Number).
	EncryptedPINBlock *ContentInformationType5 `xml:"NcrptdPINBlck"`

	// PIN format before encryption.
	PINFormat *PINFormat2Code `xml:"PINFrmt"`

	// Additional information required to verify the PIN.
	AdditionalInput *Max35Text `xml:"AddtlInpt,omitempty"`
}

func (o *OnLinePIN2) AddEncryptedPINBlock() *ContentInformationType5 {
	o.EncryptedPINBlock = new(ContentInformationType5)
	return o.EncryptedPINBlock
}

func (o *OnLinePIN2) SetPINFormat(value string) {
	o.PINFormat = (*PINFormat2Code)(&value)
}

func (o *OnLinePIN2) SetAdditionalInput(value string) {
	o.AdditionalInput = (*Max35Text)(&value)
}

// Encrypted personal identification number (PIN) and related information.
type OnLinePIN3 struct {

	// Encrypted PIN (Personal Identification Number).
	EncryptedPINBlock *ContentInformationType7 `xml:"NcrptdPINBlck"`

	// PIN (Personal Identification Number) format before encryption.
	PINFormat *PINFormat3Code `xml:"PINFrmt"`

	// Additional information required to verify the PIN (Personal Identification Number.
	AdditionalInput *Max35Text `xml:"AddtlInpt,omitempty"`
}

func (o *OnLinePIN3) AddEncryptedPINBlock() *ContentInformationType7 {
	o.EncryptedPINBlock = new(ContentInformationType7)
	return o.EncryptedPINBlock
}

func (o *OnLinePIN3) SetPINFormat(value string) {
	o.PINFormat = (*PINFormat3Code)(&value)
}

func (o *OnLinePIN3) SetAdditionalInput(value string) {
	o.AdditionalInput = (*Max35Text)(&value)
}

// Encrypted personal identification number (PIN) and related information.
type OnLinePIN4 struct {

	// Encrypted PIN (Personal Identification Number).
	EncryptedPINBlock *ContentInformationType10 `xml:"NcrptdPINBlck"`

	// PIN (Personal Identification Number) format before encryption.
	PINFormat *PINFormat3Code `xml:"PINFrmt"`

	// Additional information required to verify the PIN (Personal Identification Number).
	AdditionalInput *Max35Text `xml:"AddtlInpt,omitempty"`
}

func (o *OnLinePIN4) AddEncryptedPINBlock() *ContentInformationType10 {
	o.EncryptedPINBlock = new(ContentInformationType10)
	return o.EncryptedPINBlock
}

func (o *OnLinePIN4) SetPINFormat(value string) {
	o.PINFormat = (*PINFormat3Code)(&value)
}

func (o *OnLinePIN4) SetAdditionalInput(value string) {
	o.AdditionalInput = (*Max35Text)(&value)
}

// Encrypted personal identification number (PIN) and related information.
type OnLinePIN5 struct {

	// Encrypted PIN (Personal Identification Number).
	EncryptedPINBlock *ContentInformationType10 `xml:"NcrptdPINBlck"`

	// PIN (Personal Identification Number) format before encryption.
	PINFormat *PINFormat4Code `xml:"PINFrmt"`

	// Additional information required to verify the PIN (Personal Identification Number).
	AdditionalInput *Max35Text `xml:"AddtlInpt,omitempty"`
}

func (o *OnLinePIN5) AddEncryptedPINBlock() *ContentInformationType10 {
	o.EncryptedPINBlock = new(ContentInformationType10)
	return o.EncryptedPINBlock
}

func (o *OnLinePIN5) SetPINFormat(value string) {
	o.PINFormat = (*PINFormat4Code)(&value)
}

func (o *OnLinePIN5) SetAdditionalInput(value string) {
	o.AdditionalInput = (*Max35Text)(&value)
}
