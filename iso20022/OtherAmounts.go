package iso20022

// Identifies other amounts pertaining to the transaction.
type OtherAmounts12 struct {

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection9 `xml:"AcrdIntrstAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection9 `xml:"ChrgsFees,omitempty"`

	// Amount of country, national or federal tax charged by the jurisdiction in which the account servicer is located.
	CountryNationalFederalTax *AmountAndDirection9 `xml:"CtryNtlFdrlTax,omitempty"`

	// Principal amount of a trade (price multiplied by quantity).
	TradeAmount *AmountAndDirection9 `xml:"TradAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *AmountAndDirection9 `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of money defined as a discount on a new issue or on a tranche of an existing issue.
	IssueDiscountAllowance *AmountAndDirection9 `xml:"IsseDscntAllwnc,omitempty"`

	// Amount of payment levy tax.
	PaymentLevyTax *AmountAndDirection9 `xml:"PmtLevyTax,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection9 `xml:"LclTax,omitempty"`

	// Local tax country specific.
	LocalTaxCountrySpecific *AmountAndDirection9 `xml:"LclTaxCtrySpcfc,omitempty"`

	// Amount of commission paid to a local broker.
	LocalBrokerCommission *AmountAndDirection9 `xml:"LclBrkrComssn,omitempty"`

	// Amount of money deposited by the trading party in a margin account.
	Margin *AmountAndDirection9 `xml:"Mrgn,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection9 `xml:"Othr,omitempty"`

	// Amount of money paid for delivery by regular post mail.
	PostageAmount *AmountAndDirection9 `xml:"PstgAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryAmount *AmountAndDirection9 `xml:"RgltryAmt,omitempty"`

	// Amount of money (including insurance) paid for delivery by carrier.
	ShippingAmount *AmountAndDirection9 `xml:"ShppgAmt,omitempty"`

	// Amount of drawdown or other reduction from or in addition to the deal price.
	SpecialConcession *AmountAndDirection9 `xml:"SpclCncssn,omitempty"`

	// Amount of stamp duty.
	StampDuty *AmountAndDirection9 `xml:"StmpDty,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *AmountAndDirection9 `xml:"StockXchgTax,omitempty"`

	// Amount of tax levied on a transfer of ownership of financial instrument.
	TransferTax *AmountAndDirection9 `xml:"TrfTax,omitempty"`

	// Amount of transaction tax.
	TransactionTax *AmountAndDirection9 `xml:"TxTax,omitempty"`

	// Amount of value-added tax.
	ValueAddedTax *AmountAndDirection9 `xml:"ValAddedTax,omitempty"`

	// Amount of money that will be withheld by a tax authority.
	WithholdingTax *AmountAndDirection9 `xml:"WhldgTax,omitempty"`

	// Amount representing the difference between the cost and the current price of a security. In the context of securities settlement, it is the amount paid or received when the instructions are netted or paired off.
	NetGainLoss *AmountAndDirection9 `xml:"NetGnLoss,omitempty"`

	// Amount of consumption tax.
	ConsumptionTax *AmountAndDirection9 `xml:"CsmptnTax,omitempty"`

	// Amount of unpaid interest (on bonds which have defaulted and have subsequently
	// restructured), which is capitalized and added to the original principal amount of the bond.
	AccruedCapitalisationAmount *AmountAndDirection9 `xml:"AcrdCptlstnAmt,omitempty"`
}

func (o *OtherAmounts12) AddAccruedInterestAmount() *AmountAndDirection9 {
	o.AccruedInterestAmount = new(AmountAndDirection9)
	return o.AccruedInterestAmount
}

func (o *OtherAmounts12) AddChargesFees() *AmountAndDirection9 {
	o.ChargesFees = new(AmountAndDirection9)
	return o.ChargesFees
}

func (o *OtherAmounts12) AddCountryNationalFederalTax() *AmountAndDirection9 {
	o.CountryNationalFederalTax = new(AmountAndDirection9)
	return o.CountryNationalFederalTax
}

func (o *OtherAmounts12) AddTradeAmount() *AmountAndDirection9 {
	o.TradeAmount = new(AmountAndDirection9)
	return o.TradeAmount
}

func (o *OtherAmounts12) AddExecutingBrokerAmount() *AmountAndDirection9 {
	o.ExecutingBrokerAmount = new(AmountAndDirection9)
	return o.ExecutingBrokerAmount
}

func (o *OtherAmounts12) AddIssueDiscountAllowance() *AmountAndDirection9 {
	o.IssueDiscountAllowance = new(AmountAndDirection9)
	return o.IssueDiscountAllowance
}

func (o *OtherAmounts12) AddPaymentLevyTax() *AmountAndDirection9 {
	o.PaymentLevyTax = new(AmountAndDirection9)
	return o.PaymentLevyTax
}

func (o *OtherAmounts12) AddLocalTax() *AmountAndDirection9 {
	o.LocalTax = new(AmountAndDirection9)
	return o.LocalTax
}

func (o *OtherAmounts12) AddLocalTaxCountrySpecific() *AmountAndDirection9 {
	o.LocalTaxCountrySpecific = new(AmountAndDirection9)
	return o.LocalTaxCountrySpecific
}

func (o *OtherAmounts12) AddLocalBrokerCommission() *AmountAndDirection9 {
	o.LocalBrokerCommission = new(AmountAndDirection9)
	return o.LocalBrokerCommission
}

func (o *OtherAmounts12) AddMargin() *AmountAndDirection9 {
	o.Margin = new(AmountAndDirection9)
	return o.Margin
}

func (o *OtherAmounts12) AddOther() *AmountAndDirection9 {
	o.Other = new(AmountAndDirection9)
	return o.Other
}

func (o *OtherAmounts12) AddPostageAmount() *AmountAndDirection9 {
	o.PostageAmount = new(AmountAndDirection9)
	return o.PostageAmount
}

func (o *OtherAmounts12) AddRegulatoryAmount() *AmountAndDirection9 {
	o.RegulatoryAmount = new(AmountAndDirection9)
	return o.RegulatoryAmount
}

func (o *OtherAmounts12) AddShippingAmount() *AmountAndDirection9 {
	o.ShippingAmount = new(AmountAndDirection9)
	return o.ShippingAmount
}

func (o *OtherAmounts12) AddSpecialConcession() *AmountAndDirection9 {
	o.SpecialConcession = new(AmountAndDirection9)
	return o.SpecialConcession
}

func (o *OtherAmounts12) AddStampDuty() *AmountAndDirection9 {
	o.StampDuty = new(AmountAndDirection9)
	return o.StampDuty
}

func (o *OtherAmounts12) AddStockExchangeTax() *AmountAndDirection9 {
	o.StockExchangeTax = new(AmountAndDirection9)
	return o.StockExchangeTax
}

func (o *OtherAmounts12) AddTransferTax() *AmountAndDirection9 {
	o.TransferTax = new(AmountAndDirection9)
	return o.TransferTax
}

func (o *OtherAmounts12) AddTransactionTax() *AmountAndDirection9 {
	o.TransactionTax = new(AmountAndDirection9)
	return o.TransactionTax
}

func (o *OtherAmounts12) AddValueAddedTax() *AmountAndDirection9 {
	o.ValueAddedTax = new(AmountAndDirection9)
	return o.ValueAddedTax
}

func (o *OtherAmounts12) AddWithholdingTax() *AmountAndDirection9 {
	o.WithholdingTax = new(AmountAndDirection9)
	return o.WithholdingTax
}

func (o *OtherAmounts12) AddNetGainLoss() *AmountAndDirection9 {
	o.NetGainLoss = new(AmountAndDirection9)
	return o.NetGainLoss
}

func (o *OtherAmounts12) AddConsumptionTax() *AmountAndDirection9 {
	o.ConsumptionTax = new(AmountAndDirection9)
	return o.ConsumptionTax
}

func (o *OtherAmounts12) AddAccruedCapitalisationAmount() *AmountAndDirection9 {
	o.AccruedCapitalisationAmount = new(AmountAndDirection9)
	return o.AccruedCapitalisationAmount
}

// Identifies other amounts pertaining to the transaction.
type OtherAmounts14 struct {

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection9 `xml:"AcrdIntrstAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection9 `xml:"ChrgsFees,omitempty"`

	// Amount of country, national or federal tax charged by the jurisdiction in which the account servicer is located.
	CountryNationalFederalTax *AmountAndDirection9 `xml:"CtryNtlFdrlTax,omitempty"`

	// Principal amount of a trade (price multiplied by quantity).
	TradeAmount *AmountAndDirection9 `xml:"TradAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *AmountAndDirection9 `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of money defined as a discount on a new issue or on a tranche of an existing issue.
	IssueDiscountAllowance *AmountAndDirection9 `xml:"IsseDscntAllwnc,omitempty"`

	// Amount of payment levy tax.
	PaymentLevyTax *AmountAndDirection9 `xml:"PmtLevyTax,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection9 `xml:"LclTax,omitempty"`

	// Local tax country specific.
	LocalTaxCountrySpecific *AmountAndDirection9 `xml:"LclTaxCtrySpcfc,omitempty"`

	// Amount of commission paid to a local broker.
	LocalBrokerCommission *AmountAndDirection9 `xml:"LclBrkrComssn,omitempty"`

	// Amount of money deposited by the trading party in a margin account.
	Margin *AmountAndDirection9 `xml:"Mrgn,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection9 `xml:"Othr,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryAmount *AmountAndDirection9 `xml:"RgltryAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingAmount *AmountAndDirection9 `xml:"ShppgAmt,omitempty"`

	// Amount of drawdown or other reduction from or in addition to the deal price.
	SpecialConcession *AmountAndDirection9 `xml:"SpclCncssn,omitempty"`

	// Amount of stamp duty.
	StampDuty *AmountAndDirection9 `xml:"StmpDty,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *AmountAndDirection9 `xml:"StockXchgTax,omitempty"`

	// Amount of tax levied on a transfer of ownership of financial instrument.
	TransferTax *AmountAndDirection9 `xml:"TrfTax,omitempty"`

	// Amount of transaction tax.
	TransactionTax *AmountAndDirection9 `xml:"TxTax,omitempty"`

	// Amount of value-added tax.
	ValueAddedTax *AmountAndDirection9 `xml:"ValAddedTax,omitempty"`

	// Amount of money that will be withheld by a tax authority.
	WithholdingTax *AmountAndDirection9 `xml:"WhldgTax,omitempty"`

	// Amount representing the difference between the cost and the current price of a security. In the context of securities settlement, it is the amount paid or received when the instructions are netted or paired off.
	NetGainLoss *AmountAndDirection9 `xml:"NetGnLoss,omitempty"`

	// Amount of consumption tax.
	ConsumptionTax *AmountAndDirection9 `xml:"CsmptnTax,omitempty"`

	// Amount of unpaid interest (on bonds which have defaulted and have subsequently
	// restructured), which is capitalized and added to the original principal amount of the bond.
	AccruedCapitalisationAmount *AmountAndDirection9 `xml:"AcrdCptlstnAmt,omitempty"`
}

func (o *OtherAmounts14) AddAccruedInterestAmount() *AmountAndDirection9 {
	o.AccruedInterestAmount = new(AmountAndDirection9)
	return o.AccruedInterestAmount
}

func (o *OtherAmounts14) AddChargesFees() *AmountAndDirection9 {
	o.ChargesFees = new(AmountAndDirection9)
	return o.ChargesFees
}

func (o *OtherAmounts14) AddCountryNationalFederalTax() *AmountAndDirection9 {
	o.CountryNationalFederalTax = new(AmountAndDirection9)
	return o.CountryNationalFederalTax
}

func (o *OtherAmounts14) AddTradeAmount() *AmountAndDirection9 {
	o.TradeAmount = new(AmountAndDirection9)
	return o.TradeAmount
}

func (o *OtherAmounts14) AddExecutingBrokerAmount() *AmountAndDirection9 {
	o.ExecutingBrokerAmount = new(AmountAndDirection9)
	return o.ExecutingBrokerAmount
}

func (o *OtherAmounts14) AddIssueDiscountAllowance() *AmountAndDirection9 {
	o.IssueDiscountAllowance = new(AmountAndDirection9)
	return o.IssueDiscountAllowance
}

func (o *OtherAmounts14) AddPaymentLevyTax() *AmountAndDirection9 {
	o.PaymentLevyTax = new(AmountAndDirection9)
	return o.PaymentLevyTax
}

func (o *OtherAmounts14) AddLocalTax() *AmountAndDirection9 {
	o.LocalTax = new(AmountAndDirection9)
	return o.LocalTax
}

func (o *OtherAmounts14) AddLocalTaxCountrySpecific() *AmountAndDirection9 {
	o.LocalTaxCountrySpecific = new(AmountAndDirection9)
	return o.LocalTaxCountrySpecific
}

func (o *OtherAmounts14) AddLocalBrokerCommission() *AmountAndDirection9 {
	o.LocalBrokerCommission = new(AmountAndDirection9)
	return o.LocalBrokerCommission
}

func (o *OtherAmounts14) AddMargin() *AmountAndDirection9 {
	o.Margin = new(AmountAndDirection9)
	return o.Margin
}

func (o *OtherAmounts14) AddOther() *AmountAndDirection9 {
	o.Other = new(AmountAndDirection9)
	return o.Other
}

func (o *OtherAmounts14) AddRegulatoryAmount() *AmountAndDirection9 {
	o.RegulatoryAmount = new(AmountAndDirection9)
	return o.RegulatoryAmount
}

func (o *OtherAmounts14) AddShippingAmount() *AmountAndDirection9 {
	o.ShippingAmount = new(AmountAndDirection9)
	return o.ShippingAmount
}

func (o *OtherAmounts14) AddSpecialConcession() *AmountAndDirection9 {
	o.SpecialConcession = new(AmountAndDirection9)
	return o.SpecialConcession
}

func (o *OtherAmounts14) AddStampDuty() *AmountAndDirection9 {
	o.StampDuty = new(AmountAndDirection9)
	return o.StampDuty
}

func (o *OtherAmounts14) AddStockExchangeTax() *AmountAndDirection9 {
	o.StockExchangeTax = new(AmountAndDirection9)
	return o.StockExchangeTax
}

func (o *OtherAmounts14) AddTransferTax() *AmountAndDirection9 {
	o.TransferTax = new(AmountAndDirection9)
	return o.TransferTax
}

func (o *OtherAmounts14) AddTransactionTax() *AmountAndDirection9 {
	o.TransactionTax = new(AmountAndDirection9)
	return o.TransactionTax
}

func (o *OtherAmounts14) AddValueAddedTax() *AmountAndDirection9 {
	o.ValueAddedTax = new(AmountAndDirection9)
	return o.ValueAddedTax
}

func (o *OtherAmounts14) AddWithholdingTax() *AmountAndDirection9 {
	o.WithholdingTax = new(AmountAndDirection9)
	return o.WithholdingTax
}

func (o *OtherAmounts14) AddNetGainLoss() *AmountAndDirection9 {
	o.NetGainLoss = new(AmountAndDirection9)
	return o.NetGainLoss
}

func (o *OtherAmounts14) AddConsumptionTax() *AmountAndDirection9 {
	o.ConsumptionTax = new(AmountAndDirection9)
	return o.ConsumptionTax
}

func (o *OtherAmounts14) AddAccruedCapitalisationAmount() *AmountAndDirection9 {
	o.AccruedCapitalisationAmount = new(AmountAndDirection9)
	return o.AccruedCapitalisationAmount
}

