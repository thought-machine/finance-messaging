package iso20022

// General information about the corporate action event.
type CorporateActionGeneralInformation10 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Specifies the type of narrative related to the message.
	NarrativeType *CorporateActionNarrative1Choice `xml:"NrrtvTp,omitempty"`
}

func (c *CorporateActionGeneralInformation10) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation10) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation10) AddNarrativeType() *CorporateActionNarrative1Choice {
	c.NarrativeType = new(CorporateActionNarrative1Choice)
	return c.NarrativeType
}

// General information about the corporate action event.
type CorporateActionGeneralInformation100 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *RestrictedFINXMax16Text `xml:"ClssActnNb,omitempty"`

	// Type of processing involved by a corporate action.
	EventProcessingType *CorporateActionEventProcessingType3Choice `xml:"EvtPrcgTp,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType42Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary4Choice `xml:"MndtryVlntryEvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes70 `xml:"UndrlygScty"`
}

func (c *CorporateActionGeneralInformation100) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation100) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation100) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation100) AddEventProcessingType() *CorporateActionEventProcessingType3Choice {
	c.EventProcessingType = new(CorporateActionEventProcessingType3Choice)
	return c.EventProcessingType
}

func (c *CorporateActionGeneralInformation100) AddEventType() *CorporateActionEventType42Choice {
	c.EventType = new(CorporateActionEventType42Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation100) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary4Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary4Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation100) AddUnderlyingSecurity() *FinancialInstrumentAttributes70 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes70)
	return c.UnderlyingSecurity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation102 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Specifies the type of narrative related to the message.
	NarrativeType *CorporateActionNarrative4Choice `xml:"NrrtvTp,omitempty"`
}

func (c *CorporateActionGeneralInformation102) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation102) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation102) AddNarrativeType() *CorporateActionNarrative4Choice {
	c.NarrativeType = new(CorporateActionNarrative4Choice)
	return c.NarrativeType
}

// General information about the corporate action event.
type CorporateActionGeneralInformation103 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *RestrictedFINXMax16Text `xml:"ClssActnNb,omitempty"`

	// Type of processing involved by a corporate action.
	EventProcessingType *CorporateActionEventProcessingType3Choice `xml:"EvtPrcgTp,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType35Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary4Choice `xml:"MndtryVlntryEvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes70 `xml:"UndrlygScty"`
}

func (c *CorporateActionGeneralInformation103) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation103) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation103) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation103) AddEventProcessingType() *CorporateActionEventProcessingType3Choice {
	c.EventProcessingType = new(CorporateActionEventProcessingType3Choice)
	return c.EventProcessingType
}

func (c *CorporateActionGeneralInformation103) AddEventType() *CorporateActionEventType35Choice {
	c.EventType = new(CorporateActionEventType35Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation103) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary4Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary4Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation103) AddUnderlyingSecurity() *FinancialInstrumentAttributes70 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes70)
	return c.UnderlyingSecurity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation104 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType50Choice `xml:"EvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification21 `xml:"FinInstrmId,omitempty"`
}

func (c *CorporateActionGeneralInformation104) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation104) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation104) AddEventType() *CorporateActionEventType50Choice {
	c.EventType = new(CorporateActionEventType50Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation104) AddFinancialInstrumentIdentification() *SecurityIdentification21 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification21)
	return c.FinancialInstrumentIdentification
}

// General information about the corporate action event.
type CorporateActionGeneralInformation105 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of processing involved by a corporate action.
	EventProcessingType *CorporateActionEventProcessingType2Choice `xml:"EvtPrcgTp,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType51Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary3Choice `xml:"MndtryVlntryEvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes79 `xml:"UndrlygScty"`
}

func (c *CorporateActionGeneralInformation105) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation105) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation105) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation105) AddEventProcessingType() *CorporateActionEventProcessingType2Choice {
	c.EventProcessingType = new(CorporateActionEventProcessingType2Choice)
	return c.EventProcessingType
}

func (c *CorporateActionGeneralInformation105) AddEventType() *CorporateActionEventType51Choice {
	c.EventType = new(CorporateActionEventType51Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation105) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary3Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary3Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation105) AddUnderlyingSecurity() *FinancialInstrumentAttributes79 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes79)
	return c.UnderlyingSecurity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation106 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of processing involved by a corporate action.
	EventProcessingType *CorporateActionEventProcessingType2Choice `xml:"EvtPrcgTp,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType54Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary3Choice `xml:"MndtryVlntryEvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes79 `xml:"UndrlygScty"`
}

func (c *CorporateActionGeneralInformation106) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation106) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation106) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation106) AddEventProcessingType() *CorporateActionEventProcessingType2Choice {
	c.EventProcessingType = new(CorporateActionEventProcessingType2Choice)
	return c.EventProcessingType
}

func (c *CorporateActionGeneralInformation106) AddEventType() *CorporateActionEventType54Choice {
	c.EventType = new(CorporateActionEventType54Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation106) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary3Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary3Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation106) AddUnderlyingSecurity() *FinancialInstrumentAttributes79 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes79)
	return c.UnderlyingSecurity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation107 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType52Choice `xml:"EvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes81 `xml:"UndrlygScty,omitempty"`
}

func (c *CorporateActionGeneralInformation107) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation107) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation107) AddEventType() *CorporateActionEventType52Choice {
	c.EventType = new(CorporateActionEventType52Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation107) AddUnderlyingSecurity() *FinancialInstrumentAttributes81 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes81)
	return c.UnderlyingSecurity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation108 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType51Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary3Choice `xml:"MndtryVlntryEvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`
}

func (c *CorporateActionGeneralInformation108) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation108) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation108) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation108) AddEventType() *CorporateActionEventType51Choice {
	c.EventType = new(CorporateActionEventType51Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation108) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary3Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary3Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation108) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return c.FinancialInstrumentIdentification
}

