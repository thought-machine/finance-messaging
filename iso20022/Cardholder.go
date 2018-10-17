package iso20022

// Data related to the cardholder.
type Cardholder1 struct {

	// Identification of the cardholder involved in a transaction.
	Identification []*CardholderIdentification1 `xml:"Id,omitempty"`

	// Cardholder name associated with the card.
	Name *Max45Text `xml:"Nm,omitempty"`

	// Language selected for the cardholder interface during the transaction.
	Language *ISO2ALanguageCode `xml:"Lang,omitempty"`

	// Data related to the authentication of the cardholder.
	Authentication []*CardholderAuthentication1 `xml:"Authntcn,omitempty"`

	// Numeric characters of the cardholder's address for verification.
	AddressVerification *AddressVerification1 `xml:"AdrVrfctn,omitempty"`

	// Identifies personal data related to the cardholder.
	PersonalData *Max70Text `xml:"PrsnlData,omitempty"`
}

func (c *Cardholder1) AddIdentification() *CardholderIdentification1 {
	newValue := new(CardholderIdentification1)
	c.Identification = append(c.Identification, newValue)
	return newValue
}

func (c *Cardholder1) SetName(value string) {
	c.Name = (*Max45Text)(&value)
}

func (c *Cardholder1) SetLanguage(value string) {
	c.Language = (*ISO2ALanguageCode)(&value)
}

func (c *Cardholder1) AddAuthentication() *CardholderAuthentication1 {
	newValue := new(CardholderAuthentication1)
	c.Authentication = append(c.Authentication, newValue)
	return newValue
}

func (c *Cardholder1) AddAddressVerification() *AddressVerification1 {
	c.AddressVerification = new(AddressVerification1)
	return c.AddressVerification
}

func (c *Cardholder1) SetPersonalData(value string) {
	c.PersonalData = (*Max70Text)(&value)
}

// Data related to the cardholder.
type Cardholder10 struct {

	// Identification of the cardholder involved in a transaction.
	Identification *PersonIdentification11 `xml:"Id,omitempty"`

	// Cardholder name associated with the card.
	Name *Max45Text `xml:"Nm,omitempty"`

	// Language selected for the cardholder interface during the transaction.
	// Reference ISO 639-1 (alpha-2) et ISO 639-2 (alpha-3).
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Postal address of the owner of the payment card.
	BillingAddress *PostalAddress18 `xml:"BllgAdr,omitempty"`

	// Postal address for delivery of goods or services.
	ShippingAddress *PostalAddress18 `xml:"ShppgAdr,omitempty"`

	// Identification of the trip.
	TripNumber *Max35Text `xml:"TripNb,omitempty"`

	// Information related to the vehicle used for the transaction.
	Vehicle *Vehicle1 `xml:"Vhcl,omitempty"`

	// Method and data intended to be used for this transaction to authenticate the cardholder and its card.
	Authentication []*CardholderAuthentication9 `xml:"Authntcn,omitempty"`

	// Result of performed verifications for the transaction.
	TransactionVerificationResult []*TransactionVerificationResult4 `xml:"TxVrfctnRslt,omitempty"`

	// Identifies personal data related to the cardholder.
	PersonalData *Max70Text `xml:"PrsnlData,omitempty"`
}

func (c *Cardholder10) AddIdentification() *PersonIdentification11 {
	c.Identification = new(PersonIdentification11)
	return c.Identification
}

func (c *Cardholder10) SetName(value string) {
	c.Name = (*Max45Text)(&value)
}

func (c *Cardholder10) SetLanguage(value string) {
	c.Language = (*LanguageCode)(&value)
}

func (c *Cardholder10) AddBillingAddress() *PostalAddress18 {
	c.BillingAddress = new(PostalAddress18)
	return c.BillingAddress
}

func (c *Cardholder10) AddShippingAddress() *PostalAddress18 {
	c.ShippingAddress = new(PostalAddress18)
	return c.ShippingAddress
}

func (c *Cardholder10) SetTripNumber(value string) {
	c.TripNumber = (*Max35Text)(&value)
}

func (c *Cardholder10) AddVehicle() *Vehicle1 {
	c.Vehicle = new(Vehicle1)
	return c.Vehicle
}

