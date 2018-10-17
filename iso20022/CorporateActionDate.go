package iso20022

// Specifies corporate action dates.
type CorporateActionDate1 struct {

	// Date/time at which the issuer announced that a corporate action event will occur.
	AnnouncementDate *DateFormat6Choice `xml:"AnncmntDt,omitempty"`

	// Deadline by which the beneficial ownership of securities must be declared.
	CertificationDeadline *DateFormat6Choice `xml:"CertfctnDdln,omitempty"`

	// Date upon which the court provided approval.
	CourtApprovalDate *DateFormat6Choice `xml:"CrtApprvlDt,omitempty"`

	// First possible early closing date of an offer if different from the expiry date.
	EarlyClosingDate *DateFormat6Choice `xml:"EarlyClsgDt,omitempty"`

	// Date/time at which an event is officially effective from the issuer's perspective.
	EffectiveDate *DateFormat6Choice `xml:"FctvDt,omitempty"`

	// Date/Time on which all or part of any holding bought in a unit trust is subject to being treated as capital rather than income. This is normally one day after the previous distribution's ex date.
	EqualisationDate *DateFormat6Choice `xml:"EqulstnDt,omitempty"`

	// Date/time at which additional information on the event will be announced, for example, exchange ratio announcement date.
	FurtherDetailedAnnouncementDate *DateFormat6Choice `xml:"FrthrDtldAnncmntDt,omitempty"`

	// Date/time at which an index rate will be determined .
	IndexFixingDate *DateFormat6Choice `xml:"IndxFxgDt,omitempty"`

	// Date/time on which the lottery is run and applied to the holder's positions. This is also applicable to partial calls.
	LotteryDate *DateFormat6Choice `xml:"LtryDt,omitempty"`

	// Date/time upon on which an interest bearing financial instrument becomes due and principal is repaid by the issuer to the investor.
	MaturityDate *DateFormat6Choice `xml:"MtrtyDt,omitempty"`

	// Date/time on which the bondholder's or shareholder's meeting will take place.
	MeetingDate *DateFormat6Choice `xml:"MtgDt,omitempty"`

	// Date/time at which the margin rate will be determined .
	MarginFixingDate *DateFormat6Choice `xml:"MrgnFxgDt,omitempty"`

	// Date/time (and time) at which an issuer will determine the proration amount/quantity of an offer.
	ProrationDate *DateFormat6Choice `xml:"PrratnDt,omitempty"`

	// Date/time at which positions are struck at the end of the day to note which parties will receive the relevant amount of entitlement, due to be distributed on payment date.
	RecordDate *DateFormat6Choice `xml:"RcrdDt,omitempty"`

	// Date/time on which instructions to register or registration details will be accepted.
	RegistrationDeadline *DateFormat6Choice `xml:"RegnDdln,omitempty"`

	// Date/time on which results are published, for example, results of an offer.
	ResultsPublicationDate *DateFormat6Choice `xml:"RsltsPblctnDt,omitempty"`

	// Deadline by which instructions must be received to split securities, for example, of physical certificates.
	DeadlineToSplit *DateFormat6Choice `xml:"DdlnToSplt,omitempty"`

	// Date/time on until which tax breakdown instructions will be accepted.
	DeadlineForTaxBreakdownInstruction *DateFormat6Choice `xml:"DdlnForTaxBrkdwnInstr,omitempty"`

	// Date/time at which trading of a security is suspended as the result of an event.
	TradingSuspendedDate *DateFormat6Choice `xml:"TradgSspdDt,omitempty"`

	// Date/time upon which the terms of the take-over become unconditional as to acceptances.
	UnconditionalDate *DateFormat6Choice `xml:"UcondlDt,omitempty"`

	// Date/time at on which all conditions, including regulatory, legal etc. pertaining to the take-over, have been met.
	WhollyUnconditionalDate *DateFormat6Choice `xml:"WhlyUcondlDt,omitempty"`

	// Date/time as from which trading (including exchange and OTC trading) occurs on the underlying security without the benefit.
	ExDividendDate *DateFormat6Choice `xml:"ExDvddDt,omitempty"`

	// Date/time at which the corporate action is legally announced by an official body, for example, publication by a governmental administration.
	OfficialAnnouncementPublicationDate *DateFormat6Choice `xml:"OffclAnncmntPblctnDt,omitempty"`

	// Date/time as from which 'special processing' can start to be used by participants for that event. Special processing is a means of marking a transaction, that would normally be traded ex or cum, as being traded cum or ex respectively,  for example, a transaction dealt 'special' after the ex date would result in the buyer being eligible for the entitlement. This is typically used in the UK and Irish markets.
	SpecialExDate *DateFormat6Choice `xml:"SpclExDt,omitempty"`

	// Last date/time by which a buying counterparty to a trade can be sure that it will have the right to participate in an event.
	GuaranteedParticipationDate *DateFormat6Choice `xml:"GrntedPrtcptnDt,omitempty"`

	// Deadline by which an entitled holder needs to advise their counterparty to a transaction of their election for a corporate action event.
	ElectionToCounterpartyDeadline *DateFormat6Choice `xml:"ElctnToCtrPtyDdln,omitempty"`

	// Date/time at which an event/offer is terminated or lapsed.
	LapsedDate *DateFormat6Choice `xml:"LpsdDt,omitempty"`

	// Date/time at which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat6Choice `xml:"PmtDt,omitempty"`

	// Date/Time by which the account owner must instruct directly another party, for example to provide documentation to an issuer agent.
	ThirdPartyDeadline *DateFormat6Choice `xml:"ThrdPtyDdln,omitempty"`

	// Date/Time set by the issuer agent as a first early deadline by which the account owner must instruct directly another party, possibly giving the holder eligibility to incentives. For example, to provide documentation to an issuer agent.
	EarlyThirdPartyDeadline *DateFormat6Choice `xml:"EarlyThrdPtyDdln,omitempty"`

	// Date by which the depository stops monitoring activities of the event, for instance, accounting and tracking activities for due bills end.
	MarketClaimTrackingEndDate *DateFormat6Choice `xml:"MktClmTrckgEndDt,omitempty"`
}

func (c *CorporateActionDate1) AddAnnouncementDate() *DateFormat6Choice {
	c.AnnouncementDate = new(DateFormat6Choice)
	return c.AnnouncementDate
}

func (c *CorporateActionDate1) AddCertificationDeadline() *DateFormat6Choice {
	c.CertificationDeadline = new(DateFormat6Choice)
	return c.CertificationDeadline
}

func (c *CorporateActionDate1) AddCourtApprovalDate() *DateFormat6Choice {
	c.CourtApprovalDate = new(DateFormat6Choice)
	return c.CourtApprovalDate
}

func (c *CorporateActionDate1) AddEarlyClosingDate() *DateFormat6Choice {
	c.EarlyClosingDate = new(DateFormat6Choice)
	return c.EarlyClosingDate
}

func (c *CorporateActionDate1) AddEffectiveDate() *DateFormat6Choice {
	c.EffectiveDate = new(DateFormat6Choice)
	return c.EffectiveDate
}

func (c *CorporateActionDate1) AddEqualisationDate() *DateFormat6Choice {
	c.EqualisationDate = new(DateFormat6Choice)
	return c.EqualisationDate
}

func (c *CorporateActionDate1) AddFurtherDetailedAnnouncementDate() *DateFormat6Choice {
	c.FurtherDetailedAnnouncementDate = new(DateFormat6Choice)
	return c.FurtherDetailedAnnouncementDate
}

func (c *CorporateActionDate1) AddIndexFixingDate() *DateFormat6Choice {
	c.IndexFixingDate = new(DateFormat6Choice)
	return c.IndexFixingDate
}

func (c *CorporateActionDate1) AddLotteryDate() *DateFormat6Choice {
	c.LotteryDate = new(DateFormat6Choice)
	return c.LotteryDate
}

func (c *CorporateActionDate1) AddMaturityDate() *DateFormat6Choice {
	c.MaturityDate = new(DateFormat6Choice)
	return c.MaturityDate
}

func (c *CorporateActionDate1) AddMeetingDate() *DateFormat6Choice {
	c.MeetingDate = new(DateFormat6Choice)
	return c.MeetingDate
}

func (c *CorporateActionDate1) AddMarginFixingDate() *DateFormat6Choice {
	c.MarginFixingDate = new(DateFormat6Choice)
	return c.MarginFixingDate
}

func (c *CorporateActionDate1) AddProrationDate() *DateFormat6Choice {
	c.ProrationDate = new(DateFormat6Choice)
	return c.ProrationDate
}

func (c *CorporateActionDate1) AddRecordDate() *DateFormat6Choice {
	c.RecordDate = new(DateFormat6Choice)
	return c.RecordDate
}

func (c *CorporateActionDate1) AddRegistrationDeadline() *DateFormat6Choice {
	c.RegistrationDeadline = new(DateFormat6Choice)
	return c.RegistrationDeadline
}

func (c *CorporateActionDate1) AddResultsPublicationDate() *DateFormat6Choice {
	c.ResultsPublicationDate = new(DateFormat6Choice)
	return c.ResultsPublicationDate
}

func (c *CorporateActionDate1) AddDeadlineToSplit() *DateFormat6Choice {
	c.DeadlineToSplit = new(DateFormat6Choice)
	return c.DeadlineToSplit
}

func (c *CorporateActionDate1) AddDeadlineForTaxBreakdownInstruction() *DateFormat6Choice {
	c.DeadlineForTaxBreakdownInstruction = new(DateFormat6Choice)
	return c.DeadlineForTaxBreakdownInstruction
}

func (c *CorporateActionDate1) AddTradingSuspendedDate() *DateFormat6Choice {
	c.TradingSuspendedDate = new(DateFormat6Choice)
	return c.TradingSuspendedDate
}

func (c *CorporateActionDate1) AddUnconditionalDate() *DateFormat6Choice {
	c.UnconditionalDate = new(DateFormat6Choice)
	return c.UnconditionalDate
}

func (c *CorporateActionDate1) AddWhollyUnconditionalDate() *DateFormat6Choice {
	c.WhollyUnconditionalDate = new(DateFormat6Choice)
	return c.WhollyUnconditionalDate
}

func (c *CorporateActionDate1) AddExDividendDate() *DateFormat6Choice {
	c.ExDividendDate = new(DateFormat6Choice)
	return c.ExDividendDate
}

func (c *CorporateActionDate1) AddOfficialAnnouncementPublicationDate() *DateFormat6Choice {
	c.OfficialAnnouncementPublicationDate = new(DateFormat6Choice)
	return c.OfficialAnnouncementPublicationDate
}

func (c *CorporateActionDate1) AddSpecialExDate() *DateFormat6Choice {
	c.SpecialExDate = new(DateFormat6Choice)
	return c.SpecialExDate
}

func (c *CorporateActionDate1) AddGuaranteedParticipationDate() *DateFormat6Choice {
	c.GuaranteedParticipationDate = new(DateFormat6Choice)
	return c.GuaranteedParticipationDate
}

func (c *CorporateActionDate1) AddElectionToCounterpartyDeadline() *DateFormat6Choice {
	c.ElectionToCounterpartyDeadline = new(DateFormat6Choice)
	return c.ElectionToCounterpartyDeadline
}

func (c *CorporateActionDate1) AddLapsedDate() *DateFormat6Choice {
	c.LapsedDate = new(DateFormat6Choice)
	return c.LapsedDate
}

func (c *CorporateActionDate1) AddPaymentDate() *DateFormat6Choice {
	c.PaymentDate = new(DateFormat6Choice)
	return c.PaymentDate
}

func (c *CorporateActionDate1) AddThirdPartyDeadline() *DateFormat6Choice {
	c.ThirdPartyDeadline = new(DateFormat6Choice)
	return c.ThirdPartyDeadline
}

func (c *CorporateActionDate1) AddEarlyThirdPartyDeadline() *DateFormat6Choice {
	c.EarlyThirdPartyDeadline = new(DateFormat6Choice)
	return c.EarlyThirdPartyDeadline
}

func (c *CorporateActionDate1) AddMarketClaimTrackingEndDate() *DateFormat6Choice {
	c.MarketClaimTrackingEndDate = new(DateFormat6Choice)
	return c.MarketClaimTrackingEndDate
}

// Specifies corporate action dates.
type CorporateActionDate14 struct {

	// Date/time at which the issuer announced that a corporate action event will occur.
	AnnouncementDate *DateFormat19Choice `xml:"AnncmntDt,omitempty"`

	// Deadline by which the beneficial ownership of securities must be declared.
	CertificationDeadline *DateFormat19Choice `xml:"CertfctnDdln,omitempty"`

	// Date upon which the court provided approval.
	CourtApprovalDate *DateFormat19Choice `xml:"CrtApprvlDt,omitempty"`

	// First possible early closing date of an offer if different from the expiry date.
	EarlyClosingDate *DateFormat19Choice `xml:"EarlyClsgDt,omitempty"`

	// Date/time at which an event is officially effective from the issuer's perspective.
	EffectiveDate *DateFormat19Choice `xml:"FctvDt,omitempty"`

	// Date/Time on which all or part of any holding bought in a unit trust is subject to being treated as capital rather than income. This is normally one day after the previous distribution's ex date.
	EqualisationDate *DateFormat19Choice `xml:"EqulstnDt,omitempty"`

	// Date/time at which additional information on the event will be announced, for example, exchange ratio announcement date.
	FurtherDetailedAnnouncementDate *DateFormat19Choice `xml:"FrthrDtldAnncmntDt,omitempty"`

	// Date/time at which an index rate will be determined .
	IndexFixingDate *DateFormat19Choice `xml:"IndxFxgDt,omitempty"`

	// Date/time on which the lottery is run and applied to the holder's positions. This is also applicable to partial calls.
	LotteryDate *DateFormat19Choice `xml:"LtryDt,omitempty"`

	// Date/time to which the maturity date of an interest bearing security is extended.
	NewMaturityDate *DateFormat19Choice `xml:"NewMtrtyDt,omitempty"`

	// Date/time on which the bondholder's or shareholder's meeting will take place.
	MeetingDate *DateFormat19Choice `xml:"MtgDt,omitempty"`

	// Date/time at which the margin rate will be determined .
	MarginFixingDate *DateFormat19Choice `xml:"MrgnFxgDt,omitempty"`

	// Date/time (and time) at which an issuer will determine the proration amount/quantity of an offer.
	ProrationDate *DateFormat19Choice `xml:"PrratnDt,omitempty"`

	// Date/time at which positions are struck at the end of the day to note which parties will receive the relevant amount of entitlement, due to be distributed on payment date.
	RecordDate *DateFormat19Choice `xml:"RcrdDt,omitempty"`

	// Date/time on which instructions to register or registration details will be accepted.
	RegistrationDeadline *DateFormat19Choice `xml:"RegnDdln,omitempty"`

	// Date/time on which results are published, for example, results of an offer.
	ResultsPublicationDate *DateFormat19Choice `xml:"RsltsPblctnDt,omitempty"`

	// Deadline by which instructions must be received to split securities, for example, of physical certificates.
	DeadlineToSplit *DateFormat19Choice `xml:"DdlnToSplt,omitempty"`

	// Date/time on until which tax breakdown instructions will be accepted.
	DeadlineForTaxBreakdownInstruction *DateFormat19Choice `xml:"DdlnForTaxBrkdwnInstr,omitempty"`

	// Date/time at which trading of a security is suspended as the result of an event.
	TradingSuspendedDate *DateFormat19Choice `xml:"TradgSspdDt,omitempty"`

	// Date/time upon which the terms of the take-over become unconditional as to acceptances.
	UnconditionalDate *DateFormat19Choice `xml:"UcondlDt,omitempty"`

	// Date/time at on which all conditions, including regulatory, legal etc. pertaining to the take-over, have been met.
	WhollyUnconditionalDate *DateFormat19Choice `xml:"WhlyUcondlDt,omitempty"`

	// Date/time as from which trading (including exchange and OTC trading) occurs on the underlying security without the benefit.
	ExDividendDate *DateFormat19Choice `xml:"ExDvddDt,omitempty"`

	// Date/time at which the corporate action is legally announced by an official body, for example, publication by a governmental administration.
	OfficialAnnouncementPublicationDate *DateFormat19Choice `xml:"OffclAnncmntPblctnDt,omitempty"`

	// Date/time as from which 'special processing' can start to be used by participants for that event. Special processing is a means of marking a transaction, that would normally be traded ex or cum, as being traded cum or ex respectively,  for example, a transaction dealt 'special' after the ex date would result in the buyer being eligible for the entitlement. This is typically used in the UK and Irish markets.
	SpecialExDate *DateFormat19Choice `xml:"SpclExDt,omitempty"`

	// Last date/time by which a buying counterparty to a trade can be sure that it will have the right to participate in an event.
	GuaranteedParticipationDate *DateFormat19Choice `xml:"GrntedPrtcptnDt,omitempty"`

	// Deadline by which an entitled holder needs to advise their counterparty to a transaction of their election for a corporate action event.
	ElectionToCounterpartyDeadline *DateFormat19Choice `xml:"ElctnToCtrPtyDdln,omitempty"`

	// Date/time at which an event/offer is terminated or lapsed.
	LapsedDate *DateFormat19Choice `xml:"LpsdDt,omitempty"`

	// Date/time at which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat19Choice `xml:"PmtDt,omitempty"`

	// Date/Time by which the account owner must instruct directly another party, for example to provide documentation to an issuer agent.
	ThirdPartyDeadline *DateFormat19Choice `xml:"ThrdPtyDdln,omitempty"`

	// Date/Time set by the issuer agent as a first early deadline by which the account owner must instruct directly another party, possibly giving the holder eligibility to incentives. For example, to provide documentation to an issuer agent.
	EarlyThirdPartyDeadline *DateFormat19Choice `xml:"EarlyThrdPtyDdln,omitempty"`

	// Date by which the depository stops monitoring activities of the event, for instance, accounting and tracking activities for due bills end.
	MarketClaimTrackingEndDate *DateFormat19Choice `xml:"MktClmTrckgEndDt,omitempty"`

	// Last day an investor can become a lead plaintiff.
	LeadPlaintiffDeadline *DateFormat19Choice `xml:"LeadPlntffDdln,omitempty"`
}

