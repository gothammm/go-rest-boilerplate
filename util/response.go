package util


import (
	"encoding/json"
	"net/http"
	"errors"
)


type Response struct {
	Code int    `json:"code,omitempty"`
	statusCode int
	writer http.ResponseWriter
	Message string  `json:"message,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

type key int

const HttpStatusKey key = 1

func (r Response) Json() {
	r.writer.WriteHeader(r.statusCode)
	r.writer.Header().Set("Content-Type", "application/json")
	r.writer.Write(ToJSON(r))
}

func (r Response) ToJSON() string {
	return string(ToJSON(r))
}

func (r Response) Writer(w http.ResponseWriter) Response {
	if w == nil {
		panic(errors.New("Writer cannot be nil"))
	}

	r.writer = w
	return r
}

func (r Response) Status(clientCode int, statusCode int) Response {
	r.Code = clientCode
	r.statusCode = statusCode
	return r
}

func ToJSONString(obj interface{}) string {

	payload, err := json.MarshalIndent(obj, "", "   ")

	if err != nil {
		panic(err)
	}


	if payload != nil {
		return string(payload)
	}

	return ""
}

func ToJSON(obj interface{}) []byte {
	return []byte(ToJSONString(obj))
}