func (c *Cardholder10) AddAuthentication() *CardholderAuthentication9 {
	newValue := new(CardholderAuthentication9)
	c.Authentication = append(c.Authentication, newValue)
	return newValue
}

func (c *Cardholder10) AddTransactionVerificationResult() *TransactionVerificationResult4 {
	newValue := new(TransactionVerificationResult4)
	c.TransactionVerificationResult = append(c.TransactionVerificationResult, newValue)
	return newValue
}

func (c *Cardholder10) SetPersonalData(value string) {
	c.PersonalData = (*Max70Text)(&value)
}

// Data related to the cardholder.
type Cardholder11 struct {

	// Identification of the cardholder involved in a transaction.
	Identification *PersonIdentification11 `xml:"Id,omitempty"`

	// Cardholder name associated with the card.
	Name *Max45Text `xml:"Nm,omitempty"`

	// Language selected for the cardholder interface during the transaction.
	// Reference ISO 639-1 (alpha-2) et ISO 639-2 (alpha-3).
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Postal address of the owner of the payment card.
	BillingAddress *PostalAddress18 `xml:"BllgAdr,omitempty"`

	// Postal address for delivery of goods or services.
	ShippingAddress *PostalAddress18 `xml:"ShppgAdr,omitempty"`

	// Identification of the trip.
	TripNumber *Max35Text `xml:"TripNb,omitempty"`

	// Information related to the vehicle used for the transaction.
	Vehicle *Vehicle1 `xml:"Vhcl,omitempty"`

	// Identifies personal data related to the cardholder.
	PersonalData *Max70Text `xml:"PrsnlData,omitempty"`
}

func (c *Cardholder11) AddIdentification() *PersonIdentification11 {
	c.Identification = new(PersonIdentification11)
	return c.Identification
}

func (c *Cardholder11) SetName(value string) {
	c.Name = (*Max45Text)(&value)
}

func (c *Cardholder11) SetLanguage(value string) {
	c.Language = (*LanguageCode)(&value)
}

func (c *Cardholder11) AddBillingAddress() *PostalAddress18 {
	c.BillingAddress = new(PostalAddress18)
	return c.BillingAddress
}

func (c *Cardholder11) AddShippingAddress() *PostalAddress18 {
	c.ShippingAddress = new(PostalAddress18)
	return c.ShippingAddress
}

func (c *Cardholder11) SetTripNumber(value string) {
	c.TripNumber = (*Max35Text)(&value)
}

func (c *Cardholder11) AddVehicle() *Vehicle1 {
	c.Vehicle = new(Vehicle1)
	return c.Vehicle
}

func (c *Cardholder11) SetPersonalData(value string) {
	c.PersonalData = (*Max70Text)(&value)
}

// Data related to the cardholder.
type Cardholder2 struct {

	// Identification of the cardholder involved in a transaction.
	Identification []*CardholderIdentification1 `xml:"Id,omitempty"`

	// Cardholder name associated with the card.
	Name *Max45Text `xml:"Nm,omitempty"`

	// Data related to the authentication of the cardholder.
	Authentication []*CardholderAuthentication2 `xml:"Authntcn,omitempty"`

	// Numeric characters of the cardholder's address for verification.
	AddressVerification *AddressVerification1 `xml:"AdrVrfctn,omitempty"`

	// Identifies personal data related to the cardholder.
	PersonalData *Max70Text `xml:"PrsnlData,omitempty"`
}

func (c *Cardholder2) AddIdentification() *CardholderIdentification1 {
	newValue := new(CardholderIdentification1)
	c.Identification = append(c.Identification, newValue)
	return newValue
}

func (c *Cardholder2) SetName(value string) {
	c.Name = (*Max45Text)(&value)
}

func (c *Cardholder2) AddAuthentication() *CardholderAuthentication2 {
	newValue := new(CardholderAuthentication2)
	c.Authentication = append(c.Authentication, newValue)
	return newValue
}

func (c *Cardholder2) AddAddressVerification() *AddressVerification1 {
	c.AddressVerification = new(AddressVerification1)
	return c.AddressVerification
}

func (c *Cardholder2) SetPersonalData(value string) {
	c.PersonalData = (*Max70Text)(&value)
}