func (c *CorporateActionDate14) AddAnnouncementDate() *DateFormat19Choice {
	c.AnnouncementDate = new(DateFormat19Choice)
	return c.AnnouncementDate
}

func (c *CorporateActionDate14) AddCertificationDeadline() *DateFormat19Choice {
	c.CertificationDeadline = new(DateFormat19Choice)
	return c.CertificationDeadline
}

func (c *CorporateActionDate14) AddCourtApprovalDate() *DateFormat19Choice {
	c.CourtApprovalDate = new(DateFormat19Choice)
	return c.CourtApprovalDate
}

func (c *CorporateActionDate14) AddEarlyClosingDate() *DateFormat19Choice {
	c.EarlyClosingDate = new(DateFormat19Choice)
	return c.EarlyClosingDate
}

func (c *CorporateActionDate14) AddEffectiveDate() *DateFormat19Choice {
	c.EffectiveDate = new(DateFormat19Choice)
	return c.EffectiveDate
}

func (c *CorporateActionDate14) AddEqualisationDate() *DateFormat19Choice {
	c.EqualisationDate = new(DateFormat19Choice)
	return c.EqualisationDate
}

func (c *CorporateActionDate14) AddFurtherDetailedAnnouncementDate() *DateFormat19Choice {
	c.FurtherDetailedAnnouncementDate = new(DateFormat19Choice)
	return c.FurtherDetailedAnnouncementDate
}

func (c *CorporateActionDate14) AddIndexFixingDate() *DateFormat19Choice {
	c.IndexFixingDate = new(DateFormat19Choice)
	return c.IndexFixingDate
}

func (c *CorporateActionDate14) AddLotteryDate() *DateFormat19Choice {
	c.LotteryDate = new(DateFormat19Choice)
	return c.LotteryDate
}

func (c *CorporateActionDate14) AddNewMaturityDate() *DateFormat19Choice {
	c.NewMaturityDate = new(DateFormat19Choice)
	return c.NewMaturityDate
}

func (c *CorporateActionDate14) AddMeetingDate() *DateFormat19Choice {
	c.MeetingDate = new(DateFormat19Choice)
	return c.MeetingDate
}

func (c *CorporateActionDate14) AddMarginFixingDate() *DateFormat19Choice {
	c.MarginFixingDate = new(DateFormat19Choice)
	return c.MarginFixingDate
}

func (c *CorporateActionDate14) AddProrationDate() *DateFormat19Choice {
	c.ProrationDate = new(DateFormat19Choice)
	return c.ProrationDate
}

func (c *CorporateActionDate14) AddRecordDate() *DateFormat19Choice {
	c.RecordDate = new(DateFormat19Choice)
	return c.RecordDate
}

func (c *CorporateActionDate14) AddRegistrationDeadline() *DateFormat19Choice {
	c.RegistrationDeadline = new(DateFormat19Choice)
	return c.RegistrationDeadline
}

func (c *CorporateActionDate14) AddResultsPublicationDate() *DateFormat19Choice {
	c.ResultsPublicationDate = new(DateFormat19Choice)
	return c.ResultsPublicationDate
}

func (c *CorporateActionDate14) AddDeadlineToSplit() *DateFormat19Choice {
	c.DeadlineToSplit = new(DateFormat19Choice)
	return c.DeadlineToSplit
}

func (c *CorporateActionDate14) AddDeadlineForTaxBreakdownInstruction() *DateFormat19Choice {
	c.DeadlineForTaxBreakdownInstruction = new(DateFormat19Choice)
	return c.DeadlineForTaxBreakdownInstruction
}

func (c *CorporateActionDate14) AddTradingSuspendedDate() *DateFormat19Choice {
	c.TradingSuspendedDate = new(DateFormat19Choice)
	return c.TradingSuspendedDate
}

func (c *CorporateActionDate14) AddUnconditionalDate() *DateFormat19Choice {
	c.UnconditionalDate = new(DateFormat19Choice)
	return c.UnconditionalDate
}

func (c *CorporateActionDate14) AddWhollyUnconditionalDate() *DateFormat19Choice {
	c.WhollyUnconditionalDate = new(DateFormat19Choice)
	return c.WhollyUnconditionalDate
}

func (c *CorporateActionDate14) AddExDividendDate() *DateFormat19Choice {
	c.ExDividendDate = new(DateFormat19Choice)
	return c.ExDividendDate
}

func (c *CorporateActionDate14) AddOfficialAnnouncementPublicationDate() *DateFormat19Choice {
	c.OfficialAnnouncementPublicationDate = new(DateFormat19Choice)
	return c.OfficialAnnouncementPublicationDate
}

func (c *CorporateActionDate14) AddSpecialExDate() *DateFormat19Choice {
	c.SpecialExDate = new(DateFormat19Choice)
	return c.SpecialExDate
}

func (c *CorporateActionDate14) AddGuaranteedParticipationDate() *DateFormat19Choice {
	c.GuaranteedParticipationDate = new(DateFormat19Choice)
	return c.GuaranteedParticipationDate
}

func (c *CorporateActionDate14) AddElectionToCounterpartyDeadline() *DateFormat19Choice {
	c.ElectionToCounterpartyDeadline = new(DateFormat19Choice)
	return c.ElectionToCounterpartyDeadline
}

func (c *CorporateActionDate14) AddLapsedDate() *DateFormat19Choice {
	c.LapsedDate = new(DateFormat19Choice)
	return c.LapsedDate
}

func (c *CorporateActionDate14) AddPaymentDate() *DateFormat19Choice {
	c.PaymentDate = new(DateFormat19Choice)
	return c.PaymentDate
}

func (c *CorporateActionDate14) AddThirdPartyDeadline() *DateFormat19Choice {
	c.ThirdPartyDeadline = new(DateFormat19Choice)
	return c.ThirdPartyDeadline
}

func (c *CorporateActionDate14) AddEarlyThirdPartyDeadline() *DateFormat19Choice {
	c.EarlyThirdPartyDeadline = new(DateFormat19Choice)
	return c.EarlyThirdPartyDeadline
}

func (c *CorporateActionDate14) AddMarketClaimTrackingEndDate() *DateFormat19Choice {
	c.MarketClaimTrackingEndDate = new(DateFormat19Choice)
	return c.MarketClaimTrackingEndDate
}

func (c *CorporateActionDate14) AddLeadPlaintiffDeadline() *DateFormat19Choice {
	c.LeadPlaintiffDeadline = new(DateFormat19Choice)
	return c.LeadPlaintiffDeadline
}

// Specifies corporate action dates.
type CorporateActionDate15 struct {

	// Date/time that the account servicer has set as the deadline to respond, with instructions, to an outstanding event, giving the holder eligibility to incentives. This time is dependent on the reference time zone of the account servicer as specified in a Service Level Agreement (SLA).
	EarlyResponseDeadline *DateFormat19Choice `xml:"EarlyRspnDdln,omitempty"`

	// Last day a holder can deliver the securities that it had elected on and/or previously protected.
	CoverExpirationDate *DateFormat19Choice `xml:"CoverXprtnDt,omitempty"`

	// Last date/time a holder can request to defer delivery of securities pursuant to a notice of guaranteed delivery or other required documentation.
	ProtectDate *DateFormat19Choice `xml:"PrtctDt,omitempty"`

	// Issuer or issuer's agent deadline to respond, with an instruction, to an outstanding offer or privilege.
	MarketDeadline *DateFormat19Choice `xml:"MktDdln,omitempty"`

	// Date/time at which the account servicer has set as the deadline to respond, with instructions, to an outstanding event. This time is dependent on the reference time zone of the account servicer as specified in a Service Level Agreement (SLA).
	ResponseDeadline *DateFormat20Choice `xml:"RspnDdln,omitempty"`

	// Date/time at which an order expires or on which a privilege or offer terminates.
	ExpiryDate *DateFormat19Choice `xml:"XpryDt,omitempty"`

	// Date/time by which cash must be in place in order to take part in the event.
	SubscriptionCostDebitDate *DateFormat19Choice `xml:"SbcptCostDbtDt,omitempty"`

	// Last day that a participant of the depository can deliver securities that it had elected on and/or previously protected.
	DepositoryCoverExpirationDate *DateFormat19Choice `xml:"DpstryCoverXprtnDt,omitempty"`
}

func (c *CorporateActionDate15) AddEarlyResponseDeadline() *DateFormat19Choice {
	c.EarlyResponseDeadline = new(DateFormat19Choice)
	return c.EarlyResponseDeadline
}

func (c *CorporateActionDate15) AddCoverExpirationDate() *DateFormat19Choice {
	c.CoverExpirationDate = new(DateFormat19Choice)
	return c.CoverExpirationDate
}

func (c *CorporateActionDate15) AddProtectDate() *DateFormat19Choice {
	c.ProtectDate = new(DateFormat19Choice)
	return c.ProtectDate
}

func (c *CorporateActionDate15) AddMarketDeadline() *DateFormat19Choice {
	c.MarketDeadline = new(DateFormat19Choice)
	return c.MarketDeadline
}

func (c *CorporateActionDate15) AddResponseDeadline() *DateFormat20Choice {
	c.ResponseDeadline = new(DateFormat20Choice)
	return c.ResponseDeadline
}

func (c *CorporateActionDate15) AddExpiryDate() *DateFormat19Choice {
	c.ExpiryDate = new(DateFormat19Choice)
	return c.ExpiryDate
}

func (c *CorporateActionDate15) AddSubscriptionCostDebitDate() *DateFormat19Choice {
	c.SubscriptionCostDebitDate = new(DateFormat19Choice)
	return c.SubscriptionCostDebitDate
}

func (c *CorporateActionDate15) AddDepositoryCoverExpirationDate() *DateFormat19Choice {
	c.DepositoryCoverExpirationDate = new(DateFormat19Choice)
	return c.DepositoryCoverExpirationDate
}

// Specifies corporate action dates.
type CorporateActionDate17 struct {

	// Date on which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat19Choice `xml:"PmtDt"`

	// Date at which assets become available to the account owner (in a credit entry), or cease to be available to the account owner (in a debit entry).
	ValueDate *DateFormat11Choice `xml:"ValDt,omitempty"`

	// Date/time at which a foreign exchange rate will be determined.
	ForeignExchangeRateFixingDate *DateFormat19Choice `xml:"FXRateFxgDt,omitempty"`

	// Date on which a payment can be made, for example, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat19Choice `xml:"EarlstPmtDt,omitempty"`
}

func (c *CorporateActionDate17) AddPaymentDate() *DateFormat19Choice {
	c.PaymentDate = new(DateFormat19Choice)
	return c.PaymentDate
}

func (c *CorporateActionDate17) AddValueDate() *DateFormat11Choice {
	c.ValueDate = new(DateFormat11Choice)
	return c.ValueDate
}

func (c *CorporateActionDate17) AddForeignExchangeRateFixingDate() *DateFormat19Choice {
	c.ForeignExchangeRateFixingDate = new(DateFormat19Choice)
	return c.ForeignExchangeRateFixingDate
}

func (c *CorporateActionDate17) AddEarliestPaymentDate() *DateFormat19Choice {
	c.EarliestPaymentDate = new(DateFormat19Choice)
	return c.EarliestPaymentDate
}

// Specifies corporate action date.
type CorporateActionDate18 struct {

	// Date/time at which the account servicer has set as the deadline to respond, with instructions, to an outstanding event. This time is dependent on the reference time zone of the account servicer as specified in a Service Level Agreement (SLA).
	ResponseDeadline *DateFormat19Choice `xml:"RspnDdln,omitempty"`

	// Date/time by which cash must be in place in order to take part in the event.
	SubscriptionCostDebitDate *DateFormat19Choice `xml:"SbcptCostDbtDt,omitempty"`

	// Issuer or issuer's agent deadline to respond, with an instruction, to an outstanding offer or privilege.
	MarketDeadline *DateFormat19Choice `xml:"MktDdln,omitempty"`

	// Date/time at which an order expires or on which a privilege or offer terminates.
	ExpiryDate *DateFormat19Choice `xml:"XpryDt,omitempty"`

	// Last day a holder can deliver the securities that it had elected on and/or previously protected.
	CoverExpirationDate *DateFormat19Choice `xml:"CoverXprtnDt,omitempty"`

	// Last date/time a holder can request to defer delivery of securities pursuant to a notice of guaranteed delivery or other required documentation.
	ProtectDate *DateFormat19Choice `xml:"PrtctDt,omitempty"`

	// Date/time at which the deal (rights) was agreed.
	TradingDate *DateFormat19Choice `xml:"TradgDt,omitempty"`
}

func (c *CorporateActionDate18) AddResponseDeadline() *DateFormat19Choice {
	c.ResponseDeadline = new(DateFormat19Choice)
	return c.ResponseDeadline
}

func (c *CorporateActionDate18) AddSubscriptionCostDebitDate() *DateFormat19Choice {
	c.SubscriptionCostDebitDate = new(DateFormat19Choice)
	return c.SubscriptionCostDebitDate
}

func (c *CorporateActionDate18) AddMarketDeadline() *DateFormat19Choice {
	c.MarketDeadline = new(DateFormat19Choice)
	return c.MarketDeadline
}

func (c *CorporateActionDate18) AddExpiryDate() *DateFormat19Choice {
	c.ExpiryDate = new(DateFormat19Choice)
	return c.ExpiryDate
}

func (c *CorporateActionDate18) AddCoverExpirationDate() *DateFormat19Choice {
	c.CoverExpirationDate = new(DateFormat19Choice)
	return c.CoverExpirationDate
}

func (c *CorporateActionDate18) AddProtectDate() *DateFormat19Choice {
	c.ProtectDate = new(DateFormat19Choice)
	return c.ProtectDate
}

func (c *CorporateActionDate18) AddTradingDate() *DateFormat19Choice {
	c.TradingDate = new(DateFormat19Choice)
	return c.TradingDate
}

// Specifies corporate action dates.
type CorporateActionDate2 struct {

	// Date on which positions are struck at the end of the day to note which parties will receive the relevant amount of entitlement, due to be distributed on payment date.
	RecordDate *DateFormat4Choice `xml:"RcrdDt,omitempty"`

	// Date on which a process is to be completed or becomes effective.
	EffectiveDate *DateFormat4Choice `xml:"FctvDt,omitempty"`

	// Last day a holder can deliver the securities that it had previously protected.
	CoverExpirationDate *DateFormat4Choice `xml:"CoverXprtnDt,omitempty"`

	// Date on which all or part of any holding bought in a unit trust is subject to being treated as capital rather than income. This is normally one day after the previous distribution's ex date.
	EqualisationDate *DateFormat4Choice `xml:"EqulstnDt,omitempty"`

	// Date/time at which the margin rate will be determined .
	MarginFixingDate *DateFormat4Choice `xml:"MrgnFxgDt,omitempty"`

	// Date on which the lottery is run and applied to the holder's positions. This is also applicable to partial calls.
	LotteryDate *DateFormat4Choice `xml:"LtryDt,omitempty"`

	// Last date a holder can request to defer delivery of securities pursuant to a notice of guaranteed delivery or other required documentation.
	ProtectDate *DateFormat4Choice `xml:"PrtctDt,omitempty"`

	// Date upon which the terms of the take-over become unconditional as to acceptances.
	UnconditionalDate *DateFormat4Choice `xml:"UcondlDt,omitempty"`

	// Date on which all conditions, including regulatory, legal etc. pertaining to the take-over, have been met.
	WhollyUnconditionalDate *DateFormat4Choice `xml:"WhlyUcondlDt,omitempty"`

	// Date on which results are published, eg, results of an offer.
	ResultsPublicationDate *DateFormat4Choice `xml:"RsltsPblctnDt,omitempty"`

	// Date/time upon which the High Court provided approval.
	CourtApprovalDate *DateFormat4Choice `xml:"CrtApprvlDt,omitempty"`

	// First possible early closing date of an offer if different from the expiry date.
	EarlyClosingDate *DateFormat4Choice `xml:"EarlyClsgDt,omitempty"`

	// Date/time as from which trading (including exchange and OTC trading) occurs on the underlying security without the benefit.
	ExDividendDate *DateFormat4Choice `xml:"ExDvddDt,omitempty"`

	// Date/time at which an index rate will be determined .
	IndexFixingDate *DateFormat4Choice `xml:"IndxFxgDt,omitempty"`

	// Date on which an interest bearing financial instrument becomes due and principal is repaid by the issuer to the investor.
	MaturityDate *DateFormat4Choice `xml:"MtrtyDt,omitempty"`

	// Date on which trading of a security is suspended as the result of an event.
	TradingSuspendedDate *DateFormat4Choice `xml:"TradgSspdDt,omitempty"`

	// Deadline by which the beneficial ownership of securities must be declared.
	CertificationDeadline *DateFormat4Choice `xml:"CertfctnDdln,omitempty"`

	// Date/time at which the securities will be redeemed (early) for payment of principal.
	RedemptionDate *DateFormat4Choice `xml:"RedDt,omitempty"`

	// Date on which instructions to register or registration details will be accepted.
	RegistrationDeadline *DateFormat4Choice `xml:"RegnDdln,omitempty"`

	// Date (and time) at which an issuer will determine the proration amount/quantity of an offer.
	ProrationDate *DateFormat4Choice `xml:"PrratnDt,omitempty"`

	// Date on until which tax breakdown instructions will be accepted.
	DeadlineForTaxBreakdownInstruction *DateFormat4Choice `xml:"DdlnForTaxBrkdwnInstr,omitempty"`

	// Date/time at which an event/offer is terminated or lapsed.
	LapsedDate *DateFormat4Choice `xml:"LpsdDt,omitempty"`

	// Last date/time by which a buying counterparty to a trade can be sure that it will have the right to participate in an event.
	GuaranteedParticipationDate *DateFormat4Choice `xml:"GrntedPrtcptnDt,omitempty"`

	// Deadline by which an entitled holder needs to advise their counterparty to a transaction of their election for a corporate action event.
	ElectionToCounterpartyDeadline *DateFormat4Choice `xml:"ElctnToCtrPtyDdln,omitempty"`

	// Date/time as from which 'special processing' can start to be used by participants for that event. Special processing is a means of marking a transaction, that would normally be traded ex or cum, as being traded cum or ex respectively, eg, a transaction dealt 'special' after the ex date would result in the buyer being eligible for the entitlement. This is typically used in the UK and Irish markets.
	SpecialExDate *DateFormat4Choice `xml:"SpclExDt,omitempty"`
}

