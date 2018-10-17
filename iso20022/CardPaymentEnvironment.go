package iso20022

// Environment of the transaction.
type CardPaymentEnvironment1 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer1 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment transaction.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation5 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction1 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard1 `xml:"Card"`

	// Cardholder involved in the card payment.
	Cardholder *Cardholder1 `xml:"Crdhldr,omitempty"`
}

func (c *CardPaymentEnvironment1) AddAcquirer() *Acquirer1 {
	c.Acquirer = new(Acquirer1)
	return c.Acquirer
}

func (c *CardPaymentEnvironment1) AddMerchant() *Organisation5 {
	c.Merchant = new(Organisation5)
	return c.Merchant
}

func (c *CardPaymentEnvironment1) AddPOI() *PointOfInteraction1 {
	c.POI = new(PointOfInteraction1)
	return c.POI
}

func (c *CardPaymentEnvironment1) AddCard() *PaymentCard1 {
	c.Card = new(PaymentCard1)
	return c.Card
}

func (c *CardPaymentEnvironment1) AddCardholder() *Cardholder1 {
	c.Cardholder = new(Cardholder1)
	return c.Cardholder
}

// Environment of the transaction.
type CardPaymentEnvironment10 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer2 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation8 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction2 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard6 `xml:"Card"`

	// Cardholder involved in the card payment.
	Cardholder *Cardholder4 `xml:"Crdhldr,omitempty"`
}

func (c *CardPaymentEnvironment10) AddAcquirer() *Acquirer2 {
	c.Acquirer = new(Acquirer2)
	return c.Acquirer
}

func (c *CardPaymentEnvironment10) AddMerchant() *Organisation8 {
	c.Merchant = new(Organisation8)
	return c.Merchant
}

func (c *CardPaymentEnvironment10) AddPOI() *PointOfInteraction2 {
	c.POI = new(PointOfInteraction2)
	return c.POI
}

func (c *CardPaymentEnvironment10) AddCard() *PaymentCard6 {
	c.Card = new(PaymentCard6)
	return c.Card
}

func (c *CardPaymentEnvironment10) AddCardholder() *Cardholder4 {
	c.Cardholder = new(Cardholder4)
	return c.Cardholder
}

// Environment of the transaction given in a response to a request.
type CardPaymentEnvironment11 struct {

	// Acquirer involved in the card payment.
	AcquirerIdentification *GenericIdentification32 `xml:"AcqrrId,omitempty"`

	// Identification of the merchant.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI performing the transaction.
	POIIdentification *GenericIdentification32 `xml:"POIId"`

	// Sensitive data of the card (PlainCardData1 including the envelope), encrypted with a cryptographic key.
	ProtectedCardData *ContentInformationType5 `xml:"PrtctdCardData,omitempty"`

	// Payment card performing the transaction.
	PlainCardData *PlainCardData3 `xml:"PlainCardData,omitempty"`
}

func (c *CardPaymentEnvironment11) AddAcquirerIdentification() *GenericIdentification32 {
	c.AcquirerIdentification = new(GenericIdentification32)
	return c.AcquirerIdentification
}

func (c *CardPaymentEnvironment11) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment11) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentEnvironment11) AddProtectedCardData() *ContentInformationType5 {
	c.ProtectedCardData = new(ContentInformationType5)
	return c.ProtectedCardData
}

func (c *CardPaymentEnvironment11) AddPlainCardData() *PlainCardData3 {
	c.PlainCardData = new(PlainCardData3)
	return c.PlainCardData
}

// Environment of the cancellation.
type CardPaymentEnvironment12 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer2 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment cancellation.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation8 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction2 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard6 `xml:"Card"`
}

func (c *CardPaymentEnvironment12) AddAcquirer() *Acquirer2 {
	c.Acquirer = new(Acquirer2)
	return c.Acquirer
}

func (c *CardPaymentEnvironment12) AddMerchant() *Organisation8 {
	c.Merchant = new(Organisation8)
	return c.Merchant
}

func (c *CardPaymentEnvironment12) AddPOI() *PointOfInteraction2 {
	c.POI = new(PointOfInteraction2)
	return c.POI
}

func (c *CardPaymentEnvironment12) AddCard() *PaymentCard6 {
	c.Card = new(PaymentCard6)
	return c.Card
}

// Environment common to a collection of transactions.
type CardPaymentEnvironment13 struct {

	// Acquirer involved in the card payment transactions.
	Acquirer *Acquirer3 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment transactions.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation9 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction2 `xml:"POI,omitempty"`
}

func (c *CardPaymentEnvironment13) AddAcquirer() *Acquirer3 {
	c.Acquirer = new(Acquirer3)
	return c.Acquirer
}

func (c *CardPaymentEnvironment13) AddMerchant() *Organisation9 {
	c.Merchant = new(Organisation9)
	return c.Merchant
}

func (c *CardPaymentEnvironment13) AddPOI() *PointOfInteraction2 {
	c.POI = new(PointOfInteraction2)
	return c.POI
}

// Environment of the card payment transaction.
type CardPaymentEnvironment14 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer2 `xml:"Acqrr,omitempty"`

	// Merchant performing the transaction.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation8 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction2 `xml:"POI,omitempty"`

	// Payment card performing the transaction.
	Card *PaymentCard6 `xml:"Card"`

	// Cardholder involved in the card payment.
	Cardholder *Cardholder4 `xml:"Crdhldr,omitempty"`
}

func (c *CardPaymentEnvironment14) AddAcquirer() *Acquirer2 {
	c.Acquirer = new(Acquirer2)
	return c.Acquirer
}

func (c *CardPaymentEnvironment14) AddMerchant() *Organisation8 {
	c.Merchant = new(Organisation8)
	return c.Merchant
}

func (c *CardPaymentEnvironment14) AddPOI() *PointOfInteraction2 {
	c.POI = new(PointOfInteraction2)
	return c.POI
}

func (c *CardPaymentEnvironment14) AddCard() *PaymentCard6 {
	c.Card = new(PaymentCard6)
	return c.Card
}

func (c *CardPaymentEnvironment14) AddCardholder() *Cardholder4 {
	c.Cardholder = new(Cardholder4)
	return c.Cardholder
}

// Environment of the reconciliation exchange.
type CardPaymentEnvironment15 struct {

	// Acquirer involved in the card payment reconciliation.
	Acquirer *Acquirer2 `xml:"Acqrr"`

	// Identification of the merchant requesting the reconciliation.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI requesting the reconciliation.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`
}

func (c *CardPaymentEnvironment15) AddAcquirer() *Acquirer2 {
	c.Acquirer = new(Acquirer2)
	return c.Acquirer
}

func (c *CardPaymentEnvironment15) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment15) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

// Environment of the transaction given in a response to a request in a batch.
type CardPaymentEnvironment16 struct {

	// Acquirer involved in the card payment.
	AcquirerIdentification *GenericIdentification32 `xml:"Acqrr,omitempty"`

	// Identification of the merchant.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI performing the transaction.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`

	// Sensitive data of the card (PlainCardData1 including the envelope), encrypted with a cryptographic key.
	ProtectedCardData *ContentInformationType5 `xml:"PrtctdCardData,omitempty"`

	// Payment card performing the transaction.
	PlainCardData *PlainCardData3 `xml:"PlainCardData,omitempty"`
}

func (c *CardPaymentEnvironment16) AddAcquirerIdentification() *GenericIdentification32 {
	c.AcquirerIdentification = new(GenericIdentification32)
	return c.AcquirerIdentification
}

func (c *CardPaymentEnvironment16) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment16) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentEnvironment16) AddProtectedCardData() *ContentInformationType5 {
	c.ProtectedCardData = new(ContentInformationType5)
	return c.ProtectedCardData
}

func (c *CardPaymentEnvironment16) AddPlainCardData() *PlainCardData3 {
	c.PlainCardData = new(PlainCardData3)
	return c.PlainCardData
}

// Environment of the diagnostic exchange.
type CardPaymentEnvironment17 struct {

	// Version of acquirer configuration parameters.
	AcquirerParametersVersion *Max35Text `xml:"AcqrrParamsVrsn"`

	// Identification of the merchant requesting the diagnostic.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI requesting the diagnostic.
	POIIdentification *GenericIdentification32 `xml:"POIId"`
}

func (c *CardPaymentEnvironment17) SetAcquirerParametersVersion(value string) {
	c.AcquirerParametersVersion = (*Max35Text)(&value)
}

func (c *CardPaymentEnvironment17) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment17) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

// Environment of the transaction.
type CardPaymentEnvironment18 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer2 `xml:"AcqrrId,omitempty"`

	// Merchant performing the card payment.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation8 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction2 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard6 `xml:"Card"`
}

