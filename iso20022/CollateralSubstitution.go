package iso20022

// Provides details about the collateral that will be substituted.
type CollateralSubstitution2 struct {

	// Indicates whether the collateral substitution request is new or updated.
	CollateralSubstitutionSequence *CollateralSubstitutionSequence1Code `xml:"CollSbstitnSeq"`

	// Cash value of the requested collateral substitution transfer in the base currency.
	SubstitutionRequirement *ActiveCurrencyAndAmount `xml:"SbstitnRqrmnt"`

	// Specifies if the collateral that is substituted was posted against the variation margin or the independent amount.
	CollateralSubstitutionType *CollateralSubstitutionType1Code `xml:"CollSbstitnTp"`

	// Identifies the standard settlement instructions.
	StandardSettlementInstructions *Max140Text `xml:"StdSttlmInstrs,omitempty"`

	// Collateral type is securities.
	SecuritiesCollateral []*SecuritiesCollateral4 `xml:"SctiesColl,omitempty"`

	// Collateral type is cash.
	CashCollateral []*CashCollateral5 `xml:"CshColl,omitempty"`

	// Collateral type is other than securities or cash for example letter of credit.
	OtherCollateral []*OtherCollateral4 `xml:"OthrColl,omitempty"`

	// Provides details on the identification of previously sent and/or received message(s), in case of updated substitution request.
	LinkedReferences *Reference17 `xml:"LkdRefs,omitempty"`
}

func (c *CollateralSubstitution2) SetCollateralSubstitutionSequence(value string) {
	c.CollateralSubstitutionSequence = (*CollateralSubstitutionSequence1Code)(&value)
}

func (c *CollateralSubstitution2) SetSubstitutionRequirement(value, currency string) {
	c.SubstitutionRequirement = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CollateralSubstitution2) SetCollateralSubstitutionType(value string) {
	c.CollateralSubstitutionType = (*CollateralSubstitutionType1Code)(&value)
}

func (c *CollateralSubstitution2) SetStandardSettlementInstructions(value string) {
	c.StandardSettlementInstructions = (*Max140Text)(&value)
}

func (c *CollateralSubstitution2) AddSecuritiesCollateral() *SecuritiesCollateral4 {
	newValue := new(SecuritiesCollateral4)
	c.SecuritiesCollateral = append(c.SecuritiesCollateral, newValue)
	return newValue
}

func (c *CollateralSubstitution2) AddCashCollateral() *CashCollateral5 {
	newValue := new(CashCollateral5)
	c.CashCollateral = append(c.CashCollateral, newValue)
	return newValue
}

func (c *CollateralSubstitution2) AddOtherCollateral() *OtherCollateral4 {
	newValue := new(OtherCollateral4)
	c.OtherCollateral = append(c.OtherCollateral, newValue)
	return newValue
}

func (c *CollateralSubstitution2) AddLinkedReferences() *Reference17 {
	c.LinkedReferences = new(Reference17)
	return c.LinkedReferences
}

// Provides details about the collateral that will be substituted.
type CollateralSubstitution3 struct {

	// Indicates whether the collateral substitution request is new or updated.
	CollateralSubstitutionSequence *CollateralSubstitutionSequence1Code `xml:"CollSbstitnSeq"`

	// Cash value of the requested collateral substitution transfer in the base currency.
	SubstitutionRequirement *ActiveCurrencyAndAmount `xml:"SbstitnRqrmnt"`

	// Specifies if the collateral that is substituted was posted against the variation margin or the independent amount.
	CollateralSubstitutionType *CollateralSubstitutionType1Code `xml:"CollSbstitnTp"`

	// Identifies the standard settlement instructions.
	StandardSettlementInstructions *Max140Text `xml:"StdSttlmInstrs,omitempty"`

	// Collateral type is securities.
	SecuritiesCollateral []*SecuritiesCollateral3 `xml:"SctiesColl,omitempty"`

	// Collateral type is cash.
	CashCollateral []*CashCollateral3 `xml:"CshColl,omitempty"`

	// Collateral type is other than securities or cash for example letter of credit.
	OtherCollateral []*OtherCollateral2 `xml:"OthrColl,omitempty"`

	// Provides details on the identification of previously sent and/or received message(s), in case of updated substitution request.
	LinkedReferences *Reference17 `xml:"LkdRefs,omitempty"`
}

