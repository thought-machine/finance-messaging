package iso20022

// Specifies security date details.
type SecurityDate1 struct {

	// Date of the posting (credit or debit) to the account.
	PostingDate *DateAndDateTimeChoice `xml:"PstngDt"`

	// Date/time at which securities become available for trading, for example first dealing date.
	AvailableDate *DateFormat6Choice `xml:"AvlblDt,omitempty"`

	// Date/time at which security will assimilate, become fungible, or have the same rights to dividends as the parent issue.
	PariPassuDate *DateFormat6Choice `xml:"PrpssDt,omitempty"`

	// Date/time at which a security will be entitled to a dividend.
	DividendRankingDate *DateFormat6Choice `xml:"DvddRnkgDt,omitempty"`

	// Date/time at which a payment can be made, for example, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat6Choice `xml:"EarlstPmtDt,omitempty"`

	// Date/time at which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat6Choice `xml:"PmtDt,omitempty"`
}

func (s *SecurityDate1) AddPostingDate() *DateAndDateTimeChoice {
	s.PostingDate = new(DateAndDateTimeChoice)
	return s.PostingDate
}

func (s *SecurityDate1) AddAvailableDate() *DateFormat6Choice {
	s.AvailableDate = new(DateFormat6Choice)
	return s.AvailableDate
}

func (s *SecurityDate1) AddPariPassuDate() *DateFormat6Choice {
	s.PariPassuDate = new(DateFormat6Choice)
	return s.PariPassuDate
}

func (s *SecurityDate1) AddDividendRankingDate() *DateFormat6Choice {
	s.DividendRankingDate = new(DateFormat6Choice)
	return s.DividendRankingDate
}

func (s *SecurityDate1) AddEarliestPaymentDate() *DateFormat6Choice {
	s.EarliestPaymentDate = new(DateFormat6Choice)
	return s.EarliestPaymentDate
}

func (s *SecurityDate1) AddPaymentDate() *DateFormat6Choice {
	s.PaymentDate = new(DateFormat6Choice)
	return s.PaymentDate
}

// Specifies security date details.
type SecurityDate11 struct {

	// Date of the posting (credit or debit) to the account.
	PostingDate *DateAndDateTimeChoice `xml:"PstngDt"`

	// Date/time at which securities become available for trading, for example first dealing date.
	AvailableDate *DateFormat31Choice `xml:"AvlblDt,omitempty"`

	// Date/time at which security will assimilate, become fungible, or have the same rights to dividends as the parent issue.
	PariPassuDate *DateFormat31Choice `xml:"PrpssDt,omitempty"`

	// Date/time at which a security will be entitled to a dividend.
	DividendRankingDate *DateFormat31Choice `xml:"DvddRnkgDt,omitempty"`

	// Date/time at which a payment can be made, for example, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat31Choice `xml:"EarlstPmtDt,omitempty"`

	// Date/time at which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat31Choice `xml:"PmtDt,omitempty"`
}

func (s *SecurityDate11) AddPostingDate() *DateAndDateTimeChoice {
	s.PostingDate = new(DateAndDateTimeChoice)
	return s.PostingDate
}

func (s *SecurityDate11) AddAvailableDate() *DateFormat31Choice {
	s.AvailableDate = new(DateFormat31Choice)
	return s.AvailableDate
}

func (s *SecurityDate11) AddPariPassuDate() *DateFormat31Choice {
	s.PariPassuDate = new(DateFormat31Choice)
	return s.PariPassuDate
}

func (s *SecurityDate11) AddDividendRankingDate() *DateFormat31Choice {
	s.DividendRankingDate = new(DateFormat31Choice)
	return s.DividendRankingDate
}

func (s *SecurityDate11) AddEarliestPaymentDate() *DateFormat31Choice {
	s.EarliestPaymentDate = new(DateFormat31Choice)
	return s.EarliestPaymentDate
}

func (s *SecurityDate11) AddPaymentDate() *DateFormat31Choice {
	s.PaymentDate = new(DateFormat31Choice)
	return s.PaymentDate
}

