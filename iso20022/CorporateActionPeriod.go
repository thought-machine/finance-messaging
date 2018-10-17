package iso20022

// Specifies periods.
type CorporateActionPeriod1 struct {

	// Period during which the specified option, or all options of the event, remains valid, eg, offer period.
	ActionPeriod *Period1 `xml:"ActnPrd,omitempty"`

	// Period during a take-over where any outstanding equity must be purchased by the take-over company.
	CompulsoryPurchasePeriod *Period1 `xml:"CmplsryPurchsPrd,omitempty"`

	// Period during which the interest rate has been applied.
	InterestPeriod *Period1 `xml:"IntrstPrd,omitempty"`

	// Period during which the security is blocked.
	BlockingPeriod *Period1 `xml:"BlckgPrd,omitempty"`

	// Period during which the price of a security is determined.
	PriceCalculationPeriod *Period1 `xml:"PricClctnPrd,omitempty"`
}

func (c *CorporateActionPeriod1) AddActionPeriod() *Period1 {
	c.ActionPeriod = new(Period1)
	return c.ActionPeriod
}

func (c *CorporateActionPeriod1) AddCompulsoryPurchasePeriod() *Period1 {
	c.CompulsoryPurchasePeriod = new(Period1)
	return c.CompulsoryPurchasePeriod
}

func (c *CorporateActionPeriod1) AddInterestPeriod() *Period1 {
	c.InterestPeriod = new(Period1)
	return c.InterestPeriod
}

func (c *CorporateActionPeriod1) AddBlockingPeriod() *Period1 {
	c.BlockingPeriod = new(Period1)
	return c.BlockingPeriod
}

func (c *CorporateActionPeriod1) AddPriceCalculationPeriod() *Period1 {
	c.PriceCalculationPeriod = new(Period1)
	return c.PriceCalculationPeriod
}

// Specifies periods of a corporate action.
type CorporateActionPeriod10 struct {

	// Period during which the price of a security is determined.
	PriceCalculationPeriod *Period3Choice `xml:"PricClctnPrd,omitempty"`

	// Period during which the interest rate has been applied.
	InterestPeriod *Period3Choice `xml:"IntrstPrd,omitempty"`

	// Period during a take-over where any outstanding equity must be purchased by the take-over company.
	CompulsoryPurchasePeriod *Period3Choice `xml:"CmplsryPurchsPrd,omitempty"`

	// Period during which the security is blocked.
	BlockingPeriod *Period3Choice `xml:"BlckgPrd,omitempty"`

	// Period assigned by the court in a class action. It determines the client's eligible transactions that will be included in the class action and used to determine the resulting entitlement.
	ClaimPeriod *Period3Choice `xml:"ClmPrd,omitempty"`

	// Period defining the last date for which book entry transfers will be accepted and the date on which the suspension will be released and book entry transfer processing will resume.
	DepositorySuspensionPeriodForBookEntryTransfer *Period3Choice `xml:"DpstrySspnsnPrdForBookNtryTrf,omitempty"`

	// Period defining the last date for which deposits, into nominee name, at the agent will be accepted and the date on which the suspension will be released and deposits at agent will resume.
	DepositorySuspensionPeriodForDepositAtAgent *Period3Choice `xml:"DpstrySspnsnPrdForDpstAtAgt,omitempty"`

	// Period defining the last date for which deposits will be accepted and the date on which the suspension will be released and deposits will resume.
	DepositorySuspensionPeriodForDeposit *Period3Choice `xml:"DpstrySspnsnPrdForDpst,omitempty"`

	// Period defining the last date for which pledges will be accepted and the date on which the suspension will be released and pledge processing will resume.
	DepositorySuspensionPeriodForPledge *Period3Choice `xml:"DpstrySspnsnPrdForPldg,omitempty"`

	// Period defining the last date for which intra-position balances can be segregated and the date on which the suspension will be released and the ability to segregate intra-position balances will resume.
	DepositorySuspensionPeriodForSegregation *Period3Choice `xml:"DpstrySspnsnPrdForSgrtn,omitempty"`

	// Period defining the last date for which withdrawals, from nominee name at the agent will be accepted and the date on which the suspension will be released and withdrawals at agent processing will resume.
	DepositorySuspensionPeriodForWithdrawalAtAgent *Period3Choice `xml:"DpstrySspnsnPrdForWdrwlAtAgt,omitempty"`

	// Period defining the last date for which physical withdrawals in the nominee's name will be accepted and the date on which the suspension will be released and physical withdrawals in the nominee's name will resume.
	DepositorySuspensionPeriodForWithdrawalInNomineeName *Period3Choice `xml:"DpstrySspnsnPrdForWdrwlInNmneeNm,omitempty"`

	// Period defining the last date on which withdrawal requests in street name's will be accepted on the event security and the date on which the suspension will be released and withdrawal in street name's processing on the event security will resume.
	DepositorySuspensionPeriodForWithdrawalInStreetName *Period3Choice `xml:"DpstrySspnsnPrdForWdrwlInStrtNm,omitempty"`

	// Period defining the last date on which shareholder registration will be accepted by the issuer and the date on which shareholder registration will resume.
	BookClosurePeriod *Period3Choice `xml:"BookClsrPrd,omitempty"`

	// Period during which the settlement activities at the co-depositories are suspended in order to stabilise the holdings at the CSD.
	CoDepositoriesSuspensionPeriod *Period3Choice `xml:"CoDpstriesSspnsnPrd,omitempty"`

	// Period during which a physical certificate can be split.
	SplitPeriod *Period3Choice `xml:"SpltPrd,omitempty"`
}

