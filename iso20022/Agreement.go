package iso20022

// Agreement details for the over the counter market.
type Agreement2 struct {

	// Full details of the supporting legal agreement under which the margin call can be issued and/or governed.
	AgreementDetails *Max140Text `xml:"AgrmtDtls"`

	// Common reference to the agreement between the two counterparties.
	AgreementIdentification *Max140Text `xml:"AgrmtId,omitempty"`

	// Date on which the collateral agreement was signed.
	AgreementDate *ISODate `xml:"AgrmtDt"`

	// Denomination currency as specified in the collateral agreement.
	BaseCurrency *CurrencyCode `xml:"BaseCcy"`

	// Specifies the underlying master agreement.
	AgreementFramework *AgreementFramework1Choice `xml:"AgrmtFrmwk,omitempty"`
}

func (a *Agreement2) SetAgreementDetails(value string) {
	a.AgreementDetails = (*Max140Text)(&value)
}

func (a *Agreement2) SetAgreementIdentification(value string) {
	a.AgreementIdentification = (*Max140Text)(&value)
}

func (a *Agreement2) SetAgreementDate(value string) {
	a.AgreementDate = (*ISODate)(&value)
}

func (a *Agreement2) SetBaseCurrency(value string) {
	a.BaseCurrency = (*CurrencyCode)(&value)
}

func (a *Agreement2) AddAgreementFramework() *AgreementFramework1Choice {
	a.AgreementFramework = new(AgreementFramework1Choice)
	return a.AgreementFramework
}

// Contractual details related to the agreement between parties.
type Agreement3 struct {

	// Full name of the base standard agreement, annexes and amendments in place between the principals and applicable to this deal.
	Description *Max350Text `xml:"Desc,omitempty"`

	// Numeric representation of the day of the month and year.
	Date *ISODateTime `xml:"Dt,omitempty"`

	// Contractual currency forming the basis of a financing agreement and associated transactions. Usually, but not always, the same as the trade currency.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Type of financing closing.
	ClosingType *ClosingType1Code `xml:"ClsgTp,omitempty"`

	// Start date of a financing deal that is the date the buyer pays the seller cash and takes control of the collateral.
	StartDate *ISODateTime `xml:"StartDt,omitempty"`

	// Identifies type of settlement.
	DeliveryType *DeliveryType2Code `xml:"DlvryTp,omitempty"`

	// Fraction of the cash consideration that must be collateralized, expressed as a percent. A margin ratio of 02% indicates that the value of the collateral (after deducting for "haircut") must exceed the cash consideration by 2%.
	MarginRatio *PercentageRate `xml:"MrgnRatio,omitempty"`
}

func (a *Agreement3) SetDescription(value string) {
	a.Description = (*Max350Text)(&value)
}

func (a *Agreement3) SetDate(value string) {
	a.Date = (*ISODateTime)(&value)
}

func (a *Agreement3) SetCurrency(value string) {
	a.Currency = (*CurrencyCode)(&value)
}

func (a *Agreement3) SetClosingType(value string) {
	a.ClosingType = (*ClosingType1Code)(&value)
}

func (a *Agreement3) SetStartDate(value string) {
	a.StartDate = (*ISODateTime)(&value)
}

func (a *Agreement3) SetDeliveryType(value string) {
	a.DeliveryType = (*DeliveryType2Code)(&value)
}

func (a *Agreement3) SetMarginRatio(value string) {
	a.MarginRatio = (*PercentageRate)(&value)
}

// Agreement details for the over the counter market.
type Agreement4 struct {

	// Full details of the supporting legal agreement under which the margin call can be issued and/or governed.
	AgreementDetails *Max140Text `xml:"AgrmtDtls"`

	// Common reference to the agreement between the two counterparties.
	AgreementIdentification *Max140Text `xml:"AgrmtId,omitempty"`

	// Date on which the collateral agreement was signed.
	AgreementDate *ISODate `xml:"AgrmtDt"`

	// Denomination currency as specified in the collateral agreement.
	BaseCurrency *ActiveCurrencyCode `xml:"BaseCcy"`

	// Specifies the underlying master agreement.
	AgreementFramework *AgreementFramework1Choice `xml:"AgrmtFrmwk,omitempty"`
}

func (a *Agreement4) SetAgreementDetails(value string) {
	a.AgreementDetails = (*Max140Text)(&value)
}

func (a *Agreement4) SetAgreementIdentification(value string) {
	a.AgreementIdentification = (*Max140Text)(&value)
}

func (a *Agreement4) SetAgreementDate(value string) {
	a.AgreementDate = (*ISODate)(&value)
}

func (a *Agreement4) SetBaseCurrency(value string) {
	a.BaseCurrency = (*ActiveCurrencyCode)(&value)
}

func (a *Agreement4) AddAgreementFramework() *AgreementFramework1Choice {
	a.AgreementFramework = new(AgreementFramework1Choice)
	return a.AgreementFramework
}