func (c *CollateralSubstitution3) SetCollateralSubstitutionSequence(value string) {
	c.CollateralSubstitutionSequence = (*CollateralSubstitutionSequence1Code)(&value)
}

func (c *CollateralSubstitution3) SetSubstitutionRequirement(value, currency string) {
	c.SubstitutionRequirement = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CollateralSubstitution3) SetCollateralSubstitutionType(value string) {
	c.CollateralSubstitutionType = (*CollateralSubstitutionType1Code)(&value)
}

func (c *CollateralSubstitution3) SetStandardSettlementInstructions(value string) {
	c.StandardSettlementInstructions = (*Max140Text)(&value)
}

func (c *CollateralSubstitution3) AddSecuritiesCollateral() *SecuritiesCollateral3 {
	newValue := new(SecuritiesCollateral3)
	c.SecuritiesCollateral = append(c.SecuritiesCollateral, newValue)
	return newValue
}

func (c *CollateralSubstitution3) AddCashCollateral() *CashCollateral3 {
	newValue := new(CashCollateral3)
	c.CashCollateral = append(c.CashCollateral, newValue)
	return newValue
}

func (c *CollateralSubstitution3) AddOtherCollateral() *OtherCollateral2 {
	newValue := new(OtherCollateral2)
	c.OtherCollateral = append(c.OtherCollateral, newValue)
	return newValue
}

func (c *CollateralSubstitution3) AddLinkedReferences() *Reference17 {
	c.LinkedReferences = new(Reference17)
	return c.LinkedReferences
}

// Provides details about the collateral that will be substituted.
type CollateralSubstitution4 struct {

	// Indicates whether the collateral substitution request is new or updated.
	CollateralSubstitutionSequence *CollateralSubstitutionSequence1Code `xml:"CollSbstitnSeq"`

	// Cash value of the requested collateral substitution transfer in the base currency.
	SubstitutionRequirement *ActiveCurrencyAndAmount `xml:"SbstitnRqrmnt"`

	// Specifies if the collateral that is substituted was posted against the variation margin or the independent amount.
	CollateralSubstitutionType *CollateralSubstitutionType1Code `xml:"CollSbstitnTp"`

	// Identifies the standard settlement instructions.
	StandardSettlementInstructions *Max140Text `xml:"StdSttlmInstrs,omitempty"`

	// Collateral type is securities.
	SecuritiesCollateral []*SecuritiesCollateral5 `xml:"SctiesColl,omitempty"`

	// Collateral type is cash.
	CashCollateral []*CashCollateral3 `xml:"CshColl,omitempty"`

	// Collateral type is other than securities or cash for example letter of credit.
	OtherCollateral []*OtherCollateral5 `xml:"OthrColl,omitempty"`

	// Provides details on the identification of previously sent and/or received message(s), in case of updated substitution request.
	LinkedReferences *Reference17 `xml:"LkdRefs,omitempty"`
}

func (c *CollateralSubstitution4) SetCollateralSubstitutionSequence(value string) {
	c.CollateralSubstitutionSequence = (*CollateralSubstitutionSequence1Code)(&value)
}

func (c *CollateralSubstitution4) SetSubstitutionRequirement(value, currency string) {
	c.SubstitutionRequirement = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CollateralSubstitution4) SetCollateralSubstitutionType(value string) {
	c.CollateralSubstitutionType = (*CollateralSubstitutionType1Code)(&value)
}

