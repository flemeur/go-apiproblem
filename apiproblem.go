package apiproblem

import (
	"encoding/xml"
	"fmt"
)

const (
	JSONMimeType = "application/problem+json"
	XMLMimeType  = "application/problem+xml"
)

type APIProblem struct {
	XMLName xml.Name `json:"-" xml:"problem"`
	Type    string   `json:"type,omitempty" xml:"type,omitempty"`
	Status  int      `json:"status" xml:"status"`
	Title   string   `json:"title" xml:"title"`
	Detail  string   `json:"detail,omitempty" xml:"detail,omitempty"`
}

func New(status int, title, detail string) *APIProblem {
	return &APIProblem{
		Type:   "http://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html",
		Status: status,
		Title:  title,
		Detail: detail,
	}
}

func (p APIProblem) Error() string {
	return fmt.Sprintf("%d %s: %s", p.Status, p.Title, p.Detail)
}
