package iso20022

// Provides details about the collateral against variation margin and optionally the segregated independent amount.
type Collateral1 struct {

	// Provides details about the collateral held, in transit or that still needs to be agreed by both parties, against the variation margin.
	VariationMargin *MarginCollateral1 `xml:"VartnMrgn"`

	// Provides details about the collateral held, in transit or that still needs to be agreed by both parties, against the segregated independent amount.
	SegregatedIndependentAmount *MarginCollateral1 `xml:"SgrtdIndpdntAmt,omitempty"`
}

func (c *Collateral1) AddVariationMargin() *MarginCollateral1 {
	c.VariationMargin = new(MarginCollateral1)
	return c.VariationMargin
}

func (c *Collateral1) AddSegregatedIndependentAmount() *MarginCollateral1 {
	c.SegregatedIndependentAmount = new(MarginCollateral1)
	return c.SegregatedIndependentAmount
}

// Provides details on the collateral that will be either delivered, returned or both.
type Collateral11 struct {

	// Specifies the reference to the unambiguous identification of the margin call request.
	MarginCallRequestIdentification *Max35Text `xml:"MrgnCallReqId"`

	// Specifies the reference to the unambiguous identification of the margin call response.
	MarginCallResponseIdentification *Max35Text `xml:"MrgnCallRspnId,omitempty"`

	// Specifies the standard settlement instructions.
	StandardSettlementInstructions *Max140Text `xml:"StdSttlmInstrs,omitempty"`

	// Specifies the reference to the unambiguous identification of the collateral proposal response (in case of counter proposal).
	CollateralProposalResponseIdentification *Max35Text `xml:"CollPrpslRspnId,omitempty"`

	// Collateral type is securities.
	SecuritiesCollateral []*SecuritiesCollateral5 `xml:"SctiesColl,omitempty"`

	// Collateral type is cash.
	CashCollateral []*CashCollateral2 `xml:"CshColl,omitempty"`

	// Collateral type is other than securities or cash for example letter of credit.
	OtherCollateral []*OtherCollateral5 `xml:"OthrColl,omitempty"`
}

func (c *Collateral11) SetMarginCallRequestIdentification(value string) {
	c.MarginCallRequestIdentification = (*Max35Text)(&value)
}

func (c *Collateral11) SetMarginCallResponseIdentification(value string) {
	c.MarginCallResponseIdentification = (*Max35Text)(&value)
}

func (c *Collateral11) SetStandardSettlementInstructions(value string) {
	c.StandardSettlementInstructions = (*Max140Text)(&value)
}

func (c *Collateral11) SetCollateralProposalResponseIdentification(value string) {
	c.CollateralProposalResponseIdentification = (*Max35Text)(&value)
}

func (c *Collateral11) AddSecuritiesCollateral() *SecuritiesCollateral5 {
	newValue := new(SecuritiesCollateral5)
	c.SecuritiesCollateral = append(c.SecuritiesCollateral, newValue)
	return newValue
}

func (c *Collateral11) AddCashCollateral() *CashCollateral2 {
	newValue := new(CashCollateral2)
	c.CashCollateral = append(c.CashCollateral, newValue)
	return newValue
}

func (c *Collateral11) AddOtherCollateral() *OtherCollateral5 {
	newValue := new(OtherCollateral5)
	c.OtherCollateral = append(c.OtherCollateral, newValue)
	return newValue
}

// Provides details on the collateral that will be either delivered, returned or both.
type Collateral12 struct {

	// Specifies the reference to the unambiguous identification of the margin call request.
	MarginCallRequestIdentification *Max35Text `xml:"MrgnCallReqId"`

	// Specifies the reference to the unambiguous identification of the margin call response.
	MarginCallResponseIdentification *Max35Text `xml:"MrgnCallRspnId,omitempty"`

	// Specifies the standard settlement instructions.
	StandardSettlementInstructions *Max140Text `xml:"StdSttlmInstrs,omitempty"`

	// Specifies the reference to the unambiguous identification of the collateral proposal response (in case of counter proposal).
	CollateralProposalResponseIdentification *Max35Text `xml:"CollPrpslRspnId,omitempty"`

	// Collateral type is securities.
	SecuritiesCollateral []*SecuritiesCollateral5 `xml:"SctiesColl,omitempty"`

	// Collateral type is cash.
	CashCollateral []*CashCollateral3 `xml:"CshColl,omitempty"`

	// Collateral type is other than securities or cash for example letter of credit.
	OtherCollateral []*OtherCollateral5 `xml:"OthrColl,omitempty"`
}

