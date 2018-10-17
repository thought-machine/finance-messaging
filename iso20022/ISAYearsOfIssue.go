package iso20022

// Year in which the ISA plan is issued.
type ISAYearsOfIssue1 struct {

	// Current year ISA is an ISA that was issued during the current fiscal year.
	CurrentYearType *ISAType1Code `xml:"CurYrTp,omitempty"`

	// Current year ISA is an ISA that was issued during the current fiscal year.
	ExtendedCurrentYearType *Extended350Code `xml:"XtndedCurYrTp,omitempty"`

	// Indicates whether the ISA contains a cash component asset for transfer.
	CashComponentIndicator *YesNoIndicator `xml:"CshCmpntInd"`

	// Selection of investment plans issued during previous years.
	PreviousYears *PreviousYear1 `xml:"PrvsYrs,omitempty"`

	// Specifies the amounts already subscribed for the current year.
	CurrentYearSubscriptionDetails *SubscriptionInformation1 `xml:"CurYrSbcptDtls,omitempty"`
}

func (i *ISAYearsOfIssue1) SetCurrentYearType(value string) {
	i.CurrentYearType = (*ISAType1Code)(&value)
}

func (i *ISAYearsOfIssue1) SetExtendedCurrentYearType(value string) {
	i.ExtendedCurrentYearType = (*Extended350Code)(&value)
}

func (i *ISAYearsOfIssue1) SetCashComponentIndicator(value string) {
	i.CashComponentIndicator = (*YesNoIndicator)(&value)
}

func (i *ISAYearsOfIssue1) AddPreviousYears() *PreviousYear1 {
	i.PreviousYears = new(PreviousYear1)
	return i.PreviousYears
}

func (i *ISAYearsOfIssue1) AddCurrentYearSubscriptionDetails() *SubscriptionInformation1 {
	i.CurrentYearSubscriptionDetails = new(SubscriptionInformation1)
	return i.CurrentYearSubscriptionDetails
}

// Year in which the ISA plan is issued.
type ISAYearsOfIssue2 struct {

	// ISA that was issued during the current fiscal year.
	CurrentYearType *ISAType2Code `xml:"CurYrTp,omitempty"`

	// Current year ISA is an ISA that was issued during the current fiscal year.
	ExtendedCurrentYearType *Extended350Code `xml:"XtndedCurYrTp,omitempty"`

	// Selection of investment plans issued during previous years.
	PreviousYears *PreviousYearChoice `xml:"PrvsYrs,omitempty"`
}

func (i *ISAYearsOfIssue2) SetCurrentYearType(value string) {
	i.CurrentYearType = (*ISAType2Code)(&value)
}

func (i *ISAYearsOfIssue2) SetExtendedCurrentYearType(value string) {
	i.ExtendedCurrentYearType = (*Extended350Code)(&value)
}

func (i *ISAYearsOfIssue2) AddPreviousYears() *PreviousYearChoice {
	i.PreviousYears = new(PreviousYearChoice)
	return i.PreviousYears
}

// Year in which the ISA plan is issued.
type ISAYearsOfIssue3 struct {

	// ISA that was issued during the current fiscal year.
	CurrentYearType *ISAType1Code `xml:"CurYrTp,omitempty"`

	// Current year ISA is an ISA that was issued during the current fiscal year.
	ExtendedCurrentYearType *Extended350Code `xml:"XtndedCurYrTp,omitempty"`

	// Indicates whether the ISA contains a cash component asset for transfer.
	CashComponentIndicator *YesNoIndicator `xml:"CshCmpntInd"`

	// Specifies the amounts already subscribed for the current year.
	CurrentYearSubscriptionDetails *SubscriptionInformation1 `xml:"CurYrSbcptDtls"`

	// Selection of investment plans issued during previous years.
	PreviousYears *PreviousYear1 `xml:"PrvsYrs,omitempty"`
}

func (i *ISAYearsOfIssue3) SetCurrentYearType(value string) {
	i.CurrentYearType = (*ISAType1Code)(&value)
}

func (i *ISAYearsOfIssue3) SetExtendedCurrentYearType(value string) {
	i.ExtendedCurrentYearType = (*Extended350Code)(&value)
}

func (i *ISAYearsOfIssue3) SetCashComponentIndicator(value string) {
	i.CashComponentIndicator = (*YesNoIndicator)(&value)
}

