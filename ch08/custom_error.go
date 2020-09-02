package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	HttpCode int    `json:"-"`
	Code     int    `json:"code,omitempty"`
	Message  string `json:"message"`
}

func (e *Error) Error() string {
	fs := "httpStatus: %d, errorCode: %d, message: %s\n"
	return fmt.Sprintf(fs, e.HttpCode, e.Code, e.Message)
}

func JsonError(resp http.ResponseWriter, err Error) {
	data := struct {
		Err Error `json:"err"`
	}{err}

	b, e := json.Marshal(data)
	if e != nil {
		http.Error(resp, "Internal Server Error", 500)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(err.HttpCode)
	fmt.Fprint(resp, string(b))
}

func displayError(resp http.ResponseWriter, req *http.Request) {
	err := Error{
		HttpCode: http.StatusForbidden,
		Code:     123,
		Message:  "You are not allowed to see this",
	}
	JsonError(resp, err)
}

func main() {
	http.HandleFunc("/", displayError)
	http.ListenAndServe(":8080", nil)
}
