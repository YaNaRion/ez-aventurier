package gateway

type EventLoginRequest struct {
	UniqueID string `json:"uniqueID"`
}

type EventLoginResponse struct {
	UniqueID     string `json:"uniqueID"`
	IsConnection bool   `json:"isConnection"`
}
