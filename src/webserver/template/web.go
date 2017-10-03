package template

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strings"
)

type Response struct {
	Message string      `json:"message,omitempty"`
	Error   []string    `json:"error,omitempty"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data, omitempty"`
}

func (r *Response) SetMessage(msg string) *Response {
	r.Message = msg
	return r
}

func (r *Response) AddError(msgs ...string) *Response {
	for _, val := range msgs {
		r.Error = append(r.Error, val)
	}
	return r
}

func (r *Response) SetCode(code int) *Response {
	r.Code = code
	return r
}

func (r *Response) SetData(data interface{}) *Response {
	r.Data = data
	return r
}

// RenderJSONResponse this is as per new response scheme
func RenderJSONResponse(w http.ResponseWriter, data interface{}, errs ...string) {
	var g interface{}
	w.Header().Set("Content-Type", "application/json")
	statuscode := 200

	x, _ := json.Marshal(data)
	json.Unmarshal(x, &g)

	if g != nil && reflect.TypeOf(g).Kind() == reflect.Map {
		if value, ok := g.(map[string]interface{})["code"]; ok {
			statuscode = int(value.(float64))
		}

		if http.StatusText(statuscode) == "" {
			log.Println(statuscode)
		}

		if statuscode == 500 {
			errList := append(errs, string(x))

			if len(errList) > 0 {
				for _, e := range errList {

					log.Println(string(x), e)
				}

			}
			x = []byte(`{"errors":["Internal Error"], "code": "500"}`)
		}
		if statuscode > 200 {
			log.Println(string(x))
		}
	}
	x = []byte(strings.Replace(string(x), `\n`, `\u003cbr\u003e`, -1))
	w.WriteHeader(statuscode)
	w.Write(x)
	return
}
