package iso20022

// General cryptographic message syntax (CMS) containing protected data.
type ContentInformationType1 struct {

	// Type of data protection.
	ContentType *ContentType1Code `xml:"CnttTp"`

	// Data protection by encryption, with a session key.
	EnvelopedData *EnvelopedData1 `xml:"EnvlpdData,omitempty"`

	// Data protection by a message authentication code (MAC).
	AuthenticatedData *AuthenticatedData1 `xml:"AuthntcdData,omitempty"`

	// Data protected by digital signatures.
	SignedData *SignedData1 `xml:"SgndData,omitempty"`

	// Data protected by a digest.
	DigestedData *DigestedData1 `xml:"DgstdData,omitempty"`

	// Data protection by encryption with a previously exchanged key identified by a name.
	NamedKeyEncryptedData *NamedKeyEncryptedData1 `xml:"NmdKeyNcrptdData,omitempty"`
}

func (c *ContentInformationType1) SetContentType(value string) {
	c.ContentType = (*ContentType1Code)(&value)
}

func (c *ContentInformationType1) AddEnvelopedData() *EnvelopedData1 {
	c.EnvelopedData = new(EnvelopedData1)
	return c.EnvelopedData
}

func (c *ContentInformationType1) AddAuthenticatedData() *AuthenticatedData1 {
	c.AuthenticatedData = new(AuthenticatedData1)
	return c.AuthenticatedData
}

func (c *ContentInformationType1) AddSignedData() *SignedData1 {
	c.SignedData = new(SignedData1)
	return c.SignedData
}

func (c *ContentInformationType1) AddDigestedData() *DigestedData1 {
	c.DigestedData = new(DigestedData1)
	return c.DigestedData
}

func (c *ContentInformationType1) AddNamedKeyEncryptedData() *NamedKeyEncryptedData1 {
	c.NamedKeyEncryptedData = new(NamedKeyEncryptedData1)
	return c.NamedKeyEncryptedData
}

// General cryptographic message syntax (CMS) containing encrypted data.
type ContentInformationType10 struct {

	// Type of data protection.
	ContentType *ContentType2Code `xml:"CnttTp"`

	// Data protection by encryption or by a digital envelope, with an encryption key.
	EnvelopedData *EnvelopedData4 `xml:"EnvlpdData"`
}

func (c *ContentInformationType10) SetContentType(value string) {
	c.ContentType = (*ContentType2Code)(&value)
}

func (c *ContentInformationType10) AddEnvelopedData() *EnvelopedData4 {
	c.EnvelopedData = new(EnvelopedData4)
	return c.EnvelopedData
}

// General cryptographic message syntax (CMS) containing authenticated data.
type ContentInformationType11 struct {

	// Type of data protection.
	ContentType *ContentType2Code `xml:"CnttTp"`

	// Data protection by a message authentication code (MAC).
	AuthenticatedData []*AuthenticatedData4 `xml:"AuthntcdData,omitempty"`
}

func (c *ContentInformationType11) SetContentType(value string) {
	c.ContentType = (*ContentType2Code)(&value)
}

func (c *ContentInformationType11) AddAuthenticatedData() *AuthenticatedData4 {
	newValue := new(AuthenticatedData4)
	c.AuthenticatedData = append(c.AuthenticatedData, newValue)
	return newValue
}

// General cryptographic message syntax (CMS) containing protected data.
type ContentInformationType12 struct {

	// Type of data protection.
	ContentType *ContentType2Code `xml:"CnttTp"`

	// Data protection by encryption, with a session key.
	EnvelopedData *EnvelopedData4 `xml:"EnvlpdData,omitempty"`

	// Data protection by a message authentication code (MAC).
	AuthenticatedData *AuthenticatedData4 `xml:"AuthntcdData,omitempty"`

	// Data protected by a digital signatures.
	SignedData *SignedData4 `xml:"SgndData,omitempty"`

	// Data protected by a digest.
	DigestedData *DigestedData4 `xml:"DgstdData,omitempty"`
}

func (c *ContentInformationType12) SetContentType(value string) {
	c.ContentType = (*ContentType2Code)(&value)
}

func (c *ContentInformationType12) AddEnvelopedData() *EnvelopedData4 {
	c.EnvelopedData = new(EnvelopedData4)
	return c.EnvelopedData
}

func (c *ContentInformationType12) AddAuthenticatedData() *AuthenticatedData4 {
	c.AuthenticatedData = new(AuthenticatedData4)
	return c.AuthenticatedData
}

func (c *ContentInformationType12) AddSignedData() *SignedData4 {
	c.SignedData = new(SignedData4)
	return c.SignedData
}

func (c *ContentInformationType12) AddDigestedData() *DigestedData4 {
	c.DigestedData = new(DigestedData4)
	return c.DigestedData
}

