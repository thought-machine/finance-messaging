package iso20022

// Status of a POI component (Point of Interaction).
type PointOfInteractionComponentStatus1 struct {

	// Current version of the component that might include the release number.
	VersionNumber *Max35Text `xml:"VrsnNb,omitempty"`

	// Current status of the component.
	Status *POIComponentStatus1Code `xml:"Sts,omitempty"`
}

func (p *PointOfInteractionComponentStatus1) SetVersionNumber(value string) {
	p.VersionNumber = (*Max35Text)(&value)
}

func (p *PointOfInteractionComponentStatus1) SetStatus(value string) {
	p.Status = (*POIComponentStatus1Code)(&value)
}

// Status of a POI component (Point of Interaction).
type PointOfInteractionComponentStatus2 struct {

	// Current version of the component that might include the release number.
	VersionNumber *Max256Text `xml:"VrsnNb,omitempty"`

	// Current status of the component.
	Status *POIComponentStatus1Code `xml:"Sts,omitempty"`
}

func (p *PointOfInteractionComponentStatus2) SetVersionNumber(value string) {
	p.VersionNumber = (*Max256Text)(&value)
}

func (p *PointOfInteractionComponentStatus2) SetStatus(value string) {
	p.Status = (*POIComponentStatus1Code)(&value)
}

// Status of a POI component (Point of Interaction).
type PointOfInteractionComponentStatus3 struct {

	// Current version of the component that might include the release number.
	VersionNumber *Max256Text `xml:"VrsnNb,omitempty"`

	// Current status of the component.
	Status *POIComponentStatus1Code `xml:"Sts,omitempty"`

	// Expiration date of the component.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`
}

func (p *PointOfInteractionComponentStatus3) SetVersionNumber(value string) {
	p.VersionNumber = (*Max256Text)(&value)
}

func (p *PointOfInteractionComponentStatus3) SetStatus(value string) {
	p.Status = (*POIComponentStatus1Code)(&value)
}

func (p *PointOfInteractionComponentStatus3) SetExpiryDate(value string) {
	p.ExpiryDate = (*ISODate)(&value)
}
