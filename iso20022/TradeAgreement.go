package iso20022

// Date and identification of a trade.
type TradeAgreement1 struct {

	// Date at which the trading parties agree on a treasury trade.
	TradeDate *ISODate `xml:"TradDt"`

	// Identification of a notification.This identification must be unique amongst all notifications of same type confirmed by the same party.
	NotificationIdentification *Max35Text `xml:"NtfctnId"`

	// Reference common to the parties of a trade.
	CommonReference *Max35Text `xml:"CmonRef,omitempty"`
}

func (t *TradeAgreement1) SetTradeDate(value string) {
	t.TradeDate = (*ISODate)(&value)
}

func (t *TradeAgreement1) SetNotificationIdentification(value string) {
	t.NotificationIdentification = (*Max35Text)(&value)
}

func (t *TradeAgreement1) SetCommonReference(value string) {
	t.CommonReference = (*Max35Text)(&value)
}

// Date and identification of a trade.
type TradeAgreement10 struct {

	// Date on which the trading parties agreed on the trade.
	TradeDate *ISODate `xml:"TradDt"`

	// Reference of the present instruction assigned by the party issuing the message. This reference must be unique amongst all messages of same type sent by the same party.
	OriginatorReference *Max35Text `xml:"OrgtrRef"`

	// Reference common to both parties of the trade.
	CommonReference *Max35Text `xml:"CmonRef,omitempty"`

	// Specifies the type of underlying transaction, for example cancellation (CANC).
	OperationType *Max4Text `xml:"OprTp,omitempty"`

	// Specifies the business role between the submitter and the trade party, for example Agent (AGNT).
	OperationScope *Max4Text `xml:"OprScp,omitempty"`

	// To indicate the requested CLS settlement session that the related trade is part of.
	SettlementSessionIdentifier *Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty"`

	// Specifies if the FX transaction is PVP settlement. Payment versus payment (PvP) settlement arrangement allows for two currencies in a foreign exchange (FX) contract to exchange simultaneously on a central settlement platform to eliminate the settlement risk. To apply PvP, the two parties in the FX contract need to have a pre-agreement with the central settlement platform, for example, USD/MYR FX deals require both parties to have an agreement to settle via HK Interbank Clearing Ltd settlement platform.
	PaymentVersusPaymentIndicator *YesNoIndicator `xml:"PmtVrssPmtInd,omitempty"`
}

func (t *TradeAgreement10) SetTradeDate(value string) {
	t.TradeDate = (*ISODate)(&value)
}

func (t *TradeAgreement10) SetOriginatorReference(value string) {
	t.OriginatorReference = (*Max35Text)(&value)
}

func (t *TradeAgreement10) SetCommonReference(value string) {
	t.CommonReference = (*Max35Text)(&value)
}

func (t *TradeAgreement10) SetOperationType(value string) {
	t.OperationType = (*Max4Text)(&value)
}

func (t *TradeAgreement10) SetOperationScope(value string) {
	t.OperationScope = (*Max4Text)(&value)
}

func (t *TradeAgreement10) SetSettlementSessionIdentifier(value string) {
	t.SettlementSessionIdentifier = (*Exact4AlphaNumericText)(&value)
}

func (t *TradeAgreement10) SetPaymentVersusPaymentIndicator(value string) {
	t.PaymentVersusPaymentIndicator = (*YesNoIndicator)(&value)
}

// Date and identification of a trade together with references to previous events in its life.
type TradeAgreement11 struct {

	// Date on which the trading parties agreed to amend or cancel the trade.
	TradeDate *ISODate `xml:"TradDt"`

	// Reference of the present instruction assigned by the party issuing the message. This reference must be unique amongst all messages of same type sent by the same party.
	OriginatorReference *Max35Text `xml:"OrgtrRef"`

	// Identification of a matching system reference by a choice between a matching system unique identification or the related reference.
	MatchingSystemReference *MatchingSystemReference1Choice `xml:"MtchgSysRef"`

	// Reference common to both parties of the trade.
	CommonReference *Max35Text `xml:"CmonRef,omitempty"`

	// Describes the reason for the cancellation or the amendment.
	AmendOrCancelReason *Max35Text `xml:"AmdOrCclRsn,omitempty"`

	// Specifies the type of  underlying transaction, for example cancellation (CANC).
	OperationType *Max4Text `xml:"OprTp,omitempty"`

	// Specifies the business role between the submitter and the trade party, for example Agent (AGNT).
	OperationScope *Max4Text `xml:"OprScp,omitempty"`

	// To indicate the requested CLS settlement session that the related trade is part of.
	SettlementSessionIdentifier *Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty"`

	// Specifies if the FX transaction is PVP settlement. Payment versus payment (PvP) settlement arrangement allows for two currencies in a foreign exchange (FX) contract to exchange simultaneously on a central settlement platform to eliminate the settlement risk. To apply PvP, the two parties in the FX contract need to have a pre-agreement with the central settlement platform, for example, USD/MYR FX deals require both parties to have an agreement to settle via HK Interbank Clearing Ltd settlement platform.
	PaymentVersusPaymentIndicator *YesNoIndicator `xml:"PmtVrssPmtInd,omitempty"`
}

