package iso20022

// Reason to reject the message.
type RejectionReason1 struct {

	// Reason to reject the message.
	Reason *MessageRejectedReason1Code `xml:"Rsn"`

	// Additional information about the rejection reason.
	AdditionalInformation *Max140Text `xml:"AddtlInf,omitempty"`

	// Linked previous reference that is invalid or unrecognised, of the message being rejected.
	LinkedMessagePreviousReference *AdditionalReference2 `xml:"LkdMsgPrvsRef,omitempty"`

	// Linked other reference that is invalid or unrecognised, of the message being rejected.
	LinkedMessageOtherReference *AdditionalReference2 `xml:"LkdMsgOthrRef,omitempty"`

	// Linked related reference that is invalid or unrecognised, of the message being rejected.
	LinkedMessageRelatedReference *AdditionalReference2 `xml:"LkdMsgRltdRef,omitempty"`
}

func (r *RejectionReason1) SetReason(value string) {
	r.Reason = (*MessageRejectedReason1Code)(&value)
}

func (r *RejectionReason1) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max140Text)(&value)
}

func (r *RejectionReason1) AddLinkedMessagePreviousReference() *AdditionalReference2 {
	r.LinkedMessagePreviousReference = new(AdditionalReference2)
	return r.LinkedMessagePreviousReference
}

func (r *RejectionReason1) AddLinkedMessageOtherReference() *AdditionalReference2 {
	r.LinkedMessageOtherReference = new(AdditionalReference2)
	return r.LinkedMessageOtherReference
}

func (r *RejectionReason1) AddLinkedMessageRelatedReference() *AdditionalReference2 {
	r.LinkedMessageRelatedReference = new(AdditionalReference2)
	return r.LinkedMessageRelatedReference
}

// The status of an instruction, advice or request.
type RejectionReason10 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionReason10Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason10) AddCode() *RejectionReason10Choice {
	r.Code = new(RejectionReason10Choice)
	return r.Code
}

func (r *RejectionReason10) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// The status of an instruction, advice or request.
type RejectionReason11 struct {

	// Reason provided for the status.
	Code *RejectionReason11Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason11) AddCode() *RejectionReason11Choice {
	r.Code = new(RejectionReason11Choice)
	return r.Code
}

func (r *RejectionReason11) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the underlying reason for the status of an object.
type RejectionReason12 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *ConsentOrRejectionReason2Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason12) AddCode() *ConsentOrRejectionReason2Choice {
	r.Code = new(ConsentOrRejectionReason2Choice)
	return r.Code
}

func (r *RejectionReason12) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Information about a rejected status.
type RejectionReason16 struct {

	// Reason for the rejected status.
	Reason *RejectedReason4Choice `xml:"Rsn"`

	// Additional information about the rejected status reason.
	AdditionalReasonInformation *Max350Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason16) AddReason() *RejectedReason4Choice {
	r.Reason = new(RejectedReason4Choice)
	return r.Reason
}

func (r *RejectionReason16) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max350Text)(&value)
}

// Status of an instruction, advice or request.
type RejectionReason17 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionReason14Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason17) AddCode() *RejectionReason14Choice {
	r.Code = new(RejectionReason14Choice)
	return r.Code
}

func (r *RejectionReason17) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// The status of an instruction, advice or request.
type RejectionReason18 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionReason15Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason18) AddCode() *RejectionReason15Choice {
	r.Code = new(RejectionReason15Choice)
	return r.Code
}

func (r *RejectionReason18) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the underlying rejection reason for the status of an object.
type RejectionReason19 struct {

	// Specifies the reason why the instruction/request has a rejected status.
	Code *RejectionReason17Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason19) AddCode() *RejectionReason17Choice {
	r.Code = new(RejectionReason17Choice)
	return r.Code
}

func (r *RejectionReason19) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// General information about the reason of the rejection.
type RejectionReason2 struct {

	// Reason of the rejection provided by the rejecting party.
	RejectingPartyReason *Max35Text `xml:"RjctgPtyRsn"`

	// Date and time at which the rejection was generated.
	RejectionDateTime *ISODateTime `xml:"RjctnDtTm,omitempty"`

	// Description of the precise location of the potential error in a message.
	ErrorLocation *Max350Text `xml:"ErrLctn,omitempty"`

	// Detailed description of the rejection reason.
	ReasonDescription *Max350Text `xml:"RsnDesc,omitempty"`

	// Additional information related to the rejection and meant to allow for the precise identification of the rejection reason. This could include a copy of the rejected message in part or in full.
	AdditionalData *Max20000Text `xml:"AddtlData,omitempty"`
}

