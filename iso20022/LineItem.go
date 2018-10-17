package iso20022

// Unit of information  showing the related  provision of products and/or services and monetary summations reported as a discrete line items.
type LineItem10 struct {

	// The unique identification of this invoice line item.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Something that is produced and sold as the result of an industrial process.
	TradeProduct *TradeProduct1 `xml:"TradPdct,omitempty"`

	// Purchase order reference assigned by the buyer related to the provision of products and/or services for this line item.
	BuyerOrderIdentification *DocumentIdentification23 `xml:"BuyrOrdrId,omitempty"`

	// Contract reference related to the provision of products and/or services for this line item.
	ContractIdentification *DocumentIdentification22 `xml:"CtrctId,omitempty"`

	// Specific purchase account for recording debits and credits for accounting purposes.
	PurchaseAccountingAccount []*AccountingAccount1 `xml:"PurchsAcctgAcct,omitempty"`

	// Value of the price, eg, as a currency and value.
	NetPrice []*CurrencyAndAmount `xml:"NetPric,omitempty"`

	// Quantity and conversion factor on which the net price is based for this line item product and/or service.
	NetPriceQuantity *Quantity4 `xml:"NetPricQty,omitempty"`

	// Allowance or charge applied to the net price.
	NetPriceAllowanceCharge []*LineItemAllowanceCharge1 `xml:"NetPricAllwncChrg,omitempty"`

	// Net weight of the product.
	NetWeight *Quantity3 `xml:"NetWght,omitempty"`

	// Gross price of the product and/or service.
	GrossPrice []*CurrencyAndAmount `xml:"GrssPric,omitempty"`

	// Quantity and conversion factor on which the gross price is based for this line item product and/or service.
	GrossPriceQuantity *Quantity4 `xml:"GrssPricQty,omitempty"`

	// Gross weight of the product.
	GrossWeight *Quantity3 `xml:"GrssWght,omitempty"`

	// Logistics service charge for this line item.
	LogisticsCharge []*ChargesDetails2 `xml:"LogstcsChrg,omitempty"`

	// Amount of money due to the government or tax authority, according to various pre-defined parameters such as thresholds or income.
	Tax []*LineItemTax1 `xml:"Tax,omitempty"`

	// Allowance or charge specified for this line item.
	AllowanceCharge []*LineItemAllowanceCharge1 `xml:"AllwncChrg,omitempty"`

	// Modification on the value of goods and / or services. For example: rebate, discount, surcharge
	FinancialAdjustment []*Adjustment4 `xml:"FinAdjstmnt,omitempty"`

	// Quantity billed for this line item.
	BilledQuantity *Quantity3 `xml:"BlldQty,omitempty"`

	// Number of product packages delivered.
	PackageQuantity *DecimalNumber `xml:"PackgQty,omitempty"`

	// Number of units per package in this line item for a supply chain trade delivery.
	PerPackageUnitQuantity *Quantity3 `xml:"PerPackgUnitQty,omitempty"`

	// Physical packaging of the product.
	Packaging []*Packaging1 `xml:"Packgng,omitempty"`

	// Quantity that is free of charge for this line item.
	ChargeFreeQuantity *Quantity3 `xml:"ChrgFreeQty,omitempty"`

	// Quantity value on which the quantity measurement started for a line item. For instance the start amount of a meter reading for an electricity supplier.
	MeasureQuantityStart *Quantity3 `xml:"MeasrQtyStart,omitempty"`

	// Quantity value on which the quantity measurement ended for a line item. For instance the end amount of a meter reading for an electricity supplier.
	MeasureQuantityEnd *Quantity3 `xml:"MeasrQtyEnd,omitempty"`

	// Date/time on which the clock time measure started for a line item.
	MeasureDateTimeStart *ISODateTime `xml:"MeasrDtTmStart,omitempty"`

	// Date/time on which the clock time measure ended for a line item.
	MeasureDateTimeEnd *ISODateTime `xml:"MeasrDtTmEnd,omitempty"`

	// Party to whom the goods must be delivered in the end.
	ShipTo *TradeParty1 `xml:"ShipTo,omitempty"`

	// Specifies the applicable Incoterm and associated location.
	Incoterms *Incoterms3 `xml:"Incotrms,omitempty"`

	// Actual delivery date/time of the products and/or services for this line item.
	DeliveryDateTime *ISODateTime `xml:"DlvryDtTm,omitempty"`

	// Delivery note related to the delivery of the products and/or services for this line item.
	DeliveryNoteIdentification *DocumentIdentification22 `xml:"DlvryNoteId,omitempty"`

	// Monetary totals for this line item.
	MonetarySummation *LineItemMonetarySummation1 `xml:"MntrySummtn,omitempty"`

	// Note included in this line item.
	IncludedNote []*AdditionalInformation1 `xml:"InclNote,omitempty"`
}

func (l *LineItem10) SetIdentification(value string) {
	l.Identification = (*Max35Text)(&value)
}

func (l *LineItem10) AddTradeProduct() *TradeProduct1 {
	l.TradeProduct = new(TradeProduct1)
	return l.TradeProduct
}

func (l *LineItem10) AddBuyerOrderIdentification() *DocumentIdentification23 {
	l.BuyerOrderIdentification = new(DocumentIdentification23)
	return l.BuyerOrderIdentification
}

func (l *LineItem10) AddContractIdentification() *DocumentIdentification22 {
	l.ContractIdentification = new(DocumentIdentification22)
	return l.ContractIdentification
}

func (l *LineItem10) AddPurchaseAccountingAccount() *AccountingAccount1 {
	newValue := new(AccountingAccount1)
	l.PurchaseAccountingAccount = append(l.PurchaseAccountingAccount, newValue)
	return newValue
}

func (l *LineItem10) AddNetPrice(value, currency string) {
	l.NetPrice = append(l.NetPrice, NewCurrencyAndAmount(value, currency))
}

func (l *LineItem10) AddNetPriceQuantity() *Quantity4 {
	l.NetPriceQuantity = new(Quantity4)
	return l.NetPriceQuantity
}

func (l *LineItem10) AddNetPriceAllowanceCharge() *LineItemAllowanceCharge1 {
	newValue := new(LineItemAllowanceCharge1)
	l.NetPriceAllowanceCharge = append(l.NetPriceAllowanceCharge, newValue)
	return newValue
}

func (l *LineItem10) AddNetWeight() *Quantity3 {
	l.NetWeight = new(Quantity3)
	return l.NetWeight
}