func (t *TradeAgreement11) SetTradeDate(value string) {
	t.TradeDate = (*ISODate)(&value)
}

func (t *TradeAgreement11) SetOriginatorReference(value string) {
	t.OriginatorReference = (*Max35Text)(&value)
}

func (t *TradeAgreement11) AddMatchingSystemReference() *MatchingSystemReference1Choice {
	t.MatchingSystemReference = new(MatchingSystemReference1Choice)
	return t.MatchingSystemReference
}

func (t *TradeAgreement11) SetCommonReference(value string) {
	t.CommonReference = (*Max35Text)(&value)
}

func (t *TradeAgreement11) SetAmendOrCancelReason(value string) {
	t.AmendOrCancelReason = (*Max35Text)(&value)
}

func (t *TradeAgreement11) SetOperationType(value string) {
	t.OperationType = (*Max4Text)(&value)
}

func (t *TradeAgreement11) SetOperationScope(value string) {
	t.OperationScope = (*Max4Text)(&value)
}

func (t *TradeAgreement11) SetSettlementSessionIdentifier(value string) {
	t.SettlementSessionIdentifier = (*Exact4AlphaNumericText)(&value)
}

func (t *TradeAgreement11) SetPaymentVersusPaymentIndicator(value string) {
	t.PaymentVersusPaymentIndicator = (*YesNoIndicator)(&value)
}

// Date and identification of a trade together with references to previous events in its life.
type TradeAgreement12 struct {

	// Date on which the trading parties agreed on the trade.
	TradeDate *ISODate `xml:"TradDt"`

	// Identification of the present message assigned by the party issuing the message. This identification must be unique amongst all messages of same type sent by the same party.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Represents the original reference of the instruction for which the status is given, as assigned by the participant that submitted the foreign exchange trade.
	OriginatorReference *Max35Text `xml:"OrgtrRef"`

	// Reference common to both parties of the trade.
	CommonReference *Max35Text `xml:"CmonRef,omitempty"`

	// Specifies the reason for the cancellation or the amendment.
	AmendOrCancelReason *Max35Text `xml:"AmdOrCclRsn,omitempty"`

	// Reference to the identification of a previous event in the life of a trade which is amended or cancelled.
	RelatedReference *Max35Text `xml:"RltdRef,omitempty"`

	// Specifies the product for which the status of the confirmation is reported.
	ProductType *Max35Text `xml:"PdctTp,omitempty"`

	// Specifies the type of underlying transaction, for example cancellation (CANC).
	OperationType *Max4Text `xml:"OprTp,omitempty"`

	// Specifies the business role between the submitter and the trade party, for example, agent (AGNT).
	OperationScope *Max4Text `xml:"OprScp,omitempty"`

	// To indicate the requested CLS settlement session that the related trade is part of.
	SettlementSessionIdentifier *Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty"`

	// To indicate if the trade is split.
	SplitTradeIndicator *YesNoIndicator `xml:"SpltTradInd"`

	// Specifies if the FX transaction is PVP settlement. Payment versus payment (PvP) settlement arrangement allows for two currencies in a foreign exchange (FX) contract to exchange simultaneously on a central settlement platform to eliminate the settlement risk. To apply PvP, the two parties in the FX contract need to have a pre-agreement with the central settlement platform, for example, USD/MYR FX deals require both parties to have an agreement to settle via HK Interbank Clearing Ltd settlement platform.
	PaymentVersusPaymentIndicator *YesNoIndicator `xml:"PmtVrssPmtInd,omitempty"`
}

func (t *TradeAgreement12) SetTradeDate(value string) {
	t.TradeDate = (*ISODate)(&value)
}

func (t *TradeAgreement12) SetMessageIdentification(value string) {
	t.MessageIdentification = (*Max35Text)(&value)
}

func (t *TradeAgreement12) SetOriginatorReference(value string) {
	t.OriginatorReference = (*Max35Text)(&value)
}

func (t *TradeAgreement12) SetCommonReference(value string) {
	t.CommonReference = (*Max35Text)(&value)
}