func (c *CorporateActionDate2) AddRecordDate() *DateFormat4Choice {
	c.RecordDate = new(DateFormat4Choice)
	return c.RecordDate
}

func (c *CorporateActionDate2) AddEffectiveDate() *DateFormat4Choice {
	c.EffectiveDate = new(DateFormat4Choice)
	return c.EffectiveDate
}

func (c *CorporateActionDate2) AddCoverExpirationDate() *DateFormat4Choice {
	c.CoverExpirationDate = new(DateFormat4Choice)
	return c.CoverExpirationDate
}

func (c *CorporateActionDate2) AddEqualisationDate() *DateFormat4Choice {
	c.EqualisationDate = new(DateFormat4Choice)
	return c.EqualisationDate
}

func (c *CorporateActionDate2) AddMarginFixingDate() *DateFormat4Choice {
	c.MarginFixingDate = new(DateFormat4Choice)
	return c.MarginFixingDate
}

func (c *CorporateActionDate2) AddLotteryDate() *DateFormat4Choice {
	c.LotteryDate = new(DateFormat4Choice)
	return c.LotteryDate
}

func (c *CorporateActionDate2) AddProtectDate() *DateFormat4Choice {
	c.ProtectDate = new(DateFormat4Choice)
	return c.ProtectDate
}

func (c *CorporateActionDate2) AddUnconditionalDate() *DateFormat4Choice {
	c.UnconditionalDate = new(DateFormat4Choice)
	return c.UnconditionalDate
}

func (c *CorporateActionDate2) AddWhollyUnconditionalDate() *DateFormat4Choice {
	c.WhollyUnconditionalDate = new(DateFormat4Choice)
	return c.WhollyUnconditionalDate
}

func (c *CorporateActionDate2) AddResultsPublicationDate() *DateFormat4Choice {
	c.ResultsPublicationDate = new(DateFormat4Choice)
	return c.ResultsPublicationDate
}

func (c *CorporateActionDate2) AddCourtApprovalDate() *DateFormat4Choice {
	c.CourtApprovalDate = new(DateFormat4Choice)
	return c.CourtApprovalDate
}

func (c *CorporateActionDate2) AddEarlyClosingDate() *DateFormat4Choice {
	c.EarlyClosingDate = new(DateFormat4Choice)
	return c.EarlyClosingDate
}

func (c *CorporateActionDate2) AddExDividendDate() *DateFormat4Choice {
	c.ExDividendDate = new(DateFormat4Choice)
	return c.ExDividendDate
}

func (c *CorporateActionDate2) AddIndexFixingDate() *DateFormat4Choice {
	c.IndexFixingDate = new(DateFormat4Choice)
	return c.IndexFixingDate
}

func (c *CorporateActionDate2) AddMaturityDate() *DateFormat4Choice {
	c.MaturityDate = new(DateFormat4Choice)
	return c.MaturityDate
}

func (c *CorporateActionDate2) AddTradingSuspendedDate() *DateFormat4Choice {
	c.TradingSuspendedDate = new(DateFormat4Choice)
	return c.TradingSuspendedDate
}

func (c *CorporateActionDate2) AddCertificationDeadline() *DateFormat4Choice {
	c.CertificationDeadline = new(DateFormat4Choice)
	return c.CertificationDeadline
}

func (c *CorporateActionDate2) AddRedemptionDate() *DateFormat4Choice {
	c.RedemptionDate = new(DateFormat4Choice)
	return c.RedemptionDate
}

func (c *CorporateActionDate2) AddRegistrationDeadline() *DateFormat4Choice {
	c.RegistrationDeadline = new(DateFormat4Choice)
	return c.RegistrationDeadline
}

func (c *CorporateActionDate2) AddProrationDate() *DateFormat4Choice {
	c.ProrationDate = new(DateFormat4Choice)
	return c.ProrationDate
}

func (c *CorporateActionDate2) AddDeadlineForTaxBreakdownInstruction() *DateFormat4Choice {
	c.DeadlineForTaxBreakdownInstruction = new(DateFormat4Choice)
	return c.DeadlineForTaxBreakdownInstruction
}

func (c *CorporateActionDate2) AddLapsedDate() *DateFormat4Choice {
	c.LapsedDate = new(DateFormat4Choice)
	return c.LapsedDate
}

func (c *CorporateActionDate2) AddGuaranteedParticipationDate() *DateFormat4Choice {
	c.GuaranteedParticipationDate = new(DateFormat4Choice)
	return c.GuaranteedParticipationDate
}

func (c *CorporateActionDate2) AddElectionToCounterpartyDeadline() *DateFormat4Choice {
	c.ElectionToCounterpartyDeadline = new(DateFormat4Choice)
	return c.ElectionToCounterpartyDeadline
}

func (c *CorporateActionDate2) AddSpecialExDate() *DateFormat4Choice {
	c.SpecialExDate = new(DateFormat4Choice)
	return c.SpecialExDate
}

// Specifies corporate action dates.
type CorporateActionDate22 struct {

	// Date/time at which the issuer announced that a corporate action event will occur.
	AnnouncementDate *DateFormat19Choice `xml:"AnncmntDt,omitempty"`

	// Deadline by which the beneficial ownership of securities must be declared.
	CertificationDeadline *DateFormat19Choice `xml:"CertfctnDdln,omitempty"`

	// Date upon which the court provided approval.
	CourtApprovalDate *DateFormat19Choice `xml:"CrtApprvlDt,omitempty"`

	// First possible early closing date of an offer if different from the expiry date.
	EarlyClosingDate *DateFormat19Choice `xml:"EarlyClsgDt,omitempty"`

	// Date/time at which an event is officially effective from the issuer's perspective.
	EffectiveDate *DateFormat19Choice `xml:"FctvDt,omitempty"`

	// Date/Time on which all or part of any holding bought in a unit trust is subject to being treated as capital rather than income. This is normally one day after the previous distribution's ex date.
	EqualisationDate *DateFormat19Choice `xml:"EqulstnDt,omitempty"`

	// Date/time at which additional information on the event will be announced, for example, exchange ratio announcement date.
	FurtherDetailedAnnouncementDate *DateFormat19Choice `xml:"FrthrDtldAnncmntDt,omitempty"`

	// Date/time at which an index / rate / price / value will be determined.
	FixingDate *DateFormat19Choice `xml:"FxgDt,omitempty"`

	// Date/time on which the lottery is run and applied to the holder's positions. This is also applicable to partial calls.
	LotteryDate *DateFormat19Choice `xml:"LtryDt,omitempty"`

	// Date/time to which the maturity date of an interest bearing security is extended.
	NewMaturityDate *DateFormat19Choice `xml:"NewMtrtyDt,omitempty"`

	// Date/time on which the bondholder's or shareholder's meeting will take place.
	MeetingDate *DateFormat19Choice `xml:"MtgDt,omitempty"`

	// Date/time at which the margin rate will be determined .
	MarginFixingDate *DateFormat19Choice `xml:"MrgnFxgDt,omitempty"`

	// Date/time (and time) at which an issuer will determine the proration amount/quantity of an offer.
	ProrationDate *DateFormat19Choice `xml:"PrratnDt,omitempty"`

	// Date/time at which positions are struck at the end of the day to note which parties will receive the relevant amount of entitlement, due to be distributed on payment date.
	RecordDate *DateFormat19Choice `xml:"RcrdDt,omitempty"`

	// Date/time on which instructions to register or registration details will be accepted.
	RegistrationDeadline *DateFormat19Choice `xml:"RegnDdln,omitempty"`

	// Date/time on which results are published, for example, results of an offer.
	ResultsPublicationDate *DateFormat19Choice `xml:"RsltsPblctnDt,omitempty"`

	// Deadline by which instructions must be received to split securities, for example, of physical certificates.
	DeadlineToSplit *DateFormat19Choice `xml:"DdlnToSplt,omitempty"`

	// Date/time on until which tax breakdown instructions will be accepted.
	DeadlineForTaxBreakdownInstruction *DateFormat19Choice `xml:"DdlnForTaxBrkdwnInstr,omitempty"`

	// Date/time at which trading of a security is suspended as the result of an event.
	TradingSuspendedDate *DateFormat19Choice `xml:"TradgSspdDt,omitempty"`

	// Date/time upon which the terms of the take-over become unconditional as to acceptances.
	UnconditionalDate *DateFormat19Choice `xml:"UcondlDt,omitempty"`

	// Date/time at on which all conditions, including regulatory, legal etc. pertaining to the take-over, have been met.
	WhollyUnconditionalDate *DateFormat19Choice `xml:"WhlyUcondlDt,omitempty"`

	// Date/time as from which trading (including exchange and OTC trading) occurs on the underlying security without the benefit.
	ExDividendDate *DateFormat19Choice `xml:"ExDvddDt,omitempty"`

	// Date/time at which the corporate action is legally announced by an official body, for example, publication by a governmental administration.
	OfficialAnnouncementPublicationDate *DateFormat19Choice `xml:"OffclAnncmntPblctnDt,omitempty"`

	// Date/time as from which 'special processing' can start to be used by participants for that event. Special processing is a means of marking a transaction, that would normally be traded ex or cum, as being traded cum or ex respectively,  for example, a transaction dealt 'special' after the ex date would result in the buyer being eligible for the entitlement. This is typically used in the UK and Irish markets.
	SpecialExDate *DateFormat19Choice `xml:"SpclExDt,omitempty"`

	// Last date/time by which a buying counterparty to a trade can be sure that it will have the right to participate in an event.
	GuaranteedParticipationDate *DateFormat19Choice `xml:"GrntedPrtcptnDt,omitempty"`

	// Deadline by which an entitled holder needs to advise their counterparty to a transaction of their election for a corporate action event.
	ElectionToCounterpartyDeadline *DateFormat19Choice `xml:"ElctnToCtrPtyDdln,omitempty"`

	// Date/time at which an event/offer is terminated or lapsed.
	LapsedDate *DateFormat19Choice `xml:"LpsdDt,omitempty"`

	// Date/time at which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat19Choice `xml:"PmtDt,omitempty"`

	// Date/Time by which the account owner must instruct directly another party, for example to provide documentation to an issuer agent.
	ThirdPartyDeadline *DateFormat19Choice `xml:"ThrdPtyDdln,omitempty"`

	// Date/Time set by the issuer agent as a first early deadline by which the account owner must instruct directly another party, possibly giving the holder eligibility to incentives. For example, to provide documentation to an issuer agent.
	EarlyThirdPartyDeadline *DateFormat19Choice `xml:"EarlyThrdPtyDdln,omitempty"`

	// Date by which the depository stops monitoring activities of the event, for instance, accounting and tracking activities for due bills end.
	MarketClaimTrackingEndDate *DateFormat19Choice `xml:"MktClmTrckgEndDt,omitempty"`

	// Last day an investor can become a lead plaintiff.
	LeadPlaintiffDeadline *DateFormat19Choice `xml:"LeadPlntffDdln,omitempty"`

	// Date on which the action was filed at the applicable court.
	FilingDate *DateFormat16Choice `xml:"FilgDt,omitempty"`

	// Date for the hearing between the plaintiff and defendant, as set by the court.
	HearingDate *DateFormat16Choice `xml:"HrgDt,omitempty"`
}

func (c *CorporateActionDate22) AddAnnouncementDate() *DateFormat19Choice {
	c.AnnouncementDate = new(DateFormat19Choice)
	return c.AnnouncementDate
}

func (c *CorporateActionDate22) AddCertificationDeadline() *DateFormat19Choice {
	c.CertificationDeadline = new(DateFormat19Choice)
	return c.CertificationDeadline
}

func (c *CorporateActionDate22) AddCourtApprovalDate() *DateFormat19Choice {
	c.CourtApprovalDate = new(DateFormat19Choice)
	return c.CourtApprovalDate
}

func (c *CorporateActionDate22) AddEarlyClosingDate() *DateFormat19Choice {
	c.EarlyClosingDate = new(DateFormat19Choice)
	return c.EarlyClosingDate
}

func (c *CorporateActionDate22) AddEffectiveDate() *DateFormat19Choice {
	c.EffectiveDate = new(DateFormat19Choice)
	return c.EffectiveDate
}

func (c *CorporateActionDate22) AddEqualisationDate() *DateFormat19Choice {
	c.EqualisationDate = new(DateFormat19Choice)
	return c.EqualisationDate
}

func (c *CorporateActionDate22) AddFurtherDetailedAnnouncementDate() *DateFormat19Choice {
	c.FurtherDetailedAnnouncementDate = new(DateFormat19Choice)
	return c.FurtherDetailedAnnouncementDate
}

func (c *CorporateActionDate22) AddFixingDate() *DateFormat19Choice {
	c.FixingDate = new(DateFormat19Choice)
	return c.FixingDate
}

func (c *CorporateActionDate22) AddLotteryDate() *DateFormat19Choice {
	c.LotteryDate = new(DateFormat19Choice)
	return c.LotteryDate
}

func (c *CorporateActionDate22) AddNewMaturityDate() *DateFormat19Choice {
	c.NewMaturityDate = new(DateFormat19Choice)
	return c.NewMaturityDate
}

func (c *CorporateActionDate22) AddMeetingDate() *DateFormat19Choice {
	c.MeetingDate = new(DateFormat19Choice)
	return c.MeetingDate
}

func (c *CorporateActionDate22) AddMarginFixingDate() *DateFormat19Choice {
	c.MarginFixingDate = new(DateFormat19Choice)
	return c.MarginFixingDate
}

func (c *CorporateActionDate22) AddProrationDate() *DateFormat19Choice {
	c.ProrationDate = new(DateFormat19Choice)
	return c.ProrationDate
}

func (c *CorporateActionDate22) AddRecordDate() *DateFormat19Choice {
	c.RecordDate = new(DateFormat19Choice)
	return c.RecordDate
}

func (c *CorporateActionDate22) AddRegistrationDeadline() *DateFormat19Choice {
	c.RegistrationDeadline = new(DateFormat19Choice)
	return c.RegistrationDeadline
}

func (c *CorporateActionDate22) AddResultsPublicationDate() *DateFormat19Choice {
	c.ResultsPublicationDate = new(DateFormat19Choice)
	return c.ResultsPublicationDate
}

func (c *CorporateActionDate22) AddDeadlineToSplit() *DateFormat19Choice {
	c.DeadlineToSplit = new(DateFormat19Choice)
	return c.DeadlineToSplit
}

func (c *CorporateActionDate22) AddDeadlineForTaxBreakdownInstruction() *DateFormat19Choice {
	c.DeadlineForTaxBreakdownInstruction = new(DateFormat19Choice)
	return c.DeadlineForTaxBreakdownInstruction
}

func (c *CorporateActionDate22) AddTradingSuspendedDate() *DateFormat19Choice {
	c.TradingSuspendedDate = new(DateFormat19Choice)
	return c.TradingSuspendedDate
}

func (c *CorporateActionDate22) AddUnconditionalDate() *DateFormat19Choice {
	c.UnconditionalDate = new(DateFormat19Choice)
	return c.UnconditionalDate
}

func (c *CorporateActionDate22) AddWhollyUnconditionalDate() *DateFormat19Choice {
	c.WhollyUnconditionalDate = new(DateFormat19Choice)
	return c.WhollyUnconditionalDate
}

func (c *CorporateActionDate22) AddExDividendDate() *DateFormat19Choice {
	c.ExDividendDate = new(DateFormat19Choice)
	return c.ExDividendDate
}

func (c *CorporateActionDate22) AddOfficialAnnouncementPublicationDate() *DateFormat19Choice {
	c.OfficialAnnouncementPublicationDate = new(DateFormat19Choice)
	return c.OfficialAnnouncementPublicationDate
}

func (c *CorporateActionDate22) AddSpecialExDate() *DateFormat19Choice {
	c.SpecialExDate = new(DateFormat19Choice)
	return c.SpecialExDate
}

func (c *CorporateActionDate22) AddGuaranteedParticipationDate() *DateFormat19Choice {
	c.GuaranteedParticipationDate = new(DateFormat19Choice)
	return c.GuaranteedParticipationDate
}

func (c *CorporateActionDate22) AddElectionToCounterpartyDeadline() *DateFormat19Choice {
	c.ElectionToCounterpartyDeadline = new(DateFormat19Choice)
	return c.ElectionToCounterpartyDeadline
}

func (c *CorporateActionDate22) AddLapsedDate() *DateFormat19Choice {
	c.LapsedDate = new(DateFormat19Choice)
	return c.LapsedDate
}

