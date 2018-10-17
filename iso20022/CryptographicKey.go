package iso20022

// Cryptographic Key to exchange.
type CryptographicKey1 struct {

	// Name of the cryptographic key.
	Identification *Max140Text `xml:"Id"`

	// Additional identification of the key.
	// Usage
	// For derived unique key per transaction (DUKPT) keys, the key serial number (KSN) with the 21 bits of the transaction counter set to zero.
	AdditionalIdentification *Max35Binary `xml:"AddtlId,omitempty"`

	// Version of the cryptographic key.
	Version *Exact10Text `xml:"Vrsn"`

	// Type of algorithm used by the cryptographic key.
	Type *CryptographicKeyType1Code `xml:"Tp,omitempty"`

	// Allowed usage of the key.
	Function []*KeyUsage1Code `xml:"Fctn"`

	// Date and time on which the key must be activated.
	ActivationDate *ISODateTime `xml:"ActvtnDt,omitempty"`

	// Date and time on which the key must be deactivated.
	DeactivationDate *ISODateTime `xml:"DeactvtnDt,omitempty"`

	// Encrypted cryptographic key.
	KeyValue *ContentInformationType2 `xml:"KeyVal"`
}

func (c *CryptographicKey1) SetIdentification(value string) {
	c.Identification = (*Max140Text)(&value)
}

func (c *CryptographicKey1) SetAdditionalIdentification(value string) {
	c.AdditionalIdentification = (*Max35Binary)(&value)
}

func (c *CryptographicKey1) SetVersion(value string) {
	c.Version = (*Exact10Text)(&value)
}

func (c *CryptographicKey1) SetType(value string) {
	c.Type = (*CryptographicKeyType1Code)(&value)
}

func (c *CryptographicKey1) AddFunction(value string) {
	c.Function = append(c.Function, (*KeyUsage1Code)(&value))
}

func (c *CryptographicKey1) SetActivationDate(value string) {
	c.ActivationDate = (*ISODateTime)(&value)
}

func (c *CryptographicKey1) SetDeactivationDate(value string) {
	c.DeactivationDate = (*ISODateTime)(&value)
}

func (c *CryptographicKey1) AddKeyValue() *ContentInformationType2 {
	c.KeyValue = new(ContentInformationType2)
	return c.KeyValue
}

// Cryptographic Key to exchange.
type CryptographicKey2 struct {

	// Name of the cryptographic key.
	Identification *Max140Text `xml:"Id"`

	// Additional identification of the key.
	// Usage
	// For derived unique key per transaction (DUKPT) keys, the key serial number (KSN) with the 21 bits of the transaction counter set to zero.
	AdditionalIdentification *Max35Binary `xml:"AddtlId,omitempty"`

	// Version of the cryptographic key.
	Version *Exact10Text `xml:"Vrsn"`

	// Type of algorithm used by the cryptographic key.
	Type *CryptographicKeyType2Code `xml:"Tp,omitempty"`

	// Allowed usage of the key.
	Function []*KeyUsage1Code `xml:"Fctn"`

	// Date and time on which the key must be activated.
	ActivationDate *ISODateTime `xml:"ActvtnDt,omitempty"`

	// Date and time on which the key must be deactivated.
	DeactivationDate *ISODateTime `xml:"DeactvtnDt,omitempty"`

	// Encrypted cryptographic key.
	KeyValue *ContentInformationType5 `xml:"KeyVal"`
}

func (c *CryptographicKey2) SetIdentification(value string) {
	c.Identification = (*Max140Text)(&value)
}

func (c *CryptographicKey2) SetAdditionalIdentification(value string) {
	c.AdditionalIdentification = (*Max35Binary)(&value)
}

func (c *CryptographicKey2) SetVersion(value string) {
	c.Version = (*Exact10Text)(&value)
}

func (c *CryptographicKey2) SetType(value string) {
	c.Type = (*CryptographicKeyType2Code)(&value)
}

func (c *CryptographicKey2) AddFunction(value string) {
	c.Function = append(c.Function, (*KeyUsage1Code)(&value))
}

func (c *CryptographicKey2) SetActivationDate(value string) {
	c.ActivationDate = (*ISODateTime)(&value)
}

func (c *CryptographicKey2) SetDeactivationDate(value string) {
	c.DeactivationDate = (*ISODateTime)(&value)
}

func (c *CryptographicKey2) AddKeyValue() *ContentInformationType5 {
	c.KeyValue = new(ContentInformationType5)
	return c.KeyValue
}