// General information about the corporate action event.
type CorporateActionGeneralInformation109 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType52Choice `xml:"EvtTp"`
}

func (c *CorporateActionGeneralInformation109) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation109) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation109) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation109) AddEventType() *CorporateActionEventType52Choice {
	c.EventType = new(CorporateActionEventType52Choice)
	return c.EventType
}

// General information about the corporate action event.
type CorporateActionGeneralInformation11 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of processing involved by a corporate action.
	EventProcessingType *CorporateActionEventProcessingTypeChoice `xml:"EvtPrcgTp,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType3Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary1Choice `xml:"MndtryVlntryEvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes7 `xml:"UndrlygScty"`
}

func (c *CorporateActionGeneralInformation11) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation11) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation11) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation11) AddEventProcessingType() *CorporateActionEventProcessingTypeChoice {
	c.EventProcessingType = new(CorporateActionEventProcessingTypeChoice)
	return c.EventProcessingType
}

func (c *CorporateActionGeneralInformation11) AddEventType() *CorporateActionEventType3Choice {
	c.EventType = new(CorporateActionEventType3Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation11) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary1Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary1Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation11) AddUnderlyingSecurity() *FinancialInstrumentAttributes7 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes7)
	return c.UnderlyingSecurity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation110 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType52Choice `xml:"EvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId,omitempty"`
}

func (c *CorporateActionGeneralInformation110) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation110) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation110) AddEventType() *CorporateActionEventType52Choice {
	c.EventType = new(CorporateActionEventType52Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation110) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return c.FinancialInstrumentIdentification
}

// General information about the corporate action event.
type CorporateActionGeneralInformation111 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType53Choice `xml:"EvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat16Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Fractional quantity resulting from an event which will be paid with cash in lieu.
	FractionalQuantity *FinancialInstrumentQuantity1Choice `xml:"FrctnlQty,omitempty"`
}

func (c *CorporateActionGeneralInformation111) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation111) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation111) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation111) AddEventType() *CorporateActionEventType53Choice {
	c.EventType = new(CorporateActionEventType53Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation111) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return c.FinancialInstrumentIdentification
}

func (c *CorporateActionGeneralInformation111) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat16Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat16Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateActionGeneralInformation111) AddFractionalQuantity() *FinancialInstrumentQuantity1Choice {
	c.FractionalQuantity = new(FinancialInstrumentQuantity1Choice)
	return c.FractionalQuantity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation112 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType54Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary3Choice `xml:"MndtryVlntryEvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`
}

func (c *CorporateActionGeneralInformation112) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation112) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation112) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation112) AddEventType() *CorporateActionEventType54Choice {
	c.EventType = new(CorporateActionEventType54Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation112) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary3Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary3Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation112) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return c.FinancialInstrumentIdentification
}

// General information about the corporate action event.
type CorporateActionGeneralInformation113 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *RestrictedFINXMax16Text `xml:"ClssActnNb,omitempty"`

	// Type of processing involved by a corporate action.
	EventProcessingType *CorporateActionEventProcessingType3Choice `xml:"EvtPrcgTp,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType57Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary4Choice `xml:"MndtryVlntryEvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes85 `xml:"UndrlygScty"`
}

func (c *CorporateActionGeneralInformation113) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation113) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation113) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation113) AddEventProcessingType() *CorporateActionEventProcessingType3Choice {
	c.EventProcessingType = new(CorporateActionEventProcessingType3Choice)
	return c.EventProcessingType
}

func (c *CorporateActionGeneralInformation113) AddEventType() *CorporateActionEventType57Choice {
	c.EventType = new(CorporateActionEventType57Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation113) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary4Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary4Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation113) AddUnderlyingSecurity() *FinancialInstrumentAttributes85 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes85)
	return c.UnderlyingSecurity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation114 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *RestrictedFINXMax16Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType58Choice `xml:"EvtTp"`
}

func (c *CorporateActionGeneralInformation114) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation114) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation114) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation114) AddEventType() *CorporateActionEventType58Choice {
	c.EventType = new(CorporateActionEventType58Choice)
	return c.EventType
}

// General information about the corporate action event.
type CorporateActionGeneralInformation115 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType58Choice `xml:"EvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes84 `xml:"UndrlygScty,omitempty"`
}

func (c *CorporateActionGeneralInformation115) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation115) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation115) AddEventType() *CorporateActionEventType58Choice {
	c.EventType = new(CorporateActionEventType58Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation115) AddUnderlyingSecurity() *FinancialInstrumentAttributes84 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes84)
	return c.UnderlyingSecurity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation117 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *RestrictedFINXMax16Text `xml:"ClssActnNb,omitempty"`

	// Type of processing involved by a corporate action.
	EventProcessingType *CorporateActionEventProcessingType3Choice `xml:"EvtPrcgTp,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType61Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary4Choice `xml:"MndtryVlntryEvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes85 `xml:"UndrlygScty"`
}

func (c *CorporateActionGeneralInformation117) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation117) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation117) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation117) AddEventProcessingType() *CorporateActionEventProcessingType3Choice {
	c.EventProcessingType = new(CorporateActionEventProcessingType3Choice)
	return c.EventProcessingType
}

