package iso20022

// Provides information on the status of a trade.
type TradeData1 struct {

	// Refers to the identification of a notification.
	NotificationIdentification *Max35Text `xml:"NtfctnId"`

	// Reference to the unique identification assigned to a trade by a central matching system.
	MatchingSystemUniqueReference *Max35Text `xml:"MtchgSysUnqRef"`

	// Identifies the party which assigned a status to a treasury trade.
	StatusOriginator *Max35Text `xml:"StsOrgtr,omitempty"`

	// Specifies the new status of a trade.
	CurrentStatus *TradeStatus1Code `xml:"CurSts"`

	// Description of the status of a trade when no coded form is available.
	ExtendedCurrentStatus *Extended350Code `xml:"XtndedCurSts"`

	// Additional information on the current status of a trade in a central system.
	CurrentStatusSubType *Max70Text `xml:"CurStsSubTp,omitempty"`

	// Specifies the time at which the current status was assigned.
	CurrentStatusTime *ISODateTime `xml:"CurStsTm,omitempty"`

	// Specifies the previous status of a trade.
	PreviousStatus *TradeStatus1Code `xml:"PrvsSts,omitempty"`

	// Description of the status of a trade when no coded form is available.
	ExtendedPreviousStatus *Extended350Code `xml:"XtndedPrvsSts,omitempty"`

	// Additional information on the previous status of a trade in a central system.
	PreviousStatusSubType *Max70Text `xml:"PrvsStsSubTp,omitempty"`

	// Specifies the time at which the previous status was assigned.
	PreviousStatusTime *ISODateTime `xml:"PrvsStsTm,omitempty"`

	// Specifies the product for which the status of the confirmation is reported.
	ProductType *Max4AlphaNumericText `xml:"PdctTp,omitempty"`
}

func (t *TradeData1) SetNotificationIdentification(value string) {
	t.NotificationIdentification = (*Max35Text)(&value)
}

func (t *TradeData1) SetMatchingSystemUniqueReference(value string) {
	t.MatchingSystemUniqueReference = (*Max35Text)(&value)
}

func (t *TradeData1) SetStatusOriginator(value string) {
	t.StatusOriginator = (*Max35Text)(&value)
}

func (t *TradeData1) SetCurrentStatus(value string) {
	t.CurrentStatus = (*TradeStatus1Code)(&value)
}

func (t *TradeData1) SetExtendedCurrentStatus(value string) {
	t.ExtendedCurrentStatus = (*Extended350Code)(&value)
}

func (t *TradeData1) SetCurrentStatusSubType(value string) {
	t.CurrentStatusSubType = (*Max70Text)(&value)
}

func (t *TradeData1) SetCurrentStatusTime(value string) {
	t.CurrentStatusTime = (*ISODateTime)(&value)
}

func (t *TradeData1) SetPreviousStatus(value string) {
	t.PreviousStatus = (*TradeStatus1Code)(&value)
}

func (t *TradeData1) SetExtendedPreviousStatus(value string) {
	t.ExtendedPreviousStatus = (*Extended350Code)(&value)
}

func (t *TradeData1) SetPreviousStatusSubType(value string) {
	t.PreviousStatusSubType = (*Max70Text)(&value)
}

func (t *TradeData1) SetPreviousStatusTime(value string) {
	t.PreviousStatusTime = (*ISODateTime)(&value)
}

func (t *TradeData1) SetProductType(value string) {
	t.ProductType = (*Max4AlphaNumericText)(&value)
}

