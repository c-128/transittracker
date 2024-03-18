package handlers

import (
	"io"
	"net/http"
)

type Error func(writer http.ResponseWriter, req *http.Request) error

func ErrorHandler(handler Error) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		err := handler(writer, req)
		if err != nil {
			io.WriteString(writer, err.Error())
		}
	}
}
