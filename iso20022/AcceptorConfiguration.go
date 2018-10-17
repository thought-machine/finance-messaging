package iso20022

// Acceptor configuration to be downloaded from the terminal management system.
type AcceptorConfiguration1 struct {

	// Identification of the point of interaction for terminal management.
	POIIdentification *GenericIdentification35 `xml:"POIId,omitempty"`

	// Identification of the terminal management system (TMS) sending the acceptor parameters.
	TerminalManagerIdentification *GenericIdentification35 `xml:"TermnlMgrId"`

	// Data set containing the acceptor parameters of a point of interaction (POI).
	DataSet []*TerminalManagementDataSet3 `xml:"DataSet"`
}

func (a *AcceptorConfiguration1) AddPOIIdentification() *GenericIdentification35 {
	a.POIIdentification = new(GenericIdentification35)
	return a.POIIdentification
}

func (a *AcceptorConfiguration1) AddTerminalManagerIdentification() *GenericIdentification35 {
	a.TerminalManagerIdentification = new(GenericIdentification35)
	return a.TerminalManagerIdentification
}

func (a *AcceptorConfiguration1) AddDataSet() *TerminalManagementDataSet3 {
	newValue := new(TerminalManagementDataSet3)
	a.DataSet = append(a.DataSet, newValue)
	return newValue
}

// Acceptor configuration to be downloaded from the terminal management system.
type AcceptorConfiguration2 struct {

	// Identification of the point of interaction for terminal management.
	POIIdentification *GenericIdentification35 `xml:"POIId,omitempty"`

	// Identification of the terminal management system (TMS) sending the acceptor parameters.
	TerminalManagerIdentification *GenericIdentification35 `xml:"TermnlMgrId"`

	// Data set containing the acceptor parameters of a point of interaction (POI).
	DataSet []*TerminalManagementDataSet6 `xml:"DataSet"`
}

func (a *AcceptorConfiguration2) AddPOIIdentification() *GenericIdentification35 {
	a.POIIdentification = new(GenericIdentification35)
	return a.POIIdentification
}

func (a *AcceptorConfiguration2) AddTerminalManagerIdentification() *GenericIdentification35 {
	a.TerminalManagerIdentification = new(GenericIdentification35)
	return a.TerminalManagerIdentification
}

func (a *AcceptorConfiguration2) AddDataSet() *TerminalManagementDataSet6 {
	newValue := new(TerminalManagementDataSet6)
	a.DataSet = append(a.DataSet, newValue)
	return newValue
}

// Acceptor configuration to be downloaded from the terminal management system.
type AcceptorConfiguration3 struct {

	// Identification of the point of interaction for terminal management.
	POIIdentification *GenericIdentification35 `xml:"POIId,omitempty"`

	// Identification of the terminal management system (TMS) sending the acceptor parameters.
	TerminalManagerIdentification *GenericIdentification35 `xml:"TermnlMgrId"`

	// Data set containing the acceptor parameters of a point of interaction (POI).
	DataSet []*TerminalManagementDataSet11 `xml:"DataSet"`
}

func (a *AcceptorConfiguration3) AddPOIIdentification() *GenericIdentification35 {
	a.POIIdentification = new(GenericIdentification35)
	return a.POIIdentification
}

func (a *AcceptorConfiguration3) AddTerminalManagerIdentification() *GenericIdentification35 {
	a.TerminalManagerIdentification = new(GenericIdentification35)
	return a.TerminalManagerIdentification
}

func (a *AcceptorConfiguration3) AddDataSet() *TerminalManagementDataSet11 {
	newValue := new(TerminalManagementDataSet11)
	a.DataSet = append(a.DataSet, newValue)
	return newValue
}

// Acceptor configuration to be downloaded from the terminal management system.
type AcceptorConfiguration4 struct {

	// Identification of the point of interaction for terminal management.
	POIIdentification *GenericIdentification71 `xml:"POIId,omitempty"`

	// Identification of the terminal management system (TMS) sending the acceptor parameters.
	TerminalManagerIdentification *GenericIdentification71 `xml:"TermnlMgrId"`

	// Data set containing the acceptor parameters of a point of interaction (POI).
	DataSet []*TerminalManagementDataSet14 `xml:"DataSet"`
}

func (a *AcceptorConfiguration4) AddPOIIdentification() *GenericIdentification71 {
	a.POIIdentification = new(GenericIdentification71)
	return a.POIIdentification
}

func (a *AcceptorConfiguration4) AddTerminalManagerIdentification() *GenericIdentification71 {
	a.TerminalManagerIdentification = new(GenericIdentification71)
	return a.TerminalManagerIdentification
}

func (a *AcceptorConfiguration4) AddDataSet() *TerminalManagementDataSet14 {
	newValue := new(TerminalManagementDataSet14)
	a.DataSet = append(a.DataSet, newValue)
	return newValue
}

// Acceptor configuration to be downloaded from the terminal management system.
type AcceptorConfiguration5 struct {

	// Identification of the terminal management system (TMS) sending the acceptor parameters.
	TerminalManagerIdentification *GenericIdentification71 `xml:"TermnlMgrId"`

	// Data set containing the acceptor parameters of a point of interaction (POI).
	DataSet []*TerminalManagementDataSet19 `xml:"DataSet"`
}

func (a *AcceptorConfiguration5) AddTerminalManagerIdentification() *GenericIdentification71 {
	a.TerminalManagerIdentification = new(GenericIdentification71)
	return a.TerminalManagerIdentification
}

func (a *AcceptorConfiguration5) AddDataSet() *TerminalManagementDataSet19 {
	newValue := new(TerminalManagementDataSet19)
	a.DataSet = append(a.DataSet, newValue)
	return newValue
}