// Identifies other amounts pertaining to the transaction.
type OtherAmounts15 struct {

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection9 `xml:"AcrdIntrstAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection9 `xml:"ChrgsFees,omitempty"`

	// Amount of country, national or federal tax charged by the jurisdiction in which the account servicer is located.
	CountryNationalFederalTax *AmountAndDirection9 `xml:"CtryNtlFdrlTax,omitempty"`

	// Amount of payment levy tax.
	PaymentLevyTax *AmountAndDirection9 `xml:"PmtLevyTax,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection9 `xml:"LclTax,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection9 `xml:"Othr,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryAmount *AmountAndDirection9 `xml:"RgltryAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingAmount *AmountAndDirection9 `xml:"ShppgAmt,omitempty"`

	// Amount of stamp duty.
	StampDuty *AmountAndDirection9 `xml:"StmpDty,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *AmountAndDirection9 `xml:"StockXchgTax,omitempty"`

	// Amount of tax levied on a transfer of ownership of financial instrument.
	TransferTax *AmountAndDirection9 `xml:"TrfTax,omitempty"`

	// Amount of transaction tax.
	TransactionTax *AmountAndDirection9 `xml:"TxTax,omitempty"`

	// Amount of value-added tax.
	ValueAddedTax *AmountAndDirection9 `xml:"ValAddedTax,omitempty"`

	// Amount of money that will be withheld by a tax authority.
	WithholdingTax *AmountAndDirection9 `xml:"WhldgTax,omitempty"`

	// Amount of consumption tax.
	ConsumptionTax *AmountAndDirection9 `xml:"CsmptnTax,omitempty"`

	// Amount of unpaid interest (on bonds which have defaulted and have subsequently
	// restructured), which is capitalized and added to the original principal amount of the bond.
	AccruedCapitalisationAmount *AmountAndDirection9 `xml:"AcrdCptlstnAmt,omitempty"`
}

func (o *OtherAmounts15) AddAccruedInterestAmount() *AmountAndDirection9 {
	o.AccruedInterestAmount = new(AmountAndDirection9)
	return o.AccruedInterestAmount
}

func (o *OtherAmounts15) AddChargesFees() *AmountAndDirection9 {
	o.ChargesFees = new(AmountAndDirection9)
	return o.ChargesFees
}

func (o *OtherAmounts15) AddCountryNationalFederalTax() *AmountAndDirection9 {
	o.CountryNationalFederalTax = new(AmountAndDirection9)
	return o.CountryNationalFederalTax
}

func (o *OtherAmounts15) AddPaymentLevyTax() *AmountAndDirection9 {
	o.PaymentLevyTax = new(AmountAndDirection9)
	return o.PaymentLevyTax
}

func (o *OtherAmounts15) AddLocalTax() *AmountAndDirection9 {
	o.LocalTax = new(AmountAndDirection9)
	return o.LocalTax
}

func (o *OtherAmounts15) AddOther() *AmountAndDirection9 {
	o.Other = new(AmountAndDirection9)
	return o.Other
}

func (o *OtherAmounts15) AddRegulatoryAmount() *AmountAndDirection9 {
	o.RegulatoryAmount = new(AmountAndDirection9)
	return o.RegulatoryAmount
}

func (o *OtherAmounts15) AddShippingAmount() *AmountAndDirection9 {
	o.ShippingAmount = new(AmountAndDirection9)
	return o.ShippingAmount
}

func (o *OtherAmounts15) AddStampDuty() *AmountAndDirection9 {
	o.StampDuty = new(AmountAndDirection9)
	return o.StampDuty
}

func (o *OtherAmounts15) AddStockExchangeTax() *AmountAndDirection9 {
	o.StockExchangeTax = new(AmountAndDirection9)
	return o.StockExchangeTax
}

func (o *OtherAmounts15) AddTransferTax() *AmountAndDirection9 {
	o.TransferTax = new(AmountAndDirection9)
	return o.TransferTax
}

func (o *OtherAmounts15) AddTransactionTax() *AmountAndDirection9 {
	o.TransactionTax = new(AmountAndDirection9)
	return o.TransactionTax
}

func (o *OtherAmounts15) AddValueAddedTax() *AmountAndDirection9 {
	o.ValueAddedTax = new(AmountAndDirection9)
	return o.ValueAddedTax
}

func (o *OtherAmounts15) AddWithholdingTax() *AmountAndDirection9 {
	o.WithholdingTax = new(AmountAndDirection9)
	return o.WithholdingTax
}

func (o *OtherAmounts15) AddConsumptionTax() *AmountAndDirection9 {
	o.ConsumptionTax = new(AmountAndDirection9)
	return o.ConsumptionTax
}

func (o *OtherAmounts15) AddAccruedCapitalisationAmount() *AmountAndDirection9 {
	o.AccruedCapitalisationAmount = new(AmountAndDirection9)
	return o.AccruedCapitalisationAmount
}

// Identifies other amounts pertaining to the transaction.
type OtherAmounts16 struct {

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection29 `xml:"ChrgsFees,omitempty"`

	// Amount of country, national or federal tax charged by the jurisdiction in which the account servicer is located.
	CountryNationalFederalTax *AmountAndDirection29 `xml:"CtryNtlFdrlTax,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *AmountAndDirection29 `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of money defined as a discount on a new issue or on a tranche of an existing issue.
	IssueDiscountAllowance *AmountAndDirection29 `xml:"IsseDscntAllwnc,omitempty"`

	// Amount of payment levy tax.
	PaymentLevyTax *AmountAndDirection29 `xml:"PmtLevyTax,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection29 `xml:"LclTax,omitempty"`

	// Amount of commission paid to a local broker.
	LocalBrokerCommission *AmountAndDirection29 `xml:"LclBrkrComssn,omitempty"`

	// Amount of money deposited by the trading party in a margin account.
	Margin *AmountAndDirection29 `xml:"Mrgn,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection29 `xml:"Othr,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryAmount *AmountAndDirection29 `xml:"RgltryAmt,omitempty"`

	// Amount of drawdown or other reduction from or in addition to the deal price.
	SpecialConcession *AmountAndDirection29 `xml:"SpclCncssn,omitempty"`

	// Amount of stamp duty.
	StampDuty *AmountAndDirection29 `xml:"StmpDty,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *AmountAndDirection29 `xml:"StockXchgTax,omitempty"`

	// Amount of tax levied on a transfer of ownership of financial instrument.
	TransferTax *AmountAndDirection29 `xml:"TrfTax,omitempty"`

	// Amount of transaction tax.
	TransactionTax *AmountAndDirection29 `xml:"TxTax,omitempty"`

	// Amount of value-added tax.
	ValueAddedTax *AmountAndDirection29 `xml:"ValAddedTax,omitempty"`

	// Amount of money that will be withheld by a tax authority.
	WithholdingTax *AmountAndDirection29 `xml:"WhldgTax,omitempty"`

	// Amount representing the difference between the cost and the current price of a security. In the context of securities settlement, it is the amount paid or received when the instructions are netted or paired off.
	NetGainLoss *AmountAndDirection29 `xml:"NetGnLoss,omitempty"`

	// A tax on spending on goods and services.
	ConsumptionTax *AmountAndDirection29 `xml:"CsmptnTax,omitempty"`

	// Amount of money charged for matching and/or confirmation.
	MatchingConfirmationFee *AmountAndDirection29 `xml:"MtchgConfFee,omitempty"`

	// Amount following a foreign exchange conversion.
	ConvertedAmount *AmountAndDirection29 `xml:"ConvtdAmt,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAmount *AmountAndDirection29 `xml:"OrgnlCcyAmt,omitempty"`

	// Cost of the securities. May be requested in some countries for tax purposes.
	BookValue *AmountAndDirection29 `xml:"BookVal,omitempty"`

	// Amount of unpaid interest (on bonds which have defaulted and have subsequently
	// restructured), which is capitalized and added to the original principal amount of the bond.
	AccruedCapitalisationAmount *AmountAndDirection29 `xml:"AcrdCptlstnAmt,omitempty"`

	// Local tax as defined by the country in its market practice document.
	LocalTaxCountrySpecific1 *AmountAndDirection29 `xml:"LclTaxCtrySpcfc1,omitempty"`

	// Local tax as defined by the country in its market practice document.
	LocalTaxCountrySpecific2 *AmountAndDirection29 `xml:"LclTaxCtrySpcfc2,omitempty"`

	// Local tax as defined by the country in its market practice document.
	LocalTaxCountrySpecific3 *AmountAndDirection29 `xml:"LclTaxCtrySpcfc3,omitempty"`

	// Local tax as defined by the country in its market practice document.
	LocalTaxCountrySpecific4 *AmountAndDirection29 `xml:"LclTaxCtrySpcfc4,omitempty"`

	// Amount paid as result of transactions subject to shared brokerage commissions.
	SharedBrokerageAmount *AmountAndDirection29 `xml:"ShrdBrkrgAmt,omitempty"`

	// Indicates the total fees related to the trade, with deduction of rebates (on brokerage, commission or differential) granted by the market member (fee amount with tax excluded).
	MarketMemberFeeAmount *AmountAndDirection29 `xml:"MktMmbFeeAmt,omitempty"`

	// Specifies that this information is available upon request.
	RemunerationAmountRequest *YesNoIndicator `xml:"RmnrtnAmtReq,omitempty"`

	// Specifies the source and amount of any other remuneration received or to be received by the broker in connection with the transaction.
	RemunerationAmount *AmountAndDirection29 `xml:"RmnrtnAmt,omitempty"`

	// Amount to be paid by the lender to the borrower calculated based on the interest rate.
	BorrowingInterestAmount *AmountAndDirection29 `xml:"BrrwgIntrstAmt,omitempty"`

	// Amount to be paid by the Borrower to the Lender for the securities borrowed calculated based on the bond loan rate.
	BorrowingFee *AmountAndDirection29 `xml:"BrrwgFee,omitempty"`

	// Net market value of the securities lent used to calculate the collateral amount.
	NetMarketValue *AmountAndDirection29 `xml:"NetMktVal,omitempty"`

	// Remaining nominal value of a security.
	RemainingFaceValue *AmountAndDirection29 `xml:"RmngFaceVal,omitempty"`

	// Remaining value at which an asset is carried on a balance sheet.
	RemainingBookValue *AmountAndDirection29 `xml:"RmngBookVal,omitempty"`

	// Amount of commission paid to a clearing broker.
	ClearingBrokerCommission *AmountAndDirection29 `xml:"ClrBrkrComssn,omitempty"`

	// Difference between the deal price and another reference price.
	DifferenceInPrice *AmountAndDirection29 `xml:"DiffInPric,omitempty"`

	// Specifies that the odd-lot differential or equivalent fee has been paid by such customer in connection with the execution of an order for an odd-lot number of shares or units (or principal amount) of a security and the fact that the amount of any such differential or fee will be furnished upon oral or written request.
	OddLotFee *YesNoIndicator `xml:"OddLotFee,omitempty"`
}

func (o *OtherAmounts16) AddChargesFees() *AmountAndDirection29 {
	o.ChargesFees = new(AmountAndDirection29)
	return o.ChargesFees
}

func (o *OtherAmounts16) AddCountryNationalFederalTax() *AmountAndDirection29 {
	o.CountryNationalFederalTax = new(AmountAndDirection29)
	return o.CountryNationalFederalTax
}

func (o *OtherAmounts16) AddExecutingBrokerAmount() *AmountAndDirection29 {
	o.ExecutingBrokerAmount = new(AmountAndDirection29)
	return o.ExecutingBrokerAmount
}

func (o *OtherAmounts16) AddIssueDiscountAllowance() *AmountAndDirection29 {
	o.IssueDiscountAllowance = new(AmountAndDirection29)
	return o.IssueDiscountAllowance
}

func (o *OtherAmounts16) AddPaymentLevyTax() *AmountAndDirection29 {
	o.PaymentLevyTax = new(AmountAndDirection29)
	return o.PaymentLevyTax
}

func (o *OtherAmounts16) AddLocalTax() *AmountAndDirection29 {
	o.LocalTax = new(AmountAndDirection29)
	return o.LocalTax
}

func (o *OtherAmounts16) AddLocalBrokerCommission() *AmountAndDirection29 {
	o.LocalBrokerCommission = new(AmountAndDirection29)
	return o.LocalBrokerCommission
}

func (o *OtherAmounts16) AddMargin() *AmountAndDirection29 {
	o.Margin = new(AmountAndDirection29)
	return o.Margin
}

func (o *OtherAmounts16) AddOther() *AmountAndDirection29 {
	o.Other = new(AmountAndDirection29)
	return o.Other
}

func (o *OtherAmounts16) AddRegulatoryAmount() *AmountAndDirection29 {
	o.RegulatoryAmount = new(AmountAndDirection29)
	return o.RegulatoryAmount
}

func (o *OtherAmounts16) AddSpecialConcession() *AmountAndDirection29 {
	o.SpecialConcession = new(AmountAndDirection29)
	return o.SpecialConcession
}

func (o *OtherAmounts16) AddStampDuty() *AmountAndDirection29 {
	o.StampDuty = new(AmountAndDirection29)
	return o.StampDuty
}

func (o *OtherAmounts16) AddStockExchangeTax() *AmountAndDirection29 {
	o.StockExchangeTax = new(AmountAndDirection29)
	return o.StockExchangeTax
}

func (o *OtherAmounts16) AddTransferTax() *AmountAndDirection29 {
	o.TransferTax = new(AmountAndDirection29)
	return o.TransferTax
}

func (o *OtherAmounts16) AddTransactionTax() *AmountAndDirection29 {
	o.TransactionTax = new(AmountAndDirection29)
	return o.TransactionTax
}

func (o *OtherAmounts16) AddValueAddedTax() *AmountAndDirection29 {
	o.ValueAddedTax = new(AmountAndDirection29)
	return o.ValueAddedTax
}

func (o *OtherAmounts16) AddWithholdingTax() *AmountAndDirection29 {
	o.WithholdingTax = new(AmountAndDirection29)
	return o.WithholdingTax
}

func (o *OtherAmounts16) AddNetGainLoss() *AmountAndDirection29 {
	o.NetGainLoss = new(AmountAndDirection29)
	return o.NetGainLoss
}

func (o *OtherAmounts16) AddConsumptionTax() *AmountAndDirection29 {
	o.ConsumptionTax = new(AmountAndDirection29)
	return o.ConsumptionTax
}

func (o *OtherAmounts16) AddMatchingConfirmationFee() *AmountAndDirection29 {
	o.MatchingConfirmationFee = new(AmountAndDirection29)
	return o.MatchingConfirmationFee
}

func (o *OtherAmounts16) AddConvertedAmount() *AmountAndDirection29 {
	o.ConvertedAmount = new(AmountAndDirection29)
	return o.ConvertedAmount
}

func (o *OtherAmounts16) AddOriginalCurrencyAmount() *AmountAndDirection29 {
	o.OriginalCurrencyAmount = new(AmountAndDirection29)
	return o.OriginalCurrencyAmount
}

func (o *OtherAmounts16) AddBookValue() *AmountAndDirection29 {
	o.BookValue = new(AmountAndDirection29)
	return o.BookValue
}

func (o *OtherAmounts16) AddAccruedCapitalisationAmount() *AmountAndDirection29 {
	o.AccruedCapitalisationAmount = new(AmountAndDirection29)
	return o.AccruedCapitalisationAmount
}

func (o *OtherAmounts16) AddLocalTaxCountrySpecific1() *AmountAndDirection29 {
	o.LocalTaxCountrySpecific1 = new(AmountAndDirection29)
	return o.LocalTaxCountrySpecific1
}

func (o *OtherAmounts16) AddLocalTaxCountrySpecific2() *AmountAndDirection29 {
	o.LocalTaxCountrySpecific2 = new(AmountAndDirection29)
	return o.LocalTaxCountrySpecific2
}

func (o *OtherAmounts16) AddLocalTaxCountrySpecific3() *AmountAndDirection29 {
	o.LocalTaxCountrySpecific3 = new(AmountAndDirection29)
	return o.LocalTaxCountrySpecific3
}

func (o *OtherAmounts16) AddLocalTaxCountrySpecific4() *AmountAndDirection29 {
	o.LocalTaxCountrySpecific4 = new(AmountAndDirection29)
	return o.LocalTaxCountrySpecific4
}

func (o *OtherAmounts16) AddSharedBrokerageAmount() *AmountAndDirection29 {
	o.SharedBrokerageAmount = new(AmountAndDirection29)
	return o.SharedBrokerageAmount
}

func (o *OtherAmounts16) AddMarketMemberFeeAmount() *AmountAndDirection29 {
	o.MarketMemberFeeAmount = new(AmountAndDirection29)
	return o.MarketMemberFeeAmount
}

func (o *OtherAmounts16) SetRemunerationAmountRequest(value string) {
	o.RemunerationAmountRequest = (*YesNoIndicator)(&value)
}

func (o *OtherAmounts16) AddRemunerationAmount() *AmountAndDirection29 {
	o.RemunerationAmount = new(AmountAndDirection29)
	return o.RemunerationAmount
}

func (o *OtherAmounts16) AddBorrowingInterestAmount() *AmountAndDirection29 {
	o.BorrowingInterestAmount = new(AmountAndDirection29)
	return o.BorrowingInterestAmount
}

func (o *OtherAmounts16) AddBorrowingFee() *AmountAndDirection29 {
	o.BorrowingFee = new(AmountAndDirection29)
	return o.BorrowingFee
}

func (o *OtherAmounts16) AddNetMarketValue() *AmountAndDirection29 {
	o.NetMarketValue = new(AmountAndDirection29)
	return o.NetMarketValue
}

func (o *OtherAmounts16) AddRemainingFaceValue() *AmountAndDirection29 {
	o.RemainingFaceValue = new(AmountAndDirection29)
	return o.RemainingFaceValue
}

func (o *OtherAmounts16) AddRemainingBookValue() *AmountAndDirection29 {
	o.RemainingBookValue = new(AmountAndDirection29)
	return o.RemainingBookValue
}

func (o *OtherAmounts16) AddClearingBrokerCommission() *AmountAndDirection29 {
	o.ClearingBrokerCommission = new(AmountAndDirection29)
	return o.ClearingBrokerCommission
}

func (o *OtherAmounts16) AddDifferenceInPrice() *AmountAndDirection29 {
	o.DifferenceInPrice = new(AmountAndDirection29)
	return o.DifferenceInPrice
}

func (o *OtherAmounts16) SetOddLotFee(value string) {
	o.OddLotFee = (*YesNoIndicator)(&value)
}

// Identifies other amounts pertaining to the transaction.
type OtherAmounts17 struct {

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection9 `xml:"AcrdIntrstAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection9 `xml:"ChrgsFees,omitempty"`

	// Amount of country, national or federal tax charged by the jurisdiction in which the account servicer is located.
	CountryNationalFederalTax *AmountAndDirection9 `xml:"CtryNtlFdrlTax,omitempty"`

	// Principal amount of a trade (price multiplied by quantity).
	TradeAmount *AmountAndDirection9 `xml:"TradAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *AmountAndDirection9 `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of money defined as a discount on a new issue or on a tranche of an existing issue.
	IssueDiscountAllowance *AmountAndDirection9 `xml:"IsseDscntAllwnc,omitempty"`

	// Amount of payment levy tax.
	PaymentLevyTax *AmountAndDirection9 `xml:"PmtLevyTax,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection9 `xml:"LclTax,omitempty"`

	// Amount of commission paid to a local broker.
	LocalBrokerCommission *AmountAndDirection9 `xml:"LclBrkrComssn,omitempty"`

	// Amount of money deposited by the trading party in a margin account.
	Margin *AmountAndDirection9 `xml:"Mrgn,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection9 `xml:"Othr,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryAmount *AmountAndDirection9 `xml:"RgltryAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingAmount *AmountAndDirection9 `xml:"ShppgAmt,omitempty"`

	// Amount of drawdown or other reduction from or in addition to the deal price.
	SpecialConcession *AmountAndDirection9 `xml:"SpclCncssn,omitempty"`

	// Amount of stamp duty.
	StampDuty *AmountAndDirection9 `xml:"StmpDty,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *AmountAndDirection9 `xml:"StockXchgTax,omitempty"`

	// Amount of tax levied on a transfer of ownership of financial instrument.
	TransferTax *AmountAndDirection9 `xml:"TrfTax,omitempty"`

	// Amount of transaction tax.
	TransactionTax *AmountAndDirection9 `xml:"TxTax,omitempty"`

	// Amount of value-added tax.
	ValueAddedTax *AmountAndDirection9 `xml:"ValAddedTax,omitempty"`

	// Amount of money that will be withheld by a tax authority.
	WithholdingTax *AmountAndDirection9 `xml:"WhldgTax,omitempty"`

	// Amount representing the difference between the cost and the current price of a security. In the context of securities settlement, it is the amount paid or received when the instructions are netted or paired off.
	NetGainLoss *AmountAndDirection9 `xml:"NetGnLoss,omitempty"`

	// Amount of consumption tax.
	ConsumptionTax *AmountAndDirection9 `xml:"CsmptnTax,omitempty"`

	// Amount of unpaid interest (on bonds which have defaulted and have subsequently
	// restructured), which is capitalized and added to the original principal amount of the bond.
	AccruedCapitalisationAmount *AmountAndDirection9 `xml:"AcrdCptlstnAmt,omitempty"`

	// Cost of the securities. May be requested in some countries for tax purposes.
	BookValue *AmountAndDirection9 `xml:"BookVal,omitempty"`
}