// Provides information on the status of a trade.
type TradeData10 struct {

	// Identification of the present message assigned by the party issuing the message. This identification must be unique amongst all messages of same type sent by the same party.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Party that assigned the status to the foreign exchange trade.
	StatusOriginator *Max35Text `xml:"StsOrgtr,omitempty"`

	// Specifies the new status of the trade.
	CurrentStatus *StatusAndSubStatus1 `xml:"CurSts"`

	// Additional information about the current status of the trade.
	CurrentStatusSubType *StatusSubType1Code `xml:"CurStsSubTp,omitempty"`

	// Specifies the date and time at which the current status was assigned to all the trades, unless overwritten by a date and time assigned to an individual trade.
	CurrentStatusDateTime *ISODateTime `xml:"CurStsDtTm"`

	// Specifies the previous status of a trade.
	PreviousStatus *Status5Choice `xml:"PrvsSts,omitempty"`

	// Additional information on the previous status of a trade in a central system.
	PreviousStatusSubType *StatusSubType1Code `xml:"PrvsStsSubTp,omitempty"`

	// Specifies the product for which the status of the confirmation is reported, unless overwritten by a product type assigned to an individual trade.
	ProductType *Max35Text `xml:"PdctTp,omitempty"`
}

func (t *TradeData10) SetMessageIdentification(value string) {
	t.MessageIdentification = (*Max35Text)(&value)
}

func (t *TradeData10) SetStatusOriginator(value string) {
	t.StatusOriginator = (*Max35Text)(&value)
}

func (t *TradeData10) AddCurrentStatus() *StatusAndSubStatus1 {
	t.CurrentStatus = new(StatusAndSubStatus1)
	return t.CurrentStatus
}

func (t *TradeData10) SetCurrentStatusSubType(value string) {
	t.CurrentStatusSubType = (*StatusSubType1Code)(&value)
}

func (t *TradeData10) SetCurrentStatusDateTime(value string) {
	t.CurrentStatusDateTime = (*ISODateTime)(&value)
}

func (t *TradeData10) AddPreviousStatus() *Status5Choice {
	t.PreviousStatus = new(Status5Choice)
	return t.PreviousStatus
}

func (t *TradeData10) SetPreviousStatusSubType(value string) {
	t.PreviousStatusSubType = (*StatusSubType1Code)(&value)
}

func (t *TradeData10) SetProductType(value string) {
	t.ProductType = (*Max35Text)(&value)
}

// Provides information and details on the status of a trade.
type TradeData11 struct {

	// Represents the original reference of the instruction for which the status is given, as assigned by the participant that submitted the foreign exchange trade.
	OriginatorReference *Max35Text `xml:"OrgtrRef,omitempty"`

	// Reference to the unique system identification assigned to the trade  by the central matching system.
	MatchingSystemUniqueReference *Max35Text `xml:"MtchgSysUnqRef"`

	// Reference to the unique matching identification assigned to the trade and to the matching trade from the counterparty by the central matching system.
	MatchingSystemMatchingReference *Max35Text `xml:"MtchgSysMtchgRef,omitempty"`

	// Unique reference from the central settlement system that allows the removal of alleged trades once the matched status notification for the matching side has been received.
	MatchingSystemMatchedSideReference *Max35Text `xml:"MtchgSysMtchdSdRef,omitempty"`

	// The current settlement date of the notification.
	CurrentSettlementDate *ISODate `xml:"CurSttlmDt,omitempty"`

	// Settlement date has been amended.
	NewSettlementDate *ISODate `xml:"NewSttlmDt,omitempty"`

	// Specifies the date and time at which the current status was assigned to the individual trade.
	CurrentStatusDateTime *ISODateTime `xml:"CurStsDtTm,omitempty"`

	// Product type of the individual trade.
	ProductType *Max35Text `xml:"PdctTp,omitempty"`

	// To indicate the requested CLS settlement session that the related trade is part of.
	SettlementSessionIdentifier *Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty"`

	// Information that is to be provided to trade repositories in the context of the regulatory standards around over-the-counter (OTC) derivatives, central counterparties and trade repositories.
	RegulatoryReporting *RegulatoryReporting4 `xml:"RgltryRptg,omitempty"`
}

