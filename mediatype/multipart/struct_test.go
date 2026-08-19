package multipart

type MultipartRelatedTestBody struct {
	JsonData *ContentDescriptor `multipart:"contentType:application/json"`
	Message  *RelatedContent    `multipart:"contentType:application/text"`
	UserData *RelatedContent    `multipart:"contentType:application/x-www-form-urlencoded"`
}

type ContentDescriptor struct {
	MessageContainer  *RefToBinaryData `json:"messageContainer,omitempty"`
	UserDataContainer *RefToBinaryData `json:"userDataContainer,omitempty"`
}

type RefToBinaryData struct {
	ContentID string `json:"contentId,omitempty"`
}

type MultipartRelatedTestBody2 struct {
	JsonData  *ContentDescriptor2 `multipart:"contentType:application/json"`
	Message1  *RelatedContent     `multipart:"contentType:application/text"`
	Message2  *RelatedContent     `multipart:"contentType:application/text"`
	Message3  *RelatedContent     `multipart:"contentType:application/text"`
	UserData1 *RelatedContent     `multipart:"contentType:application/x-www-form-urlencoded"`
	UserData2 *RelatedContent     `multipart:"contentType:application/x-www-form-urlencoded"`
	UserData3 *RelatedContent     `multipart:"contentType:application/x-www-form-urlencoded"`
}

type ContentDescriptor2 struct {
	MessageContainer  []*RefToBinaryData `json:"messageContainer,omitempty"`
	UserDataContainer []*RefToBinaryData `json:"userDataContainer,omitempty"`
}
