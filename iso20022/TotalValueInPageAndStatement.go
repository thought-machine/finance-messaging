package iso20022

// Value of total holdings reported.
type TotalValueInPageAndStatement struct {

	// Total value of positions reported in this message.
	TotalHoldingsValueOfPage *ActiveCurrencyAndAmount `xml:"TtlHldgsValOfPg,omitempty"`

	// Total value of positions reported in this statement (a statement may comprise one or more messages).
	TotalHoldingsValueOfStatement *ActiveCurrencyAndAmount `xml:"TtlHldgsValOfStmt"`
}

func (t *TotalValueInPageAndStatement) SetTotalHoldingsValueOfPage(value, currency string) {
	t.TotalHoldingsValueOfPage = NewActiveCurrencyAndAmount(value, currency)
}

func (t *TotalValueInPageAndStatement) SetTotalHoldingsValueOfStatement(value, currency string) {
	t.TotalHoldingsValueOfStatement = NewActiveCurrencyAndAmount(value, currency)
}

// Totals for the value of the holdings reported in the statement or page.
type TotalValueInPageAndStatement1 struct {

	// Total value of positions reported in this message.
	TotalHoldingsValueOfPage *AmountAndDirection6 `xml:"TtlHldgsValOfPg,omitempty"`

	// Total value of positions reported in this statement (a statement may comprise one or more messages).
	TotalHoldingsValueOfStatement *AmountAndDirection6 `xml:"TtlHldgsValOfStmt"`

	// Total book value of positions reported in this statement (a statement may comprise one or more messages).
	TotalBookValueOfStatement *AmountAndDirection6 `xml:"TtlBookValOfStmt,omitempty"`

	// Total value of the holdings eligible for collateral purposes reported in this statement (a statement may comprise one or more messages).
	TotalEligibleCollateralValue *AmountAndDirection6 `xml:"TtlElgblCollVal,omitempty"`
}

func (t *TotalValueInPageAndStatement1) AddTotalHoldingsValueOfPage() *AmountAndDirection6 {
	t.TotalHoldingsValueOfPage = new(AmountAndDirection6)
	return t.TotalHoldingsValueOfPage
}

func (t *TotalValueInPageAndStatement1) AddTotalHoldingsValueOfStatement() *AmountAndDirection6 {
	t.TotalHoldingsValueOfStatement = new(AmountAndDirection6)
	return t.TotalHoldingsValueOfStatement
}

func (t *TotalValueInPageAndStatement1) AddTotalBookValueOfStatement() *AmountAndDirection6 {
	t.TotalBookValueOfStatement = new(AmountAndDirection6)
	return t.TotalBookValueOfStatement
}

func (t *TotalValueInPageAndStatement1) AddTotalEligibleCollateralValue() *AmountAndDirection6 {
	t.TotalEligibleCollateralValue = new(AmountAndDirection6)
	return t.TotalEligibleCollateralValue
}

// Totals for the value of the holdings reported in the statement or page.
type TotalValueInPageAndStatement2 struct {

	// Total value of positions reported in this message.
	TotalHoldingsValueOfPage *AmountAndDirection6 `xml:"TtlHldgsValOfPg,omitempty"`

	// Total value of positions reported in this statement (a statement may comprise one or more messages).
	TotalHoldingsValueOfStatement *AmountAndDirection6 `xml:"TtlHldgsValOfStmt"`

	// Total book value of positions reported in this statement (a statement may comprise one or more messages).
	TotalBookValueOfStatement *AmountAndDirection6 `xml:"TtlBookValOfStmt,omitempty"`
}

func (t *TotalValueInPageAndStatement2) AddTotalHoldingsValueOfPage() *AmountAndDirection6 {
	t.TotalHoldingsValueOfPage = new(AmountAndDirection6)
	return t.TotalHoldingsValueOfPage
}

func (t *TotalValueInPageAndStatement2) AddTotalHoldingsValueOfStatement() *AmountAndDirection6 {
	t.TotalHoldingsValueOfStatement = new(AmountAndDirection6)
	return t.TotalHoldingsValueOfStatement
}

func (t *TotalValueInPageAndStatement2) AddTotalBookValueOfStatement() *AmountAndDirection6 {
	t.TotalBookValueOfStatement = new(AmountAndDirection6)
	return t.TotalBookValueOfStatement
}

// Totals for the value of the holdings reported in the statement or page.
type TotalValueInPageAndStatement3 struct {

	// Total value of positions reported in this message.
	TotalHoldingsValueOfPage *AmountAndDirection14 `xml:"TtlHldgsValOfPg,omitempty"`

	// Total value of positions reported in this statement (a statement may comprise one or more messages).
	TotalHoldingsValueOfStatement *AmountAndDirection14 `xml:"TtlHldgsValOfStmt"`

	// Total book value of positions reported in this statement (a statement may comprise one or more messages).
	TotalBookValueOfStatement *AmountAndDirection14 `xml:"TtlBookValOfStmt,omitempty"`

	// Total value of the holdings eligible for collateral purposes reported in this statement (a statement may comprise one or more messages).
	TotalEligibleCollateralValue *AmountAndDirection14 `xml:"TtlElgblCollVal,omitempty"`
}

func (t *TotalValueInPageAndStatement3) AddTotalHoldingsValueOfPage() *AmountAndDirection14 {
	t.TotalHoldingsValueOfPage = new(AmountAndDirection14)
	return t.TotalHoldingsValueOfPage
}

func (t *TotalValueInPageAndStatement3) AddTotalHoldingsValueOfStatement() *AmountAndDirection14 {
	t.TotalHoldingsValueOfStatement = new(AmountAndDirection14)
	return t.TotalHoldingsValueOfStatement
}

func (t *TotalValueInPageAndStatement3) AddTotalBookValueOfStatement() *AmountAndDirection14 {
	t.TotalBookValueOfStatement = new(AmountAndDirection14)
	return t.TotalBookValueOfStatement
}

func (t *TotalValueInPageAndStatement3) AddTotalEligibleCollateralValue() *AmountAndDirection14 {
	t.TotalEligibleCollateralValue = new(AmountAndDirection14)
	return t.TotalEligibleCollateralValue
}

// Totals for the value of the holdings reported in the statement or page.
type TotalValueInPageAndStatement4 struct {

	// Total value of positions reported in this message.
	TotalHoldingsValueOfPage *AmountAndDirection14 `xml:"TtlHldgsValOfPg,omitempty"`

	// Total value of positions reported in this statement (a statement may comprise one or more messages).
	TotalHoldingsValueOfStatement *AmountAndDirection14 `xml:"TtlHldgsValOfStmt"`

	// Total book value of positions reported in this statement (a statement may comprise one or more messages).
	TotalBookValueOfStatement *AmountAndDirection14 `xml:"TtlBookValOfStmt,omitempty"`
}

func (t *TotalValueInPageAndStatement4) AddTotalHoldingsValueOfPage() *AmountAndDirection14 {
	t.TotalHoldingsValueOfPage = new(AmountAndDirection14)
	return t.TotalHoldingsValueOfPage
}

func (t *TotalValueInPageAndStatement4) AddTotalHoldingsValueOfStatement() *AmountAndDirection14 {
	t.TotalHoldingsValueOfStatement = new(AmountAndDirection14)
	return t.TotalHoldingsValueOfStatement
}

func (t *TotalValueInPageAndStatement4) AddTotalBookValueOfStatement() *AmountAndDirection14 {
	t.TotalBookValueOfStatement = new(AmountAndDirection14)
	return t.TotalBookValueOfStatement
}