func (t *TradeData11) SetOriginatorReference(value string) {
	t.OriginatorReference = (*Max35Text)(&value)
}

func (t *TradeData11) SetMatchingSystemUniqueReference(value string) {
	t.MatchingSystemUniqueReference = (*Max35Text)(&value)
}

func (t *TradeData11) SetMatchingSystemMatchingReference(value string) {
	t.MatchingSystemMatchingReference = (*Max35Text)(&value)
}

func (t *TradeData11) SetMatchingSystemMatchedSideReference(value string) {
	t.MatchingSystemMatchedSideReference = (*Max35Text)(&value)
}

func (t *TradeData11) SetCurrentSettlementDate(value string) {
	t.CurrentSettlementDate = (*ISODate)(&value)
}

func (t *TradeData11) SetNewSettlementDate(value string) {
	t.NewSettlementDate = (*ISODate)(&value)
}

func (t *TradeData11) SetCurrentStatusDateTime(value string) {
	t.CurrentStatusDateTime = (*ISODateTime)(&value)
}

func (t *TradeData11) SetProductType(value string) {
	t.ProductType = (*Max35Text)(&value)
}

func (t *TradeData11) SetSettlementSessionIdentifier(value string) {
	t.SettlementSessionIdentifier = (*Exact4AlphaNumericText)(&value)
}

func (t *TradeData11) AddRegulatoryReporting() *RegulatoryReporting4 {
	t.RegulatoryReporting = new(RegulatoryReporting4)
	return t.RegulatoryReporting
}

// Provides information on the status of a trade.
type TradeData12 struct {

	// Identification of the present message assigned by the party issuing the message. This identification must be unique amongst all messages of same type sent by the same party.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Party that assigned the status to the foreign exchange trade.
	StatusOriginator *Max35Text `xml:"StsOrgtr,omitempty"`

	// Specifies the new status of the trade.
	CurrentStatus *StatusAndSubStatus2 `xml:"CurSts"`

	// Additional information about the current status of the trade.
	CurrentStatusSubType *StatusSubType2Code `xml:"CurStsSubTp,omitempty"`

	// Specifies the date and time at which the current status was assigned to all the trades, unless overwritten by a date and time assigned to an individual trade.
	CurrentStatusDateTime *ISODateTime `xml:"CurStsDtTm"`

	// Specifies the previous status of a trade.
	PreviousStatus *Status28Choice `xml:"PrvsSts,omitempty"`

	// Additional information on the previous status of a trade in a central system.
	PreviousStatusSubType *StatusSubType2Code `xml:"PrvsStsSubTp,omitempty"`

	// Specifies the product for which the status of the confirmation is reported, unless overwritten by a product type assigned to an individual trade.
	ProductType *Max35Text `xml:"PdctTp,omitempty"`

	// To indicate the requested CLS settlement session that the related trade is part of.
	SettlementSessionIdentifier *Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty"`

	// The identification that links the quoted trades with a submitted Report issued by a central system.
	LinkedReportIdentification *Max35Text `xml:"LkdRptId,omitempty"`
}

func (t *TradeData12) SetMessageIdentification(value string) {
	t.MessageIdentification = (*Max35Text)(&value)
}

func (t *TradeData12) SetStatusOriginator(value string) {
	t.StatusOriginator = (*Max35Text)(&value)
}

func (t *TradeData12) AddCurrentStatus() *StatusAndSubStatus2 {
	t.CurrentStatus = new(StatusAndSubStatus2)
	return t.CurrentStatus
}

func (t *TradeData12) SetCurrentStatusSubType(value string) {
	t.CurrentStatusSubType = (*StatusSubType2Code)(&value)
}

func (t *TradeData12) SetCurrentStatusDateTime(value string) {
	t.CurrentStatusDateTime = (*ISODateTime)(&value)
}

func (t *TradeData12) AddPreviousStatus() *Status28Choice {
	t.PreviousStatus = new(Status28Choice)
	return t.PreviousStatus
}