// General cryptographic message syntax (CMS) containing data. protected by a MAC or a digital signature
type ContentInformationType13 struct {

	// Type of data protection.
	ContentType *ContentType2Code `xml:"CnttTp"`

	// Data protection by a message authentication code (MAC).
	AuthenticatedData *AuthenticatedData4 `xml:"AuthntcdData,omitempty"`

	// Data protected by a digital signatures.
	SignedData *SignedData4 `xml:"SgndData,omitempty"`
}

func (c *ContentInformationType13) SetContentType(value string) {
	c.ContentType = (*ContentType2Code)(&value)
}

func (c *ContentInformationType13) AddAuthenticatedData() *AuthenticatedData4 {
	c.AuthenticatedData = new(AuthenticatedData4)
	return c.AuthenticatedData
}

func (c *ContentInformationType13) AddSignedData() *SignedData4 {
	c.SignedData = new(SignedData4)
	return c.SignedData
}

// General cryptographic message syntax (CMS) containing data. protected by a digital signature
type ContentInformationType14 struct {

	// Type of data protection.
	ContentType *ContentType2Code `xml:"CnttTp"`

	// Data protected by a digital signatures.
	SignedData *SignedData4 `xml:"SgndData"`
}

func (c *ContentInformationType14) SetContentType(value string) {
	c.ContentType = (*ContentType2Code)(&value)
}

func (c *ContentInformationType14) AddSignedData() *SignedData4 {
	c.SignedData = new(SignedData4)
	return c.SignedData
}

// General cryptographic message syntax (CMS) containing authenticated data.
type ContentInformationType15 struct {

	// Type of data protection.
	ContentType *ContentType2Code `xml:"CnttTp"`

	// Data protection by a message authentication code (MAC).
	AuthenticatedData *AuthenticatedData4 `xml:"AuthntcdData"`
}

func (c *ContentInformationType15) SetContentType(value string) {
	c.ContentType = (*ContentType2Code)(&value)
}

func (c *ContentInformationType15) AddAuthenticatedData() *AuthenticatedData4 {
	c.AuthenticatedData = new(AuthenticatedData4)
	return c.AuthenticatedData
}

// General cryptographic message syntax (CMS) containing encrypted data.
type ContentInformationType2 struct {

	// Type of data protection.
	ContentType *ContentType1Code `xml:"CnttTp"`

	// Data protection by encryption, with a session key.
	EnvelopedData *EnvelopedData1 `xml:"EnvlpdData"`
}

func (c *ContentInformationType2) SetContentType(value string) {
	c.ContentType = (*ContentType1Code)(&value)
}

func (c *ContentInformationType2) AddEnvelopedData() *EnvelopedData1 {
	c.EnvelopedData = new(EnvelopedData1)
	return c.EnvelopedData
}

// General cryptographic message syntax (CMS) containing authenticated data.
type ContentInformationType3 struct {

	// Type of data protection.
	ContentType *ContentType1Code `xml:"CnttTp"`

	// Data protection by a message authentication code (MAC).
	AuthenticatedData []*AuthenticatedData1 `xml:"AuthntcdData,omitempty"`
}

func (c *ContentInformationType3) SetContentType(value string) {
	c.ContentType = (*ContentType1Code)(&value)
}

func (c *ContentInformationType3) AddAuthenticatedData() *AuthenticatedData1 {
	newValue := new(AuthenticatedData1)
	c.AuthenticatedData = append(c.AuthenticatedData, newValue)
	return newValue
}

// General cryptographic message syntax (CMS) containing protected data.
type ContentInformationType4 struct {

	// Type of data protection.
	ContentType *ContentType1Code `xml:"CnttTp"`

	// Data protection by encryption, with a session key.
	EnvelopedData *EnvelopedData2 `xml:"EnvlpdData,omitempty"`

	// Data protection by a message authentication code (MAC).
	AuthenticatedData *AuthenticatedData2 `xml:"AuthntcdData,omitempty"`

	// Data protected by digital signatures.
	SignedData *SignedData2 `xml:"SgndData,omitempty"`

	// Data protected by a digest.
	DigestedData *DigestedData2 `xml:"DgstdData,omitempty"`

	// Data protection by encryption with a previously exchanged key identified by a name.
	NamedKeyEncryptedData *NamedKeyEncryptedData2 `xml:"NmdKeyNcrptdData,omitempty"`
}

func (c *ContentInformationType4) SetContentType(value string) {
	c.ContentType = (*ContentType1Code)(&value)
}

func (c *ContentInformationType4) AddEnvelopedData() *EnvelopedData2 {
	c.EnvelopedData = new(EnvelopedData2)
	return c.EnvelopedData
}

func (c *ContentInformationType4) AddAuthenticatedData() *AuthenticatedData2 {
	c.AuthenticatedData = new(AuthenticatedData2)
	return c.AuthenticatedData
}

