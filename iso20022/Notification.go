package iso20022

// Structured information to be communicated to other parties in the transaction.
type Notification1 struct {

	// Type of the notification.
	Type *NotificationType1Code `xml:"Tp"`

	// Additional and important information to qualify and describe the notification.
	AdditionalInformation *Max140Text `xml:"AddtlInf"`
}

func (n *Notification1) SetType(value string) {
	n.Type = (*NotificationType1Code)(&value)
}

func (n *Notification1) SetAdditionalInformation(value string) {
	n.AdditionalInformation = (*Max140Text)(&value)
}

// Information about the type of notification required.
type Notification2 struct {

	// Type of notification.
	NotificationType *Max35Text `xml:"NtfctnTp"`

	// Indicates whether the notification is required.
	Required *YesNoIndicator `xml:"Reqrd"`

	// Specifies how the notification is sent.
	DistributionType *InformationDistribution1Choice `xml:"DstrbtnTp,omitempty"`
}

func (n *Notification2) SetNotificationType(value string) {
	n.NotificationType = (*Max35Text)(&value)
}

func (n *Notification2) SetRequired(value string) {
	n.Required = (*YesNoIndicator)(&value)
}

func (n *Notification2) AddDistributionType() *InformationDistribution1Choice {
	n.DistributionType = new(InformationDistribution1Choice)
	return n.DistributionType
}
