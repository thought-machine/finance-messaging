package iso20022

// Specifies the role of the Issuer agent.
type IssuerAgent1 struct {

	// Identifies Issuer Agent.
	Identification *PartyIdentification9Choice `xml:"Id"`

	// Specifies the role of the Issuer agent.
	Role *AgentRole1Code `xml:"Role,omitempty"`
}

func (i *IssuerAgent1) AddIdentification() *PartyIdentification9Choice {
	i.Identification = new(PartyIdentification9Choice)
	return i.Identification
}

func (i *IssuerAgent1) SetRole(value string) {
	i.Role = (*AgentRole1Code)(&value)
}

// Specifies the role of the issuer agent.
type IssuerAgent2 struct {

	// Identifies issuer agent.
	Identification *PartyIdentification40Choice `xml:"Id"`

	// Specifies the role of the issuer agent.
	Role *AgentRole1Code `xml:"Role,omitempty"`
}

func (i *IssuerAgent2) AddIdentification() *PartyIdentification40Choice {
	i.Identification = new(PartyIdentification40Choice)
	return i.Identification
}

func (i *IssuerAgent2) SetRole(value string) {
	i.Role = (*AgentRole1Code)(&value)
}