func (l *LineItem10) AddGrossPrice(value, currency string) {
	l.GrossPrice = append(l.GrossPrice, NewCurrencyAndAmount(value, currency))
}

func (l *LineItem10) AddGrossPriceQuantity() *Quantity4 {
	l.GrossPriceQuantity = new(Quantity4)
	return l.GrossPriceQuantity
}

func (l *LineItem10) AddGrossWeight() *Quantity3 {
	l.GrossWeight = new(Quantity3)
	return l.GrossWeight
}

func (l *LineItem10) AddLogisticsCharge() *ChargesDetails2 {
	newValue := new(ChargesDetails2)
	l.LogisticsCharge = append(l.LogisticsCharge, newValue)
	return newValue
}

func (l *LineItem10) AddTax() *LineItemTax1 {
	newValue := new(LineItemTax1)
	l.Tax = append(l.Tax, newValue)
	return newValue
}

func (l *LineItem10) AddAllowanceCharge() *LineItemAllowanceCharge1 {
	newValue := new(LineItemAllowanceCharge1)
	l.AllowanceCharge = append(l.AllowanceCharge, newValue)
	return newValue
}

func (l *LineItem10) AddFinancialAdjustment() *Adjustment4 {
	newValue := new(Adjustment4)
	l.FinancialAdjustment = append(l.FinancialAdjustment, newValue)
	return newValue
}

func (l *LineItem10) AddBilledQuantity() *Quantity3 {
	l.BilledQuantity = new(Quantity3)
	return l.BilledQuantity
}

func (l *LineItem10) SetPackageQuantity(value string) {
	l.PackageQuantity = (*DecimalNumber)(&value)
}

func (l *LineItem10) AddPerPackageUnitQuantity() *Quantity3 {
	l.PerPackageUnitQuantity = new(Quantity3)
	return l.PerPackageUnitQuantity
}

func (l *LineItem10) AddPackaging() *Packaging1 {
	newValue := new(Packaging1)
	l.Packaging = append(l.Packaging, newValue)
	return newValue
}

func (l *LineItem10) AddChargeFreeQuantity() *Quantity3 {
	l.ChargeFreeQuantity = new(Quantity3)
	return l.ChargeFreeQuantity
}

func (l *LineItem10) AddMeasureQuantityStart() *Quantity3 {
	l.MeasureQuantityStart = new(Quantity3)
	return l.MeasureQuantityStart
}

func (l *LineItem10) AddMeasureQuantityEnd() *Quantity3 {
	l.MeasureQuantityEnd = new(Quantity3)
	return l.MeasureQuantityEnd
}

func (l *LineItem10) SetMeasureDateTimeStart(value string) {
	l.MeasureDateTimeStart = (*ISODateTime)(&value)
}

func (l *LineItem10) SetMeasureDateTimeEnd(value string) {
	l.MeasureDateTimeEnd = (*ISODateTime)(&value)
}

func (l *LineItem10) AddShipTo() *TradeParty1 {
	l.ShipTo = new(TradeParty1)
	return l.ShipTo
}

func (l *LineItem10) AddIncoterms() *Incoterms3 {
	l.Incoterms = new(Incoterms3)
	return l.Incoterms
}

func (l *LineItem10) SetDeliveryDateTime(value string) {
	l.DeliveryDateTime = (*ISODateTime)(&value)
}

func (l *LineItem10) AddDeliveryNoteIdentification() *DocumentIdentification22 {
	l.DeliveryNoteIdentification = new(DocumentIdentification22)
	return l.DeliveryNoteIdentification
}

func (l *LineItem10) AddMonetarySummation() *LineItemMonetarySummation1 {
	l.MonetarySummation = new(LineItemMonetarySummation1)
	return l.MonetarySummation
}

func (l *LineItem10) AddIncludedNote() *AdditionalInformation1 {
	newValue := new(AdditionalInformation1)
	l.IncludedNote = append(l.IncludedNote, newValue)
	return newValue
}

// Commercial details of a trade transaction between a buyer and a seller.
type LineItem11 struct {

	// Information about the goods and/or services of a trade transaction.
	GoodsDescription *Max70Text `xml:"GoodsDesc,omitempty"`

	// Indicates whether or not partial shipments are allowed.
	PartialShipment *YesNoIndicator `xml:"PrtlShipmnt"`

	// Indicates whether or not transshipment of goods is allowed.
	TransShipment *YesNoIndicator `xml:"TrnsShipmnt,omitempty"`

	// Specifies an earliest shipment date and a latest shipment date.
	ShipmentDateRange *ShipmentDateRange1 `xml:"ShipmntDtRg,omitempty"`

	// Goods or services that are purchased.
	LineItemDetails []*LineItemDetails10 `xml:"LineItmDtls"`

	// Specifies the total amount of all line items (incl. their adjustments).
	LineItemsTotalAmount *CurrencyAndAmount `xml:"LineItmsTtlAmt"`

	// Information related to the conveyance of goods.
	RoutingSummary *TransportMeans5 `xml:"RtgSummry,omitempty"`

	// Specifies the applicable Incoterms and associated location. Latest version of Incoterms in effect at the date of message creation.
	Incoterms *Incoterms4 `xml:"Incotrms,omitempty"`

	// Variance on price for the goods.
	Adjustment []*Adjustment7 `xml:"Adjstmnt,omitempty"`

	// Maximum charges related to the conveyance of goods.
	FreightCharges *Charge24 `xml:"FrghtChrgs,omitempty"`

	// Amount of money due to the government or tax authority, according to various pre-defined parameters linked to the value of the goods in a trade transaction.
	Tax []*Tax23 `xml:"Tax,omitempty"`

	// Total net amount of a trade transaction. Total amount resulting from the gross amount plus freight charges, tax and plus/minus Adjustments.
	TotalNetAmount *CurrencyAndAmount `xml:"TtlNetAmt"`

	// Information important for the users of the message/service, which cannot be captured in any other message component/element. For example: Warehouse number.
	BuyerDefinedInformation []*UserDefinedInformation1 `xml:"BuyrDfndInf,omitempty"`

	// Information important for the users of the message/service, which cannot be captured in any other message component/element. For example: Warehouse number.
	SellerDefinedInformation []*UserDefinedInformation1 `xml:"SellrDfndInf,omitempty"`
}

func (l *LineItem11) SetGoodsDescription(value string) {
	l.GoodsDescription = (*Max70Text)(&value)
}