// Indetification of a cryptographic Key to use.
type CryptographicKey3 struct {

	// Name of the cryptographic key.
	Identification *Max140Text `xml:"Id"`

	// Additional identification of the key.
	// Usage
	// For derived unique key per transaction (DUKPT) keys, the key serial number (KSN) with the 21 bits of the transaction counter set to zero.
	AdditionalIdentification *Max35Binary `xml:"AddtlId,omitempty"`

	// Version of the cryptographic key.
	Version *Exact10Text `xml:"Vrsn"`
}

func (c *CryptographicKey3) SetIdentification(value string) {
	c.Identification = (*Max140Text)(&value)
}

func (c *CryptographicKey3) SetAdditionalIdentification(value string) {
	c.AdditionalIdentification = (*Max35Binary)(&value)
}

func (c *CryptographicKey3) SetVersion(value string) {
	c.Version = (*Exact10Text)(&value)
}

// Cryptographic Key to exchange.
type CryptographicKey4 struct {

	// Name of the cryptographic key.
	Identification *Max140Text `xml:"Id"`

	// Additional identification of the key.
	// Usage
	// For derived unique key per transaction (DUKPT) keys, the key serial number (KSN) with the 21 bits of the transaction counter set to zero.
	AdditionalIdentification *Max35Binary `xml:"AddtlId,omitempty"`

	// Version of the cryptographic key.
	Version *Exact10Text `xml:"Vrsn"`

	// Type of algorithm used by the cryptographic key.
	Type *CryptographicKeyType2Code `xml:"Tp,omitempty"`

	// Allowed usage of the key.
	Function []*KeyUsage1Code `xml:"Fctn"`

	// Date and time on which the key must be activated.
	ActivationDate *ISODateTime `xml:"ActvtnDt,omitempty"`

	// Date and time on which the key must be deactivated.
	DeactivationDate *ISODateTime `xml:"DeactvtnDt,omitempty"`

	// Encrypted cryptographic key.
	KeyValue *ContentInformationType7 `xml:"KeyVal"`
}

func (c *CryptographicKey4) SetIdentification(value string) {
	c.Identification = (*Max140Text)(&value)
}

func (c *CryptographicKey4) SetAdditionalIdentification(value string) {
	c.AdditionalIdentification = (*Max35Binary)(&value)
}

func (c *CryptographicKey4) SetVersion(value string) {
	c.Version = (*Exact10Text)(&value)
}

func (c *CryptographicKey4) SetType(value string) {
	c.Type = (*CryptographicKeyType2Code)(&value)
}

func (c *CryptographicKey4) AddFunction(value string) {
	c.Function = append(c.Function, (*KeyUsage1Code)(&value))
}

func (c *CryptographicKey4) SetActivationDate(value string) {
	c.ActivationDate = (*ISODateTime)(&value)
}

func (c *CryptographicKey4) SetDeactivationDate(value string) {
	c.DeactivationDate = (*ISODateTime)(&value)
}

func (c *CryptographicKey4) AddKeyValue() *ContentInformationType7 {
	c.KeyValue = new(ContentInformationType7)
	return c.KeyValue
}

// Cryptographic Key.
type CryptographicKey5 struct {

	// Name of the cryptographic key.
	Identification *Max140Text `xml:"Id"`

	// Additional identification of the key.
	// Usage
	// For derived unique key per transaction (DUKPT) keys, the key serial number (KSN) with the 21 bits of the transaction counter set to zero.
	AdditionalIdentification *Max35Binary `xml:"AddtlId,omitempty"`

	// Version of the cryptographic key.
	Version *Max256Text `xml:"Vrsn"`

	// Type of algorithm used by the cryptographic key.
	Type *CryptographicKeyType3Code `xml:"Tp"`

	// Allowed usage of the key.
	Function []*KeyUsage1Code `xml:"Fctn"`

	// Date and time on which the key must be activated.
	ActivationDate *ISODateTime `xml:"ActvtnDt,omitempty"`

	// Date and time on which the key must be deactivated.
	DeactivationDate *ISODateTime `xml:"DeactvtnDt,omitempty"`

	// Encrypted cryptographic key.
	KeyValue *ContentInformationType10 `xml:"KeyVal"`
}

func (c *CryptographicKey5) SetIdentification(value string) {
	c.Identification = (*Max140Text)(&value)
}

func (c *CryptographicKey5) SetAdditionalIdentification(value string) {
	c.AdditionalIdentification = (*Max35Binary)(&value)
}

func (c *CryptographicKey5) SetVersion(value string) {
	c.Version = (*Max256Text)(&value)
}

func (c *CryptographicKey5) SetType(value string) {
	c.Type = (*CryptographicKeyType3Code)(&value)
}

func (c *CryptographicKey5) AddFunction(value string) {
	c.Function = append(c.Function, (*KeyUsage1Code)(&value))
}

func (c *CryptographicKey5) SetActivationDate(value string) {
	c.ActivationDate = (*ISODateTime)(&value)
}

