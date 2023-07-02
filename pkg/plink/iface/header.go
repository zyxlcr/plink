package iface

import (
	"net/url"

	jsoniter "github.com/json-iterator/go"
)

type Header struct {
	Url   string `json:"url"`
	From  string `json:"from"`
	To    string `json:"to"`
	Token string `json:"token"`
}

func NewHeader(url string) *Header {
	return &Header{
		Url: url,
	}
}

func (h *Header) ToJson() ([]byte, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(h)
}

func FromJsonTo(jsonData []byte, h *Header) *Header {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(jsonData, h)
	return h
}

func (h *Header) ToReq(conn IConnection, msg IMessage) *Request {

	return &Request{
		conn:   conn,
		msg:    msg,
		Method: "POST",
		URL:    &url.URL{Path: h.Url},
	}
}
