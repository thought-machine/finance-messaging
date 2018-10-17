package iso20022

// Identifies the original notification, to which the cancellation advice refers.
type OriginalNotification10 struct {

	// Point to point reference, as assigned by the original sender, to unambiguously identify the original notification to receive message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Identification of the original notification.
	OriginalNotificationIdentification *Max35Text `xml:"OrgnlNtfctnId"`

	// Indicates whether the cancellation applies to the complete original notification or to individual items within the original notification.
	NotificationCancellation *GroupCancellationIndicator `xml:"NtfctnCxl,omitempty"`

	// Identifies the original notification item, to which the cancellation advice refers.
	OriginalNotificationReference []*OriginalNotificationReference8 `xml:"OrgnlNtfctnRef,omitempty"`
}

func (o *OriginalNotification10) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalNotification10) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalNotification10) SetOriginalNotificationIdentification(value string) {
	o.OriginalNotificationIdentification = (*Max35Text)(&value)
}

func (o *OriginalNotification10) SetNotificationCancellation(value string) {
	o.NotificationCancellation = (*GroupCancellationIndicator)(&value)
}

func (o *OriginalNotification10) AddOriginalNotificationReference() *OriginalNotificationReference8 {
	newValue := new(OriginalNotificationReference8)
	o.OriginalNotificationReference = append(o.OriginalNotificationReference, newValue)
	return newValue
}

// Identify the original notification and to provide the status.
type OriginalNotification3 struct {

	// Point to point reference, as assigned by the original sender, to unambiguously identify the original notification to receive message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Identification of the original notification.
	OriginalNotificationIdentification *Max35Text `xml:"OrgnlNtfctnId"`

	// Specifies the status of the notification in a coded form.
	NotificationStatus *NotificationStatus3Code `xml:"NtfctnSts,omitempty"`

	// Further details of the notification status.
	AdditionalStatusInformation *Max140Text `xml:"AddtlStsInf,omitempty"`

	// Identifies the original notification item and provides the status.
	OriginalNotificationReference []*OriginalNotificationReference2 `xml:"OrgnlNtfctnRef,omitempty"`
}

func (o *OriginalNotification3) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalNotification3) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalNotification3) SetOriginalNotificationIdentification(value string) {
	o.OriginalNotificationIdentification = (*Max35Text)(&value)
}

func (o *OriginalNotification3) SetNotificationStatus(value string) {
	o.NotificationStatus = (*NotificationStatus3Code)(&value)
}

func (o *OriginalNotification3) SetAdditionalStatusInformation(value string) {
	o.AdditionalStatusInformation = (*Max140Text)(&value)
}

func (o *OriginalNotification3) AddOriginalNotificationReference() *OriginalNotificationReference2 {
	newValue := new(OriginalNotificationReference2)
	o.OriginalNotificationReference = append(o.OriginalNotificationReference, newValue)
	return newValue
}

// Identify the original notification, to which the cancellation advice refers.
type OriginalNotification4 struct {

	// Point to point reference, as assigned by the original sender, to unambiguously identify the original notification to receive message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Identification of the original notification.
	OriginalNotificationIdentification *Max35Text `xml:"OrgnlNtfctnId"`

	// Indicates whether the cancellation applies to the complete original notification or to individual items within the original notification.
	NotificationCancellation *GroupCancellationIndicator `xml:"NtfctnCxl,omitempty"`

	// Identifies the original notification item, to which the cancellation advice refers.
	OriginalNotificationReference []*OriginalNotificationReference1 `xml:"OrgnlNtfctnRef,omitempty"`
}

func (o *OriginalNotification4) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalNotification4) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalNotification4) SetOriginalNotificationIdentification(value string) {
	o.OriginalNotificationIdentification = (*Max35Text)(&value)
}

func (o *OriginalNotification4) SetNotificationCancellation(value string) {
	o.NotificationCancellation = (*GroupCancellationIndicator)(&value)
}

func (o *OriginalNotification4) AddOriginalNotificationReference() *OriginalNotificationReference1 {
	newValue := new(OriginalNotificationReference1)
	o.OriginalNotificationReference = append(o.OriginalNotificationReference, newValue)
	return newValue
}

// Identifies the original notification and to provide the status.
type OriginalNotification5 struct {

	// Point to point reference, as assigned by the original sender, to unambiguously identify the original notification to receive message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Identification of the original notification.
	OriginalNotificationIdentification *Max35Text `xml:"OrgnlNtfctnId"`

	// Specifies the status of the notification in a coded form.
	NotificationStatus *NotificationStatus3Code `xml:"NtfctnSts,omitempty"`

	// Further details of the notification status.
	AdditionalStatusInformation *Max140Text `xml:"AddtlStsInf,omitempty"`

	// Identifies the original notification item and provides the status.
	OriginalNotificationReference []*OriginalNotificationReference3 `xml:"OrgnlNtfctnRef,omitempty"`
}

func (o *OriginalNotification5) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalNotification5) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalNotification5) SetOriginalNotificationIdentification(value string) {
	o.OriginalNotificationIdentification = (*Max35Text)(&value)
}

func (o *OriginalNotification5) SetNotificationStatus(value string) {
	o.NotificationStatus = (*NotificationStatus3Code)(&value)
}

func (o *OriginalNotification5) SetAdditionalStatusInformation(value string) {
	o.AdditionalStatusInformation = (*Max140Text)(&value)
}

func (o *OriginalNotification5) AddOriginalNotificationReference() *OriginalNotificationReference3 {
	newValue := new(OriginalNotificationReference3)
	o.OriginalNotificationReference = append(o.OriginalNotificationReference, newValue)
	return newValue
}

