package models

type AllowedMtcProviderInfo struct {
	MtcProviderInformation MtcProviderInformation `json:"mtcProviderInformation,omitempty"`
	AfId                   string                 `json:"afId,omitempty"`
}
