package utilities

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

var httpClient *fasthttp.Client

func init() {
	httpClient = &fasthttp.Client{}
}

func Request(method string, url string, body []byte, modifiers ...RequestModifier) (*fasthttp.Response, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod(method)
	req.SetRequestURI(url)
	req.SetBody(body)

	req = ApplyModifiers(req, modifiers...)

	err := httpClient.Do(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type RequestModifier interface {
	Modify(req *fasthttp.Request) *fasthttp.Request
}

type HeaderKVModifier struct {
	Key   string
	Value string
}

func (h HeaderKVModifier) Modify(req *fasthttp.Request) *fasthttp.Request {
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
		Value: fmt.Sprintf("Bearer %s", apiKey),
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

func ApplyModifiers(req *fasthttp.Request, modifiers ...RequestModifier) *fasthttp.Request {
	for _, modifier := range modifiers {
		req = modifier.Modify(req)
	}
	return req
}