func (c *CorporateActionGeneralInformation117) AddEventType() *CorporateActionEventType61Choice {
	c.EventType = new(CorporateActionEventType61Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation117) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary4Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary4Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation117) AddUnderlyingSecurity() *FinancialInstrumentAttributes85 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes85)
	return c.UnderlyingSecurity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation118 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *RestrictedFINXMax16Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType62Choice `xml:"EvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat17Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Fractional quantity resulting from an event which will be paid with cash in lieu.
	FractionalQuantity *FinancialInstrumentQuantity15Choice `xml:"FrctnlQty,omitempty"`
}

func (c *CorporateActionGeneralInformation118) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation118) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation118) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation118) AddEventType() *CorporateActionEventType62Choice {
	c.EventType = new(CorporateActionEventType62Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation118) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return c.FinancialInstrumentIdentification
}

func (c *CorporateActionGeneralInformation118) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat17Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat17Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateActionGeneralInformation118) AddFractionalQuantity() *FinancialInstrumentQuantity15Choice {
	c.FractionalQuantity = new(FinancialInstrumentQuantity15Choice)
	return c.FractionalQuantity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation120 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *RestrictedFINXMax16Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType57Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary4Choice `xml:"MndtryVlntryEvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`
}

func (c *CorporateActionGeneralInformation120) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation120) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation120) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation120) AddEventType() *CorporateActionEventType57Choice {
	c.EventType = new(CorporateActionEventType57Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation120) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary4Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary4Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation120) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return c.FinancialInstrumentIdentification
}

// General information about the corporate action event.
type CorporateActionGeneralInformation121 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType58Choice `xml:"EvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification21 `xml:"FinInstrmId,omitempty"`
}

func (c *CorporateActionGeneralInformation121) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation121) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation121) AddEventType() *CorporateActionEventType58Choice {
	c.EventType = new(CorporateActionEventType58Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation121) AddFinancialInstrumentIdentification() *SecurityIdentification21 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification21)
	return c.FinancialInstrumentIdentification
}

// General information about the corporate action event.
type CorporateActionGeneralInformation123 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *RestrictedFINXMax16Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType61Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary4Choice `xml:"MndtryVlntryEvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`
}

func (c *CorporateActionGeneralInformation123) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation123) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation123) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation123) AddEventType() *CorporateActionEventType61Choice {
	c.EventType = new(CorporateActionEventType61Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation123) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary4Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary4Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation123) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return c.FinancialInstrumentIdentification
}

// General information about the corporate action event.
type CorporateActionGeneralInformation21 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType3Choice `xml:"EvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes18 `xml:"UndrlygScty,omitempty"`
}

func (c *CorporateActionGeneralInformation21) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation21) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation21) AddEventType() *CorporateActionEventType3Choice {
	c.EventType = new(CorporateActionEventType3Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation21) AddUnderlyingSecurity() *FinancialInstrumentAttributes18 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes18)
	return c.UnderlyingSecurity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation22 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of processing involved by a corporate action.
	EventProcessingType *CorporateActionEventProcessingTypeChoice `xml:"EvtPrcgTp,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType3Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary1Choice `xml:"MndtryVlntryEvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes19 `xml:"UndrlygScty"`
}

func (c *CorporateActionGeneralInformation22) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation22) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation22) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation22) AddEventProcessingType() *CorporateActionEventProcessingTypeChoice {
	c.EventProcessingType = new(CorporateActionEventProcessingTypeChoice)
	return c.EventProcessingType
}

func (c *CorporateActionGeneralInformation22) AddEventType() *CorporateActionEventType3Choice {
	c.EventType = new(CorporateActionEventType3Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation22) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary1Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary1Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation22) AddUnderlyingSecurity() *FinancialInstrumentAttributes19 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes19)
	return c.UnderlyingSecurity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation23 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of processing involved by a corporate action.
	EventProcessingType *CorporateActionEventProcessingTypeChoice `xml:"EvtPrcgTp,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType3Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary1Choice `xml:"MndtryVlntryEvtTp"`

	// Indicates whether the message is related to a claim on the associated corporate action event.
	AdditionalBusinessProcessIndicator *AdditionalBusinessProcessFormat2Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes19 `xml:"UndrlygScty"`
}

func (c *CorporateActionGeneralInformation23) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation23) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation23) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation23) AddEventProcessingType() *CorporateActionEventProcessingTypeChoice {
	c.EventProcessingType = new(CorporateActionEventProcessingTypeChoice)
	return c.EventProcessingType
}

func (c *CorporateActionGeneralInformation23) AddEventType() *CorporateActionEventType3Choice {
	c.EventType = new(CorporateActionEventType3Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation23) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary1Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary1Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation23) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat2Choice {
	c.AdditionalBusinessProcessIndicator = new(AdditionalBusinessProcessFormat2Choice)
	return c.AdditionalBusinessProcessIndicator
}

func (c *CorporateActionGeneralInformation23) AddUnderlyingSecurity() *FinancialInstrumentAttributes19 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes19)
	return c.UnderlyingSecurity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation24 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType5Choice `xml:"EvtTp"`

	// Identification of the security concerned by the corporate action.
	UnderlyingSecurityIdentification *SecurityIdentification14 `xml:"UndrlygSctyId"`

	// Indicates that the additional business process relates to a claim on the associated corporate action event.
	AdditionalBusinessProcessIndicator *AdditionalBusinessProcessFormat3Choice `xml:"AddtlBizPrcInd,omitempty"`
}

func (c *CorporateActionGeneralInformation24) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation24) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation24) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation24) AddEventType() *CorporateActionEventType5Choice {
	c.EventType = new(CorporateActionEventType5Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation24) AddUnderlyingSecurityIdentification() *SecurityIdentification14 {
	c.UnderlyingSecurityIdentification = new(SecurityIdentification14)
	return c.UnderlyingSecurityIdentification
}

func (c *CorporateActionGeneralInformation24) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat3Choice {
	c.AdditionalBusinessProcessIndicator = new(AdditionalBusinessProcessFormat3Choice)
	return c.AdditionalBusinessProcessIndicator
}

// General information about the corporate action event.
type CorporateActionGeneralInformation25 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType3Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary1Choice `xml:"MndtryVlntryEvtTp"`

	// Identification of the security concerned by the corporate action.
	UnderlyingSecurityIdentification *SecurityIdentification14 `xml:"UndrlygSctyId"`
}

func (c *CorporateActionGeneralInformation25) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation25) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation25) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation25) AddEventType() *CorporateActionEventType3Choice {
	c.EventType = new(CorporateActionEventType3Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation25) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary1Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary1Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation25) AddUnderlyingSecurityIdentification() *SecurityIdentification14 {
	c.UnderlyingSecurityIdentification = new(SecurityIdentification14)
	return c.UnderlyingSecurityIdentification
}

// General information about the corporate action event.
type CorporateActionGeneralInformation26 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType3Choice `xml:"EvtTp"`

	// Identification of the security concerned by the corporate action.
	UnderlyingSecurityIdentification *SecurityIdentification14 `xml:"UndrlygSctyId,omitempty"`
}

func (c *CorporateActionGeneralInformation26) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation26) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation26) AddEventType() *CorporateActionEventType3Choice {
	c.EventType = new(CorporateActionEventType3Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation26) AddUnderlyingSecurityIdentification() *SecurityIdentification14 {
	c.UnderlyingSecurityIdentification = new(SecurityIdentification14)
	return c.UnderlyingSecurityIdentification
}

// General information about the corporate action event.
type CorporateActionGeneralInformation3 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of processing involved by a corporate action.
	EventProcessingType *CorporateActionEventProcessingTypeChoice `xml:"EvtPrcgTp,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType3Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary1Choice `xml:"MndtryVlntryEvtTp"`

	// Indicates whether the message is related to a claim on the associated corporate action event.
	AdditionalBusinessProcessIndicator *AdditionalBusinessProcessFormat2Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes7 `xml:"UndrlygScty"`
}

