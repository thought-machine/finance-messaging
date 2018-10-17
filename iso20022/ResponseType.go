package iso20022

// Response of a requested service.
type ResponseType1 struct {

	// Result of the transaction.
	Response *Response1Code `xml:"Rspn"`

	// Detailed result of the transaction.
	ResponseReason *Max35Text `xml:"RspnRsn,omitempty"`
}

func (r *ResponseType1) SetResponse(value string) {
	r.Response = (*Response1Code)(&value)
}

func (r *ResponseType1) SetResponseReason(value string) {
	r.ResponseReason = (*Max35Text)(&value)
}

// Response of a requested service.
type ResponseType2 struct {

	// Result of the request message or advice message.
	Result *Response3Code `xml:"Rslt"`

	// Detail of the result.
	ResultDetails *ResultDetail1Code `xml:"RsltDtls,omitempty"`

	// Additional information to be logged for further examination.
	AdditionalResultInformation *Max140Text `xml:"AddtlRsltInf,omitempty"`
}

func (r *ResponseType2) SetResult(value string) {
	r.Result = (*Response3Code)(&value)
}

func (r *ResponseType2) SetResultDetails(value string) {
	r.ResultDetails = (*ResultDetail1Code)(&value)
}

func (r *ResponseType2) SetAdditionalResultInformation(value string) {
	r.AdditionalResultInformation = (*Max140Text)(&value)
}

// Response of a requested service.
type ResponseType3 struct {

	// Result of the requested transaction.
	Response *Response4Code `xml:"Rspn"`

	// Detail of the response.
	ResponseReason *ResultDetail2Code `xml:"RspnRsn,omitempty"`

	// Additional information to be logged for further examination.
	AdditionalResponseInformation *Max140Text `xml:"AddtlRspnInf,omitempty"`
}

func (r *ResponseType3) SetResponse(value string) {
	r.Response = (*Response4Code)(&value)
}

func (r *ResponseType3) SetResponseReason(value string) {
	r.ResponseReason = (*ResultDetail2Code)(&value)
}

func (r *ResponseType3) SetAdditionalResponseInformation(value string) {
	r.AdditionalResponseInformation = (*Max140Text)(&value)
}

// Trace of response by the entities in the path from the issuer to the ATM.
type ResponseType4 struct {

	// Identification of the responder.
	ResponderIdentification *Max35Text `xml:"RspndrId"`

	// Codification of the response (for instance ISO 8583, IFX).
	Codification *Max35Text `xml:"Cdfctn,omitempty"`

	// Result of the request withdrawal message.
	Response *Max35Text `xml:"Rspn"`

	// Detail of the response.
	ResponseReason *Max35Text `xml:"RspnRsn,omitempty"`

	// Additional information to be logged for further examination.
	AdditionalResponseInformation *Max35Text `xml:"AddtlRspnInf,omitempty"`
}

func (r *ResponseType4) SetResponderIdentification(value string) {
	r.ResponderIdentification = (*Max35Text)(&value)
}

func (r *ResponseType4) SetCodification(value string) {
	r.Codification = (*Max35Text)(&value)
}

func (r *ResponseType4) SetResponse(value string) {
	r.Response = (*Max35Text)(&value)
}

func (r *ResponseType4) SetResponseReason(value string) {
	r.ResponseReason = (*Max35Text)(&value)
}

func (r *ResponseType4) SetAdditionalResponseInformation(value string) {
	r.AdditionalResponseInformation = (*Max35Text)(&value)
}

// Response of a requested service.
type ResponseType5 struct {

	// Result of the transaction.
	Response *Response4Code `xml:"Rspn"`

	// Detailed result of the transaction.
	ResponseReason *Max35Text `xml:"RspnRsn,omitempty"`

	// Additional information on the response for further examination.
	AdditionalResponseInformation *Max140Text `xml:"AddtlRspnInf,omitempty"`
}

func (r *ResponseType5) SetResponse(value string) {
	r.Response = (*Response4Code)(&value)
}

func (r *ResponseType5) SetResponseReason(value string) {
	r.ResponseReason = (*Max35Text)(&value)
}

func (r *ResponseType5) SetAdditionalResponseInformation(value string) {
	r.AdditionalResponseInformation = (*Max140Text)(&value)
}

// Response of a requested service.
type ResponseType6 struct {

	// Response of the terminal manager.
	Response *Response2Code `xml:"Rspn"`

	// Detail of the response.
	ResponseDetail *ResultDetail3Code `xml:"RspnDtl,omitempty"`

	// Additional information on the response for further examination.
	AdditionalResponse *Max140Text `xml:"AddtlRspn,omitempty"`
}

func (r *ResponseType6) SetResponse(value string) {
	r.Response = (*Response2Code)(&value)
}

func (r *ResponseType6) SetResponseDetail(value string) {
	r.ResponseDetail = (*ResultDetail3Code)(&value)
}

func (r *ResponseType6) SetAdditionalResponse(value string) {
	r.AdditionalResponse = (*Max140Text)(&value)
}

// Response of a requested service.
type ResponseType7 struct {

	// Result of the requested transaction.
	Response *Response4Code `xml:"Rspn"`

	// Detail of the response.
	ResponseReason *ResultDetail4Code `xml:"RspnRsn,omitempty"`

	// Additional information to be logged for further examination.
	AdditionalResponseInformation *Max140Text `xml:"AddtlRspnInf,omitempty"`
}

func (r *ResponseType7) SetResponse(value string) {
	r.Response = (*Response4Code)(&value)
}

func (r *ResponseType7) SetResponseReason(value string) {
	r.ResponseReason = (*ResultDetail4Code)(&value)
}

func (r *ResponseType7) SetAdditionalResponseInformation(value string) {
	r.AdditionalResponseInformation = (*Max140Text)(&value)
}

// Trace of response by the entities in the path from the issuer to the ATM.
type ResponseType8 struct {

	// Identification of the responder.
	ResponderIdentification *Max35Text `xml:"RspndrId"`

	// Codification of the response (for instance ISO 8583, IFX).
	Codification *Max35Text `xml:"Cdfctn,omitempty"`

	// Result of the request.
	Response *Max35Text `xml:"Rspn"`

	// Detail of the response.
	ResponseReason *Max35Text `xml:"RspnRsn,omitempty"`

	// Additional information to be logged for further examination.
	AdditionalResponseInformation *Max35Text `xml:"AddtlRspnInf,omitempty"`
}

func (r *ResponseType8) SetResponderIdentification(value string) {
	r.ResponderIdentification = (*Max35Text)(&value)
}

func (r *ResponseType8) SetCodification(value string) {
	r.Codification = (*Max35Text)(&value)
}

func (r *ResponseType8) SetResponse(value string) {
	r.Response = (*Max35Text)(&value)
}

func (r *ResponseType8) SetResponseReason(value string) {
	r.ResponseReason = (*Max35Text)(&value)
}

func (r *ResponseType8) SetAdditionalResponseInformation(value string) {
	r.AdditionalResponseInformation = (*Max35Text)(&value)
}
