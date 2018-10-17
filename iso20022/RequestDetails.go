package iso20022

// Details of the settlement condition modification request
type RequestDetails1 struct {

	// References of the transaction for which the securities settlement condition modification is requested.
	Reference *References1 `xml:"Ref"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing2Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether the instruction due to expire is confirmed for settlement
	RetainIndicator *YesNoIndicator `xml:"RtnInd,omitempty"`

	// Specifies the type of linkage requested.
	Linkage *LinkageType1Choice `xml:"Lkg,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Specifies another type of processing change request
	OtherProcessing []*GenericIdentification20 `xml:"OthrPrcg,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *YesNoIndicator `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Specifies the matching processing change requested.
	MatchingDenial *MatchingDenied1Choice `xml:"MtchgDnl,omitempty"`

	// Specifies that the transaction is requested to be unilaterally split.
	UnilateralSplit *UnilateralSplit1Choice `xml:"UnltrlSplt,omitempty"`

	// Information regarding the linkage requested.
	Linkages []*Linkages3 `xml:"Lnkgs,omitempty"`
}

func (r *RequestDetails1) AddReference() *References1 {
	r.Reference = new(References1)
	return r.Reference
}

func (r *RequestDetails1) AddAutomaticBorrowing() *AutomaticBorrowing2Choice {
	r.AutomaticBorrowing = new(AutomaticBorrowing2Choice)
	return r.AutomaticBorrowing
}

func (r *RequestDetails1) SetRetainIndicator(value string) {
	r.RetainIndicator = (*YesNoIndicator)(&value)
}

func (r *RequestDetails1) AddLinkage() *LinkageType1Choice {
	r.Linkage = new(LinkageType1Choice)
	return r.Linkage
}

func (r *RequestDetails1) AddPriority() *PriorityNumeric1Choice {
	r.Priority = new(PriorityNumeric1Choice)
	return r.Priority
}

func (r *RequestDetails1) AddOtherProcessing() *GenericIdentification20 {
	newValue := new(GenericIdentification20)
	r.OtherProcessing = append(r.OtherProcessing, newValue)
	return newValue
}

func (r *RequestDetails1) SetPartialSettlementIndicator(value string) {
	r.PartialSettlementIndicator = (*YesNoIndicator)(&value)
}

func (r *RequestDetails1) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	r.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return r.SecuritiesRTGS
}

func (r *RequestDetails1) SetHoldIndicator(value string) {
	r.HoldIndicator = (*YesNoIndicator)(&value)
}

func (r *RequestDetails1) AddMatchingDenial() *MatchingDenied1Choice {
	r.MatchingDenial = new(MatchingDenied1Choice)
	return r.MatchingDenial
}

func (r *RequestDetails1) AddUnilateralSplit() *UnilateralSplit1Choice {
	r.UnilateralSplit = new(UnilateralSplit1Choice)
	return r.UnilateralSplit
}

func (r *RequestDetails1) AddLinkages() *Linkages3 {
	newValue := new(Linkages3)
	r.Linkages = append(r.Linkages, newValue)
	return newValue
}

