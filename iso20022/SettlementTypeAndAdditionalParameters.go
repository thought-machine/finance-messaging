package iso20022

// Provides transaction type and identification information.
type SettlementTypeAndAdditionalParameters1 struct {

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Indicates whether the settlement transaction was already sent on the market and that it is only sent by an account owner to an account servicer for reconciliation purposes.
	ReconciliationIndicator *YesNoIndicator `xml:"RcncltnInd,omitempty"`
}

func (s *SettlementTypeAndAdditionalParameters1) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters1) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters1) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters1) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters1) SetReconciliationIndicator(value string) {
	s.ReconciliationIndicator = (*YesNoIndicator)(&value)
}

// Provides transaction type and identification information.
type SettlementTypeAndAdditionalParameters11 struct {

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Indicates whether the settlement transaction was already sent on the market and that it is only sent by an account owner to an account servicer for reconciliation purposes.
	ReconciliationIndicator *YesNoIndicator `xml:"RcncltnInd,omitempty"`

	// Unique identification assigned to the instruction by the client.
	ClientCollateralInstructionIdentification *Max35Text `xml:"ClntCollInstrId,omitempty"`

	// Unique identification identifying the triparty collateral management transaction from the client's point of view.
	ClientCollateralTransactionIdentification *Max35Text `xml:"ClntCollTxId,omitempty"`

	// Unique identification identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique identification assigned to the instruction by the triparty agent.
	TripartyCollateralInstructionIdentification *Max35Text `xml:"TrptyCollInstrId,omitempty"`
}

func (s *SettlementTypeAndAdditionalParameters11) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters11) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters11) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters11) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters11) SetReconciliationIndicator(value string) {
	s.ReconciliationIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementTypeAndAdditionalParameters11) SetClientCollateralInstructionIdentification(value string) {
	s.ClientCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters11) SetClientCollateralTransactionIdentification(value string) {
	s.ClientCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters11) SetTripartyAgentCollateralTransactionIdentification(value string) {
	s.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters11) SetTripartyCollateralInstructionIdentification(value string) {
	s.TripartyCollateralInstructionIdentification = (*Max35Text)(&value)
}

// Provides transaction type and identification information.
type SettlementTypeAndAdditionalParameters12 struct {

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`
}

func (s *SettlementTypeAndAdditionalParameters12) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters12) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters12) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

// Provides transaction type and identification information.
type SettlementTypeAndAdditionalParameters13 struct {

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Indicates whether the settlement transaction was already sent on the market and that it is only sent by an account owner to an account servicer for reconciliation purposes.
	ReconciliationIndicator *YesNoIndicator `xml:"RcncltnInd,omitempty"`
}

func (s *SettlementTypeAndAdditionalParameters13) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters13) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters13) SetReconciliationIndicator(value string) {
	s.ReconciliationIndicator = (*YesNoIndicator)(&value)
}

// Provides transaction type and identification information.
type SettlementTypeAndAdditionalParameters14 struct {

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp,omitempty"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Indicates whether the settlement transaction was already sent on the market and that it is only sent by an account owner to an account servicer for reconciliation purposes.
	ReconciliationIndicator *YesNoIndicator `xml:"RcncltnInd,omitempty"`
}

func (s *SettlementTypeAndAdditionalParameters14) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters14) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters14) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters14) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters14) SetReconciliationIndicator(value string) {
	s.ReconciliationIndicator = (*YesNoIndicator)(&value)
}

// Provides transaction type and identification information.
type SettlementTypeAndAdditionalParameters15 struct {

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`
}

