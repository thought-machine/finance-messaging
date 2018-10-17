package iso20022

// ATM information.
type AutomatedTellerMachine1 struct {

	// ATM terminal device identification for the acquirer and the issuer.
	Identification *Max35Text `xml:"Id"`

	// ATM terminal device identification for the ATM manager.
	AdditionalIdentification *Max35Text `xml:"AddtlId,omitempty"`

	// ATM terminal device identification for the branch.
	SequenceNumber *Max35Text `xml:"SeqNb,omitempty"`

	// Reference currency of the ATM.
	BaseCurrency *ActiveCurrencyCode `xml:"BaseCcy"`

	// Location of the ATM.
	Location *PostalAddress17 `xml:"Lctn,omitempty"`

	// Indicates the environment of the transaction.
	LocationCategory *TransactionEnvironment2Code `xml:"LctnCtgy,omitempty"`

	// Capabilities of the ATM terminal performing the transaction.
	Capabilities *PointOfInteractionCapabilities5 `xml:"Cpblties,omitempty"`

	// ATM terminal equipment.
	Equipment *ATMEquipment1 `xml:"Eqpmnt,omitempty"`
}

func (a *AutomatedTellerMachine1) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine1) SetAdditionalIdentification(value string) {
	a.AdditionalIdentification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine1) SetSequenceNumber(value string) {
	a.SequenceNumber = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine1) SetBaseCurrency(value string) {
	a.BaseCurrency = (*ActiveCurrencyCode)(&value)
}

func (a *AutomatedTellerMachine1) AddLocation() *PostalAddress17 {
	a.Location = new(PostalAddress17)
	return a.Location
}

func (a *AutomatedTellerMachine1) SetLocationCategory(value string) {
	a.LocationCategory = (*TransactionEnvironment2Code)(&value)
}

func (a *AutomatedTellerMachine1) AddCapabilities() *PointOfInteractionCapabilities5 {
	a.Capabilities = new(PointOfInteractionCapabilities5)
	return a.Capabilities
}

func (a *AutomatedTellerMachine1) AddEquipment() *ATMEquipment1 {
	a.Equipment = new(ATMEquipment1)
	return a.Equipment
}

// ATM information.
type AutomatedTellerMachine10 struct {

	// ATM terminal device identification for the acquirer and the issuer.
	Identification *Max35Text `xml:"Id"`

	// ATM terminal device identification for the ATM manager.
	AdditionalIdentification *Max35Text `xml:"AddtlId,omitempty"`

	// ATM terminal device identification for the branch.
	SequenceNumber *Max35Text `xml:"SeqNb,omitempty"`

	// Reference currency of the ATM.
	BaseCurrency *ActiveCurrencyCode `xml:"BaseCcy"`

	// Location of the ATM.
	Location *PostalAddress17 `xml:"Lctn,omitempty"`

	// Indicates the environment of the transaction.
	LocationCategory *TransactionEnvironment2Code `xml:"LctnCtgy,omitempty"`

	// Capabilities of the ATM terminal performing the transaction.
	Capabilities *PointOfInteractionCapabilities7 `xml:"Cpblties,omitempty"`

	// ATM terminal equipment.
	Equipment *ATMEquipment1 `xml:"Eqpmnt,omitempty"`

	// List of ATM devices out of service.
	AvailableDevice []*ATMDevice2Code `xml:"AvlblDvc,omitempty"`
}

func (a *AutomatedTellerMachine10) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine10) SetAdditionalIdentification(value string) {
	a.AdditionalIdentification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine10) SetSequenceNumber(value string) {
	a.SequenceNumber = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine10) SetBaseCurrency(value string) {
	a.BaseCurrency = (*ActiveCurrencyCode)(&value)
}

func (a *AutomatedTellerMachine10) AddLocation() *PostalAddress17 {
	a.Location = new(PostalAddress17)
	return a.Location
}

func (a *AutomatedTellerMachine10) SetLocationCategory(value string) {
	a.LocationCategory = (*TransactionEnvironment2Code)(&value)
}

func (a *AutomatedTellerMachine10) AddCapabilities() *PointOfInteractionCapabilities7 {
	a.Capabilities = new(PointOfInteractionCapabilities7)
	return a.Capabilities
}

func (a *AutomatedTellerMachine10) AddEquipment() *ATMEquipment1 {
	a.Equipment = new(ATMEquipment1)
	return a.Equipment
}

