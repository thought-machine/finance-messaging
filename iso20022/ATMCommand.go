package iso20022

// Maintenance command to perform on an ATM.
type ATMCommand1 struct {

	// Type of command to be performed by the ATM.
	Type *ATMCommand1Code `xml:"Tp"`

	// Urgency of the command.
	Urgency *TMSContactLevel2Code `xml:"Urgcy"`

	// Date time on which the command must be performed.
	DateTime *ISODateTime `xml:"DtTm,omitempty"`

	// Identification of the entity issuing the command.
	CommandIdentification *ATMCommandIdentification1 `xml:"CmdId,omitempty"`

	// Specific parameters attached to the command.
	CommandParameters *ATMCommandParameters1Choice `xml:"CmdParams,omitempty"`
}

func (a *ATMCommand1) SetType(value string) {
	a.Type = (*ATMCommand1Code)(&value)
}

func (a *ATMCommand1) SetUrgency(value string) {
	a.Urgency = (*TMSContactLevel2Code)(&value)
}

func (a *ATMCommand1) SetDateTime(value string) {
	a.DateTime = (*ISODateTime)(&value)
}

func (a *ATMCommand1) AddCommandIdentification() *ATMCommandIdentification1 {
	a.CommandIdentification = new(ATMCommandIdentification1)
	return a.CommandIdentification
}

func (a *ATMCommand1) AddCommandParameters() *ATMCommandParameters1Choice {
	a.CommandParameters = new(ATMCommandParameters1Choice)
	return a.CommandParameters
}

// Result of a maintenance command performed by the ATM.
type ATMCommand2 struct {

	// Type of command to be performed by the ATM.
	Type *ATMCommand2Code `xml:"Tp"`

	// Date time on which the command has been requested to be performed.
	RequiredDateTime *ISODateTime `xml:"ReqrdDtTm,omitempty"`

	// Date time on which the command has been performed.
	ProcessedDateTime *ISODateTime `xml:"PrcdDtTm"`

	// Identification of the entity issuing the command.
	CommandIdentification *ATMCommandIdentification1 `xml:"CmdId,omitempty"`

	// Final result of the processed command at the ATM.
	Result *TerminalManagementActionResult2Code `xml:"Rslt"`

	// Additional information on the failure to be logged for further examination.
	AdditionalErrorInformation *Max140Text `xml:"AddtlErrInf,omitempty"`
}

func (a *ATMCommand2) SetType(value string) {
	a.Type = (*ATMCommand2Code)(&value)
}

func (a *ATMCommand2) SetRequiredDateTime(value string) {
	a.RequiredDateTime = (*ISODateTime)(&value)
}

func (a *ATMCommand2) SetProcessedDateTime(value string) {
	a.ProcessedDateTime = (*ISODateTime)(&value)
}

func (a *ATMCommand2) AddCommandIdentification() *ATMCommandIdentification1 {
	a.CommandIdentification = new(ATMCommandIdentification1)
	return a.CommandIdentification
}

func (a *ATMCommand2) SetResult(value string) {
	a.Result = (*TerminalManagementActionResult2Code)(&value)
}

func (a *ATMCommand2) SetAdditionalErrorInformation(value string) {
	a.AdditionalErrorInformation = (*Max140Text)(&value)
}

// Maintenance command which has requested the device report.
type ATMCommand3 struct {

	// Type of command to be performed by the ATM.
	Type *ATMCommand2Code `xml:"Tp"`

	// Identification of the entity issuing the command.
	CommandIdentification *ATMCommandIdentification1 `xml:"CmdId,omitempty"`
}

func (a *ATMCommand3) SetType(value string) {
	a.Type = (*ATMCommand2Code)(&value)
}

func (a *ATMCommand3) AddCommandIdentification() *ATMCommandIdentification1 {
	a.CommandIdentification = new(ATMCommandIdentification1)
	return a.CommandIdentification
}