func (c *Collateral12) SetMarginCallRequestIdentification(value string) {
	c.MarginCallRequestIdentification = (*Max35Text)(&value)
}

func (c *Collateral12) SetMarginCallResponseIdentification(value string) {
	c.MarginCallResponseIdentification = (*Max35Text)(&value)
}

func (c *Collateral12) SetStandardSettlementInstructions(value string) {
	c.StandardSettlementInstructions = (*Max140Text)(&value)
}

func (c *Collateral12) SetCollateralProposalResponseIdentification(value string) {
	c.CollateralProposalResponseIdentification = (*Max35Text)(&value)
}

func (c *Collateral12) AddSecuritiesCollateral() *SecuritiesCollateral5 {
	newValue := new(SecuritiesCollateral5)
	c.SecuritiesCollateral = append(c.SecuritiesCollateral, newValue)
	return newValue
}

func (c *Collateral12) AddCashCollateral() *CashCollateral3 {
	newValue := new(CashCollateral3)
	c.CashCollateral = append(c.CashCollateral, newValue)
	return newValue
}

func (c *Collateral12) AddOtherCollateral() *OtherCollateral5 {
	newValue := new(OtherCollateral5)
	c.OtherCollateral = append(c.OtherCollateral, newValue)
	return newValue
}

// Provides for each collateral account the report summary and the valuation of each piece of collateral.
type Collateral13 struct {

	// Provides information about the collateral account, that is the identification, the type and optionally the name.
	AccountIdentification *CollateralAccount2 `xml:"AcctId"`

	// Provides the summary of the collateral valuation report.
	ReportSummary *Summary1 `xml:"RptSummry"`

	// Provides additionnal information about the collateral valuation that has been posted.
	CollateralValuation []*CollateralValuation5 `xml:"CollValtn,omitempty"`
}

func (c *Collateral13) AddAccountIdentification() *CollateralAccount2 {
	c.AccountIdentification = new(CollateralAccount2)
	return c.AccountIdentification
}

func (c *Collateral13) AddReportSummary() *Summary1 {
	c.ReportSummary = new(Summary1)
	return c.ReportSummary
}

func (c *Collateral13) AddCollateralValuation() *CollateralValuation5 {
	newValue := new(CollateralValuation5)
	c.CollateralValuation = append(c.CollateralValuation, newValue)
	return newValue
}

// Provides the details of the security pledge as collateral.
type Collateral14 struct {

	// Provides the values of the security pledged as collateral.
	Valuation *SecuredCollateral2Choice `xml:"Valtn"`

	// Risk control measure applied to underlying collateral whereby the value of that underlying collateral is calculated as the market value of the assets reduced by a certain percentage.
	//
	// For reporting purposes the collateral haircut will be calculated as 100 minus the ratio between the cash lent/borrowed and the market value including accrued interest of the collateral pledged times 100.
	//
	// In the case of multi-collateral repos the haircut will be based on the ratio between the cash borrowed/lent and the market value, including accrued interest of each of the individual collateral pledged.
	//
	// Only actual values, as opposed to estimated or default values will be reported for this variable.
	//
	Haircut *PercentageRate `xml:"Hrcut,omitempty"`

	// Identifies all repurchase agreements conducted against general collateral and those conducted against special collateral.
	// - General Collateral is a repurchase transaction in which the security lender may choose the security to pledge as collateral with the cash provider amongst a relatively wide range of securities meeting predefined criteria;
	// - Special Collateral is a repurchase transaction in which the cash provider requests a specific security (individual ISIN) to be provided by the cash borrower.
	//
	// Usage:
	// This field is optional and it should be provided only in case it is feasible for the reporting agent.
	SpecialCollateralIndicator *SpecialCollateral1Code `xml:"SpclCollInd,omitempty"`
}

func (c *Collateral14) AddValuation() *SecuredCollateral2Choice {
	c.Valuation = new(SecuredCollateral2Choice)
	return c.Valuation
}

func (c *Collateral14) SetHaircut(value string) {
	c.Haircut = (*PercentageRate)(&value)
}