func (t *TradeAgreement12) SetAmendOrCancelReason(value string) {
	t.AmendOrCancelReason = (*Max35Text)(&value)
}

func (t *TradeAgreement12) SetRelatedReference(value string) {
	t.RelatedReference = (*Max35Text)(&value)
}

func (t *TradeAgreement12) SetProductType(value string) {
	t.ProductType = (*Max35Text)(&value)
}

func (t *TradeAgreement12) SetOperationType(value string) {
	t.OperationType = (*Max4Text)(&value)
}

func (t *TradeAgreement12) SetOperationScope(value string) {
	t.OperationScope = (*Max4Text)(&value)
}

func (t *TradeAgreement12) SetSettlementSessionIdentifier(value string) {
	t.SettlementSessionIdentifier = (*Exact4AlphaNumericText)(&value)
}

func (t *TradeAgreement12) SetSplitTradeIndicator(value string) {
	t.SplitTradeIndicator = (*YesNoIndicator)(&value)
}

func (t *TradeAgreement12) SetPaymentVersusPaymentIndicator(value string) {
	t.PaymentVersusPaymentIndicator = (*YesNoIndicator)(&value)
}

// Contractual details related to the agreement between parties.
type TradeAgreement13 struct {

	// Party that is specified as the buyer for this trade agreement.
	Buyer *TradeParty3 `xml:"Buyr"`

	// Party that is specified as the seller for this trade agreement.
	Seller *TradeParty3 `xml:"Sellr"`

	// Quotation document referenced from this trade agreement.
	QuotationDocumentIdentification *DocumentIdentification22 `xml:"QtnDocId,omitempty"`

	// Contract document referenced from this trade agreement.
	ContractDocumentIdentification *DocumentIdentification22 `xml:"CtrctDocId,omitempty"`

	// Buyer order document referenced from this trade agreement.
	BuyerOrderIdentificationDocument *DocumentIdentification22 `xml:"BuyrOrdrIdDoc,omitempty"`

	// Additional document referenced from this trade agreement.
	AdditionalReferenceDocument []*DocumentGeneralInformation2 `xml:"AddtlRefDoc,omitempty"`

	// Specifies the applicable Incoterm and associated location.
	Incoterms *Incoterms3 `xml:"Incotrms,omitempty"`
}

func (t *TradeAgreement13) AddBuyer() *TradeParty3 {
	t.Buyer = new(TradeParty3)
	return t.Buyer
}

func (t *TradeAgreement13) AddSeller() *TradeParty3 {
	t.Seller = new(TradeParty3)
	return t.Seller
}

func (t *TradeAgreement13) AddQuotationDocumentIdentification() *DocumentIdentification22 {
	t.QuotationDocumentIdentification = new(DocumentIdentification22)
	return t.QuotationDocumentIdentification
}

func (t *TradeAgreement13) AddContractDocumentIdentification() *DocumentIdentification22 {
	t.ContractDocumentIdentification = new(DocumentIdentification22)
	return t.ContractDocumentIdentification
}

func (t *TradeAgreement13) AddBuyerOrderIdentificationDocument() *DocumentIdentification22 {
	t.BuyerOrderIdentificationDocument = new(DocumentIdentification22)
	return t.BuyerOrderIdentificationDocument
}

func (t *TradeAgreement13) AddAdditionalReferenceDocument() *DocumentGeneralInformation2 {
	newValue := new(DocumentGeneralInformation2)
	t.AdditionalReferenceDocument = append(t.AdditionalReferenceDocument, newValue)
	return newValue
}

func (t *TradeAgreement13) AddIncoterms() *Incoterms3 {
	t.Incoterms = new(Incoterms3)
	return t.Incoterms
}

// Date and identification of a trade.
type TradeAgreement14 struct {

	// Date on which the trading parties agreed on the trade.
	TradeDate *ISODate `xml:"TradDt"`

	// Reference of the present instruction assigned by the party issuing the message. This reference must be unique amongst all messages of same type sent by the same party.
	OriginatorReference *Max35Text `xml:"OrgtrRef"`

	// Reference common to both parties of the trade.
	CommonReference *Max35Text `xml:"CmonRef,omitempty"`

	// Specifies the type of underlying transaction, for example cancellation (CANC).
	OperationType *Max4Text `xml:"OprTp,omitempty"`

	// Specifies the business role between the submitter and the trade party, for example Agent (AGNT).
	OperationScope *Max4Text `xml:"OprScp,omitempty"`

	// Specifies the product for which the status of the confirmation is reported.
	ProductType *Max35Text `xml:"PdctTp,omitempty"`

	// To indicate the requested CLS settlement session that the related trade is part of.
	SettlementSessionIdentifier *Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty"`

	// Specifies if the FX transaction is PVP settlement. Payment versus payment (PvP) settlement arrangement allows for two currencies in a foreign exchange (FX) contract to exchange simultaneously on a central settlement platform to eliminate the settlement risk. To apply PvP, the two parties in the FX contract need to have a pre-agreement with the central settlement platform, for example, USD/MYR FX deals require both parties to have an agreement to settle via HK Interbank Clearing Ltd settlement platform.
	PaymentVersusPaymentIndicator *YesNoIndicator `xml:"PmtVrssPmtInd,omitempty"`
}

