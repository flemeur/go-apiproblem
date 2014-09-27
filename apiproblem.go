// Copyright (c) 2014 Flemming Andersen.
// Distributed under the MIT License (MIT)
// which can be found in the LICENSE file.

package apiproblem

import (
	"encoding/xml"
	"fmt"
)

// Constants containing the mime types for JSON and XML
const (
	JSONMimeType = "application/problem+json"
	XMLMimeType  = "application/problem+xml"
)

// APIProblem data structure
type APIProblem struct {
	XMLName xml.Name `json:"-" xml:"problem"`
	Type    string   `json:"type,omitempty" xml:"type,omitempty"`
	// Status code from the API
	Status int `json:"status" xml:"status"`
	// Title of error
	Title string `json:"title" xml:"title"`
	// Details describing the error
	Detail string `json:"detail,omitempty" xml:"detail,omitempty"`
}

// New returns a new *APIProblem
func New(status int, title, detail string) *APIProblem {
	return &APIProblem{
		Type:   "http://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html",
		Status: status,
		Title:  title,
		Detail: detail,
	}
}

// Implements the error interface
func (p *APIProblem) Error() string {
	return fmt.Sprintf("%d %s: %s", p.Status, p.Title, p.Detail)
}
