package iso20022

// Set of actions to be performed by the POI (Point Of Interaction) system.
type Action1 struct {

	// Type of action to be performed by the POI (Point Of Interaction) system.
	ActionType *ActionType1Code `xml:"ActnTp"`

	// Message to be displayed to the cardholder or the cashier.
	MessageToPresent *ActionMessage1 `xml:"MsgToPres,omitempty"`
}

func (a *Action1) SetActionType(value string) {
	a.ActionType = (*ActionType1Code)(&value)
}

func (a *Action1) AddMessageToPresent() *ActionMessage1 {
	a.MessageToPresent = new(ActionMessage1)
	return a.MessageToPresent
}

// Set of actions to be performed by the POI (Point Of Interaction) system.
type Action2 struct {

	// Type of action to be performed by the POI (Point Of Interaction) system.
	ActionType *ActionType2Code `xml:"ActnTp"`

	// Message to be displayed to the cardholder or the cashier.
	MessageToPresent *ActionMessage1 `xml:"MsgToPres,omitempty"`
}

func (a *Action2) SetActionType(value string) {
	a.ActionType = (*ActionType2Code)(&value)
}

func (a *Action2) AddMessageToPresent() *ActionMessage1 {
	a.MessageToPresent = new(ActionMessage1)
	return a.MessageToPresent
}

// Set of actions to be performed by the POI (Point Of Interaction) system.
type Action3 struct {

	// Type of action to be performed by the POI (Point Of Interaction) system.
	ActionType *ActionType3Code `xml:"ActnTp"`

	// Message to be displayed to the cardholder or the cashier.
	MessageToPresent *ActionMessage1 `xml:"MsgToPres,omitempty"`
}

func (a *Action3) SetActionType(value string) {
	a.ActionType = (*ActionType3Code)(&value)
}

func (a *Action3) AddMessageToPresent() *ActionMessage1 {
	a.MessageToPresent = new(ActionMessage1)
	return a.MessageToPresent
}

// Set of actions to be performed by the card acceptor.
type Action4 struct {

	// Type of action to be performed by the card acceptor.
	ActionType *ActionType5Code `xml:"ActnTp"`

	// Message to be displayed to the cardholder or the cashier.
	MessageToPresent *ActionMessage2 `xml:"MsgToPres,omitempty"`
}

func (a *Action4) SetActionType(value string) {
	a.ActionType = (*ActionType5Code)(&value)
}

func (a *Action4) AddMessageToPresent() *ActionMessage2 {
	a.MessageToPresent = new(ActionMessage2)
	return a.MessageToPresent
}

// Set of actions to be performed by the card acceptor.
type Action5 struct {

	// Type of action to be performed by the card acceptor.
	ActionType *ActionType6Code `xml:"ActnTp"`

	// Information to display, print or log.
	MessageToPresent *ActionMessage4 `xml:"MsgToPres,omitempty"`

	// Message to send before the completion of the transaction.
	RequestToPerform *MessageFunction7Code `xml:"ReqToPrfrm,omitempty"`
}

func (a *Action5) SetActionType(value string) {
	a.ActionType = (*ActionType6Code)(&value)
}

func (a *Action5) AddMessageToPresent() *ActionMessage4 {
	a.MessageToPresent = new(ActionMessage4)
	return a.MessageToPresent
}

func (a *Action5) SetRequestToPerform(value string) {
	a.RequestToPerform = (*MessageFunction7Code)(&value)
}

// Set of actions to be performed by the POI (Point Of Interaction) system.
type Action6 struct {

	// Type of action to be performed by the POI (Point Of Interaction) system.
	ActionType *ActionType3Code `xml:"ActnTp"`

	// Message to be displayed to the cardholder or the cashier.
	MessageToPresent *ActionMessage2 `xml:"MsgToPres,omitempty"`
}

func (a *Action6) SetActionType(value string) {
	a.ActionType = (*ActionType3Code)(&value)
}

func (a *Action6) AddMessageToPresent() *ActionMessage2 {
	a.MessageToPresent = new(ActionMessage2)
	return a.MessageToPresent
}

// Set of actions to be performed by the card acceptor.
type Action7 struct {

	// Type of action to be performed by the card acceptor.
	ActionType *ActionType6Code `xml:"ActnTp"`

	// Information to display, print or log.
	MessageToPresent *ActionMessage4 `xml:"MsgToPres,omitempty"`

	// Message to send before the completion of the transaction.
	RequestToPerform *MessageFunction11Code `xml:"ReqToPrfrm,omitempty"`
}

func (a *Action7) SetActionType(value string) {
	a.ActionType = (*ActionType6Code)(&value)
}

func (a *Action7) AddMessageToPresent() *ActionMessage4 {
	a.MessageToPresent = new(ActionMessage4)
	return a.MessageToPresent
}

func (a *Action7) SetRequestToPerform(value string) {
	a.RequestToPerform = (*MessageFunction11Code)(&value)
}
