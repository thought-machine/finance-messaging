package iso20022

// Decision of the voting party for one resolution.
type Vote3 struct {

	// Numbering of the resolution as specified by the issuer or  its agent.
	IssuerLabel *Max35Text `xml:"IssrLabl"`

	// Specifies the different instructions which can be used to vote.
	VoteOption *VoteInstruction2Code `xml:"VoteOptn"`
}

func (v *Vote3) SetIssuerLabel(value string) {
	v.IssuerLabel = (*Max35Text)(&value)
}

func (v *Vote3) SetVoteOption(value string) {
	v.VoteOption = (*VoteInstruction2Code)(&value)
}

// Decision of the voting party for one resolution. Several types of decisions can be indicated to allow for split vote specification.
type Vote4 struct {

	// Numbering of the resolution as specified by the issuer or  its agent.
	IssuerLabel *Max35Text `xml:"IssrLabl"`

	// Number of votes in favour of one resolution
	For *Number `xml:"For,omitempty"`

	// Number of votes against one resolution
	Against *Number `xml:"Agnst,omitempty"`

	// Number of votes expressed as abstain for one resolution.
	Abstain *Number `xml:"Abstn,omitempty"`

	// Number of votes withheld for one resolution
	Withhold *Number `xml:"Wthhld,omitempty"`

	// Number of votes in line with the votes of the management.
	WithManagement *Number `xml:"WthMgmt,omitempty"`

	// Number of votes against the voting recommendation of the management.
	AgainstManagement *Number `xml:"AgnstMgmt,omitempty"`

	// Number of votes for which decision is left to the party that will exercise the voting right.
	Discretionary *Number `xml:"Dscrtnry,omitempty"`

	// Number of votes for which no action has been taken.
	NoAction *Number `xml:"NoActn,omitempty"`
}

func (v *Vote4) SetIssuerLabel(value string) {
	v.IssuerLabel = (*Max35Text)(&value)
}

func (v *Vote4) SetFor(value string) {
	v.For = (*Number)(&value)
}

func (v *Vote4) SetAgainst(value string) {
	v.Against = (*Number)(&value)
}

func (v *Vote4) SetAbstain(value string) {
	v.Abstain = (*Number)(&value)
}

func (v *Vote4) SetWithhold(value string) {
	v.Withhold = (*Number)(&value)
}

func (v *Vote4) SetWithManagement(value string) {
	v.WithManagement = (*Number)(&value)
}

func (v *Vote4) SetAgainstManagement(value string) {
	v.AgainstManagement = (*Number)(&value)
}

func (v *Vote4) SetDiscretionary(value string) {
	v.Discretionary = (*Number)(&value)
}

func (v *Vote4) SetNoAction(value string) {
	v.NoAction = (*Number)(&value)
}

// Indicates the number of voting rights cast to a resolution.
type Vote5 struct {

	// Numbering of the resolution as specified by the issuer or its agent.
	IssuerLabel *Max35Text `xml:"IssrLabl"`

	// Specifies whether a resolution is accepted or not.
	Accepted *YesNoIndicator `xml:"Accptd"`

	// Number of votes in favour of one resolution
	For *Number `xml:"For,omitempty"`

	// Number of votes against one resolution
	Against *Number `xml:"Agnst,omitempty"`

	// Number of votes expressed as abstain.
	Abstain *Number `xml:"Abstn,omitempty"`

	// Total votes withheld, eg in the case where a shareholder wishes not to endorse the election of a board member.
	Withhold *Number `xml:"Wthhld,omitempty"`

	// Number of votes for which no action has been taken.
	NoAction *Number `xml:"NoActn,omitempty"`
}

func (v *Vote5) SetIssuerLabel(value string) {
	v.IssuerLabel = (*Max35Text)(&value)
}

func (v *Vote5) SetAccepted(value string) {
	v.Accepted = (*YesNoIndicator)(&value)
}

func (v *Vote5) SetFor(value string) {
	v.For = (*Number)(&value)
}

func (v *Vote5) SetAgainst(value string) {
	v.Against = (*Number)(&value)
}

func (v *Vote5) SetAbstain(value string) {
	v.Abstain = (*Number)(&value)
}

func (v *Vote5) SetWithhold(value string) {
	v.Withhold = (*Number)(&value)
}

func (v *Vote5) SetNoAction(value string) {
	v.NoAction = (*Number)(&value)
}