func (l *LineItem11) SetPartialShipment(value string) {
	l.PartialShipment = (*YesNoIndicator)(&value)
}

func (l *LineItem11) SetTransShipment(value string) {
	l.TransShipment = (*YesNoIndicator)(&value)
}

func (l *LineItem11) AddShipmentDateRange() *ShipmentDateRange1 {
	l.ShipmentDateRange = new(ShipmentDateRange1)
	return l.ShipmentDateRange
}

func (l *LineItem11) AddLineItemDetails() *LineItemDetails10 {
	newValue := new(LineItemDetails10)
	l.LineItemDetails = append(l.LineItemDetails, newValue)
	return newValue
}

func (l *LineItem11) SetLineItemsTotalAmount(value, currency string) {
	l.LineItemsTotalAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem11) AddRoutingSummary() *TransportMeans5 {
	l.RoutingSummary = new(TransportMeans5)
	return l.RoutingSummary
}

func (l *LineItem11) AddIncoterms() *Incoterms4 {
	l.Incoterms = new(Incoterms4)
	return l.Incoterms
}

func (l *LineItem11) AddAdjustment() *Adjustment7 {
	newValue := new(Adjustment7)
	l.Adjustment = append(l.Adjustment, newValue)
	return newValue
}

func (l *LineItem11) AddFreightCharges() *Charge24 {
	l.FreightCharges = new(Charge24)
	return l.FreightCharges
}

func (l *LineItem11) AddTax() *Tax23 {
	newValue := new(Tax23)
	l.Tax = append(l.Tax, newValue)
	return newValue
}

func (l *LineItem11) SetTotalNetAmount(value, currency string) {
	l.TotalNetAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem11) AddBuyerDefinedInformation() *UserDefinedInformation1 {
	newValue := new(UserDefinedInformation1)
	l.BuyerDefinedInformation = append(l.BuyerDefinedInformation, newValue)
	return newValue
}

func (l *LineItem11) AddSellerDefinedInformation() *UserDefinedInformation1 {
	newValue := new(UserDefinedInformation1)
	l.SellerDefinedInformation = append(l.SellerDefinedInformation, newValue)
	return newValue
}

// Commercial details of a trade transaction between a buyer and a seller.
type LineItem12 struct {

	// Reference to the purchase order of the underlying transaction.
	PurchaseOrderReference *DocumentIdentification7 `xml:"PurchsOrdrRef"`

	// Specifies whether this current invoice is the last submission against the purchase order referenced.
	FinalSubmission *YesNoIndicator `xml:"FnlSubmissn"`

	// Goods which are the subject of the invoice.
	CommercialLineItems []*LineItemDetails11 `xml:"ComrclLineItms"`

	// Specifies the total amount of all line items (incl. their adjustments).
	LineItemsTotalAmount *CurrencyAndAmount `xml:"LineItmsTtlAmt"`

	// Variance on price for the goods.
	Adjustment []*Adjustment6 `xml:"Adjstmnt,omitempty"`

	// Charges related to the conveyance of goods.
	FreightCharges *Charge25 `xml:"FrghtChrgs,omitempty"`

	// Amount of money due to the government or tax authority, according to various pre-defined parameters linked to the value of the goods in a trade transaction.
	Tax []*Tax22 `xml:"Tax,omitempty"`

	// Total net amount of a trade transaction. Total amount resulting from the gross amount plus freight charges, tax and plus/minus Adjustments.
	TotalNetAmount *CurrencyAndAmount `xml:"TtlNetAmt"`

	// Information important for the users of the message/service, which cannot be captured in any other message component/element. For example: Warehouse number.
	BuyerDefinedInformation []*UserDefinedInformation1 `xml:"BuyrDfndInf,omitempty"`

	// Information important for the users of the message/service, which cannot be captured in any other message component/element. For example: Warehouse number.
	SellerDefinedInformation []*UserDefinedInformation1 `xml:"SellrDfndInf,omitempty"`

	// Specifies the applicable Incoterm and associated location.
	Incoterms *Incoterms4 `xml:"Incotrms,omitempty"`
}

func (l *LineItem12) AddPurchaseOrderReference() *DocumentIdentification7 {
	l.PurchaseOrderReference = new(DocumentIdentification7)
	return l.PurchaseOrderReference
}

func (l *LineItem12) SetFinalSubmission(value string) {
	l.FinalSubmission = (*YesNoIndicator)(&value)
}

func (l *LineItem12) AddCommercialLineItems() *LineItemDetails11 {
	newValue := new(LineItemDetails11)
	l.CommercialLineItems = append(l.CommercialLineItems, newValue)
	return newValue
}

func (l *LineItem12) SetLineItemsTotalAmount(value, currency string) {
	l.LineItemsTotalAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem12) AddAdjustment() *Adjustment6 {
	newValue := new(Adjustment6)
	l.Adjustment = append(l.Adjustment, newValue)
	return newValue
}

func (l *LineItem12) AddFreightCharges() *Charge25 {
	l.FreightCharges = new(Charge25)
	return l.FreightCharges
}

func (l *LineItem12) AddTax() *Tax22 {
	newValue := new(Tax22)
	l.Tax = append(l.Tax, newValue)
	return newValue
}

func (l *LineItem12) SetTotalNetAmount(value, currency string) {
	l.TotalNetAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem12) AddBuyerDefinedInformation() *UserDefinedInformation1 {
	newValue := new(UserDefinedInformation1)
	l.BuyerDefinedInformation = append(l.BuyerDefinedInformation, newValue)
	return newValue
}

func (l *LineItem12) AddSellerDefinedInformation() *UserDefinedInformation1 {
	newValue := new(UserDefinedInformation1)
	l.SellerDefinedInformation = append(l.SellerDefinedInformation, newValue)
	return newValue
}

func (l *LineItem12) AddIncoterms() *Incoterms4 {
	l.Incoterms = new(Incoterms4)
	return l.Incoterms
}