func (c *Collateral14) SetSpecialCollateralIndicator(value string) {
	c.SpecialCollateralIndicator = (*SpecialCollateral1Code)(&value)
}

// Provides details on the collateral that will be either delivered, returned or both.
type Collateral16 struct {

	// Specifies the reference to the unambiguous identification of the margin call request.
	MarginCallRequestIdentification *Max35Text `xml:"MrgnCallReqId"`

	// Specifies the reference to the unambiguous identification of the margin call response.
	MarginCallResponseIdentification *Max35Text `xml:"MrgnCallRspnId,omitempty"`

	// Specifies the standard settlement instructions.
	StandardSettlementInstructions *Max140Text `xml:"StdSttlmInstrs,omitempty"`

	// Specifies the reference to the unambiguous identification of the collateral proposal response (in case of counter proposal).
	CollateralProposalResponseIdentification *Max35Text `xml:"CollPrpslRspnId,omitempty"`

	// Collateral type is securities.
	SecuritiesCollateral []*SecuritiesCollateral8 `xml:"SctiesColl,omitempty"`

	// Collateral type is cash.
	CashCollateral []*CashCollateral3 `xml:"CshColl,omitempty"`

	// Collateral type is other than securities or cash for example letter of credit.
	OtherCollateral []*OtherCollateral5 `xml:"OthrColl,omitempty"`
}

func (c *Collateral16) SetMarginCallRequestIdentification(value string) {
	c.MarginCallRequestIdentification = (*Max35Text)(&value)
}

func (c *Collateral16) SetMarginCallResponseIdentification(value string) {
	c.MarginCallResponseIdentification = (*Max35Text)(&value)
}

func (c *Collateral16) SetStandardSettlementInstructions(value string) {
	c.StandardSettlementInstructions = (*Max140Text)(&value)
}

func (c *Collateral16) SetCollateralProposalResponseIdentification(value string) {
	c.CollateralProposalResponseIdentification = (*Max35Text)(&value)
}

func (c *Collateral16) AddSecuritiesCollateral() *SecuritiesCollateral8 {
	newValue := new(SecuritiesCollateral8)
	c.SecuritiesCollateral = append(c.SecuritiesCollateral, newValue)
	return newValue
}

func (c *Collateral16) AddCashCollateral() *CashCollateral3 {
	newValue := new(CashCollateral3)
	c.CashCollateral = append(c.CashCollateral, newValue)
	return newValue
}

func (c *Collateral16) AddOtherCollateral() *OtherCollateral5 {
	newValue := new(OtherCollateral5)
	c.OtherCollateral = append(c.OtherCollateral, newValue)
	return newValue
}

// Provides details on the collateral that will be either delivered, returned or both.
type Collateral17 struct {

	// Specifies the reference to the unambiguous identification of the margin call request.
	MarginCallRequestIdentification *Max35Text `xml:"MrgnCallReqId"`

	// Specifies the reference to the unambiguous identification of the margin call response.
	MarginCallResponseIdentification *Max35Text `xml:"MrgnCallRspnId,omitempty"`

	// Specifies the standard settlement instructions.
	StandardSettlementInstructions *Max140Text `xml:"StdSttlmInstrs,omitempty"`

	// Specifies the reference to the unambiguous identification of the collateral proposal response (in case of counter proposal).
	CollateralProposalResponseIdentification *Max35Text `xml:"CollPrpslRspnId,omitempty"`

	// Collateral type is securities.
	SecuritiesCollateral []*SecuritiesCollateral8 `xml:"SctiesColl,omitempty"`

	// Collateral type is cash.
	CashCollateral []*CashCollateral2 `xml:"CshColl,omitempty"`

	// Collateral type is other than securities or cash for example letter of credit.
	OtherCollateral []*OtherCollateral5 `xml:"OthrColl,omitempty"`
}

func (c *Collateral17) SetMarginCallRequestIdentification(value string) {
	c.MarginCallRequestIdentification = (*Max35Text)(&value)
}

func (c *Collateral17) SetMarginCallResponseIdentification(value string) {
	c.MarginCallResponseIdentification = (*Max35Text)(&value)
}

func (c *Collateral17) SetStandardSettlementInstructions(value string) {
	c.StandardSettlementInstructions = (*Max140Text)(&value)
}

func (c *Collateral17) SetCollateralProposalResponseIdentification(value string) {
	c.CollateralProposalResponseIdentification = (*Max35Text)(&value)
}