func (c *CardPaymentEnvironment18) AddAcquirer() *Acquirer2 {
	c.Acquirer = new(Acquirer2)
	return c.Acquirer
}

func (c *CardPaymentEnvironment18) AddMerchant() *Organisation8 {
	c.Merchant = new(Organisation8)
	return c.Merchant
}

func (c *CardPaymentEnvironment18) AddPOI() *PointOfInteraction2 {
	c.POI = new(PointOfInteraction2)
	return c.POI
}

func (c *CardPaymentEnvironment18) AddCard() *PaymentCard6 {
	c.Card = new(PaymentCard6)
	return c.Card
}

// Environment of the reconciliation for the acquirer.
type CardPaymentEnvironment19 struct {

	// Acquirer involved in the card payment reconciliation.
	AcquirerIdentification *GenericIdentification32 `xml:"AcqrrId"`

	// Identification of the merchant requesting the reconciliation.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI requesting the reconciliation.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`
}

func (c *CardPaymentEnvironment19) AddAcquirerIdentification() *GenericIdentification32 {
	c.AcquirerIdentification = new(GenericIdentification32)
	return c.AcquirerIdentification
}

func (c *CardPaymentEnvironment19) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment19) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

// Environment of the transaction.
type CardPaymentEnvironment2 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer1 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation5 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction1 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard3 `xml:"Card"`

	// Cardholder involved in the card payment.
	Cardholder *Cardholder2 `xml:"Crdhldr,omitempty"`
}

func (c *CardPaymentEnvironment2) AddAcquirer() *Acquirer1 {
	c.Acquirer = new(Acquirer1)
	return c.Acquirer
}

func (c *CardPaymentEnvironment2) AddMerchant() *Organisation5 {
	c.Merchant = new(Organisation5)
	return c.Merchant
}

func (c *CardPaymentEnvironment2) AddPOI() *PointOfInteraction1 {
	c.POI = new(PointOfInteraction1)
	return c.POI
}

func (c *CardPaymentEnvironment2) AddCard() *PaymentCard3 {
	c.Card = new(PaymentCard3)
	return c.Card
}

func (c *CardPaymentEnvironment2) AddCardholder() *Cardholder2 {
	c.Cardholder = new(Cardholder2)
	return c.Cardholder
}

// Environment of the transaction.
type CardPaymentEnvironment20 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer2 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment transaction.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation8 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction3 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard7 `xml:"Card"`

	// Cardholder involved in the card payment.
	Cardholder *Cardholder5 `xml:"Crdhldr,omitempty"`

	// Replacement of the message element Cardholder by a digital envelope using a cryptographic key.
	ProtectedCardholderData *ContentInformationType7 `xml:"PrtctdCrdhldrData,omitempty"`
}

func (c *CardPaymentEnvironment20) AddAcquirer() *Acquirer2 {
	c.Acquirer = new(Acquirer2)
	return c.Acquirer
}

func (c *CardPaymentEnvironment20) AddMerchant() *Organisation8 {
	c.Merchant = new(Organisation8)
	return c.Merchant
}

func (c *CardPaymentEnvironment20) AddPOI() *PointOfInteraction3 {
	c.POI = new(PointOfInteraction3)
	return c.POI
}

func (c *CardPaymentEnvironment20) AddCard() *PaymentCard7 {
	c.Card = new(PaymentCard7)
	return c.Card
}

func (c *CardPaymentEnvironment20) AddCardholder() *Cardholder5 {
	c.Cardholder = new(Cardholder5)
	return c.Cardholder
}

func (c *CardPaymentEnvironment20) AddProtectedCardholderData() *ContentInformationType7 {
	c.ProtectedCardholderData = new(ContentInformationType7)
	return c.ProtectedCardholderData
}

// Environment of the transaction given in a response to a request.
type CardPaymentEnvironment21 struct {

	// Acquirer involved in the card payment.
	AcquirerIdentification *GenericIdentification32 `xml:"AcqrrId,omitempty"`

	// Identification of the merchant.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI (Point Of Interaction) performing the transaction.
	POIIdentification *GenericIdentification32 `xml:"POIId"`

	// Replacement of the message element PlainCardData by a digital envelope using a cryptographic key.
	ProtectedCardData *ContentInformationType7 `xml:"PrtctdCardData,omitempty"`

	// Payment card performing the transaction.
	PlainCardData *PlainCardData5 `xml:"PlainCardData,omitempty"`
}

func (c *CardPaymentEnvironment21) AddAcquirerIdentification() *GenericIdentification32 {
	c.AcquirerIdentification = new(GenericIdentification32)
	return c.AcquirerIdentification
}

func (c *CardPaymentEnvironment21) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment21) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentEnvironment21) AddProtectedCardData() *ContentInformationType7 {
	c.ProtectedCardData = new(ContentInformationType7)
	return c.ProtectedCardData
}

func (c *CardPaymentEnvironment21) AddPlainCardData() *PlainCardData5 {
	c.PlainCardData = new(PlainCardData5)
	return c.PlainCardData
}

// Environment of the transaction.
type CardPaymentEnvironment22 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer2 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation8 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction3 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard8 `xml:"Card,omitempty"`

	// Cardholder involved in the card payment.
	Cardholder *Cardholder6 `xml:"Crdhldr,omitempty"`

	// Replacement of the message element Cardholder by a digital envelope using a cryptographic key.
	ProtectedCardholderData *ContentInformationType7 `xml:"PrtctdCrdhldrData,omitempty"`
}

func (c *CardPaymentEnvironment22) AddAcquirer() *Acquirer2 {
	c.Acquirer = new(Acquirer2)
	return c.Acquirer
}

func (c *CardPaymentEnvironment22) AddMerchant() *Organisation8 {
	c.Merchant = new(Organisation8)
	return c.Merchant
}

func (c *CardPaymentEnvironment22) AddPOI() *PointOfInteraction3 {
	c.POI = new(PointOfInteraction3)
	return c.POI
}

func (c *CardPaymentEnvironment22) AddCard() *PaymentCard8 {
	c.Card = new(PaymentCard8)
	return c.Card
}

func (c *CardPaymentEnvironment22) AddCardholder() *Cardholder6 {
	c.Cardholder = new(Cardholder6)
	return c.Cardholder
}

func (c *CardPaymentEnvironment22) AddProtectedCardholderData() *ContentInformationType7 {
	c.ProtectedCardholderData = new(ContentInformationType7)
	return c.ProtectedCardholderData
}

// Environment of the cancellation.
type CardPaymentEnvironment23 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer2 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment cancellation.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation8 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction3 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard8 `xml:"Card"`
}

func (c *CardPaymentEnvironment23) AddAcquirer() *Acquirer2 {
	c.Acquirer = new(Acquirer2)
	return c.Acquirer
}

func (c *CardPaymentEnvironment23) AddMerchant() *Organisation8 {
	c.Merchant = new(Organisation8)
	return c.Merchant
}

func (c *CardPaymentEnvironment23) AddPOI() *PointOfInteraction3 {
	c.POI = new(PointOfInteraction3)
	return c.POI
}

func (c *CardPaymentEnvironment23) AddCard() *PaymentCard8 {
	c.Card = new(PaymentCard8)
	return c.Card
}

// Environment of the transaction.
type CardPaymentEnvironment24 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer2 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation8 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction3 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard8 `xml:"Card"`
}

func (c *CardPaymentEnvironment24) AddAcquirer() *Acquirer2 {
	c.Acquirer = new(Acquirer2)
	return c.Acquirer
}

func (c *CardPaymentEnvironment24) AddMerchant() *Organisation8 {
	c.Merchant = new(Organisation8)
	return c.Merchant
}

func (c *CardPaymentEnvironment24) AddPOI() *PointOfInteraction3 {
	c.POI = new(PointOfInteraction3)
	return c.POI
}

func (c *CardPaymentEnvironment24) AddCard() *PaymentCard8 {
	c.Card = new(PaymentCard8)
	return c.Card
}

// Environment of the transaction.
type CardPaymentEnvironment25 struct {

	// Acquirer involved in the card payment reconciliation.
	Acquirer *Acquirer2 `xml:"Acqrr"`

	// Identification of the merchant requesting the reconciliation.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI (Point Of Interaction) requesting the reconciliation.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`

	// Data related to the components of the POI (Point Of Interaction) that have been performed the payment transactions.
	POIComponent []*PointOfInteractionComponent4 `xml:"POICmpnt,omitempty"`
}

func (c *CardPaymentEnvironment25) AddAcquirer() *Acquirer2 {
	c.Acquirer = new(Acquirer2)
	return c.Acquirer
}

func (c *CardPaymentEnvironment25) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment25) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentEnvironment25) AddPOIComponent() *PointOfInteractionComponent4 {
	newValue := new(PointOfInteractionComponent4)
	c.POIComponent = append(c.POIComponent, newValue)
	return newValue
}