func (t *TradeData12) SetPreviousStatusSubType(value string) {
	t.PreviousStatusSubType = (*StatusSubType2Code)(&value)
}

func (t *TradeData12) SetProductType(value string) {
	t.ProductType = (*Max35Text)(&value)
}

func (t *TradeData12) SetSettlementSessionIdentifier(value string) {
	t.SettlementSessionIdentifier = (*Exact4AlphaNumericText)(&value)
}

func (t *TradeData12) SetLinkedReportIdentification(value string) {
	t.LinkedReportIdentification = (*Max35Text)(&value)
}

// Provides information on the status of a trade.
type TradeData14 struct {

	// Reference to the unique system identification assigned to the trade by the central matching system.
	MatchingSystemUniqueReference *Max35Text `xml:"MtchgSysUnqRef"`

	// Reference to the unique matching identification assigned to the trade and to the matching trade from the counterparty by the central matching system.
	MatchingSystemMatchingReference *Max35Text `xml:"MtchgSysMtchgRef,omitempty"`

	// Unique reference from the central settlement system that allows the removal of alleged trades once the matched status notification for the matching side has been received.
	MatchingSystemMatchedSideReference *Max35Text `xml:"MtchgSysMtchdSdRef,omitempty"`

	// Party that assigned the status to the trade.
	StatusOriginator *Max35Text `xml:"StsOrgtr,omitempty"`

	// Specifies the new status of the trade.
	CurrentStatus *StatusAndSubStatus2 `xml:"CurSts"`

	// Additional information about the current status of the trade.
	CurrentStatusSubType *StatusSubType2Code `xml:"CurStsSubTp,omitempty"`

	// Specifies the date and time at which the current status was assigned.
	CurrentStatusDateTime *ISODateTime `xml:"CurStsDtTm,omitempty"`

	// Specifies the previous status of the trade.
	PreviousStatus *Status28Choice `xml:"PrvsSts,omitempty"`

	// Specifies whether a trade is alleged or not.
	AllegedTrade *YesNoIndicator `xml:"AllgdTrad,omitempty"`

	// Additional information on the previous status of a trade in a central system.
	PreviousStatusSubType *StatusSubType2Code `xml:"PrvsStsSubTp,omitempty"`
}

func (t *TradeData14) SetMatchingSystemUniqueReference(value string) {
	t.MatchingSystemUniqueReference = (*Max35Text)(&value)
}

func (t *TradeData14) SetMatchingSystemMatchingReference(value string) {
	t.MatchingSystemMatchingReference = (*Max35Text)(&value)
}

func (t *TradeData14) SetMatchingSystemMatchedSideReference(value string) {
	t.MatchingSystemMatchedSideReference = (*Max35Text)(&value)
}

func (t *TradeData14) SetStatusOriginator(value string) {
	t.StatusOriginator = (*Max35Text)(&value)
}

func (t *TradeData14) AddCurrentStatus() *StatusAndSubStatus2 {
	t.CurrentStatus = new(StatusAndSubStatus2)
	return t.CurrentStatus
}

func (t *TradeData14) SetCurrentStatusSubType(value string) {
	t.CurrentStatusSubType = (*StatusSubType2Code)(&value)
}

func (t *TradeData14) SetCurrentStatusDateTime(value string) {
	t.CurrentStatusDateTime = (*ISODateTime)(&value)
}

func (t *TradeData14) AddPreviousStatus() *Status28Choice {
	t.PreviousStatus = new(Status28Choice)
	return t.PreviousStatus
}

func (t *TradeData14) SetAllegedTrade(value string) {
	t.AllegedTrade = (*YesNoIndicator)(&value)
}

func (t *TradeData14) SetPreviousStatusSubType(value string) {
	t.PreviousStatusSubType = (*StatusSubType2Code)(&value)
}