// Identifies the original notification, to which the cancellation advice refers.
type OriginalNotification6 struct {

	// Point to point reference, as assigned by the original sender, to unambiguously identify the original notification to receive message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Identification of the original notification.
	OriginalNotificationIdentification *Max35Text `xml:"OrgnlNtfctnId"`

	// Indicates whether the cancellation applies to the complete original notification or to individual items within the original notification.
	NotificationCancellation *GroupCancellationIndicator `xml:"NtfctnCxl,omitempty"`

	// Identifies the original notification item, to which the cancellation advice refers.
	OriginalNotificationReference []*OriginalNotificationReference4 `xml:"OrgnlNtfctnRef,omitempty"`
}

func (o *OriginalNotification6) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalNotification6) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalNotification6) SetOriginalNotificationIdentification(value string) {
	o.OriginalNotificationIdentification = (*Max35Text)(&value)
}

func (o *OriginalNotification6) SetNotificationCancellation(value string) {
	o.NotificationCancellation = (*GroupCancellationIndicator)(&value)
}

func (o *OriginalNotification6) AddOriginalNotificationReference() *OriginalNotificationReference4 {
	newValue := new(OriginalNotificationReference4)
	o.OriginalNotificationReference = append(o.OriginalNotificationReference, newValue)
	return newValue
}

// Identifies the original notification and to provide the status.
type OriginalNotification7 struct {

	// Point to point reference, as assigned by the original sender, to unambiguously identify the original notification to receive message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Identification of the original notification.
	OriginalNotificationIdentification *Max35Text `xml:"OrgnlNtfctnId"`

	// Specifies the status of the notification in a coded form.
	NotificationStatus *NotificationStatus3Code `xml:"NtfctnSts,omitempty"`

	// Further details of the notification status.
	AdditionalStatusInformation *Max140Text `xml:"AddtlStsInf,omitempty"`

	// Identifies the original notification item and provides the status.
	OriginalNotificationReference []*OriginalNotificationReference5 `xml:"OrgnlNtfctnRef,omitempty"`
}

func (o *OriginalNotification7) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalNotification7) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalNotification7) SetOriginalNotificationIdentification(value string) {
	o.OriginalNotificationIdentification = (*Max35Text)(&value)
}

func (o *OriginalNotification7) SetNotificationStatus(value string) {
	o.NotificationStatus = (*NotificationStatus3Code)(&value)
}

func (o *OriginalNotification7) SetAdditionalStatusInformation(value string) {
	o.AdditionalStatusInformation = (*Max140Text)(&value)
}

func (o *OriginalNotification7) AddOriginalNotificationReference() *OriginalNotificationReference5 {
	newValue := new(OriginalNotificationReference5)
	o.OriginalNotificationReference = append(o.OriginalNotificationReference, newValue)
	return newValue
}

// Identifies the original notification, to which the cancellation advice refers.
type OriginalNotification8 struct {

	// Point to point reference, as assigned by the original sender, to unambiguously identify the original notification to receive message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Identification of the original notification.
	OriginalNotificationIdentification *Max35Text `xml:"OrgnlNtfctnId"`

	// Indicates whether the cancellation applies to the complete original notification or to individual items within the original notification.
	NotificationCancellation *GroupCancellationIndicator `xml:"NtfctnCxl,omitempty"`

	// Identifies the original notification item, to which the cancellation advice refers.
	OriginalNotificationReference []*OriginalNotificationReference6 `xml:"OrgnlNtfctnRef,omitempty"`
}

func (o *OriginalNotification8) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalNotification8) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalNotification8) SetOriginalNotificationIdentification(value string) {
	o.OriginalNotificationIdentification = (*Max35Text)(&value)
}

func (o *OriginalNotification8) SetNotificationCancellation(value string) {
	o.NotificationCancellation = (*GroupCancellationIndicator)(&value)
}

func (o *OriginalNotification8) AddOriginalNotificationReference() *OriginalNotificationReference6 {
	newValue := new(OriginalNotificationReference6)
	o.OriginalNotificationReference = append(o.OriginalNotificationReference, newValue)
	return newValue
}

// Identifies the original notification and to provide the status.
type OriginalNotification9 struct {

	// Point to point reference, as assigned by the original sender, to unambiguously identify the original notification to receive message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Identification of the original notification.
	OriginalNotificationIdentification *Max35Text `xml:"OrgnlNtfctnId"`

	// Specifies the status of the notification in a coded form.
	NotificationStatus *NotificationStatus3Code `xml:"NtfctnSts,omitempty"`

	// Further details of the notification status.
	AdditionalStatusInformation *Max140Text `xml:"AddtlStsInf,omitempty"`

	// Identifies the original notification item and provides the status.
	OriginalNotificationReference []*OriginalNotificationReference7 `xml:"OrgnlNtfctnRef,omitempty"`
}

func (o *OriginalNotification9) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalNotification9) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalNotification9) SetOriginalNotificationIdentification(value string) {
	o.OriginalNotificationIdentification = (*Max35Text)(&value)
}

func (o *OriginalNotification9) SetNotificationStatus(value string) {
	o.NotificationStatus = (*NotificationStatus3Code)(&value)
}

func (o *OriginalNotification9) SetAdditionalStatusInformation(value string) {
	o.AdditionalStatusInformation = (*Max140Text)(&value)
}

func (o *OriginalNotification9) AddOriginalNotificationReference() *OriginalNotificationReference7 {
	newValue := new(OriginalNotificationReference7)
	o.OriginalNotificationReference = append(o.OriginalNotificationReference, newValue)
	return newValue
}