func (c *CorporateActionGeneralInformation3) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation3) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation3) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation3) AddEventProcessingType() *CorporateActionEventProcessingTypeChoice {
	c.EventProcessingType = new(CorporateActionEventProcessingTypeChoice)
	return c.EventProcessingType
}

func (c *CorporateActionGeneralInformation3) AddEventType() *CorporateActionEventType3Choice {
	c.EventType = new(CorporateActionEventType3Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation3) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary1Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary1Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation3) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat2Choice {
	c.AdditionalBusinessProcessIndicator = new(AdditionalBusinessProcessFormat2Choice)
	return c.AdditionalBusinessProcessIndicator
}

func (c *CorporateActionGeneralInformation3) AddUnderlyingSecurity() *FinancialInstrumentAttributes7 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes7)
	return c.UnderlyingSecurity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation33 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType7Choice `xml:"EvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId,omitempty"`
}

func (c *CorporateActionGeneralInformation33) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation33) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation33) AddEventType() *CorporateActionEventType7Choice {
	c.EventType = new(CorporateActionEventType7Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation33) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return c.FinancialInstrumentIdentification
}

// General information about the corporate action event.
type CorporateActionGeneralInformation34 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType7Choice `xml:"EvtTp"`
}

func (c *CorporateActionGeneralInformation34) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation34) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation34) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation34) AddEventType() *CorporateActionEventType7Choice {
	c.EventType = new(CorporateActionEventType7Choice)
	return c.EventType
}

// General information about the corporate action event.
type CorporateActionGeneralInformation35 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType7Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary1Choice `xml:"MndtryVlntryEvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`
}

func (c *CorporateActionGeneralInformation35) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation35) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation35) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation35) AddEventType() *CorporateActionEventType7Choice {
	c.EventType = new(CorporateActionEventType7Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation35) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary1Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary1Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation35) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return c.FinancialInstrumentIdentification
}

// General information about the corporate action event.
type CorporateActionGeneralInformation36 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType7Choice `xml:"EvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes32 `xml:"UndrlygScty,omitempty"`
}

func (c *CorporateActionGeneralInformation36) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation36) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation36) AddEventType() *CorporateActionEventType7Choice {
	c.EventType = new(CorporateActionEventType7Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation36) AddUnderlyingSecurity() *FinancialInstrumentAttributes32 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes32)
	return c.UnderlyingSecurity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation37 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of processing involved by a corporate action.
	EventProcessingType *CorporateActionEventProcessingTypeChoice `xml:"EvtPrcgTp,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType7Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary1Choice `xml:"MndtryVlntryEvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes33 `xml:"UndrlygScty"`
}

func (c *CorporateActionGeneralInformation37) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation37) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation37) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation37) AddEventProcessingType() *CorporateActionEventProcessingTypeChoice {
	c.EventProcessingType = new(CorporateActionEventProcessingTypeChoice)
	return c.EventProcessingType
}