func (c *ContentInformationType4) AddSignedData() *SignedData2 {
	c.SignedData = new(SignedData2)
	return c.SignedData
}

func (c *ContentInformationType4) AddDigestedData() *DigestedData2 {
	c.DigestedData = new(DigestedData2)
	return c.DigestedData
}

func (c *ContentInformationType4) AddNamedKeyEncryptedData() *NamedKeyEncryptedData2 {
	c.NamedKeyEncryptedData = new(NamedKeyEncryptedData2)
	return c.NamedKeyEncryptedData
}

// General cryptographic message syntax (CMS) containing encrypted data.
type ContentInformationType5 struct {

	// Type of data protection.
	ContentType *ContentType1Code `xml:"CnttTp"`

	// Data protection by encryption, with a session key.
	EnvelopedData *EnvelopedData2 `xml:"EnvlpdData"`
}

func (c *ContentInformationType5) SetContentType(value string) {
	c.ContentType = (*ContentType1Code)(&value)
}

func (c *ContentInformationType5) AddEnvelopedData() *EnvelopedData2 {
	c.EnvelopedData = new(EnvelopedData2)
	return c.EnvelopedData
}

// General cryptographic message syntax (CMS) containing authenticated data.
type ContentInformationType6 struct {

	// Type of data protection.
	ContentType *ContentType1Code `xml:"CnttTp"`

	// Data protection by a message authentication code (MAC).
	AuthenticatedData []*AuthenticatedData2 `xml:"AuthntcdData,omitempty"`
}

func (c *ContentInformationType6) SetContentType(value string) {
	c.ContentType = (*ContentType1Code)(&value)
}

func (c *ContentInformationType6) AddAuthenticatedData() *AuthenticatedData2 {
	newValue := new(AuthenticatedData2)
	c.AuthenticatedData = append(c.AuthenticatedData, newValue)
	return newValue
}

// General cryptographic message syntax (CMS) containing encrypted data.
type ContentInformationType7 struct {

	// Type of data protection.
	ContentType *ContentType1Code `xml:"CnttTp"`

	// Data protection by encryption, with a session key.
	EnvelopedData *EnvelopedData3 `xml:"EnvlpdData"`
}

func (c *ContentInformationType7) SetContentType(value string) {
	c.ContentType = (*ContentType1Code)(&value)
}

func (c *ContentInformationType7) AddEnvelopedData() *EnvelopedData3 {
	c.EnvelopedData = new(EnvelopedData3)
	return c.EnvelopedData
}

// General cryptographic message syntax (CMS) containing authenticated data.
type ContentInformationType8 struct {

	// Type of data protection.
	ContentType *ContentType1Code `xml:"CnttTp"`

	// Data protection by a message authentication code (MAC).
	AuthenticatedData []*AuthenticatedData3 `xml:"AuthntcdData,omitempty"`
}

func (c *ContentInformationType8) SetContentType(value string) {
	c.ContentType = (*ContentType1Code)(&value)
}

func (c *ContentInformationType8) AddAuthenticatedData() *AuthenticatedData3 {
	newValue := new(AuthenticatedData3)
	c.AuthenticatedData = append(c.AuthenticatedData, newValue)
	return newValue
}

// General cryptographic message syntax (CMS) containing protected data.
type ContentInformationType9 struct {

	// Type of data protection.
	ContentType *ContentType1Code `xml:"CnttTp"`

	// Data protection by encryption, with a session key.
	EnvelopedData *EnvelopedData3 `xml:"EnvlpdData,omitempty"`

	// Data protection by a message authentication code (MAC).
	AuthenticatedData *AuthenticatedData3 `xml:"AuthntcdData,omitempty"`

	// Data protected by digital signatures.
	SignedData *SignedData3 `xml:"SgndData,omitempty"`

	// Data protected by a digest.
	DigestedData *DigestedData3 `xml:"DgstdData,omitempty"`
}

func (c *ContentInformationType9) SetContentType(value string) {
	c.ContentType = (*ContentType1Code)(&value)
}

func (c *ContentInformationType9) AddEnvelopedData() *EnvelopedData3 {
	c.EnvelopedData = new(EnvelopedData3)
	return c.EnvelopedData
}

func (c *ContentInformationType9) AddAuthenticatedData() *AuthenticatedData3 {
	c.AuthenticatedData = new(AuthenticatedData3)
	return c.AuthenticatedData
}

func (c *ContentInformationType9) AddSignedData() *SignedData3 {
	c.SignedData = new(SignedData3)
	return c.SignedData
}

func (c *ContentInformationType9) AddDigestedData() *DigestedData3 {
	c.DigestedData = new(DigestedData3)
	return c.DigestedData
}