func (c *CorporateActionDate22) AddPaymentDate() *DateFormat19Choice {
	c.PaymentDate = new(DateFormat19Choice)
	return c.PaymentDate
}

func (c *CorporateActionDate22) AddThirdPartyDeadline() *DateFormat19Choice {
	c.ThirdPartyDeadline = new(DateFormat19Choice)
	return c.ThirdPartyDeadline
}

func (c *CorporateActionDate22) AddEarlyThirdPartyDeadline() *DateFormat19Choice {
	c.EarlyThirdPartyDeadline = new(DateFormat19Choice)
	return c.EarlyThirdPartyDeadline
}

func (c *CorporateActionDate22) AddMarketClaimTrackingEndDate() *DateFormat19Choice {
	c.MarketClaimTrackingEndDate = new(DateFormat19Choice)
	return c.MarketClaimTrackingEndDate
}

func (c *CorporateActionDate22) AddLeadPlaintiffDeadline() *DateFormat19Choice {
	c.LeadPlaintiffDeadline = new(DateFormat19Choice)
	return c.LeadPlaintiffDeadline
}

func (c *CorporateActionDate22) AddFilingDate() *DateFormat16Choice {
	c.FilingDate = new(DateFormat16Choice)
	return c.FilingDate
}

func (c *CorporateActionDate22) AddHearingDate() *DateFormat16Choice {
	c.HearingDate = new(DateFormat16Choice)
	return c.HearingDate
}

// Specifies corporate action dates.
type CorporateActionDate23 struct {

	// Date on which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat19Choice `xml:"PmtDt"`

	// Date/time when calculating economic benefit for a cash amount.
	ValueDate *DateFormat11Choice `xml:"ValDt,omitempty"`

	// Date/time at which a foreign exchange rate will be determined.
	ForeignExchangeRateFixingDate *DateFormat19Choice `xml:"FXRateFxgDt,omitempty"`

	// Date on which a payment can be made, for example, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat19Choice `xml:"EarlstPmtDt,omitempty"`
}

func (c *CorporateActionDate23) AddPaymentDate() *DateFormat19Choice {
	c.PaymentDate = new(DateFormat19Choice)
	return c.PaymentDate
}

func (c *CorporateActionDate23) AddValueDate() *DateFormat11Choice {
	c.ValueDate = new(DateFormat11Choice)
	return c.ValueDate
}

func (c *CorporateActionDate23) AddForeignExchangeRateFixingDate() *DateFormat19Choice {
	c.ForeignExchangeRateFixingDate = new(DateFormat19Choice)
	return c.ForeignExchangeRateFixingDate
}

func (c *CorporateActionDate23) AddEarliestPaymentDate() *DateFormat19Choice {
	c.EarliestPaymentDate = new(DateFormat19Choice)
	return c.EarliestPaymentDate
}

