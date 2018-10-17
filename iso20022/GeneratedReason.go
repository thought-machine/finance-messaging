package iso20022

// The status of an instruction, advice or request.
type GeneratedReason1 struct {

	// Specifies the reason why the transaction was generated.
	Code *GeneratedReasons1Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (g *GeneratedReason1) AddCode() *GeneratedReasons1Choice {
	g.Code = new(GeneratedReasons1Choice)
	return g.Code
}

func (g *GeneratedReason1) SetAdditionalReasonInformation(value string) {
	g.AdditionalReasonInformation = (*Max210Text)(&value)
}

// The status of an instruction, advice or request.
type GeneratedReason3 struct {

	// Specifies the reason why the transaction was generated.
	Code *GeneratedReasons3Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (g *GeneratedReason3) AddCode() *GeneratedReasons3Choice {
	g.Code = new(GeneratedReasons3Choice)
	return g.Code
}

func (g *GeneratedReason3) SetAdditionalReasonInformation(value string) {
	g.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the transaction was generated.
type GeneratedReason5 struct {

	// Specifies the reason why the transaction was generated.
	Code *GeneratedReasons5Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (g *GeneratedReason5) AddCode() *GeneratedReasons5Choice {
	g.Code = new(GeneratedReasons5Choice)
	return g.Code
}

func (g *GeneratedReason5) SetAdditionalReasonInformation(value string) {
	g.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the transaction was generated.
type GeneratedReason6 struct {

	// Specifies the reason why the transaction was generated.
	Code *GeneratedReasons6Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (g *GeneratedReason6) AddCode() *GeneratedReasons6Choice {
	g.Code = new(GeneratedReasons6Choice)
	return g.Code
}

func (g *GeneratedReason6) SetAdditionalReasonInformation(value string) {
	g.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
