package iso20022

// Encrypted data with encryption key.
type EnvelopedData1 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Transport key or key encryption key (KEK) identification for the recipient.
	Recipient []*Recipient1Choice `xml:"Rcpt"`

	// Encrypted data with an encryption key.
	EncryptedContent *EncryptedContent1 `xml:"NcrptdCntt"`
}

func (e *EnvelopedData1) SetVersion(value string) {
	e.Version = (*Number)(&value)
}

func (e *EnvelopedData1) AddRecipient() *Recipient1Choice {
	newValue := new(Recipient1Choice)
	e.Recipient = append(e.Recipient, newValue)
	return newValue
}

func (e *EnvelopedData1) AddEncryptedContent() *EncryptedContent1 {
	e.EncryptedContent = new(EncryptedContent1)
	return e.EncryptedContent
}

// Encrypted data with encryption key.
type EnvelopedData2 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Transport key or key encryption key (KEK) identification for the recipient.
	Recipient []*Recipient2Choice `xml:"Rcpt"`

	// Encrypted data with an encryption key.
	EncryptedContent *EncryptedContent2 `xml:"NcrptdCntt"`
}

func (e *EnvelopedData2) SetVersion(value string) {
	e.Version = (*Number)(&value)
}

func (e *EnvelopedData2) AddRecipient() *Recipient2Choice {
	newValue := new(Recipient2Choice)
	e.Recipient = append(e.Recipient, newValue)
	return newValue
}

func (e *EnvelopedData2) AddEncryptedContent() *EncryptedContent2 {
	e.EncryptedContent = new(EncryptedContent2)
	return e.EncryptedContent
}

// Encrypted data with encryption key.
type EnvelopedData3 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Transport key or key encryption key (KEK) identification for the recipient.
	Recipient []*Recipient3Choice `xml:"Rcpt"`

	// Encrypted data with an encryption key.
	EncryptedContent *EncryptedContent2 `xml:"NcrptdCntt"`
}

func (e *EnvelopedData3) SetVersion(value string) {
	e.Version = (*Number)(&value)
}

func (e *EnvelopedData3) AddRecipient() *Recipient3Choice {
	newValue := new(Recipient3Choice)
	e.Recipient = append(e.Recipient, newValue)
	return newValue
}

func (e *EnvelopedData3) AddEncryptedContent() *EncryptedContent2 {
	e.EncryptedContent = new(EncryptedContent2)
	return e.EncryptedContent
}

// Encrypted data with encryption key.
type EnvelopedData4 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Session key or identification of the protection key used by the recipient.
	Recipient []*Recipient4Choice `xml:"Rcpt"`

	// Data protection by encryption (digital envelope), with an encryption key.
	EncryptedContent *EncryptedContent3 `xml:"NcrptdCntt,omitempty"`
}

func (e *EnvelopedData4) SetVersion(value string) {
	e.Version = (*Number)(&value)
}

func (e *EnvelopedData4) AddRecipient() *Recipient4Choice {
	newValue := new(Recipient4Choice)
	e.Recipient = append(e.Recipient, newValue)
	return newValue
}

func (e *EnvelopedData4) AddEncryptedContent() *EncryptedContent3 {
	e.EncryptedContent = new(EncryptedContent3)
	return e.EncryptedContent
}
