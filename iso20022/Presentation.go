package iso20022

// Information for the presentation of documents.
type Presentation1 struct {

	// Medium through which the presentation can be submitted such as paper, electronic or both.
	Medium *PresentationMedium1Choice `xml:"Mdm,omitempty"`

	// Choice of representation for the place of presentation.
	PlaceOfPresentationOrUnderConfirmationChoice *PlaceOrUnderConfirmationChoice1 `xml:"PlcOfPresntnOrUdrConfChc,omitempty"`

	// Document required to be presented.
	Document []*Document8 `xml:"Doc,omitempty"`

	// Additional information related to the presentation.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (p *Presentation1) AddMedium() *PresentationMedium1Choice {
	p.Medium = new(PresentationMedium1Choice)
	return p.Medium
}

func (p *Presentation1) AddPlaceOfPresentationOrUnderConfirmationChoice() *PlaceOrUnderConfirmationChoice1 {
	p.PlaceOfPresentationOrUnderConfirmationChoice = new(PlaceOrUnderConfirmationChoice1)
	return p.PlaceOfPresentationOrUnderConfirmationChoice
}

func (p *Presentation1) AddDocument() *Document8 {
	newValue := new(Document8)
	p.Document = append(p.Document, newValue)
	return newValue
}

func (p *Presentation1) AddAdditionalInformation(value string) {
	p.AdditionalInformation = append(p.AdditionalInformation, (*Max2000Text)(&value))
}

// Information for the presentation of documents.
type Presentation2 struct {

	// Party, other than beneficiary, forwarding the documents.
	Presenter *PartyIdentification43 `xml:"Presntr,omitempty"`

	// Date on which the beneficiary presented the demand.
	BeneficiaryPresentationDate *ISODate `xml:"BnfcryPresntnDt,omitempty"`
}

func (p *Presentation2) AddPresenter() *PartyIdentification43 {
	p.Presenter = new(PartyIdentification43)
	return p.Presenter
}

func (p *Presentation2) SetBeneficiaryPresentationDate(value string) {
	p.BeneficiaryPresentationDate = (*ISODate)(&value)
}

// Electronic presentation information.
type Presentation3 struct {

	// Format for presentation documents that are submitted electronically.
	Format *DocumentFormat1Choice `xml:"Frmt,omitempty"`

	// Channel through which presentation documents are submitted electronically, such as SWIFT,  Web upload, or secure email.
	Channel *Channel1Choice `xml:"Chanl,omitempty"`

	// Uniform Resource Identifier (URI), such as a web or an email address, specifying where the presentation can be addressed.
	Address *Max256Text `xml:"Adr,omitempty"`
}

func (p *Presentation3) AddFormat() *DocumentFormat1Choice {
	p.Format = new(DocumentFormat1Choice)
	return p.Format
}

func (p *Presentation3) AddChannel() *Channel1Choice {
	p.Channel = new(Channel1Choice)
	return p.Channel
}

func (p *Presentation3) SetAddress(value string) {
	p.Address = (*Max256Text)(&value)
}

// Information for the presentation of documents.
type Presentation4 struct {

	// Medium through which the presentation can be submitted such as paper, electronic or both.
	Medium *PresentationMedium1Choice `xml:"Mdm,omitempty"`

	// Document required to be presented.
	Document []*Document11 `xml:"Doc,omitempty"`

	// Additional information related to the presentation.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (p *Presentation4) AddMedium() *PresentationMedium1Choice {
	p.Medium = new(PresentationMedium1Choice)
	return p.Medium
}

func (p *Presentation4) AddDocument() *Document11 {
	newValue := new(Document11)
	p.Document = append(p.Document, newValue)
	return newValue
}

func (p *Presentation4) AddAdditionalInformation(value string) {
	p.AdditionalInformation = append(p.AdditionalInformation, (*Max2000Text)(&value))
}
