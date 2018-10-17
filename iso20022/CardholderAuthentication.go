package iso20022

// Data related to the authentication of the cardholder.
type CardholderAuthentication1 struct {

	// Method used to authenticate a cardholder.
	AuthenticationMethod *AuthenticationMethod1Code `xml:"AuthntcnMtd"`

	// Entity or object in charge of verifying the cardholder authenticity.
	AuthenticationEntity *AuthenticationEntity1Code `xml:"AuthntcnNtty"`

	// Value used to authenticate the cardholder.
	AuthenticationValue *Max40Text `xml:"AuthntcnVal,omitempty"`

	// Encrypted personal identification number (PIN) and related information.
	CardholderOnLinePIN *OnLinePIN1 `xml:"CrdhldrOnLinePIN,omitempty"`

	// Identifies in electronic commerce transactions whether customer authentication is supported and data is available.
	AuthenticationCollectionIndicator *Max35Text `xml:"AuthntcnColltnInd,omitempty"`
}

func (c *CardholderAuthentication1) SetAuthenticationMethod(value string) {
	c.AuthenticationMethod = (*AuthenticationMethod1Code)(&value)
}

func (c *CardholderAuthentication1) SetAuthenticationEntity(value string) {
	c.AuthenticationEntity = (*AuthenticationEntity1Code)(&value)
}

func (c *CardholderAuthentication1) SetAuthenticationValue(value string) {
	c.AuthenticationValue = (*Max40Text)(&value)
}

func (c *CardholderAuthentication1) AddCardholderOnLinePIN() *OnLinePIN1 {
	c.CardholderOnLinePIN = new(OnLinePIN1)
	return c.CardholderOnLinePIN
}

func (c *CardholderAuthentication1) SetAuthenticationCollectionIndicator(value string) {
	c.AuthenticationCollectionIndicator = (*Max35Text)(&value)
}

// Data related to the authentication of the cardholder.
type CardholderAuthentication2 struct {

	// Method used to authenticate the cardholder.
	AuthenticationMethod *AuthenticationMethod1Code `xml:"AuthntcnMtd"`

	// Entity or object in charge of verifying the cardholder authenticity.
	AuthenticationEntity *AuthenticationEntity1Code `xml:"AuthntcnNtty"`
}

func (c *CardholderAuthentication2) SetAuthenticationMethod(value string) {
	c.AuthenticationMethod = (*AuthenticationMethod1Code)(&value)
}

func (c *CardholderAuthentication2) SetAuthenticationEntity(value string) {
	c.AuthenticationEntity = (*AuthenticationEntity1Code)(&value)
}

// Data related to the authentication of the cardholder.
type CardholderAuthentication3 struct {

	// Method used to authenticate a cardholder.
	AuthenticationMethod *AuthenticationMethod2Code `xml:"AuthntcnMtd"`

	// Entity or object in charge of verifying the cardholder authenticity.
	AuthenticationEntity *AuthenticationEntity1Code `xml:"AuthntcnNtty,omitempty"`

	// Value used to authenticate the cardholder.
	AuthenticationValue *Max40Text `xml:"AuthntcnVal,omitempty"`

	// Encrypted personal identification number (PIN) and related information.
	CardholderOnLinePIN *OnLinePIN2 `xml:"CrdhldrOnLinePIN,omitempty"`

	// Identifies in electronic commerce transactions whether customer authentication is supported and data is available.
	AuthenticationCollectionIndicator *Max35Text `xml:"AuthntcnColltnInd,omitempty"`
}

func (c *CardholderAuthentication3) SetAuthenticationMethod(value string) {
	c.AuthenticationMethod = (*AuthenticationMethod2Code)(&value)
}

func (c *CardholderAuthentication3) SetAuthenticationEntity(value string) {
	c.AuthenticationEntity = (*AuthenticationEntity1Code)(&value)
}

func (c *CardholderAuthentication3) SetAuthenticationValue(value string) {
	c.AuthenticationValue = (*Max40Text)(&value)
}

func (c *CardholderAuthentication3) AddCardholderOnLinePIN() *OnLinePIN2 {
	c.CardholderOnLinePIN = new(OnLinePIN2)
	return c.CardholderOnLinePIN
}

func (c *CardholderAuthentication3) SetAuthenticationCollectionIndicator(value string) {
	c.AuthenticationCollectionIndicator = (*Max35Text)(&value)
}

// Data related to the authentication of the cardholder.
type CardholderAuthentication4 struct {

	// Method used to authenticate the cardholder.
	AuthenticationMethod *AuthenticationMethod2Code `xml:"AuthntcnMtd"`

	// Entity or object in charge of verifying the cardholder authenticity.
	AuthenticationEntity *AuthenticationEntity1Code `xml:"AuthntcnNtty"`
}

func (c *CardholderAuthentication4) SetAuthenticationMethod(value string) {
	c.AuthenticationMethod = (*AuthenticationMethod2Code)(&value)
}