// Provides information on the status of a trade.
type TradeData15 struct {

	// Identification of the present message assigned by the party issuing the message. This identification must be unique amongst all messages of same type sent by the same party.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Represents the original reference of the instruction for which the status is given, as assigned by the participant that submitted the foreign exchange trade.
	OriginatorReference *Max35Text `xml:"OrgtrRef,omitempty"`

	// Reference to the unique system identification assigned to the trade  by the central matching system.
	MatchingSystemUniqueReference *Max35Text `xml:"MtchgSysUnqRef"`

	// Unique matching identification assigned to the trade and to the matching trade from the counterparty by the central matching system.
	MatchingSystemMatchingReference *Max35Text `xml:"MtchgSysMtchgRef,omitempty"`

	// Identification to the unique reference from the central settlement system that allows the removal of alleged trades once the matched status notification for the matching side has been received.
	MatchingSystemMatchedSideReference *Max35Text `xml:"MtchgSysMtchdSdRef,omitempty"`

	// Party that assigned the status to the foreign exchange or derivative trade.
	StatusOriginator *Max20Text `xml:"StsOrgtr,omitempty"`

	// Specifies the new status of a trade.
	CurrentStatus *StatusAndSubStatus2 `xml:"CurSts"`

	// Additional information on the current status of a trade in a central system.
	CurrentStatusSubType *StatusSubType2Code `xml:"CurStsSubTp,omitempty"`

	// Specifies the date and time at which the current status was assigned.
	CurrentStatusDateTime *ISODateTime `xml:"CurStsDtTm,omitempty"`

	// Specifies the previous status of a trade.
	PreviousStatus *Status28Choice `xml:"PrvsSts,omitempty"`

	// Additional information on the previous status of a trade in a central system.
	PreviousStatusSubType *StatusSubType2Code `xml:"PrvsStsSubTp,omitempty"`

	// Specifies the date and time at which the previous status was assigned.
	PreviousStatusDateTime *ISODateTime `xml:"PrvsStsDtTm,omitempty"`

	// Specifies the product for which the status of the confirmation is reported.
	ProductType *Max35Text `xml:"PdctTp,omitempty"`

	// To indicate the requested CLS Settlement Session that the related trade is part of.
	SettlementSessionIdentifier *Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty"`

	// To indicate if the trade is split.
	SplitTradeIndicator *YesNoIndicator `xml:"SpltTradInd,omitempty"`
}

func (t *TradeData15) SetMessageIdentification(value string) {
	t.MessageIdentification = (*Max35Text)(&value)
}

func (t *TradeData15) SetOriginatorReference(value string) {
	t.OriginatorReference = (*Max35Text)(&value)
}

func (t *TradeData15) SetMatchingSystemUniqueReference(value string) {
	t.MatchingSystemUniqueReference = (*Max35Text)(&value)
}

func (t *TradeData15) SetMatchingSystemMatchingReference(value string) {
	t.MatchingSystemMatchingReference = (*Max35Text)(&value)
}

func (t *TradeData15) SetMatchingSystemMatchedSideReference(value string) {
	t.MatchingSystemMatchedSideReference = (*Max35Text)(&value)
}

func (t *TradeData15) SetStatusOriginator(value string) {
	t.StatusOriginator = (*Max20Text)(&value)
}

func (t *TradeData15) AddCurrentStatus() *StatusAndSubStatus2 {
	t.CurrentStatus = new(StatusAndSubStatus2)
	return t.CurrentStatus
}

func (t *TradeData15) SetCurrentStatusSubType(value string) {
	t.CurrentStatusSubType = (*StatusSubType2Code)(&value)
}

func (t *TradeData15) SetCurrentStatusDateTime(value string) {
	t.CurrentStatusDateTime = (*ISODateTime)(&value)
}

func (t *TradeData15) AddPreviousStatus() *Status28Choice {
	t.PreviousStatus = new(Status28Choice)
	return t.PreviousStatus
}

