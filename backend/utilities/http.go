package utilities

import "net/http"

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