func (c *CorporateActionGeneralInformation37) AddEventType() *CorporateActionEventType7Choice {
	c.EventType = new(CorporateActionEventType7Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation37) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary1Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary1Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation37) AddUnderlyingSecurity() *FinancialInstrumentAttributes33 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes33)
	return c.UnderlyingSecurity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation38 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of processing involved by a corporate action.
	EventProcessingType *CorporateActionEventProcessingTypeChoice `xml:"EvtPrcgTp,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType7Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary1Choice `xml:"MndtryVlntryEvtTp"`

	// Indicates whether the message is related to a claim on the associated corporate action event.
	AdditionalBusinessProcessIndicator *AdditionalBusinessProcessFormat2Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes33 `xml:"UndrlygScty"`
}

func (c *CorporateActionGeneralInformation38) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation38) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation38) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation38) AddEventProcessingType() *CorporateActionEventProcessingTypeChoice {
	c.EventProcessingType = new(CorporateActionEventProcessingTypeChoice)
	return c.EventProcessingType
}

func (c *CorporateActionGeneralInformation38) AddEventType() *CorporateActionEventType7Choice {
	c.EventType = new(CorporateActionEventType7Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation38) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary1Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary1Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation38) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat2Choice {
	c.AdditionalBusinessProcessIndicator = new(AdditionalBusinessProcessFormat2Choice)
	return c.AdditionalBusinessProcessIndicator
}

func (c *CorporateActionGeneralInformation38) AddUnderlyingSecurity() *FinancialInstrumentAttributes33 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes33)
	return c.UnderlyingSecurity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation39 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType8Choice `xml:"EvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Indicates that the additional business process relates to a claim on the associated corporate action event.
	AdditionalBusinessProcessIndicator *AdditionalBusinessProcessFormat7Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat6Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`
}

func (c *CorporateActionGeneralInformation39) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation39) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation39) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation39) AddEventType() *CorporateActionEventType8Choice {
	c.EventType = new(CorporateActionEventType8Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation39) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return c.FinancialInstrumentIdentification
}

func (c *CorporateActionGeneralInformation39) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat7Choice {
	c.AdditionalBusinessProcessIndicator = new(AdditionalBusinessProcessFormat7Choice)
	return c.AdditionalBusinessProcessIndicator
}

func (c *CorporateActionGeneralInformation39) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat6Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat6Choice)
	return c.IntermediateSecuritiesDistributionType
}

// General information about the corporate action event.
type CorporateActionGeneralInformation4 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType3Choice `xml:"EvtTp"`

	// Identification of the security concerned by the corporate action.
	UnderlyingSecurityIdentification *SecurityIdentification11 `xml:"UndrlygSctyId"`

	// Indicates that the additional business process relates to a claim on the associated corporate action event.
	AdditionalBusinessProcessIndicator *AdditionalBusinessProcessFormat3Choice `xml:"AddtlBizPrcInd,omitempty"`
}

func (c *CorporateActionGeneralInformation4) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation4) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation4) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation4) AddEventType() *CorporateActionEventType3Choice {
	c.EventType = new(CorporateActionEventType3Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation4) AddUnderlyingSecurityIdentification() *SecurityIdentification11 {
	c.UnderlyingSecurityIdentification = new(SecurityIdentification11)
	return c.UnderlyingSecurityIdentification
}

func (c *CorporateActionGeneralInformation4) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat3Choice {
	c.AdditionalBusinessProcessIndicator = new(AdditionalBusinessProcessFormat3Choice)
	return c.AdditionalBusinessProcessIndicator
}

// General information about the corporate action event.
type CorporateActionGeneralInformation40 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Specifies the type of narrative related to the message.
	NarrativeType *CorporateActionNarrative1Choice `xml:"NrrtvTp,omitempty"`
}

func (c *CorporateActionGeneralInformation40) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation40) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation40) AddNarrativeType() *CorporateActionNarrative1Choice {
	c.NarrativeType = new(CorporateActionNarrative1Choice)
	return c.NarrativeType
}

// General information about the corporate action event.
type CorporateActionGeneralInformation49 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType11Choice `xml:"EvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId,omitempty"`
}

func (c *CorporateActionGeneralInformation49) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation49) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation49) AddEventType() *CorporateActionEventType11Choice {
	c.EventType = new(CorporateActionEventType11Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation49) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return c.FinancialInstrumentIdentification
}

// General information about the corporate action event.
type CorporateActionGeneralInformation50 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType12Choice `xml:"EvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Indicates that the additional business process relates to a claim on the associated corporate action event.
	AdditionalBusinessProcessIndicator *AdditionalBusinessProcessFormat7Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat6Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Fractional quantity resulting from an event which will be paid with cash in lieu.
	FractionalQuantity *FinancialInstrumentQuantity1Choice `xml:"FrctnlQty,omitempty"`
}

func (c *CorporateActionGeneralInformation50) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation50) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation50) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation50) AddEventType() *CorporateActionEventType12Choice {
	c.EventType = new(CorporateActionEventType12Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation50) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return c.FinancialInstrumentIdentification
}

func (c *CorporateActionGeneralInformation50) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat7Choice {
	c.AdditionalBusinessProcessIndicator = new(AdditionalBusinessProcessFormat7Choice)
	return c.AdditionalBusinessProcessIndicator
}

func (c *CorporateActionGeneralInformation50) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat6Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat6Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateActionGeneralInformation50) AddFractionalQuantity() *FinancialInstrumentQuantity1Choice {
	c.FractionalQuantity = new(FinancialInstrumentQuantity1Choice)
	return c.FractionalQuantity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation51 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of processing involved by a corporate action.
	EventProcessingType *CorporateActionEventProcessingTypeChoice `xml:"EvtPrcgTp,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType13Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary1Choice `xml:"MndtryVlntryEvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes43 `xml:"UndrlygScty"`
}

func (c *CorporateActionGeneralInformation51) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation51) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation51) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation51) AddEventProcessingType() *CorporateActionEventProcessingTypeChoice {
	c.EventProcessingType = new(CorporateActionEventProcessingTypeChoice)
	return c.EventProcessingType
}

func (c *CorporateActionGeneralInformation51) AddEventType() *CorporateActionEventType13Choice {
	c.EventType = new(CorporateActionEventType13Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation51) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary1Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary1Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation51) AddUnderlyingSecurity() *FinancialInstrumentAttributes43 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes43)
	return c.UnderlyingSecurity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation52 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType11Choice `xml:"EvtTp"`
}

func (c *CorporateActionGeneralInformation52) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation52) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation52) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation52) AddEventType() *CorporateActionEventType11Choice {
	c.EventType = new(CorporateActionEventType11Choice)
	return c.EventType
}