func (c *Collateral17) AddSecuritiesCollateral() *SecuritiesCollateral8 {
	newValue := new(SecuritiesCollateral8)
	c.SecuritiesCollateral = append(c.SecuritiesCollateral, newValue)
	return newValue
}

func (c *Collateral17) AddCashCollateral() *CashCollateral2 {
	newValue := new(CashCollateral2)
	c.CashCollateral = append(c.CashCollateral, newValue)
	return newValue
}

func (c *Collateral17) AddOtherCollateral() *OtherCollateral5 {
	newValue := new(OtherCollateral5)
	c.OtherCollateral = append(c.OtherCollateral, newValue)
	return newValue
}

// Provides the current and market value of the collateral held.
type Collateral3 struct {

	// Value of the collateral after deduction of a percentage (the haircut) that reflects the perceived risk associated with holding this collateral.
	PostHaircutValue *ActiveCurrencyAndAmount `xml:"PstHrcutVal"`

	// Value of the underlying collateral (cash, securities, LoC) based on current market prices.
	MarketValue *ActiveCurrencyAndAmount `xml:"MktVal"`

	// Provides the type of collateral, such as securities or cash.
	CollateralType *CollateralType2Code `xml:"CollTp"`
}

func (c *Collateral3) SetPostHaircutValue(value, currency string) {
	c.PostHaircutValue = NewActiveCurrencyAndAmount(value, currency)
}

func (c *Collateral3) SetMarketValue(value, currency string) {
	c.MarketValue = NewActiveCurrencyAndAmount(value, currency)
}

func (c *Collateral3) SetCollateralType(value string) {
	c.CollateralType = (*CollateralType2Code)(&value)
}

// Provides the current and market value of the collateral held.
type Collateral6 struct {

	// Value of the collateral after deduction of a percentage (the haircut) that reflects the perceived risk associated with holding this collateral.
	PostHaircutValue *ActiveCurrencyAndAmount `xml:"PstHrcutVal"`

	// Value of the underlying collateral (cash, securities, Letter of credit..) based on current market prices.
	MarketValue *ActiveCurrencyAndAmount `xml:"MktVal"`

	// Provides the type of collateral, such as securities or cash.
	CollateralType *CollateralType1Code `xml:"CollTp"`
}

func (c *Collateral6) SetPostHaircutValue(value, currency string) {
	c.PostHaircutValue = NewActiveCurrencyAndAmount(value, currency)
}

func (c *Collateral6) SetMarketValue(value, currency string) {
	c.MarketValue = NewActiveCurrencyAndAmount(value, currency)
}

func (c *Collateral6) SetCollateralType(value string) {
	c.CollateralType = (*CollateralType1Code)(&value)
}

// Provides details on the collateral that will be either delivered, returned or both.
type Collateral7 struct {

	// Specifies the reference to the unambiguous identification of the margin call request.
	MarginCallRequestIdentification *Max35Text `xml:"MrgnCallReqId"`

	// Specifies the reference to the unambiguous identification of the margin call response.
	MarginCallResponseIdentification *Max35Text `xml:"MrgnCallRspnId,omitempty"`

	// Specifies the standard settlement instructions.
	StandardSettlementInstructions *Max140Text `xml:"StdSttlmInstrs,omitempty"`

	// Specifies the reference to the unambiguous identification of the collateral proposal response (in case of counter proposal).
	CollateralProposalResponseIdentification *Max35Text `xml:"CollPrpslRspnId,omitempty"`

	// Collateral type is securities.
	SecuritiesCollateral []*SecuritiesCollateral3 `xml:"SctiesColl,omitempty"`

	// Collateral type is cash.
	CashCollateral []*CashCollateral2 `xml:"CshColl,omitempty"`

	// Collateral type is other than securities or cash for example letter of credit.
	OtherCollateral []*OtherCollateral2 `xml:"OthrColl,omitempty"`
}

func (c *Collateral7) SetMarginCallRequestIdentification(value string) {
	c.MarginCallRequestIdentification = (*Max35Text)(&value)
}

func (c *Collateral7) SetMarginCallResponseIdentification(value string) {
	c.MarginCallResponseIdentification = (*Max35Text)(&value)
}

func (c *Collateral7) SetStandardSettlementInstructions(value string) {
	c.StandardSettlementInstructions = (*Max140Text)(&value)
}