// Specifies security date details.
type SecurityDate12 struct {

	// Date/time at which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat31Choice `xml:"PmtDt"`

	// Date/time at which securities become available for trading, for example first dealing date.
	AvailableDate *DateFormat31Choice `xml:"AvlblDt,omitempty"`

	// Date/time at which a security will be entitled to a dividend.
	DividendRankingDate *DateFormat31Choice `xml:"DvddRnkgDt,omitempty"`

	// Date/time at which a payment can be made, for example, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat31Choice `xml:"EarlstPmtDt,omitempty"`

	// Date/time at which security will assimilate, become fungible, or have the same rights to dividends as the parent issue.
	PariPassuDate *DateFormat31Choice `xml:"PrpssDt,omitempty"`

	// Date/time at which the securities to be reorganised will cease to be tradeable.
	LastTradingDate *DateFormat31Choice `xml:"LastTradgDt,omitempty"`
}

func (s *SecurityDate12) AddPaymentDate() *DateFormat31Choice {
	s.PaymentDate = new(DateFormat31Choice)
	return s.PaymentDate
}

func (s *SecurityDate12) AddAvailableDate() *DateFormat31Choice {
	s.AvailableDate = new(DateFormat31Choice)
	return s.AvailableDate
}

func (s *SecurityDate12) AddDividendRankingDate() *DateFormat31Choice {
	s.DividendRankingDate = new(DateFormat31Choice)
	return s.DividendRankingDate
}

func (s *SecurityDate12) AddEarliestPaymentDate() *DateFormat31Choice {
	s.EarliestPaymentDate = new(DateFormat31Choice)
	return s.EarliestPaymentDate
}

func (s *SecurityDate12) AddPariPassuDate() *DateFormat31Choice {
	s.PariPassuDate = new(DateFormat31Choice)
	return s.PariPassuDate
}

func (s *SecurityDate12) AddLastTradingDate() *DateFormat31Choice {
	s.LastTradingDate = new(DateFormat31Choice)
	return s.LastTradingDate
}

// Specifies security date details.
type SecurityDate13 struct {

	// Date of the posting (credit or debit) to the account.
	PostingDate *DateAndDateTimeChoice `xml:"PstngDt"`

	// Date/time at which securities become available for trading, for example first dealing date.
	AvailableDate *DateFormat34Choice `xml:"AvlblDt,omitempty"`

	// Date/time at which security will assimilate, become fungible, or have the same rights to dividends as the parent issue.
	PariPassuDate *DateFormat34Choice `xml:"PrpssDt,omitempty"`

	// Date/time at which a security will be entitled to a dividend.
	DividendRankingDate *DateFormat34Choice `xml:"DvddRnkgDt,omitempty"`

	// Date/time at which a payment can be made, for example, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat34Choice `xml:"EarlstPmtDt,omitempty"`

	// Date/time at which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat34Choice `xml:"PmtDt,omitempty"`
}

func (s *SecurityDate13) AddPostingDate() *DateAndDateTimeChoice {
	s.PostingDate = new(DateAndDateTimeChoice)
	return s.PostingDate
}

func (s *SecurityDate13) AddAvailableDate() *DateFormat34Choice {
	s.AvailableDate = new(DateFormat34Choice)
	return s.AvailableDate
}

func (s *SecurityDate13) AddPariPassuDate() *DateFormat34Choice {
	s.PariPassuDate = new(DateFormat34Choice)
	return s.PariPassuDate
}

func (s *SecurityDate13) AddDividendRankingDate() *DateFormat34Choice {
	s.DividendRankingDate = new(DateFormat34Choice)
	return s.DividendRankingDate
}

func (s *SecurityDate13) AddEarliestPaymentDate() *DateFormat34Choice {
	s.EarliestPaymentDate = new(DateFormat34Choice)
	return s.EarliestPaymentDate
}

func (s *SecurityDate13) AddPaymentDate() *DateFormat34Choice {
	s.PaymentDate = new(DateFormat34Choice)
	return s.PaymentDate
}

// Specifies security date details.
type SecurityDate14 struct {

	// Date/time at which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat34Choice `xml:"PmtDt"`

	// Date/time at which securities become available for trading, for example first dealing date.
	AvailableDate *DateFormat34Choice `xml:"AvlblDt,omitempty"`

	// Date/time at which a security will be entitled to a dividend.
	DividendRankingDate *DateFormat34Choice `xml:"DvddRnkgDt,omitempty"`

	// Date/time at which a payment can be made, for example, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat34Choice `xml:"EarlstPmtDt,omitempty"`

	// Date/time at which security will assimilate, become fungible, or have the same rights to dividends as the parent issue.
	PariPassuDate *DateFormat34Choice `xml:"PrpssDt,omitempty"`

	// Date/time at which the securities to be reorganised will cease to be tradeable.
	LastTradingDate *DateFormat34Choice `xml:"LastTradgDt,omitempty"`
}