// Commercial details of a trade transaction between a buyer and a seller.
type LineItem13 struct {

	// Information about the goods and/or services of a trade transaction.
	GoodsAndOrServicesDescription *Max70Text `xml:"GoodsAndOrSvcsDesc,omitempty"`

	// Indicates whether or not partial shipments are allowed.
	PartialShipment *YesNoIndicator `xml:"PrtlShipmnt"`

	// Indicates whether or not transshipment of goods is allowed.
	TransShipment *YesNoIndicator `xml:"TrnsShipmnt,omitempty"`

	// Specifies an earliest shipment date and a latest shipment date.
	ShipmentDateRange *ShipmentDateRange1 `xml:"ShipmntDtRg,omitempty"`

	// Goods or services that are purchased.
	LineItemDetails []*LineItemDetails13 `xml:"LineItmDtls"`

	// Specifies the total amount of all line items (incl. their adjustments).
	LineItemsTotalAmount *CurrencyAndAmount `xml:"LineItmsTtlAmt"`

	// Information related to the conveyance of goods.
	RoutingSummary *TransportMeans5 `xml:"RtgSummry,omitempty"`

	// Specifies the applicable Incoterms and associated location. Latest version of Incoterms in effect at the date of message creation.
	Incoterms *Incoterms4 `xml:"Incotrms,omitempty"`

	// Variance on price for the goods.
	Adjustment []*Adjustment7 `xml:"Adjstmnt,omitempty"`

	// Maximum charges related to the conveyance of goods.
	FreightCharges *Charge24 `xml:"FrghtChrgs,omitempty"`

	// Amount of money due to the government or tax authority, according to various pre-defined parameters linked to the value of the goods in a trade transaction.
	Tax []*Tax23 `xml:"Tax,omitempty"`

	// Total net amount of a trade transaction. Total amount resulting from the gross amount plus freight charges, tax and plus/minus Adjustments.
	TotalNetAmount *CurrencyAndAmount `xml:"TtlNetAmt"`

	// Information important for the users of the message/service, which cannot be captured in any other message component/element. For example: Warehouse number.
	BuyerDefinedInformation []*UserDefinedInformation1 `xml:"BuyrDfndInf,omitempty"`

	// Information important for the users of the message/service, which cannot be captured in any other message component/element. For example: Warehouse number.
	SellerDefinedInformation []*UserDefinedInformation1 `xml:"SellrDfndInf,omitempty"`
}

func (l *LineItem13) SetGoodsAndOrServicesDescription(value string) {
	l.GoodsAndOrServicesDescription = (*Max70Text)(&value)
}

func (l *LineItem13) SetPartialShipment(value string) {
	l.PartialShipment = (*YesNoIndicator)(&value)
}

func (l *LineItem13) SetTransShipment(value string) {
	l.TransShipment = (*YesNoIndicator)(&value)
}

func (l *LineItem13) AddShipmentDateRange() *ShipmentDateRange1 {
	l.ShipmentDateRange = new(ShipmentDateRange1)
	return l.ShipmentDateRange
}

func (l *LineItem13) AddLineItemDetails() *LineItemDetails13 {
	newValue := new(LineItemDetails13)
	l.LineItemDetails = append(l.LineItemDetails, newValue)
	return newValue
}

func (l *LineItem13) SetLineItemsTotalAmount(value, currency string) {
	l.LineItemsTotalAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem13) AddRoutingSummary() *TransportMeans5 {
	l.RoutingSummary = new(TransportMeans5)
	return l.RoutingSummary
}

func (l *LineItem13) AddIncoterms() *Incoterms4 {
	l.Incoterms = new(Incoterms4)
	return l.Incoterms
}

func (l *LineItem13) AddAdjustment() *Adjustment7 {
	newValue := new(Adjustment7)
	l.Adjustment = append(l.Adjustment, newValue)
	return newValue
}

func (l *LineItem13) AddFreightCharges() *Charge24 {
	l.FreightCharges = new(Charge24)
	return l.FreightCharges
}

func (l *LineItem13) AddTax() *Tax23 {
	newValue := new(Tax23)
	l.Tax = append(l.Tax, newValue)
	return newValue
}

func (l *LineItem13) SetTotalNetAmount(value, currency string) {
	l.TotalNetAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem13) AddBuyerDefinedInformation() *UserDefinedInformation1 {
	newValue := new(UserDefinedInformation1)
	l.BuyerDefinedInformation = append(l.BuyerDefinedInformation, newValue)
	return newValue
}

func (l *LineItem13) AddSellerDefinedInformation() *UserDefinedInformation1 {
	newValue := new(UserDefinedInformation1)
	l.SellerDefinedInformation = append(l.SellerDefinedInformation, newValue)
	return newValue
}

// Calculation of the current situation of a baseline as a result of the submission of a commercial data set.
type LineItem14 struct {

	// Calculated information about the goods of the underlying transaction.
	LineItemDetails []*LineItemDetails12 `xml:"LineItmDtls"`

	// Line items total amount as indicated in the baseline.
	OrderedLineItemsTotalAmount *CurrencyAndAmount `xml:"OrdrdLineItmsTtlAmt"`

	// Line items total amount accepted by a data set submission(s).
	AcceptedLineItemsTotalAmount *CurrencyAndAmount `xml:"AccptdLineItmsTtlAmt"`

	// Difference between the ordered and the accepted line items total amount.
	OutstandingLineItemsTotalAmount *CurrencyAndAmount `xml:"OutsdngLineItmsTtlAmt"`

	// Line item total amount for which a mismatched data set has been submitted and has not yet been accepted or rejected.
	PendingLineItemsTotalAmount *CurrencyAndAmount `xml:"PdgLineItmsTtlAmt"`

	// Total net amount as indicated in the baseline.
	OrderedTotalNetAmount *CurrencyAndAmount `xml:"OrdrdTtlNetAmt"`

	// Total net amount accepted by a data set submission.
	AcceptedTotalNetAmount *CurrencyAndAmount `xml:"AccptdTtlNetAmt"`

	// Total net amount for which a mismatched data set has been submitted and has not yet been accepted or rejected.
	OutstandingTotalNetAmount *CurrencyAndAmount `xml:"OutsdngTtlNetAmt"`

	// Difference between the ordered and the accepted total net amount.
	PendingTotalNetAmount *CurrencyAndAmount `xml:"PdgTtlNetAmt"`
}

func (l *LineItem14) AddLineItemDetails() *LineItemDetails12 {
	newValue := new(LineItemDetails12)
	l.LineItemDetails = append(l.LineItemDetails, newValue)
	return newValue
}