// Environment common to a collection of transactions.
type CardPaymentEnvironment26 struct {

	// Acquirer involved in the card payment transactions.
	Acquirer *Acquirer3 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment transactions.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation9 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction3 `xml:"POI,omitempty"`
}

func (c *CardPaymentEnvironment26) AddAcquirer() *Acquirer3 {
	c.Acquirer = new(Acquirer3)
	return c.Acquirer
}

func (c *CardPaymentEnvironment26) AddMerchant() *Organisation9 {
	c.Merchant = new(Organisation9)
	return c.Merchant
}

func (c *CardPaymentEnvironment26) AddPOI() *PointOfInteraction3 {
	c.POI = new(PointOfInteraction3)
	return c.POI
}

// Environment of the card payment transaction.
type CardPaymentEnvironment27 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer2 `xml:"Acqrr,omitempty"`

	// Merchant performing the transaction.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation8 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction3 `xml:"POI,omitempty"`

	// Payment card performing the transaction.
	Card *PaymentCard8 `xml:"Card"`

	// Cardholder involved in the card payment.
	Cardholder *Cardholder6 `xml:"Crdhldr,omitempty"`

	// Replacement of the message element Cardholder by a digital envelope using a cryptographic key.
	ProtectedCardholderData *ContentInformationType7 `xml:"PrtctdCrdhldrData,omitempty"`
}

func (c *CardPaymentEnvironment27) AddAcquirer() *Acquirer2 {
	c.Acquirer = new(Acquirer2)
	return c.Acquirer
}

func (c *CardPaymentEnvironment27) AddMerchant() *Organisation8 {
	c.Merchant = new(Organisation8)
	return c.Merchant
}

func (c *CardPaymentEnvironment27) AddPOI() *PointOfInteraction3 {
	c.POI = new(PointOfInteraction3)
	return c.POI
}

func (c *CardPaymentEnvironment27) AddCard() *PaymentCard8 {
	c.Card = new(PaymentCard8)
	return c.Card
}

func (c *CardPaymentEnvironment27) AddCardholder() *Cardholder6 {
	c.Cardholder = new(Cardholder6)
	return c.Cardholder
}

func (c *CardPaymentEnvironment27) AddProtectedCardholderData() *ContentInformationType7 {
	c.ProtectedCardholderData = new(ContentInformationType7)
	return c.ProtectedCardholderData
}

// Environment of the transaction given in a response to a request in a batch.
type CardPaymentEnvironment28 struct {

	// Acquirer involved in the card payment.
	AcquirerIdentification *GenericIdentification32 `xml:"AcqrrId,omitempty"`

	// Identification of the merchant.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI (Point Of Interaction) performing the transaction.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`

	// Replacement of the message element PlainCardData by a digital envelope using a cryptographic key.
	ProtectedCardData *ContentInformationType7 `xml:"PrtctdCardData,omitempty"`

	// Payment card performing the transaction.
	PlainCardData *PlainCardData3 `xml:"PlainCardData,omitempty"`
}

func (c *CardPaymentEnvironment28) AddAcquirerIdentification() *GenericIdentification32 {
	c.AcquirerIdentification = new(GenericIdentification32)
	return c.AcquirerIdentification
}

func (c *CardPaymentEnvironment28) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment28) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentEnvironment28) AddProtectedCardData() *ContentInformationType7 {
	c.ProtectedCardData = new(ContentInformationType7)
	return c.ProtectedCardData
}

func (c *CardPaymentEnvironment28) AddPlainCardData() *PlainCardData3 {
	c.PlainCardData = new(PlainCardData3)
	return c.PlainCardData
}

// Environment of the diagnostic exchange.
type CardPaymentEnvironment29 struct {

	// Acquirer involved in the card payment transaction.
	Acquirer *Acquirer2 `xml:"Acqrr"`

	// The availability of the acquirer to process transaction must be provided.
	AcquirerAvailabilityRequested *TrueFalseIndicator `xml:"AcqrrAvlbtyReqd,omitempty"`

	// Identification of the merchant requesting the diagnostic.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI (Point Of Interaction) requesting the diagnostic.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`

	// Data related to the components of the POI (Point Of Interaction) performing the payment transactions.
	POIComponent []*PointOfInteractionComponent4 `xml:"POICmpnt,omitempty"`
}

func (c *CardPaymentEnvironment29) AddAcquirer() *Acquirer2 {
	c.Acquirer = new(Acquirer2)
	return c.Acquirer
}

func (c *CardPaymentEnvironment29) SetAcquirerAvailabilityRequested(value string) {
	c.AcquirerAvailabilityRequested = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentEnvironment29) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment29) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentEnvironment29) AddPOIComponent() *PointOfInteractionComponent4 {
	newValue := new(PointOfInteractionComponent4)
	c.POIComponent = append(c.POIComponent, newValue)
	return newValue
}

// Environment of the transaction given in a response to a request.
type CardPaymentEnvironment3 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer1 `xml:"Acqrr,omitempty"`

	// Identification of the merchant.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI performing the transaction.
	POIIdentification *GenericIdentification32 `xml:"POIId"`

	// Sensitive data of the card (PlainCardData1 including the envelope), encrypted with a cryptographic key.
	ProtectedCardData *ContentInformationType1 `xml:"PrtctdCardData,omitempty"`

	// Payment card performing the transaction.
	PlainCardData *PlainCardData3 `xml:"PlainCardData,omitempty"`
}

func (c *CardPaymentEnvironment3) AddAcquirer() *Acquirer1 {
	c.Acquirer = new(Acquirer1)
	return c.Acquirer
}

func (c *CardPaymentEnvironment3) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment3) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentEnvironment3) AddProtectedCardData() *ContentInformationType1 {
	c.ProtectedCardData = new(ContentInformationType1)
	return c.ProtectedCardData
}

func (c *CardPaymentEnvironment3) AddPlainCardData() *PlainCardData3 {
	c.PlainCardData = new(PlainCardData3)
	return c.PlainCardData
}

// Environment of the transaction.
type CardPaymentEnvironment30 struct {

	// Acquirer involved in the card payment transaction.
	AcquirerIdentification *GenericIdentification32 `xml:"AcqrrId,omitempty"`

	// Merchant involved in the card payment transaction.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction3 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard7 `xml:"Card"`

	// Language selected for the cardholder interface during the transaction.
	CardholderLanguage *ISO2ALanguageCode `xml:"CrdhldrLang,omitempty"`
}

func (c *CardPaymentEnvironment30) AddAcquirerIdentification() *GenericIdentification32 {
	c.AcquirerIdentification = new(GenericIdentification32)
	return c.AcquirerIdentification
}

func (c *CardPaymentEnvironment30) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment30) AddPOI() *PointOfInteraction3 {
	c.POI = new(PointOfInteraction3)
	return c.POI
}

func (c *CardPaymentEnvironment30) AddCard() *PaymentCard7 {
	c.Card = new(PaymentCard7)
	return c.Card
}

func (c *CardPaymentEnvironment30) SetCardholderLanguage(value string) {
	c.CardholderLanguage = (*ISO2ALanguageCode)(&value)
}

// Environment of the diagnostic exchange.
type CardPaymentEnvironment31 struct {

	// Acquirer involved in the card payment transaction.
	Acquirer *Acquirer2 `xml:"Acqrr"`

	// Indicates if the acquirer is available to perform payment transactions.
	AcquirerAvailable *TrueFalseIndicator `xml:"AcqrrAvlbl,omitempty"`

	// Identification of the merchant requesting the diagnostic.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI (Point Of Interaction) requesting the diagnostic.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`
}

func (c *CardPaymentEnvironment31) AddAcquirer() *Acquirer2 {
	c.Acquirer = new(Acquirer2)
	return c.Acquirer
}

func (c *CardPaymentEnvironment31) SetAcquirerAvailable(value string) {
	c.AcquirerAvailable = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentEnvironment31) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment31) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

// Environment of the transaction.
type CardPaymentEnvironment32 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer4 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment transaction.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation8 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction4 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard9 `xml:"Card"`

	// Device used by the customer to perform the payment transaction.
	CustomerDevice *CustomerDevice1 `xml:"CstmrDvc,omitempty"`

	// Container for tenders used by the customer to perform the payment transaction.
	Wallet *CustomerDevice1 `xml:"Wllt,omitempty"`

	// Payment token information.
	PaymentToken *CardPaymentToken1 `xml:"PmtTkn,omitempty"`

	// Cardholder involved in the card payment.
	Cardholder *Cardholder7 `xml:"Crdhldr,omitempty"`

	// Replacement of the message element Cardholder by a digital envelope using a cryptographic key.
	ProtectedCardholderData *ContentInformationType10 `xml:"PrtctdCrdhldrData,omitempty"`
}

