package iso20022

// Identification of a cryptographic algorithm.
type AlgorithmIdentification1 struct {

	// Identification of the algorithm.
	Algorithm *Algorithm1Code `xml:"Algo"`

	// Parameters associated to the algorithm.
	Parameter *Parameter1 `xml:"Param,omitempty"`
}

func (a *AlgorithmIdentification1) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm1Code)(&value)
}

func (a *AlgorithmIdentification1) AddParameter() *Parameter1 {
	a.Parameter = new(Parameter1)
	return a.Parameter
}

// Identification of a cryptographic algorithm and parameters for the MAC computation.
type AlgorithmIdentification10 struct {

	// Identification of the algorithm.
	Algorithm *Algorithm10Code `xml:"Algo"`

	// Parameters associated to the algorithm.
	Parameter *Parameter1 `xml:"Param,omitempty"`
}

func (a *AlgorithmIdentification10) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm10Code)(&value)
}

func (a *AlgorithmIdentification10) AddParameter() *Parameter1 {
	a.Parameter = new(Parameter1)
	return a.Parameter
}

// Cryptographic algorithms and parameters for the protection of transported keys by an asymmetric key.
type AlgorithmIdentification11 struct {

	// Asymmetric encryption algorithm of a transport key.
	Algorithm *Algorithm7Code `xml:"Algo"`

	// Parameters of the encryption algorithm.
	Parameter *Parameter4 `xml:"Param,omitempty"`
}

func (a *AlgorithmIdentification11) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm7Code)(&value)
}

func (a *AlgorithmIdentification11) AddParameter() *Parameter4 {
	a.Parameter = new(Parameter4)
	return a.Parameter
}

// Mask generator function cryptographic algorithm and parameters.
type AlgorithmIdentification12 struct {

	// Mask generator function cryptographic algorithm.
	Algorithm *Algorithm8Code `xml:"Algo"`

	// Parameters associated to the mask generator function cryptographic algorithm
	Parameter *Parameter5 `xml:"Param,omitempty"`
}

func (a *AlgorithmIdentification12) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm8Code)(&value)
}

func (a *AlgorithmIdentification12) AddParameter() *Parameter5 {
	a.Parameter = new(Parameter5)
	return a.Parameter
}

// Cryptographic algorithm and parameters for the protection of the transported key.
type AlgorithmIdentification13 struct {

	// Identification of the algorithm.
	Algorithm *Algorithm13Code `xml:"Algo"`

	// Parameters associated to the encryption algorithm.
	Parameter *Parameter6 `xml:"Param,omitempty"`
}

func (a *AlgorithmIdentification13) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm13Code)(&value)
}

func (a *AlgorithmIdentification13) AddParameter() *Parameter6 {
	a.Parameter = new(Parameter6)
	return a.Parameter
}

// Cryptographic algorithm and parameters for encryptions with a symmetric cryptographic key.
type AlgorithmIdentification14 struct {

	// Identification of the encryption algorithm.
	Algorithm *Algorithm15Code `xml:"Algo"`

	// Parameters associated with the CBC (Chain Block Chaining) encryption algorithm.
	Parameter *Parameter6 `xml:"Param,omitempty"`
}

func (a *AlgorithmIdentification14) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm15Code)(&value)
}

func (a *AlgorithmIdentification14) AddParameter() *Parameter6 {
	a.Parameter = new(Parameter6)
	return a.Parameter
}

// Identification of a cryptographic algorithm and parameters for the MAC computation.
type AlgorithmIdentification15 struct {

	// Identification of the MAC algorithm.
	Algorithm *Algorithm12Code `xml:"Algo"`

	// Parameters associated to the MAC algorithm.
	Parameter *Parameter7 `xml:"Param,omitempty"`
}

func (a *AlgorithmIdentification15) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm12Code)(&value)
}

func (a *AlgorithmIdentification15) AddParameter() *Parameter7 {
	a.Parameter = new(Parameter7)
	return a.Parameter
}

// Cryptographic algorithm and parameters of digests.
type AlgorithmIdentification16 struct {

	// Identification of the digest algorithm.
	Algorithm *Algorithm11Code `xml:"Algo"`
}

func (a *AlgorithmIdentification16) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm11Code)(&value)
}

// Identification of a cryptographic algorithm and parameters for digital signatures.
type AlgorithmIdentification17 struct {

	// Identification of the algorithm.
	Algorithm *Algorithm14Code `xml:"Algo"`

	// Parameters of the RSASSA-PSS digital signature algorithm (RSA signature algorithm with appendix: Probabilistic Signature Scheme).
	Parameter *Parameter8 `xml:"Param,omitempty"`
}

