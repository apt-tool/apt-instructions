package http

type ExecuteRequest struct {
	Params     []string `json:"params"`
	Path       string   `json:"path"`
	DocumentID uint     `json:"document_id"`
}