func (o *OtherAmounts17) AddAccruedInterestAmount() *AmountAndDirection9 {
	o.AccruedInterestAmount = new(AmountAndDirection9)
	return o.AccruedInterestAmount
}

func (o *OtherAmounts17) AddChargesFees() *AmountAndDirection9 {
	o.ChargesFees = new(AmountAndDirection9)
	return o.ChargesFees
}

func (o *OtherAmounts17) AddCountryNationalFederalTax() *AmountAndDirection9 {
	o.CountryNationalFederalTax = new(AmountAndDirection9)
	return o.CountryNationalFederalTax
}

func (o *OtherAmounts17) AddTradeAmount() *AmountAndDirection9 {
	o.TradeAmount = new(AmountAndDirection9)
	return o.TradeAmount
}

func (o *OtherAmounts17) AddExecutingBrokerAmount() *AmountAndDirection9 {
	o.ExecutingBrokerAmount = new(AmountAndDirection9)
	return o.ExecutingBrokerAmount
}

func (o *OtherAmounts17) AddIssueDiscountAllowance() *AmountAndDirection9 {
	o.IssueDiscountAllowance = new(AmountAndDirection9)
	return o.IssueDiscountAllowance
}

func (o *OtherAmounts17) AddPaymentLevyTax() *AmountAndDirection9 {
	o.PaymentLevyTax = new(AmountAndDirection9)
	return o.PaymentLevyTax
}

func (o *OtherAmounts17) AddLocalTax() *AmountAndDirection9 {
	o.LocalTax = new(AmountAndDirection9)
	return o.LocalTax
}

func (o *OtherAmounts17) AddLocalBrokerCommission() *AmountAndDirection9 {
	o.LocalBrokerCommission = new(AmountAndDirection9)
	return o.LocalBrokerCommission
}

func (o *OtherAmounts17) AddMargin() *AmountAndDirection9 {
	o.Margin = new(AmountAndDirection9)
	return o.Margin
}

func (o *OtherAmounts17) AddOther() *AmountAndDirection9 {
	o.Other = new(AmountAndDirection9)
	return o.Other
}

func (o *OtherAmounts17) AddRegulatoryAmount() *AmountAndDirection9 {
	o.RegulatoryAmount = new(AmountAndDirection9)
	return o.RegulatoryAmount
}

func (o *OtherAmounts17) AddShippingAmount() *AmountAndDirection9 {
	o.ShippingAmount = new(AmountAndDirection9)
	return o.ShippingAmount
}

func (o *OtherAmounts17) AddSpecialConcession() *AmountAndDirection9 {
	o.SpecialConcession = new(AmountAndDirection9)
	return o.SpecialConcession
}

func (o *OtherAmounts17) AddStampDuty() *AmountAndDirection9 {
	o.StampDuty = new(AmountAndDirection9)
	return o.StampDuty
}

func (o *OtherAmounts17) AddStockExchangeTax() *AmountAndDirection9 {
	o.StockExchangeTax = new(AmountAndDirection9)
	return o.StockExchangeTax
}

func (o *OtherAmounts17) AddTransferTax() *AmountAndDirection9 {
	o.TransferTax = new(AmountAndDirection9)
	return o.TransferTax
}

func (o *OtherAmounts17) AddTransactionTax() *AmountAndDirection9 {
	o.TransactionTax = new(AmountAndDirection9)
	return o.TransactionTax
}

func (o *OtherAmounts17) AddValueAddedTax() *AmountAndDirection9 {
	o.ValueAddedTax = new(AmountAndDirection9)
	return o.ValueAddedTax
}

func (o *OtherAmounts17) AddWithholdingTax() *AmountAndDirection9 {
	o.WithholdingTax = new(AmountAndDirection9)
	return o.WithholdingTax
}

func (o *OtherAmounts17) AddNetGainLoss() *AmountAndDirection9 {
	o.NetGainLoss = new(AmountAndDirection9)
	return o.NetGainLoss
}

func (o *OtherAmounts17) AddConsumptionTax() *AmountAndDirection9 {
	o.ConsumptionTax = new(AmountAndDirection9)
	return o.ConsumptionTax
}

func (o *OtherAmounts17) AddAccruedCapitalisationAmount() *AmountAndDirection9 {
	o.AccruedCapitalisationAmount = new(AmountAndDirection9)
	return o.AccruedCapitalisationAmount
}

func (o *OtherAmounts17) AddBookValue() *AmountAndDirection9 {
	o.BookValue = new(AmountAndDirection9)
	return o.BookValue
}

// Identifies other amounts pertaining to the transaction.
type OtherAmounts18 struct {

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection9 `xml:"AcrdIntrstAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection9 `xml:"ChrgsFees,omitempty"`

	// Amount of country, national or federal tax charged by the jurisdiction in which the account servicer is located.
	CountryNationalFederalTax *AmountAndDirection9 `xml:"CtryNtlFdrlTax,omitempty"`

	// Principal amount of a trade (price multiplied by quantity).
	TradeAmount *AmountAndDirection9 `xml:"TradAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *AmountAndDirection9 `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of money defined as a discount on a new issue or on a tranche of an existing issue.
	IssueDiscountAllowance *AmountAndDirection9 `xml:"IsseDscntAllwnc,omitempty"`

	// Amount of payment levy tax.
	PaymentLevyTax *AmountAndDirection9 `xml:"PmtLevyTax,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection9 `xml:"LclTax,omitempty"`

	// Local tax country specific.
	LocalTaxCountrySpecific *AmountAndDirection9 `xml:"LclTaxCtrySpcfc,omitempty"`

	// Amount of commission paid to a local broker.
	LocalBrokerCommission *AmountAndDirection9 `xml:"LclBrkrComssn,omitempty"`

	// Amount of money deposited by the trading party in a margin account.
	Margin *AmountAndDirection9 `xml:"Mrgn,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection9 `xml:"Othr,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryAmount *AmountAndDirection9 `xml:"RgltryAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingAmount *AmountAndDirection9 `xml:"ShppgAmt,omitempty"`

	// Amount of drawdown or other reduction from or in addition to the deal price.
	SpecialConcession *AmountAndDirection9 `xml:"SpclCncssn,omitempty"`

	// Amount of stamp duty.
	StampDuty *AmountAndDirection9 `xml:"StmpDty,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *AmountAndDirection9 `xml:"StockXchgTax,omitempty"`

	// Amount of tax levied on a transfer of ownership of financial instrument.
	TransferTax *AmountAndDirection9 `xml:"TrfTax,omitempty"`

	// Amount of transaction tax.
	TransactionTax *AmountAndDirection9 `xml:"TxTax,omitempty"`

	// Amount of value-added tax.
	ValueAddedTax *AmountAndDirection9 `xml:"ValAddedTax,omitempty"`

	// Amount of money that will be withheld by a tax authority.
	WithholdingTax *AmountAndDirection9 `xml:"WhldgTax,omitempty"`

	// Amount representing the difference between the cost and the current price of a security. In the context of securities settlement, it is the amount paid or received when the instructions are netted or paired off.
	NetGainLoss *AmountAndDirection9 `xml:"NetGnLoss,omitempty"`

	// Amount of consumption tax.
	ConsumptionTax *AmountAndDirection9 `xml:"CsmptnTax,omitempty"`

	// Amount of unpaid interest (on bonds which have defaulted and have subsequently
	// restructured), which is capitalized and added to the original principal amount of the bond.
	AccruedCapitalisationAmount *AmountAndDirection9 `xml:"AcrdCptlstnAmt,omitempty"`

	// Cost of the securities. May be requested in some countries for tax purposes.
	BookValue *AmountAndDirection9 `xml:"BookVal,omitempty"`

	// Value of the collateral available for the delivery settlement process at the account level.
	CollateralMonitorAmount *AmountAndDirection9 `xml:"CollMntrAmt,omitempty"`
}

func (o *OtherAmounts18) AddAccruedInterestAmount() *AmountAndDirection9 {
	o.AccruedInterestAmount = new(AmountAndDirection9)
	return o.AccruedInterestAmount
}

func (o *OtherAmounts18) AddChargesFees() *AmountAndDirection9 {
	o.ChargesFees = new(AmountAndDirection9)
	return o.ChargesFees
}

func (o *OtherAmounts18) AddCountryNationalFederalTax() *AmountAndDirection9 {
	o.CountryNationalFederalTax = new(AmountAndDirection9)
	return o.CountryNationalFederalTax
}

func (o *OtherAmounts18) AddTradeAmount() *AmountAndDirection9 {
	o.TradeAmount = new(AmountAndDirection9)
	return o.TradeAmount
}

func (o *OtherAmounts18) AddExecutingBrokerAmount() *AmountAndDirection9 {
	o.ExecutingBrokerAmount = new(AmountAndDirection9)
	return o.ExecutingBrokerAmount
}

func (o *OtherAmounts18) AddIssueDiscountAllowance() *AmountAndDirection9 {
	o.IssueDiscountAllowance = new(AmountAndDirection9)
	return o.IssueDiscountAllowance
}

func (o *OtherAmounts18) AddPaymentLevyTax() *AmountAndDirection9 {
	o.PaymentLevyTax = new(AmountAndDirection9)
	return o.PaymentLevyTax
}

func (o *OtherAmounts18) AddLocalTax() *AmountAndDirection9 {
	o.LocalTax = new(AmountAndDirection9)
	return o.LocalTax
}

func (o *OtherAmounts18) AddLocalTaxCountrySpecific() *AmountAndDirection9 {
	o.LocalTaxCountrySpecific = new(AmountAndDirection9)
	return o.LocalTaxCountrySpecific
}

func (o *OtherAmounts18) AddLocalBrokerCommission() *AmountAndDirection9 {
	o.LocalBrokerCommission = new(AmountAndDirection9)
	return o.LocalBrokerCommission
}

func (o *OtherAmounts18) AddMargin() *AmountAndDirection9 {
	o.Margin = new(AmountAndDirection9)
	return o.Margin
}

func (o *OtherAmounts18) AddOther() *AmountAndDirection9 {
	o.Other = new(AmountAndDirection9)
	return o.Other
}

func (o *OtherAmounts18) AddRegulatoryAmount() *AmountAndDirection9 {
	o.RegulatoryAmount = new(AmountAndDirection9)
	return o.RegulatoryAmount
}

func (o *OtherAmounts18) AddShippingAmount() *AmountAndDirection9 {
	o.ShippingAmount = new(AmountAndDirection9)
	return o.ShippingAmount
}

func (o *OtherAmounts18) AddSpecialConcession() *AmountAndDirection9 {
	o.SpecialConcession = new(AmountAndDirection9)
	return o.SpecialConcession
}

func (o *OtherAmounts18) AddStampDuty() *AmountAndDirection9 {
	o.StampDuty = new(AmountAndDirection9)
	return o.StampDuty
}

func (o *OtherAmounts18) AddStockExchangeTax() *AmountAndDirection9 {
	o.StockExchangeTax = new(AmountAndDirection9)
	return o.StockExchangeTax
}

func (o *OtherAmounts18) AddTransferTax() *AmountAndDirection9 {
	o.TransferTax = new(AmountAndDirection9)
	return o.TransferTax
}

func (o *OtherAmounts18) AddTransactionTax() *AmountAndDirection9 {
	o.TransactionTax = new(AmountAndDirection9)
	return o.TransactionTax
}

func (o *OtherAmounts18) AddValueAddedTax() *AmountAndDirection9 {
	o.ValueAddedTax = new(AmountAndDirection9)
	return o.ValueAddedTax
}

func (o *OtherAmounts18) AddWithholdingTax() *AmountAndDirection9 {
	o.WithholdingTax = new(AmountAndDirection9)
	return o.WithholdingTax
}

func (o *OtherAmounts18) AddNetGainLoss() *AmountAndDirection9 {
	o.NetGainLoss = new(AmountAndDirection9)
	return o.NetGainLoss
}

func (o *OtherAmounts18) AddConsumptionTax() *AmountAndDirection9 {
	o.ConsumptionTax = new(AmountAndDirection9)
	return o.ConsumptionTax
}

func (o *OtherAmounts18) AddAccruedCapitalisationAmount() *AmountAndDirection9 {
	o.AccruedCapitalisationAmount = new(AmountAndDirection9)
	return o.AccruedCapitalisationAmount
}

func (o *OtherAmounts18) AddBookValue() *AmountAndDirection9 {
	o.BookValue = new(AmountAndDirection9)
	return o.BookValue
}

func (o *OtherAmounts18) AddCollateralMonitorAmount() *AmountAndDirection9 {
	o.CollateralMonitorAmount = new(AmountAndDirection9)
	return o.CollateralMonitorAmount
}

// Identifies other amounts pertaining to the transaction.
type OtherAmounts2 struct {

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection9 `xml:"AcrdIntrstAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection9 `xml:"ChrgsFees,omitempty"`

	// Amount of country, national or federal tax charged by the jurisdiction in which the account servicer is located.
	CountryNationalFederalTax *AmountAndDirection9 `xml:"CtryNtlFdrlTax,omitempty"`

	// Amount of payment levy tax.
	PaymentLevyTax *AmountAndDirection9 `xml:"PmtLevyTax,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection9 `xml:"LclTax,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection9 `xml:"Othr,omitempty"`

	// Amount of money paid for delivery by regular post mail.
	PostageAmount *AmountAndDirection9 `xml:"PstgAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryAmount *AmountAndDirection9 `xml:"RgltryAmt,omitempty"`

	// Amount of money (including insurance) paid for delivery by carrier.
	ShippingAmount *AmountAndDirection9 `xml:"ShppgAmt,omitempty"`

	// Amount of stamp duty.
	StampDuty *AmountAndDirection9 `xml:"StmpDty,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *AmountAndDirection9 `xml:"StockXchgTax,omitempty"`

	// Amount of tax levied on a transfer of ownership of financial instrument.
	TransferTax *AmountAndDirection9 `xml:"TrfTax,omitempty"`

	// Amount of transaction tax.
	TransactionTax *AmountAndDirection9 `xml:"TxTax,omitempty"`

	// Amount of value-added tax.
	ValueAddedTax *AmountAndDirection9 `xml:"ValAddedTax,omitempty"`

	// Amount of money that will be withheld by a tax authority.
	WithholdingTax *AmountAndDirection9 `xml:"WhldgTax,omitempty"`

	// Amount of consumption tax.
	ConsumptionTax *AmountAndDirection9 `xml:"CsmptnTax,omitempty"`

	// Amount of unpaid interest (on bonds which have defaulted and have subsequently
	// restructured), which is capitalized and added to the original principal amount of the bond.
	AccruedCapitalisationAmount *AmountAndDirection9 `xml:"AcrdCptlstnAmt,omitempty"`
}

func (o *OtherAmounts2) AddAccruedInterestAmount() *AmountAndDirection9 {
	o.AccruedInterestAmount = new(AmountAndDirection9)
	return o.AccruedInterestAmount
}

func (o *OtherAmounts2) AddChargesFees() *AmountAndDirection9 {
	o.ChargesFees = new(AmountAndDirection9)
	return o.ChargesFees
}

func (o *OtherAmounts2) AddCountryNationalFederalTax() *AmountAndDirection9 {
	o.CountryNationalFederalTax = new(AmountAndDirection9)
	return o.CountryNationalFederalTax
}

func (o *OtherAmounts2) AddPaymentLevyTax() *AmountAndDirection9 {
	o.PaymentLevyTax = new(AmountAndDirection9)
	return o.PaymentLevyTax
}

func (o *OtherAmounts2) AddLocalTax() *AmountAndDirection9 {
	o.LocalTax = new(AmountAndDirection9)
	return o.LocalTax
}

func (o *OtherAmounts2) AddOther() *AmountAndDirection9 {
	o.Other = new(AmountAndDirection9)
	return o.Other
}

func (o *OtherAmounts2) AddPostageAmount() *AmountAndDirection9 {
	o.PostageAmount = new(AmountAndDirection9)
	return o.PostageAmount
}

func (o *OtherAmounts2) AddRegulatoryAmount() *AmountAndDirection9 {
	o.RegulatoryAmount = new(AmountAndDirection9)
	return o.RegulatoryAmount
}

func (o *OtherAmounts2) AddShippingAmount() *AmountAndDirection9 {
	o.ShippingAmount = new(AmountAndDirection9)
	return o.ShippingAmount
}

func (o *OtherAmounts2) AddStampDuty() *AmountAndDirection9 {
	o.StampDuty = new(AmountAndDirection9)
	return o.StampDuty
}

func (o *OtherAmounts2) AddStockExchangeTax() *AmountAndDirection9 {
	o.StockExchangeTax = new(AmountAndDirection9)
	return o.StockExchangeTax
}

func (o *OtherAmounts2) AddTransferTax() *AmountAndDirection9 {
	o.TransferTax = new(AmountAndDirection9)
	return o.TransferTax
}

func (o *OtherAmounts2) AddTransactionTax() *AmountAndDirection9 {
	o.TransactionTax = new(AmountAndDirection9)
	return o.TransactionTax
}

func (o *OtherAmounts2) AddValueAddedTax() *AmountAndDirection9 {
	o.ValueAddedTax = new(AmountAndDirection9)
	return o.ValueAddedTax
}

