package iso20022

// Outcome of the authorisation, and actions to perform.
type AuthorisationResult1 struct {

	// Type of party that has delivered or declined the card payment authorisation (the party is not identified).
	AuthorisationEntity *GenericIdentification33 `xml:"AuthstnNtty,omitempty"`

	// Response to an authorisation request.
	ResponseToAuthorisation *ResponseType1 `xml:"RspnToAuthstn"`

	// Value assigned by the authorising party.
	AuthorisationCode *Min6Max8Text `xml:"AuthstnCd,omitempty"`

	// Indicates whether the acquirer requires a further exchange completion after the completion of the transaction.
	CompletionRequired *TrueFalseIndicator `xml:"CmpltnReqrd,omitempty"`

	// Instructs the point of interaction (POI) how to contact the host to initiate the maintenance of the terminal.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AuthorisationResult1) AddAuthorisationEntity() *GenericIdentification33 {
	a.AuthorisationEntity = new(GenericIdentification33)
	return a.AuthorisationEntity
}

func (a *AuthorisationResult1) AddResponseToAuthorisation() *ResponseType1 {
	a.ResponseToAuthorisation = new(ResponseType1)
	return a.ResponseToAuthorisation
}

func (a *AuthorisationResult1) SetAuthorisationCode(value string) {
	a.AuthorisationCode = (*Min6Max8Text)(&value)
}

func (a *AuthorisationResult1) SetCompletionRequired(value string) {
	a.CompletionRequired = (*TrueFalseIndicator)(&value)
}

func (a *AuthorisationResult1) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}

// Outcome of the authorisation, and actions to perform.
type AuthorisationResult10 struct {

	// Type of party that has delivered or declined the card payment authorisation (the party is not identified).
	AuthorisationEntity *GenericIdentification90 `xml:"AuthstnNtty,omitempty"`

	// Response to an authorisation request.
	ResponseToAuthorisation *ResponseType5 `xml:"RspnToAuthstn"`

	// Value assigned by the authorising party.
	AuthorisationCode *Min6Max8Text `xml:"AuthstnCd,omitempty"`

	// Indicates whether the acquirer requires a further exchange completion after the completion of the transaction.
	CompletionRequired *TrueFalseIndicator `xml:"CmpltnReqrd,omitempty"`

	// Instructs the point of interaction (POI) how to contact the host to initiate the maintenance of the terminal.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AuthorisationResult10) AddAuthorisationEntity() *GenericIdentification90 {
	a.AuthorisationEntity = new(GenericIdentification90)
	return a.AuthorisationEntity
}

func (a *AuthorisationResult10) AddResponseToAuthorisation() *ResponseType5 {
	a.ResponseToAuthorisation = new(ResponseType5)
	return a.ResponseToAuthorisation
}

func (a *AuthorisationResult10) SetAuthorisationCode(value string) {
	a.AuthorisationCode = (*Min6Max8Text)(&value)
}

func (a *AuthorisationResult10) SetCompletionRequired(value string) {
	a.CompletionRequired = (*TrueFalseIndicator)(&value)
}

func (a *AuthorisationResult10) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}

// Outcome of the authorisation.
type AuthorisationResult11 struct {

	// Type of party that has delivered or declined the card payment authorisation (the party is not identified).
	AuthorisationEntity *GenericIdentification90 `xml:"AuthstnNtty,omitempty"`

	// Response to an authorisation request.
	ResponseToAuthorisation *ResponseType5 `xml:"RspnToAuthstn"`

	// Value assigned by the authorising party.
	AuthorisationCode *Min6Max8Text `xml:"AuthstnCd,omitempty"`
}

func (a *AuthorisationResult11) AddAuthorisationEntity() *GenericIdentification90 {
	a.AuthorisationEntity = new(GenericIdentification90)
	return a.AuthorisationEntity
}

func (a *AuthorisationResult11) AddResponseToAuthorisation() *ResponseType5 {
	a.ResponseToAuthorisation = new(ResponseType5)
	return a.ResponseToAuthorisation
}

func (a *AuthorisationResult11) SetAuthorisationCode(value string) {
	a.AuthorisationCode = (*Min6Max8Text)(&value)
}

