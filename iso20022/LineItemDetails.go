package iso20022

// Goods or services that are part of a commercial trade agreement.
type LineItemDetails10 struct {

	// Identification assigned to a line item.
	LineItemIdentification *Max70Text `xml:"LineItmId"`

	// Specifies the quantity of goods on a line in a trade transaction.
	Quantity *Quantity9 `xml:"Qty"`

	// Variance allowed in the quantity of goods.
	QuantityTolerance *PercentageTolerance1 `xml:"QtyTlrnce,omitempty"`

	// Amount of money for which goods or services are offered, sold, or bought.
	UnitPrice *UnitPrice18 `xml:"UnitPric,omitempty"`

	// Variance allowed on a price.
	PriceTolerance *PercentageTolerance1 `xml:"PricTlrnce,omitempty"`

	// Name of the product detailed in the corresponding line item.
	ProductName *Max70Text `xml:"PdctNm,omitempty"`

	// Identifies the product of the corresponding line item.
	ProductIdentifier []*ProductIdentifier2Choice `xml:"PdctIdr,omitempty"`

	// Identifies the characteristics of a product.
	ProductCharacteristics []*ProductCharacteristics1Choice `xml:"PdctChrtcs,omitempty"`

	// Identifies the category of product.
	ProductCategory []*ProductCategory1Choice `xml:"PdctCtgy,omitempty"`

	// Country from which the product originates.
	ProductOrigin []*CountryCode `xml:"PdctOrgn,omitempty"`

	// Specifies the shipment schedule for the goods.
	ShipmentSchedule *ShipmentSchedule2Choice `xml:"ShipmntSchdl,omitempty"`

	// Information related to the conveyance of goods.
	RoutingSummary *TransportMeans5 `xml:"RtgSummry,omitempty"`

	// Variance on price for the goods.
	Adjustment []*Adjustment7 `xml:"Adjstmnt,omitempty"`

	// Maximum charges related to the conveyance of goods.
	FreightCharges *Charge24 `xml:"FrghtChrgs,omitempty"`

	// Amount of money due to the government or tax authority, according to various pre-defined parameters linked to the value of the goods in a trade transaction.
	Tax []*Tax23 `xml:"Tax,omitempty"`

	// Total amount of the line item after adjustments have been applied.
	TotalAmount *CurrencyAndAmount `xml:"TtlAmt"`

	// Specifies the applicable Incoterms and associated location. Latest version of Incoterms in effect at the date of message creation.
	Incoterms *Incoterms4 `xml:"Incotrms,omitempty"`
}

func (l *LineItemDetails10) SetLineItemIdentification(value string) {
	l.LineItemIdentification = (*Max70Text)(&value)
}

func (l *LineItemDetails10) AddQuantity() *Quantity9 {
	l.Quantity = new(Quantity9)
	return l.Quantity
}

func (l *LineItemDetails10) AddQuantityTolerance() *PercentageTolerance1 {
	l.QuantityTolerance = new(PercentageTolerance1)
	return l.QuantityTolerance
}

func (l *LineItemDetails10) AddUnitPrice() *UnitPrice18 {
	l.UnitPrice = new(UnitPrice18)
	return l.UnitPrice
}

func (l *LineItemDetails10) AddPriceTolerance() *PercentageTolerance1 {
	l.PriceTolerance = new(PercentageTolerance1)
	return l.PriceTolerance
}

func (l *LineItemDetails10) SetProductName(value string) {
	l.ProductName = (*Max70Text)(&value)
}

func (l *LineItemDetails10) AddProductIdentifier() *ProductIdentifier2Choice {
	newValue := new(ProductIdentifier2Choice)
	l.ProductIdentifier = append(l.ProductIdentifier, newValue)
	return newValue
}

func (l *LineItemDetails10) AddProductCharacteristics() *ProductCharacteristics1Choice {
	newValue := new(ProductCharacteristics1Choice)
	l.ProductCharacteristics = append(l.ProductCharacteristics, newValue)
	return newValue
}

func (l *LineItemDetails10) AddProductCategory() *ProductCategory1Choice {
	newValue := new(ProductCategory1Choice)
	l.ProductCategory = append(l.ProductCategory, newValue)
	return newValue
}

func (l *LineItemDetails10) AddProductOrigin(value string) {
	l.ProductOrigin = append(l.ProductOrigin, (*CountryCode)(&value))
}

func (l *LineItemDetails10) AddShipmentSchedule() *ShipmentSchedule2Choice {
	l.ShipmentSchedule = new(ShipmentSchedule2Choice)
	return l.ShipmentSchedule
}