func (o *OtherAmounts2) AddWithholdingTax() *AmountAndDirection9 {
	o.WithholdingTax = new(AmountAndDirection9)
	return o.WithholdingTax
}

func (o *OtherAmounts2) AddConsumptionTax() *AmountAndDirection9 {
	o.ConsumptionTax = new(AmountAndDirection9)
	return o.ConsumptionTax
}

func (o *OtherAmounts2) AddAccruedCapitalisationAmount() *AmountAndDirection9 {
	o.AccruedCapitalisationAmount = new(AmountAndDirection9)
	return o.AccruedCapitalisationAmount
}

// Identifies other amounts pertaining to the transaction.
type OtherAmounts28 struct {

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection44 `xml:"AcrdIntrstAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection44 `xml:"ChrgsFees,omitempty"`

	// Amount of country, national or federal tax charged by the jurisdiction in which the account servicer is located.
	CountryNationalFederalTax *AmountAndDirection44 `xml:"CtryNtlFdrlTax,omitempty"`

	// Principal amount of a trade (price multiplied by quantity).
	TradeAmount *AmountAndDirection44 `xml:"TradAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *AmountAndDirection44 `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of money defined as a discount on a new issue or on a tranche of an existing issue.
	IssueDiscountAllowance *AmountAndDirection44 `xml:"IsseDscntAllwnc,omitempty"`

	// Amount of payment levy tax.
	PaymentLevyTax *AmountAndDirection44 `xml:"PmtLevyTax,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection44 `xml:"LclTax,omitempty"`

	// Local tax country specific.
	LocalTaxCountrySpecific *AmountAndDirection44 `xml:"LclTaxCtrySpcfc,omitempty"`

	// Amount of commission paid to a local broker.
	LocalBrokerCommission *AmountAndDirection44 `xml:"LclBrkrComssn,omitempty"`

	// Amount of money deposited by the trading party in a margin account.
	Margin *AmountAndDirection44 `xml:"Mrgn,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection44 `xml:"Othr,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryAmount *AmountAndDirection44 `xml:"RgltryAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingAmount *AmountAndDirection44 `xml:"ShppgAmt,omitempty"`

	// Amount of drawdown or other reduction from or in addition to the deal price.
	SpecialConcession *AmountAndDirection44 `xml:"SpclCncssn,omitempty"`

	// Amount of stamp duty.
	StampDuty *AmountAndDirection44 `xml:"StmpDty,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *AmountAndDirection44 `xml:"StockXchgTax,omitempty"`

	// Amount of tax levied on a transfer of ownership of financial instrument.
	TransferTax *AmountAndDirection44 `xml:"TrfTax,omitempty"`

	// Amount of transaction tax.
	TransactionTax *AmountAndDirection44 `xml:"TxTax,omitempty"`

	// Amount of value-added tax.
	ValueAddedTax *AmountAndDirection44 `xml:"ValAddedTax,omitempty"`

	// Amount of money that will be withheld by a tax authority.
	WithholdingTax *AmountAndDirection44 `xml:"WhldgTax,omitempty"`

	// Amount representing the difference between the cost and the current price of a security. In the context of securities settlement, it is the amount paid or received when the instructions are netted or paired off.
	NetGainLoss *AmountAndDirection44 `xml:"NetGnLoss,omitempty"`

	// Amount of consumption tax.
	ConsumptionTax *AmountAndDirection44 `xml:"CsmptnTax,omitempty"`

	// Amount of unpaid interest (on bonds which have defaulted and have subsequently
	// restructured), which is capitalized and added to the original principal amount of the bond.
	AccruedCapitalisationAmount *AmountAndDirection44 `xml:"AcrdCptlstnAmt,omitempty"`
}

func (o *OtherAmounts28) AddAccruedInterestAmount() *AmountAndDirection44 {
	o.AccruedInterestAmount = new(AmountAndDirection44)
	return o.AccruedInterestAmount
}

func (o *OtherAmounts28) AddChargesFees() *AmountAndDirection44 {
	o.ChargesFees = new(AmountAndDirection44)
	return o.ChargesFees
}

func (o *OtherAmounts28) AddCountryNationalFederalTax() *AmountAndDirection44 {
	o.CountryNationalFederalTax = new(AmountAndDirection44)
	return o.CountryNationalFederalTax
}

func (o *OtherAmounts28) AddTradeAmount() *AmountAndDirection44 {
	o.TradeAmount = new(AmountAndDirection44)
	return o.TradeAmount
}

func (o *OtherAmounts28) AddExecutingBrokerAmount() *AmountAndDirection44 {
	o.ExecutingBrokerAmount = new(AmountAndDirection44)
	return o.ExecutingBrokerAmount
}

func (o *OtherAmounts28) AddIssueDiscountAllowance() *AmountAndDirection44 {
	o.IssueDiscountAllowance = new(AmountAndDirection44)
	return o.IssueDiscountAllowance
}

func (o *OtherAmounts28) AddPaymentLevyTax() *AmountAndDirection44 {
	o.PaymentLevyTax = new(AmountAndDirection44)
	return o.PaymentLevyTax
}

func (o *OtherAmounts28) AddLocalTax() *AmountAndDirection44 {
	o.LocalTax = new(AmountAndDirection44)
	return o.LocalTax
}

func (o *OtherAmounts28) AddLocalTaxCountrySpecific() *AmountAndDirection44 {
	o.LocalTaxCountrySpecific = new(AmountAndDirection44)
	return o.LocalTaxCountrySpecific
}

func (o *OtherAmounts28) AddLocalBrokerCommission() *AmountAndDirection44 {
	o.LocalBrokerCommission = new(AmountAndDirection44)
	return o.LocalBrokerCommission
}

func (o *OtherAmounts28) AddMargin() *AmountAndDirection44 {
	o.Margin = new(AmountAndDirection44)
	return o.Margin
}

func (o *OtherAmounts28) AddOther() *AmountAndDirection44 {
	o.Other = new(AmountAndDirection44)
	return o.Other
}

func (o *OtherAmounts28) AddRegulatoryAmount() *AmountAndDirection44 {
	o.RegulatoryAmount = new(AmountAndDirection44)
	return o.RegulatoryAmount
}

func (o *OtherAmounts28) AddShippingAmount() *AmountAndDirection44 {
	o.ShippingAmount = new(AmountAndDirection44)
	return o.ShippingAmount
}

func (o *OtherAmounts28) AddSpecialConcession() *AmountAndDirection44 {
	o.SpecialConcession = new(AmountAndDirection44)
	return o.SpecialConcession
}

func (o *OtherAmounts28) AddStampDuty() *AmountAndDirection44 {
	o.StampDuty = new(AmountAndDirection44)
	return o.StampDuty
}

func (o *OtherAmounts28) AddStockExchangeTax() *AmountAndDirection44 {
	o.StockExchangeTax = new(AmountAndDirection44)
	return o.StockExchangeTax
}

func (o *OtherAmounts28) AddTransferTax() *AmountAndDirection44 {
	o.TransferTax = new(AmountAndDirection44)
	return o.TransferTax
}

func (o *OtherAmounts28) AddTransactionTax() *AmountAndDirection44 {
	o.TransactionTax = new(AmountAndDirection44)
	return o.TransactionTax
}

func (o *OtherAmounts28) AddValueAddedTax() *AmountAndDirection44 {
	o.ValueAddedTax = new(AmountAndDirection44)
	return o.ValueAddedTax
}

func (o *OtherAmounts28) AddWithholdingTax() *AmountAndDirection44 {
	o.WithholdingTax = new(AmountAndDirection44)
	return o.WithholdingTax
}

func (o *OtherAmounts28) AddNetGainLoss() *AmountAndDirection44 {
	o.NetGainLoss = new(AmountAndDirection44)
	return o.NetGainLoss
}

func (o *OtherAmounts28) AddConsumptionTax() *AmountAndDirection44 {
	o.ConsumptionTax = new(AmountAndDirection44)
	return o.ConsumptionTax
}

func (o *OtherAmounts28) AddAccruedCapitalisationAmount() *AmountAndDirection44 {
	o.AccruedCapitalisationAmount = new(AmountAndDirection44)
	return o.AccruedCapitalisationAmount
}

// Identifies other amounts pertaining to the transaction.
type OtherAmounts29 struct {

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection44 `xml:"AcrdIntrstAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection44 `xml:"ChrgsFees,omitempty"`

	// Amount of country, national or federal tax charged by the jurisdiction in which the account servicer is located.
	CountryNationalFederalTax *AmountAndDirection44 `xml:"CtryNtlFdrlTax,omitempty"`

	// Amount of payment levy tax.
	PaymentLevyTax *AmountAndDirection44 `xml:"PmtLevyTax,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection44 `xml:"LclTax,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection44 `xml:"Othr,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryAmount *AmountAndDirection44 `xml:"RgltryAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingAmount *AmountAndDirection44 `xml:"ShppgAmt,omitempty"`

	// Amount of stamp duty.
	StampDuty *AmountAndDirection44 `xml:"StmpDty,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *AmountAndDirection44 `xml:"StockXchgTax,omitempty"`

	// Amount of tax levied on a transfer of ownership of financial instrument.
	TransferTax *AmountAndDirection44 `xml:"TrfTax,omitempty"`

	// Amount of transaction tax.
	TransactionTax *AmountAndDirection44 `xml:"TxTax,omitempty"`

	// Amount of value-added tax.
	ValueAddedTax *AmountAndDirection44 `xml:"ValAddedTax,omitempty"`

	// Amount of money that will be withheld by a tax authority.
	WithholdingTax *AmountAndDirection44 `xml:"WhldgTax,omitempty"`

	// Amount of consumption tax.
	ConsumptionTax *AmountAndDirection44 `xml:"CsmptnTax,omitempty"`

	// Amount of unpaid interest (on bonds which have defaulted and have subsequently
	// restructured), which is capitalized and added to the original principal amount of the bond.
	AccruedCapitalisationAmount *AmountAndDirection44 `xml:"AcrdCptlstnAmt,omitempty"`
}

func (o *OtherAmounts29) AddAccruedInterestAmount() *AmountAndDirection44 {
	o.AccruedInterestAmount = new(AmountAndDirection44)
	return o.AccruedInterestAmount
}

func (o *OtherAmounts29) AddChargesFees() *AmountAndDirection44 {
	o.ChargesFees = new(AmountAndDirection44)
	return o.ChargesFees
}

func (o *OtherAmounts29) AddCountryNationalFederalTax() *AmountAndDirection44 {
	o.CountryNationalFederalTax = new(AmountAndDirection44)
	return o.CountryNationalFederalTax
}

func (o *OtherAmounts29) AddPaymentLevyTax() *AmountAndDirection44 {
	o.PaymentLevyTax = new(AmountAndDirection44)
	return o.PaymentLevyTax
}

func (o *OtherAmounts29) AddLocalTax() *AmountAndDirection44 {
	o.LocalTax = new(AmountAndDirection44)
	return o.LocalTax
}

func (o *OtherAmounts29) AddOther() *AmountAndDirection44 {
	o.Other = new(AmountAndDirection44)
	return o.Other
}

func (o *OtherAmounts29) AddRegulatoryAmount() *AmountAndDirection44 {
	o.RegulatoryAmount = new(AmountAndDirection44)
	return o.RegulatoryAmount
}

func (o *OtherAmounts29) AddShippingAmount() *AmountAndDirection44 {
	o.ShippingAmount = new(AmountAndDirection44)
	return o.ShippingAmount
}

func (o *OtherAmounts29) AddStampDuty() *AmountAndDirection44 {
	o.StampDuty = new(AmountAndDirection44)
	return o.StampDuty
}

func (o *OtherAmounts29) AddStockExchangeTax() *AmountAndDirection44 {
	o.StockExchangeTax = new(AmountAndDirection44)
	return o.StockExchangeTax
}

func (o *OtherAmounts29) AddTransferTax() *AmountAndDirection44 {
	o.TransferTax = new(AmountAndDirection44)
	return o.TransferTax
}

func (o *OtherAmounts29) AddTransactionTax() *AmountAndDirection44 {
	o.TransactionTax = new(AmountAndDirection44)
	return o.TransactionTax
}

func (o *OtherAmounts29) AddValueAddedTax() *AmountAndDirection44 {
	o.ValueAddedTax = new(AmountAndDirection44)
	return o.ValueAddedTax
}

func (o *OtherAmounts29) AddWithholdingTax() *AmountAndDirection44 {
	o.WithholdingTax = new(AmountAndDirection44)
	return o.WithholdingTax
}

func (o *OtherAmounts29) AddConsumptionTax() *AmountAndDirection44 {
	o.ConsumptionTax = new(AmountAndDirection44)
	return o.ConsumptionTax
}

func (o *OtherAmounts29) AddAccruedCapitalisationAmount() *AmountAndDirection44 {
	o.AccruedCapitalisationAmount = new(AmountAndDirection44)
	return o.AccruedCapitalisationAmount
}

// Identifies other amounts pertaining to the transaction.
type OtherAmounts3 struct {

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection9 `xml:"AcrdIntrstAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection9 `xml:"ChrgsFees,omitempty"`

	// Amount of country, national or federal tax charged by the jurisdiction in which the account servicer is located.
	CountryNationalFederalTax *AmountAndDirection9 `xml:"CtryNtlFdrlTax,omitempty"`

	// Principal amount of a trade (price multiplied by quantity).
	TradeAmount *AmountAndDirection9 `xml:"TradAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *AmountAndDirection9 `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of money defined as a discount on a new issue or on a tranche of an existing issue.
	IssueDiscountAllowance *AmountAndDirection9 `xml:"IsseDscntAllwnc,omitempty"`

	// Amount of payment levy tax.
	PaymentLevyTax *AmountAndDirection9 `xml:"PmtLevyTax,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection9 `xml:"LclTax,omitempty"`

	// Amount of commission paid to a local broker.
	LocalBrokerCommission *AmountAndDirection9 `xml:"LclBrkrComssn,omitempty"`

	// Amount of money deposited by the trading party in a margin account.
	Margin *AmountAndDirection9 `xml:"Mrgn,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection9 `xml:"Othr,omitempty"`

	// Amount of money paid for delivery by regular post mail.
	PostageAmount *AmountAndDirection9 `xml:"PstgAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryAmount *AmountAndDirection9 `xml:"RgltryAmt,omitempty"`

	// Amount of money (including insurance) paid for delivery by carrier.
	ShippingAmount *AmountAndDirection9 `xml:"ShppgAmt,omitempty"`

	// Amount of drawdown or other reduction from or in addition to the deal price.
	SpecialConcession *AmountAndDirection9 `xml:"SpclCncssn,omitempty"`

	// Amount of stamp duty.
	StampDuty *AmountAndDirection9 `xml:"StmpDty,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *AmountAndDirection9 `xml:"StockXchgTax,omitempty"`

	// Amount of tax levied on a transfer of ownership of financial instrument.
	TransferTax *AmountAndDirection9 `xml:"TrfTax,omitempty"`

	// Amount of transaction tax.
	TransactionTax *AmountAndDirection9 `xml:"TxTax,omitempty"`

	// Amount of value-added tax.
	ValueAddedTax *AmountAndDirection9 `xml:"ValAddedTax,omitempty"`

	// Amount of money that will be withheld by a tax authority.
	WithholdingTax *AmountAndDirection9 `xml:"WhldgTax,omitempty"`

	// Amount representing the difference between the cost and the current price of a security. In the context of securities settlement, it is the amount paid or received when the instructions are netted or paired off.
	NetGainLoss *AmountAndDirection9 `xml:"NetGnLoss,omitempty"`

	// Amount of consumption tax.
	ConsumptionTax *AmountAndDirection9 `xml:"CsmptnTax,omitempty"`

	// Amount of unpaid interest (on bonds which have defaulted and have subsequently
	// restructured), which is capitalized and added to the original principal amount of the bond.
	AccruedCapitalisationAmount *AmountAndDirection9 `xml:"AcrdCptlstnAmt,omitempty"`
}

func (o *OtherAmounts3) AddAccruedInterestAmount() *AmountAndDirection9 {
	o.AccruedInterestAmount = new(AmountAndDirection9)
	return o.AccruedInterestAmount
}

func (o *OtherAmounts3) AddChargesFees() *AmountAndDirection9 {
	o.ChargesFees = new(AmountAndDirection9)
	return o.ChargesFees
}

func (o *OtherAmounts3) AddCountryNationalFederalTax() *AmountAndDirection9 {
	o.CountryNationalFederalTax = new(AmountAndDirection9)
	return o.CountryNationalFederalTax
}

func (o *OtherAmounts3) AddTradeAmount() *AmountAndDirection9 {
	o.TradeAmount = new(AmountAndDirection9)
	return o.TradeAmount
}

func (o *OtherAmounts3) AddExecutingBrokerAmount() *AmountAndDirection9 {
	o.ExecutingBrokerAmount = new(AmountAndDirection9)
	return o.ExecutingBrokerAmount
}

func (o *OtherAmounts3) AddIssueDiscountAllowance() *AmountAndDirection9 {
	o.IssueDiscountAllowance = new(AmountAndDirection9)
	return o.IssueDiscountAllowance
}

func (o *OtherAmounts3) AddPaymentLevyTax() *AmountAndDirection9 {
	o.PaymentLevyTax = new(AmountAndDirection9)
	return o.PaymentLevyTax
}

func (o *OtherAmounts3) AddLocalTax() *AmountAndDirection9 {
	o.LocalTax = new(AmountAndDirection9)
	return o.LocalTax
}

func (o *OtherAmounts3) AddLocalBrokerCommission() *AmountAndDirection9 {
	o.LocalBrokerCommission = new(AmountAndDirection9)
	return o.LocalBrokerCommission
}

func (o *OtherAmounts3) AddMargin() *AmountAndDirection9 {
	o.Margin = new(AmountAndDirection9)
	return o.Margin
}

func (o *OtherAmounts3) AddOther() *AmountAndDirection9 {
	o.Other = new(AmountAndDirection9)
	return o.Other
}

func (o *OtherAmounts3) AddPostageAmount() *AmountAndDirection9 {
	o.PostageAmount = new(AmountAndDirection9)
	return o.PostageAmount
}