// Outcome of the authorisation.
type AuthorisationResult12 struct {

	// Type of party that has delivered or declined the card payment authorisation (the party is not identified).
	AuthorisationEntity *GenericIdentification90 `xml:"AuthstnNtty,omitempty"`

	// Response to an authorisation request.
	ResponseToAuthorisation *ResponseType5 `xml:"RspnToAuthstn"`

	// Value assigned by the authorising party.
	AuthorisationCode *Min6Max8Text `xml:"AuthstnCd,omitempty"`

	// Acquirer has requested a contact to the maintenance host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AuthorisationResult12) AddAuthorisationEntity() *GenericIdentification90 {
	a.AuthorisationEntity = new(GenericIdentification90)
	return a.AuthorisationEntity
}

func (a *AuthorisationResult12) AddResponseToAuthorisation() *ResponseType5 {
	a.ResponseToAuthorisation = new(ResponseType5)
	return a.ResponseToAuthorisation
}

func (a *AuthorisationResult12) SetAuthorisationCode(value string) {
	a.AuthorisationCode = (*Min6Max8Text)(&value)
}

func (a *AuthorisationResult12) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}

// Outcome of the withdrawal authorisation.
type AuthorisationResult13 struct {

	// Type of party that has delivered or declined the card payment authorisation (the party is not identified).
	AuthorisationEntity *PartyType16Code `xml:"AuthstnNtty,omitempty"`

	// Result of the authorisation.
	AuthorisationResponse *ResponseType7 `xml:"AuthstnRspn"`

	// Trace of response by the entities in the path from the issuer to the ATM.
	ResponseTrace []*ResponseType8 `xml:"RspnTrac,omitempty"`

	// Value assigned by the authorising party.
	AuthorisationCode *Min6Max8Text `xml:"AuthstnCd,omitempty"`

	// Sequence of actions to be performed by the ATM to complete the transaction.
	Action []*Action7 `xml:"Actn,omitempty"`
}

func (a *AuthorisationResult13) SetAuthorisationEntity(value string) {
	a.AuthorisationEntity = (*PartyType16Code)(&value)
}

func (a *AuthorisationResult13) AddAuthorisationResponse() *ResponseType7 {
	a.AuthorisationResponse = new(ResponseType7)
	return a.AuthorisationResponse
}

func (a *AuthorisationResult13) AddResponseTrace() *ResponseType8 {
	newValue := new(ResponseType8)
	a.ResponseTrace = append(a.ResponseTrace, newValue)
	return newValue
}

func (a *AuthorisationResult13) SetAuthorisationCode(value string) {
	a.AuthorisationCode = (*Min6Max8Text)(&value)
}

func (a *AuthorisationResult13) AddAction() *Action7 {
	newValue := new(Action7)
	a.Action = append(a.Action, newValue)
	return newValue
}

// Outcome of the authorisation.
type AuthorisationResult2 struct {

	// Type of party that has delivered or declined the card payment authorisation (the party is not identified).
	AuthorisationEntity *GenericIdentification33 `xml:"AuthstnNtty,omitempty"`

	// Response to an authorisation request.
	ResponseToAuthorisation *ResponseType1 `xml:"RspnToAuthstn"`

	// Value assigned by the authorising party.
	AuthorisationCode *Min6Max8Text `xml:"AuthstnCd,omitempty"`
}

func (a *AuthorisationResult2) AddAuthorisationEntity() *GenericIdentification33 {
	a.AuthorisationEntity = new(GenericIdentification33)
	return a.AuthorisationEntity
}

func (a *AuthorisationResult2) AddResponseToAuthorisation() *ResponseType1 {
	a.ResponseToAuthorisation = new(ResponseType1)
	return a.ResponseToAuthorisation
}

func (a *AuthorisationResult2) SetAuthorisationCode(value string) {
	a.AuthorisationCode = (*Min6Max8Text)(&value)
}

// Outcome of the authorisation.
type AuthorisationResult3 struct {

	// Type of party that has delivered or declined the card payment authorisation (the party is not identified).
	AuthorisationEntity *GenericIdentification33 `xml:"AuthstnNtty,omitempty"`

	// Response to an authorisation request.
	ResponseToAuthorisation *ResponseType1 `xml:"RspnToAuthstn"`

	// Value assigned by the authorising party.
	AuthorisationCode *Min6Max8Text `xml:"AuthstnCd,omitempty"`

	// Acquirer has requested a contact to the maintenance host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AuthorisationResult3) AddAuthorisationEntity() *GenericIdentification33 {
	a.AuthorisationEntity = new(GenericIdentification33)
	return a.AuthorisationEntity
}

func (a *AuthorisationResult3) AddResponseToAuthorisation() *ResponseType1 {
	a.ResponseToAuthorisation = new(ResponseType1)
	return a.ResponseToAuthorisation
}

