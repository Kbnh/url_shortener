package dto

type SaveURLRequest struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type SaveURLResponse struct {
	Params
	Alias string `json:"alias"`
}
