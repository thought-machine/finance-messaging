package iso20022

// Data related to a component of the POI performing the transaction.
type PointOfInteractionComponent1 struct {

	// Type of component belonging to a POI Terminal.
	POIComponentType *POIComponentType1Code `xml:"POICmpntTp"`

	// Identification of the software, hardware or system provider of the POI component.
	ManufacturerIdentification *Max35Text `xml:"ManfctrId,omitempty"`

	// Identification of a model of POI component for a given manufacturer.
	Model *Max35Text `xml:"Mdl,omitempty"`

	// Version of component belonging to a given model.
	VersionNumber *Max16Text `xml:"VrsnNb,omitempty"`

	// Serial number of a component.
	SerialNumber *Max35Text `xml:"SrlNb,omitempty"`

	// Unique approval number for a component, delivered by a certification body.
	// Usage: More than one approval number could be present, when assigned by different bodies. The certification body identification must be provided within the approval number (for example at the beginning of the value).
	ApprovalNumber []*Max70Text `xml:"ApprvlNb,omitempty"`
}

func (p *PointOfInteractionComponent1) SetPOIComponentType(value string) {
	p.POIComponentType = (*POIComponentType1Code)(&value)
}

func (p *PointOfInteractionComponent1) SetManufacturerIdentification(value string) {
	p.ManufacturerIdentification = (*Max35Text)(&value)
}

func (p *PointOfInteractionComponent1) SetModel(value string) {
	p.Model = (*Max35Text)(&value)
}

func (p *PointOfInteractionComponent1) SetVersionNumber(value string) {
	p.VersionNumber = (*Max16Text)(&value)
}

func (p *PointOfInteractionComponent1) SetSerialNumber(value string) {
	p.SerialNumber = (*Max35Text)(&value)
}

func (p *PointOfInteractionComponent1) AddApprovalNumber(value string) {
	p.ApprovalNumber = append(p.ApprovalNumber, (*Max70Text)(&value))
}

// Data related to a component of the POI.
type PointOfInteractionComponent2 struct {

	// Type of component belonging to a POI Terminal.
	POIComponentType *POIComponentType2Code `xml:"POICmpntTp"`

	// Identification of the software, hardware or system provider of the POI component.
	ManufacturerIdentification *Max35Text `xml:"ManfctrId,omitempty"`

	// Identification of a model of POI component for a given manufacturer.
	Model *Max35Text `xml:"Mdl,omitempty"`

	// Version of component belonging to a given model.
	VersionNumber *Max16Text `xml:"VrsnNb,omitempty"`

	// Serial number of a component.
	SerialNumber *Max35Text `xml:"SrlNb,omitempty"`

	// Unique approval number for a component, delivered by a certification body.
	// Usage: More than one approval number could be present, when assigned by different bodies. The certification body identification must be provided within the approval number (for example at the beginning of the value).
	ApprovalNumber []*Max70Text `xml:"ApprvlNb,omitempty"`
}

func (p *PointOfInteractionComponent2) SetPOIComponentType(value string) {
	p.POIComponentType = (*POIComponentType2Code)(&value)
}

func (p *PointOfInteractionComponent2) SetManufacturerIdentification(value string) {
	p.ManufacturerIdentification = (*Max35Text)(&value)
}

func (p *PointOfInteractionComponent2) SetModel(value string) {
	p.Model = (*Max35Text)(&value)
}

func (p *PointOfInteractionComponent2) SetVersionNumber(value string) {
	p.VersionNumber = (*Max16Text)(&value)
}

func (p *PointOfInteractionComponent2) SetSerialNumber(value string) {
	p.SerialNumber = (*Max35Text)(&value)
}

func (p *PointOfInteractionComponent2) AddApprovalNumber(value string) {
	p.ApprovalNumber = append(p.ApprovalNumber, (*Max70Text)(&value))
}

