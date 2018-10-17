package iso20022

// Provides details on the original group, to which the message refers.
type OriginalGroupHeader1 struct {

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Number of individual transactions contained in the original message.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original message, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of a group of transactions.
	GroupStatus *TransactionGroupStatus3Code `xml:"GrpSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical transaction status.
	NumberOfTransactionsPerStatus []*NumberOfTransactionsPerStatus3 `xml:"NbOfTxsPerSts,omitempty"`
}

func (o *OriginalGroupHeader1) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader1) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader1) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupHeader1) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalGroupHeader1) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalGroupHeader1) SetGroupStatus(value string) {
	o.GroupStatus = (*TransactionGroupStatus3Code)(&value)
}

func (o *OriginalGroupHeader1) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	o.StatusReasonInformation = append(o.StatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalGroupHeader1) AddNumberOfTransactionsPerStatus() *NumberOfTransactionsPerStatus3 {
	newValue := new(NumberOfTransactionsPerStatus3)
	o.NumberOfTransactionsPerStatus = append(o.NumberOfTransactionsPerStatus, newValue)
	return newValue
}

// Provides information on the original group, to which the message refers.
type OriginalGroupHeader2 struct {

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Provides detailed information on the return reason.
	ReturnReasonInformation []*PaymentReturnReason1 `xml:"RtrRsnInf,omitempty"`
}

func (o *OriginalGroupHeader2) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader2) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader2) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupHeader2) AddReturnReasonInformation() *PaymentReturnReason1 {
	newValue := new(PaymentReturnReason1)
	o.ReturnReasonInformation = append(o.ReturnReasonInformation, newValue)
	return newValue
}

// Provides details on the original group, to which the message refers.
type OriginalGroupHeader3 struct {

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Provides detailed information on the reversal reason.
	ReversalReasonInformation []*PaymentReversalReason7 `xml:"RvslRsnInf,omitempty"`
}

func (o *OriginalGroupHeader3) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader3) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader3) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupHeader3) AddReversalReasonInformation() *PaymentReversalReason7 {
	newValue := new(PaymentReversalReason7)
	o.ReversalReasonInformation = append(o.ReversalReasonInformation, newValue)
	return newValue
}

// Provides details on the original group, to which the message refers.
type OriginalGroupHeader4 struct {

	// Unique identification, as assigned by the assigner, to unambiguously identify the group cancellation request.
	//
	// Usage: The group cancellation request identification can be used for reconciliation or to link tasks related to the cancellation request.
	GroupCancellationIdentification *Max35Text `xml:"GrpCxlId,omitempty"`

	// Uniquely and unambiguously identifies an exception or an investigation workflow.
	Case *Case3 `xml:"Case,omitempty"`

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Number of individual transactions contained in the original message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs,omitempty"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Indicates whether the cancellation request applies to a whole group of transactions or to individual transactions within an original group.
	GroupCancellation *GroupCancellationIndicator `xml:"GrpCxl,omitempty"`

	// Provides detailed information on the cancellation reason.
	CancellationReasonInformation []*PaymentCancellationReason2 `xml:"CxlRsnInf,omitempty"`
}

func (o *OriginalGroupHeader4) SetGroupCancellationIdentification(value string) {
	o.GroupCancellationIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader4) AddCase() *Case3 {
	o.Case = new(Case3)
	return o.Case
}

func (o *OriginalGroupHeader4) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader4) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader4) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupHeader4) SetNumberOfTransactions(value string) {
	o.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalGroupHeader4) SetControlSum(value string) {
	o.ControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalGroupHeader4) SetGroupCancellation(value string) {
	o.GroupCancellation = (*GroupCancellationIndicator)(&value)
}

func (o *OriginalGroupHeader4) AddCancellationReasonInformation() *PaymentCancellationReason2 {
	newValue := new(PaymentCancellationReason2)
	o.CancellationReasonInformation = append(o.CancellationReasonInformation, newValue)
	return newValue
}

// Provides the details on the original group, to which the message refers.
type OriginalGroupHeader5 struct {

	// Unique identification, as assigned by the original assigner, to unambiguously identify the original group cancellation request.
	OriginalGroupCancellationIdentification *Max35Text `xml:"OrgnlGrpCxlId,omitempty"`

	// Identifies the case.
	ResolvedCase *Case3 `xml:"RslvdCase,omitempty"`

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Number of individual transactions contained in the original message.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original message, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of a group cancellation request.
	GroupCancellationStatus *GroupCancellationStatus1Code `xml:"GrpCxlSts,omitempty"`

	// Provides detailed information on the cancellation status reason.
	CancellationStatusReasonInformation []*CancellationStatusReason2 `xml:"CxlStsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical cancellation status.
	NumberOfTransactionsPerCancellationStatus []*NumberOfTransactionsPerStatus1 `xml:"NbOfTxsPerCxlSts,omitempty"`
}

