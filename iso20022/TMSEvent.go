package iso20022

// Result of an individual terminal management action performed by the point of interaction.
type TMSEvent1 struct {

	// Date time of the terminal management action performed by the point of interaction.
	TimeStamp *ISODateTime `xml:"TmStmp"`

	// Final result of the processed terminal management action.
	Result *TerminalManagementActionResult1Code `xml:"Rslt"`

	// Identification of the terminal management action performed by the point of interaction.
	ActionIdentification *TMSActionIdentification1 `xml:"ActnId"`

	// Additional information related to a failure.
	AdditionalErrorInformation *Max70Text `xml:"AddtlErrInf,omitempty"`
}

func (t *TMSEvent1) SetTimeStamp(value string) {
	t.TimeStamp = (*ISODateTime)(&value)
}

func (t *TMSEvent1) SetResult(value string) {
	t.Result = (*TerminalManagementActionResult1Code)(&value)
}

func (t *TMSEvent1) AddActionIdentification() *TMSActionIdentification1 {
	t.ActionIdentification = new(TMSActionIdentification1)
	return t.ActionIdentification
}

func (t *TMSEvent1) SetAdditionalErrorInformation(value string) {
	t.AdditionalErrorInformation = (*Max70Text)(&value)
}

// Result of an individual terminal management action performed by the point of interaction.
type TMSEvent2 struct {

	// Date time of the terminal management action performed by the point of interaction.
	TimeStamp *ISODateTime `xml:"TmStmp"`

	// Final result of the processed terminal management action.
	Result *TerminalManagementActionResult1Code `xml:"Rslt"`

	// Identification of the terminal management action performed by the point of interaction.
	ActionIdentification *TMSActionIdentification2 `xml:"ActnId"`

	// Additional information related to a failure.
	AdditionalErrorInformation *Max70Text `xml:"AddtlErrInf,omitempty"`
}

func (t *TMSEvent2) SetTimeStamp(value string) {
	t.TimeStamp = (*ISODateTime)(&value)
}

func (t *TMSEvent2) SetResult(value string) {
	t.Result = (*TerminalManagementActionResult1Code)(&value)
}

func (t *TMSEvent2) AddActionIdentification() *TMSActionIdentification2 {
	t.ActionIdentification = new(TMSActionIdentification2)
	return t.ActionIdentification
}

func (t *TMSEvent2) SetAdditionalErrorInformation(value string) {
	t.AdditionalErrorInformation = (*Max70Text)(&value)
}

// Result of an individual terminal management action performed by the point of interaction.
type TMSEvent3 struct {

	// Date time of the terminal management action performed by the point of interaction.
	TimeStamp *ISODateTime `xml:"TmStmp"`

	// Final result of the processed terminal management action.
	Result *TerminalManagementActionResult1Code `xml:"Rslt"`

	// Identification of the terminal management action performed by the point of interaction.
	ActionIdentification *TMSActionIdentification3 `xml:"ActnId"`

	// Additional information related to a failure.
	AdditionalErrorInformation *Max70Text `xml:"AddtlErrInf,omitempty"`
}

func (t *TMSEvent3) SetTimeStamp(value string) {
	t.TimeStamp = (*ISODateTime)(&value)
}

func (t *TMSEvent3) SetResult(value string) {
	t.Result = (*TerminalManagementActionResult1Code)(&value)
}

func (t *TMSEvent3) AddActionIdentification() *TMSActionIdentification3 {
	t.ActionIdentification = new(TMSActionIdentification3)
	return t.ActionIdentification
}

func (t *TMSEvent3) SetAdditionalErrorInformation(value string) {
	t.AdditionalErrorInformation = (*Max70Text)(&value)
}

// Result of an individual terminal management action performed by the point of interaction.
type TMSEvent4 struct {

	// Date time of the terminal management action performed by the point of interaction.
	TimeStamp *ISODateTime `xml:"TmStmp"`

	// Final result of the processed terminal management action.
	Result *TerminalManagementActionResult1Code `xml:"Rslt"`

	// Identification of the terminal management action performed by the point of interaction.
	ActionIdentification *TMSActionIdentification4 `xml:"ActnId"`

	// Additional information related to a failure.
	AdditionalErrorInformation *Max70Text `xml:"AddtlErrInf,omitempty"`
}

func (t *TMSEvent4) SetTimeStamp(value string) {
	t.TimeStamp = (*ISODateTime)(&value)
}

func (t *TMSEvent4) SetResult(value string) {
	t.Result = (*TerminalManagementActionResult1Code)(&value)
}

func (t *TMSEvent4) AddActionIdentification() *TMSActionIdentification4 {
	t.ActionIdentification = new(TMSActionIdentification4)
	return t.ActionIdentification
}

func (t *TMSEvent4) SetAdditionalErrorInformation(value string) {
	t.AdditionalErrorInformation = (*Max70Text)(&value)
}
