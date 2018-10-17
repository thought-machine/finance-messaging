package iso20022

// Specifies the details relative to the submission of a data set.
type RequiredSubmission2 struct {

	// Specifies with party(ies) is authorised to submit the data set as part of the transaction.
	Submitter []*BICIdentification1 `xml:"Submitr"`
}

func (r *RequiredSubmission2) AddSubmitter() *BICIdentification1 {
	newValue := new(BICIdentification1)
	r.Submitter = append(r.Submitter, newValue)
	return newValue
}

// Specifies the details relative to the submission of the insurance data set.
type RequiredSubmission3 struct {

	// Specifies with party(ies) is authorised to submit the data set as part of the transaction.
	Submitter []*BICIdentification1 `xml:"Submitr"`

	// Specifies if the issuer must be matched as part of the validation of the data set.
	MatchIssuer *PartyIdentification27 `xml:"MtchIssr,omitempty"`

	// Specifies if the issue date must be matched as part of the validation of the data set.
	MatchIssueDate *YesNoIndicator `xml:"MtchIsseDt"`

	// Specifies if the transport information must be matched as part of the validation of the data set.
	MatchTransport *YesNoIndicator `xml:"MtchTrnsprt"`

	// Specifies if the insured amount must be matched as part of the validation of the data set.
	MatchAmount *YesNoIndicator `xml:"MtchAmt"`

	// Specifies which clauses are required in the insurance data set.
	ClausesRequired []*InsuranceClauses1Code `xml:"ClausesReqrd,omitempty"`

	// Specifies if the assured (insured) party must be matched as part of the validation of the data set.
	MatchAssuredParty *AssuredType1Code `xml:"MtchAssrdPty,omitempty"`
}

func (r *RequiredSubmission3) AddSubmitter() *BICIdentification1 {
	newValue := new(BICIdentification1)
	r.Submitter = append(r.Submitter, newValue)
	return newValue
}

func (r *RequiredSubmission3) AddMatchIssuer() *PartyIdentification27 {
	r.MatchIssuer = new(PartyIdentification27)
	return r.MatchIssuer
}

func (r *RequiredSubmission3) SetMatchIssueDate(value string) {
	r.MatchIssueDate = (*YesNoIndicator)(&value)
}

func (r *RequiredSubmission3) SetMatchTransport(value string) {
	r.MatchTransport = (*YesNoIndicator)(&value)
}

func (r *RequiredSubmission3) SetMatchAmount(value string) {
	r.MatchAmount = (*YesNoIndicator)(&value)
}

func (r *RequiredSubmission3) AddClausesRequired(value string) {
	r.ClausesRequired = append(r.ClausesRequired, (*InsuranceClauses1Code)(&value))
}

func (r *RequiredSubmission3) SetMatchAssuredParty(value string) {
	r.MatchAssuredParty = (*AssuredType1Code)(&value)
}

// Specifies the details relative to the submission of the certificate data set.
type RequiredSubmission4 struct {

	// Specifies with party(ies) is authorised to submit the data set as part of the transaction.
	Submitter []*BICIdentification1 `xml:"Submitr"`

	// Specifies the type of the certificate.
	CertificateType *TradeCertificateType1Code `xml:"CertTp"`

	// Specifies if the issuer must be matched as part of the validation of the data set.
	MatchIssuer *PartyIdentification27 `xml:"MtchIssr,omitempty"`

	// Specifies if the issue date must be matched as part of the validation of the data set.
	MatchIssueDate *YesNoIndicator `xml:"MtchIsseDt"`

	// Specifies if the inspection date must be matched as part of the validation of the data set.
	MatchInspectionDate *YesNoIndicator `xml:"MtchInspctnDt"`

	// Specifies if the indication of an authorised inspector must be present as part of the validation of the data set.
	AuthorisedInspectorIndicator *YesNoIndicator `xml:"AuthrsdInspctrInd"`

	// Specifies if the consignee must be matched as part of the validation of the data set.
	MatchConsignee *YesNoIndicator `xml:"MtchConsgn"`

	// Specifies if the manufacturer must be matched as part of the validation of the data set.
	MatchManufacturer *PartyIdentification27 `xml:"MtchManfctr,omitempty"`

	// Specifies if the certificate data set is required in relation to specific line items, and which line items.
	LineItemIdentification []*Max70Text `xml:"LineItmId,omitempty"`
}

func (r *RequiredSubmission4) AddSubmitter() *BICIdentification1 {
	newValue := new(BICIdentification1)
	r.Submitter = append(r.Submitter, newValue)
	return newValue
}

func (r *RequiredSubmission4) SetCertificateType(value string) {
	r.CertificateType = (*TradeCertificateType1Code)(&value)
}

func (r *RequiredSubmission4) AddMatchIssuer() *PartyIdentification27 {
	r.MatchIssuer = new(PartyIdentification27)
	return r.MatchIssuer
}

func (r *RequiredSubmission4) SetMatchIssueDate(value string) {
	r.MatchIssueDate = (*YesNoIndicator)(&value)
}

func (r *RequiredSubmission4) SetMatchInspectionDate(value string) {
	r.MatchInspectionDate = (*YesNoIndicator)(&value)
}

func (r *RequiredSubmission4) SetAuthorisedInspectorIndicator(value string) {
	r.AuthorisedInspectorIndicator = (*YesNoIndicator)(&value)
}

func (r *RequiredSubmission4) SetMatchConsignee(value string) {
	r.MatchConsignee = (*YesNoIndicator)(&value)
}

func (r *RequiredSubmission4) AddMatchManufacturer() *PartyIdentification27 {
	r.MatchManufacturer = new(PartyIdentification27)
	return r.MatchManufacturer
}

func (r *RequiredSubmission4) AddLineItemIdentification(value string) {
	r.LineItemIdentification = append(r.LineItemIdentification, (*Max70Text)(&value))
}

// Specifies the details relative to the submission of the certificate data set.
type RequiredSubmission5 struct {

	// Specifies with party(ies) is authorised to submit the data set as part of the transaction.
	Submitter []*BICIdentification1 `xml:"Submitr"`

	// Specifies the type of the certificate.
	CertificateType *TradeCertificateType2Code `xml:"CertTp"`
}

func (r *RequiredSubmission5) AddSubmitter() *BICIdentification1 {
	newValue := new(BICIdentification1)
	r.Submitter = append(r.Submitter, newValue)
	return newValue
}

func (r *RequiredSubmission5) SetCertificateType(value string) {
	r.CertificateType = (*TradeCertificateType2Code)(&value)
}

// Specifies the details relative to the submission of the certificate data set.
type RequiredSubmission6 struct {

	// Specifies with party(ies) is authorised to submit the data set as part of the transaction.
	Submitter []*BICIdentification1 `xml:"Submitr"`

	// Specifies the type of the certificate, in 4 letters, for example BENE for beneficiary certificate, SHIP for shipping line certifcate.
	CertificateType *Exact4AlphaNumericText `xml:"CertTp"`

	// Description of the certificate type required.
	CertificateTypeDescription *Max140Text `xml:"CertTpDesc"`
}

func (r *RequiredSubmission6) AddSubmitter() *BICIdentification1 {
	newValue := new(BICIdentification1)
	r.Submitter = append(r.Submitter, newValue)
	return newValue
}

func (r *RequiredSubmission6) SetCertificateType(value string) {
	r.CertificateType = (*Exact4AlphaNumericText)(&value)
}

func (r *RequiredSubmission6) SetCertificateTypeDescription(value string) {
	r.CertificateTypeDescription = (*Max140Text)(&value)
}