// Decision of the voting party for one resolution. Several types of decisions can be indicated to allow for split vote specification.
type Vote6 struct {

	// Numbering of the resolution as specified by the issuer or  its agent.
	IssuerLabel *Max35Text `xml:"IssrLabl"`

	// Number of votes in favour of one resolution.
	For *Number `xml:"For,omitempty"`

	// Number of votes against one resolution.
	Against *Number `xml:"Agnst,omitempty"`

	// Number of votes expressed as abstain for one resolution.
	Abstain *Number `xml:"Abstn,omitempty"`

	// Number of votes withheld for one resolution.
	Withhold *Number `xml:"Wthhld,omitempty"`

	// Number of votes in line with the votes of the management.
	WithManagement *Number `xml:"WthMgmt,omitempty"`

	// Number of votes against the voting recommendation of the management.
	AgainstManagement *Number `xml:"AgnstMgmt,omitempty"`

	// Number of votes for which decision is left to the party that will exercise the voting right.
	Discretionary *Number `xml:"Dscrtnry,omitempty"`

	// Number of votes in favour for one year for "say on pay" type of resolution.
	OneYear *Number `xml:"OneYr,omitempty"`

	// Number of votes in favour of two years for "say on pay" type of resolution.
	TwoYears *Number `xml:"TwoYrs,omitempty"`

	// Number of votes in favour of three years for "say on pay" type of resolution.
	ThreeYears *Number `xml:"ThreeYrs,omitempty"`

	// Number of votes for which no action has been taken.
	NoAction *Number `xml:"NoActn,omitempty"`

	// Resolution withdrawn at the meeting.
	Withdrawn *YesNoIndicator `xml:"Wdrwn,omitempty"`
}

func (v *Vote6) SetIssuerLabel(value string) {
	v.IssuerLabel = (*Max35Text)(&value)
}

func (v *Vote6) SetFor(value string) {
	v.For = (*Number)(&value)
}

func (v *Vote6) SetAgainst(value string) {
	v.Against = (*Number)(&value)
}

func (v *Vote6) SetAbstain(value string) {
	v.Abstain = (*Number)(&value)
}

func (v *Vote6) SetWithhold(value string) {
	v.Withhold = (*Number)(&value)
}

func (v *Vote6) SetWithManagement(value string) {
	v.WithManagement = (*Number)(&value)
}

func (v *Vote6) SetAgainstManagement(value string) {
	v.AgainstManagement = (*Number)(&value)
}

func (v *Vote6) SetDiscretionary(value string) {
	v.Discretionary = (*Number)(&value)
}

func (v *Vote6) SetOneYear(value string) {
	v.OneYear = (*Number)(&value)
}

func (v *Vote6) SetTwoYears(value string) {
	v.TwoYears = (*Number)(&value)
}

func (v *Vote6) SetThreeYears(value string) {
	v.ThreeYears = (*Number)(&value)
}

func (v *Vote6) SetNoAction(value string) {
	v.NoAction = (*Number)(&value)
}

func (v *Vote6) SetWithdrawn(value string) {
	v.Withdrawn = (*YesNoIndicator)(&value)
}

// Indicates the number of voting rights cast to a resolution.
type Vote7 struct {

	// Numbering of the resolution as specified by the issuer or its agent.
	IssuerLabel *Max35Text `xml:"IssrLabl"`

	// Specifies the acceptance status of a resolution.
	ResolutionStatus *ResolutionStatus2Code `xml:"RsltnSts"`

	// Number of votes in favour of one resolution.
	For *Number `xml:"For,omitempty"`

	// Number of votes against one resolution.
	Against *Number `xml:"Agnst,omitempty"`

	// Number of votes expressed as abstain.
	Abstain *Number `xml:"Abstn,omitempty"`

	// Total votes withheld, for example, in the case where a shareholder wishes not to endorse the election of a board member.
	Withhold *Number `xml:"Wthhld,omitempty"`

	// Number of votes in favour for one year for "say on pay" type of resolution.
	OneYear *Number `xml:"OneYr,omitempty"`

	// Number of votes in favour of two years for "say on pay" type of resolution.
	TwoYears *Number `xml:"TwoYrs,omitempty"`

	// Number of votes in favour of three years for "say on pay" type of resolution.
	ThreeYears *Number `xml:"ThreeYrs,omitempty"`

	// Number of votes for which no action has been taken.
	NoAction *Number `xml:"NoActn,omitempty"`
}

func (v *Vote7) SetIssuerLabel(value string) {
	v.IssuerLabel = (*Max35Text)(&value)
}