// Specifies corporate action dates.
type CorporateActionDate24 struct {

	// Date/Time of the posting (credit or debit) to the account.
	PostingDate *DateAndDateTimeChoice `xml:"PstngDt"`

	// Date/time when calculating economic benefit for a cash amount.
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`

	// Date/time at which a foreign exchange rate will be determined.
	ForeignExchangeRateFixingDate *DateAndDateTimeChoice `xml:"FXRateFxgDt,omitempty"`

	// Date on which a payment can be made, for example, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateAndDateTimeChoice `xml:"EarlstPmtDt,omitempty"`

	// Date on which the distribution is due to take place (cash and/or securities).
	PaymentDate *DateAndDateTimeChoice `xml:"PmtDt,omitempty"`
}

func (c *CorporateActionDate24) AddPostingDate() *DateAndDateTimeChoice {
	c.PostingDate = new(DateAndDateTimeChoice)
	return c.PostingDate
}

func (c *CorporateActionDate24) AddValueDate() *DateAndDateTimeChoice {
	c.ValueDate = new(DateAndDateTimeChoice)
	return c.ValueDate
}

func (c *CorporateActionDate24) AddForeignExchangeRateFixingDate() *DateAndDateTimeChoice {
	c.ForeignExchangeRateFixingDate = new(DateAndDateTimeChoice)
	return c.ForeignExchangeRateFixingDate
}

func (c *CorporateActionDate24) AddEarliestPaymentDate() *DateAndDateTimeChoice {
	c.EarliestPaymentDate = new(DateAndDateTimeChoice)
	return c.EarliestPaymentDate
}

func (c *CorporateActionDate24) AddPaymentDate() *DateAndDateTimeChoice {
	c.PaymentDate = new(DateAndDateTimeChoice)
	return c.PaymentDate
}

// Specifies corporate action dates.
type CorporateActionDate27 struct {

	// Date/time at which the issuer announced that a corporate action event will occur.
	AnnouncementDate *DateFormat19Choice `xml:"AnncmntDt,omitempty"`

	// Deadline by which the beneficial ownership of securities must be declared.
	CertificationDeadline *DateFormat19Choice `xml:"CertfctnDdln,omitempty"`

	// Date upon which the court provided approval.
	CourtApprovalDate *DateFormat19Choice `xml:"CrtApprvlDt,omitempty"`

	// First possible early closing date of an offer if different from the expiry date.
	EarlyClosingDate *DateFormat19Choice `xml:"EarlyClsgDt,omitempty"`

	// Date/time at which an event is officially effective from the issuer's perspective.
	EffectiveDate *DateFormat19Choice `xml:"FctvDt,omitempty"`

	// Date/Time on which all or part of any holding bought in a unit trust is subject to being treated as capital rather than income. This is normally one day after the previous distribution's ex date.
	EqualisationDate *DateFormat19Choice `xml:"EqulstnDt,omitempty"`

	// Date/time at which additional information on the event will be announced, for example, exchange ratio announcement date.
	FurtherDetailedAnnouncementDate *DateFormat19Choice `xml:"FrthrDtldAnncmntDt,omitempty"`

	// Date/time at which an index / rate / price / value will be determined.
	FixingDate *DateFormat19Choice `xml:"FxgDt,omitempty"`

	// Date/time on which the lottery is run and applied to the holder's positions. This is also applicable to partial calls.
	LotteryDate *DateFormat19Choice `xml:"LtryDt,omitempty"`

	// Date/time to which the maturity date of an interest bearing security is extended.
	NewMaturityDate *DateFormat19Choice `xml:"NewMtrtyDt,omitempty"`

	// Date/time on which the bondholder's or shareholder's meeting will take place.
	MeetingDate *DateFormat19Choice `xml:"MtgDt,omitempty"`

	// Date/time at which the margin rate will be determined .
	MarginFixingDate *DateFormat19Choice `xml:"MrgnFxgDt,omitempty"`

	// Date/time (and time) at which an issuer will determine the proration amount/quantity of an offer.
	ProrationDate *DateFormat19Choice `xml:"PrratnDt,omitempty"`

	// Date/time at which positions are struck at the end of the day to note which parties will receive the relevant amount of entitlement, due to be distributed on payment date.
	RecordDate *DateFormat19Choice `xml:"RcrdDt,omitempty"`

	// Date/time on which instructions to register or registration details will be accepted.
	RegistrationDeadline *DateFormat19Choice `xml:"RegnDdln,omitempty"`

	// Date/time on which results are published, for example, results of an offer.
	ResultsPublicationDate *DateFormat19Choice `xml:"RsltsPblctnDt,omitempty"`

	// Deadline by which instructions must be received to split securities, for example, of physical certificates.
	DeadlineToSplit *DateFormat19Choice `xml:"DdlnToSplt,omitempty"`

	// Date/time on until which tax breakdown instructions will be accepted.
	DeadlineForTaxBreakdownInstruction *DateFormat19Choice `xml:"DdlnForTaxBrkdwnInstr,omitempty"`

	// Date/time at which trading of a security is suspended as the result of an event.
	TradingSuspendedDate *DateFormat19Choice `xml:"TradgSspdDt,omitempty"`

	// Date/time upon which the terms of the take-over become unconditional as to acceptances.
	UnconditionalDate *DateFormat19Choice `xml:"UcondlDt,omitempty"`

	// Date/time at on which all conditions, including regulatory, legal etc. pertaining to the take-over, have been met.
	WhollyUnconditionalDate *DateFormat19Choice `xml:"WhlyUcondlDt,omitempty"`

	// Date/time as from which trading (including exchange and OTC trading) occurs on the underlying security without the benefit.
	ExDividendDate *DateFormat19Choice `xml:"ExDvddDt,omitempty"`

	// Date/time at which the corporate action is legally announced by an official body, for example, publication by a governmental administration.
	OfficialAnnouncementPublicationDate *DateFormat19Choice `xml:"OffclAnncmntPblctnDt,omitempty"`

	// Date/time as from which 'special processing' can start to be used by participants for that event. Special processing is a means of marking a transaction, that would normally be traded ex or cum, as being traded cum or ex respectively,  for example, a transaction dealt 'special' after the ex date would result in the buyer being eligible for the entitlement. This is typically used in the UK and Irish markets.
	SpecialExDate *DateFormat19Choice `xml:"SpclExDt,omitempty"`

	// Last date/time by which a buying counterparty to a trade can be sure that it will have the right to participate in an event.
	GuaranteedParticipationDate *DateFormat19Choice `xml:"GrntedPrtcptnDt,omitempty"`

	// Deadline by which an entitled holder needs to advise their counterparty to a transaction of their election for a corporate action event, also known as Buyer Protection Deadline.
	ElectionToCounterpartyDeadline *DateFormat19Choice `xml:"ElctnToCtrPtyDdln,omitempty"`

	// Date/time at which an event/offer is terminated or lapsed.
	LapsedDate *DateFormat19Choice `xml:"LpsdDt,omitempty"`

	// Date/time at which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat19Choice `xml:"PmtDt,omitempty"`

	// Date/Time by which the account owner must instruct directly another party, for example to provide documentation to an issuer agent.
	ThirdPartyDeadline *DateFormat19Choice `xml:"ThrdPtyDdln,omitempty"`

	// Date/Time set by the issuer agent as a first early deadline by which the account owner must instruct directly another party, possibly giving the holder eligibility to incentives. For example, to provide documentation to an issuer agent.
	EarlyThirdPartyDeadline *DateFormat19Choice `xml:"EarlyThrdPtyDdln,omitempty"`

	// Date by which the depository stops monitoring activities of the event, for instance, accounting and tracking activities for due bills end.
	MarketClaimTrackingEndDate *DateFormat19Choice `xml:"MktClmTrckgEndDt,omitempty"`

	// Last day an investor can become a lead plaintiff.
	LeadPlaintiffDeadline *DateFormat19Choice `xml:"LeadPlntffDdln,omitempty"`

	// Date on which the action was filed at the applicable court.
	FilingDate *DateFormat16Choice `xml:"FilgDt,omitempty"`

	// Date for the hearing between the plaintiff and defendant, as set by the court.
	HearingDate *DateFormat16Choice `xml:"HrgDt,omitempty"`
}

func (c *CorporateActionDate27) AddAnnouncementDate() *DateFormat19Choice {
	c.AnnouncementDate = new(DateFormat19Choice)
	return c.AnnouncementDate
}

func (c *CorporateActionDate27) AddCertificationDeadline() *DateFormat19Choice {
	c.CertificationDeadline = new(DateFormat19Choice)
	return c.CertificationDeadline
}

func (c *CorporateActionDate27) AddCourtApprovalDate() *DateFormat19Choice {
	c.CourtApprovalDate = new(DateFormat19Choice)
	return c.CourtApprovalDate
}

func (c *CorporateActionDate27) AddEarlyClosingDate() *DateFormat19Choice {
	c.EarlyClosingDate = new(DateFormat19Choice)
	return c.EarlyClosingDate
}

func (c *CorporateActionDate27) AddEffectiveDate() *DateFormat19Choice {
	c.EffectiveDate = new(DateFormat19Choice)
	return c.EffectiveDate
}

func (c *CorporateActionDate27) AddEqualisationDate() *DateFormat19Choice {
	c.EqualisationDate = new(DateFormat19Choice)
	return c.EqualisationDate
}

func (c *CorporateActionDate27) AddFurtherDetailedAnnouncementDate() *DateFormat19Choice {
	c.FurtherDetailedAnnouncementDate = new(DateFormat19Choice)
	return c.FurtherDetailedAnnouncementDate
}

func (c *CorporateActionDate27) AddFixingDate() *DateFormat19Choice {
	c.FixingDate = new(DateFormat19Choice)
	return c.FixingDate
}

func (c *CorporateActionDate27) AddLotteryDate() *DateFormat19Choice {
	c.LotteryDate = new(DateFormat19Choice)
	return c.LotteryDate
}

func (c *CorporateActionDate27) AddNewMaturityDate() *DateFormat19Choice {
	c.NewMaturityDate = new(DateFormat19Choice)
	return c.NewMaturityDate
}

func (c *CorporateActionDate27) AddMeetingDate() *DateFormat19Choice {
	c.MeetingDate = new(DateFormat19Choice)
	return c.MeetingDate
}

func (c *CorporateActionDate27) AddMarginFixingDate() *DateFormat19Choice {
	c.MarginFixingDate = new(DateFormat19Choice)
	return c.MarginFixingDate
}

func (c *CorporateActionDate27) AddProrationDate() *DateFormat19Choice {
	c.ProrationDate = new(DateFormat19Choice)
	return c.ProrationDate
}

func (c *CorporateActionDate27) AddRecordDate() *DateFormat19Choice {
	c.RecordDate = new(DateFormat19Choice)
	return c.RecordDate
}

func (c *CorporateActionDate27) AddRegistrationDeadline() *DateFormat19Choice {
	c.RegistrationDeadline = new(DateFormat19Choice)
	return c.RegistrationDeadline
}

func (c *CorporateActionDate27) AddResultsPublicationDate() *DateFormat19Choice {
	c.ResultsPublicationDate = new(DateFormat19Choice)
	return c.ResultsPublicationDate
}

func (c *CorporateActionDate27) AddDeadlineToSplit() *DateFormat19Choice {
	c.DeadlineToSplit = new(DateFormat19Choice)
	return c.DeadlineToSplit
}

func (c *CorporateActionDate27) AddDeadlineForTaxBreakdownInstruction() *DateFormat19Choice {
	c.DeadlineForTaxBreakdownInstruction = new(DateFormat19Choice)
	return c.DeadlineForTaxBreakdownInstruction
}

func (c *CorporateActionDate27) AddTradingSuspendedDate() *DateFormat19Choice {
	c.TradingSuspendedDate = new(DateFormat19Choice)
	return c.TradingSuspendedDate
}

func (c *CorporateActionDate27) AddUnconditionalDate() *DateFormat19Choice {
	c.UnconditionalDate = new(DateFormat19Choice)
	return c.UnconditionalDate
}

func (c *CorporateActionDate27) AddWhollyUnconditionalDate() *DateFormat19Choice {
	c.WhollyUnconditionalDate = new(DateFormat19Choice)
	return c.WhollyUnconditionalDate
}

func (c *CorporateActionDate27) AddExDividendDate() *DateFormat19Choice {
	c.ExDividendDate = new(DateFormat19Choice)
	return c.ExDividendDate
}

func (c *CorporateActionDate27) AddOfficialAnnouncementPublicationDate() *DateFormat19Choice {
	c.OfficialAnnouncementPublicationDate = new(DateFormat19Choice)
	return c.OfficialAnnouncementPublicationDate
}

func (c *CorporateActionDate27) AddSpecialExDate() *DateFormat19Choice {
	c.SpecialExDate = new(DateFormat19Choice)
	return c.SpecialExDate
}

func (c *CorporateActionDate27) AddGuaranteedParticipationDate() *DateFormat19Choice {
	c.GuaranteedParticipationDate = new(DateFormat19Choice)
	return c.GuaranteedParticipationDate
}

func (c *CorporateActionDate27) AddElectionToCounterpartyDeadline() *DateFormat19Choice {
	c.ElectionToCounterpartyDeadline = new(DateFormat19Choice)
	return c.ElectionToCounterpartyDeadline
}

func (c *CorporateActionDate27) AddLapsedDate() *DateFormat19Choice {
	c.LapsedDate = new(DateFormat19Choice)
	return c.LapsedDate
}

func (c *CorporateActionDate27) AddPaymentDate() *DateFormat19Choice {
	c.PaymentDate = new(DateFormat19Choice)
	return c.PaymentDate
}

func (c *CorporateActionDate27) AddThirdPartyDeadline() *DateFormat19Choice {
	c.ThirdPartyDeadline = new(DateFormat19Choice)
	return c.ThirdPartyDeadline
}

func (c *CorporateActionDate27) AddEarlyThirdPartyDeadline() *DateFormat19Choice {
	c.EarlyThirdPartyDeadline = new(DateFormat19Choice)
	return c.EarlyThirdPartyDeadline
}

func (c *CorporateActionDate27) AddMarketClaimTrackingEndDate() *DateFormat19Choice {
	c.MarketClaimTrackingEndDate = new(DateFormat19Choice)
	return c.MarketClaimTrackingEndDate
}

func (c *CorporateActionDate27) AddLeadPlaintiffDeadline() *DateFormat19Choice {
	c.LeadPlaintiffDeadline = new(DateFormat19Choice)
	return c.LeadPlaintiffDeadline
}

func (c *CorporateActionDate27) AddFilingDate() *DateFormat16Choice {
	c.FilingDate = new(DateFormat16Choice)
	return c.FilingDate
}

func (c *CorporateActionDate27) AddHearingDate() *DateFormat16Choice {
	c.HearingDate = new(DateFormat16Choice)
	return c.HearingDate
}

// Specifies corporate action dates.
type CorporateActionDate29 struct {

	// Date/time that the account servicer has set as the deadline to respond, with instructions, to an outstanding event, giving the holder eligibility to incentives. This time is dependent on the reference time zone of the account servicer as specified in a Service Level Agreement (SLA).
	EarlyResponseDeadline *DateFormat19Choice `xml:"EarlyRspnDdln,omitempty"`

	// Last day a holder can deliver the securities that it had elected on and/or previously protected.
	CoverExpirationDate *DateFormat19Choice `xml:"CoverXprtnDt,omitempty"`

	// Last date/time a holder can request to defer delivery of securities pursuant to a notice of guaranteed delivery or other required documentation.
	ProtectDate *DateFormat19Choice `xml:"PrtctDt,omitempty"`

	// Issuer or issuer's agent deadline to respond, with an instruction, to an outstanding offer or privilege.
	MarketDeadline *DateFormat19Choice `xml:"MktDdln,omitempty"`

	// Date/time at which the account servicer has set as the deadline to respond, with instructions, to an outstanding event. This time is dependent on the reference time zone of the account servicer as specified in a Service Level Agreement (SLA).
	ResponseDeadline *DateFormat20Choice `xml:"RspnDdln,omitempty"`

	// Date/time at which an order expires or on which a privilege or offer terminates.
	ExpiryDate *DateFormat19Choice `xml:"XpryDt,omitempty"`

	// Date/time by which cash must be in place in order to take part in the event.
	SubscriptionCostDebitDate *DateFormat19Choice `xml:"SbcptCostDbtDt,omitempty"`

	// Last day that a participant of the depository can deliver securities that it had elected on and/or previously protected.
	DepositoryCoverExpirationDate *DateFormat19Choice `xml:"DpstryCoverXprtnDt,omitempty"`

	// Date/time that the account servicer has set as the deadline to respond, with instructions, to an outstanding event, for which the underlying security is out on loan. This time is dependent on the reference time zone of the account servicer as specified in an SLA.
	StockLendingDeadline *DateFormat19Choice `xml:"StockLndgDdln,omitempty"`
}

func (c *CorporateActionDate29) AddEarlyResponseDeadline() *DateFormat19Choice {
	c.EarlyResponseDeadline = new(DateFormat19Choice)
	return c.EarlyResponseDeadline
}

func (c *CorporateActionDate29) AddCoverExpirationDate() *DateFormat19Choice {
	c.CoverExpirationDate = new(DateFormat19Choice)
	return c.CoverExpirationDate
}

func (c *CorporateActionDate29) AddProtectDate() *DateFormat19Choice {
	c.ProtectDate = new(DateFormat19Choice)
	return c.ProtectDate
}

func (c *CorporateActionDate29) AddMarketDeadline() *DateFormat19Choice {
	c.MarketDeadline = new(DateFormat19Choice)
	return c.MarketDeadline
}

func (c *CorporateActionDate29) AddResponseDeadline() *DateFormat20Choice {
	c.ResponseDeadline = new(DateFormat20Choice)
	return c.ResponseDeadline
}

func (c *CorporateActionDate29) AddExpiryDate() *DateFormat19Choice {
	c.ExpiryDate = new(DateFormat19Choice)
	return c.ExpiryDate
}

func (c *CorporateActionDate29) AddSubscriptionCostDebitDate() *DateFormat19Choice {
	c.SubscriptionCostDebitDate = new(DateFormat19Choice)
	return c.SubscriptionCostDebitDate
}

func (c *CorporateActionDate29) AddDepositoryCoverExpirationDate() *DateFormat19Choice {
	c.DepositoryCoverExpirationDate = new(DateFormat19Choice)
	return c.DepositoryCoverExpirationDate
}

func (c *CorporateActionDate29) AddStockLendingDeadline() *DateFormat19Choice {
	c.StockLendingDeadline = new(DateFormat19Choice)
	return c.StockLendingDeadline
}

// Specifies coprorate action dates.
type CorporateActionDate3 struct {

	// Date/time at which the distribution is due to take place (cash and/or securities).
	PaymentDate *DateFormat4Choice `xml:"PmtDt,omitempty"`

	// Date/time at which securities become available for sale.
	AvailableDate *DateFormat4Choice `xml:"AvlblDt,omitempty"`

	// Date/time at which a security will be entitled to a dividend.
	DividendRankingDate *DateFormat4Choice `xml:"DvddRnkgDt,omitempty"`

	// Date on which security will assimilate, become fungible, or have the same rights to dividends as the parent issue.
	PariPassuDate *DateFormat4Choice `xml:"PrpssDt,omitempty"`

	// Date/time at which new securities begin trading.
	FirstDealingDate *DateFormat4Choice `xml:"FrstDealgDt,omitempty"`

	// Date/time at which a payment can be made, eg, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat4Choice `xml:"EarlstPmtDt,omitempty"`
}

func (c *CorporateActionDate3) AddPaymentDate() *DateFormat4Choice {
	c.PaymentDate = new(DateFormat4Choice)
	return c.PaymentDate
}

func (c *CorporateActionDate3) AddAvailableDate() *DateFormat4Choice {
	c.AvailableDate = new(DateFormat4Choice)
	return c.AvailableDate
}

func (c *CorporateActionDate3) AddDividendRankingDate() *DateFormat4Choice {
	c.DividendRankingDate = new(DateFormat4Choice)
	return c.DividendRankingDate
}

func (c *CorporateActionDate3) AddPariPassuDate() *DateFormat4Choice {
	c.PariPassuDate = new(DateFormat4Choice)
	return c.PariPassuDate
}

func (c *CorporateActionDate3) AddFirstDealingDate() *DateFormat4Choice {
	c.FirstDealingDate = new(DateFormat4Choice)
	return c.FirstDealingDate
}

func (c *CorporateActionDate3) AddEarliestPaymentDate() *DateFormat4Choice {
	c.EarliestPaymentDate = new(DateFormat4Choice)
	return c.EarliestPaymentDate
}

// Specifies corporate action dates.
type CorporateActionDate30 struct {

	// Date/time at which positions are struck at the end of the day to note which parties will receive the relevant amount of entitlement, due to be distributed on payment date.
	RecordDate *DateFormat19Choice `xml:"RcrdDt,omitempty"`

	// Date/time as from which trading (including exchange and OTC trading) occurs on the underlying security without the benefit.
	ExDividendDate *DateFormat19Choice `xml:"ExDvddDt,omitempty"`
}

func (c *CorporateActionDate30) AddRecordDate() *DateFormat19Choice {
	c.RecordDate = new(DateFormat19Choice)
	return c.RecordDate
}

func (c *CorporateActionDate30) AddExDividendDate() *DateFormat19Choice {
	c.ExDividendDate = new(DateFormat19Choice)
	return c.ExDividendDate
}

// Specifies corporate action dates.
type CorporateActionDate4 struct {

	// Date/time at which the coupons are to be/were submitted for payment of interest.
	CouponClippingDate *DateFormat4Choice `xml:"CpnClpngDt,omitempty"`

	// Last date/time at which a holder can consent to the changes sought by the corporation.
	ConsentExpirationDate *DateFormat4Choice `xml:"CnsntXprtnDt,omitempty"`

	// Date/time used by the offeror to determine the beneficiary eligible to participate in a consent based on the registered owner of the securities, eg, beneficial owner of consent record. The consent record date qualifier is used to indicate that a record date only applies to a certain part of the offer, not the entire offer.
	ConsentRecordDate *DateFormat4Choice `xml:"CnsntRcrdDt,omitempty"`

	// Date/time at which the distribution is due to take place (cash and/or securities).
	PaymentDate *DateFormat4Choice `xml:"PmtDt,omitempty"`

	// Date/time at which a payment can be made, eg, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat4Choice `xml:"EarlstPmtDt,omitempty"`

	// Issuer or issuer's agent deadline to respond, with an instruction, to an outstanding offer or privilege.
	MarketDeadline *DateFormat4Choice `xml:"MktDdln,omitempty"`

	// Date/time at which the account servicer has set as the deadline to respond, with instructions, to an outstanding event. This time is dependent on the reference time zone of the account servicer as specified in an SLA.
	ResponseDeadline *DateFormat4Choice `xml:"RspnDdln,omitempty"`

	// Deadline by which instructions must be received to split securities, eg, of physical certificates.
	DeadlineToSplit *DateFormat4Choice `xml:"DdlnToSplt,omitempty"`

	// Date/time at which an order expires or on which a privilege or offer terminates.
	ExpiryDate *DateFormat4Choice `xml:"XpryDt,omitempty"`

	// Date/time at which the price of a security is determined.
	QuotationSettingDate *DateFormat4Choice `xml:"QtnSetngDt,omitempty"`

	// Date/time by which cash must be in place in order to take part in the event.
	SubscriptionCostDebitDate *DateFormat4Choice `xml:"SbcptCostDbtDt,omitempty"`
}

func (c *CorporateActionDate4) AddCouponClippingDate() *DateFormat4Choice {
	c.CouponClippingDate = new(DateFormat4Choice)
	return c.CouponClippingDate
}

func (c *CorporateActionDate4) AddConsentExpirationDate() *DateFormat4Choice {
	c.ConsentExpirationDate = new(DateFormat4Choice)
	return c.ConsentExpirationDate
}

func (c *CorporateActionDate4) AddConsentRecordDate() *DateFormat4Choice {
	c.ConsentRecordDate = new(DateFormat4Choice)
	return c.ConsentRecordDate
}

func (c *CorporateActionDate4) AddPaymentDate() *DateFormat4Choice {
	c.PaymentDate = new(DateFormat4Choice)
	return c.PaymentDate
}

func (c *CorporateActionDate4) AddEarliestPaymentDate() *DateFormat4Choice {
	c.EarliestPaymentDate = new(DateFormat4Choice)
	return c.EarliestPaymentDate
}

func (c *CorporateActionDate4) AddMarketDeadline() *DateFormat4Choice {
	c.MarketDeadline = new(DateFormat4Choice)
	return c.MarketDeadline
}

func (c *CorporateActionDate4) AddResponseDeadline() *DateFormat4Choice {
	c.ResponseDeadline = new(DateFormat4Choice)
	return c.ResponseDeadline
}

func (c *CorporateActionDate4) AddDeadlineToSplit() *DateFormat4Choice {
	c.DeadlineToSplit = new(DateFormat4Choice)
	return c.DeadlineToSplit
}

func (c *CorporateActionDate4) AddExpiryDate() *DateFormat4Choice {
	c.ExpiryDate = new(DateFormat4Choice)
	return c.ExpiryDate
}

func (c *CorporateActionDate4) AddQuotationSettingDate() *DateFormat4Choice {
	c.QuotationSettingDate = new(DateFormat4Choice)
	return c.QuotationSettingDate
}

func (c *CorporateActionDate4) AddSubscriptionCostDebitDate() *DateFormat4Choice {
	c.SubscriptionCostDebitDate = new(DateFormat4Choice)
	return c.SubscriptionCostDebitDate
}

// Specifies corporate action dates.
type CorporateActionDate41 struct {

	// Date/time at which positions are struck at the end of the day to note which parties will receive the relevant amount of entitlement, due to be distributed on payment date.
	RecordDate *DateFormat19Choice `xml:"RcrdDt,omitempty"`

	// Date/time as from which trading (including exchange and OTC trading) occurs on the underlying security without the benefit.
	ExDividendDate *DateFormat19Choice `xml:"ExDvddDt,omitempty"`

	// Date/time on which the lottery is run and applied to the holder's positions. This is also applicable to partial calls.
	LotteryDate *DateFormat19Choice `xml:"LtryDt,omitempty"`
}

func (c *CorporateActionDate41) AddRecordDate() *DateFormat19Choice {
	c.RecordDate = new(DateFormat19Choice)
	return c.RecordDate
}

func (c *CorporateActionDate41) AddExDividendDate() *DateFormat19Choice {
	c.ExDividendDate = new(DateFormat19Choice)
	return c.ExDividendDate
}

func (c *CorporateActionDate41) AddLotteryDate() *DateFormat19Choice {
	c.LotteryDate = new(DateFormat19Choice)
	return c.LotteryDate
}

// Specifies corporate action dates.
type CorporateActionDate44 struct {

	// Date/time at which the issuer announced that a corporate action event will occur.
	AnnouncementDate *DateFormat31Choice `xml:"AnncmntDt,omitempty"`

	// Deadline by which the beneficial ownership of securities must be declared.
	CertificationDeadline *DateFormat31Choice `xml:"CertfctnDdln,omitempty"`

	// Date upon which the court provided approval.
	CourtApprovalDate *DateFormat31Choice `xml:"CrtApprvlDt,omitempty"`

	// First possible early closing date of an offer if different from the expiry date.
	EarlyClosingDate *DateFormat31Choice `xml:"EarlyClsgDt,omitempty"`

	// Date/time at which an event is officially effective from the issuer's perspective.
	EffectiveDate *DateFormat31Choice `xml:"FctvDt,omitempty"`

	// Date/Time on which all or part of any holding bought in a unit trust is subject to being treated as capital rather than income. This is normally one day after the previous distribution's ex date.
	EqualisationDate *DateFormat31Choice `xml:"EqulstnDt,omitempty"`

	// Date/time at which additional information on the event will be announced, for example, exchange ratio announcement date.
	FurtherDetailedAnnouncementDate *DateFormat31Choice `xml:"FrthrDtldAnncmntDt,omitempty"`

	// Date/time at which an index / rate / price / value will be determined.
	FixingDate *DateFormat31Choice `xml:"FxgDt,omitempty"`

	// Date/time on which the lottery is run and applied to the holder's positions. This is also applicable to partial calls.
	LotteryDate *DateFormat31Choice `xml:"LtryDt,omitempty"`

	// Date/time to which the maturity date of an interest bearing security is extended.
	NewMaturityDate *DateFormat31Choice `xml:"NewMtrtyDt,omitempty"`

	// Date/time on which the bondholder's or shareholder's meeting will take place.
	MeetingDate *DateFormat31Choice `xml:"MtgDt,omitempty"`

	// Date/time at which the margin rate will be determined .
	MarginFixingDate *DateFormat31Choice `xml:"MrgnFxgDt,omitempty"`

	// Date/time (and time) at which an issuer will determine the proration amount/quantity of an offer.
	ProrationDate *DateFormat31Choice `xml:"PrratnDt,omitempty"`

	// Date/time at which positions are struck at the end of the day to note which parties will receive the relevant amount of entitlement, due to be distributed on payment date.
	RecordDate *DateFormat31Choice `xml:"RcrdDt,omitempty"`

	// Date/time on which instructions to register or registration details will be accepted.
	RegistrationDeadline *DateFormat31Choice `xml:"RegnDdln,omitempty"`

	// Date/time on which results are published, for example, results of an offer.
	ResultsPublicationDate *DateFormat31Choice `xml:"RsltsPblctnDt,omitempty"`

	// Deadline by which instructions must be received to split securities, for example, of physical certificates.
	DeadlineToSplit *DateFormat31Choice `xml:"DdlnToSplt,omitempty"`

	// Date/time on until which tax breakdown instructions will be accepted.
	DeadlineForTaxBreakdownInstruction *DateFormat31Choice `xml:"DdlnForTaxBrkdwnInstr,omitempty"`

	// Date/time at which trading of a security is suspended as the result of an event.
	TradingSuspendedDate *DateFormat31Choice `xml:"TradgSspdDt,omitempty"`

	// Date/time upon which the terms of the take-over become unconditional as to acceptances.
	UnconditionalDate *DateFormat31Choice `xml:"UcondlDt,omitempty"`

	// Date/time at on which all conditions, including regulatory, legal etc. pertaining to the take-over, have been met.
	WhollyUnconditionalDate *DateFormat31Choice `xml:"WhlyUcondlDt,omitempty"`

	// Date/time as from which trading (including exchange and OTC trading) occurs on the underlying security without the benefit.
	ExDividendDate *DateFormat31Choice `xml:"ExDvddDt,omitempty"`

	// Date/time at which the corporate action is legally announced by an official body, for example, publication by a governmental administration.
	OfficialAnnouncementPublicationDate *DateFormat31Choice `xml:"OffclAnncmntPblctnDt,omitempty"`

	// Date/time as from which 'special processing' can start to be used by participants for that event. Special processing is a means of marking a transaction, that would normally be traded ex or cum, as being traded cum or ex respectively,  for example, a transaction dealt 'special' after the ex date would result in the buyer being eligible for the entitlement. This is typically used in the UK and Irish markets.
	SpecialExDate *DateFormat31Choice `xml:"SpclExDt,omitempty"`

	// Last date/time by which a buying counterparty to a trade can be sure that it will have the right to participate in an event.
	GuaranteedParticipationDate *DateFormat31Choice `xml:"GrntedPrtcptnDt,omitempty"`

	// Deadline by which an entitled holder needs to advise their counterparty to a transaction of their election for a corporate action event, also known as Buyer Protection Deadline.
	ElectionToCounterpartyMarketDeadline *DateFormat31Choice `xml:"ElctnToCtrPtyMktDdln,omitempty"`

	// Date/time the account servicer has set as the deadline to respond, with instructions, prior to the election to counterparty market deadline
	ElectionToCounterpartyResponseDeadline *DateFormat31Choice `xml:"ElctnToCtrPtyRspnDdln,omitempty"`

	// Date/time at which an event/offer is terminated or lapsed.
	LapsedDate *DateFormat31Choice `xml:"LpsdDt,omitempty"`

	// Date/time at which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat31Choice `xml:"PmtDt,omitempty"`

	// Date/Time by which the account owner must instruct directly another party, for example to provide documentation to an issuer agent.
	ThirdPartyDeadline *DateFormat31Choice `xml:"ThrdPtyDdln,omitempty"`

	// Date/Time set by the issuer agent as a first early deadline by which the account owner must instruct directly another party, possibly giving the holder eligibility to incentives. For example, to provide documentation to an issuer agent.
	EarlyThirdPartyDeadline *DateFormat31Choice `xml:"EarlyThrdPtyDdln,omitempty"`

	// Date by which the depository stops monitoring activities of the event, for instance, accounting and tracking activities for due bills end.
	MarketClaimTrackingEndDate *DateFormat31Choice `xml:"MktClmTrckgEndDt,omitempty"`

	// Last day an investor can become a lead plaintiff.
	LeadPlaintiffDeadline *DateFormat31Choice `xml:"LeadPlntffDdln,omitempty"`

	// Date on which the action was filed at the applicable court.
	FilingDate *DateFormat30Choice `xml:"FilgDt,omitempty"`

	// Date for the hearing between the plaintiff and defendant, as set by the court.
	HearingDate *DateFormat30Choice `xml:"HrgDt,omitempty"`
}

func (c *CorporateActionDate44) AddAnnouncementDate() *DateFormat31Choice {
	c.AnnouncementDate = new(DateFormat31Choice)
	return c.AnnouncementDate
}

func (c *CorporateActionDate44) AddCertificationDeadline() *DateFormat31Choice {
	c.CertificationDeadline = new(DateFormat31Choice)
	return c.CertificationDeadline
}

func (c *CorporateActionDate44) AddCourtApprovalDate() *DateFormat31Choice {
	c.CourtApprovalDate = new(DateFormat31Choice)
	return c.CourtApprovalDate
}

func (c *CorporateActionDate44) AddEarlyClosingDate() *DateFormat31Choice {
	c.EarlyClosingDate = new(DateFormat31Choice)
	return c.EarlyClosingDate
}

func (c *CorporateActionDate44) AddEffectiveDate() *DateFormat31Choice {
	c.EffectiveDate = new(DateFormat31Choice)
	return c.EffectiveDate
}

func (c *CorporateActionDate44) AddEqualisationDate() *DateFormat31Choice {
	c.EqualisationDate = new(DateFormat31Choice)
	return c.EqualisationDate
}

func (c *CorporateActionDate44) AddFurtherDetailedAnnouncementDate() *DateFormat31Choice {
	c.FurtherDetailedAnnouncementDate = new(DateFormat31Choice)
	return c.FurtherDetailedAnnouncementDate
}

func (c *CorporateActionDate44) AddFixingDate() *DateFormat31Choice {
	c.FixingDate = new(DateFormat31Choice)
	return c.FixingDate
}

func (c *CorporateActionDate44) AddLotteryDate() *DateFormat31Choice {
	c.LotteryDate = new(DateFormat31Choice)
	return c.LotteryDate
}

func (c *CorporateActionDate44) AddNewMaturityDate() *DateFormat31Choice {
	c.NewMaturityDate = new(DateFormat31Choice)
	return c.NewMaturityDate
}

func (c *CorporateActionDate44) AddMeetingDate() *DateFormat31Choice {
	c.MeetingDate = new(DateFormat31Choice)
	return c.MeetingDate
}

func (c *CorporateActionDate44) AddMarginFixingDate() *DateFormat31Choice {
	c.MarginFixingDate = new(DateFormat31Choice)
	return c.MarginFixingDate
}

func (c *CorporateActionDate44) AddProrationDate() *DateFormat31Choice {
	c.ProrationDate = new(DateFormat31Choice)
	return c.ProrationDate
}

func (c *CorporateActionDate44) AddRecordDate() *DateFormat31Choice {
	c.RecordDate = new(DateFormat31Choice)
	return c.RecordDate
}

func (c *CorporateActionDate44) AddRegistrationDeadline() *DateFormat31Choice {
	c.RegistrationDeadline = new(DateFormat31Choice)
	return c.RegistrationDeadline
}

func (c *CorporateActionDate44) AddResultsPublicationDate() *DateFormat31Choice {
	c.ResultsPublicationDate = new(DateFormat31Choice)
	return c.ResultsPublicationDate
}

func (c *CorporateActionDate44) AddDeadlineToSplit() *DateFormat31Choice {
	c.DeadlineToSplit = new(DateFormat31Choice)
	return c.DeadlineToSplit
}

func (c *CorporateActionDate44) AddDeadlineForTaxBreakdownInstruction() *DateFormat31Choice {
	c.DeadlineForTaxBreakdownInstruction = new(DateFormat31Choice)
	return c.DeadlineForTaxBreakdownInstruction
}

func (c *CorporateActionDate44) AddTradingSuspendedDate() *DateFormat31Choice {
	c.TradingSuspendedDate = new(DateFormat31Choice)
	return c.TradingSuspendedDate
}

func (c *CorporateActionDate44) AddUnconditionalDate() *DateFormat31Choice {
	c.UnconditionalDate = new(DateFormat31Choice)
	return c.UnconditionalDate
}

func (c *CorporateActionDate44) AddWhollyUnconditionalDate() *DateFormat31Choice {
	c.WhollyUnconditionalDate = new(DateFormat31Choice)
	return c.WhollyUnconditionalDate
}

func (c *CorporateActionDate44) AddExDividendDate() *DateFormat31Choice {
	c.ExDividendDate = new(DateFormat31Choice)
	return c.ExDividendDate
}

func (c *CorporateActionDate44) AddOfficialAnnouncementPublicationDate() *DateFormat31Choice {
	c.OfficialAnnouncementPublicationDate = new(DateFormat31Choice)
	return c.OfficialAnnouncementPublicationDate
}

func (c *CorporateActionDate44) AddSpecialExDate() *DateFormat31Choice {
	c.SpecialExDate = new(DateFormat31Choice)
	return c.SpecialExDate
}

func (c *CorporateActionDate44) AddGuaranteedParticipationDate() *DateFormat31Choice {
	c.GuaranteedParticipationDate = new(DateFormat31Choice)
	return c.GuaranteedParticipationDate
}

func (c *CorporateActionDate44) AddElectionToCounterpartyMarketDeadline() *DateFormat31Choice {
	c.ElectionToCounterpartyMarketDeadline = new(DateFormat31Choice)
	return c.ElectionToCounterpartyMarketDeadline
}

func (c *CorporateActionDate44) AddElectionToCounterpartyResponseDeadline() *DateFormat31Choice {
	c.ElectionToCounterpartyResponseDeadline = new(DateFormat31Choice)
	return c.ElectionToCounterpartyResponseDeadline
}

func (c *CorporateActionDate44) AddLapsedDate() *DateFormat31Choice {
	c.LapsedDate = new(DateFormat31Choice)
	return c.LapsedDate
}

func (c *CorporateActionDate44) AddPaymentDate() *DateFormat31Choice {
	c.PaymentDate = new(DateFormat31Choice)
	return c.PaymentDate
}

func (c *CorporateActionDate44) AddThirdPartyDeadline() *DateFormat31Choice {
	c.ThirdPartyDeadline = new(DateFormat31Choice)
	return c.ThirdPartyDeadline
}

func (c *CorporateActionDate44) AddEarlyThirdPartyDeadline() *DateFormat31Choice {
	c.EarlyThirdPartyDeadline = new(DateFormat31Choice)
	return c.EarlyThirdPartyDeadline
}

func (c *CorporateActionDate44) AddMarketClaimTrackingEndDate() *DateFormat31Choice {
	c.MarketClaimTrackingEndDate = new(DateFormat31Choice)
	return c.MarketClaimTrackingEndDate
}

func (c *CorporateActionDate44) AddLeadPlaintiffDeadline() *DateFormat31Choice {
	c.LeadPlaintiffDeadline = new(DateFormat31Choice)
	return c.LeadPlaintiffDeadline
}

func (c *CorporateActionDate44) AddFilingDate() *DateFormat30Choice {
	c.FilingDate = new(DateFormat30Choice)
	return c.FilingDate
}

func (c *CorporateActionDate44) AddHearingDate() *DateFormat30Choice {
	c.HearingDate = new(DateFormat30Choice)
	return c.HearingDate
}

// Specifies corporate action dates.
type CorporateActionDate45 struct {

	// Date/time at which positions are struck at the end of the day to note which parties will receive the relevant amount of entitlement, due to be distributed on payment date.
	RecordDate *DateFormat31Choice `xml:"RcrdDt,omitempty"`

	// Date/time as from which trading (including exchange and OTC trading) occurs on the underlying security without the benefit.
	ExDividendDate *DateFormat31Choice `xml:"ExDvddDt,omitempty"`

	// Date/time on which the lottery is run and applied to the holder's positions. This is also applicable to partial calls.
	LotteryDate *DateFormat31Choice `xml:"LtryDt,omitempty"`
}

func (c *CorporateActionDate45) AddRecordDate() *DateFormat31Choice {
	c.RecordDate = new(DateFormat31Choice)
	return c.RecordDate
}

func (c *CorporateActionDate45) AddExDividendDate() *DateFormat31Choice {
	c.ExDividendDate = new(DateFormat31Choice)
	return c.ExDividendDate
}

func (c *CorporateActionDate45) AddLotteryDate() *DateFormat31Choice {
	c.LotteryDate = new(DateFormat31Choice)
	return c.LotteryDate
}

// Specifies corporate action date.
type CorporateActionDate46 struct {

	// Date/time at which the account servicer has set as the deadline to respond, with instructions, to an outstanding event. This time is dependent on the reference time zone of the account servicer as specified in a Service Level Agreement (SLA).
	ResponseDeadline *DateFormat31Choice `xml:"RspnDdln,omitempty"`

	// Date/time by which cash must be in place in order to take part in the event.
	SubscriptionCostDebitDate *DateFormat31Choice `xml:"SbcptCostDbtDt,omitempty"`

	// Issuer or issuer's agent deadline to respond, with an instruction, to an outstanding offer or privilege.
	MarketDeadline *DateFormat31Choice `xml:"MktDdln,omitempty"`

	// Date/time at which an order expires or on which a privilege or offer terminates.
	ExpiryDate *DateFormat31Choice `xml:"XpryDt,omitempty"`

	// Last day a holder can deliver the securities that it had elected on and/or previously protected.
	CoverExpirationDate *DateFormat31Choice `xml:"CoverXprtnDt,omitempty"`

	// Last date/time a holder can request to defer delivery of securities pursuant to a notice of guaranteed delivery or other required documentation.
	ProtectDate *DateFormat31Choice `xml:"PrtctDt,omitempty"`

	// Date/time at which the deal (rights) was agreed.
	TradingDate *DateFormat31Choice `xml:"TradgDt,omitempty"`
}

func (c *CorporateActionDate46) AddResponseDeadline() *DateFormat31Choice {
	c.ResponseDeadline = new(DateFormat31Choice)
	return c.ResponseDeadline
}

func (c *CorporateActionDate46) AddSubscriptionCostDebitDate() *DateFormat31Choice {
	c.SubscriptionCostDebitDate = new(DateFormat31Choice)
	return c.SubscriptionCostDebitDate
}

func (c *CorporateActionDate46) AddMarketDeadline() *DateFormat31Choice {
	c.MarketDeadline = new(DateFormat31Choice)
	return c.MarketDeadline
}

func (c *CorporateActionDate46) AddExpiryDate() *DateFormat31Choice {
	c.ExpiryDate = new(DateFormat31Choice)
	return c.ExpiryDate
}

func (c *CorporateActionDate46) AddCoverExpirationDate() *DateFormat31Choice {
	c.CoverExpirationDate = new(DateFormat31Choice)
	return c.CoverExpirationDate
}

func (c *CorporateActionDate46) AddProtectDate() *DateFormat31Choice {
	c.ProtectDate = new(DateFormat31Choice)
	return c.ProtectDate
}

func (c *CorporateActionDate46) AddTradingDate() *DateFormat31Choice {
	c.TradingDate = new(DateFormat31Choice)
	return c.TradingDate
}

// Specifies corporate action dates.
type CorporateActionDate47 struct {

	// Date on which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat31Choice `xml:"PmtDt"`

	// Date/time when calculating economic benefit for a cash amount.
	ValueDate *DateFormat33Choice `xml:"ValDt,omitempty"`

	// Date/time at which a foreign exchange rate will be determined.
	ForeignExchangeRateFixingDate *DateFormat31Choice `xml:"FXRateFxgDt,omitempty"`

	// Date on which a payment can be made, for example, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat31Choice `xml:"EarlstPmtDt,omitempty"`
}

func (c *CorporateActionDate47) AddPaymentDate() *DateFormat31Choice {
	c.PaymentDate = new(DateFormat31Choice)
	return c.PaymentDate
}

func (c *CorporateActionDate47) AddValueDate() *DateFormat33Choice {
	c.ValueDate = new(DateFormat33Choice)
	return c.ValueDate
}

func (c *CorporateActionDate47) AddForeignExchangeRateFixingDate() *DateFormat31Choice {
	c.ForeignExchangeRateFixingDate = new(DateFormat31Choice)
	return c.ForeignExchangeRateFixingDate
}

func (c *CorporateActionDate47) AddEarliestPaymentDate() *DateFormat31Choice {
	c.EarliestPaymentDate = new(DateFormat31Choice)
	return c.EarliestPaymentDate
}

// Specifies corporate action dates.
type CorporateActionDate48 struct {

	// Date/time that the account servicer has set as the deadline to respond, with instructions, to an outstanding event, giving the holder eligibility to incentives. This time is dependent on the reference time zone of the account servicer as specified in a Service Level Agreement (SLA).
	EarlyResponseDeadline *DateFormat31Choice `xml:"EarlyRspnDdln,omitempty"`

	// Last day a holder can deliver the securities that it had elected on and/or previously protected.
	CoverExpirationDate *DateFormat31Choice `xml:"CoverXprtnDt,omitempty"`

	// Last date/time a holder can request to defer delivery of securities pursuant to a notice of guaranteed delivery or other required documentation.
	ProtectDate *DateFormat31Choice `xml:"PrtctDt,omitempty"`

	// Issuer or issuer's agent deadline to respond, with an instruction, to an outstanding offer or privilege.
	MarketDeadline *DateFormat31Choice `xml:"MktDdln,omitempty"`

	// Date/time at which the account servicer has set as the deadline to respond, with instructions, to an outstanding event. This time is dependent on the reference time zone of the account servicer as specified in a Service Level Agreement (SLA).
	ResponseDeadline *DateFormat32Choice `xml:"RspnDdln,omitempty"`

	// Date/time at which an order expires or on which a privilege or offer terminates.
	ExpiryDate *DateFormat31Choice `xml:"XpryDt,omitempty"`

	// Date/time by which cash must be in place in order to take part in the event.
	SubscriptionCostDebitDate *DateFormat31Choice `xml:"SbcptCostDbtDt,omitempty"`

	// Last day that a participant of the depository can deliver securities that it had elected on and/or previously protected.
	DepositoryCoverExpirationDate *DateFormat31Choice `xml:"DpstryCoverXprtnDt,omitempty"`

	// Date/time set as the deadline to respond, with instructions, to an outstanding event, for which the underlying security is out on loan.
	StockLendingDeadline *DateFormat31Choice `xml:"StockLndgDdln,omitempty"`

	// Specifies the party borrowing stocks and the associated stock lending deadline assigned to the borrower.
	BorrowerStockLendingDeadline []*BorrowerLendingDeadline1 `xml:"BrrwrStockLndgDdln,omitempty"`
}

func (c *CorporateActionDate48) AddEarlyResponseDeadline() *DateFormat31Choice {
	c.EarlyResponseDeadline = new(DateFormat31Choice)
	return c.EarlyResponseDeadline
}

func (c *CorporateActionDate48) AddCoverExpirationDate() *DateFormat31Choice {
	c.CoverExpirationDate = new(DateFormat31Choice)
	return c.CoverExpirationDate
}

func (c *CorporateActionDate48) AddProtectDate() *DateFormat31Choice {
	c.ProtectDate = new(DateFormat31Choice)
	return c.ProtectDate
}

func (c *CorporateActionDate48) AddMarketDeadline() *DateFormat31Choice {
	c.MarketDeadline = new(DateFormat31Choice)
	return c.MarketDeadline
}

func (c *CorporateActionDate48) AddResponseDeadline() *DateFormat32Choice {
	c.ResponseDeadline = new(DateFormat32Choice)
	return c.ResponseDeadline
}

func (c *CorporateActionDate48) AddExpiryDate() *DateFormat31Choice {
	c.ExpiryDate = new(DateFormat31Choice)
	return c.ExpiryDate
}

func (c *CorporateActionDate48) AddSubscriptionCostDebitDate() *DateFormat31Choice {
	c.SubscriptionCostDebitDate = new(DateFormat31Choice)
	return c.SubscriptionCostDebitDate
}

func (c *CorporateActionDate48) AddDepositoryCoverExpirationDate() *DateFormat31Choice {
	c.DepositoryCoverExpirationDate = new(DateFormat31Choice)
	return c.DepositoryCoverExpirationDate
}

func (c *CorporateActionDate48) AddStockLendingDeadline() *DateFormat31Choice {
	c.StockLendingDeadline = new(DateFormat31Choice)
	return c.StockLendingDeadline
}

func (c *CorporateActionDate48) AddBorrowerStockLendingDeadline() *BorrowerLendingDeadline1 {
	newValue := new(BorrowerLendingDeadline1)
	c.BorrowerStockLendingDeadline = append(c.BorrowerStockLendingDeadline, newValue)
	return newValue
}

// Specifies corporate action dates.
type CorporateActionDate49 struct {

	// Date/time at which positions are struck at the end of the day to note which parties will receive the relevant amount of entitlement, due to be distributed on payment date.
	RecordDate *DateFormat31Choice `xml:"RcrdDt,omitempty"`

	// Date/time as from which trading (including exchange and OTC trading) occurs on the underlying security without the benefit.
	ExDividendDate *DateFormat31Choice `xml:"ExDvddDt,omitempty"`
}

func (c *CorporateActionDate49) AddRecordDate() *DateFormat31Choice {
	c.RecordDate = new(DateFormat31Choice)
	return c.RecordDate
}

func (c *CorporateActionDate49) AddExDividendDate() *DateFormat31Choice {
	c.ExDividendDate = new(DateFormat31Choice)
	return c.ExDividendDate
}

// Specifies corporate action dates.
type CorporateActionDate5 struct {

	// Date/time at which a foreign exchange rate will be determined.
	ForeignExchangeRateFixingDate *DateFormat4Choice `xml:"FXRateFxgDt,omitempty"`

	// Date/time at which assets become available to the account owner (in a credit entry), or cease to be available to the account owner (in a debit entry).
	ValueDate *DateFormat4Choice `xml:"ValDt,omitempty"`

	// Date/time at which the distribution is due to take place (cash and/or securities).
	PaymentDate *DateFormat4Choice `xml:"PmtDt,omitempty"`

	// Date/time at which a payment can be made, eg, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat4Choice `xml:"EarlstPmtDt,omitempty"`
}

func (c *CorporateActionDate5) AddForeignExchangeRateFixingDate() *DateFormat4Choice {
	c.ForeignExchangeRateFixingDate = new(DateFormat4Choice)
	return c.ForeignExchangeRateFixingDate
}

func (c *CorporateActionDate5) AddValueDate() *DateFormat4Choice {
	c.ValueDate = new(DateFormat4Choice)
	return c.ValueDate
}

func (c *CorporateActionDate5) AddPaymentDate() *DateFormat4Choice {
	c.PaymentDate = new(DateFormat4Choice)
	return c.PaymentDate
}

func (c *CorporateActionDate5) AddEarliestPaymentDate() *DateFormat4Choice {
	c.EarliestPaymentDate = new(DateFormat4Choice)
	return c.EarliestPaymentDate
}

// Specifies corporate action dates.
type CorporateActionDate50 struct {

	// Date/time at which positions are struck at the end of the day to note which parties will receive the relevant amount of entitlement, due to be distributed on payment date.
	RecordDate *DateFormat34Choice `xml:"RcrdDt,omitempty"`

	// Date/time as from which trading (including exchange and OTC trading) occurs on the underlying security without the benefit.
	ExDividendDate *DateFormat34Choice `xml:"ExDvddDt,omitempty"`
}

func (c *CorporateActionDate50) AddRecordDate() *DateFormat34Choice {
	c.RecordDate = new(DateFormat34Choice)
	return c.RecordDate
}

func (c *CorporateActionDate50) AddExDividendDate() *DateFormat34Choice {
	c.ExDividendDate = new(DateFormat34Choice)
	return c.ExDividendDate
}

// Specifies corporate action date.
type CorporateActionDate52 struct {

	// Date/time at which the account servicer has set as the deadline to respond, with instructions, to an outstanding event. This time is dependent on the reference time zone of the account servicer as specified in a Service Level Agreement (SLA).
	ResponseDeadline *DateFormat34Choice `xml:"RspnDdln,omitempty"`

	// Date/time by which cash must be in place in order to take part in the event.
	SubscriptionCostDebitDate *DateFormat34Choice `xml:"SbcptCostDbtDt,omitempty"`

	// Issuer or issuer's agent deadline to respond, with an instruction, to an outstanding offer or privilege.
	MarketDeadline *DateFormat34Choice `xml:"MktDdln,omitempty"`

	// Date/time at which an order expires or on which a privilege or offer terminates.
	ExpiryDate *DateFormat34Choice `xml:"XpryDt,omitempty"`

	// Last day a holder can deliver the securities that it had elected on and/or previously protected.
	CoverExpirationDate *DateFormat34Choice `xml:"CoverXprtnDt,omitempty"`

	// Last date/time a holder can request to defer delivery of securities pursuant to a notice of guaranteed delivery or other required documentation.
	ProtectDate *DateFormat34Choice `xml:"PrtctDt,omitempty"`

	// Date/time at which the deal (rights) was agreed.
	TradingDate *DateFormat34Choice `xml:"TradgDt,omitempty"`
}

func (c *CorporateActionDate52) AddResponseDeadline() *DateFormat34Choice {
	c.ResponseDeadline = new(DateFormat34Choice)
	return c.ResponseDeadline
}

func (c *CorporateActionDate52) AddSubscriptionCostDebitDate() *DateFormat34Choice {
	c.SubscriptionCostDebitDate = new(DateFormat34Choice)
	return c.SubscriptionCostDebitDate
}

func (c *CorporateActionDate52) AddMarketDeadline() *DateFormat34Choice {
	c.MarketDeadline = new(DateFormat34Choice)
	return c.MarketDeadline
}

func (c *CorporateActionDate52) AddExpiryDate() *DateFormat34Choice {
	c.ExpiryDate = new(DateFormat34Choice)
	return c.ExpiryDate
}

func (c *CorporateActionDate52) AddCoverExpirationDate() *DateFormat34Choice {
	c.CoverExpirationDate = new(DateFormat34Choice)
	return c.CoverExpirationDate
}

func (c *CorporateActionDate52) AddProtectDate() *DateFormat34Choice {
	c.ProtectDate = new(DateFormat34Choice)
	return c.ProtectDate
}

func (c *CorporateActionDate52) AddTradingDate() *DateFormat34Choice {
	c.TradingDate = new(DateFormat34Choice)
	return c.TradingDate
}

// Specifies corporate action dates.
type CorporateActionDate54 struct {

	// Date/time at which positions are struck at the end of the day to note which parties will receive the relevant amount of entitlement, due to be distributed on payment date.
	RecordDate *DateFormat34Choice `xml:"RcrdDt,omitempty"`

	// Date/time as from which trading (including exchange and OTC trading) occurs on the underlying security without the benefit.
	ExDividendDate *DateFormat34Choice `xml:"ExDvddDt,omitempty"`

	// Date/time on which the lottery is run and applied to the holder's positions. This is also applicable to partial calls.
	LotteryDate *DateFormat34Choice `xml:"LtryDt,omitempty"`
}

func (c *CorporateActionDate54) AddRecordDate() *DateFormat34Choice {
	c.RecordDate = new(DateFormat34Choice)
	return c.RecordDate
}

func (c *CorporateActionDate54) AddExDividendDate() *DateFormat34Choice {
	c.ExDividendDate = new(DateFormat34Choice)
	return c.ExDividendDate
}

func (c *CorporateActionDate54) AddLotteryDate() *DateFormat34Choice {
	c.LotteryDate = new(DateFormat34Choice)
	return c.LotteryDate
}

// Specifies corporate action dates.
type CorporateActionDate55 struct {

	// Date/time that the account servicer has set as the deadline to respond, with instructions, to an outstanding event, giving the holder eligibility to incentives. This time is dependent on the reference time zone of the account servicer as specified in a Service Level Agreement (SLA).
	EarlyResponseDeadline *DateFormat34Choice `xml:"EarlyRspnDdln,omitempty"`

	// Last day a holder can deliver the securities that it had elected on and/or previously protected.
	CoverExpirationDate *DateFormat34Choice `xml:"CoverXprtnDt,omitempty"`

	// Last date/time a holder can request to defer delivery of securities pursuant to a notice of guaranteed delivery or other required documentation.
	ProtectDate *DateFormat34Choice `xml:"PrtctDt,omitempty"`

	// Issuer or issuer's agent deadline to respond, with an instruction, to an outstanding offer or privilege.
	MarketDeadline *DateFormat34Choice `xml:"MktDdln,omitempty"`

	// Date/time at which the account servicer has set as the deadline to respond, with instructions, to an outstanding event. This time is dependent on the reference time zone of the account servicer as specified in a Service Level Agreement (SLA).
	ResponseDeadline *DateFormat38Choice `xml:"RspnDdln,omitempty"`

	// Date/time at which an order expires or on which a privilege or offer terminates.
	ExpiryDate *DateFormat34Choice `xml:"XpryDt,omitempty"`

	// Date/time by which cash must be in place in order to take part in the event.
	SubscriptionCostDebitDate *DateFormat34Choice `xml:"SbcptCostDbtDt,omitempty"`

	// Last day that a participant of the depository can deliver securities that it had elected on and/or previously protected.
	DepositoryCoverExpirationDate *DateFormat34Choice `xml:"DpstryCoverXprtnDt,omitempty"`

	// Date/time set as the deadline to respond, with instructions, to an outstanding event, for which the underlying security is out on loan.
	StockLendingDeadline *DateFormat34Choice `xml:"StockLndgDdln,omitempty"`

	// Specifies the party borrowing stocks and the associated stock lending deadline assigned to the borrower.
	BorrowerStockLendingDeadline []*BorrowerLendingDeadline2 `xml:"BrrwrStockLndgDdln,omitempty"`
}

func (c *CorporateActionDate55) AddEarlyResponseDeadline() *DateFormat34Choice {
	c.EarlyResponseDeadline = new(DateFormat34Choice)
	return c.EarlyResponseDeadline
}

func (c *CorporateActionDate55) AddCoverExpirationDate() *DateFormat34Choice {
	c.CoverExpirationDate = new(DateFormat34Choice)
	return c.CoverExpirationDate
}

func (c *CorporateActionDate55) AddProtectDate() *DateFormat34Choice {
	c.ProtectDate = new(DateFormat34Choice)
	return c.ProtectDate
}

func (c *CorporateActionDate55) AddMarketDeadline() *DateFormat34Choice {
	c.MarketDeadline = new(DateFormat34Choice)
	return c.MarketDeadline
}

func (c *CorporateActionDate55) AddResponseDeadline() *DateFormat38Choice {
	c.ResponseDeadline = new(DateFormat38Choice)
	return c.ResponseDeadline
}

func (c *CorporateActionDate55) AddExpiryDate() *DateFormat34Choice {
	c.ExpiryDate = new(DateFormat34Choice)
	return c.ExpiryDate
}

func (c *CorporateActionDate55) AddSubscriptionCostDebitDate() *DateFormat34Choice {
	c.SubscriptionCostDebitDate = new(DateFormat34Choice)
	return c.SubscriptionCostDebitDate
}

func (c *CorporateActionDate55) AddDepositoryCoverExpirationDate() *DateFormat34Choice {
	c.DepositoryCoverExpirationDate = new(DateFormat34Choice)
	return c.DepositoryCoverExpirationDate
}

func (c *CorporateActionDate55) AddStockLendingDeadline() *DateFormat34Choice {
	c.StockLendingDeadline = new(DateFormat34Choice)
	return c.StockLendingDeadline
}

func (c *CorporateActionDate55) AddBorrowerStockLendingDeadline() *BorrowerLendingDeadline2 {
	newValue := new(BorrowerLendingDeadline2)
	c.BorrowerStockLendingDeadline = append(c.BorrowerStockLendingDeadline, newValue)
	return newValue
}

// Specifies corporate action dates.
type CorporateActionDate56 struct {

	// Date on which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat34Choice `xml:"PmtDt"`

	// Date/time when calculating economic benefit for a cash amount.
	ValueDate *DateFormat39Choice `xml:"ValDt,omitempty"`

	// Date/time at which a foreign exchange rate will be determined.
	ForeignExchangeRateFixingDate *DateFormat34Choice `xml:"FXRateFxgDt,omitempty"`

	// Date on which a payment can be made, for example, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat34Choice `xml:"EarlstPmtDt,omitempty"`
}

func (c *CorporateActionDate56) AddPaymentDate() *DateFormat34Choice {
	c.PaymentDate = new(DateFormat34Choice)
	return c.PaymentDate
}

func (c *CorporateActionDate56) AddValueDate() *DateFormat39Choice {
	c.ValueDate = new(DateFormat39Choice)
	return c.ValueDate
}

func (c *CorporateActionDate56) AddForeignExchangeRateFixingDate() *DateFormat34Choice {
	c.ForeignExchangeRateFixingDate = new(DateFormat34Choice)
	return c.ForeignExchangeRateFixingDate
}

func (c *CorporateActionDate56) AddEarliestPaymentDate() *DateFormat34Choice {
	c.EarliestPaymentDate = new(DateFormat34Choice)
	return c.EarliestPaymentDate
}

// Specifies corporate action dates.
type CorporateActionDate58 struct {

	// Date/time at which the issuer announced that a corporate action event will occur.
	AnnouncementDate *DateFormat34Choice `xml:"AnncmntDt,omitempty"`

	// Deadline by which the beneficial ownership of securities must be declared.
	CertificationDeadline *DateFormat34Choice `xml:"CertfctnDdln,omitempty"`

	// Date upon which the court provided approval.
	CourtApprovalDate *DateFormat34Choice `xml:"CrtApprvlDt,omitempty"`

	// First possible early closing date of an offer if different from the expiry date.
	EarlyClosingDate *DateFormat34Choice `xml:"EarlyClsgDt,omitempty"`

	// Date/time at which an event is officially effective from the issuer's perspective.
	EffectiveDate *DateFormat34Choice `xml:"FctvDt,omitempty"`

	// Date/Time on which all or part of any holding bought in a unit trust is subject to being treated as capital rather than income. This is normally one day after the previous distribution's ex date.
	EqualisationDate *DateFormat34Choice `xml:"EqulstnDt,omitempty"`

	// Date/time at which additional information on the event will be announced, for example, exchange ratio announcement date.
	FurtherDetailedAnnouncementDate *DateFormat34Choice `xml:"FrthrDtldAnncmntDt,omitempty"`

	// Date/time at which an index / rate / price / value will be determined.
	FixingDate *DateFormat34Choice `xml:"FxgDt,omitempty"`

	// Date/time on which the lottery is run and applied to the holder's positions. This is also applicable to partial calls.
	LotteryDate *DateFormat34Choice `xml:"LtryDt,omitempty"`

	// Date/time to which the maturity date of an interest bearing security is extended.
	NewMaturityDate *DateFormat34Choice `xml:"NewMtrtyDt,omitempty"`

	// Date/time on which the bondholder's or shareholder's meeting will take place.
	MeetingDate *DateFormat34Choice `xml:"MtgDt,omitempty"`

	// Date/time at which the margin rate will be determined .
	MarginFixingDate *DateFormat34Choice `xml:"MrgnFxgDt,omitempty"`

	// Date/time (and time) at which an issuer will determine the proration amount/quantity of an offer.
	ProrationDate *DateFormat34Choice `xml:"PrratnDt,omitempty"`

	// Date/time at which positions are struck at the end of the day to note which parties will receive the relevant amount of entitlement, due to be distributed on payment date.
	RecordDate *DateFormat34Choice `xml:"RcrdDt,omitempty"`

	// Date/time on which instructions to register or registration details will be accepted.
	RegistrationDeadline *DateFormat34Choice `xml:"RegnDdln,omitempty"`

	// Date/time on which results are published, for example, results of an offer.
	ResultsPublicationDate *DateFormat34Choice `xml:"RsltsPblctnDt,omitempty"`

	// Deadline by which instructions must be received to split securities, for example, of physical certificates.
	DeadlineToSplit *DateFormat34Choice `xml:"DdlnToSplt,omitempty"`

	// Date/time on until which tax breakdown instructions will be accepted.
	DeadlineForTaxBreakdownInstruction *DateFormat34Choice `xml:"DdlnForTaxBrkdwnInstr,omitempty"`

	// Date/time at which trading of a security is suspended as the result of an event.
	TradingSuspendedDate *DateFormat34Choice `xml:"TradgSspdDt,omitempty"`

	// Date/time upon which the terms of the take-over become unconditional as to acceptances.
	UnconditionalDate *DateFormat34Choice `xml:"UcondlDt,omitempty"`

	// Date/time at on which all conditions, including regulatory, legal etc. pertaining to the take-over, have been met.
	WhollyUnconditionalDate *DateFormat34Choice `xml:"WhlyUcondlDt,omitempty"`

	// Date/time as from which trading (including exchange and OTC trading) occurs on the underlying security without the benefit.
	ExDividendDate *DateFormat34Choice `xml:"ExDvddDt,omitempty"`

	// Date/time at which the corporate action is legally announced by an official body, for example, publication by a governmental administration.
	OfficialAnnouncementPublicationDate *DateFormat34Choice `xml:"OffclAnncmntPblctnDt,omitempty"`

	// Date/time as from which 'special processing' can start to be used by participants for that event. Special processing is a means of marking a transaction, that would normally be traded ex or cum, as being traded cum or ex respectively,  for example, a transaction dealt 'special' after the ex date would result in the buyer being eligible for the entitlement. This is typically used in the UK and Irish markets.
	SpecialExDate *DateFormat34Choice `xml:"SpclExDt,omitempty"`

	// Last date/time by which a buying counterparty to a trade can be sure that it will have the right to participate in an event.
	GuaranteedParticipationDate *DateFormat34Choice `xml:"GrntedPrtcptnDt,omitempty"`

	// Deadline by which an entitled holder needs to advise their counterparty to a transaction of their election for a corporate action event, also known as Buyer Protection Deadline.
	ElectionToCounterpartyMarketDeadline *DateFormat34Choice `xml:"ElctnToCtrPtyMktDdln,omitempty"`

	// Date/time the account servicer has set as the deadline to respond, with instructions, prior to the election to counterparty market deadline
	ElectionToCounterpartyResponseDeadline *DateFormat34Choice `xml:"ElctnToCtrPtyRspnDdln,omitempty"`

	// Date/time at which an event/offer is terminated or lapsed.
	LapsedDate *DateFormat34Choice `xml:"LpsdDt,omitempty"`

	// Date/time at which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat34Choice `xml:"PmtDt,omitempty"`

	// Date/Time by which the account owner must instruct directly another party, for example to provide documentation to an issuer agent.
	ThirdPartyDeadline *DateFormat34Choice `xml:"ThrdPtyDdln,omitempty"`

	// Date/Time set by the issuer agent as a first early deadline by which the account owner must instruct directly another party, possibly giving the holder eligibility to incentives. For example, to provide documentation to an issuer agent.
	EarlyThirdPartyDeadline *DateFormat34Choice `xml:"EarlyThrdPtyDdln,omitempty"`

	// Date by which the depository stops monitoring activities of the event, for instance, accounting and tracking activities for due bills end.
	MarketClaimTrackingEndDate *DateFormat34Choice `xml:"MktClmTrckgEndDt,omitempty"`

	// Last day an investor can become a lead plaintiff.
	LeadPlaintiffDeadline *DateFormat34Choice `xml:"LeadPlntffDdln,omitempty"`

	// Date on which the action was filed at the applicable court.
	FilingDate *DateFormat41Choice `xml:"FilgDt,omitempty"`

	// Date for the hearing between the plaintiff and defendant, as set by the court.
	HearingDate *DateFormat41Choice `xml:"HrgDt,omitempty"`
}

func (c *CorporateActionDate58) AddAnnouncementDate() *DateFormat34Choice {
	c.AnnouncementDate = new(DateFormat34Choice)
	return c.AnnouncementDate
}

func (c *CorporateActionDate58) AddCertificationDeadline() *DateFormat34Choice {
	c.CertificationDeadline = new(DateFormat34Choice)
	return c.CertificationDeadline
}

func (c *CorporateActionDate58) AddCourtApprovalDate() *DateFormat34Choice {
	c.CourtApprovalDate = new(DateFormat34Choice)
	return c.CourtApprovalDate
}

func (c *CorporateActionDate58) AddEarlyClosingDate() *DateFormat34Choice {
	c.EarlyClosingDate = new(DateFormat34Choice)
	return c.EarlyClosingDate
}

func (c *CorporateActionDate58) AddEffectiveDate() *DateFormat34Choice {
	c.EffectiveDate = new(DateFormat34Choice)
	return c.EffectiveDate
}

func (c *CorporateActionDate58) AddEqualisationDate() *DateFormat34Choice {
	c.EqualisationDate = new(DateFormat34Choice)
	return c.EqualisationDate
}

func (c *CorporateActionDate58) AddFurtherDetailedAnnouncementDate() *DateFormat34Choice {
	c.FurtherDetailedAnnouncementDate = new(DateFormat34Choice)
	return c.FurtherDetailedAnnouncementDate
}

func (c *CorporateActionDate58) AddFixingDate() *DateFormat34Choice {
	c.FixingDate = new(DateFormat34Choice)
	return c.FixingDate
}

func (c *CorporateActionDate58) AddLotteryDate() *DateFormat34Choice {
	c.LotteryDate = new(DateFormat34Choice)
	return c.LotteryDate
}

func (c *CorporateActionDate58) AddNewMaturityDate() *DateFormat34Choice {
	c.NewMaturityDate = new(DateFormat34Choice)
	return c.NewMaturityDate
}

func (c *CorporateActionDate58) AddMeetingDate() *DateFormat34Choice {
	c.MeetingDate = new(DateFormat34Choice)
	return c.MeetingDate
}

func (c *CorporateActionDate58) AddMarginFixingDate() *DateFormat34Choice {
	c.MarginFixingDate = new(DateFormat34Choice)
	return c.MarginFixingDate
}

func (c *CorporateActionDate58) AddProrationDate() *DateFormat34Choice {
	c.ProrationDate = new(DateFormat34Choice)
	return c.ProrationDate
}

func (c *CorporateActionDate58) AddRecordDate() *DateFormat34Choice {
	c.RecordDate = new(DateFormat34Choice)
	return c.RecordDate
}

func (c *CorporateActionDate58) AddRegistrationDeadline() *DateFormat34Choice {
	c.RegistrationDeadline = new(DateFormat34Choice)
	return c.RegistrationDeadline
}

func (c *CorporateActionDate58) AddResultsPublicationDate() *DateFormat34Choice {
	c.ResultsPublicationDate = new(DateFormat34Choice)
	return c.ResultsPublicationDate
}

func (c *CorporateActionDate58) AddDeadlineToSplit() *DateFormat34Choice {
	c.DeadlineToSplit = new(DateFormat34Choice)
	return c.DeadlineToSplit
}

func (c *CorporateActionDate58) AddDeadlineForTaxBreakdownInstruction() *DateFormat34Choice {
	c.DeadlineForTaxBreakdownInstruction = new(DateFormat34Choice)
	return c.DeadlineForTaxBreakdownInstruction
}

func (c *CorporateActionDate58) AddTradingSuspendedDate() *DateFormat34Choice {
	c.TradingSuspendedDate = new(DateFormat34Choice)
	return c.TradingSuspendedDate
}

func (c *CorporateActionDate58) AddUnconditionalDate() *DateFormat34Choice {
	c.UnconditionalDate = new(DateFormat34Choice)
	return c.UnconditionalDate
}

func (c *CorporateActionDate58) AddWhollyUnconditionalDate() *DateFormat34Choice {
	c.WhollyUnconditionalDate = new(DateFormat34Choice)
	return c.WhollyUnconditionalDate
}

func (c *CorporateActionDate58) AddExDividendDate() *DateFormat34Choice {
	c.ExDividendDate = new(DateFormat34Choice)
	return c.ExDividendDate
}

func (c *CorporateActionDate58) AddOfficialAnnouncementPublicationDate() *DateFormat34Choice {
	c.OfficialAnnouncementPublicationDate = new(DateFormat34Choice)
	return c.OfficialAnnouncementPublicationDate
}

func (c *CorporateActionDate58) AddSpecialExDate() *DateFormat34Choice {
	c.SpecialExDate = new(DateFormat34Choice)
	return c.SpecialExDate
}

func (c *CorporateActionDate58) AddGuaranteedParticipationDate() *DateFormat34Choice {
	c.GuaranteedParticipationDate = new(DateFormat34Choice)
	return c.GuaranteedParticipationDate
}

func (c *CorporateActionDate58) AddElectionToCounterpartyMarketDeadline() *DateFormat34Choice {
	c.ElectionToCounterpartyMarketDeadline = new(DateFormat34Choice)
	return c.ElectionToCounterpartyMarketDeadline
}

func (c *CorporateActionDate58) AddElectionToCounterpartyResponseDeadline() *DateFormat34Choice {
	c.ElectionToCounterpartyResponseDeadline = new(DateFormat34Choice)
	return c.ElectionToCounterpartyResponseDeadline
}

func (c *CorporateActionDate58) AddLapsedDate() *DateFormat34Choice {
	c.LapsedDate = new(DateFormat34Choice)
	return c.LapsedDate
}

func (c *CorporateActionDate58) AddPaymentDate() *DateFormat34Choice {
	c.PaymentDate = new(DateFormat34Choice)
	return c.PaymentDate
}

func (c *CorporateActionDate58) AddThirdPartyDeadline() *DateFormat34Choice {
	c.ThirdPartyDeadline = new(DateFormat34Choice)
	return c.ThirdPartyDeadline
}

func (c *CorporateActionDate58) AddEarlyThirdPartyDeadline() *DateFormat34Choice {
	c.EarlyThirdPartyDeadline = new(DateFormat34Choice)
	return c.EarlyThirdPartyDeadline
}

func (c *CorporateActionDate58) AddMarketClaimTrackingEndDate() *DateFormat34Choice {
	c.MarketClaimTrackingEndDate = new(DateFormat34Choice)
	return c.MarketClaimTrackingEndDate
}

func (c *CorporateActionDate58) AddLeadPlaintiffDeadline() *DateFormat34Choice {
	c.LeadPlaintiffDeadline = new(DateFormat34Choice)
	return c.LeadPlaintiffDeadline
}

func (c *CorporateActionDate58) AddFilingDate() *DateFormat41Choice {
	c.FilingDate = new(DateFormat41Choice)
	return c.FilingDate
}

func (c *CorporateActionDate58) AddHearingDate() *DateFormat41Choice {
	c.HearingDate = new(DateFormat41Choice)
	return c.HearingDate
}

// Specifies corporate action date.
type CorporateActionDate6 struct {

	// Date/time at which the account servicer has set as the deadline to respond, with instructions, to an outstanding event. This time is dependent on the reference time zone of the account servicer as specified in a Service Level Agreement (SLA).
	ResponseDeadline *DateFormat6Choice `xml:"RspnDdln,omitempty"`

	// Date/time by which cash must be in place in order to take part in the event.
	SubscriptionCostDebitDate *DateFormat6Choice `xml:"SbcptCostDbtDt,omitempty"`

	// Issuer or issuer's agent deadline to respond, with an instruction, to an outstanding offer or privilege.
	MarketDeadline *DateFormat6Choice `xml:"MktDdln,omitempty"`

	// Date/time at which an order expires or on which a privilege or offer terminates.
	ExpiryDate *DateFormat6Choice `xml:"XpryDt,omitempty"`

	// Last day a holder can deliver the securities that it had elected on and/or previously protected.
	CoverExpirationDate *DateFormat6Choice `xml:"CoverXprtnDt,omitempty"`

	// Last date/time a holder can request to defer delivery of securities pursuant to a notice of guaranteed delivery or other required documentation.
	ProtectDate *DateFormat6Choice `xml:"PrtctDt,omitempty"`

	// Date/time at which the deal (rights) was agreed.
	TradingDate *DateFormat6Choice `xml:"TradgDt,omitempty"`
}

func (c *CorporateActionDate6) AddResponseDeadline() *DateFormat6Choice {
	c.ResponseDeadline = new(DateFormat6Choice)
	return c.ResponseDeadline
}

func (c *CorporateActionDate6) AddSubscriptionCostDebitDate() *DateFormat6Choice {
	c.SubscriptionCostDebitDate = new(DateFormat6Choice)
	return c.SubscriptionCostDebitDate
}

func (c *CorporateActionDate6) AddMarketDeadline() *DateFormat6Choice {
	c.MarketDeadline = new(DateFormat6Choice)
	return c.MarketDeadline
}

func (c *CorporateActionDate6) AddExpiryDate() *DateFormat6Choice {
	c.ExpiryDate = new(DateFormat6Choice)
	return c.ExpiryDate
}

func (c *CorporateActionDate6) AddCoverExpirationDate() *DateFormat6Choice {
	c.CoverExpirationDate = new(DateFormat6Choice)
	return c.CoverExpirationDate
}

func (c *CorporateActionDate6) AddProtectDate() *DateFormat6Choice {
	c.ProtectDate = new(DateFormat6Choice)
	return c.ProtectDate
}

func (c *CorporateActionDate6) AddTradingDate() *DateFormat6Choice {
	c.TradingDate = new(DateFormat6Choice)
	return c.TradingDate
}

// Specifies corporate action dates.
type CorporateActionDate7 struct {

	// Date/Time of the posting (credit or debit) to the account.
	PostingDate *DateAndDateTimeChoice `xml:"PstngDt"`

	// Date/time at which assets become available to the account owner (in a credit entry), or cease to be available to the account owner (in a debit entry).
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`

	// Date/time at which a foreign exchange rate will be determined.
	ForeignExchangeRateFixingDate *DateAndDateTimeChoice `xml:"FXRateFxgDt,omitempty"`

	// Date on which a payment can be made, for example, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateAndDateTimeChoice `xml:"EarlstPmtDt,omitempty"`

	// Date on which the distribution is due to take place (cash and/or securities).
	PaymentDate *DateAndDateTimeChoice `xml:"PmtDt,omitempty"`
}