func (t *TradeData15) SetPreviousStatusSubType(value string) {
	t.PreviousStatusSubType = (*StatusSubType2Code)(&value)
}

func (t *TradeData15) SetPreviousStatusDateTime(value string) {
	t.PreviousStatusDateTime = (*ISODateTime)(&value)
}

func (t *TradeData15) SetProductType(value string) {
	t.ProductType = (*Max35Text)(&value)
}

func (t *TradeData15) SetSettlementSessionIdentifier(value string) {
	t.SettlementSessionIdentifier = (*Exact4AlphaNumericText)(&value)
}

func (t *TradeData15) SetSplitTradeIndicator(value string) {
	t.SplitTradeIndicator = (*YesNoIndicator)(&value)
}

// Provides information on the status of a trade.
type TradeData16 struct {

	// Reference to the unique system identification assigned to the trade by the central matching system.
	MatchingSystemUniqueReference *Max35Text `xml:"MtchgSysUnqRef"`

	// Reference to the unique matching identification assigned to the trade and to the matching trade from the counterparty by the central matching system.
	MatchingSystemMatchingReference *Max35Text `xml:"MtchgSysMtchgRef,omitempty"`

	// Unique reference from the central settlement system that allows the removal of alleged trades once the matched status notification for the matching side has been received.
	MatchingSystemMatchedSideReference *Max35Text `xml:"MtchgSysMtchdSdRef,omitempty"`

	// Party that assigned the status to the trade.
	StatusOriginator *Max35Text `xml:"StsOrgtr,omitempty"`

	// Specifies the new status of a trade.
	CurrentStatus *StatusAndSubStatus2 `xml:"CurSts"`

	// Additional information on the current status of a trade in a central system.
	CurrentStatusSubType *StatusSubType2Code `xml:"CurStsSubTp,omitempty"`

	// Specifies the date and time at which the current status was assigned.
	CurrentStatusDateTime *ISODateTime `xml:"CurStsDtTm,omitempty"`

	// Specifies the previous status of a trade.
	PreviousStatus *Status28Choice `xml:"PrvsSts,omitempty"`

	// Additional information on the previous status of a trade in a central system.
	PreviousStatusSubType *StatusSubType2Code `xml:"PrvsStsSubTp,omitempty"`
}

func (t *TradeData16) SetMatchingSystemUniqueReference(value string) {
	t.MatchingSystemUniqueReference = (*Max35Text)(&value)
}

func (t *TradeData16) SetMatchingSystemMatchingReference(value string) {
	t.MatchingSystemMatchingReference = (*Max35Text)(&value)
}

func (t *TradeData16) SetMatchingSystemMatchedSideReference(value string) {
	t.MatchingSystemMatchedSideReference = (*Max35Text)(&value)
}

func (t *TradeData16) SetStatusOriginator(value string) {
	t.StatusOriginator = (*Max35Text)(&value)
}

func (t *TradeData16) AddCurrentStatus() *StatusAndSubStatus2 {
	t.CurrentStatus = new(StatusAndSubStatus2)
	return t.CurrentStatus
}

func (t *TradeData16) SetCurrentStatusSubType(value string) {
	t.CurrentStatusSubType = (*StatusSubType2Code)(&value)
}

func (t *TradeData16) SetCurrentStatusDateTime(value string) {
	t.CurrentStatusDateTime = (*ISODateTime)(&value)
}

func (t *TradeData16) AddPreviousStatus() *Status28Choice {
	t.PreviousStatus = new(Status28Choice)
	return t.PreviousStatus
}

func (t *TradeData16) SetPreviousStatusSubType(value string) {
	t.PreviousStatusSubType = (*StatusSubType2Code)(&value)
}