func (c *CardPaymentEnvironment32) AddAcquirer() *Acquirer4 {
	c.Acquirer = new(Acquirer4)
	return c.Acquirer
}

func (c *CardPaymentEnvironment32) AddMerchant() *Organisation8 {
	c.Merchant = new(Organisation8)
	return c.Merchant
}

func (c *CardPaymentEnvironment32) AddPOI() *PointOfInteraction4 {
	c.POI = new(PointOfInteraction4)
	return c.POI
}

func (c *CardPaymentEnvironment32) AddCard() *PaymentCard9 {
	c.Card = new(PaymentCard9)
	return c.Card
}

func (c *CardPaymentEnvironment32) AddCustomerDevice() *CustomerDevice1 {
	c.CustomerDevice = new(CustomerDevice1)
	return c.CustomerDevice
}

func (c *CardPaymentEnvironment32) AddWallet() *CustomerDevice1 {
	c.Wallet = new(CustomerDevice1)
	return c.Wallet
}

func (c *CardPaymentEnvironment32) AddPaymentToken() *CardPaymentToken1 {
	c.PaymentToken = new(CardPaymentToken1)
	return c.PaymentToken
}

func (c *CardPaymentEnvironment32) AddCardholder() *Cardholder7 {
	c.Cardholder = new(Cardholder7)
	return c.Cardholder
}

func (c *CardPaymentEnvironment32) AddProtectedCardholderData() *ContentInformationType10 {
	c.ProtectedCardholderData = new(ContentInformationType10)
	return c.ProtectedCardholderData
}

// Environment of the transaction given in a response to a request.
type CardPaymentEnvironment33 struct {

	// Acquirer involved in the card payment.
	AcquirerIdentification *GenericIdentification53 `xml:"AcqrrId,omitempty"`

	// Identification of the merchant.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI (Point Of Interaction) performing the transaction.
	POIIdentification *GenericIdentification32 `xml:"POIId"`

	// Card performing the transaction.
	Card *PaymentCard10 `xml:"Card,omitempty"`

	// Payment token information.
	PaymentToken *CardPaymentToken2 `xml:"PmtTkn,omitempty"`
}

func (c *CardPaymentEnvironment33) AddAcquirerIdentification() *GenericIdentification53 {
	c.AcquirerIdentification = new(GenericIdentification53)
	return c.AcquirerIdentification
}

func (c *CardPaymentEnvironment33) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment33) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentEnvironment33) AddCard() *PaymentCard10 {
	c.Card = new(PaymentCard10)
	return c.Card
}

func (c *CardPaymentEnvironment33) AddPaymentToken() *CardPaymentToken2 {
	c.PaymentToken = new(CardPaymentToken2)
	return c.PaymentToken
}

// Environment of the transaction.
type CardPaymentEnvironment34 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer4 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation8 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction4 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard11 `xml:"Card,omitempty"`

	// Device used by the customer to perform the payment.
	CustomerDevice *CustomerDevice1 `xml:"CstmrDvc,omitempty"`

	// Container of tenders used by the customer to perform the payment.
	Wallet *CustomerDevice1 `xml:"Wllt,omitempty"`

	// Payment token information.
	PaymentToken *CardPaymentToken3 `xml:"PmtTkn,omitempty"`

	// Cardholder involved in the card payment.
	Cardholder *Cardholder8 `xml:"Crdhldr,omitempty"`

	// Replacement of the message element Cardholder by a digital envelope using a cryptographic key.
	ProtectedCardholderData *ContentInformationType10 `xml:"PrtctdCrdhldrData,omitempty"`
}

func (c *CardPaymentEnvironment34) AddAcquirer() *Acquirer4 {
	c.Acquirer = new(Acquirer4)
	return c.Acquirer
}

func (c *CardPaymentEnvironment34) AddMerchant() *Organisation8 {
	c.Merchant = new(Organisation8)
	return c.Merchant
}

func (c *CardPaymentEnvironment34) AddPOI() *PointOfInteraction4 {
	c.POI = new(PointOfInteraction4)
	return c.POI
}

func (c *CardPaymentEnvironment34) AddCard() *PaymentCard11 {
	c.Card = new(PaymentCard11)
	return c.Card
}

func (c *CardPaymentEnvironment34) AddCustomerDevice() *CustomerDevice1 {
	c.CustomerDevice = new(CustomerDevice1)
	return c.CustomerDevice
}

func (c *CardPaymentEnvironment34) AddWallet() *CustomerDevice1 {
	c.Wallet = new(CustomerDevice1)
	return c.Wallet
}

func (c *CardPaymentEnvironment34) AddPaymentToken() *CardPaymentToken3 {
	c.PaymentToken = new(CardPaymentToken3)
	return c.PaymentToken
}

func (c *CardPaymentEnvironment34) AddCardholder() *Cardholder8 {
	c.Cardholder = new(Cardholder8)
	return c.Cardholder
}

func (c *CardPaymentEnvironment34) AddProtectedCardholderData() *ContentInformationType10 {
	c.ProtectedCardholderData = new(ContentInformationType10)
	return c.ProtectedCardholderData
}

// Environment of the cancellation.
type CardPaymentEnvironment35 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer4 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment cancellation.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation8 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction4 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard11 `xml:"Card"`

	// Device used by the customer to perform the payment.
	CustomerDevice *CustomerDevice1 `xml:"CstmrDvc,omitempty"`

	// Container of tenders used by the customer to perform the payment.
	Wallet *CustomerDevice1 `xml:"Wllt,omitempty"`

	// Payment token information.
	PaymentToken *CardPaymentToken3 `xml:"PmtTkn,omitempty"`
}

func (c *CardPaymentEnvironment35) AddAcquirer() *Acquirer4 {
	c.Acquirer = new(Acquirer4)
	return c.Acquirer
}

func (c *CardPaymentEnvironment35) AddMerchant() *Organisation8 {
	c.Merchant = new(Organisation8)
	return c.Merchant
}

func (c *CardPaymentEnvironment35) AddPOI() *PointOfInteraction4 {
	c.POI = new(PointOfInteraction4)
	return c.POI
}

func (c *CardPaymentEnvironment35) AddCard() *PaymentCard11 {
	c.Card = new(PaymentCard11)
	return c.Card
}

func (c *CardPaymentEnvironment35) AddCustomerDevice() *CustomerDevice1 {
	c.CustomerDevice = new(CustomerDevice1)
	return c.CustomerDevice
}

func (c *CardPaymentEnvironment35) AddWallet() *CustomerDevice1 {
	c.Wallet = new(CustomerDevice1)
	return c.Wallet
}

func (c *CardPaymentEnvironment35) AddPaymentToken() *CardPaymentToken3 {
	c.PaymentToken = new(CardPaymentToken3)
	return c.PaymentToken
}

// Environment of the transaction.
type CardPaymentEnvironment36 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer4 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation8 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction4 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard11 `xml:"Card"`

	// Device used by the customer to perform the payment.
	CustomerDevice *CustomerDevice1 `xml:"CstmrDvc,omitempty"`

	// Container of tenders used by the customer to perform the payment.
	Wallet *CustomerDevice1 `xml:"Wllt,omitempty"`

	// Payment token information.
	PaymentToken *CardPaymentToken3 `xml:"PmtTkn,omitempty"`
}

func (c *CardPaymentEnvironment36) AddAcquirer() *Acquirer4 {
	c.Acquirer = new(Acquirer4)
	return c.Acquirer
}

func (c *CardPaymentEnvironment36) AddMerchant() *Organisation8 {
	c.Merchant = new(Organisation8)
	return c.Merchant
}

func (c *CardPaymentEnvironment36) AddPOI() *PointOfInteraction4 {
	c.POI = new(PointOfInteraction4)
	return c.POI
}

func (c *CardPaymentEnvironment36) AddCard() *PaymentCard11 {
	c.Card = new(PaymentCard11)
	return c.Card
}

func (c *CardPaymentEnvironment36) AddCustomerDevice() *CustomerDevice1 {
	c.CustomerDevice = new(CustomerDevice1)
	return c.CustomerDevice
}

func (c *CardPaymentEnvironment36) AddWallet() *CustomerDevice1 {
	c.Wallet = new(CustomerDevice1)
	return c.Wallet
}

func (c *CardPaymentEnvironment36) AddPaymentToken() *CardPaymentToken3 {
	c.PaymentToken = new(CardPaymentToken3)
	return c.PaymentToken
}

