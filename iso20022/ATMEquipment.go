package iso20022

// ATM terminal equipment.
type ATMEquipment1 struct {

	// ATM Manufacturer.
	Manufacturer *Max35Text `xml:"Manfctr,omitempty"`

	// Model of ATM.
	Model *Max35Text `xml:"Mdl,omitempty"`

	// Serial number of the ATM.
	SerialNumber *Max35Text `xml:"SrlNb,omitempty"`

	// Provider of the ATM application software.
	ApplicationProvider *Max35Text `xml:"ApplPrvdr,omitempty"`

	// Name of the software product.
	ApplicationName *Max35Text `xml:"ApplNm,omitempty"`

	// Current version of the software that might include the release number.
	ApplicationVersion *Max35Text `xml:"ApplVrsn,omitempty"`

	// Unique assessment number for the component.
	ApprovalNumber *Max35Text `xml:"ApprvlNb,omitempty"`

	// Configuration parameter version.
	ConfigurationParameter []*ATMConfigurationParameter1 `xml:"CfgtnParam,omitempty"`
}

func (a *ATMEquipment1) SetManufacturer(value string) {
	a.Manufacturer = (*Max35Text)(&value)
}

func (a *ATMEquipment1) SetModel(value string) {
	a.Model = (*Max35Text)(&value)
}

func (a *ATMEquipment1) SetSerialNumber(value string) {
	a.SerialNumber = (*Max35Text)(&value)
}

func (a *ATMEquipment1) SetApplicationProvider(value string) {
	a.ApplicationProvider = (*Max35Text)(&value)
}

func (a *ATMEquipment1) SetApplicationName(value string) {
	a.ApplicationName = (*Max35Text)(&value)
}

func (a *ATMEquipment1) SetApplicationVersion(value string) {
	a.ApplicationVersion = (*Max35Text)(&value)
}

func (a *ATMEquipment1) SetApprovalNumber(value string) {
	a.ApprovalNumber = (*Max35Text)(&value)
}

func (a *ATMEquipment1) AddConfigurationParameter() *ATMConfigurationParameter1 {
	newValue := new(ATMConfigurationParameter1)
	a.ConfigurationParameter = append(a.ConfigurationParameter, newValue)
	return newValue
}

// Hardware security module information, so called EPP for Encrypted PIN Pad.
type ATMEquipment2 struct {

	// ATM Manufacturer.
	Manufacturer *Max35Text `xml:"Manfctr,omitempty"`

	// Model of ATM.
	Model *Max35Text `xml:"Mdl,omitempty"`

	// Version of the device model.
	Version *Max35Text `xml:"Vrsn,omitempty"`

	// Serial number of the ATM.
	SerialNumber *Max35Text `xml:"SrlNb,omitempty"`

	// Provider of the firmware.
	FirmwareProvider *Max35Text `xml:"FrmwrPrvdr,omitempty"`

	// Identification of the firmware.
	FirmwareIdentification *Max35Text `xml:"FrmwrId,omitempty"`

	// Version of the firmware.
	FirmwareVersion *Max35Text `xml:"FrmwrVrsn,omitempty"`
}

func (a *ATMEquipment2) SetManufacturer(value string) {
	a.Manufacturer = (*Max35Text)(&value)
}

func (a *ATMEquipment2) SetModel(value string) {
	a.Model = (*Max35Text)(&value)
}

func (a *ATMEquipment2) SetVersion(value string) {
	a.Version = (*Max35Text)(&value)
}

func (a *ATMEquipment2) SetSerialNumber(value string) {
	a.SerialNumber = (*Max35Text)(&value)
}

func (a *ATMEquipment2) SetFirmwareProvider(value string) {
	a.FirmwareProvider = (*Max35Text)(&value)
}

func (a *ATMEquipment2) SetFirmwareIdentification(value string) {
	a.FirmwareIdentification = (*Max35Text)(&value)
}

func (a *ATMEquipment2) SetFirmwareVersion(value string) {
	a.FirmwareVersion = (*Max35Text)(&value)
}

// Hardware security module information, so called EPP for Encrypted PIN Pad.
type ATMEquipment3 struct {

	// ATM Manufacturer.
	Manufacturer *Max35Text `xml:"Manfctr,omitempty"`

	// Model of ATM.
	Model *Max35Text `xml:"Mdl,omitempty"`

	// Version of the device model.
	Version *Max35Text `xml:"Vrsn,omitempty"`

	// Serial number of the ATM.
	SerialNumber *Max35Text `xml:"SrlNb,omitempty"`

	// Signature of the serial number of the device. The signature may contain the serial number with the signature.
	SignedSerialNumber *ContentInformationType14 `xml:"SgndSrlNb,omitempty"`

	// Provider of the firmware.
	FirmwareProvider *Max35Text `xml:"FrmwrPrvdr,omitempty"`

	// Identification of the firmware.
	FirmwareIdentification *Max35Text `xml:"FrmwrId,omitempty"`

	// Version of the firmware.
	FirmwareVersion *Max35Text `xml:"FrmwrVrsn,omitempty"`
}

func (a *ATMEquipment3) SetManufacturer(value string) {
	a.Manufacturer = (*Max35Text)(&value)
}

func (a *ATMEquipment3) SetModel(value string) {
	a.Model = (*Max35Text)(&value)
}

func (a *ATMEquipment3) SetVersion(value string) {
	a.Version = (*Max35Text)(&value)
}

func (a *ATMEquipment3) SetSerialNumber(value string) {
	a.SerialNumber = (*Max35Text)(&value)
}

func (a *ATMEquipment3) AddSignedSerialNumber() *ContentInformationType14 {
	a.SignedSerialNumber = new(ContentInformationType14)
	return a.SignedSerialNumber
}

func (a *ATMEquipment3) SetFirmwareProvider(value string) {
	a.FirmwareProvider = (*Max35Text)(&value)
}

func (a *ATMEquipment3) SetFirmwareIdentification(value string) {
	a.FirmwareIdentification = (*Max35Text)(&value)
}

func (a *ATMEquipment3) SetFirmwareVersion(value string) {
	a.FirmwareVersion = (*Max35Text)(&value)
}