func (c *CorporateActionPeriod10) AddPriceCalculationPeriod() *Period3Choice {
	c.PriceCalculationPeriod = new(Period3Choice)
	return c.PriceCalculationPeriod
}

func (c *CorporateActionPeriod10) AddInterestPeriod() *Period3Choice {
	c.InterestPeriod = new(Period3Choice)
	return c.InterestPeriod
}

func (c *CorporateActionPeriod10) AddCompulsoryPurchasePeriod() *Period3Choice {
	c.CompulsoryPurchasePeriod = new(Period3Choice)
	return c.CompulsoryPurchasePeriod
}

func (c *CorporateActionPeriod10) AddBlockingPeriod() *Period3Choice {
	c.BlockingPeriod = new(Period3Choice)
	return c.BlockingPeriod
}

func (c *CorporateActionPeriod10) AddClaimPeriod() *Period3Choice {
	c.ClaimPeriod = new(Period3Choice)
	return c.ClaimPeriod
}

func (c *CorporateActionPeriod10) AddDepositorySuspensionPeriodForBookEntryTransfer() *Period3Choice {
	c.DepositorySuspensionPeriodForBookEntryTransfer = new(Period3Choice)
	return c.DepositorySuspensionPeriodForBookEntryTransfer
}

func (c *CorporateActionPeriod10) AddDepositorySuspensionPeriodForDepositAtAgent() *Period3Choice {
	c.DepositorySuspensionPeriodForDepositAtAgent = new(Period3Choice)
	return c.DepositorySuspensionPeriodForDepositAtAgent
}

func (c *CorporateActionPeriod10) AddDepositorySuspensionPeriodForDeposit() *Period3Choice {
	c.DepositorySuspensionPeriodForDeposit = new(Period3Choice)
	return c.DepositorySuspensionPeriodForDeposit
}

func (c *CorporateActionPeriod10) AddDepositorySuspensionPeriodForPledge() *Period3Choice {
	c.DepositorySuspensionPeriodForPledge = new(Period3Choice)
	return c.DepositorySuspensionPeriodForPledge
}

func (c *CorporateActionPeriod10) AddDepositorySuspensionPeriodForSegregation() *Period3Choice {
	c.DepositorySuspensionPeriodForSegregation = new(Period3Choice)
	return c.DepositorySuspensionPeriodForSegregation
}

func (c *CorporateActionPeriod10) AddDepositorySuspensionPeriodForWithdrawalAtAgent() *Period3Choice {
	c.DepositorySuspensionPeriodForWithdrawalAtAgent = new(Period3Choice)
	return c.DepositorySuspensionPeriodForWithdrawalAtAgent
}

func (c *CorporateActionPeriod10) AddDepositorySuspensionPeriodForWithdrawalInNomineeName() *Period3Choice {
	c.DepositorySuspensionPeriodForWithdrawalInNomineeName = new(Period3Choice)
	return c.DepositorySuspensionPeriodForWithdrawalInNomineeName
}

func (c *CorporateActionPeriod10) AddDepositorySuspensionPeriodForWithdrawalInStreetName() *Period3Choice {
	c.DepositorySuspensionPeriodForWithdrawalInStreetName = new(Period3Choice)
	return c.DepositorySuspensionPeriodForWithdrawalInStreetName
}

func (c *CorporateActionPeriod10) AddBookClosurePeriod() *Period3Choice {
	c.BookClosurePeriod = new(Period3Choice)
	return c.BookClosurePeriod
}

func (c *CorporateActionPeriod10) AddCoDepositoriesSuspensionPeriod() *Period3Choice {
	c.CoDepositoriesSuspensionPeriod = new(Period3Choice)
	return c.CoDepositoriesSuspensionPeriod
}

func (c *CorporateActionPeriod10) AddSplitPeriod() *Period3Choice {
	c.SplitPeriod = new(Period3Choice)
	return c.SplitPeriod
}

// Specifies periods.
type CorporateActionPeriod11 struct {

	// Period during which the price of a security is determined.
	PriceCalculationPeriod *Period4 `xml:"PricClctnPrd,omitempty"`

	// Period during which the specified option, or all options of the event, remains valid, for example, offer period.
	ActionPeriod *Period4 `xml:"ActnPrd,omitempty"`

	// Period during which both old and new equity may be traded simultaneously, for example, consolidation of equity or splitting of equity.
	ParallelTradingPeriod *Period4 `xml:"ParllTradgPrd,omitempty"`
}

func (c *CorporateActionPeriod11) AddPriceCalculationPeriod() *Period4 {
	c.PriceCalculationPeriod = new(Period4)
	return c.PriceCalculationPeriod
}

func (c *CorporateActionPeriod11) AddActionPeriod() *Period4 {
	c.ActionPeriod = new(Period4)
	return c.ActionPeriod
}

func (c *CorporateActionPeriod11) AddParallelTradingPeriod() *Period4 {
	c.ParallelTradingPeriod = new(Period4)
	return c.ParallelTradingPeriod
}

