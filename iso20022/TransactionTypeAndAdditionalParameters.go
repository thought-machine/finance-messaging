package iso20022

// Securities financing transaction identification information, type (repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing) and other parameters.
type TransactionTypeAndAdditionalParameters1 struct {

	// Specifies the type of securities financing transaction, that is, repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing.
	SecuritiesFinancingTransactionType *SecuritiesFinancingTransactionType1Code `xml:"SctiesFincgTxTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Settlement transaction has already been sent on the market. It is sent by an account owner to an account servicer for reconciliation purposes.
	ReconciliationIndicator *YesNoIndicator `xml:"RcncltnInd,omitempty"`
}

func (t *TransactionTypeAndAdditionalParameters1) SetSecuritiesFinancingTransactionType(value string) {
	t.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType1Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters1) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters1) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters1) SetReconciliationIndicator(value string) {
	t.ReconciliationIndicator = (*YesNoIndicator)(&value)
}

// Specifies the conditions under which the order/trade is to be settled.
type TransactionTypeAndAdditionalParameters10 struct {

	// Unambiguous identification of the transaction (unique per piece of collateral) as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction (unique per piece of collateral) as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Specifies the type of securities financing transaction, that is, repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing.
	SecuritiesFinancingTransactionType *SecuritiesFinancingTransactionType1Code `xml:"SctiesFincgTxTp"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`
}

func (t *TransactionTypeAndAdditionalParameters10) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters10) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters10) SetSecuritiesFinancingTransactionType(value string) {
	t.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType1Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters10) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters10) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters10) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters10) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters10) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

// Specifies the conditions under which the order/trade is to be settled.
type TransactionTypeAndAdditionalParameters11 struct {

	// Unambiguous identification of the transaction (unique per piece of collateral) as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction (unique per piece of collateral) as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Specifies the type of securities financing transaction, that is, repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing.
	SecuritiesFinancingTransactionType *SecuritiesFinancingTransactionType1Code `xml:"SctiesFincgTxTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Identifies the type of securities financing modification requested.
	ModificationType *RepurchaseType15Choice `xml:"ModTp,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`
}

func (t *TransactionTypeAndAdditionalParameters11) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters11) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters11) SetSecuritiesFinancingTransactionType(value string) {
	t.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType1Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters11) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters11) AddModificationType() *RepurchaseType15Choice {
	t.ModificationType = new(RepurchaseType15Choice)
	return t.ModificationType
}

func (t *TransactionTypeAndAdditionalParameters11) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters11) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

// Specifies the conditions under which the order/trade is to be settled.
type TransactionTypeAndAdditionalParameters12 struct {

	// Unambiguous identification of the transaction (unique per piece of collateral) as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction (unique per piece of collateral) as known by the account servicer.
	AccountServicerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctSvcrTxId,omitempty"`

	// Specifies the type of securities financing transaction, that is, repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing.
	SecuritiesFinancingTransactionType *SecuritiesFinancingTransactionType1Code `xml:"SctiesFincgTxTp"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId,omitempty"`
}

func (t *TransactionTypeAndAdditionalParameters12) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters12) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters12) SetSecuritiesFinancingTransactionType(value string) {
	t.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType1Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters12) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters12) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters12) SetCommonIdentification(value string) {
	t.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters12) SetPoolIdentification(value string) {
	t.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters12) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

// Securities financing transaction identification information, type (repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing) and other parameters.
type TransactionTypeAndAdditionalParameters13 struct {

	// Specifies the type of securities financing transaction, that is, repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing.
	SecuritiesFinancingTransactionType *SecuritiesFinancingTransactionType1Code `xml:"SctiesFincgTxTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

	// Settlement transaction has already been sent on the market. It is sent by an account owner to an account servicer for reconciliation purposes.
	ReconciliationIndicator *YesNoIndicator `xml:"RcncltnInd,omitempty"`
}

func (t *TransactionTypeAndAdditionalParameters13) SetSecuritiesFinancingTransactionType(value string) {
	t.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType1Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters13) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters13) SetCommonIdentification(value string) {
	t.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters13) SetReconciliationIndicator(value string) {
	t.ReconciliationIndicator = (*YesNoIndicator)(&value)
}

// Specifies the conditions under which the order/trade is to be settled.
type TransactionTypeAndAdditionalParameters14 struct {

	// Unambiguous identification of the transaction (unique per piece of collateral) as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction (unique per piece of collateral) as known by the account servicer.
	AccountServicerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctSvcrTxId,omitempty"`

	// Specifies the type of securities financing transaction, that is, repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing.
	SecuritiesFinancingTransactionType *SecuritiesFinancingTransactionType1Code `xml:"SctiesFincgTxTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Identifies the type of securities financing modification requested.
	ModificationType *RepurchaseType16Choice `xml:"ModTp,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId,omitempty"`
}

func (t *TransactionTypeAndAdditionalParameters14) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters14) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters14) SetSecuritiesFinancingTransactionType(value string) {
	t.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType1Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters14) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters14) AddModificationType() *RepurchaseType16Choice {
	t.ModificationType = new(RepurchaseType16Choice)
	return t.ModificationType
}

func (t *TransactionTypeAndAdditionalParameters14) SetCommonIdentification(value string) {
	t.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters14) SetPoolIdentification(value string) {
	t.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

// Securities financing transaction identification information, type (repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing) and other parameters.
type TransactionTypeAndAdditionalParameters15 struct {

	// Specifies the type of securities financing transaction, that is, repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing.
	SecuritiesFinancingTransactionType *SecuritiesFinancingTransactionType2Code `xml:"SctiesFincgTxTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Settlement transaction has already been sent on the market. It is sent by an account owner to an account servicer for reconciliation purposes.
	ReconciliationIndicator *YesNoIndicator `xml:"RcncltnInd,omitempty"`
}

