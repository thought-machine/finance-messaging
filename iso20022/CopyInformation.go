package iso20022

// Information provided when the message is a copy of a previous message.
type CopyInformation1 struct {

	// Indicates whether the message is a copy.
	CopyIndicator *YesNoIndicator `xml:"CpyInd"`

	// Original receiver of the message, if this message is a copy.
	OriginalReceiver *BICIdentification1 `xml:"OrgnlRcvr"`
}

func (c *CopyInformation1) SetCopyIndicator(value string) {
	c.CopyIndicator = (*YesNoIndicator)(&value)
}

func (c *CopyInformation1) AddOriginalReceiver() *BICIdentification1 {
	c.OriginalReceiver = new(BICIdentification1)
	return c.OriginalReceiver
}

// Information provided when the message is a copy of a previous message.
type CopyInformation2 struct {

	// Indicates whether the message is a copy.
	CopyIndicator *YesNoIndicator `xml:"CpyInd"`

	// Original receiver of the message, if this message is a copy.
	OriginalReceiver *BICIdentification1 `xml:"OrgnlRcvr,omitempty"`
}

func (c *CopyInformation2) SetCopyIndicator(value string) {
	c.CopyIndicator = (*YesNoIndicator)(&value)
}

func (c *CopyInformation2) AddOriginalReceiver() *BICIdentification1 {
	c.OriginalReceiver = new(BICIdentification1)
	return c.OriginalReceiver
}

// Information provided when the message is a copy of a previous message.
type CopyInformation4 struct {

	// Indicates whether the message is a copy.
	CopyIndicator *YesNoIndicator `xml:"CpyInd"`

	// Original receiver of the message, if this message is a copy.
	OriginalReceiver *AnyBICIdentifier `xml:"OrgnlRcvr,omitempty"`
}

func (c *CopyInformation4) SetCopyIndicator(value string) {
	c.CopyIndicator = (*YesNoIndicator)(&value)
}

func (c *CopyInformation4) SetOriginalReceiver(value string) {
	c.OriginalReceiver = (*AnyBICIdentifier)(&value)
}
