package iso20022

// Profile of the customer selected by an ATM.
type ATMCustomerProfile1 struct {

	// Describes the main way customer information was collected to build up the customer menu and the withdrawal request.
	RetrievalMode *ATMCustomerProfile1Code `xml:"RtrvlMd"`

	// Reference of the customer profile.
	ProfileReference *Max35Text `xml:"PrflRef,omitempty"`

	// Identification of the customer for the bank.
	CustomerIdentification *Max35Text `xml:"CstmrId,omitempty"`
}

func (a *ATMCustomerProfile1) SetRetrievalMode(value string) {
	a.RetrievalMode = (*ATMCustomerProfile1Code)(&value)
}

func (a *ATMCustomerProfile1) SetProfileReference(value string) {
	a.ProfileReference = (*Max35Text)(&value)
}

func (a *ATMCustomerProfile1) SetCustomerIdentification(value string) {
	a.CustomerIdentification = (*Max35Text)(&value)
}

// Profile of the customer selected by an ATM.
type ATMCustomerProfile2 struct {

	// Reference of the customer profile.
	ProfileReference *Max35Text `xml:"PrflRef,omitempty"`

	// Identification of the customer for the bank.
	CustomerIdentification *Max35Text `xml:"CstmrId,omitempty"`
}

func (a *ATMCustomerProfile2) SetProfileReference(value string) {
	a.ProfileReference = (*Max35Text)(&value)
}

func (a *ATMCustomerProfile2) SetCustomerIdentification(value string) {
	a.CustomerIdentification = (*Max35Text)(&value)
}

// Profile of the customer with the allowed services and restrictions.
type ATMCustomerProfile3 struct {

	// Reference of the customer profile.
	ProfileReference *Max35Text `xml:"PrflRef,omitempty"`

	// Identification of the customer for the bank.
	CustomerIdentification *Max35Text `xml:"CstmrId,omitempty"`

	// Description of the customer's profile in plaintext.
	ProfileDescription *Max70Text `xml:"PrflDesc,omitempty"`

	// Services allowed for the customer's profile.
	AllowedServices []*ATMService7 `xml:"AllwdSvcs,omitempty"`
}

func (a *ATMCustomerProfile3) SetProfileReference(value string) {
	a.ProfileReference = (*Max35Text)(&value)
}

func (a *ATMCustomerProfile3) SetCustomerIdentification(value string) {
	a.CustomerIdentification = (*Max35Text)(&value)
}

func (a *ATMCustomerProfile3) SetProfileDescription(value string) {
	a.ProfileDescription = (*Max70Text)(&value)
}

func (a *ATMCustomerProfile3) AddAllowedServices() *ATMService7 {
	newValue := new(ATMService7)
	a.AllowedServices = append(a.AllowedServices, newValue)
	return newValue
}

// Profile of the customer selected by an ATM.
type ATMCustomerProfile4 struct {

	// Describes the main way customer information was collected to build up the customer menu and to provide the service.
	RetrievalMode *ATMCustomerProfile1Code `xml:"RtrvlMd"`

	// Reference of the customer profile.
	ProfileReference *Max35Text `xml:"PrflRef,omitempty"`

	// Identification of the customer for the bank.
	CustomerIdentification *Max35Text `xml:"CstmrId,omitempty"`
}

func (a *ATMCustomerProfile4) SetRetrievalMode(value string) {
	a.RetrievalMode = (*ATMCustomerProfile1Code)(&value)
}

func (a *ATMCustomerProfile4) SetProfileReference(value string) {
	a.ProfileReference = (*Max35Text)(&value)
}

func (a *ATMCustomerProfile4) SetCustomerIdentification(value string) {
	a.CustomerIdentification = (*Max35Text)(&value)
}

// Profile of the customer with the allowed services and restrictions.
type ATMCustomerProfile5 struct {

	// Reference of the customer profile.
	ProfileReference *Max35Text `xml:"PrflRef,omitempty"`

	// Identification of the customer for the bank.
	CustomerIdentification *Max35Text `xml:"CstmrId,omitempty"`

	// Description of the customer's profile in plaintext.
	ProfileDescription *Max70Text `xml:"PrflDesc,omitempty"`

	// Services allowed for the customer's profile.
	AllowedServices []*ATMService17 `xml:"AllwdSvcs,omitempty"`
}

func (a *ATMCustomerProfile5) SetProfileReference(value string) {
	a.ProfileReference = (*Max35Text)(&value)
}

func (a *ATMCustomerProfile5) SetCustomerIdentification(value string) {
	a.CustomerIdentification = (*Max35Text)(&value)
}

func (a *ATMCustomerProfile5) SetProfileDescription(value string) {
	a.ProfileDescription = (*Max70Text)(&value)
}

func (a *ATMCustomerProfile5) AddAllowedServices() *ATMService17 {
	newValue := new(ATMService17)
	a.AllowedServices = append(a.AllowedServices, newValue)
	return newValue
}
