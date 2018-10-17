package iso20022

// Key encryption key (KEK), encrypted previously distributed symmetric key.
type KeyTransport1 struct {

	// Version of the cryptographic key.
	Version *Number `xml:"Vrsn"`

	// Transport key or key encryption key (KEK) for the recipient.
	RecipientIdentification *CertificateIdentifier1 `xml:"RcptId"`

	// Algorithm to encrypt the key encryption key (KEK).
	KeyEncryptionAlgorithm *AlgorithmIdentification1 `xml:"KeyNcrptnAlgo"`

	// Encrypted key encryption key (KEK).
	EncryptedKey *Max140Binary `xml:"NcrptdKey"`
}

func (k *KeyTransport1) SetVersion(value string) {
	k.Version = (*Number)(&value)
}

func (k *KeyTransport1) AddRecipientIdentification() *CertificateIdentifier1 {
	k.RecipientIdentification = new(CertificateIdentifier1)
	return k.RecipientIdentification
}

func (k *KeyTransport1) AddKeyEncryptionAlgorithm() *AlgorithmIdentification1 {
	k.KeyEncryptionAlgorithm = new(AlgorithmIdentification1)
	return k.KeyEncryptionAlgorithm
}

func (k *KeyTransport1) SetEncryptedKey(value string) {
	k.EncryptedKey = (*Max140Binary)(&value)
}

// Key encryption key (KEK), encrypted with a previously distributed asymmetric public key.
type KeyTransport2 struct {

	// Version of the cryptographic key.
	Version *Number `xml:"Vrsn"`

	// Transport key or key encryption key (KEK) for the recipient.
	RecipientIdentification *CertificateIdentifier1 `xml:"RcptId"`

	// Algorithm to encrypt the key encryption key (KEK).
	KeyEncryptionAlgorithm *AlgorithmIdentification7 `xml:"KeyNcrptnAlgo"`

	// Encrypted key encryption key (KEK).
	EncryptedKey *Max140Binary `xml:"NcrptdKey"`
}

func (k *KeyTransport2) SetVersion(value string) {
	k.Version = (*Number)(&value)
}

func (k *KeyTransport2) AddRecipientIdentification() *CertificateIdentifier1 {
	k.RecipientIdentification = new(CertificateIdentifier1)
	return k.RecipientIdentification
}

func (k *KeyTransport2) AddKeyEncryptionAlgorithm() *AlgorithmIdentification7 {
	k.KeyEncryptionAlgorithm = new(AlgorithmIdentification7)
	return k.KeyEncryptionAlgorithm
}

func (k *KeyTransport2) SetEncryptedKey(value string) {
	k.EncryptedKey = (*Max140Binary)(&value)
}

// Key encryption key (KEK), encrypted with a previously distributed asymmetric public key.
type KeyTransport3 struct {

	// Version of the cryptographic key.
	Version *Number `xml:"Vrsn,omitempty"`

	// Transport key or key encryption key (KEK) for the recipient.
	RecipientIdentification *CertificateIdentifier1 `xml:"RcptId"`

	// Algorithm to encrypt the key encryption key (KEK).
	KeyEncryptionAlgorithm *AlgorithmIdentification7 `xml:"KeyNcrptnAlgo"`

	// Encrypted key encryption key (KEK).
	EncryptedKey *Max3000Binary `xml:"NcrptdKey"`
}

func (k *KeyTransport3) SetVersion(value string) {
	k.Version = (*Number)(&value)
}

func (k *KeyTransport3) AddRecipientIdentification() *CertificateIdentifier1 {
	k.RecipientIdentification = new(CertificateIdentifier1)
	return k.RecipientIdentification
}

func (k *KeyTransport3) AddKeyEncryptionAlgorithm() *AlgorithmIdentification7 {
	k.KeyEncryptionAlgorithm = new(AlgorithmIdentification7)
	return k.KeyEncryptionAlgorithm
}

func (k *KeyTransport3) SetEncryptedKey(value string) {
	k.EncryptedKey = (*Max3000Binary)(&value)
}

// Key encryption key (KEK), encrypted with a previously distributed asymmetric public key.
type KeyTransport4 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of a cryptographic asymmetric key for the recipient.
	RecipientIdentification *Recipient5Choice `xml:"RcptId"`

	// Algorithm to encrypt the key encryption key (KEK).
	KeyEncryptionAlgorithm *AlgorithmIdentification11 `xml:"KeyNcrptnAlgo"`

	// Encrypted key encryption key (KEK).
	EncryptedKey *Max5000Binary `xml:"NcrptdKey"`
}

func (k *KeyTransport4) SetVersion(value string) {
	k.Version = (*Number)(&value)
}

func (k *KeyTransport4) AddRecipientIdentification() *Recipient5Choice {
	k.RecipientIdentification = new(Recipient5Choice)
	return k.RecipientIdentification
}

func (k *KeyTransport4) AddKeyEncryptionAlgorithm() *AlgorithmIdentification11 {
	k.KeyEncryptionAlgorithm = new(AlgorithmIdentification11)
	return k.KeyEncryptionAlgorithm
}

func (k *KeyTransport4) SetEncryptedKey(value string) {
	k.EncryptedKey = (*Max5000Binary)(&value)
}
