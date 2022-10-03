package request

type BookUpdateRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
