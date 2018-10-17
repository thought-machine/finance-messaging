package iso20022

// Reference information concerning the original payment initiation, to which the investigation message refers.
type UnderlyingPaymentInstruction1 struct {

	// Set of elements used to provide information on the original messsage.
	OriginalGroupInformation *UnderlyingGroupInformation1 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency, as ordered by the original initiating party.
	OriginalInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt"`

	// Date at which the initiating party requests the clearing agent to process the payment.
	// Usage: This is the date on which the debtor's account is to be debited. If payment by cheque, the date when the cheque must be generated by the bank.
	RequestedExecutionDate *ISODate `xml:"ReqdExctnDt,omitempty"`

	// Date at which the creditor requests the amount of money to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt,omitempty"`
}

func (u *UnderlyingPaymentInstruction1) AddOriginalGroupInformation() *UnderlyingGroupInformation1 {
	u.OriginalGroupInformation = new(UnderlyingGroupInformation1)
	return u.OriginalGroupInformation
}

func (u *UnderlyingPaymentInstruction1) SetOriginalPaymentInformationIdentification(value string) {
	u.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (u *UnderlyingPaymentInstruction1) SetOriginalInstructionIdentification(value string) {
	u.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (u *UnderlyingPaymentInstruction1) SetOriginalEndToEndIdentification(value string) {
	u.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (u *UnderlyingPaymentInstruction1) SetOriginalInstructedAmount(value, currency string) {
	u.OriginalInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (u *UnderlyingPaymentInstruction1) SetRequestedExecutionDate(value string) {
	u.RequestedExecutionDate = (*ISODate)(&value)
}

func (u *UnderlyingPaymentInstruction1) SetRequestedCollectionDate(value string) {
	u.RequestedCollectionDate = (*ISODate)(&value)
}

// Reference information concerning the original payment initiation, to which the investigation message refers.
type UnderlyingPaymentInstruction2 struct {

	// Set of elements used to provide information on the original message.
	OriginalGroupInformation *UnderlyingGroupInformation1 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency, as ordered by the original initiating party.
	OriginalInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt"`

	// Date at which the initiating party requests the clearing agent to process the payment.
	// Usage: This is the date on which the debtor's account is to be debited. If payment by cheque, the date when the cheque must be generated by the bank.
	RequestedExecutionDate *ISODate `xml:"ReqdExctnDt,omitempty"`

	// Date at which the creditor requests the amount of money to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt,omitempty"`
}

func (u *UnderlyingPaymentInstruction2) AddOriginalGroupInformation() *UnderlyingGroupInformation1 {
	u.OriginalGroupInformation = new(UnderlyingGroupInformation1)
	return u.OriginalGroupInformation
}

func (u *UnderlyingPaymentInstruction2) SetOriginalPaymentInformationIdentification(value string) {
	u.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (u *UnderlyingPaymentInstruction2) SetOriginalInstructionIdentification(value string) {
	u.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (u *UnderlyingPaymentInstruction2) SetOriginalEndToEndIdentification(value string) {
	u.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (u *UnderlyingPaymentInstruction2) SetOriginalInstructedAmount(value, currency string) {
	u.OriginalInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (u *UnderlyingPaymentInstruction2) SetRequestedExecutionDate(value string) {
	u.RequestedExecutionDate = (*ISODate)(&value)
}

func (u *UnderlyingPaymentInstruction2) SetRequestedCollectionDate(value string) {
	u.RequestedCollectionDate = (*ISODate)(&value)
}

// Reference information concerning the original payment initiation, to which the investigation message refers.
type UnderlyingPaymentInstruction3 struct {

	// Set of elements used to provide information on the original message.
	OriginalGroupInformation *UnderlyingGroupInformation1 `xml:"OrgnlGrpInf,omitempty"`

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId,omitempty"`

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionIdentification *Max35Text `xml:"OrgnlInstrId,omitempty"`

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndIdentification *Max35Text `xml:"OrgnlEndToEndId,omitempty"`

	// Amount of money, as provided in the original transaction, to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency, as ordered by the original initiating party.
	OriginalInstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt"`

	// Date at which the initiating party requests the clearing agent to process the payment.
	// Usage: This is the date on which the debtor's account is to be debited. If payment by cheque, the date when the cheque must be generated by the bank.
	RequestedExecutionDate *DateAndDateTimeChoice `xml:"ReqdExctnDt,omitempty"`

	// Date at which the creditor requests the amount of money to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt,omitempty"`
}

func (u *UnderlyingPaymentInstruction3) AddOriginalGroupInformation() *UnderlyingGroupInformation1 {
	u.OriginalGroupInformation = new(UnderlyingGroupInformation1)
	return u.OriginalGroupInformation
}

func (u *UnderlyingPaymentInstruction3) SetOriginalPaymentInformationIdentification(value string) {
	u.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (u *UnderlyingPaymentInstruction3) SetOriginalInstructionIdentification(value string) {
	u.OriginalInstructionIdentification = (*Max35Text)(&value)
}

func (u *UnderlyingPaymentInstruction3) SetOriginalEndToEndIdentification(value string) {
	u.OriginalEndToEndIdentification = (*Max35Text)(&value)
}

func (u *UnderlyingPaymentInstruction3) SetOriginalInstructedAmount(value, currency string) {
	u.OriginalInstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (u *UnderlyingPaymentInstruction3) AddRequestedExecutionDate() *DateAndDateTimeChoice {
	u.RequestedExecutionDate = new(DateAndDateTimeChoice)
	return u.RequestedExecutionDate
}

func (u *UnderlyingPaymentInstruction3) SetRequestedCollectionDate(value string) {
	u.RequestedCollectionDate = (*ISODate)(&value)
}
