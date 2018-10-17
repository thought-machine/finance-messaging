package iso20022

// Point of interaction (POI) performing the transaction.
type PointOfInteraction1 struct {

	// Identification of the POI (Point Of Interaction) for the acquirer or its agent.
	Identification *GenericIdentification32 `xml:"Id"`

	// Common name assigned by the acquirer to the POI system.
	SystemName *Max70Text `xml:"SysNm,omitempty"`

	// Identifier assigned by the merchant identifying a set of POI terminals performing some categories of transactions.
	GroupIdentification *Max35Text `xml:"GrpId,omitempty"`

	// Capabilities of the POI performing the transaction.
	Capabilities *PointOfInteractionCapabilities1 `xml:"Cpblties,omitempty"`

	// Data related to a component of the POI performing the transaction.
	Component []*PointOfInteractionComponent1 `xml:"Cmpnt,omitempty"`
}

func (p *PointOfInteraction1) AddIdentification() *GenericIdentification32 {
	p.Identification = new(GenericIdentification32)
	return p.Identification
}

func (p *PointOfInteraction1) SetSystemName(value string) {
	p.SystemName = (*Max70Text)(&value)
}

func (p *PointOfInteraction1) SetGroupIdentification(value string) {
	p.GroupIdentification = (*Max35Text)(&value)
}

func (p *PointOfInteraction1) AddCapabilities() *PointOfInteractionCapabilities1 {
	p.Capabilities = new(PointOfInteractionCapabilities1)
	return p.Capabilities
}

func (p *PointOfInteraction1) AddComponent() *PointOfInteractionComponent1 {
	newValue := new(PointOfInteractionComponent1)
	p.Component = append(p.Component, newValue)
	return newValue
}

// Point of interaction (POI) performing the transaction.
type PointOfInteraction2 struct {

	// Identification of the POI (Point Of Interaction) for the acquirer or its agent.
	Identification *GenericIdentification32 `xml:"Id"`

	// Common name assigned by the acquirer to the POI system.
	SystemName *Max70Text `xml:"SysNm,omitempty"`

	// Identifier assigned by the merchant identifying a set of POI terminals performing some categories of transactions.
	GroupIdentification *Max35Text `xml:"GrpId,omitempty"`

	// Capabilities of the POI performing the transaction.
	Capabilities *PointOfInteractionCapabilities1 `xml:"Cpblties,omitempty"`

	// Data related to a component of the POI performing the transaction.
	Component []*PointOfInteractionComponent3 `xml:"Cmpnt,omitempty"`
}

func (p *PointOfInteraction2) AddIdentification() *GenericIdentification32 {
	p.Identification = new(GenericIdentification32)
	return p.Identification
}

func (p *PointOfInteraction2) SetSystemName(value string) {
	p.SystemName = (*Max70Text)(&value)
}

func (p *PointOfInteraction2) SetGroupIdentification(value string) {
	p.GroupIdentification = (*Max35Text)(&value)
}

func (p *PointOfInteraction2) AddCapabilities() *PointOfInteractionCapabilities1 {
	p.Capabilities = new(PointOfInteractionCapabilities1)
	return p.Capabilities
}

func (p *PointOfInteraction2) AddComponent() *PointOfInteractionComponent3 {
	newValue := new(PointOfInteractionComponent3)
	p.Component = append(p.Component, newValue)
	return newValue
}

// Point of interaction (POI) performing the transaction.
type PointOfInteraction3 struct {

	// Identification of the POI (Point Of Interaction) for the acquirer or its agent.
	Identification *GenericIdentification32 `xml:"Id"`

	// Common name assigned by the acquirer to the POI (Point Of Interaction) system.
	SystemName *Max70Text `xml:"SysNm,omitempty"`

	// Identifier assigned by the merchant identifying a set ofPOI (Point Of Interaction) terminals performing some categories of transactions.
	GroupIdentification *Max35Text `xml:"GrpId,omitempty"`

	// Capabilities of the POI (Point Of Interaction) performing the transaction.
	Capabilities *PointOfInteractionCapabilities2 `xml:"Cpblties,omitempty"`

	// Data related to a component of the POI (Point Of Interaction) performing the transaction.
	Component []*PointOfInteractionComponent4 `xml:"Cmpnt,omitempty"`
}

func (p *PointOfInteraction3) AddIdentification() *GenericIdentification32 {
	p.Identification = new(GenericIdentification32)
	return p.Identification
}

func (p *PointOfInteraction3) SetSystemName(value string) {
	p.SystemName = (*Max70Text)(&value)
}

func (p *PointOfInteraction3) SetGroupIdentification(value string) {
	p.GroupIdentification = (*Max35Text)(&value)
}

func (p *PointOfInteraction3) AddCapabilities() *PointOfInteractionCapabilities2 {
	p.Capabilities = new(PointOfInteractionCapabilities2)
	return p.Capabilities
}

func (p *PointOfInteraction3) AddComponent() *PointOfInteractionComponent4 {
	newValue := new(PointOfInteractionComponent4)
	p.Component = append(p.Component, newValue)
	return newValue
}

