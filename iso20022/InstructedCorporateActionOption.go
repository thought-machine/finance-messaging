package iso20022

// Provides corporate action option details about total instructed balance.
type InstructedCorporateActionOption1 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb,omitempty"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption2Choice `xml:"OptnTp"`

	// Balance of instructed position.
	InstructedBalance *BalanceFormat1Choice `xml:"InstdBal"`

	// Indicates the default action related to a corporate action event.
	DefaultAction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltActn,omitempty"`

	// Any deadline chosen by the account servicer (service level agreement).
	DeadlineDateTime *ISODateTime `xml:"DdlnDtTm"`

	// Specifies the type of deadline for instructing.
	DeadlineType *DeadlineCode1Choice `xml:"DdlnTp"`
}

func (i *InstructedCorporateActionOption1) SetOptionNumber(value string) {
	i.OptionNumber = (*Exact3NumericText)(&value)
}

func (i *InstructedCorporateActionOption1) AddOptionType() *CorporateActionOption2Choice {
	i.OptionType = new(CorporateActionOption2Choice)
	return i.OptionType
}

func (i *InstructedCorporateActionOption1) AddInstructedBalance() *BalanceFormat1Choice {
	i.InstructedBalance = new(BalanceFormat1Choice)
	return i.InstructedBalance
}

func (i *InstructedCorporateActionOption1) AddDefaultAction() *DefaultProcessingOrStandingInstruction1Choice {
	i.DefaultAction = new(DefaultProcessingOrStandingInstruction1Choice)
	return i.DefaultAction
}

func (i *InstructedCorporateActionOption1) SetDeadlineDateTime(value string) {
	i.DeadlineDateTime = (*ISODateTime)(&value)
}

func (i *InstructedCorporateActionOption1) AddDeadlineType() *DeadlineCode1Choice {
	i.DeadlineType = new(DeadlineCode1Choice)
	return i.DeadlineType
}

// Provides corporate action option details about total instructed balance.
type InstructedCorporateActionOption4 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb,omitempty"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption10Choice `xml:"OptnTp"`

	// Balance of instructed position.
	InstructedBalance *BalanceFormat1Choice `xml:"InstdBal"`

	// Indicates the default action related to a corporate action event.
	DefaultAction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltActn,omitempty"`

	// Any deadline chosen by the account servicer (service level agreement).
	DeadlineDateTime *ISODateTime `xml:"DdlnDtTm"`

	// Specifies the type of deadline for instructing.
	DeadlineType *DeadlineCode1Choice `xml:"DdlnTp"`
}

func (i *InstructedCorporateActionOption4) SetOptionNumber(value string) {
	i.OptionNumber = (*Exact3NumericText)(&value)
}

func (i *InstructedCorporateActionOption4) AddOptionType() *CorporateActionOption10Choice {
	i.OptionType = new(CorporateActionOption10Choice)
	return i.OptionType
}

func (i *InstructedCorporateActionOption4) AddInstructedBalance() *BalanceFormat1Choice {
	i.InstructedBalance = new(BalanceFormat1Choice)
	return i.InstructedBalance
}

func (i *InstructedCorporateActionOption4) AddDefaultAction() *DefaultProcessingOrStandingInstruction1Choice {
	i.DefaultAction = new(DefaultProcessingOrStandingInstruction1Choice)
	return i.DefaultAction
}

func (i *InstructedCorporateActionOption4) SetDeadlineDateTime(value string) {
	i.DeadlineDateTime = (*ISODateTime)(&value)
}

func (i *InstructedCorporateActionOption4) AddDeadlineType() *DeadlineCode1Choice {
	i.DeadlineType = new(DeadlineCode1Choice)
	return i.DeadlineType
}

// Provides corporate action option details about total instructed balance.
type InstructedCorporateActionOption6 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb,omitempty"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption18Choice `xml:"OptnTp"`

	// Balance of instructed position.
	InstructedBalance *BalanceFormat5Choice `xml:"InstdBal"`

	// Indicates the default action related to a corporate action event.
	DefaultAction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltActn,omitempty"`

	// Any deadline chosen by the account servicer (service level agreement).
	DeadlineDateTime *ISODateTime `xml:"DdlnDtTm"`

	// Specifies the type of deadline for instructing.
	DeadlineType *DeadlineCode3Choice `xml:"DdlnTp"`
}

func (i *InstructedCorporateActionOption6) SetOptionNumber(value string) {
	i.OptionNumber = (*Exact3NumericText)(&value)
}

func (i *InstructedCorporateActionOption6) AddOptionType() *CorporateActionOption18Choice {
	i.OptionType = new(CorporateActionOption18Choice)
	return i.OptionType
}

func (i *InstructedCorporateActionOption6) AddInstructedBalance() *BalanceFormat5Choice {
	i.InstructedBalance = new(BalanceFormat5Choice)
	return i.InstructedBalance
}

func (i *InstructedCorporateActionOption6) AddDefaultAction() *DefaultProcessingOrStandingInstruction1Choice {
	i.DefaultAction = new(DefaultProcessingOrStandingInstruction1Choice)
	return i.DefaultAction
}

func (i *InstructedCorporateActionOption6) SetDeadlineDateTime(value string) {
	i.DeadlineDateTime = (*ISODateTime)(&value)
}

func (i *InstructedCorporateActionOption6) AddDeadlineType() *DeadlineCode3Choice {
	i.DeadlineType = new(DeadlineCode3Choice)
	return i.DeadlineType
}

// Provides corporate action option details about total instructed balance.
type InstructedCorporateActionOption7 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb,omitempty"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption23Choice `xml:"OptnTp"`

	// Balance of instructed position.
	InstructedBalance *BalanceFormat7Choice `xml:"InstdBal"`

	// Indicates the default action related to a corporate action event.
	DefaultAction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltActn,omitempty"`

	// Any deadline chosen by the account servicer (service level agreement).
	DeadlineDateTime *ISODateTime `xml:"DdlnDtTm"`

	// Specifies the type of deadline for instructing.
	DeadlineType *DeadlineCode4Choice `xml:"DdlnTp"`
}

func (i *InstructedCorporateActionOption7) SetOptionNumber(value string) {
	i.OptionNumber = (*Exact3NumericText)(&value)
}

func (i *InstructedCorporateActionOption7) AddOptionType() *CorporateActionOption23Choice {
	i.OptionType = new(CorporateActionOption23Choice)
	return i.OptionType
}

func (i *InstructedCorporateActionOption7) AddInstructedBalance() *BalanceFormat7Choice {
	i.InstructedBalance = new(BalanceFormat7Choice)
	return i.InstructedBalance
}

func (i *InstructedCorporateActionOption7) AddDefaultAction() *DefaultProcessingOrStandingInstruction1Choice {
	i.DefaultAction = new(DefaultProcessingOrStandingInstruction1Choice)
	return i.DefaultAction
}

func (i *InstructedCorporateActionOption7) SetDeadlineDateTime(value string) {
	i.DeadlineDateTime = (*ISODateTime)(&value)
}

func (i *InstructedCorporateActionOption7) AddDeadlineType() *DeadlineCode4Choice {
	i.DeadlineType = new(DeadlineCode4Choice)
	return i.DeadlineType
}