func (c *CardholderAuthentication4) SetAuthenticationEntity(value string) {
	c.AuthenticationEntity = (*AuthenticationEntity1Code)(&value)
}

// Data related to the authentication of the cardholder.
type CardholderAuthentication5 struct {

	// Method used to authenticate a cardholder.
	AuthenticationMethod *AuthenticationMethod2Code `xml:"AuthntcnMtd"`

	// Entity or object in charge of verifying the cardholder authenticity.
	AuthenticationEntity *AuthenticationEntity1Code `xml:"AuthntcnNtty,omitempty"`

	// Value used to authenticate the cardholder.
	AuthenticationValue *Max40Text `xml:"AuthntcnVal,omitempty"`

	// Encrypted personal identification number (PIN) and related information.
	CardholderOnLinePIN *OnLinePIN3 `xml:"CrdhldrOnLinePIN,omitempty"`

	// Identifies in electronic commerce transactions whether customer authentication is supported and data is available.
	AuthenticationCollectionIndicator *Max35Text `xml:"AuthntcnColltnInd,omitempty"`
}

func (c *CardholderAuthentication5) SetAuthenticationMethod(value string) {
	c.AuthenticationMethod = (*AuthenticationMethod2Code)(&value)
}

func (c *CardholderAuthentication5) SetAuthenticationEntity(value string) {
	c.AuthenticationEntity = (*AuthenticationEntity1Code)(&value)
}

func (c *CardholderAuthentication5) SetAuthenticationValue(value string) {
	c.AuthenticationValue = (*Max40Text)(&value)
}

func (c *CardholderAuthentication5) AddCardholderOnLinePIN() *OnLinePIN3 {
	c.CardholderOnLinePIN = new(OnLinePIN3)
	return c.CardholderOnLinePIN
}

func (c *CardholderAuthentication5) SetAuthenticationCollectionIndicator(value string) {
	c.AuthenticationCollectionIndicator = (*Max35Text)(&value)
}

// Data related to the authentication of the cardholder.
type CardholderAuthentication6 struct {

	// Method to authenticate the cardholder.
	AuthenticationMethod *AuthenticationMethod3Code `xml:"AuthntcnMtd"`

	// Value used to authenticate the cardholder.
	AuthenticationValue *Max5000Binary `xml:"AuthntcnVal,omitempty"`

	// Protection of the authentication value.
	ProtectedAuthenticationValue *ContentInformationType10 `xml:"PrtctdAuthntcnVal,omitempty"`

	// Encrypted personal identification number (PIN) and related information.
	CardholderOnLinePIN *OnLinePIN4 `xml:"CrdhldrOnLinePIN,omitempty"`

	// Numeric characters of the cardholder's billing or shipping address for verification.
	AddressVerification *AddressVerification1 `xml:"AdrVrfctn,omitempty"`
}

func (c *CardholderAuthentication6) SetAuthenticationMethod(value string) {
	c.AuthenticationMethod = (*AuthenticationMethod3Code)(&value)
}

func (c *CardholderAuthentication6) SetAuthenticationValue(value string) {
	c.AuthenticationValue = (*Max5000Binary)(&value)
}

func (c *CardholderAuthentication6) AddProtectedAuthenticationValue() *ContentInformationType10 {
	c.ProtectedAuthenticationValue = new(ContentInformationType10)
	return c.ProtectedAuthenticationValue
}

func (c *CardholderAuthentication6) AddCardholderOnLinePIN() *OnLinePIN4 {
	c.CardholderOnLinePIN = new(OnLinePIN4)
	return c.CardholderOnLinePIN
}

func (c *CardholderAuthentication6) AddAddressVerification() *AddressVerification1 {
	c.AddressVerification = new(AddressVerification1)
	return c.AddressVerification
}

// Data related to the authentication of the cardholder.
type CardholderAuthentication7 struct {

	// Method and data intended to be used for this transaction to authenticate the cardholder or its card.
	AuthenticationMethod *AuthenticationMethod5Code `xml:"AuthntcnMtd"`

	// Value used to authenticate the cardholder.
	AuthenticationValue *Max5000Binary `xml:"AuthntcnVal,omitempty"`

	// Protection of the authentication value.
	ProtectedAuthenticationValue *ContentInformationType10 `xml:"PrtctdAuthntcnVal,omitempty"`

	// Encrypted personal identification number (PIN) and related information.
	CardholderOnLinePIN *OnLinePIN4 `xml:"CrdhldrOnLinePIN,omitempty"`

	// Identification of the cardholder to verify.
	CardholderIdentification *PersonIdentification7 `xml:"CrdhldrId,omitempty"`

	// Numeric characters of the cardholder's billing or shipping address for verification.
	AddressVerification *AddressVerification1 `xml:"AdrVrfctn,omitempty"`
}

func (c *CardholderAuthentication7) SetAuthenticationMethod(value string) {
	c.AuthenticationMethod = (*AuthenticationMethod5Code)(&value)
}

