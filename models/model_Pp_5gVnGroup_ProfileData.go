package models

// A map (list of key-value pairs where ExtGroupId (see 3GPP TS 29.503 [6]) serves as key of a list of
// AllowedMtcProviderInfos which include MTC provider informations or AF IDs that are allowed to
// create, update and delete a 5G VN Group for the user using UDM ParameterProvision service.

// In addition to defined external group identifier, the key value "ALL" may be used to identify a map entry
// which contains a list of AllowedMtcProviderInfos that are allowed operating all the external group identifiers.

type Pp5gVnGroupProfileData struct {
	AllowedMtcProviders map[string][]AllowedMtcProviderInfo `json:"allowedMtcProviders,omitempty"`
	SupportedFeatures   string                              `json:"supportedFeatures,omitempty"`
}