// Environment of the transaction.
type CardPaymentEnvironment37 struct {

	// Acquirer involved in the card payment reconciliation.
	Acquirer *Acquirer4 `xml:"Acqrr"`

	// Identification of the merchant requesting the reconciliation.
	MerchantIdentification *GenericIdentification53 `xml:"MrchntId,omitempty"`

	// Identification of the POI (Point Of Interaction) requesting the reconciliation.
	POIIdentification *GenericIdentification53 `xml:"POIId,omitempty"`

	// Data related to the components of the POI (Point Of Interaction) that have been performed the payment transactions.
	POIComponent []*PointOfInteractionComponent5 `xml:"POICmpnt,omitempty"`
}

func (c *CardPaymentEnvironment37) AddAcquirer() *Acquirer4 {
	c.Acquirer = new(Acquirer4)
	return c.Acquirer
}

func (c *CardPaymentEnvironment37) AddMerchantIdentification() *GenericIdentification53 {
	c.MerchantIdentification = new(GenericIdentification53)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment37) AddPOIIdentification() *GenericIdentification53 {
	c.POIIdentification = new(GenericIdentification53)
	return c.POIIdentification
}

func (c *CardPaymentEnvironment37) AddPOIComponent() *PointOfInteractionComponent5 {
	newValue := new(PointOfInteractionComponent5)
	c.POIComponent = append(c.POIComponent, newValue)
	return newValue
}

// Environment of the reconciliation for the acquirer.
type CardPaymentEnvironment38 struct {

	// Acquirer involved in the card payment reconciliation.
	AcquirerIdentification *GenericIdentification53 `xml:"AcqrrId"`

	// Identification of the merchant requesting the reconciliation.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI requesting the reconciliation.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`
}

func (c *CardPaymentEnvironment38) AddAcquirerIdentification() *GenericIdentification53 {
	c.AcquirerIdentification = new(GenericIdentification53)
	return c.AcquirerIdentification
}

func (c *CardPaymentEnvironment38) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment38) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

// Environment common to a collection of transactions.
type CardPaymentEnvironment39 struct {

	// Acquirer involved in the card payment transactions.
	Acquirer *Acquirer5 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment transactions.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation9 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction4 `xml:"POI,omitempty"`
}

func (c *CardPaymentEnvironment39) AddAcquirer() *Acquirer5 {
	c.Acquirer = new(Acquirer5)
	return c.Acquirer
}

func (c *CardPaymentEnvironment39) AddMerchant() *Organisation9 {
	c.Merchant = new(Organisation9)
	return c.Merchant
}

func (c *CardPaymentEnvironment39) AddPOI() *PointOfInteraction4 {
	c.POI = new(PointOfInteraction4)
	return c.POI
}

// Environment of the cancellation.
type CardPaymentEnvironment4 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer1 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment cancellation.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation5 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction1 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard3 `xml:"Card"`
}

func (c *CardPaymentEnvironment4) AddAcquirer() *Acquirer1 {
	c.Acquirer = new(Acquirer1)
	return c.Acquirer
}

func (c *CardPaymentEnvironment4) AddMerchant() *Organisation5 {
	c.Merchant = new(Organisation5)
	return c.Merchant
}

func (c *CardPaymentEnvironment4) AddPOI() *PointOfInteraction1 {
	c.POI = new(PointOfInteraction1)
	return c.POI
}

func (c *CardPaymentEnvironment4) AddCard() *PaymentCard3 {
	c.Card = new(PaymentCard3)
	return c.Card
}

// Environment of the card payment transaction.
type CardPaymentEnvironment40 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer4 `xml:"Acqrr,omitempty"`

	// Merchant performing the transaction.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation8 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction4 `xml:"POI,omitempty"`

	// Payment card performing the transaction.
	Card *PaymentCard11 `xml:"Card"`

	// Device used by the customer to perform the payment.
	CustomerDevice *CustomerDevice1 `xml:"CstmrDvc,omitempty"`

	// Container of tenders used by the customer to perform the payment.
	Wallet *CustomerDevice1 `xml:"Wllt,omitempty"`

	// Payment token information
	PaymentToken *CardPaymentToken3 `xml:"PmtTkn,omitempty"`

	// Cardholder involved in the card payment.
	Cardholder *Cardholder8 `xml:"Crdhldr,omitempty"`

	// Replacement of the message element Cardholder by a digital envelope using a cryptographic key.
	ProtectedCardholderData *ContentInformationType10 `xml:"PrtctdCrdhldrData,omitempty"`
}

func (c *CardPaymentEnvironment40) AddAcquirer() *Acquirer4 {
	c.Acquirer = new(Acquirer4)
	return c.Acquirer
}

func (c *CardPaymentEnvironment40) AddMerchant() *Organisation8 {
	c.Merchant = new(Organisation8)
	return c.Merchant
}

func (c *CardPaymentEnvironment40) AddPOI() *PointOfInteraction4 {
	c.POI = new(PointOfInteraction4)
	return c.POI
}

func (c *CardPaymentEnvironment40) AddCard() *PaymentCard11 {
	c.Card = new(PaymentCard11)
	return c.Card
}

func (c *CardPaymentEnvironment40) AddCustomerDevice() *CustomerDevice1 {
	c.CustomerDevice = new(CustomerDevice1)
	return c.CustomerDevice
}

func (c *CardPaymentEnvironment40) AddWallet() *CustomerDevice1 {
	c.Wallet = new(CustomerDevice1)
	return c.Wallet
}

func (c *CardPaymentEnvironment40) AddPaymentToken() *CardPaymentToken3 {
	c.PaymentToken = new(CardPaymentToken3)
	return c.PaymentToken
}

func (c *CardPaymentEnvironment40) AddCardholder() *Cardholder8 {
	c.Cardholder = new(Cardholder8)
	return c.Cardholder
}

func (c *CardPaymentEnvironment40) AddProtectedCardholderData() *ContentInformationType10 {
	c.ProtectedCardholderData = new(ContentInformationType10)
	return c.ProtectedCardholderData
}

// Environment of the transaction given in a response to a request in a batch.
type CardPaymentEnvironment41 struct {

	// Acquirer involved in the card payment.
	AcquirerIdentification *GenericIdentification53 `xml:"AcqrrId,omitempty"`

	// Identification of the merchant.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI (Point Of Interaction) performing the transaction.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`

	// Payment card performing the transaction.
	Card *PaymentCard10 `xml:"Card,omitempty"`

	// Payment token information.
	PaymentToken *CardPaymentToken2 `xml:"PmtTkn,omitempty"`
}

func (c *CardPaymentEnvironment41) AddAcquirerIdentification() *GenericIdentification53 {
	c.AcquirerIdentification = new(GenericIdentification53)
	return c.AcquirerIdentification
}

func (c *CardPaymentEnvironment41) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment41) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentEnvironment41) AddCard() *PaymentCard10 {
	c.Card = new(PaymentCard10)
	return c.Card
}

func (c *CardPaymentEnvironment41) AddPaymentToken() *CardPaymentToken2 {
	c.PaymentToken = new(CardPaymentToken2)
	return c.PaymentToken
}

// Environment of the diagnostic exchange.
type CardPaymentEnvironment42 struct {

	// Acquirer involved in the card payment transaction.
	Acquirer *Acquirer4 `xml:"Acqrr"`

	// The availability of the acquirer to process transaction must be provided.
	AcquirerAvailabilityRequested *TrueFalseIndicator `xml:"AcqrrAvlbtyReqd,omitempty"`

	// Identification of the merchant requesting the diagnostic.
	MerchantIdentification *GenericIdentification53 `xml:"MrchntId,omitempty"`

	// Identification of the POI (Point Of Interaction) requesting the diagnostic.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`

	// Data related to the components of the POI (Point Of Interaction) performing the payment transactions.
	POIComponent []*PointOfInteractionComponent5 `xml:"POICmpnt,omitempty"`
}

func (c *CardPaymentEnvironment42) AddAcquirer() *Acquirer4 {
	c.Acquirer = new(Acquirer4)
	return c.Acquirer
}

func (c *CardPaymentEnvironment42) SetAcquirerAvailabilityRequested(value string) {
	c.AcquirerAvailabilityRequested = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentEnvironment42) AddMerchantIdentification() *GenericIdentification53 {
	c.MerchantIdentification = new(GenericIdentification53)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment42) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentEnvironment42) AddPOIComponent() *PointOfInteractionComponent5 {
	newValue := new(PointOfInteractionComponent5)
	c.POIComponent = append(c.POIComponent, newValue)
	return newValue
}

// Environment of the diagnostic exchange.
type CardPaymentEnvironment43 struct {

	// Acquirer involved in the card payment transaction.
	Acquirer *Acquirer4 `xml:"Acqrr"`

	// Indicates if the acquirer is available to perform payment transactions.
	AcquirerAvailable *TrueFalseIndicator `xml:"AcqrrAvlbl,omitempty"`

	// Identification of the merchant requesting the diagnostic.
	MerchantIdentification *GenericIdentification53 `xml:"MrchntId,omitempty"`

	// Identification of the POI (Point Of Interaction) requesting the diagnostic.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`
}

