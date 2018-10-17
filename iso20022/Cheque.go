package iso20022

// Set of characteristics related to a cheque instruction, such as cheque type or cheque number.
type Cheque3 struct {

	// Unique and unambiguous identifier for a cheque as assigned by the financial institution.
	Number *Max35Text `xml:"Nb,omitempty"`

	// Party to which a cheque is made payable.
	PayeeIdentification *PartyIdentification2Choice `xml:"PyeeId"`

	// Financial institution on which a cheque is drawn, ie, the financial institution that services the account of the entity that issued the cheque.
	DraweeIdentification *FinancialInstitutionIdentification3Choice `xml:"DrweeId,omitempty"`

	// Account owner that issues a cheque ordering the drawee bank to pay a specific amount, upon demand, to the payee.
	DrawerIdentification *PartyIdentification2Choice `xml:"DrwrId,omitempty"`
}

func (c *Cheque3) SetNumber(value string) {
	c.Number = (*Max35Text)(&value)
}

func (c *Cheque3) AddPayeeIdentification() *PartyIdentification2Choice {
	c.PayeeIdentification = new(PartyIdentification2Choice)
	return c.PayeeIdentification
}

func (c *Cheque3) AddDraweeIdentification() *FinancialInstitutionIdentification3Choice {
	c.DraweeIdentification = new(FinancialInstitutionIdentification3Choice)
	return c.DraweeIdentification
}

func (c *Cheque3) AddDrawerIdentification() *PartyIdentification2Choice {
	c.DrawerIdentification = new(PartyIdentification2Choice)
	return c.DrawerIdentification
}

// Set of characteristics related to a cheque instruction, such as cheque type or cheque number.
type Cheque4 struct {

	// Party to which a cheque is made payable.
	PayeeIdentification *NameAndAddress5 `xml:"PyeeId"`
}

func (c *Cheque4) AddPayeeIdentification() *NameAndAddress5 {
	c.PayeeIdentification = new(NameAndAddress5)
	return c.PayeeIdentification
}

// Set of characteristics related to a cheque instruction, such as cheque type or cheque number.
type Cheque5 struct {

	// Specifies the type of cheque to be issued by the first agent.
	ChequeType *ChequeType2Code `xml:"ChqTp,omitempty"`

	// Identifies the cheque number.
	ChequeNumber *Max35Text `xml:"ChqNb,omitempty"`

	// Identifies the party that ordered the issuance of the cheque.
	ChequeFrom *NameAndAddress3 `xml:"ChqFr,omitempty"`

	// Specifies the delivery method of the cheque by the debtor's agent.
	DeliveryMethod *ChequeDeliveryMethod1Choice `xml:"DlvryMtd,omitempty"`

	// Identifies the party to whom the debtor's agent should send the cheque.
	DeliverTo *NameAndAddress3 `xml:"DlvrTo,omitempty"`

	// Urgency or order of importance that the originator would like the recipient of the payment instruction to apply to the processing of the payment instruction.
	InstructionPriority *Priority2Code `xml:"InstrPrty,omitempty"`

	// Date when the draft becomes payable and the debtor's account is debited.
	ChequeMaturityDate *ISODate `xml:"ChqMtrtyDt,omitempty"`

	// Code agreed between the initiating party and the debtor's agent, that specifies the cheque layout, company logo and digitised signature to be used to print the cheque.
	FormsCode *Max35Text `xml:"FrmsCd,omitempty"`

	// Information that needs to be printed on a cheque, used by the payer to add miscellaneous information.
	MemoField *Max35Text `xml:"MemoFld,omitempty"`

	// Regional area in which the cheque can be cleared, when a country has no nation-wide cheque clearing organisation.
	RegionalClearingZone *Max35Text `xml:"RgnlClrZone,omitempty"`

	// Specifies the print location of the cheque.
	PrintLocation *Max35Text `xml:"PrtLctn,omitempty"`
}

func (c *Cheque5) SetChequeType(value string) {
	c.ChequeType = (*ChequeType2Code)(&value)
}

func (c *Cheque5) SetChequeNumber(value string) {
	c.ChequeNumber = (*Max35Text)(&value)
}

func (c *Cheque5) AddChequeFrom() *NameAndAddress3 {
	c.ChequeFrom = new(NameAndAddress3)
	return c.ChequeFrom
}

func (c *Cheque5) AddDeliveryMethod() *ChequeDeliveryMethod1Choice {
	c.DeliveryMethod = new(ChequeDeliveryMethod1Choice)
	return c.DeliveryMethod
}

func (c *Cheque5) AddDeliverTo() *NameAndAddress3 {
	c.DeliverTo = new(NameAndAddress3)
	return c.DeliverTo
}