func (a *AlgorithmIdentification17) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm14Code)(&value)
}

func (a *AlgorithmIdentification17) AddParameter() *Parameter8 {
	a.Parameter = new(Parameter8)
	return a.Parameter
}

// Cryptographic algorithm and parameters for the protection of the transported key.
type AlgorithmIdentification2 struct {

	// Identification of the algorithm.
	Algorithm *Algorithm2Code `xml:"Algo"`

	// Parameters associated to the algorithm.
	Parameter *Parameter1 `xml:"Param,omitempty"`
}

func (a *AlgorithmIdentification2) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm2Code)(&value)
}

func (a *AlgorithmIdentification2) AddParameter() *Parameter1 {
	a.Parameter = new(Parameter1)
	return a.Parameter
}

// Identification of a cryptographic algorithm and parameters for the MAC computation.
type AlgorithmIdentification3 struct {

	// Identification of the algorithm.
	Algorithm *Algorithm3Code `xml:"Algo"`

	// Parameters associated to the algorithm.
	Parameter *Parameter1 `xml:"Param,omitempty"`
}

func (a *AlgorithmIdentification3) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm3Code)(&value)
}

func (a *AlgorithmIdentification3) AddParameter() *Parameter1 {
	a.Parameter = new(Parameter1)
	return a.Parameter
}

// Identification of a cryptographic algorithm and parameters for digital signatures.
type AlgorithmIdentification4 struct {

	// Identification of the algorithm.
	Algorithm *Algorithm4Code `xml:"Algo"`
}

func (a *AlgorithmIdentification4) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm4Code)(&value)
}

// Cryptographic algorithm and parameters for digests.
type AlgorithmIdentification5 struct {

	// Identification of the algorithm.
	Algorithm *Algorithm5Code `xml:"Algo"`
}

func (a *AlgorithmIdentification5) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm5Code)(&value)
}

// Cryptographic algorithm and parameters for encryptions with a symmetric cryptographic key.
type AlgorithmIdentification6 struct {

	// Identification of the algorithm.
	Algorithm *Algorithm6Code `xml:"Algo"`

	// Parameters associated with the algorithm.
	Parameter *Parameter1 `xml:"Param,omitempty"`
}

func (a *AlgorithmIdentification6) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm6Code)(&value)
}

func (a *AlgorithmIdentification6) AddParameter() *Parameter1 {
	a.Parameter = new(Parameter1)
	return a.Parameter
}

// Cryptographic algorithms and parameters for the protection of transported keys by an asymmetric key.
type AlgorithmIdentification7 struct {

	// Asymmetric encryption algorithm of a transport key.
	Algorithm *Algorithm7Code `xml:"Algo"`

	// Parameters of the RSAES-OAEP encryption algorithm (RSA Encryption Scheme: Optimal Asymmetric Encryption Padding).
	Parameter *Parameter2 `xml:"Param,omitempty"`
}

func (a *AlgorithmIdentification7) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm7Code)(&value)
}

func (a *AlgorithmIdentification7) AddParameter() *Parameter2 {
	a.Parameter = new(Parameter2)
	return a.Parameter
}

// Mask generator function cryptographic algorithm and parameters.
type AlgorithmIdentification8 struct {

	// Mask generator function cryptographic algorithm.
	Algorithm *Algorithm8Code `xml:"Algo"`

	// Parameters associated to the mask generator function cryptographic algorithm
	Parameter *Parameter3 `xml:"Param,omitempty"`
}

func (a *AlgorithmIdentification8) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm8Code)(&value)
}

func (a *AlgorithmIdentification8) AddParameter() *Parameter3 {
	a.Parameter = new(Parameter3)
	return a.Parameter
}

// Cryptographic algorithm and parameters for the protection of the transported key.
type AlgorithmIdentification9 struct {

	// Identification of the algorithm.
	Algorithm *Algorithm9Code `xml:"Algo"`

	// Parameters associated to the algorithm.
	Parameter *Parameter1 `xml:"Param,omitempty"`
}

func (a *AlgorithmIdentification9) SetAlgorithm(value string) {
	a.Algorithm = (*Algorithm9Code)(&value)
}

func (a *AlgorithmIdentification9) AddParameter() *Parameter1 {
	a.Parameter = new(Parameter1)
	return a.Parameter
}