func (l *LineItemDetails10) AddRoutingSummary() *TransportMeans5 {
	l.RoutingSummary = new(TransportMeans5)
	return l.RoutingSummary
}

func (l *LineItemDetails10) AddAdjustment() *Adjustment7 {
	newValue := new(Adjustment7)
	l.Adjustment = append(l.Adjustment, newValue)
	return newValue
}

func (l *LineItemDetails10) AddFreightCharges() *Charge24 {
	l.FreightCharges = new(Charge24)
	return l.FreightCharges
}

func (l *LineItemDetails10) AddTax() *Tax23 {
	newValue := new(Tax23)
	l.Tax = append(l.Tax, newValue)
	return newValue
}

func (l *LineItemDetails10) SetTotalAmount(value, currency string) {
	l.TotalAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItemDetails10) AddIncoterms() *Incoterms4 {
	l.Incoterms = new(Incoterms4)
	return l.Incoterms
}

// Goods or services that are part of a commercial trade agreement.
type LineItemDetails11 struct {

	// Sequential number assigned to a line item.
	LineItemIdentification *Max70Text `xml:"LineItmId"`

	// Specifies the quantity of a product in a trade transaction.
	Quantity *Quantity9 `xml:"Qty"`

	// Amount of money for which goods or services are offered, sold, or bought.
	UnitPrice *UnitPrice18 `xml:"UnitPric,omitempty"`

	// Name of the product detailed in the corresponding line item.
	ProductName *Max70Text `xml:"PdctNm,omitempty"`

	// Identifies the product of the corresponding line item.
	ProductIdentifier []*ProductIdentifier2Choice `xml:"PdctIdr,omitempty"`

	// Identifies the characteristics of product.
	ProductCharacteristics []*ProductCharacteristics1Choice `xml:"PdctChrtcs,omitempty"`

	// Identifies the category of product.
	ProductCategory []*ProductCategory1Choice `xml:"PdctCtgy,omitempty"`

	// Country of origin of the goods.
	ProductOrigin *CountryCode `xml:"PdctOrgn,omitempty"`

	// Variance on price for the goods.
	Adjustment []*Adjustment6 `xml:"Adjstmnt,omitempty"`

	// Charges related to the conveyance of goods.
	FreightCharges *Charge25 `xml:"FrghtChrgs,omitempty"`

	// Amount of money due to the government or tax authority, according to various pre-defined parameters linked to the value of the goods in a trade transaction.
	Tax []*Tax22 `xml:"Tax,omitempty"`

	// Total amount of the line item after adjustments have been applied.
	TotalAmount *CurrencyAndAmount `xml:"TtlAmt"`
}

func (l *LineItemDetails11) SetLineItemIdentification(value string) {
	l.LineItemIdentification = (*Max70Text)(&value)
}

func (l *LineItemDetails11) AddQuantity() *Quantity9 {
	l.Quantity = new(Quantity9)
	return l.Quantity
}

func (l *LineItemDetails11) AddUnitPrice() *UnitPrice18 {
	l.UnitPrice = new(UnitPrice18)
	return l.UnitPrice
}

func (l *LineItemDetails11) SetProductName(value string) {
	l.ProductName = (*Max70Text)(&value)
}

func (l *LineItemDetails11) AddProductIdentifier() *ProductIdentifier2Choice {
	newValue := new(ProductIdentifier2Choice)
	l.ProductIdentifier = append(l.ProductIdentifier, newValue)
	return newValue
}

func (l *LineItemDetails11) AddProductCharacteristics() *ProductCharacteristics1Choice {
	newValue := new(ProductCharacteristics1Choice)
	l.ProductCharacteristics = append(l.ProductCharacteristics, newValue)
	return newValue
}

func (l *LineItemDetails11) AddProductCategory() *ProductCategory1Choice {
	newValue := new(ProductCategory1Choice)
	l.ProductCategory = append(l.ProductCategory, newValue)
	return newValue
}

func (l *LineItemDetails11) SetProductOrigin(value string) {
	l.ProductOrigin = (*CountryCode)(&value)
}

func (l *LineItemDetails11) AddAdjustment() *Adjustment6 {
	newValue := new(Adjustment6)
	l.Adjustment = append(l.Adjustment, newValue)
	return newValue
}

func (l *LineItemDetails11) AddFreightCharges() *Charge25 {
	l.FreightCharges = new(Charge25)
	return l.FreightCharges
}

func (l *LineItemDetails11) AddTax() *Tax22 {
	newValue := new(Tax22)
	l.Tax = append(l.Tax, newValue)
	return newValue
}