func (c *CollateralSubstitution4) SetStandardSettlementInstructions(value string) {
	c.StandardSettlementInstructions = (*Max140Text)(&value)
}

func (c *CollateralSubstitution4) AddSecuritiesCollateral() *SecuritiesCollateral5 {
	newValue := new(SecuritiesCollateral5)
	c.SecuritiesCollateral = append(c.SecuritiesCollateral, newValue)
	return newValue
}

func (c *CollateralSubstitution4) AddCashCollateral() *CashCollateral3 {
	newValue := new(CashCollateral3)
	c.CashCollateral = append(c.CashCollateral, newValue)
	return newValue
}

func (c *CollateralSubstitution4) AddOtherCollateral() *OtherCollateral5 {
	newValue := new(OtherCollateral5)
	c.OtherCollateral = append(c.OtherCollateral, newValue)
	return newValue
}

func (c *CollateralSubstitution4) AddLinkedReferences() *Reference17 {
	c.LinkedReferences = new(Reference17)
	return c.LinkedReferences
}

// Provides details about the collateral that will be substituted.
type CollateralSubstitution5 struct {

	// Indicates whether the collateral substitution request is new or updated.
	CollateralSubstitutionSequence *CollateralSubstitutionSequence1Code `xml:"CollSbstitnSeq"`

	// Cash value of the requested collateral substitution transfer in the base currency.
	SubstitutionRequirement *ActiveCurrencyAndAmount `xml:"SbstitnRqrmnt"`

	// Specifies if the collateral that is substituted was posted against the variation margin or the independent amount.
	CollateralSubstitutionType *CollateralSubstitutionType1Code `xml:"CollSbstitnTp"`

	// Identifies the standard settlement instructions.
	StandardSettlementInstructions *Max140Text `xml:"StdSttlmInstrs,omitempty"`

	// Collateral type is securities.
	SecuritiesCollateral []*SecuritiesCollateral7 `xml:"SctiesColl,omitempty"`

	// Collateral type is cash.
	CashCollateral []*CashCollateral5 `xml:"CshColl,omitempty"`

	// Collateral type is other than securities or cash for example letter of credit.
	OtherCollateral []*OtherCollateral7 `xml:"OthrColl,omitempty"`

	// Provides details on the identification of previously sent and/or received message(s), in case of updated substitution request.
	LinkedReferences *Reference17 `xml:"LkdRefs,omitempty"`
}

func (c *CollateralSubstitution5) SetCollateralSubstitutionSequence(value string) {
	c.CollateralSubstitutionSequence = (*CollateralSubstitutionSequence1Code)(&value)
}

func (c *CollateralSubstitution5) SetSubstitutionRequirement(value, currency string) {
	c.SubstitutionRequirement = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CollateralSubstitution5) SetCollateralSubstitutionType(value string) {
	c.CollateralSubstitutionType = (*CollateralSubstitutionType1Code)(&value)
}

func (c *CollateralSubstitution5) SetStandardSettlementInstructions(value string) {
	c.StandardSettlementInstructions = (*Max140Text)(&value)
}

func (c *CollateralSubstitution5) AddSecuritiesCollateral() *SecuritiesCollateral7 {
	newValue := new(SecuritiesCollateral7)
	c.SecuritiesCollateral = append(c.SecuritiesCollateral, newValue)
	return newValue
}

func (c *CollateralSubstitution5) AddCashCollateral() *CashCollateral5 {
	newValue := new(CashCollateral5)
	c.CashCollateral = append(c.CashCollateral, newValue)
	return newValue
}

func (c *CollateralSubstitution5) AddOtherCollateral() *OtherCollateral7 {
	newValue := new(OtherCollateral7)
	c.OtherCollateral = append(c.OtherCollateral, newValue)
	return newValue
}

func (c *CollateralSubstitution5) AddLinkedReferences() *Reference17 {
	c.LinkedReferences = new(Reference17)
	return c.LinkedReferences
}