func (r *RejectionReason2) SetRejectingPartyReason(value string) {
	r.RejectingPartyReason = (*Max35Text)(&value)
}

func (r *RejectionReason2) SetRejectionDateTime(value string) {
	r.RejectionDateTime = (*ISODateTime)(&value)
}

func (r *RejectionReason2) SetErrorLocation(value string) {
	r.ErrorLocation = (*Max350Text)(&value)
}

func (r *RejectionReason2) SetReasonDescription(value string) {
	r.ReasonDescription = (*Max350Text)(&value)
}

func (r *RejectionReason2) SetAdditionalData(value string) {
	r.AdditionalData = (*Max20000Text)(&value)
}

// Reason to reject the message.
type RejectionReason23 struct {

	// Reason to reject the message.
	Reason *MessageRejectedReason1Code `xml:"Rsn"`

	// Additional information about the rejection reason.
	AdditionalInformation *Max140Text `xml:"AddtlInf,omitempty"`

	// Identification of the invalid or unrecognised reference.
	LinkedMessage *LinkedMessage1Choice `xml:"LkdMsg,omitempty"`
}

func (r *RejectionReason23) SetReason(value string) {
	r.Reason = (*MessageRejectedReason1Code)(&value)
}

func (r *RejectionReason23) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max140Text)(&value)
}

func (r *RejectionReason23) AddLinkedMessage() *LinkedMessage1Choice {
	r.LinkedMessage = new(LinkedMessage1Choice)
	return r.LinkedMessage
}

// Specifies the reason why the instruction or request has a rejected status.
type RejectionReason25 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionReason23Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason25) AddCode() *RejectionReason23Choice {
	r.Code = new(RejectionReason23Choice)
	return r.Code
}

func (r *RejectionReason25) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the instruction or request has a rejected status.
//
type RejectionReason26 struct {

	// Reason provided for the status.
	Code *RejectionReason24Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason26) AddCode() *RejectionReason24Choice {
	r.Code = new(RejectionReason24Choice)
	return r.Code
}

func (r *RejectionReason26) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the underlying rejection reason for the status of an object.
type RejectionReason27 struct {

	// Specifies the reason why the instruction/request has a rejected status.
	Code *RejectionReason25Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason27) AddCode() *RejectionReason25Choice {
	r.Code = new(RejectionReason25Choice)
	return r.Code
}

func (r *RejectionReason27) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the instruction or request has a rejected status.
type RejectionReason28 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionReason26Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason28) AddCode() *RejectionReason26Choice {
	r.Code = new(RejectionReason26Choice)
	return r.Code
}

func (r *RejectionReason28) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Specifies the reason why the instruction or request has a rejected status.
type RejectionReason29 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *ConsentOrRejectionReason4Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason29) AddCode() *ConsentOrRejectionReason4Choice {
	r.Code = new(ConsentOrRejectionReason4Choice)
	return r.Code
}

func (r *RejectionReason29) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Reason to reject the message.
type RejectionReason3 struct {

	// Reason to reject the message.
	Reason *MessageRejectedReason1Code `xml:"Rsn"`

	// Additional information about the rejection reason.
	AdditionalInformation *Max140Text `xml:"AddtlInf,omitempty"`

	// Linked previous reference that is invalid or unrecognised, of the message being rejected.
	LinkedMessagePreviousReference *AdditionalReference3 `xml:"LkdMsgPrvsRef,omitempty"`

	// Linked other reference that is invalid or unrecognised, of the message being rejected.
	LinkedMessageOtherReference *AdditionalReference3 `xml:"LkdMsgOthrRef,omitempty"`

	// Linked related reference that is invalid or unrecognised, of the message being rejected.
	LinkedMessageRelatedReference *AdditionalReference3 `xml:"LkdMsgRltdRef,omitempty"`
}

func (r *RejectionReason3) SetReason(value string) {
	r.Reason = (*MessageRejectedReason1Code)(&value)
}

func (r *RejectionReason3) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max140Text)(&value)
}

func (r *RejectionReason3) AddLinkedMessagePreviousReference() *AdditionalReference3 {
	r.LinkedMessagePreviousReference = new(AdditionalReference3)
	return r.LinkedMessagePreviousReference
}

func (r *RejectionReason3) AddLinkedMessageOtherReference() *AdditionalReference3 {
	r.LinkedMessageOtherReference = new(AdditionalReference3)
	return r.LinkedMessageOtherReference
}