func (c *CryptographicKey5) SetDeactivationDate(value string) {
	c.DeactivationDate = (*ISODateTime)(&value)
}

func (c *CryptographicKey5) AddKeyValue() *ContentInformationType10 {
	c.KeyValue = new(ContentInformationType10)
	return c.KeyValue
}

// Cryptographic Key.
type CryptographicKey6 struct {

	// Name or label of the key.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Name of the cryptographic key.
	Identification *Max140Text `xml:"Id"`

	// Version of the cryptographic key.
	Version *Max256Text `xml:"Vrsn,omitempty"`

	// Type of algorithm used by the cryptographic key.
	Type *CryptographicKeyType3Code `xml:"Tp"`

	// Allowed usage of the key.
	Function []*KeyUsage1Code `xml:"Fctn"`

	// Date and time on which the key must be activated.
	ActivationDate *ISODateTime `xml:"ActvtnDt,omitempty"`

	// Date and time on which the key must be deactivated.
	DeactivationDate *ISODateTime `xml:"DeactvtnDt,omitempty"`

	// Encrypted value of the key present as CMS structure EnvelopedData.
	EncryptedKeyValue *ContentInformationType10 `xml:"NcrptdKeyVal,omitempty"`

	// Certificate to protect the key.
	Certificate []*Max5000Binary `xml:"Cert,omitempty"`

	// Chip card protection of the key.
	ICCRelatedData *Max5000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CryptographicKey6) SetName(value string) {
	c.Name = (*Max140Text)(&value)
}

func (c *CryptographicKey6) SetIdentification(value string) {
	c.Identification = (*Max140Text)(&value)
}

func (c *CryptographicKey6) SetVersion(value string) {
	c.Version = (*Max256Text)(&value)
}

func (c *CryptographicKey6) SetType(value string) {
	c.Type = (*CryptographicKeyType3Code)(&value)
}

func (c *CryptographicKey6) AddFunction(value string) {
	c.Function = append(c.Function, (*KeyUsage1Code)(&value))
}

func (c *CryptographicKey6) SetActivationDate(value string) {
	c.ActivationDate = (*ISODateTime)(&value)
}

func (c *CryptographicKey6) SetDeactivationDate(value string) {
	c.DeactivationDate = (*ISODateTime)(&value)
}

func (c *CryptographicKey6) AddEncryptedKeyValue() *ContentInformationType10 {
	c.EncryptedKeyValue = new(ContentInformationType10)
	return c.EncryptedKeyValue
}

func (c *CryptographicKey6) AddCertificate(value string) {
	c.Certificate = append(c.Certificate, (*Max5000Binary)(&value))
}

func (c *CryptographicKey6) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max5000Binary)(&value)
}

// Cryptographic Key component.
type CryptographicKey7 struct {

	// Name or label of the key.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Name of the cryptographic key.
	Identification *Max140Text `xml:"Id,omitempty"`

	// Identification of the security domain.
	SecurityDomainIdentification *Max35Text `xml:"SctyDomnId,omitempty"`

	// Additional identification of the key, for instance to derive the key.
	AdditionalIdentification *Max35Binary `xml:"AddtlId,omitempty"`

	// Version of the cryptographic key.
	Version *Max256Text `xml:"Vrsn"`

	// Sequence counter of the cryptographic key.
	SequenceCounter *Number `xml:"SeqCntr,omitempty"`

	// Type of algorithm used by the cryptographic key.
	Type *CryptographicKeyType3Code `xml:"Tp"`

	// Allowed usage of the key.
	Function []*KeyUsage1Code `xml:"Fctn"`

	// Date and time on which the key must be activated.
	ActivationDate *ISODateTime `xml:"ActvtnDt,omitempty"`

	// Date and time on which the key must be deactivated.
	DeactivationDate *ISODateTime `xml:"DeactvtnDt,omitempty"`

	// Value for checking a cryptographic key.
	KeyCheckValue *Max35Binary `xml:"KeyChckVal,omitempty"`

	// Current status of the key.
	CurrentStatus *ATMStatus3Code `xml:"CurSts"`

	// Reason for which the key has been stopped.
	FailureReason *FailureReason6Code `xml:"FailrRsn,omitempty"`
}

func (c *CryptographicKey7) SetName(value string) {
	c.Name = (*Max140Text)(&value)
}

func (c *CryptographicKey7) SetIdentification(value string) {
	c.Identification = (*Max140Text)(&value)
}

func (c *CryptographicKey7) SetSecurityDomainIdentification(value string) {
	c.SecurityDomainIdentification = (*Max35Text)(&value)
}

