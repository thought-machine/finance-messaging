package iso20022

// Digital signature of data, with an asymmetric key.
type SignedData1 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of a digest algorithm to apply before signature.
	DigestAlgorithm []*AlgorithmIdentification1 `xml:"DgstAlgo"`

	// Data to sign.
	EncapsulatedContent *EncapsulatedContent1 `xml:"NcpsltdCntt"`

	// Chain of X.509 certificates.
	Certificate []*Max3000Binary `xml:"Cert,omitempty"`

	// Entity who has signed the data.
	Signer []*Signer1 `xml:"Sgnr"`
}

func (s *SignedData1) SetVersion(value string) {
	s.Version = (*Number)(&value)
}

func (s *SignedData1) AddDigestAlgorithm() *AlgorithmIdentification1 {
	newValue := new(AlgorithmIdentification1)
	s.DigestAlgorithm = append(s.DigestAlgorithm, newValue)
	return newValue
}

func (s *SignedData1) AddEncapsulatedContent() *EncapsulatedContent1 {
	s.EncapsulatedContent = new(EncapsulatedContent1)
	return s.EncapsulatedContent
}

func (s *SignedData1) AddCertificate(value string) {
	s.Certificate = append(s.Certificate, (*Max3000Binary)(&value))
}

func (s *SignedData1) AddSigner() *Signer1 {
	newValue := new(Signer1)
	s.Signer = append(s.Signer, newValue)
	return newValue
}

// Digital signature of data, with an asymmetric key.
type SignedData2 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of a digest algorithm to apply before signature.
	DigestAlgorithm []*AlgorithmIdentification5 `xml:"DgstAlgo"`

	// Data to sign.
	EncapsulatedContent *EncapsulatedContent1 `xml:"NcpsltdCntt"`

	// Chain of X.509 certificates.
	Certificate []*Max3000Binary `xml:"Cert,omitempty"`

	// Entity who has signed the data.
	Signer []*Signer2 `xml:"Sgnr"`
}

func (s *SignedData2) SetVersion(value string) {
	s.Version = (*Number)(&value)
}

func (s *SignedData2) AddDigestAlgorithm() *AlgorithmIdentification5 {
	newValue := new(AlgorithmIdentification5)
	s.DigestAlgorithm = append(s.DigestAlgorithm, newValue)
	return newValue
}

func (s *SignedData2) AddEncapsulatedContent() *EncapsulatedContent1 {
	s.EncapsulatedContent = new(EncapsulatedContent1)
	return s.EncapsulatedContent
}

func (s *SignedData2) AddCertificate(value string) {
	s.Certificate = append(s.Certificate, (*Max3000Binary)(&value))
}

func (s *SignedData2) AddSigner() *Signer2 {
	newValue := new(Signer2)
	s.Signer = append(s.Signer, newValue)
	return newValue
}

// Digital signature of data, with an asymmetric key.
type SignedData3 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of a digest algorithm to apply before signature.
	DigestAlgorithm []*AlgorithmIdentification5 `xml:"DgstAlgo"`

	// Data to sign.
	EncapsulatedContent *EncapsulatedContent2 `xml:"NcpsltdCntt"`

	// Chain of X.509 certificates.
	Certificate []*Max3000Binary `xml:"Cert,omitempty"`

	// Entity who has signed the data.
	Signer []*Signer2 `xml:"Sgnr"`
}

func (s *SignedData3) SetVersion(value string) {
	s.Version = (*Number)(&value)
}

func (s *SignedData3) AddDigestAlgorithm() *AlgorithmIdentification5 {
	newValue := new(AlgorithmIdentification5)
	s.DigestAlgorithm = append(s.DigestAlgorithm, newValue)
	return newValue
}

func (s *SignedData3) AddEncapsulatedContent() *EncapsulatedContent2 {
	s.EncapsulatedContent = new(EncapsulatedContent2)
	return s.EncapsulatedContent
}

func (s *SignedData3) AddCertificate(value string) {
	s.Certificate = append(s.Certificate, (*Max3000Binary)(&value))
}

func (s *SignedData3) AddSigner() *Signer2 {
	newValue := new(Signer2)
	s.Signer = append(s.Signer, newValue)
	return newValue
}

// Digital signatures of data from one or several signers.
type SignedData4 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of digest algorithm applied before signature.
	DigestAlgorithm []*AlgorithmIdentification16 `xml:"DgstAlgo"`

	// Data to sign.
	EncapsulatedContent *EncapsulatedContent3 `xml:"NcpsltdCntt"`

	// Chain of X.509 certificates.
	Certificate []*Max5000Binary `xml:"Cert,omitempty"`

	// Digital signature and identification of a signer.
	Signer []*Signer3 `xml:"Sgnr"`
}

func (s *SignedData4) SetVersion(value string) {
	s.Version = (*Number)(&value)
}

func (s *SignedData4) AddDigestAlgorithm() *AlgorithmIdentification16 {
	newValue := new(AlgorithmIdentification16)
	s.DigestAlgorithm = append(s.DigestAlgorithm, newValue)
	return newValue
}

func (s *SignedData4) AddEncapsulatedContent() *EncapsulatedContent3 {
	s.EncapsulatedContent = new(EncapsulatedContent3)
	return s.EncapsulatedContent
}

func (s *SignedData4) AddCertificate(value string) {
	s.Certificate = append(s.Certificate, (*Max5000Binary)(&value))
}

func (s *SignedData4) AddSigner() *Signer3 {
	newValue := new(Signer3)
	s.Signer = append(s.Signer, newValue)
	return newValue
}
