package iso20022

// Provides information about the notification advice.
type CorporateActionNotification1 struct {

	// Date/time at which the issuer announced that a corporate action event will occur.
	AnnouncementDate *DateFormat4Choice `xml:"AnncmntDt,omitempty"`

	// Date/time at which additional information on the event will be announced, eg, exchange ratio announcement date.
	FurtherDetailedAnnouncementDate *DateFormat4Choice `xml:"FrthrDtldAnncmntDt,omitempty"`

	// Date/time at which the corporate action is legally announced by an official body, eg, publication by a governmental administration.
	OfficialAnnouncementPublicationDate *DateFormat4Choice `xml:"OffclAnncmntPblctnDt,omitempty"`

	// Specifies the status of the details of the event.
	ProcessingStatus *ProcessingStatus1FormatChoice `xml:"PrcgSts"`
}

func (c *CorporateActionNotification1) AddAnnouncementDate() *DateFormat4Choice {
	c.AnnouncementDate = new(DateFormat4Choice)
	return c.AnnouncementDate
}

func (c *CorporateActionNotification1) AddFurtherDetailedAnnouncementDate() *DateFormat4Choice {
	c.FurtherDetailedAnnouncementDate = new(DateFormat4Choice)
	return c.FurtherDetailedAnnouncementDate
}

func (c *CorporateActionNotification1) AddOfficialAnnouncementPublicationDate() *DateFormat4Choice {
	c.OfficialAnnouncementPublicationDate = new(DateFormat4Choice)
	return c.OfficialAnnouncementPublicationDate
}

func (c *CorporateActionNotification1) AddProcessingStatus() *ProcessingStatus1FormatChoice {
	c.ProcessingStatus = new(ProcessingStatus1FormatChoice)
	return c.ProcessingStatus
}

// Corporate action event notification status and contents.
type CorporateActionNotification2 struct {

	// Specifies the type of notification.
	NotificationType *CorporateActionNotificationType1Code `xml:"NtfctnTp"`

	// Specifies the status of the details of the corporate action event.
	ProcessingStatus *CorporateActionProcessingStatus1Choice `xml:"PrcgSts"`

	// Indicates whether the eligible balance is final except for a voluntary corporate action event where it can represent the current eligible balance when communicated before expiration date of that event.
	EligibleBalanceIndicator *YesNoIndicator `xml:"ElgblBalInd,omitempty"`
}

func (c *CorporateActionNotification2) SetNotificationType(value string) {
	c.NotificationType = (*CorporateActionNotificationType1Code)(&value)
}

func (c *CorporateActionNotification2) AddProcessingStatus() *CorporateActionProcessingStatus1Choice {
	c.ProcessingStatus = new(CorporateActionProcessingStatus1Choice)
	return c.ProcessingStatus
}

func (c *CorporateActionNotification2) SetEligibleBalanceIndicator(value string) {
	c.EligibleBalanceIndicator = (*YesNoIndicator)(&value)
}

// Corporate action event notification status and contents.
type CorporateActionNotification3 struct {

	// Specifies the type of notification.
	NotificationType *CorporateActionNotificationType1Code `xml:"NtfctnTp"`

	// Specifies the status of the details of the corporate action event.
	ProcessingStatus *CorporateActionProcessingStatus3Choice `xml:"PrcgSts"`

	// Indicates whether the eligible balance is final except for a voluntary corporate action event where it can represent the current eligible balance when communicated before expiration date of that event.
	EligibleBalanceIndicator *YesNoIndicator `xml:"ElgblBalInd,omitempty"`
}

func (c *CorporateActionNotification3) SetNotificationType(value string) {
	c.NotificationType = (*CorporateActionNotificationType1Code)(&value)
}

func (c *CorporateActionNotification3) AddProcessingStatus() *CorporateActionProcessingStatus3Choice {
	c.ProcessingStatus = new(CorporateActionProcessingStatus3Choice)
	return c.ProcessingStatus
}

func (c *CorporateActionNotification3) SetEligibleBalanceIndicator(value string) {
	c.EligibleBalanceIndicator = (*YesNoIndicator)(&value)
}

// Corporate action event notification status and contents.
type CorporateActionNotification4 struct {

	// Specifies the type of notification.
	NotificationType *CorporateActionNotificationType1Code `xml:"NtfctnTp"`

	// Specifies the status of the details of the corporate action event.
	ProcessingStatus *CorporateActionProcessingStatus4Choice `xml:"PrcgSts"`

	// Indicates whether the eligible balance is final except for a voluntary corporate action event where it can represent the current eligible balance when communicated before expiration date of that event.
	EligibleBalanceIndicator *YesNoIndicator `xml:"ElgblBalInd,omitempty"`
}

func (c *CorporateActionNotification4) SetNotificationType(value string) {
	c.NotificationType = (*CorporateActionNotificationType1Code)(&value)
}

func (c *CorporateActionNotification4) AddProcessingStatus() *CorporateActionProcessingStatus4Choice {
	c.ProcessingStatus = new(CorporateActionProcessingStatus4Choice)
	return c.ProcessingStatus
}

func (c *CorporateActionNotification4) SetEligibleBalanceIndicator(value string) {
	c.EligibleBalanceIndicator = (*YesNoIndicator)(&value)
}

// Corporate action event notification status and contents.
type CorporateActionNotification5 struct {

	// Specifies the type of notification.
	NotificationType *CorporateActionNotificationType1Code `xml:"NtfctnTp"`

	// Specifies the status of the details of the corporate action event.
	ProcessingStatus *CorporateActionProcessingStatus5Choice `xml:"PrcgSts"`

	// Indicates whether the eligible balance is final except for a voluntary corporate action event where it can represent the current eligible balance when communicated before expiration date of that event.
	EligibleBalanceIndicator *YesNoIndicator `xml:"ElgblBalInd,omitempty"`
}

func (c *CorporateActionNotification5) SetNotificationType(value string) {
	c.NotificationType = (*CorporateActionNotificationType1Code)(&value)
}

func (c *CorporateActionNotification5) AddProcessingStatus() *CorporateActionProcessingStatus5Choice {
	c.ProcessingStatus = new(CorporateActionProcessingStatus5Choice)
	return c.ProcessingStatus
}

func (c *CorporateActionNotification5) SetEligibleBalanceIndicator(value string) {
	c.EligibleBalanceIndicator = (*YesNoIndicator)(&value)
}

// Corporate action event notification status and contents.
type CorporateActionNotification6 struct {

	// Specifies the type of notification.
	NotificationType *CorporateActionNotificationType1Code `xml:"NtfctnTp"`

	// Specifies the status of the details of the corporate action event.
	ProcessingStatus *CorporateActionProcessingStatus6Choice `xml:"PrcgSts"`

	// Indicates whether the eligible balance is final except for a voluntary corporate action event where it can represent the current eligible balance when communicated before expiration date of that event.
	EligibleBalanceIndicator *YesNoIndicator `xml:"ElgblBalInd,omitempty"`
}

func (c *CorporateActionNotification6) SetNotificationType(value string) {
	c.NotificationType = (*CorporateActionNotificationType1Code)(&value)
}

func (c *CorporateActionNotification6) AddProcessingStatus() *CorporateActionProcessingStatus6Choice {
	c.ProcessingStatus = new(CorporateActionProcessingStatus6Choice)
	return c.ProcessingStatus
}

func (c *CorporateActionNotification6) SetEligibleBalanceIndicator(value string) {
	c.EligibleBalanceIndicator = (*YesNoIndicator)(&value)
}
