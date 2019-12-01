package responses

type CheckConnectionResponse struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Time    string `json:"time"`
}