func (l *LineItem14) SetOrderedLineItemsTotalAmount(value, currency string) {
	l.OrderedLineItemsTotalAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem14) SetAcceptedLineItemsTotalAmount(value, currency string) {
	l.AcceptedLineItemsTotalAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem14) SetOutstandingLineItemsTotalAmount(value, currency string) {
	l.OutstandingLineItemsTotalAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem14) SetPendingLineItemsTotalAmount(value, currency string) {
	l.PendingLineItemsTotalAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem14) SetOrderedTotalNetAmount(value, currency string) {
	l.OrderedTotalNetAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem14) SetAcceptedTotalNetAmount(value, currency string) {
	l.AcceptedTotalNetAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem14) SetOutstandingTotalNetAmount(value, currency string) {
	l.OutstandingTotalNetAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem14) SetPendingTotalNetAmount(value, currency string) {
	l.PendingTotalNetAmount = NewCurrencyAndAmount(value, currency)
}

// Commercial details of a trade transaction between a buyer and a seller.
type LineItem15 struct {

	// Reference to the purchase order of the underlying transaction.
	PurchaseOrderReference *DocumentIdentification7 `xml:"PurchsOrdrRef"`

	// Specifies whether this current invoice is the last submission against the purchase order referenced.
	FinalSubmission *YesNoIndicator `xml:"FnlSubmissn"`

	// Goods which are the subject of the invoice.
	CommercialLineItems []*LineItemDetails14 `xml:"ComrclLineItms"`

	// Specifies the total amount of all line items (incl. their adjustments).
	LineItemsTotalAmount *CurrencyAndAmount `xml:"LineItmsTtlAmt"`

	// Variance on price for the goods.
	Adjustment []*Adjustment6 `xml:"Adjstmnt,omitempty"`

	// Charges related to the conveyance of goods.
	FreightCharges *Charge25 `xml:"FrghtChrgs,omitempty"`

	// Amount of money due to the government or tax authority, according to various pre-defined parameters linked to the value of the goods in a trade transaction.
	Tax []*Tax22 `xml:"Tax,omitempty"`

	// Total net amount of a trade transaction. Total amount resulting from the gross amount plus freight charges, tax and plus/minus Adjustments.
	TotalNetAmount *CurrencyAndAmount `xml:"TtlNetAmt"`

	// Information important for the users of the message/service, which cannot be captured in any other message component/element. For example: Warehouse number.
	BuyerDefinedInformation []*UserDefinedInformation1 `xml:"BuyrDfndInf,omitempty"`

	// Information important for the users of the message/service, which cannot be captured in any other message component/element. For example: Warehouse number.
	SellerDefinedInformation []*UserDefinedInformation1 `xml:"SellrDfndInf,omitempty"`

	// Specifies the applicable Incoterm and associated location.
	Incoterms *Incoterms4 `xml:"Incotrms,omitempty"`
}

func (l *LineItem15) AddPurchaseOrderReference() *DocumentIdentification7 {
	l.PurchaseOrderReference = new(DocumentIdentification7)
	return l.PurchaseOrderReference
}

func (l *LineItem15) SetFinalSubmission(value string) {
	l.FinalSubmission = (*YesNoIndicator)(&value)
}

func (l *LineItem15) AddCommercialLineItems() *LineItemDetails14 {
	newValue := new(LineItemDetails14)
	l.CommercialLineItems = append(l.CommercialLineItems, newValue)
	return newValue
}

func (l *LineItem15) SetLineItemsTotalAmount(value, currency string) {
	l.LineItemsTotalAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem15) AddAdjustment() *Adjustment6 {
	newValue := new(Adjustment6)
	l.Adjustment = append(l.Adjustment, newValue)
	return newValue
}

func (l *LineItem15) AddFreightCharges() *Charge25 {
	l.FreightCharges = new(Charge25)
	return l.FreightCharges
}

func (l *LineItem15) AddTax() *Tax22 {
	newValue := new(Tax22)
	l.Tax = append(l.Tax, newValue)
	return newValue
}

func (l *LineItem15) SetTotalNetAmount(value, currency string) {
	l.TotalNetAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem15) AddBuyerDefinedInformation() *UserDefinedInformation1 {
	newValue := new(UserDefinedInformation1)
	l.BuyerDefinedInformation = append(l.BuyerDefinedInformation, newValue)
	return newValue
}

func (l *LineItem15) AddSellerDefinedInformation() *UserDefinedInformation1 {
	newValue := new(UserDefinedInformation1)
	l.SellerDefinedInformation = append(l.SellerDefinedInformation, newValue)
	return newValue
}

func (l *LineItem15) AddIncoterms() *Incoterms4 {
	l.Incoterms = new(Incoterms4)
	return l.Incoterms
}

