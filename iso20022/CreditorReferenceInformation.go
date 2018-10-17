package iso20022

// Structured information supplied to enable the matching, i.e.  reconciliation, of a payment with the items that the payment is intended to settle, eg, commercial invoices in an accounts receivable system.
type CreditorReferenceInformation1 struct {

	// Provides the type of the creditor reference.
	CreditorReferenceType *CreditorReferenceType1 `xml:"CdtrRefTp,omitempty"`

	// Unique and unambiguous reference assigned by the creditor to refer to the payment transaction.
	//
	// Usage: if available, the initiating party should provide this reference in the structured remittance information, to enable reconciliation by the creditor upon receipt of the cash.
	//
	// If the business context requires the use of a creditor reference or a payment remit identification, and only one identifier can be passed through the end-to-end chain, the creditor's reference or payment remittance identification should be quoted in the end-to-end transaction identification.
	CreditorReference *Max35Text `xml:"CdtrRef,omitempty"`
}

func (c *CreditorReferenceInformation1) AddCreditorReferenceType() *CreditorReferenceType1 {
	c.CreditorReferenceType = new(CreditorReferenceType1)
	return c.CreditorReferenceType
}

func (c *CreditorReferenceInformation1) SetCreditorReference(value string) {
	c.CreditorReference = (*Max35Text)(&value)
}

// Reference information provided by the creditor to allow the identification of the underlying documents.
type CreditorReferenceInformation2 struct {

	// Specifies the type of creditor reference.
	Type *CreditorReferenceType2 `xml:"Tp,omitempty"`

	// Unique reference, as assigned by the creditor, to unambiguously refer to the payment transaction.
	//
	// Usage: If available, the initiating party should provide this reference in the structured remittance information, to enable reconciliation by the creditor upon receipt of the amount of money.
	//
	// If the business context requires the use of a creditor reference or a payment remit identification, and only one identifier can be passed through the end-to-end chain, the creditor's reference or payment remittance identification should be quoted in the end-to-end transaction identification.
	Reference *Max35Text `xml:"Ref,omitempty"`
}

func (c *CreditorReferenceInformation2) AddType() *CreditorReferenceType2 {
	c.Type = new(CreditorReferenceType2)
	return c.Type
}

func (c *CreditorReferenceInformation2) SetReference(value string) {
	c.Reference = (*Max35Text)(&value)
}