func (c *Collateral7) SetCollateralProposalResponseIdentification(value string) {
	c.CollateralProposalResponseIdentification = (*Max35Text)(&value)
}

func (c *Collateral7) AddSecuritiesCollateral() *SecuritiesCollateral3 {
	newValue := new(SecuritiesCollateral3)
	c.SecuritiesCollateral = append(c.SecuritiesCollateral, newValue)
	return newValue
}

func (c *Collateral7) AddCashCollateral() *CashCollateral2 {
	newValue := new(CashCollateral2)
	c.CashCollateral = append(c.CashCollateral, newValue)
	return newValue
}

func (c *Collateral7) AddOtherCollateral() *OtherCollateral2 {
	newValue := new(OtherCollateral2)
	c.OtherCollateral = append(c.OtherCollateral, newValue)
	return newValue
}

// Provides details on the collateral that will be either delivered, returned or both.
type Collateral8 struct {

	// Specifies the reference to the unambiguous identification of the margin call request.
	MarginCallRequestIdentification *Max35Text `xml:"MrgnCallReqId"`

	// Specifies the reference to the unambiguous identification of the margin call response.
	MarginCallResponseIdentification *Max35Text `xml:"MrgnCallRspnId,omitempty"`

	// Specifies the standard settlement instructions.
	StandardSettlementInstructions *Max140Text `xml:"StdSttlmInstrs,omitempty"`

	// Specifies the reference to the unambiguous identification of the collateral proposal response (in case of counter proposal).
	CollateralProposalResponseIdentification *Max35Text `xml:"CollPrpslRspnId,omitempty"`

	// Collateral type is securities.
	SecuritiesCollateral []*SecuritiesCollateral3 `xml:"SctiesColl,omitempty"`

	// Collateral type is cash.
	CashCollateral []*CashCollateral3 `xml:"CshColl,omitempty"`

	// Collateral type is other than securities or cash for example letter of credit.
	OtherCollateral []*OtherCollateral2 `xml:"OthrColl,omitempty"`
}

func (c *Collateral8) SetMarginCallRequestIdentification(value string) {
	c.MarginCallRequestIdentification = (*Max35Text)(&value)
}

func (c *Collateral8) SetMarginCallResponseIdentification(value string) {
	c.MarginCallResponseIdentification = (*Max35Text)(&value)
}

func (c *Collateral8) SetStandardSettlementInstructions(value string) {
	c.StandardSettlementInstructions = (*Max140Text)(&value)
}

func (c *Collateral8) SetCollateralProposalResponseIdentification(value string) {
	c.CollateralProposalResponseIdentification = (*Max35Text)(&value)
}

func (c *Collateral8) AddSecuritiesCollateral() *SecuritiesCollateral3 {
	newValue := new(SecuritiesCollateral3)
	c.SecuritiesCollateral = append(c.SecuritiesCollateral, newValue)
	return newValue
}

func (c *Collateral8) AddCashCollateral() *CashCollateral3 {
	newValue := new(CashCollateral3)
	c.CashCollateral = append(c.CashCollateral, newValue)
	return newValue
}

func (c *Collateral8) AddOtherCollateral() *OtherCollateral2 {
	newValue := new(OtherCollateral2)
	c.OtherCollateral = append(c.OtherCollateral, newValue)
	return newValue
}

// Provides for each collateral account the report summary and the valuation of each piece of collateral.
type Collateral9 struct {

	// Provides information about the collateral account, that is the identification, the type and optionally the name.
	AccountIdentification *CollateralAccount1 `xml:"AcctId"`

	// Provides the summary of the collateral valuation report.
	ReportSummary *Summary1 `xml:"RptSummry"`

	// Provides additionnal information about the collateral valuation that has been posted.
	CollateralValuation []*CollateralValuation2 `xml:"CollValtn,omitempty"`
}

func (c *Collateral9) AddAccountIdentification() *CollateralAccount1 {
	c.AccountIdentification = new(CollateralAccount1)
	return c.AccountIdentification
}

func (c *Collateral9) AddReportSummary() *Summary1 {
	c.ReportSummary = new(Summary1)
	return c.ReportSummary
}

func (c *Collateral9) AddCollateralValuation() *CollateralValuation2 {
	newValue := new(CollateralValuation2)
	c.CollateralValuation = append(c.CollateralValuation, newValue)
	return newValue
}
