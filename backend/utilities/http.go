package utilities

import (
	"io"
	"net/http"

	"github.com/valyala/fasthttp"
)

func Request(method string, url string, body io.Reader, modifiers ...RequestModifier) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req = ApplyModifiers(req, modifiers...)
	client := fasthttp.Client{}

	return client.Do(req)
	"net/http"
)

var HttpClient *http.Client

func init() {
	HttpClient = http.DefaultClient
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

func Authorization(apiKey string) RequestModifier {
	return HeaderKVModifier{
		Key:   "Authorization",
		Value: "Bearer " + apiKey,
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
