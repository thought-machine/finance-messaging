package iso20022

// Provides the proprietary status and reason of an instruction or an instruction cancellation.
type ProprietaryStatusAndReason1 struct {

	// Proprietary identification of the status related to an instruction.
	ProprietaryStatus *GenericIdentification20 `xml:"PrtrySts"`

	// Proprietary identification of the reason related to a proprietary status.
	ProprietaryReason []*ProprietaryReason1 `xml:"PrtryRsn,omitempty"`
}

func (p *ProprietaryStatusAndReason1) AddProprietaryStatus() *GenericIdentification20 {
	p.ProprietaryStatus = new(GenericIdentification20)
	return p.ProprietaryStatus
}

func (p *ProprietaryStatusAndReason1) AddProprietaryReason() *ProprietaryReason1 {
	newValue := new(ProprietaryReason1)
	p.ProprietaryReason = append(p.ProprietaryReason, newValue)
	return newValue
}

// Proprietary status and reason of an instruction or an instruction cancellation.
type ProprietaryStatusAndReason5 struct {

	// Proprietary identification of the status of the instruction.
	Status *GenericIdentification36 `xml:"Sts"`

	// Proprietary identification of the reason for the status.
	Reason *ProprietaryReason1Choice `xml:"Rsn"`

	// Additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *ProprietaryStatusAndReason5) AddStatus() *GenericIdentification36 {
	p.Status = new(GenericIdentification36)
	return p.Status
}

func (p *ProprietaryStatusAndReason5) AddReason() *ProprietaryReason1Choice {
	p.Reason = new(ProprietaryReason1Choice)
	return p.Reason
}

func (p *ProprietaryStatusAndReason5) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Provides the proprietary status and reason of an instruction or an instruction cancellation.
type ProprietaryStatusAndReason6 struct {

	// Proprietary identification of the status related to an instruction.
	ProprietaryStatus *GenericIdentification30 `xml:"PrtrySts"`

	// Proprietary identification of the reason related to a proprietary status.
	ProprietaryReason []*ProprietaryReason4 `xml:"PrtryRsn,omitempty"`
}

func (p *ProprietaryStatusAndReason6) AddProprietaryStatus() *GenericIdentification30 {
	p.ProprietaryStatus = new(GenericIdentification30)
	return p.ProprietaryStatus
}

func (p *ProprietaryStatusAndReason6) AddProprietaryReason() *ProprietaryReason4 {
	newValue := new(ProprietaryReason4)
	p.ProprietaryReason = append(p.ProprietaryReason, newValue)
	return newValue
}

// Provides the proprietary status and reason of an instruction or an instruction cancellation.
type ProprietaryStatusAndReason7 struct {

	// Proprietary identification of the status related to an instruction.
	ProprietaryStatus *GenericIdentification47 `xml:"PrtrySts"`

	// Proprietary identification of the reason related to a proprietary status.
	ProprietaryReason []*ProprietaryReason5 `xml:"PrtryRsn,omitempty"`
}

func (p *ProprietaryStatusAndReason7) AddProprietaryStatus() *GenericIdentification47 {
	p.ProprietaryStatus = new(GenericIdentification47)
	return p.ProprietaryStatus
}

func (p *ProprietaryStatusAndReason7) AddProprietaryReason() *ProprietaryReason5 {
	newValue := new(ProprietaryReason5)
	p.ProprietaryReason = append(p.ProprietaryReason, newValue)
	return newValue
}
