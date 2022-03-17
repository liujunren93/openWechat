package utils

import (
	"encoding/xml"
	"strconv"
)

type CDATAStr string

func (s CDATAStr) String() string {
	return string(s)
}
func (c CDATAStr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}

type CDATAInt int

func (c CDATAInt) String() string {
	return strconv.FormatInt(int64(c), 10)
}
func (c CDATAInt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{c.String()}, start)
}