func (a *AutomatedTellerMachine10) AddAvailableDevice(value string) {
	a.AvailableDevice = append(a.AvailableDevice, (*ATMDevice2Code)(&value))
}

// ATM information.
type AutomatedTellerMachine2 struct {

	// ATM terminal device identification for the acquirer and the issuer.
	Identification *Max35Text `xml:"Id"`

	// ATM terminal device identification for the ATM manager.
	AdditionalIdentification *Max35Text `xml:"AddtlId,omitempty"`

	// ATM terminal device identification for the branch.
	SequenceNumber *Max35Text `xml:"SeqNb,omitempty"`

	// Reference currency of the ATM.
	BaseCurrency *ActiveCurrencyCode `xml:"BaseCcy,omitempty"`

	// Location of the ATM.
	Location *PostalAddress17 `xml:"Lctn,omitempty"`
}

func (a *AutomatedTellerMachine2) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine2) SetAdditionalIdentification(value string) {
	a.AdditionalIdentification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine2) SetSequenceNumber(value string) {
	a.SequenceNumber = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine2) SetBaseCurrency(value string) {
	a.BaseCurrency = (*ActiveCurrencyCode)(&value)
}

func (a *AutomatedTellerMachine2) AddLocation() *PostalAddress17 {
	a.Location = new(PostalAddress17)
	return a.Location
}

// ATM information.
type AutomatedTellerMachine3 struct {

	// ATM terminal device identification for the acquirer and the issuer.
	Identification *Max35Text `xml:"Id"`

	// ATM terminal device identification for the ATM manager.
	AdditionalIdentification *Max35Text `xml:"AddtlId,omitempty"`

	// ATM terminal device identification for the branch.
	SequenceNumber *Max35Text `xml:"SeqNb,omitempty"`

	// Location of the ATM.
	Location *PostalAddress17 `xml:"Lctn,omitempty"`
}

func (a *AutomatedTellerMachine3) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine3) SetAdditionalIdentification(value string) {
	a.AdditionalIdentification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine3) SetSequenceNumber(value string) {
	a.SequenceNumber = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine3) AddLocation() *PostalAddress17 {
	a.Location = new(PostalAddress17)
	return a.Location
}

// ATM information.
type AutomatedTellerMachine4 struct {

	// ATM terminal device identification for the acquirer and the issuer.
	Identification *Max35Text `xml:"Id"`

	// ATM terminal device identification for the ATM manager.
	AdditionalIdentification *Max35Text `xml:"AddtlId,omitempty"`

	// ATM terminal device identification for the branch.
	SequenceNumber *Max35Text `xml:"SeqNb,omitempty"`

	// Reference currency of the ATM.
	BaseCurrency *ActiveCurrencyCode `xml:"BaseCcy"`

	// Location of the ATM.
	Location *PostalAddress17 `xml:"Lctn,omitempty"`

	// Indicates the environment of the transaction.
	LocationCategory *TransactionEnvironment2Code `xml:"LctnCtgy,omitempty"`

	// Capabilities of the ATM terminal performing the transaction.
	Capabilities *PointOfInteractionCapabilities5 `xml:"Cpblties,omitempty"`

	// ATM terminal equipment.
	Equipment *ATMEquipment1 `xml:"Eqpmnt,omitempty"`

	// List of ATM devices out of service.
	AvailableDevice []*ATMDevice2Code `xml:"AvlblDvc,omitempty"`
}

func (a *AutomatedTellerMachine4) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine4) SetAdditionalIdentification(value string) {
	a.AdditionalIdentification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine4) SetSequenceNumber(value string) {
	a.SequenceNumber = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine4) SetBaseCurrency(value string) {
	a.BaseCurrency = (*ActiveCurrencyCode)(&value)
}

func (a *AutomatedTellerMachine4) AddLocation() *PostalAddress17 {
	a.Location = new(PostalAddress17)
	return a.Location
}

func (a *AutomatedTellerMachine4) SetLocationCategory(value string) {
	a.LocationCategory = (*TransactionEnvironment2Code)(&value)
}

func (a *AutomatedTellerMachine4) AddCapabilities() *PointOfInteractionCapabilities5 {
	a.Capabilities = new(PointOfInteractionCapabilities5)
	return a.Capabilities
}