func (t *TradeAgreement14) SetTradeDate(value string) {
	t.TradeDate = (*ISODate)(&value)
}

func (t *TradeAgreement14) SetOriginatorReference(value string) {
	t.OriginatorReference = (*Max35Text)(&value)
}

func (t *TradeAgreement14) SetCommonReference(value string) {
	t.CommonReference = (*Max35Text)(&value)
}

func (t *TradeAgreement14) SetOperationType(value string) {
	t.OperationType = (*Max4Text)(&value)
}

func (t *TradeAgreement14) SetOperationScope(value string) {
	t.OperationScope = (*Max4Text)(&value)
}

func (t *TradeAgreement14) SetProductType(value string) {
	t.ProductType = (*Max35Text)(&value)
}

func (t *TradeAgreement14) SetSettlementSessionIdentifier(value string) {
	t.SettlementSessionIdentifier = (*Exact4AlphaNumericText)(&value)
}

func (t *TradeAgreement14) SetPaymentVersusPaymentIndicator(value string) {
	t.PaymentVersusPaymentIndicator = (*YesNoIndicator)(&value)
}

// Date and identification of a trade together with references to previous events in its life.
type TradeAgreement15 struct {

	// Date on which the trading parties agreed to amend or cancel the trade.
	TradeDate *ISODate `xml:"TradDt"`

	// Reference of the present instruction assigned by the party issuing the message. This reference must be unique amongst all messages of same type sent by the same party.
	OriginatorReference *Max35Text `xml:"OrgtrRef"`

	// Identification of a matching system reference by a choice between a matching system unique identification or the related reference.
	MatchingSystemReference *MatchingSystemReference1Choice `xml:"MtchgSysRef"`

	// Reference common to both parties of the trade.
	CommonReference *Max35Text `xml:"CmonRef,omitempty"`

	// Describes the reason for the cancellation or the amendment.
	AmendOrCancelReason *Max35Text `xml:"AmdOrCclRsn,omitempty"`

	// Specifies the type of  underlying transaction, for example cancellation (CANC).
	OperationType *Max4Text `xml:"OprTp,omitempty"`

	// Specifies the business role between the submitter and the trade party, for example Agent (AGNT).
	OperationScope *Max4Text `xml:"OprScp,omitempty"`

	// Specifies the product for which the status of the confirmation is reported.
	ProductType *Max35Text `xml:"PdctTp,omitempty"`

	// To indicate the requested CLS settlement session that the related trade is part of.
	SettlementSessionIdentifier *Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty"`

	// Specifies if the FX transaction is PVP settlement. Payment versus payment (PvP) settlement arrangement allows for two currencies in a foreign exchange (FX) contract to exchange simultaneously on a central settlement platform to eliminate the settlement risk. To apply PvP, the two parties in the FX contract need to have a pre-agreement with the central settlement platform, for example, USD/MYR FX deals require both parties to have an agreement to settle via HK Interbank Clearing Ltd settlement platform.
	PaymentVersusPaymentIndicator *YesNoIndicator `xml:"PmtVrssPmtInd,omitempty"`
}

func (t *TradeAgreement15) SetTradeDate(value string) {
	t.TradeDate = (*ISODate)(&value)
}

func (t *TradeAgreement15) SetOriginatorReference(value string) {
	t.OriginatorReference = (*Max35Text)(&value)
}

func (t *TradeAgreement15) AddMatchingSystemReference() *MatchingSystemReference1Choice {
	t.MatchingSystemReference = new(MatchingSystemReference1Choice)
	return t.MatchingSystemReference
}

func (t *TradeAgreement15) SetCommonReference(value string) {
	t.CommonReference = (*Max35Text)(&value)
}

func (t *TradeAgreement15) SetAmendOrCancelReason(value string) {
	t.AmendOrCancelReason = (*Max35Text)(&value)
}

func (t *TradeAgreement15) SetOperationType(value string) {
	t.OperationType = (*Max4Text)(&value)
}