func (s *SecurityDate14) AddPaymentDate() *DateFormat34Choice {
	s.PaymentDate = new(DateFormat34Choice)
	return s.PaymentDate
}

func (s *SecurityDate14) AddAvailableDate() *DateFormat34Choice {
	s.AvailableDate = new(DateFormat34Choice)
	return s.AvailableDate
}

func (s *SecurityDate14) AddDividendRankingDate() *DateFormat34Choice {
	s.DividendRankingDate = new(DateFormat34Choice)
	return s.DividendRankingDate
}

func (s *SecurityDate14) AddEarliestPaymentDate() *DateFormat34Choice {
	s.EarliestPaymentDate = new(DateFormat34Choice)
	return s.EarliestPaymentDate
}

func (s *SecurityDate14) AddPariPassuDate() *DateFormat34Choice {
	s.PariPassuDate = new(DateFormat34Choice)
	return s.PariPassuDate
}

func (s *SecurityDate14) AddLastTradingDate() *DateFormat34Choice {
	s.LastTradingDate = new(DateFormat34Choice)
	return s.LastTradingDate
}

// Specifies security date details.
type SecurityDate2 struct {

	// Date/time at which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat6Choice `xml:"PmtDt"`

	// Date/time at which securities become available for trading, for example first dealing date.
	AvailableDate *DateFormat6Choice `xml:"AvlblDt,omitempty"`

	// Date/time at which a security will be entitled to a dividend.
	DividendRankingDate *DateFormat6Choice `xml:"DvddRnkgDt,omitempty"`

	// Date/time at which a payment can be made, for example, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat6Choice `xml:"EarlstPmtDt,omitempty"`

	// Date/time at which security will assimilate, become fungible, or have the same rights to dividends as the parent issue.
	PariPassuDate *DateFormat6Choice `xml:"PrpssDt,omitempty"`
}

func (s *SecurityDate2) AddPaymentDate() *DateFormat6Choice {
	s.PaymentDate = new(DateFormat6Choice)
	return s.PaymentDate
}

func (s *SecurityDate2) AddAvailableDate() *DateFormat6Choice {
	s.AvailableDate = new(DateFormat6Choice)
	return s.AvailableDate
}

func (s *SecurityDate2) AddDividendRankingDate() *DateFormat6Choice {
	s.DividendRankingDate = new(DateFormat6Choice)
	return s.DividendRankingDate
}

func (s *SecurityDate2) AddEarliestPaymentDate() *DateFormat6Choice {
	s.EarliestPaymentDate = new(DateFormat6Choice)
	return s.EarliestPaymentDate
}

func (s *SecurityDate2) AddPariPassuDate() *DateFormat6Choice {
	s.PariPassuDate = new(DateFormat6Choice)
	return s.PariPassuDate
}

// Specifies security date details.
type SecurityDate5 struct {

	// Date/time at which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat19Choice `xml:"PmtDt"`

	// Date/time at which securities become available for trading, for example first dealing date.
	AvailableDate *DateFormat19Choice `xml:"AvlblDt,omitempty"`

	// Date/time at which a security will be entitled to a dividend.
	DividendRankingDate *DateFormat19Choice `xml:"DvddRnkgDt,omitempty"`

	// Date/time at which a payment can be made, for example, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat19Choice `xml:"EarlstPmtDt,omitempty"`

	// Date/time at which security will assimilate, become fungible, or have the same rights to dividends as the parent issue.
	PariPassuDate *DateFormat19Choice `xml:"PrpssDt,omitempty"`
}

func (s *SecurityDate5) AddPaymentDate() *DateFormat19Choice {
	s.PaymentDate = new(DateFormat19Choice)
	return s.PaymentDate
}

func (s *SecurityDate5) AddAvailableDate() *DateFormat19Choice {
	s.AvailableDate = new(DateFormat19Choice)
	return s.AvailableDate
}

func (s *SecurityDate5) AddDividendRankingDate() *DateFormat19Choice {
	s.DividendRankingDate = new(DateFormat19Choice)
	return s.DividendRankingDate
}

