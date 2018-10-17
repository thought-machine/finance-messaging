package iso20022

// Result of an individual terminal management action by the point of interaction.
type TMSActionIdentification1 struct {

	// Types of terminal management action performed by a point of interaction.
	ActionType *TerminalManagementAction1Code `xml:"ActnTp"`

	// Data set on which the action has been performed.
	DataSetIdentification *DataSetIdentification2 `xml:"DataSetId,omitempty"`
}

func (t *TMSActionIdentification1) SetActionType(value string) {
	t.ActionType = (*TerminalManagementAction1Code)(&value)
}

func (t *TMSActionIdentification1) AddDataSetIdentification() *DataSetIdentification2 {
	t.DataSetIdentification = new(DataSetIdentification2)
	return t.DataSetIdentification
}

// Result of an individual terminal management action by the point of interaction.
type TMSActionIdentification2 struct {

	// Types of terminal management action performed by a point of interaction.
	ActionType *TerminalManagementAction1Code `xml:"ActnTp"`

	// Data set on which the action has been performed.
	DataSetIdentification *DataSetIdentification3 `xml:"DataSetId,omitempty"`
}

func (t *TMSActionIdentification2) SetActionType(value string) {
	t.ActionType = (*TerminalManagementAction1Code)(&value)
}

func (t *TMSActionIdentification2) AddDataSetIdentification() *DataSetIdentification3 {
	t.DataSetIdentification = new(DataSetIdentification3)
	return t.DataSetIdentification
}

// Result of an individual terminal management action by the point of interaction.
type TMSActionIdentification3 struct {

	// Types of terminal management action performed by a point of interaction.
	ActionType *TerminalManagementAction1Code `xml:"ActnTp"`

	// Data set on which the action has been performed.
	DataSetIdentification *DataSetIdentification4 `xml:"DataSetId,omitempty"`
}

func (t *TMSActionIdentification3) SetActionType(value string) {
	t.ActionType = (*TerminalManagementAction1Code)(&value)
}

func (t *TMSActionIdentification3) AddDataSetIdentification() *DataSetIdentification4 {
	t.DataSetIdentification = new(DataSetIdentification4)
	return t.DataSetIdentification
}

// Result of an individual terminal management action by the point of interaction.
type TMSActionIdentification4 struct {

	// Types of terminal management action performed by a point of interaction.
	ActionType *TerminalManagementAction2Code `xml:"ActnTp"`

	// Data set on which the action has been performed.
	DataSetIdentification *DataSetIdentification6 `xml:"DataSetId,omitempty"`
}

func (t *TMSActionIdentification4) SetActionType(value string) {
	t.ActionType = (*TerminalManagementAction2Code)(&value)
}

func (t *TMSActionIdentification4) AddDataSetIdentification() *DataSetIdentification6 {
	t.DataSetIdentification = new(DataSetIdentification6)
	return t.DataSetIdentification
}