// Specifies periods.
type CorporateActionPeriod2 struct {

	// Period during which the assented line is available.
	AssentedLinePeriod *Period1 `xml:"AssntdLinePrd,omitempty"`

	// Period during which the specified option, or all options of the event, remains valid, eg, offer period.
	ActionPeriod *Period1 `xml:"ActnPrd,omitempty"`

	// Period during which the privilege is not available, eg, this can happen whenever a meeting takes place or whenever a coupon payment is due.
	PrivilegeSuspensionPeriod *Period1 `xml:"PrvlgSspnsnPrd,omitempty"`

	// Period during which both old and new equity may be traded simultaneously, eg, consolidation of equity or splitting of equity.
	ParallelTradingPeriod *Period1 `xml:"ParllTradgPrd,omitempty"`

	// Period (last day included) during which an account owner can surrender or sell securities to the issuer and receive the sale proceeds.
	SellThruIssuerPeriod *Period1 `xml:"SellThruIssrPrd,omitempty"`

	// Period during which the shareholder can revoke, change or withdraw its instruction.
	RevocabilityPeriod *Period1 `xml:"RvcbltyPrd,omitempty"`

	// Period during which the price of a security is determined (for  outturn securities).
	PriceCalculationPeriod *Period1 `xml:"PricClctnPrd,omitempty"`
}

func (c *CorporateActionPeriod2) AddAssentedLinePeriod() *Period1 {
	c.AssentedLinePeriod = new(Period1)
	return c.AssentedLinePeriod
}

func (c *CorporateActionPeriod2) AddActionPeriod() *Period1 {
	c.ActionPeriod = new(Period1)
	return c.ActionPeriod
}

func (c *CorporateActionPeriod2) AddPrivilegeSuspensionPeriod() *Period1 {
	c.PrivilegeSuspensionPeriod = new(Period1)
	return c.PrivilegeSuspensionPeriod
}

func (c *CorporateActionPeriod2) AddParallelTradingPeriod() *Period1 {
	c.ParallelTradingPeriod = new(Period1)
	return c.ParallelTradingPeriod
}

func (c *CorporateActionPeriod2) AddSellThruIssuerPeriod() *Period1 {
	c.SellThruIssuerPeriod = new(Period1)
	return c.SellThruIssuerPeriod
}

func (c *CorporateActionPeriod2) AddRevocabilityPeriod() *Period1 {
	c.RevocabilityPeriod = new(Period1)
	return c.RevocabilityPeriod
}

func (c *CorporateActionPeriod2) AddPriceCalculationPeriod() *Period1 {
	c.PriceCalculationPeriod = new(Period1)
	return c.PriceCalculationPeriod
}

// Specifies periods of a corporate action.
type CorporateActionPeriod3 struct {

	// Period during which the price of a security is determined.
	PriceCalculationPeriod *Period1Choice `xml:"PricClctnPrd,omitempty"`

	// Period during which the interest rate has been applied.
	InterestPeriod *Period1Choice `xml:"IntrstPrd,omitempty"`

	// Period during a take-over where any outstanding equity must be purchased by the take-over company.
	CompulsoryPurchasePeriod *Period1Choice `xml:"CmplsryPurchsPrd,omitempty"`

	// Period during which the security is blocked.
	BlockingPeriod *Period1Choice `xml:"BlckgPrd,omitempty"`

	// Period assigned by the court in a class action. It determines the client's eligible transactions that will be included in the class action and used to determine the resulting entitlement.
	ClaimPeriod *Period1Choice `xml:"ClmPrd,omitempty"`

	// Period defining the last date for which book entry transfers will be accepted and the date on which the suspension will be released and book entry transfer processing will resume.
	DepositorySuspensionPeriodForBookEntryTransfer *Period1Choice `xml:"DpstrySspnsnPrdForBookNtryTrf,omitempty"`

	// Period defining the last date for which deposits, into nominee name, at the agent will be accepted and the date on which the suspension will be released and deposits at agent will resume.
	DepositorySuspensionPeriodForDepositAtAgent *Period1Choice `xml:"DpstrySspnsnPrdForDpstAtAgt,omitempty"`

	// Period defining the last date for which deposits will be accepted and the date on which the suspension will be released and deposits will resume.
	DepositorySuspensionPeriodForDeposit *Period1Choice `xml:"DpstrySspnsnPrdForDpst,omitempty"`

	// Period defining the last date for which pledges will be accepted and the date on which the suspension will be released and pledge processing will resume.
	DepositorySuspensionPeriodForPledge *Period1Choice `xml:"DpstrySspnsnPrdForPldg,omitempty"`

	// Period defining the last date for which intra-position balances can be segregated and the date on which the suspension will be released and the ability to segregate intra-position balances will resume.
	DepositorySuspensionPeriodForSegregation *Period1Choice `xml:"DpstrySspnsnPrdForSgrtn,omitempty"`

	// Period defining the last date for which withdrawals, from nominee name at the agent will be accepted and the date on which the suspension will be released and withdrawals at agent processing will resume.
	DepositorySuspensionPeriodForWithdrawalAtAgent *Period1Choice `xml:"DpstrySspnsnPrdForWdrwlAtAgt,omitempty"`

	// Period defining the last date for which physical withdrawals in the nominee's name will be accepted and the date on which the suspension will be released and physical withdrawals in the nominee's name will resume.
	DepositorySuspensionPeriodForWithdrawalInNomineeName *Period1Choice `xml:"DpstrySspnsnPrdForWdrwlInNmneeNm,omitempty"`

	// Period defining the last date on which withdrawal requests in street name's will be accepted on the event security and the date on which the suspension will be released and withdrawal in street name's processing on the event security will resume.
	DepositorySuspensionPeriodForWithdrawalInStreetName *Period1Choice `xml:"DpstrySspnsnPrdForWdrwlInStrtNm,omitempty"`

	// Period defining the last date on which shareholder registration will be accepted by the issuer and the date on which shareholder registration will resume.
	BookClosurePeriod *Period1Choice `xml:"BookClsrPrd,omitempty"`
}