// Data related to a component of the POI performing the transaction.
type PointOfInteractionComponent3 struct {

	// Type of component belonging to a POI Terminal.
	Type *POIComponentType3Code `xml:"Tp"`

	// Identification of the POI component.
	Identification *PointOfInteractionComponentIdentification1 `xml:"Id"`

	// Status of the POI component.
	Status *PointOfInteractionComponentStatus1 `xml:"Sts,omitempty"`

	// Identification of the standard for which the component complies with.
	StandardCompliance []*GenericIdentification48 `xml:"StdCmplc,omitempty"`

	// Characteristics of a POI component.
	Characteristics *PointOfInteractionComponentCharacteristics1 `xml:"Chrtcs,omitempty"`

	// Assessments for the component of the point of interaction.
	Assessment []*PointOfInteractionComponentAssessment1 `xml:"Assmnt,omitempty"`
}

func (p *PointOfInteractionComponent3) SetType(value string) {
	p.Type = (*POIComponentType3Code)(&value)
}

func (p *PointOfInteractionComponent3) AddIdentification() *PointOfInteractionComponentIdentification1 {
	p.Identification = new(PointOfInteractionComponentIdentification1)
	return p.Identification
}

func (p *PointOfInteractionComponent3) AddStatus() *PointOfInteractionComponentStatus1 {
	p.Status = new(PointOfInteractionComponentStatus1)
	return p.Status
}

func (p *PointOfInteractionComponent3) AddStandardCompliance() *GenericIdentification48 {
	newValue := new(GenericIdentification48)
	p.StandardCompliance = append(p.StandardCompliance, newValue)
	return newValue
}

func (p *PointOfInteractionComponent3) AddCharacteristics() *PointOfInteractionComponentCharacteristics1 {
	p.Characteristics = new(PointOfInteractionComponentCharacteristics1)
	return p.Characteristics
}

func (p *PointOfInteractionComponent3) AddAssessment() *PointOfInteractionComponentAssessment1 {
	newValue := new(PointOfInteractionComponentAssessment1)
	p.Assessment = append(p.Assessment, newValue)
	return newValue
}

// Data related to a component of the POI (Point Of Interaction) performing the transaction.
type PointOfInteractionComponent4 struct {

	// Type of component belonging to a POI (Point Of Interaction) Terminal.
	Type *POIComponentType3Code `xml:"Tp"`

	// Identification of the POI (Point Of Interaction) component.
	Identification *PointOfInteractionComponentIdentification1 `xml:"Id"`

	// Status of the POI (Point Of Interaction) component.
	Status *PointOfInteractionComponentStatus1 `xml:"Sts,omitempty"`

	// Identification of the standard for which the component complies with.
	StandardCompliance []*GenericIdentification48 `xml:"StdCmplc,omitempty"`

	// Characteristics of a POI (Point Of Interaction) component.
	Characteristics *PointOfInteractionComponentCharacteristics2 `xml:"Chrtcs,omitempty"`

	// Assessments for the component of the point of interaction.
	Assessment []*PointOfInteractionComponentAssessment1 `xml:"Assmnt,omitempty"`
}

func (p *PointOfInteractionComponent4) SetType(value string) {
	p.Type = (*POIComponentType3Code)(&value)
}

func (p *PointOfInteractionComponent4) AddIdentification() *PointOfInteractionComponentIdentification1 {
	p.Identification = new(PointOfInteractionComponentIdentification1)
	return p.Identification
}

func (p *PointOfInteractionComponent4) AddStatus() *PointOfInteractionComponentStatus1 {
	p.Status = new(PointOfInteractionComponentStatus1)
	return p.Status
}

func (p *PointOfInteractionComponent4) AddStandardCompliance() *GenericIdentification48 {
	newValue := new(GenericIdentification48)
	p.StandardCompliance = append(p.StandardCompliance, newValue)
	return newValue
}

func (p *PointOfInteractionComponent4) AddCharacteristics() *PointOfInteractionComponentCharacteristics2 {
	p.Characteristics = new(PointOfInteractionComponentCharacteristics2)
	return p.Characteristics
}

func (p *PointOfInteractionComponent4) AddAssessment() *PointOfInteractionComponentAssessment1 {
	newValue := new(PointOfInteractionComponentAssessment1)
	p.Assessment = append(p.Assessment, newValue)
	return newValue
}

