package utilities

import (
	"bytes"
	"fmt"
	"net/http"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{}
}

func Request(method string, url string, body []byte, modifiers ...RequestModifier) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req = ApplyModifiers(req, modifiers...)

	return httpClient.Do(req)
}

func Client() *http.Client {
	return httpClient
}

func IsOk(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}

type RequestModifier interface {
	Modify(req *http.Request) *http.Request
}

type HeaderKVModifier struct {
	Key   string
	Value string
}

func (h HeaderKVModifier) Modify(req *http.Request) *http.Request {
	req.Header.Set(h.Key, h.Value)
	return req
}

func HeaderKV(key, value string) RequestModifier {
	return HeaderKVModifier{
		Key:   key,
		Value: value,
	}
}

func Authorization(value string) RequestModifier {
	return HeaderKVModifier{
		Key:   "Authorization",
		Value: fmt.Sprintf("Bearer %s", value),
	}
}

func JSON() RequestModifier {
	return HeaderKVModifier{
		Key:   "Content-Type",
		Value: "application/json",
	}
}

func FormURLEncoded() RequestModifier {
	return HeaderKVModifier{
		Key:   "Content-Type",
		Value: "application/x-www-form-urlencoded",
	}
}

func AcceptJSON() RequestModifier {
	return HeaderKVModifier{
		Key:   "Accept",
		Value: "application/json",
	}
}

func ApplyModifiers(req *http.Request, modifiers ...RequestModifier) *http.Request {
	for _, modifier := range modifiers {
		req = modifier.Modify(req)
	}
	return req
}