func (s *SecurityDate5) AddEarliestPaymentDate() *DateFormat19Choice {
	s.EarliestPaymentDate = new(DateFormat19Choice)
	return s.EarliestPaymentDate
}

func (s *SecurityDate5) AddPariPassuDate() *DateFormat19Choice {
	s.PariPassuDate = new(DateFormat19Choice)
	return s.PariPassuDate
}

// Specifies security date details.
type SecurityDate6 struct {

	// Date of the posting (credit or debit) to the account.
	PostingDate *DateAndDateTimeChoice `xml:"PstngDt"`

	// Date/time at which securities become available for trading, for example first dealing date.
	AvailableDate *DateFormat19Choice `xml:"AvlblDt,omitempty"`

	// Date/time at which security will assimilate, become fungible, or have the same rights to dividends as the parent issue.
	PariPassuDate *DateFormat19Choice `xml:"PrpssDt,omitempty"`

	// Date/time at which a security will be entitled to a dividend.
	DividendRankingDate *DateFormat19Choice `xml:"DvddRnkgDt,omitempty"`

	// Date/time at which a payment can be made, for example, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat19Choice `xml:"EarlstPmtDt,omitempty"`

	// Date/time at which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat19Choice `xml:"PmtDt,omitempty"`
}

func (s *SecurityDate6) AddPostingDate() *DateAndDateTimeChoice {
	s.PostingDate = new(DateAndDateTimeChoice)
	return s.PostingDate
}

func (s *SecurityDate6) AddAvailableDate() *DateFormat19Choice {
	s.AvailableDate = new(DateFormat19Choice)
	return s.AvailableDate
}

func (s *SecurityDate6) AddPariPassuDate() *DateFormat19Choice {
	s.PariPassuDate = new(DateFormat19Choice)
	return s.PariPassuDate
}

func (s *SecurityDate6) AddDividendRankingDate() *DateFormat19Choice {
	s.DividendRankingDate = new(DateFormat19Choice)
	return s.DividendRankingDate
}

func (s *SecurityDate6) AddEarliestPaymentDate() *DateFormat19Choice {
	s.EarliestPaymentDate = new(DateFormat19Choice)
	return s.EarliestPaymentDate
}

func (s *SecurityDate6) AddPaymentDate() *DateFormat19Choice {
	s.PaymentDate = new(DateFormat19Choice)
	return s.PaymentDate
}

// Specifies security date details.
type SecurityDate9 struct {

	// Date/time at which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat19Choice `xml:"PmtDt"`

	// Date/time at which securities become available for trading, for example first dealing date.
	AvailableDate *DateFormat19Choice `xml:"AvlblDt,omitempty"`

	// Date/time at which a security will be entitled to a dividend.
	DividendRankingDate *DateFormat19Choice `xml:"DvddRnkgDt,omitempty"`

	// Date/time at which a payment can be made, for example, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat19Choice `xml:"EarlstPmtDt,omitempty"`

	// Date/time at which security will assimilate, become fungible, or have the same rights to dividends as the parent issue.
	PariPassuDate *DateFormat19Choice `xml:"PrpssDt,omitempty"`

	// Date/time at which the securities to be reorganised will cease to be tradeable.
	LastTradingDate *DateFormat19Choice `xml:"LastTradgDt,omitempty"`
}

func (s *SecurityDate9) AddPaymentDate() *DateFormat19Choice {
	s.PaymentDate = new(DateFormat19Choice)
	return s.PaymentDate
}

func (s *SecurityDate9) AddAvailableDate() *DateFormat19Choice {
	s.AvailableDate = new(DateFormat19Choice)
	return s.AvailableDate
}

func (s *SecurityDate9) AddDividendRankingDate() *DateFormat19Choice {
	s.DividendRankingDate = new(DateFormat19Choice)
	return s.DividendRankingDate
}

func (s *SecurityDate9) AddEarliestPaymentDate() *DateFormat19Choice {
	s.EarliestPaymentDate = new(DateFormat19Choice)
	return s.EarliestPaymentDate
}

func (s *SecurityDate9) AddPariPassuDate() *DateFormat19Choice {
	s.PariPassuDate = new(DateFormat19Choice)
	return s.PariPassuDate
}

func (s *SecurityDate9) AddLastTradingDate() *DateFormat19Choice {
	s.LastTradingDate = new(DateFormat19Choice)
	return s.LastTradingDate
}