// Data related to the cardholder.
type Cardholder3 struct {

	// Identification of the cardholder involved in a transaction.
	Identification []*CardholderIdentification1 `xml:"Id,omitempty"`

	// Cardholder name associated with the card.
	Name *Max45Text `xml:"Nm,omitempty"`

	// Language selected for the cardholder interface during the transaction.
	Language *ISO2ALanguageCode `xml:"Lang,omitempty"`

	// Data related to the authentication of the cardholder.
	Authentication []*CardholderAuthentication3 `xml:"Authntcn,omitempty"`

	// Numeric characters of the cardholder's address for verification.
	AddressVerification *AddressVerification1 `xml:"AdrVrfctn,omitempty"`

	// Identifies personal data related to the cardholder.
	PersonalData *Max70Text `xml:"PrsnlData,omitempty"`
}

func (c *Cardholder3) AddIdentification() *CardholderIdentification1 {
	newValue := new(CardholderIdentification1)
	c.Identification = append(c.Identification, newValue)
	return newValue
}

func (c *Cardholder3) SetName(value string) {
	c.Name = (*Max45Text)(&value)
}

func (c *Cardholder3) SetLanguage(value string) {
	c.Language = (*ISO2ALanguageCode)(&value)
}

func (c *Cardholder3) AddAuthentication() *CardholderAuthentication3 {
	newValue := new(CardholderAuthentication3)
	c.Authentication = append(c.Authentication, newValue)
	return newValue
}

func (c *Cardholder3) AddAddressVerification() *AddressVerification1 {
	c.AddressVerification = new(AddressVerification1)
	return c.AddressVerification
}

func (c *Cardholder3) SetPersonalData(value string) {
	c.PersonalData = (*Max70Text)(&value)
}

// Data related to the cardholder.
type Cardholder4 struct {

	// Identification of the cardholder involved in a transaction.
	Identification []*CardholderIdentification1 `xml:"Id,omitempty"`

	// Cardholder name associated with the card.
	Name *Max45Text `xml:"Nm,omitempty"`

	// Data related to the authentication of the cardholder.
	Authentication []*CardholderAuthentication4 `xml:"Authntcn,omitempty"`

	// Numeric characters of the cardholder's address for verification.
	AddressVerification *AddressVerification1 `xml:"AdrVrfctn,omitempty"`

	// Identifies personal data related to the cardholder.
	PersonalData *Max70Text `xml:"PrsnlData,omitempty"`
}

func (c *Cardholder4) AddIdentification() *CardholderIdentification1 {
	newValue := new(CardholderIdentification1)
	c.Identification = append(c.Identification, newValue)
	return newValue
}

func (c *Cardholder4) SetName(value string) {
	c.Name = (*Max45Text)(&value)
}

func (c *Cardholder4) AddAuthentication() *CardholderAuthentication4 {
	newValue := new(CardholderAuthentication4)
	c.Authentication = append(c.Authentication, newValue)
	return newValue
}

func (c *Cardholder4) AddAddressVerification() *AddressVerification1 {
	c.AddressVerification = new(AddressVerification1)
	return c.AddressVerification
}

func (c *Cardholder4) SetPersonalData(value string) {
	c.PersonalData = (*Max70Text)(&value)
}

// Data related to the cardholder.
type Cardholder5 struct {

	// Identification of the cardholder involved in a transaction.
	Identification *PersonIdentification7 `xml:"Id,omitempty"`

	// Cardholder name associated with the card.
	Name *Max45Text `xml:"Nm,omitempty"`

	// Language selected for the cardholder interface during the transaction.
	Language *ISO2ALanguageCode `xml:"Lang,omitempty"`

	// Postal address of the owner of the payment card.
	BillingAddress *PostalAddress13 `xml:"BllgAdr,omitempty"`

	// Postal address for delivery of goods or services.
	ShippingAddress *PostalAddress13 `xml:"ShppgAdr,omitempty"`

	// Data related to the authentication of the cardholder.
	Authentication []*CardholderAuthentication5 `xml:"Authntcn,omitempty"`

	// Numeric characters of the cardholder's address for verification.
	AddressVerification *AddressVerification1 `xml:"AdrVrfctn,omitempty"`

	// Identifies personal data related to the cardholder.
	PersonalData *Max70Text `xml:"PrsnlData,omitempty"`
}

