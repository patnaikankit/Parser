package pkg

type Section struct {
	Header string
	Body   []string
}

type Structure struct {
	Title    string
	Data     map[string]string
	Sections []Section
}