func (c *CryptographicKey7) SetAdditionalIdentification(value string) {
	c.AdditionalIdentification = (*Max35Binary)(&value)
}

func (c *CryptographicKey7) SetVersion(value string) {
	c.Version = (*Max256Text)(&value)
}

func (c *CryptographicKey7) SetSequenceCounter(value string) {
	c.SequenceCounter = (*Number)(&value)
}

func (c *CryptographicKey7) SetType(value string) {
	c.Type = (*CryptographicKeyType3Code)(&value)
}

func (c *CryptographicKey7) AddFunction(value string) {
	c.Function = append(c.Function, (*KeyUsage1Code)(&value))
}

func (c *CryptographicKey7) SetActivationDate(value string) {
	c.ActivationDate = (*ISODateTime)(&value)
}

func (c *CryptographicKey7) SetDeactivationDate(value string) {
	c.DeactivationDate = (*ISODateTime)(&value)
}

func (c *CryptographicKey7) SetKeyCheckValue(value string) {
	c.KeyCheckValue = (*Max35Binary)(&value)
}

func (c *CryptographicKey7) SetCurrentStatus(value string) {
	c.CurrentStatus = (*ATMStatus3Code)(&value)
}

func (c *CryptographicKey7) SetFailureReason(value string) {
	c.FailureReason = (*FailureReason6Code)(&value)
}

// Cryptographic key.
type CryptographicKey8 struct {

	// Name or label of the key.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Name of the cryptographic key.
	Identification *Max140Text `xml:"Id,omitempty"`

	// Identification of the security domain.
	SecurityDomainIdentification *Max35Text `xml:"SctyDomnId,omitempty"`

	// Additional identification of the key, for instance to derive the key.
	AdditionalIdentification *Max35Binary `xml:"AddtlId,omitempty"`

	// Version of the cryptographic key.
	Version *Max256Text `xml:"Vrsn,omitempty"`

	// Sequence counter of the cryptographic key.
	SequenceCounter *Number `xml:"SeqCntr,omitempty"`

	// Type of algorithm used by the cryptographic key.
	Type *CryptographicKeyType3Code `xml:"Tp"`

	// Allowed usage of the key.
	Function []*KeyUsage1Code `xml:"Fctn"`

	// Date and time on which the key must be activated.
	ActivationDate *ISODateTime `xml:"ActvtnDt,omitempty"`

	// Date and time on which the key must be deactivated.
	DeactivationDate *ISODateTime `xml:"DeactvtnDt,omitempty"`

	// Value for checking a cryptographic key.
	KeyCheckValue *Max35Binary `xml:"KeyChckVal,omitempty"`

	// Encrypted value of the cryptographic key.
	EncryptedKeyValue *ContentInformationType10 `xml:"NcrptdKeyVal,omitempty"`

	// Value of the public component of a RSA key.
	PublicKeyValue *PublicRSAKey1 `xml:"PblcKeyVal,omitempty"`
}

func (c *CryptographicKey8) SetName(value string) {
	c.Name = (*Max140Text)(&value)
}

func (c *CryptographicKey8) SetIdentification(value string) {
	c.Identification = (*Max140Text)(&value)
}

func (c *CryptographicKey8) SetSecurityDomainIdentification(value string) {
	c.SecurityDomainIdentification = (*Max35Text)(&value)
}

func (c *CryptographicKey8) SetAdditionalIdentification(value string) {
	c.AdditionalIdentification = (*Max35Binary)(&value)
}

func (c *CryptographicKey8) SetVersion(value string) {
	c.Version = (*Max256Text)(&value)
}

func (c *CryptographicKey8) SetSequenceCounter(value string) {
	c.SequenceCounter = (*Number)(&value)
}

func (c *CryptographicKey8) SetType(value string) {
	c.Type = (*CryptographicKeyType3Code)(&value)
}

func (c *CryptographicKey8) AddFunction(value string) {
	c.Function = append(c.Function, (*KeyUsage1Code)(&value))
}

func (c *CryptographicKey8) SetActivationDate(value string) {
	c.ActivationDate = (*ISODateTime)(&value)
}

func (c *CryptographicKey8) SetDeactivationDate(value string) {
	c.DeactivationDate = (*ISODateTime)(&value)
}

func (c *CryptographicKey8) SetKeyCheckValue(value string) {
	c.KeyCheckValue = (*Max35Binary)(&value)
}

func (c *CryptographicKey8) AddEncryptedKeyValue() *ContentInformationType10 {
	c.EncryptedKeyValue = new(ContentInformationType10)
	return c.EncryptedKeyValue
}

func (c *CryptographicKey8) AddPublicKeyValue() *PublicRSAKey1 {
	c.PublicKeyValue = new(PublicRSAKey1)
	return c.PublicKeyValue
}