func (i *ISAYearsOfIssue3) AddCurrentYearSubscriptionDetails() *SubscriptionInformation1 {
	i.CurrentYearSubscriptionDetails = new(SubscriptionInformation1)
	return i.CurrentYearSubscriptionDetails
}

func (i *ISAYearsOfIssue3) AddPreviousYears() *PreviousYear1 {
	i.PreviousYears = new(PreviousYear1)
	return i.PreviousYears
}

// Year in which the ISA plan is issued.
type ISAYearsOfIssue4 struct {

	// Current year of the Investment Saving Plan (ISA) that was issued during the current fiscal year.
	CurrentYear *CurrentYearType1Choice `xml:"CurYr,omitempty"`

	// Indicates whether the ISA contains a cash component asset for transfer.
	CashComponentIndicator *YesNoIndicator `xml:"CshCmpntInd"`

	// Selection of investment plans issued during previous years.
	PreviousYears *PreviousYear2 `xml:"PrvsYrs,omitempty"`

	// Specifies the amounts already subscribed for the current year.
	CurrentYearSubscriptionDetails *SubscriptionInformation1 `xml:"CurYrSbcptDtls,omitempty"`
}

func (i *ISAYearsOfIssue4) AddCurrentYear() *CurrentYearType1Choice {
	i.CurrentYear = new(CurrentYearType1Choice)
	return i.CurrentYear
}

func (i *ISAYearsOfIssue4) SetCashComponentIndicator(value string) {
	i.CashComponentIndicator = (*YesNoIndicator)(&value)
}

func (i *ISAYearsOfIssue4) AddPreviousYears() *PreviousYear2 {
	i.PreviousYears = new(PreviousYear2)
	return i.PreviousYears
}

func (i *ISAYearsOfIssue4) AddCurrentYearSubscriptionDetails() *SubscriptionInformation1 {
	i.CurrentYearSubscriptionDetails = new(SubscriptionInformation1)
	return i.CurrentYearSubscriptionDetails
}

// Year in which the ISA plan is issued.
type ISAYearsOfIssue5 struct {

	// ISA that was issued during the current fiscal year.
	CurrentYear *CurrentYearType1Choice `xml:"CurYr,omitempty"`

	// Indicates whether the ISA contains a cash component asset for transfer.
	CashComponentIndicator *YesNoIndicator `xml:"CshCmpntInd"`

	// Selection of investment plans issued during previous years.
	PreviousYears *PreviousYear3 `xml:"PrvsYrs,omitempty"`

	// Specifies the amounts already subscribed for the current year.
	CurrentYearSubscriptionDetails *SubscriptionInformation1 `xml:"CurYrSbcptDtls"`
}

func (i *ISAYearsOfIssue5) AddCurrentYear() *CurrentYearType1Choice {
	i.CurrentYear = new(CurrentYearType1Choice)
	return i.CurrentYear
}

func (i *ISAYearsOfIssue5) SetCashComponentIndicator(value string) {
	i.CashComponentIndicator = (*YesNoIndicator)(&value)
}

func (i *ISAYearsOfIssue5) AddPreviousYears() *PreviousYear3 {
	i.PreviousYears = new(PreviousYear3)
	return i.PreviousYears
}

func (i *ISAYearsOfIssue5) AddCurrentYearSubscriptionDetails() *SubscriptionInformation1 {
	i.CurrentYearSubscriptionDetails = new(SubscriptionInformation1)
	return i.CurrentYearSubscriptionDetails
}

// Year in which the ISA plan is issued.
type ISAYearsOfIssue6 struct {

	// ISA that was issued during the current fiscal year.
	CurrentYear *CurrentYearType2Choice `xml:"CurYr,omitempty"`

	// Selection of investment plans issued during previous years.
	PreviousYears *PreviousYearChoice `xml:"PrvsYrs,omitempty"`
}

func (i *ISAYearsOfIssue6) AddCurrentYear() *CurrentYearType2Choice {
	i.CurrentYear = new(CurrentYearType2Choice)
	return i.CurrentYear
}

func (i *ISAYearsOfIssue6) AddPreviousYears() *PreviousYearChoice {
	i.PreviousYears = new(PreviousYearChoice)
	return i.PreviousYears
}
