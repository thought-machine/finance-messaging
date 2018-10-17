package iso20022

// Status applying to individual instructions of a MeetingInstruction.
type DetailedInstructionStatus1 struct {

	// Identifies the detailed instruction within an instruction message.
	InstructionIdentification *Max35Text `xml:"InstrId"`

	// Identifies the safekeeping account.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Identifies the subaccount of the safekeeping account.
	SubAccountIdentification *Max35Text `xml:"SubAcctId,omitempty"`

	// Status applying to individual instructions of a MeetingInstruction.
	InstructionStatus *InstructionStatus2Choice `xml:"InstrSts"`
}

func (d *DetailedInstructionStatus1) SetInstructionIdentification(value string) {
	d.InstructionIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus1) SetAccountIdentification(value string) {
	d.AccountIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus1) SetSubAccountIdentification(value string) {
	d.SubAccountIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus1) AddInstructionStatus() *InstructionStatus2Choice {
	d.InstructionStatus = new(InstructionStatus2Choice)
	return d.InstructionStatus
}

// Status applying to individual instructions of a MeetingInstruction.
type DetailedInstructionStatus10 struct {

	// Identifies the detailed instruction with an instruction message.
	InstructionIdentification *Max35Text `xml:"InstrId"`

	// Identifies the safekeeping account.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification40Choice `xml:"AcctOwnr,omitempty"`

	// Identifies the subaccount of the safekeeping account.
	SubAccountIdentification *Max35Text `xml:"SubAcctId,omitempty"`

	// Owner of the voting rights.
	RightsHolder []*PartyIdentification40Choice `xml:"RghtsHldr,omitempty"`

	// Indicates whether standing instructions have been applied or not.
	StandingInstruction *YesNoIndicator `xml:"StgInstr"`

	// Details of the vote.
	VotePerResolution []*Vote6 `xml:"VotePerRsltn,omitempty"`
}

func (d *DetailedInstructionStatus10) SetInstructionIdentification(value string) {
	d.InstructionIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus10) SetAccountIdentification(value string) {
	d.AccountIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus10) AddAccountOwner() *PartyIdentification40Choice {
	d.AccountOwner = new(PartyIdentification40Choice)
	return d.AccountOwner
}

func (d *DetailedInstructionStatus10) SetSubAccountIdentification(value string) {
	d.SubAccountIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus10) AddRightsHolder() *PartyIdentification40Choice {
	newValue := new(PartyIdentification40Choice)
	d.RightsHolder = append(d.RightsHolder, newValue)
	return newValue
}

func (d *DetailedInstructionStatus10) SetStandingInstruction(value string) {
	d.StandingInstruction = (*YesNoIndicator)(&value)
}

func (d *DetailedInstructionStatus10) AddVotePerResolution() *Vote6 {
	newValue := new(Vote6)
	d.VotePerResolution = append(d.VotePerResolution, newValue)
	return newValue
}

// Status applying to individual instructions of a MeetingInstruction.
type DetailedInstructionStatus11 struct {

	// Identifies the detailed instruction within an instruction message.
	InstructionIdentification *Max35Text `xml:"InstrId"`

	// Identifies the safekeeping account.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Identifies the subaccount of the safekeeping account.
	SubAccountIdentification *Max35Text `xml:"SubAcctId,omitempty"`

	// Status applying to individual instructions of a MeetingInstruction.
	InstructionStatus *InstructionStatus6Choice `xml:"InstrSts"`
}

func (d *DetailedInstructionStatus11) SetInstructionIdentification(value string) {
	d.InstructionIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus11) SetAccountIdentification(value string) {
	d.AccountIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus11) SetSubAccountIdentification(value string) {
	d.SubAccountIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus11) AddInstructionStatus() *InstructionStatus6Choice {
	d.InstructionStatus = new(InstructionStatus6Choice)
	return d.InstructionStatus
}

// Status applying to individual instructions of a MeetingInstruction.
type DetailedInstructionStatus2 struct {

	// Identifies the detailed instruction with an instruction message.
	InstructionIdentification *Max35Text `xml:"InstrId"`

	// Identifies the safekeeping account.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification9Choice `xml:"AcctOwnr,omitempty"`

	// Identifies the subaccount of the safekeeping account.
	SubAccountIdentification *Max35Text `xml:"SubAcctId,omitempty"`

	// Owner of the voting rights.
	RightsHolder []*PartyIdentification9Choice `xml:"RghtsHldr,omitempty"`

	// Indicates whether standing instructions have been applied or not.
	StandingInstruction *YesNoIndicator `xml:"StgInstr"`

	// Details of the vote.
	VotePerResolution []*Vote4 `xml:"VotePerRsltn"`
}