func (c *Cardholder5) AddIdentification() *PersonIdentification7 {
	c.Identification = new(PersonIdentification7)
	return c.Identification
}

func (c *Cardholder5) SetName(value string) {
	c.Name = (*Max45Text)(&value)
}

func (c *Cardholder5) SetLanguage(value string) {
	c.Language = (*ISO2ALanguageCode)(&value)
}

func (c *Cardholder5) AddBillingAddress() *PostalAddress13 {
	c.BillingAddress = new(PostalAddress13)
	return c.BillingAddress
}

func (c *Cardholder5) AddShippingAddress() *PostalAddress13 {
	c.ShippingAddress = new(PostalAddress13)
	return c.ShippingAddress
}

func (c *Cardholder5) AddAuthentication() *CardholderAuthentication5 {
	newValue := new(CardholderAuthentication5)
	c.Authentication = append(c.Authentication, newValue)
	return newValue
}

func (c *Cardholder5) AddAddressVerification() *AddressVerification1 {
	c.AddressVerification = new(AddressVerification1)
	return c.AddressVerification
}

func (c *Cardholder5) SetPersonalData(value string) {
	c.PersonalData = (*Max70Text)(&value)
}

// Data related to the cardholder.
type Cardholder6 struct {

	// Identification of the cardholder involved in a transaction.
	Identification []*PersonIdentification7 `xml:"Id,omitempty"`

	// Cardholder name associated with the card.
	Name *Max45Text `xml:"Nm,omitempty"`

	// Data related to the authentication of the cardholder.
	Authentication []*CardholderAuthentication4 `xml:"Authntcn,omitempty"`

	// Numeric characters of the cardholder's address for verification.
	AddressVerification *AddressVerification1 `xml:"AdrVrfctn,omitempty"`

	// Identifies personal data related to the cardholder.
	PersonalData *Max70Text `xml:"PrsnlData,omitempty"`
}

func (c *Cardholder6) AddIdentification() *PersonIdentification7 {
	newValue := new(PersonIdentification7)
	c.Identification = append(c.Identification, newValue)
	return newValue
}

func (c *Cardholder6) SetName(value string) {
	c.Name = (*Max45Text)(&value)
}

func (c *Cardholder6) AddAuthentication() *CardholderAuthentication4 {
	newValue := new(CardholderAuthentication4)
	c.Authentication = append(c.Authentication, newValue)
	return newValue
}

func (c *Cardholder6) AddAddressVerification() *AddressVerification1 {
	c.AddressVerification = new(AddressVerification1)
	return c.AddressVerification
}

func (c *Cardholder6) SetPersonalData(value string) {
	c.PersonalData = (*Max70Text)(&value)
}

// Data related to the cardholder.
type Cardholder7 struct {

	// Identification of the cardholder involved in a transaction.
	Identification *PersonIdentification7 `xml:"Id,omitempty"`

	// Cardholder name associated with the card.
	Name *Max45Text `xml:"Nm,omitempty"`

	// Language selected for the cardholder interface during the transaction.
	// Reference ISO 639-1 (alpha-2) et ISO 639-2 (alpha-3).
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Postal address of the owner of the payment card.
	BillingAddress *PostalAddress13 `xml:"BllgAdr,omitempty"`

	// Postal address for delivery of goods or services.
	ShippingAddress *PostalAddress13 `xml:"ShppgAdr,omitempty"`

	// Method and data intended to be used for this transaction to authenticate the cardholder.
	Authentication []*CardholderAuthentication6 `xml:"Authntcn,omitempty"`

	// Result of performed verifications for the transaction.
	TransactionVerificationResult []*TransactionVerificationResult3 `xml:"TxVrfctnRslt,omitempty"`

	// Identifies personal data related to the cardholder.
	PersonalData *Max70Text `xml:"PrsnlData,omitempty"`
}

func (c *Cardholder7) AddIdentification() *PersonIdentification7 {
	c.Identification = new(PersonIdentification7)
	return c.Identification
}

func (c *Cardholder7) SetName(value string) {
	c.Name = (*Max45Text)(&value)
}