func (t *TransactionTypeAndAdditionalParameters15) SetSecuritiesFinancingTransactionType(value string) {
	t.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType2Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters15) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters15) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters15) SetReconciliationIndicator(value string) {
	t.ReconciliationIndicator = (*YesNoIndicator)(&value)
}

// Specifies the conditions under which the order/trade is to be settled.
type TransactionTypeAndAdditionalParameters16 struct {

	// Unambiguous identification of the transaction (unique per piece of collateral) as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction (unique per piece of collateral) as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Specifies the type of securities financing transaction, that is, repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing.
	SecuritiesFinancingTransactionType *SecuritiesFinancingTransactionType2Code `xml:"SctiesFincgTxTp"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`
}

func (t *TransactionTypeAndAdditionalParameters16) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters16) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters16) SetSecuritiesFinancingTransactionType(value string) {
	t.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType2Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters16) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters16) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters16) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters16) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters16) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

// Specifies the conditions under which the order/trade is to be settled.
type TransactionTypeAndAdditionalParameters17 struct {

	// Unambiguous identification of the transaction (unique per piece of collateral) as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction (unique per piece of collateral) as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Specifies the type of securities financing transaction, that is, repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing.
	SecuritiesFinancingTransactionType *SecuritiesFinancingTransactionType2Code `xml:"SctiesFincgTxTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Identifies the type of securities financing modification requested.
	ModificationType *RepurchaseType21Choice `xml:"ModTp,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`
}

func (t *TransactionTypeAndAdditionalParameters17) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters17) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters17) SetSecuritiesFinancingTransactionType(value string) {
	t.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType2Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters17) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters17) AddModificationType() *RepurchaseType21Choice {
	t.ModificationType = new(RepurchaseType21Choice)
	return t.ModificationType
}

func (t *TransactionTypeAndAdditionalParameters17) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters17) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

// Securities financing transaction identification information, type (repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing) and other parameters.
type TransactionTypeAndAdditionalParameters18 struct {

	// Specifies the type of securities financing transaction, that is, repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing.
	SecuritiesFinancingTransactionType *SecuritiesFinancingTransactionType2Code `xml:"SctiesFincgTxTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

	// Settlement transaction has already been sent on the market. It is sent by an account owner to an account servicer for reconciliation purposes.
	ReconciliationIndicator *YesNoIndicator `xml:"RcncltnInd,omitempty"`
}

func (t *TransactionTypeAndAdditionalParameters18) SetSecuritiesFinancingTransactionType(value string) {
	t.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType2Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters18) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters18) SetCommonIdentification(value string) {
	t.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters18) SetReconciliationIndicator(value string) {
	t.ReconciliationIndicator = (*YesNoIndicator)(&value)
}

// Specifies the conditions under which the order/trade is to be settled.
type TransactionTypeAndAdditionalParameters19 struct {

	// Unambiguous identification of the transaction (unique per piece of collateral) as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction (unique per piece of collateral) as known by the account servicer.
	AccountServicerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctSvcrTxId,omitempty"`

	// Specifies the type of securities financing transaction, that is, repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing.
	SecuritiesFinancingTransactionType *SecuritiesFinancingTransactionType2Code `xml:"SctiesFincgTxTp"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId,omitempty"`
}

func (t *TransactionTypeAndAdditionalParameters19) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters19) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters19) SetSecuritiesFinancingTransactionType(value string) {
	t.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType2Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters19) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters19) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters19) SetCommonIdentification(value string) {
	t.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters19) SetPoolIdentification(value string) {
	t.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters19) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