func (v *Vote7) SetResolutionStatus(value string) {
	v.ResolutionStatus = (*ResolutionStatus2Code)(&value)
}

func (v *Vote7) SetFor(value string) {
	v.For = (*Number)(&value)
}

func (v *Vote7) SetAgainst(value string) {
	v.Against = (*Number)(&value)
}

func (v *Vote7) SetAbstain(value string) {
	v.Abstain = (*Number)(&value)
}

func (v *Vote7) SetWithhold(value string) {
	v.Withhold = (*Number)(&value)
}

func (v *Vote7) SetOneYear(value string) {
	v.OneYear = (*Number)(&value)
}

func (v *Vote7) SetTwoYears(value string) {
	v.TwoYears = (*Number)(&value)
}

func (v *Vote7) SetThreeYears(value string) {
	v.ThreeYears = (*Number)(&value)
}

func (v *Vote7) SetNoAction(value string) {
	v.NoAction = (*Number)(&value)
}

// Decision of the voting party for one resolution. Several types of decisions can be indicated to allow for split vote specification.
type Vote8 struct {

	// Numbering of the resolution as specified by the issuer or  its agent.
	IssuerLabel *Max35Text `xml:"IssrLabl"`

	// Number of votes in favour of one resolution.
	For *Number `xml:"For,omitempty"`

	// Number of votes against one resolution
	Against *Number `xml:"Agnst,omitempty"`

	// Number of votes expressed as abstain for one resolution.
	Abstain *Number `xml:"Abstn,omitempty"`

	// Number of votes withheld for one resolution.
	Withhold *Number `xml:"Wthhld,omitempty"`

	// Number of votes in line with the votes of the management.
	WithManagement *Number `xml:"WthMgmt,omitempty"`

	// Number of votes against the voting recommendation of the management.
	AgainstManagement *Number `xml:"AgnstMgmt,omitempty"`

	// Number of votes for which decision is left to the party that will exercise the voting right.
	Discretionary *Number `xml:"Dscrtnry,omitempty"`

	// Number of votes in favour for one year for "say on pay" type of resolution.
	OneYear *Number `xml:"OneYr,omitempty"`

	// Number of votes in favour of two years for "say on pay" type of resolution.
	TwoYears *Number `xml:"TwoYrs,omitempty"`

	// Number of votes in favour of three years for "say on pay" type of resolution.
	ThreeYears *Number `xml:"ThreeYrs,omitempty"`

	// Number of votes for which no action has been taken.
	NoAction *Number `xml:"NoActn,omitempty"`
}

func (v *Vote8) SetIssuerLabel(value string) {
	v.IssuerLabel = (*Max35Text)(&value)
}

func (v *Vote8) SetFor(value string) {
	v.For = (*Number)(&value)
}

func (v *Vote8) SetAgainst(value string) {
	v.Against = (*Number)(&value)
}

func (v *Vote8) SetAbstain(value string) {
	v.Abstain = (*Number)(&value)
}

func (v *Vote8) SetWithhold(value string) {
	v.Withhold = (*Number)(&value)
}

func (v *Vote8) SetWithManagement(value string) {
	v.WithManagement = (*Number)(&value)
}

func (v *Vote8) SetAgainstManagement(value string) {
	v.AgainstManagement = (*Number)(&value)
}

func (v *Vote8) SetDiscretionary(value string) {
	v.Discretionary = (*Number)(&value)
}

func (v *Vote8) SetOneYear(value string) {
	v.OneYear = (*Number)(&value)
}

func (v *Vote8) SetTwoYears(value string) {
	v.TwoYears = (*Number)(&value)
}

func (v *Vote8) SetThreeYears(value string) {
	v.ThreeYears = (*Number)(&value)
}

func (v *Vote8) SetNoAction(value string) {
	v.NoAction = (*Number)(&value)
}

// Decision of the voting party for one resolution.
type Vote9 struct {

	// Numbering of the resolution as specified by the issuer or  its agent.
	IssuerLabel *Max35Text `xml:"IssrLabl"`

	// Specifies the different instructions that can be used to vote.
	VoteOption *VoteInstruction3Code `xml:"VoteOptn"`
}

func (v *Vote9) SetIssuerLabel(value string) {
	v.IssuerLabel = (*Max35Text)(&value)
}

func (v *Vote9) SetVoteOption(value string) {
	v.VoteOption = (*VoteInstruction3Code)(&value)
}
