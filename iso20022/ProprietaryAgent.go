package iso20022

// Set of elements to identify a proprietary party.
type ProprietaryAgent1 struct {

	// Identifies the type of proprietary agent reported.
	Type *Max35Text `xml:"Tp"`

	// Proprietary agent.
	Agent *BranchAndFinancialInstitutionIdentification3 `xml:"Agt"`
}

func (p *ProprietaryAgent1) SetType(value string) {
	p.Type = (*Max35Text)(&value)
}

func (p *ProprietaryAgent1) AddAgent() *BranchAndFinancialInstitutionIdentification3 {
	p.Agent = new(BranchAndFinancialInstitutionIdentification3)
	return p.Agent
}

// Set of elements used to identify a proprietary party.
type ProprietaryAgent2 struct {

	// Specifies the type of proprietary agent.
	Type *Max35Text `xml:"Tp"`

	// Organisation established primarily to provide financial services.
	Agent *BranchAndFinancialInstitutionIdentification4 `xml:"Agt"`
}

func (p *ProprietaryAgent2) SetType(value string) {
	p.Type = (*Max35Text)(&value)
}

func (p *ProprietaryAgent2) AddAgent() *BranchAndFinancialInstitutionIdentification4 {
	p.Agent = new(BranchAndFinancialInstitutionIdentification4)
	return p.Agent
}

// Identifies a proprietary party.
type ProprietaryAgent3 struct {

	// Specifies the type of proprietary agent.
	Type *Max35Text `xml:"Tp"`

	// Organisation established primarily to provide financial services.
	Agent *BranchAndFinancialInstitutionIdentification5 `xml:"Agt"`
}

func (p *ProprietaryAgent3) SetType(value string) {
	p.Type = (*Max35Text)(&value)
}

func (p *ProprietaryAgent3) AddAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.Agent = new(BranchAndFinancialInstitutionIdentification5)
	return p.Agent
}