func (o *OriginalGroupHeader5) SetOriginalGroupCancellationIdentification(value string) {
	o.OriginalGroupCancellationIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader5) AddResolvedCase() *Case3 {
	o.ResolvedCase = new(Case3)
	return o.ResolvedCase
}

func (o *OriginalGroupHeader5) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader5) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader5) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupHeader5) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalGroupHeader5) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalGroupHeader5) SetGroupCancellationStatus(value string) {
	o.GroupCancellationStatus = (*GroupCancellationStatus1Code)(&value)
}

func (o *OriginalGroupHeader5) AddCancellationStatusReasonInformation() *CancellationStatusReason2 {
	newValue := new(CancellationStatusReason2)
	o.CancellationStatusReasonInformation = append(o.CancellationStatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalGroupHeader5) AddNumberOfTransactionsPerCancellationStatus() *NumberOfTransactionsPerStatus1 {
	newValue := new(NumberOfTransactionsPerStatus1)
	o.NumberOfTransactionsPerCancellationStatus = append(o.NumberOfTransactionsPerCancellationStatus, newValue)
	return newValue
}

// Provides details on the original group, to which the message refers.
type OriginalGroupHeader6 struct {

	// Unique identification, as assigned by the assigner, to unambiguously identify the group cancellation request.
	//
	// Usage: The group cancellation request identification can be used for reconciliation or to link tasks related to the cancellation request.
	GroupCancellationIdentification *Max35Text `xml:"GrpCxlId,omitempty"`

	// Uniquely and unambiguously identifies an exception or an investigation workflow.
	Case *Case3 `xml:"Case,omitempty"`

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Number of individual transactions contained in the original message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs,omitempty"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Indicates whether the cancellation request applies to a whole group of transactions or to individual transactions within an original group.
	GroupCancellation *GroupCancellationIndicator `xml:"GrpCxl,omitempty"`

	// Provides detailed information on the cancellation reason.
	CancellationReasonInformation []*PaymentCancellationReason3 `xml:"CxlRsnInf,omitempty"`
}

func (o *OriginalGroupHeader6) SetGroupCancellationIdentification(value string) {
	o.GroupCancellationIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader6) AddCase() *Case3 {
	o.Case = new(Case3)
	return o.Case
}

func (o *OriginalGroupHeader6) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader6) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader6) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupHeader6) SetNumberOfTransactions(value string) {
	o.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalGroupHeader6) SetControlSum(value string) {
	o.ControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalGroupHeader6) SetGroupCancellation(value string) {
	o.GroupCancellation = (*GroupCancellationIndicator)(&value)
}

func (o *OriginalGroupHeader6) AddCancellationReasonInformation() *PaymentCancellationReason3 {
	newValue := new(PaymentCancellationReason3)
	o.CancellationReasonInformation = append(o.CancellationReasonInformation, newValue)
	return newValue
}

// Provides details on the original group, to which the message refers.
type OriginalGroupHeader7 struct {

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Number of individual transactions contained in the original message.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original message, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of a group of transactions.
	GroupStatus *ExternalPaymentGroupStatus1Code `xml:"GrpSts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical transaction status.
	NumberOfTransactionsPerStatus []*NumberOfTransactionsPerStatus5 `xml:"NbOfTxsPerSts,omitempty"`
}

func (o *OriginalGroupHeader7) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader7) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader7) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupHeader7) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalGroupHeader7) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalGroupHeader7) SetGroupStatus(value string) {
	o.GroupStatus = (*ExternalPaymentGroupStatus1Code)(&value)
}

func (o *OriginalGroupHeader7) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	o.StatusReasonInformation = append(o.StatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalGroupHeader7) AddNumberOfTransactionsPerStatus() *NumberOfTransactionsPerStatus5 {
	newValue := new(NumberOfTransactionsPerStatus5)
	o.NumberOfTransactionsPerStatus = append(o.NumberOfTransactionsPerStatus, newValue)
	return newValue
}
