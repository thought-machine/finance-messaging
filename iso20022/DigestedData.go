package iso20022

// Digest computed on the identified data.
type DigestedData1 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of a digest algorithm.
	DigestAlgorithm []*AlgorithmIdentification1 `xml:"DgstAlgo"`

	// Data on which the digest is computed.
	EncapsulatedContent *EncapsulatedContent1 `xml:"NcpsltdCntt"`

	// Result of data-digesting process.
	Digest *Max140Text `xml:"Dgst"`
}

func (d *DigestedData1) SetVersion(value string) {
	d.Version = (*Number)(&value)
}

func (d *DigestedData1) AddDigestAlgorithm() *AlgorithmIdentification1 {
	newValue := new(AlgorithmIdentification1)
	d.DigestAlgorithm = append(d.DigestAlgorithm, newValue)
	return newValue
}

func (d *DigestedData1) AddEncapsulatedContent() *EncapsulatedContent1 {
	d.EncapsulatedContent = new(EncapsulatedContent1)
	return d.EncapsulatedContent
}

func (d *DigestedData1) SetDigest(value string) {
	d.Digest = (*Max140Text)(&value)
}

// Digest computed on the identified data.
type DigestedData2 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of a digest algorithm.
	DigestAlgorithm []*AlgorithmIdentification5 `xml:"DgstAlgo"`

	// Data on which the digest is computed.
	EncapsulatedContent *EncapsulatedContent1 `xml:"NcpsltdCntt"`

	// Result of data-digesting process.
	Digest *Max140Text `xml:"Dgst"`
}

func (d *DigestedData2) SetVersion(value string) {
	d.Version = (*Number)(&value)
}

func (d *DigestedData2) AddDigestAlgorithm() *AlgorithmIdentification5 {
	newValue := new(AlgorithmIdentification5)
	d.DigestAlgorithm = append(d.DigestAlgorithm, newValue)
	return newValue
}

func (d *DigestedData2) AddEncapsulatedContent() *EncapsulatedContent1 {
	d.EncapsulatedContent = new(EncapsulatedContent1)
	return d.EncapsulatedContent
}

func (d *DigestedData2) SetDigest(value string) {
	d.Digest = (*Max140Text)(&value)
}

// Digest computed on the identified data.
type DigestedData3 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of a digest algorithm.
	DigestAlgorithm []*AlgorithmIdentification5 `xml:"DgstAlgo"`

	// Data on which the digest is computed.
	EncapsulatedContent *EncapsulatedContent2 `xml:"NcpsltdCntt"`

	// Result of data-digesting process.
	Digest *Max140Text `xml:"Dgst"`
}

func (d *DigestedData3) SetVersion(value string) {
	d.Version = (*Number)(&value)
}

func (d *DigestedData3) AddDigestAlgorithm() *AlgorithmIdentification5 {
	newValue := new(AlgorithmIdentification5)
	d.DigestAlgorithm = append(d.DigestAlgorithm, newValue)
	return newValue
}

func (d *DigestedData3) AddEncapsulatedContent() *EncapsulatedContent2 {
	d.EncapsulatedContent = new(EncapsulatedContent2)
	return d.EncapsulatedContent
}

func (d *DigestedData3) SetDigest(value string) {
	d.Digest = (*Max140Text)(&value)
}

// Digest computed on the identified data.
type DigestedData4 struct {

	// Version of the data structure.
	Version *Number `xml:"Vrsn,omitempty"`

	// Identification of the digest algorithm.
	DigestAlgorithm *AlgorithmIdentification16 `xml:"DgstAlgo"`

	// Data on which the digest is computed.
	EncapsulatedContent *EncapsulatedContent3 `xml:"NcpsltdCntt"`

	// Result of data-digesting process.
	Digest *Max140Binary `xml:"Dgst"`
}

func (d *DigestedData4) SetVersion(value string) {
	d.Version = (*Number)(&value)
}

func (d *DigestedData4) AddDigestAlgorithm() *AlgorithmIdentification16 {
	d.DigestAlgorithm = new(AlgorithmIdentification16)
	return d.DigestAlgorithm
}

func (d *DigestedData4) AddEncapsulatedContent() *EncapsulatedContent3 {
	d.EncapsulatedContent = new(EncapsulatedContent3)
	return d.EncapsulatedContent
}

func (d *DigestedData4) SetDigest(value string) {
	d.Digest = (*Max140Binary)(&value)
}