func (c *CorporateActionPeriod3) AddPriceCalculationPeriod() *Period1Choice {
	c.PriceCalculationPeriod = new(Period1Choice)
	return c.PriceCalculationPeriod
}

func (c *CorporateActionPeriod3) AddInterestPeriod() *Period1Choice {
	c.InterestPeriod = new(Period1Choice)
	return c.InterestPeriod
}

func (c *CorporateActionPeriod3) AddCompulsoryPurchasePeriod() *Period1Choice {
	c.CompulsoryPurchasePeriod = new(Period1Choice)
	return c.CompulsoryPurchasePeriod
}

func (c *CorporateActionPeriod3) AddBlockingPeriod() *Period1Choice {
	c.BlockingPeriod = new(Period1Choice)
	return c.BlockingPeriod
}

func (c *CorporateActionPeriod3) AddClaimPeriod() *Period1Choice {
	c.ClaimPeriod = new(Period1Choice)
	return c.ClaimPeriod
}

func (c *CorporateActionPeriod3) AddDepositorySuspensionPeriodForBookEntryTransfer() *Period1Choice {
	c.DepositorySuspensionPeriodForBookEntryTransfer = new(Period1Choice)
	return c.DepositorySuspensionPeriodForBookEntryTransfer
}

func (c *CorporateActionPeriod3) AddDepositorySuspensionPeriodForDepositAtAgent() *Period1Choice {
	c.DepositorySuspensionPeriodForDepositAtAgent = new(Period1Choice)
	return c.DepositorySuspensionPeriodForDepositAtAgent
}

func (c *CorporateActionPeriod3) AddDepositorySuspensionPeriodForDeposit() *Period1Choice {
	c.DepositorySuspensionPeriodForDeposit = new(Period1Choice)
	return c.DepositorySuspensionPeriodForDeposit
}

func (c *CorporateActionPeriod3) AddDepositorySuspensionPeriodForPledge() *Period1Choice {
	c.DepositorySuspensionPeriodForPledge = new(Period1Choice)
	return c.DepositorySuspensionPeriodForPledge
}

func (c *CorporateActionPeriod3) AddDepositorySuspensionPeriodForSegregation() *Period1Choice {
	c.DepositorySuspensionPeriodForSegregation = new(Period1Choice)
	return c.DepositorySuspensionPeriodForSegregation
}

func (c *CorporateActionPeriod3) AddDepositorySuspensionPeriodForWithdrawalAtAgent() *Period1Choice {
	c.DepositorySuspensionPeriodForWithdrawalAtAgent = new(Period1Choice)
	return c.DepositorySuspensionPeriodForWithdrawalAtAgent
}

func (c *CorporateActionPeriod3) AddDepositorySuspensionPeriodForWithdrawalInNomineeName() *Period1Choice {
	c.DepositorySuspensionPeriodForWithdrawalInNomineeName = new(Period1Choice)
	return c.DepositorySuspensionPeriodForWithdrawalInNomineeName
}

func (c *CorporateActionPeriod3) AddDepositorySuspensionPeriodForWithdrawalInStreetName() *Period1Choice {
	c.DepositorySuspensionPeriodForWithdrawalInStreetName = new(Period1Choice)
	return c.DepositorySuspensionPeriodForWithdrawalInStreetName
}

func (c *CorporateActionPeriod3) AddBookClosurePeriod() *Period1Choice {
	c.BookClosurePeriod = new(Period1Choice)
	return c.BookClosurePeriod
}

// Specifies periods.
type CorporateActionPeriod4 struct {

	// Period during which the price of a security is determined.
	PriceCalculationPeriod *Period3 `xml:"PricClctnPrd,omitempty"`

	// Period during which the specified option, or all options of the event, remains valid, for example, offer period.
	ActionPeriod *Period3 `xml:"ActnPrd,omitempty"`

	// Period during which both old and new equity may be traded simultaneously, for example, consolidation of equity or splitting of equity.
	ParallelTradingPeriod *Period3 `xml:"ParllTradgPrd,omitempty"`
}

func (c *CorporateActionPeriod4) AddPriceCalculationPeriod() *Period3 {
	c.PriceCalculationPeriod = new(Period3)
	return c.PriceCalculationPeriod
}

func (c *CorporateActionPeriod4) AddActionPeriod() *Period3 {
	c.ActionPeriod = new(Period3)
	return c.ActionPeriod
}

func (c *CorporateActionPeriod4) AddParallelTradingPeriod() *Period3 {
	c.ParallelTradingPeriod = new(Period3)
	return c.ParallelTradingPeriod
}