func (l *LineItemDetails11) SetTotalAmount(value, currency string) {
	l.TotalAmount = NewCurrencyAndAmount(value, currency)
}

// Calculation of the current situation of a line item as a result of the submission of a commercial dataset.
type LineItemDetails12 struct {

	// Sequential number assigned to a line item.
	LineItemIdentification *Max70Text `xml:"LineItmId"`

	// Name of the product detailed in the corresponding line item.
	ProductName *Max70Text `xml:"PdctNm,omitempty"`

	// Identifies the product of the corresponding line item.
	ProductIdentifier []*ProductIdentifier2Choice `xml:"PdctIdr,omitempty"`

	// Identifies the characteristic of a product.
	ProductCharacteristics []*ProductCharacteristics1Choice `xml:"PdctChrtcs,omitempty"`

	// Identifies the category of product.
	ProductCategory []*ProductCategory1Choice `xml:"PdctCtgy,omitempty"`

	// Quantity ordered for a line as indicated in the baseline.
	OrderedQuantity *Quantity9 `xml:"OrdrdQty"`

	// Quantity accepted by data set submission.
	AcceptedQuantity *Quantity9 `xml:"AccptdQty"`

	// Difference between the ordered quantity and the accepted one.
	OutstandingQuantity *Quantity9 `xml:"OutsdngQty"`

	// Quantity of a product for which a mismatched data set has been submitted and has not yet been accepted or rejected..
	PendingQuantity *Quantity9 `xml:"PdgQty"`

	// Variance allowed on the quantity of goods.
	QuantityTolerance *PercentageTolerance1 `xml:"QtyTlrnce,omitempty"`

	// Price multiplied by the ordered quantity for a line as indicated in the baseline.
	OrderedAmount *CurrencyAndAmount `xml:"OrdrdAmt"`

	// Price multiplied by the accepted quantity for a line.
	AcceptedAmount *CurrencyAndAmount `xml:"AccptdAmt"`

	// Price multiplied by the outstanding quantity for a line.
	OutstandingAmount *CurrencyAndAmount `xml:"OutsdngAmt"`

	// Price multiplied by the pending quantity for a line.
	PendingAmount *CurrencyAndAmount `xml:"PdgAmt"`

	// Variance on price for the goods.
	PriceTolerance *PercentageTolerance1 `xml:"PricTlrnce,omitempty"`
}

func (l *LineItemDetails12) SetLineItemIdentification(value string) {
	l.LineItemIdentification = (*Max70Text)(&value)
}

func (l *LineItemDetails12) SetProductName(value string) {
	l.ProductName = (*Max70Text)(&value)
}

func (l *LineItemDetails12) AddProductIdentifier() *ProductIdentifier2Choice {
	newValue := new(ProductIdentifier2Choice)
	l.ProductIdentifier = append(l.ProductIdentifier, newValue)
	return newValue
}

func (l *LineItemDetails12) AddProductCharacteristics() *ProductCharacteristics1Choice {
	newValue := new(ProductCharacteristics1Choice)
	l.ProductCharacteristics = append(l.ProductCharacteristics, newValue)
	return newValue
}

func (l *LineItemDetails12) AddProductCategory() *ProductCategory1Choice {
	newValue := new(ProductCategory1Choice)
	l.ProductCategory = append(l.ProductCategory, newValue)
	return newValue
}

func (l *LineItemDetails12) AddOrderedQuantity() *Quantity9 {
	l.OrderedQuantity = new(Quantity9)
	return l.OrderedQuantity
}

func (l *LineItemDetails12) AddAcceptedQuantity() *Quantity9 {
	l.AcceptedQuantity = new(Quantity9)
	return l.AcceptedQuantity
}

func (l *LineItemDetails12) AddOutstandingQuantity() *Quantity9 {
	l.OutstandingQuantity = new(Quantity9)
	return l.OutstandingQuantity
}

func (l *LineItemDetails12) AddPendingQuantity() *Quantity9 {
	l.PendingQuantity = new(Quantity9)
	return l.PendingQuantity
}

func (l *LineItemDetails12) AddQuantityTolerance() *PercentageTolerance1 {
	l.QuantityTolerance = new(PercentageTolerance1)
	return l.QuantityTolerance
}

func (l *LineItemDetails12) SetOrderedAmount(value, currency string) {
	l.OrderedAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItemDetails12) SetAcceptedAmount(value, currency string) {
	l.AcceptedAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItemDetails12) SetOutstandingAmount(value, currency string) {
	l.OutstandingAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItemDetails12) SetPendingAmount(value, currency string) {
	l.PendingAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItemDetails12) AddPriceTolerance() *PercentageTolerance1 {
	l.PriceTolerance = new(PercentageTolerance1)
	return l.PriceTolerance
}

