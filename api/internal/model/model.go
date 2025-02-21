package model

type Country struct {
	Name       Name                `json:"name"`
	Capital    []string            `json:"capital"`
	Currencies map[string]Currency `json:"currencies"`
	Population int                 `json:"population"`
}

type Name struct {
	Common string `json:"common"`
}
type ErrorResponse struct {
	Error string `json:"error"`
}

type Currency struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}
