package iso20022

// Modification on the value of goods and / or services. For example: rebate, discount, surcharge
type Adjustment3 struct {

	// Specifies the type of adjustment applied to the amount of goods/services by means of a code.
	Type *AdjustmentType2Code `xml:"Tp"`

	// Specifies a type of adjustment not present in the code list.
	OtherAdjustmentType *Max35Text `xml:"OthrAdjstmntTp"`

	// Specifies the monetary amount of the adjustment.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Specifies the percentage rate of the adjustment.
	Rate *PercentageRate `xml:"Rate"`

	// Specifies whether the adjustment must be subtracted or added to the total amount.
	Direction *AdjustmentDirection1Code `xml:"Drctn"`
}

func (a *Adjustment3) SetType(value string) {
	a.Type = (*AdjustmentType2Code)(&value)
}

func (a *Adjustment3) SetOtherAdjustmentType(value string) {
	a.OtherAdjustmentType = (*Max35Text)(&value)
}

func (a *Adjustment3) SetAmount(value, currency string) {
	a.Amount = NewCurrencyAndAmount(value, currency)
}

func (a *Adjustment3) SetRate(value string) {
	a.Rate = (*PercentageRate)(&value)
}

func (a *Adjustment3) SetDirection(value string) {
	a.Direction = (*AdjustmentDirection1Code)(&value)
}

// Modification on the value of goods and / or services. For example: rebate, discount, surcharge
type Adjustment4 struct {

	// Specifies the type of adjustment applied to the amount of goods/services by means of a code.
	Type *AdjustmentType2Code `xml:"Tp"`

	// Specifies a type of adjustment not present in the code list.
	OtherAdjustmentType *Max35Text `xml:"OthrAdjstmntTp"`

	// Specifies whether the adjustment must be subtracted or added to the total amount.
	Direction *AdjustmentDirection1Code `xml:"Drctn"`

	// Specifies the monetary amount of the adjustment.
	Amount *CurrencyAndAmount `xml:"Amt"`
}

func (a *Adjustment4) SetType(value string) {
	a.Type = (*AdjustmentType2Code)(&value)
}

func (a *Adjustment4) SetOtherAdjustmentType(value string) {
	a.OtherAdjustmentType = (*Max35Text)(&value)
}

func (a *Adjustment4) SetDirection(value string) {
	a.Direction = (*AdjustmentDirection1Code)(&value)
}

func (a *Adjustment4) SetAmount(value, currency string) {
	a.Amount = NewCurrencyAndAmount(value, currency)
}

// Modification on the value of goods and / or services. For example: rebate, discount, surcharge
type Adjustment5 struct {

	// Specifies whether the adjustment must be substracted or added to the total amount.
	Direction *AdjustmentDirection1Code `xml:"Drctn"`

	// Specifies the monetary amount of the adjustment.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`
}

func (a *Adjustment5) SetDirection(value string) {
	a.Direction = (*AdjustmentDirection1Code)(&value)
}

func (a *Adjustment5) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

// Modification on the value of goods and / or services. For example: rebate, discount, surcharge
type Adjustment6 struct {

	// Specifies the type of adjustment.
	Type *AdjustmentType1Choice `xml:"Tp"`

	// Specifies whether the adjustment must be subtracted or added to the total amount.
	Direction *AdjustmentDirection1Code `xml:"Drctn"`

	// Specifies the monetary amount of the adjustment.
	Amount *CurrencyAndAmount `xml:"Amt"`
}

func (a *Adjustment6) AddType() *AdjustmentType1Choice {
	a.Type = new(AdjustmentType1Choice)
	return a.Type
}

func (a *Adjustment6) SetDirection(value string) {
	a.Direction = (*AdjustmentDirection1Code)(&value)
}

func (a *Adjustment6) SetAmount(value, currency string) {
	a.Amount = NewCurrencyAndAmount(value, currency)
}

// Modification on the value of goods and / or services. For example: rebate, discount, surcharge
type Adjustment7 struct {

	// Specifies the type of adjustment.
	Type *AdjustmentType1Choice `xml:"Tp"`

	// Specifies the monetary amount or rate of the adjustment.
	AmountOrPercentage *AmountOrPercentage2Choice `xml:"AmtOrPctg"`

	// Specifies whether the adjustment must be subtracted or added to the total amount.
	Direction *AdjustmentDirection1Code `xml:"Drctn"`
}

func (a *Adjustment7) AddType() *AdjustmentType1Choice {
	a.Type = new(AdjustmentType1Choice)
	return a.Type
}

func (a *Adjustment7) AddAmountOrPercentage() *AmountOrPercentage2Choice {
	a.AmountOrPercentage = new(AmountOrPercentage2Choice)
	return a.AmountOrPercentage
}

func (a *Adjustment7) SetDirection(value string) {
	a.Direction = (*AdjustmentDirection1Code)(&value)
}