// General information about the corporate action event.
type CorporateActionGeneralInformation53 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType14Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary1Choice `xml:"MndtryVlntryEvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`
}

func (c *CorporateActionGeneralInformation53) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation53) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation53) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation53) AddEventType() *CorporateActionEventType14Choice {
	c.EventType = new(CorporateActionEventType14Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation53) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary1Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary1Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation53) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return c.FinancialInstrumentIdentification
}

// General information about the corporate action event.
type CorporateActionGeneralInformation54 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of processing involved by a corporate action.
	EventProcessingType *CorporateActionEventProcessingTypeChoice `xml:"EvtPrcgTp,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType14Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary1Choice `xml:"MndtryVlntryEvtTp"`

	// Indicates whether the message is related to a claim on the associated corporate action event.
	AdditionalBusinessProcessIndicator *AdditionalBusinessProcessFormat2Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes43 `xml:"UndrlygScty"`
}

func (c *CorporateActionGeneralInformation54) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation54) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation54) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation54) AddEventProcessingType() *CorporateActionEventProcessingTypeChoice {
	c.EventProcessingType = new(CorporateActionEventProcessingTypeChoice)
	return c.EventProcessingType
}

func (c *CorporateActionGeneralInformation54) AddEventType() *CorporateActionEventType14Choice {
	c.EventType = new(CorporateActionEventType14Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation54) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary1Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary1Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation54) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat2Choice {
	c.AdditionalBusinessProcessIndicator = new(AdditionalBusinessProcessFormat2Choice)
	return c.AdditionalBusinessProcessIndicator
}

func (c *CorporateActionGeneralInformation54) AddUnderlyingSecurity() *FinancialInstrumentAttributes43 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes43)
	return c.UnderlyingSecurity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation55 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType11Choice `xml:"EvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes32 `xml:"UndrlygScty,omitempty"`
}

func (c *CorporateActionGeneralInformation55) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation55) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation55) AddEventType() *CorporateActionEventType11Choice {
	c.EventType = new(CorporateActionEventType11Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation55) AddUnderlyingSecurity() *FinancialInstrumentAttributes32 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes32)
	return c.UnderlyingSecurity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation56 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType13Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary1Choice `xml:"MndtryVlntryEvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`
}

func (c *CorporateActionGeneralInformation56) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation56) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation56) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation56) AddEventType() *CorporateActionEventType13Choice {
	c.EventType = new(CorporateActionEventType13Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation56) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary1Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary1Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation56) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return c.FinancialInstrumentIdentification
}

// General information about the corporate action event.
type CorporateActionGeneralInformation6 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType3Choice `xml:"EvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes6 `xml:"UndrlygScty,omitempty"`
}

func (c *CorporateActionGeneralInformation6) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation6) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation6) AddEventType() *CorporateActionEventType3Choice {
	c.EventType = new(CorporateActionEventType3Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation6) AddUnderlyingSecurity() *FinancialInstrumentAttributes6 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes6)
	return c.UnderlyingSecurity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation69 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of processing involved by a corporate action.
	EventProcessingType *CorporateActionEventProcessingTypeChoice `xml:"EvtPrcgTp,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType14Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary1Choice `xml:"MndtryVlntryEvtTp"`

	// Indicates whether the message is related to a claim on the associated corporate action event.
	AdditionalBusinessProcessIndicator *AdditionalBusinessProcessFormat2Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes48 `xml:"UndrlygScty"`
}

func (c *CorporateActionGeneralInformation69) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation69) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation69) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation69) AddEventProcessingType() *CorporateActionEventProcessingTypeChoice {
	c.EventProcessingType = new(CorporateActionEventProcessingTypeChoice)
	return c.EventProcessingType
}

func (c *CorporateActionGeneralInformation69) AddEventType() *CorporateActionEventType14Choice {
	c.EventType = new(CorporateActionEventType14Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation69) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary1Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary1Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation69) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat2Choice {
	c.AdditionalBusinessProcessIndicator = new(AdditionalBusinessProcessFormat2Choice)
	return c.AdditionalBusinessProcessIndicator
}

func (c *CorporateActionGeneralInformation69) AddUnderlyingSecurity() *FinancialInstrumentAttributes48 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes48)
	return c.UnderlyingSecurity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation7 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType3Choice `xml:"EvtTp"`

	// Identification of the security concerned by the corporate action.
	UnderlyingSecurityIdentification *SecurityIdentification11 `xml:"UndrlygSctyId,omitempty"`
}

func (c *CorporateActionGeneralInformation7) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation7) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation7) AddEventType() *CorporateActionEventType3Choice {
	c.EventType = new(CorporateActionEventType3Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation7) AddUnderlyingSecurityIdentification() *SecurityIdentification11 {
	c.UnderlyingSecurityIdentification = new(SecurityIdentification11)
	return c.UnderlyingSecurityIdentification
}

// General information about the corporate action event.
type CorporateActionGeneralInformation70 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of processing involved by a corporate action.
	EventProcessingType *CorporateActionEventProcessingTypeChoice `xml:"EvtPrcgTp,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType13Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary1Choice `xml:"MndtryVlntryEvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes48 `xml:"UndrlygScty"`
}

func (c *CorporateActionGeneralInformation70) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation70) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation70) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation70) AddEventProcessingType() *CorporateActionEventProcessingTypeChoice {
	c.EventProcessingType = new(CorporateActionEventProcessingTypeChoice)
	return c.EventProcessingType
}

