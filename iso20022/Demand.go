package iso20022

// Details of the demand.
type Demand1 struct {

	// Unique and unambiguous identifier assigned by the presenting party to the demand.
	Identification *Max35Text `xml:"Id"`

	// Type of demand.
	Type *DemandType1Code `xml:"Tp"`

	// Details related to the undertaking.
	UndertakingIdentification *Undertaking6 `xml:"UdrtkgId"`

	// Details related to the demand amount.
	DemandAmount *UndertakingAmount3 `xml:"DmndAmt"`

	// Unique and unambiguous identifier assigned by the advising party to the undertaking.
	AdvisingPartyReferenceNumber *Max35Text `xml:"AdvsgPtyRefNb,omitempty"`

	// Unique and unambiguous identifier assigned by the second advising party to the undertaking.
	SecondAdvisingPartyReferenceNumber *Max35Text `xml:"ScndAdvsgPtyRefNb,omitempty"`

	// Unique and unambiguous identifier assigned by the confirmer to the undertaking.
	ConfirmerReferenceNumber *Max35Text `xml:"CnfrmrRefNb,omitempty"`

	// Details related to the settlement account.
	SettlementAccount []*CashAccount27 `xml:"SttlmAcct,omitempty"`

	// Details of the beneficiary's presentation of documents.
	PresentationDetails *Presentation2 `xml:"PresntnDtls,omitempty"`

	// Requested new expiry date as an alternative to payment of the demand.
	RequestedExpiryDate *ISODate `xml:"ReqdXpryDt,omitempty"`

	// Document(s) presented for examination.
	DemandDocumentation *DemandDocumentation1 `xml:"DmndDcmnttn,omitempty"`

	// Additional information related to the demand.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (d *Demand1) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *Demand1) SetType(value string) {
	d.Type = (*DemandType1Code)(&value)
}

func (d *Demand1) AddUndertakingIdentification() *Undertaking6 {
	d.UndertakingIdentification = new(Undertaking6)
	return d.UndertakingIdentification
}

func (d *Demand1) AddDemandAmount() *UndertakingAmount3 {
	d.DemandAmount = new(UndertakingAmount3)
	return d.DemandAmount
}

func (d *Demand1) SetAdvisingPartyReferenceNumber(value string) {
	d.AdvisingPartyReferenceNumber = (*Max35Text)(&value)
}

func (d *Demand1) SetSecondAdvisingPartyReferenceNumber(value string) {
	d.SecondAdvisingPartyReferenceNumber = (*Max35Text)(&value)
}

func (d *Demand1) SetConfirmerReferenceNumber(value string) {
	d.ConfirmerReferenceNumber = (*Max35Text)(&value)
}

func (d *Demand1) AddSettlementAccount() *CashAccount27 {
	newValue := new(CashAccount27)
	d.SettlementAccount = append(d.SettlementAccount, newValue)
	return newValue
}

func (d *Demand1) AddPresentationDetails() *Presentation2 {
	d.PresentationDetails = new(Presentation2)
	return d.PresentationDetails
}

func (d *Demand1) SetRequestedExpiryDate(value string) {
	d.RequestedExpiryDate = (*ISODate)(&value)
}

func (d *Demand1) AddDemandDocumentation() *DemandDocumentation1 {
	d.DemandDocumentation = new(DemandDocumentation1)
	return d.DemandDocumentation
}

func (d *Demand1) AddAdditionalInformation(value string) {
	d.AdditionalInformation = append(d.AdditionalInformation, (*Max2000Text)(&value))
}

// Information about the demand.
type Demand2 struct {

	// Unique and unambiguous identifier assigned by the presenting party to the demand.
	Identification *Max35Text `xml:"Id"`

	// Date and time the demand is submitted.
	SubmissionDateTime *ISODateTime `xml:"SubmissnDtTm"`

	// Amount and currency of the demand.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Additional information related to the demand.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (d *Demand2) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *Demand2) SetSubmissionDateTime(value string) {
	d.SubmissionDateTime = (*ISODateTime)(&value)
}

func (d *Demand2) SetAmount(value, currency string) {
	d.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *Demand2) AddAdditionalInformation(value string) {
	d.AdditionalInformation = append(d.AdditionalInformation, (*Max2000Text)(&value))
}

// Details related to the demand.
type Demand3 struct {

	// Unique and unambiguous identifier assigned by the presenting party to the demand.
	Identification *Max35Text `xml:"Id"`

	// Date and time of demand submission.
	SubmissionDateTime *ISODateTime `xml:"SubmissnDtTm"`

	// Amount and currency of the demand.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`
}

func (d *Demand3) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *Demand3) SetSubmissionDateTime(value string) {
	d.SubmissionDateTime = (*ISODateTime)(&value)
}

func (d *Demand3) SetAmount(value, currency string) {
	d.Amount = NewActiveCurrencyAndAmount(value, currency)
}

// Information about a demand.
type Demand4 struct {

	// Unique and unambiguous identifier assigned by the presenting party to the demand.
	Identification *Max35Text `xml:"Id"`

	// Amount and currency of the demand.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`
}

func (d *Demand4) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *Demand4) SetAmount(value, currency string) {
	d.Amount = NewActiveCurrencyAndAmount(value, currency)
}
