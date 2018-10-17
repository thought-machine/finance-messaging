package iso20022

// Content of the acceptor configuration.
type AcceptorConfigurationContent1 struct {

	// Acceptor parameters dedicated to an acquirer protocol.
	AcquirerProtocolParameters []*AcquirerProtocolParameters1 `xml:"AcqrrPrtcolParams,omitempty"`

	// Acceptor parameters dedicated to the merchant.
	MerchantParameters []*Max10000Binary `xml:"MrchntParams,omitempty"`

	// Acceptor parameters dedicated to a payment application of the point of interaction.
	ApplicationParameters []*ApplicationParameters1 `xml:"ApplParams,omitempty"`

	// Acceptor parameters dedicated to the communication with an acquirer host.
	HostCommunicationParameters []*HostCommunicationParameter1 `xml:"HstComParams,omitempty"`
}

func (a *AcceptorConfigurationContent1) AddAcquirerProtocolParameters() *AcquirerProtocolParameters1 {
	newValue := new(AcquirerProtocolParameters1)
	a.AcquirerProtocolParameters = append(a.AcquirerProtocolParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent1) AddMerchantParameters(value string) {
	a.MerchantParameters = append(a.MerchantParameters, (*Max10000Binary)(&value))
}

func (a *AcceptorConfigurationContent1) AddApplicationParameters() *ApplicationParameters1 {
	newValue := new(ApplicationParameters1)
	a.ApplicationParameters = append(a.ApplicationParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent1) AddHostCommunicationParameters() *HostCommunicationParameter1 {
	newValue := new(HostCommunicationParameter1)
	a.HostCommunicationParameters = append(a.HostCommunicationParameters, newValue)
	return newValue
}

// Content of the acceptor configuration.
type AcceptorConfigurationContent2 struct {

	// Acceptor parameters dedicated to an acquirer protocol.
	AcquirerProtocolParameters []*AcquirerProtocolParameters3 `xml:"AcqrrPrtcolParams,omitempty"`

	// Acceptor parameters dedicated to the merchant.
	MerchantParameters []*Max10000Binary `xml:"MrchntParams,omitempty"`

	// Manufacturer configuration parameters of the point of interaction.
	TerminalParameters []*Max10000Binary `xml:"TermnlParams,omitempty"`

	// Acceptor parameters dedicated to a payment application of the point of interaction.
	ApplicationParameters []*ApplicationParameters2 `xml:"ApplParams,omitempty"`

	// Acceptor parameters dedicated to the communication with an acquirer host.
	HostCommunicationParameters []*HostCommunicationParameter2 `xml:"HstComParams,omitempty"`

	// Point of interaction parameters related to the security of software application and application protocol.
	SecurityParameters []*SecurityParameters1 `xml:"SctyParams,omitempty"`
}

func (a *AcceptorConfigurationContent2) AddAcquirerProtocolParameters() *AcquirerProtocolParameters3 {
	newValue := new(AcquirerProtocolParameters3)
	a.AcquirerProtocolParameters = append(a.AcquirerProtocolParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent2) AddMerchantParameters(value string) {
	a.MerchantParameters = append(a.MerchantParameters, (*Max10000Binary)(&value))
}

func (a *AcceptorConfigurationContent2) AddTerminalParameters(value string) {
	a.TerminalParameters = append(a.TerminalParameters, (*Max10000Binary)(&value))
}

func (a *AcceptorConfigurationContent2) AddApplicationParameters() *ApplicationParameters2 {
	newValue := new(ApplicationParameters2)
	a.ApplicationParameters = append(a.ApplicationParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent2) AddHostCommunicationParameters() *HostCommunicationParameter2 {
	newValue := new(HostCommunicationParameter2)
	a.HostCommunicationParameters = append(a.HostCommunicationParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent2) AddSecurityParameters() *SecurityParameters1 {
	newValue := new(SecurityParameters1)
	a.SecurityParameters = append(a.SecurityParameters, newValue)
	return newValue
}

// Content of the acceptor configuration.
type AcceptorConfigurationContent3 struct {

	// Acceptor parameters dedicated to an acquirer protocol.
	AcquirerProtocolParameters []*AcquirerProtocolParameters6 `xml:"AcqrrPrtcolParams,omitempty"`

	// Acceptor parameters dedicated to the merchant.
	MerchantParameters []*Max10000Binary `xml:"MrchntParams,omitempty"`

	// Manufacturer configuration parameters of the point of interaction.
	TerminalParameters []*PaymentTerminalParameters1 `xml:"TermnlParams,omitempty"`

	// Acceptor parameters dedicated to a payment application of the point of interaction.
	ApplicationParameters []*ApplicationParameters3 `xml:"ApplParams,omitempty"`

	// Acceptor parameters dedicated to the communication with an acquirer host.
	HostCommunicationParameters []*HostCommunicationParameter2 `xml:"HstComParams,omitempty"`

	// Point of interaction parameters related to the security of software application and application protocol.
	SecurityParameters []*SecurityParameters2 `xml:"SctyParams,omitempty"`
}

func (a *AcceptorConfigurationContent3) AddAcquirerProtocolParameters() *AcquirerProtocolParameters6 {
	newValue := new(AcquirerProtocolParameters6)
	a.AcquirerProtocolParameters = append(a.AcquirerProtocolParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent3) AddMerchantParameters(value string) {
	a.MerchantParameters = append(a.MerchantParameters, (*Max10000Binary)(&value))
}

func (a *AcceptorConfigurationContent3) AddTerminalParameters() *PaymentTerminalParameters1 {
	newValue := new(PaymentTerminalParameters1)
	a.TerminalParameters = append(a.TerminalParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent3) AddApplicationParameters() *ApplicationParameters3 {
	newValue := new(ApplicationParameters3)
	a.ApplicationParameters = append(a.ApplicationParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent3) AddHostCommunicationParameters() *HostCommunicationParameter2 {
	newValue := new(HostCommunicationParameter2)
	a.HostCommunicationParameters = append(a.HostCommunicationParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent3) AddSecurityParameters() *SecurityParameters2 {
	newValue := new(SecurityParameters2)
	a.SecurityParameters = append(a.SecurityParameters, newValue)
	return newValue
}

// Content of the acceptor configuration.
type AcceptorConfigurationContent4 struct {

	// Configuration parameters of the TMS protocol between a POI and a terminal manager.
	TMSProtocolParameters []*TMSProtocolParameters1 `xml:"TMSPrtcolParams,omitempty"`

	// Acceptor parameters dedicated to an acquirer protocol.
	AcquirerProtocolParameters []*AcquirerProtocolParameters7 `xml:"AcqrrPrtcolParams,omitempty"`

	// Acceptor parameters dedicated to the merchant.
	MerchantParameters []*MerchantConfigurationParameters1 `xml:"MrchntParams,omitempty"`

	// Manufacturer configuration parameters of the point of interaction.
	TerminalParameters []*PaymentTerminalParameters2 `xml:"TermnlParams,omitempty"`

	// Acceptor parameters dedicated to a payment application of the point of interaction.
	ApplicationParameters []*ApplicationParameters4 `xml:"ApplParams,omitempty"`

	// Acceptor parameters dedicated to the communication with an acquirer host or a terminal manager host.
	HostCommunicationParameters []*HostCommunicationParameter3 `xml:"HstComParams,omitempty"`

	// Point of interaction parameters related to the security of software application and application protocol.
	SecurityParameters []*SecurityParameters3 `xml:"SctyParams,omitempty"`
}

func (a *AcceptorConfigurationContent4) AddTMSProtocolParameters() *TMSProtocolParameters1 {
	newValue := new(TMSProtocolParameters1)
	a.TMSProtocolParameters = append(a.TMSProtocolParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent4) AddAcquirerProtocolParameters() *AcquirerProtocolParameters7 {
	newValue := new(AcquirerProtocolParameters7)
	a.AcquirerProtocolParameters = append(a.AcquirerProtocolParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent4) AddMerchantParameters() *MerchantConfigurationParameters1 {
	newValue := new(MerchantConfigurationParameters1)
	a.MerchantParameters = append(a.MerchantParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent4) AddTerminalParameters() *PaymentTerminalParameters2 {
	newValue := new(PaymentTerminalParameters2)
	a.TerminalParameters = append(a.TerminalParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent4) AddApplicationParameters() *ApplicationParameters4 {
	newValue := new(ApplicationParameters4)
	a.ApplicationParameters = append(a.ApplicationParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent4) AddHostCommunicationParameters() *HostCommunicationParameter3 {
	newValue := new(HostCommunicationParameter3)
	a.HostCommunicationParameters = append(a.HostCommunicationParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent4) AddSecurityParameters() *SecurityParameters3 {
	newValue := new(SecurityParameters3)
	a.SecurityParameters = append(a.SecurityParameters, newValue)
	return newValue
}

// Content of the acceptor configuration.
type AcceptorConfigurationContent5 struct {

	// True if the whole configuration related to the terminal manager has to be replaced by the configuration included in the message content.
	ReplaceConfiguration *TrueFalseIndicator `xml:"RplcCfgtn,omitempty"`

	// Configuration parameters of the TMS protocol between a POI and a terminal manager.
	TMSProtocolParameters []*TMSProtocolParameters2 `xml:"TMSPrtcolParams,omitempty"`

	// Acceptor parameters dedicated to an acquirer protocol.
	AcquirerProtocolParameters []*AcquirerProtocolParameters9 `xml:"AcqrrPrtcolParams,omitempty"`

	// Acceptor parameters dedicated to the merchant.
	MerchantParameters []*MerchantConfigurationParameters2 `xml:"MrchntParams,omitempty"`

	// Manufacturer configuration parameters of the point of interaction.
	TerminalParameters []*PaymentTerminalParameters3 `xml:"TermnlParams,omitempty"`

	// Acceptor parameters dedicated to a payment application of the point of interaction.
	ApplicationParameters []*ApplicationParameters5 `xml:"ApplParams,omitempty"`

	// Acceptor parameters dedicated to the communication with an acquirer host or a terminal manager host.
	HostCommunicationParameters []*HostCommunicationParameter4 `xml:"HstComParams,omitempty"`

	// Point of interaction parameters related to the security of software application and application protocol.
	SecurityParameters []*SecurityParameters6 `xml:"SctyParams,omitempty"`
}

func (a *AcceptorConfigurationContent5) SetReplaceConfiguration(value string) {
	a.ReplaceConfiguration = (*TrueFalseIndicator)(&value)
}

func (a *AcceptorConfigurationContent5) AddTMSProtocolParameters() *TMSProtocolParameters2 {
	newValue := new(TMSProtocolParameters2)
	a.TMSProtocolParameters = append(a.TMSProtocolParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent5) AddAcquirerProtocolParameters() *AcquirerProtocolParameters9 {
	newValue := new(AcquirerProtocolParameters9)
	a.AcquirerProtocolParameters = append(a.AcquirerProtocolParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent5) AddMerchantParameters() *MerchantConfigurationParameters2 {
	newValue := new(MerchantConfigurationParameters2)
	a.MerchantParameters = append(a.MerchantParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent5) AddTerminalParameters() *PaymentTerminalParameters3 {
	newValue := new(PaymentTerminalParameters3)
	a.TerminalParameters = append(a.TerminalParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent5) AddApplicationParameters() *ApplicationParameters5 {
	newValue := new(ApplicationParameters5)
	a.ApplicationParameters = append(a.ApplicationParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent5) AddHostCommunicationParameters() *HostCommunicationParameter4 {
	newValue := new(HostCommunicationParameter4)
	a.HostCommunicationParameters = append(a.HostCommunicationParameters, newValue)
	return newValue
}

func (a *AcceptorConfigurationContent5) AddSecurityParameters() *SecurityParameters6 {
	newValue := new(SecurityParameters6)
	a.SecurityParameters = append(a.SecurityParameters, newValue)
	return newValue
}