// Maintenance command the ATM must perform.
type ATMCommand4 struct {

	// Type of command to be performed by the ATM.
	Type *ATMCommand2Code `xml:"Tp"`

	// Urgency of the command.
	Urgency *TMSContactLevel2Code `xml:"Urgcy"`

	// Date time on which the command must be performed.
	DateTime *ISODateTime `xml:"DtTm,omitempty"`

	// Identification of the entity issuing the command.
	CommandIdentification *ATMCommandIdentification1 `xml:"CmdId,omitempty"`

	// Reason for sending the command.
	Reason *ATMCommandReason1Code `xml:"Rsn,omitempty"`

	// Trace of reasons by the entities in the path from the origin of the command to the ATM.
	TraceReason []*ATMCommandReason1Code `xml:"TracRsn,omitempty"`

	// Additional information about the reason to request this command.
	AdditionalReasonInformation *Max70Text `xml:"AddtlRsnInf,omitempty"`

	// Specific parameters attached to the command.
	CommandParameters *ATMCommandParameters2Choice `xml:"CmdParams,omitempty"`
}

func (a *ATMCommand4) SetType(value string) {
	a.Type = (*ATMCommand2Code)(&value)
}

func (a *ATMCommand4) SetUrgency(value string) {
	a.Urgency = (*TMSContactLevel2Code)(&value)
}

func (a *ATMCommand4) SetDateTime(value string) {
	a.DateTime = (*ISODateTime)(&value)
}

func (a *ATMCommand4) AddCommandIdentification() *ATMCommandIdentification1 {
	a.CommandIdentification = new(ATMCommandIdentification1)
	return a.CommandIdentification
}

func (a *ATMCommand4) SetReason(value string) {
	a.Reason = (*ATMCommandReason1Code)(&value)
}

func (a *ATMCommand4) AddTraceReason(value string) {
	a.TraceReason = append(a.TraceReason, (*ATMCommandReason1Code)(&value))
}

func (a *ATMCommand4) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max70Text)(&value)
}

func (a *ATMCommand4) AddCommandParameters() *ATMCommandParameters2Choice {
	a.CommandParameters = new(ATMCommandParameters2Choice)
	return a.CommandParameters
}

// Command result for reinitialization of the transaction counters.
type ATMCommand5 struct {

	// Type of command to be performed by the ATM.
	Type *ATMCommand3Code `xml:"Tp"`

	// Date time on which the command has been requested to be performed.
	RequiredDateTime *ISODateTime `xml:"ReqrdDtTm,omitempty"`

	// Date time on which the command has been performed.
	ProcessedDateTime *ISODateTime `xml:"PrcdDtTm"`

	// Identification of the entity issuing the command.
	CommandIdentification *ATMCommandIdentification1 `xml:"CmdId,omitempty"`

	// Final result of the processed command at the ATM.
	Result *TerminalManagementActionResult2Code `xml:"Rslt"`

	// Additional information on the failure to be logged for further examination.
	AdditionalErrorInformation *Max140Text `xml:"AddtlErrInf,omitempty"`
}

func (a *ATMCommand5) SetType(value string) {
	a.Type = (*ATMCommand3Code)(&value)
}

func (a *ATMCommand5) SetRequiredDateTime(value string) {
	a.RequiredDateTime = (*ISODateTime)(&value)
}

func (a *ATMCommand5) SetProcessedDateTime(value string) {
	a.ProcessedDateTime = (*ISODateTime)(&value)
}

func (a *ATMCommand5) AddCommandIdentification() *ATMCommandIdentification1 {
	a.CommandIdentification = new(ATMCommandIdentification1)
	return a.CommandIdentification
}

func (a *ATMCommand5) SetResult(value string) {
	a.Result = (*TerminalManagementActionResult2Code)(&value)
}

func (a *ATMCommand5) SetAdditionalErrorInformation(value string) {
	a.AdditionalErrorInformation = (*Max140Text)(&value)
}

// Party which has requested the reconciliation.
type ATMCommand6 struct {

	// Type of command to be performed by the ATM.
	Type *ATMCommand3Code `xml:"Tp"`

	// Identification of the entity issuing the command.
	CommandIdentification *ATMCommandIdentification1 `xml:"CmdId,omitempty"`
}

