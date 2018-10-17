package iso20022

// Key encryption key (KEK), using previously distributed symmetric key.
type KEK1 struct {

	// Version of the cryptographic key.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of the key encryption key (KEK).
	KEKIdentification *KEKIdentifier1 `xml:"KEKId"`

	// Algorithm to encrypt the key encryption key (KEK).
	KeyEncryptionAlgorithm *AlgorithmIdentification1 `xml:"KeyNcrptnAlgo"`

	// Encrypted key encryption key (KEK).
	EncryptedKey *Max140Binary `xml:"NcrptdKey"`
}

func (k *KEK1) SetVersion(value string) {
	k.Version = (*Number)(&value)
}

func (k *KEK1) AddKEKIdentification() *KEKIdentifier1 {
	k.KEKIdentification = new(KEKIdentifier1)
	return k.KEKIdentification
}

func (k *KEK1) AddKeyEncryptionAlgorithm() *AlgorithmIdentification1 {
	k.KeyEncryptionAlgorithm = new(AlgorithmIdentification1)
	return k.KeyEncryptionAlgorithm
}

func (k *KEK1) SetEncryptedKey(value string) {
	k.EncryptedKey = (*Max140Binary)(&value)
}

// Key encryption key (KEK), using previously distributed symmetric key.
type KEK2 struct {

	// Version of the cryptographic key.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of the key encryption key (KEK).
	KEKIdentification *KEKIdentifier1 `xml:"KEKId"`

	// Algorithm to encrypt the key encryption key (KEK).
	KeyEncryptionAlgorithm *AlgorithmIdentification2 `xml:"KeyNcrptnAlgo"`

	// Encrypted key encryption key (KEK).
	EncryptedKey *Max140Binary `xml:"NcrptdKey"`
}

func (k *KEK2) SetVersion(value string) {
	k.Version = (*Number)(&value)
}

func (k *KEK2) AddKEKIdentification() *KEKIdentifier1 {
	k.KEKIdentification = new(KEKIdentifier1)
	return k.KEKIdentification
}

func (k *KEK2) AddKeyEncryptionAlgorithm() *AlgorithmIdentification2 {
	k.KeyEncryptionAlgorithm = new(AlgorithmIdentification2)
	return k.KeyEncryptionAlgorithm
}

func (k *KEK2) SetEncryptedKey(value string) {
	k.EncryptedKey = (*Max140Binary)(&value)
}

// Key encryption key (KEK), using previously distributed symmetric key.
type KEK3 struct {

	// Version of the cryptographic key.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of the key encryption key (KEK).
	KEKIdentification *KEKIdentifier1 `xml:"KEKId"`

	// Algorithm to encrypt the key encryption key (KEK).
	KeyEncryptionAlgorithm *AlgorithmIdentification9 `xml:"KeyNcrptnAlgo"`

	// Encrypted key encryption key (KEK).
	EncryptedKey *Max140Binary `xml:"NcrptdKey"`
}

func (k *KEK3) SetVersion(value string) {
	k.Version = (*Number)(&value)
}

func (k *KEK3) AddKEKIdentification() *KEKIdentifier1 {
	k.KEKIdentification = new(KEKIdentifier1)
	return k.KEKIdentification
}

func (k *KEK3) AddKeyEncryptionAlgorithm() *AlgorithmIdentification9 {
	k.KeyEncryptionAlgorithm = new(AlgorithmIdentification9)
	return k.KeyEncryptionAlgorithm
}

func (k *KEK3) SetEncryptedKey(value string) {
	k.EncryptedKey = (*Max140Binary)(&value)
}

// Key encryption key (KEK), using previously distributed symmetric key.
type KEK4 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of the key encryption key (KEK).
	KEKIdentification *KEKIdentifier2 `xml:"KEKId"`

	// Algorithm to encrypt the key encryption key (KEK).
	KeyEncryptionAlgorithm *AlgorithmIdentification13 `xml:"KeyNcrptnAlgo"`

	// Encrypted key encryption key (KEK).
	EncryptedKey *Max500Binary `xml:"NcrptdKey"`
}

func (k *KEK4) SetVersion(value string) {
	k.Version = (*Number)(&value)
}

func (k *KEK4) AddKEKIdentification() *KEKIdentifier2 {
	k.KEKIdentification = new(KEKIdentifier2)
	return k.KEKIdentification
}

func (k *KEK4) AddKeyEncryptionAlgorithm() *AlgorithmIdentification13 {
	k.KeyEncryptionAlgorithm = new(AlgorithmIdentification13)
	return k.KeyEncryptionAlgorithm
}

func (k *KEK4) SetEncryptedKey(value string) {
	k.EncryptedKey = (*Max500Binary)(&value)
}
