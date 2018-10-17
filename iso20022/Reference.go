package iso20022

// Provides the identification of the collateral message cancellation request.
type Reference16 struct {

	// Identification of the collateral message cancellation request.
	CollateralMessageCancellationRequestIdentification *Max35Text `xml:"CollMsgCxlReqId"`
}

func (r *Reference16) SetCollateralMessageCancellationRequestIdentification(value string) {
	r.CollateralMessageCancellationRequestIdentification = (*Max35Text)(&value)
}

// Information related to a linked transaction.
type Reference17 struct {

	// Identification of the collateral substitution request.
	CollateralSubstitutionRequestIdentification *Max35Text `xml:"CollSbstitnReqId"`

	// Identification of the collateral substitution response.
	CollateralSubstitutionResponseIdentification *Max35Text `xml:"CollSbstitnRspnId,omitempty"`
}

func (r *Reference17) SetCollateralSubstitutionRequestIdentification(value string) {
	r.CollateralSubstitutionRequestIdentification = (*Max35Text)(&value)
}

func (r *Reference17) SetCollateralSubstitutionResponseIdentification(value string) {
	r.CollateralSubstitutionResponseIdentification = (*Max35Text)(&value)
}

// Provides the references of the underlying trade leg(s) and/or the reference to the related NetPosition report message.
type Reference19 struct {

	// Reference allocated by the central counterparty - central counterpatry trade leg reference identification that uniquely identifies the trade.
	TradeLegNotificationIdentification []*Max35Text `xml:"TradLegNtfctnId,omitempty"`

	// After netting, reference that is common to a net transaction to settle and all its underlying trades
	NetPositionIdentification *Max35Text `xml:"NetPosId,omitempty"`
}

func (r *Reference19) AddTradeLegNotificationIdentification(value string) {
	r.TradeLegNotificationIdentification = append(r.TradeLegNotificationIdentification, (*Max35Text)(&value))
}

func (r *Reference19) SetNetPositionIdentification(value string) {
	r.NetPositionIdentification = (*Max35Text)(&value)
}

// Additional references linked to the updated interest request such the original InterestRequest identification, and optionaly the InterestResponse identification.
type Reference20 struct {

	// Provides the reference to the interest payment request.
	InterestPaymentRequestIdentification *Max35Text `xml:"IntrstPmtReqId"`

	// Provides the reference to the interest payment response.
	InterestPaymentResponseIdentification *Max35Text `xml:"IntrstPmtRspnId,omitempty"`
}

func (r *Reference20) SetInterestPaymentRequestIdentification(value string) {
	r.InterestPaymentRequestIdentification = (*Max35Text)(&value)
}

func (r *Reference20) SetInterestPaymentResponseIdentification(value string) {
	r.InterestPaymentResponseIdentification = (*Max35Text)(&value)
}
