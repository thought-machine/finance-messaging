package iso20022

// Acquirer involved in the card payment.
type Acquirer1 struct {

	// Identification of the acquirer (for example the bank identification number BIN).
	Identification *GenericIdentification32 `xml:"Id,omitempty"`

	// Version of the payment acquirer parameters of the POI.
	ParametersVersion *ISODateTime `xml:"ParamsVrsn"`
}

func (a *Acquirer1) AddIdentification() *GenericIdentification32 {
	a.Identification = new(GenericIdentification32)
	return a.Identification
}

func (a *Acquirer1) SetParametersVersion(value string) {
	a.ParametersVersion = (*ISODateTime)(&value)
}

// Acquirer involved in the card payment.
type Acquirer2 struct {

	// Identification of the acquirer (for example the bank identification number BIN).
	Identification *GenericIdentification32 `xml:"Id,omitempty"`

	// Version of the payment acquirer parameters of the POI.
	ParametersVersion *Max35Text `xml:"ParamsVrsn"`
}

func (a *Acquirer2) AddIdentification() *GenericIdentification32 {
	a.Identification = new(GenericIdentification32)
	return a.Identification
}

func (a *Acquirer2) SetParametersVersion(value string) {
	a.ParametersVersion = (*Max35Text)(&value)
}

// Acquirer involved in the card payment.
type Acquirer3 struct {

	// Identification of the acquirer (for example the bank identification number BIN).
	Identification *GenericIdentification32 `xml:"Id"`

	// Version of the payment acquirer parameters of the POI.
	ParametersVersion *Max35Text `xml:"ParamsVrsn,omitempty"`
}

func (a *Acquirer3) AddIdentification() *GenericIdentification32 {
	a.Identification = new(GenericIdentification32)
	return a.Identification
}

func (a *Acquirer3) SetParametersVersion(value string) {
	a.ParametersVersion = (*Max35Text)(&value)
}

// Acquirer involved in the card payment.
type Acquirer4 struct {

	// Identification of the acquirer (for example the bank identification number BIN).
	Identification *GenericIdentification53 `xml:"Id,omitempty"`

	// Version of the payment acquirer parameters of the POI.
	ParametersVersion *Max256Text `xml:"ParamsVrsn"`
}

func (a *Acquirer4) AddIdentification() *GenericIdentification53 {
	a.Identification = new(GenericIdentification53)
	return a.Identification
}

func (a *Acquirer4) SetParametersVersion(value string) {
	a.ParametersVersion = (*Max256Text)(&value)
}

// Acquirer involved in the card payment.
type Acquirer5 struct {

	// Identification of the acquirer (for example the bank identification number BIN).
	Identification *GenericIdentification53 `xml:"Id"`

	// Version of the payment acquirer parameters of the POI.
	ParametersVersion *Max256Text `xml:"ParamsVrsn,omitempty"`
}

func (a *Acquirer5) AddIdentification() *GenericIdentification53 {
	a.Identification = new(GenericIdentification53)
	return a.Identification
}

func (a *Acquirer5) SetParametersVersion(value string) {
	a.ParametersVersion = (*Max256Text)(&value)
}

// Acquirer of the card transaction.
type Acquirer6 struct {

	// Identification of the acquirer.
	// It correspond to the ISO 8583 field number 32.
	Identification *Max35Text `xml:"Id"`

	// Identification of the entity assigning the acquirer identification.
	Issuer *Max35Text `xml:"Issr,omitempty"`

	// Country of the acquirer.
	// It correspond to the ISO 8583 field number 19.
	CountryCode *ISO3NumericCountryCode `xml:"CtryCd,omitempty"`
}

func (a *Acquirer6) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *Acquirer6) SetIssuer(value string) {
	a.Issuer = (*Max35Text)(&value)
}

func (a *Acquirer6) SetCountryCode(value string) {
	a.CountryCode = (*ISO3NumericCountryCode)(&value)
}

// Acquirer of the withdrawal transaction, in charge of the funds settlement with the issuer.
type Acquirer7 struct {

	// Identification of the acquirer.
	AcquiringInstitution *Max35Text `xml:"AcqrgInstn,omitempty"`

	// Identification of the acquirer branch.
	Branch *Max35Text `xml:"Brnch,omitempty"`
}

func (a *Acquirer7) SetAcquiringInstitution(value string) {
	a.AcquiringInstitution = (*Max35Text)(&value)
}

func (a *Acquirer7) SetBranch(value string) {
	a.Branch = (*Max35Text)(&value)
}

// Institution in charge of managing the ATM.
type Acquirer8 struct {

	// Identification of the ATM manager.
	Identification *Max35Text `xml:"Id"`

	// Software version of the application.
	ApplicationVersion *Max35Text `xml:"ApplVrsn,omitempty"`
}

func (a *Acquirer8) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *Acquirer8) SetApplicationVersion(value string) {
	a.ApplicationVersion = (*Max35Text)(&value)
}
