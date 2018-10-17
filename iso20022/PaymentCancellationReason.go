package iso20022

// Provides further details on the reason of the mandate cancellation request.
type PaymentCancellationReason1 struct {

	// Party that issues the cancellation request.
	Originator *PartyIdentification43 `xml:"Orgtr,omitempty"`

	// Specifies the reason for the cancellation request.
	Reason *MandateReason1Choice `xml:"Rsn"`

	// Further details on the cancellation request reason.
	AdditionalInformation []*Max105Text `xml:"AddtlInf,omitempty"`
}

func (p *PaymentCancellationReason1) AddOriginator() *PartyIdentification43 {
	p.Originator = new(PartyIdentification43)
	return p.Originator
}

func (p *PaymentCancellationReason1) AddReason() *MandateReason1Choice {
	p.Reason = new(MandateReason1Choice)
	return p.Reason
}

func (p *PaymentCancellationReason1) AddAdditionalInformation(value string) {
	p.AdditionalInformation = append(p.AdditionalInformation, (*Max105Text)(&value))
}

// Provides further details on the reason of the cancellation request.
type PaymentCancellationReason2 struct {

	// Party that issues the cancellation request.
	Originator *PartyIdentification43 `xml:"Orgtr,omitempty"`

	// Specifies the reason for the cancellation.
	Reason *CancellationReason14Choice `xml:"Rsn,omitempty"`

	// Further details on the cancellation request reason.
	AdditionalInformation []*Max105Text `xml:"AddtlInf,omitempty"`
}

func (p *PaymentCancellationReason2) AddOriginator() *PartyIdentification43 {
	p.Originator = new(PartyIdentification43)
	return p.Originator
}

func (p *PaymentCancellationReason2) AddReason() *CancellationReason14Choice {
	p.Reason = new(CancellationReason14Choice)
	return p.Reason
}

func (p *PaymentCancellationReason2) AddAdditionalInformation(value string) {
	p.AdditionalInformation = append(p.AdditionalInformation, (*Max105Text)(&value))
}

// Provides further details on the reason of the cancellation request.
type PaymentCancellationReason3 struct {

	// Party that issues the cancellation request.
	Originator *PartyIdentification43 `xml:"Orgtr,omitempty"`

	// Specifies the reason for the cancellation.
	Reason *CancellationReason33Choice `xml:"Rsn,omitempty"`

	// Further details on the cancellation request reason.
	AdditionalInformation []*Max105Text `xml:"AddtlInf,omitempty"`
}

func (p *PaymentCancellationReason3) AddOriginator() *PartyIdentification43 {
	p.Originator = new(PartyIdentification43)
	return p.Originator
}

func (p *PaymentCancellationReason3) AddReason() *CancellationReason33Choice {
	p.Reason = new(CancellationReason33Choice)
	return p.Reason
}

func (p *PaymentCancellationReason3) AddAdditionalInformation(value string) {
	p.AdditionalInformation = append(p.AdditionalInformation, (*Max105Text)(&value))
}
