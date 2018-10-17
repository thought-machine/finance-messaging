package iso20022

// Further information on the return reason of the transaction.
type ReturnReasonInformation1 struct {

	// Party issuing the return.
	ReturnOriginator *PartyIdentification8 `xml:"RtrOrgtr,omitempty"`

	// Specifies the reason for the return.
	ReturnReason *ReturnReason1Choice `xml:"RtrRsn,omitempty"`

	// Further details on the return reason.
	AdditionalReturnReasonInformation []*Max105Text `xml:"AddtlRtrRsnInf,omitempty"`
}

func (r *ReturnReasonInformation1) AddReturnOriginator() *PartyIdentification8 {
	r.ReturnOriginator = new(PartyIdentification8)
	return r.ReturnOriginator
}

func (r *ReturnReasonInformation1) AddReturnReason() *ReturnReason1Choice {
	r.ReturnReason = new(ReturnReason1Choice)
	return r.ReturnReason
}

func (r *ReturnReasonInformation1) AddAdditionalReturnReasonInformation(value string) {
	r.AdditionalReturnReasonInformation = append(r.AdditionalReturnReasonInformation, (*Max105Text)(&value))
}

// Set of elements used to provide information on the reason of the return of the transaction.
type ReturnReasonInformation10 struct {

	// Bank transaction code included in the original entry for the transaction.
	OriginalBankTransactionCode *BankTransactionCodeStructure4 `xml:"OrgnlBkTxCd,omitempty"`

	// Party that issues the return.
	Originator *PartyIdentification32 `xml:"Orgtr,omitempty"`

	// Specifies the reason for the return.
	Reason *ReturnReason5Choice `xml:"Rsn,omitempty"`

	// Further details on the return reason.
	AdditionalInformation []*Max105Text `xml:"AddtlInf,omitempty"`
}

func (r *ReturnReasonInformation10) AddOriginalBankTransactionCode() *BankTransactionCodeStructure4 {
	r.OriginalBankTransactionCode = new(BankTransactionCodeStructure4)
	return r.OriginalBankTransactionCode
}

func (r *ReturnReasonInformation10) AddOriginator() *PartyIdentification32 {
	r.Originator = new(PartyIdentification32)
	return r.Originator
}

func (r *ReturnReasonInformation10) AddReason() *ReturnReason5Choice {
	r.Reason = new(ReturnReason5Choice)
	return r.Reason
}

func (r *ReturnReasonInformation10) AddAdditionalInformation(value string) {
	r.AdditionalInformation = append(r.AdditionalInformation, (*Max105Text)(&value))
}

// Further information on the return reason of the transaction.
type ReturnReasonInformation5 struct {

	// Bank transaction code included in the original entry for the transaction.
	OriginalBankTransactionCode *BankTransactionCodeStructure1 `xml:"OrgnlBkTxCd,omitempty"`

	// Party issuing the return.
	ReturnOriginator *PartyIdentification8 `xml:"RtrOrgtr,omitempty"`

	// Specifies the reason for the return.
	ReturnReason *ReturnReason1Choice `xml:"RtrRsn,omitempty"`

	// Further details on the return reason.
	AdditionalReturnReasonInformation []*Max105Text `xml:"AddtlRtrRsnInf,omitempty"`
}

func (r *ReturnReasonInformation5) AddOriginalBankTransactionCode() *BankTransactionCodeStructure1 {
	r.OriginalBankTransactionCode = new(BankTransactionCodeStructure1)
	return r.OriginalBankTransactionCode
}

func (r *ReturnReasonInformation5) AddReturnOriginator() *PartyIdentification8 {
	r.ReturnOriginator = new(PartyIdentification8)
	return r.ReturnOriginator
}

func (r *ReturnReasonInformation5) AddReturnReason() *ReturnReason1Choice {
	r.ReturnReason = new(ReturnReason1Choice)
	return r.ReturnReason
}

func (r *ReturnReasonInformation5) AddAdditionalReturnReasonInformation(value string) {
	r.AdditionalReturnReasonInformation = append(r.AdditionalReturnReasonInformation, (*Max105Text)(&value))
}

// Set of elements used to provide information on the reason of the return of the transaction.
type ReturnReasonInformation9 struct {

	// Party that issues the return.
	Originator *PartyIdentification32 `xml:"Orgtr,omitempty"`

	// Specifies the reason for the return.
	Reason *ReturnReason5Choice `xml:"Rsn,omitempty"`

	// Further details on the return reason.
	AdditionalInformation []*Max105Text `xml:"AddtlInf,omitempty"`
}

func (r *ReturnReasonInformation9) AddOriginator() *PartyIdentification32 {
	r.Originator = new(PartyIdentification32)
	return r.Originator
}

func (r *ReturnReasonInformation9) AddReason() *ReturnReason5Choice {
	r.Reason = new(ReturnReason5Choice)
	return r.Reason
}

func (r *ReturnReasonInformation9) AddAdditionalInformation(value string) {
	r.AdditionalInformation = append(r.AdditionalInformation, (*Max105Text)(&value))
}