func (a *AuthorisationResult3) SetAuthorisationCode(value string) {
	a.AuthorisationCode = (*Min6Max8Text)(&value)
}

func (a *AuthorisationResult3) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}

// Outcome of the authorisation, and actions to perform.
type AuthorisationResult4 struct {

	// Type of party that has delivered or declined the card payment authorisation (the party is not identified).
	AuthorisationEntity *GenericIdentification70 `xml:"AuthstnNtty,omitempty"`

	// Response to an authorisation request.
	ResponseToAuthorisation *ResponseType1 `xml:"RspnToAuthstn"`

	// Value assigned by the authorising party.
	AuthorisationCode *Min6Max8Text `xml:"AuthstnCd,omitempty"`

	// Indicates whether the acquirer requires a further exchange completion after the completion of the transaction.
	CompletionRequired *TrueFalseIndicator `xml:"CmpltnReqrd,omitempty"`

	// Instructs the point of interaction (POI) how to contact the host to initiate the maintenance of the terminal.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AuthorisationResult4) AddAuthorisationEntity() *GenericIdentification70 {
	a.AuthorisationEntity = new(GenericIdentification70)
	return a.AuthorisationEntity
}

func (a *AuthorisationResult4) AddResponseToAuthorisation() *ResponseType1 {
	a.ResponseToAuthorisation = new(ResponseType1)
	return a.ResponseToAuthorisation
}

func (a *AuthorisationResult4) SetAuthorisationCode(value string) {
	a.AuthorisationCode = (*Min6Max8Text)(&value)
}

func (a *AuthorisationResult4) SetCompletionRequired(value string) {
	a.CompletionRequired = (*TrueFalseIndicator)(&value)
}

func (a *AuthorisationResult4) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}

// Outcome of the authorisation.
type AuthorisationResult5 struct {

	// Type of party that has delivered or declined the card payment authorisation (the party is not identified).
	AuthorisationEntity *GenericIdentification70 `xml:"AuthstnNtty,omitempty"`

	// Response to an authorisation request.
	ResponseToAuthorisation *ResponseType1 `xml:"RspnToAuthstn"`

	// Value assigned by the authorising party.
	AuthorisationCode *Min6Max8Text `xml:"AuthstnCd,omitempty"`
}

func (a *AuthorisationResult5) AddAuthorisationEntity() *GenericIdentification70 {
	a.AuthorisationEntity = new(GenericIdentification70)
	return a.AuthorisationEntity
}

func (a *AuthorisationResult5) AddResponseToAuthorisation() *ResponseType1 {
	a.ResponseToAuthorisation = new(ResponseType1)
	return a.ResponseToAuthorisation
}

func (a *AuthorisationResult5) SetAuthorisationCode(value string) {
	a.AuthorisationCode = (*Min6Max8Text)(&value)
}

// Outcome of the authorisation.
type AuthorisationResult6 struct {

	// Type of party that has delivered or declined the card payment authorisation (the party is not identified).
	AuthorisationEntity *GenericIdentification70 `xml:"AuthstnNtty,omitempty"`

	// Response to an authorisation request.
	ResponseToAuthorisation *ResponseType1 `xml:"RspnToAuthstn"`

	// Value assigned by the authorising party.
	AuthorisationCode *Min6Max8Text `xml:"AuthstnCd,omitempty"`

	// Acquirer has requested a contact to the maintenance host.
	TMSTrigger *TMSTrigger1 `xml:"TMSTrggr,omitempty"`
}

func (a *AuthorisationResult6) AddAuthorisationEntity() *GenericIdentification70 {
	a.AuthorisationEntity = new(GenericIdentification70)
	return a.AuthorisationEntity
}

func (a *AuthorisationResult6) AddResponseToAuthorisation() *ResponseType1 {
	a.ResponseToAuthorisation = new(ResponseType1)
	return a.ResponseToAuthorisation
}

func (a *AuthorisationResult6) SetAuthorisationCode(value string) {
	a.AuthorisationCode = (*Min6Max8Text)(&value)
}

func (a *AuthorisationResult6) AddTMSTrigger() *TMSTrigger1 {
	a.TMSTrigger = new(TMSTrigger1)
	return a.TMSTrigger
}