func (c *CorporateActionGeneralInformation70) AddEventType() *CorporateActionEventType13Choice {
	c.EventType = new(CorporateActionEventType13Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation70) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary1Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary1Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation70) AddUnderlyingSecurity() *FinancialInstrumentAttributes48 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes48)
	return c.UnderlyingSecurity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation71 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType11Choice `xml:"EvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes50 `xml:"UndrlygScty,omitempty"`
}

func (c *CorporateActionGeneralInformation71) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation71) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation71) AddEventType() *CorporateActionEventType11Choice {
	c.EventType = new(CorporateActionEventType11Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation71) AddUnderlyingSecurity() *FinancialInstrumentAttributes50 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes50)
	return c.UnderlyingSecurity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation79 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType12Choice `xml:"EvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Indicates that the additional business process relates to a claim on the associated corporate action event.
	AdditionalBusinessProcessIndicator *AdditionalBusinessProcessFormat7Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat5Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Fractional quantity resulting from an event which will be paid with cash in lieu.
	FractionalQuantity *FinancialInstrumentQuantity1Choice `xml:"FrctnlQty,omitempty"`
}

func (c *CorporateActionGeneralInformation79) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation79) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation79) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation79) AddEventType() *CorporateActionEventType12Choice {
	c.EventType = new(CorporateActionEventType12Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation79) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return c.FinancialInstrumentIdentification
}

func (c *CorporateActionGeneralInformation79) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat7Choice {
	c.AdditionalBusinessProcessIndicator = new(AdditionalBusinessProcessFormat7Choice)
	return c.AdditionalBusinessProcessIndicator
}

func (c *CorporateActionGeneralInformation79) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat5Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat5Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateActionGeneralInformation79) AddFractionalQuantity() *FinancialInstrumentQuantity1Choice {
	c.FractionalQuantity = new(FinancialInstrumentQuantity1Choice)
	return c.FractionalQuantity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation8 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType3Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary1Choice `xml:"MndtryVlntryEvtTp"`

	// Identification of the security concerned by the corporate action.
	UnderlyingSecurityIdentification *SecurityIdentification11 `xml:"UndrlygSctyId"`
}

func (c *CorporateActionGeneralInformation8) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation8) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation8) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation8) AddEventType() *CorporateActionEventType3Choice {
	c.EventType = new(CorporateActionEventType3Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation8) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary1Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary1Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation8) AddUnderlyingSecurityIdentification() *SecurityIdentification11 {
	c.UnderlyingSecurityIdentification = new(SecurityIdentification11)
	return c.UnderlyingSecurityIdentification
}

// General information about the corporate action event.
type CorporateActionGeneralInformation84 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of processing involved by a corporate action.
	EventProcessingType *CorporateActionEventProcessingType2Choice `xml:"EvtPrcgTp,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType34Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary3Choice `xml:"MndtryVlntryEvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes66 `xml:"UndrlygScty"`
}

func (c *CorporateActionGeneralInformation84) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation84) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation84) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation84) AddEventProcessingType() *CorporateActionEventProcessingType2Choice {
	c.EventProcessingType = new(CorporateActionEventProcessingType2Choice)
	return c.EventProcessingType
}

func (c *CorporateActionGeneralInformation84) AddEventType() *CorporateActionEventType34Choice {
	c.EventType = new(CorporateActionEventType34Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation84) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary3Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary3Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation84) AddUnderlyingSecurity() *FinancialInstrumentAttributes66 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes66)
	return c.UnderlyingSecurity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation85 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of processing involved by a corporate action.
	EventProcessingType *CorporateActionEventProcessingType2Choice `xml:"EvtPrcgTp,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType33Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary3Choice `xml:"MndtryVlntryEvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes66 `xml:"UndrlygScty"`
}

func (c *CorporateActionGeneralInformation85) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation85) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation85) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation85) AddEventProcessingType() *CorporateActionEventProcessingType2Choice {
	c.EventProcessingType = new(CorporateActionEventProcessingType2Choice)
	return c.EventProcessingType
}

func (c *CorporateActionGeneralInformation85) AddEventType() *CorporateActionEventType33Choice {
	c.EventType = new(CorporateActionEventType33Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation85) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary3Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary3Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation85) AddUnderlyingSecurity() *FinancialInstrumentAttributes66 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes66)
	return c.UnderlyingSecurity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation86 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType33Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary3Choice `xml:"MndtryVlntryEvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`
}

func (c *CorporateActionGeneralInformation86) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation86) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation86) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation86) AddEventType() *CorporateActionEventType33Choice {
	c.EventType = new(CorporateActionEventType33Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation86) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary3Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary3Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation86) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return c.FinancialInstrumentIdentification
}

// General information about the corporate action event.
type CorporateActionGeneralInformation87 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType34Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary3Choice `xml:"MndtryVlntryEvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`
}

func (c *CorporateActionGeneralInformation87) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation87) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation87) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation87) AddEventType() *CorporateActionEventType34Choice {
	c.EventType = new(CorporateActionEventType34Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation87) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary3Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary3Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation87) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return c.FinancialInstrumentIdentification
}

// General information about the corporate action event.
type CorporateActionGeneralInformation88 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType32Choice `xml:"EvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes65 `xml:"UndrlygScty,omitempty"`
}

func (c *CorporateActionGeneralInformation88) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation88) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation88) AddEventType() *CorporateActionEventType32Choice {
	c.EventType = new(CorporateActionEventType32Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation88) AddUnderlyingSecurity() *FinancialInstrumentAttributes65 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes65)
	return c.UnderlyingSecurity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation89 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType31Choice `xml:"EvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat16Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Fractional quantity resulting from an event which will be paid with cash in lieu.
	FractionalQuantity *FinancialInstrumentQuantity1Choice `xml:"FrctnlQty,omitempty"`
}

func (c *CorporateActionGeneralInformation89) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation89) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation89) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation89) AddEventType() *CorporateActionEventType31Choice {
	c.EventType = new(CorporateActionEventType31Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation89) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return c.FinancialInstrumentIdentification
}