func (c *CorporateActionDate7) AddPostingDate() *DateAndDateTimeChoice {
	c.PostingDate = new(DateAndDateTimeChoice)
	return c.PostingDate
}

func (c *CorporateActionDate7) AddValueDate() *DateAndDateTimeChoice {
	c.ValueDate = new(DateAndDateTimeChoice)
	return c.ValueDate
}

func (c *CorporateActionDate7) AddForeignExchangeRateFixingDate() *DateAndDateTimeChoice {
	c.ForeignExchangeRateFixingDate = new(DateAndDateTimeChoice)
	return c.ForeignExchangeRateFixingDate
}

func (c *CorporateActionDate7) AddEarliestPaymentDate() *DateAndDateTimeChoice {
	c.EarliestPaymentDate = new(DateAndDateTimeChoice)
	return c.EarliestPaymentDate
}

func (c *CorporateActionDate7) AddPaymentDate() *DateAndDateTimeChoice {
	c.PaymentDate = new(DateAndDateTimeChoice)
	return c.PaymentDate
}

// Specifies corporate action dates.
type CorporateActionDate8 struct {

	// Date/time that the account servicer has set as the deadline to respond, with instructions, to an outstanding event, giving the holder eligibility to incentives. This time is dependent on the reference time zone of the account servicer as specified in a Service Level Agreement (SLA).
	EarlyResponseDeadline *DateFormat6Choice `xml:"EarlyRspnDdln,omitempty"`

	// Last day a holder can deliver the securities that it had elected on and/or previously protected.
	CoverExpirationDate *DateFormat6Choice `xml:"CoverXprtnDt,omitempty"`

	// Last date/time a holder can request to defer delivery of securities pursuant to a notice of guaranteed delivery or other required documentation.
	ProtectDate *DateFormat6Choice `xml:"PrtctDt,omitempty"`

	// Issuer or issuer's agent deadline to respond, with an instruction, to an outstanding offer or privilege.
	MarketDeadline *DateFormat6Choice `xml:"MktDdln,omitempty"`

	// Date/time at which the account servicer has set as the deadline to respond, with instructions, to an outstanding event. This time is dependent on the reference time zone of the account servicer as specified in a Service Level Agreement (SLA).
	ResponseDeadline *DateFormat7Choice `xml:"RspnDdln,omitempty"`

	// Date/time at which an order expires or on which a privilege or offer terminates.
	ExpiryDate *DateFormat6Choice `xml:"XpryDt,omitempty"`

	// Date/time by which cash must be in place in order to take part in the event.
	SubscriptionCostDebitDate *DateFormat6Choice `xml:"SbcptCostDbtDt,omitempty"`

	// Last day that a participant of the depository can deliver securities that it had elected on and/or previously protected.
	DepositoryCoverExpirationDate *DateFormat6Choice `xml:"DpstryCoverXprtnDt,omitempty"`

	// Last day an investor can become a lead plaintiff.
	LeadPlaintiffDeadline *DateFormat6Choice `xml:"LeadPlntffDdln,omitempty"`
}

