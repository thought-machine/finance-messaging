package iso20022

// Acceptor parameters dedicated to a payment application of the point of interaction.
type ApplicationParameters1 struct {

	// Identification of the payment application.
	ApplicationIdentification *Max35Text `xml:"ApplId"`

	// Version of the payment application configuration parameters.
	Version *Max16Text `xml:"Vrsn"`

	// Configuration parameters used by the related payment application.
	Parameters []*Max10000Binary `xml:"Params,omitempty"`

	// Sensitive parameters (sequence of Parameters including the enveloppes) encrypted with a cryptographic key.
	EncryptedParameters *ContentInformationType2 `xml:"NcrptdParams,omitempty"`
}

func (a *ApplicationParameters1) SetApplicationIdentification(value string) {
	a.ApplicationIdentification = (*Max35Text)(&value)
}

func (a *ApplicationParameters1) SetVersion(value string) {
	a.Version = (*Max16Text)(&value)
}

func (a *ApplicationParameters1) AddParameters(value string) {
	a.Parameters = append(a.Parameters, (*Max10000Binary)(&value))
}

func (a *ApplicationParameters1) AddEncryptedParameters() *ContentInformationType2 {
	a.EncryptedParameters = new(ContentInformationType2)
	return a.EncryptedParameters
}

// Acceptor parameters dedicated to a payment application of the point of interaction.
type ApplicationParameters2 struct {

	// Identification of the payment application.
	ApplicationIdentification *Max35Text `xml:"ApplId"`

	// Version of the payment application configuration parameters.
	Version *Max16Text `xml:"Vrsn"`

	// Configuration parameters used by the related payment application.
	Parameters []*Max10000Binary `xml:"Params,omitempty"`

	// Sensitive parameters (sequence of Parameters including the enveloppes) encrypted with a cryptographic key.
	EncryptedParameters *ContentInformationType5 `xml:"NcrptdParams,omitempty"`
}

func (a *ApplicationParameters2) SetApplicationIdentification(value string) {
	a.ApplicationIdentification = (*Max35Text)(&value)
}

func (a *ApplicationParameters2) SetVersion(value string) {
	a.Version = (*Max16Text)(&value)
}

func (a *ApplicationParameters2) AddParameters(value string) {
	a.Parameters = append(a.Parameters, (*Max10000Binary)(&value))
}

func (a *ApplicationParameters2) AddEncryptedParameters() *ContentInformationType5 {
	a.EncryptedParameters = new(ContentInformationType5)
	return a.EncryptedParameters
}

// Acceptor parameters dedicated to a payment application of the point of interaction.
type ApplicationParameters3 struct {

	// Identification of the payment application.
	ApplicationIdentification *Max35Text `xml:"ApplId"`

	// Version of the payment application configuration parameters.
	Version *Max16Text `xml:"Vrsn"`

	// Configuration parameters used by the related payment application.
	Parameters []*Max100KBinary `xml:"Params,omitempty"`

	// Sensitive parameters (sequence of parameters including the enveloppes) encrypted with a cryptographic key.
	EncryptedParameters *ContentInformationType7 `xml:"NcrptdParams,omitempty"`
}

func (a *ApplicationParameters3) SetApplicationIdentification(value string) {
	a.ApplicationIdentification = (*Max35Text)(&value)
}

func (a *ApplicationParameters3) SetVersion(value string) {
	a.Version = (*Max16Text)(&value)
}

func (a *ApplicationParameters3) AddParameters(value string) {
	a.Parameters = append(a.Parameters, (*Max100KBinary)(&value))
}

func (a *ApplicationParameters3) AddEncryptedParameters() *ContentInformationType7 {
	a.EncryptedParameters = new(ContentInformationType7)
	return a.EncryptedParameters
}

// Acceptor parameters dedicated to a payment application of the point of interaction.
type ApplicationParameters4 struct {

	// Identification of the payment application.
	ApplicationIdentification *Max35Text `xml:"ApplId"`

	// Version of the payment application configuration parameters.
	Version *Max256Text `xml:"Vrsn"`

	// Configuration parameters used by the related payment application.
	Parameters []*Max100KBinary `xml:"Params,omitempty"`

	// Sensitive parameters (sequence of parameters including the envelope) encrypted with a cryptographic key.
	EncryptedParameters *ContentInformationType10 `xml:"NcrptdParams,omitempty"`
}

func (a *ApplicationParameters4) SetApplicationIdentification(value string) {
	a.ApplicationIdentification = (*Max35Text)(&value)
}

func (a *ApplicationParameters4) SetVersion(value string) {
	a.Version = (*Max256Text)(&value)
}

func (a *ApplicationParameters4) AddParameters(value string) {
	a.Parameters = append(a.Parameters, (*Max100KBinary)(&value))
}

func (a *ApplicationParameters4) AddEncryptedParameters() *ContentInformationType10 {
	a.EncryptedParameters = new(ContentInformationType10)
	return a.EncryptedParameters
}

// Acceptor parameters dedicated to a payment application of the point of interaction.
type ApplicationParameters5 struct {

	// Type of action for the configuration parameters.
	ActionType *TerminalManagementAction3Code `xml:"ActnTp"`

	// Identification of the payment application.
	ApplicationIdentification *Max35Text `xml:"ApplId"`

	// Version of the payment application configuration parameters.
	Version *Max256Text `xml:"Vrsn"`

	// Configuration parameters used by the related payment application.
	Parameters []*Max100KBinary `xml:"Params,omitempty"`

	// Sensitive parameters (sequence of parameters including the envelope) encrypted with a cryptographic key.
	EncryptedParameters *ContentInformationType10 `xml:"NcrptdParams,omitempty"`
}

func (a *ApplicationParameters5) SetActionType(value string) {
	a.ActionType = (*TerminalManagementAction3Code)(&value)
}

func (a *ApplicationParameters5) SetApplicationIdentification(value string) {
	a.ApplicationIdentification = (*Max35Text)(&value)
}

func (a *ApplicationParameters5) SetVersion(value string) {
	a.Version = (*Max256Text)(&value)
}

func (a *ApplicationParameters5) AddParameters(value string) {
	a.Parameters = append(a.Parameters, (*Max100KBinary)(&value))
}

func (a *ApplicationParameters5) AddEncryptedParameters() *ContentInformationType10 {
	a.EncryptedParameters = new(ContentInformationType10)
	return a.EncryptedParameters
}
