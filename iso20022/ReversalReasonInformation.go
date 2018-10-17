package iso20022

// Further information on the reversal reason of the transaction.
type ReversalReasonInformation1 struct {

	// Party issuing the reversal.
	ReversalOriginator *PartyIdentification8 `xml:"RvslOrgtr,omitempty"`

	// Specifies the reason for the reversal.
	ReversalReason *ReversalReason1Choice `xml:"RvslRsn,omitempty"`

	// Further details on the reversal reason.
	AdditionalReversalReasonInformation []*Max105Text `xml:"AddtlRvslRsnInf,omitempty"`
}

func (r *ReversalReasonInformation1) AddReversalOriginator() *PartyIdentification8 {
	r.ReversalOriginator = new(PartyIdentification8)
	return r.ReversalOriginator
}

func (r *ReversalReasonInformation1) AddReversalReason() *ReversalReason1Choice {
	r.ReversalReason = new(ReversalReason1Choice)
	return r.ReversalReason
}

func (r *ReversalReasonInformation1) AddAdditionalReversalReasonInformation(value string) {
	r.AdditionalReversalReasonInformation = append(r.AdditionalReversalReasonInformation, (*Max105Text)(&value))
}

// Set of elements used to provide information on the reason of the reversal of the transaction.
type ReversalReasonInformation6 struct {

	// Party that issues the reversal.
	Originator *PartyIdentification32 `xml:"Orgtr,omitempty"`

	// Specifies the reason for the reversal.
	Reason *ReversalReason4Choice `xml:"Rsn,omitempty"`

	// Further details on the reversal reason.
	AdditionalInformation []*Max105Text `xml:"AddtlInf,omitempty"`
}

func (r *ReversalReasonInformation6) AddOriginator() *PartyIdentification32 {
	r.Originator = new(PartyIdentification32)
	return r.Originator
}

func (r *ReversalReasonInformation6) AddReason() *ReversalReason4Choice {
	r.Reason = new(ReversalReason4Choice)
	return r.Reason
}

func (r *ReversalReasonInformation6) AddAdditionalInformation(value string) {
	r.AdditionalInformation = append(r.AdditionalInformation, (*Max105Text)(&value))
}