// Details of the settlement condition modification request
type RequestDetails11 struct {

	// References of the transaction for which the securities settlement condition modification is requested.
	Reference *References9 `xml:"Ref"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing2Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether the instruction due to expire is confirmed for settlement
	RetainIndicator *YesNoIndicator `xml:"RtnInd,omitempty"`

	// Specifies the type of linkage requested.
	Linkage *LinkageType1Choice `xml:"Lkg,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Specifies another type of processing change request
	OtherProcessing []*GenericIdentification20 `xml:"OthrPrcg,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator4 `xml:"HldInd,omitempty"`

	// Specifies the matching processing change requested.
	MatchingDenial *MatchingDenied1Choice `xml:"MtchgDnl,omitempty"`

	// Specifies that the transaction is requested to be unilaterally split.
	UnilateralSplit *UnilateralSplit1Choice `xml:"UnltrlSplt,omitempty"`

	// Information regarding the linkage requested.
	Linkages []*Linkages27 `xml:"Lnkgs,omitempty"`
}

func (r *RequestDetails11) AddReference() *References9 {
	r.Reference = new(References9)
	return r.Reference
}

func (r *RequestDetails11) AddAutomaticBorrowing() *AutomaticBorrowing2Choice {
	r.AutomaticBorrowing = new(AutomaticBorrowing2Choice)
	return r.AutomaticBorrowing
}

func (r *RequestDetails11) SetRetainIndicator(value string) {
	r.RetainIndicator = (*YesNoIndicator)(&value)
}

func (r *RequestDetails11) AddLinkage() *LinkageType1Choice {
	r.Linkage = new(LinkageType1Choice)
	return r.Linkage
}

func (r *RequestDetails11) AddPriority() *PriorityNumeric1Choice {
	r.Priority = new(PriorityNumeric1Choice)
	return r.Priority
}

func (r *RequestDetails11) AddOtherProcessing() *GenericIdentification20 {
	newValue := new(GenericIdentification20)
	r.OtherProcessing = append(r.OtherProcessing, newValue)
	return newValue
}

func (r *RequestDetails11) SetPartialSettlementIndicator(value string) {
	r.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (r *RequestDetails11) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	r.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return r.SecuritiesRTGS
}

func (r *RequestDetails11) AddHoldIndicator() *HoldIndicator4 {
	r.HoldIndicator = new(HoldIndicator4)
	return r.HoldIndicator
}

func (r *RequestDetails11) AddMatchingDenial() *MatchingDenied1Choice {
	r.MatchingDenial = new(MatchingDenied1Choice)
	return r.MatchingDenial
}

func (r *RequestDetails11) AddUnilateralSplit() *UnilateralSplit1Choice {
	r.UnilateralSplit = new(UnilateralSplit1Choice)
	return r.UnilateralSplit
}

func (r *RequestDetails11) AddLinkages() *Linkages27 {
	newValue := new(Linkages27)
	r.Linkages = append(r.Linkages, newValue)
	return newValue
}

// Details of the settlement condition modification request
type RequestDetails13 struct {

	// References of the transaction for which the securities settlement condition modification is requested.
	Reference *References16 `xml:"Ref"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing2Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether the instruction due to expire is confirmed for settlement
	RetainIndicator *YesNoIndicator `xml:"RtnInd,omitempty"`

	// Specifies the type of linkage requested.
	Linkage *LinkageType1Choice `xml:"Lkg,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Specifies another type of processing change request
	OtherProcessing []*GenericIdentification20 `xml:"OthrPrcg,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator4 `xml:"HldInd,omitempty"`

	// Specifies the matching processing change requested.
	MatchingDenial *MatchingDenied1Choice `xml:"MtchgDnl,omitempty"`

	// Specifies that the transaction is requested to be unilaterally split.
	UnilateralSplit *UnilateralSplit1Choice `xml:"UnltrlSplt,omitempty"`

	// Information regarding the linkage requested.
	Linkages []*Linkages27 `xml:"Lnkgs,omitempty"`
}

func (r *RequestDetails13) AddReference() *References16 {
	r.Reference = new(References16)
	return r.Reference
}

func (r *RequestDetails13) AddAutomaticBorrowing() *AutomaticBorrowing2Choice {
	r.AutomaticBorrowing = new(AutomaticBorrowing2Choice)
	return r.AutomaticBorrowing
}

func (r *RequestDetails13) SetRetainIndicator(value string) {
	r.RetainIndicator = (*YesNoIndicator)(&value)
}

func (r *RequestDetails13) AddLinkage() *LinkageType1Choice {
	r.Linkage = new(LinkageType1Choice)
	return r.Linkage
}

func (r *RequestDetails13) AddPriority() *PriorityNumeric1Choice {
	r.Priority = new(PriorityNumeric1Choice)
	return r.Priority
}

func (r *RequestDetails13) AddOtherProcessing() *GenericIdentification20 {
	newValue := new(GenericIdentification20)
	r.OtherProcessing = append(r.OtherProcessing, newValue)
	return newValue
}

