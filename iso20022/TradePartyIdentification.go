package iso20022

// Entity involved in an activity.
type TradePartyIdentification3 struct {

	// Identifies the fund which is one of the parties in a treasury trade.
	FundInformation *FundIdentification2 `xml:"FndInf,omitempty"`

	// Specifies the party which submits a treasury trade to a matching system or to a settlement system or to a counterparty.
	SubmittingParty *PartyIdentification8Choice `xml:"SubmitgPty"`

	// Specifies the party which originated a treasury trade. This party may be the same as the submitting party.
	TradeParty *PartyIdentification8Choice `xml:"TradPty"`
}

func (t *TradePartyIdentification3) AddFundInformation() *FundIdentification2 {
	t.FundInformation = new(FundIdentification2)
	return t.FundInformation
}

func (t *TradePartyIdentification3) AddSubmittingParty() *PartyIdentification8Choice {
	t.SubmittingParty = new(PartyIdentification8Choice)
	return t.SubmittingParty
}

func (t *TradePartyIdentification3) AddTradeParty() *PartyIdentification8Choice {
	t.TradeParty = new(PartyIdentification8Choice)
	return t.TradeParty
}

// Entity involved in an activity.
type TradePartyIdentification4 struct {

	// Identifies the fund which is one of the parties in a treasury trade.
	FundInformation *FundIdentification2 `xml:"FndInf,omitempty"`

	// Specifies the party which is the buyer or the seller.
	BuyerOrSellerIndicator *OptionParty1Code `xml:"BuyrOrSellrInd"`

	// Specifies the party which submits a treasury trade to a matching system or to a settlement system or to a counterparty.
	SubmittingParty *PartyIdentification8Choice `xml:"SubmitgPty"`

	// Specifies the party which originated a treasury trade. This party may be the same as the submitting party.
	TradeParty *PartyIdentification8Choice `xml:"TradPty"`
}

func (t *TradePartyIdentification4) AddFundInformation() *FundIdentification2 {
	t.FundInformation = new(FundIdentification2)
	return t.FundInformation
}

func (t *TradePartyIdentification4) SetBuyerOrSellerIndicator(value string) {
	t.BuyerOrSellerIndicator = (*OptionParty1Code)(&value)
}

func (t *TradePartyIdentification4) AddSubmittingParty() *PartyIdentification8Choice {
	t.SubmittingParty = new(PartyIdentification8Choice)
	return t.SubmittingParty
}

func (t *TradePartyIdentification4) AddTradeParty() *PartyIdentification8Choice {
	t.TradeParty = new(PartyIdentification8Choice)
	return t.TradeParty
}

// Entity involved in an activity.
type TradePartyIdentification6 struct {

	// Party that submits the foreign exchange trade to the matching system or to the settlement system or to the counterparty.
	SubmittingParty *PartyIdentification73Choice `xml:"SubmitgPty"`

	// Party that originated the foreign exchange trade. This party may be the same as the submitting party.
	TradeParty *PartyIdentification73Choice `xml:"TradPty,omitempty"`

	// Identifies the fund that is one of the parties in the foreign exchange trade.
	//
	FundIdentification []*FundIdentification4 `xml:"FndId,omitempty"`
}

func (t *TradePartyIdentification6) AddSubmittingParty() *PartyIdentification73Choice {
	t.SubmittingParty = new(PartyIdentification73Choice)
	return t.SubmittingParty
}

func (t *TradePartyIdentification6) AddTradeParty() *PartyIdentification73Choice {
	t.TradeParty = new(PartyIdentification73Choice)
	return t.TradeParty
}

func (t *TradePartyIdentification6) AddFundIdentification() *FundIdentification4 {
	newValue := new(FundIdentification4)
	t.FundIdentification = append(t.FundIdentification, newValue)
	return newValue
}