func (c *CardholderAuthentication7) SetAuthenticationValue(value string) {
	c.AuthenticationValue = (*Max5000Binary)(&value)
}

func (c *CardholderAuthentication7) AddProtectedAuthenticationValue() *ContentInformationType10 {
	c.ProtectedAuthenticationValue = new(ContentInformationType10)
	return c.ProtectedAuthenticationValue
}

func (c *CardholderAuthentication7) AddCardholderOnLinePIN() *OnLinePIN4 {
	c.CardholderOnLinePIN = new(OnLinePIN4)
	return c.CardholderOnLinePIN
}

func (c *CardholderAuthentication7) AddCardholderIdentification() *PersonIdentification7 {
	c.CardholderIdentification = new(PersonIdentification7)
	return c.CardholderIdentification
}

func (c *CardholderAuthentication7) AddAddressVerification() *AddressVerification1 {
	c.AddressVerification = new(AddressVerification1)
	return c.AddressVerification
}

// Data related to the authentication of the card and the cardholder.
type CardholderAuthentication8 struct {

	// Method and data intended to be used for this transaction to authenticate the customer or its card.
	AuthenticationMethod *AuthenticationMethod7Code `xml:"AuthntcnMtd"`

	// True if an authentication token is requested to the host. This token will be provided to the ATM for further authentication.
	TokenRequested *TrueFalseIndicator `xml:"TknReqd,omitempty"`

	// Value or token to be used for customer or card authentication.
	AuthenticationValue *Max5000Binary `xml:"AuthntcnVal,omitempty"`

	// Protection of the authentication value.
	ProtectedAuthenticationValue *ContentInformationType10 `xml:"PrtctdAuthntcnVal,omitempty"`

	// Encrypted personal identification number (PIN) and related information.
	CardholderOnLinePIN *OnLinePIN5 `xml:"CrdhldrOnLinePIN,omitempty"`
}

func (c *CardholderAuthentication8) SetAuthenticationMethod(value string) {
	c.AuthenticationMethod = (*AuthenticationMethod7Code)(&value)
}

func (c *CardholderAuthentication8) SetTokenRequested(value string) {
	c.TokenRequested = (*TrueFalseIndicator)(&value)
}

func (c *CardholderAuthentication8) SetAuthenticationValue(value string) {
	c.AuthenticationValue = (*Max5000Binary)(&value)
}

func (c *CardholderAuthentication8) AddProtectedAuthenticationValue() *ContentInformationType10 {
	c.ProtectedAuthenticationValue = new(ContentInformationType10)
	return c.ProtectedAuthenticationValue
}

func (c *CardholderAuthentication8) AddCardholderOnLinePIN() *OnLinePIN5 {
	c.CardholderOnLinePIN = new(OnLinePIN5)
	return c.CardholderOnLinePIN
}

// Data related to the authentication of the cardholder.
type CardholderAuthentication9 struct {

	// Method and data intended to be used for this transaction to authenticate the cardholder or its card.
	AuthenticationMethod *AuthenticationMethod5Code `xml:"AuthntcnMtd"`

	// Value used to authenticate the cardholder.
	AuthenticationValue *Max5000Binary `xml:"AuthntcnVal,omitempty"`

	// Protection of the authentication value.
	ProtectedAuthenticationValue *ContentInformationType10 `xml:"PrtctdAuthntcnVal,omitempty"`

	// Encrypted personal identification number (PIN) and related information.
	CardholderOnLinePIN *OnLinePIN4 `xml:"CrdhldrOnLinePIN,omitempty"`

	// Identification of the cardholder to verify.
	CardholderIdentification *PersonIdentification11 `xml:"CrdhldrId,omitempty"`

	// Numeric characters of the cardholder's billing or shipping address for verification.
	AddressVerification *AddressVerification1 `xml:"AdrVrfctn,omitempty"`
}

func (c *CardholderAuthentication9) SetAuthenticationMethod(value string) {
	c.AuthenticationMethod = (*AuthenticationMethod5Code)(&value)
}

func (c *CardholderAuthentication9) SetAuthenticationValue(value string) {
	c.AuthenticationValue = (*Max5000Binary)(&value)
}

func (c *CardholderAuthentication9) AddProtectedAuthenticationValue() *ContentInformationType10 {
	c.ProtectedAuthenticationValue = new(ContentInformationType10)
	return c.ProtectedAuthenticationValue
}

func (c *CardholderAuthentication9) AddCardholderOnLinePIN() *OnLinePIN4 {
	c.CardholderOnLinePIN = new(OnLinePIN4)
	return c.CardholderOnLinePIN
}

func (c *CardholderAuthentication9) AddCardholderIdentification() *PersonIdentification11 {
	c.CardholderIdentification = new(PersonIdentification11)
	return c.CardholderIdentification
}

func (c *CardholderAuthentication9) AddAddressVerification() *AddressVerification1 {
	c.AddressVerification = new(AddressVerification1)
	return c.AddressVerification
}