func (t *TradeAgreement15) SetOperationScope(value string) {
	t.OperationScope = (*Max4Text)(&value)
}

func (t *TradeAgreement15) SetProductType(value string) {
	t.ProductType = (*Max35Text)(&value)
}

func (t *TradeAgreement15) SetSettlementSessionIdentifier(value string) {
	t.SettlementSessionIdentifier = (*Exact4AlphaNumericText)(&value)
}

func (t *TradeAgreement15) SetPaymentVersusPaymentIndicator(value string) {
	t.PaymentVersusPaymentIndicator = (*YesNoIndicator)(&value)
}

// Date and identification of a trade together with references to previous events in its life.
type TradeAgreement2 struct {

	// Date at which the trading parties have agreed to amend or cancel a treasury trade.
	TradeDate *ISODate `xml:"TradDt"`

	// Identification of a notification.This identification must be unique amongst all notifications of same type confirmed by the same party.
	NotificationIdentification *Max35Text `xml:"NtfctnId"`

	// Reference common to the parties of a trade.
	CommonReference *Max35Text `xml:"CmonRef,omitempty"`

	// Describes the reason for the cancellation or the amendment.
	AmendOrCancelReason *Max35Text `xml:"AmdOrCclRsn,omitempty"`

	// Refers to the identification of a previous event in the life of a trade which is amended or cancelled.
	RelatedReference *Max35Text `xml:"RltdRef"`
}

func (t *TradeAgreement2) SetTradeDate(value string) {
	t.TradeDate = (*ISODate)(&value)
}

func (t *TradeAgreement2) SetNotificationIdentification(value string) {
	t.NotificationIdentification = (*Max35Text)(&value)
}

func (t *TradeAgreement2) SetCommonReference(value string) {
	t.CommonReference = (*Max35Text)(&value)
}

func (t *TradeAgreement2) SetAmendOrCancelReason(value string) {
	t.AmendOrCancelReason = (*Max35Text)(&value)
}

func (t *TradeAgreement2) SetRelatedReference(value string) {
	t.RelatedReference = (*Max35Text)(&value)
}

// Contractual details related to the agreement between parties.
type TradeAgreement6 struct {

	// Party that is specified as the buyer for this trade agreement.
	Buyer *TradeParty1 `xml:"Buyr"`

	// Party that is specified as the seller for this trade agreement.
	Seller *TradeParty1 `xml:"Sellr"`

	// Quotation document referenced from this trade agreement.
	QuotationDocumentIdentification *DocumentIdentification22 `xml:"QtnDocId,omitempty"`

	// Contract document referenced from this trade agreement.
	ContractDocumentIdentification *DocumentIdentification22 `xml:"CtrctDocId,omitempty"`

	// Buyer order document referenced from this trade agreement.
	BuyerOrderIdentificationDocument *DocumentIdentification22 `xml:"BuyrOrdrIdDoc,omitempty"`

	// Additional document referenced from this trade agreement.
	AdditionalReferenceDocument []*DocumentGeneralInformation2 `xml:"AddtlRefDoc,omitempty"`

	// Specifies the applicable Incoterm and associated location.
	Incoterms *Incoterms3 `xml:"Incotrms,omitempty"`
}

func (t *TradeAgreement6) AddBuyer() *TradeParty1 {
	t.Buyer = new(TradeParty1)
	return t.Buyer
}

func (t *TradeAgreement6) AddSeller() *TradeParty1 {
	t.Seller = new(TradeParty1)
	return t.Seller
}

func (t *TradeAgreement6) AddQuotationDocumentIdentification() *DocumentIdentification22 {
	t.QuotationDocumentIdentification = new(DocumentIdentification22)
	return t.QuotationDocumentIdentification
}

func (t *TradeAgreement6) AddContractDocumentIdentification() *DocumentIdentification22 {
	t.ContractDocumentIdentification = new(DocumentIdentification22)
	return t.ContractDocumentIdentification
}

func (t *TradeAgreement6) AddBuyerOrderIdentificationDocument() *DocumentIdentification22 {
	t.BuyerOrderIdentificationDocument = new(DocumentIdentification22)
	return t.BuyerOrderIdentificationDocument
}

func (t *TradeAgreement6) AddAdditionalReferenceDocument() *DocumentGeneralInformation2 {
	newValue := new(DocumentGeneralInformation2)
	t.AdditionalReferenceDocument = append(t.AdditionalReferenceDocument, newValue)
	return newValue
}

func (t *TradeAgreement6) AddIncoterms() *Incoterms3 {
	t.Incoterms = new(Incoterms3)
	return t.Incoterms
}