func (d *DetailedInstructionStatus2) SetInstructionIdentification(value string) {
	d.InstructionIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus2) SetAccountIdentification(value string) {
	d.AccountIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus2) AddAccountOwner() *PartyIdentification9Choice {
	d.AccountOwner = new(PartyIdentification9Choice)
	return d.AccountOwner
}

func (d *DetailedInstructionStatus2) SetSubAccountIdentification(value string) {
	d.SubAccountIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus2) AddRightsHolder() *PartyIdentification9Choice {
	newValue := new(PartyIdentification9Choice)
	d.RightsHolder = append(d.RightsHolder, newValue)
	return newValue
}

func (d *DetailedInstructionStatus2) SetStandingInstruction(value string) {
	d.StandingInstruction = (*YesNoIndicator)(&value)
}

func (d *DetailedInstructionStatus2) AddVotePerResolution() *Vote4 {
	newValue := new(Vote4)
	d.VotePerResolution = append(d.VotePerResolution, newValue)
	return newValue
}

// Status applying to individual instructions of a MeetingInstruction.
type DetailedInstructionStatus8 struct {

	// Identifies the detailed instruction within an instruction message.
	InstructionIdentification *Max35Text `xml:"InstrId"`

	// Identifies the safekeeping account.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Identifies the subaccount of the safekeeping account.
	SubAccountIdentification *Max35Text `xml:"SubAcctId,omitempty"`

	// Status applying to individual instructions of a MeetingInstruction.
	InstructionStatus *InstructionStatus4Choice `xml:"InstrSts"`
}

func (d *DetailedInstructionStatus8) SetInstructionIdentification(value string) {
	d.InstructionIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus8) SetAccountIdentification(value string) {
	d.AccountIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus8) SetSubAccountIdentification(value string) {
	d.SubAccountIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus8) AddInstructionStatus() *InstructionStatus4Choice {
	d.InstructionStatus = new(InstructionStatus4Choice)
	return d.InstructionStatus
}

// Status applying to individual instructions of a MeetingInstruction.
type DetailedInstructionStatus9 struct {

	// Identifies the detailed instruction with an instruction message.
	InstructionIdentification *Max35Text `xml:"InstrId"`

	// Identifies the safekeeping account.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification9Choice `xml:"AcctOwnr,omitempty"`

	// Identifies the subaccount of the safekeeping account.
	SubAccountIdentification *Max35Text `xml:"SubAcctId,omitempty"`

	// Owner of the voting rights.
	RightsHolder []*PartyIdentification9Choice `xml:"RghtsHldr,omitempty"`

	// Indicates whether standing instructions have been applied or not.
	StandingInstruction *YesNoIndicator `xml:"StgInstr"`

	// Details of the vote.
	VotePerResolution []*Vote4 `xml:"VotePerRsltn"`
}

func (d *DetailedInstructionStatus9) SetInstructionIdentification(value string) {
	d.InstructionIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus9) SetAccountIdentification(value string) {
	d.AccountIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus9) AddAccountOwner() *PartyIdentification9Choice {
	d.AccountOwner = new(PartyIdentification9Choice)
	return d.AccountOwner
}

func (d *DetailedInstructionStatus9) SetSubAccountIdentification(value string) {
	d.SubAccountIdentification = (*Max35Text)(&value)
}

func (d *DetailedInstructionStatus9) AddRightsHolder() *PartyIdentification9Choice {
	newValue := new(PartyIdentification9Choice)
	d.RightsHolder = append(d.RightsHolder, newValue)
	return newValue
}

func (d *DetailedInstructionStatus9) SetStandingInstruction(value string) {
	d.StandingInstruction = (*YesNoIndicator)(&value)
}

func (d *DetailedInstructionStatus9) AddVotePerResolution() *Vote4 {
	newValue := new(Vote4)
	d.VotePerResolution = append(d.VotePerResolution, newValue)
	return newValue
}
