package iso20022

// Message authentication code (MAC), computed on the data to protect with an encryption key.
type AuthenticatedData1 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Information related to the transport key.
	Recipient []*Recipient1Choice `xml:"Rcpt"`

	// Algorithm to compute message authentication code (MAC).
	MACAlgorithm *AlgorithmIdentification1 `xml:"MACAlgo"`

	// Data to authenticate.
	EncapsulatedContent *EncapsulatedContent1 `xml:"NcpsltdCntt"`

	// Encrypted data which authenticates the data.
	MAC *Max35Binary `xml:"MAC"`
}

func (a *AuthenticatedData1) SetVersion(value string) {
	a.Version = (*Number)(&value)
}

func (a *AuthenticatedData1) AddRecipient() *Recipient1Choice {
	newValue := new(Recipient1Choice)
	a.Recipient = append(a.Recipient, newValue)
	return newValue
}

func (a *AuthenticatedData1) AddMACAlgorithm() *AlgorithmIdentification1 {
	a.MACAlgorithm = new(AlgorithmIdentification1)
	return a.MACAlgorithm
}

func (a *AuthenticatedData1) AddEncapsulatedContent() *EncapsulatedContent1 {
	a.EncapsulatedContent = new(EncapsulatedContent1)
	return a.EncapsulatedContent
}

func (a *AuthenticatedData1) SetMAC(value string) {
	a.MAC = (*Max35Binary)(&value)
}

// Message authentication code (MAC), computed on the data to protect with an encryption key.
type AuthenticatedData2 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Information related to the transport key.
	Recipient []*Recipient2Choice `xml:"Rcpt"`

	// Algorithm to compute message authentication code (MAC).
	MACAlgorithm *AlgorithmIdentification3 `xml:"MACAlgo"`

	// Data to authenticate.
	EncapsulatedContent *EncapsulatedContent1 `xml:"NcpsltdCntt"`

	// Encrypted data which authenticates the data.
	MAC *Max35Binary `xml:"MAC"`
}

func (a *AuthenticatedData2) SetVersion(value string) {
	a.Version = (*Number)(&value)
}

func (a *AuthenticatedData2) AddRecipient() *Recipient2Choice {
	newValue := new(Recipient2Choice)
	a.Recipient = append(a.Recipient, newValue)
	return newValue
}

func (a *AuthenticatedData2) AddMACAlgorithm() *AlgorithmIdentification3 {
	a.MACAlgorithm = new(AlgorithmIdentification3)
	return a.MACAlgorithm
}

func (a *AuthenticatedData2) AddEncapsulatedContent() *EncapsulatedContent1 {
	a.EncapsulatedContent = new(EncapsulatedContent1)
	return a.EncapsulatedContent
}

func (a *AuthenticatedData2) SetMAC(value string) {
	a.MAC = (*Max35Binary)(&value)
}

// Message authentication code (MAC), computed on the data to protect with an encryption key.
type AuthenticatedData3 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Information related to the transport key.
	Recipient []*Recipient3Choice `xml:"Rcpt"`

	// Algorithm to compute message authentication code (MAC).
	MACAlgorithm *AlgorithmIdentification10 `xml:"MACAlgo"`

	// Data to authenticate.
	EncapsulatedContent *EncapsulatedContent2 `xml:"NcpsltdCntt"`

	// Encrypted data which authenticates the data.
	MAC *Max35Binary `xml:"MAC"`
}

func (a *AuthenticatedData3) SetVersion(value string) {
	a.Version = (*Number)(&value)
}

func (a *AuthenticatedData3) AddRecipient() *Recipient3Choice {
	newValue := new(Recipient3Choice)
	a.Recipient = append(a.Recipient, newValue)
	return newValue
}

func (a *AuthenticatedData3) AddMACAlgorithm() *AlgorithmIdentification10 {
	a.MACAlgorithm = new(AlgorithmIdentification10)
	return a.MACAlgorithm
}

func (a *AuthenticatedData3) AddEncapsulatedContent() *EncapsulatedContent2 {
	a.EncapsulatedContent = new(EncapsulatedContent2)
	return a.EncapsulatedContent
}

func (a *AuthenticatedData3) SetMAC(value string) {
	a.MAC = (*Max35Binary)(&value)
}

// Message authentication code (MAC), computed on the data to protect with an encryption key.
type AuthenticatedData4 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Session key or protection key identification used by the recipient.
	Recipient []*Recipient4Choice `xml:"Rcpt"`

	// Algorithm to compute message authentication code (MAC).
	MACAlgorithm *AlgorithmIdentification15 `xml:"MACAlgo"`

	// Data to authenticate.
	EncapsulatedContent *EncapsulatedContent3 `xml:"NcpsltdCntt"`

	// Message authentication code value.
	MAC *Max140Binary `xml:"MAC"`
}

func (a *AuthenticatedData4) SetVersion(value string) {
	a.Version = (*Number)(&value)
}

func (a *AuthenticatedData4) AddRecipient() *Recipient4Choice {
	newValue := new(Recipient4Choice)
	a.Recipient = append(a.Recipient, newValue)
	return newValue
}

func (a *AuthenticatedData4) AddMACAlgorithm() *AlgorithmIdentification15 {
	a.MACAlgorithm = new(AlgorithmIdentification15)
	return a.MACAlgorithm
}

func (a *AuthenticatedData4) AddEncapsulatedContent() *EncapsulatedContent3 {
	a.EncapsulatedContent = new(EncapsulatedContent3)
	return a.EncapsulatedContent
}

func (a *AuthenticatedData4) SetMAC(value string) {
	a.MAC = (*Max140Binary)(&value)
}