// Unit of information  showing the related  provision of products and/or services and monetary summations reported as a discrete line items.
type LineItem16 struct {

	// The unique identification of this invoice line item.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Something that is produced and sold as the result of an industrial process.
	TradeProduct *TradeProduct2 `xml:"TradPdct,omitempty"`

	// Purchase order reference assigned by the buyer related to the provision of products and/or services for this line item.
	BuyerOrderIdentification *DocumentIdentification23 `xml:"BuyrOrdrId,omitempty"`

	// Contract reference related to the provision of products and/or services for this line item.
	ContractIdentification *DocumentIdentification22 `xml:"CtrctId,omitempty"`

	// Specific purchase account for recording debits and credits for accounting purposes.
	PurchaseAccountingAccount []*AccountingAccount1 `xml:"PurchsAcctgAcct,omitempty"`

	// Value of the price, eg, as a currency and value.
	NetPrice []*CurrencyAndAmount `xml:"NetPric,omitempty"`

	// Quantity and conversion factor on which the net price is based for this line item product and/or service.
	NetPriceQuantity *Quantity9 `xml:"NetPricQty,omitempty"`

	// Allowance or charge applied to the net price.
	NetPriceAllowanceCharge []*LineItemAllowanceCharge2 `xml:"NetPricAllwncChrg,omitempty"`

	// Net weight of the product.
	NetWeight *Quantity10 `xml:"NetWght,omitempty"`

	// Gross price of the product and/or service.
	GrossPrice []*CurrencyAndAmount `xml:"GrssPric,omitempty"`

	// Quantity and conversion factor on which the gross price is based for this line item product and/or service.
	GrossPriceQuantity *Quantity9 `xml:"GrssPricQty,omitempty"`

	// Gross weight of the product.
	GrossWeight *Quantity10 `xml:"GrssWght,omitempty"`

	// Logistics service charge for this line item.
	LogisticsCharge []*ChargesDetails4 `xml:"LogstcsChrg,omitempty"`

	// Amount of money due to the government or tax authority, according to various pre-defined parameters such as thresholds or income.
	Tax []*LineItemTax1 `xml:"Tax,omitempty"`

	// Allowance or charge specified for this line item.
	AllowanceCharge []*LineItemAllowanceCharge2 `xml:"AllwncChrg,omitempty"`

	// Modification on the value of goods and / or services. For example: rebate, discount, surcharge
	FinancialAdjustment []*Adjustment6 `xml:"FinAdjstmnt,omitempty"`

	// Quantity billed for this line item.
	BilledQuantity *Quantity10 `xml:"BlldQty,omitempty"`

	// Number of product packages delivered.
	PackageQuantity *DecimalNumber `xml:"PackgQty,omitempty"`

	// Number of units per package in this line item for a supply chain trade delivery.
	PerPackageUnitQuantity *Quantity10 `xml:"PerPackgUnitQty,omitempty"`

	// Physical packaging of the product.
	Packaging []*Packaging1 `xml:"Packgng,omitempty"`

	// Quantity that is free of charge for this line item.
	ChargeFreeQuantity *Quantity10 `xml:"ChrgFreeQty,omitempty"`

	// Quantity value on which the quantity measurement started for a line item. For instance the start amount of a meter reading for an electricity supplier.
	MeasureQuantityStart *Quantity10 `xml:"MeasrQtyStart,omitempty"`

	// Quantity value on which the quantity measurement ended for a line item. For instance the end amount of a meter reading for an electricity supplier.
	MeasureQuantityEnd *Quantity10 `xml:"MeasrQtyEnd,omitempty"`

	// Date/time on which the clock time measure started for a line item.
	MeasureDateTimeStart *ISODateTime `xml:"MeasrDtTmStart,omitempty"`

	// Date/time on which the clock time measure ended for a line item.
	MeasureDateTimeEnd *ISODateTime `xml:"MeasrDtTmEnd,omitempty"`

	// Party to whom the goods must be delivered in the end.
	ShipTo *TradeParty3 `xml:"ShipTo,omitempty"`

	// Specifies the applicable Incoterm and associated location.
	Incoterms *Incoterms3 `xml:"Incotrms,omitempty"`

	// Actual delivery date/time of the products and/or services for this line item.
	DeliveryDateTime *ISODateTime `xml:"DlvryDtTm,omitempty"`

	// Delivery note related to the delivery of the products and/or services for this line item.
	DeliveryNoteIdentification *DocumentIdentification22 `xml:"DlvryNoteId,omitempty"`

	// Monetary totals for this line item.
	MonetarySummation *LineItemMonetarySummation1 `xml:"MntrySummtn,omitempty"`

	// Note included in this line item.
	IncludedNote []*AdditionalInformation1 `xml:"InclNote,omitempty"`
}

func (l *LineItem16) SetIdentification(value string) {
	l.Identification = (*Max35Text)(&value)
}

func (l *LineItem16) AddTradeProduct() *TradeProduct2 {
	l.TradeProduct = new(TradeProduct2)
	return l.TradeProduct
}

func (l *LineItem16) AddBuyerOrderIdentification() *DocumentIdentification23 {
	l.BuyerOrderIdentification = new(DocumentIdentification23)
	return l.BuyerOrderIdentification
}

func (l *LineItem16) AddContractIdentification() *DocumentIdentification22 {
	l.ContractIdentification = new(DocumentIdentification22)
	return l.ContractIdentification
}

func (l *LineItem16) AddPurchaseAccountingAccount() *AccountingAccount1 {
	newValue := new(AccountingAccount1)
	l.PurchaseAccountingAccount = append(l.PurchaseAccountingAccount, newValue)
	return newValue
}

func (l *LineItem16) AddNetPrice(value, currency string) {
	l.NetPrice = append(l.NetPrice, NewCurrencyAndAmount(value, currency))
}

func (l *LineItem16) AddNetPriceQuantity() *Quantity9 {
	l.NetPriceQuantity = new(Quantity9)
	return l.NetPriceQuantity
}

func (l *LineItem16) AddNetPriceAllowanceCharge() *LineItemAllowanceCharge2 {
	newValue := new(LineItemAllowanceCharge2)
	l.NetPriceAllowanceCharge = append(l.NetPriceAllowanceCharge, newValue)
	return newValue
}

func (l *LineItem16) AddNetWeight() *Quantity10 {
	l.NetWeight = new(Quantity10)
	return l.NetWeight
}

func (l *LineItem16) AddGrossPrice(value, currency string) {
	l.GrossPrice = append(l.GrossPrice, NewCurrencyAndAmount(value, currency))
}

func (l *LineItem16) AddGrossPriceQuantity() *Quantity9 {
	l.GrossPriceQuantity = new(Quantity9)
	return l.GrossPriceQuantity
}

func (l *LineItem16) AddGrossWeight() *Quantity10 {
	l.GrossWeight = new(Quantity10)
	return l.GrossWeight
}

func (l *LineItem16) AddLogisticsCharge() *ChargesDetails4 {
	newValue := new(ChargesDetails4)
	l.LogisticsCharge = append(l.LogisticsCharge, newValue)
	return newValue
}

func (l *LineItem16) AddTax() *LineItemTax1 {
	newValue := new(LineItemTax1)
	l.Tax = append(l.Tax, newValue)
	return newValue
}

func (l *LineItem16) AddAllowanceCharge() *LineItemAllowanceCharge2 {
	newValue := new(LineItemAllowanceCharge2)
	l.AllowanceCharge = append(l.AllowanceCharge, newValue)
	return newValue
}

func (l *LineItem16) AddFinancialAdjustment() *Adjustment6 {
	newValue := new(Adjustment6)
	l.FinancialAdjustment = append(l.FinancialAdjustment, newValue)
	return newValue
}

func (l *LineItem16) AddBilledQuantity() *Quantity10 {
	l.BilledQuantity = new(Quantity10)
	return l.BilledQuantity
}

func (l *LineItem16) SetPackageQuantity(value string) {
	l.PackageQuantity = (*DecimalNumber)(&value)
}

func (l *LineItem16) AddPerPackageUnitQuantity() *Quantity10 {
	l.PerPackageUnitQuantity = new(Quantity10)
	return l.PerPackageUnitQuantity
}

func (l *LineItem16) AddPackaging() *Packaging1 {
	newValue := new(Packaging1)
	l.Packaging = append(l.Packaging, newValue)
	return newValue
}

func (l *LineItem16) AddChargeFreeQuantity() *Quantity10 {
	l.ChargeFreeQuantity = new(Quantity10)
	return l.ChargeFreeQuantity
}

