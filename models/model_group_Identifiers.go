package models

// refer to TS29505 v17.9.0, section 5.4.2.35 Type: GroupIdentifiers
type GroupIdentifiers struct {
	ExtGroupId string   `json:"extGroupId,omitempty"`
	IntGroupId string   `json:"intGroupId,omitempty"`
	UeIdList   []string `json:"ueIdList,omitempty"`
	//AllowedAfIds: A list of Application Function Identifiers authorized to retrieve this Identities lists.
	// The absence of this IE indicates that any AF is allowed to retrieve them.
	AllowedAfIds []string `json:"allowedAfIds,omitempty"`
}