func (o *OtherAmounts3) AddRegulatoryAmount() *AmountAndDirection9 {
	o.RegulatoryAmount = new(AmountAndDirection9)
	return o.RegulatoryAmount
}

func (o *OtherAmounts3) AddShippingAmount() *AmountAndDirection9 {
	o.ShippingAmount = new(AmountAndDirection9)
	return o.ShippingAmount
}

func (o *OtherAmounts3) AddSpecialConcession() *AmountAndDirection9 {
	o.SpecialConcession = new(AmountAndDirection9)
	return o.SpecialConcession
}

func (o *OtherAmounts3) AddStampDuty() *AmountAndDirection9 {
	o.StampDuty = new(AmountAndDirection9)
	return o.StampDuty
}

func (o *OtherAmounts3) AddStockExchangeTax() *AmountAndDirection9 {
	o.StockExchangeTax = new(AmountAndDirection9)
	return o.StockExchangeTax
}

func (o *OtherAmounts3) AddTransferTax() *AmountAndDirection9 {
	o.TransferTax = new(AmountAndDirection9)
	return o.TransferTax
}

func (o *OtherAmounts3) AddTransactionTax() *AmountAndDirection9 {
	o.TransactionTax = new(AmountAndDirection9)
	return o.TransactionTax
}

func (o *OtherAmounts3) AddValueAddedTax() *AmountAndDirection9 {
	o.ValueAddedTax = new(AmountAndDirection9)
	return o.ValueAddedTax
}

func (o *OtherAmounts3) AddWithholdingTax() *AmountAndDirection9 {
	o.WithholdingTax = new(AmountAndDirection9)
	return o.WithholdingTax
}

func (o *OtherAmounts3) AddNetGainLoss() *AmountAndDirection9 {
	o.NetGainLoss = new(AmountAndDirection9)
	return o.NetGainLoss
}

func (o *OtherAmounts3) AddConsumptionTax() *AmountAndDirection9 {
	o.ConsumptionTax = new(AmountAndDirection9)
	return o.ConsumptionTax
}

func (o *OtherAmounts3) AddAccruedCapitalisationAmount() *AmountAndDirection9 {
	o.AccruedCapitalisationAmount = new(AmountAndDirection9)
	return o.AccruedCapitalisationAmount
}

// Identifies other amounts pertaining to the transaction.
type OtherAmounts30 struct {

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection44 `xml:"AcrdIntrstAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection44 `xml:"ChrgsFees,omitempty"`

	// Amount of country, national or federal tax charged by the jurisdiction in which the account servicer is located.
	CountryNationalFederalTax *AmountAndDirection44 `xml:"CtryNtlFdrlTax,omitempty"`

	// Principal amount of a trade (price multiplied by quantity).
	TradeAmount *AmountAndDirection44 `xml:"TradAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *AmountAndDirection44 `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of money defined as a discount on a new issue or on a tranche of an existing issue.
	IssueDiscountAllowance *AmountAndDirection44 `xml:"IsseDscntAllwnc,omitempty"`

	// Amount of payment levy tax.
	PaymentLevyTax *AmountAndDirection44 `xml:"PmtLevyTax,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection44 `xml:"LclTax,omitempty"`

	// Local tax country specific.
	LocalTaxCountrySpecific *AmountAndDirection44 `xml:"LclTaxCtrySpcfc,omitempty"`

	// Amount of commission paid to a local broker.
	LocalBrokerCommission *AmountAndDirection44 `xml:"LclBrkrComssn,omitempty"`

	// Amount of money deposited by the trading party in a margin account.
	Margin *AmountAndDirection44 `xml:"Mrgn,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection44 `xml:"Othr,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryAmount *AmountAndDirection44 `xml:"RgltryAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingAmount *AmountAndDirection44 `xml:"ShppgAmt,omitempty"`

	// Amount of drawdown or other reduction from or in addition to the deal price.
	SpecialConcession *AmountAndDirection44 `xml:"SpclCncssn,omitempty"`

	// Amount of stamp duty.
	StampDuty *AmountAndDirection44 `xml:"StmpDty,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *AmountAndDirection44 `xml:"StockXchgTax,omitempty"`

	// Amount of tax levied on a transfer of ownership of financial instrument.
	TransferTax *AmountAndDirection44 `xml:"TrfTax,omitempty"`

	// Amount of transaction tax.
	TransactionTax *AmountAndDirection44 `xml:"TxTax,omitempty"`

	// Amount of value-added tax.
	ValueAddedTax *AmountAndDirection44 `xml:"ValAddedTax,omitempty"`

	// Amount of money that will be withheld by a tax authority.
	WithholdingTax *AmountAndDirection44 `xml:"WhldgTax,omitempty"`

	// Amount representing the difference between the cost and the current price of a security. In the context of securities settlement, it is the amount paid or received when the instructions are netted or paired off.
	NetGainLoss *AmountAndDirection44 `xml:"NetGnLoss,omitempty"`

	// Amount of consumption tax.
	ConsumptionTax *AmountAndDirection44 `xml:"CsmptnTax,omitempty"`

	// Amount of unpaid interest (on bonds which have defaulted and have subsequently
	// restructured), which is capitalized and added to the original principal amount of the bond.
	AccruedCapitalisationAmount *AmountAndDirection44 `xml:"AcrdCptlstnAmt,omitempty"`

	// Cost of the securities. May be requested in some countries for tax purposes.
	BookValue *AmountAndDirection44 `xml:"BookVal,omitempty"`

	// Value of the collateral available for the delivery settlement process at the account level.
	CollateralMonitorAmount *AmountAndDirection44 `xml:"CollMntrAmt,omitempty"`
}

func (o *OtherAmounts30) AddAccruedInterestAmount() *AmountAndDirection44 {
	o.AccruedInterestAmount = new(AmountAndDirection44)
	return o.AccruedInterestAmount
}

func (o *OtherAmounts30) AddChargesFees() *AmountAndDirection44 {
	o.ChargesFees = new(AmountAndDirection44)
	return o.ChargesFees
}

func (o *OtherAmounts30) AddCountryNationalFederalTax() *AmountAndDirection44 {
	o.CountryNationalFederalTax = new(AmountAndDirection44)
	return o.CountryNationalFederalTax
}

func (o *OtherAmounts30) AddTradeAmount() *AmountAndDirection44 {
	o.TradeAmount = new(AmountAndDirection44)
	return o.TradeAmount
}

func (o *OtherAmounts30) AddExecutingBrokerAmount() *AmountAndDirection44 {
	o.ExecutingBrokerAmount = new(AmountAndDirection44)
	return o.ExecutingBrokerAmount
}

func (o *OtherAmounts30) AddIssueDiscountAllowance() *AmountAndDirection44 {
	o.IssueDiscountAllowance = new(AmountAndDirection44)
	return o.IssueDiscountAllowance
}

func (o *OtherAmounts30) AddPaymentLevyTax() *AmountAndDirection44 {
	o.PaymentLevyTax = new(AmountAndDirection44)
	return o.PaymentLevyTax
}

func (o *OtherAmounts30) AddLocalTax() *AmountAndDirection44 {
	o.LocalTax = new(AmountAndDirection44)
	return o.LocalTax
}

func (o *OtherAmounts30) AddLocalTaxCountrySpecific() *AmountAndDirection44 {
	o.LocalTaxCountrySpecific = new(AmountAndDirection44)
	return o.LocalTaxCountrySpecific
}

func (o *OtherAmounts30) AddLocalBrokerCommission() *AmountAndDirection44 {
	o.LocalBrokerCommission = new(AmountAndDirection44)
	return o.LocalBrokerCommission
}

func (o *OtherAmounts30) AddMargin() *AmountAndDirection44 {
	o.Margin = new(AmountAndDirection44)
	return o.Margin
}

func (o *OtherAmounts30) AddOther() *AmountAndDirection44 {
	o.Other = new(AmountAndDirection44)
	return o.Other
}

func (o *OtherAmounts30) AddRegulatoryAmount() *AmountAndDirection44 {
	o.RegulatoryAmount = new(AmountAndDirection44)
	return o.RegulatoryAmount
}

func (o *OtherAmounts30) AddShippingAmount() *AmountAndDirection44 {
	o.ShippingAmount = new(AmountAndDirection44)
	return o.ShippingAmount
}

func (o *OtherAmounts30) AddSpecialConcession() *AmountAndDirection44 {
	o.SpecialConcession = new(AmountAndDirection44)
	return o.SpecialConcession
}

func (o *OtherAmounts30) AddStampDuty() *AmountAndDirection44 {
	o.StampDuty = new(AmountAndDirection44)
	return o.StampDuty
}

func (o *OtherAmounts30) AddStockExchangeTax() *AmountAndDirection44 {
	o.StockExchangeTax = new(AmountAndDirection44)
	return o.StockExchangeTax
}

func (o *OtherAmounts30) AddTransferTax() *AmountAndDirection44 {
	o.TransferTax = new(AmountAndDirection44)
	return o.TransferTax
}

func (o *OtherAmounts30) AddTransactionTax() *AmountAndDirection44 {
	o.TransactionTax = new(AmountAndDirection44)
	return o.TransactionTax
}

func (o *OtherAmounts30) AddValueAddedTax() *AmountAndDirection44 {
	o.ValueAddedTax = new(AmountAndDirection44)
	return o.ValueAddedTax
}

func (o *OtherAmounts30) AddWithholdingTax() *AmountAndDirection44 {
	o.WithholdingTax = new(AmountAndDirection44)
	return o.WithholdingTax
}

func (o *OtherAmounts30) AddNetGainLoss() *AmountAndDirection44 {
	o.NetGainLoss = new(AmountAndDirection44)
	return o.NetGainLoss
}

func (o *OtherAmounts30) AddConsumptionTax() *AmountAndDirection44 {
	o.ConsumptionTax = new(AmountAndDirection44)
	return o.ConsumptionTax
}

func (o *OtherAmounts30) AddAccruedCapitalisationAmount() *AmountAndDirection44 {
	o.AccruedCapitalisationAmount = new(AmountAndDirection44)
	return o.AccruedCapitalisationAmount
}

func (o *OtherAmounts30) AddBookValue() *AmountAndDirection44 {
	o.BookValue = new(AmountAndDirection44)
	return o.BookValue
}

func (o *OtherAmounts30) AddCollateralMonitorAmount() *AmountAndDirection44 {
	o.CollateralMonitorAmount = new(AmountAndDirection44)
	return o.CollateralMonitorAmount
}

// Identifies other amounts pertaining to the transaction.
type OtherAmounts31 struct {

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection44 `xml:"AcrdIntrstAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection44 `xml:"ChrgsFees,omitempty"`

	// Amount of country, national or federal tax charged by the jurisdiction in which the account servicer is located.
	CountryNationalFederalTax *AmountAndDirection44 `xml:"CtryNtlFdrlTax,omitempty"`

	// Principal amount of a trade (price multiplied by quantity).
	TradeAmount *AmountAndDirection44 `xml:"TradAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *AmountAndDirection44 `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of money defined as a discount on a new issue or on a tranche of an existing issue.
	IssueDiscountAllowance *AmountAndDirection44 `xml:"IsseDscntAllwnc,omitempty"`

	// Amount of payment levy tax.
	PaymentLevyTax *AmountAndDirection44 `xml:"PmtLevyTax,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection44 `xml:"LclTax,omitempty"`

	// Amount of commission paid to a local broker.
	LocalBrokerCommission *AmountAndDirection44 `xml:"LclBrkrComssn,omitempty"`

	// Amount of money deposited by the trading party in a margin account.
	Margin *AmountAndDirection44 `xml:"Mrgn,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection44 `xml:"Othr,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryAmount *AmountAndDirection44 `xml:"RgltryAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingAmount *AmountAndDirection44 `xml:"ShppgAmt,omitempty"`

	// Amount of drawdown or other reduction from or in addition to the deal price.
	SpecialConcession *AmountAndDirection44 `xml:"SpclCncssn,omitempty"`

	// Amount of stamp duty.
	StampDuty *AmountAndDirection44 `xml:"StmpDty,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *AmountAndDirection44 `xml:"StockXchgTax,omitempty"`

	// Amount of tax levied on a transfer of ownership of financial instrument.
	TransferTax *AmountAndDirection44 `xml:"TrfTax,omitempty"`

	// Amount of transaction tax.
	TransactionTax *AmountAndDirection44 `xml:"TxTax,omitempty"`

	// Amount of value-added tax.
	ValueAddedTax *AmountAndDirection44 `xml:"ValAddedTax,omitempty"`

	// Amount of money that will be withheld by a tax authority.
	WithholdingTax *AmountAndDirection44 `xml:"WhldgTax,omitempty"`

	// Amount representing the difference between the cost and the current price of a security. In the context of securities settlement, it is the amount paid or received when the instructions are netted or paired off.
	NetGainLoss *AmountAndDirection44 `xml:"NetGnLoss,omitempty"`

	// Amount of consumption tax.
	ConsumptionTax *AmountAndDirection44 `xml:"CsmptnTax,omitempty"`

	// Amount of unpaid interest (on bonds which have defaulted and have subsequently
	// restructured), which is capitalized and added to the original principal amount of the bond.
	AccruedCapitalisationAmount *AmountAndDirection44 `xml:"AcrdCptlstnAmt,omitempty"`

	// Cost of the securities. May be requested in some countries for tax purposes.
	BookValue *AmountAndDirection44 `xml:"BookVal,omitempty"`
}

func (o *OtherAmounts31) AddAccruedInterestAmount() *AmountAndDirection44 {
	o.AccruedInterestAmount = new(AmountAndDirection44)
	return o.AccruedInterestAmount
}

func (o *OtherAmounts31) AddChargesFees() *AmountAndDirection44 {
	o.ChargesFees = new(AmountAndDirection44)
	return o.ChargesFees
}

func (o *OtherAmounts31) AddCountryNationalFederalTax() *AmountAndDirection44 {
	o.CountryNationalFederalTax = new(AmountAndDirection44)
	return o.CountryNationalFederalTax
}

func (o *OtherAmounts31) AddTradeAmount() *AmountAndDirection44 {
	o.TradeAmount = new(AmountAndDirection44)
	return o.TradeAmount
}

func (o *OtherAmounts31) AddExecutingBrokerAmount() *AmountAndDirection44 {
	o.ExecutingBrokerAmount = new(AmountAndDirection44)
	return o.ExecutingBrokerAmount
}

func (o *OtherAmounts31) AddIssueDiscountAllowance() *AmountAndDirection44 {
	o.IssueDiscountAllowance = new(AmountAndDirection44)
	return o.IssueDiscountAllowance
}

func (o *OtherAmounts31) AddPaymentLevyTax() *AmountAndDirection44 {
	o.PaymentLevyTax = new(AmountAndDirection44)
	return o.PaymentLevyTax
}

func (o *OtherAmounts31) AddLocalTax() *AmountAndDirection44 {
	o.LocalTax = new(AmountAndDirection44)
	return o.LocalTax
}

func (o *OtherAmounts31) AddLocalBrokerCommission() *AmountAndDirection44 {
	o.LocalBrokerCommission = new(AmountAndDirection44)
	return o.LocalBrokerCommission
}

func (o *OtherAmounts31) AddMargin() *AmountAndDirection44 {
	o.Margin = new(AmountAndDirection44)
	return o.Margin
}

func (o *OtherAmounts31) AddOther() *AmountAndDirection44 {
	o.Other = new(AmountAndDirection44)
	return o.Other
}

func (o *OtherAmounts31) AddRegulatoryAmount() *AmountAndDirection44 {
	o.RegulatoryAmount = new(AmountAndDirection44)
	return o.RegulatoryAmount
}

func (o *OtherAmounts31) AddShippingAmount() *AmountAndDirection44 {
	o.ShippingAmount = new(AmountAndDirection44)
	return o.ShippingAmount
}

func (o *OtherAmounts31) AddSpecialConcession() *AmountAndDirection44 {
	o.SpecialConcession = new(AmountAndDirection44)
	return o.SpecialConcession
}

func (o *OtherAmounts31) AddStampDuty() *AmountAndDirection44 {
	o.StampDuty = new(AmountAndDirection44)
	return o.StampDuty
}

func (o *OtherAmounts31) AddStockExchangeTax() *AmountAndDirection44 {
	o.StockExchangeTax = new(AmountAndDirection44)
	return o.StockExchangeTax
}

func (o *OtherAmounts31) AddTransferTax() *AmountAndDirection44 {
	o.TransferTax = new(AmountAndDirection44)
	return o.TransferTax
}

func (o *OtherAmounts31) AddTransactionTax() *AmountAndDirection44 {
	o.TransactionTax = new(AmountAndDirection44)
	return o.TransactionTax
}

func (o *OtherAmounts31) AddValueAddedTax() *AmountAndDirection44 {
	o.ValueAddedTax = new(AmountAndDirection44)
	return o.ValueAddedTax
}

func (o *OtherAmounts31) AddWithholdingTax() *AmountAndDirection44 {
	o.WithholdingTax = new(AmountAndDirection44)
	return o.WithholdingTax
}

func (o *OtherAmounts31) AddNetGainLoss() *AmountAndDirection44 {
	o.NetGainLoss = new(AmountAndDirection44)
	return o.NetGainLoss
}

func (o *OtherAmounts31) AddConsumptionTax() *AmountAndDirection44 {
	o.ConsumptionTax = new(AmountAndDirection44)
	return o.ConsumptionTax
}

func (o *OtherAmounts31) AddAccruedCapitalisationAmount() *AmountAndDirection44 {
	o.AccruedCapitalisationAmount = new(AmountAndDirection44)
	return o.AccruedCapitalisationAmount
}

func (o *OtherAmounts31) AddBookValue() *AmountAndDirection44 {
	o.BookValue = new(AmountAndDirection44)
	return o.BookValue
}