func (a *AutomatedTellerMachine4) AddEquipment() *ATMEquipment1 {
	a.Equipment = new(ATMEquipment1)
	return a.Equipment
}

func (a *AutomatedTellerMachine4) AddAvailableDevice(value string) {
	a.AvailableDevice = append(a.AvailableDevice, (*ATMDevice2Code)(&value))
}

// ATM information.
type AutomatedTellerMachine5 struct {

	// ATM terminal device identification for the acquirer and the issuer.
	Identification *Max35Text `xml:"Id"`

	// ATM terminal device identification for the ATM manager.
	AdditionalIdentification *Max35Text `xml:"AddtlId,omitempty"`

	// ATM terminal device identification for the branch.
	SequenceNumber *Max35Text `xml:"SeqNb,omitempty"`

	// Reference currency of the ATM.
	BaseCurrency *ActiveCurrencyCode `xml:"BaseCcy"`

	// Location of the ATM.
	Location *PostalAddress17 `xml:"Lctn,omitempty"`

	// Indicates the environment of the transaction.
	LocationCategory *TransactionEnvironment2Code `xml:"LctnCtgy,omitempty"`

	// ATM terminal equipment.
	Equipment *ATMEquipment1 `xml:"Eqpmnt,omitempty"`

	// List of ATM devices out of service.
	OutOfServiceDevice []*ATMDevice2Code `xml:"OutOfSvcDvc,omitempty"`

	// Mechanism used to protect the message of the ATM protocol.
	MessageProtection *MessageProtection1Code `xml:"MsgPrtcn,omitempty"`
}

func (a *AutomatedTellerMachine5) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine5) SetAdditionalIdentification(value string) {
	a.AdditionalIdentification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine5) SetSequenceNumber(value string) {
	a.SequenceNumber = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine5) SetBaseCurrency(value string) {
	a.BaseCurrency = (*ActiveCurrencyCode)(&value)
}

func (a *AutomatedTellerMachine5) AddLocation() *PostalAddress17 {
	a.Location = new(PostalAddress17)
	return a.Location
}

func (a *AutomatedTellerMachine5) SetLocationCategory(value string) {
	a.LocationCategory = (*TransactionEnvironment2Code)(&value)
}

func (a *AutomatedTellerMachine5) AddEquipment() *ATMEquipment1 {
	a.Equipment = new(ATMEquipment1)
	return a.Equipment
}

func (a *AutomatedTellerMachine5) AddOutOfServiceDevice(value string) {
	a.OutOfServiceDevice = append(a.OutOfServiceDevice, (*ATMDevice2Code)(&value))
}

func (a *AutomatedTellerMachine5) SetMessageProtection(value string) {
	a.MessageProtection = (*MessageProtection1Code)(&value)
}

// ATM information.
type AutomatedTellerMachine6 struct {

	// ATM terminal device identification for the acquirer and the issuer.
	Identification *Max35Text `xml:"Id"`

	// ATM terminal device identification for the ATM manager.
	AdditionalIdentification *Max35Text `xml:"AddtlId,omitempty"`

	// ATM terminal device identification for the branch.
	SequenceNumber *Max35Text `xml:"SeqNb,omitempty"`

	// Location of the ATM.
	Location *PostalAddress17 `xml:"Lctn,omitempty"`

	// Indicates the environment of the transaction.
	LocationCategory *TransactionEnvironment2Code `xml:"LctnCtgy,omitempty"`

	// ATM terminal equipment.
	Equipment *ATMEquipment1 `xml:"Eqpmnt,omitempty"`
}

func (a *AutomatedTellerMachine6) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine6) SetAdditionalIdentification(value string) {
	a.AdditionalIdentification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine6) SetSequenceNumber(value string) {
	a.SequenceNumber = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine6) AddLocation() *PostalAddress17 {
	a.Location = new(PostalAddress17)
	return a.Location
}

func (a *AutomatedTellerMachine6) SetLocationCategory(value string) {
	a.LocationCategory = (*TransactionEnvironment2Code)(&value)
}

func (a *AutomatedTellerMachine6) AddEquipment() *ATMEquipment1 {
	a.Equipment = new(ATMEquipment1)
	return a.Equipment
}