func (s *SettlementTypeAndAdditionalParameters15) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters15) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters15) SetCommonIdentification(value string) {
	s.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

// Provides transaction type and identification information.
type SettlementTypeAndAdditionalParameters16 struct {

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId,omitempty"`

	// Indicates whether the settlement transaction was already sent on the market and that it is only sent by an account owner to an account servicer for reconciliation purposes.
	ReconciliationIndicator *YesNoIndicator `xml:"RcncltnInd,omitempty"`

	// Unique identification assigned to the instruction by the client.
	ClientCollateralInstructionIdentification *RestrictedFINXMax16Text `xml:"ClntCollInstrId,omitempty"`

	// Unique identification identifying the triparty collateral management transaction from the client's point of view.
	ClientCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"ClntCollTxId,omitempty"`

	// Unique identification identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique identification assigned to the instruction by the triparty agent.
	TripartyCollateralInstructionIdentification *RestrictedFINXMax16Text `xml:"TrptyCollInstrId,omitempty"`
}

func (s *SettlementTypeAndAdditionalParameters16) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters16) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters16) SetCommonIdentification(value string) {
	s.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters16) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters16) SetReconciliationIndicator(value string) {
	s.ReconciliationIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementTypeAndAdditionalParameters16) SetClientCollateralInstructionIdentification(value string) {
	s.ClientCollateralInstructionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters16) SetClientCollateralTransactionIdentification(value string) {
	s.ClientCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters16) SetTripartyAgentCollateralTransactionIdentification(value string) {
	s.TripartyAgentCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters16) SetTripartyCollateralInstructionIdentification(value string) {
	s.TripartyCollateralInstructionIdentification = (*RestrictedFINXMax16Text)(&value)
}

// Provides transaction type and identification information.
type SettlementTypeAndAdditionalParameters17 struct {

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId,omitempty"`

	// Indicates whether the settlement transaction was already sent on the market and that it is only sent by an account owner to an account servicer for reconciliation purposes.
	ReconciliationIndicator *YesNoIndicator `xml:"RcncltnInd,omitempty"`
}

func (s *SettlementTypeAndAdditionalParameters17) SetCommonIdentification(value string) {
	s.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters17) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters17) SetReconciliationIndicator(value string) {
	s.ReconciliationIndicator = (*YesNoIndicator)(&value)
}

// Provides transaction type and identification information.
type SettlementTypeAndAdditionalParameters18 struct {

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp,omitempty"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId,omitempty"`

	// Indicates whether the settlement transaction was already sent on the market and that it is only sent by an account owner to an account servicer for reconciliation purposes.
	ReconciliationIndicator *YesNoIndicator `xml:"RcncltnInd,omitempty"`
}

func (s *SettlementTypeAndAdditionalParameters18) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters18) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters18) SetCommonIdentification(value string) {
	s.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters18) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters18) SetReconciliationIndicator(value string) {
	s.ReconciliationIndicator = (*YesNoIndicator)(&value)
}

// Provides transaction type and identification information.
type SettlementTypeAndAdditionalParameters19 struct {

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Indicates whether the settlement transaction was already sent on the market and that it is only sent by an account owner to an account servicer for reconciliation purposes.
	ReconciliationIndicator *YesNoIndicator `xml:"RcncltnInd,omitempty"`

	// Unique identification assigned to the instruction by the client.
	ClientCollateralInstructionIdentification *Max35Text `xml:"ClntCollInstrId,omitempty"`

	// Unique identification identifying the triparty collateral management transaction from the client's point of view.
	ClientCollateralTransactionIdentification *Max35Text `xml:"ClntCollTxId,omitempty"`

	// Unique identification identifying the triparty collateral management transaction from the triparty-agent's/service-provider's point of view.
	TripartyAgentServiceProviderCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtSvcPrvdrCollTxId,omitempty"`

	// Unique identification assigned to the instruction by the triparty-agent/service-provider.
	TripartyAgentServiceProviderCollateralInstructionIdentification *Max35Text `xml:"TrptyAgtSvcPrvdrCollInstrId,omitempty"`
}

func (s *SettlementTypeAndAdditionalParameters19) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters19) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters19) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters19) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters19) SetReconciliationIndicator(value string) {
	s.ReconciliationIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementTypeAndAdditionalParameters19) SetClientCollateralInstructionIdentification(value string) {
	s.ClientCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters19) SetClientCollateralTransactionIdentification(value string) {
	s.ClientCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters19) SetTripartyAgentServiceProviderCollateralTransactionIdentification(value string) {
	s.TripartyAgentServiceProviderCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters19) SetTripartyAgentServiceProviderCollateralInstructionIdentification(value string) {
	s.TripartyAgentServiceProviderCollateralInstructionIdentification = (*Max35Text)(&value)
}

