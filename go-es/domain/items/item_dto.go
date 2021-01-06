package items

type Item struct {
	Id          string      `json:"id"`
	Title       string      `json:"title"`
	Description Description `json:"description"`
	Pictures    []Picture   `json:"pictures"`
	Status      string      `json:"status"`
}

type Description struct {
	PlainText string `json:"plain_text"`
	Html      string `json:"html"`
}

type Picture struct {
	Id  int64  `json:"id"`
	Url string `json:"url"`
}