func (c *Cardholder7) SetLanguage(value string) {
	c.Language = (*LanguageCode)(&value)
}

func (c *Cardholder7) AddBillingAddress() *PostalAddress13 {
	c.BillingAddress = new(PostalAddress13)
	return c.BillingAddress
}

func (c *Cardholder7) AddShippingAddress() *PostalAddress13 {
	c.ShippingAddress = new(PostalAddress13)
	return c.ShippingAddress
}

func (c *Cardholder7) AddAuthentication() *CardholderAuthentication6 {
	newValue := new(CardholderAuthentication6)
	c.Authentication = append(c.Authentication, newValue)
	return newValue
}

func (c *Cardholder7) AddTransactionVerificationResult() *TransactionVerificationResult3 {
	newValue := new(TransactionVerificationResult3)
	c.TransactionVerificationResult = append(c.TransactionVerificationResult, newValue)
	return newValue
}

func (c *Cardholder7) SetPersonalData(value string) {
	c.PersonalData = (*Max70Text)(&value)
}

// Data related to the cardholder.
type Cardholder8 struct {

	// Identification of the cardholder involved in a transaction.
	Identification *PersonIdentification7 `xml:"Id,omitempty"`

	// Cardholder name associated with the card.
	Name *Max45Text `xml:"Nm,omitempty"`

	// Identifies personal data related to the cardholder.
	PersonalData *Max70Text `xml:"PrsnlData,omitempty"`
}

func (c *Cardholder8) AddIdentification() *PersonIdentification7 {
	c.Identification = new(PersonIdentification7)
	return c.Identification
}

func (c *Cardholder8) SetName(value string) {
	c.Name = (*Max45Text)(&value)
}

func (c *Cardholder8) SetPersonalData(value string) {
	c.PersonalData = (*Max70Text)(&value)
}

// Data related to the cardholder.
type Cardholder9 struct {

	// Identification of the cardholder involved in a transaction.
	Identification *PersonIdentification7 `xml:"Id,omitempty"`

	// Cardholder name associated with the card.
	Name *Max45Text `xml:"Nm,omitempty"`

	// Language selected for the cardholder interface during the transaction.
	// Reference ISO 639-1 (alpha-2) et ISO 639-2 (alpha-3).
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Postal address of the owner of the payment card.
	BillingAddress *PostalAddress18 `xml:"BllgAdr,omitempty"`

	// Postal address for delivery of goods or services.
	ShippingAddress *PostalAddress18 `xml:"ShppgAdr,omitempty"`

	// Method and data intended to be used for this transaction to authenticate the cardholder and its card.
	Authentication []*CardholderAuthentication7 `xml:"Authntcn,omitempty"`

	// Result of performed verifications for the transaction.
	TransactionVerificationResult []*TransactionVerificationResult4 `xml:"TxVrfctnRslt,omitempty"`

	// Identifies personal data related to the cardholder.
	PersonalData *Max70Text `xml:"PrsnlData,omitempty"`
}

func (c *Cardholder9) AddIdentification() *PersonIdentification7 {
	c.Identification = new(PersonIdentification7)
	return c.Identification
}

func (c *Cardholder9) SetName(value string) {
	c.Name = (*Max45Text)(&value)
}

func (c *Cardholder9) SetLanguage(value string) {
	c.Language = (*LanguageCode)(&value)
}

func (c *Cardholder9) AddBillingAddress() *PostalAddress18 {
	c.BillingAddress = new(PostalAddress18)
	return c.BillingAddress
}

func (c *Cardholder9) AddShippingAddress() *PostalAddress18 {
	c.ShippingAddress = new(PostalAddress18)
	return c.ShippingAddress
}

func (c *Cardholder9) AddAuthentication() *CardholderAuthentication7 {
	newValue := new(CardholderAuthentication7)
	c.Authentication = append(c.Authentication, newValue)
	return newValue
}

func (c *Cardholder9) AddTransactionVerificationResult() *TransactionVerificationResult4 {
	newValue := new(TransactionVerificationResult4)
	c.TransactionVerificationResult = append(c.TransactionVerificationResult, newValue)
	return newValue
}

func (c *Cardholder9) SetPersonalData(value string) {
	c.PersonalData = (*Max70Text)(&value)
}