func (l *LineItem16) AddMeasureQuantityStart() *Quantity10 {
	l.MeasureQuantityStart = new(Quantity10)
	return l.MeasureQuantityStart
}

func (l *LineItem16) AddMeasureQuantityEnd() *Quantity10 {
	l.MeasureQuantityEnd = new(Quantity10)
	return l.MeasureQuantityEnd
}

func (l *LineItem16) SetMeasureDateTimeStart(value string) {
	l.MeasureDateTimeStart = (*ISODateTime)(&value)
}

func (l *LineItem16) SetMeasureDateTimeEnd(value string) {
	l.MeasureDateTimeEnd = (*ISODateTime)(&value)
}

func (l *LineItem16) AddShipTo() *TradeParty3 {
	l.ShipTo = new(TradeParty3)
	return l.ShipTo
}

func (l *LineItem16) AddIncoterms() *Incoterms3 {
	l.Incoterms = new(Incoterms3)
	return l.Incoterms
}

func (l *LineItem16) SetDeliveryDateTime(value string) {
	l.DeliveryDateTime = (*ISODateTime)(&value)
}

func (l *LineItem16) AddDeliveryNoteIdentification() *DocumentIdentification22 {
	l.DeliveryNoteIdentification = new(DocumentIdentification22)
	return l.DeliveryNoteIdentification
}

func (l *LineItem16) AddMonetarySummation() *LineItemMonetarySummation1 {
	l.MonetarySummation = new(LineItemMonetarySummation1)
	return l.MonetarySummation
}

func (l *LineItem16) AddIncludedNote() *AdditionalInformation1 {
	newValue := new(AdditionalInformation1)
	l.IncludedNote = append(l.IncludedNote, newValue)
	return newValue
}

// Commercial details of a trade transaction between a buyer and a seller.
type LineItem7 struct {

	// Information about the goods and/or services of a trade transaction.
	GoodsDescription *Max70Text `xml:"GoodsDesc,omitempty"`

	// Indicates whether or not partial shipments are allowed.
	PartialShipment *YesNoIndicator `xml:"PrtlShipmnt"`

	// Indicates whether or not transshipment of goods is allowed.
	TransShipment *YesNoIndicator `xml:"TrnsShipmnt,omitempty"`

	// Specifies an earliest shipment date and a latest shipment date.
	ShipmentDateRange *ShipmentDateRange1 `xml:"ShipmntDtRg,omitempty"`

	// Goods or services that are purchased.
	LineItemDetails []*LineItemDetails7 `xml:"LineItmDtls"`

	// Specifies the total amount of all line items (incl. their adjustments).
	LineItemsTotalAmount *CurrencyAndAmount `xml:"LineItmsTtlAmt"`

	// Information related to the conveyance of goods.
	RoutingSummary *TransportMeans1 `xml:"RtgSummry,omitempty"`

	// Specifies the applicable Incoterms and associated location. Latest version of Incoterms in effect at the date of message creation.
	Incoterms []*Incoterms1 `xml:"Incotrms,omitempty"`

	// Variance on price for the goods.
	Adjustment []*Adjustment3 `xml:"Adjstmnt,omitempty"`

	// Maximum charges related to the conveyance of goods.
	FreightCharges *Charge12 `xml:"FrghtChrgs,omitempty"`

	// Amount of money due to the government or tax authority, according to various pre-defined parameters linked to the value of the goods in a trade transaction.
	Tax []*Tax13 `xml:"Tax,omitempty"`

	// Total net amount of a trade transaction. Total amount resulting from the gross amount plus freight charges, tax and plus/minus Adjustments.
	TotalNetAmount *CurrencyAndAmount `xml:"TtlNetAmt"`

	// Information important for the users of the message/service, which cannot be captured in any other message component/element. For example: Warehouse number.
	BuyerDefinedInformation []*UserDefinedInformation1 `xml:"BuyrDfndInf,omitempty"`

	// Information important for the users of the message/service, which cannot be captured in any other message component/element. For example: Warehouse number.
	SellerDefinedInformation []*UserDefinedInformation1 `xml:"SellrDfndInf,omitempty"`
}

func (l *LineItem7) SetGoodsDescription(value string) {
	l.GoodsDescription = (*Max70Text)(&value)
}

func (l *LineItem7) SetPartialShipment(value string) {
	l.PartialShipment = (*YesNoIndicator)(&value)
}

func (l *LineItem7) SetTransShipment(value string) {
	l.TransShipment = (*YesNoIndicator)(&value)
}

func (l *LineItem7) AddShipmentDateRange() *ShipmentDateRange1 {
	l.ShipmentDateRange = new(ShipmentDateRange1)
	return l.ShipmentDateRange
}

func (l *LineItem7) AddLineItemDetails() *LineItemDetails7 {
	newValue := new(LineItemDetails7)
	l.LineItemDetails = append(l.LineItemDetails, newValue)
	return newValue
}

func (l *LineItem7) SetLineItemsTotalAmount(value, currency string) {
	l.LineItemsTotalAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem7) AddRoutingSummary() *TransportMeans1 {
	l.RoutingSummary = new(TransportMeans1)
	return l.RoutingSummary
}

func (l *LineItem7) AddIncoterms() *Incoterms1 {
	newValue := new(Incoterms1)
	l.Incoterms = append(l.Incoterms, newValue)
	return newValue
}

func (l *LineItem7) AddAdjustment() *Adjustment3 {
	newValue := new(Adjustment3)
	l.Adjustment = append(l.Adjustment, newValue)
	return newValue
}

func (l *LineItem7) AddFreightCharges() *Charge12 {
	l.FreightCharges = new(Charge12)
	return l.FreightCharges
}

func (l *LineItem7) AddTax() *Tax13 {
	newValue := new(Tax13)
	l.Tax = append(l.Tax, newValue)
	return newValue
}

func (l *LineItem7) SetTotalNetAmount(value, currency string) {
	l.TotalNetAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem7) AddBuyerDefinedInformation() *UserDefinedInformation1 {
	newValue := new(UserDefinedInformation1)
	l.BuyerDefinedInformation = append(l.BuyerDefinedInformation, newValue)
	return newValue
}

func (l *LineItem7) AddSellerDefinedInformation() *UserDefinedInformation1 {
	newValue := new(UserDefinedInformation1)
	l.SellerDefinedInformation = append(l.SellerDefinedInformation, newValue)
	return newValue
}

