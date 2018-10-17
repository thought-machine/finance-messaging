package iso20022

// Details of the margin call request.
type MarginCall1 struct {

	// Sum of the exposures of all transactions which are in the favour of party A. That is, all transactions which would have an amount payable by party B to party A if they were being terminated.
	ExposedAmountPartyA *ActiveCurrencyAndAmount `xml:"XpsdAmtPtyA,omitempty"`

	// Sum of the exposures of all transactions which are in the favour of party B. That is, all transactions which would have an amount payable by party A to party B if they were being terminated.
	ExposedAmountPartyB *ActiveCurrencyAndAmount `xml:"XpsdAmtPtyB,omitempty"`

	// Determines how the variation margin requirement is to be calculated:
	// - either Net, in which the exposure of all transactions in favour of party A and the the exposure of all transactions in favour of party B will be netted together or
	// - gross in which two separate variation margin requirements will be determined.
	ExposureConvention *ExposureConventionType1Code `xml:"XpsrCnvntn,omitempty"`

	// Amount applied as an add-on to the exposure (to party A) usually intended to cover a possible increase in exposure before the next valuation date.
	IndependentAmountPartyA *AggregatedIndependentAmount1 `xml:"IndpdntAmtPtyA,omitempty"`

	// An amount applied as an add-on to the exposure (to party B) usually intended to cover a possible increase in exposure before the next valuation date.
	IndependentAmountPartyB *AggregatedIndependentAmount1 `xml:"IndpdntAmtPtyB,omitempty"`

	// Provides information like threshold amount, threshold type, minimum transfer amount, rouding amount or rounding convention, that applies to either the variation margin or the segregated independent amount.
	MarginTerms *MarginTerms1Choice `xml:"MrgnTerms,omitempty"`

	// Provides details about the collateral held, in transit or that still needs to be agreed by both parties with a segregation between variation margin and segregated independent amount.
	CollateralBalance *CollateralBalance1Choice `xml:"CollBal,omitempty"`
}

func (m *MarginCall1) SetExposedAmountPartyA(value, currency string) {
	m.ExposedAmountPartyA = NewActiveCurrencyAndAmount(value, currency)
}

func (m *MarginCall1) SetExposedAmountPartyB(value, currency string) {
	m.ExposedAmountPartyB = NewActiveCurrencyAndAmount(value, currency)
}

func (m *MarginCall1) SetExposureConvention(value string) {
	m.ExposureConvention = (*ExposureConventionType1Code)(&value)
}

func (m *MarginCall1) AddIndependentAmountPartyA() *AggregatedIndependentAmount1 {
	m.IndependentAmountPartyA = new(AggregatedIndependentAmount1)
	return m.IndependentAmountPartyA
}

func (m *MarginCall1) AddIndependentAmountPartyB() *AggregatedIndependentAmount1 {
	m.IndependentAmountPartyB = new(AggregatedIndependentAmount1)
	return m.IndependentAmountPartyB
}

func (m *MarginCall1) AddMarginTerms() *MarginTerms1Choice {
	m.MarginTerms = new(MarginTerms1Choice)
	return m.MarginTerms
}

func (m *MarginCall1) AddCollateralBalance() *CollateralBalance1Choice {
	m.CollateralBalance = new(CollateralBalance1Choice)
	return m.CollateralBalance
}

// Specifies the calculation and the resulting margin and independent amount needed to cover the risk exposure of one party versus another.
type MarginCall2 struct {

	// Provides additional information on the collateral account of the party delivering/receiving the collateral.
	CollateralAccountIdentification *CollateralAccount2 `xml:"CollAcctId,omitempty"`

	// Summation of the call amounts per margin type. It is provided for the purposes of carrying forward for future messages that are used to compare the margin call results to determine whether a call is agreed or full/partially disputed.
	MarginCallResult *MarginCallResult3 `xml:"MrgnCallRslt"`

	// Provides details about the margin calculation that would be due to party A.
	MarginDetailDueToA *MarginCall1 `xml:"MrgnDtlDueToA,omitempty"`

	// Provides details about the margin calculation that would be due to party B.
	MarginDetailDueToB *MarginCall1 `xml:"MrgnDtlDueToB,omitempty"`

	// Amount of expected margin that will be either delivered to party A by party B or recalled to party A from party B.
	RequirementDetailsDueToA *MarginRequirement1Choice `xml:"RqrmntDtlsDueToA,omitempty"`

	// Amount of expected margin that will be either delivered to party B by party A or recalled to party B from party A.
	RequirementDetailsDueToB *MarginRequirement1Choice `xml:"RqrmntDtlsDueToB,omitempty"`

	// Provides details about the type of collateral that will be either delivered to party A by party B or recalled to party A from party B.
	ExpectedCollateralDueToA *ExpectedCollateral2Choice `xml:"XpctdCollDueToA,omitempty"`

	// Provides details about the type of collateral that will be either delivered to party B by party A or recalled to party B from party A.
	ExpectedCollateralDueToB *ExpectedCollateral2Choice `xml:"XpctdCollDueToB,omitempty"`
}

func (m *MarginCall2) AddCollateralAccountIdentification() *CollateralAccount2 {
	m.CollateralAccountIdentification = new(CollateralAccount2)
	return m.CollateralAccountIdentification
}

func (m *MarginCall2) AddMarginCallResult() *MarginCallResult3 {
	m.MarginCallResult = new(MarginCallResult3)
	return m.MarginCallResult
}

func (m *MarginCall2) AddMarginDetailDueToA() *MarginCall1 {
	m.MarginDetailDueToA = new(MarginCall1)
	return m.MarginDetailDueToA
}

func (m *MarginCall2) AddMarginDetailDueToB() *MarginCall1 {
	m.MarginDetailDueToB = new(MarginCall1)
	return m.MarginDetailDueToB
}

func (m *MarginCall2) AddRequirementDetailsDueToA() *MarginRequirement1Choice {
	m.RequirementDetailsDueToA = new(MarginRequirement1Choice)
	return m.RequirementDetailsDueToA
}

func (m *MarginCall2) AddRequirementDetailsDueToB() *MarginRequirement1Choice {
	m.RequirementDetailsDueToB = new(MarginRequirement1Choice)
	return m.RequirementDetailsDueToB
}

func (m *MarginCall2) AddExpectedCollateralDueToA() *ExpectedCollateral2Choice {
	m.ExpectedCollateralDueToA = new(ExpectedCollateral2Choice)
	return m.ExpectedCollateralDueToA
}

func (m *MarginCall2) AddExpectedCollateralDueToB() *ExpectedCollateral2Choice {
	m.ExpectedCollateralDueToB = new(ExpectedCollateral2Choice)
	return m.ExpectedCollateralDueToB
}
