package iso20022

// Encrypted data with an encryption key.
type EncryptedContent1 struct {

	// Type of data which have been encrypted.
	ContentType *ContentType1Code `xml:"CnttTp"`

	// Algorithm used to encrypt the data.
	ContentEncryptionAlgorithm *AlgorithmIdentification1 `xml:"CnttNcrptnAlgo"`

	// Encrypted data, result of the content encryption.
	EncryptedData *Max10000Binary `xml:"NcrptdData"`
}

func (e *EncryptedContent1) SetContentType(value string) {
	e.ContentType = (*ContentType1Code)(&value)
}

func (e *EncryptedContent1) AddContentEncryptionAlgorithm() *AlgorithmIdentification1 {
	e.ContentEncryptionAlgorithm = new(AlgorithmIdentification1)
	return e.ContentEncryptionAlgorithm
}

func (e *EncryptedContent1) SetEncryptedData(value string) {
	e.EncryptedData = (*Max10000Binary)(&value)
}

// Encrypted data with an encryption key.
type EncryptedContent2 struct {

	// Type of data which have been encrypted.
	ContentType *ContentType1Code `xml:"CnttTp"`

	// Algorithm used to encrypt the data.
	ContentEncryptionAlgorithm *AlgorithmIdentification6 `xml:"CnttNcrptnAlgo"`

	// Encrypted data, result of the content encryption.
	EncryptedData *Max10000Binary `xml:"NcrptdData"`
}

func (e *EncryptedContent2) SetContentType(value string) {
	e.ContentType = (*ContentType1Code)(&value)
}

func (e *EncryptedContent2) AddContentEncryptionAlgorithm() *AlgorithmIdentification6 {
	e.ContentEncryptionAlgorithm = new(AlgorithmIdentification6)
	return e.ContentEncryptionAlgorithm
}

func (e *EncryptedContent2) SetEncryptedData(value string) {
	e.EncryptedData = (*Max10000Binary)(&value)
}

// Encrypted data with an encryption key.
type EncryptedContent3 struct {

	// Type of data which have been encrypted.
	ContentType *ContentType2Code `xml:"CnttTp"`

	// Algorithm used to encrypt the data.
	ContentEncryptionAlgorithm *AlgorithmIdentification14 `xml:"CnttNcrptnAlgo"`

	// Encrypted data, result of the content encryption.
	EncryptedData *Max100KBinary `xml:"NcrptdData"`
}

func (e *EncryptedContent3) SetContentType(value string) {
	e.ContentType = (*ContentType2Code)(&value)
}

func (e *EncryptedContent3) AddContentEncryptionAlgorithm() *AlgorithmIdentification14 {
	e.ContentEncryptionAlgorithm = new(AlgorithmIdentification14)
	return e.ContentEncryptionAlgorithm
}

func (e *EncryptedContent3) SetEncryptedData(value string) {
	e.EncryptedData = (*Max100KBinary)(&value)
}