// Calculation of the current situation of a baseline as a result of the submission of a commercial data set.
type LineItem8 struct {

	// Calculated information about the goods of the underlying transaction.
	LineItemDetails []*LineItemDetails8 `xml:"LineItmDtls"`

	// Line items total amount as indicated in the baseline.
	OrderedLineItemsTotalAmount *CurrencyAndAmount `xml:"OrdrdLineItmsTtlAmt"`

	// Line items total amount accepted by a data set submission(s).
	AcceptedLineItemsTotalAmount *CurrencyAndAmount `xml:"AccptdLineItmsTtlAmt"`

	// Difference between the ordered and the accepted line items total amount.
	OutstandingLineItemsTotalAmount *CurrencyAndAmount `xml:"OutsdngLineItmsTtlAmt"`

	// Line item total amount for which a mismatched data set has been submitted and has not yet been accepted or rejected.
	PendingLineItemsTotalAmount *CurrencyAndAmount `xml:"PdgLineItmsTtlAmt"`

	// Total net amount as indicated in the baseline.
	OrderedTotalNetAmount *CurrencyAndAmount `xml:"OrdrdTtlNetAmt"`

	// Total net amount accepted by a data set submission.
	AcceptedTotalNetAmount *CurrencyAndAmount `xml:"AccptdTtlNetAmt"`

	// Total net amount for which a mismatched data set has been submitted and has not yet been accepted or rejected.
	OutstandingTotalNetAmount *CurrencyAndAmount `xml:"OutsdngTtlNetAmt"`

	// Difference between the ordered and the accepted total net amount.
	PendingTotalNetAmount *CurrencyAndAmount `xml:"PdgTtlNetAmt"`
}

func (l *LineItem8) AddLineItemDetails() *LineItemDetails8 {
	newValue := new(LineItemDetails8)
	l.LineItemDetails = append(l.LineItemDetails, newValue)
	return newValue
}

func (l *LineItem8) SetOrderedLineItemsTotalAmount(value, currency string) {
	l.OrderedLineItemsTotalAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem8) SetAcceptedLineItemsTotalAmount(value, currency string) {
	l.AcceptedLineItemsTotalAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem8) SetOutstandingLineItemsTotalAmount(value, currency string) {
	l.OutstandingLineItemsTotalAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem8) SetPendingLineItemsTotalAmount(value, currency string) {
	l.PendingLineItemsTotalAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem8) SetOrderedTotalNetAmount(value, currency string) {
	l.OrderedTotalNetAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem8) SetAcceptedTotalNetAmount(value, currency string) {
	l.AcceptedTotalNetAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem8) SetOutstandingTotalNetAmount(value, currency string) {
	l.OutstandingTotalNetAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem8) SetPendingTotalNetAmount(value, currency string) {
	l.PendingTotalNetAmount = NewCurrencyAndAmount(value, currency)
}

// Commercial details of a trade transaction between a buyer and a seller.
type LineItem9 struct {

	// Reference to the purchase order of the underlying transaction.
	PurchaseOrderReference *DocumentIdentification7 `xml:"PurchsOrdrRef"`

	// Specifies whether this current invoice is the last submission against the purchase order referenced.
	FinalSubmission *YesNoIndicator `xml:"FnlSubmissn"`

	// Goods which are the subject of the invoice.
	CommercialLineItems []*LineItemDetails9 `xml:"ComrclLineItms"`

	// Specifies the total amount of all line items (incl. their adjustments).
	LineItemsTotalAmount *CurrencyAndAmount `xml:"LineItmsTtlAmt"`

	// Specifies the applicable Incoterm and associated location.
	Incoterms *Incoterms2 `xml:"Incotrms,omitempty"`

	// Variance on price for the goods.
	Adjustment []*Adjustment4 `xml:"Adjstmnt,omitempty"`

	// Charges related to the conveyance of goods.
	FreightCharges *Charge13 `xml:"FrghtChrgs,omitempty"`

	// Amount of money due to the government or tax authority, according to various pre-defined parameters linked to the value of the goods in a trade transaction.
	Tax []*Tax12 `xml:"Tax,omitempty"`

	// Total net amount of a trade transaction. Total amount resulting from the gross amount plus freight charges, tax and plus/minus Adjustments.
	TotalNetAmount *CurrencyAndAmount `xml:"TtlNetAmt"`

	// Information important for the users of the message/service, which cannot be captured in any other message component/element. For example: Warehouse number.
	BuyerDefinedInformation []*UserDefinedInformation1 `xml:"BuyrDfndInf,omitempty"`

	// Information important for the users of the message/service, which cannot be captured in any other message component/element. For example: Warehouse number.
	SellerDefinedInformation []*UserDefinedInformation1 `xml:"SellrDfndInf,omitempty"`
}

func (l *LineItem9) AddPurchaseOrderReference() *DocumentIdentification7 {
	l.PurchaseOrderReference = new(DocumentIdentification7)
	return l.PurchaseOrderReference
}

func (l *LineItem9) SetFinalSubmission(value string) {
	l.FinalSubmission = (*YesNoIndicator)(&value)
}

func (l *LineItem9) AddCommercialLineItems() *LineItemDetails9 {
	newValue := new(LineItemDetails9)
	l.CommercialLineItems = append(l.CommercialLineItems, newValue)
	return newValue
}

func (l *LineItem9) SetLineItemsTotalAmount(value, currency string) {
	l.LineItemsTotalAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem9) AddIncoterms() *Incoterms2 {
	l.Incoterms = new(Incoterms2)
	return l.Incoterms
}

func (l *LineItem9) AddAdjustment() *Adjustment4 {
	newValue := new(Adjustment4)
	l.Adjustment = append(l.Adjustment, newValue)
	return newValue
}

func (l *LineItem9) AddFreightCharges() *Charge13 {
	l.FreightCharges = new(Charge13)
	return l.FreightCharges
}

func (l *LineItem9) AddTax() *Tax12 {
	newValue := new(Tax12)
	l.Tax = append(l.Tax, newValue)
	return newValue
}

func (l *LineItem9) SetTotalNetAmount(value, currency string) {
	l.TotalNetAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItem9) AddBuyerDefinedInformation() *UserDefinedInformation1 {
	newValue := new(UserDefinedInformation1)
	l.BuyerDefinedInformation = append(l.BuyerDefinedInformation, newValue)
	return newValue
}

func (l *LineItem9) AddSellerDefinedInformation() *UserDefinedInformation1 {
	newValue := new(UserDefinedInformation1)
	l.SellerDefinedInformation = append(l.SellerDefinedInformation, newValue)
	return newValue
}
