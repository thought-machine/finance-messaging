package iso20022

// Identification of a key encryption key (KEK), using previously distributed symmetric key.
type KEKIdentifier1 struct {

	// Identification of the cryptographic key.
	KeyIdentification *Max140Text `xml:"KeyId"`

	// Version of the cryptographic key.
	KeyVersion *Exact10Text `xml:"KeyVrsn"`

	// Identification used for derivation of a unique key from a master key provided for the data protection.
	DerivationIdentification *Min5Max16Binary `xml:"DerivtnId,omitempty"`
}

func (k *KEKIdentifier1) SetKeyIdentification(value string) {
	k.KeyIdentification = (*Max140Text)(&value)
}

func (k *KEKIdentifier1) SetKeyVersion(value string) {
	k.KeyVersion = (*Exact10Text)(&value)
}

func (k *KEKIdentifier1) SetDerivationIdentification(value string) {
	k.DerivationIdentification = (*Min5Max16Binary)(&value)
}

// Identification of a key encryption key (KEK), using previously distributed symmetric key.
type KEKIdentifier2 struct {

	// Identification of the cryptographic key.
	KeyIdentification *Max140Text `xml:"KeyId"`

	// Version of the cryptographic key.
	KeyVersion *Max140Text `xml:"KeyVrsn"`

	// Number of usages of the cryptographic key.
	SequenceNumber *Number `xml:"SeqNb,omitempty"`

	// Identification used for derivation of a unique key from a master key provided for the data protection.
	DerivationIdentification *Min5Max16Binary `xml:"DerivtnId,omitempty"`
}

func (k *KEKIdentifier2) SetKeyIdentification(value string) {
	k.KeyIdentification = (*Max140Text)(&value)
}

func (k *KEKIdentifier2) SetKeyVersion(value string) {
	k.KeyVersion = (*Max140Text)(&value)
}

func (k *KEKIdentifier2) SetSequenceNumber(value string) {
	k.SequenceNumber = (*Number)(&value)
}

func (k *KEKIdentifier2) SetDerivationIdentification(value string) {
	k.DerivationIdentification = (*Min5Max16Binary)(&value)
}

// Key that must be created and sent in the response, or that must be verified..
type KEKIdentifier3 struct {

	// Name or label of the key.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Identification of the cryptographic key.
	Identification *Max140Text `xml:"Id"`

	// Version of the cryptographic key.
	Version *Max140Text `xml:"Vrsn,omitempty"`

	// Value to check the key, for instance, result of the encryption of the null binary string.
	KeyCheckValue *Max35Binary `xml:"KeyChckVal,omitempty"`
}

func (k *KEKIdentifier3) SetName(value string) {
	k.Name = (*Max140Text)(&value)
}

func (k *KEKIdentifier3) SetIdentification(value string) {
	k.Identification = (*Max140Text)(&value)
}

func (k *KEKIdentifier3) SetVersion(value string) {
	k.Version = (*Max140Text)(&value)
}

func (k *KEKIdentifier3) SetKeyCheckValue(value string) {
	k.KeyCheckValue = (*Max35Binary)(&value)
}

// Cryptographic key involved in the security command.
type KEKIdentifier4 struct {

	// Name or label of the key.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Identification of the cryptographic key.
	KeyIdentification *Max140Text `xml:"KeyId,omitempty"`

	// Version of the cryptographic key.
	KeyVersion *Max140Text `xml:"KeyVrsn,omitempty"`
}

func (k *KEKIdentifier4) SetName(value string) {
	k.Name = (*Max140Text)(&value)
}

func (k *KEKIdentifier4) SetKeyIdentification(value string) {
	k.KeyIdentification = (*Max140Text)(&value)
}

func (k *KEKIdentifier4) SetKeyVersion(value string) {
	k.KeyVersion = (*Max140Text)(&value)
}

// Identification of a key encryption key (KEK), using previously distributed symmetric key.
type KEKIdentifier5 struct {

	// Identification of the cryptographic key.
	KeyIdentification *Max140Text `xml:"KeyId"`

	// Version of the cryptographic key.
	KeyVersion *Max140Text `xml:"KeyVrsn"`

	// Number of usages of the cryptographic key.
	SequenceNumber *Number `xml:"SeqNb,omitempty"`

	// Identification used for derivation of a unique key from a master key provided for the data protection.
	DerivationIdentification *Min5Max16Binary `xml:"DerivtnId,omitempty"`

	// Type of algorithm used by the cryptographic key.
	Type *CryptographicKeyType3Code `xml:"Tp,omitempty"`

	// Allowed usage of the key.
	Function []*KeyUsage1Code `xml:"Fctn,omitempty"`
}

func (k *KEKIdentifier5) SetKeyIdentification(value string) {
	k.KeyIdentification = (*Max140Text)(&value)
}

func (k *KEKIdentifier5) SetKeyVersion(value string) {
	k.KeyVersion = (*Max140Text)(&value)
}

func (k *KEKIdentifier5) SetSequenceNumber(value string) {
	k.SequenceNumber = (*Number)(&value)
}

func (k *KEKIdentifier5) SetDerivationIdentification(value string) {
	k.DerivationIdentification = (*Min5Max16Binary)(&value)
}

func (k *KEKIdentifier5) SetType(value string) {
	k.Type = (*CryptographicKeyType3Code)(&value)
}

func (k *KEKIdentifier5) AddFunction(value string) {
	k.Function = append(k.Function, (*KeyUsage1Code)(&value))
}