// Specifies the conditions under which the order/trade is to be settled.
type TransactionTypeAndAdditionalParameters2 struct {

	// Unambiguous identification of the transaction (unique per piece of collateral) as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction (unique per piece of collateral) as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Specifies the type of securities financing transaction, that is, repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing.
	SecuritiesFinancingTransactionType *SecuritiesFinancingTransactionType1Code `xml:"SctiesFincgTxTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Identifies the type of securities financing modification requested.
	ModificationType *RepurchaseType2Choice `xml:"ModTp,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`
}

func (t *TransactionTypeAndAdditionalParameters2) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters2) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters2) SetSecuritiesFinancingTransactionType(value string) {
	t.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType1Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters2) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters2) AddModificationType() *RepurchaseType2Choice {
	t.ModificationType = new(RepurchaseType2Choice)
	return t.ModificationType
}

func (t *TransactionTypeAndAdditionalParameters2) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters2) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

// Specifies the conditions under which the order/trade is to be settled.
type TransactionTypeAndAdditionalParameters20 struct {

	// Unambiguous identification of the transaction (unique per piece of collateral) as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction (unique per piece of collateral) as known by the account servicer.
	AccountServicerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctSvcrTxId,omitempty"`

	// Specifies the type of securities financing transaction, that is, repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing.
	SecuritiesFinancingTransactionType *SecuritiesFinancingTransactionType2Code `xml:"SctiesFincgTxTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Identifies the type of securities financing modification requested.
	ModificationType *RepurchaseType31Choice `xml:"ModTp,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId,omitempty"`
}

func (t *TransactionTypeAndAdditionalParameters20) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters20) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters20) SetSecuritiesFinancingTransactionType(value string) {
	t.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType2Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters20) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters20) AddModificationType() *RepurchaseType31Choice {
	t.ModificationType = new(RepurchaseType31Choice)
	return t.ModificationType
}

func (t *TransactionTypeAndAdditionalParameters20) SetCommonIdentification(value string) {
	t.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters20) SetPoolIdentification(value string) {
	t.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

// Specifies the conditions under which the order/trade is to be settled.
type TransactionTypeAndAdditionalParameters3 struct {

	// Unambiguous identification of the transaction (unique per piece of collateral) as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction (unique per piece of collateral) as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Specifies the type of securities financing transaction, that is, repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing.
	SecuritiesFinancingTransactionType *SecuritiesFinancingTransactionType1Code `xml:"SctiesFincgTxTp"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`
}

func (t *TransactionTypeAndAdditionalParameters3) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters3) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters3) SetSecuritiesFinancingTransactionType(value string) {
	t.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType1Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters3) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters3) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters3) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters3) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters3) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

// Specifies the conditions under which the order/trade is to be settled.
type TransactionTypeAndAdditionalParameters7 struct {

	// Unambiguous identification of the transaction (unique per piece of collateral) as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction (unique per piece of collateral) as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Specifies the type of securities financing transaction, that is, repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing.
	SecuritiesFinancingTransactionType *SecuritiesFinancingTransactionType1Code `xml:"SctiesFincgTxTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Identifies the type of securities financing modification requested.
	ModificationType *RepurchaseType7Choice `xml:"ModTp,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`
}

func (t *TransactionTypeAndAdditionalParameters7) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters7) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters7) SetSecuritiesFinancingTransactionType(value string) {
	t.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType1Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters7) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters7) AddModificationType() *RepurchaseType7Choice {
	t.ModificationType = new(RepurchaseType7Choice)
	return t.ModificationType
}

func (t *TransactionTypeAndAdditionalParameters7) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters7) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

// Securities financing transaction identification information, type (repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing) and other parameters.
type TransactionTypeAndAdditionalParameters9 struct {

	// Specifies the type of securities financing transaction, that is, repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing.
	SecuritiesFinancingTransactionType *SecuritiesFinancingTransactionType1Code `xml:"SctiesFincgTxTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Settlement transaction has already been sent on the market. It is sent by an account owner to an account servicer for reconciliation purposes.
	ReconciliationIndicator *YesNoIndicator `xml:"RcncltnInd,omitempty"`
}

func (t *TransactionTypeAndAdditionalParameters9) SetSecuritiesFinancingTransactionType(value string) {
	t.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType1Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters9) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionTypeAndAdditionalParameters9) SetCommonIdentification(value string) {
	t.CommonIdentification = (*Max35Text)(&value)
}

func (t *TransactionTypeAndAdditionalParameters9) SetReconciliationIndicator(value string) {
	t.ReconciliationIndicator = (*YesNoIndicator)(&value)
}