func (c *CorporateActionDate8) AddEarlyResponseDeadline() *DateFormat6Choice {
	c.EarlyResponseDeadline = new(DateFormat6Choice)
	return c.EarlyResponseDeadline
}

func (c *CorporateActionDate8) AddCoverExpirationDate() *DateFormat6Choice {
	c.CoverExpirationDate = new(DateFormat6Choice)
	return c.CoverExpirationDate
}

func (c *CorporateActionDate8) AddProtectDate() *DateFormat6Choice {
	c.ProtectDate = new(DateFormat6Choice)
	return c.ProtectDate
}

func (c *CorporateActionDate8) AddMarketDeadline() *DateFormat6Choice {
	c.MarketDeadline = new(DateFormat6Choice)
	return c.MarketDeadline
}

func (c *CorporateActionDate8) AddResponseDeadline() *DateFormat7Choice {
	c.ResponseDeadline = new(DateFormat7Choice)
	return c.ResponseDeadline
}

func (c *CorporateActionDate8) AddExpiryDate() *DateFormat6Choice {
	c.ExpiryDate = new(DateFormat6Choice)
	return c.ExpiryDate
}

func (c *CorporateActionDate8) AddSubscriptionCostDebitDate() *DateFormat6Choice {
	c.SubscriptionCostDebitDate = new(DateFormat6Choice)
	return c.SubscriptionCostDebitDate
}

