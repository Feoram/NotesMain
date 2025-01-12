package service

type Note struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type QuoteResponse struct {
	Quote struct {
		Body string `json:"body"`
	} `json:"quote"`
}
