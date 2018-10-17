package iso20022

// Configuration parameters of the TMS protocol between a POI and a terminal manager.
type TMSProtocolParameters1 struct {

	// Identification of the master terminal manager or the terminal manager.
	TerminalManagerIdentification *GenericIdentification71 `xml:"TermnlMgrId"`

	// Maintenance services provided by the terminal manager.
	MaintenanceService []*DataSetCategory5Code `xml:"MntncSvc"`

	// Version of the TMS protocol parameters.
	Version *Max256Text `xml:"Vrsn"`

	// Identification of applications which may be managed by the TM, partially or globally.
	ApplicationIdentification []*Max35Text `xml:"ApplId,omitempty"`

	// Addresses of the terminal manager host.
	HostAddress *NetworkParameters3 `xml:"HstAdr,omitempty"`

	// Cryptographic key used to communicate with the terminal manager host.
	HostKey []*KEKIdentifier2 `xml:"HstKey,omitempty"`

	// New identification of the POI for the terminal manager.
	POIIdentification *Max35Text `xml:"POIId,omitempty"`

	// New identification of the initiating party to set in TMS messages with this terminal manager.
	InitiatingPartyIdentification *Max35Text `xml:"InitgPtyId,omitempty"`

	// New identification of the recipient party to set in TMS messages with this terminal manager
	RecipientPartyIdentification *Max35Text `xml:"RcptPtyId,omitempty"`
}

func (t *TMSProtocolParameters1) AddTerminalManagerIdentification() *GenericIdentification71 {
	t.TerminalManagerIdentification = new(GenericIdentification71)
	return t.TerminalManagerIdentification
}

func (t *TMSProtocolParameters1) AddMaintenanceService(value string) {
	t.MaintenanceService = append(t.MaintenanceService, (*DataSetCategory5Code)(&value))
}

func (t *TMSProtocolParameters1) SetVersion(value string) {
	t.Version = (*Max256Text)(&value)
}

func (t *TMSProtocolParameters1) AddApplicationIdentification(value string) {
	t.ApplicationIdentification = append(t.ApplicationIdentification, (*Max35Text)(&value))
}

func (t *TMSProtocolParameters1) AddHostAddress() *NetworkParameters3 {
	t.HostAddress = new(NetworkParameters3)
	return t.HostAddress
}

func (t *TMSProtocolParameters1) AddHostKey() *KEKIdentifier2 {
	newValue := new(KEKIdentifier2)
	t.HostKey = append(t.HostKey, newValue)
	return newValue
}

func (t *TMSProtocolParameters1) SetPOIIdentification(value string) {
	t.POIIdentification = (*Max35Text)(&value)
}

func (t *TMSProtocolParameters1) SetInitiatingPartyIdentification(value string) {
	t.InitiatingPartyIdentification = (*Max35Text)(&value)
}

func (t *TMSProtocolParameters1) SetRecipientPartyIdentification(value string) {
	t.RecipientPartyIdentification = (*Max35Text)(&value)
}

// Configuration parameters of the TMS protocol between a POI and a terminal manager.
type TMSProtocolParameters2 struct {

	// Type of action for the configuration parameters.
	ActionType *TerminalManagementAction3Code `xml:"ActnTp"`

	// Identification of the master terminal manager or the terminal manager.
	TerminalManagerIdentification *GenericIdentification71 `xml:"TermnlMgrId"`

	// Maintenance services provided by the terminal manager.
	MaintenanceService []*DataSetCategory10Code `xml:"MntncSvc"`

	// Version of the TMS protocol parameters.
	Version *Max256Text `xml:"Vrsn"`

	// Identification of applications which may be managed by the TM, partially or globally.
	ApplicationIdentification []*Max35Text `xml:"ApplId,omitempty"`

	// Identification of the terminal manager host.
	HostIdentification *Max35Text `xml:"HstId"`

	// New identification of the POI for the terminal manager.
	POIIdentification *Max35Text `xml:"POIId,omitempty"`

	// New identification of the initiating party to set in TMS messages with this terminal manager.
	InitiatingPartyIdentification *Max35Text `xml:"InitgPtyId,omitempty"`

	// New identification of the recipient party to set in TMS messages with this terminal manager
	RecipientPartyIdentification *Max35Text `xml:"RcptPtyId,omitempty"`

	// Configuration parameters are exchanged per file transfer protocol rather than per message.
	FileTransfer *TrueFalseIndicator `xml:"FileTrf,omitempty"`
}

func (t *TMSProtocolParameters2) SetActionType(value string) {
	t.ActionType = (*TerminalManagementAction3Code)(&value)
}

func (t *TMSProtocolParameters2) AddTerminalManagerIdentification() *GenericIdentification71 {
	t.TerminalManagerIdentification = new(GenericIdentification71)
	return t.TerminalManagerIdentification
}

func (t *TMSProtocolParameters2) AddMaintenanceService(value string) {
	t.MaintenanceService = append(t.MaintenanceService, (*DataSetCategory10Code)(&value))
}

func (t *TMSProtocolParameters2) SetVersion(value string) {
	t.Version = (*Max256Text)(&value)
}

func (t *TMSProtocolParameters2) AddApplicationIdentification(value string) {
	t.ApplicationIdentification = append(t.ApplicationIdentification, (*Max35Text)(&value))
}

func (t *TMSProtocolParameters2) SetHostIdentification(value string) {
	t.HostIdentification = (*Max35Text)(&value)
}

func (t *TMSProtocolParameters2) SetPOIIdentification(value string) {
	t.POIIdentification = (*Max35Text)(&value)
}

func (t *TMSProtocolParameters2) SetInitiatingPartyIdentification(value string) {
	t.InitiatingPartyIdentification = (*Max35Text)(&value)
}

func (t *TMSProtocolParameters2) SetRecipientPartyIdentification(value string) {
	t.RecipientPartyIdentification = (*Max35Text)(&value)
}

func (t *TMSProtocolParameters2) SetFileTransfer(value string) {
	t.FileTransfer = (*TrueFalseIndicator)(&value)
}