func (c *CorporateActionDate8) AddDepositoryCoverExpirationDate() *DateFormat6Choice {
	c.DepositoryCoverExpirationDate = new(DateFormat6Choice)
	return c.DepositoryCoverExpirationDate
}

func (c *CorporateActionDate8) AddLeadPlaintiffDeadline() *DateFormat6Choice {
	c.LeadPlaintiffDeadline = new(DateFormat6Choice)
	return c.LeadPlaintiffDeadline
}

// Specifies corporate action dates.
type CorporateActionDate9 struct {

	// Date on which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat6Choice `xml:"PmtDt"`

	// Date at which assets become available to the account owner (in a credit entry), or cease to be available to the account owner (in a debit entry).
	ValueDate *DateFormat11Choice `xml:"ValDt,omitempty"`

	// Date/time at which a foreign exchange rate will be determined.
	ForeignExchangeRateFixingDate *DateFormat6Choice `xml:"FXRateFxgDt,omitempty"`

	// Date on which a payment can be made, for example, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat6Choice `xml:"EarlstPmtDt,omitempty"`
}

func (c *CorporateActionDate9) AddPaymentDate() *DateFormat6Choice {
	c.PaymentDate = new(DateFormat6Choice)
	return c.PaymentDate
}

func (c *CorporateActionDate9) AddValueDate() *DateFormat11Choice {
	c.ValueDate = new(DateFormat11Choice)
	return c.ValueDate
}

func (c *CorporateActionDate9) AddForeignExchangeRateFixingDate() *DateFormat6Choice {
	c.ForeignExchangeRateFixingDate = new(DateFormat6Choice)
	return c.ForeignExchangeRateFixingDate
}

func (c *CorporateActionDate9) AddEarliestPaymentDate() *DateFormat6Choice {
	c.EarliestPaymentDate = new(DateFormat6Choice)
	return c.EarliestPaymentDate
}
