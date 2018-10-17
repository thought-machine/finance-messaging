package iso20022

// Parameters associated to a cryptographic algorithm.
type Parameter1 struct {

	// Initialisation vector of a cipher block chaining (CBC) mode encryption.
	InitialisationVector *Max500Binary `xml:"InitlstnVctr,omitempty"`
}

func (p *Parameter1) SetInitialisationVector(value string) {
	p.InitialisationVector = (*Max500Binary)(&value)
}

// Parameters of the RSAES-OAEP encryption algorithm (RSA Encryption Scheme: Optimal Asymmetric Encryption Padding).
type Parameter2 struct {

	// Digest algorithm used in the RSAES-OAEP encryption algorithm (RSA Encryption Scheme: Optimal Asymmetric Encryption Padding).
	DigestAlgorithm *Algorithm5Code `xml:"DgstAlgo,omitempty"`

	// Mask generator function cryptographic algorithm and parameters.
	MaskGeneratorAlgorithm *AlgorithmIdentification8 `xml:"MskGnrtrAlgo,omitempty"`
}

func (p *Parameter2) SetDigestAlgorithm(value string) {
	p.DigestAlgorithm = (*Algorithm5Code)(&value)
}

func (p *Parameter2) AddMaskGeneratorAlgorithm() *AlgorithmIdentification8 {
	p.MaskGeneratorAlgorithm = new(AlgorithmIdentification8)
	return p.MaskGeneratorAlgorithm
}

// Parameters associated to a mask generator cryptographic function.
type Parameter3 struct {

	// Digest algorithm used in the mask generator function.
	DigestAlgorithm *Algorithm5Code `xml:"DgstAlgo,omitempty"`
}

func (p *Parameter3) SetDigestAlgorithm(value string) {
	p.DigestAlgorithm = (*Algorithm5Code)(&value)
}

// Parameters of the asymmetric encryption algorithm.
type Parameter4 struct {

	// Format of data before encryption, if the format is not plaintext or implicit.
	EncryptionFormat *EncryptionFormat1Code `xml:"NcrptnFrmt,omitempty"`

	// Identification of the digest algorithm.
	DigestAlgorithm *Algorithm11Code `xml:"DgstAlgo,omitempty"`

	// Mask generator function cryptographic algorithm and parameters.
	MaskGeneratorAlgorithm *AlgorithmIdentification12 `xml:"MskGnrtrAlgo,omitempty"`
}

func (p *Parameter4) SetEncryptionFormat(value string) {
	p.EncryptionFormat = (*EncryptionFormat1Code)(&value)
}

func (p *Parameter4) SetDigestAlgorithm(value string) {
	p.DigestAlgorithm = (*Algorithm11Code)(&value)
}

func (p *Parameter4) AddMaskGeneratorAlgorithm() *AlgorithmIdentification12 {
	p.MaskGeneratorAlgorithm = new(AlgorithmIdentification12)
	return p.MaskGeneratorAlgorithm
}

// Parameters associated to a mask generator cryptographic function.
type Parameter5 struct {

	// Digest algorithm used in the mask generator function.
	DigestAlgorithm *Algorithm11Code `xml:"DgstAlgo,omitempty"`
}

func (p *Parameter5) SetDigestAlgorithm(value string) {
	p.DigestAlgorithm = (*Algorithm11Code)(&value)
}

// Parameters associated to a cryptographic encryption algorithm.
type Parameter6 struct {

	// Format of data before encryption, if the format is not plaintext or implicit.
	EncryptionFormat *EncryptionFormat1Code `xml:"NcrptnFrmt,omitempty"`

	// Initialisation vector of a cipher block chaining (CBC) mode encryption.
	InitialisationVector *Max500Binary `xml:"InitlstnVctr,omitempty"`

	// Byte padding for a cypher block chaining mode encryption, if the padding is not implicit.
	BytePadding *BytePadding1Code `xml:"BPddg,omitempty"`
}

func (p *Parameter6) SetEncryptionFormat(value string) {
	p.EncryptionFormat = (*EncryptionFormat1Code)(&value)
}

func (p *Parameter6) SetInitialisationVector(value string) {
	p.InitialisationVector = (*Max500Binary)(&value)
}

func (p *Parameter6) SetBytePadding(value string) {
	p.BytePadding = (*BytePadding1Code)(&value)
}

// Parameters associated to the MAC algorithm.
type Parameter7 struct {

	// Initialisation vector of a cipher block chaining (CBC) mode encryption.
	InitialisationVector *Max500Binary `xml:"InitlstnVctr,omitempty"`

	// Byte padding for a cypher block chaining mode encryption, if the padding is not implicit.
	BytePadding *BytePadding1Code `xml:"BPddg,omitempty"`
}

func (p *Parameter7) SetInitialisationVector(value string) {
	p.InitialisationVector = (*Max500Binary)(&value)
}

func (p *Parameter7) SetBytePadding(value string) {
	p.BytePadding = (*BytePadding1Code)(&value)
}

// Parameters of the RSASSA-PSS digital signature algorithm (RSA signature algorithm with appendix: Probabilistic Signature Scheme).
type Parameter8 struct {

	// Identification of the digest algorithm.
	DigestAlgorithm *Algorithm11Code `xml:"DgstAlgo"`

	// Mask generator function cryptographic algorithm and parameters.
	MaskGeneratorAlgorithm *AlgorithmIdentification12 `xml:"MskGnrtrAlgo"`

	// Length of the salt to include in the signature.
	SaltLength *Number `xml:"SaltLngth"`

	// Trailer field number.
	TrailerField *Number `xml:"TrlrFld,omitempty"`
}

func (p *Parameter8) SetDigestAlgorithm(value string) {
	p.DigestAlgorithm = (*Algorithm11Code)(&value)
}

func (p *Parameter8) AddMaskGeneratorAlgorithm() *AlgorithmIdentification12 {
	p.MaskGeneratorAlgorithm = new(AlgorithmIdentification12)
	return p.MaskGeneratorAlgorithm
}

func (p *Parameter8) SetSaltLength(value string) {
	p.SaltLength = (*Number)(&value)
}

func (p *Parameter8) SetTrailerField(value string) {
	p.TrailerField = (*Number)(&value)
}