func (r *RequestDetails13) SetPartialSettlementIndicator(value string) {
	r.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (r *RequestDetails13) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	r.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return r.SecuritiesRTGS
}

func (r *RequestDetails13) AddHoldIndicator() *HoldIndicator4 {
	r.HoldIndicator = new(HoldIndicator4)
	return r.HoldIndicator
}

func (r *RequestDetails13) AddMatchingDenial() *MatchingDenied1Choice {
	r.MatchingDenial = new(MatchingDenied1Choice)
	return r.MatchingDenial
}

func (r *RequestDetails13) AddUnilateralSplit() *UnilateralSplit1Choice {
	r.UnilateralSplit = new(UnilateralSplit1Choice)
	return r.UnilateralSplit
}

func (r *RequestDetails13) AddLinkages() *Linkages27 {
	newValue := new(Linkages27)
	r.Linkages = append(r.Linkages, newValue)
	return newValue
}

// Details of the settlement condition modification request
type RequestDetails15 struct {

	// References of the transaction for which the securities settlement condition modification is requested.
	Reference *References18 `xml:"Ref"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing7Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether the instruction due to expire is confirmed for settlement
	RetainIndicator *YesNoIndicator `xml:"RtnInd,omitempty"`

	// Specifies the type of linkage requested.
	Linkage *LinkageType3Choice `xml:"Lkg,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric4Choice `xml:"Prty,omitempty"`

	// Specifies another type of processing change request
	OtherProcessing []*GenericIdentification30 `xml:"OthrPrcg,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS4Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator6 `xml:"HldInd,omitempty"`

	// Specifies the matching processing change requested.
	MatchingDenial *MatchingDenied3Choice `xml:"MtchgDnl,omitempty"`

	// Specifies that the transaction is requested to be unilaterally split.
	UnilateralSplit *UnilateralSplit3Choice `xml:"UnltrlSplt,omitempty"`

	// Information regarding the linkage requested.
	Linkages []*Linkages39 `xml:"Lnkgs,omitempty"`
}

func (r *RequestDetails15) AddReference() *References18 {
	r.Reference = new(References18)
	return r.Reference
}

func (r *RequestDetails15) AddAutomaticBorrowing() *AutomaticBorrowing7Choice {
	r.AutomaticBorrowing = new(AutomaticBorrowing7Choice)
	return r.AutomaticBorrowing
}

func (r *RequestDetails15) SetRetainIndicator(value string) {
	r.RetainIndicator = (*YesNoIndicator)(&value)
}

func (r *RequestDetails15) AddLinkage() *LinkageType3Choice {
	r.Linkage = new(LinkageType3Choice)
	return r.Linkage
}

func (r *RequestDetails15) AddPriority() *PriorityNumeric4Choice {
	r.Priority = new(PriorityNumeric4Choice)
	return r.Priority
}

func (r *RequestDetails15) AddOtherProcessing() *GenericIdentification30 {
	newValue := new(GenericIdentification30)
	r.OtherProcessing = append(r.OtherProcessing, newValue)
	return newValue
}

func (r *RequestDetails15) SetPartialSettlementIndicator(value string) {
	r.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (r *RequestDetails15) AddSecuritiesRTGS() *SecuritiesRTGS4Choice {
	r.SecuritiesRTGS = new(SecuritiesRTGS4Choice)
	return r.SecuritiesRTGS
}

func (r *RequestDetails15) AddHoldIndicator() *HoldIndicator6 {
	r.HoldIndicator = new(HoldIndicator6)
	return r.HoldIndicator
}

func (r *RequestDetails15) AddMatchingDenial() *MatchingDenied3Choice {
	r.MatchingDenial = new(MatchingDenied3Choice)
	return r.MatchingDenial
}

func (r *RequestDetails15) AddUnilateralSplit() *UnilateralSplit3Choice {
	r.UnilateralSplit = new(UnilateralSplit3Choice)
	return r.UnilateralSplit
}

func (r *RequestDetails15) AddLinkages() *Linkages39 {
	newValue := new(Linkages39)
	r.Linkages = append(r.Linkages, newValue)
	return newValue
}

// Details of the settlement condition modification request
type RequestDetails16 struct {

	// References of the transaction for which the securities settlement condition modification is requested.
	Reference *References21 `xml:"Ref"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing11Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether the instruction due to expire is confirmed for settlement
	RetainIndicator *YesNoIndicator `xml:"RtnInd,omitempty"`

	// Specifies the type of linkage requested.
	Linkage *LinkageType4Choice `xml:"Lkg,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric5Choice `xml:"Prty,omitempty"`

	// Specifies another type of processing change request
	OtherProcessing []*GenericIdentification47 `xml:"OthrPrcg,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS5Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator7 `xml:"HldInd,omitempty"`

	// Specifies the matching processing change requested.
	MatchingDenial *MatchingDenied4Choice `xml:"MtchgDnl,omitempty"`

	// Specifies that the transaction is requested to be unilaterally split.
	UnilateralSplit *UnilateralSplit4Choice `xml:"UnltrlSplt,omitempty"`

	// Information regarding the linkage requested.
	Linkages []*Linkages44 `xml:"Lnkgs,omitempty"`
}