// Specifies periods related to a corporate action option.
type CorporateActionPeriod5 struct {

	// Period during which the price of a security is determined.
	PriceCalculationPeriod *Period1Choice `xml:"PricClctnPrd,omitempty"`

	// Period during which both old and new equity may be traded simultaneously, for example, consolidation of equity or splitting of equity.
	ParallelTradingPeriod *Period1Choice `xml:"ParllTradgPrd,omitempty"`

	// Period during which the specified option, or all options of the event, remains valid, for example, offer period.
	ActionPeriod *Period1Choice `xml:"ActnPrd,omitempty"`

	// Period during which the shareholder can revoke, change or withdraw its instruction.
	RevocabilityPeriod *Period1Choice `xml:"RvcbltyPrd,omitempty"`

	// Period during which the privilege is not available, for example, this can happen whenever a meeting takes place or whenever a coupon payment is due.
	PrivilegeSuspensionPeriod *Period1Choice `xml:"PrvlgSspnsnPrd,omitempty"`

	// Period during which the participant of the account servicer can revoke change or withdraw its instructions.
	AccountServicerRevocabilityPeriod *Period1Choice `xml:"AcctSvcrRvcbltyPrd,omitempty"`

	// Period defining the last date on which withdrawal in street name requests on the outturn security will be accepted and the date on which the suspension will be released and withdrawal by transfer processing on the outturn security will resume.
	DepositorySuspensionPeriodForWithdrawal *Period1Choice `xml:"DpstrySspnsnPrdForWdrwl,omitempty"`
}

func (c *CorporateActionPeriod5) AddPriceCalculationPeriod() *Period1Choice {
	c.PriceCalculationPeriod = new(Period1Choice)
	return c.PriceCalculationPeriod
}

func (c *CorporateActionPeriod5) AddParallelTradingPeriod() *Period1Choice {
	c.ParallelTradingPeriod = new(Period1Choice)
	return c.ParallelTradingPeriod
}

func (c *CorporateActionPeriod5) AddActionPeriod() *Period1Choice {
	c.ActionPeriod = new(Period1Choice)
	return c.ActionPeriod
}

func (c *CorporateActionPeriod5) AddRevocabilityPeriod() *Period1Choice {
	c.RevocabilityPeriod = new(Period1Choice)
	return c.RevocabilityPeriod
}

func (c *CorporateActionPeriod5) AddPrivilegeSuspensionPeriod() *Period1Choice {
	c.PrivilegeSuspensionPeriod = new(Period1Choice)
	return c.PrivilegeSuspensionPeriod
}

func (c *CorporateActionPeriod5) AddAccountServicerRevocabilityPeriod() *Period1Choice {
	c.AccountServicerRevocabilityPeriod = new(Period1Choice)
	return c.AccountServicerRevocabilityPeriod
}

func (c *CorporateActionPeriod5) AddDepositorySuspensionPeriodForWithdrawal() *Period1Choice {
	c.DepositorySuspensionPeriodForWithdrawal = new(Period1Choice)
	return c.DepositorySuspensionPeriodForWithdrawal
}

// Specifies periods of a corporate action.
type CorporateActionPeriod6 struct {

	// Period during which the price of a security is determined.
	PriceCalculationPeriod *Period3Choice `xml:"PricClctnPrd,omitempty"`

	// Period during which the interest rate has been applied.
	InterestPeriod *Period3Choice `xml:"IntrstPrd,omitempty"`

	// Period during a take-over where any outstanding equity must be purchased by the take-over company.
	CompulsoryPurchasePeriod *Period3Choice `xml:"CmplsryPurchsPrd,omitempty"`

	// Period during which the security is blocked.
	BlockingPeriod *Period3Choice `xml:"BlckgPrd,omitempty"`

	// Period assigned by the court in a class action. It determines the client's eligible transactions that will be included in the class action and used to determine the resulting entitlement.
	ClaimPeriod *Period3Choice `xml:"ClmPrd,omitempty"`

	// Period defining the last date for which book entry transfers will be accepted and the date on which the suspension will be released and book entry transfer processing will resume.
	DepositorySuspensionPeriodForBookEntryTransfer *Period3Choice `xml:"DpstrySspnsnPrdForBookNtryTrf,omitempty"`

	// Period defining the last date for which deposits, into nominee name, at the agent will be accepted and the date on which the suspension will be released and deposits at agent will resume.
	DepositorySuspensionPeriodForDepositAtAgent *Period3Choice `xml:"DpstrySspnsnPrdForDpstAtAgt,omitempty"`

	// Period defining the last date for which deposits will be accepted and the date on which the suspension will be released and deposits will resume.
	DepositorySuspensionPeriodForDeposit *Period3Choice `xml:"DpstrySspnsnPrdForDpst,omitempty"`

	// Period defining the last date for which pledges will be accepted and the date on which the suspension will be released and pledge processing will resume.
	DepositorySuspensionPeriodForPledge *Period3Choice `xml:"DpstrySspnsnPrdForPldg,omitempty"`

	// Period defining the last date for which intra-position balances can be segregated and the date on which the suspension will be released and the ability to segregate intra-position balances will resume.
	DepositorySuspensionPeriodForSegregation *Period3Choice `xml:"DpstrySspnsnPrdForSgrtn,omitempty"`

	// Period defining the last date for which withdrawals, from nominee name at the agent will be accepted and the date on which the suspension will be released and withdrawals at agent processing will resume.
	DepositorySuspensionPeriodForWithdrawalAtAgent *Period3Choice `xml:"DpstrySspnsnPrdForWdrwlAtAgt,omitempty"`

	// Period defining the last date for which physical withdrawals in the nominee's name will be accepted and the date on which the suspension will be released and physical withdrawals in the nominee's name will resume.
	DepositorySuspensionPeriodForWithdrawalInNomineeName *Period3Choice `xml:"DpstrySspnsnPrdForWdrwlInNmneeNm,omitempty"`

	// Period defining the last date on which withdrawal requests in street name's will be accepted on the event security and the date on which the suspension will be released and withdrawal in street name's processing on the event security will resume.
	DepositorySuspensionPeriodForWithdrawalInStreetName *Period3Choice `xml:"DpstrySspnsnPrdForWdrwlInStrtNm,omitempty"`

	// Period defining the last date on which shareholder registration will be accepted by the issuer and the date on which shareholder registration will resume.
	BookClosurePeriod *Period3Choice `xml:"BookClsrPrd,omitempty"`
}

