package iso20022

// Provides details about the accepted collateral substitution.
type CollateralSubstitutionResponse1 struct {

	// Reference to the collateral substitution request identification.
	CollateralSubstitutionRequestIdentification *Max35Text `xml:"CollSbstitnReqId"`

	// Provides the accepted collateral substitution amount.
	AcceptedAmount *ActiveCurrencyAndAmount `xml:"AccptdAmt,omitempty"`
}

func (c *CollateralSubstitutionResponse1) SetCollateralSubstitutionRequestIdentification(value string) {
	c.CollateralSubstitutionRequestIdentification = (*Max35Text)(&value)
}

func (c *CollateralSubstitutionResponse1) SetAcceptedAmount(value, currency string) {
	c.AcceptedAmount = NewActiveCurrencyAndAmount(value, currency)
}

// Provides details about the rejected collateral substitution.
type CollateralSubstitutionResponse2 struct {

	// Reference to the collateral substitution request identification.
	CollateralSubstitutionRequestIdentification *Max35Text `xml:"CollSbstitnReqId"`

	// Specifies the collateral substitution amount that is rejected.
	RejectedAmount *ActiveCurrencyAndAmount `xml:"RjctdAmt"`

	// Specifies the reasons why the collateral substitution is rejected.
	RejectionReason *RejectionReasonV021Code `xml:"RjctnRsn"`

	// Provides additional information about the collateral substitution request rejection.
	RejectionReasonInformation *Max140Text `xml:"RjctnRsnInf,omitempty"`
}

func (c *CollateralSubstitutionResponse2) SetCollateralSubstitutionRequestIdentification(value string) {
	c.CollateralSubstitutionRequestIdentification = (*Max35Text)(&value)
}

func (c *CollateralSubstitutionResponse2) SetRejectedAmount(value, currency string) {
	c.RejectedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CollateralSubstitutionResponse2) SetRejectionReason(value string) {
	c.RejectionReason = (*RejectionReasonV021Code)(&value)
}

func (c *CollateralSubstitutionResponse2) SetRejectionReasonInformation(value string) {
	c.RejectionReasonInformation = (*Max140Text)(&value)
}
