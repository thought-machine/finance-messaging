package iso20022

// Configuration parameter version of the ATM.
type ATMConfigurationParameter1 struct {

	// Type of the ATM configuration.
	Type *DataSetCategory7Code `xml:"Tp"`

	// Active version of the configuration.
	Version *Max35Text `xml:"Vrsn"`
}

func (a *ATMConfigurationParameter1) SetType(value string) {
	a.Type = (*DataSetCategory7Code)(&value)
}

func (a *ATMConfigurationParameter1) SetVersion(value string) {
	a.Version = (*Max35Text)(&value)
}

// Parameters to be used by the various cryptographic key commands.
type ATMConfigurationParameter2 struct {

	// Category of the cryptographic key.
	KeyCategory *CryptographicKeyType4Code `xml:"KeyCtgy,omitempty"`

	// Random value from the host provided during a previous exchange.
	HostChallenge *Max140Binary `xml:"HstChllng,omitempty"`

	// Ordered certificate chain of the asymmetric key encryption key, starting with the host certificate.
	Certificate []*Max5000Binary `xml:"Cert,omitempty"`

	// Cryptographic key involved in the security command.
	KeyProperties []*KEKIdentifier4 `xml:"KeyProps,omitempty"`
}

func (a *ATMConfigurationParameter2) SetKeyCategory(value string) {
	a.KeyCategory = (*CryptographicKeyType4Code)(&value)
}

func (a *ATMConfigurationParameter2) SetHostChallenge(value string) {
	a.HostChallenge = (*Max140Binary)(&value)
}

func (a *ATMConfigurationParameter2) AddCertificate(value string) {
	a.Certificate = append(a.Certificate, (*Max5000Binary)(&value))
}

func (a *ATMConfigurationParameter2) AddKeyProperties() *KEKIdentifier4 {
	newValue := new(KEKIdentifier4)
	a.KeyProperties = append(a.KeyProperties, newValue)
	return newValue
}
