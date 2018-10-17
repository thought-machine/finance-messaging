package iso20022

// Identification of a linked corporate action event.
type CorporateActionEventReference1 struct {

	// Identification of the linked corporate action event.
	EventIdentification *CorporateActionEventReference1Choice `xml:"EvtId"`

	// Specifies when this corporate action event is to be processed relative to a linked corporate action event.
	LinkageType *ProcessingPosition1Choice `xml:"LkgTp,omitempty"`
}

func (c *CorporateActionEventReference1) AddEventIdentification() *CorporateActionEventReference1Choice {
	c.EventIdentification = new(CorporateActionEventReference1Choice)
	return c.EventIdentification
}

func (c *CorporateActionEventReference1) AddLinkageType() *ProcessingPosition1Choice {
	c.LinkageType = new(ProcessingPosition1Choice)
	return c.LinkageType
}

// Identification of a linked corporate action event.
type CorporateActionEventReference3 struct {

	// Identification of the linked corporate action event.
	EventIdentification *CorporateActionEventReference3Choice `xml:"EvtId"`

	// Specifies when this corporate action event is to be processed relative to a linked corporate action event.
	LinkageType *ProcessingPosition7Choice `xml:"LkgTp,omitempty"`
}

func (c *CorporateActionEventReference3) AddEventIdentification() *CorporateActionEventReference3Choice {
	c.EventIdentification = new(CorporateActionEventReference3Choice)
	return c.EventIdentification
}

func (c *CorporateActionEventReference3) AddLinkageType() *ProcessingPosition7Choice {
	c.LinkageType = new(ProcessingPosition7Choice)
	return c.LinkageType
}

// Identification of a linked corporate action event.
type CorporateActionEventReference4 struct {

	// Identification of the linked corporate action event.
	EventIdentification *CorporateActionEventReference4Choice `xml:"EvtId"`

	// Specifies when this corporate action event is to be processed relative to a linked corporate action event.
	LinkageType *ProcessingPosition10Choice `xml:"LkgTp,omitempty"`
}

func (c *CorporateActionEventReference4) AddEventIdentification() *CorporateActionEventReference4Choice {
	c.EventIdentification = new(CorporateActionEventReference4Choice)
	return c.EventIdentification
}

func (c *CorporateActionEventReference4) AddLinkageType() *ProcessingPosition10Choice {
	c.LinkageType = new(ProcessingPosition10Choice)
	return c.LinkageType
}