// Point of interaction (POI) performing the transaction.
type PointOfInteraction4 struct {

	// Identification of the POI (Point Of Interaction) for the acquirer or its agent.
	Identification *GenericIdentification32 `xml:"Id"`

	// Common name assigned by the acquirer to the POI (Point Of Interaction) system.
	SystemName *Max70Text `xml:"SysNm,omitempty"`

	// Identifier assigned by the merchant identifying a set of POI (Point Of Interaction) terminals performing some categories of transactions.
	GroupIdentification *Max35Text `xml:"GrpId,omitempty"`

	// Capabilities of the POI (Point Of Interaction) performing the transaction.
	Capabilities *PointOfInteractionCapabilities3 `xml:"Cpblties,omitempty"`

	// Data related to a component of the POI (Point Of Interaction) performing the transaction.
	Component []*PointOfInteractionComponent5 `xml:"Cmpnt,omitempty"`
}

func (p *PointOfInteraction4) AddIdentification() *GenericIdentification32 {
	p.Identification = new(GenericIdentification32)
	return p.Identification
}

func (p *PointOfInteraction4) SetSystemName(value string) {
	p.SystemName = (*Max70Text)(&value)
}

func (p *PointOfInteraction4) SetGroupIdentification(value string) {
	p.GroupIdentification = (*Max35Text)(&value)
}

func (p *PointOfInteraction4) AddCapabilities() *PointOfInteractionCapabilities3 {
	p.Capabilities = new(PointOfInteractionCapabilities3)
	return p.Capabilities
}

func (p *PointOfInteraction4) AddComponent() *PointOfInteractionComponent5 {
	newValue := new(PointOfInteractionComponent5)
	p.Component = append(p.Component, newValue)
	return newValue
}

// Point of interaction (POI) performing the transaction.
type PointOfInteraction5 struct {

	// Identification of the POI (Point Of Interaction) for the acquirer or its agent.
	Identification *GenericIdentification32 `xml:"Id"`

	// Common name assigned by the acquirer to the POI (Point Of Interaction) system.
	SystemName *Max70Text `xml:"SysNm,omitempty"`

	// Identifier assigned by the merchant identifying a set of POI (Point Of Interaction) terminals performing some categories of transactions.
	GroupIdentification *Max35Text `xml:"GrpId,omitempty"`

	// Capabilities of the POI (Point Of Interaction) performing the transaction.
	Capabilities *PointOfInteractionCapabilities6 `xml:"Cpblties,omitempty"`

	// Time zone name as defined by IANA (Internet Assigned Numbers Authority) in the time zone data base. America/Chicago or Europe/Paris are examples of time zone names.
	TimeZone *Max70Text `xml:"TmZone,omitempty"`

	// Indicates the type of integration of the POI terminal in the sale environment.
	TerminalIntegration *LocationCategory3Code `xml:"TermnlIntgtn,omitempty"`

	// Data related to a component of the POI (Point Of Interaction) performing the transaction.
	Component []*PointOfInteractionComponent6 `xml:"Cmpnt,omitempty"`
}

func (p *PointOfInteraction5) AddIdentification() *GenericIdentification32 {
	p.Identification = new(GenericIdentification32)
	return p.Identification
}

func (p *PointOfInteraction5) SetSystemName(value string) {
	p.SystemName = (*Max70Text)(&value)
}

func (p *PointOfInteraction5) SetGroupIdentification(value string) {
	p.GroupIdentification = (*Max35Text)(&value)
}

func (p *PointOfInteraction5) AddCapabilities() *PointOfInteractionCapabilities6 {
	p.Capabilities = new(PointOfInteractionCapabilities6)
	return p.Capabilities
}

func (p *PointOfInteraction5) SetTimeZone(value string) {
	p.TimeZone = (*Max70Text)(&value)
}

func (p *PointOfInteraction5) SetTerminalIntegration(value string) {
	p.TerminalIntegration = (*LocationCategory3Code)(&value)
}

func (p *PointOfInteraction5) AddComponent() *PointOfInteractionComponent6 {
	newValue := new(PointOfInteractionComponent6)
	p.Component = append(p.Component, newValue)
	return newValue
}

// Identification of a point of interaction.
type PointOfInteraction6 struct {

	// Identifier of the terminal manufacturer.
	ManufacturerIdentifier *Max35Text `xml:"ManfctrIdr"`

	// Identifier of the terminal model.
	Model *Max35Text `xml:"Mdl"`

	// Serial number of the terminal manufacturer.
	SerialNumber *Max35Text `xml:"SrlNb"`
}

func (p *PointOfInteraction6) SetManufacturerIdentifier(value string) {
	p.ManufacturerIdentifier = (*Max35Text)(&value)
}

func (p *PointOfInteraction6) SetModel(value string) {
	p.Model = (*Max35Text)(&value)
}

func (p *PointOfInteraction6) SetSerialNumber(value string) {
	p.SerialNumber = (*Max35Text)(&value)
}