// Identifies other amounts pertaining to the transaction.
type OtherAmounts32 struct {

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection47 `xml:"AcrdIntrstAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection47 `xml:"ChrgsFees,omitempty"`

	// Principal amount of a trade (price multiplied by quantity).
	TradeAmount *AmountAndDirection47 `xml:"TradAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *AmountAndDirection47 `xml:"ExctgBrkrAmt,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection47 `xml:"LclTax,omitempty"`

	// Amount of commission paid to a local broker.
	LocalBrokerCommission *AmountAndDirection47 `xml:"LclBrkrComssn,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection47 `xml:"Othr,omitempty"`

	// Amount of stamp duty.
	StampDuty *AmountAndDirection47 `xml:"StmpDty,omitempty"`

	// Amount of transaction tax.
	TransactionTax *AmountAndDirection47 `xml:"TxTax,omitempty"`

	// Amount of money that will be withheld by a tax authority.
	WithholdingTax *AmountAndDirection47 `xml:"WhldgTax,omitempty"`

	// Amount of consumption tax.
	ConsumptionTax *AmountAndDirection47 `xml:"CsmptnTax,omitempty"`
}

func (o *OtherAmounts32) AddAccruedInterestAmount() *AmountAndDirection47 {
	o.AccruedInterestAmount = new(AmountAndDirection47)
	return o.AccruedInterestAmount
}

func (o *OtherAmounts32) AddChargesFees() *AmountAndDirection47 {
	o.ChargesFees = new(AmountAndDirection47)
	return o.ChargesFees
}

func (o *OtherAmounts32) AddTradeAmount() *AmountAndDirection47 {
	o.TradeAmount = new(AmountAndDirection47)
	return o.TradeAmount
}

func (o *OtherAmounts32) AddExecutingBrokerAmount() *AmountAndDirection47 {
	o.ExecutingBrokerAmount = new(AmountAndDirection47)
	return o.ExecutingBrokerAmount
}

func (o *OtherAmounts32) AddLocalTax() *AmountAndDirection47 {
	o.LocalTax = new(AmountAndDirection47)
	return o.LocalTax
}

func (o *OtherAmounts32) AddLocalBrokerCommission() *AmountAndDirection47 {
	o.LocalBrokerCommission = new(AmountAndDirection47)
	return o.LocalBrokerCommission
}

func (o *OtherAmounts32) AddOther() *AmountAndDirection47 {
	o.Other = new(AmountAndDirection47)
	return o.Other
}

func (o *OtherAmounts32) AddStampDuty() *AmountAndDirection47 {
	o.StampDuty = new(AmountAndDirection47)
	return o.StampDuty
}

func (o *OtherAmounts32) AddTransactionTax() *AmountAndDirection47 {
	o.TransactionTax = new(AmountAndDirection47)
	return o.TransactionTax
}

func (o *OtherAmounts32) AddWithholdingTax() *AmountAndDirection47 {
	o.WithholdingTax = new(AmountAndDirection47)
	return o.WithholdingTax
}

func (o *OtherAmounts32) AddConsumptionTax() *AmountAndDirection47 {
	o.ConsumptionTax = new(AmountAndDirection47)
	return o.ConsumptionTax
}

// Identifies other amounts pertaining to the transaction.
type OtherAmounts33 struct {

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection58 `xml:"AcrdIntrstAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection58 `xml:"ChrgsFees,omitempty"`

	// Amount of country, national or federal tax charged by the jurisdiction in which the account servicer is located.
	CountryNationalFederalTax *AmountAndDirection58 `xml:"CtryNtlFdrlTax,omitempty"`

	// Amount of payment levy tax.
	PaymentLevyTax *AmountAndDirection58 `xml:"PmtLevyTax,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection58 `xml:"LclTax,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection58 `xml:"Othr,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryAmount *AmountAndDirection58 `xml:"RgltryAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingAmount *AmountAndDirection58 `xml:"ShppgAmt,omitempty"`

	// Amount of stamp duty.
	StampDuty *AmountAndDirection58 `xml:"StmpDty,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *AmountAndDirection58 `xml:"StockXchgTax,omitempty"`

	// Amount of tax levied on a transfer of ownership of financial instrument.
	TransferTax *AmountAndDirection58 `xml:"TrfTax,omitempty"`

	// Amount of transaction tax.
	TransactionTax *AmountAndDirection58 `xml:"TxTax,omitempty"`

	// Amount of value-added tax.
	ValueAddedTax *AmountAndDirection58 `xml:"ValAddedTax,omitempty"`

	// Amount of money that will be withheld by a tax authority.
	WithholdingTax *AmountAndDirection58 `xml:"WhldgTax,omitempty"`

	// Amount of consumption tax.
	ConsumptionTax *AmountAndDirection58 `xml:"CsmptnTax,omitempty"`

	// Amount of unpaid interest (on bonds which have defaulted and have subsequently
	// restructured), which is capitalized and added to the original principal amount of the bond.
	AccruedCapitalisationAmount *AmountAndDirection58 `xml:"AcrdCptlstnAmt,omitempty"`
}

func (o *OtherAmounts33) AddAccruedInterestAmount() *AmountAndDirection58 {
	o.AccruedInterestAmount = new(AmountAndDirection58)
	return o.AccruedInterestAmount
}

func (o *OtherAmounts33) AddChargesFees() *AmountAndDirection58 {
	o.ChargesFees = new(AmountAndDirection58)
	return o.ChargesFees
}

func (o *OtherAmounts33) AddCountryNationalFederalTax() *AmountAndDirection58 {
	o.CountryNationalFederalTax = new(AmountAndDirection58)
	return o.CountryNationalFederalTax
}

func (o *OtherAmounts33) AddPaymentLevyTax() *AmountAndDirection58 {
	o.PaymentLevyTax = new(AmountAndDirection58)
	return o.PaymentLevyTax
}

func (o *OtherAmounts33) AddLocalTax() *AmountAndDirection58 {
	o.LocalTax = new(AmountAndDirection58)
	return o.LocalTax
}

func (o *OtherAmounts33) AddOther() *AmountAndDirection58 {
	o.Other = new(AmountAndDirection58)
	return o.Other
}

func (o *OtherAmounts33) AddRegulatoryAmount() *AmountAndDirection58 {
	o.RegulatoryAmount = new(AmountAndDirection58)
	return o.RegulatoryAmount
}

func (o *OtherAmounts33) AddShippingAmount() *AmountAndDirection58 {
	o.ShippingAmount = new(AmountAndDirection58)
	return o.ShippingAmount
}

func (o *OtherAmounts33) AddStampDuty() *AmountAndDirection58 {
	o.StampDuty = new(AmountAndDirection58)
	return o.StampDuty
}

func (o *OtherAmounts33) AddStockExchangeTax() *AmountAndDirection58 {
	o.StockExchangeTax = new(AmountAndDirection58)
	return o.StockExchangeTax
}

func (o *OtherAmounts33) AddTransferTax() *AmountAndDirection58 {
	o.TransferTax = new(AmountAndDirection58)
	return o.TransferTax
}

func (o *OtherAmounts33) AddTransactionTax() *AmountAndDirection58 {
	o.TransactionTax = new(AmountAndDirection58)
	return o.TransactionTax
}

func (o *OtherAmounts33) AddValueAddedTax() *AmountAndDirection58 {
	o.ValueAddedTax = new(AmountAndDirection58)
	return o.ValueAddedTax
}

func (o *OtherAmounts33) AddWithholdingTax() *AmountAndDirection58 {
	o.WithholdingTax = new(AmountAndDirection58)
	return o.WithholdingTax
}

func (o *OtherAmounts33) AddConsumptionTax() *AmountAndDirection58 {
	o.ConsumptionTax = new(AmountAndDirection58)
	return o.ConsumptionTax
}

func (o *OtherAmounts33) AddAccruedCapitalisationAmount() *AmountAndDirection58 {
	o.AccruedCapitalisationAmount = new(AmountAndDirection58)
	return o.AccruedCapitalisationAmount
}

// Identifies other amounts pertaining to the transaction.
type OtherAmounts34 struct {

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection58 `xml:"AcrdIntrstAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection58 `xml:"ChrgsFees,omitempty"`

	// Amount of country, national or federal tax charged by the jurisdiction in which the account servicer is located.
	CountryNationalFederalTax *AmountAndDirection58 `xml:"CtryNtlFdrlTax,omitempty"`

	// Principal amount of a trade (price multiplied by quantity).
	TradeAmount *AmountAndDirection58 `xml:"TradAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *AmountAndDirection58 `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of money defined as a discount on a new issue or on a tranche of an existing issue.
	IssueDiscountAllowance *AmountAndDirection58 `xml:"IsseDscntAllwnc,omitempty"`

	// Amount of payment levy tax.
	PaymentLevyTax *AmountAndDirection58 `xml:"PmtLevyTax,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection58 `xml:"LclTax,omitempty"`

	// Amount of commission paid to a local broker.
	LocalBrokerCommission *AmountAndDirection58 `xml:"LclBrkrComssn,omitempty"`

	// Amount of money deposited by the trading party in a margin account.
	Margin *AmountAndDirection58 `xml:"Mrgn,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection58 `xml:"Othr,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryAmount *AmountAndDirection58 `xml:"RgltryAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingAmount *AmountAndDirection58 `xml:"ShppgAmt,omitempty"`

	// Amount of drawdown or other reduction from or in addition to the deal price.
	SpecialConcession *AmountAndDirection58 `xml:"SpclCncssn,omitempty"`

	// Amount of stamp duty.
	StampDuty *AmountAndDirection58 `xml:"StmpDty,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *AmountAndDirection58 `xml:"StockXchgTax,omitempty"`

	// Amount of tax levied on a transfer of ownership of financial instrument.
	TransferTax *AmountAndDirection58 `xml:"TrfTax,omitempty"`

	// Amount of transaction tax.
	TransactionTax *AmountAndDirection58 `xml:"TxTax,omitempty"`

	// Amount of value-added tax.
	ValueAddedTax *AmountAndDirection58 `xml:"ValAddedTax,omitempty"`

	// Amount of money that will be withheld by a tax authority.
	WithholdingTax *AmountAndDirection58 `xml:"WhldgTax,omitempty"`

	// Amount representing the difference between the cost and the current price of a security. In the context of securities settlement, it is the amount paid or received when the instructions are netted or paired off.
	NetGainLoss *AmountAndDirection58 `xml:"NetGnLoss,omitempty"`

	// Amount of consumption tax.
	ConsumptionTax *AmountAndDirection58 `xml:"CsmptnTax,omitempty"`

	// Amount of unpaid interest (on bonds which have defaulted and have subsequently
	// restructured), which is capitalized and added to the original principal amount of the bond.
	AccruedCapitalisationAmount *AmountAndDirection58 `xml:"AcrdCptlstnAmt,omitempty"`

	// Cost of the securities. May be requested in some countries for tax purposes.
	BookValue *AmountAndDirection58 `xml:"BookVal,omitempty"`
}

func (o *OtherAmounts34) AddAccruedInterestAmount() *AmountAndDirection58 {
	o.AccruedInterestAmount = new(AmountAndDirection58)
	return o.AccruedInterestAmount
}

func (o *OtherAmounts34) AddChargesFees() *AmountAndDirection58 {
	o.ChargesFees = new(AmountAndDirection58)
	return o.ChargesFees
}

func (o *OtherAmounts34) AddCountryNationalFederalTax() *AmountAndDirection58 {
	o.CountryNationalFederalTax = new(AmountAndDirection58)
	return o.CountryNationalFederalTax
}

func (o *OtherAmounts34) AddTradeAmount() *AmountAndDirection58 {
	o.TradeAmount = new(AmountAndDirection58)
	return o.TradeAmount
}

func (o *OtherAmounts34) AddExecutingBrokerAmount() *AmountAndDirection58 {
	o.ExecutingBrokerAmount = new(AmountAndDirection58)
	return o.ExecutingBrokerAmount
}

func (o *OtherAmounts34) AddIssueDiscountAllowance() *AmountAndDirection58 {
	o.IssueDiscountAllowance = new(AmountAndDirection58)
	return o.IssueDiscountAllowance
}

func (o *OtherAmounts34) AddPaymentLevyTax() *AmountAndDirection58 {
	o.PaymentLevyTax = new(AmountAndDirection58)
	return o.PaymentLevyTax
}

func (o *OtherAmounts34) AddLocalTax() *AmountAndDirection58 {
	o.LocalTax = new(AmountAndDirection58)
	return o.LocalTax
}

func (o *OtherAmounts34) AddLocalBrokerCommission() *AmountAndDirection58 {
	o.LocalBrokerCommission = new(AmountAndDirection58)
	return o.LocalBrokerCommission
}

func (o *OtherAmounts34) AddMargin() *AmountAndDirection58 {
	o.Margin = new(AmountAndDirection58)
	return o.Margin
}

func (o *OtherAmounts34) AddOther() *AmountAndDirection58 {
	o.Other = new(AmountAndDirection58)
	return o.Other
}

func (o *OtherAmounts34) AddRegulatoryAmount() *AmountAndDirection58 {
	o.RegulatoryAmount = new(AmountAndDirection58)
	return o.RegulatoryAmount
}

func (o *OtherAmounts34) AddShippingAmount() *AmountAndDirection58 {
	o.ShippingAmount = new(AmountAndDirection58)
	return o.ShippingAmount
}

func (o *OtherAmounts34) AddSpecialConcession() *AmountAndDirection58 {
	o.SpecialConcession = new(AmountAndDirection58)
	return o.SpecialConcession
}

func (o *OtherAmounts34) AddStampDuty() *AmountAndDirection58 {
	o.StampDuty = new(AmountAndDirection58)
	return o.StampDuty
}

func (o *OtherAmounts34) AddStockExchangeTax() *AmountAndDirection58 {
	o.StockExchangeTax = new(AmountAndDirection58)
	return o.StockExchangeTax
}

func (o *OtherAmounts34) AddTransferTax() *AmountAndDirection58 {
	o.TransferTax = new(AmountAndDirection58)
	return o.TransferTax
}

func (o *OtherAmounts34) AddTransactionTax() *AmountAndDirection58 {
	o.TransactionTax = new(AmountAndDirection58)
	return o.TransactionTax
}

func (o *OtherAmounts34) AddValueAddedTax() *AmountAndDirection58 {
	o.ValueAddedTax = new(AmountAndDirection58)
	return o.ValueAddedTax
}

func (o *OtherAmounts34) AddWithholdingTax() *AmountAndDirection58 {
	o.WithholdingTax = new(AmountAndDirection58)
	return o.WithholdingTax
}

func (o *OtherAmounts34) AddNetGainLoss() *AmountAndDirection58 {
	o.NetGainLoss = new(AmountAndDirection58)
	return o.NetGainLoss
}

func (o *OtherAmounts34) AddConsumptionTax() *AmountAndDirection58 {
	o.ConsumptionTax = new(AmountAndDirection58)
	return o.ConsumptionTax
}

func (o *OtherAmounts34) AddAccruedCapitalisationAmount() *AmountAndDirection58 {
	o.AccruedCapitalisationAmount = new(AmountAndDirection58)
	return o.AccruedCapitalisationAmount
}

func (o *OtherAmounts34) AddBookValue() *AmountAndDirection58 {
	o.BookValue = new(AmountAndDirection58)
	return o.BookValue
}

// Identifies other amounts pertaining to the transaction.
type OtherAmounts35 struct {

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection58 `xml:"AcrdIntrstAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection58 `xml:"ChrgsFees,omitempty"`

	// Amount of country, national or federal tax charged by the jurisdiction in which the account servicer is located.
	CountryNationalFederalTax *AmountAndDirection58 `xml:"CtryNtlFdrlTax,omitempty"`

	// Principal amount of a trade (price multiplied by quantity).
	TradeAmount *AmountAndDirection58 `xml:"TradAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *AmountAndDirection58 `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of money defined as a discount on a new issue or on a tranche of an existing issue.
	IssueDiscountAllowance *AmountAndDirection58 `xml:"IsseDscntAllwnc,omitempty"`

	// Amount of payment levy tax.
	PaymentLevyTax *AmountAndDirection58 `xml:"PmtLevyTax,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection58 `xml:"LclTax,omitempty"`

	// Local tax country specific.
	LocalTaxCountrySpecific *AmountAndDirection58 `xml:"LclTaxCtrySpcfc,omitempty"`

	// Amount of commission paid to a local broker.
	LocalBrokerCommission *AmountAndDirection58 `xml:"LclBrkrComssn,omitempty"`

	// Amount of money deposited by the trading party in a margin account.
	Margin *AmountAndDirection58 `xml:"Mrgn,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection58 `xml:"Othr,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryAmount *AmountAndDirection58 `xml:"RgltryAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingAmount *AmountAndDirection58 `xml:"ShppgAmt,omitempty"`

	// Amount of drawdown or other reduction from or in addition to the deal price.
	SpecialConcession *AmountAndDirection58 `xml:"SpclCncssn,omitempty"`

	// Amount of stamp duty.
	StampDuty *AmountAndDirection58 `xml:"StmpDty,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *AmountAndDirection58 `xml:"StockXchgTax,omitempty"`

	// Amount of tax levied on a transfer of ownership of financial instrument.
	TransferTax *AmountAndDirection58 `xml:"TrfTax,omitempty"`

	// Amount of transaction tax.
	TransactionTax *AmountAndDirection58 `xml:"TxTax,omitempty"`

	// Amount of value-added tax.
	ValueAddedTax *AmountAndDirection58 `xml:"ValAddedTax,omitempty"`

	// Amount of money that will be withheld by a tax authority.
	WithholdingTax *AmountAndDirection58 `xml:"WhldgTax,omitempty"`

	// Amount representing the difference between the cost and the current price of a security. In the context of securities settlement, it is the amount paid or received when the instructions are netted or paired off.
	NetGainLoss *AmountAndDirection58 `xml:"NetGnLoss,omitempty"`

	// Amount of consumption tax.
	ConsumptionTax *AmountAndDirection58 `xml:"CsmptnTax,omitempty"`

	// Amount of unpaid interest (on bonds which have defaulted and have subsequently
	// restructured), which is capitalized and added to the original principal amount of the bond.
	AccruedCapitalisationAmount *AmountAndDirection58 `xml:"AcrdCptlstnAmt,omitempty"`
}