func (c *CorporateActionPeriod6) AddPriceCalculationPeriod() *Period3Choice {
	c.PriceCalculationPeriod = new(Period3Choice)
	return c.PriceCalculationPeriod
}

func (c *CorporateActionPeriod6) AddInterestPeriod() *Period3Choice {
	c.InterestPeriod = new(Period3Choice)
	return c.InterestPeriod
}

func (c *CorporateActionPeriod6) AddCompulsoryPurchasePeriod() *Period3Choice {
	c.CompulsoryPurchasePeriod = new(Period3Choice)
	return c.CompulsoryPurchasePeriod
}

func (c *CorporateActionPeriod6) AddBlockingPeriod() *Period3Choice {
	c.BlockingPeriod = new(Period3Choice)
	return c.BlockingPeriod
}

func (c *CorporateActionPeriod6) AddClaimPeriod() *Period3Choice {
	c.ClaimPeriod = new(Period3Choice)
	return c.ClaimPeriod
}

func (c *CorporateActionPeriod6) AddDepositorySuspensionPeriodForBookEntryTransfer() *Period3Choice {
	c.DepositorySuspensionPeriodForBookEntryTransfer = new(Period3Choice)
	return c.DepositorySuspensionPeriodForBookEntryTransfer
}

func (c *CorporateActionPeriod6) AddDepositorySuspensionPeriodForDepositAtAgent() *Period3Choice {
	c.DepositorySuspensionPeriodForDepositAtAgent = new(Period3Choice)
	return c.DepositorySuspensionPeriodForDepositAtAgent
}

func (c *CorporateActionPeriod6) AddDepositorySuspensionPeriodForDeposit() *Period3Choice {
	c.DepositorySuspensionPeriodForDeposit = new(Period3Choice)
	return c.DepositorySuspensionPeriodForDeposit
}

func (c *CorporateActionPeriod6) AddDepositorySuspensionPeriodForPledge() *Period3Choice {
	c.DepositorySuspensionPeriodForPledge = new(Period3Choice)
	return c.DepositorySuspensionPeriodForPledge
}

func (c *CorporateActionPeriod6) AddDepositorySuspensionPeriodForSegregation() *Period3Choice {
	c.DepositorySuspensionPeriodForSegregation = new(Period3Choice)
	return c.DepositorySuspensionPeriodForSegregation
}

func (c *CorporateActionPeriod6) AddDepositorySuspensionPeriodForWithdrawalAtAgent() *Period3Choice {
	c.DepositorySuspensionPeriodForWithdrawalAtAgent = new(Period3Choice)
	return c.DepositorySuspensionPeriodForWithdrawalAtAgent
}

func (c *CorporateActionPeriod6) AddDepositorySuspensionPeriodForWithdrawalInNomineeName() *Period3Choice {
	c.DepositorySuspensionPeriodForWithdrawalInNomineeName = new(Period3Choice)
	return c.DepositorySuspensionPeriodForWithdrawalInNomineeName
}

func (c *CorporateActionPeriod6) AddDepositorySuspensionPeriodForWithdrawalInStreetName() *Period3Choice {
	c.DepositorySuspensionPeriodForWithdrawalInStreetName = new(Period3Choice)
	return c.DepositorySuspensionPeriodForWithdrawalInStreetName
}

func (c *CorporateActionPeriod6) AddBookClosurePeriod() *Period3Choice {
	c.BookClosurePeriod = new(Period3Choice)
	return c.BookClosurePeriod
}

// Specifies periods related to a corporate action option.
type CorporateActionPeriod7 struct {

	// Period during which the price of a security is determined.
	PriceCalculationPeriod *Period3Choice `xml:"PricClctnPrd,omitempty"`

	// Period during which both old and new equity may be traded simultaneously, for example, consolidation of equity or splitting of equity.
	ParallelTradingPeriod *Period3Choice `xml:"ParllTradgPrd,omitempty"`

	// Period during which the specified option, or all options of the event, remains valid, for example, offer period.
	ActionPeriod *Period3Choice `xml:"ActnPrd,omitempty"`

	// Period during which the shareholder can revoke, change or withdraw its instruction.
	RevocabilityPeriod *Period3Choice `xml:"RvcbltyPrd,omitempty"`

	// Period during which the privilege is not available, for example, this can happen whenever a meeting takes place or whenever a coupon payment is due.
	PrivilegeSuspensionPeriod *Period3Choice `xml:"PrvlgSspnsnPrd,omitempty"`

	// Period during which the participant of the account servicer can revoke change or withdraw its instructions.
	AccountServicerRevocabilityPeriod *Period3Choice `xml:"AcctSvcrRvcbltyPrd,omitempty"`

	// Period defining the last date on which withdrawal in street name requests on the outturn security will be accepted and the date on which the suspension will be released and withdrawal by transfer processing on the outturn security will resume.
	DepositorySuspensionPeriodForWithdrawal *Period3Choice `xml:"DpstrySspnsnPrdForWdrwl,omitempty"`
}

