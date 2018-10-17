package iso20022

// Identification of the transaction in an unambiguous way.
type TransactionIdentifier1 struct {

	// Local date and time of the transaction assigned by the POI (Point Of Interaction).
	TransactionDateTime *ISODateTime `xml:"TxDtTm"`

	// Identification of the transaction that has to be unique for a time period.
	TransactionReference *Max35Text `xml:"TxRef"`
}

func (t *TransactionIdentifier1) SetTransactionDateTime(value string) {
	t.TransactionDateTime = (*ISODateTime)(&value)
}

func (t *TransactionIdentifier1) SetTransactionReference(value string) {
	t.TransactionReference = (*Max35Text)(&value)
}

// Identification of the reconciliation period between the acquirer and the issuer or their respective agents.
type TransactionIdentifier2 struct {

	// Date of the reconciliation.
	// It correspond to the ISO 8583 field number 28 for the versions 1993 and 2003.
	ReconciliationDate *ISODate `xml:"RcncltnDt"`

	// Identification of the reconciliation.
	// It correspond to the ISO 8583 field number 29 for the versions 1993 and 2003.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`
}

func (t *TransactionIdentifier2) SetReconciliationDate(value string) {
	t.ReconciliationDate = (*ISODate)(&value)
}

func (t *TransactionIdentifier2) SetReconciliationIdentification(value string) {
	t.ReconciliationIdentification = (*Max35Text)(&value)
}