// Goods or services that are part of a commercial trade agreement.
type LineItemDetails13 struct {

	// Identification assigned to a line item.
	LineItemIdentification *Max70Text `xml:"LineItmId"`

	// Specifies the quantity of goods on a line in a trade transaction.
	Quantity *Quantity9 `xml:"Qty"`

	// Variance allowed in the quantity of goods.
	QuantityTolerance *PercentageTolerance1 `xml:"QtyTlrnce,omitempty"`

	// Amount of money for which goods or services are offered, sold, or bought.
	UnitPrice *UnitPrice18 `xml:"UnitPric,omitempty"`

	// Variance allowed on a price.
	PriceTolerance *PercentageTolerance1 `xml:"PricTlrnce,omitempty"`

	// Name of the product detailed in the corresponding line item.
	ProductName *Max70Text `xml:"PdctNm,omitempty"`

	// Identifies the product of the corresponding line item.
	ProductIdentifier []*ProductIdentifier2Choice `xml:"PdctIdr,omitempty"`

	// Identifies the characteristics of a product.
	ProductCharacteristics []*ProductCharacteristics1Choice `xml:"PdctChrtcs,omitempty"`

	// Identifies the category of product.
	ProductCategory []*ProductCategory1Choice `xml:"PdctCtgy,omitempty"`

	// Country from which the product originates.
	ProductOrigin []*CountryCode `xml:"PdctOrgn,omitempty"`

	// Specifies the shipment schedule for the goods.
	ShipmentSchedule *ShipmentSchedule2Choice `xml:"ShipmntSchdl,omitempty"`

	// Information related to the conveyance of goods.
	RoutingSummary *TransportMeans5 `xml:"RtgSummry,omitempty"`

	// Variance on price for the goods.
	Adjustment []*Adjustment7 `xml:"Adjstmnt,omitempty"`

	// Maximum charges related to the conveyance of goods.
	FreightCharges *Charge24 `xml:"FrghtChrgs,omitempty"`

	// Amount of money due to the government or tax authority, according to various pre-defined parameters linked to the value of the goods in a trade transaction.
	Tax []*Tax23 `xml:"Tax,omitempty"`

	// Total amount of the line item after adjustments have been applied.
	TotalAmount *CurrencyAndAmount `xml:"TtlAmt"`

	// Specifies the applicable Incoterms and associated location. Latest version of Incoterms in effect at the date of message creation.
	Incoterms *Incoterms4 `xml:"Incotrms,omitempty"`
}

func (l *LineItemDetails13) SetLineItemIdentification(value string) {
	l.LineItemIdentification = (*Max70Text)(&value)
}

func (l *LineItemDetails13) AddQuantity() *Quantity9 {
	l.Quantity = new(Quantity9)
	return l.Quantity
}

func (l *LineItemDetails13) AddQuantityTolerance() *PercentageTolerance1 {
	l.QuantityTolerance = new(PercentageTolerance1)
	return l.QuantityTolerance
}

func (l *LineItemDetails13) AddUnitPrice() *UnitPrice18 {
	l.UnitPrice = new(UnitPrice18)
	return l.UnitPrice
}

func (l *LineItemDetails13) AddPriceTolerance() *PercentageTolerance1 {
	l.PriceTolerance = new(PercentageTolerance1)
	return l.PriceTolerance
}

func (l *LineItemDetails13) SetProductName(value string) {
	l.ProductName = (*Max70Text)(&value)
}

func (l *LineItemDetails13) AddProductIdentifier() *ProductIdentifier2Choice {
	newValue := new(ProductIdentifier2Choice)
	l.ProductIdentifier = append(l.ProductIdentifier, newValue)
	return newValue
}

func (l *LineItemDetails13) AddProductCharacteristics() *ProductCharacteristics1Choice {
	newValue := new(ProductCharacteristics1Choice)
	l.ProductCharacteristics = append(l.ProductCharacteristics, newValue)
	return newValue
}

func (l *LineItemDetails13) AddProductCategory() *ProductCategory1Choice {
	newValue := new(ProductCategory1Choice)
	l.ProductCategory = append(l.ProductCategory, newValue)
	return newValue
}

func (l *LineItemDetails13) AddProductOrigin(value string) {
	l.ProductOrigin = append(l.ProductOrigin, (*CountryCode)(&value))
}