func (c *CorporateActionPeriod7) AddPriceCalculationPeriod() *Period3Choice {
	c.PriceCalculationPeriod = new(Period3Choice)
	return c.PriceCalculationPeriod
}

func (c *CorporateActionPeriod7) AddParallelTradingPeriod() *Period3Choice {
	c.ParallelTradingPeriod = new(Period3Choice)
	return c.ParallelTradingPeriod
}

func (c *CorporateActionPeriod7) AddActionPeriod() *Period3Choice {
	c.ActionPeriod = new(Period3Choice)
	return c.ActionPeriod
}

func (c *CorporateActionPeriod7) AddRevocabilityPeriod() *Period3Choice {
	c.RevocabilityPeriod = new(Period3Choice)
	return c.RevocabilityPeriod
}

func (c *CorporateActionPeriod7) AddPrivilegeSuspensionPeriod() *Period3Choice {
	c.PrivilegeSuspensionPeriod = new(Period3Choice)
	return c.PrivilegeSuspensionPeriod
}

func (c *CorporateActionPeriod7) AddAccountServicerRevocabilityPeriod() *Period3Choice {
	c.AccountServicerRevocabilityPeriod = new(Period3Choice)
	return c.AccountServicerRevocabilityPeriod
}

func (c *CorporateActionPeriod7) AddDepositorySuspensionPeriodForWithdrawal() *Period3Choice {
	c.DepositorySuspensionPeriodForWithdrawal = new(Period3Choice)
	return c.DepositorySuspensionPeriodForWithdrawal
}

// Specifies periods of a corporate action.
type CorporateActionPeriod8 struct {

	// Period during which the price of a security is determined.
	PriceCalculationPeriod *Period3Choice `xml:"PricClctnPrd,omitempty"`

	// Period during which the interest rate has been applied.
	InterestPeriod *Period3Choice `xml:"IntrstPrd,omitempty"`

	// Period during a take-over where any outstanding equity must be purchased by the take-over company.
	CompulsoryPurchasePeriod *Period3Choice `xml:"CmplsryPurchsPrd,omitempty"`

	// Period during which the security is blocked.
	BlockingPeriod *Period3Choice `xml:"BlckgPrd,omitempty"`

	// Period assigned by the court in a class action. It determines the client's eligible transactions that will be included in the class action and used to determine the resulting entitlement.
	ClaimPeriod *Period3Choice `xml:"ClmPrd,omitempty"`

	// Period defining the last date for which book entry transfers will be accepted and the date on which the suspension will be released and book entry transfer processing will resume.
	DepositorySuspensionPeriodForBookEntryTransfer *Period3Choice `xml:"DpstrySspnsnPrdForBookNtryTrf,omitempty"`

	// Period defining the last date for which deposits, into nominee name, at the agent will be accepted and the date on which the suspension will be released and deposits at agent will resume.
	DepositorySuspensionPeriodForDepositAtAgent *Period3Choice `xml:"DpstrySspnsnPrdForDpstAtAgt,omitempty"`

	// Period defining the last date for which deposits will be accepted and the date on which the suspension will be released and deposits will resume.
	DepositorySuspensionPeriodForDeposit *Period3Choice `xml:"DpstrySspnsnPrdForDpst,omitempty"`

	// Period defining the last date for which pledges will be accepted and the date on which the suspension will be released and pledge processing will resume.
	DepositorySuspensionPeriodForPledge *Period3Choice `xml:"DpstrySspnsnPrdForPldg,omitempty"`

	// Period defining the last date for which intra-position balances can be segregated and the date on which the suspension will be released and the ability to segregate intra-position balances will resume.
	DepositorySuspensionPeriodForSegregation *Period3Choice `xml:"DpstrySspnsnPrdForSgrtn,omitempty"`

	// Period defining the last date for which withdrawals, from nominee name at the agent will be accepted and the date on which the suspension will be released and withdrawals at agent processing will resume.
	DepositorySuspensionPeriodForWithdrawalAtAgent *Period3Choice `xml:"DpstrySspnsnPrdForWdrwlAtAgt,omitempty"`

	// Period defining the last date for which physical withdrawals in the nominee's name will be accepted and the date on which the suspension will be released and physical withdrawals in the nominee's name will resume.
	DepositorySuspensionPeriodForWithdrawalInNomineeName *Period3Choice `xml:"DpstrySspnsnPrdForWdrwlInNmneeNm,omitempty"`

	// Period defining the last date on which withdrawal requests in street name's will be accepted on the event security and the date on which the suspension will be released and withdrawal in street name's processing on the event security will resume.
	DepositorySuspensionPeriodForWithdrawalInStreetName *Period3Choice `xml:"DpstrySspnsnPrdForWdrwlInStrtNm,omitempty"`

	// Period defining the last date on which shareholder registration will be accepted by the issuer and the date on which shareholder registration will resume.
	BookClosurePeriod *Period3Choice `xml:"BookClsrPrd,omitempty"`

	// Period during which the settlement activities at the co-depositories are suspended in order to stabilise the holdings at the CSD.
	CoDepositoriesSuspensionPeriod *Period3Choice `xml:"CoDpstriesSspnsnPrd,omitempty"`
}