func (a *ATMCommand6) SetType(value string) {
	a.Type = (*ATMCommand3Code)(&value)
}

func (a *ATMCommand6) AddCommandIdentification() *ATMCommandIdentification1 {
	a.CommandIdentification = new(ATMCommandIdentification1)
	return a.CommandIdentification
}

// Maintenance command to perform on an ATM.
type ATMCommand7 struct {

	// Type of command to be performed by the ATM.
	Type *ATMCommand4Code `xml:"Tp"`

	// Urgency of the command.
	Urgency *TMSContactLevel2Code `xml:"Urgcy"`

	// Date time on which the command must be performed.
	DateTime *ISODateTime `xml:"DtTm,omitempty"`

	// Identification of the entity issuing the command.
	CommandIdentification *ATMCommandIdentification1 `xml:"CmdId,omitempty"`

	// Specific parameters attached to the command.
	CommandParameters *ATMCommandParameters1Choice `xml:"CmdParams,omitempty"`
}

func (a *ATMCommand7) SetType(value string) {
	a.Type = (*ATMCommand4Code)(&value)
}

func (a *ATMCommand7) SetUrgency(value string) {
	a.Urgency = (*TMSContactLevel2Code)(&value)
}

func (a *ATMCommand7) SetDateTime(value string) {
	a.DateTime = (*ISODateTime)(&value)
}

func (a *ATMCommand7) AddCommandIdentification() *ATMCommandIdentification1 {
	a.CommandIdentification = new(ATMCommandIdentification1)
	return a.CommandIdentification
}

func (a *ATMCommand7) AddCommandParameters() *ATMCommandParameters1Choice {
	a.CommandParameters = new(ATMCommandParameters1Choice)
	return a.CommandParameters
}

// Command result for reinitialization of the transaction counters.
type ATMCommand8 struct {

	// Type of command to be performed by the ATM.
	Type *ATMCommand5Code `xml:"Tp"`

	// Date time on which the command has been requested to be performed.
	RequiredDateTime *ISODateTime `xml:"ReqrdDtTm,omitempty"`

	// Date time on which the command has been performed.
	ProcessedDateTime *ISODateTime `xml:"PrcdDtTm"`

	// Identification of the entity issuing the command.
	CommandIdentification *ATMCommandIdentification1 `xml:"CmdId,omitempty"`

	// Final result of the processed command at the ATM.
	Result *TerminalManagementActionResult2Code `xml:"Rslt"`

	// Additional information on the failure to be logged for further examination.
	AdditionalErrorInformation *Max140Text `xml:"AddtlErrInf,omitempty"`
}

func (a *ATMCommand8) SetType(value string) {
	a.Type = (*ATMCommand5Code)(&value)
}

func (a *ATMCommand8) SetRequiredDateTime(value string) {
	a.RequiredDateTime = (*ISODateTime)(&value)
}

func (a *ATMCommand8) SetProcessedDateTime(value string) {
	a.ProcessedDateTime = (*ISODateTime)(&value)
}

func (a *ATMCommand8) AddCommandIdentification() *ATMCommandIdentification1 {
	a.CommandIdentification = new(ATMCommandIdentification1)
	return a.CommandIdentification
}

func (a *ATMCommand8) SetResult(value string) {
	a.Result = (*TerminalManagementActionResult2Code)(&value)
}

func (a *ATMCommand8) SetAdditionalErrorInformation(value string) {
	a.AdditionalErrorInformation = (*Max140Text)(&value)
}

// Party which has requested the reconciliation.
type ATMCommand9 struct {

	// Type of command to be performed by the ATM.
	Type *ATMCommand5Code `xml:"Tp"`

	// Identification of the entity issuing the command.
	CommandIdentification *ATMCommandIdentification1 `xml:"CmdId,omitempty"`
}

func (a *ATMCommand9) SetType(value string) {
	a.Type = (*ATMCommand5Code)(&value)
}

func (a *ATMCommand9) AddCommandIdentification() *ATMCommandIdentification1 {
	a.CommandIdentification = new(ATMCommandIdentification1)
	return a.CommandIdentification
}