func (c *CorporateActionGeneralInformation89) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat16Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat16Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateActionGeneralInformation89) AddFractionalQuantity() *FinancialInstrumentQuantity1Choice {
	c.FractionalQuantity = new(FinancialInstrumentQuantity1Choice)
	return c.FractionalQuantity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation9 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType3Choice `xml:"EvtTp"`
}

func (c *CorporateActionGeneralInformation9) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation9) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation9) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation9) AddEventType() *CorporateActionEventType3Choice {
	c.EventType = new(CorporateActionEventType3Choice)
	return c.EventType
}

// General information about the corporate action event.
type CorporateActionGeneralInformation90 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType32Choice `xml:"EvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId,omitempty"`
}

func (c *CorporateActionGeneralInformation90) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation90) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation90) AddEventType() *CorporateActionEventType32Choice {
	c.EventType = new(CorporateActionEventType32Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation90) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return c.FinancialInstrumentIdentification
}

// General information about the corporate action event.
type CorporateActionGeneralInformation91 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *Max35Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType32Choice `xml:"EvtTp"`
}

func (c *CorporateActionGeneralInformation91) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation91) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation91) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation91) AddEventType() *CorporateActionEventType32Choice {
	c.EventType = new(CorporateActionEventType32Choice)
	return c.EventType
}

// General information about the corporate action event.
type CorporateActionGeneralInformation92 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *Max35Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Specifies the type of narrative related to the message.
	NarrativeType *CorporateActionNarrative3Choice `xml:"NrrtvTp,omitempty"`
}

func (c *CorporateActionGeneralInformation92) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation92) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionGeneralInformation92) AddNarrativeType() *CorporateActionNarrative3Choice {
	c.NarrativeType = new(CorporateActionNarrative3Choice)
	return c.NarrativeType
}

// General information about the corporate action event.
type CorporateActionGeneralInformation93 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *RestrictedFINXMax16Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType35Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary4Choice `xml:"MndtryVlntryEvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`
}

func (c *CorporateActionGeneralInformation93) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation93) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation93) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation93) AddEventType() *CorporateActionEventType35Choice {
	c.EventType = new(CorporateActionEventType35Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation93) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary4Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary4Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation93) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return c.FinancialInstrumentIdentification
}

// General information about the corporate action event.
type CorporateActionGeneralInformation94 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *RestrictedFINXMax16Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType36Choice `xml:"EvtTp"`
}

func (c *CorporateActionGeneralInformation94) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation94) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation94) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation94) AddEventType() *CorporateActionEventType36Choice {
	c.EventType = new(CorporateActionEventType36Choice)
	return c.EventType
}

// General information about the corporate action event.
type CorporateActionGeneralInformation97 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType36Choice `xml:"EvtTp"`

	// Security concerned by the corporate action.
	UnderlyingSecurity *FinancialInstrumentAttributes69 `xml:"UndrlygScty,omitempty"`
}

func (c *CorporateActionGeneralInformation97) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation97) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation97) AddEventType() *CorporateActionEventType36Choice {
	c.EventType = new(CorporateActionEventType36Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation97) AddUnderlyingSecurity() *FinancialInstrumentAttributes69 {
	c.UnderlyingSecurity = new(FinancialInstrumentAttributes69)
	return c.UnderlyingSecurity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation98 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *RestrictedFINXMax16Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType41Choice `xml:"EvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat17Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Fractional quantity resulting from an event which will be paid with cash in lieu.
	FractionalQuantity *FinancialInstrumentQuantity15Choice `xml:"FrctnlQty,omitempty"`
}

func (c *CorporateActionGeneralInformation98) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation98) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation98) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation98) AddEventType() *CorporateActionEventType41Choice {
	c.EventType = new(CorporateActionEventType41Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation98) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return c.FinancialInstrumentIdentification
}

func (c *CorporateActionGeneralInformation98) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat17Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat17Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateActionGeneralInformation98) AddFractionalQuantity() *FinancialInstrumentQuantity15Choice {
	c.FractionalQuantity = new(FinancialInstrumentQuantity15Choice)
	return c.FractionalQuantity
}

// General information about the corporate action event.
type CorporateActionGeneralInformation99 struct {

	// Reference assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId"`

	// Official and unique reference assigned by the official central body/entity within each market at the beginning of a corporate action event.
	OfficialCorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"OffclCorpActnEvtId,omitempty"`

	// Reference assigned by a court to a class action.
	ClassActionNumber *RestrictedFINXMax16Text `xml:"ClssActnNb,omitempty"`

	// Type of corporate action event.
	EventType *CorporateActionEventType42Choice `xml:"EvtTp"`

	// Specifies whether the event is mandatory, mandatory with options or voluntary.
	MandatoryVoluntaryEventType *CorporateActionMandatoryVoluntary4Choice `xml:"MndtryVlntryEvtTp"`

	// Identification of the security concerned by the corporate action.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`
}

func (c *CorporateActionGeneralInformation99) SetCorporateActionEventIdentification(value string) {
	c.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation99) SetOfficialCorporateActionEventIdentification(value string) {
	c.OfficialCorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation99) SetClassActionNumber(value string) {
	c.ClassActionNumber = (*RestrictedFINXMax16Text)(&value)
}

func (c *CorporateActionGeneralInformation99) AddEventType() *CorporateActionEventType42Choice {
	c.EventType = new(CorporateActionEventType42Choice)
	return c.EventType
}

func (c *CorporateActionGeneralInformation99) AddMandatoryVoluntaryEventType() *CorporateActionMandatoryVoluntary4Choice {
	c.MandatoryVoluntaryEventType = new(CorporateActionMandatoryVoluntary4Choice)
	return c.MandatoryVoluntaryEventType
}

func (c *CorporateActionGeneralInformation99) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return c.FinancialInstrumentIdentification
}
