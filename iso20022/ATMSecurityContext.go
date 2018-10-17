package iso20022

// Information related to security commands.
type ATMSecurityContext1 struct {

	// Key exchange security scheme in activation on the ATM for the host manager.
	CurrentSecurityScheme *ATMSecurityScheme1Code `xml:"CurSctySchme"`

	// Key exchange security schemes implemented in the hardware security module of the ATM.
	SecuritySchemeCapabilities []*ATMSecurityScheme2Code `xml:"SctySchmeCpblties,omitempty"`

	// Hardware security module of the ATM.
	SecurityDevice *ATMSecurityDevice1 `xml:"SctyDvc"`

	// Cryptographic keys stored in the hardware security module of the ATM.
	Key []*CryptographicKey7 `xml:"Key,omitempty"`

	// Random value from the host provided during a previous exchange.
	HostChallenge *Max140Binary `xml:"HstChllng,omitempty"`
}

func (a *ATMSecurityContext1) SetCurrentSecurityScheme(value string) {
	a.CurrentSecurityScheme = (*ATMSecurityScheme1Code)(&value)
}

func (a *ATMSecurityContext1) AddSecuritySchemeCapabilities(value string) {
	a.SecuritySchemeCapabilities = append(a.SecuritySchemeCapabilities, (*ATMSecurityScheme2Code)(&value))
}

func (a *ATMSecurityContext1) AddSecurityDevice() *ATMSecurityDevice1 {
	a.SecurityDevice = new(ATMSecurityDevice1)
	return a.SecurityDevice
}

func (a *ATMSecurityContext1) AddKey() *CryptographicKey7 {
	newValue := new(CryptographicKey7)
	a.Key = append(a.Key, newValue)
	return newValue
}

func (a *ATMSecurityContext1) SetHostChallenge(value string) {
	a.HostChallenge = (*Max140Binary)(&value)
}

// Context of the ATM for the key download.
type ATMSecurityContext2 struct {

	// Key exchange security scheme in activation on the ATM for the host manager.
	CurrentSecurityScheme *ATMSecurityScheme1Code `xml:"CurSctySchme"`

	// Hardware security module information, so called EPP for Encrypted PIN Pad.
	DeviceProperty *ATMEquipment3 `xml:"DvcPrprty,omitempty"`

	// Configuration parameters in use by the security device.
	CurrentConfiguration *ATMSecurityConfiguration1 `xml:"CurCfgtn,omitempty"`
}

func (a *ATMSecurityContext2) SetCurrentSecurityScheme(value string) {
	a.CurrentSecurityScheme = (*ATMSecurityScheme1Code)(&value)
}

func (a *ATMSecurityContext2) AddDeviceProperty() *ATMEquipment3 {
	a.DeviceProperty = new(ATMEquipment3)
	return a.DeviceProperty
}

func (a *ATMSecurityContext2) AddCurrentConfiguration() *ATMSecurityConfiguration1 {
	a.CurrentConfiguration = new(ATMSecurityConfiguration1)
	return a.CurrentConfiguration
}