// Provides information on the status of a trade.
type TradeData7 struct {

	// Identification of the present message assigned by the party issuing the message. This identification must be unique amongst all messages of same type sent by the same party.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Represents the original reference of the instruction for which the status is given, as assigned by the participant that submitted the foreign exchange trade.
	OriginatorReference *Max35Text `xml:"OrgtrRef,omitempty"`

	// Reference to the unique system identification assigned to the trade  by the central matching system.
	MatchingSystemUniqueReference *Max35Text `xml:"MtchgSysUnqRef"`

	// Unique matching identification assigned to the trade and to the matching trade from the counterparty by the central matching system.
	MatchingSystemMatchingReference *Max35Text `xml:"MtchgSysMtchgRef,omitempty"`

	// Identification to the unique reference from the central settlement system that allows the removal of alleged trades once the matched status notification for the matching side has been received.
	MatchingSystemMatchedSideReference *Max35Text `xml:"MtchgSysMtchdSdRef,omitempty"`

	// Party that assigned the status to the foreign exchange or derivative trade.
	StatusOriginator *Max20Text `xml:"StsOrgtr,omitempty"`

	// Specifies the new status of a trade.
	CurrentStatus *StatusAndSubStatus1 `xml:"CurSts"`

	// Additional information on the current status of a trade in a central system.
	CurrentStatusSubType *StatusSubType1Code `xml:"CurStsSubTp,omitempty"`

	// Specifies the date and time at which the current status was assigned.
	CurrentStatusDateTime *ISODateTime `xml:"CurStsDtTm,omitempty"`

	// Specifies the previous status of a trade.
	PreviousStatus *Status6Choice `xml:"PrvsSts,omitempty"`

	// Additional information on the previous status of a trade in a central system.
	PreviousStatusSubType *StatusSubType1Code `xml:"PrvsStsSubTp,omitempty"`

	// Specifies the date and time at which the previous status was assigned.
	PreviousStatusDateTime *ISODateTime `xml:"PrvsStsDtTm,omitempty"`

	// Specifies the product for which the status of the confirmation is reported.
	ProductType *Max35Text `xml:"PdctTp,omitempty"`

	// To indicate the requested CLS Settlement Session that the related trade is part of.
	SettlementSessionIdentifier *Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty"`

	// To indicate if the trade is split.
	SplitTradeIndicator *YesNoIndicator `xml:"SpltTradInd,omitempty"`
}

func (t *TradeData7) SetMessageIdentification(value string) {
	t.MessageIdentification = (*Max35Text)(&value)
}

func (t *TradeData7) SetOriginatorReference(value string) {
	t.OriginatorReference = (*Max35Text)(&value)
}

func (t *TradeData7) SetMatchingSystemUniqueReference(value string) {
	t.MatchingSystemUniqueReference = (*Max35Text)(&value)
}

func (t *TradeData7) SetMatchingSystemMatchingReference(value string) {
	t.MatchingSystemMatchingReference = (*Max35Text)(&value)
}

func (t *TradeData7) SetMatchingSystemMatchedSideReference(value string) {
	t.MatchingSystemMatchedSideReference = (*Max35Text)(&value)
}

func (t *TradeData7) SetStatusOriginator(value string) {
	t.StatusOriginator = (*Max20Text)(&value)
}

func (t *TradeData7) AddCurrentStatus() *StatusAndSubStatus1 {
	t.CurrentStatus = new(StatusAndSubStatus1)
	return t.CurrentStatus
}

func (t *TradeData7) SetCurrentStatusSubType(value string) {
	t.CurrentStatusSubType = (*StatusSubType1Code)(&value)
}

func (t *TradeData7) SetCurrentStatusDateTime(value string) {
	t.CurrentStatusDateTime = (*ISODateTime)(&value)
}

func (t *TradeData7) AddPreviousStatus() *Status6Choice {
	t.PreviousStatus = new(Status6Choice)
	return t.PreviousStatus
}

func (t *TradeData7) SetPreviousStatusSubType(value string) {
	t.PreviousStatusSubType = (*StatusSubType1Code)(&value)
}