func (c *CardPaymentEnvironment43) AddAcquirer() *Acquirer4 {
	c.Acquirer = new(Acquirer4)
	return c.Acquirer
}

func (c *CardPaymentEnvironment43) SetAcquirerAvailable(value string) {
	c.AcquirerAvailable = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentEnvironment43) AddMerchantIdentification() *GenericIdentification53 {
	c.MerchantIdentification = new(GenericIdentification53)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment43) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

// Environment of the transaction.
type CardPaymentEnvironment44 struct {

	// Acquirer involved in the card payment transaction.
	AcquirerIdentification *GenericIdentification53 `xml:"AcqrrId,omitempty"`

	// Merchant involved in the card payment transaction.
	MerchantIdentification *GenericIdentification53 `xml:"MrchntId,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction4 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard9 `xml:"Card"`

	// Language selected for the cardholder interface during the transaction.
	// Reference: ISO 639-1 (alpha-2) et ISO 639-2 (alpha-3).
	CardholderLanguage *LanguageCode `xml:"CrdhldrLang,omitempty"`
}

func (c *CardPaymentEnvironment44) AddAcquirerIdentification() *GenericIdentification53 {
	c.AcquirerIdentification = new(GenericIdentification53)
	return c.AcquirerIdentification
}

func (c *CardPaymentEnvironment44) AddMerchantIdentification() *GenericIdentification53 {
	c.MerchantIdentification = new(GenericIdentification53)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment44) AddPOI() *PointOfInteraction4 {
	c.POI = new(PointOfInteraction4)
	return c.POI
}

func (c *CardPaymentEnvironment44) AddCard() *PaymentCard9 {
	c.Card = new(PaymentCard9)
	return c.Card
}

func (c *CardPaymentEnvironment44) SetCardholderLanguage(value string) {
	c.CardholderLanguage = (*LanguageCode)(&value)
}

// Environment of the transaction.
type CardPaymentEnvironment45 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer4 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment transaction.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation25 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction5 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard21 `xml:"Card"`

	// Device used by the customer to perform the payment transaction.
	CustomerDevice *CustomerDevice1 `xml:"CstmrDvc,omitempty"`

	// Container for tenders used by the customer to perform the payment transaction.
	Wallet *CustomerDevice1 `xml:"Wllt,omitempty"`

	// Payment token information.
	PaymentToken *CardPaymentToken1 `xml:"PmtTkn,omitempty"`

	// Cardholder involved in the card payment.
	Cardholder *Cardholder10 `xml:"Crdhldr,omitempty"`

	// Replacement of the message element Cardholder by a digital envelope using a cryptographic key.
	ProtectedCardholderData *ContentInformationType10 `xml:"PrtctdCrdhldrData,omitempty"`
}

func (c *CardPaymentEnvironment45) AddAcquirer() *Acquirer4 {
	c.Acquirer = new(Acquirer4)
	return c.Acquirer
}

func (c *CardPaymentEnvironment45) AddMerchant() *Organisation25 {
	c.Merchant = new(Organisation25)
	return c.Merchant
}

func (c *CardPaymentEnvironment45) AddPOI() *PointOfInteraction5 {
	c.POI = new(PointOfInteraction5)
	return c.POI
}

func (c *CardPaymentEnvironment45) AddCard() *PaymentCard21 {
	c.Card = new(PaymentCard21)
	return c.Card
}

func (c *CardPaymentEnvironment45) AddCustomerDevice() *CustomerDevice1 {
	c.CustomerDevice = new(CustomerDevice1)
	return c.CustomerDevice
}

func (c *CardPaymentEnvironment45) AddWallet() *CustomerDevice1 {
	c.Wallet = new(CustomerDevice1)
	return c.Wallet
}

func (c *CardPaymentEnvironment45) AddPaymentToken() *CardPaymentToken1 {
	c.PaymentToken = new(CardPaymentToken1)
	return c.PaymentToken
}

func (c *CardPaymentEnvironment45) AddCardholder() *Cardholder10 {
	c.Cardholder = new(Cardholder10)
	return c.Cardholder
}

func (c *CardPaymentEnvironment45) AddProtectedCardholderData() *ContentInformationType10 {
	c.ProtectedCardholderData = new(ContentInformationType10)
	return c.ProtectedCardholderData
}

// Environment of the transaction given in a response to a request.
type CardPaymentEnvironment46 struct {

	// Acquirer involved in the card payment.
	AcquirerIdentification *GenericIdentification53 `xml:"AcqrrId,omitempty"`

	// Identification of the merchant.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI (Point Of Interaction) performing the transaction.
	POIIdentification *GenericIdentification32 `xml:"POIId"`

	// Card performing the transaction.
	Card *PaymentCard19 `xml:"Card,omitempty"`

	// Payment token information.
	PaymentToken *CardPaymentToken2 `xml:"PmtTkn,omitempty"`
}

func (c *CardPaymentEnvironment46) AddAcquirerIdentification() *GenericIdentification53 {
	c.AcquirerIdentification = new(GenericIdentification53)
	return c.AcquirerIdentification
}

func (c *CardPaymentEnvironment46) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment46) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentEnvironment46) AddCard() *PaymentCard19 {
	c.Card = new(PaymentCard19)
	return c.Card
}

func (c *CardPaymentEnvironment46) AddPaymentToken() *CardPaymentToken2 {
	c.PaymentToken = new(CardPaymentToken2)
	return c.PaymentToken
}

// Environment of the transaction.
type CardPaymentEnvironment47 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer4 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation25 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction5 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard20 `xml:"Card,omitempty"`

	// Device used by the customer to perform the payment.
	CustomerDevice *CustomerDevice1 `xml:"CstmrDvc,omitempty"`

	// Container of tenders used by the customer to perform the payment.
	Wallet *CustomerDevice1 `xml:"Wllt,omitempty"`

	// Payment token information.
	PaymentToken *CardPaymentToken3 `xml:"PmtTkn,omitempty"`

	// Cardholder involved in the card payment.
	Cardholder *Cardholder11 `xml:"Crdhldr,omitempty"`

	// Replacement of the message element Cardholder by a digital envelope using a cryptographic key.
	ProtectedCardholderData *ContentInformationType10 `xml:"PrtctdCrdhldrData,omitempty"`
}

func (c *CardPaymentEnvironment47) AddAcquirer() *Acquirer4 {
	c.Acquirer = new(Acquirer4)
	return c.Acquirer
}

func (c *CardPaymentEnvironment47) AddMerchant() *Organisation25 {
	c.Merchant = new(Organisation25)
	return c.Merchant
}

func (c *CardPaymentEnvironment47) AddPOI() *PointOfInteraction5 {
	c.POI = new(PointOfInteraction5)
	return c.POI
}

func (c *CardPaymentEnvironment47) AddCard() *PaymentCard20 {
	c.Card = new(PaymentCard20)
	return c.Card
}

func (c *CardPaymentEnvironment47) AddCustomerDevice() *CustomerDevice1 {
	c.CustomerDevice = new(CustomerDevice1)
	return c.CustomerDevice
}

func (c *CardPaymentEnvironment47) AddWallet() *CustomerDevice1 {
	c.Wallet = new(CustomerDevice1)
	return c.Wallet
}

func (c *CardPaymentEnvironment47) AddPaymentToken() *CardPaymentToken3 {
	c.PaymentToken = new(CardPaymentToken3)
	return c.PaymentToken
}

func (c *CardPaymentEnvironment47) AddCardholder() *Cardholder11 {
	c.Cardholder = new(Cardholder11)
	return c.Cardholder
}

func (c *CardPaymentEnvironment47) AddProtectedCardholderData() *ContentInformationType10 {
	c.ProtectedCardholderData = new(ContentInformationType10)
	return c.ProtectedCardholderData
}

// Environment of the cancellation.
type CardPaymentEnvironment48 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer4 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment cancellation.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation25 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction5 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard20 `xml:"Card"`

	// Device used by the customer to perform the payment.
	CustomerDevice *CustomerDevice1 `xml:"CstmrDvc,omitempty"`

	// Container of tenders used by the customer to perform the payment.
	Wallet *CustomerDevice1 `xml:"Wllt,omitempty"`

	// Payment token information.
	PaymentToken *CardPaymentToken3 `xml:"PmtTkn,omitempty"`
}

func (c *CardPaymentEnvironment48) AddAcquirer() *Acquirer4 {
	c.Acquirer = new(Acquirer4)
	return c.Acquirer
}

func (c *CardPaymentEnvironment48) AddMerchant() *Organisation25 {
	c.Merchant = new(Organisation25)
	return c.Merchant
}

func (c *CardPaymentEnvironment48) AddPOI() *PointOfInteraction5 {
	c.POI = new(PointOfInteraction5)
	return c.POI
}

func (c *CardPaymentEnvironment48) AddCard() *PaymentCard20 {
	c.Card = new(PaymentCard20)
	return c.Card
}

func (c *CardPaymentEnvironment48) AddCustomerDevice() *CustomerDevice1 {
	c.CustomerDevice = new(CustomerDevice1)
	return c.CustomerDevice
}

func (c *CardPaymentEnvironment48) AddWallet() *CustomerDevice1 {
	c.Wallet = new(CustomerDevice1)
	return c.Wallet
}

func (c *CardPaymentEnvironment48) AddPaymentToken() *CardPaymentToken3 {
	c.PaymentToken = new(CardPaymentToken3)
	return c.PaymentToken
}

// Environment of the transaction.
type CardPaymentEnvironment49 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer4 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation25 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction5 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard20 `xml:"Card"`

	// Device used by the customer to perform the payment.
	CustomerDevice *CustomerDevice1 `xml:"CstmrDvc,omitempty"`

	// Container of tenders used by the customer to perform the payment.
	Wallet *CustomerDevice1 `xml:"Wllt,omitempty"`

	// Payment token information.
	PaymentToken *CardPaymentToken3 `xml:"PmtTkn,omitempty"`
}

func (c *CardPaymentEnvironment49) AddAcquirer() *Acquirer4 {
	c.Acquirer = new(Acquirer4)
	return c.Acquirer
}

func (c *CardPaymentEnvironment49) AddMerchant() *Organisation25 {
	c.Merchant = new(Organisation25)
	return c.Merchant
}

func (c *CardPaymentEnvironment49) AddPOI() *PointOfInteraction5 {
	c.POI = new(PointOfInteraction5)
	return c.POI
}

func (c *CardPaymentEnvironment49) AddCard() *PaymentCard20 {
	c.Card = new(PaymentCard20)
	return c.Card
}

func (c *CardPaymentEnvironment49) AddCustomerDevice() *CustomerDevice1 {
	c.CustomerDevice = new(CustomerDevice1)
	return c.CustomerDevice
}

func (c *CardPaymentEnvironment49) AddWallet() *CustomerDevice1 {
	c.Wallet = new(CustomerDevice1)
	return c.Wallet
}

func (c *CardPaymentEnvironment49) AddPaymentToken() *CardPaymentToken3 {
	c.PaymentToken = new(CardPaymentToken3)
	return c.PaymentToken
}

// Environment common to a collection of transactions.
type CardPaymentEnvironment5 struct {

	// Acquirer involved in the card payment transactions.
	Acquirer *Acquirer1 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment transactions.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation5 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction1 `xml:"POI,omitempty"`
}

func (c *CardPaymentEnvironment5) AddAcquirer() *Acquirer1 {
	c.Acquirer = new(Acquirer1)
	return c.Acquirer
}

func (c *CardPaymentEnvironment5) AddMerchant() *Organisation5 {
	c.Merchant = new(Organisation5)
	return c.Merchant
}

func (c *CardPaymentEnvironment5) AddPOI() *PointOfInteraction1 {
	c.POI = new(PointOfInteraction1)
	return c.POI
}

// Environment of the transaction.
type CardPaymentEnvironment50 struct {

	// Acquirer involved in the card payment reconciliation.
	Acquirer *Acquirer4 `xml:"Acqrr"`

	// Identification of the merchant requesting the reconciliation.
	MerchantIdentification *GenericIdentification53 `xml:"MrchntId,omitempty"`

	// Identification of the POI (Point Of Interaction) requesting the reconciliation.
	POIIdentification *GenericIdentification53 `xml:"POIId,omitempty"`

	// Data related to the components of the POI (Point Of Interaction) that have been performed the payment transactions.
	POIComponent []*PointOfInteractionComponent6 `xml:"POICmpnt,omitempty"`
}

func (c *CardPaymentEnvironment50) AddAcquirer() *Acquirer4 {
	c.Acquirer = new(Acquirer4)
	return c.Acquirer
}

func (c *CardPaymentEnvironment50) AddMerchantIdentification() *GenericIdentification53 {
	c.MerchantIdentification = new(GenericIdentification53)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment50) AddPOIIdentification() *GenericIdentification53 {
	c.POIIdentification = new(GenericIdentification53)
	return c.POIIdentification
}

func (c *CardPaymentEnvironment50) AddPOIComponent() *PointOfInteractionComponent6 {
	newValue := new(PointOfInteractionComponent6)
	c.POIComponent = append(c.POIComponent, newValue)
	return newValue
}

// Environment common to a collection of transactions.
type CardPaymentEnvironment51 struct {

	// Acquirer involved in the card payment transactions.
	Acquirer *Acquirer5 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment transactions.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation9 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction5 `xml:"POI,omitempty"`
}

func (c *CardPaymentEnvironment51) AddAcquirer() *Acquirer5 {
	c.Acquirer = new(Acquirer5)
	return c.Acquirer
}

func (c *CardPaymentEnvironment51) AddMerchant() *Organisation9 {
	c.Merchant = new(Organisation9)
	return c.Merchant
}

func (c *CardPaymentEnvironment51) AddPOI() *PointOfInteraction5 {
	c.POI = new(PointOfInteraction5)
	return c.POI
}

// Environment of the card payment transaction.
type CardPaymentEnvironment52 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer4 `xml:"Acqrr,omitempty"`

	// Merchant performing the transaction.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation25 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction5 `xml:"POI,omitempty"`

	// Payment card performing the transaction.
	Card *PaymentCard20 `xml:"Card"`

	// Device used by the customer to perform the payment.
	CustomerDevice *CustomerDevice1 `xml:"CstmrDvc,omitempty"`

	// Container of tenders used by the customer to perform the payment.
	Wallet *CustomerDevice1 `xml:"Wllt,omitempty"`

	// Payment token information
	PaymentToken *CardPaymentToken3 `xml:"PmtTkn,omitempty"`

	// Cardholder involved in the card payment.
	Cardholder *Cardholder11 `xml:"Crdhldr,omitempty"`

	// Replacement of the message element Cardholder by a digital envelope using a cryptographic key.
	ProtectedCardholderData *ContentInformationType10 `xml:"PrtctdCrdhldrData,omitempty"`
}

func (c *CardPaymentEnvironment52) AddAcquirer() *Acquirer4 {
	c.Acquirer = new(Acquirer4)
	return c.Acquirer
}

func (c *CardPaymentEnvironment52) AddMerchant() *Organisation25 {
	c.Merchant = new(Organisation25)
	return c.Merchant
}

func (c *CardPaymentEnvironment52) AddPOI() *PointOfInteraction5 {
	c.POI = new(PointOfInteraction5)
	return c.POI
}

func (c *CardPaymentEnvironment52) AddCard() *PaymentCard20 {
	c.Card = new(PaymentCard20)
	return c.Card
}

func (c *CardPaymentEnvironment52) AddCustomerDevice() *CustomerDevice1 {
	c.CustomerDevice = new(CustomerDevice1)
	return c.CustomerDevice
}

func (c *CardPaymentEnvironment52) AddWallet() *CustomerDevice1 {
	c.Wallet = new(CustomerDevice1)
	return c.Wallet
}

func (c *CardPaymentEnvironment52) AddPaymentToken() *CardPaymentToken3 {
	c.PaymentToken = new(CardPaymentToken3)
	return c.PaymentToken
}

func (c *CardPaymentEnvironment52) AddCardholder() *Cardholder11 {
	c.Cardholder = new(Cardholder11)
	return c.Cardholder
}

func (c *CardPaymentEnvironment52) AddProtectedCardholderData() *ContentInformationType10 {
	c.ProtectedCardholderData = new(ContentInformationType10)
	return c.ProtectedCardholderData
}

// Environment of the transaction.
type CardPaymentEnvironment53 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer4 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment transaction.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation25 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction5 `xml:"POI,omitempty"`

	// Payment card performing the transaction.
	Card *PaymentCard21 `xml:"Card"`

	// Device used by the customer to perform the payment transaction.
	CustomerDevice *CustomerDevice1 `xml:"CstmrDvc,omitempty"`

	// Container for tenders used by the customer to perform the payment transaction.
	Wallet *CustomerDevice1 `xml:"Wllt,omitempty"`

	// Payment token information.
	PaymentToken *CardPaymentToken1 `xml:"PmtTkn,omitempty"`

	// Cardholder involved in the card payment.
	Cardholder *Cardholder10 `xml:"Crdhldr,omitempty"`

	// Replacement of the message element Cardholder by a digital envelope using a cryptographic key.
	ProtectedCardholderData *ContentInformationType10 `xml:"PrtctdCrdhldrData,omitempty"`
}

func (c *CardPaymentEnvironment53) AddAcquirer() *Acquirer4 {
	c.Acquirer = new(Acquirer4)
	return c.Acquirer
}

func (c *CardPaymentEnvironment53) AddMerchant() *Organisation25 {
	c.Merchant = new(Organisation25)
	return c.Merchant
}

func (c *CardPaymentEnvironment53) AddPOI() *PointOfInteraction5 {
	c.POI = new(PointOfInteraction5)
	return c.POI
}

func (c *CardPaymentEnvironment53) AddCard() *PaymentCard21 {
	c.Card = new(PaymentCard21)
	return c.Card
}

func (c *CardPaymentEnvironment53) AddCustomerDevice() *CustomerDevice1 {
	c.CustomerDevice = new(CustomerDevice1)
	return c.CustomerDevice
}

func (c *CardPaymentEnvironment53) AddWallet() *CustomerDevice1 {
	c.Wallet = new(CustomerDevice1)
	return c.Wallet
}

func (c *CardPaymentEnvironment53) AddPaymentToken() *CardPaymentToken1 {
	c.PaymentToken = new(CardPaymentToken1)
	return c.PaymentToken
}

func (c *CardPaymentEnvironment53) AddCardholder() *Cardholder10 {
	c.Cardholder = new(Cardholder10)
	return c.Cardholder
}

func (c *CardPaymentEnvironment53) AddProtectedCardholderData() *ContentInformationType10 {
	c.ProtectedCardholderData = new(ContentInformationType10)
	return c.ProtectedCardholderData
}

// Environment of the transaction.
type CardPaymentEnvironment54 struct {

	// Acquirer involved in the card payment.
	AcquirerIdentification *GenericIdentification53 `xml:"AcqrrId,omitempty"`

	// Merchant performing the card payment.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`

	// Payment card performing the transaction.
	Card *PaymentCard19 `xml:"Card,omitempty"`

	// Payment token information.
	PaymentToken *CardPaymentToken2 `xml:"PmtTkn,omitempty"`
}

