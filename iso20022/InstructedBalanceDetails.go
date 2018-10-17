package iso20022

// Provides information about total instructed balance.
type InstructedBalanceDetails1 struct {

	// Provides information about the total instructed balance.
	TotalInstructedBalance *BalanceFormat1Choice `xml:"TtlInstdBal"`

	// Provide instructed balance breakdown information per option.
	OptionDetails []*InstructedCorporateActionOption1 `xml:"OptnDtls,omitempty"`
}

func (i *InstructedBalanceDetails1) AddTotalInstructedBalance() *BalanceFormat1Choice {
	i.TotalInstructedBalance = new(BalanceFormat1Choice)
	return i.TotalInstructedBalance
}

func (i *InstructedBalanceDetails1) AddOptionDetails() *InstructedCorporateActionOption1 {
	newValue := new(InstructedCorporateActionOption1)
	i.OptionDetails = append(i.OptionDetails, newValue)
	return newValue
}

// Provides information about total instructed balance.
type InstructedBalanceDetails3 struct {

	// Provides information about the total instructed balance.
	TotalInstructedBalance *BalanceFormat1Choice `xml:"TtlInstdBal"`

	// Provide instructed balance breakdown information per option.
	OptionDetails []*InstructedCorporateActionOption4 `xml:"OptnDtls,omitempty"`
}

func (i *InstructedBalanceDetails3) AddTotalInstructedBalance() *BalanceFormat1Choice {
	i.TotalInstructedBalance = new(BalanceFormat1Choice)
	return i.TotalInstructedBalance
}

func (i *InstructedBalanceDetails3) AddOptionDetails() *InstructedCorporateActionOption4 {
	newValue := new(InstructedCorporateActionOption4)
	i.OptionDetails = append(i.OptionDetails, newValue)
	return newValue
}

// Provides information about total instructed balance.
type InstructedBalanceDetails5 struct {

	// Provides information about the total instructed balance.
	TotalInstructedBalance *BalanceFormat5Choice `xml:"TtlInstdBal"`

	// Provide instructed balance breakdown information per option.
	OptionDetails []*InstructedCorporateActionOption6 `xml:"OptnDtls,omitempty"`
}

func (i *InstructedBalanceDetails5) AddTotalInstructedBalance() *BalanceFormat5Choice {
	i.TotalInstructedBalance = new(BalanceFormat5Choice)
	return i.TotalInstructedBalance
}

func (i *InstructedBalanceDetails5) AddOptionDetails() *InstructedCorporateActionOption6 {
	newValue := new(InstructedCorporateActionOption6)
	i.OptionDetails = append(i.OptionDetails, newValue)
	return newValue
}

// Provides information about total instructed balance.
type InstructedBalanceDetails6 struct {

	// Provides information about the total instructed balance.
	TotalInstructedBalance *BalanceFormat7Choice `xml:"TtlInstdBal"`

	// Provide instructed balance breakdown information per option.
	OptionDetails []*InstructedCorporateActionOption7 `xml:"OptnDtls,omitempty"`
}

func (i *InstructedBalanceDetails6) AddTotalInstructedBalance() *BalanceFormat7Choice {
	i.TotalInstructedBalance = new(BalanceFormat7Choice)
	return i.TotalInstructedBalance
}

func (i *InstructedBalanceDetails6) AddOptionDetails() *InstructedCorporateActionOption7 {
	newValue := new(InstructedCorporateActionOption7)
	i.OptionDetails = append(i.OptionDetails, newValue)
	return newValue
}