func (l *LineItemDetails13) AddShipmentSchedule() *ShipmentSchedule2Choice {
	l.ShipmentSchedule = new(ShipmentSchedule2Choice)
	return l.ShipmentSchedule
}

func (l *LineItemDetails13) AddRoutingSummary() *TransportMeans5 {
	l.RoutingSummary = new(TransportMeans5)
	return l.RoutingSummary
}

func (l *LineItemDetails13) AddAdjustment() *Adjustment7 {
	newValue := new(Adjustment7)
	l.Adjustment = append(l.Adjustment, newValue)
	return newValue
}

func (l *LineItemDetails13) AddFreightCharges() *Charge24 {
	l.FreightCharges = new(Charge24)
	return l.FreightCharges
}

func (l *LineItemDetails13) AddTax() *Tax23 {
	newValue := new(Tax23)
	l.Tax = append(l.Tax, newValue)
	return newValue
}

func (l *LineItemDetails13) SetTotalAmount(value, currency string) {
	l.TotalAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItemDetails13) AddIncoterms() *Incoterms4 {
	l.Incoterms = new(Incoterms4)
	return l.Incoterms
}

// Goods or services that are part of a commercial trade agreement.
type LineItemDetails14 struct {

	// Sequential number assigned to a line item.
	LineItemIdentification *Max70Text `xml:"LineItmId"`

	// Specifies the quantity of a product in a trade transaction.
	Quantity *Quantity9 `xml:"Qty"`

	// Amount of money for which goods or services are offered, sold, or bought.
	UnitPrice *UnitPrice18 `xml:"UnitPric,omitempty"`

	// Name of the product detailed in the corresponding line item.
	ProductName *Max70Text `xml:"PdctNm,omitempty"`

	// Identifies the product of the corresponding line item.
	ProductIdentifier []*ProductIdentifier2Choice `xml:"PdctIdr,omitempty"`

	// Identifies the characteristics of product.
	ProductCharacteristics []*ProductCharacteristics1Choice `xml:"PdctChrtcs,omitempty"`

	// Identifies the category of product.
	ProductCategory []*ProductCategory1Choice `xml:"PdctCtgy,omitempty"`

	// Country of origin of the goods.
	ProductOrigin *CountryCode `xml:"PdctOrgn,omitempty"`

	// Variance on price for the goods.
	Adjustment []*Adjustment6 `xml:"Adjstmnt,omitempty"`

	// Charges related to the conveyance of goods.
	FreightCharges *Charge25 `xml:"FrghtChrgs,omitempty"`

	// Amount of money due to the government or tax authority, according to various pre-defined parameters linked to the value of the goods in a trade transaction.
	Tax []*Tax22 `xml:"Tax,omitempty"`

	// Total amount of the line item after adjustments have been applied.
	TotalAmount *CurrencyAndAmount `xml:"TtlAmt"`
}

func (l *LineItemDetails14) SetLineItemIdentification(value string) {
	l.LineItemIdentification = (*Max70Text)(&value)
}

func (l *LineItemDetails14) AddQuantity() *Quantity9 {
	l.Quantity = new(Quantity9)
	return l.Quantity
}

func (l *LineItemDetails14) AddUnitPrice() *UnitPrice18 {
	l.UnitPrice = new(UnitPrice18)
	return l.UnitPrice
}

func (l *LineItemDetails14) SetProductName(value string) {
	l.ProductName = (*Max70Text)(&value)
}

func (l *LineItemDetails14) AddProductIdentifier() *ProductIdentifier2Choice {
	newValue := new(ProductIdentifier2Choice)
	l.ProductIdentifier = append(l.ProductIdentifier, newValue)
	return newValue
}

func (l *LineItemDetails14) AddProductCharacteristics() *ProductCharacteristics1Choice {
	newValue := new(ProductCharacteristics1Choice)
	l.ProductCharacteristics = append(l.ProductCharacteristics, newValue)
	return newValue
}

func (l *LineItemDetails14) AddProductCategory() *ProductCategory1Choice {
	newValue := new(ProductCategory1Choice)
	l.ProductCategory = append(l.ProductCategory, newValue)
	return newValue
}

func (l *LineItemDetails14) SetProductOrigin(value string) {
	l.ProductOrigin = (*CountryCode)(&value)
}

func (l *LineItemDetails14) AddAdjustment() *Adjustment6 {
	newValue := new(Adjustment6)
	l.Adjustment = append(l.Adjustment, newValue)
	return newValue
}

func (l *LineItemDetails14) AddFreightCharges() *Charge25 {
	l.FreightCharges = new(Charge25)
	return l.FreightCharges
}