func (c *CardPaymentEnvironment54) AddAcquirerIdentification() *GenericIdentification53 {
	c.AcquirerIdentification = new(GenericIdentification53)
	return c.AcquirerIdentification
}

func (c *CardPaymentEnvironment54) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment54) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentEnvironment54) AddCard() *PaymentCard19 {
	c.Card = new(PaymentCard19)
	return c.Card
}

func (c *CardPaymentEnvironment54) AddPaymentToken() *CardPaymentToken2 {
	c.PaymentToken = new(CardPaymentToken2)
	return c.PaymentToken
}

// Environment of the diagnostic exchange.
type CardPaymentEnvironment55 struct {

	// Acquirer involved in the card payment transaction.
	Acquirer *Acquirer4 `xml:"Acqrr"`

	// The availability of the acquirer to process transaction must be provided.
	AcquirerAvailabilityRequested *TrueFalseIndicator `xml:"AcqrrAvlbtyReqd,omitempty"`

	// Identification of the merchant requesting the diagnostic.
	MerchantIdentification *GenericIdentification53 `xml:"MrchntId,omitempty"`

	// Identification of the POI (Point Of Interaction) requesting the diagnostic.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`

	// Data related to the components of the POI (Point Of Interaction) performing the payment transactions.
	POIComponent []*PointOfInteractionComponent6 `xml:"POICmpnt,omitempty"`
}

