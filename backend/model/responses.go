package model

// PostInfoHandlerResponse returned from PostInfoHandler
type PostInfoHandlerResponse struct {
	Message string `json:"message"`
	ID      int    `json:"id"`
}

// GetInfoListHandlerResponse returned from GetInfoHandler
type GetInfoListHandlerResponse struct {
	Info InfoList `json:"info"`
}

// UpdateInfoHandlerResponse returned from UpdateInfoHandler
type UpdateInfoHandlerResponse struct {
	Message string `json:"message"`
}

// GetTypesHandlerResponse returned from GetTypesHandler
type GetTypesHandlerResponse struct {
	Types TypeList `json:"types"`
}

// GetServicesHandlerResponse returned from GetServicesHandler
type GetServicesHandlerResponse struct {
	Services ServiceList `json:"services"`
}