// ATM information.
type AutomatedTellerMachine7 struct {

	// ATM terminal device identification for the acquirer and the issuer.
	Identification *Max35Text `xml:"Id"`

	// ATM terminal device identification for the ATM manager.
	AdditionalIdentification *Max35Text `xml:"AddtlId,omitempty"`

	// ATM terminal device identification for the branch.
	SequenceNumber *Max35Text `xml:"SeqNb,omitempty"`
}

func (a *AutomatedTellerMachine7) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine7) SetAdditionalIdentification(value string) {
	a.AdditionalIdentification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine7) SetSequenceNumber(value string) {
	a.SequenceNumber = (*Max35Text)(&value)
}

// ATM information.
type AutomatedTellerMachine8 struct {

	// ATM terminal device identification for the acquirer and the issuer.
	Identification *Max35Text `xml:"Id"`

	// ATM terminal device identification for the ATM manager.
	AdditionalIdentification *Max35Text `xml:"AddtlId,omitempty"`

	// ATM terminal device identification for the branch.
	SequenceNumber *Max35Text `xml:"SeqNb,omitempty"`

	// Reference currency of the ATM.
	BaseCurrency *ActiveCurrencyCode `xml:"BaseCcy"`

	// Location of the ATM.
	Location *PostalAddress17 `xml:"Lctn,omitempty"`

	// Indicates the environment of the transaction.
	LocationCategory *TransactionEnvironment2Code `xml:"LctnCtgy,omitempty"`

	// ATM terminal equipment.
	Equipment *ATMEquipment1 `xml:"Eqpmnt,omitempty"`
}

func (a *AutomatedTellerMachine8) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine8) SetAdditionalIdentification(value string) {
	a.AdditionalIdentification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine8) SetSequenceNumber(value string) {
	a.SequenceNumber = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine8) SetBaseCurrency(value string) {
	a.BaseCurrency = (*ActiveCurrencyCode)(&value)
}

func (a *AutomatedTellerMachine8) AddLocation() *PostalAddress17 {
	a.Location = new(PostalAddress17)
	return a.Location
}

func (a *AutomatedTellerMachine8) SetLocationCategory(value string) {
	a.LocationCategory = (*TransactionEnvironment2Code)(&value)
}

func (a *AutomatedTellerMachine8) AddEquipment() *ATMEquipment1 {
	a.Equipment = new(ATMEquipment1)
	return a.Equipment
}

// ATM information.
type AutomatedTellerMachine9 struct {

	// ATM terminal device identification for the acquirer and the issuer.
	Identification *Max35Text `xml:"Id"`

	// ATM terminal device identification for the ATM manager.
	AdditionalIdentification *Max35Text `xml:"AddtlId,omitempty"`

	// ATM terminal device identification for the branch.
	SequenceNumber *Max35Text `xml:"SeqNb,omitempty"`

	// Reference currency of the ATM.
	BaseCurrency *ActiveCurrencyCode `xml:"BaseCcy"`

	// Location of the ATM.
	Location *PostalAddress17 `xml:"Lctn,omitempty"`

	// Indicates the environment of the transaction.
	LocationCategory *TransactionEnvironment2Code `xml:"LctnCtgy,omitempty"`

	// Capabilities of the ATM terminal performing the transaction.
	Capabilities *PointOfInteractionCapabilities7 `xml:"Cpblties,omitempty"`

	// ATM terminal equipment.
	Equipment *ATMEquipment1 `xml:"Eqpmnt,omitempty"`
}

func (a *AutomatedTellerMachine9) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine9) SetAdditionalIdentification(value string) {
	a.AdditionalIdentification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine9) SetSequenceNumber(value string) {
	a.SequenceNumber = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine9) SetBaseCurrency(value string) {
	a.BaseCurrency = (*ActiveCurrencyCode)(&value)
}

func (a *AutomatedTellerMachine9) AddLocation() *PostalAddress17 {
	a.Location = new(PostalAddress17)
	return a.Location
}

func (a *AutomatedTellerMachine9) SetLocationCategory(value string) {
	a.LocationCategory = (*TransactionEnvironment2Code)(&value)
}

func (a *AutomatedTellerMachine9) AddCapabilities() *PointOfInteractionCapabilities7 {
	a.Capabilities = new(PointOfInteractionCapabilities7)
	return a.Capabilities
}

func (a *AutomatedTellerMachine9) AddEquipment() *ATMEquipment1 {
	a.Equipment = new(ATMEquipment1)
	return a.Equipment
}
