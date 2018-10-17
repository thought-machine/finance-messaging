package iso20022

// Configuration parameters in use by the security device.
type ATMSecurityConfiguration1 struct {

	// Configuration of the cryptographic keys.
	Keys *ATMSecurityConfiguration2 `xml:"Keys,omitempty"`

	// Configuration of the encryption or digital envelope, if the security module is able to perform encryption.
	Encryption *ATMSecurityConfiguration3 `xml:"Ncrptn,omitempty"`

	// MAC (Message Authentication Code) algorithm the security module is able to manage.
	MACAlgorithm []*Algorithm12Code `xml:"MACAlgo,omitempty"`

	// Digest algorithm the security module is able to manage.
	DigestAlgorithm []*Algorithm11Code `xml:"DgstAlgo,omitempty"`

	// Configuration of the digital signatures if the security module is able to perform digital signatures with an asymmetric key.
	DigitalSignature *ATMSecurityConfiguration4 `xml:"DgtlSgntr,omitempty"`

	// Configuration of the PIN online verification.
	PIN *ATMSecurityConfiguration5 `xml:"PIN,omitempty"`

	// Mechanism used to protect the message of the ATM protocol.
	MessageProtection []*MessageProtection1Code `xml:"MsgPrtcn,omitempty"`
}

func (a *ATMSecurityConfiguration1) AddKeys() *ATMSecurityConfiguration2 {
	a.Keys = new(ATMSecurityConfiguration2)
	return a.Keys
}

func (a *ATMSecurityConfiguration1) AddEncryption() *ATMSecurityConfiguration3 {
	a.Encryption = new(ATMSecurityConfiguration3)
	return a.Encryption
}

func (a *ATMSecurityConfiguration1) AddMACAlgorithm(value string) {
	a.MACAlgorithm = append(a.MACAlgorithm, (*Algorithm12Code)(&value))
}

func (a *ATMSecurityConfiguration1) AddDigestAlgorithm(value string) {
	a.DigestAlgorithm = append(a.DigestAlgorithm, (*Algorithm11Code)(&value))
}

func (a *ATMSecurityConfiguration1) AddDigitalSignature() *ATMSecurityConfiguration4 {
	a.DigitalSignature = new(ATMSecurityConfiguration4)
	return a.DigitalSignature
}

func (a *ATMSecurityConfiguration1) AddPIN() *ATMSecurityConfiguration5 {
	a.PIN = new(ATMSecurityConfiguration5)
	return a.PIN
}

func (a *ATMSecurityConfiguration1) AddMessageProtection(value string) {
	a.MessageProtection = append(a.MessageProtection, (*MessageProtection1Code)(&value))
}

// Configuration of the cryptographic keys.
type ATMSecurityConfiguration2 struct {

	// Maximum number of symmetric keys the security module is able to manage.
	MaximumSymmetricKey *Number `xml:"MaxSmmtrcKey,omitempty"`

	// Maximum number of asymmetric keys the security module is able to manage.
	MaximumAsymmetricKey *Number `xml:"MaxAsmmtrcKey,omitempty"`

	// Maximum RSA key length (in number of bytes), the security module is able to manage.
	MaximumRSAKeyLength *Number `xml:"MaxRSAKeyLngth,omitempty"`

	// Maximum RSA root key length (in number of bytes), the security module is able to manage, if different from the maximum RSA key length.
	MaximumRootKeyLength *Number `xml:"MaxRootKeyLngth,omitempty"`
}

func (a *ATMSecurityConfiguration2) SetMaximumSymmetricKey(value string) {
	a.MaximumSymmetricKey = (*Number)(&value)
}

func (a *ATMSecurityConfiguration2) SetMaximumAsymmetricKey(value string) {
	a.MaximumAsymmetricKey = (*Number)(&value)
}

func (a *ATMSecurityConfiguration2) SetMaximumRSAKeyLength(value string) {
	a.MaximumRSAKeyLength = (*Number)(&value)
}

func (a *ATMSecurityConfiguration2) SetMaximumRootKeyLength(value string) {
	a.MaximumRootKeyLength = (*Number)(&value)
}