func (o *OtherAmounts35) AddAccruedInterestAmount() *AmountAndDirection58 {
	o.AccruedInterestAmount = new(AmountAndDirection58)
	return o.AccruedInterestAmount
}

func (o *OtherAmounts35) AddChargesFees() *AmountAndDirection58 {
	o.ChargesFees = new(AmountAndDirection58)
	return o.ChargesFees
}

func (o *OtherAmounts35) AddCountryNationalFederalTax() *AmountAndDirection58 {
	o.CountryNationalFederalTax = new(AmountAndDirection58)
	return o.CountryNationalFederalTax
}

func (o *OtherAmounts35) AddTradeAmount() *AmountAndDirection58 {
	o.TradeAmount = new(AmountAndDirection58)
	return o.TradeAmount
}

func (o *OtherAmounts35) AddExecutingBrokerAmount() *AmountAndDirection58 {
	o.ExecutingBrokerAmount = new(AmountAndDirection58)
	return o.ExecutingBrokerAmount
}

func (o *OtherAmounts35) AddIssueDiscountAllowance() *AmountAndDirection58 {
	o.IssueDiscountAllowance = new(AmountAndDirection58)
	return o.IssueDiscountAllowance
}

func (o *OtherAmounts35) AddPaymentLevyTax() *AmountAndDirection58 {
	o.PaymentLevyTax = new(AmountAndDirection58)
	return o.PaymentLevyTax
}

func (o *OtherAmounts35) AddLocalTax() *AmountAndDirection58 {
	o.LocalTax = new(AmountAndDirection58)
	return o.LocalTax
}

func (o *OtherAmounts35) AddLocalTaxCountrySpecific() *AmountAndDirection58 {
	o.LocalTaxCountrySpecific = new(AmountAndDirection58)
	return o.LocalTaxCountrySpecific
}

func (o *OtherAmounts35) AddLocalBrokerCommission() *AmountAndDirection58 {
	o.LocalBrokerCommission = new(AmountAndDirection58)
	return o.LocalBrokerCommission
}

func (o *OtherAmounts35) AddMargin() *AmountAndDirection58 {
	o.Margin = new(AmountAndDirection58)
	return o.Margin
}

func (o *OtherAmounts35) AddOther() *AmountAndDirection58 {
	o.Other = new(AmountAndDirection58)
	return o.Other
}

func (o *OtherAmounts35) AddRegulatoryAmount() *AmountAndDirection58 {
	o.RegulatoryAmount = new(AmountAndDirection58)
	return o.RegulatoryAmount
}

func (o *OtherAmounts35) AddShippingAmount() *AmountAndDirection58 {
	o.ShippingAmount = new(AmountAndDirection58)
	return o.ShippingAmount
}

func (o *OtherAmounts35) AddSpecialConcession() *AmountAndDirection58 {
	o.SpecialConcession = new(AmountAndDirection58)
	return o.SpecialConcession
}

func (o *OtherAmounts35) AddStampDuty() *AmountAndDirection58 {
	o.StampDuty = new(AmountAndDirection58)
	return o.StampDuty
}

func (o *OtherAmounts35) AddStockExchangeTax() *AmountAndDirection58 {
	o.StockExchangeTax = new(AmountAndDirection58)
	return o.StockExchangeTax
}

func (o *OtherAmounts35) AddTransferTax() *AmountAndDirection58 {
	o.TransferTax = new(AmountAndDirection58)
	return o.TransferTax
}

func (o *OtherAmounts35) AddTransactionTax() *AmountAndDirection58 {
	o.TransactionTax = new(AmountAndDirection58)
	return o.TransactionTax
}

func (o *OtherAmounts35) AddValueAddedTax() *AmountAndDirection58 {
	o.ValueAddedTax = new(AmountAndDirection58)
	return o.ValueAddedTax
}

func (o *OtherAmounts35) AddWithholdingTax() *AmountAndDirection58 {
	o.WithholdingTax = new(AmountAndDirection58)
	return o.WithholdingTax
}

func (o *OtherAmounts35) AddNetGainLoss() *AmountAndDirection58 {
	o.NetGainLoss = new(AmountAndDirection58)
	return o.NetGainLoss
}

func (o *OtherAmounts35) AddConsumptionTax() *AmountAndDirection58 {
	o.ConsumptionTax = new(AmountAndDirection58)
	return o.ConsumptionTax
}

func (o *OtherAmounts35) AddAccruedCapitalisationAmount() *AmountAndDirection58 {
	o.AccruedCapitalisationAmount = new(AmountAndDirection58)
	return o.AccruedCapitalisationAmount
}

// Identifies other amounts pertaining to the transaction.
type OtherAmounts36 struct {

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection72 `xml:"AcrdIntrstAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection72 `xml:"ChrgsFees,omitempty"`

	// Principal amount of a trade (price multiplied by quantity).
	TradeAmount *AmountAndDirection72 `xml:"TradAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *AmountAndDirection72 `xml:"ExctgBrkrAmt,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection72 `xml:"LclTax,omitempty"`

	// Amount of commission paid to a local broker.
	LocalBrokerCommission *AmountAndDirection72 `xml:"LclBrkrComssn,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection72 `xml:"Othr,omitempty"`

	// Amount of stamp duty.
	StampDuty *AmountAndDirection72 `xml:"StmpDty,omitempty"`

	// Amount of transaction tax.
	TransactionTax *AmountAndDirection72 `xml:"TxTax,omitempty"`

	// Amount of money that will be withheld by a tax authority.
	WithholdingTax *AmountAndDirection72 `xml:"WhldgTax,omitempty"`

	// Amount of consumption tax.
	ConsumptionTax *AmountAndDirection72 `xml:"CsmptnTax,omitempty"`
}

func (o *OtherAmounts36) AddAccruedInterestAmount() *AmountAndDirection72 {
	o.AccruedInterestAmount = new(AmountAndDirection72)
	return o.AccruedInterestAmount
}

func (o *OtherAmounts36) AddChargesFees() *AmountAndDirection72 {
	o.ChargesFees = new(AmountAndDirection72)
	return o.ChargesFees
}

func (o *OtherAmounts36) AddTradeAmount() *AmountAndDirection72 {
	o.TradeAmount = new(AmountAndDirection72)
	return o.TradeAmount
}

func (o *OtherAmounts36) AddExecutingBrokerAmount() *AmountAndDirection72 {
	o.ExecutingBrokerAmount = new(AmountAndDirection72)
	return o.ExecutingBrokerAmount
}

func (o *OtherAmounts36) AddLocalTax() *AmountAndDirection72 {
	o.LocalTax = new(AmountAndDirection72)
	return o.LocalTax
}

func (o *OtherAmounts36) AddLocalBrokerCommission() *AmountAndDirection72 {
	o.LocalBrokerCommission = new(AmountAndDirection72)
	return o.LocalBrokerCommission
}

func (o *OtherAmounts36) AddOther() *AmountAndDirection72 {
	o.Other = new(AmountAndDirection72)
	return o.Other
}

func (o *OtherAmounts36) AddStampDuty() *AmountAndDirection72 {
	o.StampDuty = new(AmountAndDirection72)
	return o.StampDuty
}

func (o *OtherAmounts36) AddTransactionTax() *AmountAndDirection72 {
	o.TransactionTax = new(AmountAndDirection72)
	return o.TransactionTax
}

func (o *OtherAmounts36) AddWithholdingTax() *AmountAndDirection72 {
	o.WithholdingTax = new(AmountAndDirection72)
	return o.WithholdingTax
}

func (o *OtherAmounts36) AddConsumptionTax() *AmountAndDirection72 {
	o.ConsumptionTax = new(AmountAndDirection72)
	return o.ConsumptionTax
}

// Identifies other amounts pertaining to the transaction.
type OtherAmounts38 struct {

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection58 `xml:"AcrdIntrstAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection58 `xml:"ChrgsFees,omitempty"`

	// Amount of country, national or federal tax charged by the jurisdiction in which the account servicer is located.
	CountryNationalFederalTax *AmountAndDirection58 `xml:"CtryNtlFdrlTax,omitempty"`

	// Principal amount of a trade (price multiplied by quantity).
	TradeAmount *AmountAndDirection58 `xml:"TradAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *AmountAndDirection58 `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of money defined as a discount on a new issue or on a tranche of an existing issue.
	IssueDiscountAllowance *AmountAndDirection58 `xml:"IsseDscntAllwnc,omitempty"`

	// Amount of payment levy tax.
	PaymentLevyTax *AmountAndDirection58 `xml:"PmtLevyTax,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection58 `xml:"LclTax,omitempty"`

	// Local tax country specific.
	LocalTaxCountrySpecific *AmountAndDirection58 `xml:"LclTaxCtrySpcfc,omitempty"`

	// Amount of commission paid to a local broker.
	LocalBrokerCommission *AmountAndDirection58 `xml:"LclBrkrComssn,omitempty"`

	// Amount of money deposited by the trading party in a margin account.
	Margin *AmountAndDirection58 `xml:"Mrgn,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection58 `xml:"Othr,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryAmount *AmountAndDirection58 `xml:"RgltryAmt,omitempty"`

	// All costs related to the physical delivery of documents such as stamps, postage, carrier fees, insurances or messenger services.
	ShippingAmount *AmountAndDirection58 `xml:"ShppgAmt,omitempty"`

	// Amount of drawdown or other reduction from or in addition to the deal price.
	SpecialConcession *AmountAndDirection58 `xml:"SpclCncssn,omitempty"`

	// Amount of stamp duty.
	StampDuty *AmountAndDirection58 `xml:"StmpDty,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *AmountAndDirection58 `xml:"StockXchgTax,omitempty"`

	// Amount of tax levied on a transfer of ownership of financial instrument.
	TransferTax *AmountAndDirection58 `xml:"TrfTax,omitempty"`

	// Amount of transaction tax.
	TransactionTax *AmountAndDirection58 `xml:"TxTax,omitempty"`

	// Amount of value-added tax.
	ValueAddedTax *AmountAndDirection58 `xml:"ValAddedTax,omitempty"`

	// Amount of money that will be withheld by a tax authority.
	WithholdingTax *AmountAndDirection58 `xml:"WhldgTax,omitempty"`

	// Amount representing the difference between the cost and the current price of a security. In the context of securities settlement, it is the amount paid or received when the instructions are netted or paired off.
	NetGainLoss *AmountAndDirection58 `xml:"NetGnLoss,omitempty"`

	// Amount of consumption tax.
	ConsumptionTax *AmountAndDirection58 `xml:"CsmptnTax,omitempty"`

	// Amount of unpaid interest (on bonds which have defaulted and have subsequently
	// restructured), which is capitalized and added to the original principal amount of the bond.
	AccruedCapitalisationAmount *AmountAndDirection58 `xml:"AcrdCptlstnAmt,omitempty"`

	// Cost of the securities. May be requested in some countries for tax purposes.
	BookValue *AmountAndDirection58 `xml:"BookVal,omitempty"`

	// Value of the collateral available for the delivery settlement process at the account level.
	CollateralMonitorAmount *AmountAndDirection58 `xml:"CollMntrAmt,omitempty"`
}

func (o *OtherAmounts38) AddAccruedInterestAmount() *AmountAndDirection58 {
	o.AccruedInterestAmount = new(AmountAndDirection58)
	return o.AccruedInterestAmount
}

func (o *OtherAmounts38) AddChargesFees() *AmountAndDirection58 {
	o.ChargesFees = new(AmountAndDirection58)
	return o.ChargesFees
}

func (o *OtherAmounts38) AddCountryNationalFederalTax() *AmountAndDirection58 {
	o.CountryNationalFederalTax = new(AmountAndDirection58)
	return o.CountryNationalFederalTax
}

func (o *OtherAmounts38) AddTradeAmount() *AmountAndDirection58 {
	o.TradeAmount = new(AmountAndDirection58)
	return o.TradeAmount
}

func (o *OtherAmounts38) AddExecutingBrokerAmount() *AmountAndDirection58 {
	o.ExecutingBrokerAmount = new(AmountAndDirection58)
	return o.ExecutingBrokerAmount
}

func (o *OtherAmounts38) AddIssueDiscountAllowance() *AmountAndDirection58 {
	o.IssueDiscountAllowance = new(AmountAndDirection58)
	return o.IssueDiscountAllowance
}

func (o *OtherAmounts38) AddPaymentLevyTax() *AmountAndDirection58 {
	o.PaymentLevyTax = new(AmountAndDirection58)
	return o.PaymentLevyTax
}

func (o *OtherAmounts38) AddLocalTax() *AmountAndDirection58 {
	o.LocalTax = new(AmountAndDirection58)
	return o.LocalTax
}

func (o *OtherAmounts38) AddLocalTaxCountrySpecific() *AmountAndDirection58 {
	o.LocalTaxCountrySpecific = new(AmountAndDirection58)
	return o.LocalTaxCountrySpecific
}

func (o *OtherAmounts38) AddLocalBrokerCommission() *AmountAndDirection58 {
	o.LocalBrokerCommission = new(AmountAndDirection58)
	return o.LocalBrokerCommission
}

func (o *OtherAmounts38) AddMargin() *AmountAndDirection58 {
	o.Margin = new(AmountAndDirection58)
	return o.Margin
}

func (o *OtherAmounts38) AddOther() *AmountAndDirection58 {
	o.Other = new(AmountAndDirection58)
	return o.Other
}

func (o *OtherAmounts38) AddRegulatoryAmount() *AmountAndDirection58 {
	o.RegulatoryAmount = new(AmountAndDirection58)
	return o.RegulatoryAmount
}

func (o *OtherAmounts38) AddShippingAmount() *AmountAndDirection58 {
	o.ShippingAmount = new(AmountAndDirection58)
	return o.ShippingAmount
}

func (o *OtherAmounts38) AddSpecialConcession() *AmountAndDirection58 {
	o.SpecialConcession = new(AmountAndDirection58)
	return o.SpecialConcession
}

func (o *OtherAmounts38) AddStampDuty() *AmountAndDirection58 {
	o.StampDuty = new(AmountAndDirection58)
	return o.StampDuty
}

func (o *OtherAmounts38) AddStockExchangeTax() *AmountAndDirection58 {
	o.StockExchangeTax = new(AmountAndDirection58)
	return o.StockExchangeTax
}

func (o *OtherAmounts38) AddTransferTax() *AmountAndDirection58 {
	o.TransferTax = new(AmountAndDirection58)
	return o.TransferTax
}

func (o *OtherAmounts38) AddTransactionTax() *AmountAndDirection58 {
	o.TransactionTax = new(AmountAndDirection58)
	return o.TransactionTax
}

func (o *OtherAmounts38) AddValueAddedTax() *AmountAndDirection58 {
	o.ValueAddedTax = new(AmountAndDirection58)
	return o.ValueAddedTax
}

func (o *OtherAmounts38) AddWithholdingTax() *AmountAndDirection58 {
	o.WithholdingTax = new(AmountAndDirection58)
	return o.WithholdingTax
}

func (o *OtherAmounts38) AddNetGainLoss() *AmountAndDirection58 {
	o.NetGainLoss = new(AmountAndDirection58)
	return o.NetGainLoss
}

func (o *OtherAmounts38) AddConsumptionTax() *AmountAndDirection58 {
	o.ConsumptionTax = new(AmountAndDirection58)
	return o.ConsumptionTax
}

func (o *OtherAmounts38) AddAccruedCapitalisationAmount() *AmountAndDirection58 {
	o.AccruedCapitalisationAmount = new(AmountAndDirection58)
	return o.AccruedCapitalisationAmount
}

func (o *OtherAmounts38) AddBookValue() *AmountAndDirection58 {
	o.BookValue = new(AmountAndDirection58)
	return o.BookValue
}

func (o *OtherAmounts38) AddCollateralMonitorAmount() *AmountAndDirection58 {
	o.CollateralMonitorAmount = new(AmountAndDirection58)
	return o.CollateralMonitorAmount
}

// Identifies other amounts pertaining to the transaction.
type OtherAmounts4 struct {

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection9 `xml:"AcrdIntrstAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection9 `xml:"ChrgsFees,omitempty"`

	// Amount of country, national or federal tax charged by the jurisdiction in which the account servicer is located.
	CountryNationalFederalTax *AmountAndDirection9 `xml:"CtryNtlFdrlTax,omitempty"`

	// Principal amount of a trade (price multiplied by quantity).
	TradeAmount *AmountAndDirection9 `xml:"TradAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *AmountAndDirection9 `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of money defined as a discount on a new issue or on a tranche of an existing issue.
	IssueDiscountAllowance *AmountAndDirection9 `xml:"IsseDscntAllwnc,omitempty"`

	// Amount of payment levy tax.
	PaymentLevyTax *AmountAndDirection9 `xml:"PmtLevyTax,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection9 `xml:"LclTax,omitempty"`

	// Amount of commission paid to a local broker.
	LocalBrokerCommission *AmountAndDirection9 `xml:"LclBrkrComssn,omitempty"`

	// Amount of money deposited by the trading party in a margin account.
	Margin *AmountAndDirection9 `xml:"Mrgn,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection9 `xml:"Othr,omitempty"`

	// Amount of money paid for delivery by regular post mail.
	PostageAmount *AmountAndDirection9 `xml:"PstgAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryAmount *AmountAndDirection9 `xml:"RgltryAmt,omitempty"`

	// Amount of money (including insurance) paid for delivery by carrier.
	ShippingAmount *AmountAndDirection9 `xml:"ShppgAmt,omitempty"`

	// Amount of drawdown or other reduction from or in addition to the deal price.
	SpecialConcession *AmountAndDirection9 `xml:"SpclCncssn,omitempty"`

	// Amount of stamp duty.
	StampDuty *AmountAndDirection9 `xml:"StmpDty,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *AmountAndDirection9 `xml:"StockXchgTax,omitempty"`

	// Amount of tax levied on a transfer of ownership of financial instrument.
	TransferTax *AmountAndDirection9 `xml:"TrfTax,omitempty"`

	// Amount of transaction tax.
	TransactionTax *AmountAndDirection9 `xml:"TxTax,omitempty"`

	// Amount of value-added tax.
	ValueAddedTax *AmountAndDirection9 `xml:"ValAddedTax,omitempty"`

	// Amount of money that will be withheld by a tax authority.
	WithholdingTax *AmountAndDirection9 `xml:"WhldgTax,omitempty"`

	// Amount representing the difference between the cost and the current price of a security. In the context of securities settlement, it is the amount paid or received when the instructions are netted or paired off.
	NetGainLoss *AmountAndDirection9 `xml:"NetGnLoss,omitempty"`

	// Amount of consumption tax.
	ConsumptionTax *AmountAndDirection9 `xml:"CsmptnTax,omitempty"`

	// Amount of unpaid interest (on bonds which have defaulted and have subsequently
	// restructured), which is capitalized and added to the original principal amount of the bond.
	AccruedCapitalisationAmount *AmountAndDirection9 `xml:"AcrdCptlstnAmt,omitempty"`

	// Cost of the securities. May be requested in some countries for tax purposes.
	BookValue *AmountAndDirection9 `xml:"BookVal,omitempty"`
}