func (c *CorporateActionPeriod8) AddPriceCalculationPeriod() *Period3Choice {
	c.PriceCalculationPeriod = new(Period3Choice)
	return c.PriceCalculationPeriod
}

func (c *CorporateActionPeriod8) AddInterestPeriod() *Period3Choice {
	c.InterestPeriod = new(Period3Choice)
	return c.InterestPeriod
}

func (c *CorporateActionPeriod8) AddCompulsoryPurchasePeriod() *Period3Choice {
	c.CompulsoryPurchasePeriod = new(Period3Choice)
	return c.CompulsoryPurchasePeriod
}

func (c *CorporateActionPeriod8) AddBlockingPeriod() *Period3Choice {
	c.BlockingPeriod = new(Period3Choice)
	return c.BlockingPeriod
}

func (c *CorporateActionPeriod8) AddClaimPeriod() *Period3Choice {
	c.ClaimPeriod = new(Period3Choice)
	return c.ClaimPeriod
}

func (c *CorporateActionPeriod8) AddDepositorySuspensionPeriodForBookEntryTransfer() *Period3Choice {
	c.DepositorySuspensionPeriodForBookEntryTransfer = new(Period3Choice)
	return c.DepositorySuspensionPeriodForBookEntryTransfer
}

func (c *CorporateActionPeriod8) AddDepositorySuspensionPeriodForDepositAtAgent() *Period3Choice {
	c.DepositorySuspensionPeriodForDepositAtAgent = new(Period3Choice)
	return c.DepositorySuspensionPeriodForDepositAtAgent
}

func (c *CorporateActionPeriod8) AddDepositorySuspensionPeriodForDeposit() *Period3Choice {
	c.DepositorySuspensionPeriodForDeposit = new(Period3Choice)
	return c.DepositorySuspensionPeriodForDeposit
}

func (c *CorporateActionPeriod8) AddDepositorySuspensionPeriodForPledge() *Period3Choice {
	c.DepositorySuspensionPeriodForPledge = new(Period3Choice)
	return c.DepositorySuspensionPeriodForPledge
}

func (c *CorporateActionPeriod8) AddDepositorySuspensionPeriodForSegregation() *Period3Choice {
	c.DepositorySuspensionPeriodForSegregation = new(Period3Choice)
	return c.DepositorySuspensionPeriodForSegregation
}

func (c *CorporateActionPeriod8) AddDepositorySuspensionPeriodForWithdrawalAtAgent() *Period3Choice {
	c.DepositorySuspensionPeriodForWithdrawalAtAgent = new(Period3Choice)
	return c.DepositorySuspensionPeriodForWithdrawalAtAgent
}

func (c *CorporateActionPeriod8) AddDepositorySuspensionPeriodForWithdrawalInNomineeName() *Period3Choice {
	c.DepositorySuspensionPeriodForWithdrawalInNomineeName = new(Period3Choice)
	return c.DepositorySuspensionPeriodForWithdrawalInNomineeName
}

func (c *CorporateActionPeriod8) AddDepositorySuspensionPeriodForWithdrawalInStreetName() *Period3Choice {
	c.DepositorySuspensionPeriodForWithdrawalInStreetName = new(Period3Choice)
	return c.DepositorySuspensionPeriodForWithdrawalInStreetName
}

func (c *CorporateActionPeriod8) AddBookClosurePeriod() *Period3Choice {
	c.BookClosurePeriod = new(Period3Choice)
	return c.BookClosurePeriod
}

func (c *CorporateActionPeriod8) AddCoDepositoriesSuspensionPeriod() *Period3Choice {
	c.CoDepositoriesSuspensionPeriod = new(Period3Choice)
	return c.CoDepositoriesSuspensionPeriod
}

// Specifies periods.
type CorporateActionPeriod9 struct {

	// Period during which the price of a security is determined.
	PriceCalculationPeriod *Period5 `xml:"PricClctnPrd,omitempty"`

	// Period during which the specified option, or all options of the event, remains valid, for example, offer period.
	ActionPeriod *Period5 `xml:"ActnPrd,omitempty"`

	// Period during which both old and new equity may be traded simultaneously, for example, consolidation of equity or splitting of equity.
	ParallelTradingPeriod *Period5 `xml:"ParllTradgPrd,omitempty"`
}

func (c *CorporateActionPeriod9) AddPriceCalculationPeriod() *Period5 {
	c.PriceCalculationPeriod = new(Period5)
	return c.PriceCalculationPeriod
}

func (c *CorporateActionPeriod9) AddActionPeriod() *Period5 {
	c.ActionPeriod = new(Period5)
	return c.ActionPeriod
}

func (c *CorporateActionPeriod9) AddParallelTradingPeriod() *Period5 {
	c.ParallelTradingPeriod = new(Period5)
	return c.ParallelTradingPeriod
}
