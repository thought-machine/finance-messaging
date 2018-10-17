package iso20022

// Entity who has signed the data and its digital signature.
type Signer1 struct {

	// Version of the Cryptographic Message Syntax (CMS) data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of the entity who has signed the data.
	SignerIdentification *CertificateIdentifier1 `xml:"SgnrId"`

	// Identification of a digest algorithm to apply before signature.
	DigestAlgorithm *AlgorithmIdentification1 `xml:"DgstAlgo"`

	// Cryptographic digital signature algorithm.
	SignatureAlgorithm *AlgorithmIdentification1 `xml:"SgntrAlgo"`

	// Digital signature.
	Signature *Max500Binary `xml:"Sgntr"`
}

func (s *Signer1) SetVersion(value string) {
	s.Version = (*Number)(&value)
}

func (s *Signer1) AddSignerIdentification() *CertificateIdentifier1 {
	s.SignerIdentification = new(CertificateIdentifier1)
	return s.SignerIdentification
}

func (s *Signer1) AddDigestAlgorithm() *AlgorithmIdentification1 {
	s.DigestAlgorithm = new(AlgorithmIdentification1)
	return s.DigestAlgorithm
}

func (s *Signer1) AddSignatureAlgorithm() *AlgorithmIdentification1 {
	s.SignatureAlgorithm = new(AlgorithmIdentification1)
	return s.SignatureAlgorithm
}

func (s *Signer1) SetSignature(value string) {
	s.Signature = (*Max500Binary)(&value)
}

// Entity who has signed the data and its digital signature.
type Signer2 struct {

	// Version of the Cryptographic Message Syntax (CMS) data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of the entity who has signed the data.
	SignerIdentification *CertificateIdentifier1 `xml:"SgnrId"`

	// Identification of a digest algorithm to apply before signature.
	DigestAlgorithm *AlgorithmIdentification5 `xml:"DgstAlgo"`

	// Cryptographic digital signature algorithm.
	SignatureAlgorithm *AlgorithmIdentification4 `xml:"SgntrAlgo"`

	// Digital signature.
	Signature *Max500Binary `xml:"Sgntr"`
}

func (s *Signer2) SetVersion(value string) {
	s.Version = (*Number)(&value)
}

func (s *Signer2) AddSignerIdentification() *CertificateIdentifier1 {
	s.SignerIdentification = new(CertificateIdentifier1)
	return s.SignerIdentification
}

func (s *Signer2) AddDigestAlgorithm() *AlgorithmIdentification5 {
	s.DigestAlgorithm = new(AlgorithmIdentification5)
	return s.DigestAlgorithm
}

func (s *Signer2) AddSignatureAlgorithm() *AlgorithmIdentification4 {
	s.SignatureAlgorithm = new(AlgorithmIdentification4)
	return s.SignatureAlgorithm
}

func (s *Signer2) SetSignature(value string) {
	s.Signature = (*Max500Binary)(&value)
}

// Entity who has signed the data and its digital signature.
type Signer3 struct {

	// Version of the Cryptographic Message Syntax (CMS) data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of the entity who has signed the data.
	SignerIdentification *Recipient5Choice `xml:"SgnrId,omitempty"`

	// Identification of a digest algorithm to apply before signature.
	DigestAlgorithm *AlgorithmIdentification16 `xml:"DgstAlgo"`

	// Cryptographic digital signature algorithm.
	SignatureAlgorithm *AlgorithmIdentification17 `xml:"SgntrAlgo"`

	// Digital signature.
	Signature *Max3000Binary `xml:"Sgntr"`
}

func (s *Signer3) SetVersion(value string) {
	s.Version = (*Number)(&value)
}

func (s *Signer3) AddSignerIdentification() *Recipient5Choice {
	s.SignerIdentification = new(Recipient5Choice)
	return s.SignerIdentification
}

func (s *Signer3) AddDigestAlgorithm() *AlgorithmIdentification16 {
	s.DigestAlgorithm = new(AlgorithmIdentification16)
	return s.DigestAlgorithm
}

func (s *Signer3) AddSignatureAlgorithm() *AlgorithmIdentification17 {
	s.SignatureAlgorithm = new(AlgorithmIdentification17)
	return s.SignatureAlgorithm
}

func (s *Signer3) SetSignature(value string) {
	s.Signature = (*Max3000Binary)(&value)
}
