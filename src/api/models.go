package api

// Squad model
type Squad struct {
	XMLName string   `xml:"squad"`
	Nick    string   `xml:"nick,attr"`
	Email   string   `xml:"email"`
	Name    string   `xml:"name"`
	Picture string   `xml:"picture"`
	Title   string   `xml:"title"`
	Web     string   `xml:"web"`
	Members []Member `xml:"member"`
}

// Member model
type Member struct {
	ID   string `xml:"id,attr"`
	Nick string `xml:"nick,attr"`
	Name string `xml:"name"`
}