// Configuration of the encryption or digital envelope, if the security module is able to perform encryption.
type ATMSecurityConfiguration3 struct {

	// True if the security module is able to perform encryption with an asymmetric key.
	AsymmetricEncryption *TrueFalseIndicator `xml:"AsmmtrcNcrptn,omitempty"`

	// True if the security module is able to identify an asymmetric key with certificate issuer X.500 name and certificate serial number. False if a proprietary asymmetric key identifier is required.
	AsymmetricKeyStandardIdentification *TrueFalseIndicator `xml:"AsmmtrcKeyStdId,omitempty"`

	// Asymmetric encryption algorithm the security module is able to manage.
	AsymmetricEncryptionAlgorithm []*Algorithm7Code `xml:"AsmmtrcNcrptnAlgo,omitempty"`

	// True if the security module is able to manage a symmetric transport session key to protect cryptographic keys and data. False if only a previously exchanged symmetric key must be used; a proprietary symmetric key identifier is then used.
	SymmetricTransportKey *TrueFalseIndicator `xml:"SmmtrcTrnsprtKey,omitempty"`

	// Symmetric transport session key algorithm the security module is able to manage.
	SymmetricTransportKeyAlgorithm []*Algorithm13Code `xml:"SmmtrcTrnsprtKeyAlgo,omitempty"`

	// Symmetric encryption algorithm the security module is able to manage.
	SymmetricEncryptionAlgorithm []*Algorithm15Code `xml:"SmmtrcNcrptnAlgo,omitempty"`

	// Format of data before encryption, if the format is not plaintext or implicit.
	EncryptionFormat []*EncryptionFormat1Code `xml:"NcrptnFrmt,omitempty"`
}

func (a *ATMSecurityConfiguration3) SetAsymmetricEncryption(value string) {
	a.AsymmetricEncryption = (*TrueFalseIndicator)(&value)
}

func (a *ATMSecurityConfiguration3) SetAsymmetricKeyStandardIdentification(value string) {
	a.AsymmetricKeyStandardIdentification = (*TrueFalseIndicator)(&value)
}

func (a *ATMSecurityConfiguration3) AddAsymmetricEncryptionAlgorithm(value string) {
	a.AsymmetricEncryptionAlgorithm = append(a.AsymmetricEncryptionAlgorithm, (*Algorithm7Code)(&value))
}

func (a *ATMSecurityConfiguration3) SetSymmetricTransportKey(value string) {
	a.SymmetricTransportKey = (*TrueFalseIndicator)(&value)
}

func (a *ATMSecurityConfiguration3) AddSymmetricTransportKeyAlgorithm(value string) {
	a.SymmetricTransportKeyAlgorithm = append(a.SymmetricTransportKeyAlgorithm, (*Algorithm13Code)(&value))
}

func (a *ATMSecurityConfiguration3) AddSymmetricEncryptionAlgorithm(value string) {
	a.SymmetricEncryptionAlgorithm = append(a.SymmetricEncryptionAlgorithm, (*Algorithm15Code)(&value))
}

func (a *ATMSecurityConfiguration3) AddEncryptionFormat(value string) {
	a.EncryptionFormat = append(a.EncryptionFormat, (*EncryptionFormat1Code)(&value))
}

// Configuration of the digital signatures if the security module is able to perform digital signatures with an asymmetric key.
type ATMSecurityConfiguration4 struct {

	// Maximum number of certificates in a certificate path, the security module is able to manage.
	MaximumCertificates *Number `xml:"MaxCerts,omitempty"`

	// Maximum number of cosigners, the security module is able to manage in a digital signature.
	MaximumSignatures *Number `xml:"MaxSgntrs,omitempty"`

	// Digital signature algorithm the security module is able to manage.
	DigitalSignatureAlgorithm []*Algorithm14Code `xml:"DgtlSgntrAlgo,omitempty"`
}

func (a *ATMSecurityConfiguration4) SetMaximumCertificates(value string) {
	a.MaximumCertificates = (*Number)(&value)
}

func (a *ATMSecurityConfiguration4) SetMaximumSignatures(value string) {
	a.MaximumSignatures = (*Number)(&value)
}

func (a *ATMSecurityConfiguration4) AddDigitalSignatureAlgorithm(value string) {
	a.DigitalSignatureAlgorithm = append(a.DigitalSignatureAlgorithm, (*Algorithm14Code)(&value))
}

// Configuration of the PIN online verification.
type ATMSecurityConfiguration5 struct {

	// PIN block format the security module is able to manage for online verification of the PIN.
	PINFormat []*PINFormat4Code `xml:"PINFrmt,omitempty"`

	// Maximum number of digits the security module is able to accept when the cardholder enters its PIN.
	PINLengthCapabilities *Number `xml:"PINLngthCpblties,omitempty"`
}

func (a *ATMSecurityConfiguration5) AddPINFormat(value string) {
	a.PINFormat = append(a.PINFormat, (*PINFormat4Code)(&value))
}

func (a *ATMSecurityConfiguration5) SetPINLengthCapabilities(value string) {
	a.PINLengthCapabilities = (*Number)(&value)
}