// Data related to a component of the POI (Point Of Interaction) performing the transaction.
type PointOfInteractionComponent5 struct {

	// Type of component belonging to a POI (Point Of Interaction) Terminal.
	Type *POIComponentType3Code `xml:"Tp"`

	// Identification of the POI (Point Of Interaction) component.
	Identification *PointOfInteractionComponentIdentification1 `xml:"Id"`

	// Status of the POI (Point Of Interaction) component.
	Status *PointOfInteractionComponentStatus2 `xml:"Sts,omitempty"`

	// Identification of the standard for which the component complies with.
	StandardCompliance []*GenericIdentification48 `xml:"StdCmplc,omitempty"`

	// Characteristics of a POI (Point Of Interaction) component.
	Characteristics *PointOfInteractionComponentCharacteristics2 `xml:"Chrtcs,omitempty"`

	// Assessments for the component of the point of interaction.
	Assessment []*PointOfInteractionComponentAssessment1 `xml:"Assmnt,omitempty"`
}

func (p *PointOfInteractionComponent5) SetType(value string) {
	p.Type = (*POIComponentType3Code)(&value)
}

func (p *PointOfInteractionComponent5) AddIdentification() *PointOfInteractionComponentIdentification1 {
	p.Identification = new(PointOfInteractionComponentIdentification1)
	return p.Identification
}

func (p *PointOfInteractionComponent5) AddStatus() *PointOfInteractionComponentStatus2 {
	p.Status = new(PointOfInteractionComponentStatus2)
	return p.Status
}

func (p *PointOfInteractionComponent5) AddStandardCompliance() *GenericIdentification48 {
	newValue := new(GenericIdentification48)
	p.StandardCompliance = append(p.StandardCompliance, newValue)
	return newValue
}

func (p *PointOfInteractionComponent5) AddCharacteristics() *PointOfInteractionComponentCharacteristics2 {
	p.Characteristics = new(PointOfInteractionComponentCharacteristics2)
	return p.Characteristics
}

func (p *PointOfInteractionComponent5) AddAssessment() *PointOfInteractionComponentAssessment1 {
	newValue := new(PointOfInteractionComponentAssessment1)
	p.Assessment = append(p.Assessment, newValue)
	return newValue
}

// Data related to a component of the POI (Point Of Interaction) performing the transaction.
type PointOfInteractionComponent6 struct {

	// Type of component belonging to a POI (Point Of Interaction) Terminal.
	Type *POIComponentType4Code `xml:"Tp"`

	// Identification of the POI (Point Of Interaction) component.
	Identification *PointOfInteractionComponentIdentification1 `xml:"Id"`

	// Status of the POI (Point Of Interaction) component.
	Status *PointOfInteractionComponentStatus3 `xml:"Sts,omitempty"`

	// Identification of the standard for which the component complies with.
	StandardCompliance []*GenericIdentification48 `xml:"StdCmplc,omitempty"`

	// Characteristics of a POI (Point Of Interaction) component.
	Characteristics *PointOfInteractionComponentCharacteristics2 `xml:"Chrtcs,omitempty"`

	// Assessments for the component of the point of interaction.
	Assessment []*PointOfInteractionComponentAssessment1 `xml:"Assmnt,omitempty"`
}

func (p *PointOfInteractionComponent6) SetType(value string) {
	p.Type = (*POIComponentType4Code)(&value)
}

func (p *PointOfInteractionComponent6) AddIdentification() *PointOfInteractionComponentIdentification1 {
	p.Identification = new(PointOfInteractionComponentIdentification1)
	return p.Identification
}

func (p *PointOfInteractionComponent6) AddStatus() *PointOfInteractionComponentStatus3 {
	p.Status = new(PointOfInteractionComponentStatus3)
	return p.Status
}

func (p *PointOfInteractionComponent6) AddStandardCompliance() *GenericIdentification48 {
	newValue := new(GenericIdentification48)
	p.StandardCompliance = append(p.StandardCompliance, newValue)
	return newValue
}

func (p *PointOfInteractionComponent6) AddCharacteristics() *PointOfInteractionComponentCharacteristics2 {
	p.Characteristics = new(PointOfInteractionComponentCharacteristics2)
	return p.Characteristics
}

func (p *PointOfInteractionComponent6) AddAssessment() *PointOfInteractionComponentAssessment1 {
	newValue := new(PointOfInteractionComponentAssessment1)
	p.Assessment = append(p.Assessment, newValue)
	return newValue
}
