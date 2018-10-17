package iso20022

// Point of interaction parameters related to the security of software application and application protocol.
type SecurityParameters1 struct {

	// Point of interaction challenge for cryptographic key injection.
	POIChallenge *Max140Binary `xml:"POIChllng,omitempty"`

	// Terminal manager challenge for cryptographic key injection.
	TMChallenge *Max140Binary `xml:"TMChllng,omitempty"`

	// Key to inject in the point of interaction, protected by the temporary key previously sent.
	SymmetricKey []*CryptographicKey2 `xml:"SmmtrcKey,omitempty"`
}

func (s *SecurityParameters1) SetPOIChallenge(value string) {
	s.POIChallenge = (*Max140Binary)(&value)
}

func (s *SecurityParameters1) SetTMChallenge(value string) {
	s.TMChallenge = (*Max140Binary)(&value)
}

func (s *SecurityParameters1) AddSymmetricKey() *CryptographicKey2 {
	newValue := new(CryptographicKey2)
	s.SymmetricKey = append(s.SymmetricKey, newValue)
	return newValue
}

// Point of interaction parameters related to the security of software application and application protocol.
type SecurityParameters2 struct {

	// Point of interaction challenge for cryptographic key injection.
	POIChallenge *Max140Binary `xml:"POIChllng,omitempty"`

	// Terminal manager challenge for cryptographic key injection.
	TMChallenge *Max140Binary `xml:"TMChllng,omitempty"`

	// Key to inject in the point of interaction, protected by the temporary key previously sent.
	SymmetricKey []*CryptographicKey4 `xml:"SmmtrcKey,omitempty"`
}

func (s *SecurityParameters2) SetPOIChallenge(value string) {
	s.POIChallenge = (*Max140Binary)(&value)
}

func (s *SecurityParameters2) SetTMChallenge(value string) {
	s.TMChallenge = (*Max140Binary)(&value)
}

func (s *SecurityParameters2) AddSymmetricKey() *CryptographicKey4 {
	newValue := new(CryptographicKey4)
	s.SymmetricKey = append(s.SymmetricKey, newValue)
	return newValue
}

// Parameters related to the security of software application and application protocol.
type SecurityParameters3 struct {

	// Version of the security parameters.
	Version *Max256Text `xml:"Vrsn"`

	// Point of interaction challenge for cryptographic key injection.
	POIChallenge *Max140Binary `xml:"POIChllng,omitempty"`

	// Terminal manager challenge for cryptographic key injection.
	TMChallenge *Max140Binary `xml:"TMChllng,omitempty"`

	// Key to inject in the point of interaction, protected by the temporary key previously sent.
	SymmetricKey []*CryptographicKey5 `xml:"SmmtrcKey,omitempty"`
}

func (s *SecurityParameters3) SetVersion(value string) {
	s.Version = (*Max256Text)(&value)
}

func (s *SecurityParameters3) SetPOIChallenge(value string) {
	s.POIChallenge = (*Max140Binary)(&value)
}

func (s *SecurityParameters3) SetTMChallenge(value string) {
	s.TMChallenge = (*Max140Binary)(&value)
}

func (s *SecurityParameters3) AddSymmetricKey() *CryptographicKey5 {
	newValue := new(CryptographicKey5)
	s.SymmetricKey = append(s.SymmetricKey, newValue)
	return newValue
}

// Security parameters of the ATM for the initiated key download.
type SecurityParameters4 struct {

	// Cryptographic key used to protect the key download.
	Key *CryptographicKey8 `xml:"Key,omitempty"`

	// Digital signature of implicit data depending on the security scheme download procedure.
	DigitalSignature *ContentInformationType14 `xml:"DgtlSgntr,omitempty"`

	// Ordered certificate chain of the asymmetric key encryption key, starting with the ATM certificate.
	Certificate []*Max5000Binary `xml:"Cert,omitempty"`

	// Random value from the ATM to avoid message replay.
	ATMChallenge *Max140Binary `xml:"ATMChllng,omitempty"`

	// Requested  key for downloading, depending on the key hierarchy used by the host.
	RequestedKey *Max35Text `xml:"ReqdKey,omitempty"`
}

func (s *SecurityParameters4) AddKey() *CryptographicKey8 {
	s.Key = new(CryptographicKey8)
	return s.Key
}

func (s *SecurityParameters4) AddDigitalSignature() *ContentInformationType14 {
	s.DigitalSignature = new(ContentInformationType14)
	return s.DigitalSignature
}

func (s *SecurityParameters4) AddCertificate(value string) {
	s.Certificate = append(s.Certificate, (*Max5000Binary)(&value))
}

func (s *SecurityParameters4) SetATMChallenge(value string) {
	s.ATMChallenge = (*Max140Binary)(&value)
}

func (s *SecurityParameters4) SetRequestedKey(value string) {
	s.RequestedKey = (*Max35Text)(&value)
}

// Security parameters of the host downloading the key.
type SecurityParameters5 struct {

	// Random value from the host.
	HostChallenge *Max140Binary `xml:"HstChllng,omitempty"`

	// Cryptographic key used to store in the ATM.
	Key []*CryptographicKey8 `xml:"Key,omitempty"`

	// Digital signature of implicit data depending on the security scheme download procedure.
	DigitalSignature *ContentInformationType14 `xml:"DgtlSgntr,omitempty"`
}

func (s *SecurityParameters5) SetHostChallenge(value string) {
	s.HostChallenge = (*Max140Binary)(&value)
}

func (s *SecurityParameters5) AddKey() *CryptographicKey8 {
	newValue := new(CryptographicKey8)
	s.Key = append(s.Key, newValue)
	return newValue
}

func (s *SecurityParameters5) AddDigitalSignature() *ContentInformationType14 {
	s.DigitalSignature = new(ContentInformationType14)
	return s.DigitalSignature
}

// Parameters related to the security of software application and application protocol.
type SecurityParameters6 struct {

	// Type of action for the configuration parameters.
	ActionType *TerminalManagementAction3Code `xml:"ActnTp"`

	// Version of the security parameters.
	Version *Max256Text `xml:"Vrsn"`

	// Point of interaction challenge for cryptographic key injection.
	POIChallenge *Max140Binary `xml:"POIChllng,omitempty"`

	// Terminal manager challenge for cryptographic key injection.
	TMChallenge *Max140Binary `xml:"TMChllng,omitempty"`

	// Key to inject in the point of interaction, protected by the temporary key previously sent.
	SymmetricKey []*CryptographicKey5 `xml:"SmmtrcKey,omitempty"`
}

func (s *SecurityParameters6) SetActionType(value string) {
	s.ActionType = (*TerminalManagementAction3Code)(&value)
}

func (s *SecurityParameters6) SetVersion(value string) {
	s.Version = (*Max256Text)(&value)
}

func (s *SecurityParameters6) SetPOIChallenge(value string) {
	s.POIChallenge = (*Max140Binary)(&value)
}

func (s *SecurityParameters6) SetTMChallenge(value string) {
	s.TMChallenge = (*Max140Binary)(&value)
}

func (s *SecurityParameters6) AddSymmetricKey() *CryptographicKey5 {
	newValue := new(CryptographicKey5)
	s.SymmetricKey = append(s.SymmetricKey, newValue)
	return newValue
}