func (c *CardPaymentEnvironment55) AddAcquirer() *Acquirer4 {
	c.Acquirer = new(Acquirer4)
	return c.Acquirer
}

func (c *CardPaymentEnvironment55) SetAcquirerAvailabilityRequested(value string) {
	c.AcquirerAvailabilityRequested = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentEnvironment55) AddMerchantIdentification() *GenericIdentification53 {
	c.MerchantIdentification = new(GenericIdentification53)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment55) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentEnvironment55) AddPOIComponent() *PointOfInteractionComponent6 {
	newValue := new(PointOfInteractionComponent6)
	c.POIComponent = append(c.POIComponent, newValue)
	return newValue
}

// Environment of the transaction.
type CardPaymentEnvironment56 struct {

	// Acquirer involved in the card payment transaction.
	AcquirerIdentification *GenericIdentification53 `xml:"AcqrrId,omitempty"`

	// Merchant involved in the card payment transaction.
	MerchantIdentification *GenericIdentification53 `xml:"MrchntId,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction5 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard21 `xml:"Card"`

	// Language selected for the cardholder interface during the transaction.
	// Reference: ISO 639-1 (alpha-2) et ISO 639-2 (alpha-3).
	CardholderLanguage *LanguageCode `xml:"CrdhldrLang,omitempty"`
}

func (c *CardPaymentEnvironment56) AddAcquirerIdentification() *GenericIdentification53 {
	c.AcquirerIdentification = new(GenericIdentification53)
	return c.AcquirerIdentification
}

func (c *CardPaymentEnvironment56) AddMerchantIdentification() *GenericIdentification53 {
	c.MerchantIdentification = new(GenericIdentification53)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment56) AddPOI() *PointOfInteraction5 {
	c.POI = new(PointOfInteraction5)
	return c.POI
}

func (c *CardPaymentEnvironment56) AddCard() *PaymentCard21 {
	c.Card = new(PaymentCard21)
	return c.Card
}

func (c *CardPaymentEnvironment56) SetCardholderLanguage(value string) {
	c.CardholderLanguage = (*LanguageCode)(&value)
}

// Environment of the card payment transaction.
type CardPaymentEnvironment6 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer1 `xml:"Acqrr,omitempty"`

	// Merchant performing the transaction.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation5 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction1 `xml:"POI,omitempty"`

	// Payment card performing the transaction.
	Card *PaymentCard3 `xml:"Card"`

	// Cardholder involved in the card payment.
	Cardholder *Cardholder2 `xml:"Crdhldr,omitempty"`
}

func (c *CardPaymentEnvironment6) AddAcquirer() *Acquirer1 {
	c.Acquirer = new(Acquirer1)
	return c.Acquirer
}

func (c *CardPaymentEnvironment6) AddMerchant() *Organisation5 {
	c.Merchant = new(Organisation5)
	return c.Merchant
}

func (c *CardPaymentEnvironment6) AddPOI() *PointOfInteraction1 {
	c.POI = new(PointOfInteraction1)
	return c.POI
}

func (c *CardPaymentEnvironment6) AddCard() *PaymentCard3 {
	c.Card = new(PaymentCard3)
	return c.Card
}

func (c *CardPaymentEnvironment6) AddCardholder() *Cardholder2 {
	c.Cardholder = new(Cardholder2)
	return c.Cardholder
}

// Environment of the reconciliation exchange.
type CardPaymentEnvironment7 struct {

	// Acquirer involved in the card payment reconciliation.
	Acquirer *Acquirer1 `xml:"Acqrr"`

	// Identification of the merchant requesting the reconciliation.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI requesting the reconciliation.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`
}

func (c *CardPaymentEnvironment7) AddAcquirer() *Acquirer1 {
	c.Acquirer = new(Acquirer1)
	return c.Acquirer
}

func (c *CardPaymentEnvironment7) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment7) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

// Environment of the diagnostic exchange.
type CardPaymentEnvironment8 struct {

	// Version of the payment application configuration parameters of the POI.
	AcquirerParametersVersion *ISODateTime `xml:"AcqrrParamsVrsn"`

	// Identification of the merchant requesting the diagnostic.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI requesting the diagnostic.
	POIIdentification *GenericIdentification32 `xml:"POIId"`
}

func (c *CardPaymentEnvironment8) SetAcquirerParametersVersion(value string) {
	c.AcquirerParametersVersion = (*ISODateTime)(&value)
}

func (c *CardPaymentEnvironment8) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment8) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

// Environment of the transaction.
type CardPaymentEnvironment9 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer2 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment transaction.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation8 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction2 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard5 `xml:"Card"`

	// Cardholder involved in the card payment.
	Cardholder *Cardholder3 `xml:"Crdhldr,omitempty"`
}

func (c *CardPaymentEnvironment9) AddAcquirer() *Acquirer2 {
	c.Acquirer = new(Acquirer2)
	return c.Acquirer
}

func (c *CardPaymentEnvironment9) AddMerchant() *Organisation8 {
	c.Merchant = new(Organisation8)
	return c.Merchant
}

func (c *CardPaymentEnvironment9) AddPOI() *PointOfInteraction2 {
	c.POI = new(PointOfInteraction2)
	return c.POI
}

func (c *CardPaymentEnvironment9) AddCard() *PaymentCard5 {
	c.Card = new(PaymentCard5)
	return c.Card
}

func (c *CardPaymentEnvironment9) AddCardholder() *Cardholder3 {
	c.Cardholder = new(Cardholder3)
	return c.Cardholder
}