func (t *TradeData7) SetPreviousStatusDateTime(value string) {
	t.PreviousStatusDateTime = (*ISODateTime)(&value)
}

func (t *TradeData7) SetProductType(value string) {
	t.ProductType = (*Max35Text)(&value)
}

func (t *TradeData7) SetSettlementSessionIdentifier(value string) {
	t.SettlementSessionIdentifier = (*Exact4AlphaNumericText)(&value)
}

func (t *TradeData7) SetSplitTradeIndicator(value string) {
	t.SplitTradeIndicator = (*YesNoIndicator)(&value)
}

// Provides information on the status of a trade.
type TradeData9 struct {

	// Reference to the unique system identification assigned to the trade by the central matching system.
	MatchingSystemUniqueReference *Max35Text `xml:"MtchgSysUnqRef"`

	// Reference to the unique matching identification assigned to the trade and to the matching trade from the counterparty by the central matching system.
	MatchingSystemMatchingReference *Max35Text `xml:"MtchgSysMtchgRef,omitempty"`

	// Unique reference from the central settlement system that allows the removal of alleged trades once the matched status notification for the matching side has been received.
	MatchingSystemMatchedSideReference *Max35Text `xml:"MtchgSysMtchdSdRef,omitempty"`

	// Party that assigned the status to the trade.
	StatusOriginator *Max35Text `xml:"StsOrgtr,omitempty"`

	// Specifies the new status of the trade.
	CurrentStatus *StatusAndSubStatus1 `xml:"CurSts"`

	// Additional information about the current status of the trade.
	CurrentStatusSubType *StatusSubType1Code `xml:"CurStsSubTp,omitempty"`

	// Specifies the date and time at which the current status was assigned.
	CurrentStatusDateTime *ISODateTime `xml:"CurStsDtTm,omitempty"`

	// Specifies the previous status of the trade.
	PreviousStatus *Status5Choice `xml:"PrvsSts,omitempty"`

	// Specifies whether a trade is alleged or not.
	AllegedTrade *YesNoIndicator `xml:"AllgdTrad,omitempty"`

	// Additional information on the previous status of a trade in a central system.
	PreviousStatusSubType *StatusSubType1Code `xml:"PrvsStsSubTp,omitempty"`
}

func (t *TradeData9) SetMatchingSystemUniqueReference(value string) {
	t.MatchingSystemUniqueReference = (*Max35Text)(&value)
}

func (t *TradeData9) SetMatchingSystemMatchingReference(value string) {
	t.MatchingSystemMatchingReference = (*Max35Text)(&value)
}

func (t *TradeData9) SetMatchingSystemMatchedSideReference(value string) {
	t.MatchingSystemMatchedSideReference = (*Max35Text)(&value)
}

func (t *TradeData9) SetStatusOriginator(value string) {
	t.StatusOriginator = (*Max35Text)(&value)
}

func (t *TradeData9) AddCurrentStatus() *StatusAndSubStatus1 {
	t.CurrentStatus = new(StatusAndSubStatus1)
	return t.CurrentStatus
}

func (t *TradeData9) SetCurrentStatusSubType(value string) {
	t.CurrentStatusSubType = (*StatusSubType1Code)(&value)
}

func (t *TradeData9) SetCurrentStatusDateTime(value string) {
	t.CurrentStatusDateTime = (*ISODateTime)(&value)
}

func (t *TradeData9) AddPreviousStatus() *Status5Choice {
	t.PreviousStatus = new(Status5Choice)
	return t.PreviousStatus
}

func (t *TradeData9) SetAllegedTrade(value string) {
	t.AllegedTrade = (*YesNoIndicator)(&value)
}

func (t *TradeData9) SetPreviousStatusSubType(value string) {
	t.PreviousStatusSubType = (*StatusSubType1Code)(&value)
}