func (l *LineItemDetails14) AddTax() *Tax22 {
	newValue := new(Tax22)
	l.Tax = append(l.Tax, newValue)
	return newValue
}

func (l *LineItemDetails14) SetTotalAmount(value, currency string) {
	l.TotalAmount = NewCurrencyAndAmount(value, currency)
}

// Goods or services that are part of a commercial trade agreement.
type LineItemDetails7 struct {

	// Identification assigned to a line item.
	LineItemIdentification *Max70Text `xml:"LineItmId"`

	// Specifies the quantity of goods on a line in a trade transaction.
	Quantity *Quantity4 `xml:"Qty"`

	// Variance allowed in the quantity of goods.
	QuantityTolerance *PercentageTolerance1 `xml:"QtyTlrnce,omitempty"`

	// Amount of money for which goods or services are offered, sold, or bought.
	UnitPrice *UnitPrice9 `xml:"UnitPric,omitempty"`

	// Variance allowed on a price.
	PriceTolerance *PercentageTolerance1 `xml:"PricTlrnce,omitempty"`

	// Name of the product detailed in the corresponding line item.
	ProductName *Max70Text `xml:"PdctNm,omitempty"`

	// Identifies the product of the corresponding line item.
	ProductIdentifier []*ProductIdentifier2Choice `xml:"PdctIdr,omitempty"`

	// Identifies the characteristics of a product.
	ProductCharacteristics []*ProductCharacteristics1Choice `xml:"PdctChrtcs,omitempty"`

	// Identifies the category of product.
	ProductCategory []*ProductCategory1Choice `xml:"PdctCtgy,omitempty"`

	// Country from which the product originates.
	ProductOrigin []*CountryCode `xml:"PdctOrgn,omitempty"`

	// Specifies the shipment schedule for the goods.
	ShipmentSchedule *ShipmentSchedule1Choice `xml:"ShipmntSchdl,omitempty"`

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

	// Total amount of the line item after adjustments have been applied.
	TotalAmount *CurrencyAndAmount `xml:"TtlAmt"`
}

func (l *LineItemDetails7) SetLineItemIdentification(value string) {
	l.LineItemIdentification = (*Max70Text)(&value)
}

func (l *LineItemDetails7) AddQuantity() *Quantity4 {
	l.Quantity = new(Quantity4)
	return l.Quantity
}

func (l *LineItemDetails7) AddQuantityTolerance() *PercentageTolerance1 {
	l.QuantityTolerance = new(PercentageTolerance1)
	return l.QuantityTolerance
}

func (l *LineItemDetails7) AddUnitPrice() *UnitPrice9 {
	l.UnitPrice = new(UnitPrice9)
	return l.UnitPrice
}

func (l *LineItemDetails7) AddPriceTolerance() *PercentageTolerance1 {
	l.PriceTolerance = new(PercentageTolerance1)
	return l.PriceTolerance
}

func (l *LineItemDetails7) SetProductName(value string) {
	l.ProductName = (*Max70Text)(&value)
}

func (l *LineItemDetails7) AddProductIdentifier() *ProductIdentifier2Choice {
	newValue := new(ProductIdentifier2Choice)
	l.ProductIdentifier = append(l.ProductIdentifier, newValue)
	return newValue
}

func (l *LineItemDetails7) AddProductCharacteristics() *ProductCharacteristics1Choice {
	newValue := new(ProductCharacteristics1Choice)
	l.ProductCharacteristics = append(l.ProductCharacteristics, newValue)
	return newValue
}

func (l *LineItemDetails7) AddProductCategory() *ProductCategory1Choice {
	newValue := new(ProductCategory1Choice)
	l.ProductCategory = append(l.ProductCategory, newValue)
	return newValue
}

func (l *LineItemDetails7) AddProductOrigin(value string) {
	l.ProductOrigin = append(l.ProductOrigin, (*CountryCode)(&value))
}

func (l *LineItemDetails7) AddShipmentSchedule() *ShipmentSchedule1Choice {
	l.ShipmentSchedule = new(ShipmentSchedule1Choice)
	return l.ShipmentSchedule
}

func (l *LineItemDetails7) AddRoutingSummary() *TransportMeans1 {
	l.RoutingSummary = new(TransportMeans1)
	return l.RoutingSummary
}

func (l *LineItemDetails7) AddIncoterms() *Incoterms1 {
	newValue := new(Incoterms1)
	l.Incoterms = append(l.Incoterms, newValue)
	return newValue
}