func (r *RequestDetails16) AddReference() *References21 {
	r.Reference = new(References21)
	return r.Reference
}

func (r *RequestDetails16) AddAutomaticBorrowing() *AutomaticBorrowing11Choice {
	r.AutomaticBorrowing = new(AutomaticBorrowing11Choice)
	return r.AutomaticBorrowing
}

func (r *RequestDetails16) SetRetainIndicator(value string) {
	r.RetainIndicator = (*YesNoIndicator)(&value)
}

func (r *RequestDetails16) AddLinkage() *LinkageType4Choice {
	r.Linkage = new(LinkageType4Choice)
	return r.Linkage
}

func (r *RequestDetails16) AddPriority() *PriorityNumeric5Choice {
	r.Priority = new(PriorityNumeric5Choice)
	return r.Priority
}

func (r *RequestDetails16) AddOtherProcessing() *GenericIdentification47 {
	newValue := new(GenericIdentification47)
	r.OtherProcessing = append(r.OtherProcessing, newValue)
	return newValue
}

func (r *RequestDetails16) SetPartialSettlementIndicator(value string) {
	r.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (r *RequestDetails16) AddSecuritiesRTGS() *SecuritiesRTGS5Choice {
	r.SecuritiesRTGS = new(SecuritiesRTGS5Choice)
	return r.SecuritiesRTGS
}

func (r *RequestDetails16) AddHoldIndicator() *HoldIndicator7 {
	r.HoldIndicator = new(HoldIndicator7)
	return r.HoldIndicator
}

func (r *RequestDetails16) AddMatchingDenial() *MatchingDenied4Choice {
	r.MatchingDenial = new(MatchingDenied4Choice)
	return r.MatchingDenial
}

func (r *RequestDetails16) AddUnilateralSplit() *UnilateralSplit4Choice {
	r.UnilateralSplit = new(UnilateralSplit4Choice)
	return r.UnilateralSplit
}

func (r *RequestDetails16) AddLinkages() *Linkages44 {
	newValue := new(Linkages44)
	r.Linkages = append(r.Linkages, newValue)
	return newValue
}

// Details of the processing request.
type RequestDetails19 struct {

	// Type of data being requested, for example, a sub-member BIC.
	Type *Max35Text `xml:"Tp"`

	// Identificates the requestor.
	RequestorIdentification *PartyIdentification73Choice `xml:"RqstrId,omitempty"`

	// Additional information to support the processing request.
	AdditionalRequestInformation []*Max35Text `xml:"AddtlReqInf,omitempty"`
}

func (r *RequestDetails19) SetType(value string) {
	r.Type = (*Max35Text)(&value)
}

func (r *RequestDetails19) AddRequestorIdentification() *PartyIdentification73Choice {
	r.RequestorIdentification = new(PartyIdentification73Choice)
	return r.RequestorIdentification
}

func (r *RequestDetails19) AddAdditionalRequestInformation(value string) {
	r.AdditionalRequestInformation = append(r.AdditionalRequestInformation, (*Max35Text)(&value))
}

// Details of data request.
type RequestDetails3 struct {

	// Type of data being requested, for example, a sub-member BIC.
	Type *Max35Text `xml:"Tp"`

	// Specific report data requested, for example, a  BIC.
	Key *Max35Text `xml:"Key,omitempty"`
}

func (r *RequestDetails3) SetType(value string) {
	r.Type = (*Max35Text)(&value)
}

func (r *RequestDetails3) SetKey(value string) {
	r.Key = (*Max35Text)(&value)
}

// Details of one or several keys of the request.
type RequestDetails4 struct {

	// Key for which the specific data is returned, for example, a BIC.
	Key *Max35Text `xml:"Key"`

	// Data being returned.
	ReportData []*ReportParameter1 `xml:"RptData,omitempty"`
}

func (r *RequestDetails4) SetKey(value string) {
	r.Key = (*Max35Text)(&value)
}

func (r *RequestDetails4) AddReportData() *ReportParameter1 {
	newValue := new(ReportParameter1)
	r.ReportData = append(r.ReportData, newValue)
	return newValue
}

// Report of the requested data.
type RequestDetails5 struct {

	// Type of data requested, for example, a sub-member BIC.
	Type *Max35Text `xml:"Tp"`

	// Reference to the request for which the report is sent.
	RequestReference *Max35Text `xml:"ReqRef"`

	// Report key and returned data.
	ReportKey []*RequestDetails4 `xml:"RptKey"`
}

func (r *RequestDetails5) SetType(value string) {
	r.Type = (*Max35Text)(&value)
}

func (r *RequestDetails5) SetRequestReference(value string) {
	r.RequestReference = (*Max35Text)(&value)
}

func (r *RequestDetails5) AddReportKey() *RequestDetails4 {
	newValue := new(RequestDetails4)
	r.ReportKey = append(r.ReportKey, newValue)
	return newValue
}

// Details of the settlement condition modification request
type RequestDetails6 struct {

	// References of the transaction for which the securities settlement condition modification is requested.
	Reference *References7 `xml:"Ref"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing2Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether the instruction due to expire is confirmed for settlement
	RetainIndicator *YesNoIndicator `xml:"RtnInd,omitempty"`

	// Specifies the type of linkage requested.
	Linkage *LinkageType1Choice `xml:"Lkg,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Specifies another type of processing change request
	OtherProcessing []*GenericIdentification20 `xml:"OthrPrcg,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator2 `xml:"HldInd,omitempty"`

	// Specifies the matching processing change requested.
	MatchingDenial *MatchingDenied1Choice `xml:"MtchgDnl,omitempty"`

	// Specifies that the transaction is requested to be unilaterally split.
	UnilateralSplit *UnilateralSplit1Choice `xml:"UnltrlSplt,omitempty"`

	// Information regarding the linkage requested.
	Linkages []*Linkages10 `xml:"Lnkgs,omitempty"`
}

func (r *RequestDetails6) AddReference() *References7 {
	r.Reference = new(References7)
	return r.Reference
}

func (r *RequestDetails6) AddAutomaticBorrowing() *AutomaticBorrowing2Choice {
	r.AutomaticBorrowing = new(AutomaticBorrowing2Choice)
	return r.AutomaticBorrowing
}

func (r *RequestDetails6) SetRetainIndicator(value string) {
	r.RetainIndicator = (*YesNoIndicator)(&value)
}

func (r *RequestDetails6) AddLinkage() *LinkageType1Choice {
	r.Linkage = new(LinkageType1Choice)
	return r.Linkage
}

func (r *RequestDetails6) AddPriority() *PriorityNumeric1Choice {
	r.Priority = new(PriorityNumeric1Choice)
	return r.Priority
}

func (r *RequestDetails6) AddOtherProcessing() *GenericIdentification20 {
	newValue := new(GenericIdentification20)
	r.OtherProcessing = append(r.OtherProcessing, newValue)
	return newValue
}

func (r *RequestDetails6) SetPartialSettlementIndicator(value string) {
	r.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (r *RequestDetails6) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	r.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return r.SecuritiesRTGS
}

func (r *RequestDetails6) AddHoldIndicator() *HoldIndicator2 {
	r.HoldIndicator = new(HoldIndicator2)
	return r.HoldIndicator
}

func (r *RequestDetails6) AddMatchingDenial() *MatchingDenied1Choice {
	r.MatchingDenial = new(MatchingDenied1Choice)
	return r.MatchingDenial
}

func (r *RequestDetails6) AddUnilateralSplit() *UnilateralSplit1Choice {
	r.UnilateralSplit = new(UnilateralSplit1Choice)
	return r.UnilateralSplit
}

func (r *RequestDetails6) AddLinkages() *Linkages10 {
	newValue := new(Linkages10)
	r.Linkages = append(r.Linkages, newValue)
	return newValue
}

// Details of the settlement condition modification request
type RequestDetails8 struct {

	// References of the transaction for which the securities settlement condition modification is requested.
	Reference *References9 `xml:"Ref"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing2Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether the instruction due to expire is confirmed for settlement
	RetainIndicator *YesNoIndicator `xml:"RtnInd,omitempty"`

	// Specifies the type of linkage requested.
	Linkage *LinkageType1Choice `xml:"Lkg,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Specifies another type of processing change request
	OtherProcessing []*GenericIdentification20 `xml:"OthrPrcg,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator4 `xml:"HldInd,omitempty"`

	// Specifies the matching processing change requested.
	MatchingDenial *MatchingDenied1Choice `xml:"MtchgDnl,omitempty"`

	// Specifies that the transaction is requested to be unilaterally split.
	UnilateralSplit *UnilateralSplit1Choice `xml:"UnltrlSplt,omitempty"`

	// Information regarding the linkage requested.
	Linkages []*Linkages16 `xml:"Lnkgs,omitempty"`
}

func (r *RequestDetails8) AddReference() *References9 {
	r.Reference = new(References9)
	return r.Reference
}

func (r *RequestDetails8) AddAutomaticBorrowing() *AutomaticBorrowing2Choice {
	r.AutomaticBorrowing = new(AutomaticBorrowing2Choice)
	return r.AutomaticBorrowing
}

func (r *RequestDetails8) SetRetainIndicator(value string) {
	r.RetainIndicator = (*YesNoIndicator)(&value)
}

func (r *RequestDetails8) AddLinkage() *LinkageType1Choice {
	r.Linkage = new(LinkageType1Choice)
	return r.Linkage
}

func (r *RequestDetails8) AddPriority() *PriorityNumeric1Choice {
	r.Priority = new(PriorityNumeric1Choice)
	return r.Priority
}

func (r *RequestDetails8) AddOtherProcessing() *GenericIdentification20 {
	newValue := new(GenericIdentification20)
	r.OtherProcessing = append(r.OtherProcessing, newValue)
	return newValue
}

func (r *RequestDetails8) SetPartialSettlementIndicator(value string) {
	r.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (r *RequestDetails8) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	r.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return r.SecuritiesRTGS
}

func (r *RequestDetails8) AddHoldIndicator() *HoldIndicator4 {
	r.HoldIndicator = new(HoldIndicator4)
	return r.HoldIndicator
}

func (r *RequestDetails8) AddMatchingDenial() *MatchingDenied1Choice {
	r.MatchingDenial = new(MatchingDenied1Choice)
	return r.MatchingDenial
}

func (r *RequestDetails8) AddUnilateralSplit() *UnilateralSplit1Choice {
	r.UnilateralSplit = new(UnilateralSplit1Choice)
	return r.UnilateralSplit
}

func (r *RequestDetails8) AddLinkages() *Linkages16 {
	newValue := new(Linkages16)
	r.Linkages = append(r.Linkages, newValue)
	return newValue
}
