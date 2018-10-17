package iso20022

// Specifies the voting entitlement.
type EligiblePosition2 struct {

	// Identification of the securities account.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Identifies party that legally owns the account.
	AccountOwner *PartyIdentification9Choice `xml:"AcctOwnr,omitempty"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, eg, sub-balance per status.
	HoldingBalance []*HoldingBalance3 `xml:"HldgBal,omitempty"`

	// Identifies owner of the voting rights.
	RightsHolder []*PartyIdentification9Choice `xml:"RghtsHldr,omitempty"`
}

func (e *EligiblePosition2) SetAccountIdentification(value string) {
	e.AccountIdentification = (*Max35Text)(&value)
}

func (e *EligiblePosition2) AddAccountOwner() *PartyIdentification9Choice {
	e.AccountOwner = new(PartyIdentification9Choice)
	return e.AccountOwner
}

func (e *EligiblePosition2) AddHoldingBalance() *HoldingBalance3 {
	newValue := new(HoldingBalance3)
	e.HoldingBalance = append(e.HoldingBalance, newValue)
	return newValue
}

func (e *EligiblePosition2) AddRightsHolder() *PartyIdentification9Choice {
	newValue := new(PartyIdentification9Choice)
	e.RightsHolder = append(e.RightsHolder, newValue)
	return newValue
}

// Specifies the voting entitlement.
type EligiblePosition3 struct {

	// Identification of the securities account.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Identifies party that legally owns the account.
	AccountOwner *PartyIdentification9Choice `xml:"AcctOwnr,omitempty"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, eg, sub-balance per status.
	HoldingBalance []*HoldingBalance6 `xml:"HldgBal,omitempty"`

	// Identifies owner of the voting rights.
	RightsHolder []*PartyIdentification9Choice `xml:"RghtsHldr,omitempty"`
}

func (e *EligiblePosition3) SetAccountIdentification(value string) {
	e.AccountIdentification = (*Max35Text)(&value)
}

func (e *EligiblePosition3) AddAccountOwner() *PartyIdentification9Choice {
	e.AccountOwner = new(PartyIdentification9Choice)
	return e.AccountOwner
}

func (e *EligiblePosition3) AddHoldingBalance() *HoldingBalance6 {
	newValue := new(HoldingBalance6)
	e.HoldingBalance = append(e.HoldingBalance, newValue)
	return newValue
}

func (e *EligiblePosition3) AddRightsHolder() *PartyIdentification9Choice {
	newValue := new(PartyIdentification9Choice)
	e.RightsHolder = append(e.RightsHolder, newValue)
	return newValue
}

// Specifies the voting entitlement.
type EligiblePosition4 struct {

	// Identification of the securities account.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Identifies party that legally owns the account.
	AccountOwner *PartyIdentification39 `xml:"AcctOwnr,omitempty"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, eg, sub-balance per status.
	HoldingBalance []*HoldingBalance6 `xml:"HldgBal,omitempty"`

	// Identifies owner of the voting rights.
	RightsHolder []*PartyIdentification39 `xml:"RghtsHldr,omitempty"`
}

func (e *EligiblePosition4) SetAccountIdentification(value string) {
	e.AccountIdentification = (*Max35Text)(&value)
}

func (e *EligiblePosition4) AddAccountOwner() *PartyIdentification39 {
	e.AccountOwner = new(PartyIdentification39)
	return e.AccountOwner
}

func (e *EligiblePosition4) AddHoldingBalance() *HoldingBalance6 {
	newValue := new(HoldingBalance6)
	e.HoldingBalance = append(e.HoldingBalance, newValue)
	return newValue
}

func (e *EligiblePosition4) AddRightsHolder() *PartyIdentification39 {
	newValue := new(PartyIdentification39)
	e.RightsHolder = append(e.RightsHolder, newValue)
	return newValue
}

// Specifies the voting entitlement.
type EligiblePosition5 struct {

	// Identification of the securities account.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Identifies the party that legally owns the account.
	AccountOwner *PartyIdentification40Choice `xml:"AcctOwnr,omitempty"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, for example, sub-balance per status.
	HoldingBalance []*HoldingBalance7 `xml:"HldgBal,omitempty"`

	// Identifies the owner of the voting rights.
	RightsHolder []*PartyIdentification40Choice `xml:"RghtsHldr,omitempty"`
}

func (e *EligiblePosition5) SetAccountIdentification(value string) {
	e.AccountIdentification = (*Max35Text)(&value)
}

func (e *EligiblePosition5) AddAccountOwner() *PartyIdentification40Choice {
	e.AccountOwner = new(PartyIdentification40Choice)
	return e.AccountOwner
}

func (e *EligiblePosition5) AddHoldingBalance() *HoldingBalance7 {
	newValue := new(HoldingBalance7)
	e.HoldingBalance = append(e.HoldingBalance, newValue)
	return newValue
}

func (e *EligiblePosition5) AddRightsHolder() *PartyIdentification40Choice {
	newValue := new(PartyIdentification40Choice)
	e.RightsHolder = append(e.RightsHolder, newValue)
	return newValue
}

// Specifies the voting entitlement.
type EligiblePosition6 struct {

	// Identification of the securities account.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Identifies the party that legally owns the account.
	AccountOwner *PartyIdentification71 `xml:"AcctOwnr,omitempty"`

	// Net position of a segregated holding of a single security within the overall position held in a securities account, for example, sub-balance per status.
	HoldingBalance []*HoldingBalance7 `xml:"HldgBal,omitempty"`

	// Identifies the owner of the voting rights.
	RightsHolder []*PartyIdentification71 `xml:"RghtsHldr,omitempty"`
}

func (e *EligiblePosition6) SetAccountIdentification(value string) {
	e.AccountIdentification = (*Max35Text)(&value)
}

func (e *EligiblePosition6) AddAccountOwner() *PartyIdentification71 {
	e.AccountOwner = new(PartyIdentification71)
	return e.AccountOwner
}

func (e *EligiblePosition6) AddHoldingBalance() *HoldingBalance7 {
	newValue := new(HoldingBalance7)
	e.HoldingBalance = append(e.HoldingBalance, newValue)
	return newValue
}

func (e *EligiblePosition6) AddRightsHolder() *PartyIdentification71 {
	newValue := new(PartyIdentification71)
	e.RightsHolder = append(e.RightsHolder, newValue)
	return newValue
}