// Provides transaction type and identification information.
type SettlementTypeAndAdditionalParameters2 struct {

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`
}

func (s *SettlementTypeAndAdditionalParameters2) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters2) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters2) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

// Provides transaction type and identification information.
type SettlementTypeAndAdditionalParameters20 struct {

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId,omitempty"`

	// Indicates whether the settlement transaction was already sent on the market and that it is only sent by an account owner to an account servicer for reconciliation purposes.
	ReconciliationIndicator *YesNoIndicator `xml:"RcncltnInd,omitempty"`

	// Unique identification assigned to the instruction by the client.
	ClientCollateralInstructionIdentification *RestrictedFINXMax16Text `xml:"ClntCollInstrId,omitempty"`

	// Unique identification identifying the triparty collateral management transaction from the client's point of view.
	ClientCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"ClntCollTxId,omitempty"`

	// Unique identification identifying the triparty collateral management transaction from the triparty-agent's/service-provider's point of view.
	TripartyAgentServiceProviderCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"TrptyAgtSvcPrvdrCollTxId,omitempty"`

	// Unique identification assigned to the instruction by the triparty-agent/service-provider.
	TripartyAgentServiceProviderCollateralInstructionIdentification *RestrictedFINXMax16Text `xml:"TrptyAgtSvcPrvdrCollInstrId,omitempty"`
}

func (s *SettlementTypeAndAdditionalParameters20) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters20) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters20) SetCommonIdentification(value string) {
	s.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters20) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters20) SetReconciliationIndicator(value string) {
	s.ReconciliationIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementTypeAndAdditionalParameters20) SetClientCollateralInstructionIdentification(value string) {
	s.ClientCollateralInstructionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters20) SetClientCollateralTransactionIdentification(value string) {
	s.ClientCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters20) SetTripartyAgentServiceProviderCollateralTransactionIdentification(value string) {
	s.TripartyAgentServiceProviderCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters20) SetTripartyAgentServiceProviderCollateralInstructionIdentification(value string) {
	s.TripartyAgentServiceProviderCollateralInstructionIdentification = (*RestrictedFINXMax16Text)(&value)
}

// Provides transaction type and identification information.
type SettlementTypeAndAdditionalParameters5 struct {

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp,omitempty"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Indicates whether the settlement transaction was already sent on the market and that it is only sent by an account owner to an account servicer for reconciliation purposes.
	ReconciliationIndicator *YesNoIndicator `xml:"RcncltnInd,omitempty"`
}

func (s *SettlementTypeAndAdditionalParameters5) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters5) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters5) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters5) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters5) SetReconciliationIndicator(value string) {
	s.ReconciliationIndicator = (*YesNoIndicator)(&value)
}

// Provides transaction type and identification information.
type SettlementTypeAndAdditionalParameters6 struct {

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Indicates whether the settlement transaction was already sent on the market and that it is only sent by an account owner to an account servicer for reconciliation purposes.
	ReconciliationIndicator *YesNoIndicator `xml:"RcncltnInd,omitempty"`
}

func (s *SettlementTypeAndAdditionalParameters6) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters6) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters6) SetReconciliationIndicator(value string) {
	s.ReconciliationIndicator = (*YesNoIndicator)(&value)
}

// Provides transaction type and identification information.
type SettlementTypeAndAdditionalParameters9 struct {

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Indicates whether the settlement transaction was already sent on the market and that it is only sent by an account owner to an account servicer for reconciliation purposes.
	ReconciliationIndicator *YesNoIndicator `xml:"RcncltnInd,omitempty"`

	// Unique identification assigned to the instruction by the client.
	ClientCollateralInstructionIdentification *Max35Text `xml:"ClntCollInstrId,omitempty"`

	// Unique identification identifying the triparty collateral management transaction from the client's point of view.
	ClientCollateralTransactionIdentification *Max35Text `xml:"ClntCollTxId,omitempty"`

	// Unique identification identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique identification assigned to the instruction by the triparty agent.
	TripartyCollateralInstructionIdentification *Max35Text `xml:"TrptyCollInstrId,omitempty"`
}

func (s *SettlementTypeAndAdditionalParameters9) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters9) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters9) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters9) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters9) SetReconciliationIndicator(value string) {
	s.ReconciliationIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementTypeAndAdditionalParameters9) SetClientCollateralInstructionIdentification(value string) {
	s.ClientCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters9) SetClientCollateralTransactionIdentification(value string) {
	s.ClientCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters9) SetTripartyAgentCollateralTransactionIdentification(value string) {
	s.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters9) SetTripartyCollateralInstructionIdentification(value string) {
	s.TripartyCollateralInstructionIdentification = (*Max35Text)(&value)
}