func (l *LineItemDetails7) AddAdjustment() *Adjustment3 {
	newValue := new(Adjustment3)
	l.Adjustment = append(l.Adjustment, newValue)
	return newValue
}

func (l *LineItemDetails7) AddFreightCharges() *Charge12 {
	l.FreightCharges = new(Charge12)
	return l.FreightCharges
}

func (l *LineItemDetails7) AddTax() *Tax13 {
	newValue := new(Tax13)
	l.Tax = append(l.Tax, newValue)
	return newValue
}

func (l *LineItemDetails7) SetTotalAmount(value, currency string) {
	l.TotalAmount = NewCurrencyAndAmount(value, currency)
}

// Calculation of the current situation of a line item as a result of the submission of a commercial dataset.
type LineItemDetails8 struct {

	// Sequential number assigned to a line item.
	LineItemIdentification *Max70Text `xml:"LineItmId"`

	// Name of the product detailed in the corresponding line item.
	ProductName *Max70Text `xml:"PdctNm,omitempty"`

	// Identifies the product of the corresponding line item.
	ProductIdentifier []*ProductIdentifier2Choice `xml:"PdctIdr,omitempty"`

	// Identifies the characteristic of a product.
	ProductCharacteristics []*ProductCharacteristics1Choice `xml:"PdctChrtcs,omitempty"`

	// Identifies the category of product.
	ProductCategory []*ProductCategory1Choice `xml:"PdctCtgy,omitempty"`

	// Quantity ordered for a line as indicated in the baseline.
	OrderedQuantity *Quantity4 `xml:"OrdrdQty"`

	// Quantity accepted by data set submission.
	AcceptedQuantity *Quantity4 `xml:"AccptdQty"`

	// Difference between the ordered quantity and the accepted one.
	OutstandingQuantity *Quantity4 `xml:"OutsdngQty"`

	// Quantity of a product for which a mismatched data set has been submitted and has not yet been accepted or rejected..
	PendingQuantity *Quantity4 `xml:"PdgQty"`

	// Variance allowed on the quantity of goods.
	QuantityTolerance *PercentageTolerance1 `xml:"QtyTlrnce,omitempty"`

	// Price multiplied by the ordered quantity for a line as indicated in the baseline.
	OrderedAmount *CurrencyAndAmount `xml:"OrdrdAmt"`

	// Price multiplied by the accepted quantity for a line.
	AcceptedAmount *CurrencyAndAmount `xml:"AccptdAmt"`

	// Price multiplied by the outstanding quantity for a line.
	OutstandingAmount *CurrencyAndAmount `xml:"OutsdngAmt"`

	// Price multiplied by the pending quantity for a line.
	PendingAmount *CurrencyAndAmount `xml:"PdgAmt"`

	// Variance on price for the goods.
	PriceTolerance *PercentageTolerance1 `xml:"PricTlrnce,omitempty"`
}

func (l *LineItemDetails8) SetLineItemIdentification(value string) {
	l.LineItemIdentification = (*Max70Text)(&value)
}

func (l *LineItemDetails8) SetProductName(value string) {
	l.ProductName = (*Max70Text)(&value)
}

func (l *LineItemDetails8) AddProductIdentifier() *ProductIdentifier2Choice {
	newValue := new(ProductIdentifier2Choice)
	l.ProductIdentifier = append(l.ProductIdentifier, newValue)
	return newValue
}

func (l *LineItemDetails8) AddProductCharacteristics() *ProductCharacteristics1Choice {
	newValue := new(ProductCharacteristics1Choice)
	l.ProductCharacteristics = append(l.ProductCharacteristics, newValue)
	return newValue
}

func (l *LineItemDetails8) AddProductCategory() *ProductCategory1Choice {
	newValue := new(ProductCategory1Choice)
	l.ProductCategory = append(l.ProductCategory, newValue)
	return newValue
}

func (l *LineItemDetails8) AddOrderedQuantity() *Quantity4 {
	l.OrderedQuantity = new(Quantity4)
	return l.OrderedQuantity
}

func (l *LineItemDetails8) AddAcceptedQuantity() *Quantity4 {
	l.AcceptedQuantity = new(Quantity4)
	return l.AcceptedQuantity
}

func (l *LineItemDetails8) AddOutstandingQuantity() *Quantity4 {
	l.OutstandingQuantity = new(Quantity4)
	return l.OutstandingQuantity
}

func (l *LineItemDetails8) AddPendingQuantity() *Quantity4 {
	l.PendingQuantity = new(Quantity4)
	return l.PendingQuantity
}