// Outcome of the authorisation.
type AuthorisationResult7 struct {

	// Type of party that has delivered or declined the card payment authorisation (the party is not identified).
	AuthorisationEntity *GenericIdentification75 `xml:"AuthstnNtty,omitempty"`

	// Response to an authorisation request.
	TransactionResponse *ResponseType2 `xml:"TxRspn"`

	// Value assigned by the authorising party.
	AuthorisationCode *Min6Max8Text `xml:"AuthstnCd,omitempty"`

	// Additional information relevant for the destination.
	AdditionalInformation []*ActionMessage3 `xml:"AddtlInf,omitempty"`
}

func (a *AuthorisationResult7) AddAuthorisationEntity() *GenericIdentification75 {
	a.AuthorisationEntity = new(GenericIdentification75)
	return a.AuthorisationEntity
}

func (a *AuthorisationResult7) AddTransactionResponse() *ResponseType2 {
	a.TransactionResponse = new(ResponseType2)
	return a.TransactionResponse
}

func (a *AuthorisationResult7) SetAuthorisationCode(value string) {
	a.AuthorisationCode = (*Min6Max8Text)(&value)
}

func (a *AuthorisationResult7) AddAdditionalInformation() *ActionMessage3 {
	newValue := new(ActionMessage3)
	a.AdditionalInformation = append(a.AdditionalInformation, newValue)
	return newValue
}

// Outcome of the authorisation.
type AuthorisationResult8 struct {

	// Type of party that has delivered or declined the card payment authorisation (the party is not identified).
	AuthorisationEntity *GenericIdentification75 `xml:"AuthstnNtty,omitempty"`

	// Response to an authorisation request.
	TransactionResponse *ResponseType2 `xml:"TxRspn"`

	// Set of actions to be performed by the card acceptor.
	Action []*Action4 `xml:"Actn,omitempty"`

	// Value assigned by the authorising party.
	AuthorisationCode *Min6Max8Text `xml:"AuthstnCd,omitempty"`

	// Additional information relevant for the destination.
	AdditionalInformation []*ActionMessage3 `xml:"AddtlInf,omitempty"`
}

func (a *AuthorisationResult8) AddAuthorisationEntity() *GenericIdentification75 {
	a.AuthorisationEntity = new(GenericIdentification75)
	return a.AuthorisationEntity
}

func (a *AuthorisationResult8) AddTransactionResponse() *ResponseType2 {
	a.TransactionResponse = new(ResponseType2)
	return a.TransactionResponse
}

func (a *AuthorisationResult8) AddAction() *Action4 {
	newValue := new(Action4)
	a.Action = append(a.Action, newValue)
	return newValue
}

func (a *AuthorisationResult8) SetAuthorisationCode(value string) {
	a.AuthorisationCode = (*Min6Max8Text)(&value)
}

func (a *AuthorisationResult8) AddAdditionalInformation() *ActionMessage3 {
	newValue := new(ActionMessage3)
	a.AdditionalInformation = append(a.AdditionalInformation, newValue)
	return newValue
}

// Outcome of the withdrawal authorisation.
type AuthorisationResult9 struct {

	// Type of party that has delivered or declined the card payment authorisation (the party is not identified).
	AuthorisationEntity *PartyType13Code `xml:"AuthstnNtty,omitempty"`

	// Result of the authorisation.
	AuthorisationResponse *ResponseType3 `xml:"AuthstnRspn"`

	// Trace of response by the entities in the path from the issuer to the ATM.
	ResponseTrace []*ResponseType4 `xml:"RspnTrac,omitempty"`

	// Value assigned by the authorising party.
	AuthorisationCode *Min6Max8Text `xml:"AuthstnCd,omitempty"`

	// Sequence of actions to be performed by the ATM to complete the transaction.
	Action []*Action5 `xml:"Actn,omitempty"`
}

func (a *AuthorisationResult9) SetAuthorisationEntity(value string) {
	a.AuthorisationEntity = (*PartyType13Code)(&value)
}

func (a *AuthorisationResult9) AddAuthorisationResponse() *ResponseType3 {
	a.AuthorisationResponse = new(ResponseType3)
	return a.AuthorisationResponse
}

func (a *AuthorisationResult9) AddResponseTrace() *ResponseType4 {
	newValue := new(ResponseType4)
	a.ResponseTrace = append(a.ResponseTrace, newValue)
	return newValue
}

func (a *AuthorisationResult9) SetAuthorisationCode(value string) {
	a.AuthorisationCode = (*Min6Max8Text)(&value)
}

func (a *AuthorisationResult9) AddAction() *Action5 {
	newValue := new(Action5)
	a.Action = append(a.Action, newValue)
	return newValue
}