func (c *Cheque5) SetInstructionPriority(value string) {
	c.InstructionPriority = (*Priority2Code)(&value)
}

func (c *Cheque5) SetChequeMaturityDate(value string) {
	c.ChequeMaturityDate = (*ISODate)(&value)
}

func (c *Cheque5) SetFormsCode(value string) {
	c.FormsCode = (*Max35Text)(&value)
}

func (c *Cheque5) SetMemoField(value string) {
	c.MemoField = (*Max35Text)(&value)
}

func (c *Cheque5) SetRegionalClearingZone(value string) {
	c.RegionalClearingZone = (*Max35Text)(&value)
}

func (c *Cheque5) SetPrintLocation(value string) {
	c.PrintLocation = (*Max35Text)(&value)
}

// Set of characteristics related to a cheque instruction, such as cheque type or cheque number.
type Cheque6 struct {

	// Specifies the type of cheque to be issued.
	ChequeType *ChequeType2Code `xml:"ChqTp,omitempty"`

	// Unique and unambiguous identifier for a cheque as assigned by the agent.
	ChequeNumber *Max35Text `xml:"ChqNb,omitempty"`

	// Identifies the party that ordered the issuance of the cheque.
	ChequeFrom *NameAndAddress10 `xml:"ChqFr,omitempty"`

	// Specifies the delivery method of the cheque by the debtor's agent.
	DeliveryMethod *ChequeDeliveryMethod1Choice `xml:"DlvryMtd,omitempty"`

	// Party to whom the debtor's agent needs to send the cheque.
	DeliverTo *NameAndAddress10 `xml:"DlvrTo,omitempty"`

	// Urgency or order of importance that the originator would like the recipient of the payment instruction to apply to the processing of the payment instruction.
	InstructionPriority *Priority2Code `xml:"InstrPrty,omitempty"`

	// Date when the draft becomes payable and the debtor's account is debited.
	ChequeMaturityDate *ISODate `xml:"ChqMtrtyDt,omitempty"`

	// Identifies, in a coded form, the cheque layout, company logo and digitised signature to be used to print the cheque, as agreed between the initiating party and the debtor's agent.
	FormsCode *Max35Text `xml:"FrmsCd,omitempty"`

	// Information that needs to be printed on a cheque, used by the payer to add miscellaneous information.
	MemoField []*Max35Text `xml:"MemoFld,omitempty"`

	// Regional area in which the cheque can be cleared, when a country has no nation-wide cheque clearing organisation.
	RegionalClearingZone *Max35Text `xml:"RgnlClrZone,omitempty"`

	// Specifies the print location of the cheque.
	PrintLocation *Max35Text `xml:"PrtLctn,omitempty"`
}

func (c *Cheque6) SetChequeType(value string) {
	c.ChequeType = (*ChequeType2Code)(&value)
}

func (c *Cheque6) SetChequeNumber(value string) {
	c.ChequeNumber = (*Max35Text)(&value)
}

func (c *Cheque6) AddChequeFrom() *NameAndAddress10 {
	c.ChequeFrom = new(NameAndAddress10)
	return c.ChequeFrom
}

func (c *Cheque6) AddDeliveryMethod() *ChequeDeliveryMethod1Choice {
	c.DeliveryMethod = new(ChequeDeliveryMethod1Choice)
	return c.DeliveryMethod
}

func (c *Cheque6) AddDeliverTo() *NameAndAddress10 {
	c.DeliverTo = new(NameAndAddress10)
	return c.DeliverTo
}

func (c *Cheque6) SetInstructionPriority(value string) {
	c.InstructionPriority = (*Priority2Code)(&value)
}

func (c *Cheque6) SetChequeMaturityDate(value string) {
	c.ChequeMaturityDate = (*ISODate)(&value)
}

func (c *Cheque6) SetFormsCode(value string) {
	c.FormsCode = (*Max35Text)(&value)
}

func (c *Cheque6) AddMemoField(value string) {
	c.MemoField = append(c.MemoField, (*Max35Text)(&value))
}

func (c *Cheque6) SetRegionalClearingZone(value string) {
	c.RegionalClearingZone = (*Max35Text)(&value)
}

func (c *Cheque6) SetPrintLocation(value string) {
	c.PrintLocation = (*Max35Text)(&value)
}