func (l *LineItemDetails8) AddQuantityTolerance() *PercentageTolerance1 {
	l.QuantityTolerance = new(PercentageTolerance1)
	return l.QuantityTolerance
}

func (l *LineItemDetails8) SetOrderedAmount(value, currency string) {
	l.OrderedAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItemDetails8) SetAcceptedAmount(value, currency string) {
	l.AcceptedAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItemDetails8) SetOutstandingAmount(value, currency string) {
	l.OutstandingAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItemDetails8) SetPendingAmount(value, currency string) {
	l.PendingAmount = NewCurrencyAndAmount(value, currency)
}

func (l *LineItemDetails8) AddPriceTolerance() *PercentageTolerance1 {
	l.PriceTolerance = new(PercentageTolerance1)
	return l.PriceTolerance
}

// Goods or services that are part of a commercial trade agreement.
type LineItemDetails9 struct {

	// Sequential number assigned to a line item.
	LineItemIdentification *Max70Text `xml:"LineItmId"`

	// Specifies the quantity of a product in a trade transaction.
	Quantity *Quantity4 `xml:"Qty"`

	// Amount of money for which goods or services are offered, sold, or bought.
	UnitPrice *UnitPrice9 `xml:"UnitPric,omitempty"`

	// Name of the product detailed in the corresponding line item.
	ProductName *Max70Text `xml:"PdctNm,omitempty"`

	// Identifies the product of the corresponding line item.
	ProductIdentifier []*ProductIdentifier2Choice `xml:"PdctIdr,omitempty"`

	// Identifies the characteristics of product.
	ProductCharacteristics []*ProductCharacteristics1Choice `xml:"PdctChrtcs,omitempty"`

	// Identifies the category of product.
	ProductCategory []*ProductCategory1Choice `xml:"PdctCtgy,omitempty"`

	// Country of origin of the goods.
	ProductOrigin *CountryCode `xml:"PdctOrgn,omitempty"`

	// Variance on price for the goods.
	Adjustment []*Adjustment4 `xml:"Adjstmnt,omitempty"`

	// Charges related to the conveyance of goods.
	FreightCharges *Charge13 `xml:"FrghtChrgs,omitempty"`

	// Amount of money due to the government or tax authority, according to various pre-defined parameters linked to the value of the goods in a trade transaction.
	Tax []*Tax12 `xml:"Tax,omitempty"`

	// Total amount of the line item after adjustments have been applied.
	TotalAmount *CurrencyAndAmount `xml:"TtlAmt"`
}

func (l *LineItemDetails9) SetLineItemIdentification(value string) {
	l.LineItemIdentification = (*Max70Text)(&value)
}

func (l *LineItemDetails9) AddQuantity() *Quantity4 {
	l.Quantity = new(Quantity4)
	return l.Quantity
}

func (l *LineItemDetails9) AddUnitPrice() *UnitPrice9 {
	l.UnitPrice = new(UnitPrice9)
	return l.UnitPrice
}

func (l *LineItemDetails9) SetProductName(value string) {
	l.ProductName = (*Max70Text)(&value)
}

func (l *LineItemDetails9) AddProductIdentifier() *ProductIdentifier2Choice {
	newValue := new(ProductIdentifier2Choice)
	l.ProductIdentifier = append(l.ProductIdentifier, newValue)
	return newValue
}

func (l *LineItemDetails9) AddProductCharacteristics() *ProductCharacteristics1Choice {
	newValue := new(ProductCharacteristics1Choice)
	l.ProductCharacteristics = append(l.ProductCharacteristics, newValue)
	return newValue
}

func (l *LineItemDetails9) AddProductCategory() *ProductCategory1Choice {
	newValue := new(ProductCategory1Choice)
	l.ProductCategory = append(l.ProductCategory, newValue)
	return newValue
}

func (l *LineItemDetails9) SetProductOrigin(value string) {
	l.ProductOrigin = (*CountryCode)(&value)
}

func (l *LineItemDetails9) AddAdjustment() *Adjustment4 {
	newValue := new(Adjustment4)
	l.Adjustment = append(l.Adjustment, newValue)
	return newValue
}

func (l *LineItemDetails9) AddFreightCharges() *Charge13 {
	l.FreightCharges = new(Charge13)
	return l.FreightCharges
}

func (l *LineItemDetails9) AddTax() *Tax12 {
	newValue := new(Tax12)
	l.Tax = append(l.Tax, newValue)
	return newValue
}

func (l *LineItemDetails9) SetTotalAmount(value, currency string) {
	l.TotalAmount = NewCurrencyAndAmount(value, currency)
}