func (o *OtherAmounts4) AddAccruedInterestAmount() *AmountAndDirection9 {
	o.AccruedInterestAmount = new(AmountAndDirection9)
	return o.AccruedInterestAmount
}

func (o *OtherAmounts4) AddChargesFees() *AmountAndDirection9 {
	o.ChargesFees = new(AmountAndDirection9)
	return o.ChargesFees
}

func (o *OtherAmounts4) AddCountryNationalFederalTax() *AmountAndDirection9 {
	o.CountryNationalFederalTax = new(AmountAndDirection9)
	return o.CountryNationalFederalTax
}

func (o *OtherAmounts4) AddTradeAmount() *AmountAndDirection9 {
	o.TradeAmount = new(AmountAndDirection9)
	return o.TradeAmount
}

func (o *OtherAmounts4) AddExecutingBrokerAmount() *AmountAndDirection9 {
	o.ExecutingBrokerAmount = new(AmountAndDirection9)
	return o.ExecutingBrokerAmount
}

func (o *OtherAmounts4) AddIssueDiscountAllowance() *AmountAndDirection9 {
	o.IssueDiscountAllowance = new(AmountAndDirection9)
	return o.IssueDiscountAllowance
}

func (o *OtherAmounts4) AddPaymentLevyTax() *AmountAndDirection9 {
	o.PaymentLevyTax = new(AmountAndDirection9)
	return o.PaymentLevyTax
}

func (o *OtherAmounts4) AddLocalTax() *AmountAndDirection9 {
	o.LocalTax = new(AmountAndDirection9)
	return o.LocalTax
}

func (o *OtherAmounts4) AddLocalBrokerCommission() *AmountAndDirection9 {
	o.LocalBrokerCommission = new(AmountAndDirection9)
	return o.LocalBrokerCommission
}

func (o *OtherAmounts4) AddMargin() *AmountAndDirection9 {
	o.Margin = new(AmountAndDirection9)
	return o.Margin
}

func (o *OtherAmounts4) AddOther() *AmountAndDirection9 {
	o.Other = new(AmountAndDirection9)
	return o.Other
}

func (o *OtherAmounts4) AddPostageAmount() *AmountAndDirection9 {
	o.PostageAmount = new(AmountAndDirection9)
	return o.PostageAmount
}

func (o *OtherAmounts4) AddRegulatoryAmount() *AmountAndDirection9 {
	o.RegulatoryAmount = new(AmountAndDirection9)
	return o.RegulatoryAmount
}

func (o *OtherAmounts4) AddShippingAmount() *AmountAndDirection9 {
	o.ShippingAmount = new(AmountAndDirection9)
	return o.ShippingAmount
}

func (o *OtherAmounts4) AddSpecialConcession() *AmountAndDirection9 {
	o.SpecialConcession = new(AmountAndDirection9)
	return o.SpecialConcession
}

func (o *OtherAmounts4) AddStampDuty() *AmountAndDirection9 {
	o.StampDuty = new(AmountAndDirection9)
	return o.StampDuty
}

func (o *OtherAmounts4) AddStockExchangeTax() *AmountAndDirection9 {
	o.StockExchangeTax = new(AmountAndDirection9)
	return o.StockExchangeTax
}

func (o *OtherAmounts4) AddTransferTax() *AmountAndDirection9 {
	o.TransferTax = new(AmountAndDirection9)
	return o.TransferTax
}

func (o *OtherAmounts4) AddTransactionTax() *AmountAndDirection9 {
	o.TransactionTax = new(AmountAndDirection9)
	return o.TransactionTax
}

func (o *OtherAmounts4) AddValueAddedTax() *AmountAndDirection9 {
	o.ValueAddedTax = new(AmountAndDirection9)
	return o.ValueAddedTax
}

func (o *OtherAmounts4) AddWithholdingTax() *AmountAndDirection9 {
	o.WithholdingTax = new(AmountAndDirection9)
	return o.WithholdingTax
}

func (o *OtherAmounts4) AddNetGainLoss() *AmountAndDirection9 {
	o.NetGainLoss = new(AmountAndDirection9)
	return o.NetGainLoss
}

func (o *OtherAmounts4) AddConsumptionTax() *AmountAndDirection9 {
	o.ConsumptionTax = new(AmountAndDirection9)
	return o.ConsumptionTax
}

func (o *OtherAmounts4) AddAccruedCapitalisationAmount() *AmountAndDirection9 {
	o.AccruedCapitalisationAmount = new(AmountAndDirection9)
	return o.AccruedCapitalisationAmount
}

func (o *OtherAmounts4) AddBookValue() *AmountAndDirection9 {
	o.BookValue = new(AmountAndDirection9)
	return o.BookValue
}

// Identifies other amounts pertaining to the transaction.
type OtherAmounts8 struct {

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection23 `xml:"AcrdIntrstAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection23 `xml:"ChrgsFees,omitempty"`

	// Principal amount of a trade (price multiplied by quantity).
	TradeAmount *AmountAndDirection23 `xml:"TradAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *AmountAndDirection23 `xml:"ExctgBrkrAmt,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection23 `xml:"LclTax,omitempty"`

	// Amount of commission paid to a local broker.
	LocalBrokerCommission *AmountAndDirection23 `xml:"LclBrkrComssn,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection23 `xml:"Othr,omitempty"`

	// Amount of stamp duty.
	StampDuty *AmountAndDirection23 `xml:"StmpDty,omitempty"`

	// Amount of transaction tax.
	TransactionTax *AmountAndDirection23 `xml:"TxTax,omitempty"`

	// Amount of money that will be withheld by a tax authority.
	WithholdingTax *AmountAndDirection23 `xml:"WhldgTax,omitempty"`

	// Amount of consumption tax.
	ConsumptionTax *AmountAndDirection23 `xml:"CsmptnTax,omitempty"`
}

func (o *OtherAmounts8) AddAccruedInterestAmount() *AmountAndDirection23 {
	o.AccruedInterestAmount = new(AmountAndDirection23)
	return o.AccruedInterestAmount
}

func (o *OtherAmounts8) AddChargesFees() *AmountAndDirection23 {
	o.ChargesFees = new(AmountAndDirection23)
	return o.ChargesFees
}

func (o *OtherAmounts8) AddTradeAmount() *AmountAndDirection23 {
	o.TradeAmount = new(AmountAndDirection23)
	return o.TradeAmount
}

func (o *OtherAmounts8) AddExecutingBrokerAmount() *AmountAndDirection23 {
	o.ExecutingBrokerAmount = new(AmountAndDirection23)
	return o.ExecutingBrokerAmount
}

func (o *OtherAmounts8) AddLocalTax() *AmountAndDirection23 {
	o.LocalTax = new(AmountAndDirection23)
	return o.LocalTax
}

func (o *OtherAmounts8) AddLocalBrokerCommission() *AmountAndDirection23 {
	o.LocalBrokerCommission = new(AmountAndDirection23)
	return o.LocalBrokerCommission
}

func (o *OtherAmounts8) AddOther() *AmountAndDirection23 {
	o.Other = new(AmountAndDirection23)
	return o.Other
}

func (o *OtherAmounts8) AddStampDuty() *AmountAndDirection23 {
	o.StampDuty = new(AmountAndDirection23)
	return o.StampDuty
}

func (o *OtherAmounts8) AddTransactionTax() *AmountAndDirection23 {
	o.TransactionTax = new(AmountAndDirection23)
	return o.TransactionTax
}

func (o *OtherAmounts8) AddWithholdingTax() *AmountAndDirection23 {
	o.WithholdingTax = new(AmountAndDirection23)
	return o.WithholdingTax
}

func (o *OtherAmounts8) AddConsumptionTax() *AmountAndDirection23 {
	o.ConsumptionTax = new(AmountAndDirection23)
	return o.ConsumptionTax
}

// Identifies other amounts pertaining to the transaction.
type OtherAmounts9 struct {

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection9 `xml:"AcrdIntrstAmt,omitempty"`

	// Amount of money paid for the provision of financial services that cannot be categorised by another qualifier.
	ChargesFees *AmountAndDirection9 `xml:"ChrgsFees,omitempty"`

	// Amount of country, national or federal tax charged by the jurisdiction in which the account servicer is located.
	CountryNationalFederalTax *AmountAndDirection9 `xml:"CtryNtlFdrlTax,omitempty"`

	// Principal amount of a trade (price multiplied by quantity).
	TradeAmount *AmountAndDirection9 `xml:"TradAmt,omitempty"`

	// Amount of money paid to an executing broker as a commission.
	ExecutingBrokerAmount *AmountAndDirection9 `xml:"ExctgBrkrAmt,omitempty"`

	// Amount of money defined as a discount on a new issue or on a tranche of an existing issue.
	IssueDiscountAllowance *AmountAndDirection9 `xml:"IsseDscntAllwnc,omitempty"`

	// Amount of payment levy tax.
	PaymentLevyTax *AmountAndDirection9 `xml:"PmtLevyTax,omitempty"`

	// Tax charged by the jurisdiction in which the financial instrument settles.
	LocalTax *AmountAndDirection9 `xml:"LclTax,omitempty"`

	// Local tax country specific.
	LocalTaxCountrySpecific *AmountAndDirection9 `xml:"LclTaxCtrySpcfc,omitempty"`

	// Amount of commission paid to a local broker.
	LocalBrokerCommission *AmountAndDirection9 `xml:"LclBrkrComssn,omitempty"`

	// Amount of money deposited by the trading party in a margin account.
	Margin *AmountAndDirection9 `xml:"Mrgn,omitempty"`

	// An amount that is not indicated by a known business denomination.
	Other *AmountAndDirection9 `xml:"Othr,omitempty"`

	// Amount of money paid for delivery by regular post mail.
	PostageAmount *AmountAndDirection9 `xml:"PstgAmt,omitempty"`

	// Amount of money charged by a regulatory authority, for example, Securities and Exchange fees.
	RegulatoryAmount *AmountAndDirection9 `xml:"RgltryAmt,omitempty"`

	// Amount of money (including insurance) paid for delivery by carrier.
	ShippingAmount *AmountAndDirection9 `xml:"ShppgAmt,omitempty"`

	// Amount of drawdown or other reduction from or in addition to the deal price.
	SpecialConcession *AmountAndDirection9 `xml:"SpclCncssn,omitempty"`

	// Amount of stamp duty.
	StampDuty *AmountAndDirection9 `xml:"StmpDty,omitempty"`

	// Amount of stock exchange tax.
	StockExchangeTax *AmountAndDirection9 `xml:"StockXchgTax,omitempty"`

	// Amount of tax levied on a transfer of ownership of financial instrument.
	TransferTax *AmountAndDirection9 `xml:"TrfTax,omitempty"`

	// Amount of transaction tax.
	TransactionTax *AmountAndDirection9 `xml:"TxTax,omitempty"`

	// Amount of value-added tax.
	ValueAddedTax *AmountAndDirection9 `xml:"ValAddedTax,omitempty"`

	// Amount of money that will be withheld by a tax authority.
	WithholdingTax *AmountAndDirection9 `xml:"WhldgTax,omitempty"`

	// Amount representing the difference between the cost and the current price of a security. In the context of securities settlement, it is the amount paid or received when the instructions are netted or paired off.
	NetGainLoss *AmountAndDirection9 `xml:"NetGnLoss,omitempty"`

	// Amount of consumption tax.
	ConsumptionTax *AmountAndDirection9 `xml:"CsmptnTax,omitempty"`

	// Amount of unpaid interest (on bonds which have defaulted and have subsequently
	// restructured), which is capitalized and added to the original principal amount of the bond.
	AccruedCapitalisationAmount *AmountAndDirection9 `xml:"AcrdCptlstnAmt,omitempty"`

	// Cost of the securities. May be requested in some countries for tax purposes.
	BookValue *AmountAndDirection9 `xml:"BookVal,omitempty"`

	// Value of the collateral available for the delivery settlement process at the account level.
	CollateralMonitorAmount *AmountAndDirection9 `xml:"CollMntrAmt,omitempty"`
}

func (o *OtherAmounts9) AddAccruedInterestAmount() *AmountAndDirection9 {
	o.AccruedInterestAmount = new(AmountAndDirection9)
	return o.AccruedInterestAmount
}

func (o *OtherAmounts9) AddChargesFees() *AmountAndDirection9 {
	o.ChargesFees = new(AmountAndDirection9)
	return o.ChargesFees
}

func (o *OtherAmounts9) AddCountryNationalFederalTax() *AmountAndDirection9 {
	o.CountryNationalFederalTax = new(AmountAndDirection9)
	return o.CountryNationalFederalTax
}

func (o *OtherAmounts9) AddTradeAmount() *AmountAndDirection9 {
	o.TradeAmount = new(AmountAndDirection9)
	return o.TradeAmount
}

func (o *OtherAmounts9) AddExecutingBrokerAmount() *AmountAndDirection9 {
	o.ExecutingBrokerAmount = new(AmountAndDirection9)
	return o.ExecutingBrokerAmount
}

func (o *OtherAmounts9) AddIssueDiscountAllowance() *AmountAndDirection9 {
	o.IssueDiscountAllowance = new(AmountAndDirection9)
	return o.IssueDiscountAllowance
}

func (o *OtherAmounts9) AddPaymentLevyTax() *AmountAndDirection9 {
	o.PaymentLevyTax = new(AmountAndDirection9)
	return o.PaymentLevyTax
}

func (o *OtherAmounts9) AddLocalTax() *AmountAndDirection9 {
	o.LocalTax = new(AmountAndDirection9)
	return o.LocalTax
}

func (o *OtherAmounts9) AddLocalTaxCountrySpecific() *AmountAndDirection9 {
	o.LocalTaxCountrySpecific = new(AmountAndDirection9)
	return o.LocalTaxCountrySpecific
}

func (o *OtherAmounts9) AddLocalBrokerCommission() *AmountAndDirection9 {
	o.LocalBrokerCommission = new(AmountAndDirection9)
	return o.LocalBrokerCommission
}

func (o *OtherAmounts9) AddMargin() *AmountAndDirection9 {
	o.Margin = new(AmountAndDirection9)
	return o.Margin
}

func (o *OtherAmounts9) AddOther() *AmountAndDirection9 {
	o.Other = new(AmountAndDirection9)
	return o.Other
}

func (o *OtherAmounts9) AddPostageAmount() *AmountAndDirection9 {
	o.PostageAmount = new(AmountAndDirection9)
	return o.PostageAmount
}

func (o *OtherAmounts9) AddRegulatoryAmount() *AmountAndDirection9 {
	o.RegulatoryAmount = new(AmountAndDirection9)
	return o.RegulatoryAmount
}

func (o *OtherAmounts9) AddShippingAmount() *AmountAndDirection9 {
	o.ShippingAmount = new(AmountAndDirection9)
	return o.ShippingAmount
}

func (o *OtherAmounts9) AddSpecialConcession() *AmountAndDirection9 {
	o.SpecialConcession = new(AmountAndDirection9)
	return o.SpecialConcession
}

func (o *OtherAmounts9) AddStampDuty() *AmountAndDirection9 {
	o.StampDuty = new(AmountAndDirection9)
	return o.StampDuty
}

func (o *OtherAmounts9) AddStockExchangeTax() *AmountAndDirection9 {
	o.StockExchangeTax = new(AmountAndDirection9)
	return o.StockExchangeTax
}

func (o *OtherAmounts9) AddTransferTax() *AmountAndDirection9 {
	o.TransferTax = new(AmountAndDirection9)
	return o.TransferTax
}

func (o *OtherAmounts9) AddTransactionTax() *AmountAndDirection9 {
	o.TransactionTax = new(AmountAndDirection9)
	return o.TransactionTax
}

func (o *OtherAmounts9) AddValueAddedTax() *AmountAndDirection9 {
	o.ValueAddedTax = new(AmountAndDirection9)
	return o.ValueAddedTax
}

func (o *OtherAmounts9) AddWithholdingTax() *AmountAndDirection9 {
	o.WithholdingTax = new(AmountAndDirection9)
	return o.WithholdingTax
}

func (o *OtherAmounts9) AddNetGainLoss() *AmountAndDirection9 {
	o.NetGainLoss = new(AmountAndDirection9)
	return o.NetGainLoss
}

func (o *OtherAmounts9) AddConsumptionTax() *AmountAndDirection9 {
	o.ConsumptionTax = new(AmountAndDirection9)
	return o.ConsumptionTax
}

func (o *OtherAmounts9) AddAccruedCapitalisationAmount() *AmountAndDirection9 {
	o.AccruedCapitalisationAmount = new(AmountAndDirection9)
	return o.AccruedCapitalisationAmount
}

func (o *OtherAmounts9) AddBookValue() *AmountAndDirection9 {
	o.BookValue = new(AmountAndDirection9)
	return o.BookValue
}

func (o *OtherAmounts9) AddCollateralMonitorAmount() *AmountAndDirection9 {
	o.CollateralMonitorAmount = new(AmountAndDirection9)
	return o.CollateralMonitorAmount
}
