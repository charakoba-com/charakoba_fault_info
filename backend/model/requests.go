package model

// PostInfoHandlerRequest is request type for PostInfoHandler
type PostInfoHandlerRequest struct {
	InfoType string `json:"info_type"`
	Service  string `json:"service"`
	Begin    string `json:"begin"`
	End      string `json:"end,omitempty"`
	Detail   string `json:"detail,omitempty"`
	Token    string `json:"token"`
}

// UpdateInfoHandlerRequest is request type for UpdateInfoHandler
type UpdateInfoHandlerRequest struct {
	InfoType string `json:"info_type,omitempty"`
	Service  string `json:"service,omitempty"`
	Begin    string `json:"begin,omitempty"`
	End      string `json:"end,omitempty"`
	Detail   string `json:"detail,omitempty"`
	Token    string `json:"token"`
}
