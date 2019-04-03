package book

type Author struct {
	Firstname string `json:firstname`
	Lastname string `json:lastname`
}

type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}