func (r *RejectionReason3) AddLinkedMessageRelatedReference() *AdditionalReference3 {
	r.LinkedMessageRelatedReference = new(AdditionalReference3)
	return r.LinkedMessageRelatedReference
}

// Specifies the reason why the instruction or request has a rejected status.
type RejectionReason30 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionReason27Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason30) AddCode() *RejectionReason27Choice {
	r.Code = new(RejectionReason27Choice)
	return r.Code
}

func (r *RejectionReason30) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// Information about a rejected status.
type RejectionReason31 struct {

	// Reason for the rejected status.
	Reason *RejectedReason16Choice `xml:"Rsn"`

	// Additional information about the rejected status reason.
	AdditionalReasonInformation *Max350Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason31) AddReason() *RejectedReason16Choice {
	r.Reason = new(RejectedReason16Choice)
	return r.Reason
}

func (r *RejectionReason31) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max350Text)(&value)
}

// Information about a rejected status.
type RejectionReason32 struct {

	// Reason for the rejected status.
	Reason *RejectedReason15Choice `xml:"Rsn"`

	// Additional information about the rejected status reason.
	AdditionalReasonInformation *Max350Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason32) AddReason() *RejectedReason15Choice {
	r.Reason = new(RejectedReason15Choice)
	return r.Reason
}

func (r *RejectionReason32) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max350Text)(&value)
}

// Information about a rejected status.
type RejectionReason33 struct {

	// Reason for the rejected status.
	Reason *RejectedReason17Choice `xml:"Rsn"`

	// Additional information about the rejected status reason.
	AdditionalReasonInformation *Max350Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason33) AddReason() *RejectedReason17Choice {
	r.Reason = new(RejectedReason17Choice)
	return r.Reason
}

func (r *RejectionReason33) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max350Text)(&value)
}

// Specifies the reason why the instruction or request has a rejected status.
type RejectionReason34 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionReason28Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason34) AddCode() *RejectionReason28Choice {
	r.Code = new(RejectionReason28Choice)
	return r.Code
}

func (r *RejectionReason34) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies the underlying rejection reason for the status of an object.
type RejectionReason35 struct {

	// Specifies the reason why the instruction/request has a rejected status.
	Code *RejectionReason29Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason35) AddCode() *RejectionReason29Choice {
	r.Code = new(RejectionReason29Choice)
	return r.Code
}

func (r *RejectionReason35) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies the reason why the instruction or request has a rejected status.
type RejectionReason36 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionReason30Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason36) AddCode() *RejectionReason30Choice {
	r.Code = new(RejectionReason30Choice)
	return r.Code
}

func (r *RejectionReason36) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies the reason why the instruction or request has a rejected status.
//
type RejectionReason37 struct {

	// Reason provided for the status.
	Code *RejectionReason31Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason37) AddCode() *RejectionReason31Choice {
	r.Code = new(RejectionReason31Choice)
	return r.Code
}

func (r *RejectionReason37) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies the reason why the instruction or request has a rejected status.
type RejectionReason39 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionReason33Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason39) AddCode() *RejectionReason33Choice {
	r.Code = new(RejectionReason33Choice)
	return r.Code
}

func (r *RejectionReason39) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// Specifies the reason why the instruction or request has a rejected status.
type RejectionReason40 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *ConsentOrRejectionReason5Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason40) AddCode() *ConsentOrRejectionReason5Choice {
	r.Code = new(ConsentOrRejectionReason5Choice)
	return r.Code
}

func (r *RejectionReason40) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}

// The status of an instruction, advice or request.
type RejectionReason5 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionReason4Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason5) AddCode() *RejectionReason4Choice {
	r.Code = new(RejectionReason4Choice)
	return r.Code
}

func (r *RejectionReason5) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// The status of an instruction, advice or request.
type RejectionReason6 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RejectionReason2Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason6) AddCode() *RejectionReason2Choice {
	r.Code = new(RejectionReason2Choice)
	return r.Code
}

func (r *RejectionReason6) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}

// The status of an instruction, advice or request.
type RejectionReason9 struct {

	// Specifies the reason why the instruction/request has a rejected status.
	Code *RejectionReason9Choice `xml:"Cd"`

	// Provides additional information about the reason in narrative form.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectionReason9) AddCode() *RejectionReason9Choice {
	r.Code = new(RejectionReason9Choice)
	return r.Code
}

func (r *RejectionReason9) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