// Set of characteristics related to a cheque instruction, such as cheque type or cheque number.
type Cheque7 struct {

	// Specifies the type of cheque to be issued.
	ChequeType *ChequeType2Code `xml:"ChqTp,omitempty"`

	// Unique and unambiguous identifier for a cheque as assigned by the agent.
	ChequeNumber *Max35Text `xml:"ChqNb,omitempty"`

	// Identifies the party that ordered the issuance of the cheque.
	ChequeFrom *NameAndAddress10 `xml:"ChqFr,omitempty"`

	// Specifies the delivery method of the cheque by the debtor's agent.
	DeliveryMethod *ChequeDeliveryMethod1Choice `xml:"DlvryMtd,omitempty"`

	// Party to whom the debtor's agent needs to send the cheque.
	DeliverTo *NameAndAddress10 `xml:"DlvrTo,omitempty"`

	// Urgency or order of importance that the originator would like the recipient of the payment instruction to apply to the processing of the payment instruction.
	InstructionPriority *Priority2Code `xml:"InstrPrty,omitempty"`

	// Date when the draft becomes payable and the debtor's account is debited.
	ChequeMaturityDate *ISODate `xml:"ChqMtrtyDt,omitempty"`

	// Identifies, in a coded form, the cheque layout, company logo and digitised signature to be used to print the cheque, as agreed between the initiating party and the debtor's agent.
	FormsCode *Max35Text `xml:"FrmsCd,omitempty"`

	// Information that needs to be printed on a cheque, used by the payer to add miscellaneous information.
	MemoField []*Max35Text `xml:"MemoFld,omitempty"`

	// Regional area in which the cheque can be cleared, when a country has no nation-wide cheque clearing organisation.
	RegionalClearingZone *Max35Text `xml:"RgnlClrZone,omitempty"`

	// Specifies the print location of the cheque.
	PrintLocation *Max35Text `xml:"PrtLctn,omitempty"`

	// Signature to be used by the cheque servicer on a specific cheque to be printed.
	Signature []*Max70Text `xml:"Sgntr,omitempty"`
}

func (c *Cheque7) SetChequeType(value string) {
	c.ChequeType = (*ChequeType2Code)(&value)
}

func (c *Cheque7) SetChequeNumber(value string) {
	c.ChequeNumber = (*Max35Text)(&value)
}

func (c *Cheque7) AddChequeFrom() *NameAndAddress10 {
	c.ChequeFrom = new(NameAndAddress10)
	return c.ChequeFrom
}

func (c *Cheque7) AddDeliveryMethod() *ChequeDeliveryMethod1Choice {
	c.DeliveryMethod = new(ChequeDeliveryMethod1Choice)
	return c.DeliveryMethod
}

func (c *Cheque7) AddDeliverTo() *NameAndAddress10 {
	c.DeliverTo = new(NameAndAddress10)
	return c.DeliverTo
}

func (c *Cheque7) SetInstructionPriority(value string) {
	c.InstructionPriority = (*Priority2Code)(&value)
}

func (c *Cheque7) SetChequeMaturityDate(value string) {
	c.ChequeMaturityDate = (*ISODate)(&value)
}

func (c *Cheque7) SetFormsCode(value string) {
	c.FormsCode = (*Max35Text)(&value)
}

func (c *Cheque7) AddMemoField(value string) {
	c.MemoField = append(c.MemoField, (*Max35Text)(&value))
}

func (c *Cheque7) SetRegionalClearingZone(value string) {
	c.RegionalClearingZone = (*Max35Text)(&value)
}

func (c *Cheque7) SetPrintLocation(value string) {
	c.PrintLocation = (*Max35Text)(&value)
}

func (c *Cheque7) AddSignature(value string) {
	c.Signature = append(c.Signature, (*Max70Text)(&value))
}

// Set of characteristics related to a cheque instruction, such as cheque type or cheque number.
type Cheque9 struct {

	// Unique and unambiguous identifier for a cheque as assigned by the financial institution.
	Number *Max35Text `xml:"Nb,omitempty"`

	// Party to which a cheque is made payable.
	PayeeIdentification *PartyIdentification113 `xml:"PyeeId"`

	// Financial institution on which a cheque is drawn, that is, the financial institution that services the account of the entity that issued the cheque.
	DraweeIdentification *FinancialInstitutionIdentification10 `xml:"DrweeId,omitempty"`

	// Account owner that issues a cheque ordering the drawee bank to pay a specific amount, upon demand, to the payee.
	DrawerIdentification *PartyIdentification113 `xml:"DrwrId,omitempty"`
}

func (c *Cheque9) SetNumber(value string) {
	c.Number = (*Max35Text)(&value)
}

func (c *Cheque9) AddPayeeIdentification() *PartyIdentification113 {
	c.PayeeIdentification = new(PartyIdentification113)
	return c.PayeeIdentification
}

func (c *Cheque9) AddDraweeIdentification() *FinancialInstitutionIdentification10 {
	c.DraweeIdentification = new(FinancialInstitutionIdentification10)
	return c.DraweeIdentification
}

func (c *Cheque9) AddDrawerIdentification() *PartyIdentification113 {
	c.DrawerIdentification = new(PartyIdentification113)
	return c.DrawerIdentification
}
